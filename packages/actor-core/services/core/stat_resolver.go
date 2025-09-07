package core

import (
	"context"
	"fmt"
	"sync"

	"actor-core/constants"
	"actor-core/models/core"
)

// StatResolver implements the StatResolver interface
type StatResolver struct {
	formulas map[string]Formula
	cache    map[string]float64
	mu       sync.RWMutex
	version  int64 // FormulaRegistryVersion
}

// Formula represents a stat calculation formula
type Formula interface {
	Calculate(primary *core.PrimaryCore) float64
	GetDependencies() []string
	GetName() string
	GetType() string
}

// BasicFormula represents a basic calculation formula
type BasicFormula struct {
	Name         string
	Type         string
	Dependencies []string
	Calculator   func(primary *core.PrimaryCore) float64
}

// Calculate implements Formula interface
func (f *BasicFormula) Calculate(primary *core.PrimaryCore) float64 {
	return f.Calculator(primary)
}

// GetDependencies implements Formula interface
func (f *BasicFormula) GetDependencies() []string {
	return f.Dependencies
}

// GetName implements Formula interface
func (f *BasicFormula) GetName() string {
	return f.Name
}

// GetType implements Formula interface
func (f *BasicFormula) GetType() string {
	return f.Type
}

// NewStatResolver creates a new StatResolver
func NewStatResolver() *StatResolver {
	sr := &StatResolver{
		formulas: make(map[string]Formula),
		cache:    make(map[string]float64),
		version:  1,
	}

	// Initialize default formulas
	sr.initializeDefaultFormulas()

	return sr
}

// generateCacheKey creates a versioned cache key
func (sr *StatResolver) generateCacheKey(statName string, primary *core.PrimaryCore) string {
	// For now, use simple versioned key. In production, could include hash of primary stats
	return fmt.Sprintf("stat:%s:v%d", statName, sr.version)
}

