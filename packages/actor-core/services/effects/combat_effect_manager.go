package effects

import (
	"actor-core-v2/models/effects"
	"context"
	"fmt"
	"sync"
	"time"
)

// CombatEffectManager manages combat effects and their transition to non-combat effects
type CombatEffectManager struct {
	effectManager *EffectManagerImpl
	combatEffects map[string]*effects.CombatEffect // combatID -> effect
	combatStates  map[string]*CombatState          // combatID -> state
	mutex         sync.RWMutex
	version       int64
}

// CombatState tracks the state of a combat
type CombatState struct {
	CombatID     string   `json:"combat_id"`
	IsActive     bool     `json:"is_active"`
	StartTime    int64    `json:"start_time"`
	EndTime      int64    `json:"end_time"`
	Participants []string `json:"participants"`
	Version      int64    `json:"version"`
	LastUpdated  int64    `json:"last_updated"`
}

// NewCombatEffectManager creates a new combat effect manager
func NewCombatEffectManager(effectManager *EffectManagerImpl) *CombatEffectManager {
	return &CombatEffectManager{
		effectManager: effectManager,
		combatEffects: make(map[string]*effects.CombatEffect),
		combatStates:  make(map[string]*CombatState),
		version:       1,
	}
}

// StartCombat starts a new combat
func (cem *CombatEffectManager) StartCombat(ctx context.Context, combatID string, participants []string) error {
	cem.mutex.Lock()
	defer cem.mutex.Unlock()

	now := time.Now().Unix()
	combatState := &CombatState{
		CombatID:     combatID,
		IsActive:     true,
		StartTime:    now,
		EndTime:      0,
		Participants: participants,
		Version:      1,
		LastUpdated:  now,
	}

	cem.combatStates[combatID] = combatState
	cem.version++

	return nil
}

// EndCombat ends a combat and processes effect transitions
func (cem *CombatEffectManager) EndCombat(ctx context.Context, combatID string) error {
	cem.mutex.Lock()
	defer cem.mutex.Unlock()

	combatState, exists := cem.combatStates[combatID]
	if !exists {
		return fmt.Errorf("combat %s not found", combatID)
	}

	if !combatState.IsActive {
		return fmt.Errorf("combat %s is already ended", combatID)
	}

	now := time.Now().Unix()
	combatState.IsActive = false
	combatState.EndTime = now
	combatState.Version++
	combatState.LastUpdated = now

	// Process all combat effects for this combat
	for _, combatEffect := range cem.combatEffects {
		if combatEffect.CombatID == combatID {
			// Set combat end time
			combatEffect.SetCombatEndTime(now)

			// Convert to non-combat effect if it should persist
			if combatEffect.ShouldPersistAfterCombat() {
				nonCombatEffect := combatEffect.ConvertToNonCombatEffect()

				// Apply to each participant
				for _, participantID := range combatState.Participants {
					err := cem.effectManager.ApplyEffect(ctx, participantID, nonCombatEffect)
					if err != nil {
						// Log error but continue processing other effects
						fmt.Printf("Error applying non-combat effect to participant %s: %v\n", participantID, err)
					}
				}
			}
		}
	}

	cem.version++
	return nil
}

// ApplyCombatEffect applies a combat effect
func (cem *CombatEffectManager) ApplyCombatEffect(ctx context.Context, actorID string, combatEffect *effects.CombatEffect) error {
	cem.mutex.Lock()
	defer cem.mutex.Unlock()

	// Check if combat is active
	combatState, exists := cem.combatStates[combatEffect.CombatID]
	if !exists {
		return fmt.Errorf("combat %s not found", combatEffect.CombatID)
	}

	if !combatState.IsActive {
		return fmt.Errorf("combat %s is not active", combatEffect.CombatID)
	}

	// Check if effect already exists for this actor (by type and target)
	for _, existingEffect := range cem.combatEffects {
		if existingEffect.TargetID == actorID && existingEffect.Type == combatEffect.Type {
			// Check if effect is stackable
			if combatEffect.Stackable {
				// Check if we can add more stacks
				if existingEffect.Stacks >= combatEffect.MaxStacks {
					return fmt.Errorf("effect %s has reached maximum stacks (%d)", combatEffect.Type, combatEffect.MaxStacks)
				}
				// Increment stacks
				existingEffect.Stacks++
				cem.version++
				return nil
			} else {
				// Non-stackable effect already exists
				return fmt.Errorf("non-stackable effect %s already exists for actor %s", combatEffect.Type, actorID)
			}
		}
	}

	// Store new combat effect
	cem.combatEffects[combatEffect.ID] = combatEffect

	// Apply to target if it's a non-combat effect that should work outside combat
	if !combatEffect.CombatOnly {
		nonCombatEffect := combatEffect.ConvertToNonCombatEffect()
		err := cem.effectManager.ApplyEffect(ctx, combatEffect.TargetID, nonCombatEffect)
		if err != nil {
			return fmt.Errorf("error applying non-combat effect: %v", err)
		}
	}

	cem.version++
	return nil
}

