package test

import (
	"testing"

	"rpg-system/internal/model"
	"rpg-system/internal/registry"
)

func TestStatRegistry_GetStatDefinition(t *testing.T) {
	reg := registry.NewStatRegistry()

	tests := []struct {
		statKey     model.StatKey
		shouldExist bool
		isPrimary   bool
	}{
		{model.STR, true, true},
		{model.INT, true, true},
		{model.WIL, true, true},
		{model.AGI, true, true},
		{model.SPD, true, true},
		{model.END, true, true},
		{model.PER, true, true},
		{model.LUK, true, true},
		{model.HP_MAX, true, false},
		{model.MANA_MAX, true, false},
		{model.ATK, true, false},
		{model.MATK, true, false},
		{model.DEF, true, false},
		{model.EVASION, true, false},
		{model.MOVE_SPEED, true, false},
		{model.CRIT_CHANCE, true, false},
		{model.CRIT_DAMAGE, true, false},
		{model.StatKey("NONEXISTENT"), false, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.statKey), func(t *testing.T) {
			def, exists := reg.GetStatDefinition(tt.statKey)

			if tt.shouldExist {
				if !exists {
					t.Errorf("Expected stat definition to exist for %s", tt.statKey)
				} else {
					if def.IsPrimary != tt.isPrimary {
						t.Errorf("Expected IsPrimary=%v for %s, got %v", tt.isPrimary, tt.statKey, def.IsPrimary)
					}
					if def.Key != tt.statKey {
						t.Errorf("Expected Key=%s for %s, got %s", tt.statKey, tt.statKey, def.Key)
					}
				}
			} else {
				if exists {
					t.Errorf("Expected stat definition to not exist for %s", tt.statKey)
				}
			}
		})
	}
}

func TestStatRegistry_GetAllPrimaryStats(t *testing.T) {
	reg := registry.NewStatRegistry()
	primaryStats := reg.GetAllPrimaryStats()

	expectedCount := len(model.PrimaryStats())
	if len(primaryStats) != expectedCount {
		t.Errorf("Expected %d primary stats, got %d", expectedCount, len(primaryStats))
	}

	// Check that all returned stats are primary
	for _, stat := range primaryStats {
		if !stat.IsPrimary {
			t.Errorf("Expected all returned stats to be primary, but %s is not", stat.Key)
		}
	}
}

func TestStatRegistry_GetAllDerivedStats(t *testing.T) {
	reg := registry.NewStatRegistry()
	derivedStats := reg.GetAllDerivedStats()

	// Should have at least the known derived stats
	expectedDerivedStats := []model.StatKey{
		model.HP_MAX, model.MANA_MAX, model.ATK, model.MATK, model.DEF,
		model.EVASION, model.MOVE_SPEED, model.CRIT_CHANCE, model.CRIT_DAMAGE,
	}

	if len(derivedStats) < len(expectedDerivedStats) {
		t.Errorf("Expected at least %d derived stats, got %d", len(expectedDerivedStats), len(derivedStats))
	}

	// Check that all returned stats are derived
	for _, stat := range derivedStats {
		if stat.IsPrimary {
			t.Errorf("Expected all returned stats to be derived, but %s is primary", stat.Key)
		}
	}
}

func TestStatRegistry_CalculateLevelValue(t *testing.T) {
	reg := registry.NewStatRegistry()

	tests := []struct {
		statKey     model.StatKey
		level       int64
		expectedMin float64
		expectedMax float64
	}{
		{model.STR, 1, 9, 11},    // Base value around 10
		{model.STR, 5, 13, 15},   // Level 5
		{model.STR, 10, 18, 20},  // Level 10
		{model.STR, 50, 58, 62},  // Level 50 (soft cap)
		{model.STR, 100, 82, 86}, // Level 100 (max level)
	}

	for _, tt := range tests {
		t.Run(string(tt.statKey)+"_Level_"+string(rune(tt.level)), func(t *testing.T) {
			value := reg.CalculateLevelValue(tt.statKey, tt.level)

			if value < tt.expectedMin || value > tt.expectedMax {
				t.Errorf("Expected value between %f and %f for %s level %d, got %f",
					tt.expectedMin, tt.expectedMax, tt.statKey, tt.level, value)
			}
		})
	}
}

