package unit

import (
	"fmt"
	"testing"
	"time"

	"actor-core-v2/services/cache"
)

func TestNewMemCache(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:         "100MB",
		DefaultTTL:      5 * time.Minute,
		EvictionPolicy:  "lru",
		EnableStats:     true,
		CleanupInterval: 1 * time.Minute,
	}

	mc := cache.NewMemCache(config)

	if mc == nil {
		t.Error("Expected MemCache to be created")
	}

	stats := mc.GetStats()
	if stats == nil {
		t.Error("Expected stats to be initialized")
	}

	if stats.Hits != 0 {
		t.Errorf("Expected hits to be 0, got %d", stats.Hits)
	}

	if stats.Misses != 0 {
		t.Errorf("Expected misses to be 0, got %d", stats.Misses)
	}
}

func TestMemCacheSetAndGet(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Test setting and getting a value
	key := "test_key"
	value := "test_value"
	ttl := 5 * time.Minute

	err := mc.Set(key, value, ttl)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	retrievedValue, err := mc.Get(key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if retrievedValue != value {
		t.Errorf("Expected value to be %s, got %v", value, retrievedValue)
	}

	// Test statistics
	stats := mc.GetStats()
	if stats.Hits != 1 {
		t.Errorf("Expected hits to be 1, got %d", stats.Hits)
	}

	if stats.Misses != 0 {
		t.Errorf("Expected misses to be 0, got %d", stats.Misses)
	}
}

func TestMemCacheGetNonExistent(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Test getting non-existent key
	_, err := mc.Get("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent key")
	}

	// Test statistics
	stats := mc.GetStats()
	if stats.Misses != 1 {
		t.Errorf("Expected misses to be 1, got %d", stats.Misses)
	}
}

func TestMemCacheDelete(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Set a value
	key := "test_key"
	value := "test_value"
	mc.Set(key, value, 5*time.Minute)

	// Verify it exists
	if !mc.Exists(key) {
		t.Error("Expected key to exist")
	}

	// Delete the key
	err := mc.Delete(key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify it's deleted
	if mc.Exists(key) {
		t.Error("Expected key to be deleted")
	}

	// Test deleting non-existent key
	err = mc.Delete("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent key")
	}
}

func TestMemCacheClear(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Set multiple values
	mc.Set("key1", "value1", 5*time.Minute)
	mc.Set("key2", "value2", 5*time.Minute)
	mc.Set("key3", "value3", 5*time.Minute)

	// Verify they exist
	if !mc.Exists("key1") || !mc.Exists("key2") || !mc.Exists("key3") {
		t.Error("Expected all keys to exist")
	}

	// Clear cache
	err := mc.Clear()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify all keys are deleted
	if mc.Exists("key1") || mc.Exists("key2") || mc.Exists("key3") {
		t.Error("Expected all keys to be deleted")
	}
}

