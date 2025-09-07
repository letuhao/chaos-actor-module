package race

import (
	"sync"
	"testing"

	"actor-core/constants"
	"actor-core/models/core"
	coreServices "actor-core/services/core"
)

// TestResolveDerived_Parallel tests parallel resolution of derived stats
func TestResolveDerived_Parallel(t *testing.T) {
	sr := coreServices.NewStatResolver()

	// Create test primary cores
	primaryCores := make([]*core.PrimaryCore, 10)
	for i := 0; i < 10; i++ {
		primaryCores[i] = &core.PrimaryCore{
			Vitality:        int64(50 + i*5),
			Constitution:    int64(45 + i*3),
			Agility:         int64(60 + i*2),
			Intelligence:    int64(55 + i*4),
			Wisdom:          int64(40 + i*6),
			Charisma:        int64(35 + i*7),
			Luck:            int64(30 + i*8),
			Karma:           int64(25 + i*9),
			Endurance:       int64(70 + i*1),
			Strength:        int64(65 + i*2),
			SpiritualEnergy: int64(50 + i*3),
			PhysicalEnergy:  int64(55 + i*4),
			MentalEnergy:    int64(60 + i*5),
			Willpower:       int64(45 + i*6),
		}
	}

	// Test parallel resolution
	var wg sync.WaitGroup
	numGoroutines := 20
	results := make([][]*core.DerivedStats, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()

			// Each goroutine resolves stats for all primary cores
			results[goroutineID] = make([]*core.DerivedStats, len(primaryCores))
			for j, pc := range primaryCores {
				derivedStats, err := sr.ResolveStats(pc)
				if err != nil {
					t.Errorf("Goroutine %d, PC %d: Failed to resolve stats: %v", goroutineID, j, err)
					return
				}
				results[goroutineID][j] = derivedStats
			}
		}(i)
	}

	wg.Wait()

	// Verify that all results are identical for the same input
	for i := 0; i < len(primaryCores); i++ {
		firstResult := results[0][i]
		if firstResult == nil {
			continue
		}

		firstValues := firstResult.GetAllStats()

		for goroutineID := 1; goroutineID < numGoroutines; goroutineID++ {
			result := results[goroutineID][i]
			if result == nil {
				continue
			}

			values := result.GetAllStats()

			// Compare all stat values
			for statName, firstValue := range firstValues {
				value, exists := values[statName]
				if !exists {
					t.Errorf("Goroutine %d: Stat %s missing", goroutineID, statName)
					continue
				}

				if firstValue != value {
					t.Errorf("Goroutine %d: Stat %s mismatch - expected %f, got %f",
						goroutineID, statName, firstValue, value)
				}
			}
		}
	}
}

// TestStatResolverConcurrentOperations tests concurrent operations on StatResolver
func TestStatResolverConcurrentOperations(t *testing.T) {
	sr := coreServices.NewStatResolver()

	pc := &core.PrimaryCore{
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

	var wg sync.WaitGroup
	numGoroutines := 50

	// Test concurrent ResolveStats calls
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()

			for j := 0; j < 10; j++ {
				_, err := sr.ResolveStats(pc)
				if err != nil {
					t.Errorf("Goroutine %d, iteration %d: Failed to resolve stats: %v", goroutineID, j, err)
				}
			}
		}(i)
	}

	wg.Wait()
}

// TestStatResolverConcurrentFormulaOperations tests concurrent formula operations
func TestStatResolverConcurrentFormulaOperations(t *testing.T) {
	sr := coreServices.NewStatResolver()

	pc := &core.PrimaryCore{
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

	var wg sync.WaitGroup
	numGoroutines := 20

	// Test concurrent formula operations
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()

			// Mix of different operations
			for j := 0; j < 5; j++ {
				switch j % 4 {
				case 0:
					// Resolve stats
					_, err := sr.ResolveStats(pc)
					if err != nil {
						t.Errorf("Goroutine %d: Failed to resolve stats: %v", goroutineID, err)
					}
				case 1:
					// Get formula
					_, err := sr.GetFormula(constants.Stat_HP_MAX)
					if err != nil {
						t.Errorf("Goroutine %d: Failed to get formula: %v", goroutineID, err)
					}
				case 2:
					// Get all formulas
					_ = sr.GetAllFormulas()
				case 3:
					// Get cache size
					_ = sr.GetCacheSize()
				}
			}
		}(i)
	}

	wg.Wait()
}

// TestStatResolverConcurrentCacheOperations tests concurrent cache operations
func TestStatResolverConcurrentCacheOperations(t *testing.T) {
	sr := coreServices.NewStatResolver()

	pc := &core.PrimaryCore{
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

	var wg sync.WaitGroup
	numGoroutines := 30

	// Test concurrent cache operations
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()

			for j := 0; j < 10; j++ {
				switch j % 3 {
				case 0:
					// Resolve stats (populates cache)
					_, err := sr.ResolveStats(pc)
					if err != nil {
						t.Errorf("Goroutine %d: Failed to resolve stats: %v", goroutineID, err)
					}
				case 1:
					// Get cache size
					_ = sr.GetCacheSize()
				case 2:
					// Clear cache
					sr.ClearCache()
				}
			}
		}(i)
	}

	wg.Wait()
}

// TestStatResolverConcurrentStatResolution tests concurrent resolution of individual stats
func TestStatResolverConcurrentStatResolution(t *testing.T) {
	sr := coreServices.NewStatResolver()

	pc := &core.PrimaryCore{
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

	// Test concurrent resolution of individual stats
	statNames := []string{
		constants.Stat_HP_MAX,
		constants.Stat_MOVE_SPEED,
		constants.Stat_STAMINA,
		constants.Stat_CRIT_CHANCE,
		constants.Stat_CRIT_MULTI,
		constants.Stat_ACCURACY,
		constants.Stat_PENETRATION,
		constants.Stat_BRUTALITY,
	}

	var wg sync.WaitGroup
	numGoroutines := 20

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(goroutineID int) {
			defer wg.Done()

			for j := 0; j < 10; j++ {
				statName := statNames[j%len(statNames)]
				_, err := sr.ResolveStat(statName, pc)
				if err != nil {
					t.Errorf("Goroutine %d: Failed to resolve stat %s: %v", goroutineID, statName, err)
				}
			}
		}(i)
	}

	wg.Wait()
}
