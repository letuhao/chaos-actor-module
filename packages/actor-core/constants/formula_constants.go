package constants

// Formula calculation constants
const (
	// Core Stats Multipliers
	HPMaxVitalityMultiplier     = 10.0
	HPMaxConstitutionMultiplier = 5.0
	StaminaEnduranceMultiplier  = 10.0
	StaminaConstitutionMultiplier = 3.0
	
	// Speed Multipliers
	SpeedAgilityMultiplier      = 0.1
	HasteAgilityMultiplier      = 0.01
	HasteBaseValue              = 1.0
	
	// Critical Hit Multipliers
	CritChanceBaseValue         = 0.05
	CritChanceLuckMultiplier    = 0.001
	CritMultiBaseValue          = 1.5
	CritMultiLuckMultiplier     = 0.01
	
	// Movement Multipliers
	MoveSpeedAgilityMultiplier  = 0.1
	RegenHPVitalityMultiplier   = 0.01
	
	// Combat Stats Multipliers
	AccuracyBaseValue           = 0.8
	AccuracyIntelligenceMultiplier = 0.01
	PenetrationStrengthMultiplier = 0.01
	LethalityStrengthAgilityMultiplier = 0.005
	BrutalityStrengthMultiplier = 0.01
	ArmorClassBaseValue         = 10.0
	ArmorClassConstitutionMultiplier = 0.5
	EvasionAgilityMultiplier    = 0.01
	BlockChanceConstitutionMultiplier = 0.005
	ParryChanceAgilityMultiplier = 0.005
	DodgeChanceAgilityMultiplier = 0.01
	
	// Energy Stats Multipliers
	EnergyEfficiencyBaseValue   = 1.0
	EnergyEfficiencyIntelligenceMultiplier = 0.01
	EnergyDrainWillpowerMultiplier = 0.01
	ResourceRegenVitalityMultiplier = 0.01
	
	// Learning Stats Multipliers
	LearningRateBaseValue       = 1.0
	LearningRateIntelligenceMultiplier = 0.01
	AdaptationBaseValue         = 1.0
	AdaptationWisdomMultiplier  = 0.01
	MemoryBaseValue             = 1.0
	MemoryIntelligenceMultiplier = 0.01
	ExperienceBaseValue         = 1.0
	ExperienceWisdomMultiplier  = 0.01
	
	// Social Stats Multipliers
	LeadershipBaseValue         = 1.0
	LeadershipCharismaMultiplier = 0.01
	DiplomacyBaseValue          = 1.0
	DiplomacyCharismaIntelligenceMultiplier = 0.005
	IntimidationBaseValue       = 1.0
	IntimidationStrengthCharismaMultiplier = 0.005
	EmpathyBaseValue            = 1.0
	EmpathyWisdomCharismaMultiplier = 0.005
	DeceptionBaseValue          = 1.0
	DeceptionIntelligenceCharismaMultiplier = 0.005
	PerformanceBaseValue        = 1.0
	PerformanceCharismaMultiplier = 0.01
	
	// Mystical Stats Multipliers
	ManaEfficiencyBaseValue     = 1.0
	ManaEfficiencyIntelligenceMultiplier = 0.01
	SpellPowerBaseValue         = 1.0
	SpellPowerSpiritualEnergyMultiplier = 0.01
	MysticResonanceBaseValue    = 1.0
	MysticResonanceSpiritualMentalMultiplier = 0.005
	RealityBendBaseValue        = 1.0
	RealityBendMentalEnergyMultiplier = 0.01
	TimeSenseBaseValue          = 1.0
	TimeSenseMentalEnergyMultiplier = 0.01
	SpaceSenseBaseValue         = 1.0
	SpaceSenseMentalEnergyMultiplier = 0.01
	
	// Movement Stats Multipliers
	JumpHeightBaseValue         = 1.0
	JumpHeightStrengthAgilityMultiplier = 0.005
	ClimbSpeedBaseValue         = 1.0
	ClimbSpeedStrengthAgilityMultiplier = 0.005
	SwimSpeedBaseValue          = 1.0
	SwimSpeedStrengthAgilityMultiplier = 0.005
	FlightSpeedBaseValue        = 1.0
	FlightSpeedSpiritualEnergyMultiplier = 0.01
	TeleportRangeBaseValue      = 1.0
	TeleportRangeMentalEnergyMultiplier = 0.01
	StealthBaseValue            = 1.0
	StealthAgilityMultiplier    = 0.01
	
	// Aura Stats Multipliers
	AuraRadiusBaseValue         = 1.0
	AuraRadiusSpiritualEnergyMultiplier = 0.01
	AuraStrengthBaseValue       = 1.0
	AuraStrengthSpiritualEnergyMultiplier = 0.01
	PresenceBaseValue           = 1.0
	PresenceCharismaSpiritualMultiplier = 0.005
	AweBaseValue                = 1.0
	AweCharismaSpiritualMultiplier = 0.005
	
	// Proficiency Stats Multipliers
	WeaponMasteryBaseValue      = 1.0
	WeaponMasteryStrengthAgilityMultiplier = 0.005
	SkillLevelBaseValue         = 1.0
	SkillLevelIntelligenceMultiplier = 0.01
	LifeStealDefaultValue       = 0.0
	CastSpeedBaseValue          = 1.0
	CastSpeedIntelligenceMultiplier = 0.01
	WeightCapacityStrengthMultiplier = 10.0
	PersuasionBaseValue         = 1.0
	PersuasionCharismaMultiplier = 0.01
	MerchantPriceModifierBaseValue = 1.0
	MerchantPriceModifierCharismaMultiplier = 0.01
	FactionReputationGainBaseValue = 1.0
	FactionReputationGainCharismaMultiplier = 0.01
	
	// Talent Amplifiers Multipliers
	CultivationSpeedBaseValue   = 1.0
	CultivationSpeedAllEnergyMultiplier = 0.001
	EnergyEfficiencyAmpBaseValue = 1.0
	EnergyEfficiencyAmpIntelligenceMultiplier = 0.01
	BreakthroughSuccessBaseValue = 1.0
	BreakthroughSuccessWillpowerLuckMultiplier = 0.005
	SkillLearningBaseValue      = 1.0
	SkillLearningIntelligenceWisdomMultiplier = 0.005
	CombatEffectivenessBaseValue = 1.0
	CombatEffectivenessStrengthAgilityIntelligenceMultiplier = 0.003
	ResourceGatheringBaseValue  = 1.0
	ResourceGatheringLuckWisdomMultiplier = 0.005
	
	// Validation Constants
	MinHasteValue               = 0.1
	MaxCritChanceValue          = 1.0
	MinCritMultiValue           = 1.0
	MinStatValue                = 0.0
)
