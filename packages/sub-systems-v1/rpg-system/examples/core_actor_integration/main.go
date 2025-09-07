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
	fmt.Println("=== RPG Stats + Core Actor Integration Example ===")
	fmt.Println()

	// In a real application, you would connect to MongoDB
	// For this example, we'll use a mock adapter
	adapter := createMockAdapter()
	defer adapter.Close()

	// Create integration
	integration := integration.NewCoreActorIntegration(adapter)

	ctx := context.Background()

	// Example 1: Create a new character
	fmt.Println("1. Creating a new character...")
	actorID := "player_demo_001"

	// Create initial character data
	err := createInitialCharacter(adapter, actorID)
	if err != nil {
		log.Fatalf("Failed to create character: %v", err)
	}

	// Build Core Actor contribution
	contribution, err := integration.BuildCoreActorContribution(ctx, actorID)
	if err != nil {
		log.Fatalf("Failed to build contribution: %v", err)
	}

	fmt.Printf("Character created: %s\n", actorID)
	fmt.Printf("Primary Stats: HP=%d, Attack=%d, Defense=%d, Speed=%d\n",
		contribution.Primary.HPMax, contribution.Primary.Attack,
		contribution.Primary.Defense, contribution.Primary.Speed)
	fmt.Println()

	// Example 2: Level up and allocate stats
	fmt.Println("2. Leveling up and allocating stats...")

	updates := &integration.PlayerStatsUpdates{
		XPGranted:     1000,
		LevelIncrease: 2,
		StatAllocations: map[model.StatKey]int64{
			model.STR: 3,
			model.INT: 2,
			model.END: 2,
		},
	}

	contribution, err = integration.UpdatePlayerStats(ctx, actorID, updates)
	if err != nil {
		log.Fatalf("Failed to update stats: %v", err)
	}

	fmt.Printf("After level up: HP=%d, Attack=%d, Defense=%d, Speed=%d\n",
		contribution.Primary.HPMax, contribution.Primary.Attack,
		contribution.Primary.Defense, contribution.Primary.Speed)
	fmt.Printf("Tags: %v\n", contribution.Tags)
	fmt.Println()

	// Example 3: Equip items
	fmt.Println("3. Equipping items...")

	// Create some equipment
	sword := model.StatModifier{
		Key:   model.STR,
		Op:    model.ADD_FLAT,
		Value: 5.0,
		Source: model.ModifierSourceRef{
			Kind:  "item",
			ID:    "iron_sword",
			Label: "Iron Sword",
		},
		Priority: 1,
	}

	armor := model.StatModifier{
		Key:   model.END,
		Op:    model.ADD_FLAT,
		Value: 3.0,
		Source: model.ModifierSourceRef{
			Kind:  "item",
			ID:    "leather_armor",
			Label: "Leather Armor",
		},
		Priority: 1,
	}

	updates = &integration.PlayerStatsUpdates{
		ItemsEquipped: map[string]model.StatModifier{
			"weapon": sword,
			"armor":  armor,
		},
	}

	contribution, err = integration.UpdatePlayerStats(ctx, actorID, updates)
	if err != nil {
		log.Fatalf("Failed to equip items: %v", err)
	}

	fmt.Printf("After equipping items: HP=%d, Attack=%d, Defense=%d\n",
		contribution.Primary.HPMax, contribution.Primary.Attack,
		contribution.Primary.Defense)
	fmt.Printf("Flat modifiers: ATK=%.2f, DEF=%.2f\n",
		contribution.Flat["ATK"], contribution.Flat["DEF"])
	fmt.Println()

	// Example 4: Apply buffs
	fmt.Println("4. Applying buffs...")

	strengthBuff := model.StatModifier{
		Key:   model.STR,
		Op:    model.MULTIPLY,
		Value: 1.2, // 20% increase
		Source: model.ModifierSourceRef{
			Kind:  "buff",
			ID:    "strength_potion",
			Label: "Strength Potion",
		},
		Priority: 2,
		Conditions: &model.ModifierConditions{
			DurationMs: 300000, // 5 minutes
		},
	}

	updates = &integration.PlayerStatsUpdates{
		EffectsAdded: []model.StatModifier{strengthBuff},
	}

	contribution, err = integration.UpdatePlayerStats(ctx, actorID, updates)
	if err != nil {
		log.Fatalf("Failed to apply buff: %v", err)
	}

	fmt.Printf("After applying buff: HP=%d, Attack=%d\n",
		contribution.Primary.HPMax, contribution.Primary.Attack)
	fmt.Printf("Multiplicative modifiers: CritChance=%.2f\n",
		contribution.Mult["CritChance"])
	fmt.Println()

	// Example 5: Get stat breakdown
	fmt.Println("5. Getting detailed stat breakdown...")

	breakdown, err := integration.GetStatBreakdown(ctx, actorID, model.STR)
	if err != nil {
		log.Fatalf("Failed to get breakdown: %v", err)
	}

	if breakdown != nil {
		fmt.Printf("STR Breakdown:\n")
		fmt.Printf("  Base: %.2f\n", breakdown.Base)
		fmt.Printf("  Additive Flat: %.2f\n", breakdown.AdditiveFlat)
		fmt.Printf("  Additive %: %.2f\n", breakdown.AdditivePct)
		fmt.Printf("  Multiplicative: %.2f\n", breakdown.Multiplicative)
		if breakdown.CappedTo != nil {
			fmt.Printf("  Capped to: %.2f\n", *breakdown.CappedTo)
		}
	}
	fmt.Println()

	// Example 6: Grant title
	fmt.Println("6. Granting title...")

	title := model.StatModifier{
		Key:   model.PER,
		Op:    model.ADD_FLAT,
		Value: 2.0,
		Source: model.ModifierSourceRef{
			Kind:  "title",
			ID:    "noble_warrior",
			Label: "Noble Warrior",
		},
		Priority: 1,
	}

	updates = &integration.PlayerStatsUpdates{
		TitlesGranted: map[string]model.StatModifier{
			"noble_warrior": title,
		},
	}

	contribution, err = integration.UpdatePlayerStats(ctx, actorID, updates)
	if err != nil {
		log.Fatalf("Failed to grant title: %v", err)
	}

	fmt.Printf("After granting title: Tags=%v\n", contribution.Tags)
	fmt.Printf("Social amplifier: %.2f\n", contribution.Flat["amplifiers.social"])
	fmt.Println()

	fmt.Println("=== Integration Example Complete ===")
}

// createMockAdapter creates a mock adapter for demonstration
func createMockAdapter() integration.DatabaseAdapter {
	// In a real application, this would connect to MongoDB
	// For this example, we'll use the mock adapter
	return integration.NewMockMongoAdapter()
}

// createInitialCharacter creates initial character data
func createInitialCharacter(adapter integration.DatabaseAdapter, actorID string) error {
	// Create initial progress
	progress := &model.PlayerProgress{
		ActorID: actorID,
		Level:   1,
		XP:      0,
		Allocations: map[model.StatKey]int64{
			model.STR: 15,
			model.INT: 12,
			model.WIL: 10,
			model.AGI: 14,
			model.SPD: 13,
			model.END: 16,
			model.PER: 8,
			model.LUK: 11,
		},
		LastUpdated: time.Now().Unix(),
	}

	// In a real application, this would save to the database
	// For this example, we'll just return success
	_ = progress
	return nil
}
