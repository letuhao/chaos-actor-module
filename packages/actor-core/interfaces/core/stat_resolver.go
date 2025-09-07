package core

import "context"

// StatResolver defines the interface for resolving stat calculations
type StatResolver interface {
	// ResolveStats resolves all stats from primary stats
	ResolveStats(primaryStats map[string]int64) (*StatSnapshot, error)

	// ResolveStatsWithContext resolves all stats from primary stats with context
	ResolveStatsWithContext(ctx context.Context, primaryStats map[string]int64) (*StatSnapshot, error)

	// ResolveStat resolves a specific stat
	ResolveStat(statName string, primaryStats map[string]int64) (interface{}, error)

	// ResolveStatWithContext resolves a specific stat with context
	ResolveStatWithContext(ctx context.Context, statName string, primaryStats map[string]int64) (interface{}, error)

	// ResolveDerivedStats resolves all derived stats
	ResolveDerivedStats(primaryStats map[string]int64) (map[string]float64, error)

	// ResolveDerivedStatsWithContext resolves all derived stats with context
	ResolveDerivedStatsWithContext(ctx context.Context, primaryStats map[string]int64) (map[string]float64, error)

	// ResolveFlexibleStats resolves flexible stats
	ResolveFlexibleStats(primaryStats map[string]int64) (*FlexibleStats, error)

	// ResolveFlexibleStatsWithContext resolves flexible stats with context
	ResolveFlexibleStatsWithContext(ctx context.Context, primaryStats map[string]int64) (*FlexibleStats, error)

	// CheckDependencies checks if all dependencies are available
	CheckDependencies(statName string) ([]string, error)

	// CheckDependenciesWithContext checks if all dependencies are available with context
	CheckDependenciesWithContext(ctx context.Context, statName string) ([]string, error)

	// GetCalculationOrder returns the order in which stats should be calculated
	GetCalculationOrder() []string

	// GetCalculationOrderWithContext returns the order in which stats should be calculated with context
	GetCalculationOrderWithContext(ctx context.Context) []string

	// ValidateStats validates all stats
	ValidateStats(stats map[string]interface{}) error

	// ValidateStatsWithContext validates all stats with context
	ValidateStatsWithContext(ctx context.Context, stats map[string]interface{}) error
}

// FlexibleStats represents flexible stats that can be customized
type FlexibleStats struct {
	CustomPrimary        map[string]int64              `json:"custom_primary"`
	CustomDerived        map[string]float64            `json:"custom_derived"`
	SubSystemStats       map[string]map[string]float64 `json:"subsystem_stats"`
	SpeedSystem          *FlexibleSpeedSystem          `json:"speed_system"`
	KarmaSystem          *FlexibleKarmaSystem          `json:"karma_system"`
	AdministrativeSystem *FlexibleAdministrativeSystem `json:"administrative_system"`
	ProficiencySystem    *ProficiencySystem            `json:"proficiency_system"`
	SkillSystem          *UniversalSkillSystem         `json:"skill_system"`
}

// FlexibleSpeedSystem represents the flexible speed system
type FlexibleSpeedSystem struct {
	MovementSpeeds       map[string]float64 `json:"movement_speeds"`
	CastingSpeeds        map[string]float64 `json:"casting_speeds"`
	CraftingSpeeds       map[string]float64 `json:"crafting_speeds"`
	LearningSpeeds       map[string]float64 `json:"learning_speeds"`
	CombatSpeeds         map[string]float64 `json:"combat_speeds"`
	SocialSpeeds         map[string]float64 `json:"social_speeds"`
	AdministrativeSpeeds map[string]float64 `json:"administrative_speeds"`
}

// FlexibleKarmaSystem represents the flexible karma system
type FlexibleKarmaSystem struct {
	DivisionKarma   map[string]map[string]int64 `json:"division_karma"`
	KarmaCategories map[string][]string         `json:"karma_categories"`
	KarmaTypes      map[string]KarmaType        `json:"karma_types"`
	KarmaInfluence  map[string]float64          `json:"karma_influence"`
}

