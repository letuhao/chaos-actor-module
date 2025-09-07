package unit

import (
	"context"
	"testing"

	"actor-core-v2/models/core"
	coreService "actor-core-v2/services/core"
)

func TestNewStatResolver(t *testing.T) {
	sr := coreService.NewStatResolver()

	if sr == nil {
		t.Error("Expected StatResolver to be created")
	}

	if sr.GetStatsCount() == 0 {
		t.Error("Expected formulas to be initialized")
	}

	if sr.GetVersion() != 1 {
		t.Errorf("Expected version to be 1, got %d", sr.GetVersion())
	}
}

func TestResolveStats(t *testing.T) {
	sr := coreService.NewStatResolver()
	pc := core.NewPrimaryCore()

	// Set some test values
	pc.Vitality = 20
	pc.Constitution = 15
	pc.Agility = 25
	pc.Intelligence = 30
	pc.Charisma = 35
	pc.Luck = 40
	pc.SpiritualEnergy = 50
	pc.PhysicalEnergy = 60
	pc.MentalEnergy = 70
	pc.Strength = 80
	pc.Willpower = 90
	pc.Wisdom = 100

	derivedStats, err := sr.ResolveStats(pc)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if derivedStats == nil {
		t.Error("Expected derivedStats to be created")
	}

	// Test specific calculations
	expectedHPMax := float64(20*10 + 15*5) // Vitality * 10 + Constitution * 5
	if derivedStats.HPMax != expectedHPMax {
		t.Errorf("Expected HPMax to be %f, got %f", expectedHPMax, derivedStats.HPMax)
	}

	expectedStamina := float64(10*10 + 15*3) // Endurance * 10 + Constitution * 3
	if derivedStats.Stamina != expectedStamina {
		t.Errorf("Expected Stamina to be %f, got %f", expectedStamina, derivedStats.Stamina)
	}

	expectedSpeed := float64(25) * 0.1 // Agility * 0.1
	if derivedStats.Speed != expectedSpeed {
		t.Errorf("Expected Speed to be %f, got %f", expectedSpeed, derivedStats.Speed)
	}

	expectedHaste := 1.0 + float64(25)*0.01 // 1.0 + Agility * 0.01
	if derivedStats.Haste != expectedHaste {
		t.Errorf("Expected Haste to be %f, got %f", expectedHaste, derivedStats.Haste)
	}
}

func TestResolveStat(t *testing.T) {
	sr := coreService.NewStatResolver()
	pc := core.NewPrimaryCore()
	pc.Vitality = 20
	pc.Constitution = 15

	// Test valid stat
	value, err := sr.ResolveStat("hp_max", pc)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedValue := float64(20*10 + 15*5)
	if value != expectedValue {
		t.Errorf("Expected value to be %f, got %f", expectedValue, value)
	}

	// Test invalid stat
	_, err = sr.ResolveStat("invalid_stat", pc)
	if err == nil {
		t.Error("Expected error for invalid stat")
	}
}

func TestResolveStatsWithContext(t *testing.T) {
	sr := coreService.NewStatResolver()
	pc := core.NewPrimaryCore()

	ctx := context.Background()

	derivedStats, err := sr.ResolveStatsWithContext(ctx, pc)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if derivedStats == nil {
		t.Error("Expected derivedStats to be created")
	}
}

func TestResolveStatsWithCancelledContext(t *testing.T) {
	sr := coreService.NewStatResolver()
	pc := core.NewPrimaryCore()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := sr.ResolveStatsWithContext(ctx, pc)
	if err == nil {
		t.Error("Expected error for cancelled context")
	}
}

func TestCheckDependencies(t *testing.T) {
	sr := coreService.NewStatResolver()

	// Test valid stat
	deps, err := sr.CheckDependencies("hp_max")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedDeps := []string{"vitality", "constitution"}
	if len(deps) != len(expectedDeps) {
		t.Errorf("Expected %d dependencies, got %d", len(expectedDeps), len(deps))
	}

	// Test invalid stat
	_, err = sr.CheckDependencies("invalid_stat")
	if err == nil {
		t.Error("Expected error for invalid stat")
	}
}

func TestGetCalculationOrder(t *testing.T) {
	sr := coreService.NewStatResolver()

	order := sr.GetCalculationOrder()
	if len(order) == 0 {
		t.Error("Expected calculation order to be non-empty")
	}

	// Check that dependencies come before dependent stats
	hpMaxIndex := -1
	vitalityIndex := -1

	for i, stat := range order {
		if stat == "hp_max" {
			hpMaxIndex = i
		}
		if stat == "vitality" {
			vitalityIndex = i
		}
	}

	if hpMaxIndex == -1 || vitalityIndex == -1 {
		t.Error("Expected hp_max and vitality to be in calculation order")
	}

	if vitalityIndex >= hpMaxIndex {
		t.Error("Expected vitality to come before hp_max in calculation order")
	}
}

