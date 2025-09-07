package constants

// Primary Stats Constants
const (
	// Basic Stats
	STAT_VITALITY     = "vitality"
	STAT_ENDURANCE    = "endurance"
	STAT_CONSTITUTION = "constitution"
	STAT_INTELLIGENCE = "intelligence"
	STAT_WISDOM       = "wisdom"
	STAT_CHARISMA     = "charisma"
	STAT_WILLPOWER    = "willpower"
	STAT_LUCK         = "luck"
	STAT_FATE         = "fate"
	STAT_KARMA        = "karma"

	// Physical Stats
	STAT_STRENGTH    = "strength"
	STAT_AGILITY     = "agility"
	STAT_PERSONALITY = "personality"

	// Universal Cultivation Stats
	STAT_SPIRITUAL_ENERGY    = "spiritualEnergy"
	STAT_PHYSICAL_ENERGY     = "physicalEnergy"
	STAT_MENTAL_ENERGY       = "mentalEnergy"
	STAT_CULTIVATION_LEVEL   = "cultivationLevel"
	STAT_BREAKTHROUGH_POINTS = "breakthroughPoints"

	// Life Stats
	STAT_LIFE_SPAN = "lifeSpan"
	STAT_AGE       = "age"
)

// Primary Stats Display Names
var PrimaryStatsDisplayNames = map[string]string{
	STAT_VITALITY:            "Vitality",
	STAT_ENDURANCE:           "Endurance",
	STAT_CONSTITUTION:        "Constitution",
	STAT_INTELLIGENCE:        "Intelligence",
	STAT_WISDOM:              "Wisdom",
	STAT_CHARISMA:            "Charisma",
	STAT_WILLPOWER:           "Willpower",
	STAT_LUCK:                "Luck",
	STAT_FATE:                "Fate",
	STAT_KARMA:               "Karma",
	STAT_STRENGTH:            "Strength",
	STAT_AGILITY:             "Agility",
	STAT_PERSONALITY:         "Personality",
	STAT_SPIRITUAL_ENERGY:    "Spiritual Energy",
	STAT_PHYSICAL_ENERGY:     "Physical Energy",
	STAT_MENTAL_ENERGY:       "Mental Energy",
	STAT_CULTIVATION_LEVEL:   "Cultivation Level",
	STAT_BREAKTHROUGH_POINTS: "Breakthrough Points",
	STAT_LIFE_SPAN:           "Life Span",
	STAT_AGE:                 "Age",
}

// Primary Stats Descriptions
var PrimaryStatsDescriptions = map[string]string{
	STAT_VITALITY:            "Overall health and resilience",
	STAT_ENDURANCE:           "Physical stamina and recovery",
	STAT_CONSTITUTION:        "Body's natural resistance to damage",
	STAT_INTELLIGENCE:        "Problem solving, memory, learning",
	STAT_WISDOM:              "Insight, perception, decision making",
	STAT_CHARISMA:            "Social influence, leadership, presence",
	STAT_WILLPOWER:           "Mental resistance, focus, determination",
	STAT_LUCK:                "General luck and fortune",
	STAT_FATE:                "Destiny and fate manipulation",
	STAT_KARMA:               "Good/evil alignment influence",
	STAT_STRENGTH:            "Physical power and damage",
	STAT_AGILITY:             "Speed and precision",
	STAT_PERSONALITY:         "Social influence and charisma",
	STAT_SPIRITUAL_ENERGY:    "Universal spiritual energy for all cultivation systems",
	STAT_PHYSICAL_ENERGY:     "Universal physical energy for body cultivation",
	STAT_MENTAL_ENERGY:       "Universal mental energy for magic and mental cultivation",
	STAT_CULTIVATION_LEVEL:   "Overall cultivation level across all systems",
	STAT_BREAKTHROUGH_POINTS: "Universal breakthrough points for advancement",
	STAT_LIFE_SPAN:           "Maximum lifespan of the actor",
	STAT_AGE:                 "Current age of the actor",
}

// Primary Stats Categories
const (
	CATEGORY_BASIC_STATS       = "basic"
	CATEGORY_PHYSICAL_STATS    = "physical"
	CATEGORY_CULTIVATION_STATS = "cultivation"
	CATEGORY_LIFE_STATS        = "life"
)

// Primary Stats by Category
var PrimaryStatsByCategory = map[string][]string{
	CATEGORY_BASIC_STATS: {
		STAT_VITALITY, STAT_ENDURANCE, STAT_CONSTITUTION,
		STAT_INTELLIGENCE, STAT_WISDOM, STAT_CHARISMA,
		STAT_WILLPOWER, STAT_LUCK, STAT_FATE, STAT_KARMA,
	},
	CATEGORY_PHYSICAL_STATS: {
		STAT_STRENGTH, STAT_AGILITY, STAT_PERSONALITY,
	},
	CATEGORY_CULTIVATION_STATS: {
		STAT_SPIRITUAL_ENERGY, STAT_PHYSICAL_ENERGY, STAT_MENTAL_ENERGY,
		STAT_CULTIVATION_LEVEL, STAT_BREAKTHROUGH_POINTS,
	},
	CATEGORY_LIFE_STATS: {
		STAT_LIFE_SPAN, STAT_AGE,
	},
}
