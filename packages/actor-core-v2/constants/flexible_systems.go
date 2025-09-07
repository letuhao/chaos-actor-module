package constants

// Formula Types
const (
	FormulaTypeCalculation = "calculation"
)

// Flexible Systems Constants
const (
	// Speed System Categories
	SPEED_MOVEMENT       = "movement"
	SPEED_CASTING        = "casting"
	SPEED_CRAFTING       = "crafting"
	SPEED_LEARNING       = "learning"
	SPEED_COMBAT         = "combat"
	SPEED_SOCIAL         = "social"
	SPEED_ADMINISTRATIVE = "administrative"

	// Speed Types
	SPEED_WALKING           = "walking"
	SPEED_RUNNING           = "running"
	SPEED_SWIMMING          = "swimming"
	SPEED_CLIMBING          = "climbing"
	SPEED_FLYING            = "flying"
	SPEED_TELEPORT          = "teleport"
	SPEED_SPELL_CASTING     = "spell_casting"
	SPEED_TECHNIQUE_CASTING = "technique_casting"
	SPEED_ALCHEMY           = "alchemy"
	SPEED_FORGING           = "forging"
	SPEED_ENCHANTING        = "enchanting"
	SPEED_READING           = "reading"
	SPEED_STUDYING          = "studying"
	SPEED_PRACTICING        = "practicing"
	SPEED_ATTACK            = "attack"
	SPEED_DEFENSE           = "defense"
	SPEED_CONVERSATION      = "conversation"
	SPEED_NEGOTIATION       = "negotiation"
	SPEED_ADMINISTRATION    = "administration"
	SPEED_MANAGEMENT        = "management"

	// Karma System Categories
	KARMA_FORTUNE      = "fortune"
	KARMA_KARMA        = "karma"
	KARMA_MERIT        = "merit"
	KARMA_CONTRIBUTION = "contribution"
	KARMA_REPUTATION   = "reputation"
	KARMA_HONOR        = "honor"
	KARMA_DISHONOR     = "dishonor"

	// Administrative Division Types
	DIVISION_WORLD      = "world"
	DIVISION_CONTINENT  = "continent"
	DIVISION_NATION     = "nation"
	DIVISION_RACE       = "race"
	DIVISION_SECT       = "sect"
	DIVISION_REALM      = "realm"
	DIVISION_ZONE       = "zone"
	DIVISION_CITY       = "city"
	DIVISION_VILLAGE    = "village"
	DIVISION_CLAN       = "clan"
	DIVISION_FAMILY     = "family"
	DIVISION_INDIVIDUAL = "individual"

	// Proficiency Categories
	PROFICIENCY_COMBAT         = "combat"
	PROFICIENCY_MAGIC          = "magic"
	PROFICIENCY_CRAFTING       = "crafting"
	PROFICIENCY_SOCIAL         = "social"
	PROFICIENCY_MOVEMENT       = "movement"
	PROFICIENCY_SURVIVAL       = "survival"
	PROFICIENCY_LEARNING       = "learning"
	PROFICIENCY_ADMINISTRATION = "administration"

	// Skill Categories
	SKILL_COMBAT      = "combat"
	SKILL_MAGIC       = "magic"
	SKILL_PROFESSION  = "profession"
	SKILL_SOCIAL      = "social"
	SKILL_MOVEMENT    = "movement"
	SKILL_SURVIVAL    = "survival"
	SKILL_CULTIVATION = "cultivation"
	SKILL_CRAFTING    = "crafting"

	// Energy Types
	ENERGY_MP         = "mp"
	ENERGY_QI         = "qi"
	ENERGY_LUST       = "lust"
	ENERGY_WRATH      = "wrath"
	ENERGY_SPIRITUAL  = "spiritual"
	ENERGY_PHYSICAL   = "physical"
	ENERGY_MENTAL     = "mental"
	ENERGY_DIVINE     = "divine"
	ENERGY_DEMONIC    = "demonic"
	ENERGY_NECROMANCY = "necromancy"
	ENERGY_ELEMENTAL  = "elemental"

	// Damage Types
	DAMAGE_PHYSICAL   = "physical"
	DAMAGE_FIRE       = "fire"
	DAMAGE_WATER      = "water"
	DAMAGE_EARTH      = "earth"
	DAMAGE_AIR        = "air"
	DAMAGE_LIGHTNING  = "lightning"
	DAMAGE_ICE        = "ice"
	DAMAGE_POISON     = "poison"
	DAMAGE_DARK       = "dark"
	DAMAGE_LIGHT      = "light"
	DAMAGE_MAGICAL    = "magical"
	DAMAGE_PSYCHIC    = "psychic"
	DAMAGE_TIME       = "time"
	DAMAGE_SPACE      = "space"
	DAMAGE_REALITY    = "reality"
	DAMAGE_CONCEPTUAL = "conceptual"

	// Defence Types
	DEFENCE_ARMOR      = "armor"
	DEFENCE_SHIELD     = "shield"
	DEFENCE_MAGIC      = "magic"
	DEFENCE_SPIRITUAL  = "spiritual"
	DEFENCE_MENTAL     = "mental"
	DEFENCE_ELEMENTAL  = "elemental"
	DEFENCE_DIVINE     = "divine"
	DEFENCE_DEMONIC    = "demonic"
	DEFENCE_NECROMANCY = "necromancy"
	DEFENCE_TIME       = "time"
	DEFENCE_SPACE      = "space"
	DEFENCE_REALITY    = "reality"
	DEFENCE_CONCEPTUAL = "conceptual"

	// Amplifier Types
	AMPLIFIER_MULTIPLIER = "multiplier"
	AMPLIFIER_PIERCING   = "piercing"
	AMPLIFIER_CRITICAL   = "critical"
	AMPLIFIER_LETHAL     = "lethal"
	AMPLIFIER_BRUTAL     = "brutal"
	AMPLIFIER_ELEMENTAL  = "elemental"
	AMPLIFIER_MYSTICAL   = "mystical"
	AMPLIFIER_DIVINE     = "divine"
	AMPLIFIER_DEMONIC    = "demonic"
	AMPLIFIER_NECROMANCY = "necromancy"
	AMPLIFIER_TIME       = "time"
	AMPLIFIER_SPACE      = "space"
	AMPLIFIER_REALITY    = "reality"
	AMPLIFIER_CONCEPTUAL = "conceptual"
)

