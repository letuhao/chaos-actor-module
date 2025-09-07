package effects

import (
	"testing"
	"time"

	"actor-core-v2/models/effects"
)

func TestNonCombatEffect_NewNonCombatEffect(t *testing.T) {
	// Test creating a new non-combat effect
	effect := effects.NewNonCombatEffect("test_effect", "Test Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	if effect == nil {
		t.Errorf("Expected non-combat effect to be created")
	}

	if effect.ID != "test_effect" {
		t.Errorf("Expected ID 'test_effect', got '%s'", effect.ID)
	}

	if effect.Name != "Test Effect" {
		t.Errorf("Expected name 'Test Effect', got '%s'", effect.Name)
	}

	if effect.Type != effects.NonCombatEffectTypeStatModifier {
		t.Errorf("Expected type %s, got %s", effects.NonCombatEffectTypeStatModifier, effect.Type)
	}

	if effect.Category != effects.EffectCategoryDebuff {
		t.Errorf("Expected category %s, got %s", effects.EffectCategoryDebuff, effect.Category)
	}

	if effect.Duration != 0 {
		t.Errorf("Expected Duration 0, got %d", effect.Duration)
	}

	if effect.Stackable {
		t.Errorf("Expected Stackable to be false by default")
	}

	if effect.MaxStacks != 1 {
		t.Errorf("Expected MaxStacks 1, got %d", effect.MaxStacks)
	}

	if len(effect.Effects) != 0 {
		t.Errorf("Expected empty Effects, got %d", len(effect.Effects))
	}

	if len(effect.Conditions) != 0 {
		t.Errorf("Expected empty Conditions, got %d", len(effect.Conditions))
	}
}