func TestMemCacheTTL(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Set a value with TTL
	key := "test_key"
	value := "test_value"
	ttl := 1 * time.Second
	mc.Set(key, value, ttl)

	// Get TTL
	retrievedTTL, err := mc.GetTTL(key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if retrievedTTL <= 0 || retrievedTTL > ttl {
		t.Errorf("Expected TTL to be between 0 and %v, got %v", ttl, retrievedTTL)
	}

	// Set TTL
	newTTL := 2 * time.Second
	err = mc.SetTTL(key, newTTL)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify new TTL
	retrievedTTL, err = mc.GetTTL(key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if retrievedTTL <= 0 || retrievedTTL > newTTL {
		t.Errorf("Expected TTL to be between 0 and %v, got %v", newTTL, retrievedTTL)
	}
}

func TestMemCacheExpiration(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Set a value with very short TTL
	key := "test_key"
	value := "test_value"
	ttl := 500 * time.Millisecond
	mc.Set(key, value, ttl)

	// Verify it exists initially
	if !mc.Exists(key) {
		t.Error("Expected key to exist initially")
	}

	// Wait for expiration
	time.Sleep(600 * time.Millisecond)

	// Verify it's expired (Exists should remove expired entries)
	if mc.Exists(key) {
		t.Error("Expected key to be expired")
	}

	// Try to get expired key
	_, err := mc.Get(key)
	if err == nil {
		t.Error("Expected error for expired key")
	}
}

func TestMemCacheStats(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Perform operations
	mc.Set("key1", "value1", 5*time.Minute)
	mc.Get("key1") // Hit
	mc.Get("key2") // Miss
	mc.Set("key2", "value2", 5*time.Minute)
	mc.Get("key2") // Hit

	// Check stats
	stats := mc.GetStats()
	if stats.Hits != 2 {
		t.Errorf("Expected hits to be 2, got %d", stats.Hits)
	}

	if stats.Misses != 1 {
		t.Errorf("Expected misses to be 1, got %d", stats.Misses)
	}

	if stats.HitRatio != 2.0/3.0 {
		t.Errorf("Expected hit ratio to be %.2f, got %.2f", 2.0/3.0, stats.HitRatio)
	}
}

func TestMemCacheResetStats(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Perform operations
	mc.Set("key1", "value1", 5*time.Minute)
	mc.Get("key1")
	mc.Get("key2")

	// Reset stats
	err := mc.ResetStats()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check stats are reset
	stats := mc.GetStats()
	if stats.Hits != 0 {
		t.Errorf("Expected hits to be 0, got %d", stats.Hits)
	}

	if stats.Misses != 0 {
		t.Errorf("Expected misses to be 0, got %d", stats.Misses)
	}
}

func TestMemCacheHealth(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Test health check
	err := mc.Health()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestMemCacheConcurrentAccess(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Test concurrent access
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			key := fmt.Sprintf("key_%d", i)
			value := fmt.Sprintf("value_%d", i)

			mc.Set(key, value, 5*time.Minute)
			mc.Get(key)
			mc.Exists(key)

			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify all keys exist
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		if !mc.Exists(key) {
			t.Errorf("Expected key %s to exist", key)
		}
	}
}

func TestMemCacheGetKeys(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Set multiple keys
	keys := []string{"key1", "key2", "key3"}
	for _, key := range keys {
		mc.Set(key, "value", 5*time.Minute)
	}

	// Get all keys
	retrievedKeys := mc.GetKeys()
	if len(retrievedKeys) != len(keys) {
		t.Errorf("Expected %d keys, got %d", len(keys), len(retrievedKeys))
	}

	// Check that all expected keys are present
	keyMap := make(map[string]bool)
	for _, key := range retrievedKeys {
		keyMap[key] = true
	}

	for _, expectedKey := range keys {
		if !keyMap[expectedKey] {
			t.Errorf("Expected key %s to be present", expectedKey)
		}
	}
}

func TestMemCacheGetSize(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Initial size should be 0
	if mc.GetSize() != 0 {
		t.Errorf("Expected initial size to be 0, got %d", mc.GetSize())
	}

	// Set some values
	mc.Set("key1", "value1", 5*time.Minute)
	mc.Set("key2", "value2", 5*time.Minute)

	// Size should be greater than 0
	if mc.GetSize() <= 0 {
		t.Error("Expected size to be greater than 0")
	}
}

func TestMemCacheGetEntryCount(t *testing.T) {
	config := cache.MemCacheConfig{
		MaxSize:        "100MB",
		DefaultTTL:     5 * time.Minute,
		EvictionPolicy: "lru",
		EnableStats:    true,
	}

	mc := cache.NewMemCache(config)

	// Initial count should be 0
	if mc.GetEntryCount() != 0 {
		t.Errorf("Expected initial entry count to be 0, got %d", mc.GetEntryCount())
	}

	// Set some values
	mc.Set("key1", "value1", 5*time.Minute)
	mc.Set("key2", "value2", 5*time.Minute)

	// Count should be 2
	if mc.GetEntryCount() != 2 {
		t.Errorf("Expected entry count to be 2, got %d", mc.GetEntryCount())
	}
}
