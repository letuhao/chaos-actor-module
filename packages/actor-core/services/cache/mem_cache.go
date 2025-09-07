package cache

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"unsafe"
)

// MemCache implements in-memory cache with LRU eviction
type MemCache struct {
	data    map[string]*CacheEntry
	mutex   sync.RWMutex
	config  MemCacheConfig
	stats   *CacheStats
	created int64
}

// NewMemCache creates a new MemCache instance
func NewMemCache(config MemCacheConfig) *MemCache {
	now := time.Now().Unix()
	return &MemCache{
		data:    make(map[string]*CacheEntry),
		config:  config,
		stats:   &CacheStats{LastReset: now},
		created: now,
	}
}

// Get retrieves a value from cache
func (mc *MemCache) Get(key string) (interface{}, error) {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	entry, exists := mc.data[key]
	if !exists {
		mc.stats.Misses++
		mc.updateHitRatio()
		return nil, fmt.Errorf("key not found: %s", key)
	}

	// Check if expired
	if entry.ExpiresAt > 0 && time.Now().Unix() >= entry.ExpiresAt {
		mc.mutex.RUnlock()
		mc.mutex.Lock()
		delete(mc.data, key)
		mc.stats.Size--
		mc.stats.Evictions++
		mc.mutex.Unlock()
		mc.mutex.RLock()

		mc.stats.Misses++
		mc.updateHitRatio()
		return nil, fmt.Errorf("key expired: %s", key)
	}

	// Update access statistics
	entry.AccessCount++
	entry.LastAccess = time.Now().Unix()
	mc.stats.Hits++
	mc.updateHitRatio()

	return entry.Value, nil
}

// Set stores a value in cache
func (mc *MemCache) Set(key string, value interface{}, ttl time.Duration) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	now := time.Now().Unix()
	var expiresAt int64

	if ttl > 0 {
		expiresAt = time.Now().Add(ttl).Unix()
	}

	// Calculate size
	size := mc.calculateSize(value)

	// Check if we need to evict
	if mc.needsEviction(size) {
		mc.evictEntries(size)
	}

	// Create or update entry
	entry := &CacheEntry{
		Value:       value,
		ExpiresAt:   expiresAt,
		CreatedAt:   now,
		AccessCount: 1,
		LastAccess:  now,
		Size:        size,
	}

	// Check if key already exists
	if existingEntry, exists := mc.data[key]; exists {
		mc.stats.Size -= existingEntry.Size
	}

	mc.data[key] = entry
	mc.stats.Size += size

	return nil
}

// Delete removes a key from cache
func (mc *MemCache) Delete(key string) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	entry, exists := mc.data[key]
	if !exists {
		return fmt.Errorf("key not found: %s", key)
	}

	delete(mc.data, key)
	mc.stats.Size -= entry.Size

	return nil
}

// Clear removes all keys from cache
func (mc *MemCache) Clear() error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	mc.data = make(map[string]*CacheEntry)
	mc.stats.Size = 0

	return nil
}

// Exists checks if a key exists in cache
func (mc *MemCache) Exists(key string) bool {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	entry, exists := mc.data[key]
	if !exists {
		return false
	}

	// Check if expired
	if entry.ExpiresAt > 0 && time.Now().Unix() >= entry.ExpiresAt {
		// Key is expired, remove it
		mc.mutex.RUnlock()
		mc.mutex.Lock()
		delete(mc.data, key)
		mc.stats.Size -= entry.Size
		mc.stats.Evictions++
		mc.mutex.Unlock()
		mc.mutex.RLock()
		return false
	}

	return true
}

// GetTTL returns the time to live for a key
func (mc *MemCache) GetTTL(key string) (time.Duration, error) {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	entry, exists := mc.data[key]
	if !exists {
		return 0, fmt.Errorf("key not found: %s", key)
	}

	if entry.ExpiresAt <= 0 {
		return 0, nil // No expiration
	}

	now := time.Now().Unix()
	if now >= entry.ExpiresAt {
		return 0, fmt.Errorf("key expired: %s", key)
	}

	return time.Duration(entry.ExpiresAt-now) * time.Second, nil
}

// SetTTL sets the time to live for a key
func (mc *MemCache) SetTTL(key string, ttl time.Duration) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	entry, exists := mc.data[key]
	if !exists {
		return fmt.Errorf("key not found: %s", key)
	}

	if ttl <= 0 {
		entry.ExpiresAt = 0
	} else {
		entry.ExpiresAt = time.Now().Unix() + int64(ttl.Seconds())
	}

	return nil
}

// GetStats returns cache statistics
func (mc *MemCache) GetStats() *CacheStats {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	// Create a copy to avoid race conditions
	stats := *mc.stats
	return &stats
}

// ResetStats resets cache statistics
func (mc *MemCache) ResetStats() error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	now := time.Now().Unix()
	mc.stats = &CacheStats{
		LastReset: now,
	}

	return nil
}

