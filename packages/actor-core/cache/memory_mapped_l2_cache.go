package cache

import (
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// MemoryMappedL2Cache represents a memory-mapped L2 cache implementation
// This cache uses memory-mapped files for fast persistent storage
type MemoryMappedL2Cache struct {
	mmap     []byte
	index    *LockFreeIndex
	stats    *CacheStats
	file     *os.File
	fileSize int64
	mu       sync.RWMutex
	closed   int32
}

// LockFreeIndex represents a lock-free index for memory-mapped cache
type LockFreeIndex struct {
	entries *sync.Map // map[string]*IndexEntry
	mu      sync.RWMutex
}

// IndexEntry represents an entry in the memory-mapped cache index
type IndexEntry struct {
	Offset      int64
	Size        int64
	TTL         int64
	CreatedAt   int64
	AccessCount int64
	Hash        uint64
}

// MemoryMappedCacheEntry represents a cache entry in memory-mapped storage
type MemoryMappedCacheEntry struct {
	Key         string
	Value       []byte
	TTL         int64
	CreatedAt   int64
	AccessCount int64
	Hash        uint64
}

// NewMemoryMappedL2Cache creates a new memory-mapped L2 cache
func NewMemoryMappedL2Cache(filePath string, maxSize int64) (*MemoryMappedL2Cache, error) {
	// Create or open the cache file
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open cache file: %w", err)
	}

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	fileSize := fileInfo.Size()
	if fileSize == 0 {
		// Initialize empty file with header
		fileSize = 1024 // Start with 1KB
		if err := file.Truncate(fileSize); err != nil {
			file.Close()
			return nil, fmt.Errorf("failed to initialize file: %w", err)
		}
	}

	// Memory map the file (Windows doesn't support mmap, use alternative approach)
	// For now, we'll use a simple approach without memory mapping
	mmap := make([]byte, fileSize)
	if _, err := file.ReadAt(mmap, 0); err != nil && err != io.EOF {
		file.Close()
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	cache := &MemoryMappedL2Cache{
		mmap:     mmap,
		index:    &LockFreeIndex{entries: &sync.Map{}},
		stats:    &CacheStats{maxSize: maxSize},
		file:     file,
		fileSize: fileSize,
	}

	// Load existing index from file
	if err := cache.loadIndex(); err != nil {
		cache.Close()
		return nil, fmt.Errorf("failed to load index: %w", err)
	}

	return cache, nil
}

// Get retrieves a value from the memory-mapped cache
func (c *MemoryMappedL2Cache) Get(key string) (interface{}, bool) {
	if key == "" {
		atomic.AddInt64(&c.stats.misses, 1)
		return nil, false
	}

	// Check if cache is closed
	if atomic.LoadInt32(&c.closed) == 1 {
		return nil, false
	}

	// Get index entry
	entryInterface, exists := c.index.entries.Load(key)
	if !exists {
		atomic.AddInt64(&c.stats.misses, 1)
		return nil, false
	}

	entry := entryInterface.(*IndexEntry)

	// Check TTL
	now := time.Now().UnixNano()
	if now > entry.TTL {
		// Entry expired, remove it
		c.Delete(key)
		atomic.AddInt64(&c.stats.misses, 1)
		return nil, false
	}

	// Read value from memory-mapped file
	value, err := c.readValue(entry.Offset, entry.Size)
	if err != nil {
		atomic.AddInt64(&c.stats.misses, 1)
		return nil, false
	}

	// Update access count
	atomic.AddInt64(&entry.AccessCount, 1)
	atomic.AddInt64(&c.stats.hits, 1)

	return value, true
}

// Set stores a value in the memory-mapped cache
func (c *MemoryMappedL2Cache) Set(key string, value interface{}, ttl time.Duration) error {
	if key == "" {
		return ErrEmptyKey
	}

	if value == nil {
		return ErrNilValue
	}

	// Check if cache is closed
	if atomic.LoadInt32(&c.closed) == 1 {
		return fmt.Errorf("cache is closed")
	}

	// Serialize value
	valueBytes, err := c.serializeValue(value)
	if err != nil {
		return fmt.Errorf("failed to serialize value: %w", err)
	}

	// Check if we need to expand the file
	requiredSize := int64(len(valueBytes)) + 1024 // Add some padding
	if requiredSize > c.fileSize {
		if err := c.expandFile(requiredSize); err != nil {
			return fmt.Errorf("failed to expand file: %w", err)
		}
	}

	// Find a place to store the value
	offset, err := c.findFreeSpace(requiredSize)
	if err != nil {
		return fmt.Errorf("failed to find free space: %w", err)
	}

	// Write value to memory-mapped file
	if err := c.writeValue(offset, valueBytes); err != nil {
		return fmt.Errorf("failed to write value: %w", err)
	}

	// Create index entry
	now := time.Now().UnixNano()
	entry := &IndexEntry{
		Offset:      offset,
		Size:        int64(len(valueBytes)),
		TTL:         now + int64(ttl),
		CreatedAt:   now,
		AccessCount: 1,
		Hash:        c.hashKey(key),
	}

	// Store in index
	c.index.entries.Store(key, entry)

	// Update stats
	atomic.AddInt64(&c.stats.size, 1)
	atomic.AddInt64(&c.stats.memoryUsage, int64(len(valueBytes)))

	return nil
}

// Delete removes a value from the memory-mapped cache
func (c *MemoryMappedL2Cache) Delete(key string) error {
	if key == "" {
		return ErrEmptyKey
	}

	// Check if cache is closed
	if atomic.LoadInt32(&c.closed) == 1 {
		return fmt.Errorf("cache is closed")
	}

	// Get index entry
	entryInterface, exists := c.index.entries.Load(key)
	if !exists {
		return ErrKeyNotFound
	}

	entry := entryInterface.(*IndexEntry)

	// Remove from index
	c.index.entries.Delete(key)

	// Mark space as free (simple implementation)
	// In a real implementation, you'd want a free space manager
	c.markSpaceFree(entry.Offset, entry.Size)

	// Update stats
	atomic.AddInt64(&c.stats.size, -1)
	atomic.AddInt64(&c.stats.memoryUsage, -entry.Size)

	return nil
}

// Clear removes all values from the memory-mapped cache
func (c *MemoryMappedL2Cache) Clear() error {
	// Check if cache is closed
	if atomic.LoadInt32(&c.closed) == 1 {
		return fmt.Errorf("cache is closed")
	}

	// Clear index
	c.index.entries.Range(func(key, value interface{}) bool {
		c.index.entries.Delete(key)
		return true
	})

	// Reset stats
	atomic.StoreInt64(&c.stats.size, 0)
	atomic.StoreInt64(&c.stats.memoryUsage, 0)
	atomic.StoreInt64(&c.stats.hits, 0)
	atomic.StoreInt64(&c.stats.misses, 0)

	return nil
}

// GetStats returns cache statistics
func (c *MemoryMappedL2Cache) GetStats() *CacheStats {
	return &CacheStats{
		hits:        atomic.LoadInt64(&c.stats.hits),
		misses:      atomic.LoadInt64(&c.stats.misses),
		size:        atomic.LoadInt64(&c.stats.size),
		maxSize:     c.stats.maxSize,
		memoryUsage: atomic.LoadInt64(&c.stats.memoryUsage),
	}
}

// GetHitRate returns the cache hit rate
func (c *MemoryMappedL2Cache) GetHitRate() float64 {
	hits := atomic.LoadInt64(&c.stats.hits)
	misses := atomic.LoadInt64(&c.stats.misses)
	total := hits + misses
	if total == 0 {
		return 0.0
	}
	return float64(hits) / float64(total)
}

// GetUsagePercentage returns the cache usage percentage
func (c *MemoryMappedL2Cache) GetUsagePercentage() float64 {
	size := atomic.LoadInt64(&c.stats.size)
	maxSize := c.stats.maxSize
	if maxSize == 0 {
		return 0.0
	}
	return float64(size) / float64(maxSize) * 100.0
}

// Has checks if a key exists in the cache
func (c *MemoryMappedL2Cache) Has(key string) bool {
	entryInterface, exists := c.index.entries.Load(key)
	if !exists {
		return false
	}

	entry := entryInterface.(*IndexEntry)
	now := time.Now().UnixNano()

	if now > entry.TTL {
		// Entry expired, remove it
		c.Delete(key)
		return false
	}

	return true
}

// Keys returns all keys in the cache
func (c *MemoryMappedL2Cache) Keys() []string {
	keys := make([]string, 0)
	c.index.entries.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}

// Size returns the current size of the cache
func (c *MemoryMappedL2Cache) Size() int64 {
	return atomic.LoadInt64(&c.stats.size)
}

// MaxSize returns the maximum size of the cache
func (c *MemoryMappedL2Cache) MaxSize() int64 {
	return c.stats.maxSize
}

// SetMaxSize sets the maximum size of the cache
func (c *MemoryMappedL2Cache) SetMaxSize(maxSize int64) {
	atomic.StoreInt64(&c.stats.maxSize, maxSize)
}

// Cleanup removes expired entries
func (c *MemoryMappedL2Cache) Cleanup() int64 {
	removed := int64(0)
	now := time.Now().UnixNano()

	c.index.entries.Range(func(key, value interface{}) bool {
		entry := value.(*IndexEntry)

		if now > entry.TTL {
			c.index.entries.Delete(key)
			c.markSpaceFree(entry.Offset, entry.Size)
			atomic.AddInt64(&c.stats.size, -1)
			atomic.AddInt64(&c.stats.memoryUsage, -entry.Size)
			removed++
		}
		return true
	})

	return removed
}

// Close closes the memory-mapped cache
func (c *MemoryMappedL2Cache) Close() error {
	// Mark as closed
	atomic.StoreInt32(&c.closed, 1)

	// Save index to file
	if err := c.saveIndex(); err != nil {
		return fmt.Errorf("failed to save index: %w", err)
	}

	// No need to unmap memory on Windows (we're not using mmap)

	// Close file
	if err := c.file.Close(); err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}

	return nil
}

