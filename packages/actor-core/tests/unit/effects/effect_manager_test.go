package effects

import (
	"context"
	"fmt"
	"testing"
	"time"

	"actor-core-v2/models/effects"
	effectservices "actor-core-v2/services/effects"
)

func TestEffectManagerImpl_ApplyEffect(t *testing.T) {
	// Test creating a new effect manager
	manager := effectservices.NewEffectManager(nil)
	ctx := context.Background()
	actorID := "test_actor"

	// Create a test effect
	effect := effects.NewNonCombatEffect(
		"test_effect",
		"Test Effect",
		effects.NonCombatEffectTypeStatModifier,
		effects.EffectCategoryDebuff,
	)

	effect.SetDuration(3600) // 1 hour
	effect.SetIntensity(0.8)
	effect.Activate()

	// Apply the effect
	err := manager.ApplyEffect(ctx, actorID, effect)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the effect was applied
	activeEffects, err := manager.GetActiveEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 1 {
		t.Errorf("Expected 1 active effect, got %d", len(activeEffects))
	}

	if activeEffects[0].ID != effect.ID {
		t.Errorf("Expected effect ID %s, got %s", effect.ID, activeEffects[0].ID)
	}
}

func TestEffectManagerImpl_ApplyEffect_Stackable(t *testing.T) {
	// Test creating a new effect manager
	manager := effectservices.NewEffectManager(nil)
	ctx := context.Background()
	actorID := "test_actor"

	// Create a stackable effect
	effect := effects.NewNonCombatEffect(
		"stackable_effect",
		"Stackable Effect",
		effects.NonCombatEffectTypeStatModifier,
		effects.EffectCategoryDebuff,
	)

	effect.SetDuration(3600) // 1 hour
	effect.SetIntensity(0.8)
	effect.SetStackable(true, 3)
	effect.Activate()

	// Apply the effect multiple times
	for i := 0; i < 3; i++ {
		// Create a new effect instance for each application
		newEffect := effect.Clone()
		newEffect.ID = fmt.Sprintf("stackable_effect_%d_%d", time.Now().UnixNano(), i)
		newEffect.Activate()

		err := manager.ApplyEffect(ctx, actorID, newEffect)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	}

	// Verify the effect was stacked (should have 1 effect with 3 stacks)
	activeEffects, err := manager.GetActiveEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 1 {
		t.Errorf("Expected 1 active effect, got %d", len(activeEffects))
	}

	if activeEffects[0].Stacks != 3 {
		t.Errorf("Expected 3 stacks, got %d", activeEffects[0].Stacks)
	}

	// Try to apply more than max stacks
	err = manager.ApplyEffect(ctx, actorID, effect)
	if err == nil {
		t.Errorf("Expected error when exceeding max stacks")
	}
}

func TestEffectManagerImpl_ApplyEffect_NonStackable(t *testing.T) {
	// Test creating a new effect manager
	manager := effectservices.NewEffectManager(nil)
	ctx := context.Background()
	actorID := "test_actor"

	// Create a non-stackable effect
	effect := effects.NewNonCombatEffect(
		"non_stackable_effect",
		"Non-Stackable Effect",
		effects.NonCombatEffectTypeStatModifier,
		effects.EffectCategoryDebuff,
	)

	effect.SetDuration(3600) // 1 hour
	effect.SetIntensity(0.8)
	effect.SetStackable(false, 1)
	effect.Activate()

	// Apply the effect
	err := manager.ApplyEffect(ctx, actorID, effect)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Try to apply the same effect again (same type)
	err = manager.ApplyEffect(ctx, actorID, effect)
	if err == nil {
		t.Errorf("Expected error when applying non-stackable effect twice")
	}

	// Verify only one effect exists
	activeEffects, err := manager.GetActiveEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 1 {
		t.Errorf("Expected 1 active effect, got %d", len(activeEffects))
	}
}

