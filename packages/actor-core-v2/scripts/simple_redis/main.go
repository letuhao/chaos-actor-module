package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("🚀 Testing Redis (Memurai) Connection...")
	fmt.Println("=====================================")

	// Create Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test connection
	fmt.Println("1. Testing connection...")
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("❌ Failed to connect to Redis: %v", err)
	}
	fmt.Printf("✅ Redis connection successful! Pong: %s\n", pong)

	// Test basic operations
	fmt.Println("\n2. Testing basic operations...")

	// Set a key
	err = client.Set(ctx, "test_key", "test_value", 5*time.Minute).Err()
	if err != nil {
		log.Fatalf("❌ Failed to set key: %v", err)
	}
	fmt.Println("✅ Set key successful")

	// Get a key
	val, err := client.Get(ctx, "test_key").Result()
	if err != nil {
		log.Fatalf("❌ Failed to get key: %v", err)
	}
	fmt.Printf("✅ Get key successful: %s\n", val)

	// Test exists
	exists, err := client.Exists(ctx, "test_key").Result()
	if err != nil {
		log.Fatalf("❌ Failed to check if key exists: %v", err)
	}
	fmt.Printf("✅ Key exists: %d\n", exists)

	// Test TTL
	ttl, err := client.TTL(ctx, "test_key").Result()
	if err != nil {
		log.Fatalf("❌ Failed to get TTL: %v", err)
	}
	fmt.Printf("✅ TTL: %v\n", ttl)

	// Test delete
	err = client.Del(ctx, "test_key").Err()
	if err != nil {
		log.Fatalf("❌ Failed to delete key: %v", err)
	}
	fmt.Println("✅ Delete key successful")

	// Test complex data
	fmt.Println("\n3. Testing complex data...")

	// Test with JSON data
	jsonData := map[string]interface{}{
		"name":   "John Doe",
		"age":    30,
		"active": true,
		"scores": []int{85, 92, 78},
	}

	err = client.Set(ctx, "user:123", jsonData, 5*time.Minute).Err()
	if err != nil {
		log.Fatalf("❌ Failed to set complex data: %v", err)
	}
	fmt.Println("✅ Set complex data successful")

	// Get complex data
	var retrievedData map[string]interface{}
	err = client.Get(ctx, "user:123").Scan(&retrievedData)
	if err != nil {
		log.Fatalf("❌ Failed to get complex data: %v", err)
	}
	fmt.Printf("✅ Get complex data successful: %+v\n", retrievedData)

	// Clean up
	err = client.Del(ctx, "user:123").Err()
	if err != nil {
		log.Fatalf("❌ Failed to delete complex data: %v", err)
	}
	fmt.Println("✅ Clean up successful")

	// Test performance
	fmt.Println("\n4. Testing performance...")
	start := time.Now()

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("perf_test_%d", i)
		value := fmt.Sprintf("value_%d", i)

		// Set
		err = client.Set(ctx, key, value, 5*time.Minute).Err()
		if err != nil {
			log.Fatalf("❌ Failed to set key %s: %v", key, err)
		}

		// Get
		_, err = client.Get(ctx, key).Result()
		if err != nil {
			log.Fatalf("❌ Failed to get key %s: %v", key, err)
		}
	}

	duration := time.Since(start)
	opsPerSecond := float64(2000) / duration.Seconds()

	fmt.Printf("✅ Performance test completed:\n")
	fmt.Printf("   Operations: 2000 (set + get)\n")
	fmt.Printf("   Duration: %v\n", duration)
	fmt.Printf("   Ops/sec: %.2f\n", opsPerSecond)

	// Clean up performance test
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("perf_test_%d", i)
		client.Del(ctx, key)
	}

	// Test Redis info
	fmt.Println("\n5. Testing Redis info...")
	info, err := client.Info(ctx).Result()
	if err != nil {
		log.Fatalf("❌ Failed to get Redis info: %v", err)
	}
	fmt.Printf("✅ Redis info retrieved (length: %d characters)\n", len(info))

	// Test database size
	dbSize, err := client.DBSize(ctx).Result()
	if err != nil {
		log.Fatalf("❌ Failed to get database size: %v", err)
	}
	fmt.Printf("✅ Database size: %d keys\n", dbSize)

	// Close connection
	err = client.Close()
	if err != nil {
		log.Fatalf("❌ Failed to close Redis connection: %v", err)
	}
	fmt.Println("✅ Redis connection closed")

	fmt.Println("\n🎉 All Redis tests completed successfully!")
	fmt.Println("=====================================")
	fmt.Println("Redis (Memurai) is ready for use with Actor Core v2.0!")
}
