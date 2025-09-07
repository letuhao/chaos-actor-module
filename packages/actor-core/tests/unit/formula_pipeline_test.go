package unit

import (
	"testing"
	"time"

	"actor-core/constants"
	"actor-core/models/core"
	coreServices "actor-core/services/core"
)

func TestNewFormulaPipeline(t *testing.T) {
	fp := coreServices.NewFormulaPipeline()
	if fp == nil {
		t.Error("Expected non-nil FormulaPipeline")
	}
	if fp.GetOrder() == nil {
		t.Error("Expected non-nil order slice")
	}
}

func TestFormulaPipelineAddFormula(t *testing.T) {
	fp := coreServices.NewFormulaPipeline()

	// Create test steps
	flatStep := &coreServices.FlatStep{Formula: "vitality * 10"}
	multStep := &coreServices.MultStep{Formula: "1 + mastery * 0.02"}
	clampStep := &coreServices.ClampStep{Min: 1, Max: 1000}

	steps := []coreServices.Step{flatStep, multStep, clampStep}
	deps := []string{constants.Stat_VITALITY}

	fp.AddFormula(constants.Stat_HP_MAX, deps, steps)

	formula, exists := fp.GetFormula(constants.Stat_HP_MAX)
	if !exists {
		t.Error("Expected formula to exist")
	}
	if formula.Name != constants.Stat_HP_MAX {
		t.Errorf("Expected name %s, got %s", constants.Stat_HP_MAX, formula.Name)
	}
	if len(formula.Deps) != 1 || formula.Deps[0] != constants.Stat_VITALITY {
		t.Error("Expected correct dependencies")
	}
	if len(formula.Steps) != 3 {
		t.Error("Expected 3 steps")
	}
}

func TestBuildOrder(t *testing.T) {
	fp := coreServices.NewFormulaPipeline()

	// Add formulas with dependencies
	fp.AddFormula(constants.Stat_VITALITY, []string{}, []coreServices.Step{})
	fp.AddFormula(constants.Stat_HP_MAX, []string{constants.Stat_VITALITY}, []coreServices.Step{})
	fp.AddFormula(constants.Stat_STAMINA, []string{constants.Stat_VITALITY}, []coreServices.Step{})

	err := fp.BuildOrder()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	order := fp.GetOrder()
	if len(order) != 3 {
		t.Errorf("Expected 3 formulas in order, got %d", len(order))
	}

	// Vitality should come before HP_MAX and STAMINA
	vitalityIndex := -1
	hpMaxIndex := -1
	staminaIndex := -1

	for i, name := range order {
		switch name {
		case constants.Stat_VITALITY:
			vitalityIndex = i
		case constants.Stat_HP_MAX:
			hpMaxIndex = i
		case constants.Stat_STAMINA:
			staminaIndex = i
		}
	}

	if vitalityIndex == -1 || hpMaxIndex == -1 || staminaIndex == -1 {
		t.Error("Expected all formulas to be in order")
	}

	if vitalityIndex >= hpMaxIndex || vitalityIndex >= staminaIndex {
		t.Error("Expected vitality to come before dependent stats")
	}
}

func TestBuildOrderCircularDependency(t *testing.T) {
	fp := coreServices.NewFormulaPipeline()

	// Create circular dependency: A -> B -> C -> A
	fp.AddFormula("A", []string{"C"}, []coreServices.Step{})
	fp.AddFormula("B", []string{"A"}, []coreServices.Step{})
	fp.AddFormula("C", []string{"B"}, []coreServices.Step{})

	err := fp.BuildOrder()
	if err == nil {
		t.Error("Expected error for circular dependency")
	}
}

func TestResolveDerivedStats(t *testing.T) {
	fp := coreServices.NewFormulaPipeline()

	// Create a simple formula
	flatStep := &coreServices.FlatStep{Formula: "vitality * 10"}
	steps := []coreServices.Step{flatStep}
	deps := []string{constants.Stat_VITALITY}

	fp.AddFormula(constants.Stat_HP_MAX, deps, steps)

	err := fp.BuildOrder()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Create test primary core
	pc := core.NewPrimaryCore()
	pc.Vitality = 50

	// Resolve derived stats
	ds, err := fp.ResolveDerivedStats(pc, nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if ds == nil {
		t.Error("Expected non-nil DerivedStats")
	}

	// Check version was incremented (starts at 0, SetStat increments to 1, then pipeline increments to 2)
	if ds.Version != 3 {
		t.Errorf("Expected version 3, got %d", ds.Version)
	}
}

func TestResolveDerivedStatsWithPrevious(t *testing.T) {
	fp := coreServices.NewFormulaPipeline()

	// Create a simple formula
	flatStep := &coreServices.FlatStep{Formula: "vitality * 10"}
	steps := []coreServices.Step{flatStep}
	deps := []string{constants.Stat_VITALITY}

	fp.AddFormula(constants.Stat_HP_MAX, deps, steps)

	err := fp.BuildOrder()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Create previous derived stats
	previous := core.NewDerivedStats()
	previous.Version = 5
	previous.UpdatedAt = time.Now().Unix()

	// Create test primary core
	pc := core.NewPrimaryCore()
	pc.Vitality = 50

	// Resolve derived stats
	ds, err := fp.ResolveDerivedStats(pc, previous)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if ds == nil {
		t.Error("Expected non-nil DerivedStats")
	}

	// Check version was incremented from previous (5 + SetStat increment + pipeline increment = 7)
	if ds.Version != 7 {
		t.Errorf("Expected version 7, got %d", ds.Version)
	}
}

func TestStepTypes(t *testing.T) {
	// Test FlatStep
	flatStep := &coreServices.FlatStep{Formula: "test"}
	if flatStep.GetType() != "flat" {
		t.Error("Expected FlatStep type to be 'flat'")
	}

	// Test MultStep
	multStep := &coreServices.MultStep{Formula: "test"}
	if multStep.GetType() != "mult" {
		t.Error("Expected MultStep type to be 'mult'")
	}

	// Test ClampStep
	clampStep := &coreServices.ClampStep{Min: 0, Max: 100}
	if clampStep.GetType() != "clamp" {
		t.Error("Expected ClampStep type to be 'clamp'")
	}
}

func TestFormulaPipelineImmutable(t *testing.T) {
	fp := coreServices.NewFormulaPipeline()

	// Add formula
	flatStep := &coreServices.FlatStep{Formula: "vitality * 10"}
	steps := []coreServices.Step{flatStep}
	deps := []string{constants.Stat_VITALITY}

	fp.AddFormula(constants.Stat_HP_MAX, deps, steps)

	err := fp.BuildOrder()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Create test primary core
	pc := core.NewPrimaryCore()
	pc.Vitality = 50

	// Resolve first time
	ds1, err := fp.ResolveDerivedStats(pc, nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Resolve second time with same input
	ds2, err := fp.ResolveDerivedStats(pc, nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Results should be identical but different instances
	if ds1 == ds2 {
		t.Error("Expected different instances for immutable pipeline")
	}

	// Versions should be the same (both start from 0)
	if ds1.Version != ds2.Version {
		t.Error("Expected same version for identical inputs")
	}
}
