package cache

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"
)

func TestCacheWarmer_New(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	// Test with default config
	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)
	if warmer == nil {
		t.Fatal("Expected warmer to be created")
	}
	defer warmer.Close()

	if warmer.l1Cache != l1Cache {
		t.Fatal("Expected L1 cache to be set")
	}

	if warmer.l2Cache != l2Cache {
		t.Fatal("Expected L2 cache to be set")
	}

	if warmer.l3Cache != l3Cache {
		t.Fatal("Expected L3 cache to be set")
	}
}

func TestCacheWarmer_NewWithConfig(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	// Test with custom config
	config := &CacheWarmerConfig{
		EnableWarming:       true,
		WarmingInterval:     time.Minute * 2,
		MaxWarmingWorkers:   2,
		EnablePrediction:    true,
		PredictionWindow:    time.Minute * 30,
		MinAccessCount:      3,
		ConfidenceThreshold: 0.6,
		PreloadBatchSize:    50,
		PreloadTimeout:      time.Second * 15,
		MaxPreloadSize:      512 * 1024,
		WarmingPriority:     2,
		EnableHotPath:       true,
		HotPathThreshold:    0.7,
	}

	warmer := NewCacheWarmer(config, l1Cache, l2Cache, l3Cache)
	if warmer == nil {
		t.Fatal("Expected warmer to be created")
	}
	defer warmer.Close()

	if warmer.config != config {
		t.Fatal("Expected config to be set")
	}
}

func TestCacheWarmer_RecordAccess(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Record access
	warmer.RecordAccess("key1")

	// Check if pattern was recorded
	patterns := warmer.GetAccessPatterns()
	pattern, exists := patterns["key1"]
	if !exists {
		t.Fatal("Expected pattern to be recorded")
	}

	if pattern.Key != "key1" {
		t.Errorf("Expected key to be 'key1', got %s", pattern.Key)
	}

	if pattern.AccessCount != 1 {
		t.Errorf("Expected access count to be 1, got %d", pattern.AccessCount)
	}

	if pattern.FirstAccess.IsZero() {
		t.Fatal("Expected first access time to be set")
	}

	if pattern.LastAccess.IsZero() {
		t.Fatal("Expected last access time to be set")
	}
}

func TestCacheWarmer_RecordMultipleAccesses(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Record multiple accesses
	for i := 0; i < 5; i++ {
		warmer.RecordAccess("key1")
		time.Sleep(time.Millisecond * 10) // Small delay to ensure different timestamps
	}

	// Check pattern
	patterns := warmer.GetAccessPatterns()
	pattern, exists := patterns["key1"]
	if !exists {
		t.Fatal("Expected pattern to be recorded")
	}

	if pattern.AccessCount != 5 {
		t.Errorf("Expected access count to be 5, got %d", pattern.AccessCount)
	}

	if len(pattern.AccessTimes) != 5 {
		t.Errorf("Expected 5 access times, got %d", len(pattern.AccessTimes))
	}

	if len(pattern.AccessIntervals) != 4 {
		t.Errorf("Expected 4 access intervals, got %d", len(pattern.AccessIntervals))
	}
}

func TestCacheWarmer_PredictNextAccess(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Record multiple accesses with regular intervals
	for i := 0; i < 10; i++ {
		warmer.RecordAccess("key1")
		time.Sleep(time.Millisecond * 100)
	}

	// Predict next access
	nextAccess, confidence := warmer.PredictNextAccess("key1")
	if nextAccess.IsZero() {
		t.Fatal("Expected next access time to be predicted")
	}

	if confidence <= 0 {
		t.Errorf("Expected confidence to be > 0, got %f", confidence)
	}

	// Check that prediction is in the future
	if nextAccess.Before(time.Now()) {
		t.Fatal("Expected prediction to be in the future")
	}
}

func TestCacheWarmer_PredictNextAccessLowConfidence(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Record only one access (below minimum threshold)
	warmer.RecordAccess("key1")

	// Predict next access
	nextAccess, confidence := warmer.PredictNextAccess("key1")
	if !nextAccess.IsZero() {
		t.Fatal("Expected no prediction for low confidence")
	}

	if confidence > 0 {
		t.Errorf("Expected confidence to be 0, got %f", confidence)
	}
}

func TestCacheWarmer_GetWarmingCandidates(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Record accesses for multiple keys with different patterns
	// Key1: High frequency
	for i := 0; i < 10; i++ {
		warmer.RecordAccess("key1")
		time.Sleep(time.Millisecond * 10)
	}

	// Key2: Medium frequency
	for i := 0; i < 5; i++ {
		warmer.RecordAccess("key2")
		time.Sleep(time.Millisecond * 20)
	}

	// Key3: Low frequency
	for i := 0; i < 2; i++ {
		warmer.RecordAccess("key3")
		time.Sleep(time.Millisecond * 50)
	}

	// Get warming candidates
	candidates := warmer.GetWarmingCandidates(2)
	if len(candidates) != 2 {
		t.Errorf("Expected 2 candidates, got %d", len(candidates))
	}

	// Key1 should be first (highest frequency)
	if candidates[0] != "key1" {
		t.Errorf("Expected key1 to be first candidate, got %s", candidates[0])
	}

	// Key2 should be second
	if candidates[1] != "key2" {
		t.Errorf("Expected key2 to be second candidate, got %s", candidates[1])
	}
}

