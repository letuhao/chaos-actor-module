package cache

import (
	"fmt"
	"sync"
	"testing"
)

func TestMemoryOptimizer_New(t *testing.T) {
	config := &MemoryOptimizationConfig{
		EnableCompression:   true,
		EnableDeduplication: true,
		EnableCompaction:    true,
		MaxMemoryUsage:      1024 * 1024 * 1024, // 1GB
	}

	optimizer := NewMemoryOptimizer(config)
	if optimizer == nil {
		t.Fatal("Expected non-nil optimizer")
	}

	if !optimizer.config.EnableCompression {
		t.Error("Expected compression to be enabled")
	}
	if !optimizer.config.EnableDeduplication {
		t.Error("Expected deduplication to be enabled")
	}
	if !optimizer.config.EnableCompaction {
		t.Error("Expected compaction to be enabled")
	}
}

func TestMemoryOptimizer_DefaultConfig(t *testing.T) {
	optimizer := NewMemoryOptimizer(nil)
	if optimizer == nil {
		t.Fatal("Expected non-nil optimizer")
	}

	config := optimizer.config
	if !config.EnableCompression {
		t.Error("Expected compression to be enabled by default")
	}
	if !config.EnableDeduplication {
		t.Error("Expected deduplication to be enabled by default")
	}
	if !config.EnableCompaction {
		t.Error("Expected compaction to be enabled by default")
	}
	if config.MaxMemoryUsage <= 0 {
		t.Error("Expected positive max memory usage")
	}
}

func TestMemoryOptimizer_GetMemoryStats(t *testing.T) {
	optimizer := NewMemoryOptimizer(nil)

	stats := optimizer.GetMemoryStats()
	if stats == nil {
		t.Fatal("Expected non-nil memory stats")
	}

	// Check that stats have been initialized
	if stats.Timestamp.IsZero() {
		t.Error("Expected timestamp to be set")
	}
}

func TestMemoryOptimizer_GetOptimizationReport(t *testing.T) {
	optimizer := NewMemoryOptimizer(nil)

	report := optimizer.GetOptimizationReport()
	if report == nil {
		t.Fatal("Expected non-nil optimization report")
	}

	// Check required fields
	requiredFields := []string{
		"memory_stats", "compression", "deduplication", "compaction",
		"gc_pressure", "memory_pools", "zero_copy", "memory_reuse",
	}

	for _, field := range requiredFields {
		if _, exists := report[field]; !exists {
			t.Errorf("Missing required field: %s", field)
		}
	}

	// Check compression field
	compression, ok := report["compression"].(map[string]interface{})
	if !ok {
		t.Error("Expected compression to be a map")
	} else {
		if _, exists := compression["enabled"]; !exists {
			t.Error("Missing compression enabled field")
		}
	}
}

