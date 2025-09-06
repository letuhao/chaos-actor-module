package rules

import (
"math"
"sort"

"rpg-stats/internal/model"
)

type StackingEngine struct{}

func NewStackingEngine() *StackingEngine {
return &StackingEngine{}
}

func (se *StackingEngine) ApplyModifiers(baseValue float64, modifiers []model.StatModifier, withBreakdown bool) (float64, *model.StatBreakdown) {
if len(modifiers) == 0 {
return baseValue, nil
}

sortedModifiers := se.sortModifiers(modifiers)

var breakdown *model.StatBreakdown
if withBreakdown {
breakdown = &model.StatBreakdown{
Base: baseValue,
}
}

value := baseValue
value = se.applyAddFlat(value, sortedModifiers, breakdown)
value = se.applyAddPct(value, sortedModifiers, breakdown)
value = se.applyMultiply(value, sortedModifiers, breakdown)
value = se.applyOverrides(value, sortedModifiers, breakdown)
value = se.applyCaps(value, sortedModifiers, breakdown)
value = se.applyRounding(value, sortedModifiers, breakdown)

return value, breakdown
}

func (se *StackingEngine) sortModifiers(modifiers []model.StatModifier) []model.StatModifier {
sorted := make([]model.StatModifier, len(modifiers))
copy(sorted, modifiers)

sort.Slice(sorted, func(i, j int) bool {
if sorted[i].Op != sorted[j].Op {
return se.getOperationOrder(sorted[i].Op) < se.getOperationOrder(sorted[j].Op)
}

if sorted[i].Priority != sorted[j].Priority {
return sorted[i].Priority > sorted[j].Priority
}

if sorted[i].Op == model.OVERRIDE {
return sorted[i].Value > sorted[j].Value
}

return false
})

return sorted
}

func (se *StackingEngine) getOperationOrder(op model.ModifierStacking) int {
switch op {
case model.ADD_FLAT:
return 1
case model.ADD_PCT:
return 2
case model.MULTIPLY:
return 3
case model.OVERRIDE:
return 4
default:
return 5
}
}

func (se *StackingEngine) applyAddFlat(value float64, modifiers []model.StatModifier, breakdown *model.StatBreakdown) float64 {
var total float64
for _, mod := range modifiers {
if mod.Op == model.ADD_FLAT {
total += mod.Value
}
}

if breakdown != nil {
breakdown.AdditiveFlat = total
}

return value + total
}

func (se *StackingEngine) applyAddPct(value float64, modifiers []model.StatModifier, breakdown *model.StatBreakdown) float64 {
var totalPct float64
for _, mod := range modifiers {
if mod.Op == model.ADD_PCT {
totalPct += mod.Value
}
}

if breakdown != nil {
breakdown.AdditivePct = totalPct
}

return value * (1 + totalPct/100)
}

func (se *StackingEngine) applyMultiply(value float64, modifiers []model.StatModifier, breakdown *model.StatBreakdown) float64 {
var totalMult float64 = 1.0
for _, mod := range modifiers {
if mod.Op == model.MULTIPLY {
totalMult *= mod.Value
}
}

if breakdown != nil {
breakdown.Multiplicative = totalMult
}

return value * totalMult
}

func (se *StackingEngine) applyOverrides(value float64, modifiers []model.StatModifier, breakdown *model.StatBreakdown) float64 {
var overrides []model.OverrideEntry
var highestValue float64

for _, mod := range modifiers {
if mod.Op == model.OVERRIDE {
overrides = append(overrides, model.OverrideEntry{
Value:  mod.Value,
Source: mod.Source,
})

if mod.Value > highestValue {
highestValue = mod.Value
}
}
}

if len(overrides) > 0 {
if breakdown != nil {
breakdown.Overrides = overrides
}
return highestValue
}

return value
}

func (se *StackingEngine) applyCaps(value float64, modifiers []model.StatModifier, breakdown *model.StatBreakdown) float64 {
if value > 0.75 && se.isResistanceStat(modifiers) {
if breakdown != nil {
capped := 0.75
breakdown.CappedTo = &capped
}
return 0.75
}

return value
}

func (se *StackingEngine) applyRounding(value float64, modifiers []model.StatModifier, breakdown *model.StatBreakdown) float64 {
return math.Round(value*100) / 100
}

func (se *StackingEngine) isResistanceStat(modifiers []model.StatModifier) bool {
if len(modifiers) == 0 {
return false
}

key := modifiers[0].Key
return key == "FIRE_RES" || key == "ICE_RES" || key == "LIGHTNING_RES" || key == "POISON_RES"
}

func (se *StackingEngine) GroupModifiersByStack(modifiers []model.StatModifier) map[string][]model.StatModifier {
groups := make(map[string][]model.StatModifier)

for _, mod := range modifiers {
if mod.Conditions != nil && mod.Conditions.StackID != "" {
groups[mod.Conditions.StackID] = append(groups[mod.Conditions.StackID], mod)
} else {
key := string(mod.Key) + "_" + string(mod.Op)
groups[key] = append(groups[key], mod)
}
}

return groups
}

func (se *StackingEngine) ApplyStackLimits(groups map[string][]model.StatModifier) []model.StatModifier {
var result []model.StatModifier

for _, group := range groups {
if len(group) == 0 {
continue
}

maxStacks := -1
if group[0].Conditions != nil {
maxStacks = group[0].Conditions.MaxStacks
}

if maxStacks > 0 && len(group) > maxStacks {
sort.Slice(group, func(i, j int) bool {
if group[i].Priority != group[j].Priority {
return group[i].Priority > group[j].Priority
}
return group[i].Value > group[j].Value
})
group = group[:maxStacks]
}

result = append(result, group...)
}

return result
}
