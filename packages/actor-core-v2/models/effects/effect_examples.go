package effects

import (
	"time"
)

// EffectExamples contains predefined non-combat effects for the game
type EffectExamples struct{}

// NewEffectExamples creates a new effect examples instance
func NewEffectExamples() *EffectExamples {
	return &EffectExamples{}
}

// === INJURY EFFECTS (Bị Thương) ===

// CreateInjuryEffect creates a basic injury effect
func (ee *EffectExamples) CreateInjuryEffect(severity string, duration int64) *NonCombatEffect {
	effect := NewNonCombatEffect(
		"injury_"+severity+"_"+time.Now().Format("20060102150405"),
		"Injury - "+severity,
		"injury",
		EnvironmentalCategory,
	).SetDuration(duration).SetIntensity(1.0).SetStackable(false, 1)

	switch severity {
	case "minor":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "speed",
			Multiplier: 0.9, // -10% speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "regen_hp",
			Multiplier: 0.8, // -20% HP regeneration
		})

	case "moderate":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "speed",
			Multiplier: 0.7, // -30% speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "regen_hp",
			Multiplier: 0.5, // -50% HP regeneration
		}).AddModifier(NonCombatEffectModifier{
			Type:     AdditionModifier,
			Target:   "derived",
			Stat:     "hp_max",
			Addition: -50, // -50 max HP
		})

	case "severe":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "speed",
			Multiplier: 0.5, // -50% speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "regen_hp",
			Multiplier: 0.2, // -80% HP regeneration
		}).AddModifier(NonCombatEffectModifier{
			Type:     AdditionModifier,
			Target:   "derived",
			Stat:     "hp_max",
			Addition: -150, // -150 max HP
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "cultivation_speed",
			Multiplier: 0.8, // -20% cultivation speed
		})
	}

	return effect
}

// CreateBrokenBoneEffect creates a broken bone effect
func (ee *EffectExamples) CreateBrokenBoneEffect(bone string, duration int64) *NonCombatEffect {
	effect := NewNonCombatEffect(
		"broken_bone_"+bone+"_"+time.Now().Format("20060102150405"),
		"Broken Bone - "+bone,
		"injury",
		EnvironmentalCategory,
	).SetDuration(duration).SetIntensity(1.0).SetStackable(false, 1)

	effect.AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "derived",
		Stat:       "speed",
		Multiplier: 0.6, // -40% speed
	}).AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "derived",
		Stat:       "strength",
		Multiplier: 0.7, // -30% strength
	}).AddModifier(NonCombatEffectModifier{
		Type:     AdditionModifier,
		Target:   "derived",
		Stat:     "hp_max",
		Addition: -100, // -100 max HP
	})

	return effect
}

// === QI DEVIATION EFFECTS (Tẩu Hỏa Nhập Ma) ===

// CreateQiDeviationEffect creates a qi deviation effect
func (ee *EffectExamples) CreateQiDeviationEffect(severity string, duration int64) *NonCombatEffect {
	effect := NewNonCombatEffect(
		"qi_deviation_"+severity+"_"+time.Now().Format("20060102150405"),
		"Qi Deviation - "+severity,
		"qi_deviation",
		CultivationCategory,
	).SetDuration(duration).SetIntensity(1.0).SetStackable(true, 3)

	switch severity {
	case "minor":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "spiritual_energy",
			Multiplier: 0.9, // -10% spiritual energy
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "cultivation_speed",
			Multiplier: 0.8, // -20% cultivation speed
		})

	case "moderate":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "spiritual_energy",
			Multiplier: 0.7, // -30% spiritual energy
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "cultivation_speed",
			Multiplier: 0.5, // -50% cultivation speed
		}).AddModifier(NonCombatEffectModifier{
			Type:     AdditionModifier,
			Target:   "derived",
			Stat:     "hp_max",
			Addition: -100, // -100 max HP
		})

	case "severe":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "spiritual_energy",
			Multiplier: 0.5, // -50% spiritual energy
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "cultivation_speed",
			Multiplier: 0.2, // -80% cultivation speed
		}).AddModifier(NonCombatEffectModifier{
			Type:     AdditionModifier,
			Target:   "derived",
			Stat:     "hp_max",
			Addition: -200, // -200 max HP
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "energy_efficiency",
			Multiplier: 0.6, // -40% energy efficiency
		})
	}

	return effect
}

