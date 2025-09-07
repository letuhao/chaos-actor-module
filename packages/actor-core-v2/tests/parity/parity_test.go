package parity

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"actor-core/models/core"
	coreServices "actor-core/services/core"
)

// ParityTestData represents the structure for parity test data
type ParityTestData struct {
	Input  ParityInput  `json:"input"`
	Output ParityOutput `json:"output"`
}

type ParityInput struct {
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

type ParityOutput struct {
	DerivedStats map[string]float64 `json:"derived_stats"`
}

// TestGoTypeScriptParity tests that Go and TypeScript implementations produce identical results
func TestGoTypeScriptParity(t *testing.T) {
	// Skip if TypeScript implementation is not available
	if !isTypeScriptAvailable() {
		t.Skip("TypeScript implementation not available")
	}

	sr := coreServices.NewStatResolver()

	// Generate test cases
	testCases := generateParityTestCases()

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase_%d", i), func(t *testing.T) {
			// Test Go implementation
			goResult, err := testGoImplementation(sr, testCase.Input)
			if err != nil {
				t.Fatalf("Go implementation failed: %v", err)
			}

			// Test TypeScript implementation
			tsResult, err := testTypeScriptImplementation(testCase.Input)
			if err != nil {
				t.Fatalf("TypeScript implementation failed: %v", err)
			}

			// Compare results
			if err := compareResults(goResult, tsResult); err != nil {
				t.Errorf("Results mismatch: %v", err)
			}
		})
	}
}

// generateParityTestCases generates test cases for parity testing
func generateParityTestCases() []ParityTestData {
	var testCases []ParityTestData

	// Generate a variety of test cases
	testValues := []int64{1, 10, 25, 50, 75, 100}

	for _, vitality := range testValues {
		for _, constitution := range testValues {
			for _, agility := range testValues {
				// Limit combinations to avoid too many tests
				if len(testCases) >= 20 {
					break
				}

				testCases = append(testCases, ParityTestData{
					Input: ParityInput{
						PrimaryStats: PrimaryStatsInput{
							Vitality:        vitality,
							Constitution:    constitution,
							Agility:         agility,
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
						},
					},
				})
			}
			if len(testCases) >= 20 {
				break
			}
		}
		if len(testCases) >= 20 {
			break
		}
	}

	return testCases
}

// testGoImplementation tests the Go implementation
func testGoImplementation(sr *coreServices.StatResolver, input ParityInput) (map[string]float64, error) {
	pc := &core.PrimaryCore{
		Vitality:        input.PrimaryStats.Vitality,
		Constitution:    input.PrimaryStats.Constitution,
		Agility:         input.PrimaryStats.Agility,
		Intelligence:    input.PrimaryStats.Intelligence,
		Wisdom:          input.PrimaryStats.Wisdom,
		Charisma:        input.PrimaryStats.Charisma,
		Luck:            input.PrimaryStats.Luck,
		Karma:           input.PrimaryStats.Karma,
		Endurance:       input.PrimaryStats.Endurance,
		Strength:        input.PrimaryStats.Strength,
		SpiritualEnergy: input.PrimaryStats.SpiritualEnergy,
		PhysicalEnergy:  input.PrimaryStats.PhysicalEnergy,
		MentalEnergy:    input.PrimaryStats.MentalEnergy,
		Willpower:       input.PrimaryStats.Willpower,
	}

	derivedStats, err := sr.ResolveStats(pc)
	if err != nil {
		return nil, err
	}

	return derivedStats.GetAllStats(), nil
}

