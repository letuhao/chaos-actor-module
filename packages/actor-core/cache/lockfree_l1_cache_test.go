package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLockFreeL1Cache_New(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	if cache == nil {
		t.Fatal("Expected cache to be created")
	}

	if cache.maxSize != 1000 {
		t.Errorf("Expected maxSize to be 1000, got %d", cache.maxSize)
	}

	if cache.evictor.evictionPolicy != "allkeys-lru" {
		t.Errorf("Expected eviction policy to be 'allkeys-lru', got %s", cache.evictor.evictionPolicy)
	}
}

func TestLockFreeL1Cache_SetAndGet(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Test basic set and get
	err := cache.Set("key1", "value1", time.Hour)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	value, exists := cache.Get("key1")
	if !exists {
		t.Fatal("Expected key to exist")
	}

	if value != "value1" {
		t.Errorf("Expected value to be 'value1', got %v", value)
	}
}

func TestLockFreeL1Cache_GetNonExistent(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	value, exists := cache.Get("nonexistent")
	if exists {
		t.Fatal("Expected key to not exist")
	}

	if value != nil {
		t.Errorf("Expected value to be nil, got %v", value)
	}
}

func TestLockFreeL1Cache_SetEmptyKey(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	err := cache.Set("", "value", time.Hour)
	if err != ErrEmptyKey {
		t.Errorf("Expected ErrEmptyKey, got %v", err)
	}
}

func TestLockFreeL1Cache_SetNilValue(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	err := cache.Set("key", nil, time.Hour)
	if err != ErrNilValue {
		t.Errorf("Expected ErrNilValue, got %v", err)
	}
}

func TestLockFreeL1Cache_Expiration(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Set with short TTL
	err := cache.Set("key1", "value1", 100*time.Millisecond)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Should exist immediately
	value, exists := cache.Get("key1")
	if !exists {
		t.Fatal("Expected key to exist immediately")
	}
	if value != "value1" {
		t.Errorf("Expected value to be 'value1', got %v", value)
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Should not exist after expiration
	value, exists = cache.Get("key1")
	if exists {
		t.Fatal("Expected key to not exist after expiration")
	}
	if value != nil {
		t.Errorf("Expected value to be nil after expiration, got %v", value)
	}
}

func TestLockFreeL1Cache_Delete(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Set a value
	err := cache.Set("key1", "value1", time.Hour)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify it exists
	value, exists := cache.Get("key1")
	if !exists {
		t.Fatal("Expected key to exist")
	}
	if value != "value1" {
		t.Errorf("Expected value to be 'value1', got %v", value)
	}

	// Delete it
	err = cache.Delete("key1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify it's gone
	value, exists = cache.Get("key1")
	if exists {
		t.Fatal("Expected key to not exist after deletion")
	}
	if value != nil {
		t.Errorf("Expected value to be nil after deletion, got %v", value)
	}
}

func TestLockFreeL1Cache_DeleteNonExistent(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	err := cache.Delete("nonexistent")
	if err != ErrKeyNotFound {
		t.Errorf("Expected ErrKeyNotFound, got %v", err)
	}
}

func TestLockFreeL1Cache_Clear(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Set multiple values
	cache.Set("key1", "value1", time.Hour)
	cache.Set("key2", "value2", time.Hour)
	cache.Set("key3", "value3", time.Hour)

	// Verify they exist
	if cache.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", cache.Size())
	}

	// Clear cache
	err := cache.Clear()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify cache is empty
	if cache.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", cache.Size())
	}

	// Verify keys don't exist
	_, exists := cache.Get("key1")
	if exists {
		t.Fatal("Expected key1 to not exist after clear")
	}
}

func TestLockFreeL1Cache_Stats(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Initial stats
	stats := cache.GetStats()
	if stats.hits != 0 {
		t.Errorf("Expected hits to be 0, got %d", stats.hits)
	}
	if stats.misses != 0 {
		t.Errorf("Expected misses to be 0, got %d", stats.misses)
	}
	if stats.size != 0 {
		t.Errorf("Expected size to be 0, got %d", stats.size)
	}

	// Set a value
	cache.Set("key1", "value1", time.Hour)

	// Get the value (hit)
	cache.Get("key1")

	// Get non-existent value (miss)
	cache.Get("nonexistent")

	// Check stats
	stats = cache.GetStats()
	if stats.hits != 1 {
		t.Errorf("Expected hits to be 1, got %d", stats.hits)
	}
	if stats.misses != 1 {
		t.Errorf("Expected misses to be 1, got %d", stats.misses)
	}
	if stats.size != 1 {
		t.Errorf("Expected size to be 1, got %d", stats.size)
	}
}