// CreateMeridianBlockageEffect creates a meridian blockage effect
func (ee *EffectExamples) CreateMeridianBlockageEffect(meridian string, duration int64) *NonCombatEffect {
	effect := NewNonCombatEffect(
		"meridian_blockage_"+meridian+"_"+time.Now().Format("20060102150405"),
		"Meridian Blockage - "+meridian,
		"meridian_blockage",
		CultivationCategory,
	).SetDuration(duration).SetIntensity(1.0).SetStackable(false, 1)

	effect.AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "primary",
		Stat:       "spiritual_energy",
		Multiplier: 0.8, // -20% spiritual energy
	}).AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "derived",
		Stat:       "cultivation_speed",
		Multiplier: 0.7, // -30% cultivation speed
	}).AddModifier(NonCombatEffectModifier{
		Type:       MultiplierModifier,
		Target:     "derived",
		Stat:       "energy_efficiency",
		Multiplier: 0.8, // -20% energy efficiency
	})

	return effect
}

// === CURSE EFFECTS (Bị Nguyền Rủa) ===

// CreateCurseEffect creates a curse effect
func (ee *EffectExamples) CreateCurseEffect(curseType string, duration int64) *NonCombatEffect {
	effect := NewNonCombatEffect(
		"curse_"+curseType+"_"+time.Now().Format("20060102150405"),
		"Curse - "+curseType,
		"curse",
		MysticalCategory,
	).SetDuration(duration).SetIntensity(1.0).SetStackable(false, 1)

	switch curseType {
	case "weakness":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "strength",
			Multiplier: 0.8, // -20% strength
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "hp_max",
			Multiplier: 0.9, // -10% max HP
		})

	case "misfortune":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "luck",
			Multiplier: 0.5, // -50% luck
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "crit_chance",
			Multiplier: 0.7, // -30% crit chance
		})

	case "energy_drain":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "spiritual_energy",
			Multiplier: 0.6, // -40% spiritual energy
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "energy_efficiency",
			Multiplier: 0.5, // -50% energy efficiency
		})

	case "cultivation_block":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "cultivation_speed",
			Multiplier: 0.1, // -90% cultivation speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "breakthrough_success",
			Multiplier: 0.0, // 0% breakthrough success
		})
	}

	return effect
}

// CreateDivineCurseEffect creates a divine curse effect
func (ee *EffectExamples) CreateDivineCurseEffect(curseType string, duration int64) *NonCombatEffect {
	effect := NewNonCombatEffect(
		"divine_curse_"+curseType+"_"+time.Now().Format("20060102150405"),
		"Divine Curse - "+curseType,
		"divine_curse",
		MysticalCategory,
	).SetDuration(duration).SetIntensity(1.0).SetStackable(false, 1)

	switch curseType {
	case "karma_imbalance":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "karma",
			Multiplier: 0.3, // -70% karma
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "luck",
			Multiplier: 0.5, // -50% luck
		})

	case "fate_distortion":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "fate",
			Multiplier: 0.4, // -60% fate
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "reality_bend",
			Multiplier: 0.2, // -80% reality bend
		})
	}

	return effect
}

// === POSITIVE EFFECTS ===

// CreateCultivationBoostEffect creates a cultivation boost effect
func (ee *EffectExamples) CreateCultivationBoostEffect(boostType string, duration int64) *NonCombatEffect {
	effect := NewNonCombatEffect(
		"cultivation_boost_"+boostType+"_"+time.Now().Format("20060102150405"),
		"Cultivation Boost - "+boostType,
		CultivationBoostEffect,
		CultivationCategory,
	).SetDuration(duration).SetIntensity(1.0).SetStackable(true, 3)

	switch boostType {
	case "spiritual_insight":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "cultivation_speed",
			Multiplier: 1.5, // +50% cultivation speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "spiritual_energy",
			Multiplier: 1.2, // +20% spiritual energy
		})

	case "energy_refinement":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "energy_efficiency",
			Multiplier: 1.3, // +30% energy efficiency
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "breakthrough_success",
			Multiplier: 1.2, // +20% breakthrough success
		})
	}

	return effect
}

