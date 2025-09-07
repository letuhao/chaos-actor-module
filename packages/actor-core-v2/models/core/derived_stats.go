package core

import (
	"time"

	"actor-core/constants"
)

// DerivedStats represents the derived stats of an actor
type DerivedStats struct {
	// Core Derived Stats
	HPMax      float64 `json:"hp_max"`
	Stamina    float64 `json:"stamina"`
	Speed      float64 `json:"speed"`
	Haste      float64 `json:"haste"`
	CritChance float64 `json:"crit_chance"`
	CritMulti  float64 `json:"crit_multi"`
	MoveSpeed  float64 `json:"move_speed"`
	RegenHP    float64 `json:"regen_hp"`

	// Combat Stats
	Accuracy    float64 `json:"accuracy"`
	Penetration float64 `json:"penetration"`
	Lethality   float64 `json:"lethality"`
	Brutality   float64 `json:"brutality"`
	ArmorClass  float64 `json:"armor_class"`
	Evasion     float64 `json:"evasion"`
	BlockChance float64 `json:"block_chance"`
	ParryChance float64 `json:"parry_chance"`
	DodgeChance float64 `json:"dodge_chance"`

	// Energy Stats
	EnergyEfficiency float64 `json:"energy_efficiency"`
	EnergyCapacity   float64 `json:"energy_capacity"`
	EnergyDrain      float64 `json:"energy_drain"`
	ResourceRegen    float64 `json:"resource_regen"`
	ResourceDecay    float64 `json:"resource_decay"`

	// Learning Stats
	LearningRate float64 `json:"learning_rate"`
	Adaptation   float64 `json:"adaptation"`
	Memory       float64 `json:"memory"`
	Experience   float64 `json:"experience"`

	// Social Stats
	Leadership   float64 `json:"leadership"`
	Diplomacy    float64 `json:"diplomacy"`
	Intimidation float64 `json:"intimidation"`
	Empathy      float64 `json:"empathy"`
	Deception    float64 `json:"deception"`
	Performance  float64 `json:"performance"`

	// Mystical Stats
	ManaEfficiency  float64 `json:"mana_efficiency"`
	SpellPower      float64 `json:"spell_power"`
	MysticResonance float64 `json:"mystic_resonance"`
	RealityBend     float64 `json:"reality_bend"`
	TimeSense       float64 `json:"time_sense"`
	SpaceSense      float64 `json:"space_sense"`

	// Movement Stats
	JumpHeight    float64 `json:"jump_height"`
	ClimbSpeed    float64 `json:"climb_speed"`
	SwimSpeed     float64 `json:"swim_speed"`
	FlightSpeed   float64 `json:"flight_speed"`
	TeleportRange float64 `json:"teleport_range"`
	Stealth       float64 `json:"stealth"`

	// Aura Stats
	AuraRadius   float64 `json:"aura_radius"`
	AuraStrength float64 `json:"aura_strength"`
	Presence     float64 `json:"presence"`
	Awe          float64 `json:"awe"`

	// Proficiency Stats
	WeaponMastery         float64 `json:"weapon_mastery"`
	SkillLevel            float64 `json:"skill_level"`
	LifeSteal             float64 `json:"life_steal"`
	CastSpeed             float64 `json:"cast_speed"`
	WeightCapacity        float64 `json:"weight_capacity"`
	Persuasion            float64 `json:"persuasion"`
	MerchantPriceModifier float64 `json:"merchant_price_modifier"`
	FactionReputationGain float64 `json:"faction_reputation_gain"`

	// Talent Amplifiers
	CultivationSpeed    float64 `json:"cultivation_speed"`
	EnergyEfficiencyAmp float64 `json:"energy_efficiency_amp"`
	BreakthroughSuccess float64 `json:"breakthrough_success"`
	SkillLearning       float64 `json:"skill_learning"`
	CombatEffectiveness float64 `json:"combat_effectiveness"`
	ResourceGathering   float64 `json:"resource_gathering"`

	// Metadata
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	Version   int64 `json:"version"`
}

