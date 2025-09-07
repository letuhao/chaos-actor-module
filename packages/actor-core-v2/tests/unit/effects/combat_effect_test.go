package effects

import (
	"testing"
	"time"

	"actor-core/models/effects"
)

func TestCombatEffect_NewCombatEffect(t *testing.T) {
	// Test creating a new combat effect
	effect := effects.NewCombatEffect("test_combat_effect", "Test Combat Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	if effect == nil {
		t.Errorf("Expected combat effect to be created")
	}

	if effect.ID != "test_combat_effect" {
		t.Errorf("Expected ID 'test_combat_effect', got '%s'", effect.ID)
	}

	if effect.Name != "Test Combat Effect" {
		t.Errorf("Expected name 'Test Combat Effect', got '%s'", effect.Name)
	}

	if effect.Type != effects.NonCombatEffectTypeStatModifier {
		t.Errorf("Expected type %s, got %s", effects.NonCombatEffectTypeStatModifier, effect.Type)
	}

	if effect.Category != effects.EffectCategoryDebuff {
		t.Errorf("Expected category %s, got %s", effects.EffectCategoryDebuff, effect.Category)
	}

	if effect.CombatID != "" {
		t.Errorf("Expected empty CombatID, got '%s'", effect.CombatID)
	}

	if effect.AttackerID != "" {
		t.Errorf("Expected empty AttackerID, got '%s'", effect.AttackerID)
	}

	if effect.TargetID != "" {
		t.Errorf("Expected empty TargetID, got '%s'", effect.TargetID)
	}

	if effect.CombatOnly {
		t.Errorf("Expected CombatOnly to be false by default")
	}

	if !effect.PersistAfterCombat {
		t.Errorf("Expected PersistAfterCombat to be true by default")
	}

	if effect.CombatIntensity != 1.0 {
		t.Errorf("Expected CombatIntensity to be 1.0, got %f", effect.CombatIntensity)
	}

	if len(effect.CombatModifiers) != 0 {
		t.Errorf("Expected empty CombatModifiers, got %d", len(effect.CombatModifiers))
	}
}