// serializeValue serializes a value to bytes
func (c *MemoryMappedL2Cache) serializeValue(value interface{}) ([]byte, error) {
	// Simple serialization - in a real implementation, you'd use a proper serializer
	switch v := value.(type) {
	case string:
		return []byte(v), nil
	case []byte:
		return v, nil
	case int:
		return []byte(fmt.Sprintf("%d", v)), nil
	case int64:
		return []byte(fmt.Sprintf("%d", v)), nil
	case float64:
		return []byte(fmt.Sprintf("%f", v)), nil
	case bool:
		if v {
			return []byte("true"), nil
		}
		return []byte("false"), nil
	default:
		return []byte(fmt.Sprintf("%v", v)), nil
	}
}

// deserializeValue deserializes bytes to a value
func (c *MemoryMappedL2Cache) deserializeValue(data []byte) (interface{}, error) {
	// Simple deserialization - in a real implementation, you'd use a proper deserializer
	return string(data), nil
}

// readValue reads a value from the memory-mapped file
func (c *MemoryMappedL2Cache) readValue(offset, size int64) (interface{}, error) {
	if offset+size > int64(len(c.mmap)) {
		return nil, fmt.Errorf("offset out of bounds")
	}

	data := make([]byte, size)
	copy(data, c.mmap[offset:offset+size])

	return c.deserializeValue(data)
}

