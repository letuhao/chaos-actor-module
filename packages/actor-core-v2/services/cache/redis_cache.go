package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache implements Redis-based cache
type RedisCache struct {
	client  *redis.Client
	config  RedisConfig
	stats   *CacheStats
	ctx     context.Context
	created int64
}

// NewRedisCache creates a new RedisCache instance
func NewRedisCache(config RedisConfig) (*RedisCache, error) {
	now := time.Now().Unix()

	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password:     config.Password,
		DB:           config.DB,
		MaxRetries:   config.MaxRetries,
		DialTimeout:  config.DialTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
	})

	// Test connection
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisCache{
		client:  client,
		config:  config,
		stats:   &CacheStats{LastReset: now},
		ctx:     ctx,
		created: now,
	}, nil
}

// Get retrieves a value from Redis cache
func (rc *RedisCache) Get(key string) (interface{}, error) {
	val, err := rc.client.Get(rc.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			rc.stats.Misses++
			rc.updateHitRatio()
			return nil, fmt.Errorf("key not found: %s", key)
		}
		rc.stats.Errors++
		return nil, fmt.Errorf("Redis get error: %w", err)
	}

	// Deserialize value
	var value interface{}
	if err := json.Unmarshal([]byte(val), &value); err != nil {
		rc.stats.Errors++
		return nil, fmt.Errorf("failed to unmarshal value: %w", err)
	}

	rc.stats.Hits++
	rc.updateHitRatio()
	return value, nil
}

