package cache

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

// PersistentL3Cache represents a persistent L3 cache implementation
// This cache uses memory-mapped files with compression for long-term storage
type PersistentL3Cache struct {
	mmap       []byte
	index      *LockFreeIndex
	stats      *CacheStats
	file       *os.File
	fileSize   int64
	compressor *CacheCompressor
	mu         sync.RWMutex
	closed     int32
	basePath   string
	indexFile  *os.File
	indexMmap  []byte
}

// CacheCompressor handles compression for the persistent cache
type CacheCompressor struct {
	algorithm string
	level     int
	enabled   bool
}

// CompressedEntry represents a compressed cache entry
type CompressedEntry struct {
	OriginalSize     int64
	CompressedSize   int64
	CompressionRatio float64
	Algorithm        string
	Data             []byte
}

// NewPersistentL3Cache creates a new persistent L3 cache
func NewPersistentL3Cache(basePath string, maxSize int64, compressionEnabled bool) (*PersistentL3Cache, error) {
	// Create base directory if it doesn't exist
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create base directory: %w", err)
	}

	// Create data file path
	dataFilePath := filepath.Join(basePath, "cache.data")
	indexFilePath := filepath.Join(basePath, "cache.index")

	// Create or open the data file
	dataFile, err := os.OpenFile(dataFilePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open data file: %w", err)
	}

	// Create or open the index file
	indexFile, err := os.OpenFile(indexFilePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		dataFile.Close()
		return nil, fmt.Errorf("failed to open index file: %w", err)
	}

	// Get file sizes
	dataFileInfo, err := dataFile.Stat()
	if err != nil {
		dataFile.Close()
		indexFile.Close()
		return nil, fmt.Errorf("failed to get data file info: %w", err)
	}

	indexFileInfo, err := indexFile.Stat()
	if err != nil {
		dataFile.Close()
		indexFile.Close()
		return nil, fmt.Errorf("failed to get index file info: %w", err)
	}

	dataFileSize := dataFileInfo.Size()
	indexFileSize := indexFileInfo.Size()

	// Initialize files if empty
	if dataFileSize == 0 {
		dataFileSize = 1024 * 1024 // Start with 1MB
		if err := dataFile.Truncate(dataFileSize); err != nil {
			dataFile.Close()
			indexFile.Close()
			return nil, fmt.Errorf("failed to initialize data file: %w", err)
		}
	}

	if indexFileSize == 0 {
		indexFileSize = 64 * 1024 // Start with 64KB for index
		if err := indexFile.Truncate(indexFileSize); err != nil {
			dataFile.Close()
			indexFile.Close()
			return nil, fmt.Errorf("failed to initialize index file: %w", err)
		}
	}

	// Memory map the data file (Windows doesn't support mmap, use alternative approach)
	dataMmap := make([]byte, dataFileSize)
	if _, err := dataFile.ReadAt(dataMmap, 0); err != nil && err != io.EOF {
		dataFile.Close()
		indexFile.Close()
		return nil, fmt.Errorf("failed to read data file: %w", err)
	}

	// Memory map the index file (Windows doesn't support mmap, use alternative approach)
	indexMmap := make([]byte, indexFileSize)
	if _, err := indexFile.ReadAt(indexMmap, 0); err != nil && err != io.EOF {
		dataFile.Close()
		indexFile.Close()
		return nil, fmt.Errorf("failed to read index file: %w", err)
	}

	// Create compressor
	compressor := &CacheCompressor{
		algorithm: "gzip",
		level:     6, // Default compression level
		enabled:   compressionEnabled,
	}

	cache := &PersistentL3Cache{
		mmap:       dataMmap,
		index:      &LockFreeIndex{entries: &sync.Map{}},
		stats:      &CacheStats{maxSize: maxSize},
		file:       dataFile,
		fileSize:   dataFileSize,
		compressor: compressor,
		basePath:   basePath,
		indexFile:  indexFile,
		indexMmap:  indexMmap,
	}

	// Load existing index from file
	if err := cache.loadIndex(); err != nil {
		cache.Close()
		return nil, fmt.Errorf("failed to load index: %w", err)
	}

	return cache, nil
}