// CreateLearningBoostEffect creates a learning boost effect
func (ee *EffectExamples) CreateLearningBoostEffect(boostType string, duration int64) *NonCombatEffect {
	effect := NewNonCombatEffect(
		"learning_boost_"+boostType+"_"+time.Now().Format("20060102150405"),
		"Learning Boost - "+boostType,
		LearningBoostEffect,
		LearningCategory,
	).SetDuration(duration).SetIntensity(1.0).SetStackable(false, 1)

	switch boostType {
	case "memory_enhancement":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "learning_rate",
			Multiplier: 1.5, // +50% learning rate
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "memory",
			Multiplier: 1.3, // +30% memory
		})

	case "focus":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "learning_rate",
			Multiplier: 1.2, // +20% learning rate
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "concentration",
			Multiplier: 1.4, // +40% concentration
		})
	}

	return effect
}

// === ENVIRONMENTAL EFFECTS ===

// CreateWeatherEffect creates a weather-related effect
func (ee *EffectExamples) CreateWeatherEffect(weatherType string, duration int64) *NonCombatEffect {
	effect := NewNonCombatEffect(
		"weather_"+weatherType+"_"+time.Now().Format("20060102150405"),
		"Weather - "+weatherType,
		"weather_adaptation",
		EnvironmentalCategory,
	).SetDuration(duration).SetIntensity(1.0).SetStackable(false, 1)

	switch weatherType {
	case "heavy_rain":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "speed",
			Multiplier: 0.8, // -20% speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "fire_damage",
			Multiplier: 0.7, // -30% fire damage
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "water_damage",
			Multiplier: 1.3, // +30% water damage
		})

	case "blizzard":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "speed",
			Multiplier: 0.6, // -40% speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "ice_damage",
			Multiplier: 1.5, // +50% ice damage
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "fire_damage",
			Multiplier: 0.5, // -50% fire damage
		})
	}

	return effect
}

// === COMBAT-RELATED NON-COMBAT EFFECTS ===

// CreateCombatFatigueEffect creates a combat fatigue effect
func (ee *EffectExamples) CreateCombatFatigueEffect(severity string, duration int64) *NonCombatEffect {
	effect := NewNonCombatEffect(
		"combat_fatigue_"+severity+"_"+time.Now().Format("20060102150405"),
		"Combat Fatigue - "+severity,
		"combat_fatigue",
		EnvironmentalCategory,
	).SetDuration(duration).SetIntensity(1.0).SetStackable(true, 5)

	switch severity {
	case "light":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "speed",
			Multiplier: 0.95, // -5% speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "regen_hp",
			Multiplier: 0.9, // -10% HP regeneration
		})

	case "moderate":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "speed",
			Multiplier: 0.85, // -15% speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "regen_hp",
			Multiplier: 0.7, // -30% HP regeneration
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "cultivation_speed",
			Multiplier: 0.9, // -10% cultivation speed
		})

	case "severe":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "speed",
			Multiplier: 0.7, // -30% speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "regen_hp",
			Multiplier: 0.5, // -50% HP regeneration
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "cultivation_speed",
			Multiplier: 0.7, // -30% cultivation speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "energy_efficiency",
			Multiplier: 0.8, // -20% energy efficiency
		})
	}

	return effect
}

// CreateSpiritualExhaustionEffect creates a spiritual exhaustion effect
func (ee *EffectExamples) CreateSpiritualExhaustionEffect(severity string, duration int64) *NonCombatEffect {
	effect := NewNonCombatEffect(
		"spiritual_exhaustion_"+severity+"_"+time.Now().Format("20060102150405"),
		"Spiritual Exhaustion - "+severity,
		"spiritual_exhaustion",
		CultivationCategory,
	).SetDuration(duration).SetIntensity(1.0).SetStackable(true, 3)

	switch severity {
	case "light":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "spiritual_energy",
			Multiplier: 0.9, // -10% spiritual energy
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "energy_efficiency",
			Multiplier: 0.95, // -5% energy efficiency
		})

	case "moderate":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "spiritual_energy",
			Multiplier: 0.7, // -30% spiritual energy
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "energy_efficiency",
			Multiplier: 0.8, // -20% energy efficiency
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "cultivation_speed",
			Multiplier: 0.9, // -10% cultivation speed
		})

	case "severe":
		effect.AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "primary",
			Stat:       "spiritual_energy",
			Multiplier: 0.5, // -50% spiritual energy
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "energy_efficiency",
			Multiplier: 0.6, // -40% energy efficiency
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "cultivation_speed",
			Multiplier: 0.7, // -30% cultivation speed
		}).AddModifier(NonCombatEffectModifier{
			Type:       MultiplierModifier,
			Target:     "derived",
			Stat:       "breakthrough_success",
			Multiplier: 0.8, // -20% breakthrough success
		})
	}

	return effect
}