func TestValidateStats(t *testing.T) {
	sr := coreService.NewStatResolver()

	// Test valid stats
	validStats := map[string]float64{
		"hp_max":      100.0,
		"stamina":     100.0,
		"crit_chance": 0.05,
		"crit_multi":  1.5,
		"haste":       1.1,
	}

	err := sr.ValidateStats(validStats)
	if err != nil {
		t.Errorf("Expected no error for valid stats, got %v", err)
	}

	// Test invalid stats
	invalidStats := map[string]float64{
		"hp_max":      -100.0, // negative value
		"stamina":     100.0,
		"crit_chance": 1.5,  // exceeds 1.0
		"crit_multi":  0.5,  // less than 1.0
		"haste":       0.05, // less than 0.1
	}

	err = sr.ValidateStats(invalidStats)
	if err == nil {
		t.Error("Expected error for invalid stats")
	}
}

func TestAddFormula(t *testing.T) {
	sr := coreService.NewStatResolver()

	// Create a custom formula
	formula := &coreService.BasicFormula{
		Name:         "custom_stat",
		Type:         "calculation",
		Dependencies: []string{"vitality"},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Vitality) * 2.0
		},
	}

	err := sr.AddFormula(formula)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test that formula was added
	_, err = sr.GetFormula("custom_stat")
	if err != nil {
		t.Errorf("Expected formula to be added, got %v", err)
	}

	// Test adding nil formula
	err = sr.AddFormula(nil)
	if err == nil {
		t.Error("Expected error for nil formula")
	}

	// Test adding formula with empty name
	emptyFormula := &coreService.BasicFormula{
		Name:         "",
		Type:         "calculation",
		Dependencies: []string{},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return 0.0
		},
	}

	err = sr.AddFormula(emptyFormula)
	if err == nil {
		t.Error("Expected error for empty formula name")
	}
}

func TestRemoveFormula(t *testing.T) {
	sr := coreService.NewStatResolver()

	// Test removing existing formula
	err := sr.RemoveFormula("hp_max")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test that formula was removed
	_, err = sr.GetFormula("hp_max")
	if err == nil {
		t.Error("Expected formula to be removed")
	}

	// Test removing non-existing formula
	err = sr.RemoveFormula("non_existing")
	if err == nil {
		t.Error("Expected error for non-existing formula")
	}
}

func TestGetFormula(t *testing.T) {
	sr := coreService.NewStatResolver()

	// Test getting existing formula
	formula, err := sr.GetFormula("hp_max")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if formula == nil {
		t.Error("Expected formula to be returned")
	}

	if formula.GetName() != "hp_max" {
		t.Errorf("Expected formula name to be hp_max, got %s", formula.GetName())
	}

	// Test getting non-existing formula
	_, err = sr.GetFormula("non_existing")
	if err == nil {
		t.Error("Expected error for non-existing formula")
	}
}

func TestGetAllFormulas(t *testing.T) {
	sr := coreService.NewStatResolver()

	formulas := sr.GetAllFormulas()
	if len(formulas) == 0 {
		t.Error("Expected formulas to be returned")
	}

	// Check that hp_max formula exists
	if _, exists := formulas["hp_max"]; !exists {
		t.Error("Expected hp_max formula to exist")
	}
}

func TestClearCache(t *testing.T) {
	sr := coreService.NewStatResolver()
	pc := core.NewPrimaryCore()

	// Resolve some stats to populate cache
	_, err := sr.ResolveStats(pc)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	cacheSize := sr.GetCacheSize()
	if cacheSize == 0 {
		t.Error("Expected cache to be populated")
	}

	// Clear cache
	sr.ClearCache()

	cacheSize = sr.GetCacheSize()
	if cacheSize != 0 {
		t.Error("Expected cache to be cleared")
	}
}

func TestGetCacheSize(t *testing.T) {
	sr := coreService.NewStatResolver()
	pc := core.NewPrimaryCore()

	// Initially cache should be empty
	cacheSize := sr.GetCacheSize()
	if cacheSize != 0 {
		t.Error("Expected cache to be empty initially")
	}

	// Resolve some stats to populate cache
	_, err := sr.ResolveStats(pc)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	cacheSize = sr.GetCacheSize()
	if cacheSize == 0 {
		t.Error("Expected cache to be populated")
	}
}

func TestGetVersion(t *testing.T) {
	sr := coreService.NewStatResolver()

	if sr.GetVersion() != 1 {
		t.Errorf("Expected version to be 1, got %d", sr.GetVersion())
	}

	// Add a formula to increment version
	formula := &coreService.BasicFormula{
		Name:         "test_stat",
		Type:         "calculation",
		Dependencies: []string{},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return 0.0
		},
	}

	sr.AddFormula(formula)

	if sr.GetVersion() != 2 {
		t.Errorf("Expected version to be 2, got %d", sr.GetVersion())
	}
}

func TestGetStatsCount(t *testing.T) {
	sr := coreService.NewStatResolver()

	count := sr.GetStatsCount()
	if count == 0 {
		t.Error("Expected stats count to be non-zero")
	}

	initialCount := count

	// Add a formula
	formula := &coreService.BasicFormula{
		Name:         "test_stat2",
		Type:         "calculation",
		Dependencies: []string{},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return 0.0
		},
	}

	sr.AddFormula(formula)

	count = sr.GetStatsCount()
	if count != initialCount+1 {
		t.Errorf("Expected stats count to be %d, got %d", initialCount+1, count)
	}
}