// NewDerivedStats creates a new DerivedStats with default values
func NewDerivedStats() *DerivedStats {
	now := time.Now().Unix()
	return &DerivedStats{
		// Core Derived Stats - calculated from primary stats
		HPMax:      100.0,
		Stamina:    100.0,
		Speed:      1.0,
		Haste:      1.0,
		CritChance: 0.05,
		CritMulti:  1.5,
		MoveSpeed:  1.0,
		RegenHP:    0.1,

		// Combat Stats
		Accuracy:    0.8,
		Penetration: 0.0,
		Lethality:   0.0,
		Brutality:   0.0,
		ArmorClass:  10.0,
		Evasion:     0.1,
		BlockChance: 0.0,
		ParryChance: 0.0,
		DodgeChance: 0.0,

		// Energy Stats
		EnergyEfficiency: 1.0,
		EnergyCapacity:   100.0,
		EnergyDrain:      0.0,
		ResourceRegen:    0.1,
		ResourceDecay:    0.0,

		// Learning Stats
		LearningRate: 1.0,
		Adaptation:   1.0,
		Memory:       1.0,
		Experience:   1.0,

		// Social Stats
		Leadership:   1.0,
		Diplomacy:    1.0,
		Intimidation: 1.0,
		Empathy:      1.0,
		Deception:    1.0,
		Performance:  1.0,

		// Mystical Stats
		ManaEfficiency:  1.0,
		SpellPower:      1.0,
		MysticResonance: 1.0,
		RealityBend:     1.0,
		TimeSense:       1.0,
		SpaceSense:      1.0,

		// Movement Stats
		JumpHeight:    1.0,
		ClimbSpeed:    1.0,
		SwimSpeed:     1.0,
		FlightSpeed:   1.0,
		TeleportRange: 1.0,
		Stealth:       1.0,

		// Aura Stats
		AuraRadius:   1.0,
		AuraStrength: 1.0,
		Presence:     1.0,
		Awe:          1.0,

		// Proficiency Stats
		WeaponMastery:         1.0,
		SkillLevel:            1.0,
		LifeSteal:             0.0,
		CastSpeed:             1.0,
		WeightCapacity:        100.0,
		Persuasion:            1.0,
		MerchantPriceModifier: 1.0,
		FactionReputationGain: 1.0,

		// Talent Amplifiers
		CultivationSpeed:    1.0,
		EnergyEfficiencyAmp: 1.0,
		BreakthroughSuccess: 1.0,
		SkillLearning:       1.0,
		CombatEffectiveness: 1.0,
		ResourceGathering:   1.0,

		// Metadata
		CreatedAt: now,
		UpdatedAt: now,
		Version:   1,
	}
}

