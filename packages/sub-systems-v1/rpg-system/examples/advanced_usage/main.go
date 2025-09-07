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
	fmt.Println("=== RPG System Advanced Usage Example ===")
	fmt.Println()

	// Create integration with mock adapter
	adapter := integration.NewMockMongoAdapter()
	defer adapter.Close()

	integration := integration.NewCoreActorIntegration(adapter)
	ctx := context.Background()

	// Example 1: Character Creation and Progression
	fmt.Println("1. Character Creation and Progression")
	fmt.Println("=====================================")

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

	// Level up the character
	fmt.Println("\nLeveling up character...")
	updates := &integration.PlayerStatsUpdates{
		XPGranted: 2500, // Enough for several levels
		StatAllocations: map[model.StatKey]int64{
			model.STR: 5,
			model.END: 3,
			model.AGI: 2,
		},
	}

	contribution, err = integration.UpdatePlayerStats(ctx, actorID, updates)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("After Leveling: HP=%d, Attack=%d, Defense=%d, Speed=%d\n",
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

	enduranceBuff := model.StatModifier{
		Key:   model.END,
		Op:    model.ADD_PCT,
		Value: 0.20, // 20% increase
		Source: model.ModifierSourceRef{
			Kind:  "effect",
			ID:    "endurance_elixir",
			Label: "Endurance Elixir",
		},
		Conditions: &model.ModifierConditions{
			DurationMs: 600000, // 10 minutes
		},
	}

	// Add effects
	err = adapter.AddEffect(ctx, actorID, strengthBuff, 5*time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	err = adapter.AddEffect(ctx, actorID, enduranceBuff, 10*time.Minute)
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

	// Example 4: Titles and Achievements
	fmt.Println("\n4. Titles and Achievements")
	fmt.Println("==========================")

	// Grant titles
	warriorTitle := model.StatModifier{
		Key:   model.ATK,
		Op:    model.ADD_PCT,
		Value: 0.15, // 15% attack bonus
		Source: model.ModifierSourceRef{
			Kind:  "title",
			ID:    "warrior",
			Label: "Warrior",
		},
	}

	guardianTitle := model.StatModifier{
		Key:   model.DEF,
		Op:    model.ADD_PCT,
		Value: 0.10, // 10% defense bonus
		Source: model.ModifierSourceRef{
			Kind:  "title",
			ID:    "guardian",
			Label: "Guardian",
		},
	}

	// Grant titles
	err = adapter.GrantTitle(ctx, actorID, "warrior", warriorTitle)
	if err != nil {
		log.Fatal(err)
	}

	err = adapter.GrantTitle(ctx, actorID, "guardian", guardianTitle)
	if err != nil {
		log.Fatal(err)
	}

	// Rebuild contribution with titles
	contribution, err = integration.BuildCoreActorContribution(ctx, actorID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("With Titles: HP=%d, Attack=%d, Defense=%d, Speed=%d\n",
		contribution.Primary.HPMax,
		contribution.Primary.Attack,
		contribution.Primary.Defense,
		contribution.Primary.Speed)

	// Example 5: Stat Breakdown Analysis
	fmt.Println("\n5. Stat Breakdown Analysis")
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

	// Example 6: Performance Testing
	fmt.Println("\n6. Performance Testing")
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

	// Example 7: Core Actor Integration
	fmt.Println("\n7. Core Actor Integration")
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

	fmt.Println("\n=== Advanced Usage Example Complete ===")
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
