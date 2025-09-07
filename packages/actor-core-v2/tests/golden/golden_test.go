package golden

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"actor-core/models/core"
	coreServices "actor-core/services/core"
)

// GoldenTestData represents the structure for golden test data
type GoldenTestData struct {
	Input  GoldenInput  `json:"input"`
	Output GoldenOutput `json:"output"`
}

type GoldenInput struct {
	PrimaryStats PrimaryStatsInput `json:"primary_stats"`
}

type PrimaryStatsInput struct {
	Vitality        int64 `json:"vitality"`
	Constitution    int64 `json:"constitution"`
	Agility         int64 `json:"agility"`
	Intelligence    int64 `json:"intelligence"`
	Wisdom          int64 `json:"wisdom"`
	Charisma        int64 `json:"charisma"`
	Luck            int64 `json:"luck"`
	Karma           int64 `json:"karma"`
	Endurance       int64 `json:"endurance"`
	Strength        int64 `json:"strength"`
	SpiritualEnergy int64 `json:"spiritual_energy"`
	PhysicalEnergy  int64 `json:"physical_energy"`
	MentalEnergy    int64 `json:"mental_energy"`
	Willpower       int64 `json:"willpower"`
}

type GoldenOutput struct {
	DerivedStats map[string]float64 `json:"derived_stats"`
}

// generateGoldenTestData generates test data for all derived stats
func generateGoldenTestData() []GoldenTestData {
	var testData []GoldenTestData

	// Generate grid of primary stats (1..100 step 10)
	for vitality := 1; vitality <= 100; vitality += 10 {
		for constitution := 1; constitution <= 100; constitution += 10 {
			for agility := 1; agility <= 100; agility += 10 {
				for intelligence := 1; intelligence <= 100; intelligence += 10 {
					for wisdom := 1; wisdom <= 100; wisdom += 10 {
						// Limit combinations to avoid too many tests
						if len(testData) >= 50 {
							break
						}

						primaryStats := PrimaryStatsInput{
							Vitality:        int64(vitality),
							Constitution:    int64(constitution),
							Agility:         int64(agility),
							Intelligence:    int64(intelligence),
							Wisdom:          int64(wisdom),
							Charisma:        int64(50), // Fixed values for other stats
							Luck:            int64(50),
							Karma:           int64(50),
							Endurance:       int64(50),
							Strength:        int64(50),
							SpiritualEnergy: int64(50),
							PhysicalEnergy:  int64(50),
							MentalEnergy:    int64(50),
							Willpower:       int64(50),
						}

						testData = append(testData, GoldenTestData{
							Input: GoldenInput{
								PrimaryStats: primaryStats,
							},
						})
					}
					if len(testData) >= 50 {
						break
					}
				}
				if len(testData) >= 50 {
					break
				}
			}
			if len(testData) >= 50 {
				break
			}
		}
		if len(testData) >= 50 {
			break
		}
	}

	return testData
}

