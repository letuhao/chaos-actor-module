package cache

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"
)

func TestMultiLayerCacheManager_New(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	if manager == nil {
		t.Fatal("Expected manager to be created")
	}

	if manager.l1Cache == nil {
		t.Fatal("Expected L1 cache to be initialized")
	}

	if manager.l2Cache == nil {
		t.Fatal("Expected L2 cache to be initialized")
	}

	if manager.l3Cache == nil {
		t.Fatal("Expected L3 cache to be initialized")
	}
}

func TestMultiLayerCacheManager_NewWithNilConfig(t *testing.T) {
	manager, err := NewMultiLayerCacheManager(nil)
	if err == nil {
		t.Fatal("Expected error for nil config")
	}

	if manager != nil {
		t.Fatal("Expected manager to be nil")
	}
}

func TestMultiLayerCacheManager_SetAndGet(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Test Set
	err = manager.Set("key1", "value1", time.Hour)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test Get
	value, found := manager.Get("key1")
	if !found {
		t.Fatal("Expected key1 to be found")
	}

	if value != "value1" {
		t.Errorf("Expected value to be 'value1', got %v", value)
	}
}

func TestMultiLayerCacheManager_GetNonExistent(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Test Get non-existent key
	value, found := manager.Get("nonexistent")
	if found {
		t.Fatal("Expected key to not be found")
	}

	if value != nil {
		t.Errorf("Expected value to be nil, got %v", value)
	}
}

func TestMultiLayerCacheManager_SetEmptyKey(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Test Set empty key
	err = manager.Set("", "value", time.Hour)
	if err != ErrEmptyKey {
		t.Errorf("Expected ErrEmptyKey, got %v", err)
	}
}

func TestMultiLayerCacheManager_SetNilValue(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Test Set nil value
	err = manager.Set("key", nil, time.Hour)
	if err != ErrNilValue {
		t.Errorf("Expected ErrNilValue, got %v", err)
	}
}

func TestMultiLayerCacheManager_Delete(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Set a value
	err = manager.Set("key1", "value1", time.Hour)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify it exists
	if !manager.Has("key1") {
		t.Fatal("Expected key1 to exist")
	}

	// Delete the value
	err = manager.Delete("key1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify it's deleted
	if manager.Has("key1") {
		t.Fatal("Expected key1 to be deleted")
	}
}

func TestMultiLayerCacheManager_DeleteEmptyKey(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Test Delete empty key
	err = manager.Delete("")
	if err != ErrEmptyKey {
		t.Errorf("Expected ErrEmptyKey, got %v", err)
	}
}

func TestMultiLayerCacheManager_Clear(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Set some values
	for i := 0; i < 10; i++ {
		err = manager.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i), time.Hour)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	}

	// Verify values exist
	for i := 0; i < 10; i++ {
		if !manager.Has(fmt.Sprintf("key%d", i)) {
			t.Fatalf("Expected key%d to exist", i)
		}
	}

	// Clear all caches
	err = manager.Clear()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify values are cleared
	for i := 0; i < 10; i++ {
		if manager.Has(fmt.Sprintf("key%d", i)) {
			t.Fatalf("Expected key%d to be cleared", i)
		}
	}
}

func TestMultiLayerCacheManager_Has(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Test Has with non-existent key
	if manager.Has("nonexistent") {
		t.Fatal("Expected key to not exist")
	}

	// Set a value
	err = manager.Set("key1", "value1", time.Hour)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test Has with existing key
	if !manager.Has("key1") {
		t.Fatal("Expected key1 to exist")
	}
}

func TestMultiLayerCacheManager_Stats(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Get initial stats
	stats := manager.GetStats()
	if stats == nil {
		t.Fatal("Expected stats to be returned")
	}

	// Perform some operations
	manager.Set("key1", "value1", time.Hour)
	manager.Get("key1")
	manager.Get("nonexistent")

	// Get updated stats
	stats = manager.GetStats()
	if stats.TotalSets != 1 {
		t.Errorf("Expected TotalSets to be 1, got %d", stats.TotalSets)
	}

	if stats.TotalGets != 2 {
		t.Errorf("Expected TotalGets to be 2, got %d", stats.TotalGets)
	}

	if stats.TotalHits != 1 {
		t.Errorf("Expected TotalHits to be 1, got %d", stats.TotalHits)
	}

	if stats.TotalMisses != 1 {
		t.Errorf("Expected TotalMisses to be 1, got %d", stats.TotalMisses)
	}
}

func TestMultiLayerCacheManager_HitRate(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Test hit rate with no operations
	hitRate := manager.GetHitRate()
	if hitRate != 0.0 {
		t.Errorf("Expected hit rate to be 0.0, got %f", hitRate)
	}

	// Perform operations
	manager.Set("key1", "value1", time.Hour)
	manager.Get("key1") // Hit
	manager.Get("key2") // Miss

	// Test hit rate
	hitRate = manager.GetHitRate()
	expectedHitRate := 50.0 // 1 hit out of 2 total operations
	if hitRate != expectedHitRate {
		t.Errorf("Expected hit rate to be %f, got %f", expectedHitRate, hitRate)
	}
}

