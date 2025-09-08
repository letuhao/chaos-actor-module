package cache

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// MultiLayerCacheManager manages the three-layer cache system
type MultiLayerCacheManager struct {
	l1Cache *LockFreeL1Cache
	l2Cache *MemoryMappedL2Cache
	l3Cache *PersistentL3Cache

	// Configuration
	config *MultiLayerConfig

	// Statistics
	stats *MultiLayerStats

	// State
	closed int32
	mu     sync.RWMutex
}

// MultiLayerConfig holds configuration for the multi-layer cache
type MultiLayerConfig struct {
	// L1 Cache settings
	L1MaxSize        int64
	L1EvictionPolicy string

	// L2 Cache settings
	L2CachePath string
	L2MaxSize   int64

	// L3 Cache settings
	L3CacheDir    string
	L3MaxSize     int64
	L3Compression bool

	// Performance settings
	EnablePreloading bool
	PreloadWorkers   int
	SyncInterval     time.Duration
}

// MultiLayerStats holds statistics for the multi-layer cache
type MultiLayerStats struct {
	// Overall stats
	TotalHits   int64
	TotalMisses int64
	TotalSets   int64
	TotalGets   int64

	// Layer-specific stats
	L1Hits   int64
	L1Misses int64
	L2Hits   int64
	L2Misses int64
	L3Hits   int64
	L3Misses int64

	// Performance stats
	AverageLatency time.Duration
	LastSyncTime   time.Time
	SyncCount      int64
}

// NewMultiLayerCacheManager creates a new multi-layer cache manager
func NewMultiLayerCacheManager(config *MultiLayerConfig) (*MultiLayerCacheManager, error) {
	if config == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	// Create L1 cache
	l1Cache := NewLockFreeL1Cache(config.L1MaxSize, config.L1EvictionPolicy)

	// Create L2 cache
	l2Cache, err := NewMemoryMappedL2Cache(config.L2CachePath, config.L2MaxSize)
	if err != nil {
		return nil, fmt.Errorf("failed to create L2 cache: %w", err)
	}

	// Create L3 cache
	l3Cache, err := NewPersistentL3Cache(config.L3CacheDir, config.L3MaxSize, config.L3Compression)
	if err != nil {
		l2Cache.Close()
		return nil, fmt.Errorf("failed to create L3 cache: %w", err)
	}

	manager := &MultiLayerCacheManager{
		l1Cache: l1Cache,
		l2Cache: l2Cache,
		l3Cache: l3Cache,
		config:  config,
		stats:   &MultiLayerStats{},
	}

	// Start background sync if enabled
	if config.SyncInterval > 0 {
		go manager.startBackgroundSync()
	}

	return manager, nil
}

// Get retrieves a value from the multi-layer cache
func (m *MultiLayerCacheManager) Get(key string) (interface{}, bool) {
	if key == "" {
		atomic.AddInt64(&m.stats.TotalMisses, 1)
		return nil, false
	}

	atomic.AddInt64(&m.stats.TotalGets, 1)
	start := time.Now()

	// Try L1 cache first (fastest)
	if value, found := m.l1Cache.Get(key); found {
		atomic.AddInt64(&m.stats.L1Hits, 1)
		atomic.AddInt64(&m.stats.TotalHits, 1)
		m.updateLatency(time.Since(start))
		return value, true
	}
	atomic.AddInt64(&m.stats.L1Misses, 1)

	// Try L2 cache (fast)
	if value, found := m.l2Cache.Get(key); found {
		atomic.AddInt64(&m.stats.L2Hits, 1)
		atomic.AddInt64(&m.stats.TotalHits, 1)

		// Promote to L1 cache
		m.l1Cache.Set(key, value, time.Hour)
		m.updateLatency(time.Since(start))
		return value, true
	}
	atomic.AddInt64(&m.stats.L2Misses, 1)

	// Try L3 cache (persistent)
	if value, found := m.l3Cache.Get(key); found {
		atomic.AddInt64(&m.stats.L3Hits, 1)
		atomic.AddInt64(&m.stats.TotalHits, 1)

		// Promote to L1 and L2 caches
		m.l1Cache.Set(key, value, time.Hour)
		m.l2Cache.Set(key, value, time.Hour)
		m.updateLatency(time.Since(start))
		return value, true
	}
	atomic.AddInt64(&m.stats.L3Misses, 1)
	atomic.AddInt64(&m.stats.TotalMisses, 1)

	m.updateLatency(time.Since(start))
	return nil, false
}

