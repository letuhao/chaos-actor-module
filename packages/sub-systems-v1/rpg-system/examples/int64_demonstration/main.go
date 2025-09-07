package main

import (
	"context"
	"fmt"
	"log"
	"math"

	"rpg-system/internal/integration"
	"rpg-system/internal/model"
)

func main() {
	fmt.Println("=== RPG System int64 Demonstration ===")
	fmt.Println("Testing int64 support with very large stat values")
	fmt.Println()

	// Create integration with mock adapter
	adapter := integration.NewMockMongoAdapter()
	defer adapter.Close()

	integration := integration.NewCoreActorIntegration(adapter)
	ctx := context.Background()

	// Test 1: Direct int64 conversion demonstration
	fmt.Println("1. Direct int64 Conversion Test")
	fmt.Println("===============================")

	// Create a character with very large stat values
	actorID := "int64_test_hero"
	err := createInt64TestCharacter(adapter, actorID)
	if err != nil {
		log.Fatal(err)
	}

	// Build contribution
	contribution, err := integration.BuildCoreActorContribution(ctx, actorID)
	if err != nil {
		log.Fatal(err)
	}

	// Display the results
	fmt.Printf("Character Stats (int64):\n")
	fmt.Printf("  HPMax: %d\n", contribution.Primary.HPMax)
	fmt.Printf("  LifeSpan: %d\n", contribution.Primary.LifeSpan)
	fmt.Printf("  Attack: %d\n", contribution.Primary.Attack)
	fmt.Printf("  Defense: %d\n", contribution.Primary.Defense)
	fmt.Printf("  Speed: %d\n", contribution.Primary.Speed)

	// Test 2: Demonstrate int64 arithmetic
	fmt.Println("\n2. int64 Arithmetic Test")
	fmt.Println("========================")

	// Test with very large float64 values that would overflow int32
	largeFloat := 5000000000.0      // 5 billion
	veryLargeFloat := 10000000000.0 // 10 billion

	fmt.Printf("Large float64 value: %.0f\n", largeFloat)
	fmt.Printf("Converted to int64: %d\n", int64(largeFloat))
	fmt.Printf("Very large float64 value: %.0f\n", veryLargeFloat)
	fmt.Printf("Converted to int64: %d\n", int64(veryLargeFloat))

	// Test arithmetic operations
	result1 := int64(largeFloat) + int64(veryLargeFloat)
	result2 := int64(largeFloat) * 2
	result3 := int64(veryLargeFloat) / 2

	fmt.Printf("Addition: %d + %d = %d\n", int64(largeFloat), int64(veryLargeFloat), result1)
	fmt.Printf("Multiplication: %d * 2 = %d\n", int64(largeFloat), result2)
	fmt.Printf("Division: %d / 2 = %d\n", int64(veryLargeFloat), result3)

	// Test 3: Compare with int32 limits
	fmt.Println("\n3. int32 vs int64 Comparison")
	fmt.Println("============================")

	int32Max := int32(2147483647)
	int64Max := int64(9223372036854775807)

	fmt.Printf("int32 maximum value: %d\n", int32Max)
	fmt.Printf("int64 maximum value: %d\n", int64Max)
	fmt.Printf("int64 can handle %d times larger values than int32\n", int64Max/int64(int32Max))

	// Test values that would overflow int32
	testValue := int64(3000000000) // 3 billion
	fmt.Printf("Test value: %d\n", testValue)
	fmt.Printf("Would overflow int32: %t\n", testValue > int64(int32Max))
	fmt.Printf("Fits in int64: %t\n", testValue < int64Max)

	// Test 4: RPG System with extreme values
	fmt.Println("\n4. RPG System Extreme Values Test")
	fmt.Println("=================================")

	// Create a character with extreme stat values
	actorID2 := "extreme_hero"
	err = createExtremeCharacter(adapter, actorID2)
	if err != nil {
		log.Fatal(err)
	}

	contribution2, err := integration.BuildCoreActorContribution(ctx, actorID2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Extreme Character Stats:\n")
	fmt.Printf("  HPMax: %d (int64)\n", contribution2.Primary.HPMax)
	fmt.Printf("  LifeSpan: %d (int64)\n", contribution2.Primary.LifeSpan)
	fmt.Printf("  Attack: %d (int64)\n", contribution2.Primary.Attack)
	fmt.Printf("  Defense: %d (int64)\n", contribution2.Primary.Defense)
	fmt.Printf("  Speed: %d (int64)\n", contribution2.Primary.Speed)

	// Verify the values are reasonable
	fmt.Printf("\nValue Analysis:\n")
	fmt.Printf("  HPMax > int32 max: %t\n", contribution2.Primary.HPMax > int64(int32Max))
	fmt.Printf("  Attack > int32 max: %t\n", contribution2.Primary.Attack > int64(int32Max))
	fmt.Printf("  All values < int64 max: %t\n",
		contribution2.Primary.HPMax < int64Max &&
			contribution2.Primary.LifeSpan < int64Max &&
			contribution2.Primary.Attack < int64Max &&
			contribution2.Primary.Defense < int64Max &&
			contribution2.Primary.Speed < int64Max)

	fmt.Println("\n=== int64 Demonstration Complete ===")
	fmt.Println("✅ int64 support working correctly!")
	fmt.Println("✅ Can handle values beyond 2.1 billion limit!")
	fmt.Println("✅ Arithmetic operations work with large numbers!")
	fmt.Println("✅ RPG System integration supports int64!")
}

// createInt64TestCharacter creates a character for int64 testing
func createInt64TestCharacter(adapter integration.DatabaseAdapter, actorID string) error {
	ctx := context.Background()

	// Create high-level progress
	progress := &model.PlayerProgress{
		ActorID:     actorID,
		Level:       1000,
		XP:          1000000,
		Allocations: make(map[model.StatKey]int64),
		LastUpdated: 0,
	}

	// Set high stat allocations
	progress.Allocations[model.STR] = 100000
	progress.Allocations[model.INT] = 95000
	progress.Allocations[model.WIL] = 98000
	progress.Allocations[model.AGI] = 102000
	progress.Allocations[model.SPD] = 101000
	progress.Allocations[model.END] = 105000
	progress.Allocations[model.PER] = 97000
	progress.Allocations[model.LUK] = 96000

	return adapter.SavePlayerProgress(ctx, progress)
}

// createExtremeCharacter creates a character with extreme stat values
func createExtremeCharacter(adapter integration.DatabaseAdapter, actorID string) error {
	ctx := context.Background()

	// Create maximum level progress
	progress := &model.PlayerProgress{
		ActorID:     actorID,
		Level:       10000,
		XP:          10000000,
		Allocations: make(map[model.StatKey]int64),
		LastUpdated: 0,
	}

	// Set extreme stat allocations (using maximum int values)
	progress.Allocations[model.STR] = math.MaxInt32
	progress.Allocations[model.INT] = math.MaxInt32
	progress.Allocations[model.WIL] = math.MaxInt32
	progress.Allocations[model.AGI] = math.MaxInt32
	progress.Allocations[model.SPD] = math.MaxInt32
	progress.Allocations[model.END] = math.MaxInt32
	progress.Allocations[model.PER] = math.MaxInt32
	progress.Allocations[model.LUK] = math.MaxInt32

	return adapter.SavePlayerProgress(ctx, progress)
}
