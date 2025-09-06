package integration

import (
"rpg-stats/internal/model"
"rpg-stats/internal/resolver"
)

type SnapshotProvider struct {
resolver *resolver.StatResolver
}

func NewSnapshotProvider() *SnapshotProvider {
return &SnapshotProvider{
resolver: resolver.NewStatResolver(),
}
}

func (sp *SnapshotProvider) BuildForActor(actorID string, options *model.SnapshotOptions) (*model.StatSnapshot, error) {
// Sample character with some stat allocations
allocations := map[model.StatKey]int{
model.STR: 15,
model.INT: 12,
model.WIL: 10,
model.AGI: 14,
model.SPD: 13,
model.END: 16,
model.PER: 8,
model.LUK: 11,
}

// Sample equipment modifiers
items := []model.StatModifier{
{
Key:   model.STR,
Op:    model.ADD_FLAT,
Value: 3.0,
Source: model.ModifierSourceRef{
Kind:  "item",
ID:    "iron_sword",
Label: "Iron Sword",
},
Priority: 1,
},
{
Key:   model.ATK,
Op:    model.ADD_FLAT,
Value: 8.0,
Source: model.ModifierSourceRef{
Kind:  "item",
ID:    "iron_sword",
Label: "Iron Sword",
},
Priority: 1,
},
{
Key:   model.HP_MAX,
Op:    model.ADD_PCT,
Value: 10.0, // 10% increase
Source: model.ModifierSourceRef{
Kind:  "item",
ID:    "health_ring",
Label: "Ring of Health",
},
Priority: 1,
},
}

input := model.ComputeInput{
ActorID:         actorID,
Level:           5,
BaseAllocations: allocations,
Registry:        []model.StatDef{},
Items:           items,
Titles:          []model.StatModifier{},
Passives:        []model.StatModifier{},
Buffs:           []model.StatModifier{},
Debuffs:         []model.StatModifier{},
Auras:           []model.StatModifier{},
Environment:     []model.StatModifier{},
WithBreakdown:   false,
}

snapshot := sp.resolver.ComputeSnapshot(input)
return &snapshot, nil
}
