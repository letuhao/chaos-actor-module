package cache

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

// BenchmarkConfig holds configuration for benchmarks
type BenchmarkConfig struct {
	// Test parameters
	KeyCount    int
	ValueSize   int
	Duration    time.Duration
	Concurrency int

	// Cache settings
	L1MaxSize int64
	L2MaxSize int64
	L3MaxSize int64

	// Performance targets
	TargetLatency     time.Duration
	TargetThroughput  int64
	TargetMemoryUsage int64
}

// BenchmarkResults holds benchmark results
type BenchmarkResults struct {
	// Performance metrics
	Latency     time.Duration
	Throughput  float64
	MemoryUsage int64
	CPUUsage    float64

	// Cache metrics
	HitRate      float64
	MissRate     float64
	EvictionRate float64

	// System metrics
	GoroutineCount int
	GCStats        runtime.MemStats

	// Test configuration
	Config     *BenchmarkConfig
	Duration   time.Duration
	Operations int64
}

// PerformanceMonitor monitors cache performance
type PerformanceMonitor struct {
	// Configuration
	config *BenchmarkConfig

	// Metrics collection
	startTime  time.Time
	endTime    time.Time
	operations int64
	latencies  []time.Duration

	// Memory tracking
	startMem runtime.MemStats
	endMem   runtime.MemStats

	// Cache references
	l1Cache     *LockFreeL1Cache
	l2Cache     *MemoryMappedL2Cache
	l3Cache     *PersistentL3Cache
	multiCache  *MultiLayerCacheManager
	warmer      *CacheWarmer
	invalidator *CacheInvalidator

	// Synchronization
	mu sync.Mutex
}

// NewPerformanceMonitor creates a new performance monitor
func NewPerformanceMonitor(config *BenchmarkConfig) *PerformanceMonitor {
	return &PerformanceMonitor{
		config:    config,
		latencies: make([]time.Duration, 0, 10000),
	}
}

// Start begins performance monitoring
func (pm *PerformanceMonitor) Start() {
	pm.startTime = time.Now()
	runtime.ReadMemStats(&pm.startMem)
}

// Stop ends performance monitoring
func (pm *PerformanceMonitor) Stop() {
	pm.endTime = time.Now()
	runtime.ReadMemStats(&pm.endMem)
}

// RecordOperation records a single operation
func (pm *PerformanceMonitor) RecordOperation(latency time.Duration) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.operations++
	pm.latencies = append(pm.latencies, latency)
}

// GetResults returns benchmark results
func (pm *PerformanceMonitor) GetResults() *BenchmarkResults {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	duration := pm.endTime.Sub(pm.startTime)

	// Calculate latency statistics
	var totalLatency time.Duration
	for _, latency := range pm.latencies {
		totalLatency += latency
	}

	avgLatency := time.Duration(0)
	if len(pm.latencies) > 0 {
		avgLatency = totalLatency / time.Duration(len(pm.latencies))
	}

	// Calculate throughput
	throughput := float64(pm.operations) / duration.Seconds()

	// Calculate memory usage
	memoryUsage := int64(pm.endMem.Alloc - pm.startMem.Alloc)

	// Calculate CPU usage (simplified)
	cpuUsage := float64(pm.operations) / duration.Seconds() / 1000.0

	// Get cache metrics
	hitRate := 0.0
	missRate := 0.0
	evictionRate := 0.0

	if pm.l1Cache != nil {
		l1Stats := pm.l1Cache.GetStats()
		if l1Stats.hits+l1Stats.misses > 0 {
			hitRate = float64(l1Stats.hits) / float64(l1Stats.hits+l1Stats.misses) * 100.0
			missRate = float64(l1Stats.misses) / float64(l1Stats.hits+l1Stats.misses) * 100.0
		}
	}

	return &BenchmarkResults{
		Latency:        avgLatency,
		Throughput:     throughput,
		MemoryUsage:    memoryUsage,
		CPUUsage:       cpuUsage,
		HitRate:        hitRate,
		MissRate:       missRate,
		EvictionRate:   evictionRate,
		GoroutineCount: runtime.NumGoroutine(),
		GCStats:        pm.endMem,
		Config:         pm.config,
		Duration:       duration,
		Operations:     pm.operations,
	}
}