func TestEffectManagerImpl_RemoveEffect(t *testing.T) {
	// Test creating a new effect manager
	manager := effectservices.NewEffectManager(nil)
	ctx := context.Background()
	actorID := "test_actor"

	// Create a test effect
	effect := effects.NewNonCombatEffect(
		"test_effect",
		"Test Effect",
		effects.NonCombatEffectTypeStatModifier,
		effects.EffectCategoryDebuff,
	)

	effect.SetDuration(3600) // 1 hour
	effect.SetIntensity(0.8)
	effect.Activate()

	// Apply the effect
	err := manager.ApplyEffect(ctx, actorID, effect)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Remove the effect
	err = manager.RemoveEffect(ctx, actorID, effect.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the effect was removed
	activeEffects, err := manager.GetActiveEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 0 {
		t.Errorf("Expected 0 active effects, got %d", len(activeEffects))
	}
}

func TestEffectManagerImpl_GetEffectByID(t *testing.T) {
	// Test creating a new effect manager
	manager := effectservices.NewEffectManager(nil)
	ctx := context.Background()
	actorID := "test_actor"

	// Create a test effect
	effect := effects.NewNonCombatEffect(
		"test_effect",
		"Test Effect",
		effects.NonCombatEffectTypeStatModifier,
		effects.EffectCategoryDebuff,
	)

	effect.SetDuration(3600) // 1 hour
	effect.SetIntensity(0.8)
	effect.Activate()

	// Apply the effect
	err := manager.ApplyEffect(ctx, actorID, effect)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Get the effect by ID
	retrievedEffect, err := manager.GetEffectByID(ctx, actorID, effect.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if retrievedEffect == nil {
		t.Errorf("Expected effect to be found")
	}

	if retrievedEffect.ID != effect.ID {
		t.Errorf("Expected effect ID %s, got %s", effect.ID, retrievedEffect.ID)
	}

	// Try to get non-existent effect
	_, err = manager.GetEffectByID(ctx, actorID, "non_existent")
	if err == nil {
		t.Errorf("Expected error when getting non-existent effect")
	}
}

func TestEffectManagerImpl_GetEffectsByType(t *testing.T) {
	// Test creating a new effect manager
	manager := effectservices.NewEffectManager(nil)
	ctx := context.Background()
	actorID := "test_actor"

	// Create effects of different types
	effect1 := effects.NewNonCombatEffect(
		"effect1",
		"Effect 1",
		effects.NonCombatEffectTypeStatModifier,
		effects.EffectCategoryDebuff,
	)
	effect1.SetDuration(3600)
	effect1.Activate()

	effect2 := effects.NewNonCombatEffect(
		"effect2",
		"Effect 2",
		effects.NonCombatEffectTypeBuff,
		effects.EffectCategoryBuff,
	)
	effect2.SetDuration(3600)
	effect2.Activate()

	// Apply both effects
	err := manager.ApplyEffect(ctx, actorID, effect1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = manager.ApplyEffect(ctx, actorID, effect2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Get effects by type
	statModifierEffects, err := manager.GetEffectsByType(ctx, actorID, effects.NonCombatEffectTypeStatModifier)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(statModifierEffects) != 1 {
		t.Errorf("Expected 1 stat modifier effect, got %d", len(statModifierEffects))
	}

	if statModifierEffects[0].ID != effect1.ID {
		t.Errorf("Expected effect ID %s, got %s", effect1.ID, statModifierEffects[0].ID)
	}
}

func TestEffectManagerImpl_GetEffectsByCategory(t *testing.T) {
	// Test creating a new effect manager
	manager := effectservices.NewEffectManager(nil)
	ctx := context.Background()
	actorID := "test_actor"

	// Create effects of different categories
	effect1 := effects.NewNonCombatEffect(
		"effect1",
		"Effect 1",
		effects.NonCombatEffectTypeStatModifier,
		effects.EffectCategoryDebuff,
	)
	effect1.SetDuration(3600)
	effect1.Activate()

	effect2 := effects.NewNonCombatEffect(
		"effect2",
		"Effect 2",
		effects.NonCombatEffectTypeBuff,
		effects.EffectCategoryBuff,
	)
	effect2.SetDuration(3600)
	effect2.Activate()

	// Apply both effects
	err := manager.ApplyEffect(ctx, actorID, effect1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = manager.ApplyEffect(ctx, actorID, effect2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Get effects by category
	debuffEffects, err := manager.GetEffectsByCategory(ctx, actorID, effects.EffectCategoryDebuff)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(debuffEffects) != 1 {
		t.Errorf("Expected 1 debuff effect, got %d", len(debuffEffects))
	}

	if debuffEffects[0].ID != effect1.ID {
		t.Errorf("Expected effect ID %s, got %s", effect1.ID, debuffEffects[0].ID)
	}
}

func TestEffectManagerImpl_ProcessEffects(t *testing.T) {
	// Test creating a new effect manager
	manager := effectservices.NewEffectManager(nil)
	ctx := context.Background()
	actorID := "test_actor"

	// Create a test effect with short duration
	effect := effects.NewNonCombatEffect(
		"test_effect",
		"Test Effect",
		effects.NonCombatEffectTypeStatModifier,
		effects.EffectCategoryDebuff,
	)

	effect.SetDuration(1 * int64(time.Second.Seconds()))
	effect.Activate()

	// Apply the effect
	err := manager.ApplyEffect(ctx, actorID, effect)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Process effects immediately
	err = manager.ProcessEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Wait for effect to expire
	time.Sleep(2 * time.Second)

	// Process effects again
	err = manager.ProcessEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the effect was removed
	activeEffects, err := manager.GetActiveEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 0 {
		t.Errorf("Expected 0 active effects, got %d", len(activeEffects))
	}
}

func TestEffectManagerImpl_ClearAllEffects(t *testing.T) {
	// Test creating a new effect manager
	manager := effectservices.NewEffectManager(nil)
	ctx := context.Background()
	actorID := "test_actor"

	// Create multiple effects
	effect1 := effects.NewNonCombatEffect(
		"effect1",
		"Effect 1",
		effects.NonCombatEffectTypeStatModifier,
		effects.EffectCategoryDebuff,
	)
	effect1.SetDuration(3600)
	effect1.Activate()

	effect2 := effects.NewNonCombatEffect(
		"effect2",
		"Effect 2",
		effects.NonCombatEffectTypeBuff,
		effects.EffectCategoryBuff,
	)
	effect2.SetDuration(3600)
	effect2.Activate()

	// Apply both effects
	err := manager.ApplyEffect(ctx, actorID, effect1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = manager.ApplyEffect(ctx, actorID, effect2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Clear all effects
	err = manager.ClearAllEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify all effects were removed
	activeEffects, err := manager.GetActiveEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 0 {
		t.Errorf("Expected 0 active effects, got %d", len(activeEffects))
	}
}