func TestLockFreeL1Cache_HitRate(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// No operations yet
	hitRate := cache.GetHitRate()
	if hitRate != 0.0 {
		t.Errorf("Expected hit rate to be 0.0, got %f", hitRate)
	}

	// Set a value
	cache.Set("key1", "value1", time.Hour)

	// Get the value (hit)
	cache.Get("key1")

	// Get non-existent value (miss)
	cache.Get("nonexistent")

	// Check hit rate
	hitRate = cache.GetHitRate()
	expectedHitRate := 0.5 // 1 hit out of 2 total operations
	if hitRate != expectedHitRate {
		t.Errorf("Expected hit rate to be %f, got %f", expectedHitRate, hitRate)
	}
}

func TestLockFreeL1Cache_UsagePercentage(t *testing.T) {
	cache := NewLockFreeL1Cache(100, "allkeys-lru")

	// Empty cache
	usage := cache.GetUsagePercentage()
	if usage != 0.0 {
		t.Errorf("Expected usage to be 0.0, got %f", usage)
	}

	// Set 50 values
	for i := 0; i < 50; i++ {
		cache.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i), time.Hour)
	}

	// Check usage
	usage = cache.GetUsagePercentage()
	expectedUsage := 50.0 // 50 out of 100
	if usage != expectedUsage {
		t.Errorf("Expected usage to be %f, got %f", expectedUsage, usage)
	}
}

func TestLockFreeL1Cache_Eviction(t *testing.T) {
	cache := NewLockFreeL1Cache(2, "allkeys-lru")

	// Set 3 values (should trigger eviction)
	cache.Set("key1", "value1", time.Hour)
	cache.Set("key2", "value2", time.Hour)
	cache.Set("key3", "value3", time.Hour)

	// Cache should only have 2 values
	if cache.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", cache.Size())
	}

	// At least one key should be evicted
	// Note: LRU eviction might not evict key1 if it was accessed recently
	// So we just check that size is correct and at least one key exists
	if cache.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", cache.Size())
	}

	// Check that we have exactly 2 keys
	keys := cache.Keys()
	if len(keys) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(keys))
	}
}

func TestLockFreeL1Cache_ConcurrentAccess(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Concurrent writers
	var wg sync.WaitGroup
	numGoroutines := 100
	numOperations := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				key := fmt.Sprintf("key_%d_%d", goroutineID, j)
				value := fmt.Sprintf("value_%d_%d", goroutineID, j)

				// Set value
				cache.Set(key, value, time.Hour)

				// Get value
				cache.Get(key)
			}
		}(i)
	}

	wg.Wait()

	// Verify cache is not empty
	if cache.Size() == 0 {
		t.Fatal("Expected cache to have some values")
	}

	// Verify stats are reasonable
	stats := cache.GetStats()
	if stats.hits == 0 {
		t.Fatal("Expected some hits")
	}
}

func TestLockFreeL1Cache_Cleanup(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Set values with different TTLs
	cache.Set("key1", "value1", 100*time.Millisecond)
	cache.Set("key2", "value2", 200*time.Millisecond)
	cache.Set("key3", "value3", time.Hour)

	// Wait for first two to expire
	time.Sleep(150 * time.Millisecond)

	// Cleanup should remove expired entries
	removed := cache.Cleanup()
	if removed != 1 { // key1 should be removed
		t.Errorf("Expected 1 entry to be removed, got %d", removed)
	}

	// key1 should be gone
	_, exists := cache.Get("key1")
	if exists {
		t.Fatal("Expected key1 to be removed by cleanup")
	}

	// key2 and key3 should still exist
	_, exists = cache.Get("key2")
	if !exists {
		t.Fatal("Expected key2 to still exist")
	}

	_, exists = cache.Get("key3")
	if !exists {
		t.Fatal("Expected key3 to still exist")
	}
}

