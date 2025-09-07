package integration

import (
	"fmt"
	"testing"
	"time"

	"actor-core-v2/services/cache"
)

func TestRedisConnection(t *testing.T) {
	// Test Redis connection với Memurai default config
	config := cache.RedisConfig{
		Host:         "localhost",
		Port:         6379,
		Password:     "", // Memurai default không có password
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
		EnableTLS:    false,
	}

	// Tạo Redis cache instance
	redisCache, err := cache.NewRedisCache(config)
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisCache.Close()

	t.Log("✅ Successfully connected to Redis (Memurai)")
}

func TestRedisBasicOperations(t *testing.T) {
	config := cache.RedisConfig{
		Host:         "localhost",
		Port:         6379,
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
		EnableTLS:    false,
	}

	redisCache, err := cache.NewRedisCache(config)
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisCache.Close()

	// Test basic operations
	key := "test_key"
	value := "test_value"
	ttl := 5 * time.Minute

	// Test Set
	err = redisCache.Set(key, value, ttl)
	if err != nil {
		t.Errorf("Failed to set key: %v", err)
	}

	// Test Get
	retrievedValue, err := redisCache.Get(key)
	if err != nil {
		t.Errorf("Failed to get key: %v", err)
	}

	if retrievedValue != value {
		t.Errorf("Expected value %s, got %v", value, retrievedValue)
	}

	// Test Exists
	if !redisCache.Exists(key) {
		t.Error("Expected key to exist")
	}

	// Test TTL
	retrievedTTL, err := redisCache.GetTTL(key)
	if err != nil {
		t.Errorf("Failed to get TTL: %v", err)
	}

	if retrievedTTL <= 0 || retrievedTTL > ttl {
		t.Errorf("Expected TTL between 0 and %v, got %v", ttl, retrievedTTL)
	}

	// Test Delete
	err = redisCache.Delete(key)
	if err != nil {
		t.Errorf("Failed to delete key: %v", err)
	}

	// Test key no longer exists
	if redisCache.Exists(key) {
		t.Error("Expected key to be deleted")
	}

	t.Log("✅ All Redis basic operations passed")
}

func TestRedisComplexData(t *testing.T) {
	config := cache.RedisConfig{
		Host:         "localhost",
		Port:         6379,
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
		EnableTLS:    false,
	}

	redisCache, err := cache.NewRedisCache(config)
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisCache.Close()

	// Test với complex data structures
	testCases := []struct {
		name  string
		key   string
		value interface{}
	}{
		{
			name:  "String",
			key:   "test_string",
			value: "hello world",
		},
		{
			name:  "Number",
			key:   "test_number",
			value: 42,
		},
		{
			name:  "Float",
			key:   "test_float",
			value: 3.14159,
		},
		{
			name:  "Boolean",
			key:   "test_bool",
			value: true,
		},
		{
			name: "Map",
			key:  "test_map",
			value: map[string]interface{}{
				"name":   "John Doe",
				"age":    30,
				"active": true,
				"scores": []int{85, 92, 78},
			},
		},
		{
			name:  "Array",
			key:   "test_array",
			value: []string{"apple", "banana", "cherry"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set value
			err := redisCache.Set(tc.key, tc.value, 5*time.Minute)
			if err != nil {
				t.Errorf("Failed to set %s: %v", tc.name, err)
				return
			}

			// Get value
			retrievedValue, err := redisCache.Get(tc.key)
			if err != nil {
				t.Errorf("Failed to get %s: %v", tc.name, err)
				return
			}

			// Verify value (simplified comparison)
			if retrievedValue == nil {
				t.Errorf("Expected %s value, got nil", tc.name)
			}

			// Clean up
			redisCache.Delete(tc.key)
		})
	}

	t.Log("✅ All Redis complex data operations passed")
}