// initializeDefaultFormulas initializes default calculation formulas
func (sr *StatResolver) initializeDefaultFormulas() {
	// HPMax formula: Vitality * 10 + Constitution * 5
	sr.formulas[constants.Stat_HP_MAX] = &BasicFormula{
		Name:         constants.Stat_HP_MAX,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_VITALITY, constants.Stat_CONSTITUTION},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Vitality)*constants.HPMaxVitalityMultiplier + float64(primary.Constitution)*constants.HPMaxConstitutionMultiplier
		},
	}

	// Stamina formula: Endurance * 10 + Constitution * 3
	sr.formulas[constants.Stat_STAMINA] = &BasicFormula{
		Name:         constants.Stat_STAMINA,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_ENDURANCE, constants.Stat_CONSTITUTION},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(float64(primary.Endurance)*constants.StaminaEnduranceMultiplier + float64(primary.Constitution)*constants.StaminaConstitutionMultiplier)
		},
	}

	// Speed formula: Agility * 0.1
	sr.formulas[constants.Stat_SPEED] = &BasicFormula{
		Name:         constants.Stat_SPEED,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Agility) * constants.SpeedAgilityMultiplier
		},
	}

	// Haste formula: 1.0 + Agility * 0.01
	sr.formulas[constants.Stat_HASTE] = &BasicFormula{
		Name:         constants.Stat_HASTE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.HasteBaseValue + float64(primary.Agility)*constants.HasteAgilityMultiplier
		},
	}

	// CritChance formula: 0.05 + Luck * 0.001
	sr.formulas[constants.Stat_CRIT_CHANCE] = &BasicFormula{
		Name:         constants.Stat_CRIT_CHANCE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_LUCK},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.CritChanceBaseValue + float64(primary.Luck)*constants.CritChanceLuckMultiplier
		},
	}

	// CritMulti formula: 1.5 + Luck * 0.01
	sr.formulas[constants.Stat_CRIT_MULTI] = &BasicFormula{
		Name:         constants.Stat_CRIT_MULTI,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_LUCK},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.CritMultiBaseValue + float64(primary.Luck)*constants.CritMultiLuckMultiplier
		},
	}

	// MoveSpeed formula: Agility * 0.1
	sr.formulas[constants.Stat_MOVE_SPEED] = &BasicFormula{
		Name:         constants.Stat_MOVE_SPEED,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Agility) * constants.SpeedAgilityMultiplier
		},
	}

	// RegenHP formula: Vitality * 0.01
	sr.formulas[constants.Stat_REGEN_HP] = &BasicFormula{
		Name:         constants.Stat_REGEN_HP,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_VITALITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Vitality) * constants.RegenHPVitalityMultiplier
		},
	}

	// Accuracy formula: 0.8 + Intelligence * 0.01
	sr.formulas[constants.Stat_ACCURACY] = &BasicFormula{
		Name:         constants.Stat_ACCURACY,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_INTELLIGENCE},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.AccuracyBaseValue + float64(primary.Intelligence)*constants.AccuracyIntelligenceMultiplier
		},
	}

	// Penetration formula: Strength * 0.01
	sr.formulas[constants.Stat_PENETRATION] = &BasicFormula{
		Name:         constants.Stat_PENETRATION,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_STRENGTH},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Strength) * constants.PenetrationStrengthMultiplier
		},
	}

	// Lethality formula: (Strength + Agility) * 0.005
	sr.formulas[constants.Stat_LETHALITY] = &BasicFormula{
		Name:         constants.Stat_LETHALITY,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_STRENGTH, constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Strength+primary.Agility) * constants.LethalityStrengthAgilityMultiplier
		},
	}

	// Brutality formula: Strength * 0.01
	sr.formulas[constants.Stat_BRUTALITY] = &BasicFormula{
		Name:         constants.Stat_BRUTALITY,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_STRENGTH},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Strength) * constants.PenetrationStrengthMultiplier
		},
	}

	// ArmorClass formula: 10.0 + Constitution * 0.5
	sr.formulas[constants.Stat_ARMOR_CLASS] = &BasicFormula{
		Name:         constants.Stat_ARMOR_CLASS,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_CONSTITUTION},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.ArmorClassBaseValue + float64(primary.Constitution)*constants.ArmorClassConstitutionMultiplier
		},
	}

	// Evasion formula: Agility * 0.01
	sr.formulas[constants.Stat_EVASION] = &BasicFormula{
		Name:         constants.Stat_EVASION,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Agility) * constants.EvasionAgilityMultiplier
		},
	}

	// BlockChance formula: Constitution * 0.005
	sr.formulas[constants.Stat_BLOCK_CHANCE] = &BasicFormula{
		Name:         constants.Stat_BLOCK_CHANCE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_CONSTITUTION},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Constitution) * constants.BlockChanceConstitutionMultiplier
		},
	}

	// ParryChance formula: Agility * 0.005
	sr.formulas[constants.Stat_PARRY_CHANCE] = &BasicFormula{
		Name:         constants.Stat_PARRY_CHANCE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Agility) * constants.ParryChanceAgilityMultiplier
		},
	}

	// DodgeChance formula: Agility * 0.01
	sr.formulas[constants.Stat_DODGE_CHANCE] = &BasicFormula{
		Name:         constants.Stat_DODGE_CHANCE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Agility) * constants.EvasionAgilityMultiplier
		},
	}

	// EnergyEfficiency formula: 1.0 + Intelligence * 0.01
	sr.formulas[constants.Stat_ENERGY_EFFICIENCY] = &BasicFormula{
		Name:         constants.Stat_ENERGY_EFFICIENCY,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_INTELLIGENCE},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// EnergyCapacity formula: SpiritualEnergy + PhysicalEnergy + MentalEnergy
	sr.formulas[constants.Stat_ENERGY_CAPACITY] = &BasicFormula{
		Name:         constants.Stat_ENERGY_CAPACITY,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_SPIRITUAL_ENERGY, constants.Stat_PHYSICAL_ENERGY, constants.Stat_MENTAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.SpiritualEnergy + primary.PhysicalEnergy + primary.MentalEnergy)
		},
	}

	// EnergyDrain formula: Willpower * 0.01
	sr.formulas[constants.Stat_ENERGY_DRAIN] = &BasicFormula{
		Name:         constants.Stat_ENERGY_DRAIN,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_WILLPOWER},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Willpower) * constants.EnergyDrainWillpowerMultiplier
		},
	}

	// ResourceRegen formula: Vitality * 0.01
	sr.formulas[constants.Stat_RESOURCE_REGEN] = &BasicFormula{
		Name:         constants.Stat_RESOURCE_REGEN,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_VITALITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Vitality) * constants.RegenHPVitalityMultiplier
		},
	}

	// LearningRate formula: 1.0 + Intelligence * 0.01
	sr.formulas[constants.Stat_LEARNING_RATE] = &BasicFormula{
		Name:         constants.Stat_LEARNING_RATE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_INTELLIGENCE},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// Adaptation formula: 1.0 + Wisdom * 0.01
	sr.formulas[constants.Stat_ADAPTATION] = &BasicFormula{
		Name:         constants.Stat_ADAPTATION,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_WISDOM},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.AdaptationBaseValue + float64(primary.Wisdom)*constants.AdaptationWisdomMultiplier
		},
	}

	// Memory formula: 1.0 + Intelligence * 0.01
	sr.formulas[constants.Stat_MEMORY] = &BasicFormula{
		Name:         constants.Stat_MEMORY,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_INTELLIGENCE},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// Experience formula: 1.0 + Wisdom * 0.01
	sr.formulas[constants.Stat_EXPERIENCE] = &BasicFormula{
		Name:         constants.Stat_EXPERIENCE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_WISDOM},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.AdaptationBaseValue + float64(primary.Wisdom)*constants.AdaptationWisdomMultiplier
		},
	}

	// Leadership formula: 1.0 + Charisma * 0.01
	sr.formulas[constants.Stat_LEADERSHIP] = &BasicFormula{
		Name:         constants.Stat_LEADERSHIP,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_CHARISMA},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LeadershipBaseValue + float64(primary.Charisma)*constants.LeadershipCharismaMultiplier
		},
	}

	// Diplomacy formula: 1.0 + (Charisma + Intelligence) * 0.005
	sr.formulas[constants.Stat_DIPLOMACY] = &BasicFormula{
		Name:         constants.Stat_DIPLOMACY,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_CHARISMA, constants.Stat_INTELLIGENCE},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.DiplomacyBaseValue + float64(primary.Charisma+primary.Intelligence)*constants.DiplomacyCharismaIntelligenceMultiplier
		},
	}

	// Intimidation formula: 1.0 + (Strength + Charisma) * 0.005
	sr.formulas[constants.Stat_INTIMIDATION] = &BasicFormula{
		Name:         constants.Stat_INTIMIDATION,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_STRENGTH, constants.Stat_CHARISMA},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.IntimidationBaseValue + float64(primary.Strength+primary.Charisma)*constants.IntimidationStrengthCharismaMultiplier
		},
	}

	// Empathy formula: 1.0 + (Wisdom + Charisma) * 0.005
	sr.formulas[constants.Stat_EMPATHY] = &BasicFormula{
		Name:         constants.Stat_EMPATHY,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_WISDOM, constants.Stat_CHARISMA},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EmpathyBaseValue + float64(primary.Wisdom+primary.Charisma)*constants.EmpathyWisdomCharismaMultiplier
		},
	}

	// Deception formula: 1.0 + (Intelligence + Charisma) * 0.005
	sr.formulas[constants.Stat_DECEPTION] = &BasicFormula{
		Name:         constants.Stat_DECEPTION,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_INTELLIGENCE, constants.Stat_CHARISMA},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.DeceptionBaseValue + float64(primary.Intelligence+primary.Charisma)*constants.DeceptionIntelligenceCharismaMultiplier
		},
	}

	// Performance formula: 1.0 + Charisma * 0.01
	sr.formulas[constants.Stat_PERFORMANCE] = &BasicFormula{
		Name:         constants.Stat_PERFORMANCE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_CHARISMA},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LeadershipBaseValue + float64(primary.Charisma)*constants.LeadershipCharismaMultiplier
		},
	}

	// ManaEfficiency formula: 1.0 + Intelligence * 0.01
	sr.formulas[constants.Stat_MANA_EFFICIENCY] = &BasicFormula{
		Name:         constants.Stat_MANA_EFFICIENCY,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_INTELLIGENCE},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// SpellPower formula: 1.0 + SpiritualEnergy * 0.01
	sr.formulas[constants.Stat_SPELL_POWER] = &BasicFormula{
		Name:         constants.Stat_SPELL_POWER,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_SPIRITUAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.SpellPowerBaseValue + float64(primary.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
		},
	}

	// MysticResonance formula: 1.0 + (SpiritualEnergy + MentalEnergy) * 0.005
	sr.formulas[constants.Stat_MYSTIC_RESONANCE] = &BasicFormula{
		Name:         constants.Stat_MYSTIC_RESONANCE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_SPIRITUAL_ENERGY, constants.Stat_MENTAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.MysticResonanceBaseValue + float64(primary.SpiritualEnergy+primary.MentalEnergy)*constants.MysticResonanceSpiritualMentalMultiplier
		},
	}

	// RealityBend formula: 1.0 + MentalEnergy * 0.01
	sr.formulas[constants.Stat_REALITY_BEND] = &BasicFormula{
		Name:         constants.Stat_REALITY_BEND,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_MENTAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.RealityBendBaseValue + float64(primary.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier
		},
	}

	// TimeSense formula: 1.0 + MentalEnergy * 0.01
	sr.formulas[constants.Stat_TIME_SENSE] = &BasicFormula{
		Name:         constants.Stat_TIME_SENSE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_MENTAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.RealityBendBaseValue + float64(primary.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier
		},
	}

	// SpaceSense formula: 1.0 + MentalEnergy * 0.01
	sr.formulas[constants.Stat_SPACE_SENSE] = &BasicFormula{
		Name:         constants.Stat_SPACE_SENSE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_MENTAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.RealityBendBaseValue + float64(primary.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier
		},
	}

	// JumpHeight formula: 1.0 + (Strength + Agility) * 0.005
	sr.formulas[constants.Stat_JUMP_HEIGHT] = &BasicFormula{
		Name:         constants.Stat_JUMP_HEIGHT,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_STRENGTH, constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.JumpHeightBaseValue + float64(primary.Strength+primary.Agility)*constants.JumpHeightStrengthAgilityMultiplier
		},
	}

	// ClimbSpeed formula: 1.0 + (Strength + Agility) * 0.005
	sr.formulas[constants.Stat_CLIMB_SPEED] = &BasicFormula{
		Name:         constants.Stat_CLIMB_SPEED,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_STRENGTH, constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.JumpHeightBaseValue + float64(primary.Strength+primary.Agility)*constants.JumpHeightStrengthAgilityMultiplier
		},
	}

	// SwimSpeed formula: 1.0 + (Strength + Agility) * 0.005
	sr.formulas[constants.Stat_SWIM_SPEED] = &BasicFormula{
		Name:         constants.Stat_SWIM_SPEED,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_STRENGTH, constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.JumpHeightBaseValue + float64(primary.Strength+primary.Agility)*constants.JumpHeightStrengthAgilityMultiplier
		},
	}

	// FlightSpeed formula: 1.0 + SpiritualEnergy * 0.01
	sr.formulas[constants.Stat_FLIGHT_SPEED] = &BasicFormula{
		Name:         constants.Stat_FLIGHT_SPEED,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_SPIRITUAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.SpellPowerBaseValue + float64(primary.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
		},
	}

	// TeleportRange formula: 1.0 + MentalEnergy * 0.01
	sr.formulas[constants.Stat_TELEPORT_RANGE] = &BasicFormula{
		Name:         constants.Stat_TELEPORT_RANGE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_MENTAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.RealityBendBaseValue + float64(primary.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier
		},
	}

	// Stealth formula: 1.0 + Agility * 0.01
	sr.formulas[constants.Stat_STEALTH] = &BasicFormula{
		Name:         constants.Stat_STEALTH,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.HasteBaseValue + float64(primary.Agility)*constants.HasteAgilityMultiplier
		},
	}

	// AuraRadius formula: 1.0 + SpiritualEnergy * 0.01
	sr.formulas[constants.Stat_AURA_RADIUS] = &BasicFormula{
		Name:         constants.Stat_AURA_RADIUS,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_SPIRITUAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.SpellPowerBaseValue + float64(primary.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
		},
	}

	// AuraStrength formula: 1.0 + SpiritualEnergy * 0.01
	sr.formulas[constants.Stat_AURA_STRENGTH] = &BasicFormula{
		Name:         constants.Stat_AURA_STRENGTH,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_SPIRITUAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.SpellPowerBaseValue + float64(primary.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
		},
	}

	// Presence formula: 1.0 + (Charisma + SpiritualEnergy) * 0.005
	sr.formulas[constants.Stat_PRESENCE] = &BasicFormula{
		Name:         constants.Stat_PRESENCE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_CHARISMA, constants.Stat_SPIRITUAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.PresenceBaseValue + float64(primary.Charisma+primary.SpiritualEnergy)*constants.PresenceCharismaSpiritualMultiplier
		},
	}

	// Awe formula: 1.0 + (Charisma + SpiritualEnergy) * 0.005
	sr.formulas[constants.Stat_AWE] = &BasicFormula{
		Name:         constants.Stat_AWE,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_CHARISMA, constants.Stat_SPIRITUAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.PresenceBaseValue + float64(primary.Charisma+primary.SpiritualEnergy)*constants.PresenceCharismaSpiritualMultiplier
		},
	}

	// WeaponMastery formula: 1.0 + (Strength + Agility) * 0.005
	sr.formulas[constants.Stat_WEAPON_MASTERY] = &BasicFormula{
		Name:         constants.Stat_WEAPON_MASTERY,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_STRENGTH, constants.Stat_AGILITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.JumpHeightBaseValue + float64(primary.Strength+primary.Agility)*constants.JumpHeightStrengthAgilityMultiplier
		},
	}

	// SkillLevel formula: 1.0 + Intelligence * 0.01
	sr.formulas[constants.Stat_SKILL_LEVEL] = &BasicFormula{
		Name:         constants.Stat_SKILL_LEVEL,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_INTELLIGENCE},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// LifeSteal formula: 0.0 (default)
	sr.formulas[constants.Stat_LIFE_STEAL] = &BasicFormula{
		Name:         constants.Stat_LIFE_STEAL,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LifeStealDefaultValue
		},
	}

	// CastSpeed formula: 1.0 + Intelligence * 0.01
	sr.formulas[constants.Stat_CAST_SPEED] = &BasicFormula{
		Name:         constants.Stat_CAST_SPEED,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_INTELLIGENCE},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// WeightCapacity formula: Strength * 10.0
	sr.formulas[constants.Stat_WEIGHT_CAPACITY] = &BasicFormula{
		Name:         constants.Stat_WEIGHT_CAPACITY,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_STRENGTH},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Strength) * constants.WeightCapacityStrengthMultiplier
		},
	}

	// Persuasion formula: 1.0 + Charisma * 0.01
	sr.formulas[constants.Stat_PERSUASION] = &BasicFormula{
		Name:         constants.Stat_PERSUASION,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_CHARISMA},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LeadershipBaseValue + float64(primary.Charisma)*constants.LeadershipCharismaMultiplier
		},
	}

	// MerchantPriceModifier formula: 1.0 + Charisma * 0.01
	sr.formulas[constants.Stat_MERCHANT_PRICE_MODIFIER] = &BasicFormula{
		Name:         constants.Stat_MERCHANT_PRICE_MODIFIER,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_CHARISMA},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LeadershipBaseValue + float64(primary.Charisma)*constants.LeadershipCharismaMultiplier
		},
	}

	// FactionReputationGain formula: 1.0 + Charisma * 0.01
	sr.formulas[constants.Stat_FACTION_REPUTATION_GAIN] = &BasicFormula{
		Name:         constants.Stat_FACTION_REPUTATION_GAIN,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_CHARISMA},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LeadershipBaseValue + float64(primary.Charisma)*constants.LeadershipCharismaMultiplier
		},
	}

	// CultivationSpeed formula: 1.0 + (SpiritualEnergy + PhysicalEnergy + MentalEnergy) * 0.001
	sr.formulas[constants.Stat_CULTIVATION_SPEED] = &BasicFormula{
		Name:         constants.Stat_CULTIVATION_SPEED,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_SPIRITUAL_ENERGY, constants.Stat_PHYSICAL_ENERGY, constants.Stat_MENTAL_ENERGY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.CultivationSpeedBaseValue + float64(primary.SpiritualEnergy+primary.PhysicalEnergy+primary.MentalEnergy)*constants.CultivationSpeedAllEnergyMultiplier
		},
	}

	// EnergyEfficiencyAmp formula: 1.0 + Intelligence * 0.01
	sr.formulas[constants.Stat_ENERGY_EFFICIENCY_AMP] = &BasicFormula{
		Name:         constants.Stat_ENERGY_EFFICIENCY_AMP,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_INTELLIGENCE},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// BreakthroughSuccess formula: 1.0 + (Willpower + Luck) * 0.005
	sr.formulas[constants.Stat_BREAKTHROUGH_SUCCESS] = &BasicFormula{
		Name:         constants.Stat_BREAKTHROUGH_SUCCESS,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_WILLPOWER, constants.Stat_LUCK},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.BreakthroughSuccessBaseValue + float64(primary.Willpower+primary.Luck)*constants.BreakthroughSuccessWillpowerLuckMultiplier
		},
	}

	// SkillLearning formula: 1.0 + (Intelligence + Wisdom) * 0.005
	sr.formulas[constants.Stat_SKILL_LEARNING] = &BasicFormula{
		Name:         constants.Stat_SKILL_LEARNING,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_INTELLIGENCE, constants.Stat_WISDOM},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.SkillLearningBaseValue + float64(primary.Intelligence+primary.Wisdom)*constants.SkillLearningIntelligenceWisdomMultiplier
		},
	}

	// CombatEffectiveness formula: 1.0 + (Strength + Agility + Intelligence) * 0.003
	sr.formulas[constants.Stat_COMBAT_EFFECTIVENESS] = &BasicFormula{
		Name:         constants.Stat_COMBAT_EFFECTIVENESS,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_STRENGTH, constants.Stat_AGILITY, constants.Stat_INTELLIGENCE},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.CombatEffectivenessBaseValue + float64(primary.Strength+primary.Agility+primary.Intelligence)*constants.CombatEffectivenessStrengthAgilityIntelligenceMultiplier
		},
	}

	// ResourceGathering formula: 1.0 + (Luck + Wisdom) * 0.005
	sr.formulas[constants.Stat_RESOURCE_GATHERING] = &BasicFormula{
		Name:         constants.Stat_RESOURCE_GATHERING,
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_LUCK, constants.Stat_WISDOM},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.ResourceGatheringBaseValue + float64(primary.Luck+primary.Wisdom)*constants.ResourceGatheringLuckWisdomMultiplier
		},
	}
}