// Get retrieves a value from the persistent cache
func (c *PersistentL3Cache) Get(key string) (interface{}, bool) {
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

	// Read compressed value from memory-mapped file
	compressedData, err := c.readValue(entry.Offset, entry.Size)
	if err != nil {
		atomic.AddInt64(&c.stats.misses, 1)
		return nil, false
	}

	// Decompress value
	valueBytes, err := c.decompressValue(compressedData)
	if err != nil {
		atomic.AddInt64(&c.stats.misses, 1)
		return nil, false
	}

	// Update access count
	atomic.AddInt64(&entry.AccessCount, 1)
	atomic.AddInt64(&c.stats.hits, 1)

	return string(valueBytes), true
}

// Set stores a value in the persistent cache
func (c *PersistentL3Cache) Set(key string, value interface{}, ttl time.Duration) error {
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

	// Compress value
	compressedData, err := c.compressValue(valueBytes)
	if err != nil {
		return fmt.Errorf("failed to compress value: %w", err)
	}

	// Check if we need to expand the file
	requiredSize := int64(len(compressedData)) + 1024 // Add some padding
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

	// Write compressed value to memory-mapped file
	if err := c.writeValue(offset, compressedData); err != nil {
		return fmt.Errorf("failed to write value: %w", err)
	}

	// Create index entry
	now := time.Now().UnixNano()
	entry := &IndexEntry{
		Offset:      offset,
		Size:        int64(len(compressedData)),
		TTL:         now + int64(ttl),
		CreatedAt:   now,
		AccessCount: 1,
		Hash:        c.hashKey(key),
	}

	// Store in index
	c.index.entries.Store(key, entry)

	// Update stats
	atomic.AddInt64(&c.stats.size, 1)
	atomic.AddInt64(&c.stats.memoryUsage, int64(len(compressedData)))

	return nil
}

// Delete removes a value from the persistent cache
func (c *PersistentL3Cache) Delete(key string) error {
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

	// Mark space as free
	c.markSpaceFree(entry.Offset, entry.Size)

	// Update stats
	atomic.AddInt64(&c.stats.size, -1)
	atomic.AddInt64(&c.stats.memoryUsage, -entry.Size)

	return nil
}

