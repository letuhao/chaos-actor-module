package effects

import (
	"testing"
	"time"

	"actor-core-v2/models/effects"
)

func TestCombatInjuryEffect(t *testing.T) {
	effect := effects.CombatInjuryEffect()

	// Test basic properties
	if effect.ID != "combat_injury" {
		t.Errorf("Expected ID 'combat_injury', got '%s'", effect.ID)
	}

	if effect.Name != "Combat Injury" {
		t.Errorf("Expected Name 'Combat Injury', got '%s'", effect.Name)
	}

	if effect.Type != effects.NonCombatEffectTypeStatModifier {
		t.Errorf("Expected Type StatModifier, got %v", effect.Type)
	}

	if effect.Category != effects.EffectCategoryDebuff {
		t.Errorf("Expected Category Debuff, got %v", effect.Category)
	}

	if effect.Duration != 24*int64(time.Hour.Seconds()) {
		t.Errorf("Expected Duration 24h, got %v", effect.Duration)
	}

	if effect.Stackable != true {
		t.Errorf("Expected Stackable true, got %v", effect.Stackable)
	}

	if effect.MaxStacks != 3 {
		t.Errorf("Expected MaxStacks 3, got %v", effect.MaxStacks)
	}

	if effect.PersistAfterCombat != true {
		t.Errorf("Expected PersistAfterCombat true, got %v", effect.PersistAfterCombat)
	}

	if effect.CombatOnly != false {
		t.Errorf("Expected CombatOnly false, got %v", effect.CombatOnly)
	}

	if effect.CombatIntensity != 0.8 {
		t.Errorf("Expected CombatIntensity 0.8, got %v", effect.CombatIntensity)
	}

	// Test modifiers
	if len(effect.Effects) != 2 {
		t.Errorf("Expected 2 modifiers, got %d", len(effect.Effects))
	}

	// Test health modifier
	healthMod := effect.Effects[0]
	if healthMod.Stat != "health" {
		t.Errorf("Expected health stat, got '%s'", healthMod.Stat)
	}

	if healthMod.Type != effects.MultiplierModifier {
		t.Errorf("Expected Multiplier type, got %v", healthMod.Type)
	}

	if healthMod.Multiplier != 0.7 {
		t.Errorf("Expected multiplier 0.7, got %v", healthMod.Multiplier)
	}

	// Test movement speed modifier
	speedMod := effect.Effects[1]
	if speedMod.Stat != "movement_speed" {
		t.Errorf("Expected movement_speed stat, got '%s'", speedMod.Stat)
	}

	if speedMod.Type != effects.MultiplierModifier {
		t.Errorf("Expected Multiplier type, got %v", speedMod.Type)
	}

	if speedMod.Multiplier != 0.5 {
		t.Errorf("Expected multiplier 0.5, got %v", speedMod.Multiplier)
	}
}

func TestPoisonedEffect(t *testing.T) {
	effect := effects.PoisonedEffect()

	// Test basic properties
	if effect.ID != "poisoned" {
		t.Errorf("Expected ID 'poisoned', got '%s'", effect.ID)
	}

	if effect.Name != "Poisoned" {
		t.Errorf("Expected Name 'Poisoned', got '%s'", effect.Name)
	}

	if effect.Type != effects.NonCombatEffectTypeDamageOverTime {
		t.Errorf("Expected Type DamageOverTime, got %v", effect.Type)
	}

	if effect.Category != effects.EffectCategoryDebuff {
		t.Errorf("Expected Category Debuff, got %v", effect.Category)
	}

	if effect.Duration != 2*int64(time.Hour.Seconds()) {
		t.Errorf("Expected Duration 2h, got %v", effect.Duration)
	}

	if effect.Stackable != true {
		t.Errorf("Expected Stackable true, got %v", effect.Stackable)
	}

	if effect.MaxStacks != 5 {
		t.Errorf("Expected MaxStacks 5, got %v", effect.MaxStacks)
	}

	if effect.PersistAfterCombat != true {
		t.Errorf("Expected PersistAfterCombat true, got %v", effect.PersistAfterCombat)
	}

	if effect.CombatOnly != false {
		t.Errorf("Expected CombatOnly false, got %v", effect.CombatOnly)
	}

	if effect.CombatIntensity != 0.6 {
		t.Errorf("Expected CombatIntensity 0.6, got %v", effect.CombatIntensity)
	}

	// Test modifiers
	if len(effect.Effects) != 1 {
		t.Errorf("Expected 1 modifier, got %d", len(effect.Effects))
	}

	// Test health modifier
	healthMod := effect.Effects[0]
	if healthMod.Stat != "health" {
		t.Errorf("Expected health stat, got '%s'", healthMod.Stat)
	}

	if healthMod.Type != effects.AdditionModifier {
		t.Errorf("Expected Addition type, got %v", healthMod.Type)
	}

	if healthMod.Addition != -10.0 {
		t.Errorf("Expected addition -10.0, got %v", healthMod.Addition)
	}
}

