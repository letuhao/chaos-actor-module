package effects

import (
	"time"
)

// NonCombatEffect represents a non-combat effect that can be applied to an actor
type NonCombatEffect struct {
	ID         string                    `json:"id"`
	Name       string                    `json:"name"`
	Type       NonCombatEffectType       `json:"type"`
	Category   EffectCategory            `json:"category"`
	Duration   int64                     `json:"duration"`   // Duration in seconds
	Intensity  float64                   `json:"intensity"`  // Effect intensity multiplier
	Stackable  bool                      `json:"stackable"`  // Can this effect stack?
	MaxStacks  int                       `json:"max_stacks"` // Maximum number of stacks
	Stacks     int                       `json:"stacks"`     // Current number of stacks
	Effects    []NonCombatEffectModifier `json:"effects"`    // What this effect modifies
	Conditions []EffectCondition         `json:"conditions"` // Conditions for activation
	Source     string                    `json:"source"`     // What applied this effect
	StartTime  int64                     `json:"start_time"` // When effect started (Unix timestamp)
	EndTime    int64                     `json:"end_time"`   // When effect ends (Unix timestamp)
	IsActive   bool                      `json:"is_active"`  // Is effect currently active
	Version    int64                     `json:"version"`    // Version for tracking changes
	UpdatedAt  int64                     `json:"updated_at"` // Last update timestamp
}

// NonCombatEffectType defines the type of non-combat effect
type NonCombatEffectType string

const (
	// Cultivation Effects
	CultivationBoostEffect NonCombatEffectType = "cultivation_boost"
	BreakthroughEffect     NonCombatEffectType = "breakthrough"
	QiRefinementEffect     NonCombatEffectType = "qi_refinement"
	SpiritualInsightEffect NonCombatEffectType = "spiritual_insight"

	// Learning Effects
	LearningBoostEffect     NonCombatEffectType = "learning_boost"
	MemoryEnhancementEffect NonCombatEffectType = "memory_enhancement"
	FocusEffect             NonCombatEffectType = "focus"
	ConcentrationEffect     NonCombatEffectType = "concentration"

	// Social Effects
	CharismaBoostEffect NonCombatEffectType = "charisma_boost"
	LeadershipEffect    NonCombatEffectType = "leadership"
	DiplomacyEffect     NonCombatEffectType = "diplomacy"
	ReputationEffect    NonCombatEffectType = "reputation"

	// Crafting Effects
	CraftingMasteryEffect NonCombatEffectType = "crafting_mastery"
	AlchemyBoostEffect    NonCombatEffectType = "alchemy_boost"
	ForgingBoostEffect    NonCombatEffectType = "forging_boost"
	EnchantingBoostEffect NonCombatEffectType = "enchanting_boost"

	// Movement Effects
	SpeedBoostEffect    NonCombatEffectType = "speed_boost"
	FlightEffect        NonCombatEffectType = "flight"
	TeleportationEffect NonCombatEffectType = "teleportation"
	StealthEffect       NonCombatEffectType = "stealth"

	// Environmental Effects
	WeatherAdaptationEffect NonCombatEffectType = "weather_adaptation"
	TerrainMasteryEffect    NonCombatEffectType = "terrain_mastery"
	ClimateResistanceEffect NonCombatEffectType = "climate_resistance"

	// Mystical Effects
	FateManipulationEffect NonCombatEffectType = "fate_manipulation"
	KarmaInfluenceEffect   NonCombatEffectType = "karma_influence"
	LuckBoostEffect        NonCombatEffectType = "luck_boost"
	FortuneEffect          NonCombatEffectType = "fortune"

	// General Effect Types
	NonCombatEffectTypeStatModifier   NonCombatEffectType = "stat_modifier"
	NonCombatEffectTypeBuff           NonCombatEffectType = "buff"
	NonCombatEffectTypeDamageOverTime NonCombatEffectType = "damage_over_time"
	NonCombatEffectTypeControl        NonCombatEffectType = "control"
)

// EffectCategory defines the category of effect
type EffectCategory string

const (
	CultivationCategory   EffectCategory = "cultivation"
	LearningCategory      EffectCategory = "learning"
	SocialCategory        EffectCategory = "social"
	CraftingCategory      EffectCategory = "crafting"
	MovementCategory      EffectCategory = "movement"
	EnvironmentalCategory EffectCategory = "environmental"
	MysticalCategory      EffectCategory = "mystical"
	EffectCategoryBuff    EffectCategory = "buff"
	EffectCategoryDebuff  EffectCategory = "debuff"
)

// NonCombatEffectModifier defines how an effect modifies stats
type NonCombatEffectModifier struct {
	Type       ModifierType      `json:"type"`       // Type of modification
	Target     string            `json:"target"`     // "primary", "derived", "custom"
	Stat       string            `json:"stat"`       // Stat name to modify
	Value      float64           `json:"value"`      // Modifier value
	Multiplier float64           `json:"multiplier"` // Multiplier (default 1.0)
	Addition   float64           `json:"addition"`   // Addition (default 0.0)
	Conditions []EffectCondition `json:"conditions"` // Conditions for this modifier
}

