package registry

import (
	"chaos-actor-module/packages/actor-core/interfaces"
	"fmt"
	"sync"
	"time"
)

// CacheEntry represents a cache entry
type CacheEntry struct {
	Value     interface{}
	ExpiresAt time.Time
	CreatedAt time.Time
}

// IsExpired checks if the cache entry is expired
func (ce *CacheEntry) IsExpired() bool {
	return time.Now().After(ce.ExpiresAt)
}

// CacheImpl implements the Cache interface
type CacheImpl struct {
	entries        map[string]*CacheEntry
	mu             sync.RWMutex
	maxSize        int64
	evictionPolicy string
	stats          *interfaces.CacheStats
}

// NewCache creates a new cache
func NewCache(maxSize int64, evictionPolicy string) interfaces.Cache {
	return &CacheImpl{
		entries:        make(map[string]*CacheEntry),
		maxSize:        maxSize,
		evictionPolicy: evictionPolicy,
		stats: &interfaces.CacheStats{
			Hits:        0,
			Misses:      0,
			Size:        0,
			MaxSize:     maxSize,
			MemoryUsage: 0,
		},
	}
}

// Get gets a value from the cache
func (c *CacheImpl) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	entry, exists := c.entries[key]
	c.mu.RUnlock()

	if !exists {
		c.mu.Lock()
		c.stats.Misses++
		c.mu.Unlock()
		return nil, false
	}

	// Check if expired
	if entry.IsExpired() {
		c.mu.Lock()
		delete(c.entries, key)
		c.stats.Misses++
		c.stats.Size = int64(len(c.entries))
		c.mu.Unlock()
		return nil, false
	}

	c.mu.Lock()
	c.stats.Hits++
	c.mu.Unlock()

	return entry.Value, true
}

// Set sets a value in the cache
func (c *CacheImpl) Set(key string, value interface{}, ttl string) error {
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	if value == nil {
		return fmt.Errorf("value cannot be nil")
	}

	// Parse TTL
	duration, err := time.ParseDuration(ttl)
	if err != nil {
		return fmt.Errorf("invalid TTL: %w", err)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if we need to evict
	if int64(len(c.entries)) >= c.maxSize {
		if err := c.evict(); err != nil {
			return fmt.Errorf("failed to evict: %w", err)
		}
	}

	// Create entry
	entry := &CacheEntry{
		Value:     value,
		ExpiresAt: time.Now().Add(duration),
		CreatedAt: time.Now(),
	}

	c.entries[key] = entry
	c.stats.Size = int64(len(c.entries))

	return nil
}

// Delete deletes a value from the cache
func (c *CacheImpl) Delete(key string) error {
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.entries[key]; !exists {
		return fmt.Errorf("key %s not found", key)
	}

	delete(c.entries, key)
	c.stats.Size = int64(len(c.entries))

	return nil
}

// Clear clears all values from the cache
func (c *CacheImpl) Clear() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries = make(map[string]*CacheEntry)
	c.stats.Size = 0
	c.stats.Hits = 0
	c.stats.Misses = 0

	return nil
}

// GetStats returns cache statistics
func (c *CacheImpl) GetStats() *interfaces.CacheStats {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Return a copy to prevent external modification
	return &interfaces.CacheStats{
		Hits:        c.stats.Hits,
		Misses:      c.stats.Misses,
		Size:        c.stats.Size,
		MaxSize:     c.stats.MaxSize,
		MemoryUsage: c.stats.MemoryUsage,
	}
}

// evict evicts entries based on the eviction policy
func (c *CacheImpl) evict() error {
	switch c.evictionPolicy {
	case "allkeys-lru":
		return c.evictLRU()
	case "allkeys-lfu":
		return c.evictLFU()
	case "volatile-lru":
		return c.evictVolatileLRU()
	case "volatile-lfu":
		return c.evictVolatileLFU()
	case "volatile-ttl":
		return c.evictVolatileTTL()
	case "noeviction":
		return fmt.Errorf("no eviction policy set")
	default:
		return c.evictLRU() // Default to LRU
	}
}

