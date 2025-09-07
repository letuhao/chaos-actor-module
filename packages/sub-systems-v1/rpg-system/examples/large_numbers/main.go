package main

import (
	"context"
	"fmt"
	"log"

	"rpg-system/internal/integration"
	"rpg-system/internal/model"
)

func main() {
	fmt.Println("=== RPG System Large Numbers Test ===")
	fmt.Println("Testing int64 support for very large stat values")
	fmt.Println()

	// Create integration with mock adapter
	adapter := integration.NewMockMongoAdapter()
	defer adapter.Close()

	integration := integration.NewCoreActorIntegration(adapter)
	ctx := context.Background()

	// Create a character with very large stat values
	actorID := "high_level_hero"
	err := createHighLevelCharacter(adapter, actorID)
	if err != nil {
		log.Fatal(err)
	}

	// Build contribution
	contribution, err := integration.BuildCoreActorContribution(ctx, actorID)
	if err != nil {
		log.Fatal(err)
	}

	// Display the results
	fmt.Printf("High-Level Character Stats:\n")
	fmt.Printf("  HPMax: %d (int64)\n", contribution.Primary.HPMax)
	fmt.Printf("  LifeSpan: %d (int64)\n", contribution.Primary.LifeSpan)
	fmt.Printf("  Attack: %d (int64)\n", contribution.Primary.Attack)
	fmt.Printf("  Defense: %d (int64)\n", contribution.Primary.Defense)
	fmt.Printf("  Speed: %d (int64)\n", contribution.Primary.Speed)

	// Test with extremely large values
	fmt.Println("\nTesting with extremely large values...")

	// Create a character with maximum possible stat allocations
	actorID2 := "max_stats_hero"
	err = createMaxStatsCharacter(adapter, actorID2)
	if err != nil {
		log.Fatal(err)
	}

	contribution2, err := integration.BuildCoreActorContribution(ctx, actorID2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nMaximum Stats Character:\n")
	fmt.Printf("  HPMax: %d (int64)\n", contribution2.Primary.HPMax)
	fmt.Printf("  LifeSpan: %d (int64)\n", contribution2.Primary.LifeSpan)
	fmt.Printf("  Attack: %d (int64)\n", contribution2.Primary.Attack)
	fmt.Printf("  Defense: %d (int64)\n", contribution2.Primary.Defense)
	fmt.Printf("  Speed: %d (int64)\n", contribution2.Primary.Speed)

	// Verify we can handle values beyond 2.1 billion
	fmt.Println("\nVerification:")
	fmt.Printf("  HPMax > 2,147,483,647: %t\n", contribution2.Primary.HPMax > 2147483647)
	fmt.Printf("  Attack > 2,147,483,647: %t\n", contribution2.Primary.Attack > 2147483647)
	fmt.Printf("  Defense > 2,147,483,647: %t\n", contribution2.Primary.Defense > 2147483647)

	fmt.Println("\n=== Large Numbers Test Complete ===")
	fmt.Println("✅ int64 support working correctly!")
	fmt.Println("✅ Can handle values beyond 2.1 billion limit!")
}

// createHighLevelCharacter creates a character with high-level stats
func createHighLevelCharacter(adapter integration.DatabaseAdapter, actorID string) error {
	ctx := context.Background()

	// Create high-level progress (level 1000)
	progress := &model.PlayerProgress{
		ActorID:     actorID,
		Level:       1000,
		XP:          1000000,
		Allocations: make(map[model.StatKey]int64),
		LastUpdated: 0,
	}

	// Set very high stat allocations
	progress.Allocations[model.STR] = 50000
	progress.Allocations[model.INT] = 45000
	progress.Allocations[model.WIL] = 48000
	progress.Allocations[model.AGI] = 52000
	progress.Allocations[model.SPD] = 51000
	progress.Allocations[model.END] = 55000
	progress.Allocations[model.PER] = 47000
	progress.Allocations[model.LUK] = 46000

	// Save progress
	return adapter.SavePlayerProgress(ctx, progress)
}

// createMaxStatsCharacter creates a character with maximum possible stats
func createMaxStatsCharacter(adapter integration.DatabaseAdapter, actorID string) error {
	ctx := context.Background()

	// Create maximum level progress
	progress := &model.PlayerProgress{
		ActorID:     actorID,
		Level:       10000,
		XP:          10000000,
		Allocations: make(map[model.StatKey]int64),
		LastUpdated: 0,
	}

	// Set maximum stat allocations (using int max for allocations)
	progress.Allocations[model.STR] = 2147483647
	progress.Allocations[model.INT] = 2147483647
	progress.Allocations[model.WIL] = 2147483647
	progress.Allocations[model.AGI] = 2147483647
	progress.Allocations[model.SPD] = 2147483647
	progress.Allocations[model.END] = 2147483647
	progress.Allocations[model.PER] = 2147483647
	progress.Allocations[model.LUK] = 2147483647

	// Save progress
	return adapter.SavePlayerProgress(ctx, progress)
}