// GetCombatEffects returns all combat effects for a specific combat
func (cem *CombatEffectManager) GetCombatEffects(ctx context.Context, combatID string) ([]*effects.CombatEffect, error) {
	cem.mutex.RLock()
	defer cem.mutex.RUnlock()

	var combatEffects []*effects.CombatEffect
	for _, effect := range cem.combatEffects {
		if effect.CombatID == combatID {
			combatEffects = append(combatEffects, effect)
		}
	}

	return combatEffects, nil
}

// GetCombatEffect returns a specific combat effect
func (cem *CombatEffectManager) GetCombatEffect(ctx context.Context, effectID string) (*effects.CombatEffect, error) {
	cem.mutex.RLock()
	defer cem.mutex.RUnlock()

	effect, exists := cem.combatEffects[effectID]
	if !exists {
		return nil, fmt.Errorf("combat effect %s not found", effectID)
	}

	return effect, nil
}

// GetCombatState returns the state of a combat
func (cem *CombatEffectManager) GetCombatState(ctx context.Context, combatID string) (*CombatState, error) {
	cem.mutex.RLock()
	defer cem.mutex.RUnlock()

	state, exists := cem.combatStates[combatID]
	if !exists {
		return nil, fmt.Errorf("combat %s not found", combatID)
	}

	return state, nil
}

// ProcessAllCombatEffects processes all combat effects
func (cem *CombatEffectManager) ProcessAllCombatEffects(ctx context.Context) error {
	cem.mutex.Lock()
	defer cem.mutex.Unlock()

	// Process all active combats
	for combatID, combatState := range cem.combatStates {
		if combatState.IsActive {
			// Process each participant's effects
			for _, participantID := range combatState.Participants {
				err := cem.effectManager.ProcessEffects(ctx, participantID)
				if err != nil {
					// Log error but continue processing other participants
					fmt.Printf("Error processing effects for participant %s in combat %s: %v\n", participantID, combatID, err)
				}
			}
		}
	}

	cem.version++
	return nil
}

// GetCombatStats returns statistics about combat effects
func (cem *CombatEffectManager) GetCombatStats(ctx context.Context) (*CombatStats, error) {
	cem.mutex.RLock()
	defer cem.mutex.RUnlock()

	stats := &CombatStats{
		TotalCombats:       len(cem.combatStates),
		ActiveCombats:      0,
		EndedCombats:       0,
		TotalCombatEffects: len(cem.combatEffects),
		EffectsByCombat:    make(map[string]int),
		Version:            cem.version,
		LastProcessed:      time.Now().Unix(),
	}

	// Count active and ended combats
	for _, combatState := range cem.combatStates {
		if combatState.IsActive {
			stats.ActiveCombats++
		} else {
			stats.EndedCombats++
		}
	}

	// Count effects by combat
	for _, effect := range cem.combatEffects {
		stats.EffectsByCombat[effect.CombatID]++
	}

	return stats, nil
}

// ClearCombatEffects removes all combat effects for a specific combat
func (cem *CombatEffectManager) ClearCombatEffects(ctx context.Context, combatID string) error {
	cem.mutex.Lock()
	defer cem.mutex.Unlock()

	var toRemove []string
	for effectID, effect := range cem.combatEffects {
		if effect.CombatID == combatID {
			toRemove = append(toRemove, effectID)
		}
	}

	// Remove combat effects
	for _, effectID := range toRemove {
		delete(cem.combatEffects, effectID)
	}

	// Remove from participants' non-combat effects
	combatState, exists := cem.combatStates[combatID]
	if exists {
		for _, participantID := range combatState.Participants {
			for _, effectID := range toRemove {
				err := cem.effectManager.RemoveEffect(ctx, participantID, effectID)
				if err != nil {
					// Log error but continue
					fmt.Printf("Error removing effect %s from participant %s: %v\n", effectID, participantID, err)
				}
			}
		}
	}

	cem.version++
	return nil
}

// CombatStats contains statistics about combat effects
type CombatStats struct {
	TotalCombats       int            `json:"total_combats"`
	ActiveCombats      int            `json:"active_combats"`
	EndedCombats       int            `json:"ended_combats"`
	TotalCombatEffects int            `json:"total_combat_effects"`
	EffectsByCombat    map[string]int `json:"effects_by_combat"`
	Version            int64          `json:"version"`
	LastProcessed      int64          `json:"last_processed"`
}

// GetActiveCombatEffects returns all active combat effects for an actor
func (cem *CombatEffectManager) GetActiveCombatEffects(ctx context.Context, actorID string) ([]*effects.CombatEffect, error) {
	cem.mutex.RLock()
	defer cem.mutex.RUnlock()

	var activeEffects []*effects.CombatEffect
	for _, effect := range cem.combatEffects {
		if effect.TargetID == actorID && effect.IsActive {
			activeEffects = append(activeEffects, effect)
		}
	}

	return activeEffects, nil
}

