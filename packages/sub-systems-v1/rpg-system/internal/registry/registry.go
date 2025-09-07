package registry

import (
	"rpg-system/internal/model"
)

// StatRegistry manages all stat definitions and formulas
type StatRegistry struct {
	statDefinitions map[model.StatKey]*model.StatDef
	levelCurves     map[model.StatKey]*LevelCurve
	derivedFormulas map[model.StatKey]*DerivedFormula
}

// LevelCurve defines how a stat scales with level
type LevelCurve struct {
	BaseValue    float64
	PerLevel     float64
	MaxLevel     int64
	SoftCapLevel int64
	SoftCapValue float64
}

// DerivedFormula defines how a derived stat is calculated
type DerivedFormula struct {
	StatKey    model.StatKey
	Formula    string
	Components []model.StatKey
	BaseValue  float64
}

// NewStatRegistry creates a new stat registry with default definitions
func NewStatRegistry() *StatRegistry {
	registry := &StatRegistry{
		statDefinitions: make(map[model.StatKey]*model.StatDef),
		levelCurves:     make(map[model.StatKey]*LevelCurve),
		derivedFormulas: make(map[model.StatKey]*DerivedFormula),
	}

	registry.initializePrimaryStats()
	registry.initializeDerivedStats()
	registry.initializeLevelCurves()
	registry.initializeDerivedFormulas()

	return registry
}

// GetStatDefinition returns the definition for a stat
func (r *StatRegistry) GetStatDefinition(statKey model.StatKey) (*model.StatDef, bool) {
	def, exists := r.statDefinitions[statKey]
	return def, exists
}

// GetLevelCurve returns the level curve for a stat
func (r *StatRegistry) GetLevelCurve(statKey model.StatKey) (*LevelCurve, bool) {
	curve, exists := r.levelCurves[statKey]
	return curve, exists
}

// GetDerivedFormula returns the formula for a derived stat
func (r *StatRegistry) GetDerivedFormula(statKey model.StatKey) (*DerivedFormula, bool) {
	formula, exists := r.derivedFormulas[statKey]
	return formula, exists
}

// GetAllPrimaryStats returns all primary stat definitions
func (r *StatRegistry) GetAllPrimaryStats() []*model.StatDef {
	var primaryStats []*model.StatDef
	for _, statKey := range model.PrimaryStats() {
		if def, exists := r.statDefinitions[statKey]; exists {
			primaryStats = append(primaryStats, def)
		}
	}
	return primaryStats
}

// GetAllDerivedStats returns all derived stat definitions
func (r *StatRegistry) GetAllDerivedStats() []*model.StatDef {
	var derivedStats []*model.StatDef
	for statKey, def := range r.statDefinitions {
		if !statKey.IsPrimary() {
			derivedStats = append(derivedStats, def)
		}
	}
	return derivedStats
}

// CalculateLevelValue calculates the value of a stat at a given level
func (r *StatRegistry) CalculateLevelValue(statKey model.StatKey, level int64) float64 {
	curve, exists := r.GetLevelCurve(statKey)
	if !exists {
		return 0
	}

	// Apply level scaling
	value := curve.BaseValue + float64(level-1)*curve.PerLevel

	// Apply soft cap if applicable
	if curve.SoftCapLevel > 0 && level > curve.SoftCapLevel {
		softCapValue := curve.BaseValue + float64(curve.SoftCapLevel-1)*curve.PerLevel
		excessLevels := level - curve.SoftCapLevel
		value = softCapValue + float64(excessLevels)*curve.SoftCapValue
	}

	// Apply max level cap
	if curve.MaxLevel > 0 && level > curve.MaxLevel {
		maxValue := curve.BaseValue + float64(curve.MaxLevel-1)*curve.PerLevel
		return maxValue
	}

	return value
}

// initializePrimaryStats sets up the 8 primary stats
func (r *StatRegistry) initializePrimaryStats() {
	primaryStats := []struct {
		key          model.StatKey
		name         string
		description  string
		minValue     float64
		maxValue     float64
		defaultValue float64
	}{
		{model.STR, "Strength", "Physical power and melee damage", 1, 100, 15},
		{model.INT, "Intelligence", "Magical power and mana", 1, 100, 12},
		{model.WIL, "Willpower", "Mental fortitude and spell resistance", 1, 100, 14},
		{model.AGI, "Agility", "Speed and evasion", 1, 100, 13},
		{model.SPD, "Speed", "Movement and action speed", 1, 100, 16},
		{model.END, "Endurance", "Health and stamina", 1, 100, 16},
		{model.PER, "Personality", "Social influence and merchant prices", 1, 100, 11},
		{model.LUK, "Luck", "Critical hits and random events", 1, 100, 10},
	}

	for _, stat := range primaryStats {
		r.statDefinitions[stat.key] = &model.StatDef{
			Key:          stat.key,
			DisplayName:  stat.name,
			Description:  stat.description,
			MinValue:     stat.minValue,
			MaxValue:     stat.maxValue,
			DefaultValue: stat.defaultValue,
			IsPrimary:    true,
		}
	}
}

