package core

import (
	"context"
	"fmt"

	"actor-core-v2/constants"
	"actor-core-v2/models/core"
)

// StatResolver implements the StatResolver interface
type StatResolver struct {
	formulas map[string]Formula
	cache    map[string]float64
	version  int64
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

// initializeDefaultFormulas initializes default calculation formulas
func (sr *StatResolver) initializeDefaultFormulas() {
	// HPMax formula: Vitality * 10 + Constitution * 5
	sr.formulas["hp_max"] = &BasicFormula{
		Name:         "hp_max",
		Type:         "calculation",
		Dependencies: []string{"vitality", "constitution"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Vitality)*constants.HPMaxVitalityMultiplier + float64(primary.Constitution)*constants.HPMaxConstitutionMultiplier
		},
	}

	// Stamina formula: Endurance * 10 + Constitution * 3
	sr.formulas["stamina"] = &BasicFormula{
		Name:         "stamina",
		Type:         "calculation",
		Dependencies: []string{"endurance", "constitution"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(float64(primary.Endurance)*constants.StaminaEnduranceMultiplier + float64(primary.Constitution)*constants.StaminaConstitutionMultiplier)
		},
	}

	// Speed formula: Agility * 0.1
	sr.formulas["speed"] = &BasicFormula{
		Name:         "speed",
		Type:         "calculation",
		Dependencies: []string{"agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Agility) * constants.SpeedAgilityMultiplier
		},
	}

	// Haste formula: 1.0 + Agility * 0.01
	sr.formulas["haste"] = &BasicFormula{
		Name:         "haste",
		Type:         "calculation",
		Dependencies: []string{"agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.HasteBaseValue + float64(primary.Agility)*constants.HasteAgilityMultiplier
		},
	}

	// CritChance formula: 0.05 + Luck * 0.001
	sr.formulas["crit_chance"] = &BasicFormula{
		Name:         "crit_chance",
		Type:         "calculation",
		Dependencies: []string{"luck"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.CritChanceBaseValue + float64(primary.Luck)*constants.CritChanceLuckMultiplier
		},
	}

	// CritMulti formula: 1.5 + Luck * 0.01
	sr.formulas["crit_multi"] = &BasicFormula{
		Name:         "crit_multi",
		Type:         "calculation",
		Dependencies: []string{"luck"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.CritMultiBaseValue + float64(primary.Luck)*constants.CritMultiLuckMultiplier
		},
	}

	// MoveSpeed formula: Agility * 0.1
	sr.formulas["move_speed"] = &BasicFormula{
		Name:         "move_speed",
		Type:         "calculation",
		Dependencies: []string{"agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Agility) * constants.SpeedAgilityMultiplier
		},
	}

	// RegenHP formula: Vitality * 0.01
	sr.formulas["regen_hp"] = &BasicFormula{
		Name:         "regen_hp",
		Type:         "calculation",
		Dependencies: []string{"vitality"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Vitality) * constants.RegenHPVitalityMultiplier
		},
	}

	// Accuracy formula: 0.8 + Intelligence * 0.01
	sr.formulas["accuracy"] = &BasicFormula{
		Name:         "accuracy",
		Type:         "calculation",
		Dependencies: []string{"intelligence"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.AccuracyBaseValue + float64(primary.Intelligence)*constants.AccuracyIntelligenceMultiplier
		},
	}

	// Penetration formula: Strength * 0.01
	sr.formulas["penetration"] = &BasicFormula{
		Name:         "penetration",
		Type:         "calculation",
		Dependencies: []string{"strength"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Strength) * constants.PenetrationStrengthMultiplier
		},
	}

	// Lethality formula: (Strength + Agility) * 0.005
	sr.formulas["lethality"] = &BasicFormula{
		Name:         "lethality",
		Type:         "calculation",
		Dependencies: []string{"strength", "agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Strength+primary.Agility) * constants.LethalityStrengthAgilityMultiplier
		},
	}

	// Brutality formula: Strength * 0.01
	sr.formulas["brutality"] = &BasicFormula{
		Name:         "brutality",
		Type:         "calculation",
		Dependencies: []string{"strength"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Strength) * constants.PenetrationStrengthMultiplier
		},
	}

	// ArmorClass formula: 10.0 + Constitution * 0.5
	sr.formulas["armor_class"] = &BasicFormula{
		Name:         "armor_class",
		Type:         "calculation",
		Dependencies: []string{"constitution"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.ArmorClassBaseValue + float64(primary.Constitution)*constants.ArmorClassConstitutionMultiplier
		},
	}

	// Evasion formula: Agility * 0.01
	sr.formulas["evasion"] = &BasicFormula{
		Name:         "evasion",
		Type:         "calculation",
		Dependencies: []string{"agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Agility) * constants.EvasionAgilityMultiplier
		},
	}

	// BlockChance formula: Constitution * 0.005
	sr.formulas["block_chance"] = &BasicFormula{
		Name:         "block_chance",
		Type:         "calculation",
		Dependencies: []string{"constitution"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Constitution) * constants.BlockChanceConstitutionMultiplier
		},
	}

	// ParryChance formula: Agility * 0.005
	sr.formulas["parry_chance"] = &BasicFormula{
		Name:         "parry_chance",
		Type:         "calculation",
		Dependencies: []string{"agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Agility) * constants.ParryChanceAgilityMultiplier
		},
	}

	// DodgeChance formula: Agility * 0.01
	sr.formulas["dodge_chance"] = &BasicFormula{
		Name:         "dodge_chance",
		Type:         "calculation",
		Dependencies: []string{"agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Agility) * constants.EvasionAgilityMultiplier
		},
	}

	// EnergyEfficiency formula: 1.0 + Intelligence * 0.01
	sr.formulas["energy_efficiency"] = &BasicFormula{
		Name:         "energy_efficiency",
		Type:         "calculation",
		Dependencies: []string{"intelligence"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// EnergyCapacity formula: SpiritualEnergy + PhysicalEnergy + MentalEnergy
	sr.formulas["energy_capacity"] = &BasicFormula{
		Name:         "energy_capacity",
		Type:         "calculation",
		Dependencies: []string{"spiritual_energy", "physical_energy", "mental_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.SpiritualEnergy + primary.PhysicalEnergy + primary.MentalEnergy)
		},
	}

	// EnergyDrain formula: Willpower * 0.01
	sr.formulas["energy_drain"] = &BasicFormula{
		Name:         "energy_drain",
		Type:         "calculation",
		Dependencies: []string{"willpower"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Willpower) * constants.EnergyDrainWillpowerMultiplier
		},
	}

	// ResourceRegen formula: Vitality * 0.01
	sr.formulas["resource_regen"] = &BasicFormula{
		Name:         "resource_regen",
		Type:         "calculation",
		Dependencies: []string{"vitality"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Vitality) * constants.RegenHPVitalityMultiplier
		},
	}

	// LearningRate formula: 1.0 + Intelligence * 0.01
	sr.formulas["learning_rate"] = &BasicFormula{
		Name:         "learning_rate",
		Type:         "calculation",
		Dependencies: []string{"intelligence"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// Adaptation formula: 1.0 + Wisdom * 0.01
	sr.formulas["adaptation"] = &BasicFormula{
		Name:         "adaptation",
		Type:         "calculation",
		Dependencies: []string{"wisdom"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.AdaptationBaseValue + float64(primary.Wisdom)*constants.AdaptationWisdomMultiplier
		},
	}

	// Memory formula: 1.0 + Intelligence * 0.01
	sr.formulas["memory"] = &BasicFormula{
		Name:         "memory",
		Type:         "calculation",
		Dependencies: []string{"intelligence"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// Experience formula: 1.0 + Wisdom * 0.01
	sr.formulas["experience"] = &BasicFormula{
		Name:         "experience",
		Type:         "calculation",
		Dependencies: []string{"wisdom"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.AdaptationBaseValue + float64(primary.Wisdom)*constants.AdaptationWisdomMultiplier
		},
	}

	// Leadership formula: 1.0 + Charisma * 0.01
	sr.formulas["leadership"] = &BasicFormula{
		Name:         "leadership",
		Type:         "calculation",
		Dependencies: []string{"charisma"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LeadershipBaseValue + float64(primary.Charisma)*constants.LeadershipCharismaMultiplier
		},
	}

	// Diplomacy formula: 1.0 + (Charisma + Intelligence) * 0.005
	sr.formulas["diplomacy"] = &BasicFormula{
		Name:         "diplomacy",
		Type:         "calculation",
		Dependencies: []string{"charisma", "intelligence"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.DiplomacyBaseValue + float64(primary.Charisma+primary.Intelligence)*constants.DiplomacyCharismaIntelligenceMultiplier
		},
	}

	// Intimidation formula: 1.0 + (Strength + Charisma) * 0.005
	sr.formulas["intimidation"] = &BasicFormula{
		Name:         "intimidation",
		Type:         "calculation",
		Dependencies: []string{"strength", "charisma"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.IntimidationBaseValue + float64(primary.Strength+primary.Charisma)*constants.IntimidationStrengthCharismaMultiplier
		},
	}

	// Empathy formula: 1.0 + (Wisdom + Charisma) * 0.005
	sr.formulas["empathy"] = &BasicFormula{
		Name:         "empathy",
		Type:         "calculation",
		Dependencies: []string{"wisdom", "charisma"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EmpathyBaseValue + float64(primary.Wisdom+primary.Charisma)*constants.EmpathyWisdomCharismaMultiplier
		},
	}

	// Deception formula: 1.0 + (Intelligence + Charisma) * 0.005
	sr.formulas["deception"] = &BasicFormula{
		Name:         "deception",
		Type:         "calculation",
		Dependencies: []string{"intelligence", "charisma"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.DeceptionBaseValue + float64(primary.Intelligence+primary.Charisma)*constants.DeceptionIntelligenceCharismaMultiplier
		},
	}

	// Performance formula: 1.0 + Charisma * 0.01
	sr.formulas["performance"] = &BasicFormula{
		Name:         "performance",
		Type:         "calculation",
		Dependencies: []string{"charisma"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LeadershipBaseValue + float64(primary.Charisma)*constants.LeadershipCharismaMultiplier
		},
	}

	// ManaEfficiency formula: 1.0 + Intelligence * 0.01
	sr.formulas["mana_efficiency"] = &BasicFormula{
		Name:         "mana_efficiency",
		Type:         "calculation",
		Dependencies: []string{"intelligence"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// SpellPower formula: 1.0 + SpiritualEnergy * 0.01
	sr.formulas["spell_power"] = &BasicFormula{
		Name:         "spell_power",
		Type:         "calculation",
		Dependencies: []string{"spiritual_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.SpellPowerBaseValue + float64(primary.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
		},
	}

	// MysticResonance formula: 1.0 + (SpiritualEnergy + MentalEnergy) * 0.005
	sr.formulas["mystic_resonance"] = &BasicFormula{
		Name:         "mystic_resonance",
		Type:         "calculation",
		Dependencies: []string{"spiritual_energy", "mental_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.MysticResonanceBaseValue + float64(primary.SpiritualEnergy+primary.MentalEnergy)*constants.MysticResonanceSpiritualMentalMultiplier
		},
	}

	// RealityBend formula: 1.0 + MentalEnergy * 0.01
	sr.formulas["reality_bend"] = &BasicFormula{
		Name:         "reality_bend",
		Type:         "calculation",
		Dependencies: []string{"mental_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.RealityBendBaseValue + float64(primary.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier
		},
	}

	// TimeSense formula: 1.0 + MentalEnergy * 0.01
	sr.formulas["time_sense"] = &BasicFormula{
		Name:         "time_sense",
		Type:         "calculation",
		Dependencies: []string{"mental_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.RealityBendBaseValue + float64(primary.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier
		},
	}

	// SpaceSense formula: 1.0 + MentalEnergy * 0.01
	sr.formulas["space_sense"] = &BasicFormula{
		Name:         "space_sense",
		Type:         "calculation",
		Dependencies: []string{"mental_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.RealityBendBaseValue + float64(primary.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier
		},
	}

	// JumpHeight formula: 1.0 + (Strength + Agility) * 0.005
	sr.formulas["jump_height"] = &BasicFormula{
		Name:         "jump_height",
		Type:         "calculation",
		Dependencies: []string{"strength", "agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.JumpHeightBaseValue + float64(primary.Strength+primary.Agility)*constants.JumpHeightStrengthAgilityMultiplier
		},
	}

	// ClimbSpeed formula: 1.0 + (Strength + Agility) * 0.005
	sr.formulas["climb_speed"] = &BasicFormula{
		Name:         "climb_speed",
		Type:         "calculation",
		Dependencies: []string{"strength", "agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.JumpHeightBaseValue + float64(primary.Strength+primary.Agility)*constants.JumpHeightStrengthAgilityMultiplier
		},
	}

	// SwimSpeed formula: 1.0 + (Strength + Agility) * 0.005
	sr.formulas["swim_speed"] = &BasicFormula{
		Name:         "swim_speed",
		Type:         "calculation",
		Dependencies: []string{"strength", "agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.JumpHeightBaseValue + float64(primary.Strength+primary.Agility)*constants.JumpHeightStrengthAgilityMultiplier
		},
	}

	// FlightSpeed formula: 1.0 + SpiritualEnergy * 0.01
	sr.formulas["flight_speed"] = &BasicFormula{
		Name:         "flight_speed",
		Type:         "calculation",
		Dependencies: []string{"spiritual_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.SpellPowerBaseValue + float64(primary.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
		},
	}

	// TeleportRange formula: 1.0 + MentalEnergy * 0.01
	sr.formulas["teleport_range"] = &BasicFormula{
		Name:         "teleport_range",
		Type:         "calculation",
		Dependencies: []string{"mental_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.RealityBendBaseValue + float64(primary.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier
		},
	}

	// Stealth formula: 1.0 + Agility * 0.01
	sr.formulas["stealth"] = &BasicFormula{
		Name:         "stealth",
		Type:         "calculation",
		Dependencies: []string{"agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.HasteBaseValue + float64(primary.Agility)*constants.HasteAgilityMultiplier
		},
	}

	// AuraRadius formula: 1.0 + SpiritualEnergy * 0.01
	sr.formulas["aura_radius"] = &BasicFormula{
		Name:         "aura_radius",
		Type:         "calculation",
		Dependencies: []string{"spiritual_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.SpellPowerBaseValue + float64(primary.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
		},
	}

	// AuraStrength formula: 1.0 + SpiritualEnergy * 0.01
	sr.formulas["aura_strength"] = &BasicFormula{
		Name:         "aura_strength",
		Type:         "calculation",
		Dependencies: []string{"spiritual_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.SpellPowerBaseValue + float64(primary.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
		},
	}

	// Presence formula: 1.0 + (Charisma + SpiritualEnergy) * 0.005
	sr.formulas["presence"] = &BasicFormula{
		Name:         "presence",
		Type:         "calculation",
		Dependencies: []string{"charisma", "spiritual_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.PresenceBaseValue + float64(primary.Charisma+primary.SpiritualEnergy)*constants.PresenceCharismaSpiritualMultiplier
		},
	}

	// Awe formula: 1.0 + (Charisma + SpiritualEnergy) * 0.005
	sr.formulas["awe"] = &BasicFormula{
		Name:         "awe",
		Type:         "calculation",
		Dependencies: []string{"charisma", "spiritual_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.PresenceBaseValue + float64(primary.Charisma+primary.SpiritualEnergy)*constants.PresenceCharismaSpiritualMultiplier
		},
	}

	// WeaponMastery formula: 1.0 + (Strength + Agility) * 0.005
	sr.formulas["weapon_mastery"] = &BasicFormula{
		Name:         "weapon_mastery",
		Type:         "calculation",
		Dependencies: []string{"strength", "agility"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.JumpHeightBaseValue + float64(primary.Strength+primary.Agility)*constants.JumpHeightStrengthAgilityMultiplier
		},
	}

	// SkillLevel formula: 1.0 + Intelligence * 0.01
	sr.formulas["skill_level"] = &BasicFormula{
		Name:         "skill_level",
		Type:         "calculation",
		Dependencies: []string{"intelligence"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// LifeSteal formula: 0.0 (default)
	sr.formulas["life_steal"] = &BasicFormula{
		Name:         "life_steal",
		Type:         "calculation",
		Dependencies: []string{},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LifeStealDefaultValue
		},
	}

	// CastSpeed formula: 1.0 + Intelligence * 0.01
	sr.formulas["cast_speed"] = &BasicFormula{
		Name:         "cast_speed",
		Type:         "calculation",
		Dependencies: []string{"intelligence"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// WeightCapacity formula: Strength * 10.0
	sr.formulas["weight_capacity"] = &BasicFormula{
		Name:         "weight_capacity",
		Type:         "calculation",
		Dependencies: []string{"strength"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Strength) * constants.WeightCapacityStrengthMultiplier
		},
	}

	// Persuasion formula: 1.0 + Charisma * 0.01
	sr.formulas["persuasion"] = &BasicFormula{
		Name:         "persuasion",
		Type:         "calculation",
		Dependencies: []string{"charisma"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LeadershipBaseValue + float64(primary.Charisma)*constants.LeadershipCharismaMultiplier
		},
	}

	// MerchantPriceModifier formula: 1.0 + Charisma * 0.01
	sr.formulas["merchant_price_modifier"] = &BasicFormula{
		Name:         "merchant_price_modifier",
		Type:         "calculation",
		Dependencies: []string{"charisma"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LeadershipBaseValue + float64(primary.Charisma)*constants.LeadershipCharismaMultiplier
		},
	}

	// FactionReputationGain formula: 1.0 + Charisma * 0.01
	sr.formulas["faction_reputation_gain"] = &BasicFormula{
		Name:         "faction_reputation_gain",
		Type:         "calculation",
		Dependencies: []string{"charisma"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.LeadershipBaseValue + float64(primary.Charisma)*constants.LeadershipCharismaMultiplier
		},
	}

	// CultivationSpeed formula: 1.0 + (SpiritualEnergy + PhysicalEnergy + MentalEnergy) * 0.001
	sr.formulas["cultivation_speed"] = &BasicFormula{
		Name:         "cultivation_speed",
		Type:         "calculation",
		Dependencies: []string{"spiritual_energy", "physical_energy", "mental_energy"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.CultivationSpeedBaseValue + float64(primary.SpiritualEnergy+primary.PhysicalEnergy+primary.MentalEnergy)*constants.CultivationSpeedAllEnergyMultiplier
		},
	}

	// EnergyEfficiencyAmp formula: 1.0 + Intelligence * 0.01
	sr.formulas["energy_efficiency_amp"] = &BasicFormula{
		Name:         "energy_efficiency_amp",
		Type:         "calculation",
		Dependencies: []string{"intelligence"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.EnergyEfficiencyBaseValue + float64(primary.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
		},
	}

	// BreakthroughSuccess formula: 1.0 + (Willpower + Luck) * 0.005
	sr.formulas["breakthrough_success"] = &BasicFormula{
		Name:         "breakthrough_success",
		Type:         "calculation",
		Dependencies: []string{"willpower", "luck"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.BreakthroughSuccessBaseValue + float64(primary.Willpower+primary.Luck)*constants.BreakthroughSuccessWillpowerLuckMultiplier
		},
	}

	// SkillLearning formula: 1.0 + (Intelligence + Wisdom) * 0.005
	sr.formulas["skill_learning"] = &BasicFormula{
		Name:         "skill_learning",
		Type:         "calculation",
		Dependencies: []string{"intelligence", "wisdom"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.SkillLearningBaseValue + float64(primary.Intelligence+primary.Wisdom)*constants.SkillLearningIntelligenceWisdomMultiplier
		},
	}

	// CombatEffectiveness formula: 1.0 + (Strength + Agility + Intelligence) * 0.003
	sr.formulas["combat_effectiveness"] = &BasicFormula{
		Name:         "combat_effectiveness",
		Type:         "calculation",
		Dependencies: []string{"strength", "agility", "intelligence"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.CombatEffectivenessBaseValue + float64(primary.Strength+primary.Agility+primary.Intelligence)*constants.CombatEffectivenessStrengthAgilityIntelligenceMultiplier
		},
	}

	// ResourceGathering formula: 1.0 + (Luck + Wisdom) * 0.005
	sr.formulas["resource_gathering"] = &BasicFormula{
		Name:         "resource_gathering",
		Type:         "calculation",
		Dependencies: []string{"luck", "wisdom"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return constants.ResourceGatheringBaseValue + float64(primary.Luck+primary.Wisdom)*constants.ResourceGatheringLuckWisdomMultiplier
		},
	}
}

// ResolveStats resolves all stats from primary stats
func (sr *StatResolver) ResolveStats(primaryStats *core.PrimaryCore) (*core.DerivedStats, error) {
	derivedStats := core.NewDerivedStats()

	// Calculate all derived stats using formulas
	for statName, formula := range sr.formulas {
		value := formula.Calculate(primaryStats)

		// Set the calculated value
		if err := derivedStats.SetStat(statName, value); err != nil {
			return nil, fmt.Errorf("failed to set stat %s: %w", statName, err)
		}

		// Cache the result
		sr.cache[statName] = value
	}

	sr.version++
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
	formula, exists := sr.formulas[statName]
	if !exists {
		return constants.LifeStealDefaultValue, fmt.Errorf("formula for stat %s not found", statName)
	}

	value := formula.Calculate(primaryStats)

	// Cache the result
	sr.cache[statName] = value

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
		case "crit_chance":
			if value > constants.MaxCritChanceValue {
				return fmt.Errorf("crit_chance cannot exceed %f: %f", constants.MaxCritChanceValue, value)
			}
		case "crit_multi":
			if value < constants.MinCritMultiValue {
				return fmt.Errorf("crit_multi cannot be less than %f: %f", constants.MinCritMultiValue, value)
			}
		case "haste":
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

	sr.formulas[formula.GetName()] = formula
	sr.version++

	return nil
}

// RemoveFormula removes a formula
func (sr *StatResolver) RemoveFormula(statName string) error {
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
	formula, exists := sr.formulas[statName]
	if !exists {
		return nil, fmt.Errorf("formula for stat %s not found", statName)
	}

	return formula, nil
}

// GetAllFormulas returns all formulas
func (sr *StatResolver) GetAllFormulas() map[string]Formula {
	result := make(map[string]Formula)
	for name, formula := range sr.formulas {
		result[name] = formula
	}
	return result
}

// ClearCache clears the calculation cache
func (sr *StatResolver) ClearCache() {
	sr.cache = make(map[string]float64)
}

// GetCacheSize returns the cache size
func (sr *StatResolver) GetCacheSize() int {
	return len(sr.cache)
}

// GetVersion returns the current version
func (sr *StatResolver) GetVersion() int64 {
	return sr.version
}

// GetStatsCount returns the number of stats
func (sr *StatResolver) GetStatsCount() int {
	return len(sr.formulas)
}
