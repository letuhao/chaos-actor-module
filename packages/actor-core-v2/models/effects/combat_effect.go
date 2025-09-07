package effects

// CombatEffect extends NonCombatEffect with combat-specific properties
type CombatEffect struct {
	*NonCombatEffect

	// Combat-specific properties
	CombatID        string `json:"combat_id"`         // ID of the combat where this effect was applied
	AttackerID      string `json:"attacker_id"`       // ID of the actor who applied this effect
	TargetID        string `json:"target_id"`         // ID of the actor who received this effect
	CombatStartTime int64  `json:"combat_start_time"` // When the combat started
	CombatEndTime   int64  `json:"combat_end_time"`   // When the combat ended (0 if still active)

	// Combat-specific behavior
	PersistAfterCombat bool    `json:"persist_after_combat"` // Whether this effect persists after combat ends
	CombatOnly         bool    `json:"combat_only"`          // Whether this effect only works in combat
	CombatIntensity    float64 `json:"combat_intensity"`     // Intensity multiplier during combat

	// Combat-specific modifiers
	CombatModifiers    []CombatEffectModifier    `json:"combat_modifiers"`     // Modifiers that only apply during combat
	NonCombatModifiers []NonCombatEffectModifier `json:"non_combat_modifiers"` // Modifiers that apply outside combat
}

// CombatEffectModifier defines how an effect modifies stats during combat
type CombatEffectModifier struct {
	Type       ModifierType      `json:"type"`
	Target     string            `json:"target"`     // "primary", "derived", "custom"
	Stat       string            `json:"stat"`       // Stat name
	Value      float64           `json:"value"`      // Base value for the modifier
	Multiplier float64           `json:"multiplier"` // Multiplier to apply
	Addition   float64           `json:"addition"`   // Addition to apply
	Conditions []EffectCondition `json:"conditions"`

	// Combat-specific properties
	CombatOnly         bool `json:"combat_only"`          // Only applies during combat
	PersistAfterCombat bool `json:"persist_after_combat"` // Persists after combat ends
}

// NewCombatEffect creates a new combat effect
func NewCombatEffect(id, name string, effectType NonCombatEffectType, category EffectCategory) *CombatEffect {
	baseEffect := NewNonCombatEffect(id, name, effectType, category)

	return &CombatEffect{
		NonCombatEffect:    baseEffect,
		CombatID:           "",
		AttackerID:         "",
		TargetID:           "",
		CombatStartTime:    0,
		CombatEndTime:      0,
		PersistAfterCombat: true,  // Default to persisting after combat
		CombatOnly:         false, // Default to working outside combat too
		CombatIntensity:    1.0,   // Default intensity multiplier
		CombatModifiers:    []CombatEffectModifier{},
		NonCombatModifiers: []NonCombatEffectModifier{},
	}
}

// SetCombatInfo sets combat-specific information
func (ce *CombatEffect) SetCombatInfo(combatID, attackerID, targetID string, combatStartTime int64) *CombatEffect {
	ce.CombatID = combatID
	ce.AttackerID = attackerID
	ce.TargetID = targetID
	ce.CombatStartTime = combatStartTime
	return ce
}

// SetCombatEndTime sets when the combat ended
func (ce *CombatEffect) SetCombatEndTime(combatEndTime int64) *CombatEffect {
	ce.CombatEndTime = combatEndTime
	return ce
}

// SetPersistAfterCombat sets whether this effect persists after combat
func (ce *CombatEffect) SetPersistAfterCombat(persist bool) *CombatEffect {
	ce.PersistAfterCombat = persist
	return ce
}

// SetCombatOnly sets whether this effect only works in combat
func (ce *CombatEffect) SetCombatOnly(combatOnly bool) *CombatEffect {
	ce.CombatOnly = combatOnly
	return ce
}

// SetCombatIntensity sets the intensity multiplier during combat
func (ce *CombatEffect) SetCombatIntensity(intensity float64) *CombatEffect {
	ce.CombatIntensity = intensity
	return ce
}

// AddCombatModifier adds a combat-specific modifier
func (ce *CombatEffect) AddCombatModifier(modifier CombatEffectModifier) *CombatEffect {
	ce.CombatModifiers = append(ce.CombatModifiers, modifier)
	return ce
}

// AddNonCombatModifier adds a non-combat modifier
func (ce *CombatEffect) AddNonCombatModifier(modifier NonCombatEffectModifier) *CombatEffect {
	ce.NonCombatModifiers = append(ce.NonCombatModifiers, modifier)
	return ce
}