// ModifierType defines how the effect modifies stats
type ModifierType string

const (
	MultiplierModifier ModifierType = "multiplier"
	AdditionModifier   ModifierType = "addition"
	OverrideModifier   ModifierType = "override"
	CapModifier        ModifierType = "cap"
)

// EffectCondition defines conditions for effect activation
type EffectCondition struct {
	Type     string      `json:"type"`     // Condition type
	Variable string      `json:"variable"` // Variable to check
	Operator string      `json:"operator"` // Comparison operator
	Value    interface{} `json:"value"`    // Value to compare against
}

// NewNonCombatEffect creates a new non-combat effect
func NewNonCombatEffect(id, name string, effectType NonCombatEffectType, category EffectCategory) *NonCombatEffect {
	now := time.Now().Unix()
	return &NonCombatEffect{
		ID:         id,
		Name:       name,
		Type:       effectType,
		Category:   category,
		Duration:   0,
		Intensity:  1.0,
		Stackable:  false,
		MaxStacks:  1,
		Stacks:     1,
		Effects:    []NonCombatEffectModifier{},
		Conditions: []EffectCondition{},
		Source:     "",
		StartTime:  now,
		EndTime:    now,
		IsActive:   false,
		Version:    1,
		UpdatedAt:  now,
	}
}

// SetDuration sets the duration of the effect
func (e *NonCombatEffect) SetDuration(duration int64) *NonCombatEffect {
	e.Duration = duration
	e.EndTime = e.StartTime + duration
	e.Version++
	e.UpdatedAt = time.Now().Unix()
	return e
}

// SetIntensity sets the intensity of the effect
func (e *NonCombatEffect) SetIntensity(intensity float64) *NonCombatEffect {
	e.Intensity = intensity
	e.Version++
	e.UpdatedAt = time.Now().Unix()
	return e
}

// SetStackable sets whether the effect can stack
func (e *NonCombatEffect) SetStackable(stackable bool, maxStacks int) *NonCombatEffect {
	e.Stackable = stackable
	e.MaxStacks = maxStacks
	e.Version++
	e.UpdatedAt = time.Now().Unix()
	return e
}

// AddModifier adds a modifier to the effect
func (e *NonCombatEffect) AddModifier(modifier NonCombatEffectModifier) *NonCombatEffect {
	e.Effects = append(e.Effects, modifier)
	e.Version++
	e.UpdatedAt = time.Now().Unix()
	return e
}

// AddCondition adds a condition to the effect
func (e *NonCombatEffect) AddCondition(condition EffectCondition) *NonCombatEffect {
	e.Conditions = append(e.Conditions, condition)
	e.Version++
	e.UpdatedAt = time.Now().Unix()
	return e
}

// SetSource sets the source of the effect
func (e *NonCombatEffect) SetSource(source string) *NonCombatEffect {
	e.Source = source
	e.Version++
	e.UpdatedAt = time.Now().Unix()
	return e
}

// Activate activates the effect
func (e *NonCombatEffect) Activate() *NonCombatEffect {
	now := time.Now().Unix()
	e.IsActive = true
	e.StartTime = now
	if e.Duration > 0 {
		e.EndTime = now + e.Duration
	}
	e.Version++
	e.UpdatedAt = now
	return e
}

// Deactivate deactivates the effect
func (e *NonCombatEffect) Deactivate() *NonCombatEffect {
	e.IsActive = false
	e.Version++
	e.UpdatedAt = time.Now().Unix()
	return e
}

// IsExpired checks if the effect has expired
func (e *NonCombatEffect) IsExpired() bool {
	if e.Duration <= 0 {
		return false // Permanent effect
	}
	return time.Now().Unix() >= e.EndTime
}

// GetRemainingDuration returns the remaining duration in seconds
func (e *NonCombatEffect) GetRemainingDuration() int64 {
	if e.Duration <= 0 {
		return -1 // Permanent effect
	}
	remaining := e.EndTime - time.Now().Unix()
	if remaining < 0 {
		return 0
	}
	return remaining
}

// Clone creates a deep copy of the effect
func (e *NonCombatEffect) Clone() *NonCombatEffect {
	clone := &NonCombatEffect{
		ID:        e.ID,
		Name:      e.Name,
		Type:      e.Type,
		Category:  e.Category,
		Duration:  e.Duration,
		Intensity: e.Intensity,
		Stackable: e.Stackable,
		MaxStacks: e.MaxStacks,
		Stacks:    e.Stacks,
		Source:    e.Source,
		StartTime: e.StartTime,
		EndTime:   e.EndTime,
		IsActive:  e.IsActive,
		Version:   e.Version,
		UpdatedAt: e.UpdatedAt,
	}

	// Deep copy effects
	clone.Effects = make([]NonCombatEffectModifier, len(e.Effects))
	copy(clone.Effects, e.Effects)

	// Deep copy conditions
	clone.Conditions = make([]EffectCondition, len(e.Conditions))
	copy(clone.Conditions, e.Conditions)

	return clone
}