// BenchmarkL1Cache benchmarks L1 cache performance
func BenchmarkL1Cache(b *testing.B) {
	config := &BenchmarkConfig{
		KeyCount:         1000,
		ValueSize:        1024,
		Concurrency:      10,
		L1MaxSize:        10000,
		TargetLatency:    time.Microsecond * 100,
		TargetThroughput: 100000,
	}

	cache := NewLockFreeL1Cache(config.L1MaxSize, "allkeys-lru")
	defer cache.Clear()

	// Generate test data
	keys := generateKeys(config.KeyCount)
	values := generateValues(config.KeyCount, config.ValueSize)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := keys[i%len(keys)]
			value := values[i%len(values)]

			start := time.Now()
			cache.Set(key, value, time.Hour)
			latency := time.Since(start)

			// Record operation
			_ = latency

			i++
		}
	})
}

// BenchmarkL2Cache benchmarks L2 cache performance
func BenchmarkL2Cache(b *testing.B) {
	config := &BenchmarkConfig{
		KeyCount:         1000,
		ValueSize:        1024,
		Concurrency:      10,
		L2MaxSize:        100000,
		TargetLatency:    time.Microsecond * 500,
		TargetThroughput: 50000,
	}

	cache, err := NewMemoryMappedL2Cache("test_l2_cache.dat", config.L2MaxSize)
	if err != nil {
		b.Fatalf("Failed to create L2 cache: %v", err)
	}
	defer cache.Close()

	// Generate test data
	keys := generateKeys(config.KeyCount)
	values := generateValues(config.KeyCount, config.ValueSize)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := keys[i%len(keys)]
			value := values[i%len(values)]

			start := time.Now()
			cache.Set(key, value, time.Hour)
			latency := time.Since(start)

			// Record operation
			_ = latency

			i++
		}
	})
}

// BenchmarkL3Cache benchmarks L3 cache performance
func BenchmarkL3Cache(b *testing.B) {
	config := &BenchmarkConfig{
		KeyCount:         1000,
		ValueSize:        1024,
		Concurrency:      10,
		L3MaxSize:        1000000,
		TargetLatency:    time.Millisecond * 1,
		TargetThroughput: 10000,
	}

	cache, err := NewPersistentL3Cache("test_l3_cache", config.L3MaxSize, true)
	if err != nil {
		b.Fatalf("Failed to create L3 cache: %v", err)
	}
	defer cache.Close()

	// Generate test data
	keys := generateKeys(config.KeyCount)
	values := generateValues(config.KeyCount, config.ValueSize)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := keys[i%len(keys)]
			value := values[i%len(values)]

			start := time.Now()
			cache.Set(key, value, time.Hour)
			latency := time.Since(start)

			// Record operation
			_ = latency

			i++
		}
	})
}