// evictLRU evicts the least recently used entry
func (c *CacheImpl) evictLRU() error {
	if len(c.entries) == 0 {
		return nil
	}

	var oldestKey string
	var oldestTime time.Time

	for key, entry := range c.entries {
		if oldestKey == "" || entry.CreatedAt.Before(oldestTime) {
			oldestKey = key
			oldestTime = entry.CreatedAt
		}
	}

	if oldestKey != "" {
		delete(c.entries, oldestKey)
		c.stats.Size = int64(len(c.entries))
	}

	return nil
}

// evictLFU evicts the least frequently used entry
func (c *CacheImpl) evictLFU() error {
	// For simplicity, we'll use LRU as LFU requires additional tracking
	return c.evictLRU()
}

// evictVolatileLRU evicts the least recently used volatile entry
func (c *CacheImpl) evictVolatileLRU() error {
	// For simplicity, we'll use LRU as volatile requires additional tracking
	return c.evictLRU()
}

// evictVolatileLFU evicts the least frequently used volatile entry
func (c *CacheImpl) evictVolatileLFU() error {
	// For simplicity, we'll use LRU as volatile LFU requires additional tracking
	return c.evictLRU()
}

// evictVolatileTTL evicts the entry with the shortest TTL
func (c *CacheImpl) evictVolatileTTL() error {
	if len(c.entries) == 0 {
		return nil
	}

	var shortestKey string
	var shortestTTL time.Duration

	for key, entry := range c.entries {
		ttl := time.Until(entry.ExpiresAt)
		if shortestKey == "" || ttl < shortestTTL {
			shortestKey = key
			shortestTTL = ttl
		}
	}

	if shortestKey != "" {
		delete(c.entries, shortestKey)
		c.stats.Size = int64(len(c.entries))
	}

	return nil
}

// Has checks if a key exists in the cache
func (c *CacheImpl) Has(key string) bool {
	c.mu.RLock()
	entry, exists := c.entries[key]
	c.mu.RUnlock()

	if !exists {
		return false
	}

	// Check if expired
	if entry.IsExpired() {
		c.mu.Lock()
		delete(c.entries, key)
		c.stats.Size = int64(len(c.entries))
		c.mu.Unlock()
		return false
	}

	return true
}

// Keys returns all keys in the cache
func (c *CacheImpl) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]string, 0, len(c.entries))
	for key := range c.entries {
		keys = append(keys, key)
	}

	return keys
}

// Size returns the current size of the cache
func (c *CacheImpl) Size() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return int64(len(c.entries))
}

// MaxSize returns the maximum size of the cache
func (c *CacheImpl) MaxSize() int64 {
	return c.maxSize
}

// SetMaxSize sets the maximum size of the cache
func (c *CacheImpl) SetMaxSize(maxSize int64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.maxSize = maxSize
	c.stats.MaxSize = maxSize
}

// GetEvictionPolicy returns the eviction policy
func (c *CacheImpl) GetEvictionPolicy() string {
	return c.evictionPolicy
}

// SetEvictionPolicy sets the eviction policy
func (c *CacheImpl) SetEvictionPolicy(policy string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.evictionPolicy = policy
}

// Cleanup removes expired entries
func (c *CacheImpl) Cleanup() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	removed := int64(0)
	for key, entry := range c.entries {
		if entry.IsExpired() {
			delete(c.entries, key)
			removed++
		}
	}

	c.stats.Size = int64(len(c.entries))
	return removed
}

// Reset resets the cache statistics
func (c *CacheImpl) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.stats.Hits = 0
	c.stats.Misses = 0
}

// GetHitRate returns the cache hit rate
func (c *CacheImpl) GetHitRate() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	total := c.stats.Hits + c.stats.Misses
	if total == 0 {
		return 0.0
	}

	return float64(c.stats.Hits) / float64(total)
}

// GetUsagePercentage returns the cache usage percentage
func (c *CacheImpl) GetUsagePercentage() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.stats.MaxSize == 0 {
		return 0.0
	}

	return float64(c.stats.Size) / float64(c.stats.MaxSize) * 100.0
}