// KarmaType represents a type of karma
type KarmaType struct {
	Name        string             `json:"name"`
	Category    string             `json:"category"`
	Description string             `json:"description"`
	Influence   map[string]float64 `json:"influence"`
	DecayRate   float64            `json:"decay_rate"`
	MaxValue    int64              `json:"max_value"`
	MinValue    int64              `json:"min_value"`
}

// FlexibleAdministrativeSystem represents the flexible administrative system
type FlexibleAdministrativeSystem struct {
	Divisions     map[string]map[string]AdministrativeDivision `json:"divisions"`
	DivisionTypes map[string]DivisionType                      `json:"division_types"`
	Hierarchy     map[string][]string                          `json:"hierarchy"`
	Relationships map[string]map[string]DivisionRelationship   `json:"relationships"`
}

// AdministrativeDivision represents an administrative division
type AdministrativeDivision struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	Parent     string                 `json:"parent"`
	Children   []string               `json:"children"`
	Attributes map[string]interface{} `json:"attributes"`
	Level      int                    `json:"level"`
	IsActive   bool                   `json:"is_active"`
	CreatedAt  int64                  `json:"created_at"`
	UpdatedAt  int64                  `json:"updated_at"`
}

// DivisionType represents a type of administrative division
type DivisionType struct {
	Name        string                 `json:"name"`
	Category    string                 `json:"category"`
	Description string                 `json:"description"`
	Attributes  map[string]interface{} `json:"attributes"`
	MaxLevel    int                    `json:"max_level"`
	MinLevel    int                    `json:"min_level"`
}

// DivisionRelationship represents a relationship between divisions
type DivisionRelationship struct {
	Type        string  `json:"type"`
	Strength    float64 `json:"strength"`
	Description string  `json:"description"`
	IsActive    bool    `json:"is_active"`
}

// ProficiencySystem represents the proficiency system
type ProficiencySystem struct {
	Proficiencies map[string]*Proficiency `json:"proficiencies"`
	Categories    map[string][]string     `json:"categories"`
	MaxSkills     int                     `json:"max_skills"`
}

// Proficiency represents a proficiency
type Proficiency struct {
	SkillName  string  `json:"skill_name"`
	Category   string  `json:"category"`
	Level      int64   `json:"level"`
	Experience int64   `json:"experience"`
	MaxLevel   int64   `json:"max_level"`
	Multiplier float64 `json:"multiplier"`
	LastUsed   int64   `json:"last_used"`
	TotalUses  int64   `json:"total_uses"`
}

// UniversalSkillSystem represents the universal skill system
type UniversalSkillSystem struct {
	Skills     map[string]*UniversalSkill `json:"skills"`
	Categories map[string][]string        `json:"categories"`
	SkillTrees map[string]*SkillTree      `json:"skill_trees"`
	MaxSkills  int                        `json:"max_skills"`
}

// UniversalSkill represents a universal skill
type UniversalSkill struct {
	Name         string             `json:"name"`
	Category     string             `json:"category"`
	SubCategory  string             `json:"sub_category"`
	Level        int64              `json:"level"`
	Experience   int64              `json:"experience"`
	MaxLevel     int64              `json:"max_level"`
	Requirements []string           `json:"requirements"`
	Bonuses      map[string]float64 `json:"bonuses"`
	Cooldown     int64              `json:"cooldown"`
	ManaCost     float64            `json:"mana_cost"`
	StaminaCost  float64            `json:"stamina_cost"`
}

// SkillTree represents a skill tree
type SkillTree struct {
	Name          string                     `json:"name"`
	Category      string                     `json:"category"`
	Skills        map[string]*UniversalSkill `json:"skills"`
	Prerequisites map[string][]string        `json:"prerequisites"`
	Unlocks       map[string][]string        `json:"unlocks"`
}