// ResolveStats resolves all stats from primary stats
func (sr *StatResolver) ResolveStats(primaryStats *core.PrimaryCore) (*core.DerivedStats, error) {
	derivedStats := core.NewDerivedStats()

	// Get formulas with read lock
	sr.mu.RLock()
	formulas := make(map[string]Formula)
	for name, formula := range sr.formulas {
		formulas[name] = formula
	}
	sr.mu.RUnlock()

	// Calculate all derived stats using formulas
	for statName, formula := range formulas {
		value := formula.Calculate(primaryStats)

		// Set the calculated value
		if err := derivedStats.SetStat(statName, value); err != nil {
			return nil, fmt.Errorf("failed to set stat %s: %w", statName, err)
		}

		// Cache the result with write lock
		sr.mu.Lock()
		sr.cache[statName] = value
		sr.mu.Unlock()
	}

	// Increment version with write lock
	sr.mu.Lock()
	sr.version++
	sr.mu.Unlock()

	return derivedStats, nil
}

// ResolveStatsWithContext resolves all stats from primary stats with context
func (sr *StatResolver) ResolveStatsWithContext(ctx context.Context, primaryStats *core.PrimaryCore) (*core.DerivedStats, error) {
	// Check if context is cancelled
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return sr.ResolveStats(primaryStats)
}

