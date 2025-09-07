package test

import (
	"context"
	"testing"
	"time"

	"rpg-system/internal/integration"
	"rpg-system/internal/model"
)

func TestMongoAdapter_PlayerProgress(t *testing.T) {
	adapter := integration.NewMockMongoAdapter()
	ctx := context.Background()

	// Test saving player progress
	progress := &model.PlayerProgress{
		ActorID: "test_player",
		Level:   5,
		XP:      1000,
		Allocations: map[model.StatKey]int64{
			model.STR: 15,
			model.INT: 12,
		},
		LastUpdated: time.Now().Unix(),
	}

	err := adapter.SavePlayerProgress(ctx, progress)
	if err != nil {
		t.Errorf("Failed to save player progress: %v", err)
	}

	// Test retrieving player progress
	retrieved, err := adapter.GetPlayerProgress(ctx, "test_player")
	if err != nil {
		t.Errorf("Failed to retrieve player progress: %v", err)
	}

	if retrieved.ActorID != progress.ActorID {
		t.Errorf("Expected actor ID %s, got %s", progress.ActorID, retrieved.ActorID)
	}

	if retrieved.Level != progress.Level {
		t.Errorf("Expected level %d, got %d", progress.Level, retrieved.Level)
	}

	// Test non-existent player
	_, err = adapter.GetPlayerProgress(ctx, "nonexistent")
	if err == nil {
		t.Errorf("Expected error for non-existent player")
	}
}

func TestMongoAdapter_Effects(t *testing.T) {
	adapter := integration.NewMockMongoAdapter()
	ctx := context.Background()

	// Test adding effect
	effect := model.StatModifier{
		Key:   model.STR,
		Op:    model.ADD_FLAT,
		Value: 5.0,
		Source: model.ModifierSourceRef{
			Kind:  "buff",
			ID:    "strength_potion",
			Label: "Strength Potion",
		},
		Priority: 1,
	}

	err := adapter.AddEffect(ctx, "test_player", effect, 5*time.Minute)
	if err != nil {
		t.Errorf("Failed to add effect: %v", err)
	}

	// Test retrieving effects
	effects, err := adapter.GetActiveEffects(ctx, "test_player")
	if err != nil {
		t.Errorf("Failed to retrieve effects: %v", err)
	}

	if len(effects) != 1 {
		t.Errorf("Expected 1 effect, got %d", len(effects))
	}

	if effects[0].Source.ID != "strength_potion" {
		t.Errorf("Expected effect ID 'strength_potion', got '%s'", effects[0].Source.ID)
	}

	// Test removing effect
	err = adapter.RemoveEffect(ctx, "test_player", "strength_potion")
	if err != nil {
		t.Errorf("Failed to remove effect: %v", err)
	}

	effects, err = adapter.GetActiveEffects(ctx, "test_player")
	if err != nil {
		t.Errorf("Failed to retrieve effects after removal: %v", err)
	}

	if len(effects) != 0 {
		t.Errorf("Expected 0 effects after removal, got %d", len(effects))
	}
}

func TestMongoAdapter_Equipment(t *testing.T) {
	adapter := integration.NewMockMongoAdapter()
	ctx := context.Background()

	// Test equipping item
	item := model.StatModifier{
		Key:   model.ATK,
		Op:    model.ADD_FLAT,
		Value: 10.0,
		Source: model.ModifierSourceRef{
			Kind:  "item",
			ID:    "iron_sword",
			Label: "Iron Sword",
		},
		Priority: 1,
	}

	err := adapter.EquipItem(ctx, "test_player", "weapon", item)
	if err != nil {
		t.Errorf("Failed to equip item: %v", err)
	}

	// Test retrieving equipment
	equipment, err := adapter.GetEquippedItems(ctx, "test_player")
	if err != nil {
		t.Errorf("Failed to retrieve equipment: %v", err)
	}

	if len(equipment) != 1 {
		t.Errorf("Expected 1 equipment item, got %d", len(equipment))
	}

	if equipment[0].Source.ID != "iron_sword" {
		t.Errorf("Expected item ID 'iron_sword', got '%s'", equipment[0].Source.ID)
	}

	// Test unequipping item
	err = adapter.UnequipItem(ctx, "test_player", "weapon")
	if err != nil {
		t.Errorf("Failed to unequip item: %v", err)
	}

	equipment, err = adapter.GetEquippedItems(ctx, "test_player")
	if err != nil {
		t.Errorf("Failed to retrieve equipment after unequipping: %v", err)
	}

	if len(equipment) != 0 {
		t.Errorf("Expected 0 equipment items after unequipping, got %d", len(equipment))
	}
}