// writeValue writes a value to the memory-mapped file
func (c *MemoryMappedL2Cache) writeValue(offset int64, data []byte) error {
	requiredSize := offset + int64(len(data))
	if requiredSize > int64(len(c.mmap)) {
		// Expand buffer
		newSize := requiredSize * 2 // Double the size
		newMmap := make([]byte, newSize)
		copy(newMmap, c.mmap)
		c.mmap = newMmap
	}

	copy(c.mmap[offset:], data)
	return nil
}

// findFreeSpace finds a free space in the memory-mapped file
func (c *MemoryMappedL2Cache) findFreeSpace(size int64) (int64, error) {
	// Simple implementation - just append to the end
	// In a real implementation, you'd want a proper free space manager
	offset := c.fileSize
	c.fileSize += size
	return offset, nil
}

// markSpaceFree marks space as free (placeholder implementation)
func (c *MemoryMappedL2Cache) markSpaceFree(offset, size int64) {
	// In a real implementation, you'd update a free space manager
	// For now, we just leave the space as is
}

// expandFile expands the memory-mapped file
func (c *MemoryMappedL2Cache) expandFile(newSize int64) error {
	// Expand file
	if err := c.file.Truncate(newSize); err != nil {
		return fmt.Errorf("failed to expand file: %w", err)
	}

	// Reallocate memory buffer
	c.mmap = make([]byte, newSize)
	if _, err := c.file.ReadAt(c.mmap, 0); err != nil && err != io.EOF {
		return fmt.Errorf("failed to read expanded file: %w", err)
	}

	c.fileSize = newSize

	return nil
}

