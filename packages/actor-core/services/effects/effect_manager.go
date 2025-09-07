package effects

import (
	"context"
	"fmt"
	"sync"
	"time"

	"actor-core-v2/models/effects"
)

// EffectConfig contains configuration for the effect manager
type EffectConfig struct {
	MaxEffectsPerActor int `json:"max_effects_per_actor"`
}

// DefaultEffectConfig returns default effect configuration
func DefaultEffectConfig() *EffectConfig {
	return &EffectConfig{
		MaxEffectsPerActor: 50,
	}
}

// EffectStats contains statistics about effects
type EffectStats struct {
	ActorID           string                              `json:"actor_id"`
	TotalEffects      int                                 `json:"total_effects"`
	ActiveEffects     int                                 `json:"active_effects"`
	ExpiredEffects    int                                 `json:"expired_effects"`
	EffectsByType     map[effects.NonCombatEffectType]int `json:"effects_by_type"`
	EffectsByCategory map[effects.EffectCategory]int      `json:"effects_by_category"`
	LastProcessed     int64                               `json:"last_processed"`
	Version           int64                               `json:"version"`
}

// EffectManagerImpl implements the EffectManager interface
type EffectManagerImpl struct {
	actors        map[string]*ActorEffects                           // actorID -> effects
	resistances   map[string]map[effects.NonCombatEffectType]float64 // actorID -> effectType -> resistance
	immunities    map[string]map[effects.NonCombatEffectType]bool    // actorID -> effectType -> immune
	config        *EffectConfig
	mutex         sync.RWMutex
	version       int64
	lastProcessed int64
}

// ActorEffects contains all effects for a specific actor
type ActorEffects struct {
	ActorID     string                                  `json:"actor_id"`
	Effects     map[string]*effects.NonCombatEffect     `json:"effects"`     // effectID -> effect
	Resistances map[effects.NonCombatEffectType]float64 `json:"resistances"` // effectType -> resistance
	Immunities  map[effects.NonCombatEffectType]bool    `json:"immunities"`  // effectType -> immune
	Version     int64                                   `json:"version"`
	LastUpdated int64                                   `json:"last_updated"`
}

// NewEffectManager creates a new effect manager
func NewEffectManager(config *EffectConfig) *EffectManagerImpl {
	if config == nil {
		config = DefaultEffectConfig()
	}

	return &EffectManagerImpl{
		actors:        make(map[string]*ActorEffects),
		resistances:   make(map[string]map[effects.NonCombatEffectType]float64),
		immunities:    make(map[string]map[effects.NonCombatEffectType]bool),
		config:        config,
		version:       1,
		lastProcessed: time.Now().Unix(),
	}
}

// ApplyEffect applies a non-combat effect to an actor
func (em *EffectManagerImpl) ApplyEffect(ctx context.Context, actorID string, effect *effects.NonCombatEffect) error {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	// Get or create actor effects
	actorEffects := em.getOrCreateActorEffects(actorID)

	// Check if actor has too many effects
	if len(actorEffects.Effects) >= em.config.MaxEffectsPerActor {
		return fmt.Errorf("actor %s has too many effects (max: %d)", actorID, em.config.MaxEffectsPerActor)
	}

	// Check immunity
	if actorEffects.Immunities[effect.Type] {
		return fmt.Errorf("actor %s is immune to effect type %s", actorID, effect.Type)
	}

	// Apply resistance
	resistance := actorEffects.Resistances[effect.Type]
	if resistance > 0 {
		effect.Intensity *= (1.0 - resistance)
		effect.Duration = int64(float64(effect.Duration) * (1.0 - resistance))
	}

	// Check if effect already exists (by type)
	for _, existing := range actorEffects.Effects {
		if existing.Type == effect.Type {
			if effect.Stackable && existing.Stackable && existing.Stacks < effect.MaxStacks {
				// Stack the effect
				existing.Stacks++
				existing.Intensity += effect.Intensity
				existing.EndTime = effect.EndTime
				existing.Version++
				existing.UpdatedAt = time.Now().Unix()
				em.version++
				return nil
			} else {
				// Non-stackable effect already exists
				return fmt.Errorf("non-stackable effect %s already exists for actor %s", effect.Type, actorID)
			}
		}
	}

	// Add new effect
	effect.Activate()
	actorEffects.Effects[effect.ID] = effect

	// Update actor effects
	actorEffects.Version++
	actorEffects.LastUpdated = time.Now().Unix()
	em.version++

	return nil
}