func TestMongoAdapter_Titles(t *testing.T) {
	adapter := integration.NewMockMongoAdapter()
	ctx := context.Background()

	// Test granting title
	title := model.StatModifier{
		Key:   model.PER,
		Op:    model.ADD_FLAT,
		Value: 3.0,
		Source: model.ModifierSourceRef{
			Kind:  "title",
			ID:    "noble_warrior",
			Label: "Noble Warrior",
		},
		Priority: 1,
	}

	err := adapter.GrantTitle(ctx, "test_player", "noble_warrior", title)
	if err != nil {
		t.Errorf("Failed to grant title: %v", err)
	}

	// Test retrieving titles
	titles, err := adapter.GetOwnedTitles(ctx, "test_player")
	if err != nil {
		t.Errorf("Failed to retrieve titles: %v", err)
	}

	if len(titles) != 1 {
		t.Errorf("Expected 1 title, got %d", len(titles))
	}

	if titles[0].Source.ID != "noble_warrior" {
		t.Errorf("Expected title ID 'noble_warrior', got '%s'", titles[0].Source.ID)
	}
}

func TestMongoAdapter_StatRegistry(t *testing.T) {
	adapter := integration.NewMockMongoAdapter()
	ctx := context.Background()

	// Test saving stat registry
	registry := []model.StatDef{
		{
			Key:          model.STR,
			Category:     "primary",
			DisplayName:  "Strength",
			Description:  "Physical power",
			IsPrimary:    true,
			MinValue:     1,
			MaxValue:     100,
			DefaultValue: 10,
		},
		{
			Key:          model.HP_MAX,
			Category:     "derived",
			DisplayName:  "Health Points",
			Description:  "Maximum health",
			IsPrimary:    false,
			MinValue:     1,
			MaxValue:     10000,
			DefaultValue: 100,
		},
	}

	err := adapter.SaveStatRegistry(ctx, registry)
	if err != nil {
		t.Errorf("Failed to save stat registry: %v", err)
	}

	// Test retrieving stat registry
	retrieved, err := adapter.GetStatRegistry(ctx)
	if err != nil {
		t.Errorf("Failed to retrieve stat registry: %v", err)
	}

	if len(retrieved) != 2 {
		t.Errorf("Expected 2 stat definitions, got %d", len(retrieved))
	}

	if retrieved[0].Key != model.STR {
		t.Errorf("Expected first stat to be STR, got %s", retrieved[0].Key)
	}
}

func TestMongoAdapter_PlayerStatsSummary(t *testing.T) {
	adapter := integration.NewMockMongoAdapter()
	ctx := context.Background()

	// Set up test data
	progress := &model.PlayerProgress{
		ActorID: "test_player",
		Level:   5,
		XP:      1000,
		Allocations: map[model.StatKey]int64{
			model.STR: 15,
			model.INT: 12,
		},
		LastUpdated: time.Now().Unix(),
	}
	adapter.SavePlayerProgress(ctx, progress)

	effect := model.StatModifier{
		Key:   model.STR,
		Op:    model.ADD_FLAT,
		Value: 5.0,
		Source: model.ModifierSourceRef{
			Kind:  "buff",
			ID:    "strength_potion",
			Label: "Strength Potion",
		},
		Priority: 1,
	}
	adapter.AddEffect(ctx, "test_player", effect, 5*time.Minute)

	item := model.StatModifier{
		Key:   model.ATK,
		Op:    model.ADD_FLAT,
		Value: 10.0,
		Source: model.ModifierSourceRef{
			Kind:  "item",
			ID:    "iron_sword",
			Label: "Iron Sword",
		},
		Priority: 1,
	}
	adapter.EquipItem(ctx, "test_player", "weapon", item)

	// Test getting player stats summary
	summary, err := adapter.GetPlayerStatsSummary(ctx, "test_player")
	if err != nil {
		t.Errorf("Failed to get player stats summary: %v", err)
	}

	if summary.Progress.ActorID != "test_player" {
		t.Errorf("Expected actor ID 'test_player', got '%s'", summary.Progress.ActorID)
	}

	if len(summary.Effects) != 1 {
		t.Errorf("Expected 1 effect, got %d", len(summary.Effects))
	}

	if len(summary.Equipment) != 1 {
		t.Errorf("Expected 1 equipment item, got %d", len(summary.Equipment))
	}
}
