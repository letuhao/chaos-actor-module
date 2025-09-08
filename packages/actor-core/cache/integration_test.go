package cache

import (
	"fmt"
	"testing"
	"time"
)

// TestComprehensiveIntegration tests the integration of all optimization systems
func TestComprehensiveIntegration(t *testing.T) {
	// Create all optimization systems
	simdConfig := DefaultSIMDConfig()
	simdOptimizer := NewSIMDOptimizer(simdConfig)

	monitoringConfig := DefaultMonitoringConfig()
	monitor := NewAdvancedMonitor(monitoringConfig)
	defer monitor.Close()

	memoryConfig := DefaultMemoryOptimizationConfig()
	memoryOptimizer := NewMemoryOptimizer(memoryConfig)

	networkConfig := DefaultNetworkOptimizationConfig()
	networkOptimizer := NewNetworkOptimizer(networkConfig)

	// Test data
	testData := make([][]byte, 1000)
	for i := range testData {
		testData[i] = make([]byte, 1024)
		for j := range testData[i] {
			testData[i][j] = byte((i + j) % 256)
		}
	}

	// Test SIMD optimizations
	t.Run("SIMD_Optimizations", func(t *testing.T) {
		// Test batch hashing
		hashes := simdOptimizer.BatchHash(testData[:100])
		if len(hashes) != 100 {
			t.Errorf("Expected 100 hashes, got %d", len(hashes))
		}

		// Test fast comparison
		result := simdOptimizer.FastCompare(testData[0], testData[1])
		if result == 0 {
			t.Error("Expected different data to compare as different")
		}

		// Test fast search
		pattern := []byte{0, 1, 2, 3}
		index := simdOptimizer.FastSearch(testData[0], pattern)
		if index == -1 {
			t.Log("Pattern not found - this is expected for random data")
		}

		// Test fast sorting
		sortData := []uint64{5, 2, 8, 1, 9, 3, 7, 4, 6}
		simdOptimizer.FastSort(sortData)
		for i := 1; i < len(sortData); i++ {
			if sortData[i-1] > sortData[i] {
				t.Error("Data not sorted correctly")
			}
		}
	})

	// Test advanced monitoring
	t.Run("Advanced_Monitoring", func(t *testing.T) {
		// Record metrics
		monitor.RecordMetric("test_metric", MetricTypeGauge, 42.0, map[string]string{"test": "value"})
		monitor.RecordMetric("cpu_usage", MetricTypeGauge, 75.0, map[string]string{"type": "system"})

		// Record alert
		monitor.RecordAlert(AlertLevelWarning, "Test alert", "test_metric", 85.0, 80.0, map[string]string{"test": "value"})

		// Record trace
		span := TraceSpan{
			TraceID:   "test_trace",
			SpanID:    "test_span",
			Operation: "test_operation",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Millisecond * 100),
			Duration:  time.Millisecond * 100,
			Tags:      map[string]string{"test": "value"},
		}
		monitor.RecordTrace(span)

		// Get dashboard data
		dashboardData := monitor.GetDashboardData()
		if dashboardData == nil {
			t.Fatal("Expected non-nil dashboard data")
		}

		// Check required fields
		requiredFields := []string{"timestamp", "metrics", "alerts", "profiling", "statistics", "system_info", "cache_performance"}
		for _, field := range requiredFields {
			if _, exists := dashboardData[field]; !exists {
				t.Errorf("Missing required field: %s", field)
			}
		}
	})

	// Test memory optimization
	t.Run("Memory_Optimization", func(t *testing.T) {
		// Test compression
		compressed, err := memoryOptimizer.CompressData("test_key", testData[0])
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if compressed == nil {
			t.Error("Expected non-nil compressed data")
		}

		// Test deduplication
		deduplicated, err := memoryOptimizer.DeduplicateData(testData[0])
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(deduplicated) != len(testData[0]) {
			t.Errorf("Expected deduplicated length %d, got %d", len(testData[0]), len(deduplicated))
		}

		// Test memory pooling
		obj := memoryOptimizer.GetPooledObject("pool_1024", 1024)
		if obj != nil {
			memoryOptimizer.PutPooledObject("pool_1024", obj)
		}

		// Test zero-copy operation
		executed := false
		memoryOptimizer.ZeroCopyOperation(func() {
			executed = true
		})
		if !executed {
			t.Error("Expected operation to be executed")
		}

		// Test memory reuse
		reused := memoryOptimizer.ReuseMemory(testData[0])
		if len(reused) != len(testData[0]) {
			t.Errorf("Expected reused length %d, got %d", len(testData[0]), len(reused))
		}

		// Test memory optimization
		err = memoryOptimizer.OptimizeMemory()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Get optimization report
		report := memoryOptimizer.GetOptimizationReport()
		if report == nil {
			t.Fatal("Expected non-nil optimization report")
		}
	})

	// Test network optimization
	t.Run("Network_Optimization", func(t *testing.T) {
		// Test compression
		compressed, err := networkOptimizer.CompressData(testData[0])
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if compressed == nil {
			t.Error("Expected non-nil compressed data")
		}

		// Test encryption
		encrypted, err := networkOptimizer.EncryptData(testData[0])
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if encrypted == nil {
			t.Error("Expected non-nil encrypted data")
		}

		// Test batching
		message := &NetworkMessage{
			ID:       "test_message",
			Data:     testData[0],
			Priority: 1,
			Timeout:  time.Second * 5,
		}
		err = networkOptimizer.BatchMessage(message)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Test pipelining
		request := &NetworkRequest{
			ID:        "test_request",
			Data:      testData[0],
			Response:  make(chan *NetworkResponse, 1),
			Timestamp: time.Now(),
			Timeout:   time.Second * 5,
		}
		err = networkOptimizer.PipelineRequest(request)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Get optimization report
		report := networkOptimizer.GetOptimizationReport()
		if report == nil {
			t.Fatal("Expected non-nil optimization report")
		}
	})

	// Test combined performance
	t.Run("Combined_Performance", func(t *testing.T) {
		start := time.Now()

		// Perform combined operations
		for i := 0; i < 100; i++ {
			// SIMD operations
			simdOptimizer.FastHash(testData[i%len(testData)])
			simdOptimizer.FastCompare(testData[i%len(testData)], testData[(i+1)%len(testData)])

			// Memory operations
			memoryOptimizer.CompressData(fmt.Sprintf("key_%d", i), testData[i%len(testData)])
			memoryOptimizer.DeduplicateData(testData[i%len(testData)])

			// Network operations
			networkOptimizer.CompressData(testData[i%len(testData)])
			networkOptimizer.EncryptData(testData[i%len(testData)])

			// Monitoring
			monitor.RecordMetric("performance_test", MetricTypeGauge, float64(i), nil)
		}

		duration := time.Since(start)
		t.Logf("Combined operations completed in %v", duration)

		// Verify performance is reasonable (less than 1 second for 100 operations)
		if duration > time.Second {
			t.Errorf("Performance too slow: %v", duration)
		}
	})
}

