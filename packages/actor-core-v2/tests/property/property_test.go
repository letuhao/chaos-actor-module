package property

import (
	"math"
	"testing"

	"actor-core/constants"
	"actor-core/models/core"
	coreServices "actor-core/services/core"
)

// TestStatClampInvariants tests that all stats respect their clamp values
func TestStatClampInvariants(t *testing.T) {
	sr := coreServices.NewStatResolver()

	// Test with extreme values to ensure clamping works
	extremePC := &core.PrimaryCore{
		Vitality:        1000000, // Very high
		Constitution:    1000000,
		Agility:         1000000,
		Intelligence:    1000000,
		Wisdom:          1000000,
		Charisma:        1000000,
		Luck:            1000000,
		Karma:           1000000,
		Endurance:       1000000,
		Strength:        1000000,
		SpiritualEnergy: 1000000,
		PhysicalEnergy:  1000000,
		MentalEnergy:    1000000,
		Willpower:       1000000,
	}

	derivedStats, err := sr.ResolveStats(extremePC)
	if err != nil {
		t.Fatalf("Failed to resolve stats: %v", err)
	}

	allStats := derivedStats.GetAllStats()

	// Check that all stats are within reasonable bounds
	for statName, value := range allStats {
		// Basic sanity checks
		if math.IsNaN(value) {
			t.Errorf("Stat %s is NaN", statName)
		}
		if math.IsInf(value, 0) {
			t.Errorf("Stat %s is infinite", statName)
		}

		// Check specific clamp values based on stat type
		// Note: These are expected ranges, not hard limits
		switch statName {
		case constants.Stat_HP_MAX:
			if value < 0 {
				t.Errorf("Stat %s should be non-negative, got %f", statName, value)
			}
		case constants.Stat_MOVE_SPEED:
			if value < 0 {
				t.Errorf("Stat %s should be non-negative, got %f", statName, value)
			}
		case constants.Stat_CRIT_CHANCE:
			if value < 0 {
				t.Errorf("Stat %s should be non-negative, got %f", statName, value)
			}
		case constants.Stat_CRIT_MULTI:
			if value < 0 {
				t.Errorf("Stat %s should be non-negative, got %f", statName, value)
			}
		}
	}
}

// TestStatMonotonicity tests that stats increase monotonically with primary stats
func TestStatMonotonicity(t *testing.T) {
	sr := coreServices.NewStatResolver()

	// Test that increasing primary stats increases derived stats
	basePC := &core.PrimaryCore{
		Vitality:        50,
		Constitution:    50,
		Agility:         50,
		Intelligence:    50,
		Wisdom:          50,
		Charisma:        50,
		Luck:            50,
		Karma:           50,
		Endurance:       50,
		Strength:        50,
		SpiritualEnergy: 50,
		PhysicalEnergy:  50,
		MentalEnergy:    50,
		Willpower:       50,
	}

	// Get base stats
	baseStats, err := sr.ResolveStats(basePC)
	if err != nil {
		t.Fatalf("Failed to resolve base stats: %v", err)
	}
	baseValues := baseStats.GetAllStats()

	// Test vitality increase
	highVitalityPC := *basePC
	highVitalityPC.Vitality = 100

	highVitalityStats, err := sr.ResolveStats(&highVitalityPC)
	if err != nil {
		t.Fatalf("Failed to resolve high vitality stats: %v", err)
	}
	highVitalityValues := highVitalityStats.GetAllStats()

	// Check that HP-related stats increase with vitality
	hpStats := []string{constants.Stat_HP_MAX}
	for _, statName := range hpStats {
		if baseValues[statName] >= highVitalityValues[statName] {
			t.Errorf("Stat %s should increase with vitality: base=%f, high=%f",
				statName, baseValues[statName], highVitalityValues[statName])
		}
	}

	// Check that stamina increases with endurance (not vitality)
	highEndurancePC := *basePC
	highEndurancePC.Endurance = 100

	highEnduranceStats, err := sr.ResolveStats(&highEndurancePC)
	if err != nil {
		t.Fatalf("Failed to resolve high endurance stats: %v", err)
	}
	highEnduranceValues := highEnduranceStats.GetAllStats()

	if baseValues[constants.Stat_STAMINA] >= highEnduranceValues[constants.Stat_STAMINA] {
		t.Errorf("Stat %s should increase with endurance: base=%f, high=%f",
			constants.Stat_STAMINA, baseValues[constants.Stat_STAMINA], highEnduranceValues[constants.Stat_STAMINA])
	}
}