// ResolveStat resolves a specific stat
func (sr *StatResolver) ResolveStat(statName string, primaryStats *core.PrimaryCore) (float64, error) {
	// Generate versioned cache key
	cacheKey := sr.generateCacheKey(statName, primaryStats)

	// Check cache first
	if value, exists := sr.getCachedValue(cacheKey); exists {
		return value, nil
	}

	// Get formula with read lock
	sr.mu.RLock()
	formula, exists := sr.formulas[statName]
	sr.mu.RUnlock()

	if !exists {
		return constants.LifeStealDefaultValue, fmt.Errorf("formula for stat %s not found", statName)
	}

	// Calculate value
	value := formula.Calculate(primaryStats)

	// Cache the result with write lock
	sr.setCachedValue(cacheKey, value)

	return value, nil
}

// ResolveStatWithContext resolves a specific stat with context
func (sr *StatResolver) ResolveStatWithContext(ctx context.Context, statName string, primaryStats *core.PrimaryCore) (float64, error) {
	// Check if context is cancelled
	select {
	case <-ctx.Done():
		return constants.LifeStealDefaultValue, ctx.Err()
	default:
	}

	return sr.ResolveStat(statName, primaryStats)
}

// ResolveDerivedStats resolves all derived stats
func (sr *StatResolver) ResolveDerivedStats(primaryStats *core.PrimaryCore) (map[string]float64, error) {
	derivedStats, err := sr.ResolveStats(primaryStats)
	if err != nil {
		return nil, err
	}

	return derivedStats.GetAllStats(), nil
}

