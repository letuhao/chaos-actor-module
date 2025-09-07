package benchmark

import (
	"testing"

	"actor-core/services/cache"
)

func BenchmarkMemCache_Set(b *testing.B) {
	mc := cache.NewMemCache(cache.MemCacheConfig{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "test_key"
		value := "test_value"
		err := mc.Set(key, value, 0)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMemCache_Get(b *testing.B) {
	mc := cache.NewMemCache(cache.MemCacheConfig{})

	// Pre-populate cache
	for i := 0; i < 1000; i++ {
		key := "test_key"
		value := "test_value"
		mc.Set(key, value, 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "test_key"
		_, err := mc.Get(key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMemCache_SetAndGet(b *testing.B) {
	mc := cache.NewMemCache(cache.MemCacheConfig{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "test_key"
		value := "test_value"

		err := mc.Set(key, value, 0)
		if err != nil {
			b.Fatal(err)
		}

		_, err = mc.Get(key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMemCache_ConcurrentAccess(b *testing.B) {
	mc := cache.NewMemCache(cache.MemCacheConfig{})

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := "test_key"
			value := "test_value"

			if i%2 == 0 {
				mc.Set(key, value, 0)
			} else {
				mc.Get(key)
			}
			i++
		}
	})
}

func BenchmarkMemCache_Exists(b *testing.B) {
	mc := cache.NewMemCache(cache.MemCacheConfig{})

	// Pre-populate cache
	for i := 0; i < 1000; i++ {
		key := "test_key"
		value := "test_value"
		mc.Set(key, value, 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "test_key"
		mc.Exists(key)
	}
}

func BenchmarkMemCache_Delete(b *testing.B) {
	mc := cache.NewMemCache(cache.MemCacheConfig{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "test_key"
		value := "test_value"

		// Set first
		mc.Set(key, value, 0)

		// Then delete
		err := mc.Delete(key)
		if err != nil {
			b.Fatal(err)
		}
	}
}