// RemoveEffect removes a specific effect from an actor
func (em *EffectManagerImpl) RemoveEffect(ctx context.Context, actorID string, effectID string) error {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	actorEffects, exists := em.actors[actorID]
	if !exists {
		return fmt.Errorf("actor %s not found", actorID)
	}

	if _, exists := actorEffects.Effects[effectID]; !exists {
		return fmt.Errorf("effect %s not found for actor %s", effectID, actorID)
	}

	delete(actorEffects.Effects, effectID)
	actorEffects.Version++
	actorEffects.LastUpdated = time.Now().Unix()
	em.version++

	return nil
}

// RemoveEffectByType removes all effects of a specific type from an actor
func (em *EffectManagerImpl) RemoveEffectByType(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) error {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	actorEffects, exists := em.actors[actorID]
	if !exists {
		return fmt.Errorf("actor %s not found", actorID)
	}

	var toRemove []string
	for effectID, effect := range actorEffects.Effects {
		if effect.Type == effectType {
			toRemove = append(toRemove, effectID)
		}
	}

	for _, effectID := range toRemove {
		delete(actorEffects.Effects, effectID)
	}

	if len(toRemove) > 0 {
		actorEffects.Version++
		actorEffects.LastUpdated = time.Now().Unix()
		em.version++
	}

	return nil
}

// GetActiveEffects returns all active effects for an actor
func (em *EffectManagerImpl) GetActiveEffects(ctx context.Context, actorID string) ([]*effects.NonCombatEffect, error) {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	actorEffects, exists := em.actors[actorID]
	if !exists {
		return []*effects.NonCombatEffect{}, nil
	}

	var activeEffects []*effects.NonCombatEffect
	for _, effect := range actorEffects.Effects {
		if effect.IsActive && !effect.IsExpired() {
			activeEffects = append(activeEffects, effect)
		}
	}

	return activeEffects, nil
}

// GetEffectByID returns a specific effect by ID
func (em *EffectManagerImpl) GetEffectByID(ctx context.Context, actorID string, effectID string) (*effects.NonCombatEffect, error) {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	actorEffects, exists := em.actors[actorID]
	if !exists {
		return nil, fmt.Errorf("actor %s not found", actorID)
	}

	effect, exists := actorEffects.Effects[effectID]
	if !exists {
		return nil, fmt.Errorf("effect %s not found for actor %s", effectID, actorID)
	}

	return effect, nil
}

// GetEffectsByType returns all effects of a specific type for an actor
func (em *EffectManagerImpl) GetEffectsByType(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) ([]*effects.NonCombatEffect, error) {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	actorEffects, exists := em.actors[actorID]
	if !exists {
		return []*effects.NonCombatEffect{}, nil
	}

	var effects []*effects.NonCombatEffect
	for _, effect := range actorEffects.Effects {
		if effect.Type == effectType {
			effects = append(effects, effect)
		}
	}

	return effects, nil
}

// GetEffectsByCategory returns all effects of a specific category for an actor
func (em *EffectManagerImpl) GetEffectsByCategory(ctx context.Context, actorID string, category effects.EffectCategory) ([]*effects.NonCombatEffect, error) {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	actorEffects, exists := em.actors[actorID]
	if !exists {
		return []*effects.NonCombatEffect{}, nil
	}

	var effects []*effects.NonCombatEffect
	for _, effect := range actorEffects.Effects {
		if effect.Category == category {
			effects = append(effects, effect)
		}
	}

	return effects, nil
}

// ProcessEffects processes all effects for an actor
func (em *EffectManagerImpl) ProcessEffects(ctx context.Context, actorID string) error {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	actorEffects, exists := em.actors[actorID]
	if !exists {
		return nil
	}

	now := time.Now().Unix()
	var toRemove []string

	for effectID, effect := range actorEffects.Effects {
		if effect.IsExpired() {
			toRemove = append(toRemove, effectID)
		} else if effect.IsActive {
			// Process effect tick
			em.processEffectTick(effect)
		}
	}

	// Remove expired effects
	for _, effectID := range toRemove {
		delete(actorEffects.Effects, effectID)
	}

	if len(toRemove) > 0 {
		actorEffects.Version++
		actorEffects.LastUpdated = now
		em.version++
	}

	return nil
}

