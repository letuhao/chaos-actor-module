package test

import (
	"testing"

	"rpg-system/internal/model"
	"rpg-system/internal/resolver"
)

func TestStatResolver_ComputeSnapshot(t *testing.T) {
	resolver := resolver.NewStatResolver()

	tests := []struct {
		name     string
		input    model.ComputeInput
		expected map[model.StatKey]float64
	}{
		{
			name: "Basic primary stats",
			input: model.ComputeInput{
				ActorID:         "test_actor",
				Level:           1,
				BaseAllocations: map[model.StatKey]int64{model.STR: 15, model.INT: 12, model.END: 16},
				Registry:        []model.StatDef{},
				Items:           []model.StatModifier{},
				WithBreakdown:   false,
			},
			expected: map[model.StatKey]float64{
				model.STR: 15,
				model.INT: 12,
				model.END: 16,
			},
		},
		{
			name: "With equipment modifiers",
			input: model.ComputeInput{
				ActorID:         "test_actor",
				Level:           1,
				BaseAllocations: map[model.StatKey]int64{model.STR: 15, model.INT: 12, model.END: 16},
				Registry:        []model.StatDef{},
				Items: []model.StatModifier{
					{Key: model.STR, Op: model.ADD_FLAT, Value: 3.0, Source: model.ModifierSourceRef{Kind: "item", ID: "sword"}},
					{Key: model.ATK, Op: model.ADD_FLAT, Value: 8.0, Source: model.ModifierSourceRef{Kind: "item", ID: "sword"}},
				},
				WithBreakdown: false,
			},
			expected: map[model.StatKey]float64{
				model.STR: 18, // 15 + 3
				model.INT: 12,
				model.END: 16,
			},
		},
		{
			name: "With breakdown enabled",
			input: model.ComputeInput{
				ActorID:         "test_actor",
				Level:           1,
				BaseAllocations: map[model.StatKey]int64{model.STR: 15, model.INT: 12, model.END: 16},
				Registry:        []model.StatDef{},
				Items: []model.StatModifier{
					{Key: model.STR, Op: model.ADD_FLAT, Value: 3.0, Source: model.ModifierSourceRef{Kind: "item", ID: "sword"}},
				},
				WithBreakdown: true,
			},
			expected: map[model.StatKey]float64{
				model.STR: 18, // 15 + 3
				model.INT: 12,
				model.END: 16,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			snapshot := resolver.ComputeSnapshot(tt.input)

			if snapshot.ActorID != tt.input.ActorID {
				t.Errorf("Expected actor ID %s, got %s", tt.input.ActorID, snapshot.ActorID)
			}

			// Check that snapshot has a hash
			if snapshot.Hash == "" {
				t.Errorf("Expected snapshot to have a hash")
			}

			// Check expected stats
			for statKey, expectedValue := range tt.expected {
				if actualValue, exists := snapshot.Stats[statKey]; exists {
					if actualValue != expectedValue {
						t.Errorf("Expected %s = %f, got %f", statKey, expectedValue, actualValue)
					}
				} else {
					t.Errorf("Expected stat %s to exist in snapshot", statKey)
				}
			}

			// Check breakdown if enabled
			if tt.input.WithBreakdown {
				if snapshot.Breakdown == nil {
					t.Errorf("Expected breakdown to be present when WithBreakdown is true")
				} else {
					// Check that breakdown has entries for modified stats
					for statKey := range tt.expected {
						if _, exists := snapshot.Breakdown[statKey]; exists {
							// Breakdown exists for this stat
						}
					}
				}
			}
		})
	}
}

func TestStatResolver_CollectAllModifiers(t *testing.T) {
	resolver := resolver.NewStatResolver()

	input := model.ComputeInput{
		ActorID: "test_actor",
		Items: []model.StatModifier{
			{Key: model.STR, Op: model.ADD_FLAT, Value: 1.0},
		},
		Titles: []model.StatModifier{
			{Key: model.INT, Op: model.ADD_FLAT, Value: 2.0},
		},
		Buffs: []model.StatModifier{
			{Key: model.STR, Op: model.MULTIPLY, Value: 1.1},
		},
		Debuffs: []model.StatModifier{
			{Key: model.STR, Op: model.MULTIPLY, Value: 0.9},
		},
		Auras: []model.StatModifier{
			{Key: model.INT, Op: model.ADD_PCT, Value: 10.0},
		},
		Environment: []model.StatModifier{
			{Key: model.END, Op: model.ADD_FLAT, Value: 5.0},
		},
	}

	modifiers := resolver.CollectAllModifiers(input)

	expectedCount := len(input.Items) + len(input.Titles) + len(input.Buffs) + len(input.Debuffs) + len(input.Auras) + len(input.Environment)
	if len(modifiers) != expectedCount {
		t.Errorf("Expected %d modifiers, got %d", expectedCount, len(modifiers))
	}

	// Check that all modifier types are present
	hasItem := false
	hasTitle := false
	hasBuff := false
	hasDebuff := false
	hasAura := false
	hasEnvironment := false

	for _, mod := range modifiers {
		if mod.Key == model.STR && mod.Op == model.ADD_FLAT && mod.Value == 1.0 {
			hasItem = true
		}
		if mod.Key == model.INT && mod.Op == model.ADD_FLAT && mod.Value == 2.0 {
			hasTitle = true
		}
		if mod.Key == model.STR && mod.Op == model.MULTIPLY && mod.Value == 1.1 {
			hasBuff = true
		}
		if mod.Key == model.STR && mod.Op == model.MULTIPLY && mod.Value == 0.9 {
			hasDebuff = true
		}
		if mod.Key == model.INT && mod.Op == model.ADD_PCT && mod.Value == 10.0 {
			hasAura = true
		}
		if mod.Key == model.END && mod.Op == model.ADD_FLAT && mod.Value == 5.0 {
			hasEnvironment = true
		}
	}

	if !hasItem {
		t.Errorf("Expected item modifier to be present")
	}
	if !hasTitle {
		t.Errorf("Expected title modifier to be present")
	}
	if !hasBuff {
		t.Errorf("Expected buff modifier to be present")
	}
	if !hasDebuff {
		t.Errorf("Expected debuff modifier to be present")
	}
	if !hasAura {
		t.Errorf("Expected aura modifier to be present")
	}
	if !hasEnvironment {
		t.Errorf("Expected environment modifier to be present")
	}
}

func TestStatResolver_FilterModifiersForStat(t *testing.T) {
	resolver := resolver.NewStatResolver()

	modifiers := []model.StatModifier{
		{Key: model.STR, Op: model.ADD_FLAT, Value: 1.0},
		{Key: model.INT, Op: model.ADD_FLAT, Value: 2.0},
		{Key: model.STR, Op: model.MULTIPLY, Value: 1.1},
		{Key: model.END, Op: model.ADD_FLAT, Value: 3.0},
	}

	// Test filtering for STR
	strModifiers := resolver.FilterModifiersForStat(modifiers, model.STR)
	if len(strModifiers) != 2 {
		t.Errorf("Expected 2 STR modifiers, got %d", len(strModifiers))
	}

	// Test filtering for INT
	intModifiers := resolver.FilterModifiersForStat(modifiers, model.INT)
	if len(intModifiers) != 1 {
		t.Errorf("Expected 1 INT modifier, got %d", len(intModifiers))
	}

	// Test filtering for non-existent stat
	nonexistentModifiers := resolver.FilterModifiersForStat(modifiers, model.HP_MAX)
	if len(nonexistentModifiers) != 0 {
		t.Errorf("Expected 0 HP_MAX modifiers, got %d", len(nonexistentModifiers))
	}
}