// Set stores a value in the multi-layer cache
func (m *MultiLayerCacheManager) Set(key string, value interface{}, ttl time.Duration) error {
	if key == "" {
		return ErrEmptyKey
	}

	if value == nil {
		return ErrNilValue
	}

	if atomic.LoadInt32(&m.closed) == 1 {
		return fmt.Errorf("cache manager is closed")
	}

	atomic.AddInt64(&m.stats.TotalSets, 1)
	start := time.Now()

	// Set in all layers
	if err := m.l1Cache.Set(key, value, ttl); err != nil {
		return fmt.Errorf("failed to set in L1 cache: %w", err)
	}

	if err := m.l2Cache.Set(key, value, ttl); err != nil {
		return fmt.Errorf("failed to set in L2 cache: %w", err)
	}

	if err := m.l3Cache.Set(key, value, ttl); err != nil {
		return fmt.Errorf("failed to set in L3 cache: %w", err)
	}

	m.updateLatency(time.Since(start))
	return nil
}

// Delete removes a value from all cache layers
func (m *MultiLayerCacheManager) Delete(key string) error {
	if key == "" {
		return ErrEmptyKey
	}

	// Delete from all layers
	m.l1Cache.Delete(key)
	m.l2Cache.Delete(key)
	m.l3Cache.Delete(key)

	return nil
}

// Clear clears all cache layers
func (m *MultiLayerCacheManager) Clear() error {
	m.l1Cache.Clear()
	m.l2Cache.Clear()
	m.l3Cache.Clear()
	return nil
}

// Has checks if a key exists in any cache layer
func (m *MultiLayerCacheManager) Has(key string) bool {
	return m.l1Cache.Has(key) || m.l2Cache.Has(key) || m.l3Cache.Has(key)
}

// GetStats returns comprehensive statistics for all layers
func (m *MultiLayerCacheManager) GetStats() *MultiLayerStats {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Get individual layer stats (for future use)
	_ = m.l1Cache.GetStats()
	_ = m.l2Cache.GetStats()
	_ = m.l3Cache.GetStats()

	// Calculate hit rates
	totalHits := atomic.LoadInt64(&m.stats.TotalHits)
	totalMisses := atomic.LoadInt64(&m.stats.TotalMisses)

	stats := &MultiLayerStats{
		TotalHits:      totalHits,
		TotalMisses:    totalMisses,
		TotalSets:      atomic.LoadInt64(&m.stats.TotalSets),
		TotalGets:      atomic.LoadInt64(&m.stats.TotalGets),
		L1Hits:         atomic.LoadInt64(&m.stats.L1Hits),
		L1Misses:       atomic.LoadInt64(&m.stats.L1Misses),
		L2Hits:         atomic.LoadInt64(&m.stats.L2Hits),
		L2Misses:       atomic.LoadInt64(&m.stats.L2Misses),
		L3Hits:         atomic.LoadInt64(&m.stats.L3Hits),
		L3Misses:       atomic.LoadInt64(&m.stats.L3Misses),
		AverageLatency: m.stats.AverageLatency,
		LastSyncTime:   m.stats.LastSyncTime,
		SyncCount:      atomic.LoadInt64(&m.stats.SyncCount),
	}

	return stats
}

// GetHitRate returns the overall hit rate
func (m *MultiLayerCacheManager) GetHitRate() float64 {
	totalHits := atomic.LoadInt64(&m.stats.TotalHits)
	totalMisses := atomic.LoadInt64(&m.stats.TotalMisses)

	if totalHits+totalMisses == 0 {
		return 0.0
	}

	return float64(totalHits) / float64(totalHits+totalMisses) * 100.0
}

// GetLayerHitRates returns hit rates for each layer
func (m *MultiLayerCacheManager) GetLayerHitRates() map[string]float64 {
	rates := make(map[string]float64)

	// L1 hit rate
	l1Hits := atomic.LoadInt64(&m.stats.L1Hits)
	l1Misses := atomic.LoadInt64(&m.stats.L1Misses)
	if l1Hits+l1Misses > 0 {
		rates["L1"] = float64(l1Hits) / float64(l1Hits+l1Misses) * 100.0
	} else {
		rates["L1"] = 0.0
	}

	// L2 hit rate
	l2Hits := atomic.LoadInt64(&m.stats.L2Hits)
	l2Misses := atomic.LoadInt64(&m.stats.L2Misses)
	if l2Hits+l2Misses > 0 {
		rates["L2"] = float64(l2Hits) / float64(l2Hits+l2Misses) * 100.0
	} else {
		rates["L2"] = 0.0
	}

	// L3 hit rate
	l3Hits := atomic.LoadInt64(&m.stats.L3Hits)
	l3Misses := atomic.LoadInt64(&m.stats.L3Misses)
	if l3Hits+l3Misses > 0 {
		rates["L3"] = float64(l3Hits) / float64(l3Hits+l3Misses) * 100.0
	} else {
		rates["L3"] = 0.0
	}

	return rates
}