func TestFearEffect(t *testing.T) {
	effect := effects.FearEffect()

	// Test basic properties
	if effect.ID != "fear" {
		t.Errorf("Expected ID 'fear', got '%s'", effect.ID)
	}

	if effect.Name != "Fear" {
		t.Errorf("Expected Name 'Fear', got '%s'", effect.Name)
	}

	if effect.Type != effects.NonCombatEffectTypeControl {
		t.Errorf("Expected Type Control, got %v", effect.Type)
	}

	if effect.Category != effects.EffectCategoryDebuff {
		t.Errorf("Expected Category Debuff, got %v", effect.Category)
	}

	if effect.Duration != 30*int64(time.Minute.Seconds()) {
		t.Errorf("Expected Duration 30m, got %v", effect.Duration)
	}

	if effect.Stackable != false {
		t.Errorf("Expected Stackable false, got %v", effect.Stackable)
	}

	if effect.MaxStacks != 1 {
		t.Errorf("Expected MaxStacks 1, got %v", effect.MaxStacks)
	}

	if effect.PersistAfterCombat != true {
		t.Errorf("Expected PersistAfterCombat true, got %v", effect.PersistAfterCombat)
	}

	if effect.CombatOnly != false {
		t.Errorf("Expected CombatOnly false, got %v", effect.CombatOnly)
	}

	if effect.CombatIntensity != 0.9 {
		t.Errorf("Expected CombatIntensity 0.9, got %v", effect.CombatIntensity)
	}

	// Test modifiers
	if len(effect.Effects) != 2 {
		t.Errorf("Expected 2 modifiers, got %d", len(effect.Effects))
	}

	// Test attack modifier
	attackMod := effect.Effects[0]
	if attackMod.Stat != "attack" {
		t.Errorf("Expected attack stat, got '%s'", attackMod.Stat)
	}

	if attackMod.Type != effects.MultiplierModifier {
		t.Errorf("Expected Multiplier type, got %v", attackMod.Type)
	}

	if attackMod.Multiplier != 0.6 {
		t.Errorf("Expected multiplier 0.6, got %v", attackMod.Multiplier)
	}

	// Test defense modifier
	defenseMod := effect.Effects[1]
	if defenseMod.Stat != "defense" {
		t.Errorf("Expected defense stat, got '%s'", defenseMod.Stat)
	}

	if defenseMod.Type != effects.MultiplierModifier {
		t.Errorf("Expected Multiplier type, got %v", defenseMod.Type)
	}

	if defenseMod.Multiplier != 0.8 {
		t.Errorf("Expected multiplier 0.8, got %v", defenseMod.Multiplier)
	}
}

func TestBerserkEffect(t *testing.T) {
	effect := effects.BerserkEffect()

	// Test basic properties
	if effect.ID != "berserk" {
		t.Errorf("Expected ID 'berserk', got '%s'", effect.ID)
	}

	if effect.Name != "Berserk" {
		t.Errorf("Expected Name 'Berserk', got '%s'", effect.Name)
	}

	if effect.Type != effects.NonCombatEffectTypeBuff {
		t.Errorf("Expected Type Buff, got %v", effect.Type)
	}

	if effect.Category != effects.EffectCategoryBuff {
		t.Errorf("Expected Category Buff, got %v", effect.Category)
	}

	if effect.Duration != 10*int64(time.Minute.Seconds()) {
		t.Errorf("Expected Duration 10m, got %v", effect.Duration)
	}

	if effect.Stackable != false {
		t.Errorf("Expected Stackable false, got %v", effect.Stackable)
	}

	if effect.MaxStacks != 1 {
		t.Errorf("Expected MaxStacks 1, got %v", effect.MaxStacks)
	}

	if effect.PersistAfterCombat != false {
		t.Errorf("Expected PersistAfterCombat false, got %v", effect.PersistAfterCombat)
	}

	if effect.CombatOnly != true {
		t.Errorf("Expected CombatOnly true, got %v", effect.CombatOnly)
	}

	if effect.CombatIntensity != 1.0 {
		t.Errorf("Expected CombatIntensity 1.0, got %v", effect.CombatIntensity)
	}

	// Test modifiers
	if len(effect.Effects) != 3 {
		t.Errorf("Expected 3 modifiers, got %d", len(effect.Effects))
	}

	// Test attack modifier
	attackMod := effect.Effects[0]
	if attackMod.Stat != "attack" {
		t.Errorf("Expected attack stat, got '%s'", attackMod.Stat)
	}

	if attackMod.Type != effects.MultiplierModifier {
		t.Errorf("Expected Multiplier type, got %v", attackMod.Type)
	}

	if attackMod.Multiplier != 1.5 {
		t.Errorf("Expected multiplier 1.5, got %v", attackMod.Multiplier)
	}

	// Test speed modifier
	speedMod := effect.Effects[1]
	if speedMod.Stat != "movement_speed" {
		t.Errorf("Expected movement_speed stat, got '%s'", speedMod.Stat)
	}

	if speedMod.Type != effects.MultiplierModifier {
		t.Errorf("Expected Multiplier type, got %v", speedMod.Type)
	}

	if speedMod.Multiplier != 1.3 {
		t.Errorf("Expected multiplier 1.3, got %v", speedMod.Multiplier)
	}

	// Test defense modifier
	defenseMod := effect.Effects[2]
	if defenseMod.Stat != "defense" {
		t.Errorf("Expected defense stat, got '%s'", defenseMod.Stat)
	}

	if defenseMod.Type != effects.MultiplierModifier {
		t.Errorf("Expected Multiplier type, got %v", defenseMod.Type)
	}

	if defenseMod.Multiplier != 0.7 {
		t.Errorf("Expected multiplier 0.7, got %v", defenseMod.Multiplier)
	}
}