// ResolveDerivedStatsWithContext resolves all derived stats with context
func (sr *StatResolver) ResolveDerivedStatsWithContext(ctx context.Context, primaryStats *core.PrimaryCore) (map[string]float64, error) {
	// Check if context is cancelled
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return sr.ResolveDerivedStats(primaryStats)
}

// CheckDependencies checks if all dependencies are available
func (sr *StatResolver) CheckDependencies(statName string) ([]string, error) {
	formula, exists := sr.formulas[statName]
	if !exists {
		return nil, fmt.Errorf("formula for stat %s not found", statName)
	}

	return formula.GetDependencies(), nil
}

// CheckDependenciesWithContext checks if all dependencies are available with context
func (sr *StatResolver) CheckDependenciesWithContext(ctx context.Context, statName string) ([]string, error) {
	// Check if context is cancelled
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return sr.CheckDependencies(statName)
}

// GetCalculationOrder returns the order in which stats should be calculated
func (sr *StatResolver) GetCalculationOrder() []string {
	// Simple topological sort based on dependencies
	visited := make(map[string]bool)
	order := make([]string, 0, len(sr.formulas))

	var visit func(string)
	visit = func(statName string) {
		if visited[statName] {
			return
		}
		visited[statName] = true

		// Visit dependencies first
		if formula, exists := sr.formulas[statName]; exists {
			for _, dep := range formula.GetDependencies() {
				visit(dep)
			}
		}

		// Add to order
		order = append(order, statName)
	}

	// Visit all stats
	for statName := range sr.formulas {
		visit(statName)
	}

	return order
}

