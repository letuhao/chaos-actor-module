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
ADD_FLAT    ModifierStacking = "ADD_FLAT"
ADD_PCT     ModifierStacking = "ADD_PCT"
MULTIPLY    ModifierStacking = "MULTIPLY"
OVERRIDE    ModifierStacking = "OVERRIDE"
)

type ModifierSourceRef struct {
Kind  string
ID    string
Label string
}

type StatModifier struct {
Key        StatKey
Op         ModifierStacking
Value      float64
Source     ModifierSourceRef
Priority   int
}

type StatSnapshot struct {
ActorID string
Stats   map[StatKey]float64
Version int
Ts      int64
}

type StatDef struct {
Key         StatKey
Category    string
DisplayName string
Description string
Rounding    string
}

type SnapshotOptions struct {
WithBreakdown bool
}

type ComputeInput struct {
ActorID         string
Level           int
BaseAllocations map[StatKey]int
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
