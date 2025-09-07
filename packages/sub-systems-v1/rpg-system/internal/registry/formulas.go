package registry

import (
	"rpg-system/internal/model"
)

// FormulaCalculator handles calculation of derived stats
type FormulaCalculator struct {
	registry *StatRegistry
}

// NewFormulaCalculator creates a new formula calculator
func NewFormulaCalculator(registry *StatRegistry) *FormulaCalculator {
	return &FormulaCalculator{
		registry: registry,
	}
}

// CalculateDerivedStat calculates a derived stat value based on primary stats
func (fc *FormulaCalculator) CalculateDerivedStat(statKey model.StatKey, primaryStats map[model.StatKey]float64) float64 {
	formula, exists := fc.registry.GetDerivedFormula(statKey)
	if !exists {
		return 0
	}

	switch statKey {
	case model.HP_MAX:
		return fc.calculateHP(primaryStats, formula)
	case model.MANA_MAX:
		return fc.calculateMana(primaryStats, formula)
	case model.ATK:
		return fc.calculateATK(primaryStats, formula)
	case model.MATK:
		return fc.calculateMATK(primaryStats, formula)
	case model.DEF:
		return fc.calculateDEF(primaryStats, formula)
	case model.EVASION:
		return fc.calculateEvasion(primaryStats, formula)
	case model.MOVE_SPEED:
		return fc.calculateMoveSpeed(primaryStats, formula)
	case model.CRIT_CHANCE:
		return fc.calculateCritChance(primaryStats, formula)
	case model.CRIT_DAMAGE:
		return fc.calculateCritDamage(primaryStats, formula)
	default:
		return 0
	}
}

// calculateHP calculates HP_MAX = Base + STR*10 + END*5
func (fc *FormulaCalculator) calculateHP(primaryStats map[model.StatKey]float64, formula *DerivedFormula) float64 {
	str := primaryStats[model.STR]
	end := primaryStats[model.END]
	return formula.BaseValue + str*10 + end*5
}

// calculateMana calculates MANA_MAX = Base + INT*15 + WIL*5
func (fc *FormulaCalculator) calculateMana(primaryStats map[model.StatKey]float64, formula *DerivedFormula) float64 {
	intel := primaryStats[model.INT]
	wil := primaryStats[model.WIL]
	return formula.BaseValue + intel*15 + wil*5
}

// calculateATK calculates ATK = STR*2 + AGI*0.3
func (fc *FormulaCalculator) calculateATK(primaryStats map[model.StatKey]float64, formula *DerivedFormula) float64 {
	str := primaryStats[model.STR]
	agi := primaryStats[model.AGI]
	return formula.BaseValue + str*2 + agi*0.3
}

// calculateMATK calculates MATK = INT*2 + WIL*0.4
func (fc *FormulaCalculator) calculateMATK(primaryStats map[model.StatKey]float64, formula *DerivedFormula) float64 {
	intel := primaryStats[model.INT]
	wil := primaryStats[model.WIL]
	return formula.BaseValue + intel*2 + wil*0.4
}

// calculateDEF calculates DEF = END*1.5 + STR*0.5
func (fc *FormulaCalculator) calculateDEF(primaryStats map[model.StatKey]float64, formula *DerivedFormula) float64 {
	end := primaryStats[model.END]
	str := primaryStats[model.STR]
	return formula.BaseValue + end*1.5 + str*0.5
}

// calculateEvasion calculates EVASION = AGI*0.8 + SPD*0.2
func (fc *FormulaCalculator) calculateEvasion(primaryStats map[model.StatKey]float64, formula *DerivedFormula) float64 {
	agi := primaryStats[model.AGI]
	spd := primaryStats[model.SPD]
	return formula.BaseValue + agi*0.8 + spd*0.2
}

// calculateMoveSpeed calculates MOVE_SPEED = Base + SPD*0.1
func (fc *FormulaCalculator) calculateMoveSpeed(primaryStats map[model.StatKey]float64, formula *DerivedFormula) float64 {
	spd := primaryStats[model.SPD]
	return formula.BaseValue + spd*0.1
}

// calculateCritChance calculates CRIT_CHANCE = Base + LUK*0.003 + AGI*0.001
func (fc *FormulaCalculator) calculateCritChance(primaryStats map[model.StatKey]float64, formula *DerivedFormula) float64 {
	luk := primaryStats[model.LUK]
	agi := primaryStats[model.AGI]
	chance := formula.BaseValue + luk*0.003 + agi*0.001

	// Cap at 100% (1.0)
	if chance > 1.0 {
		chance = 1.0
	}

	return chance
}

// calculateCritDamage calculates CRIT_DAMAGE = Base + LUK*0.02
func (fc *FormulaCalculator) calculateCritDamage(primaryStats map[model.StatKey]float64, formula *DerivedFormula) float64 {
	luk := primaryStats[model.LUK]
	return formula.BaseValue + luk*0.02
}

// CalculateAllDerivedStats calculates all derived stats for given primary stats
func (fc *FormulaCalculator) CalculateAllDerivedStats(primaryStats map[model.StatKey]float64) map[model.StatKey]float64 {
	derivedStats := make(map[model.StatKey]float64)

	// Calculate each derived stat
	derivedStatKeys := []model.StatKey{
		model.HP_MAX, model.MANA_MAX, model.ATK, model.MATK, model.DEF,
		model.EVASION, model.MOVE_SPEED, model.CRIT_CHANCE, model.CRIT_DAMAGE,
	}

	for _, statKey := range derivedStatKeys {
		derivedStats[statKey] = fc.CalculateDerivedStat(statKey, primaryStats)
	}

	return derivedStats
}

// ValidateStatValue checks if a stat value is within valid bounds
func (fc *FormulaCalculator) ValidateStatValue(statKey model.StatKey, value float64) bool {
	def, exists := fc.registry.GetStatDefinition(statKey)
	if !exists {
		return false
	}

	return value >= def.MinValue && value <= def.MaxValue
}

// ApplyStatCaps applies min/max caps to a stat value
func (fc *FormulaCalculator) ApplyStatCaps(statKey model.StatKey, value float64) float64 {
	def, exists := fc.registry.GetStatDefinition(statKey)
	if !exists {
		return value
	}

	// Apply special caps for certain stats
	switch statKey {
	case model.CRIT_CHANCE:
		// Cap critical chance at 75%
		if value > 0.75 {
			value = 0.75
		}
	}

	// Apply general min/max caps
	if value < def.MinValue {
		return def.MinValue
	}
	if value > def.MaxValue {
		return def.MaxValue
	}

	return value
}
