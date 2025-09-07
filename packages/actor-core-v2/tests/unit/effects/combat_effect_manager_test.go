package effects

import (
	"context"
	"fmt"
	"testing"
	"time"

	"actor-core/models/effects"
	effectservices "actor-core/services/effects"
)

func TestCombatEffectManager_NewCombatEffectManager(t *testing.T) {
	// Test creating a new combat effect manager
	manager := effectservices.NewCombatEffectManager(nil)

	if manager == nil {
		t.Errorf("Expected combat effect manager to be created")
	}
}

func TestCombatEffectManager_StartCombat(t *testing.T) {
	// Create a mock effect manager
	effectManager := effectservices.NewEffectManager(nil)
	manager := effectservices.NewCombatEffectManager(effectManager)
	ctx := context.Background()
	combatID := "combat_123"
	participants := []string{"actor_1", "actor_2"}

	// Start combat
	err := manager.StartCombat(ctx, combatID, participants)
	if err != nil {
		t.Errorf("Expected no error starting combat, got %v", err)
	}

	// Verify combat is active
	combatState, err := manager.GetCombatState(ctx, combatID)
	if err != nil {
		t.Errorf("Expected no error getting combat state, got %v", err)
	}

	if combatState == nil {
		t.Errorf("Expected combat state to be returned")
	}

	if combatState.CombatID != combatID {
		t.Errorf("Expected CombatID %s, got %s", combatID, combatState.CombatID)
	}

	if !combatState.IsActive {
		t.Errorf("Expected combat to be active")
	}
}

func TestCombatEffectManager_ApplyCombatEffect(t *testing.T) {
	// Create a mock effect manager
	effectManager := effectservices.NewEffectManager(nil)
	manager := effectservices.NewCombatEffectManager(effectManager)
	ctx := context.Background()
	actorID := "test_actor"
	combatID := "combat_123"

	// Start combat first
	err := manager.StartCombat(ctx, combatID, []string{actorID, "attacker_456"})
	if err != nil {
		t.Errorf("Expected no error starting combat, got %v", err)
	}

	// Create a test combat effect with unique ID
	effect := effects.CombatInjuryEffect()
	effect.ID = fmt.Sprintf("combat_injury_%d", time.Now().UnixNano())
	effect.SetCombatInfo(combatID, "attacker_456", actorID, time.Now().Unix())
	effect.Activate() // Activate the effect

	// Apply the combat effect
	err = manager.ApplyCombatEffect(ctx, actorID, effect)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the effect was applied
	activeEffects, err := manager.GetActiveCombatEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 1 {
		t.Errorf("Expected 1 active combat effect, got %d", len(activeEffects))
	}

	if activeEffects[0].ID != effect.ID {
		t.Errorf("Expected effect ID %s, got %s", effect.ID, activeEffects[0].ID)
	}

	if activeEffects[0].CombatID != combatID {
		t.Errorf("Expected CombatID %s, got %s", combatID, activeEffects[0].CombatID)
	}
}

func TestCombatEffectManager_ApplyCombatEffect_Stackable(t *testing.T) {
	// Create a mock effect manager
	effectManager := effectservices.NewEffectManager(nil)
	manager := effectservices.NewCombatEffectManager(effectManager)
	ctx := context.Background()
	actorID := "test_actor"
	combatID := "combat_123"

	// Start combat first
	err := manager.StartCombat(ctx, combatID, []string{actorID, "attacker_456"})
	if err != nil {
		t.Errorf("Expected no error starting combat, got %v", err)
	}

	// Create a stackable combat effect with unique ID
	effect := effects.PoisonedEffect()
	effect.ID = fmt.Sprintf("poisoned_%d", time.Now().UnixNano())
	effect.SetCombatInfo(combatID, "attacker_456", actorID, time.Now().Unix())
	effect.Activate() // Activate the effect

	// Apply the effect multiple times (up to max stacks)
	for i := 0; i < 5; i++ {
		// Create a new effect instance for each application
		newEffect := effect.Clone()
		newEffect.ID = fmt.Sprintf("poisoned_%d_%d", time.Now().UnixNano(), i)
		newEffect.SetCombatInfo(combatID, "attacker_456", actorID, time.Now().Unix())
		newEffect.Activate()

		err := manager.ApplyCombatEffect(ctx, actorID, newEffect)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	}

	// Verify the effect was stacked (should have 1 effect with 5 stacks)
	activeEffects, err := manager.GetActiveCombatEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 1 {
		t.Errorf("Expected 1 active combat effect, got %d", len(activeEffects))
	}

	if activeEffects[0].Stacks != 5 {
		t.Errorf("Expected 5 stacks, got %d", activeEffects[0].Stacks)
	}

	// Try to apply more than max stacks
	err = manager.ApplyCombatEffect(ctx, actorID, effect)
	if err == nil {
		t.Errorf("Expected error when exceeding max stacks")
	}
}

