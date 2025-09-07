package constants

// Derived Stats Constants
const (
	// Core Derived Stats
	STAT_HP_MAX      = "hpMax"
	STAT_STAMINA     = "stamina"
	STAT_SPEED       = "speed"
	STAT_HASTE       = "haste"
	STAT_CRIT_CHANCE = "critChance"
	STAT_CRIT_MULTI  = "critMulti"
	STAT_MOVE_SPEED  = "moveSpeed"
	STAT_REGEN_HP    = "regenHP"

	// Combat Stats
	STAT_ACCURACY     = "accuracy"
	STAT_PENETRATION  = "penetration"
	STAT_LETHALITY    = "lethality"
	STAT_BRUTALITY    = "brutality"
	STAT_ARMOR_CLASS  = "armorClass"
	STAT_EVASION      = "evasion"
	STAT_BLOCK_CHANCE = "blockChance"
	STAT_PARRY_CHANCE = "parryChance"
	STAT_DODGE_CHANCE = "dodgeChance"

	// Energy Stats
	STAT_ENERGY_EFFICIENCY = "energyEfficiency"
	STAT_ENERGY_CAPACITY   = "energyCapacity"
	STAT_ENERGY_DRAIN      = "energyDrain"
	STAT_RESOURCE_REGEN    = "resourceRegen"
	STAT_RESOURCE_DECAY    = "resourceDecay"

	// Learning Stats
	STAT_LEARNING_RATE = "learningRate"
	STAT_ADAPTATION    = "adaptation"
	STAT_MEMORY        = "memory"
	STAT_EXPERIENCE    = "experience"

	// Social Stats
	STAT_LEADERSHIP   = "leadership"
	STAT_DIPLOMACY    = "diplomacy"
	STAT_INTIMIDATION = "intimidation"
	STAT_EMPATHY      = "empathy"
	STAT_DECEPTION    = "deception"
	STAT_PERFORMANCE  = "performance"

	// Mystical Stats
	STAT_MANA_EFFICIENCY  = "manaEfficiency"
	STAT_SPELL_POWER      = "spellPower"
	STAT_MYSTIC_RESONANCE = "mysticResonance"
	STAT_REALITY_BEND     = "realityBend"
	STAT_TIME_SENSE       = "timeSense"
	STAT_SPACE_SENSE      = "spaceSense"

	// Movement Stats
	STAT_JUMP_HEIGHT    = "jumpHeight"
	STAT_CLIMB_SPEED    = "climbSpeed"
	STAT_SWIM_SPEED     = "swimSpeed"
	STAT_FLIGHT_SPEED   = "flightSpeed"
	STAT_TELEPORT_RANGE = "teleportRange"
	STAT_STEALTH        = "stealth"

	// Aura Stats
	STAT_AURA_RADIUS   = "auraRadius"
	STAT_AURA_STRENGTH = "auraStrength"
	STAT_PRESENCE      = "presence"
	STAT_AWE           = "awe"

	// Proficiency Stats
	STAT_WEAPON_MASTERY          = "weaponMastery"
	STAT_SKILL_LEVEL             = "skillLevel"
	STAT_LIFE_STEAL              = "lifeSteal"
	STAT_CAST_SPEED              = "castSpeed"
	STAT_WEIGHT_CAPACITY         = "weightCapacity"
	STAT_PERSUASION              = "persuasion"
	STAT_MERCHANT_PRICE_MODIFIER = "merchantPriceModifier"
	STAT_FACTION_REPUTATION_GAIN = "factionReputationGain"

	// Talent Amplifiers
	STAT_CULTIVATION_SPEED     = "cultivationSpeed"
	STAT_ENERGY_EFFICIENCY_AMP = "energyEfficiencyAmp"
	STAT_BREAKTHROUGH_SUCCESS  = "breakthroughSuccess"
	STAT_SKILL_LEARNING        = "skillLearning"
	STAT_COMBAT_EFFECTIVENESS  = "combatEffectiveness"
	STAT_RESOURCE_GATHERING    = "resourceGathering"
)