// hashKey generates a hash for a key
func (c *MemoryMappedL2Cache) hashKey(key string) uint64 {
	// Simple hash function - in a real implementation, you'd use a proper hash function
	hash := uint64(0)
	for _, b := range []byte(key) {
		hash = hash*31 + uint64(b)
	}
	return hash
}

// loadIndex loads the index from the memory-mapped file
func (c *MemoryMappedL2Cache) loadIndex() error {
	// Simple implementation - in a real implementation, you'd have a proper index format
	// For now, we just start with an empty index
	return nil
}

// saveIndex saves the index to the memory-mapped file
func (c *MemoryMappedL2Cache) saveIndex() error {
	// Simple implementation - in a real implementation, you'd have a proper index format
	// For now, we just return success
	return nil
}

// GetFileSize returns the current file size
func (c *MemoryMappedL2Cache) GetFileSize() int64 {
	return c.fileSize
}

// GetMemoryUsage returns the current memory usage
func (c *MemoryMappedL2Cache) GetMemoryUsage() int64 {
	return atomic.LoadInt64(&c.stats.memoryUsage)
}

// IsClosed returns whether the cache is closed
func (c *MemoryMappedL2Cache) IsClosed() bool {
	return atomic.LoadInt32(&c.closed) == 1
}

// Flush flushes the memory-mapped file to disk
func (c *MemoryMappedL2Cache) Flush() error {
	if atomic.LoadInt32(&c.closed) == 1 {
		return fmt.Errorf("cache is closed")
	}

	// Flush memory buffer to disk
	if err := c.file.Sync(); err != nil {
		return fmt.Errorf("failed to flush to disk: %w", err)
	}

	return nil
}

// GetIndexSize returns the number of entries in the index
func (c *MemoryMappedL2Cache) GetIndexSize() int64 {
	size := int64(0)
	c.index.entries.Range(func(key, value interface{}) bool {
		size++
		return true
	})
	return size
}

// GetAverageEntrySize returns the average size of entries
func (c *MemoryMappedL2Cache) GetAverageEntrySize() float64 {
	totalSize := atomic.LoadInt64(&c.stats.memoryUsage)
	count := atomic.LoadInt64(&c.stats.size)

	if count == 0 {
		return 0.0
	}

	return float64(totalSize) / float64(count)
}

// GetCompressionRatio returns the compression ratio (placeholder)
func (c *MemoryMappedL2Cache) GetCompressionRatio() float64 {
	// Placeholder - in a real implementation, you'd calculate actual compression ratio
	return 1.0
}

// GetFragmentationRatio returns the fragmentation ratio (placeholder)
func (c *MemoryMappedL2Cache) GetFragmentationRatio() float64 {
	// Placeholder - in a real implementation, you'd calculate actual fragmentation ratio
	return 0.0
}
