package effects

import "time"

// CombatInjuryEffect creates a combat injury effect
func CombatInjuryEffect() *CombatEffect {
	effect := NewCombatEffect(
		"combat_injury",
		"Combat Injury",
		NonCombatEffectTypeStatModifier,
		EffectCategoryDebuff,
	)

	effect.SetDuration(24 * int64(time.Hour.Seconds()))
	effect.SetIntensity(0.8)
	effect.SetStackable(true, 3)
	effect.SetCombatIntensity(0.8)
	effect.SetPersistAfterCombat(true)

	// Add modifiers
	effect.NonCombatEffect.AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "primary",
		Stat:       "health",
		Multiplier: 0.7,
	})

	effect.NonCombatEffect.AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "derived",
		Stat:       "movement_speed",
		Multiplier: 0.5,
	})

	return effect
}

// PoisonedEffect creates a poisoned effect
func PoisonedEffect() *CombatEffect {
	effect := NewCombatEffect(
		"poisoned",
		"Poisoned",
		NonCombatEffectTypeDamageOverTime,
		EffectCategoryDebuff,
	)

	effect.SetDuration(2 * int64(time.Hour.Seconds()))
	effect.SetIntensity(0.6)
	effect.SetStackable(true, 5)
	effect.SetCombatIntensity(0.6)
	effect.SetPersistAfterCombat(true)

	// Add modifiers
	effect.NonCombatEffect.AddModifier(NonCombatEffectModifier{
		Type:     AdditionModifier,
		Target:   "primary",
		Stat:     "health",
		Addition: -10.0,
	})

	return effect
}

// FearEffect creates a fear effect
func FearEffect() *CombatEffect {
	effect := NewCombatEffect(
		"fear",
		"Fear",
		NonCombatEffectTypeControl,
		EffectCategoryDebuff,
	)

	effect.SetDuration(30 * int64(time.Minute.Seconds()))
	effect.SetIntensity(0.9)
	effect.SetStackable(false, 1)
	effect.SetCombatIntensity(0.9)
	effect.SetPersistAfterCombat(true)

	// Add modifiers
	effect.NonCombatEffect.AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "derived",
		Stat:       "attack",
		Multiplier: 0.6,
	})

	effect.NonCombatEffect.AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "derived",
		Stat:       "defense",
		Multiplier: 0.8,
	})

	return effect
}

// BerserkEffect creates a berserk effect
func BerserkEffect() *CombatEffect {
	effect := NewCombatEffect(
		"berserk",
		"Berserk",
		NonCombatEffectTypeBuff,
		EffectCategoryBuff,
	)

	effect.SetDuration(10 * int64(time.Minute.Seconds()))
	effect.SetIntensity(1.0)
	effect.SetStackable(false, 1)
	effect.SetCombatOnly(true)
	effect.SetCombatIntensity(1.0)
	effect.SetPersistAfterCombat(false)

	// Add modifiers
	effect.NonCombatEffect.AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "derived",
		Stat:       "attack",
		Multiplier: 1.5,
	})

	effect.NonCombatEffect.AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "derived",
		Stat:       "movement_speed",
		Multiplier: 1.3,
	})

	effect.NonCombatEffect.AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "derived",
		Stat:       "defense",
		Multiplier: 0.7,
	})

	return effect
}

// StunEffect creates a stun effect
func StunEffect() *CombatEffect {
	effect := NewCombatEffect(
		"stun",
		"Stun",
		NonCombatEffectTypeControl,
		EffectCategoryDebuff,
	)

	effect.SetDuration(5 * int64(time.Minute.Seconds()))
	effect.SetIntensity(1.0)
	effect.SetStackable(false, 1)
	effect.SetCombatOnly(true)
	effect.SetCombatIntensity(1.0)
	effect.SetPersistAfterCombat(false)

	// Add modifiers
	effect.NonCombatEffect.AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "derived",
		Stat:       "movement_speed",
		Multiplier: 0.0,
	})

	return effect
}

// BleedingEffect creates a bleeding effect
func BleedingEffect() *CombatEffect {
	effect := NewCombatEffect(
		"bleeding",
		"Bleeding",
		NonCombatEffectTypeDamageOverTime,
		EffectCategoryDebuff,
	)

	effect.SetDuration(1 * int64(time.Hour.Seconds()))
	effect.SetIntensity(0.7)
	effect.SetStackable(true, 3)
	effect.SetCombatIntensity(0.7)
	effect.SetPersistAfterCombat(true)

	// Add modifiers
	effect.NonCombatEffect.AddModifier(NonCombatEffectModifier{
		Type:     AdditionModifier,
		Target:   "primary",
		Stat:     "health",
		Addition: -5.0,
	})

	return effect
}