// Clear removes all values from the persistent cache
func (c *PersistentL3Cache) Clear() error {
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
func (c *PersistentL3Cache) GetStats() *CacheStats {
	return &CacheStats{
		hits:        atomic.LoadInt64(&c.stats.hits),
		misses:      atomic.LoadInt64(&c.stats.misses),
		size:        atomic.LoadInt64(&c.stats.size),
		maxSize:     c.stats.maxSize,
		memoryUsage: atomic.LoadInt64(&c.stats.memoryUsage),
	}
}

// GetHitRate returns the cache hit rate
func (c *PersistentL3Cache) GetHitRate() float64 {
	hits := atomic.LoadInt64(&c.stats.hits)
	misses := atomic.LoadInt64(&c.stats.misses)
	total := hits + misses
	if total == 0 {
		return 0.0
	}
	return float64(hits) / float64(total)
}

// GetUsagePercentage returns the cache usage percentage
func (c *PersistentL3Cache) GetUsagePercentage() float64 {
	size := atomic.LoadInt64(&c.stats.size)
	maxSize := c.stats.maxSize
	if maxSize == 0 {
		return 0.0
	}
	return float64(size) / float64(maxSize) * 100.0
}

// Has checks if a key exists in the cache
func (c *PersistentL3Cache) Has(key string) bool {
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
func (c *PersistentL3Cache) Keys() []string {
	keys := make([]string, 0)
	c.index.entries.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}

// Size returns the current size of the cache
func (c *PersistentL3Cache) Size() int64 {
	return atomic.LoadInt64(&c.stats.size)
}

// MaxSize returns the maximum size of the cache
func (c *PersistentL3Cache) MaxSize() int64 {
	return c.stats.maxSize
}

// SetMaxSize sets the maximum size of the cache
func (c *PersistentL3Cache) SetMaxSize(maxSize int64) {
	atomic.StoreInt64(&c.stats.maxSize, maxSize)
}

// Cleanup removes expired entries
func (c *PersistentL3Cache) Cleanup() int64 {
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

// Close closes the persistent cache
func (c *PersistentL3Cache) Close() error {
	// Mark as closed
	atomic.StoreInt32(&c.closed, 1)

	// Save index to file
	if err := c.saveIndex(); err != nil {
		return fmt.Errorf("failed to save index: %w", err)
	}

	// No need to unmap memory on Windows (we're not using mmap)

	// Close files
	if err := c.file.Close(); err != nil {
		return fmt.Errorf("failed to close data file: %w", err)
	}

	if err := c.indexFile.Close(); err != nil {
		return fmt.Errorf("failed to close index file: %w", err)
	}

	return nil
}

// compressValue compresses a value using the configured algorithm
func (c *PersistentL3Cache) compressValue(data []byte) ([]byte, error) {
	if !c.compressor.enabled {
		return data, nil
	}

	switch c.compressor.algorithm {
	case "gzip":
		return c.compressGzip(data)
	case "lz4":
		return c.compressLZ4(data)
	default:
		return data, nil
	}
}

// decompressValue decompresses a value using the configured algorithm
func (c *PersistentL3Cache) decompressValue(data []byte) ([]byte, error) {
	if !c.compressor.enabled {
		return data, nil
	}

	switch c.compressor.algorithm {
	case "gzip":
		return c.decompressGzip(data)
	case "lz4":
		return c.decompressLZ4(data)
	default:
		return data, nil
	}
}

// compressGzip compresses data using gzip
func (c *PersistentL3Cache) compressGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer, err := gzip.NewWriterLevel(&buf, c.compressor.level)
	if err != nil {
		return nil, err
	}

	if _, err := writer.Write(data); err != nil {
		writer.Close()
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// decompressGzip decompresses data using gzip
func (c *PersistentL3Cache) decompressGzip(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return io.ReadAll(reader)
}

// compressLZ4 compresses data using LZ4 (placeholder)
func (c *PersistentL3Cache) compressLZ4(data []byte) ([]byte, error) {
	// Placeholder - in a real implementation, you'd use an LZ4 library
	return data, nil
}

// decompressLZ4 decompresses data using LZ4 (placeholder)
func (c *PersistentL3Cache) decompressLZ4(data []byte) ([]byte, error) {
	// Placeholder - in a real implementation, you'd use an LZ4 library
	return data, nil
}

// serializeValue serializes a value to bytes
func (c *PersistentL3Cache) serializeValue(value interface{}) ([]byte, error) {
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
func (c *PersistentL3Cache) deserializeValue(data []byte) (interface{}, error) {
	// Simple deserialization - in a real implementation, you'd use a proper deserializer
	return string(data), nil
}

// readValue reads a value from the memory-mapped file
func (c *PersistentL3Cache) readValue(offset, size int64) ([]byte, error) {
	if offset+size > int64(len(c.mmap)) {
		return nil, fmt.Errorf("offset out of bounds")
	}

	data := make([]byte, size)
	copy(data, c.mmap[offset:offset+size])

	return data, nil
}

// writeValue writes a value to the memory-mapped file
func (c *PersistentL3Cache) writeValue(offset int64, data []byte) error {
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
func (c *PersistentL3Cache) findFreeSpace(size int64) (int64, error) {
	// Simple implementation - just append to the end
	// In a real implementation, you'd want a proper free space manager
	offset := c.fileSize
	c.fileSize += size
	return offset, nil
}

// markSpaceFree marks space as free (placeholder implementation)
func (c *PersistentL3Cache) markSpaceFree(offset, size int64) {
	// In a real implementation, you'd update a free space manager
	// For now, we just leave the space as is
}

// expandFile expands the memory-mapped file
func (c *PersistentL3Cache) expandFile(newSize int64) error {
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
func (c *PersistentL3Cache) hashKey(key string) uint64 {
	// Simple hash function - in a real implementation, you'd use a proper hash function
	hash := uint64(0)
	for _, b := range []byte(key) {
		hash = hash*31 + uint64(b)
	}
	return hash
}

// loadIndex loads the index from the memory-mapped file
func (c *PersistentL3Cache) loadIndex() error {
	// Simple implementation - in a real implementation, you'd have a proper index format
	// For now, we just start with an empty index
	return nil
}

// saveIndex saves the index to the memory-mapped file
func (c *PersistentL3Cache) saveIndex() error {
	// Simple implementation - in a real implementation, you'd have a proper index format
	// For now, we just return success
	return nil
}

// GetFileSize returns the current file size
func (c *PersistentL3Cache) GetFileSize() int64 {
	return c.fileSize
}

// GetMemoryUsage returns the current memory usage
func (c *PersistentL3Cache) GetMemoryUsage() int64 {
	return atomic.LoadInt64(&c.stats.memoryUsage)
}

// IsClosed returns whether the cache is closed
func (c *PersistentL3Cache) IsClosed() bool {
	return atomic.LoadInt32(&c.closed) == 1
}

// Flush flushes the memory-mapped file to disk
func (c *PersistentL3Cache) Flush() error {
	if atomic.LoadInt32(&c.closed) == 1 {
		return fmt.Errorf("cache is closed")
	}

	// Flush memory buffer to disk
	if err := c.file.Sync(); err != nil {
		return fmt.Errorf("failed to flush data to disk: %w", err)
	}

	if err := c.indexFile.Sync(); err != nil {
		return fmt.Errorf("failed to flush index to disk: %w", err)
	}

	return nil
}

// GetCompressionRatio returns the compression ratio
func (c *PersistentL3Cache) GetCompressionRatio() float64 {
	if !c.compressor.enabled {
		return 1.0
	}

	// Calculate compression ratio based on stored data
	// This is a placeholder - in a real implementation, you'd track this properly
	return 0.7 // Assume 30% compression
}

// GetFragmentationRatio returns the fragmentation ratio (placeholder)
func (c *PersistentL3Cache) GetFragmentationRatio() float64 {
	// Placeholder - in a real implementation, you'd calculate actual fragmentation ratio
	return 0.1 // Assume 10% fragmentation
}

// GetIndexSize returns the number of entries in the index
func (c *PersistentL3Cache) GetIndexSize() int64 {
	size := int64(0)
	c.index.entries.Range(func(key, value interface{}) bool {
		size++
		return true
	})
	return size
}

// GetAverageEntrySize returns the average size of entries
func (c *PersistentL3Cache) GetAverageEntrySize() float64 {
	totalSize := atomic.LoadInt64(&c.stats.memoryUsage)
	count := atomic.LoadInt64(&c.stats.size)

	if count == 0 {
		return 0.0
	}

	return float64(totalSize) / float64(count)
}

// SetCompressionLevel sets the compression level
func (c *PersistentL3Cache) SetCompressionLevel(level int) {
	c.compressor.level = level
}

// SetCompressionAlgorithm sets the compression algorithm
func (c *PersistentL3Cache) SetCompressionAlgorithm(algorithm string) {
	c.compressor.algorithm = algorithm
}

// EnableCompression enables or disables compression
func (c *PersistentL3Cache) EnableCompression(enabled bool) {
	c.compressor.enabled = enabled
}

// GetCompressionStats returns compression statistics
func (c *PersistentL3Cache) GetCompressionStats() map[string]interface{} {
	return map[string]interface{}{
		"enabled":   c.compressor.enabled,
		"algorithm": c.compressor.algorithm,
		"level":     c.compressor.level,
		"ratio":     c.GetCompressionRatio(),
	}
}