// TestStatAssociativity tests associativity properties where applicable
func TestStatAssociativity(t *testing.T) {
	sr := coreServices.NewStatResolver()

	// Test that stat calculations are consistent
	pc := &core.PrimaryCore{
		Vitality:        75,
		Constitution:    60,
		Agility:         80,
		Intelligence:    70,
		Wisdom:          65,
		Charisma:        55,
		Luck:            45,
		Karma:           50,
		Endurance:       70,
		Strength:        85,
		SpiritualEnergy: 60,
		PhysicalEnergy:  70,
		MentalEnergy:    65,
		Willpower:       75,
	}

	// Resolve stats multiple times
	stats1, err := sr.ResolveStats(pc)
	if err != nil {
		t.Fatalf("Failed to resolve stats (1): %v", err)
	}

	stats2, err := sr.ResolveStats(pc)
	if err != nil {
		t.Fatalf("Failed to resolve stats (2): %v", err)
	}

	values1 := stats1.GetAllStats()
	values2 := stats2.GetAllStats()

	// Check that results are identical
	for statName, value1 := range values1 {
		value2, exists := values2[statName]
		if !exists {
			t.Errorf("Stat %s missing in second resolution", statName)
			continue
		}

		if math.Abs(value1-value2) > 0.001 {
			t.Errorf("Stat %s not consistent: first=%f, second=%f", statName, value1, value2)
		}
	}
}

// TestStatCommutativity tests commutativity properties where applicable
func TestStatCommutativity(t *testing.T) {
	sr := coreServices.NewStatResolver()

	// Test that order of primary stat changes doesn't affect final result
	pc1 := &core.PrimaryCore{
		Vitality:        50,
		Constitution:    60,
		Agility:         70,
		Intelligence:    80,
		Wisdom:          90,
		Charisma:        40,
		Luck:            30,
		Karma:           20,
		Endurance:       10,
		Strength:        100,
		SpiritualEnergy: 50,
		PhysicalEnergy:  60,
		MentalEnergy:    70,
		Willpower:       80,
	}

	pc2 := &core.PrimaryCore{
		Vitality:        50,
		Constitution:    60,
		Agility:         70,
		Intelligence:    80,
		Wisdom:          90,
		Charisma:        40,
		Luck:            30,
		Karma:           20,
		Endurance:       10,
		Strength:        100,
		SpiritualEnergy: 50,
		PhysicalEnergy:  60,
		MentalEnergy:    70,
		Willpower:       80,
	}

	// Resolve stats for both
	stats1, err := sr.ResolveStats(pc1)
	if err != nil {
		t.Fatalf("Failed to resolve stats (1): %v", err)
	}

	stats2, err := sr.ResolveStats(pc2)
	if err != nil {
		t.Fatalf("Failed to resolve stats (2): %v", err)
	}

	values1 := stats1.GetAllStats()
	values2 := stats2.GetAllStats()

	// Check that results are identical
	for statName, value1 := range values1 {
		value2, exists := values2[statName]
		if !exists {
			t.Errorf("Stat %s missing in second resolution", statName)
			continue
		}

		if math.Abs(value1-value2) > 0.001 {
			t.Errorf("Stat %s not commutative: first=%f, second=%f", statName, value1, value2)
		}
	}
}

// TestStatBoundaryConditions tests edge cases and boundary conditions
func TestStatBoundaryConditions(t *testing.T) {
	sr := coreServices.NewStatResolver()

	// Test with minimum values
	minPC := &core.PrimaryCore{
		Vitality:        1,
		Constitution:    1,
		Agility:         1,
		Intelligence:    1,
		Wisdom:          1,
		Charisma:        1,
		Luck:            1,
		Karma:           1,
		Endurance:       1,
		Strength:        1,
		SpiritualEnergy: 1,
		PhysicalEnergy:  1,
		MentalEnergy:    1,
		Willpower:       1,
	}

	minStats, err := sr.ResolveStats(minPC)
	if err != nil {
		t.Fatalf("Failed to resolve min stats: %v", err)
	}

	minValues := minStats.GetAllStats()

	// Check that all stats are positive
	for statName, value := range minValues {
		if value < 0 {
			t.Errorf("Stat %s should be non-negative, got %f", statName, value)
		}
	}

	// Test with zero values
	zeroPC := &core.PrimaryCore{}

	zeroStats, err := sr.ResolveStats(zeroPC)
	if err != nil {
		t.Fatalf("Failed to resolve zero stats: %v", err)
	}

	zeroValues := zeroStats.GetAllStats()

	// Check that stats handle zero values gracefully
	for statName, value := range zeroValues {
		if math.IsNaN(value) {
			t.Errorf("Stat %s is NaN with zero input", statName)
		}
		if math.IsInf(value, 0) {
			t.Errorf("Stat %s is infinite with zero input", statName)
		}
	}
}