func TestStunEffect(t *testing.T) {
	effect := effects.StunEffect()

	// Test basic properties
	if effect.ID != "stun" {
		t.Errorf("Expected ID 'stun', got '%s'", effect.ID)
	}

	if effect.Name != "Stun" {
		t.Errorf("Expected Name 'Stun', got '%s'", effect.Name)
	}

	if effect.Type != effects.NonCombatEffectTypeControl {
		t.Errorf("Expected Type Control, got %v", effect.Type)
	}

	if effect.Category != effects.EffectCategoryDebuff {
		t.Errorf("Expected Category Debuff, got %v", effect.Category)
	}

	if effect.Duration != 5*int64(time.Minute.Seconds()) {
		t.Errorf("Expected Duration 5m, got %v", effect.Duration)
	}

	if effect.Stackable != false {
		t.Errorf("Expected Stackable false, got %v", effect.Stackable)
	}

	if effect.MaxStacks != 1 {
		t.Errorf("Expected MaxStacks 1, got %v", effect.MaxStacks)
	}

	if effect.PersistAfterCombat != false {
		t.Errorf("Expected PersistAfterCombat false, got %v", effect.PersistAfterCombat)
	}

	if effect.CombatOnly != true {
		t.Errorf("Expected CombatOnly true, got %v", effect.CombatOnly)
	}

	if effect.CombatIntensity != 1.0 {
		t.Errorf("Expected CombatIntensity 1.0, got %v", effect.CombatIntensity)
	}

	// Test modifiers
	if len(effect.Effects) != 1 {
		t.Errorf("Expected 1 modifier, got %d", len(effect.Effects))
	}

	// Test movement speed modifier
	speedMod := effect.Effects[0]
	if speedMod.Stat != "movement_speed" {
		t.Errorf("Expected movement_speed stat, got '%s'", speedMod.Stat)
	}

	if speedMod.Type != effects.MultiplierModifier {
		t.Errorf("Expected Multiplier type, got %v", speedMod.Type)
	}

	if speedMod.Multiplier != 0.0 {
		t.Errorf("Expected multiplier 0.0, got %v", speedMod.Multiplier)
	}
}

func TestBleedingEffect(t *testing.T) {
	effect := effects.BleedingEffect()

	// Test basic properties
	if effect.ID != "bleeding" {
		t.Errorf("Expected ID 'bleeding', got '%s'", effect.ID)
	}

	if effect.Name != "Bleeding" {
		t.Errorf("Expected Name 'Bleeding', got '%s'", effect.Name)
	}

	if effect.Type != effects.NonCombatEffectTypeDamageOverTime {
		t.Errorf("Expected Type DamageOverTime, got %v", effect.Type)
	}

	if effect.Category != effects.EffectCategoryDebuff {
		t.Errorf("Expected Category Debuff, got %v", effect.Category)
	}

	if effect.Duration != 1*int64(time.Hour.Seconds()) {
		t.Errorf("Expected Duration 1h, got %v", effect.Duration)
	}

	if effect.Stackable != true {
		t.Errorf("Expected Stackable true, got %v", effect.Stackable)
	}

	if effect.MaxStacks != 3 {
		t.Errorf("Expected MaxStacks 3, got %v", effect.MaxStacks)
	}

	if effect.PersistAfterCombat != true {
		t.Errorf("Expected PersistAfterCombat true, got %v", effect.PersistAfterCombat)
	}

	if effect.CombatOnly != false {
		t.Errorf("Expected CombatOnly false, got %v", effect.CombatOnly)
	}

	if effect.CombatIntensity != 0.7 {
		t.Errorf("Expected CombatIntensity 0.7, got %v", effect.CombatIntensity)
	}

	// Test modifiers
	if len(effect.Effects) != 1 {
		t.Errorf("Expected 1 modifier, got %d", len(effect.Effects))
	}

	// Test health modifier
	healthMod := effect.Effects[0]
	if healthMod.Stat != "health" {
		t.Errorf("Expected health stat, got '%s'", healthMod.Stat)
	}

	if healthMod.Type != effects.AdditionModifier {
		t.Errorf("Expected Addition type, got %v", healthMod.Type)
	}

	if healthMod.Addition != -5.0 {
		t.Errorf("Expected addition -5.0, got %v", healthMod.Addition)
	}
}