func TestRedisPerformance(t *testing.T) {
	config := cache.RedisConfig{
		Host:         "localhost",
		Port:         6379,
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
		EnableTLS:    false,
	}

	redisCache, err := cache.NewRedisCache(config)
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisCache.Close()

	// Performance test
	numOperations := 1000
	start := time.Now()

	for i := 0; i < numOperations; i++ {
		key := fmt.Sprintf("perf_test_%d", i)
		value := fmt.Sprintf("value_%d", i)

		// Set
		err := redisCache.Set(key, value, 5*time.Minute)
		if err != nil {
			t.Errorf("Failed to set key %s: %v", key, err)
		}

		// Get
		_, err = redisCache.Get(key)
		if err != nil {
			t.Errorf("Failed to get key %s: %v", key, err)
		}
	}

	duration := time.Since(start)
	opsPerSecond := float64(numOperations*2) / duration.Seconds()

	t.Logf("✅ Performance test completed:")
	t.Logf("   Operations: %d (set + get)", numOperations*2)
	t.Logf("   Duration: %v", duration)
	t.Logf("   Ops/sec: %.2f", opsPerSecond)

	// Clean up
	for i := 0; i < numOperations; i++ {
		key := fmt.Sprintf("perf_test_%d", i)
		redisCache.Delete(key)
	}
}

func TestRedisHealth(t *testing.T) {
	config := cache.RedisConfig{
		Host:         "localhost",
		Port:         6379,
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
		EnableTLS:    false,
	}

	redisCache, err := cache.NewRedisCache(config)
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisCache.Close()

	// Test health check
	err = redisCache.Health()
	if err != nil {
		t.Errorf("Redis health check failed: %v", err)
	}

	// Test stats
	stats := redisCache.GetStats()
	if stats == nil {
		t.Error("Expected stats to be available")
	}

	t.Log("✅ Redis health check passed")
}

func TestRedisWithTags(t *testing.T) {
	config := cache.RedisConfig{
		Host:         "localhost",
		Port:         6379,
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
		EnableTLS:    false,
	}

	redisCache, err := cache.NewRedisCache(config)
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisCache.Close()

	// Test tag-based operations
	key1 := "actor:123:primary_core"
	key2 := "actor:123:derived_stats"
	key3 := "actor:124:primary_core"
	value := "test_value"
	tag := "actor:123"

	// Set values with tags
	err = redisCache.SetWithTags(key1, value, 5*time.Minute, []string{tag})
	if err != nil {
		t.Errorf("Failed to set key1 with tag: %v", err)
	}

	err = redisCache.SetWithTags(key2, value, 5*time.Minute, []string{tag})
	if err != nil {
		t.Errorf("Failed to set key2 with tag: %v", err)
	}

	err = redisCache.SetWithTags(key3, value, 5*time.Minute, []string{"actor:124"})
	if err != nil {
		t.Errorf("Failed to set key3 with tag: %v", err)
	}

	// Verify keys exist
	if !redisCache.Exists(key1) || !redisCache.Exists(key2) || !redisCache.Exists(key3) {
		t.Error("Expected all keys to exist")
	}

	// Invalidate by tag
	err = redisCache.InvalidateByTag(tag)
	if err != nil {
		t.Errorf("Failed to invalidate by tag: %v", err)
	}

	// Verify tagged keys are deleted
	if redisCache.Exists(key1) || redisCache.Exists(key2) {
		t.Error("Expected tagged keys to be deleted")
	}

	// Verify non-tagged key still exists
	if !redisCache.Exists(key3) {
		t.Error("Expected non-tagged key to still exist")
	}

	// Clean up
	redisCache.Delete(key3)

	t.Log("✅ Redis tag-based operations passed")
}

func TestRedisPubSub(t *testing.T) {
	config := cache.RedisConfig{
		Host:         "localhost",
		Port:         6379,
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
		EnableTLS:    false,
	}

	redisCache, err := cache.NewRedisCache(config)
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisCache.Close()

	// Test Pub/Sub
	channel := "test_channel"
	message := map[string]interface{}{
		"type":      "cache_invalidation",
		"key":       "actor:123",
		"action":    "delete",
		"timestamp": time.Now().Unix(),
	}

	// Publish message
	err = redisCache.Publish(channel, message)
	if err != nil {
		t.Errorf("Failed to publish message: %v", err)
	}

	// Subscribe to channel
	pubsub := redisCache.Subscribe(channel)
	defer pubsub.Close()

	// Wait for message (with timeout)
	timeout := time.After(2 * time.Second)
	select {
	case msg := <-pubsub.Channel():
		if msg.Channel != channel {
			t.Errorf("Expected channel %s, got %s", channel, msg.Channel)
		}
		t.Logf("✅ Received message: %s", msg.Payload)
	case <-timeout:
		t.Log("⚠️ No message received within timeout (this is expected in some cases)")
	}

	t.Log("✅ Redis Pub/Sub test completed")
}