func TestCombatEffectManager_ApplyCombatEffect_NonStackable(t *testing.T) {
	// Create a mock effect manager
	effectManager := effectservices.NewEffectManager(nil)
	manager := effectservices.NewCombatEffectManager(effectManager)
	ctx := context.Background()
	actorID := "test_actor"
	combatID := "combat_123"

	// Start combat first
	err := manager.StartCombat(ctx, combatID, []string{actorID, "attacker_456"})
	if err != nil {
		t.Errorf("Expected no error starting combat, got %v", err)
	}

	// Create a non-stackable combat effect with unique ID
	effect := effects.FearEffect()
	effect.ID = fmt.Sprintf("fear_%d", time.Now().UnixNano())
	effect.SetCombatInfo(combatID, "attacker_456", actorID, time.Now().Unix())
	effect.Activate() // Activate the effect

	// Apply the effect
	err = manager.ApplyCombatEffect(ctx, actorID, effect)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Try to apply the same effect again (same type)
	err = manager.ApplyCombatEffect(ctx, actorID, effect)
	if err == nil {
		t.Errorf("Expected error when applying non-stackable effect twice")
	}

	// Verify only one effect exists
	activeEffects, err := manager.GetActiveCombatEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 1 {
		t.Errorf("Expected 1 active combat effect, got %d", len(activeEffects))
	}
}

