package cache

import (
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// LockFreeL1Cache represents a lock-free L1 cache implementation
// This is the fastest cache layer, designed for ultra-low latency
type LockFreeL1Cache struct {
	cache     *sync.Map
	maxSize   int64
	stats     *CacheStats
	evictor   *LockFreeEvictor
	preloader *CachePreloader
}

// CacheEntry represents a cache entry with atomic operations
type CacheEntry struct {
	value       unsafe.Pointer // atomic pointer to value
	expiresAt   int64          // atomic timestamp
	createdAt   int64          // atomic timestamp
	accessCount int64          // atomic access count
	size        int64          // atomic size in bytes
}

// CacheStats represents cache statistics with atomic operations
type CacheStats struct {
	hits        int64
	misses      int64
	size        int64
	maxSize     int64
	memoryUsage int64
}

// LockFreeEvictor handles lock-free eviction
type LockFreeEvictor struct {
	accessCounts   *sync.Map // map[string]*int64
	evictionPolicy string
}

// CachePreloader handles aggressive preloading
type CachePreloader struct {
	preloadQueue   chan string
	preloadWorkers int
	stats          *PreloadStats
	actorCore      interface{} // Will be replaced with actual ActorCore type
}

// PreloadStats represents preload statistics
type PreloadStats struct {
	preloadedCount int64
	preloadTime    int64 // nanoseconds
	hitRate        int64 // percentage * 100
}

// NewLockFreeL1Cache creates a new lock-free L1 cache
func NewLockFreeL1Cache(maxSize int64, evictionPolicy string) *LockFreeL1Cache {
	return &LockFreeL1Cache{
		cache:     &sync.Map{},
		maxSize:   maxSize,
		stats:     &CacheStats{maxSize: maxSize},
		evictor:   &LockFreeEvictor{accessCounts: &sync.Map{}, evictionPolicy: evictionPolicy},
		preloader: &CachePreloader{preloadQueue: make(chan string, 1000), preloadWorkers: 4},
	}
}

// Get retrieves a value from the cache (lock-free)
func (c *LockFreeL1Cache) Get(key string) (interface{}, bool) {
	if key == "" {
		atomic.AddInt64(&c.stats.misses, 1)
		return nil, false
	}

	// Load entry atomically
	entryInterface, exists := c.cache.Load(key)
	if !exists {
		atomic.AddInt64(&c.stats.misses, 1)
		return nil, false
	}

	entry := entryInterface.(*CacheEntry)

	// Check expiration atomically
	now := time.Now().UnixNano()
	expiresAt := atomic.LoadInt64(&entry.expiresAt)
	if now > expiresAt {
		// Entry expired, remove it
		c.cache.Delete(key)
		atomic.AddInt64(&c.stats.misses, 1)
		atomic.AddInt64(&c.stats.size, -1)
		return nil, false
	}

	// Update access count atomically
	atomic.AddInt64(&entry.accessCount, 1)
	atomic.AddInt64(&c.stats.hits, 1)

	// Load value atomically
	valuePtr := atomic.LoadPointer(&entry.value)
	if valuePtr == nil {
		atomic.AddInt64(&c.stats.misses, 1)
		return nil, false
	}

	return *(*interface{})(valuePtr), true
}

// Set stores a value in the cache (lock-free)
func (c *LockFreeL1Cache) Set(key string, value interface{}, ttl time.Duration) error {
	if key == "" {
		return ErrEmptyKey
	}

	if value == nil {
		return ErrNilValue
	}

	// Check if we need to evict
	currentSize := atomic.LoadInt64(&c.stats.size)
	if currentSize >= c.maxSize {
		if err := c.evict(); err != nil {
			return err
		}
	}

	// Create new entry
	now := time.Now().UnixNano()
	expiresAt := now + int64(ttl)

	// Calculate size (approximate)
	size := int64(unsafe.Sizeof(value)) + int64(len(key))

	entry := &CacheEntry{
		expiresAt:   expiresAt,
		createdAt:   now,
		accessCount: 1,
		size:        size,
	}

	// Store value atomically
	valuePtr := unsafe.Pointer(&value)
	atomic.StorePointer(&entry.value, valuePtr)

	// Store entry in cache
	c.cache.Store(key, entry)

	// Update stats atomically
	atomic.AddInt64(&c.stats.size, 1)
	atomic.AddInt64(&c.stats.memoryUsage, size)

	return nil
}

// Delete removes a value from the cache (lock-free)
func (c *LockFreeL1Cache) Delete(key string) error {
	if key == "" {
		return ErrEmptyKey
	}

	// Load entry to get size
	entryInterface, exists := c.cache.Load(key)
	if !exists {
		return ErrKeyNotFound
	}

	entry := entryInterface.(*CacheEntry)
	size := atomic.LoadInt64(&entry.size)

	// Delete from cache
	c.cache.Delete(key)

	// Update stats atomically
	atomic.AddInt64(&c.stats.size, -1)
	atomic.AddInt64(&c.stats.memoryUsage, -size)

	return nil
}

// Clear removes all values from the cache (lock-free)
func (c *LockFreeL1Cache) Clear() error {
	// Clear all entries
	c.cache.Range(func(key, value interface{}) bool {
		c.cache.Delete(key)
		return true
	})

	// Reset stats atomically
	atomic.StoreInt64(&c.stats.size, 0)
	atomic.StoreInt64(&c.stats.memoryUsage, 0)
	atomic.StoreInt64(&c.stats.hits, 0)
	atomic.StoreInt64(&c.stats.misses, 0)

	return nil
}

// GetStats returns cache statistics
func (c *LockFreeL1Cache) GetStats() *CacheStats {
	return &CacheStats{
		hits:        atomic.LoadInt64(&c.stats.hits),
		misses:      atomic.LoadInt64(&c.stats.misses),
		size:        atomic.LoadInt64(&c.stats.size),
		maxSize:     c.stats.maxSize,
		memoryUsage: atomic.LoadInt64(&c.stats.memoryUsage),
	}
}

// GetHitRate returns the cache hit rate
func (c *LockFreeL1Cache) GetHitRate() float64 {
	hits := atomic.LoadInt64(&c.stats.hits)
	misses := atomic.LoadInt64(&c.stats.misses)
	total := hits + misses
	if total == 0 {
		return 0.0
	}
	return float64(hits) / float64(total)
}

// GetUsagePercentage returns the cache usage percentage
func (c *LockFreeL1Cache) GetUsagePercentage() float64 {
	size := atomic.LoadInt64(&c.stats.size)
	maxSize := c.stats.maxSize
	if maxSize == 0 {
		return 0.0
	}
	return float64(size) / float64(maxSize) * 100.0
}

// evict removes entries based on the eviction policy (lock-free)
func (c *LockFreeL1Cache) evict() error {
	switch c.evictor.evictionPolicy {
	case "allkeys-lru":
		return c.evictLRU()
	case "allkeys-lfu":
		return c.evictLFU()
	case "volatile-ttl":
		return c.evictTTL()
	default:
		return c.evictLRU() // Default to LRU
	}
}

// evictLRU evicts the least recently used entry (lock-free)
func (c *LockFreeL1Cache) evictLRU() error {
	var oldestKey string
	var oldestTime int64 = 0

	c.cache.Range(func(key, value interface{}) bool {
		entry := value.(*CacheEntry)
		createdAt := atomic.LoadInt64(&entry.createdAt)

		if oldestKey == "" || createdAt < oldestTime {
			oldestKey = key.(string)
			oldestTime = createdAt
		}
		return true
	})

	if oldestKey != "" {
		return c.Delete(oldestKey)
	}

	return nil
}

// evictLFU evicts the least frequently used entry (lock-free)
func (c *LockFreeL1Cache) evictLFU() error {
	var leastUsedKey string
	var leastUsedCount int64 = 0

	c.cache.Range(func(key, value interface{}) bool {
		entry := value.(*CacheEntry)
		accessCount := atomic.LoadInt64(&entry.accessCount)

		if leastUsedKey == "" || accessCount < leastUsedCount {
			leastUsedKey = key.(string)
			leastUsedCount = accessCount
		}
		return true
	})

	if leastUsedKey != "" {
		return c.Delete(leastUsedKey)
	}

	return nil
}

// evictTTL evicts the entry with the shortest TTL (lock-free)
func (c *LockFreeL1Cache) evictTTL() error {
	var shortestKey string
	var shortestTTL int64 = 0

	now := time.Now().UnixNano()

	c.cache.Range(func(key, value interface{}) bool {
		entry := value.(*CacheEntry)
		expiresAt := atomic.LoadInt64(&entry.expiresAt)
		ttl := expiresAt - now

		if shortestKey == "" || ttl < shortestTTL {
			shortestKey = key.(string)
			shortestTTL = ttl
		}
		return true
	})

	if shortestKey != "" {
		return c.Delete(shortestKey)
	}

	return nil
}

// Has checks if a key exists in the cache (lock-free)
func (c *LockFreeL1Cache) Has(key string) bool {
	entryInterface, exists := c.cache.Load(key)
	if !exists {
		return false
	}

	entry := entryInterface.(*CacheEntry)
	now := time.Now().UnixNano()
	expiresAt := atomic.LoadInt64(&entry.expiresAt)

	if now > expiresAt {
		// Entry expired, remove it
		c.cache.Delete(key)
		atomic.AddInt64(&c.stats.size, -1)
		return false
	}

	return true
}

// Keys returns all keys in the cache
func (c *LockFreeL1Cache) Keys() []string {
	keys := make([]string, 0)
	c.cache.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}

// Size returns the current size of the cache
func (c *LockFreeL1Cache) Size() int64 {
	return atomic.LoadInt64(&c.stats.size)
}

// MaxSize returns the maximum size of the cache
func (c *LockFreeL1Cache) MaxSize() int64 {
	return c.maxSize
}

// SetMaxSize sets the maximum size of the cache
func (c *LockFreeL1Cache) SetMaxSize(maxSize int64) {
	c.maxSize = maxSize
	atomic.StoreInt64(&c.stats.maxSize, maxSize)
}

// GetEvictionPolicy returns the eviction policy
func (c *LockFreeL1Cache) GetEvictionPolicy() string {
	return c.evictor.evictionPolicy
}

// SetEvictionPolicy sets the eviction policy
func (c *LockFreeL1Cache) SetEvictionPolicy(policy string) {
	c.evictor.evictionPolicy = policy
}

// Cleanup removes expired entries (lock-free)
func (c *LockFreeL1Cache) Cleanup() int64 {
	removed := int64(0)
	now := time.Now().UnixNano()

	c.cache.Range(func(key, value interface{}) bool {
		entry := value.(*CacheEntry)
		expiresAt := atomic.LoadInt64(&entry.expiresAt)

		if now > expiresAt {
			c.cache.Delete(key)
			size := atomic.LoadInt64(&entry.size)
			atomic.AddInt64(&c.stats.size, -1)
			atomic.AddInt64(&c.stats.memoryUsage, -size)
			removed++
		}
		return true
	})

	return removed
}

// Reset resets the cache statistics
func (c *LockFreeL1Cache) Reset() {
	atomic.StoreInt64(&c.stats.hits, 0)
	atomic.StoreInt64(&c.stats.misses, 0)
}

// StartPreloading starts the preloading workers
func (c *LockFreeL1Cache) StartPreloading() {
	for i := 0; i < c.preloader.preloadWorkers; i++ {
		go c.preloadWorker()
	}
}

// preloadWorker processes preload requests
func (c *LockFreeL1Cache) preloadWorker() {
	for actorID := range c.preloader.preloadQueue {
		start := time.Now()

		// Preload actor data (placeholder implementation)
		// This will be implemented when we integrate with ActorCore
		_ = actorID

		// Update stats (check for nil pointer)
		if c.preloader != nil && c.preloader.stats != nil {
			atomic.AddInt64(&c.preloader.stats.preloadedCount, 1)
			atomic.AddInt64(&c.preloader.stats.preloadTime, int64(time.Since(start)))
		}
	}
}

// Preload adds an actor ID to the preload queue
func (c *LockFreeL1Cache) Preload(actorID string) {
	select {
	case c.preloader.preloadQueue <- actorID:
	default:
		// Queue is full, skip preload
	}
}

// GetPreloadStats returns preload statistics
func (c *LockFreeL1Cache) GetPreloadStats() *PreloadStats {
	if c.preloader == nil || c.preloader.stats == nil {
		return &PreloadStats{}
	}

	return &PreloadStats{
		preloadedCount: atomic.LoadInt64(&c.preloader.stats.preloadedCount),
		preloadTime:    atomic.LoadInt64(&c.preloader.stats.preloadTime),
		hitRate:        atomic.LoadInt64(&c.preloader.stats.hitRate),
	}
}

// Errors
var (
	ErrEmptyKey    = &CacheError{message: "key cannot be empty"}
	ErrNilValue    = &CacheError{message: "value cannot be nil"}
	ErrKeyNotFound = &CacheError{message: "key not found"}
)

// CacheError represents a cache error
type CacheError struct {
	message string
}

func (e *CacheError) Error() string {
	return e.message
}