func TestNonCombatEffect_AddModifier(t *testing.T) {
	effect := effects.NewNonCombatEffect("test_effect", "Test Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Add a modifier
	modifier := effects.NonCombatEffectModifier{
		Type:       effects.AdditionModifier,
		Target:     "primary",
		Stat:       "attack_power",
		Value:      10.0,
		Multiplier: 1.0,
		Addition:   10.0,
		Conditions: []effects.EffectCondition{},
	}

	effect.AddModifier(modifier)

	if len(effect.Effects) != 1 {
		t.Errorf("Expected 1 modifier, got %d", len(effect.Effects))
	}

	if effect.Effects[0].Type != effects.AdditionModifier {
		t.Errorf("Expected modifier type %s, got %s", effects.AdditionModifier, effect.Effects[0].Type)
	}

	if effect.Effects[0].Target != "primary" {
		t.Errorf("Expected modifier target 'primary', got '%s'", effect.Effects[0].Target)
	}

	if effect.Effects[0].Stat != "attack_power" {
		t.Errorf("Expected modifier stat 'attack_power', got '%s'", effect.Effects[0].Stat)
	}

	if effect.Effects[0].Value != 10.0 {
		t.Errorf("Expected modifier value 10.0, got %f", effect.Effects[0].Value)
	}

	if effect.Effects[0].Multiplier != 1.0 {
		t.Errorf("Expected modifier multiplier 1.0, got %f", effect.Effects[0].Multiplier)
	}

	if effect.Effects[0].Addition != 10.0 {
		t.Errorf("Expected modifier addition 10.0, got %f", effect.Effects[0].Addition)
	}
}

func TestNonCombatEffect_SetDuration(t *testing.T) {
	effect := effects.NewNonCombatEffect("test_effect", "Test Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Test setting duration
	duration := int64(3600)
	effect.SetDuration(duration)

	if effect.Duration != duration {
		t.Errorf("Expected Duration %d, got %d", duration, effect.Duration)
	}

	// Test setting zero duration
	effect.SetDuration(0)
	if effect.Duration != 0 {
		t.Errorf("Expected Duration 0, got %d", effect.Duration)
	}

	// Test setting negative duration
	effect.SetDuration(-1)
	if effect.Duration != -1 {
		t.Errorf("Expected Duration -1, got %d", effect.Duration)
	}
}

func TestNonCombatEffect_SetIntensity(t *testing.T) {
	effect := effects.NewNonCombatEffect("test_effect", "Test Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Test setting intensity
	intensity := 1.5
	effect.SetIntensity(intensity)

	if effect.Intensity != intensity {
		t.Errorf("Expected Intensity %f, got %f", intensity, effect.Intensity)
	}

	// Test setting zero intensity
	effect.SetIntensity(0.0)
	if effect.Intensity != 0.0 {
		t.Errorf("Expected Intensity 0.0, got %f", effect.Intensity)
	}

	// Test setting negative intensity
	effect.SetIntensity(-1.0)
	if effect.Intensity != -1.0 {
		t.Errorf("Expected Intensity -1.0, got %f", effect.Intensity)
	}
}

func TestNonCombatEffect_SetStackable(t *testing.T) {
	effect := effects.NewNonCombatEffect("test_effect", "Test Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Test setting to true
	effect.SetStackable(true, 3)
	if !effect.Stackable {
		t.Errorf("Expected Stackable to be true")
	}

	if effect.MaxStacks != 3 {
		t.Errorf("Expected MaxStacks 3, got %d", effect.MaxStacks)
	}

	// Test setting to false
	effect.SetStackable(false, 1)
	if effect.Stackable {
		t.Errorf("Expected Stackable to be false")
	}

	if effect.MaxStacks != 1 {
		t.Errorf("Expected MaxStacks 1, got %d", effect.MaxStacks)
	}
}

func TestNonCombatEffect_AddCondition(t *testing.T) {
	effect := effects.NewNonCombatEffect("test_effect", "Test Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Add a condition
	condition := effects.EffectCondition{
		Type:     "health_percentage",
		Operator: "less_than",
		Value:    0.5,
		Variable: "health",
	}

	effect.AddCondition(condition)

	if len(effect.Conditions) != 1 {
		t.Errorf("Expected 1 condition, got %d", len(effect.Conditions))
	}

	if effect.Conditions[0].Type != "health_percentage" {
		t.Errorf("Expected condition type 'health_percentage', got '%s'", effect.Conditions[0].Type)
	}

	if effect.Conditions[0].Operator != "less_than" {
		t.Errorf("Expected condition operator 'less_than', got '%s'", effect.Conditions[0].Operator)
	}

	if effect.Conditions[0].Value != 0.5 {
		t.Errorf("Expected condition value 0.5, got %v", effect.Conditions[0].Value)
	}

	if effect.Conditions[0].Variable != "health" {
		t.Errorf("Expected condition variable 'health', got '%s'", effect.Conditions[0].Variable)
	}
}

func TestNonCombatEffect_ModifierTypes(t *testing.T) {
	effect := effects.NewNonCombatEffect("test_effect", "Test Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Test different modifier types
	modifiers := []effects.NonCombatEffectModifier{
		{
			Type:       effects.AdditionModifier,
			Target:     "primary",
			Stat:       "attack_power",
			Value:      10.0,
			Multiplier: 1.0,
			Addition:   10.0,
			Conditions: []effects.EffectCondition{},
		},
		{
			Type:       effects.MultiplierModifier,
			Target:     "derived",
			Stat:       "defense_power",
			Value:      1.5,
			Multiplier: 1.5,
			Addition:   0.0,
			Conditions: []effects.EffectCondition{},
		},
		{
			Type:       effects.OverrideModifier,
			Target:     "custom",
			Stat:       "armor_penetration",
			Value:      0.3,
			Multiplier: 1.0,
			Addition:   0.3,
			Conditions: []effects.EffectCondition{},
		},
	}

	for _, modifier := range modifiers {
		effect.AddModifier(modifier)
	}

	if len(effect.Effects) != 3 {
		t.Errorf("Expected 3 modifiers, got %d", len(effect.Effects))
	}

	// Verify each modifier type
	if effect.Effects[0].Type != effects.AdditionModifier {
		t.Errorf("Expected first modifier type %s, got %s", effects.AdditionModifier, effect.Effects[0].Type)
	}

	if effect.Effects[1].Type != effects.MultiplierModifier {
		t.Errorf("Expected second modifier type %s, got %s", effects.MultiplierModifier, effect.Effects[1].Type)
	}

	if effect.Effects[2].Type != effects.OverrideModifier {
		t.Errorf("Expected third modifier type %s, got %s", effects.OverrideModifier, effect.Effects[2].Type)
	}
}

func TestNonCombatEffect_IsExpired(t *testing.T) {
	effect := effects.NewNonCombatEffect("test_effect", "Test Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Test permanent effect (no expiration)
	if effect.IsExpired() {
		t.Errorf("Expected permanent effect to not be expired")
	}

	// Test effect with duration
	effect.SetDuration(1) // 1 second
	effect.Activate()

	// Should not be expired immediately
	if effect.IsExpired() {
		t.Errorf("Expected effect to not be expired immediately")
	}

	// Wait for expiration
	time.Sleep(2 * time.Second)

	// Should be expired now
	if !effect.IsExpired() {
		t.Errorf("Expected effect to be expired after duration")
	}
}

func TestNonCombatEffect_Clone(t *testing.T) {
	effect := effects.NewNonCombatEffect("test_effect", "Test Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)
	effect.SetDuration(3600)
	effect.SetIntensity(1.5)
	effect.SetStackable(true, 3)

	// Add a modifier
	modifier := effects.NonCombatEffectModifier{
		Type:       effects.AdditionModifier,
		Target:     "primary",
		Stat:       "attack_power",
		Value:      10.0,
		Multiplier: 1.0,
		Addition:   10.0,
		Conditions: []effects.EffectCondition{},
	}
	effect.AddModifier(modifier)

	// Clone the effect
	clone := effect.Clone()

	if clone == nil {
		t.Errorf("Expected clone to be created")
	}

	if clone.ID != effect.ID {
		t.Errorf("Expected clone ID %s, got %s", effect.ID, clone.ID)
	}

	if clone.Name != effect.Name {
		t.Errorf("Expected clone name %s, got %s", effect.Name, clone.Name)
	}

	if clone.Type != effect.Type {
		t.Errorf("Expected clone type %s, got %s", effect.Type, clone.Type)
	}

	if clone.Category != effect.Category {
		t.Errorf("Expected clone category %s, got %s", effect.Category, clone.Category)
	}

	if clone.Duration != effect.Duration {
		t.Errorf("Expected clone duration %d, got %d", effect.Duration, clone.Duration)
	}

	if clone.Intensity != effect.Intensity {
		t.Errorf("Expected clone intensity %f, got %f", effect.Intensity, clone.Intensity)
	}

	if clone.Stackable != effect.Stackable {
		t.Errorf("Expected clone stackable %t, got %t", effect.Stackable, clone.Stackable)
	}

	if clone.MaxStacks != effect.MaxStacks {
		t.Errorf("Expected clone max stacks %d, got %d", effect.MaxStacks, clone.MaxStacks)
	}

	if len(clone.Effects) != len(effect.Effects) {
		t.Errorf("Expected clone effects length %d, got %d", len(effect.Effects), len(clone.Effects))
	}

	if len(clone.Conditions) != len(effect.Conditions) {
		t.Errorf("Expected clone conditions length %d, got %d", len(effect.Conditions), len(clone.Conditions))
	}
}