// Speed System Display Names
var SpeedSystemDisplayNames = map[string]string{
	SPEED_MOVEMENT:       "Movement Speed",
	SPEED_CASTING:        "Casting Speed",
	SPEED_CRAFTING:       "Crafting Speed",
	SPEED_LEARNING:       "Learning Speed",
	SPEED_COMBAT:         "Combat Speed",
	SPEED_SOCIAL:         "Social Speed",
	SPEED_ADMINISTRATIVE: "Administrative Speed",
}

// Speed Types Display Names
var SpeedTypesDisplayNames = map[string]string{
	SPEED_WALKING:           "Walking",
	SPEED_RUNNING:           "Running",
	SPEED_SWIMMING:          "Swimming",
	SPEED_CLIMBING:          "Climbing",
	SPEED_FLYING:            "Flying",
	SPEED_TELEPORT:          "Teleport",
	SPEED_SPELL_CASTING:     "Spell Casting",
	SPEED_TECHNIQUE_CASTING: "Technique Casting",
	SPEED_ALCHEMY:           "Alchemy",
	SPEED_FORGING:           "Forging",
	SPEED_ENCHANTING:        "Enchanting",
	SPEED_READING:           "Reading",
	SPEED_STUDYING:          "Studying",
	SPEED_PRACTICING:        "Practicing",
	SPEED_ATTACK:            "Attack",
	SPEED_DEFENSE:           "Defense",
	SPEED_CONVERSATION:      "Conversation",
	SPEED_NEGOTIATION:       "Negotiation",
	SPEED_ADMINISTRATION:    "Administration",
	SPEED_MANAGEMENT:        "Management",
}

// Karma System Display Names
var KarmaSystemDisplayNames = map[string]string{
	KARMA_FORTUNE:      "Fortune",
	KARMA_KARMA:        "Karma",
	KARMA_MERIT:        "Merit",
	KARMA_CONTRIBUTION: "Contribution",
	KARMA_REPUTATION:   "Reputation",
	KARMA_HONOR:        "Honor",
	KARMA_DISHONOR:     "Dishonor",
}

// Administrative Division Display Names
var AdministrativeDivisionDisplayNames = map[string]string{
	DIVISION_WORLD:      "World",
	DIVISION_CONTINENT:  "Continent",
	DIVISION_NATION:     "Nation",
	DIVISION_RACE:       "Race",
	DIVISION_SECT:       "Sect",
	DIVISION_REALM:      "Realm",
	DIVISION_ZONE:       "Zone",
	DIVISION_CITY:       "City",
	DIVISION_VILLAGE:    "Village",
	DIVISION_CLAN:       "Clan",
	DIVISION_FAMILY:     "Family",
	DIVISION_INDIVIDUAL: "Individual",
}