// BenchmarkMultiLayerCache benchmarks multi-layer cache performance
func BenchmarkMultiLayerCache(b *testing.B) {
	config := &BenchmarkConfig{
		KeyCount:         1000,
		ValueSize:        1024,
		Concurrency:      10,
		L1MaxSize:        10000,
		L2MaxSize:        100000,
		L3MaxSize:        1000000,
		TargetLatency:    time.Microsecond * 200,
		TargetThroughput: 75000,
	}

	// Create cache layers
	_ = NewLockFreeL1Cache(config.L1MaxSize, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache("test_multi_l2.dat", config.L2MaxSize)
	if err != nil {
		b.Fatalf("Failed to create L2 cache: %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache("test_multi_l3", config.L3MaxSize, true)
	if err != nil {
		b.Fatalf("Failed to create L3 cache: %v", err)
	}
	defer l3Cache.Close()

	// Create multi-layer cache manager
	multiConfig := &MultiLayerConfig{
		L1MaxSize:        config.L1MaxSize,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      "test_multi_l2.dat",
		L2MaxSize:        config.L2MaxSize,
		L3CacheDir:       "test_multi_l3",
		L3MaxSize:        config.L3MaxSize,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	multiCache, err := NewMultiLayerCacheManager(multiConfig)
	if err != nil {
		b.Fatalf("Failed to create multi-layer cache: %v", err)
	}
	defer multiCache.Close()

	// Generate test data
	keys := generateKeys(config.KeyCount)
	values := generateValues(config.KeyCount, config.ValueSize)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := keys[i%len(keys)]
			value := values[i%len(values)]

			start := time.Now()
			multiCache.Set(key, value, time.Hour)
			latency := time.Since(start)

			// Record operation
			_ = latency

			i++
		}
	})
}

// BenchmarkCacheWarmer benchmarks cache warmer performance
func BenchmarkCacheWarmer(b *testing.B) {
	config := &BenchmarkConfig{
		KeyCount:         1000,
		ValueSize:        1024,
		Concurrency:      10,
		L1MaxSize:        10000,
		L2MaxSize:        100000,
		L3MaxSize:        1000000,
		TargetLatency:    time.Microsecond * 300,
		TargetThroughput: 50000,
	}

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(config.L1MaxSize, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache("test_warmer_l2.dat", config.L2MaxSize)
	if err != nil {
		b.Fatalf("Failed to create L2 cache: %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache("test_warmer_l3", config.L3MaxSize, true)
	if err != nil {
		b.Fatalf("Failed to create L3 cache: %v", err)
	}
	defer l3Cache.Close()

	// Create cache warmer
	warmerConfig := &CacheWarmerConfig{
		EnableWarming:       true,
		WarmingInterval:     time.Second,
		MaxWarmingWorkers:   4,
		EnablePrediction:    true,
		PredictionWindow:    time.Minute,
		MinAccessCount:      5,
		ConfidenceThreshold: 0.7,
		PreloadBatchSize:    100,
		PreloadTimeout:      time.Second * 30,
		MaxPreloadSize:      1024 * 1024,
		WarmingPriority:     1,
		EnableHotPath:       true,
		HotPathThreshold:    0.8,
	}

	warmer := NewCacheWarmer(warmerConfig, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Generate test data
	keys := generateKeys(config.KeyCount)
	values := generateValues(config.KeyCount, config.ValueSize)

	// Pre-populate L3 cache
	for i, key := range keys {
		l3Cache.Set(key, values[i], time.Hour)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := keys[i%len(keys)]

			start := time.Now()
			warmer.RecordAccess(key)
			latency := time.Since(start)

			// Record operation
			_ = latency

			i++
		}
	})
}

// BenchmarkCacheInvalidator benchmarks cache invalidator performance
func BenchmarkCacheInvalidator(b *testing.B) {
	config := &BenchmarkConfig{
		KeyCount:         1000,
		ValueSize:        1024,
		Concurrency:      10,
		L1MaxSize:        10000,
		L2MaxSize:        100000,
		L3MaxSize:        1000000,
		TargetLatency:    time.Microsecond * 400,
		TargetThroughput: 40000,
	}

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(config.L1MaxSize, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache("test_invalidator_l2.dat", config.L2MaxSize)
	if err != nil {
		b.Fatalf("Failed to create L2 cache: %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache("test_invalidator_l3", config.L3MaxSize, true)
	if err != nil {
		b.Fatalf("Failed to create L3 cache: %v", err)
	}
	defer l3Cache.Close()

	// Create cache invalidator
	invalidatorConfig := &CacheInvalidatorConfig{
		EnableInvalidation:     true,
		InvalidationInterval:   time.Second,
		MaxInvalidationWorkers: 4,
		EnableTTL:              true,
		TTLCheckInterval:       time.Second,
		DefaultTTL:             time.Hour,
		MaxTTL:                 time.Hour * 24,
		EnableDependencies:     true,
		MaxDependencyDepth:     10,
		DependencyTimeout:      time.Second * 30,
		InvalidationPriority:   1,
		BatchSize:              100,
		EnableLazyInvalidation: true,
		LazyThreshold:          1000,
	}

	invalidator := NewCacheInvalidator(invalidatorConfig, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Generate test data
	keys := generateKeys(config.KeyCount)
	values := generateValues(config.KeyCount, config.ValueSize)

	// Pre-populate caches
	for i, key := range keys {
		l1Cache.Set(key, values[i], time.Hour)
		l2Cache.Set(key, values[i], time.Hour)
		l3Cache.Set(key, values[i], time.Hour)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := keys[i%len(keys)]

			start := time.Now()
			invalidator.Invalidate(key, InvalidationReasonExplicit)
			latency := time.Since(start)

			// Record operation
			_ = latency

			i++
		}
	})
}

// BenchmarkConcurrentAccess benchmarks concurrent access patterns
func BenchmarkConcurrentAccess(b *testing.B) {
	config := &BenchmarkConfig{
		KeyCount:         1000,
		ValueSize:        1024,
		Concurrency:      100,
		L1MaxSize:        10000,
		L2MaxSize:        100000,
		L3MaxSize:        1000000,
		TargetLatency:    time.Microsecond * 500,
		TargetThroughput: 20000,
	}

	// Create multi-layer cache
	_ = NewLockFreeL1Cache(config.L1MaxSize, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache("test_concurrent_l2.dat", config.L2MaxSize)
	if err != nil {
		b.Fatalf("Failed to create L2 cache: %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache("test_concurrent_l3", config.L3MaxSize, true)
	if err != nil {
		b.Fatalf("Failed to create L3 cache: %v", err)
	}
	defer l3Cache.Close()

	multiConfig := &MultiLayerConfig{
		L1MaxSize:        config.L1MaxSize,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      "test_concurrent_l2.dat",
		L2MaxSize:        config.L2MaxSize,
		L3CacheDir:       "test_concurrent_l3",
		L3MaxSize:        config.L3MaxSize,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	multiCache, err := NewMultiLayerCacheManager(multiConfig)
	if err != nil {
		b.Fatalf("Failed to create multi-layer cache: %v", err)
	}
	defer multiCache.Close()

	// Generate test data
	keys := generateKeys(config.KeyCount)
	values := generateValues(config.KeyCount, config.ValueSize)

	// Pre-populate cache
	for i, key := range keys {
		multiCache.Set(key, values[i], time.Hour)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := keys[i%len(keys)]

			start := time.Now()
			multiCache.Get(key)
			latency := time.Since(start)

			// Record operation
			_ = latency

			i++
		}
	})
}

// generateKeys generates test keys
func generateKeys(count int) []string {
	keys := make([]string, count)
	for i := 0; i < count; i++ {
		keys[i] = fmt.Sprintf("key_%d", i)
	}
	return keys
}

// generateValues generates test values
func generateValues(count int, size int) []string {
	values := make([]string, count)
	for i := 0; i < count; i++ {
		value := make([]byte, size)
		for j := 0; j < size; j++ {
			value[j] = byte(rand.Intn(256))
		}
		values[i] = string(value)
	}
	return values
}

// RunPerformanceTest runs a comprehensive performance test
func RunPerformanceTest(config *BenchmarkConfig) (*BenchmarkResults, error) {
	// Create performance monitor
	monitor := NewPerformanceMonitor(config)

	// Create multi-layer cache manager
	multiConfig := &MultiLayerConfig{
		L1MaxSize:        config.L1MaxSize,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      "perf_test_l2.dat",
		L2MaxSize:        config.L2MaxSize,
		L3CacheDir:       "perf_test_l3",
		L3MaxSize:        config.L3MaxSize,
		L3Compression:    true,
		EnablePreloading: true,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	multiCache, err := NewMultiLayerCacheManager(multiConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create multi-layer cache: %w", err)
	}
	defer multiCache.Close()

	// Get individual cache instances from multi-layer manager
	l1Cache := multiCache.GetL1Cache()
	l2Cache := multiCache.GetL2Cache()
	l3Cache := multiCache.GetL3Cache()

	// Create cache warmer
	warmerConfig := &CacheWarmerConfig{
		EnableWarming:       true,
		WarmingInterval:     time.Second,
		MaxWarmingWorkers:   4,
		EnablePrediction:    true,
		PredictionWindow:    time.Minute,
		MinAccessCount:      5,
		ConfidenceThreshold: 0.7,
		PreloadBatchSize:    100,
		PreloadTimeout:      time.Second * 30,
		MaxPreloadSize:      1024 * 1024,
		WarmingPriority:     1,
		EnableHotPath:       true,
		HotPathThreshold:    0.8,
	}

	warmer := NewCacheWarmer(warmerConfig, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Create cache invalidator
	invalidatorConfig := &CacheInvalidatorConfig{
		EnableInvalidation:     true,
		InvalidationInterval:   time.Second,
		MaxInvalidationWorkers: 4,
		EnableTTL:              true,
		TTLCheckInterval:       time.Second,
		DefaultTTL:             time.Hour,
		MaxTTL:                 time.Hour * 24,
		EnableDependencies:     true,
		MaxDependencyDepth:     10,
		DependencyTimeout:      time.Second * 30,
		InvalidationPriority:   1,
		BatchSize:              100,
		EnableLazyInvalidation: true,
		LazyThreshold:          1000,
	}

	invalidator := NewCacheInvalidator(invalidatorConfig, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Set monitor references
	monitor.l1Cache = l1Cache
	monitor.l2Cache = l2Cache
	monitor.l3Cache = l3Cache
	monitor.multiCache = multiCache
	monitor.warmer = warmer
	monitor.invalidator = invalidator

	// Generate test data
	keys := generateKeys(config.KeyCount)
	values := generateValues(config.KeyCount, config.ValueSize)

	// Start monitoring
	monitor.Start()

	// Run performance test
	done := make(chan bool)

	// Start workers
	for i := 0; i < config.Concurrency; i++ {
		go func(workerID int) {
			rand.Seed(time.Now().UnixNano() + int64(workerID))

			for {
				select {
				case <-done:
					return
				default:
					// Random operation
					operation := rand.Intn(4)
					key := keys[rand.Intn(len(keys))]
					value := values[rand.Intn(len(values))]

					start := time.Now()

					switch operation {
					case 0: // Set
						multiCache.Set(key, value, time.Hour)
					case 1: // Get
						multiCache.Get(key)
					case 2: // Record access
						warmer.RecordAccess(key)
					case 3: // Invalidate
						invalidator.Invalidate(key, InvalidationReasonExplicit)
					}

					latency := time.Since(start)
					monitor.RecordOperation(latency)
				}
			}
		}(i)
	}

	// Run for specified duration
	time.Sleep(config.Duration)
	close(done)

	// Stop monitoring
	monitor.Stop()

	// Get results
	results := monitor.GetResults()

	return results, nil
}

// PrintBenchmarkResults prints benchmark results
func PrintBenchmarkResults(results *BenchmarkResults) {
	fmt.Printf("=== Performance Test Results ===\n")
	fmt.Printf("Duration: %v\n", results.Duration)
	fmt.Printf("Operations: %d\n", results.Operations)
	fmt.Printf("Throughput: %.2f ops/sec\n", results.Throughput)
	fmt.Printf("Average Latency: %v\n", results.Latency)
	fmt.Printf("Memory Usage: %d bytes\n", results.MemoryUsage)
	fmt.Printf("CPU Usage: %.2f%%\n", results.CPUUsage)
	fmt.Printf("Hit Rate: %.2f%%\n", results.HitRate)
	fmt.Printf("Miss Rate: %.2f%%\n", results.MissRate)
	fmt.Printf("Goroutine Count: %d\n", results.GoroutineCount)
	fmt.Printf("GC Cycles: %d\n", results.GCStats.NumGC)
	fmt.Printf("GC Pause Total: %v\n", time.Duration(results.GCStats.PauseTotalNs))
	fmt.Printf("===============================\n")
}
