package model

type StatKey string

const (
	// Primary stats
	STR StatKey = "STR"
	INT StatKey = "INT"
	WIL StatKey = "WIL"
	AGI StatKey = "AGI"
	SPD StatKey = "SPD"
	END StatKey = "END"
	PER StatKey = "PER"
	LUK StatKey = "LUK"

	// Derived stats
	HP_MAX      StatKey = "HP_MAX"
	MANA_MAX    StatKey = "MANA_MAX"
	ATK         StatKey = "ATK"
	MATK        StatKey = "MATK"
	DEF         StatKey = "DEF"
	EVASION     StatKey = "EVASION"
	MOVE_SPEED  StatKey = "MOVE_SPEED"
	CRIT_CHANCE StatKey = "CRIT_CHANCE"
	CRIT_DAMAGE StatKey = "CRIT_DAMAGE"
)

type ModifierStacking string

const (
	ADD_FLAT ModifierStacking = "ADD_FLAT"
	ADD_PCT  ModifierStacking = "ADD_PCT"
	MULTIPLY ModifierStacking = "MULTIPLY"
	OVERRIDE ModifierStacking = "OVERRIDE"
)

type ModifierSourceRef struct {
	Kind  string
	ID    string
	Label string
}

type ModifierConditions struct {
	RequiresTagsAll []string
	RequiresTagsAny []string
	ForbidTags      []string
	DurationMs      int64
	StackID         string
	MaxStacks       int64
}

type StatModifier struct {
	Key        StatKey
	Op         ModifierStacking
	Value      float64
	Source     ModifierSourceRef
	Conditions *ModifierConditions
	Priority   int64
}

type StatSnapshot struct {
	ActorID   string
	Stats     map[StatKey]float64
	Breakdown map[StatKey]*StatBreakdown
	Version   int64
	Ts        int64
	Hash      string
}

type OverrideEntry struct {
	Value  float64
	Source ModifierSourceRef
}

type StatBreakdown struct {
	Base           float64
	AdditiveFlat   float64
	AdditivePct    float64
	Multiplicative float64
	Overrides      []OverrideEntry
	CappedTo       *float64
	Notes          []string
}

type StatDef struct {
	Key          StatKey
	Category     string
	DisplayName  string
	Description  string
	Rounding     string
	MinValue     float64
	MaxValue     float64
	DefaultValue float64
	IsPrimary    bool
}

type SnapshotOptions struct {
	WithBreakdown bool
}

type ComputeInput struct {
	ActorID         string
	Level           int64
	BaseAllocations map[StatKey]int64
	Registry        []StatDef
	Items           []StatModifier
	Titles          []StatModifier
	Passives        []StatModifier
	Buffs           []StatModifier
	Debuffs         []StatModifier
	Auras           []StatModifier
	Environment     []StatModifier
	WithBreakdown   bool
}

func PrimaryStats() []StatKey {
	return []StatKey{STR, INT, WIL, AGI, SPD, END, PER, LUK}
}

func (sk StatKey) IsPrimary() bool {
	for _, primary := range PrimaryStats() {
		if sk == primary {
			return true
		}
	}
	return false
}

// Progression types
type PlayerProgress struct {
	ActorID     string
	Level       int64
	XP          int64
	Allocations map[StatKey]int64
	LastUpdated int64
}

type ProgressionResult struct {
	NewLevel      *int64
	PointsGranted *int64
}

type LevelProgression struct {
	Level         int64
	XPRequired    int64
	PointsGranted int64
}