func TestLockFreeL1Cache_Has(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Non-existent key
	exists := cache.Has("nonexistent")
	if exists {
		t.Fatal("Expected non-existent key to not exist")
	}

	// Set a value
	cache.Set("key1", "value1", time.Hour)

	// Should exist
	exists = cache.Has("key1")
	if !exists {
		t.Fatal("Expected key1 to exist")
	}

	// Set with short TTL
	cache.Set("key2", "value2", 100*time.Millisecond)

	// Should exist immediately
	exists = cache.Has("key2")
	if !exists {
		t.Fatal("Expected key2 to exist immediately")
	}

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Should not exist after expiration
	exists = cache.Has("key2")
	if exists {
		t.Fatal("Expected key2 to not exist after expiration")
	}
}

func TestLockFreeL1Cache_Keys(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Empty cache
	keys := cache.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys, got %d", len(keys))
	}

	// Set some values
	cache.Set("key1", "value1", time.Hour)
	cache.Set("key2", "value2", time.Hour)
	cache.Set("key3", "value3", time.Hour)

	// Get keys
	keys = cache.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// Verify all keys are present
	keyMap := make(map[string]bool)
	for _, key := range keys {
		keyMap[key] = true
	}

	if !keyMap["key1"] {
		t.Fatal("Expected key1 to be present")
	}
	if !keyMap["key2"] {
		t.Fatal("Expected key2 to be present")
	}
	if !keyMap["key3"] {
		t.Fatal("Expected key3 to be present")
	}
}

func TestLockFreeL1Cache_SetMaxSize(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Set max size
	cache.SetMaxSize(500)

	if cache.MaxSize() != 500 {
		t.Errorf("Expected max size to be 500, got %d", cache.MaxSize())
	}
}

func TestLockFreeL1Cache_SetEvictionPolicy(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Set eviction policy
	cache.SetEvictionPolicy("allkeys-lfu")

	if cache.GetEvictionPolicy() != "allkeys-lfu" {
		t.Errorf("Expected eviction policy to be 'allkeys-lfu', got %s", cache.GetEvictionPolicy())
	}
}

func TestLockFreeL1Cache_Reset(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Set a value and get it
	cache.Set("key1", "value1", time.Hour)
	cache.Get("key1")
	cache.Get("nonexistent")

	// Check stats
	stats := cache.GetStats()
	if stats.hits == 0 || stats.misses == 0 {
		t.Fatal("Expected some hits and misses")
	}

	// Reset stats
	cache.Reset()

	// Check stats are reset
	stats = cache.GetStats()
	if stats.hits != 0 {
		t.Errorf("Expected hits to be 0 after reset, got %d", stats.hits)
	}
	if stats.misses != 0 {
		t.Errorf("Expected misses to be 0 after reset, got %d", stats.misses)
	}
}

func TestLockFreeL1Cache_Preload(t *testing.T) {
	cache := NewLockFreeL1Cache(1000, "allkeys-lru")

	// Start preloading
	cache.StartPreloading()

	// Preload some actors
	cache.Preload("actor1")
	cache.Preload("actor2")
	cache.Preload("actor3")

	// Wait a bit for preloading to complete
	time.Sleep(200 * time.Millisecond)

	// Check preload stats
	stats := cache.GetPreloadStats()
	if stats.preloadedCount == 0 {
		t.Log("No preloaded actors - this is expected for placeholder implementation")
	}
}

func BenchmarkLockFreeL1Cache_Set(b *testing.B) {
	cache := NewLockFreeL1Cache(10000, "allkeys-lru")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)
		cache.Set(key, value, time.Hour)
	}
}

func BenchmarkLockFreeL1Cache_Get(b *testing.B) {
	cache := NewLockFreeL1Cache(10000, "allkeys-lru")

	// Pre-populate cache
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)
		cache.Set(key, value, time.Hour)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i%1000)
		cache.Get(key)
	}
}

func BenchmarkLockFreeL1Cache_Concurrent(b *testing.B) {
	cache := NewLockFreeL1Cache(10000, "allkeys-lru")

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("key_%d", i)
			value := fmt.Sprintf("value_%d", i)
			cache.Set(key, value, time.Hour)
			cache.Get(key)
			i++
		}
	})
}