func TestCombatEffect_AddCombatModifier(t *testing.T) {
	effect := effects.NewCombatEffect("test_combat_effect", "Test Combat Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Add a combat modifier
	modifier := effects.CombatEffectModifier{
		Type:       effects.MultiplierModifier,
		Target:     "primary",
		Stat:       "attack_power",
		Value:      1.5,
		Multiplier: 1.2,
		Addition:   10.0,
		Conditions: []effects.EffectCondition{},
	}

	effect.AddCombatModifier(modifier)

	if len(effect.CombatModifiers) != 1 {
		t.Errorf("Expected 1 combat modifier, got %d", len(effect.CombatModifiers))
	}

	if effect.CombatModifiers[0].Type != effects.MultiplierModifier {
		t.Errorf("Expected modifier type %s, got %s", effects.MultiplierModifier, effect.CombatModifiers[0].Type)
	}

	if effect.CombatModifiers[0].Target != "primary" {
		t.Errorf("Expected modifier target 'primary', got '%s'", effect.CombatModifiers[0].Target)
	}

	if effect.CombatModifiers[0].Stat != "attack_power" {
		t.Errorf("Expected modifier stat 'attack_power', got '%s'", effect.CombatModifiers[0].Stat)
	}

	if effect.CombatModifiers[0].Value != 1.5 {
		t.Errorf("Expected modifier value 1.5, got %f", effect.CombatModifiers[0].Value)
	}

	if effect.CombatModifiers[0].Multiplier != 1.2 {
		t.Errorf("Expected modifier multiplier 1.2, got %f", effect.CombatModifiers[0].Multiplier)
	}

	if effect.CombatModifiers[0].Addition != 10.0 {
		t.Errorf("Expected modifier addition 10.0, got %f", effect.CombatModifiers[0].Addition)
	}
}

func TestCombatEffect_SetCombatInfo(t *testing.T) {
	effect := effects.NewCombatEffect("test_combat_effect", "Test Combat Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	combatID := "combat_123"
	attackerID := "attacker_456"
	targetID := "target_789"
	combatStartTime := time.Now().Unix()

	effect.SetCombatInfo(combatID, attackerID, targetID, combatStartTime)

	if effect.CombatID != combatID {
		t.Errorf("Expected CombatID %s, got %s", combatID, effect.CombatID)
	}

	if effect.AttackerID != attackerID {
		t.Errorf("Expected AttackerID %s, got %s", attackerID, effect.AttackerID)
	}

	if effect.TargetID != targetID {
		t.Errorf("Expected TargetID %s, got %s", targetID, effect.TargetID)
	}

	if effect.CombatStartTime != combatStartTime {
		t.Errorf("Expected CombatStartTime %d, got %d", combatStartTime, effect.CombatStartTime)
	}
}

func TestCombatEffect_SetPersistAfterCombat(t *testing.T) {
	effect := effects.NewCombatEffect("test_combat_effect", "Test Combat Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Test setting to true
	effect.SetPersistAfterCombat(true)
	if !effect.PersistAfterCombat {
		t.Errorf("Expected PersistAfterCombat to be true")
	}

	// Test setting to false
	effect.SetPersistAfterCombat(false)
	if effect.PersistAfterCombat {
		t.Errorf("Expected PersistAfterCombat to be false")
	}
}

func TestCombatEffect_SetCombatOnly(t *testing.T) {
	effect := effects.NewCombatEffect("test_combat_effect", "Test Combat Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Test setting to true
	effect.SetCombatOnly(true)
	if !effect.CombatOnly {
		t.Errorf("Expected CombatOnly to be true")
	}

	// Test setting to false
	effect.SetCombatOnly(false)
	if effect.CombatOnly {
		t.Errorf("Expected CombatOnly to be false")
	}
}

func TestCombatEffect_SetCombatIntensity(t *testing.T) {
	effect := effects.NewCombatEffect("test_combat_effect", "Test Combat Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Test setting combat intensity
	intensity := 2.5
	effect.SetCombatIntensity(intensity)

	if effect.CombatIntensity != intensity {
		t.Errorf("Expected CombatIntensity %f, got %f", intensity, effect.CombatIntensity)
	}

	// Test setting negative intensity
	effect.SetCombatIntensity(-1.0)
	if effect.CombatIntensity != -1.0 {
		t.Errorf("Expected CombatIntensity -1.0, got %f", effect.CombatIntensity)
	}

	// Test setting zero intensity
	effect.SetCombatIntensity(0.0)
	if effect.CombatIntensity != 0.0 {
		t.Errorf("Expected CombatIntensity 0.0, got %f", effect.CombatIntensity)
	}
}

func TestCombatEffect_Embedding(t *testing.T) {
	effect := effects.NewCombatEffect("test_combat_effect", "Test Combat Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Test that we can access embedded NonCombatEffect fields
	if effect.ID != "test_combat_effect" {
		t.Errorf("Expected embedded ID 'test_combat_effect', got '%s'", effect.ID)
	}

	if effect.Name != "Test Combat Effect" {
		t.Errorf("Expected embedded name 'Test Combat Effect', got '%s'", effect.Name)
	}

	if effect.Type != effects.NonCombatEffectTypeStatModifier {
		t.Errorf("Expected embedded type %s, got %s", effects.NonCombatEffectTypeStatModifier, effect.Type)
	}

	if effect.Category != effects.EffectCategoryDebuff {
		t.Errorf("Expected embedded category %s, got %s", effects.EffectCategoryDebuff, effect.Category)
	}

	// Test that we can set embedded fields
	effect.SetDuration(3600)
	effect.SetStackable(true, 3)

	if effect.Duration != 3600 {
		t.Errorf("Expected embedded Duration 3600, got %d", effect.Duration)
	}

	if !effect.Stackable {
		t.Errorf("Expected embedded Stackable to be true")
	}

	if effect.MaxStacks != 3 {
		t.Errorf("Expected embedded MaxStacks 3, got %d", effect.MaxStacks)
	}
}

func TestCombatEffect_CombatModifierTypes(t *testing.T) {
	effect := effects.NewCombatEffect("test_combat_effect", "Test Combat Effect", effects.NonCombatEffectTypeStatModifier, effects.EffectCategoryDebuff)

	// Test different modifier types
	modifiers := []effects.CombatEffectModifier{
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
		effect.AddCombatModifier(modifier)
	}

	if len(effect.CombatModifiers) != 3 {
		t.Errorf("Expected 3 combat modifiers, got %d", len(effect.CombatModifiers))
	}

	// Verify each modifier type
	if effect.CombatModifiers[0].Type != effects.AdditionModifier {
		t.Errorf("Expected first modifier type %s, got %s", effects.AdditionModifier, effect.CombatModifiers[0].Type)
	}

	if effect.CombatModifiers[1].Type != effects.MultiplierModifier {
		t.Errorf("Expected second modifier type %s, got %s", effects.MultiplierModifier, effect.CombatModifiers[1].Type)
	}

	if effect.CombatModifiers[2].Type != effects.OverrideModifier {
		t.Errorf("Expected third modifier type %s, got %s", effects.OverrideModifier, effect.CombatModifiers[2].Type)
	}
}