// GetCalculationOrderWithContext returns the order in which stats should be calculated with context
func (sr *StatResolver) GetCalculationOrderWithContext(ctx context.Context) []string {
	// Check if context is cancelled
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	return sr.GetCalculationOrder()
}

// ValidateStats validates all stats
func (sr *StatResolver) ValidateStats(stats map[string]float64) error {
	for statName, value := range stats {
		if value < 0 {
			return fmt.Errorf("stat %s has negative value: %f", statName, value)
		}

		// Add more validation rules as needed
		switch statName {
		case constants.Stat_CRIT_CHANCE:
			if value > constants.MaxCritChanceValue {
				return fmt.Errorf("crit_chance cannot exceed %f: %f", constants.MaxCritChanceValue, value)
			}
		case constants.Stat_CRIT_MULTI:
			if value < constants.MinCritMultiValue {
				return fmt.Errorf("crit_multi cannot be less than %f: %f", constants.MinCritMultiValue, value)
			}
		case constants.Stat_HASTE:
			if value < constants.MinHasteValue {
				return fmt.Errorf("haste cannot be less than %f: %f", constants.MinHasteValue, value)
			}
		}
	}

	return nil
}

// ValidateStatsWithContext validates all stats with context
func (sr *StatResolver) ValidateStatsWithContext(ctx context.Context, stats map[string]float64) error {
	// Check if context is cancelled
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	return sr.ValidateStats(stats)
}