// IsInCombat checks if this effect is currently in combat
func (ce *CombatEffect) IsInCombat() bool {
	return ce.CombatEndTime == 0
}

// IsCombatEnded checks if the combat has ended
func (ce *CombatEffect) IsCombatEnded() bool {
	return ce.CombatEndTime > 0
}

// GetEffectiveIntensity returns the effective intensity based on combat state
func (ce *CombatEffect) GetEffectiveIntensity() float64 {
	if ce.IsInCombat() {
		return ce.Intensity * ce.CombatIntensity
	}
	return ce.Intensity
}

// GetEffectiveModifiers returns the effective modifiers based on combat state
func (ce *CombatEffect) GetEffectiveModifiers() []NonCombatEffectModifier {
	var effectiveModifiers []NonCombatEffectModifier

	// Add non-combat modifiers
	effectiveModifiers = append(effectiveModifiers, ce.NonCombatModifiers...)

	// Add combat modifiers if in combat
	if ce.IsInCombat() {
		for _, combatMod := range ce.CombatModifiers {
			// Convert combat modifier to non-combat modifier
			nonCombatMod := NonCombatEffectModifier{
				Type:       combatMod.Type,
				Target:     combatMod.Target,
				Stat:       combatMod.Stat,
				Value:      combatMod.Value,
				Multiplier: combatMod.Multiplier,
				Addition:   combatMod.Addition,
				Conditions: combatMod.Conditions,
			}
			effectiveModifiers = append(effectiveModifiers, nonCombatMod)
		}
	}

	return effectiveModifiers
}

// ShouldPersistAfterCombat checks if this effect should persist after combat
func (ce *CombatEffect) ShouldPersistAfterCombat() bool {
	return ce.PersistAfterCombat && !ce.CombatOnly
}

// ConvertToNonCombatEffect converts this combat effect to a non-combat effect
func (ce *CombatEffect) ConvertToNonCombatEffect() *NonCombatEffect {
	// Create a new non-combat effect with the same base properties
	nonCombatEffect := NewNonCombatEffect(ce.ID, ce.Name, ce.Type, ce.Category)
	nonCombatEffect.Duration = ce.Duration
	nonCombatEffect.Intensity = ce.Intensity
	nonCombatEffect.Stackable = ce.Stackable
	nonCombatEffect.MaxStacks = ce.MaxStacks
	nonCombatEffect.Source = ce.Source
	nonCombatEffect.StartTime = ce.StartTime
	nonCombatEffect.EndTime = ce.EndTime
	nonCombatEffect.IsActive = ce.IsActive
	nonCombatEffect.Conditions = ce.Conditions

	// Add non-combat modifiers
	nonCombatEffect.Effects = append(nonCombatEffect.Effects, ce.NonCombatModifiers...)

	// Add combat modifiers if they should persist
	for _, combatMod := range ce.CombatModifiers {
		if combatMod.PersistAfterCombat {
			nonCombatMod := NonCombatEffectModifier{
				Type:       combatMod.Type,
				Target:     combatMod.Target,
				Stat:       combatMod.Stat,
				Value:      combatMod.Value,
				Multiplier: combatMod.Multiplier,
				Addition:   combatMod.Addition,
				Conditions: combatMod.Conditions,
			}
			nonCombatEffect.Effects = append(nonCombatEffect.Effects, nonCombatMod)
		}
	}

	return nonCombatEffect
}

// Clone creates a deep copy of the combat effect
func (ce *CombatEffect) Clone() *CombatEffect {
	clone := &CombatEffect{
		NonCombatEffect:    ce.NonCombatEffect.Clone(),
		CombatID:           ce.CombatID,
		AttackerID:         ce.AttackerID,
		TargetID:           ce.TargetID,
		CombatStartTime:    ce.CombatStartTime,
		CombatEndTime:      ce.CombatEndTime,
		PersistAfterCombat: ce.PersistAfterCombat,
		CombatOnly:         ce.CombatOnly,
		CombatIntensity:    ce.CombatIntensity,
		CombatModifiers:    make([]CombatEffectModifier, len(ce.CombatModifiers)),
		NonCombatModifiers: make([]NonCombatEffectModifier, len(ce.NonCombatModifiers)),
	}

	// Copy combat modifiers
	copy(clone.CombatModifiers, ce.CombatModifiers)

	// Copy non-combat modifiers
	copy(clone.NonCombatModifiers, ce.NonCombatModifiers)

	return clone
}
