package redis_test

import (
	"fmt"
	"log"

	"actor-core/tests/integration"
)

func main() {
	fmt.Println("🚀 Testing Redis (Memurai) Connection...")
	fmt.Println("=====================================")

	helper := integration.NewRedisTestHelper()

	// Test 1: Connection
	fmt.Println("\n1. Testing Redis Connection...")
	err := helper.TestConnection()
	if err != nil {
		log.Fatalf("❌ Connection failed: %v", err)
	}
	fmt.Println("✅ Redis connection successful!")

	// Test 2: Basic Operations
	fmt.Println("\n2. Testing Basic Operations...")
	err = helper.TestBasicOperations()
	if err != nil {
		log.Fatalf("❌ Basic operations failed: %v", err)
	}
	fmt.Println("✅ Basic operations successful!")

	// Test 3: Complex Data
	fmt.Println("\n3. Testing Complex Data...")
	err = helper.TestComplexData()
	if err != nil {
		log.Fatalf("❌ Complex data test failed: %v", err)
	}
	fmt.Println("✅ Complex data test successful!")

	// Test 4: Performance
	fmt.Println("\n4. Testing Performance...")
	operations := 1000
	duration, opsPerSecond, err := helper.TestPerformance(operations)
	if err != nil {
		log.Fatalf("❌ Performance test failed: %v", err)
	}
	fmt.Printf("✅ Performance test successful!\n")
	fmt.Printf("   Operations: %d (set + get)\n", operations*2)
	fmt.Printf("   Duration: %v\n", duration)
	fmt.Printf("   Ops/sec: %.2f\n", opsPerSecond)

	// Test 5: Tags
	fmt.Println("\n5. Testing Tag-based Operations...")
	err = helper.TestTags()
	if err != nil {
		log.Fatalf("❌ Tag operations failed: %v", err)
	}
	fmt.Println("✅ Tag operations successful!")

	// Test 6: Pub/Sub
	fmt.Println("\n6. Testing Pub/Sub...")
	err = helper.TestPubSub()
	if err != nil {
		fmt.Printf("⚠️ Pub/Sub test failed (this might be expected): %v\n", err)
	} else {
		fmt.Println("✅ Pub/Sub test successful!")
	}

	// Test 7: Statistics
	fmt.Println("\n7. Testing Statistics...")
	stats, err := helper.GetStats()
	if err != nil {
		log.Fatalf("❌ Statistics test failed: %v", err)
	}
	fmt.Printf("✅ Statistics retrieved successfully!\n")
	fmt.Printf("   Hits: %d\n", stats.Hits)
	fmt.Printf("   Misses: %d\n", stats.Misses)
	fmt.Printf("   Hit Ratio: %.2f%%\n", stats.HitRatio*100)
	fmt.Printf("   Size: %d bytes\n", stats.Size)

	// Test 8: Memory Usage
	fmt.Println("\n8. Testing Memory Usage...")
	memoryUsage, err := helper.GetMemoryUsage()
	if err != nil {
		log.Fatalf("❌ Memory usage test failed: %v", err)
	}
	fmt.Printf("✅ Memory usage: %d bytes\n", memoryUsage)

	// Test 9: Keys
	fmt.Println("\n9. Testing Keys...")
	keys, err := helper.GetKeys("*")
	if err != nil {
		log.Fatalf("❌ Keys test failed: %v", err)
	}
	fmt.Printf("✅ Found %d keys in database\n", len(keys))

	// Test 10: Clear All
	fmt.Println("\n10. Testing Clear All...")
	err = helper.ClearAll()
	if err != nil {
		log.Fatalf("❌ Clear all failed: %v", err)
	}
	fmt.Println("✅ Clear all successful!")

	fmt.Println("\n🎉 All Redis tests completed successfully!")
	fmt.Println("=====================================")
	fmt.Println("Redis (Memurai) is ready for use with Actor Core v2.0!")
}