// Derived Stats Display Names
var DerivedStatsDisplayNames = map[string]string{
	STAT_HP_MAX:                  "Maximum Health",
	STAT_STAMINA:                 "Stamina",
	STAT_SPEED:                   "Speed",
	STAT_HASTE:                   "Haste",
	STAT_CRIT_CHANCE:             "Critical Hit Chance",
	STAT_CRIT_MULTI:              "Critical Hit Multiplier",
	STAT_MOVE_SPEED:              "Movement Speed",
	STAT_REGEN_HP:                "Health Regeneration",
	STAT_ACCURACY:                "Accuracy",
	STAT_PENETRATION:             "Penetration",
	STAT_LETHALITY:               "Lethality",
	STAT_BRUTALITY:               "Brutality",
	STAT_ARMOR_CLASS:             "Armor Class",
	STAT_EVASION:                 "Evasion",
	STAT_BLOCK_CHANCE:            "Block Chance",
	STAT_PARRY_CHANCE:            "Parry Chance",
	STAT_DODGE_CHANCE:            "Dodge Chance",
	STAT_ENERGY_EFFICIENCY:       "Energy Efficiency",
	STAT_ENERGY_CAPACITY:         "Energy Capacity",
	STAT_ENERGY_DRAIN:            "Energy Drain",
	STAT_RESOURCE_REGEN:          "Resource Regeneration",
	STAT_RESOURCE_DECAY:          "Resource Decay",
	STAT_LEARNING_RATE:           "Learning Rate",
	STAT_ADAPTATION:              "Adaptation",
	STAT_MEMORY:                  "Memory",
	STAT_EXPERIENCE:              "Experience",
	STAT_LEADERSHIP:              "Leadership",
	STAT_DIPLOMACY:               "Diplomacy",
	STAT_INTIMIDATION:            "Intimidation",
	STAT_EMPATHY:                 "Empathy",
	STAT_DECEPTION:               "Deception",
	STAT_PERFORMANCE:             "Performance",
	STAT_MANA_EFFICIENCY:         "Mana Efficiency",
	STAT_SPELL_POWER:             "Spell Power",
	STAT_MYSTIC_RESONANCE:        "Mystic Resonance",
	STAT_REALITY_BEND:            "Reality Bend",
	STAT_TIME_SENSE:              "Time Sense",
	STAT_SPACE_SENSE:             "Space Sense",
	STAT_JUMP_HEIGHT:             "Jump Height",
	STAT_CLIMB_SPEED:             "Climb Speed",
	STAT_SWIM_SPEED:              "Swim Speed",
	STAT_FLIGHT_SPEED:            "Flight Speed",
	STAT_TELEPORT_RANGE:          "Teleport Range",
	STAT_STEALTH:                 "Stealth",
	STAT_AURA_RADIUS:             "Aura Radius",
	STAT_AURA_STRENGTH:           "Aura Strength",
	STAT_PRESENCE:                "Presence",
	STAT_AWE:                     "Awe",
	STAT_WEAPON_MASTERY:          "Weapon Mastery",
	STAT_SKILL_LEVEL:             "Skill Level",
	STAT_LIFE_STEAL:              "Life Steal",
	STAT_CAST_SPEED:              "Cast Speed",
	STAT_WEIGHT_CAPACITY:         "Weight Capacity",
	STAT_PERSUASION:              "Persuasion",
	STAT_MERCHANT_PRICE_MODIFIER: "Merchant Price Modifier",
	STAT_FACTION_REPUTATION_GAIN: "Faction Reputation Gain",
	STAT_CULTIVATION_SPEED:       "Cultivation Speed",
	STAT_ENERGY_EFFICIENCY_AMP:   "Energy Efficiency Amplifier",
	STAT_BREAKTHROUGH_SUCCESS:    "Breakthrough Success",
	STAT_SKILL_LEARNING:          "Skill Learning",
	STAT_COMBAT_EFFECTIVENESS:    "Combat Effectiveness",
	STAT_RESOURCE_GATHERING:      "Resource Gathering",
}