// TestCacheSystemIntegration tests the integration with the cache system
func TestCacheSystemIntegration(t *testing.T) {
	// Create multi-layer cache
	config := &MultiLayerConfig{
		L1MaxSize:        1000,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      "test_integration_l2.dat",
		L2MaxSize:        10000,
		L3CacheDir:       "test_integration_l3",
		L3MaxSize:        100000,
		L3Compression:    true,
		EnablePreloading: true,
		PreloadWorkers:   2,
		SyncInterval:     time.Second,
	}

	multiCache, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Failed to create multi-layer cache: %v", err)
	}
	defer multiCache.Close()

	// Create optimization systems
	simdOptimizer := NewSIMDOptimizer(nil)
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()
	memoryOptimizer := NewMemoryOptimizer(nil)
	_ = NewNetworkOptimizer(nil) // Network optimizer for future use

	// Test data
	testData := make([]byte, 1024)
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// Test cache operations with optimizations
	t.Run("Cache_Operations", func(t *testing.T) {
		// Set data with compression
		compressed, err := memoryOptimizer.CompressData("test_key", testData)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		multiCache.Set("test_key", compressed, time.Hour)

		// Get data with decompression
		value, exists := multiCache.Get("test_key")
		if !exists {
			t.Fatal("Expected key to exist")
		}

		compressedValue, ok := value.([]byte)
		if !ok {
			t.Fatal("Expected value to be []byte")
		}

		decompressed, err := memoryOptimizer.DecompressData("test_key", compressedValue)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(decompressed) != len(testData) {
			t.Errorf("Expected decompressed length %d, got %d", len(testData), len(decompressed))
		}
	})

	// Test performance monitoring
	t.Run("Performance_Monitoring", func(t *testing.T) {
		// Record cache operations
		monitor.RecordMetric("cache_hit_rate", MetricTypeGauge, 0.85, map[string]string{"cache": "l1"})
		monitor.RecordMetric("cache_miss_rate", MetricTypeGauge, 0.15, map[string]string{"cache": "l1"})
		monitor.RecordMetric("cache_throughput", MetricTypeCounter, 1000.0, map[string]string{"cache": "multi_layer"})

		// Record performance trace
		span := TraceSpan{
			TraceID:   "cache_operation",
			SpanID:    "cache_set",
			Operation: "cache_set",
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Millisecond * 50),
			Duration:  time.Millisecond * 50,
			Tags:      map[string]string{"operation": "set", "key": "test_key"},
		}
		monitor.RecordTrace(span)

		// Get performance report
		dashboardData := monitor.GetDashboardData()
		if dashboardData == nil {
			t.Fatal("Expected non-nil dashboard data")
		}

		// Check cache performance data
		cachePerformance, ok := dashboardData["cache_performance"].(map[string]interface{})
		if !ok {
			t.Error("Expected cache_performance to be a map")
		} else {
			requiredFields := []string{"l1_hit_rate", "l2_hit_rate", "l3_hit_rate", "total_hit_rate", "throughput", "latency_p50", "latency_p90", "latency_p99"}
			for _, field := range requiredFields {
				if _, exists := cachePerformance[field]; !exists {
					t.Errorf("Missing required cache performance field: %s", field)
				}
			}
		}
	})

	// Test memory optimization with cache
	t.Run("Memory_Optimization_With_Cache", func(t *testing.T) {
		// Test memory optimization
		err := memoryOptimizer.OptimizeMemory()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Get memory stats
		stats := memoryOptimizer.GetMemoryStats()
		if stats == nil {
			t.Fatal("Expected non-nil memory stats")
		}

		// Verify memory stats are reasonable
		if stats.HeapAlloc == 0 {
			t.Error("Expected non-zero heap allocation")
		}
		if stats.HeapSys == 0 {
			t.Error("Expected non-zero heap system memory")
		}
	})

	// Test SIMD optimizations with cache
	t.Run("SIMD_Optimizations_With_Cache", func(t *testing.T) {
		// Test batch hashing of cache keys
		keys := make([][]byte, 100)
		for i := range keys {
			keys[i] = []byte(fmt.Sprintf("key_%d", i))
		}

		hashes := simdOptimizer.BatchHash(keys)
		if len(hashes) != len(keys) {
			t.Errorf("Expected %d hashes, got %d", len(keys), len(hashes))
		}

		// Test fast comparison of cache values
		value1 := []byte("test_value_1")
		value2 := []byte("test_value_2")

		result := simdOptimizer.FastCompare(value1, value2)
		if result == 0 {
			t.Error("Expected different values to compare as different")
		}

		// Test fast search in cache data
		pattern := []byte("test")
		index := simdOptimizer.FastSearch(value1, pattern)
		if index == -1 {
			t.Error("Expected pattern to be found")
		}
	})
}