func TestCombatEffectManager_RemoveCombatEffect(t *testing.T) {
	// Create a mock effect manager
	effectManager := effectservices.NewEffectManager(nil)
	manager := effectservices.NewCombatEffectManager(effectManager)
	ctx := context.Background()
	actorID := "test_actor"
	combatID := "combat_123"

	// Start combat first
	err := manager.StartCombat(ctx, combatID, []string{actorID, "attacker_456"})
	if err != nil {
		t.Errorf("Expected no error starting combat, got %v", err)
	}

	// Create and apply a test combat effect
	effect := effects.CombatInjuryEffect()
	effect.ID = fmt.Sprintf("combat_injury_%d", time.Now().UnixNano())
	effect.SetCombatInfo(combatID, "attacker_456", actorID, time.Now().Unix())
	effect.Activate() // Activate the effect

	err = manager.ApplyCombatEffect(ctx, actorID, effect)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Remove the effect
	err = manager.RemoveCombatEffect(ctx, actorID, effect.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the effect was removed
	activeEffects, err := manager.GetActiveCombatEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 0 {
		t.Errorf("Expected 0 active combat effects, got %d", len(activeEffects))
	}
}

func TestCombatEffectManager_EndCombat(t *testing.T) {
	// Create a mock effect manager
	effectManager := effectservices.NewEffectManager(nil)
	manager := effectservices.NewCombatEffectManager(effectManager)
	ctx := context.Background()
	actorID := "test_actor"
	combatID := "combat_123"

	// Start combat first
	err := manager.StartCombat(ctx, combatID, []string{actorID, "attacker_456"})
	if err != nil {
		t.Errorf("Expected no error starting combat, got %v", err)
	}

	// Create and apply a test combat effect
	effect := effects.CombatInjuryEffect()
	effect.ID = fmt.Sprintf("combat_injury_%d", time.Now().UnixNano())
	effect.SetCombatInfo(combatID, "attacker_456", actorID, time.Now().Unix())
	effect.Activate() // Activate the effect

	err = manager.ApplyCombatEffect(ctx, actorID, effect)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// End combat
	err = manager.EndCombat(ctx, combatID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify combat is ended
	combatState, err := manager.GetCombatState(ctx, combatID)
	if err != nil {
		t.Errorf("Expected no error getting combat state, got %v", err)
	}

	if combatState == nil {
		t.Errorf("Expected combat state to be returned")
	}

	if combatState.IsActive {
		t.Errorf("Expected combat to be ended")
	}
}

func TestCombatEffectManager_GetCombatEffectByID(t *testing.T) {
	// Create a mock effect manager
	effectManager := effectservices.NewEffectManager(nil)
	manager := effectservices.NewCombatEffectManager(effectManager)
	ctx := context.Background()
	actorID := "test_actor"
	combatID := "combat_123"

	// Start combat first
	err := manager.StartCombat(ctx, combatID, []string{actorID, "attacker_456"})
	if err != nil {
		t.Errorf("Expected no error starting combat, got %v", err)
	}

	// Create and apply a test combat effect
	effect := effects.CombatInjuryEffect()
	effect.ID = fmt.Sprintf("combat_injury_%d", time.Now().UnixNano())
	effect.SetCombatInfo(combatID, "attacker_456", actorID, time.Now().Unix())
	effect.Activate() // Activate the effect

	err = manager.ApplyCombatEffect(ctx, actorID, effect)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Get the effect by ID
	retrievedEffect, err := manager.GetCombatEffectByID(ctx, actorID, effect.ID)
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
	_, err = manager.GetCombatEffectByID(ctx, actorID, "non_existent")
	if err == nil {
		t.Errorf("Expected error when getting non-existent effect")
	}
}

func TestCombatEffectManager_ProcessCombatEffects(t *testing.T) {
	// Create a mock effect manager
	effectManager := effectservices.NewEffectManager(nil)
	manager := effectservices.NewCombatEffectManager(effectManager)
	ctx := context.Background()
	actorID := "test_actor"
	combatID := "combat_123"

	// Start combat first
	err := manager.StartCombat(ctx, combatID, []string{actorID, "attacker_456"})
	if err != nil {
		t.Errorf("Expected no error starting combat, got %v", err)
	}

	// Create and apply a test combat effect with short duration
	effect := effects.CombatInjuryEffect()
	effect.ID = fmt.Sprintf("combat_injury_%d", time.Now().UnixNano())
	effect.SetDuration(1 * int64(time.Second.Seconds()))
	effect.SetCombatInfo(combatID, "attacker_456", actorID, time.Now().Unix())
	effect.Activate() // Activate the effect

	err = manager.ApplyCombatEffect(ctx, actorID, effect)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Process effects immediately
	err = manager.ProcessCombatEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Wait for effect to expire
	time.Sleep(2 * time.Second)

	// Process effects again
	err = manager.ProcessCombatEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the effect was removed
	activeEffects, err := manager.GetActiveCombatEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 0 {
		t.Errorf("Expected 0 active combat effects, got %d", len(activeEffects))
	}
}

func TestCombatEffectManager_ClearAllCombatEffects(t *testing.T) {
	// Create a mock effect manager
	effectManager := effectservices.NewEffectManager(nil)
	manager := effectservices.NewCombatEffectManager(effectManager)
	ctx := context.Background()
	actorID := "test_actor"
	combatID := "combat_123"

	// Start combat first
	err := manager.StartCombat(ctx, combatID, []string{actorID, "attacker_456"})
	if err != nil {
		t.Errorf("Expected no error starting combat, got %v", err)
	}

	// Create and apply multiple combat effects
	effect1 := effects.CombatInjuryEffect()
	effect1.ID = fmt.Sprintf("combat_injury_%d", time.Now().UnixNano())
	effect1.SetCombatInfo(combatID, "attacker_456", actorID, time.Now().Unix())
	effect1.Activate() // Activate the effect

	effect2 := effects.PoisonedEffect()
	effect2.ID = fmt.Sprintf("poisoned_%d", time.Now().UnixNano())
	effect2.SetCombatInfo(combatID, "attacker_456", actorID, time.Now().Unix())
	effect2.Activate() // Activate the effect

	err = manager.ApplyCombatEffect(ctx, actorID, effect1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = manager.ApplyCombatEffect(ctx, actorID, effect2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Clear all combat effects
	err = manager.ClearAllCombatEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify all effects were removed
	activeEffects, err := manager.GetActiveCombatEffects(ctx, actorID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(activeEffects) != 0 {
		t.Errorf("Expected 0 active combat effects, got %d", len(activeEffects))
	}
}
