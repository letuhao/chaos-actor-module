package integration

import (
	"fmt"
	"time"

	"actor-core-v2/services/cache"
)

// RedisTestHelper provides helper functions for Redis testing
type RedisTestHelper struct {
	config cache.RedisConfig
}

// NewRedisTestHelper creates a new Redis test helper
func NewRedisTestHelper() *RedisTestHelper {
	return &RedisTestHelper{
		config: cache.RedisConfig{
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
		},
	}
}

// GetRedisCache creates a Redis cache instance for testing
func (h *RedisTestHelper) GetRedisCache() (*cache.RedisCache, error) {
	return cache.NewRedisCache(h.config)
}

// TestConnection tests Redis connection
func (h *RedisTestHelper) TestConnection() error {
	redisCache, err := h.GetRedisCache()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}
	defer redisCache.Close()

	// Test basic ping
	err = redisCache.Health()
	if err != nil {
		return fmt.Errorf("Redis health check failed: %w", err)
	}

	return nil
}

// TestBasicOperations tests basic Redis operations
func (h *RedisTestHelper) TestBasicOperations() error {
	redisCache, err := h.GetRedisCache()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}
	defer redisCache.Close()

	// Test set/get/delete
	key := "test_basic_ops"
	value := "test_value"
	ttl := 5 * time.Minute

	// Set
	err = redisCache.Set(key, value, ttl)
	if err != nil {
		return fmt.Errorf("failed to set key: %w", err)
	}

	// Get
	retrievedValue, err := redisCache.Get(key)
	if err != nil {
		return fmt.Errorf("failed to get key: %w", err)
	}

	if retrievedValue != value {
		return fmt.Errorf("expected value %s, got %v", value, retrievedValue)
	}

	// Exists
	if !redisCache.Exists(key) {
		return fmt.Errorf("expected key to exist")
	}

	// Delete
	err = redisCache.Delete(key)
	if err != nil {
		return fmt.Errorf("failed to delete key: %w", err)
	}

	// Verify deleted
	if redisCache.Exists(key) {
		return fmt.Errorf("expected key to be deleted")
	}

	return nil
}

// TestComplexData tests complex data structures
func (h *RedisTestHelper) TestComplexData() error {
	redisCache, err := h.GetRedisCache()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}
	defer redisCache.Close()

	// Test complex data
	testData := map[string]interface{}{
		"string": "hello world",
		"number": 42,
		"float":  3.14159,
		"bool":   true,
		"array":  []string{"apple", "banana", "cherry"},
		"object": map[string]interface{}{
			"name": "John Doe",
			"age":  30,
		},
	}

	for dataType, data := range testData {
		key := fmt.Sprintf("test_%s", dataType)

		// Set
		err = redisCache.Set(key, data, 5*time.Minute)
		if err != nil {
			return fmt.Errorf("failed to set %s: %w", dataType, err)
		}

		// Get
		retrievedValue, err := redisCache.Get(key)
		if err != nil {
			return fmt.Errorf("failed to get %s: %w", dataType, err)
		}

		if retrievedValue == nil {
			return fmt.Errorf("expected %s value, got nil", dataType)
		}

		// Clean up
		redisCache.Delete(key)
	}

	return nil
}

// TestPerformance tests Redis performance
func (h *RedisTestHelper) TestPerformance(operations int) (time.Duration, float64, error) {
	redisCache, err := h.GetRedisCache()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	defer redisCache.Close()

	start := time.Now()

	for i := 0; i < operations; i++ {
		key := fmt.Sprintf("perf_test_%d", i)
		value := fmt.Sprintf("value_%d", i)

		// Set
		err := redisCache.Set(key, value, 5*time.Minute)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to set key %s: %w", key, err)
		}

		// Get
		_, err = redisCache.Get(key)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to get key %s: %w", key, err)
		}
	}

	duration := time.Since(start)
	opsPerSecond := float64(operations*2) / duration.Seconds()

	// Clean up
	for i := 0; i < operations; i++ {
		key := fmt.Sprintf("perf_test_%d", i)
		redisCache.Delete(key)
	}

	return duration, opsPerSecond, nil
}

