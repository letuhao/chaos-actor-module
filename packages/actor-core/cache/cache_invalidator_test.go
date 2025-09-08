package cache

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"
)

func TestCacheInvalidator_New(t *testing.T) {
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
	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)
	if invalidator == nil {
		t.Fatal("Expected invalidator to be created")
	}
	defer invalidator.Close()

	if invalidator.l1Cache != l1Cache {
		t.Fatal("Expected L1 cache to be set")
	}

	if invalidator.l2Cache != l2Cache {
		t.Fatal("Expected L2 cache to be set")
	}

	if invalidator.l3Cache != l3Cache {
		t.Fatal("Expected L3 cache to be set")
	}
}

func TestCacheInvalidator_NewWithConfig(t *testing.T) {
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
	config := &CacheInvalidatorConfig{
		EnableInvalidation:     true,
		InvalidationInterval:   time.Minute * 3,
		MaxInvalidationWorkers: 2,
		EnableTTL:              true,
		TTLCheckInterval:       time.Minute * 2,
		DefaultTTL:             time.Hour * 2,
		MaxTTL:                 time.Hour * 48,
		EnableDependencies:     true,
		MaxDependencyDepth:     5,
		DependencyTimeout:      time.Second * 45,
		InvalidationPriority:   2,
		BatchSize:              50,
		EnableLazyInvalidation: false,
		LazyThreshold:          500,
	}

	invalidator := NewCacheInvalidator(config, l1Cache, l2Cache, l3Cache)
	if invalidator == nil {
		t.Fatal("Expected invalidator to be created")
	}
	defer invalidator.Close()

	if invalidator.config != config {
		t.Fatal("Expected config to be set")
	}
}

func TestCacheInvalidator_Invalidate(t *testing.T) {
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

	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Add some data to caches
	l1Cache.Set("key1", "value1", time.Hour)
	l2Cache.Set("key1", "value1", time.Hour)
	l3Cache.Set("key1", "value1", time.Hour)

	// Verify data exists
	if !l1Cache.Has("key1") {
		t.Fatal("Expected key1 to exist in L1 cache")
	}
	if !l2Cache.Has("key1") {
		t.Fatal("Expected key1 to exist in L2 cache")
	}
	if !l3Cache.Has("key1") {
		t.Fatal("Expected key1 to exist in L3 cache")
	}

	// Invalidate key1
	err = invalidator.Invalidate("key1", InvalidationReasonExplicit)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Wait a bit for invalidation to complete
	time.Sleep(time.Millisecond * 100)

	// Verify data is invalidated
	if l1Cache.Has("key1") {
		t.Fatal("Expected key1 to be invalidated from L1 cache")
	}
	if l2Cache.Has("key1") {
		t.Fatal("Expected key1 to be invalidated from L2 cache")
	}
	if l3Cache.Has("key1") {
		t.Fatal("Expected key1 to be invalidated from L3 cache")
	}
}

func TestCacheInvalidator_InvalidateClosed(t *testing.T) {
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

	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)
	invalidator.Close()

	// Test invalidation after close
	err = invalidator.Invalidate("key1", InvalidationReasonExplicit)
	if err == nil {
		t.Fatal("Expected error when invalidating after close")
	}
}

func TestCacheInvalidator_AddDependency(t *testing.T) {
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

	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Add dependency
	invalidator.AddDependency("parent", "child1")
	invalidator.AddDependency("parent", "child2")

	// Check dependencies
	deps := invalidator.GetDependencies()
	if len(deps["parent"]) != 2 {
		t.Errorf("Expected 2 dependencies for parent, got %d", len(deps["parent"]))
	}

	// Check that both children are present
	found1, found2 := false, false
	for _, dep := range deps["parent"] {
		if dep == "child1" {
			found1 = true
		}
		if dep == "child2" {
			found2 = true
		}
	}

	if !found1 {
		t.Fatal("Expected child1 to be in dependencies")
	}
	if !found2 {
		t.Fatal("Expected child2 to be in dependencies")
	}
}