// Derived Stats Categories
const (
	CATEGORY_CORE_DERIVED      = "core_derived"
	CATEGORY_COMBAT_STATS      = "combat"
	CATEGORY_ENERGY_STATS      = "energy"
	CATEGORY_LEARNING_STATS    = "learning"
	CATEGORY_SOCIAL_STATS      = "social"
	CATEGORY_MYSTICAL_STATS    = "mystical"
	CATEGORY_MOVEMENT_STATS    = "movement"
	CATEGORY_AURA_STATS        = "aura"
	CATEGORY_PROFICIENCY_STATS = "proficiency"
	CATEGORY_TALENT_AMPLIFIERS = "talent_amplifiers"
)

// Derived Stats by Category
var DerivedStatsByCategory = map[string][]string{
	CATEGORY_CORE_DERIVED: {
		STAT_HP_MAX, STAT_STAMINA, STAT_SPEED, STAT_HASTE,
		STAT_CRIT_CHANCE, STAT_CRIT_MULTI, STAT_MOVE_SPEED, STAT_REGEN_HP,
	},
	CATEGORY_COMBAT_STATS: {
		STAT_ACCURACY, STAT_PENETRATION, STAT_LETHALITY, STAT_BRUTALITY,
		STAT_ARMOR_CLASS, STAT_EVASION, STAT_BLOCK_CHANCE, STAT_PARRY_CHANCE, STAT_DODGE_CHANCE,
	},
	CATEGORY_ENERGY_STATS: {
		STAT_ENERGY_EFFICIENCY, STAT_ENERGY_CAPACITY, STAT_ENERGY_DRAIN,
		STAT_RESOURCE_REGEN, STAT_RESOURCE_DECAY,
	},
	CATEGORY_LEARNING_STATS: {
		STAT_LEARNING_RATE, STAT_ADAPTATION, STAT_MEMORY, STAT_EXPERIENCE,
	},
	CATEGORY_SOCIAL_STATS: {
		STAT_LEADERSHIP, STAT_DIPLOMACY, STAT_INTIMIDATION, STAT_EMPATHY,
		STAT_DECEPTION, STAT_PERFORMANCE,
	},
	CATEGORY_MYSTICAL_STATS: {
		STAT_MANA_EFFICIENCY, STAT_SPELL_POWER, STAT_MYSTIC_RESONANCE,
		STAT_REALITY_BEND, STAT_TIME_SENSE, STAT_SPACE_SENSE,
	},
	CATEGORY_MOVEMENT_STATS: {
		STAT_JUMP_HEIGHT, STAT_CLIMB_SPEED, STAT_SWIM_SPEED,
		STAT_FLIGHT_SPEED, STAT_TELEPORT_RANGE, STAT_STEALTH,
	},
	CATEGORY_AURA_STATS: {
		STAT_AURA_RADIUS, STAT_AURA_STRENGTH, STAT_PRESENCE, STAT_AWE,
	},
	CATEGORY_PROFICIENCY_STATS: {
		STAT_WEAPON_MASTERY, STAT_SKILL_LEVEL, STAT_LIFE_STEAL, STAT_CAST_SPEED,
		STAT_WEIGHT_CAPACITY, STAT_PERSUASION, STAT_MERCHANT_PRICE_MODIFIER, STAT_FACTION_REPUTATION_GAIN,
	},
	CATEGORY_TALENT_AMPLIFIERS: {
		STAT_CULTIVATION_SPEED, STAT_ENERGY_EFFICIENCY_AMP, STAT_BREAKTHROUGH_SUCCESS,
		STAT_SKILL_LEARNING, STAT_COMBAT_EFFECTIVENESS, STAT_RESOURCE_GATHERING,
	},
}