func TestMultiLayerCacheManager_LayerHitRates(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Get layer hit rates
	rates := manager.GetLayerHitRates()
	if rates == nil {
		t.Fatal("Expected layer hit rates to be returned")
	}

	// Check that all layers are present
	if _, exists := rates["L1"]; !exists {
		t.Fatal("Expected L1 hit rate to be present")
	}

	if _, exists := rates["L2"]; !exists {
		t.Fatal("Expected L2 hit rate to be present")
	}

	if _, exists := rates["L3"]; !exists {
		t.Fatal("Expected L3 hit rate to be present")
	}
}

func TestMultiLayerCacheManager_MemoryUsage(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Get memory usage
	usage := manager.GetMemoryUsage()
	if usage == nil {
		t.Fatal("Expected memory usage to be returned")
	}

	// Check that all layers are present
	if _, exists := usage["L1"]; !exists {
		t.Fatal("Expected L1 memory usage to be present")
	}

	if _, exists := usage["L2"]; !exists {
		t.Fatal("Expected L2 memory usage to be present")
	}

	if _, exists := usage["L3"]; !exists {
		t.Fatal("Expected L3 memory usage to be present")
	}
}

func TestMultiLayerCacheManager_Sync(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Test sync
	err = manager.Sync()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check that sync count increased
	stats := manager.GetStats()
	if stats.SyncCount != 1 {
		t.Errorf("Expected SyncCount to be 1, got %d", stats.SyncCount)
	}
}

func TestMultiLayerCacheManager_Preload(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Prepare preload data
	data := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	// Test preload
	err = manager.Preload(data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify preloaded data
	for key, expectedValue := range data {
		value, found := manager.Get(key)
		if !found {
			t.Fatalf("Expected key %s to be found after preload", key)
		}
		if value != expectedValue {
			t.Errorf("Expected value for key %s to be %v, got %v", key, expectedValue, value)
		}
	}
}

func TestMultiLayerCacheManager_GetConfig(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Test GetConfig
	returnedConfig := manager.GetConfig()
	if returnedConfig == nil {
		t.Fatal("Expected config to be returned")
	}

	if returnedConfig.L1MaxSize != config.L1MaxSize {
		t.Errorf("Expected L1MaxSize to be %d, got %d", config.L1MaxSize, returnedConfig.L1MaxSize)
	}

	if returnedConfig.L2CachePath != config.L2CachePath {
		t.Errorf("Expected L2CachePath to be %s, got %s", config.L2CachePath, returnedConfig.L2CachePath)
	}

	if returnedConfig.L3CacheDir != config.L3CacheDir {
		t.Errorf("Expected L3CacheDir to be %s, got %s", config.L3CacheDir, returnedConfig.L3CacheDir)
	}
}

func TestMultiLayerCacheManager_SetConfig(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Test SetConfig with valid config
	newConfig := &MultiLayerConfig{
		L1MaxSize:        200,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        2000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        20000,
		L3Compression:    false,
		EnablePreloading: true,
		PreloadWorkers:   8,
		SyncInterval:     time.Second * 30,
	}

	err = manager.SetConfig(newConfig)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify config was updated
	returnedConfig := manager.GetConfig()
	if returnedConfig.L1MaxSize != newConfig.L1MaxSize {
		t.Errorf("Expected L1MaxSize to be %d, got %d", newConfig.L1MaxSize, returnedConfig.L1MaxSize)
	}

	// Test SetConfig with nil config
	err = manager.SetConfig(nil)
	if err == nil {
		t.Fatal("Expected error for nil config")
	}
}

func TestMultiLayerCacheManager_Close(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test Close
	err = manager.Close()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test operations after close
	err = manager.Set("key1", "value1", time.Hour)
	if err == nil {
		t.Fatal("Expected error when setting after close")
	}

	_, found := manager.Get("key1")
	if found {
		t.Fatal("Expected key to not be found after close")
	}
}

func TestMultiLayerCacheManager_ConcurrentAccess(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")
	l3Dir := filepath.Join(tempDir, "l3_cache")

	config := &MultiLayerConfig{
		L1MaxSize:        100,
		L1EvictionPolicy: "allkeys-lru",
		L2CachePath:      cachePath,
		L2MaxSize:        1000,
		L3CacheDir:       l3Dir,
		L3MaxSize:        10000,
		L3Compression:    true,
		EnablePreloading: false,
		PreloadWorkers:   4,
		SyncInterval:     time.Minute,
	}

	manager, err := NewMultiLayerCacheManager(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer manager.Close()

	// Test concurrent access
	done := make(chan bool)

	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key%d_%d", id, j)
				value := fmt.Sprintf("value%d_%d", id, j)

				// Set value
				manager.Set(key, value, time.Hour)

				// Get value
				manager.Get(key)
			}
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify final state
	stats := manager.GetStats()
	if stats.TotalSets != 1000 {
		t.Errorf("Expected TotalSets to be 1000, got %d", stats.TotalSets)
	}

	if stats.TotalGets != 1000 {
		t.Errorf("Expected TotalGets to be 1000, got %d", stats.TotalGets)
	}
}