func TestMemoryOptimizer_CompressData(t *testing.T) {
	optimizer := NewMemoryOptimizer(nil)

	data := []byte("test data for compression")
	key := "test_key"

	// Test compression
	compressed, err := optimizer.CompressData(key, data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if compressed == nil {
		t.Error("Expected non-nil compressed data")
	}

	// Test decompression
	decompressed, err := optimizer.DecompressData(key, compressed)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(decompressed) != len(data) {
		t.Errorf("Expected decompressed length %d, got %d", len(data), len(decompressed))
	}
}

func TestMemoryOptimizer_DeduplicateData(t *testing.T) {
	optimizer := NewMemoryOptimizer(nil)

	data := []byte("test data for deduplication")

	// Test deduplication
	deduplicated1, err := optimizer.DeduplicateData(data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(deduplicated1) != len(data) {
		t.Errorf("Expected deduplicated length %d, got %d", len(data), len(deduplicated1))
	}

	// Test deduplication of same data
	deduplicated2, err := optimizer.DeduplicateData(data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(deduplicated2) != len(data) {
		t.Errorf("Expected deduplicated length %d, got %d", len(data), len(deduplicated2))
	}
}

func TestMemoryOptimizer_GetPooledObject(t *testing.T) {
	optimizer := NewMemoryOptimizer(nil)

	// Test getting pooled object
	obj := optimizer.GetPooledObject("pool_1024", 1024)
	if obj == nil {
		t.Error("Expected non-nil pooled object")
	}

	// Test putting pooled object back
	optimizer.PutPooledObject("pool_1024", obj)
}

func TestMemoryOptimizer_ZeroCopyOperation(t *testing.T) {
	optimizer := NewMemoryOptimizer(nil)

	// Test zero-copy operation
	executed := false
	optimizer.ZeroCopyOperation(func() {
		executed = true
	})

	if !executed {
		t.Error("Expected operation to be executed")
	}
}

func TestMemoryOptimizer_ReuseMemory(t *testing.T) {
	optimizer := NewMemoryOptimizer(nil)

	data := []byte("test data for memory reuse")

	// Test memory reuse
	reused := optimizer.ReuseMemory(data)
	if len(reused) != len(data) {
		t.Errorf("Expected reused length %d, got %d", len(data), len(reused))
	}
}

func TestMemoryOptimizer_OptimizeMemory(t *testing.T) {
	optimizer := NewMemoryOptimizer(nil)

	// Test memory optimization
	err := optimizer.OptimizeMemory()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestCompressionManager_Compress(t *testing.T) {
	manager := &CompressionManager{
		enabled:    true,
		level:      6,
		compressed: make(map[string][]byte),
		original:   make(map[string][]byte),
	}

	data := []byte("test data for compression")
	key := "test_key"

	// Test compression
	compressed, err := manager.Compress(key, data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if compressed == nil {
		t.Error("Expected non-nil compressed data")
	}

	// Test decompression
	decompressed, err := manager.Decompress(key, compressed)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(decompressed) != len(data) {
		t.Errorf("Expected decompressed length %d, got %d", len(data), len(decompressed))
	}
}

func TestDeduplicationManager_Deduplicate(t *testing.T) {
	manager := &DeduplicationManager{
		enabled:      true,
		threshold:    10,
		deduplicated: make(map[string][]byte),
		references:   make(map[string]int),
	}

	data := []byte("test data for deduplication")

	// Test deduplication
	deduplicated1, err := manager.Deduplicate(data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(deduplicated1) != len(data) {
		t.Errorf("Expected deduplicated length %d, got %d", len(data), len(deduplicated1))
	}

	// Test deduplication of same data
	deduplicated2, err := manager.Deduplicate(data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(deduplicated2) != len(data) {
		t.Errorf("Expected deduplicated length %d, got %d", len(data), len(deduplicated2))
	}
}

func TestCompactionManager_Optimize(t *testing.T) {
	manager := &CompactionManager{
		enabled:   true,
		threshold: 0.7,
	}

	stats := &MemoryStats{
		HeapInuse: 800,
		HeapSys:   1000,
	}

	// Test compaction
	err := manager.Optimize(stats)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if manager.compactionCount == 0 {
		t.Error("Expected compaction count to be greater than 0")
	}
}

func TestGCPressureManager_Optimize(t *testing.T) {
	manager := &GCPressureManager{
		enabled:   true,
		threshold: 0.8,
	}

	stats := &MemoryStats{
		HeapInuse: 900,
		HeapSys:   1000,
	}

	// Test GC pressure optimization
	err := manager.Optimize(stats)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if manager.gcCount == 0 {
		t.Error("Expected GC count to be greater than 0")
	}
}

func TestMemoryPoolManager_GetAndPut(t *testing.T) {
	manager := &MemoryPoolManager{
		enabled:   true,
		pools:     make(map[string]*sync.Pool),
		poolSizes: make(map[string]int),
	}

	// Initialize a pool
	poolName := "test_pool"
	manager.pools[poolName] = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}
	manager.poolSizes[poolName] = 1024

	// Test getting pooled object
	obj := manager.Get(poolName, 1024)
	if obj == nil {
		t.Error("Expected non-nil pooled object")
	}

	// Test putting pooled object back
	manager.Put(poolName, obj)
}

func TestZeroCopyManager_RecordOperation(t *testing.T) {
	manager := &ZeroCopyManager{
		enabled: true,
	}

	// Test recording operation
	executed := false
	manager.RecordOperation(func() {
		executed = true
	})

	if !executed {
		t.Error("Expected operation to be executed")
	}

	if manager.totalOps == 0 {
		t.Error("Expected total operations to be greater than 0")
	}

	if manager.zeroCopyOps == 0 {
		t.Error("Expected zero-copy operations to be greater than 0")
	}
}

func TestMemoryReuseManager_Reuse(t *testing.T) {
	manager := &MemoryReuseManager{
		enabled:   true,
		threshold: 0.5,
	}

	data := []byte("test data for memory reuse")

	// Test memory reuse
	reused := manager.Reuse(data)
	if len(reused) != len(data) {
		t.Errorf("Expected reused length %d, got %d", len(data), len(reused))
	}

	if manager.totalBytes == 0 {
		t.Error("Expected total bytes to be greater than 0")
	}

	if manager.reusedBytes == 0 {
		t.Error("Expected reused bytes to be greater than 0")
	}
}

func TestMemoryOptimizer_Integration(t *testing.T) {
	config := &MemoryOptimizationConfig{
		EnableCompression:   true,
		EnableDeduplication: true,
		EnableCompaction:    true,
		EnableGCPressure:    true,
		EnableMemoryPooling: true,
		EnableZeroCopy:      true,
		EnableMemoryReuse:   true,
		MaxMemoryUsage:      1024 * 1024 * 1024, // 1GB
	}

	optimizer := NewMemoryOptimizer(config)

	// Test compression
	data := []byte("test data for integration test")
	key := "integration_test"

	_, err := optimizer.CompressData(key, data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test deduplication
	_, err = optimizer.DeduplicateData(data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test memory pooling
	obj := optimizer.GetPooledObject("pool_1024", 1024)
	if obj != nil {
		optimizer.PutPooledObject("pool_1024", obj)
	}

	// Test zero-copy operation
	executed := false
	optimizer.ZeroCopyOperation(func() {
		executed = true
	})

	if !executed {
		t.Error("Expected operation to be executed")
	}

	// Test memory reuse
	reused := optimizer.ReuseMemory(data)
	if len(reused) != len(data) {
		t.Errorf("Expected reused length %d, got %d", len(data), len(reused))
	}

	// Test memory optimization
	err = optimizer.OptimizeMemory()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test getting optimization report
	report := optimizer.GetOptimizationReport()
	if report == nil {
		t.Fatal("Expected non-nil optimization report")
	}
}

func BenchmarkMemoryOptimizer_CompressData(b *testing.B) {
	optimizer := NewMemoryOptimizer(nil)
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)
		optimizer.CompressData(key, data)
	}
}

func BenchmarkMemoryOptimizer_DeduplicateData(b *testing.B) {
	optimizer := NewMemoryOptimizer(nil)
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		optimizer.DeduplicateData(data)
	}
}

func BenchmarkMemoryOptimizer_GetPooledObject(b *testing.B) {
	optimizer := NewMemoryOptimizer(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj := optimizer.GetPooledObject("pool_1024", 1024)
		if obj != nil {
			optimizer.PutPooledObject("pool_1024", obj)
		}
	}
}

func BenchmarkMemoryOptimizer_ZeroCopyOperation(b *testing.B) {
	optimizer := NewMemoryOptimizer(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		optimizer.ZeroCopyOperation(func() {
			// Simple operation
			_ = i * 2
		})
	}
}

func BenchmarkMemoryOptimizer_ReuseMemory(b *testing.B) {
	optimizer := NewMemoryOptimizer(nil)
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		optimizer.ReuseMemory(data)
	}
}
