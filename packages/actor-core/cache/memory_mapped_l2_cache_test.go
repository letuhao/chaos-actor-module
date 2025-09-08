package cache

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"
)

func TestMemoryMappedL2Cache_New(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	if cache == nil {
		t.Fatal("Expected cache to be created")
	}

	if cache.GetFileSize() == 0 {
		t.Fatal("Expected file size to be > 0")
	}
}

func TestMemoryMappedL2Cache_SetAndGet(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Test basic set and get
	err = cache.Set("key1", "value1", time.Hour)
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

func TestMemoryMappedL2Cache_GetNonExistent(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	value, exists := cache.Get("nonexistent")
	if exists {
		t.Fatal("Expected key to not exist")
	}

	if value != nil {
		t.Errorf("Expected value to be nil, got %v", value)
	}
}

func TestMemoryMappedL2Cache_SetEmptyKey(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	err = cache.Set("", "value", time.Hour)
	if err != ErrEmptyKey {
		t.Errorf("Expected ErrEmptyKey, got %v", err)
	}
}

func TestMemoryMappedL2Cache_SetNilValue(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	err = cache.Set("key", nil, time.Hour)
	if err != ErrNilValue {
		t.Errorf("Expected ErrNilValue, got %v", err)
	}
}

func TestMemoryMappedL2Cache_Expiration(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Set with short TTL
	err = cache.Set("key1", "value1", 100*time.Millisecond)
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

func TestMemoryMappedL2Cache_Delete(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Set a value
	err = cache.Set("key1", "value1", time.Hour)
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

func TestMemoryMappedL2Cache_DeleteNonExistent(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	err = cache.Delete("nonexistent")
	if err != ErrKeyNotFound {
		t.Errorf("Expected ErrKeyNotFound, got %v", err)
	}
}

func TestMemoryMappedL2Cache_Clear(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Set multiple values
	cache.Set("key1", "value1", time.Hour)
	cache.Set("key2", "value2", time.Hour)
	cache.Set("key3", "value3", time.Hour)

	// Verify they exist
	if cache.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", cache.Size())
	}

	// Clear cache
	err = cache.Clear()
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

func TestMemoryMappedL2Cache_Stats(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

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

func TestMemoryMappedL2Cache_HitRate(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

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

func TestMemoryMappedL2Cache_UsagePercentage(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

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

func TestMemoryMappedL2Cache_Has(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

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

func TestMemoryMappedL2Cache_Keys(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

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

func TestMemoryMappedL2Cache_SetMaxSize(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Set max size
	cache.SetMaxSize(500)

	if cache.MaxSize() != 500 {
		t.Errorf("Expected max size to be 500, got %d", cache.MaxSize())
	}
}

func TestMemoryMappedL2Cache_Cleanup(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

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

func TestMemoryMappedL2Cache_Close(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Set a value
	cache.Set("key1", "value1", time.Hour)

	// Close cache
	err = cache.Close()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify cache is closed
	if !cache.IsClosed() {
		t.Fatal("Expected cache to be closed")
	}

	// Operations should fail after close
	_, exists := cache.Get("key1")
	if exists {
		t.Fatal("Expected operations to fail after close")
	}
}

func TestMemoryMappedL2Cache_Flush(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Set a value
	cache.Set("key1", "value1", time.Hour)

	// Flush to disk
	err = cache.Flush()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestMemoryMappedL2Cache_FileSize(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Check initial file size
	initialSize := cache.GetFileSize()
	if initialSize == 0 {
		t.Fatal("Expected initial file size to be > 0")
	}

	// Set a value
	cache.Set("key1", "value1", time.Hour)

	// File size should increase due to buffer expansion
	currentSize := cache.GetFileSize()
	if currentSize <= initialSize {
		t.Errorf("Expected file size to be > %d, got %d", initialSize, currentSize)
	}
}

func TestMemoryMappedL2Cache_MemoryUsage(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Initial memory usage should be 0
	usage := cache.GetMemoryUsage()
	if usage != 0 {
		t.Errorf("Expected initial memory usage to be 0, got %d", usage)
	}

	// Set a value
	cache.Set("key1", "value1", time.Hour)

	// Memory usage should increase
	usage = cache.GetMemoryUsage()
	if usage == 0 {
		t.Fatal("Expected memory usage to increase after setting value")
	}
}

func TestMemoryMappedL2Cache_IndexSize(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Initial index size should be 0
	size := cache.GetIndexSize()
	if size != 0 {
		t.Errorf("Expected initial index size to be 0, got %d", size)
	}

	// Set some values
	cache.Set("key1", "value1", time.Hour)
	cache.Set("key2", "value2", time.Hour)

	// Index size should increase
	size = cache.GetIndexSize()
	if size != 2 {
		t.Errorf("Expected index size to be 2, got %d", size)
	}
}

func TestMemoryMappedL2Cache_AverageEntrySize(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Initial average should be 0
	avg := cache.GetAverageEntrySize()
	if avg != 0.0 {
		t.Errorf("Expected initial average to be 0.0, got %f", avg)
	}

	// Set some values
	cache.Set("key1", "value1", time.Hour)
	cache.Set("key2", "value2", time.Hour)

	// Average should be calculated
	avg = cache.GetAverageEntrySize()
	if avg <= 0.0 {
		t.Errorf("Expected average to be > 0.0, got %f", avg)
	}
}

func TestMemoryMappedL2Cache_CompressionRatio(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Compression ratio should be 1.0 (no compression)
	ratio := cache.GetCompressionRatio()
	if ratio != 1.0 {
		t.Errorf("Expected compression ratio to be 1.0, got %f", ratio)
	}
}

func TestMemoryMappedL2Cache_FragmentationRatio(t *testing.T) {
	// Create temporary directory
	tempDir := t.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 100)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	// Fragmentation ratio should be 0.0 (no fragmentation)
	ratio := cache.GetFragmentationRatio()
	if ratio != 0.0 {
		t.Errorf("Expected fragmentation ratio to be 0.0, got %f", ratio)
	}
}

func BenchmarkMemoryMappedL2Cache_Set(b *testing.B) {
	// Create temporary directory
	tempDir := b.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 10000)
	if err != nil {
		b.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)
		cache.Set(key, value, time.Hour)
	}
}

func BenchmarkMemoryMappedL2Cache_Get(b *testing.B) {
	// Create temporary directory
	tempDir := b.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 10000)
	if err != nil {
		b.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

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

func BenchmarkMemoryMappedL2Cache_Concurrent(b *testing.B) {
	// Create temporary directory
	tempDir := b.TempDir()
	cachePath := filepath.Join(tempDir, "test_cache.dat")

	cache, err := NewMemoryMappedL2Cache(cachePath, 10000)
	if err != nil {
		b.Fatalf("Expected no error, got %v", err)
	}
	defer cache.Close()

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
