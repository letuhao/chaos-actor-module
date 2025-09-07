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
	fmt.Println("=== RPG Stats + MongoDB Integration Example ===")
	fmt.Println()

	// Connect to MongoDB
	fmt.Println("Connecting to MongoDB at localhost:27017...")
	adapter, err := integration.NewMongoAdapterLocal("rpg_stats_demo")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer adapter.Close()

	fmt.Println("Connected to MongoDB successfully!")
	fmt.Println()

	// Create integration
	integration := integration.NewCoreActorIntegration(adapter)
	ctx := context.Background()

	// Example 1: Create a new character
	fmt.Println("1. Creating a new character...")
	actorID := "player_mongodb_001"

	// Create initial character data
	err = createInitialCharacter(adapter, actorID)
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
	fmt.Printf("Tags: %v\n", contribution.Tags)
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

	// Example 5: Grant title
	fmt.Println("5. Granting title...")

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

	// Example 6: Get stat breakdown
	fmt.Println("6. Getting detailed stat breakdown...")

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

	// Example 7: Save stat registry to database
	fmt.Println("7. Saving stat registry to database...")

	registry := createStatRegistry()
	err = adapter.SaveStatRegistry(ctx, registry)
	if err != nil {
		log.Fatalf("Failed to save stat registry: %v", err)
	}

	fmt.Printf("Saved %d stat definitions to database\n", len(registry))
	fmt.Println()

	// Example 8: Cleanup expired effects
	fmt.Println("8. Cleaning up expired effects...")

	err = adapter.CleanupExpiredEffects(ctx)
	if err != nil {
		log.Fatalf("Failed to cleanup expired effects: %v", err)
	}

	fmt.Println("Expired effects cleaned up successfully")
	fmt.Println()

	fmt.Println("=== MongoDB Integration Example Complete ===")
	fmt.Println("Check your MongoDB database 'rpg_stats_demo' to see the data!")
}

// createInitialCharacter creates initial character data
func createInitialCharacter(adapter *integration.MongoAdapter, actorID string) error {
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

	ctx := context.Background()
	return adapter.SavePlayerProgress(ctx, progress)
}

// createStatRegistry creates a sample stat registry
func createStatRegistry() []model.StatDef {
	return []model.StatDef{
		// Primary stats
		{
			Key:          model.STR,
			Category:     "primary",
			DisplayName:  "Strength",
			Description:  "Physical power and melee damage",
			IsPrimary:    true,
			MinValue:     1,
			MaxValue:     100,
			DefaultValue: 10,
		},
		{
			Key:          model.INT,
			Category:     "primary",
			DisplayName:  "Intelligence",
			Description:  "Magical power and mana",
			IsPrimary:    true,
			MinValue:     1,
			MaxValue:     100,
			DefaultValue: 10,
		},
		{
			Key:          model.WIL,
			Category:     "primary",
			DisplayName:  "Willpower",
			Description:  "Mental fortitude and spell resistance",
			IsPrimary:    true,
			MinValue:     1,
			MaxValue:     100,
			DefaultValue: 10,
		},
		{
			Key:          model.AGI,
			Category:     "primary",
			DisplayName:  "Agility",
			Description:  "Speed and evasion",
			IsPrimary:    true,
			MinValue:     1,
			MaxValue:     100,
			DefaultValue: 10,
		},
		{
			Key:          model.SPD,
			Category:     "primary",
			DisplayName:  "Speed",
			Description:  "Movement and action speed",
			IsPrimary:    true,
			MinValue:     1,
			MaxValue:     100,
			DefaultValue: 10,
		},
		{
			Key:          model.END,
			Category:     "primary",
			DisplayName:  "Endurance",
			Description:  "Health and stamina",
			IsPrimary:    true,
			MinValue:     1,
			MaxValue:     100,
			DefaultValue: 10,
		},
		{
			Key:          model.PER,
			Category:     "primary",
			DisplayName:  "Personality",
			Description:  "Social influence and merchant prices",
			IsPrimary:    true,
			MinValue:     1,
			MaxValue:     100,
			DefaultValue: 10,
		},
		{
			Key:          model.LUK,
			Category:     "primary",
			DisplayName:  "Luck",
			Description:  "Critical hits and random events",
			IsPrimary:    true,
			MinValue:     1,
			MaxValue:     100,
			DefaultValue: 10,
		},
		// Derived stats
		{
			Key:          model.HP_MAX,
			Category:     "derived",
			DisplayName:  "Health Points",
			Description:  "Maximum health points",
			IsPrimary:    false,
			MinValue:     1,
			MaxValue:     10000,
			DefaultValue: 100,
		},
		{
			Key:          model.MANA_MAX,
			Category:     "derived",
			DisplayName:  "Mana Points",
			Description:  "Maximum mana points",
			IsPrimary:    false,
			MinValue:     1,
			MaxValue:     10000,
			DefaultValue: 50,
		},
		{
			Key:          model.ATK,
			Category:     "derived",
			DisplayName:  "Attack Power",
			Description:  "Physical attack damage",
			IsPrimary:    false,
			MinValue:     1,
			MaxValue:     1000,
			DefaultValue: 20,
		},
		{
			Key:          model.MATK,
			Category:     "derived",
			DisplayName:  "Magic Attack",
			Description:  "Magical attack damage",
			IsPrimary:    false,
			MinValue:     1,
			MaxValue:     1000,
			DefaultValue: 15,
		},
		{
			Key:          model.DEF,
			Category:     "derived",
			DisplayName:  "Defense",
			Description:  "Physical damage reduction",
			IsPrimary:    false,
			MinValue:     0,
			MaxValue:     1000,
			DefaultValue: 10,
		},
		{
			Key:          model.EVASION,
			Category:     "derived",
			DisplayName:  "Evasion",
			Description:  "Chance to avoid attacks",
			IsPrimary:    false,
			MinValue:     0,
			MaxValue:     100,
			DefaultValue: 5,
		},
		{
			Key:          model.MOVE_SPEED,
			Category:     "derived",
			DisplayName:  "Movement Speed",
			Description:  "Movement speed multiplier",
			IsPrimary:    false,
			MinValue:     0.1,
			MaxValue:     10,
			DefaultValue: 1.0,
		},
		{
			Key:          model.CRIT_CHANCE,
			Category:     "derived",
			DisplayName:  "Critical Chance",
			Description:  "Chance for critical hits",
			IsPrimary:    false,
			MinValue:     0,
			MaxValue:     1,
			DefaultValue: 0.01,
		},
		{
			Key:          model.CRIT_DAMAGE,
			Category:     "derived",
			DisplayName:  "Critical Damage",
			Description:  "Critical hit damage multiplier",
			IsPrimary:    false,
			MinValue:     1,
			MaxValue:     10,
			DefaultValue: 2.0,
		},
	}
}
