// Package test contains tests for the RPG Stats Sub-System
package test

import (
	"testing"

	"rpg-stats/internal/model"
	"rpg-stats/internal/rules"
)

func TestStackingEngine_ApplyModifiers(t *testing.T) {
	engine := rules.NewStackingEngine()

	tests := []struct {
		name      string
		baseValue float64
		modifiers []model.StatModifier
		expected  float64
	}{
		{
			name:      "No modifiers",
			baseValue: 100.0,
			modifiers: []model.StatModifier{},
			expected:  100.0,
		},
		{
			name:      "Add flat modifiers",
			baseValue: 100.0,
			modifiers: []model.StatModifier{
				{Key: model.STR, Op: model.ADD_FLAT, Value: 10.0, Priority: 1},
				{Key: model.STR, Op: model.ADD_FLAT, Value: 5.0, Priority: 1},
			},
			expected: 115.0,
		},
		{
			name:      "Add percentage modifiers",
			baseValue: 100.0,
			modifiers: []model.StatModifier{
				{Key: model.STR, Op: model.ADD_PCT, Value: 10.0, Priority: 1}, // 10%
				{Key: model.STR, Op: model.ADD_PCT, Value: 5.0, Priority: 1},  // 5%
			},
			expected: 115.0, // 100 * (1 + 0.15)
		},
		{
			name:      "Multiplicative modifiers",
			baseValue: 100.0,
			modifiers: []model.StatModifier{
				{Key: model.STR, Op: model.MULTIPLY, Value: 1.2, Priority: 1}, // 20% increase
				{Key: model.STR, Op: model.MULTIPLY, Value: 1.1, Priority: 1}, // 10% increase
			},
			expected: 132.0, // 100 * 1.2 * 1.1
		},
		{
			name:      "Override modifiers",
			baseValue: 100.0,
			modifiers: []model.StatModifier{
				{Key: model.STR, Op: model.OVERRIDE, Value: 50.0, Priority: 1},
				{Key: model.STR, Op: model.OVERRIDE, Value: 75.0, Priority: 2}, // Higher priority
			},
			expected: 75.0, // Highest priority and value wins
		},
		{
			name:      "Mixed modifiers",
			baseValue: 100.0,
			modifiers: []model.StatModifier{
				{Key: model.STR, Op: model.ADD_FLAT, Value: 20.0, Priority: 1}, // +20
				{Key: model.STR, Op: model.ADD_PCT, Value: 10.0, Priority: 1},  // +10%
				{Key: model.STR, Op: model.MULTIPLY, Value: 1.2, Priority: 1},  // *1.2
			},
			expected: 144.0, // (100 + 20) * 1.1 * 1.2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := engine.ApplyModifiers(tt.baseValue, tt.modifiers, false)
			if result != tt.expected {
				t.Errorf("ApplyModifiers() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStackingEngine_GroupModifiersByStack(t *testing.T) {
	engine := rules.NewStackingEngine()

	modifiers := []model.StatModifier{
		{
			Key: model.STR,
			Op:  model.ADD_FLAT,
			Conditions: &model.ModifierConditions{
				StackID: "sword_bonus",
			},
		},
		{
			Key: model.STR,
			Op:  model.ADD_FLAT,
			Conditions: &model.ModifierConditions{
				StackID: "sword_bonus",
			},
		},
		{
			Key: model.STR,
			Op:  model.ADD_FLAT,
			Conditions: &model.ModifierConditions{
				StackID: "armor_bonus",
			},
		},
		{
			Key: model.STR,
			Op:  model.ADD_FLAT,
			// No stack ID
		},
	}

	groups := engine.GroupModifiersByStack(modifiers)

	// Should have 3 groups: sword_bonus, armor_bonus, and STR_ADD_FLAT
	if len(groups) != 3 {
		t.Errorf("Expected 3 groups, got %d", len(groups))
	}

	if len(groups["sword_bonus"]) != 2 {
		t.Errorf("Expected 2 modifiers in sword_bonus group, got %d", len(groups["sword_bonus"]))
	}

	if len(groups["armor_bonus"]) != 1 {
		t.Errorf("Expected 1 modifier in armor_bonus group, got %d", len(groups["armor_bonus"]))
	}
}

func TestStackingEngine_ApplyStackLimits(t *testing.T) {
	engine := rules.NewStackingEngine()

	// Create a group with max stacks = 2
	group := []model.StatModifier{
		{Key: model.STR, Op: model.ADD_FLAT, Value: 10.0, Priority: 1},
		{Key: model.STR, Op: model.ADD_FLAT, Value: 15.0, Priority: 2},
		{Key: model.STR, Op: model.ADD_FLAT, Value: 5.0, Priority: 3},
		{Key: model.STR, Op: model.ADD_FLAT, Value: 20.0, Priority: 4},
	}

	// Set max stacks to 2
	for i := range group {
		group[i].Conditions = &model.ModifierConditions{
			MaxStacks: 2,
		}
	}

	groups := map[string][]model.StatModifier{
		"test_stack": group,
	}

	result := engine.ApplyStackLimits(groups)

	if len(result) != 2 {
		t.Errorf("Expected 2 modifiers after applying stack limit, got %d", len(result))
	}

	// Should keep the highest priority modifiers
	if result[0].Priority != 4 || result[1].Priority != 2 {
		t.Errorf("Expected highest priority modifiers to be kept")
	}
}