// Health checks cache health
func (mc *MemCache) Health() error {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	// Check if cache is not too large
	if mc.stats.Size > mc.getMaxSize() {
		return fmt.Errorf("cache size exceeds maximum: %d > %d", mc.stats.Size, mc.getMaxSize())
	}

	// Check if hit ratio is reasonable
	if mc.stats.Hits+mc.stats.Misses > 0 && mc.stats.HitRatio < 0.1 {
		return fmt.Errorf("cache hit ratio too low: %.2f", mc.stats.HitRatio)
	}

	return nil
}

// Private methods

// updateHitRatio updates the hit ratio
func (mc *MemCache) updateHitRatio() {
	total := mc.stats.Hits + mc.stats.Misses
	if total > 0 {
		mc.stats.HitRatio = float64(mc.stats.Hits) / float64(total)
	}
}

// calculateSize calculates the size of a value
func (mc *MemCache) calculateSize(value interface{}) int64 {
	// Serialize to JSON to get approximate size
	data, err := json.Marshal(value)
	if err != nil {
		// Fallback to unsafe.Sizeof
		return int64(unsafe.Sizeof(value))
	}
	return int64(len(data))
}

// needsEviction checks if eviction is needed
func (mc *MemCache) needsEviction(additionalSize int64) bool {
	maxSize := mc.getMaxSize()
	return mc.stats.Size+additionalSize > maxSize
}

// getMaxSize returns the maximum cache size in bytes
func (mc *MemCache) getMaxSize() int64 {
	// Parse size string (e.g., "100MB", "1GB")
	// For simplicity, assume it's in MB
	// In a real implementation, you'd parse the size string properly
	return 100 * 1024 * 1024 // 100MB default
}

// evictEntries evicts entries based on the eviction policy
func (mc *MemCache) evictEntries(requiredSize int64) {
	switch mc.config.EvictionPolicy {
	case "lru":
		mc.evictLRU(requiredSize)
	case "lfu":
		mc.evictLFU(requiredSize)
	case "ttl":
		mc.evictTTL(requiredSize)
	default:
		mc.evictLRU(requiredSize)
	}
}

// evictLRU evicts least recently used entries
func (mc *MemCache) evictLRU(requiredSize int64) {
	// Find entries with oldest last access time
	var oldestKey string
	var oldestTime int64 = time.Now().Unix()

	for key, entry := range mc.data {
		if entry.LastAccess < oldestTime {
			oldestTime = entry.LastAccess
			oldestKey = key
		}
	}

	if oldestKey != "" {
		entry := mc.data[oldestKey]
		delete(mc.data, oldestKey)
		mc.stats.Size -= entry.Size
		mc.stats.Evictions++
	}
}

// evictLFU evicts least frequently used entries
func (mc *MemCache) evictLFU(requiredSize int64) {
	// Find entries with lowest access count
	var leastUsedKey string
	var leastUsedCount int64 = ^int64(0) // Max int64

	for key, entry := range mc.data {
		if entry.AccessCount < leastUsedCount {
			leastUsedCount = entry.AccessCount
			leastUsedKey = key
		}
	}

	if leastUsedKey != "" {
		entry := mc.data[leastUsedKey]
		delete(mc.data, leastUsedKey)
		mc.stats.Size -= entry.Size
		mc.stats.Evictions++
	}
}

// evictTTL evicts entries closest to expiration
func (mc *MemCache) evictTTL(requiredSize int64) {
	// Find entries closest to expiration
	var closestKey string
	var closestExpiry int64 = ^int64(0) // Max int64

	for key, entry := range mc.data {
		if entry.ExpiresAt > 0 && entry.ExpiresAt < closestExpiry {
			closestExpiry = entry.ExpiresAt
			closestKey = key
		}
	}

	if closestKey != "" {
		entry := mc.data[closestKey]
		delete(mc.data, closestKey)
		mc.stats.Size -= entry.Size
		mc.stats.Evictions++
	}
}

// StartCleanup starts the cleanup goroutine
func (mc *MemCache) StartCleanup() {
	go func() {
		ticker := time.NewTicker(mc.config.CleanupInterval)
		defer ticker.Stop()

		for range ticker.C {
			mc.CleanupExpired()
		}
	}()
}

// CleanupExpired removes expired entries
func (mc *MemCache) CleanupExpired() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	now := time.Now().Unix()
	for key, entry := range mc.data {
		if entry.ExpiresAt > 0 && now >= entry.ExpiresAt {
			delete(mc.data, key)
			mc.stats.Size -= entry.Size
			mc.stats.Evictions++
		}
	}
}

// GetSize returns the current cache size
func (mc *MemCache) GetSize() int64 {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()
	return mc.stats.Size
}

// GetEntryCount returns the number of entries in cache
func (mc *MemCache) GetEntryCount() int {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()
	return len(mc.data)
}

// GetKeys returns all cache keys
func (mc *MemCache) GetKeys() []string {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	keys := make([]string, 0, len(mc.data))
	for key := range mc.data {
		keys = append(keys, key)
	}
	return keys
}