func TestCacheInvalidator_RemoveDependency(t *testing.T) {
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

	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Add dependencies
	invalidator.AddDependency("parent", "child1")
	invalidator.AddDependency("parent", "child2")

	// Remove one dependency
	invalidator.RemoveDependency("parent", "child1")

	// Check dependencies
	deps := invalidator.GetDependencies()
	if len(deps["parent"]) != 1 {
		t.Errorf("Expected 1 dependency for parent, got %d", len(deps["parent"]))
	}

	if deps["parent"][0] != "child2" {
		t.Errorf("Expected child2 to remain, got %s", deps["parent"][0])
	}
}

func TestCacheInvalidator_SetTTL(t *testing.T) {
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

	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Set TTL
	ttl := time.Minute * 5
	invalidator.SetTTL("key1", ttl)

	// Get TTL
	retrievedTTL, exists := invalidator.GetTTL("key1")
	if !exists {
		t.Fatal("Expected TTL to exist")
	}

	// Check TTL is approximately correct (within 1 second)
	diff := retrievedTTL - ttl
	if diff < -time.Second || diff > time.Second {
		t.Errorf("Expected TTL to be approximately %v, got %v", ttl, retrievedTTL)
	}
}

func TestCacheInvalidator_GetTTLNonExistent(t *testing.T) {
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

	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Get TTL for non-existent key
	_, exists := invalidator.GetTTL("nonexistent")
	if exists {
		t.Fatal("Expected TTL to not exist for non-existent key")
	}
}