// RemoveCombatEffect removes a combat effect by ID
func (cem *CombatEffectManager) RemoveCombatEffect(ctx context.Context, actorID string, effectID string) error {
	cem.mutex.Lock()
	defer cem.mutex.Unlock()

	effect, exists := cem.combatEffects[effectID]
	if !exists {
		return fmt.Errorf("combat effect %s not found", effectID)
	}

	if effect.TargetID != actorID {
		return fmt.Errorf("combat effect %s does not belong to actor %s", effectID, actorID)
	}

	delete(cem.combatEffects, effectID)
	cem.version++

	return nil
}

// RemoveCombatEffectByType removes all combat effects of a specific type for an actor
func (cem *CombatEffectManager) RemoveCombatEffectByType(ctx context.Context, actorID string, effectType string) error {
	cem.mutex.Lock()
	defer cem.mutex.Unlock()

	var toRemove []string
	for effectID, effect := range cem.combatEffects {
		if effect.TargetID == actorID && string(effect.Type) == effectType {
			toRemove = append(toRemove, effectID)
		}
	}

	for _, effectID := range toRemove {
		delete(cem.combatEffects, effectID)
	}

	if len(toRemove) > 0 {
		cem.version++
	}

	return nil
}

// GetCombatEffectByID returns a specific combat effect by ID for an actor
func (cem *CombatEffectManager) GetCombatEffectByID(ctx context.Context, actorID string, effectID string) (*effects.CombatEffect, error) {
	cem.mutex.RLock()
	defer cem.mutex.RUnlock()

	effect, exists := cem.combatEffects[effectID]
	if !exists {
		return nil, fmt.Errorf("combat effect %s not found", effectID)
	}

	if effect.TargetID != actorID {
		return nil, fmt.Errorf("combat effect %s does not belong to actor %s", effectID, actorID)
	}

	return effect, nil
}

// GetCombatEffectsByType returns all combat effects of a specific type for an actor
func (cem *CombatEffectManager) GetCombatEffectsByType(ctx context.Context, actorID string, effectType string) ([]*effects.CombatEffect, error) {
	cem.mutex.RLock()
	defer cem.mutex.RUnlock()

	var effects []*effects.CombatEffect
	for _, effect := range cem.combatEffects {
		if effect.TargetID == actorID && string(effect.Type) == effectType {
			effects = append(effects, effect)
		}
	}

	return effects, nil
}

// GetCombatEffectsByCategory returns all combat effects of a specific category for an actor
func (cem *CombatEffectManager) GetCombatEffectsByCategory(ctx context.Context, actorID string, category effects.EffectCategory) ([]*effects.CombatEffect, error) {
	cem.mutex.RLock()
	defer cem.mutex.RUnlock()

	var effects []*effects.CombatEffect
	for _, effect := range cem.combatEffects {
		if effect.TargetID == actorID && effect.Category == category {
			effects = append(effects, effect)
		}
	}

	return effects, nil
}

// ProcessCombatEffects processes all combat effects for an actor
func (cem *CombatEffectManager) ProcessCombatEffects(ctx context.Context, actorID string) error {
	cem.mutex.Lock()
	defer cem.mutex.Unlock()

	var toRemove []string

	for effectID, effect := range cem.combatEffects {
		if effect.TargetID == actorID {
			if effect.IsExpired() {
				toRemove = append(toRemove, effectID)
			}
		}
	}

	// Remove expired effects
	for _, effectID := range toRemove {
		delete(cem.combatEffects, effectID)
	}

	if len(toRemove) > 0 {
		cem.version++
	}

	return nil
}

// ClearAllCombatEffects removes all combat effects for an actor
func (cem *CombatEffectManager) ClearAllCombatEffects(ctx context.Context, actorID string) error {
	cem.mutex.Lock()
	defer cem.mutex.Unlock()

	var toRemove []string
	for effectID, effect := range cem.combatEffects {
		if effect.TargetID == actorID {
			toRemove = append(toRemove, effectID)
		}
	}

	for _, effectID := range toRemove {
		delete(cem.combatEffects, effectID)
	}

	if len(toRemove) > 0 {
		cem.version++
	}

	return nil
}

// GetCombatEffectStats returns statistics about combat effects for an actor
func (cem *CombatEffectManager) GetCombatEffectStats(ctx context.Context, actorID string) (*CombatStats, error) {
	cem.mutex.RLock()
	defer cem.mutex.RUnlock()

	stats := &CombatStats{
		TotalCombats:       0,
		ActiveCombats:      0,
		EndedCombats:       0,
		TotalCombatEffects: 0,
		EffectsByCombat:    make(map[string]int),
		Version:            cem.version,
		LastProcessed:      time.Now().Unix(),
	}

	for _, effect := range cem.combatEffects {
		if effect.TargetID == actorID {
			stats.TotalCombatEffects++
		}
	}

	return stats, nil
}