// TestTags tests tag-based operations
func (h *RedisTestHelper) TestTags() error {
	redisCache, err := h.GetRedisCache()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}
	defer redisCache.Close()

	// Set up test data with tags
	key1 := "actor:123:primary_core"
	key2 := "actor:123:derived_stats"
	key3 := "actor:124:primary_core"
	value := "test_value"
	tag1 := "actor:123"
	tag2 := "actor:124"

	// Set values with tags
	err = redisCache.SetWithTags(key1, value, 5*time.Minute, []string{tag1})
	if err != nil {
		return fmt.Errorf("failed to set key1 with tag: %w", err)
	}

	err = redisCache.SetWithTags(key2, value, 5*time.Minute, []string{tag1})
	if err != nil {
		return fmt.Errorf("failed to set key2 with tag: %w", err)
	}

	err = redisCache.SetWithTags(key3, value, 5*time.Minute, []string{tag2})
	if err != nil {
		return fmt.Errorf("failed to set key3 with tag: %w", err)
	}

	// Verify all keys exist
	if !redisCache.Exists(key1) || !redisCache.Exists(key2) || !redisCache.Exists(key3) {
		return fmt.Errorf("expected all keys to exist")
	}

	// Invalidate by tag1
	err = redisCache.InvalidateByTag(tag1)
	if err != nil {
		return fmt.Errorf("failed to invalidate by tag1: %w", err)
	}

	// Verify tagged keys are deleted
	if redisCache.Exists(key1) || redisCache.Exists(key2) {
		return fmt.Errorf("expected tagged keys to be deleted")
	}

	// Verify non-tagged key still exists
	if !redisCache.Exists(key3) {
		return fmt.Errorf("expected non-tagged key to still exist")
	}

	// Clean up
	redisCache.Delete(key3)

	return nil
}

// TestPubSub tests Redis Pub/Sub
func (h *RedisTestHelper) TestPubSub() error {
	redisCache, err := h.GetRedisCache()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}
	defer redisCache.Close()

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
		return fmt.Errorf("failed to publish message: %w", err)
	}

	// Subscribe to channel
	pubsub := redisCache.Subscribe(channel)
	defer pubsub.Close()

	// Wait for message (with timeout)
	timeout := time.After(2 * time.Second)
	select {
	case msg := <-pubsub.Channel():
		if msg.Channel != channel {
			return fmt.Errorf("expected channel %s, got %s", channel, msg.Channel)
		}
		// Message received successfully
	case <-timeout:
		// No message received within timeout - this might be expected in some cases
		return fmt.Errorf("no message received within timeout")
	}

	return nil
}

// GetStats returns Redis statistics
func (h *RedisTestHelper) GetStats() (*cache.CacheStats, error) {
	redisCache, err := h.GetRedisCache()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	defer redisCache.Close()

	return redisCache.GetStats(), nil
}

// GetMemoryUsage returns Redis memory usage
func (h *RedisTestHelper) GetMemoryUsage() (int64, error) {
	redisCache, err := h.GetRedisCache()
	if err != nil {
		return 0, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	defer redisCache.Close()

	return redisCache.GetMemoryUsage()
}

// GetKeys returns all keys matching a pattern
func (h *RedisTestHelper) GetKeys(pattern string) ([]string, error) {
	redisCache, err := h.GetRedisCache()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	defer redisCache.Close()

	return redisCache.GetKeys(pattern)
}

// ClearAll clears all keys in the current database
func (h *RedisTestHelper) ClearAll() error {
	redisCache, err := h.GetRedisCache()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}
	defer redisCache.Close()

	return redisCache.Clear()
}