// testTypeScriptImplementation tests the TypeScript implementation
func testTypeScriptImplementation(input ParityInput) (map[string]float64, error) {
	// Create temporary input file
	inputFile := "test_input.json"
	outputFile := "test_output.json"

	// Marshal input to JSON
	inputData, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input: %v", err)
	}

	// Write input file
	if err := os.WriteFile(inputFile, inputData, 0644); err != nil {
		return nil, fmt.Errorf("failed to write input file: %v", err)
	}
	defer os.Remove(inputFile)
	defer os.Remove(outputFile)

	// Run TypeScript implementation
	cmd := exec.Command("node", "packages/shared/stats.gen.js", inputFile, outputFile)
	cmd.Dir = filepath.Join("..", "..", "..", "..") // Go up to project root

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("TypeScript implementation failed: %v, output: %s", err, string(output))
	}

	// Read output file
	outputData, err := os.ReadFile(outputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read output file: %v", err)
	}

	// Parse output
	var result ParityOutput
	if err := json.Unmarshal(outputData, &result); err != nil {
		return nil, fmt.Errorf("failed to parse output: %v", err)
	}

	return result.DerivedStats, nil
}

// compareResults compares Go and TypeScript results
func compareResults(goResult, tsResult map[string]float64) error {
	// Check that all Go stats exist in TypeScript result
	for statName, goValue := range goResult {
		tsValue, exists := tsResult[statName]
		if !exists {
			return fmt.Errorf("stat %s missing in TypeScript result", statName)
		}

		// Allow small floating point differences
		diff := goValue - tsValue
		if diff < -0.001 || diff > 0.001 {
			return fmt.Errorf("stat %s mismatch - Go: %f, TypeScript: %f (diff: %f)",
				statName, goValue, tsValue, diff)
		}
	}

	// Check that all TypeScript stats exist in Go result
	for statName, tsValue := range tsResult {
		goValue, exists := goResult[statName]
		if !exists {
			return fmt.Errorf("stat %s missing in Go result", statName)
		}

		// Allow small floating point differences
		diff := goValue - tsValue
		if diff < -0.001 || diff > 0.001 {
			return fmt.Errorf("stat %s mismatch - Go: %f, TypeScript: %f (diff: %f)",
				statName, goValue, tsValue, diff)
		}
	}

	return nil
}

// isTypeScriptAvailable checks if TypeScript implementation is available
func isTypeScriptAvailable() bool {
	// Check if TypeScript files exist
	tsFile := filepath.Join("..", "..", "..", "..", "packages", "shared", "stats.gen.js")
	if _, err := os.Stat(tsFile); os.IsNotExist(err) {
		return false
	}

	// Check if Node.js is available
	cmd := exec.Command("node", "--version")
	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

// TestGoTypeScriptParityGolden tests parity using golden data
func TestGoTypeScriptParityGolden(t *testing.T) {
	// Skip if TypeScript implementation is not available
	if !isTypeScriptAvailable() {
		t.Skip("TypeScript implementation not available")
	}

	// Load golden data
	goldenDir := "../golden/testdata/derived"
	matches, err := filepath.Glob(filepath.Join(goldenDir, "*.golden.json"))
	if err != nil {
		t.Fatalf("Failed to find golden files: %v", err)
	}

	if len(matches) == 0 {
		t.Skip("No golden files found, run golden tests first")
	}

	sr := coreServices.NewStatResolver()

	// Test first few golden files
	maxFiles := 5
	if len(matches) > maxFiles {
		matches = matches[:maxFiles]
	}

	for _, goldenFile := range matches {
		t.Run(filepath.Base(goldenFile), func(t *testing.T) {
			// Read golden file
			goldenData, err := os.ReadFile(goldenFile)
			if err != nil {
				t.Fatalf("Failed to read golden file: %v", err)
			}

			// Parse golden data
			var data ParityTestData
			if err := json.Unmarshal(goldenData, &data); err != nil {
				t.Fatalf("Failed to parse golden file: %v", err)
			}

			// Test Go implementation
			goResult, err := testGoImplementation(sr, data.Input)
			if err != nil {
				t.Fatalf("Go implementation failed: %v", err)
			}

			// Test TypeScript implementation
			tsResult, err := testTypeScriptImplementation(data.Input)
			if err != nil {
				t.Fatalf("TypeScript implementation failed: %v", err)
			}

			// Compare results
			if err := compareResults(goResult, tsResult); err != nil {
				t.Errorf("Results mismatch: %v", err)
			}
		})
	}
}
