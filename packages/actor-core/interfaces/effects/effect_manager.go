package effects

import (
	"context"
	"time"

	"actor-core/models/effects"
)

// EffectManager defines the interface for managing non-combat effects
type EffectManager interface {
	// ApplyEffect applies a non-combat effect to an actor
	ApplyEffect(ctx context.Context, actorID string, effect *effects.NonCombatEffect) error

	// RemoveEffect removes a specific effect from an actor
	RemoveEffect(ctx context.Context, actorID string, effectID string) error

	// RemoveEffectByType removes all effects of a specific type from an actor
	RemoveEffectByType(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) error

	// GetActiveEffects returns all active effects for an actor
	GetActiveEffects(ctx context.Context, actorID string) ([]*effects.NonCombatEffect, error)

	// GetEffectByID returns a specific effect by ID
	GetEffectByID(ctx context.Context, actorID string, effectID string) (*effects.NonCombatEffect, error)

	// GetEffectsByType returns all effects of a specific type for an actor
	GetEffectsByType(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) ([]*effects.NonCombatEffect, error)

	// GetEffectsByCategory returns all effects of a specific category for an actor
	GetEffectsByCategory(ctx context.Context, actorID string, category effects.EffectCategory) ([]*effects.NonCombatEffect, error)

	// ProcessEffects processes all effects for an actor (tick effects, remove expired)
	ProcessEffects(ctx context.Context, actorID string) error

	// ProcessAllEffects processes effects for all actors
	ProcessAllEffects(ctx context.Context) error

	// SetResistance sets resistance to a specific effect type
	SetResistance(ctx context.Context, actorID string, effectType effects.NonCombatEffectType, resistance float64) error

	// GetResistance returns resistance to a specific effect type
	GetResistance(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) (float64, error)

	// SetImmunity sets immunity to a specific effect type
	SetImmunity(ctx context.Context, actorID string, effectType effects.NonCombatEffectType, immune bool) error

	// GetImmunity returns immunity to a specific effect type
	GetImmunity(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) (bool, error)

	// ClearAllEffects removes all effects from an actor
	ClearAllEffects(ctx context.Context, actorID string) error

	// GetEffectStats returns statistics about effects for an actor
	GetEffectStats(ctx context.Context, actorID string) (*EffectStats, error)
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

// EffectConfig contains configuration for the effect manager
type EffectConfig struct {
	MaxEffectsPerActor int           `json:"max_effects_per_actor"`
	ProcessInterval    time.Duration `json:"process_interval"`
	MaxDuration        int64         `json:"max_duration"` // Maximum effect duration in seconds
	EnableLogging      bool          `json:"enable_logging"`
	EnableMetrics      bool          `json:"enable_metrics"`
}

// DefaultEffectConfig returns default configuration
func DefaultEffectConfig() *EffectConfig {
	return &EffectConfig{
		MaxEffectsPerActor: 100,
		ProcessInterval:    time.Minute,
		MaxDuration:        86400, // 24 hours
		EnableLogging:      true,
		EnableMetrics:      true,
	}
}