// AddFormula adds a new formula
func (sr *StatResolver) AddFormula(formula Formula) error {
	if formula == nil {
		return fmt.Errorf("formula cannot be nil")
	}

	if formula.GetName() == "" {
		return fmt.Errorf("formula name cannot be empty")
	}

	sr.mu.Lock()
	defer sr.mu.Unlock()

	sr.formulas[formula.GetName()] = formula
	sr.version++

	return nil
}

// RemoveFormula removes a formula
func (sr *StatResolver) RemoveFormula(statName string) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	if _, exists := sr.formulas[statName]; !exists {
		return fmt.Errorf("formula for stat %s not found", statName)
	}

	delete(sr.formulas, statName)
	delete(sr.cache, statName)
	sr.version++

	return nil
}

// GetFormula gets a formula
func (sr *StatResolver) GetFormula(statName string) (Formula, error) {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	formula, exists := sr.formulas[statName]
	if !exists {
		return nil, fmt.Errorf("formula for stat %s not found", statName)
	}

	return formula, nil
}

// GetAllFormulas returns all formulas
func (sr *StatResolver) GetAllFormulas() map[string]Formula {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	result := make(map[string]Formula)
	for name, formula := range sr.formulas {
		result[name] = formula
	}
	return result
}

// ClearCache clears the calculation cache
func (sr *StatResolver) ClearCache() {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.cache = make(map[string]float64)
}

// GetCacheSize returns the cache size
func (sr *StatResolver) GetCacheSize() int {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	return len(sr.cache)
}

// getCachedValue gets a value from cache with read lock
func (sr *StatResolver) getCachedValue(key string) (float64, bool) {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	value, exists := sr.cache[key]
	return value, exists
}

// setCachedValue sets a value in cache with write lock
func (sr *StatResolver) setCachedValue(key string, value float64) {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.cache[key] = value
}

// GetVersion returns the current version
func (sr *StatResolver) GetVersion() int64 {
	return sr.version
}

// GetStatsCount returns the number of stats
func (sr *StatResolver) GetStatsCount() int {
	return len(sr.formulas)
}