// Proficiency Categories Display Names
var ProficiencyCategoriesDisplayNames = map[string]string{
	PROFICIENCY_COMBAT:         "Combat",
	PROFICIENCY_MAGIC:          "Magic",
	PROFICIENCY_CRAFTING:       "Crafting",
	PROFICIENCY_SOCIAL:         "Social",
	PROFICIENCY_MOVEMENT:       "Movement",
	PROFICIENCY_SURVIVAL:       "Survival",
	PROFICIENCY_LEARNING:       "Learning",
	PROFICIENCY_ADMINISTRATION: "Administration",
}

// Skill Categories Display Names
var SkillCategoriesDisplayNames = map[string]string{
	SKILL_COMBAT:      "Combat",
	SKILL_MAGIC:       "Magic",
	SKILL_PROFESSION:  "Profession",
	SKILL_SOCIAL:      "Social",
	SKILL_MOVEMENT:    "Movement",
	SKILL_SURVIVAL:    "Survival",
	SKILL_CULTIVATION: "Cultivation",
	SKILL_CRAFTING:    "Crafting",
}

// Energy Types Display Names
var EnergyTypesDisplayNames = map[string]string{
	ENERGY_MP:         "Mana Points",
	ENERGY_QI:         "Qi",
	ENERGY_LUST:       "Lust",
	ENERGY_WRATH:      "Wrath",
	ENERGY_SPIRITUAL:  "Spiritual Energy",
	ENERGY_PHYSICAL:   "Physical Energy",
	ENERGY_MENTAL:     "Mental Energy",
	ENERGY_DIVINE:     "Divine Energy",
	ENERGY_DEMONIC:    "Demonic Energy",
	ENERGY_NECROMANCY: "Necromancy Energy",
	ENERGY_ELEMENTAL:  "Elemental Energy",
}

// Damage Types Display Names
var DamageTypesDisplayNames = map[string]string{
	DAMAGE_PHYSICAL:   "Physical",
	DAMAGE_FIRE:       "Fire",
	DAMAGE_WATER:      "Water",
	DAMAGE_EARTH:      "Earth",
	DAMAGE_AIR:        "Air",
	DAMAGE_LIGHTNING:  "Lightning",
	DAMAGE_ICE:        "Ice",
	DAMAGE_POISON:     "Poison",
	DAMAGE_DARK:       "Dark",
	DAMAGE_LIGHT:      "Light",
	DAMAGE_MAGICAL:    "Magical",
	DAMAGE_PSYCHIC:    "Psychic",
	DAMAGE_TIME:       "Time",
	DAMAGE_SPACE:      "Space",
	DAMAGE_REALITY:    "Reality",
	DAMAGE_CONCEPTUAL: "Conceptual",
}

// Defence Types Display Names
var DefenceTypesDisplayNames = map[string]string{
	DEFENCE_ARMOR:      "Armor",
	DEFENCE_SHIELD:     "Shield",
	DEFENCE_MAGIC:      "Magic",
	DEFENCE_SPIRITUAL:  "Spiritual",
	DEFENCE_MENTAL:     "Mental",
	DEFENCE_ELEMENTAL:  "Elemental",
	DEFENCE_DIVINE:     "Divine",
	DEFENCE_DEMONIC:    "Demonic",
	DEFENCE_NECROMANCY: "Necromancy",
	DEFENCE_TIME:       "Time",
	DEFENCE_SPACE:      "Space",
	DEFENCE_REALITY:    "Reality",
	DEFENCE_CONCEPTUAL: "Conceptual",
}

// Amplifier Types Display Names
var AmplifierTypesDisplayNames = map[string]string{
	AMPLIFIER_MULTIPLIER: "Multiplier",
	AMPLIFIER_PIERCING:   "Piercing",
	AMPLIFIER_CRITICAL:   "Critical",
	AMPLIFIER_LETHAL:     "Lethal",
	AMPLIFIER_BRUTAL:     "Brutal",
	AMPLIFIER_ELEMENTAL:  "Elemental",
	AMPLIFIER_MYSTICAL:   "Mystical",
	AMPLIFIER_DIVINE:     "Divine",
	AMPLIFIER_DEMONIC:    "Demonic",
	AMPLIFIER_NECROMANCY: "Necromancy",
	AMPLIFIER_TIME:       "Time",
	AMPLIFIER_SPACE:      "Space",
	AMPLIFIER_REALITY:    "Reality",
	AMPLIFIER_CONCEPTUAL: "Conceptual",
}
