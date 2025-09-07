package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"rpg-system/internal/integration"
	"rpg-system/internal/model"
)

func main() {
	fmt.Println("=== RPG System Simple Usage Example ===")
	fmt.Println()

	// Create integration with mock adapter
	adapter := integration.NewMockMongoAdapter()
	defer adapter.Close()

	integration := integration.NewCoreActorIntegration(adapter)
	ctx := context.Background()

	// Example 1: Character Creation
	fmt.Println("1. Character Creation")
	fmt.Println("====================")

	actorID := "hero_001"
	err := createCharacter(adapter, actorID)
	if err != nil {
		log.Fatal(err)
	}

	// Build initial contribution
	contribution, err := integration.BuildCoreActorContribution(ctx, actorID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Initial Stats: HP=%d, Attack=%d, Defense=%d, Speed=%d\n",
		contribution.Primary.HPMax,
		contribution.Primary.Attack,
		contribution.Primary.Defense,
		contribution.Primary.Speed)

	// Example 2: Equipment System
	fmt.Println("\n2. Equipment System")
	fmt.Println("==================")

	// Equip items
	sword := model.StatModifier{
		Key:   model.STR,
		Op:    model.ADD_FLAT,
		Value: 8.0,
		Source: model.ModifierSourceRef{
			Kind:  "item",
			ID:    "iron_sword",
			Label: "Iron Sword",
		},
	}

	armor := model.StatModifier{
		Key:   model.END,
		Op:    model.ADD_FLAT,
		Value: 6.0,
		Source: model.ModifierSourceRef{
			Kind:  "item",
			ID:    "leather_armor",
			Label: "Leather Armor",
		},
	}

	// Equip items
	err = adapter.EquipItem(ctx, actorID, "weapon", sword)
	if err != nil {
		log.Fatal(err)
	}

	err = adapter.EquipItem(ctx, actorID, "armor", armor)
	if err != nil {
		log.Fatal(err)
	}

	// Rebuild contribution with equipment
	contribution, err = integration.BuildCoreActorContribution(ctx, actorID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("With Equipment: HP=%d, Attack=%d, Defense=%d, Speed=%d\n",
		contribution.Primary.HPMax,
		contribution.Primary.Attack,
		contribution.Primary.Defense,
		contribution.Primary.Speed)

	// Example 3: Effects and Buffs
	fmt.Println("\n3. Effects and Buffs")
	fmt.Println("===================")

	// Add temporary buffs
	strengthBuff := model.StatModifier{
		Key:   model.STR,
		Op:    model.ADD_PCT,
		Value: 0.25, // 25% increase
		Source: model.ModifierSourceRef{
			Kind:  "effect",
			ID:    "strength_potion",
			Label: "Strength Potion",
		},
		Conditions: &model.ModifierConditions{
			DurationMs: 300000, // 5 minutes
		},
	}

	// Add effect
	err = adapter.AddEffect(ctx, actorID, strengthBuff, 5*time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	// Rebuild contribution with effects
	contribution, err = integration.BuildCoreActorContribution(ctx, actorID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("With Buffs: HP=%d, Attack=%d, Defense=%d, Speed=%d\n",
		contribution.Primary.HPMax,
		contribution.Primary.Attack,
		contribution.Primary.Defense,
		contribution.Primary.Speed)

	// Example 4: Stat Breakdown Analysis
	fmt.Println("\n4. Stat Breakdown Analysis")
	fmt.Println("==========================")

	// Get detailed breakdown for STR
	breakdown, err := integration.GetStatBreakdown(ctx, actorID, model.STR)
	if err != nil {
		log.Fatal(err)
	}

	if breakdown != nil {
		fmt.Printf("STR Breakdown:\n")
		fmt.Printf("  Base: %.2f\n", breakdown.Base)
		fmt.Printf("  Flat Modifiers: %.2f\n", breakdown.AdditiveFlat)
		fmt.Printf("  Percentage Modifiers: %.2f\n", breakdown.AdditivePct)
		fmt.Printf("  Multiplicative: %.2f\n", breakdown.Multiplicative)
		if breakdown.CappedTo != nil {
			fmt.Printf("  Capped To: %.2f\n", *breakdown.CappedTo)
		}
		fmt.Printf("  Overrides: %d\n", len(breakdown.Overrides))
		fmt.Printf("  Notes: %d\n", len(breakdown.Notes))
	}

	// Example 5: Performance Testing
	fmt.Println("\n5. Performance Testing")
	fmt.Println("======================")

	// Test multiple character calculations
	start := time.Now()
	for i := 0; i < 1000; i++ {
		_, err := integration.BuildCoreActorContribution(ctx, actorID)
		if err != nil {
			log.Fatal(err)
		}
	}
	duration := time.Since(start)

	fmt.Printf("1000 calculations completed in %v\n", duration)
	fmt.Printf("Average time per calculation: %v\n", duration/1000)

	// Example 6: Core Actor Integration
	fmt.Println("\n6. Core Actor Integration")
	fmt.Println("=========================")

	// Demonstrate Core Actor contribution
	fmt.Printf("Core Actor Primary Stats:\n")
	fmt.Printf("  HPMax: %d\n", contribution.Primary.HPMax)
	fmt.Printf("  LifeSpan: %d\n", contribution.Primary.LifeSpan)
	fmt.Printf("  Attack: %d\n", contribution.Primary.Attack)
	fmt.Printf("  Defense: %d\n", contribution.Primary.Defense)
	fmt.Printf("  Speed: %d\n", contribution.Primary.Speed)

	fmt.Printf("\nCore Actor Flat Modifiers:\n")
	for key, value := range contribution.Flat {
		fmt.Printf("  %s: %.2f\n", key, value)
	}

	fmt.Printf("\nCore Actor Multiplicative Modifiers:\n")
	for key, value := range contribution.Mult {
		fmt.Printf("  %s: %.2f\n", key, value)
	}

	fmt.Printf("\nCore Actor Tags: %v\n", contribution.Tags)

	fmt.Println("\n=== Simple Usage Example Complete ===")
}

// createCharacter creates a new character with initial stats
func createCharacter(adapter integration.DatabaseAdapter, actorID string) error {
	ctx := context.Background()

	// Create initial progress
	progress := &model.PlayerProgress{
		ActorID:     actorID,
		Level:       1,
		XP:          0,
		Allocations: make(map[model.StatKey]int64),
		LastUpdated: time.Now().Unix(),
	}

	// Set initial stat allocations
	progress.Allocations[model.STR] = 12
	progress.Allocations[model.INT] = 10
	progress.Allocations[model.WIL] = 11
	progress.Allocations[model.AGI] = 13
	progress.Allocations[model.SPD] = 14
	progress.Allocations[model.END] = 15
	progress.Allocations[model.PER] = 10
	progress.Allocations[model.LUK] = 9

	// Save progress
	return adapter.SavePlayerProgress(ctx, progress)
}