// CalculateFromPrimary calculates derived stats from primary stats
func (ds *DerivedStats) CalculateFromPrimary(pc *PrimaryCore) {
	now := time.Now().Unix()

	// Core Derived Stats calculations
	ds.HPMax = float64(pc.Vitality)*constants.HPMaxVitalityMultiplier + float64(pc.Constitution)*constants.HPMaxConstitutionMultiplier
	ds.Stamina = float64(pc.Endurance)*constants.StaminaEnduranceMultiplier + float64(pc.Constitution)*constants.StaminaConstitutionMultiplier
	ds.Speed = float64(pc.Agility) * constants.SpeedAgilityMultiplier
	ds.Haste = constants.HasteBaseValue + float64(pc.Agility)*constants.HasteAgilityMultiplier
	ds.CritChance = constants.CritChanceBaseValue + float64(pc.Luck)*constants.CritChanceLuckMultiplier
	ds.CritMulti = constants.CritMultiBaseValue + float64(pc.Luck)*constants.CritMultiLuckMultiplier
	ds.MoveSpeed = float64(pc.Agility) * constants.MoveSpeedAgilityMultiplier
	ds.RegenHP = float64(pc.Vitality) * constants.RegenHPVitalityMultiplier

	// Combat Stats calculations
	ds.Accuracy = constants.AccuracyBaseValue + float64(pc.Intelligence)*constants.AccuracyIntelligenceMultiplier
	ds.Penetration = float64(pc.Strength) * constants.PenetrationStrengthMultiplier
	ds.Lethality = float64(pc.Strength+pc.Agility) * constants.LethalityStrengthAgilityMultiplier
	ds.Brutality = float64(pc.Strength) * constants.PenetrationStrengthMultiplier
	ds.ArmorClass = constants.ArmorClassBaseValue + float64(pc.Constitution)*constants.ArmorClassConstitutionMultiplier
	ds.Evasion = float64(pc.Agility) * constants.EvasionAgilityMultiplier
	ds.BlockChance = float64(pc.Constitution) * constants.BlockChanceConstitutionMultiplier
	ds.ParryChance = float64(pc.Agility) * constants.ParryChanceAgilityMultiplier
	ds.DodgeChance = float64(pc.Agility) * constants.EvasionAgilityMultiplier

	// Energy Stats calculations
	ds.EnergyEfficiency = constants.EnergyEfficiencyBaseValue + float64(pc.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
	ds.EnergyCapacity = float64(pc.SpiritualEnergy + pc.PhysicalEnergy + pc.MentalEnergy)
	ds.EnergyDrain = float64(pc.Willpower) * constants.EnergyDrainWillpowerMultiplier
	ds.ResourceRegen = float64(pc.Vitality) * constants.ResourceRegenVitalityMultiplier
	ds.ResourceDecay = constants.LifeStealDefaultValue

	// Learning Stats calculations
	ds.LearningRate = constants.EnergyEfficiencyBaseValue + float64(pc.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
	ds.Adaptation = constants.AdaptationBaseValue + float64(pc.Wisdom)*constants.AdaptationWisdomMultiplier
	ds.Memory = constants.EnergyEfficiencyBaseValue + float64(pc.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
	ds.Experience = constants.AdaptationBaseValue + float64(pc.Wisdom)*constants.AdaptationWisdomMultiplier

	// Social Stats calculations
	ds.Leadership = constants.LeadershipBaseValue + float64(pc.Charisma)*constants.LeadershipCharismaMultiplier
	ds.Diplomacy = constants.DiplomacyBaseValue + float64(pc.Charisma+pc.Intelligence)*constants.DiplomacyCharismaIntelligenceMultiplier
	ds.Intimidation = constants.IntimidationBaseValue + float64(pc.Strength+pc.Charisma)*constants.IntimidationStrengthCharismaMultiplier
	ds.Empathy = constants.EmpathyBaseValue + float64(pc.Wisdom+pc.Charisma)*constants.EmpathyWisdomCharismaMultiplier
	ds.Deception = constants.DeceptionBaseValue + float64(pc.Intelligence+pc.Charisma)*constants.DeceptionIntelligenceCharismaMultiplier
	ds.Performance = constants.LeadershipBaseValue + float64(pc.Charisma)*constants.LeadershipCharismaMultiplier

	// Mystical Stats calculations
	ds.ManaEfficiency = constants.EnergyEfficiencyBaseValue + float64(pc.Intelligence)*constants.EnergyEfficiencyIntelligenceMultiplier
	ds.SpellPower = constants.SpellPowerBaseValue + float64(pc.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
	ds.MysticResonance = constants.MysticResonanceBaseValue + float64(pc.SpiritualEnergy+pc.MentalEnergy)*constants.MysticResonanceSpiritualMentalMultiplier
	ds.RealityBend = constants.RealityBendBaseValue + float64(pc.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier
	ds.TimeSense = constants.RealityBendBaseValue + float64(pc.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier
	ds.SpaceSense = constants.RealityBendBaseValue + float64(pc.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier

	// Movement Stats calculations
	ds.JumpHeight = constants.JumpHeightBaseValue + float64(pc.Strength+pc.Agility)*constants.JumpHeightStrengthAgilityMultiplier
	ds.ClimbSpeed = constants.JumpHeightBaseValue + float64(pc.Strength+pc.Agility)*constants.JumpHeightStrengthAgilityMultiplier
	ds.SwimSpeed = constants.JumpHeightBaseValue + float64(pc.Strength+pc.Agility)*constants.JumpHeightStrengthAgilityMultiplier
	ds.FlightSpeed = constants.SpellPowerBaseValue + float64(pc.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
	ds.TeleportRange = constants.RealityBendBaseValue + float64(pc.MentalEnergy)*constants.RealityBendMentalEnergyMultiplier
	ds.Stealth = constants.StealthBaseValue + float64(pc.Agility)*constants.StealthAgilityMultiplier

	// Aura Stats calculations
	ds.AuraRadius = constants.SpellPowerBaseValue + float64(pc.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
	ds.AuraStrength = constants.SpellPowerBaseValue + float64(pc.SpiritualEnergy)*constants.SpellPowerSpiritualEnergyMultiplier
	ds.Presence = constants.PresenceBaseValue + float64(pc.Charisma+pc.SpiritualEnergy)*constants.PresenceCharismaSpiritualMultiplier
	ds.Awe = constants.PresenceBaseValue + float64(pc.Charisma+pc.SpiritualEnergy)*constants.PresenceCharismaSpiritualMultiplier

	// Proficiency Stats calculations
	ds.WeaponMastery = constants.WeaponMasteryBaseValue + float64(pc.Strength+pc.Agility)*constants.WeaponMasteryStrengthAgilityMultiplier
	ds.SkillLevel = constants.SkillLevelBaseValue + float64(pc.Intelligence)*constants.SkillLevelIntelligenceMultiplier
	ds.LifeSteal = constants.LifeStealDefaultValue
	ds.CastSpeed = constants.CastSpeedBaseValue + float64(pc.Intelligence)*constants.CastSpeedIntelligenceMultiplier
	ds.WeightCapacity = float64(pc.Strength) * constants.WeightCapacityStrengthMultiplier
	ds.Persuasion = constants.PersuasionBaseValue + float64(pc.Charisma)*constants.PersuasionCharismaMultiplier
	ds.MerchantPriceModifier = constants.MerchantPriceModifierBaseValue + float64(pc.Charisma)*constants.MerchantPriceModifierCharismaMultiplier
	ds.FactionReputationGain = constants.FactionReputationGainBaseValue + float64(pc.Charisma)*constants.FactionReputationGainCharismaMultiplier

	// Talent Amplifiers calculations
	ds.CultivationSpeed = constants.CultivationSpeedBaseValue + float64(pc.SpiritualEnergy+pc.PhysicalEnergy+pc.MentalEnergy)*constants.CultivationSpeedAllEnergyMultiplier
	ds.EnergyEfficiencyAmp = constants.EnergyEfficiencyAmpBaseValue + float64(pc.Intelligence)*constants.EnergyEfficiencyAmpIntelligenceMultiplier
	ds.BreakthroughSuccess = constants.BreakthroughSuccessBaseValue + float64(pc.Willpower+pc.Luck)*constants.BreakthroughSuccessWillpowerLuckMultiplier
	ds.SkillLearning = constants.SkillLearningBaseValue + float64(pc.Intelligence+pc.Wisdom)*constants.SkillLearningIntelligenceWisdomMultiplier
	ds.CombatEffectiveness = constants.CombatEffectivenessBaseValue + float64(pc.Strength+pc.Agility+pc.Intelligence)*constants.CombatEffectivenessStrengthAgilityIntelligenceMultiplier
	ds.ResourceGathering = constants.ResourceGatheringBaseValue + float64(pc.Luck+pc.Wisdom)*constants.ResourceGatheringLuckWisdomMultiplier

	ds.UpdatedAt = now
	ds.Version++
}

// GetStat gets a stat value by name
func (ds *DerivedStats) GetStat(statName string) (float64, error) {
	switch statName {
	case "hp_max":
		return ds.HPMax, nil
	case "stamina":
		return ds.Stamina, nil
	case "speed":
		return ds.Speed, nil
	case "haste":
		return ds.Haste, nil
	case "crit_chance":
		return ds.CritChance, nil
	case "crit_multi":
		return ds.CritMulti, nil
	case "move_speed":
		return ds.MoveSpeed, nil
	case "regen_hp":
		return ds.RegenHP, nil
	case "accuracy":
		return ds.Accuracy, nil
	case "penetration":
		return ds.Penetration, nil
	case "lethality":
		return ds.Lethality, nil
	case "brutality":
		return ds.Brutality, nil
	case "armor_class":
		return ds.ArmorClass, nil
	case "evasion":
		return ds.Evasion, nil
	case "block_chance":
		return ds.BlockChance, nil
	case "parry_chance":
		return ds.ParryChance, nil
	case "dodge_chance":
		return ds.DodgeChance, nil
	case "energy_efficiency":
		return ds.EnergyEfficiency, nil
	case "energy_capacity":
		return ds.EnergyCapacity, nil
	case "energy_drain":
		return ds.EnergyDrain, nil
	case "resource_regen":
		return ds.ResourceRegen, nil
	case "resource_decay":
		return ds.ResourceDecay, nil
	case "learning_rate":
		return ds.LearningRate, nil
	case "adaptation":
		return ds.Adaptation, nil
	case "memory":
		return ds.Memory, nil
	case "experience":
		return ds.Experience, nil
	case "leadership":
		return ds.Leadership, nil
	case "diplomacy":
		return ds.Diplomacy, nil
	case "intimidation":
		return ds.Intimidation, nil
	case "empathy":
		return ds.Empathy, nil
	case "deception":
		return ds.Deception, nil
	case "performance":
		return ds.Performance, nil
	case "mana_efficiency":
		return ds.ManaEfficiency, nil
	case "spell_power":
		return ds.SpellPower, nil
	case "mystic_resonance":
		return ds.MysticResonance, nil
	case "reality_bend":
		return ds.RealityBend, nil
	case "time_sense":
		return ds.TimeSense, nil
	case "space_sense":
		return ds.SpaceSense, nil
	case "jump_height":
		return ds.JumpHeight, nil
	case "climb_speed":
		return ds.ClimbSpeed, nil
	case "swim_speed":
		return ds.SwimSpeed, nil
	case "flight_speed":
		return ds.FlightSpeed, nil
	case "teleport_range":
		return ds.TeleportRange, nil
	case "stealth":
		return ds.Stealth, nil
	case "aura_radius":
		return ds.AuraRadius, nil
	case "aura_strength":
		return ds.AuraStrength, nil
	case "presence":
		return ds.Presence, nil
	case "awe":
		return ds.Awe, nil
	case "weapon_mastery":
		return ds.WeaponMastery, nil
	case "skill_level":
		return ds.SkillLevel, nil
	case "life_steal":
		return ds.LifeSteal, nil
	case "cast_speed":
		return ds.CastSpeed, nil
	case "weight_capacity":
		return ds.WeightCapacity, nil
	case "persuasion":
		return ds.Persuasion, nil
	case "merchant_price_modifier":
		return ds.MerchantPriceModifier, nil
	case "faction_reputation_gain":
		return ds.FactionReputationGain, nil
	case "cultivation_speed":
		return ds.CultivationSpeed, nil
	case "energy_efficiency_amp":
		return ds.EnergyEfficiencyAmp, nil
	case "breakthrough_success":
		return ds.BreakthroughSuccess, nil
	case "skill_learning":
		return ds.SkillLearning, nil
	case "combat_effectiveness":
		return ds.CombatEffectiveness, nil
	case "resource_gathering":
		return ds.ResourceGathering, nil
	default:
		return constants.LifeStealDefaultValue, ErrStatNotFound
	}
}

// SetStat sets a stat value by name
func (ds *DerivedStats) SetStat(statName string, value float64) error {
	switch statName {
	case "hp_max":
		ds.HPMax = value
	case "stamina":
		ds.Stamina = value
	case "speed":
		ds.Speed = value
	case "haste":
		ds.Haste = value
	case "crit_chance":
		ds.CritChance = value
	case "crit_multi":
		ds.CritMulti = value
	case "move_speed":
		ds.MoveSpeed = value
	case "regen_hp":
		ds.RegenHP = value
	case "accuracy":
		ds.Accuracy = value
	case "penetration":
		ds.Penetration = value
	case "lethality":
		ds.Lethality = value
	case "brutality":
		ds.Brutality = value
	case "armor_class":
		ds.ArmorClass = value
	case "evasion":
		ds.Evasion = value
	case "block_chance":
		ds.BlockChance = value
	case "parry_chance":
		ds.ParryChance = value
	case "dodge_chance":
		ds.DodgeChance = value
	case "energy_efficiency":
		ds.EnergyEfficiency = value
	case "energy_capacity":
		ds.EnergyCapacity = value
	case "energy_drain":
		ds.EnergyDrain = value
	case "resource_regen":
		ds.ResourceRegen = value
	case "resource_decay":
		ds.ResourceDecay = value
	case "learning_rate":
		ds.LearningRate = value
	case "adaptation":
		ds.Adaptation = value
	case "memory":
		ds.Memory = value
	case "experience":
		ds.Experience = value
	case "leadership":
		ds.Leadership = value
	case "diplomacy":
		ds.Diplomacy = value
	case "intimidation":
		ds.Intimidation = value
	case "empathy":
		ds.Empathy = value
	case "deception":
		ds.Deception = value
	case "performance":
		ds.Performance = value
	case "mana_efficiency":
		ds.ManaEfficiency = value
	case "spell_power":
		ds.SpellPower = value
	case "mystic_resonance":
		ds.MysticResonance = value
	case "reality_bend":
		ds.RealityBend = value
	case "time_sense":
		ds.TimeSense = value
	case "space_sense":
		ds.SpaceSense = value
	case "jump_height":
		ds.JumpHeight = value
	case "climb_speed":
		ds.ClimbSpeed = value
	case "swim_speed":
		ds.SwimSpeed = value
	case "flight_speed":
		ds.FlightSpeed = value
	case "teleport_range":
		ds.TeleportRange = value
	case "stealth":
		ds.Stealth = value
	case "aura_radius":
		ds.AuraRadius = value
	case "aura_strength":
		ds.AuraStrength = value
	case "presence":
		ds.Presence = value
	case "awe":
		ds.Awe = value
	case "weapon_mastery":
		ds.WeaponMastery = value
	case "skill_level":
		ds.SkillLevel = value
	case "life_steal":
		ds.LifeSteal = value
	case "cast_speed":
		ds.CastSpeed = value
	case "weight_capacity":
		ds.WeightCapacity = value
	case "persuasion":
		ds.Persuasion = value
	case "merchant_price_modifier":
		ds.MerchantPriceModifier = value
	case "faction_reputation_gain":
		ds.FactionReputationGain = value
	case "cultivation_speed":
		ds.CultivationSpeed = value
	case "energy_efficiency_amp":
		ds.EnergyEfficiencyAmp = value
	case "breakthrough_success":
		ds.BreakthroughSuccess = value
	case "skill_learning":
		ds.SkillLearning = value
	case "combat_effectiveness":
		ds.CombatEffectiveness = value
	case "resource_gathering":
		ds.ResourceGathering = value
	default:
		return ErrStatNotFound
	}

	ds.UpdatedAt = time.Now().Unix()
	ds.Version++
	return nil
}

// GetAllStats returns all derived stats as a map
func (ds *DerivedStats) GetAllStats() map[string]float64 {
	return map[string]float64{
		"hp_max":                  ds.HPMax,
		"stamina":                 ds.Stamina,
		"speed":                   ds.Speed,
		"haste":                   ds.Haste,
		"crit_chance":             ds.CritChance,
		"crit_multi":              ds.CritMulti,
		"move_speed":              ds.MoveSpeed,
		"regen_hp":                ds.RegenHP,
		"accuracy":                ds.Accuracy,
		"penetration":             ds.Penetration,
		"lethality":               ds.Lethality,
		"brutality":               ds.Brutality,
		"armor_class":             ds.ArmorClass,
		"evasion":                 ds.Evasion,
		"block_chance":            ds.BlockChance,
		"parry_chance":            ds.ParryChance,
		"dodge_chance":            ds.DodgeChance,
		"energy_efficiency":       ds.EnergyEfficiency,
		"energy_capacity":         ds.EnergyCapacity,
		"energy_drain":            ds.EnergyDrain,
		"resource_regen":          ds.ResourceRegen,
		"resource_decay":          ds.ResourceDecay,
		"learning_rate":           ds.LearningRate,
		"adaptation":              ds.Adaptation,
		"memory":                  ds.Memory,
		"experience":              ds.Experience,
		"leadership":              ds.Leadership,
		"diplomacy":               ds.Diplomacy,
		"intimidation":            ds.Intimidation,
		"empathy":                 ds.Empathy,
		"deception":               ds.Deception,
		"performance":             ds.Performance,
		"mana_efficiency":         ds.ManaEfficiency,
		"spell_power":             ds.SpellPower,
		"mystic_resonance":        ds.MysticResonance,
		"reality_bend":            ds.RealityBend,
		"time_sense":              ds.TimeSense,
		"space_sense":             ds.SpaceSense,
		"jump_height":             ds.JumpHeight,
		"climb_speed":             ds.ClimbSpeed,
		"swim_speed":              ds.SwimSpeed,
		"flight_speed":            ds.FlightSpeed,
		"teleport_range":          ds.TeleportRange,
		"stealth":                 ds.Stealth,
		"aura_radius":             ds.AuraRadius,
		"aura_strength":           ds.AuraStrength,
		"presence":                ds.Presence,
		"awe":                     ds.Awe,
		"weapon_mastery":          ds.WeaponMastery,
		"skill_level":             ds.SkillLevel,
		"life_steal":              ds.LifeSteal,
		"cast_speed":              ds.CastSpeed,
		"weight_capacity":         ds.WeightCapacity,
		"persuasion":              ds.Persuasion,
		"merchant_price_modifier": ds.MerchantPriceModifier,
		"faction_reputation_gain": ds.FactionReputationGain,
		"cultivation_speed":       ds.CultivationSpeed,
		"energy_efficiency_amp":   ds.EnergyEfficiencyAmp,
		"breakthrough_success":    ds.BreakthroughSuccess,
		"skill_learning":          ds.SkillLearning,
		"combat_effectiveness":    ds.CombatEffectiveness,
		"resource_gathering":      ds.ResourceGathering,
	}
}

// Clone creates a deep copy of the DerivedStats
func (ds *DerivedStats) Clone() *DerivedStats {
	return &DerivedStats{
		HPMax:                 ds.HPMax,
		Stamina:               ds.Stamina,
		Speed:                 ds.Speed,
		Haste:                 ds.Haste,
		CritChance:            ds.CritChance,
		CritMulti:             ds.CritMulti,
		MoveSpeed:             ds.MoveSpeed,
		RegenHP:               ds.RegenHP,
		Accuracy:              ds.Accuracy,
		Penetration:           ds.Penetration,
		Lethality:             ds.Lethality,
		Brutality:             ds.Brutality,
		ArmorClass:            ds.ArmorClass,
		Evasion:               ds.Evasion,
		BlockChance:           ds.BlockChance,
		ParryChance:           ds.ParryChance,
		DodgeChance:           ds.DodgeChance,
		EnergyEfficiency:      ds.EnergyEfficiency,
		EnergyCapacity:        ds.EnergyCapacity,
		EnergyDrain:           ds.EnergyDrain,
		ResourceRegen:         ds.ResourceRegen,
		ResourceDecay:         ds.ResourceDecay,
		LearningRate:          ds.LearningRate,
		Adaptation:            ds.Adaptation,
		Memory:                ds.Memory,
		Experience:            ds.Experience,
		Leadership:            ds.Leadership,
		Diplomacy:             ds.Diplomacy,
		Intimidation:          ds.Intimidation,
		Empathy:               ds.Empathy,
		Deception:             ds.Deception,
		Performance:           ds.Performance,
		ManaEfficiency:        ds.ManaEfficiency,
		SpellPower:            ds.SpellPower,
		MysticResonance:       ds.MysticResonance,
		RealityBend:           ds.RealityBend,
		TimeSense:             ds.TimeSense,
		SpaceSense:            ds.SpaceSense,
		JumpHeight:            ds.JumpHeight,
		ClimbSpeed:            ds.ClimbSpeed,
		SwimSpeed:             ds.SwimSpeed,
		FlightSpeed:           ds.FlightSpeed,
		TeleportRange:         ds.TeleportRange,
		Stealth:               ds.Stealth,
		AuraRadius:            ds.AuraRadius,
		AuraStrength:          ds.AuraStrength,
		Presence:              ds.Presence,
		Awe:                   ds.Awe,
		WeaponMastery:         ds.WeaponMastery,
		SkillLevel:            ds.SkillLevel,
		LifeSteal:             ds.LifeSteal,
		CastSpeed:             ds.CastSpeed,
		WeightCapacity:        ds.WeightCapacity,
		Persuasion:            ds.Persuasion,
		MerchantPriceModifier: ds.MerchantPriceModifier,
		FactionReputationGain: ds.FactionReputationGain,
		CultivationSpeed:      ds.CultivationSpeed,
		EnergyEfficiencyAmp:   ds.EnergyEfficiencyAmp,
		BreakthroughSuccess:   ds.BreakthroughSuccess,
		SkillLearning:         ds.SkillLearning,
		CombatEffectiveness:   ds.CombatEffectiveness,
		ResourceGathering:     ds.ResourceGathering,
		CreatedAt:             ds.CreatedAt,
		UpdatedAt:             ds.UpdatedAt,
		Version:               ds.Version,
	}
}

// GetVersion returns the current version
func (ds *DerivedStats) GetVersion() int64 {
	return ds.Version
}

// GetUpdatedAt returns the last update timestamp
func (ds *DerivedStats) GetUpdatedAt() int64 {
	return ds.UpdatedAt
}

// GetCreatedAt returns the creation timestamp
func (ds *DerivedStats) GetCreatedAt() int64 {
	return ds.CreatedAt
}