// TestGoldenDerivedStats tests all derived stats with golden data
func TestGoldenDerivedStats(t *testing.T) {
	sr := coreServices.NewStatResolver()
	testData := generateGoldenTestData()

	// Create testdata directory
	testdataDir := "testdata/derived"
	if err := os.MkdirAll(testdataDir, 0755); err != nil {
		t.Fatalf("Failed to create testdata directory: %v", err)
	}

	for i, data := range testData {
		// Create primary core from input
		pc := &core.PrimaryCore{
			Vitality:        data.Input.PrimaryStats.Vitality,
			Constitution:    data.Input.PrimaryStats.Constitution,
			Agility:         data.Input.PrimaryStats.Agility,
			Intelligence:    data.Input.PrimaryStats.Intelligence,
			Wisdom:          data.Input.PrimaryStats.Wisdom,
			Charisma:        data.Input.PrimaryStats.Charisma,
			Luck:            data.Input.PrimaryStats.Luck,
			Karma:           data.Input.PrimaryStats.Karma,
			Endurance:       data.Input.PrimaryStats.Endurance,
			Strength:        data.Input.PrimaryStats.Strength,
			SpiritualEnergy: data.Input.PrimaryStats.SpiritualEnergy,
			PhysicalEnergy:  data.Input.PrimaryStats.PhysicalEnergy,
			MentalEnergy:    data.Input.PrimaryStats.MentalEnergy,
			Willpower:       data.Input.PrimaryStats.Willpower,
		}

		// Resolve derived stats
		derivedStats, err := sr.ResolveStats(pc)
		if err != nil {
			t.Errorf("Test case %d: Failed to resolve stats: %v", i, err)
			continue
		}

		// Get all derived stats
		allStats := derivedStats.GetAllStats()

		// Create golden output
		data.Output.DerivedStats = allStats

		// Write golden file
		goldenFile := filepath.Join(testdataDir, fmt.Sprintf("test_%d.golden.json", i))
		goldenData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			t.Errorf("Test case %d: Failed to marshal golden data: %v", i, err)
			continue
		}

		if err := os.WriteFile(goldenFile, goldenData, 0644); err != nil {
			t.Errorf("Test case %d: Failed to write golden file: %v", i, err)
			continue
		}

		t.Logf("Generated golden file: %s", goldenFile)
	}
}

// TestGoldenDerivedStatsLoad tests loading and validating golden data
func TestGoldenDerivedStatsLoad(t *testing.T) {
	testdataDir := "testdata/derived"

	// Find all golden files
	matches, err := filepath.Glob(filepath.Join(testdataDir, "*.golden.json"))
	if err != nil {
		t.Fatalf("Failed to find golden files: %v", err)
	}

	if len(matches) == 0 {
		t.Skip("No golden files found, run TestGoldenDerivedStats first")
	}

	sr := coreServices.NewStatResolver()

	for _, goldenFile := range matches {
		// Read golden file
		goldenData, err := os.ReadFile(goldenFile)
		if err != nil {
			t.Errorf("Failed to read golden file %s: %v", goldenFile, err)
			continue
		}

		// Parse golden data
		var data GoldenTestData
		if err := json.Unmarshal(goldenData, &data); err != nil {
			t.Errorf("Failed to parse golden file %s: %v", goldenFile, err)
			continue
		}

		// Create primary core from input
		pc := &core.PrimaryCore{
			Vitality:        data.Input.PrimaryStats.Vitality,
			Constitution:    data.Input.PrimaryStats.Constitution,
			Agility:         data.Input.PrimaryStats.Agility,
			Intelligence:    data.Input.PrimaryStats.Intelligence,
			Wisdom:          data.Input.PrimaryStats.Wisdom,
			Charisma:        data.Input.PrimaryStats.Charisma,
			Luck:            data.Input.PrimaryStats.Luck,
			Karma:           data.Input.PrimaryStats.Karma,
			Endurance:       data.Input.PrimaryStats.Endurance,
			Strength:        data.Input.PrimaryStats.Strength,
			SpiritualEnergy: data.Input.PrimaryStats.SpiritualEnergy,
			PhysicalEnergy:  data.Input.PrimaryStats.PhysicalEnergy,
			MentalEnergy:    data.Input.PrimaryStats.MentalEnergy,
			Willpower:       data.Input.PrimaryStats.Willpower,
		}

		// Resolve derived stats
		derivedStats, err := sr.ResolveStats(pc)
		if err != nil {
			t.Errorf("Golden file %s: Failed to resolve stats: %v", goldenFile, err)
			continue
		}

		// Get all derived stats
		allStats := derivedStats.GetAllStats()

		// Compare with golden data
		for statName, expectedValue := range data.Output.DerivedStats {
			actualValue, exists := allStats[statName]
			if !exists {
				t.Errorf("Golden file %s: Stat %s not found in actual output", goldenFile, statName)
				continue
			}

			// Allow small floating point differences
			diff := actualValue - expectedValue
			if diff < -0.001 || diff > 0.001 {
				t.Errorf("Golden file %s: Stat %s mismatch - expected %f, got %f (diff: %f)",
					goldenFile, statName, expectedValue, actualValue, diff)
			}
		}
	}
}