// GetMemoryUsage returns memory usage for each layer
func (m *MultiLayerCacheManager) GetMemoryUsage() map[string]int64 {
	usage := make(map[string]int64)

	l1Stats := m.l1Cache.GetStats()
	l2Stats := m.l2Cache.GetStats()
	l3Stats := m.l3Cache.GetStats()

	usage["L1"] = l1Stats.memoryUsage
	usage["L2"] = l2Stats.memoryUsage
	usage["L3"] = l3Stats.memoryUsage

	return usage
}

// Sync synchronizes data between cache layers
func (m *MultiLayerCacheManager) Sync() error {
	if atomic.LoadInt32(&m.closed) == 1 {
		return fmt.Errorf("cache manager is closed")
	}

	// Sync L2 to L1 (promote frequently accessed items)
	l2Keys := m.l2Cache.Keys()
	for _, key := range l2Keys {
		if !m.l1Cache.Has(key) {
			if value, found := m.l2Cache.Get(key); found {
				m.l1Cache.Set(key, value, time.Hour)
			}
		}
	}

	// Sync L3 to L2 (promote frequently accessed items)
	l3Keys := m.l3Cache.Keys()
	for _, key := range l3Keys {
		if !m.l2Cache.Has(key) {
			if value, found := m.l3Cache.Get(key); found {
				m.l2Cache.Set(key, value, time.Hour)
			}
		}
	}

	atomic.AddInt64(&m.stats.SyncCount, 1)
	m.stats.LastSyncTime = time.Now()

	return nil
}

// Close closes all cache layers
func (m *MultiLayerCacheManager) Close() error {
	atomic.StoreInt32(&m.closed, 1)

	var err error

	if closeErr := m.l2Cache.Close(); closeErr != nil {
		err = closeErr
	}

	if closeErr := m.l3Cache.Close(); closeErr != nil {
		if err != nil {
			err = fmt.Errorf("multiple errors: %v, %v", err, closeErr)
		} else {
			err = closeErr
		}
	}

	return err
}

// startBackgroundSync starts the background synchronization process
func (m *MultiLayerCacheManager) startBackgroundSync() {
	ticker := time.NewTicker(m.config.SyncInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if atomic.LoadInt32(&m.closed) == 1 {
				return
			}
			m.Sync()
		}
	}
}

// updateLatency updates the average latency
func (m *MultiLayerCacheManager) updateLatency(duration time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Simple moving average
	if m.stats.AverageLatency == 0 {
		m.stats.AverageLatency = duration
	} else {
		m.stats.AverageLatency = (m.stats.AverageLatency + duration) / 2
	}
}

// Preload preloads data into the cache
func (m *MultiLayerCacheManager) Preload(data map[string]interface{}) error {
	if atomic.LoadInt32(&m.closed) == 1 {
		return fmt.Errorf("cache manager is closed")
	}

	// Preload into L1 cache (most frequently accessed)
	for key, value := range data {
		if err := m.l1Cache.Set(key, value, time.Hour); err != nil {
			return fmt.Errorf("failed to preload key %s: %w", key, err)
		}
	}

	return nil
}

// GetConfig returns the current configuration
func (m *MultiLayerCacheManager) GetConfig() *MultiLayerConfig {
	return m.config
}

// GetL1Cache returns the L1 cache instance
func (m *MultiLayerCacheManager) GetL1Cache() *LockFreeL1Cache {
	return m.l1Cache
}

// GetL2Cache returns the L2 cache instance
func (m *MultiLayerCacheManager) GetL2Cache() *MemoryMappedL2Cache {
	return m.l2Cache
}

// GetL3Cache returns the L3 cache instance
func (m *MultiLayerCacheManager) GetL3Cache() *PersistentL3Cache {
	return m.l3Cache
}

// SetConfig updates the configuration
func (m *MultiLayerCacheManager) SetConfig(config *MultiLayerConfig) error {
	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.config = config
	return nil
}