func TestCacheWarmer_WarmCache(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Add some data to L3 cache
	l3Cache.Set("key1", "value1", time.Hour)
	l3Cache.Set("key2", "value2", time.Hour)

	// Record accesses
	warmer.RecordAccess("key1")
	warmer.RecordAccess("key2")

	// Warm cache
	candidates := []string{"key1", "key2"}
	err = warmer.WarmCache(candidates)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Wait a bit for warming to complete
	time.Sleep(time.Millisecond * 100)

	// Check if data was promoted to L1 and L2
	if !l1Cache.Has("key1") {
		t.Fatal("Expected key1 to be in L1 cache after warming")
	}

	if !l1Cache.Has("key2") {
		t.Fatal("Expected key2 to be in L1 cache after warming")
	}

	if !l2Cache.Has("key1") {
		t.Fatal("Expected key1 to be in L2 cache after warming")
	}

	if !l2Cache.Has("key2") {
		t.Fatal("Expected key2 to be in L2 cache after warming")
	}
}

func TestCacheWarmer_GetStats(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Get initial stats
	stats := warmer.GetStats()
	if stats == nil {
		t.Fatal("Expected stats to be returned")
	}

	if stats.ActiveWorkers != 4 {
		t.Errorf("Expected 4 active workers, got %d", stats.ActiveWorkers)
	}

	// Record some accesses
	warmer.RecordAccess("key1")
	warmer.RecordAccess("key2")

	// Get updated stats
	stats = warmer.GetStats()
	if stats == nil {
		t.Fatal("Expected stats to be returned")
	}
}

func TestCacheWarmer_ClearPatterns(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Record some accesses
	warmer.RecordAccess("key1")
	warmer.RecordAccess("key2")

	// Check patterns exist
	patterns := warmer.GetAccessPatterns()
	if len(patterns) != 2 {
		t.Errorf("Expected 2 patterns, got %d", len(patterns))
	}

	// Clear patterns
	warmer.ClearPatterns()

	// Check patterns are cleared
	patterns = warmer.GetAccessPatterns()
	if len(patterns) != 0 {
		t.Errorf("Expected 0 patterns after clear, got %d", len(patterns))
	}
}

func TestCacheWarmer_SetConfig(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Test setting valid config
	newConfig := &CacheWarmerConfig{
		EnableWarming:       false,
		WarmingInterval:     time.Minute * 3,
		MaxWarmingWorkers:   6,
		EnablePrediction:    false,
		PredictionWindow:    time.Hour,
		MinAccessCount:      10,
		ConfidenceThreshold: 0.8,
		PreloadBatchSize:    200,
		PreloadTimeout:      time.Minute,
		MaxPreloadSize:      2 * 1024 * 1024,
		WarmingPriority:     3,
		EnableHotPath:       false,
		HotPathThreshold:    0.9,
	}

	err = warmer.SetConfig(newConfig)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify config was updated
	config := warmer.GetConfig()
	if config.EnableWarming != newConfig.EnableWarming {
		t.Errorf("Expected EnableWarming to be %v, got %v", newConfig.EnableWarming, config.EnableWarming)
	}

	if config.MaxWarmingWorkers != newConfig.MaxWarmingWorkers {
		t.Errorf("Expected MaxWarmingWorkers to be %d, got %d", newConfig.MaxWarmingWorkers, config.MaxWarmingWorkers)
	}

	// Test setting nil config
	err = warmer.SetConfig(nil)
	if err == nil {
		t.Fatal("Expected error for nil config")
	}
}

func TestCacheWarmer_Close(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)

	// Test close
	err = warmer.Close()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test operations after close
	warmer.RecordAccess("key1")
	// Should not panic or error
}

func TestCacheWarmer_ConcurrentAccess(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	// Create cache layers
	l1Cache := NewLockFreeL1Cache(100, "allkeys-lru")
	l2Cache, err := NewMemoryMappedL2Cache(cachePath, 1000)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l2Cache.Close()

	l3Cache, err := NewPersistentL3Cache(l3Dir, 10000, true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer l3Cache.Close()

	warmer := NewCacheWarmer(nil, l1Cache, l2Cache, l3Cache)
	defer warmer.Close()

	// Test concurrent access recording
	done := make(chan bool)

	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key%d_%d", id, j)
				warmer.RecordAccess(key)
			}
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Check final state
	patterns := warmer.GetAccessPatterns()
	if len(patterns) != 1000 {
		t.Errorf("Expected 1000 patterns, got %d", len(patterns))
	}
}
