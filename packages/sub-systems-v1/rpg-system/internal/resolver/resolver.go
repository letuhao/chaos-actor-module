package resolver

import (
	"rpg-system/internal/model"
	"rpg-system/internal/registry"
	"rpg-system/internal/rules"
	"rpg-system/internal/util"
	"time"
)

type StatResolver struct {
	registry          *registry.StatRegistry
	formulaCalculator *registry.FormulaCalculator
	stackingEngine    *rules.StackingEngine
}

func NewStatResolver() *StatResolver {
	reg := registry.NewStatRegistry()
	formulaCalc := registry.NewFormulaCalculator(reg)
	stackingEngine := rules.NewStackingEngine()

	return &StatResolver{
		registry:          reg,
		formulaCalculator: formulaCalc,
		stackingEngine:    stackingEngine,
	}
}

func (sr *StatResolver) ComputeSnapshot(input model.ComputeInput) *model.StatSnapshot {
	snapshot := &model.StatSnapshot{
		ActorID:   input.ActorID,
		Stats:     make(map[model.StatKey]float64),
		Breakdown: make(map[model.StatKey]*model.StatBreakdown),
		Version:   1,
		Ts:        time.Now().Unix(),
	}

	// Calculate primary stats from base allocations
	primaryStats := make(map[model.StatKey]float64)
	for key, value := range input.BaseAllocations {
		primaryStats[key] = float64(value)
		snapshot.Stats[key] = float64(value)
	}

	// Calculate derived stats using the registry formulas
	derivedStats := sr.formulaCalculator.CalculateAllDerivedStats(primaryStats)
	for key, value := range derivedStats {
		snapshot.Stats[key] = value
	}

	// Collect all modifiers from all sources
	allModifiers := sr.CollectAllModifiers(input)

	// Apply modifiers using the stacking engine
	for statKey, baseValue := range snapshot.Stats {
		statModifiers := sr.FilterModifiersForStat(allModifiers, statKey)
		if len(statModifiers) > 0 {
			finalValue, breakdown := sr.stackingEngine.ApplyModifiers(baseValue, statModifiers, input.WithBreakdown)
			snapshot.Stats[statKey] = finalValue

			if input.WithBreakdown && breakdown != nil {
				snapshot.Breakdown[statKey] = breakdown
			}
		}
	}

	// Apply stat caps
	for statKey, value := range snapshot.Stats {
		cappedValue := sr.formulaCalculator.ApplyStatCaps(statKey, value)
		snapshot.Stats[statKey] = cappedValue
	}

	// Generate hash for caching
	hash, err := util.BuildHash(input)
	if err == nil {
		snapshot.Hash = hash
	}

	return snapshot
}

// CollectAllModifiers collects modifiers from all sources
func (sr *StatResolver) CollectAllModifiers(input model.ComputeInput) []model.StatModifier {
	var allModifiers []model.StatModifier

	// Add modifiers from all sources
	allModifiers = append(allModifiers, input.Items...)
	allModifiers = append(allModifiers, input.Titles...)
	allModifiers = append(allModifiers, input.Passives...)
	allModifiers = append(allModifiers, input.Buffs...)
	allModifiers = append(allModifiers, input.Debuffs...)
	allModifiers = append(allModifiers, input.Auras...)
	allModifiers = append(allModifiers, input.Environment...)

	return allModifiers
}

// FilterModifiersForStat returns only modifiers that affect the given stat
func (sr *StatResolver) FilterModifiersForStat(modifiers []model.StatModifier, statKey model.StatKey) []model.StatModifier {
	var filtered []model.StatModifier
	for _, mod := range modifiers {
		if mod.Key == statKey {
			filtered = append(filtered, mod)
		}
	}
	return filtered
}