// Set stores a value in Redis cache
func (rc *RedisCache) Set(key string, value interface{}, ttl time.Duration) error {
	// Serialize value
	data, err := json.Marshal(value)
	if err != nil {
		rc.stats.Errors++
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	// Set in Redis
	err = rc.client.Set(rc.ctx, key, data, ttl).Err()
	if err != nil {
		rc.stats.Errors++
		return fmt.Errorf("Redis set error: %w", err)
	}

	rc.stats.Size += int64(len(data))
	return nil
}

// Delete removes a key from Redis cache
func (rc *RedisCache) Delete(key string) error {
	err := rc.client.Del(rc.ctx, key).Err()
	if err != nil {
		rc.stats.Errors++
		return fmt.Errorf("Redis delete error: %w", err)
	}
	return nil
}

// Clear removes all keys from Redis cache
func (rc *RedisCache) Clear() error {
	err := rc.client.FlushDB(rc.ctx).Err()
	if err != nil {
		rc.stats.Errors++
		return fmt.Errorf("Redis clear error: %w", err)
	}
	rc.stats.Size = 0
	return nil
}

// Exists checks if a key exists in Redis cache
func (rc *RedisCache) Exists(key string) bool {
	exists, err := rc.client.Exists(rc.ctx, key).Result()
	if err != nil {
		rc.stats.Errors++
		return false
	}
	return exists > 0
}

// GetTTL returns the time to live for a key
func (rc *RedisCache) GetTTL(key string) (time.Duration, error) {
	ttl, err := rc.client.TTL(rc.ctx, key).Result()
	if err != nil {
		rc.stats.Errors++
		return 0, fmt.Errorf("Redis TTL error: %w", err)
	}

	if ttl == -2 {
		return 0, fmt.Errorf("key not found: %s", key)
	}

	if ttl == -1 {
		return 0, nil // No expiration
	}

	return ttl, nil
}

// SetTTL sets the time to live for a key
func (rc *RedisCache) SetTTL(key string, ttl time.Duration) error {
	err := rc.client.Expire(rc.ctx, key, ttl).Err()
	if err != nil {
		rc.stats.Errors++
		return fmt.Errorf("Redis expire error: %w", err)
	}
	return nil
}

// GetStats returns cache statistics
func (rc *RedisCache) GetStats() *CacheStats {
	// Get Redis info
	info, err := rc.client.Info(rc.ctx, "memory").Result()
	if err != nil {
		rc.stats.Errors++
		return rc.stats
	}

	// Parse memory info (simplified)
	// In a real implementation, you'd parse the Redis info string properly
	rc.stats.MaxSize = rc.parseMemoryInfo(info)

	// Create a copy to avoid race conditions
	stats := *rc.stats
	return &stats
}

// ResetStats resets cache statistics
func (rc *RedisCache) ResetStats() error {
	now := time.Now().Unix()
	rc.stats = &CacheStats{
		LastReset: now,
	}
	return nil
}

// Health checks Redis cache health
func (rc *RedisCache) Health() error {
	// Ping Redis
	_, err := rc.client.Ping(rc.ctx).Result()
	if err != nil {
		return fmt.Errorf("Redis health check failed: %w", err)
	}

	// Check if hit ratio is reasonable
	if rc.stats.Hits+rc.stats.Misses > 0 && rc.stats.HitRatio < 0.1 {
		return fmt.Errorf("cache hit ratio too low: %.2f", rc.stats.HitRatio)
	}

	return nil
}

// Private methods

// updateHitRatio updates the hit ratio
func (rc *RedisCache) updateHitRatio() {
	total := rc.stats.Hits + rc.stats.Misses
	if total > 0 {
		rc.stats.HitRatio = float64(rc.stats.Hits) / float64(total)
	}
}

// parseMemoryInfo parses Redis memory info
func (rc *RedisCache) parseMemoryInfo(info string) int64 {
	// Simplified parsing - in a real implementation, you'd parse the full info string
	// For now, return a default value
	return 100 * 1024 * 1024 // 100MB default
}

// GetClient returns the Redis client
func (rc *RedisCache) GetClient() *redis.Client {
	return rc.client
}

// GetContext returns the Redis context
func (rc *RedisCache) GetContext() context.Context {
	return rc.ctx
}

// Close closes the Redis connection
func (rc *RedisCache) Close() error {
	return rc.client.Close()
}

// GetKeys returns all cache keys matching a pattern
func (rc *RedisCache) GetKeys(pattern string) ([]string, error) {
	keys, err := rc.client.Keys(rc.ctx, pattern).Result()
	if err != nil {
		rc.stats.Errors++
		return nil, fmt.Errorf("Redis keys error: %w", err)
	}
	return keys, nil
}

// GetSize returns the current cache size
func (rc *RedisCache) GetSize() int64 {
	// Get database size
	size, err := rc.client.DBSize(rc.ctx).Result()
	if err != nil {
		rc.stats.Errors++
		return 0
	}
	return size
}

// GetMemoryUsage returns Redis memory usage
func (rc *RedisCache) GetMemoryUsage() (int64, error) {
	info, err := rc.client.Info(rc.ctx, "memory").Result()
	if err != nil {
		rc.stats.Errors++
		return 0, fmt.Errorf("Redis memory info error: %w", err)
	}

	// Parse memory usage from info string
	// This is simplified - in a real implementation, you'd parse the full info
	return rc.parseMemoryInfo(info), nil
}

// SetWithTags sets a value with tags for easier invalidation
func (rc *RedisCache) SetWithTags(key string, value interface{}, ttl time.Duration, tags []string) error {
	// Set the main key
	if err := rc.Set(key, value, ttl); err != nil {
		return err
	}

	// Set tags for this key
	for _, tag := range tags {
		tagKey := fmt.Sprintf("tag:%s", tag)
		err := rc.client.SAdd(rc.ctx, tagKey, key).Err()
		if err != nil {
			rc.stats.Errors++
			return fmt.Errorf("Redis tag set error: %w", err)
		}

		// Set expiration for tag
		if ttl > 0 {
			rc.client.Expire(rc.ctx, tagKey, ttl)
		}
	}

	return nil
}

// InvalidateByTag invalidates all keys with a specific tag
func (rc *RedisCache) InvalidateByTag(tag string) error {
	tagKey := fmt.Sprintf("tag:%s", tag)

	// Get all keys with this tag
	keys, err := rc.client.SMembers(rc.ctx, tagKey).Result()
	if err != nil {
		rc.stats.Errors++
		return fmt.Errorf("Redis tag members error: %w", err)
	}

	// Delete all keys
	if len(keys) > 0 {
		err = rc.client.Del(rc.ctx, keys...).Err()
		if err != nil {
			rc.stats.Errors++
			return fmt.Errorf("Redis tag delete error: %w", err)
		}
	}

	// Delete the tag key
	err = rc.client.Del(rc.ctx, tagKey).Err()
	if err != nil {
		rc.stats.Errors++
		return fmt.Errorf("Redis tag key delete error: %w", err)
	}

	return nil
}

// Publish publishes a message to a Redis channel
func (rc *RedisCache) Publish(channel string, message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		rc.stats.Errors++
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	err = rc.client.Publish(rc.ctx, channel, data).Err()
	if err != nil {
		rc.stats.Errors++
		return fmt.Errorf("Redis publish error: %w", err)
	}

	return nil
}

// Subscribe subscribes to a Redis channel
func (rc *RedisCache) Subscribe(channels ...string) *redis.PubSub {
	return rc.client.Subscribe(rc.ctx, channels...)
}