// BenchmarkComprehensiveIntegration benchmarks the comprehensive integration
func BenchmarkComprehensiveIntegration(b *testing.B) {
	// Create all optimization systems
	simdOptimizer := NewSIMDOptimizer(nil)
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()
	memoryOptimizer := NewMemoryOptimizer(nil)
	networkOptimizer := NewNetworkOptimizer(nil)

	// Test data
	testData := make([]byte, 1024)
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// SIMD operations
		simdOptimizer.FastHash(testData)
		simdOptimizer.FastCompare(testData, testData)

		// Memory operations
		memoryOptimizer.CompressData(fmt.Sprintf("key_%d", i), testData)
		memoryOptimizer.DeduplicateData(testData)

		// Network operations
		networkOptimizer.CompressData(testData)
		networkOptimizer.EncryptData(testData)

		// Monitoring
		monitor.RecordMetric("benchmark_test", MetricTypeGauge, float64(i), nil)
	}
}

// BenchmarkCacheSystemIntegration benchmarks the cache system integration
func BenchmarkCacheSystemIntegration(b *testing.B) {
	// Create multi-layer cache
	config := &MultiLayerConfig{
		L1MaxSize:        1000,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      "benchmark_l2.dat",
		L2MaxSize:        10000,
		L3CacheDir:       "benchmark_l3",
		L3MaxSize:        100000,
		L3Compression:    true,
		EnablePreloading: false, // Disable for benchmark
		SyncInterval:     time.Second,
	}

	multiCache, err := NewMultiLayerCacheManager(config)
	if err != nil {
		b.Fatalf("Failed to create multi-layer cache: %v", err)
	}
	defer multiCache.Close()

	// Create optimization systems
	memoryOptimizer := NewMemoryOptimizer(nil)

	// Test data
	testData := make([]byte, 1024)
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)

		// Compress data
		compressed, err := memoryOptimizer.CompressData(key, testData)
		if err != nil {
			b.Fatalf("Expected no error, got %v", err)
		}

		// Set in cache
		multiCache.Set(key, compressed, time.Hour)

		// Get from cache
		value, exists := multiCache.Get(key)
		if !exists {
			b.Fatal("Expected key to exist")
		}

		// Decompress data
		compressedValue, ok := value.([]byte)
		if !ok {
			b.Fatal("Expected value to be []byte")
		}

		_, err = memoryOptimizer.DecompressData(key, compressedValue)
		if err != nil {
			b.Fatalf("Expected no error, got %v", err)
		}
	}
}