func TestStatRegistry_GetDerivedFormula(t *testing.T) {
	reg := registry.NewStatRegistry()

	tests := []struct {
		statKey      model.StatKey
		shouldExist  bool
		expectedBase float64
	}{
		{model.HP_MAX, true, 100},
		{model.MANA_MAX, true, 50},
		{model.ATK, true, 0},
		{model.MATK, true, 0},
		{model.DEF, true, 0},
		{model.EVASION, true, 0},
		{model.MOVE_SPEED, true, 1.0},
		{model.CRIT_CHANCE, true, 0.01},
		{model.CRIT_DAMAGE, true, 2.0},
		{model.STR, false, 0}, // Primary stat, no formula
		{model.StatKey("NONEXISTENT"), false, 0},
	}

	for _, tt := range tests {
		t.Run(string(tt.statKey), func(t *testing.T) {
			formula, exists := reg.GetDerivedFormula(tt.statKey)

			if tt.shouldExist {
				if !exists {
					t.Errorf("Expected derived formula to exist for %s", tt.statKey)
				} else {
					if formula.BaseValue != tt.expectedBase {
						t.Errorf("Expected BaseValue=%f for %s, got %f", tt.expectedBase, tt.statKey, formula.BaseValue)
					}
					if formula.StatKey != tt.statKey {
						t.Errorf("Expected StatKey=%s for %s, got %s", tt.statKey, tt.statKey, formula.StatKey)
					}
				}
			} else {
				if exists {
					t.Errorf("Expected derived formula to not exist for %s", tt.statKey)
				}
			}
		})
	}
}

func TestFormulaCalculator_CalculateAllDerivedStats(t *testing.T) {
	reg := registry.NewStatRegistry()
	calc := registry.NewFormulaCalculator(reg)

	// Test with sample primary stats
	primaryStats := map[model.StatKey]float64{
		model.STR: 15,
		model.INT: 12,
		model.WIL: 10,
		model.AGI: 14,
		model.SPD: 13,
		model.END: 16,
		model.PER: 8,
		model.LUK: 11,
	}

	derivedStats := calc.CalculateAllDerivedStats(primaryStats)

	// Check that derived stats are calculated
	expectedDerivedStats := []model.StatKey{
		model.HP_MAX, model.MANA_MAX, model.ATK, model.MATK, model.DEF,
		model.EVASION, model.MOVE_SPEED, model.CRIT_CHANCE, model.CRIT_DAMAGE,
	}

	for _, statKey := range expectedDerivedStats {
		if value, exists := derivedStats[statKey]; exists {
			if value <= 0 {
				t.Errorf("Expected positive value for %s, got %f", statKey, value)
			}
		} else {
			t.Errorf("Expected %s to be calculated", statKey)
		}
	}

	// Test specific calculations
	// HP = 100 + STR*10 + END*5 = 100 + 15*10 + 16*5 = 100 + 150 + 80 = 330
	expectedHP := 100.0 + 15.0*10.0 + 16.0*5.0
	if hp, exists := derivedStats[model.HP_MAX]; exists {
		if hp < expectedHP-1 || hp > expectedHP+1 {
			t.Errorf("Expected HP around %f, got %f", expectedHP, hp)
		}
	}

	// ATK = STR*2 + AGI*0.3 = 15*2 + 14*0.3 = 30 + 4.2 = 34.2
	expectedATK := 15.0*2.0 + 14.0*0.3
	if atk, exists := derivedStats[model.ATK]; exists {
		if atk < expectedATK-1 || atk > expectedATK+1 {
			t.Errorf("Expected ATK around %f, got %f", expectedATK, atk)
		}
	}
}

func TestFormulaCalculator_ApplyStatCaps(t *testing.T) {
	reg := registry.NewStatRegistry()
	calc := registry.NewFormulaCalculator(reg)

	tests := []struct {
		statKey     model.StatKey
		value       float64
		expectedMin float64
		expectedMax float64
	}{
		{model.STR, 50.0, 49.0, 51.0},         // Should not be capped
		{model.HP_MAX, 1000.0, 999.0, 1001.0}, // Should not be capped
		{model.CRIT_CHANCE, 0.8, 0.75, 0.75},  // Should be capped at 75%
		{model.CRIT_CHANCE, 0.5, 0.49, 0.51},  // Should not be capped
	}

	for _, tt := range tests {
		t.Run(string(tt.statKey)+"_"+string(rune(tt.value)), func(t *testing.T) {
			cappedValue := calc.ApplyStatCaps(tt.statKey, tt.value)

			if cappedValue < tt.expectedMin || cappedValue > tt.expectedMax {
				t.Errorf("Expected capped value between %f and %f for %s value %f, got %f",
					tt.expectedMin, tt.expectedMax, tt.statKey, tt.value, cappedValue)
			}
		})
	}
}