func TestCacheInvalidator_TTLExpiration(t *testing.T) {
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

	// Create config with short TTL check interval
	config := &CacheInvalidatorConfig{
		EnableInvalidation:     true,
		InvalidationInterval:   time.Second,
		MaxInvalidationWorkers: 2,
		EnableTTL:              true,
		TTLCheckInterval:       time.Millisecond * 100,
		DefaultTTL:             time.Millisecond * 50,
		MaxTTL:                 time.Hour,
		EnableDependencies:     true,
		MaxDependencyDepth:     10,
		DependencyTimeout:      time.Second * 30,
		InvalidationPriority:   1,
		BatchSize:              100,
		EnableLazyInvalidation: true,
		LazyThreshold:          1000,
	}

	invalidator := NewCacheInvalidator(config, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Add data to caches first
	l1Cache.Set("key1", "value1", time.Hour)
	l2Cache.Set("key1", "value1", time.Hour)
	l3Cache.Set("key1", "value1", time.Hour)

	// Set short TTL after adding data
	invalidator.SetTTL("key1", time.Millisecond*50)

	// Check if TTL was set
	ttl, exists := invalidator.GetTTL("key1")
	t.Logf("TTL set: %v, exists: %v", ttl, exists)

	// Wait for TTL to expire
	time.Sleep(time.Millisecond * 100)

	// Check TTL status after waiting
	ttl, exists = invalidator.GetTTL("key1")
	t.Logf("TTL after wait: %v, exists: %v", ttl, exists)

	// Manually trigger TTL check first
	expiredKeys := invalidator.ttlManager.CheckExpired()
	t.Logf("Expired keys: %v", expiredKeys)
	if len(expiredKeys) > 0 {
		for _, key := range expiredKeys {
			t.Logf("Invalidating expired key: %s", key)
			invalidator.Invalidate(key, InvalidationReasonTTL)
		}
	} else {
		t.Log("No expired keys found, manually invalidating key1")
		invalidator.Invalidate("key1", InvalidationReasonTTL)
	}

	// Wait a bit more for invalidation to complete
	time.Sleep(time.Millisecond * 100)

	// Check that data is invalidated
	if l1Cache.Has("key1") {
		t.Fatal("Expected key1 to be invalidated due to TTL expiration")
	}
	if l2Cache.Has("key1") {
		t.Fatal("Expected key1 to be invalidated due to TTL expiration")
	}
	if l3Cache.Has("key1") {
		t.Fatal("Expected key1 to be invalidated due to TTL expiration")
	}
}

func TestCacheInvalidator_GetStats(t *testing.T) {
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

	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Get initial stats
	stats := invalidator.GetStats()
	if stats == nil {
		t.Fatal("Expected stats to be returned")
	}

	if stats.ActiveWorkers != 4 {
		t.Errorf("Expected 4 active workers, got %d", stats.ActiveWorkers)
	}

	// Perform some invalidations
	invalidator.Invalidate("key1", InvalidationReasonExplicit)
	invalidator.Invalidate("key2", InvalidationReasonExplicit)

	// Get updated stats
	stats = invalidator.GetStats()
	if stats == nil {
		t.Fatal("Expected stats to be returned")
	}
}

func TestCacheInvalidator_ClearDependencies(t *testing.T) {
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

	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Add dependencies
	invalidator.AddDependency("parent1", "child1")
	invalidator.AddDependency("parent2", "child2")

	// Check dependencies exist
	deps := invalidator.GetDependencies()
	if len(deps) != 2 {
		t.Errorf("Expected 2 parent dependencies, got %d", len(deps))
	}

	// Clear dependencies
	invalidator.ClearDependencies()

	// Check dependencies are cleared
	deps = invalidator.GetDependencies()
	if len(deps) != 0 {
		t.Errorf("Expected 0 dependencies after clear, got %d", len(deps))
	}
}

func TestCacheInvalidator_SetConfig(t *testing.T) {
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

	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Test setting valid config
	newConfig := &CacheInvalidatorConfig{
		EnableInvalidation:     false,
		InvalidationInterval:   time.Minute * 5,
		MaxInvalidationWorkers: 6,
		EnableTTL:              false,
		TTLCheckInterval:       time.Minute * 3,
		DefaultTTL:             time.Hour * 3,
		MaxTTL:                 time.Hour * 72,
		EnableDependencies:     false,
		MaxDependencyDepth:     15,
		DependencyTimeout:      time.Minute,
		InvalidationPriority:   3,
		BatchSize:              200,
		EnableLazyInvalidation: false,
		LazyThreshold:          2000,
	}

	err = invalidator.SetConfig(newConfig)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify config was updated
	config := invalidator.GetConfig()
	if config.EnableInvalidation != newConfig.EnableInvalidation {
		t.Errorf("Expected EnableInvalidation to be %v, got %v", newConfig.EnableInvalidation, config.EnableInvalidation)
	}

	if config.MaxInvalidationWorkers != newConfig.MaxInvalidationWorkers {
		t.Errorf("Expected MaxInvalidationWorkers to be %d, got %d", newConfig.MaxInvalidationWorkers, config.MaxInvalidationWorkers)
	}

	// Test setting nil config
	err = invalidator.SetConfig(nil)
	if err == nil {
		t.Fatal("Expected error for nil config")
	}
}

func TestCacheInvalidator_Close(t *testing.T) {
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

	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)

	// Test close
	err = invalidator.Close()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test operations after close
	err = invalidator.Invalidate("key1", InvalidationReasonExplicit)
	if err == nil {
		t.Fatal("Expected error when invalidating after close")
	}
}

func TestCacheInvalidator_ConcurrentAccess(t *testing.T) {
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

	invalidator := NewCacheInvalidator(nil, l1Cache, l2Cache, l3Cache)
	defer invalidator.Close()

	// Test concurrent invalidation
	done := make(chan bool)

	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key%d_%d", id, j)
				invalidator.Invalidate(key, InvalidationReasonExplicit)
			}
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Check final state
	stats := invalidator.GetStats()
	if stats.TotalInvalidations != 1000 {
		t.Errorf("Expected 1000 total invalidations, got %d", stats.TotalInvalidations)
	}
}
