package resolver

import "rpg-stats/internal/model"

type StatResolver struct{}

func NewStatResolver() *StatResolver {
return &StatResolver{}
}

func (sr *StatResolver) ComputeSnapshot(input model.ComputeInput) model.StatSnapshot {
snapshot := model.StatSnapshot{
ActorID: input.ActorID,
Stats:   make(map[model.StatKey]float64),
Version: 1,
Ts:      1234567890,
}

// Calculate primary stats
for key, value := range input.BaseAllocations {
snapshot.Stats[key] = float64(value)
}

// Calculate some basic derived stats
if str, ok := snapshot.Stats[model.STR]; ok {
snapshot.Stats["HP_MAX"] = 100.0 + str*10.0
snapshot.Stats["ATK"] = str * 2.0
}

if int, ok := snapshot.Stats[model.INT]; ok {
snapshot.Stats["MANA_MAX"] = 80.0 + int*8.0
snapshot.Stats["MATK"] = int * 1.5
}

if end, ok := snapshot.Stats[model.END]; ok {
snapshot.Stats["HP_MAX"] += end * 5.0
snapshot.Stats["DEF"] = end * 1.2
}

if agi, ok := snapshot.Stats[model.AGI]; ok {
snapshot.Stats["EVASION"] = agi * 0.5
snapshot.Stats["ATK"] += agi * 0.3
}

if spd, ok := snapshot.Stats[model.SPD]; ok {
snapshot.Stats["MOVE_SPEED"] = 100.0 + spd*2.0
}

if luk, ok := snapshot.Stats[model.LUK]; ok {
snapshot.Stats["CRIT_CHANCE"] = 0.05 + luk*0.01
snapshot.Stats["CRIT_DAMAGE"] = 1.5 + luk*0.02
}

// Apply modifiers
for _, mod := range input.Items {
if current, exists := snapshot.Stats[mod.Key]; exists {
switch mod.Op {
case model.ADD_FLAT:
snapshot.Stats[mod.Key] = current + mod.Value
case model.ADD_PCT:
snapshot.Stats[mod.Key] = current * (1.0 + mod.Value/100.0)
case model.MULTIPLY:
snapshot.Stats[mod.Key] = current * mod.Value
case model.OVERRIDE:
snapshot.Stats[mod.Key] = mod.Value
}
}
}

return snapshot
}