// initializeDerivedStats sets up derived stats
func (r *StatRegistry) initializeDerivedStats() {
	derivedStats := []struct {
		key          model.StatKey
		name         string
		description  string
		minValue     float64
		maxValue     float64
		defaultValue float64
	}{
		{model.HP_MAX, "Health Points", "Maximum health points", 1, 10000, 100},
		{model.MANA_MAX, "Mana Points", "Maximum mana points", 1, 10000, 50},
		{model.ATK, "Attack Power", "Physical attack damage", 1, 1000, 20},
		{model.MATK, "Magic Attack", "Magical attack damage", 1, 1000, 15},
		{model.DEF, "Defense", "Physical damage reduction", 0, 1000, 10},
		{model.EVASION, "Evasion", "Chance to avoid attacks", 0, 100, 5},
		{model.MOVE_SPEED, "Movement Speed", "Movement speed multiplier", 0.1, 10, 1.0},
		{model.CRIT_CHANCE, "Critical Chance", "Chance for critical hits", 0, 1, 0.01},
		{model.CRIT_DAMAGE, "Critical Damage", "Critical hit damage multiplier", 1, 10, 2.0},
	}

	for _, stat := range derivedStats {
		r.statDefinitions[stat.key] = &model.StatDef{
			Key:          stat.key,
			DisplayName:  stat.name,
			Description:  stat.description,
			MinValue:     stat.minValue,
			MaxValue:     stat.maxValue,
			DefaultValue: stat.defaultValue,
			IsPrimary:    false,
		}
	}
}

// initializeLevelCurves sets up level scaling for primary stats
func (r *StatRegistry) initializeLevelCurves() {
	// Primary stats scale with level
	primaryStats := model.PrimaryStats()
	for _, statKey := range primaryStats {
		r.levelCurves[statKey] = &LevelCurve{
			BaseValue:    10,  // Base value at level 1
			PerLevel:     1,   // +1 per level
			MaxLevel:     100, // Max level
			SoftCapLevel: 50,  // Soft cap starts at level 50
			SoftCapValue: 0.5, // +0.5 per level after soft cap
		}
	}
}

// initializeDerivedFormulas sets up formulas for derived stats
func (r *StatRegistry) initializeDerivedFormulas() {
	// HP = Base + STR*10 + END*5
	r.derivedFormulas[model.HP_MAX] = &DerivedFormula{
		StatKey:    model.HP_MAX,
		Formula:    "Base + STR*10 + END*5",
		Components: []model.StatKey{model.STR, model.END},
		BaseValue:  100,
	}

	// MANA = Base + INT*15 + WIL*5
	r.derivedFormulas[model.MANA_MAX] = &DerivedFormula{
		StatKey:    model.MANA_MAX,
		Formula:    "Base + INT*15 + WIL*5",
		Components: []model.StatKey{model.INT, model.WIL},
		BaseValue:  50,
	}

	// ATK = STR*2 + AGI*0.3
	r.derivedFormulas[model.ATK] = &DerivedFormula{
		StatKey:    model.ATK,
		Formula:    "STR*2 + AGI*0.3",
		Components: []model.StatKey{model.STR, model.AGI},
		BaseValue:  0,
	}

	// MATK = INT*2 + WIL*0.4
	r.derivedFormulas[model.MATK] = &DerivedFormula{
		StatKey:    model.MATK,
		Formula:    "INT*2 + WIL*0.4",
		Components: []model.StatKey{model.INT, model.WIL},
		BaseValue:  0,
	}

	// DEF = END*1.5 + STR*0.5
	r.derivedFormulas[model.DEF] = &DerivedFormula{
		StatKey:    model.DEF,
		Formula:    "END*1.5 + STR*0.5",
		Components: []model.StatKey{model.END, model.STR},
		BaseValue:  0,
	}

	// EVASION = AGI*0.8 + SPD*0.2
	r.derivedFormulas[model.EVASION] = &DerivedFormula{
		StatKey:    model.EVASION,
		Formula:    "AGI*0.8 + SPD*0.2",
		Components: []model.StatKey{model.AGI, model.SPD},
		BaseValue:  0,
	}

	// MOVE_SPEED = Base + SPD*0.1
	r.derivedFormulas[model.MOVE_SPEED] = &DerivedFormula{
		StatKey:    model.MOVE_SPEED,
		Formula:    "Base + SPD*0.1",
		Components: []model.StatKey{model.SPD},
		BaseValue:  1.0,
	}

	// CRIT_CHANCE = Base + LUK*0.003 + AGI*0.001
	r.derivedFormulas[model.CRIT_CHANCE] = &DerivedFormula{
		StatKey:    model.CRIT_CHANCE,
		Formula:    "Base + LUK*0.003 + AGI*0.001",
		Components: []model.StatKey{model.LUK, model.AGI},
		BaseValue:  0.01,
	}

	// CRIT_DAMAGE = Base + LUK*0.02
	r.derivedFormulas[model.CRIT_DAMAGE] = &DerivedFormula{
		StatKey:    model.CRIT_DAMAGE,
		Formula:    "Base + LUK*0.02",
		Components: []model.StatKey{model.LUK},
		BaseValue:  2.0,
	}
}