// ProcessAllEffects processes effects for all actors
func (em *EffectManagerImpl) ProcessAllEffects(ctx context.Context) error {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	now := time.Now().Unix()

	for _, actorEffects := range em.actors {
		var toRemove []string

		for effectID, effect := range actorEffects.Effects {
			if effect.IsExpired() {
				toRemove = append(toRemove, effectID)
			} else if effect.IsActive {
				// Process effect tick
				em.processEffectTick(effect)
			}
		}

		// Remove expired effects
		for _, effectID := range toRemove {
			delete(actorEffects.Effects, effectID)
		}

		if len(toRemove) > 0 {
			actorEffects.Version++
			actorEffects.LastUpdated = now
		}
	}

	em.lastProcessed = now
	em.version++

	return nil
}

// SetResistance sets resistance to a specific effect type
func (em *EffectManagerImpl) SetResistance(ctx context.Context, actorID string, effectType effects.NonCombatEffectType, resistance float64) error {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	actorEffects := em.getOrCreateActorEffects(actorID)
	actorEffects.Resistances[effectType] = resistance
	actorEffects.Version++
	actorEffects.LastUpdated = time.Now().Unix()
	em.version++

	return nil
}

// GetResistance returns resistance to a specific effect type
func (em *EffectManagerImpl) GetResistance(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) (float64, error) {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	actorEffects, exists := em.actors[actorID]
	if !exists {
		return 0.0, nil
	}

	return actorEffects.Resistances[effectType], nil
}

// SetImmunity sets immunity to a specific effect type
func (em *EffectManagerImpl) SetImmunity(ctx context.Context, actorID string, effectType effects.NonCombatEffectType, immune bool) error {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	actorEffects := em.getOrCreateActorEffects(actorID)
	actorEffects.Immunities[effectType] = immune
	actorEffects.Version++
	actorEffects.LastUpdated = time.Now().Unix()
	em.version++

	return nil
}

// GetImmunity returns immunity to a specific effect type
func (em *EffectManagerImpl) GetImmunity(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) (bool, error) {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	actorEffects, exists := em.actors[actorID]
	if !exists {
		return false, nil
	}

	return actorEffects.Immunities[effectType], nil
}

// ClearAllEffects removes all effects from an actor
func (em *EffectManagerImpl) ClearAllEffects(ctx context.Context, actorID string) error {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	actorEffects, exists := em.actors[actorID]
	if !exists {
		return nil
	}

	actorEffects.Effects = make(map[string]*effects.NonCombatEffect)
	actorEffects.Version++
	actorEffects.LastUpdated = time.Now().Unix()
	em.version++

	return nil
}

// GetEffectStats returns statistics about effects for an actor
func (em *EffectManagerImpl) GetEffectStats(ctx context.Context, actorID string) (*EffectStats, error) {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	actorEffects, exists := em.actors[actorID]
	if !exists {
		return &EffectStats{
			ActorID:           actorID,
			TotalEffects:      0,
			ActiveEffects:     0,
			ExpiredEffects:    0,
			EffectsByType:     make(map[effects.NonCombatEffectType]int),
			EffectsByCategory: make(map[effects.EffectCategory]int),
			LastProcessed:     em.lastProcessed,
			Version:           em.version,
		}, nil
	}

	stats := &EffectStats{
		ActorID:           actorID,
		TotalEffects:      len(actorEffects.Effects),
		ActiveEffects:     0,
		ExpiredEffects:    0,
		EffectsByType:     make(map[effects.NonCombatEffectType]int),
		EffectsByCategory: make(map[effects.EffectCategory]int),
		LastProcessed:     em.lastProcessed,
		Version:           em.version,
	}

	for _, effect := range actorEffects.Effects {
		if effect.IsActive && !effect.IsExpired() {
			stats.ActiveEffects++
		} else {
			stats.ExpiredEffects++
		}

		stats.EffectsByType[effect.Type]++
		stats.EffectsByCategory[effect.Category]++
	}

	return stats, nil
}

// Helper methods

func (em *EffectManagerImpl) getOrCreateActorEffects(actorID string) *ActorEffects {
	if actorEffects, exists := em.actors[actorID]; exists {
		return actorEffects
	}

	now := time.Now().Unix()
	actorEffects := &ActorEffects{
		ActorID:     actorID,
		Effects:     make(map[string]*effects.NonCombatEffect),
		Resistances: make(map[effects.NonCombatEffectType]float64),
		Immunities:  make(map[effects.NonCombatEffectType]bool),
		Version:     1,
		LastUpdated: now,
	}

	em.actors[actorID] = actorEffects
	return actorEffects
}

func (em *EffectManagerImpl) processEffectTick(effect *effects.NonCombatEffect) {
	// This is where we would apply the effect's modifiers to the actor's stats
	// For now, we just update the effect's version
	effect.Version++
	effect.UpdatedAt = time.Now().Unix()
}
