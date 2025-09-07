package core

import (
	"actor-core/models/core"
	"fmt"
	"math"
	"time"
)

// Step represents a single step in the formula pipeline
type Step interface {
	Apply(ds *core.DerivedStats, pc *core.PrimaryCore) (float64, error)
	GetType() string
}

// FlatStep applies flat additions to the base value
type FlatStep struct {
	Formula string
}

func (fs *FlatStep) Apply(ds *core.DerivedStats, pc *core.PrimaryCore) (float64, error) {
	// This will be implemented to evaluate the flat formula
	// For now, return 0 as placeholder
	return 0, nil
}

func (fs *FlatStep) GetType() string {
	return "flat"
}

// MultStep applies multiplicative modifiers
type MultStep struct {
	Formula string
}

func (ms *MultStep) Apply(ds *core.DerivedStats, pc *core.PrimaryCore) (float64, error) {
	// This will be implemented to evaluate the mult formula
	// For now, return 1 as placeholder
	return 1, nil
}

func (ms *MultStep) GetType() string {
	return "mult"
}

// ClampStep applies min/max clamping
type ClampStep struct {
	Min float64
	Max float64
}

func (cs *ClampStep) Apply(ds *core.DerivedStats, pc *core.PrimaryCore) (float64, error) {
	// This will be implemented to apply clamping
	// For now, return the value as-is
	return 0, nil
}

func (cs *ClampStep) GetType() string {
	return "clamp"
}

// FormulaDefinition represents a complete formula with steps
type FormulaDefinition struct {
	Name  string
	Deps  []string
	Steps []Step
}

// FormulaPipeline manages the immutable formula evaluation pipeline
type FormulaPipeline struct {
	formulas map[string]*FormulaDefinition
	order    []string
}

// NewFormulaPipeline creates a new formula pipeline
func NewFormulaPipeline() *FormulaPipeline {
	return &FormulaPipeline{
		formulas: make(map[string]*FormulaDefinition),
		order:    make([]string, 0),
	}
}

// AddFormula adds a formula to the pipeline
func (fp *FormulaPipeline) AddFormula(name string, deps []string, steps []Step) {
	fp.formulas[name] = &FormulaDefinition{
		Name:  name,
		Deps:  deps,
		Steps: steps,
	}
}

// BuildOrder performs topological sort to determine evaluation order
func (fp *FormulaPipeline) BuildOrder() error {
	// Simple topological sort implementation
	visited := make(map[string]bool)
	tempVisited := make(map[string]bool)
	order := make([]string, 0)

	var visit func(string) error
	visit = func(node string) error {
		if tempVisited[node] {
			return fmt.Errorf("circular dependency detected: %s", node)
		}
		if visited[node] {
			return nil
		}

		tempVisited[node] = true
		if formula, exists := fp.formulas[node]; exists {
			for _, dep := range formula.Deps {
				if err := visit(dep); err != nil {
					return err
				}
			}
		}
		tempVisited[node] = false
		visited[node] = true
		order = append(order, node)
		return nil
	}

	// Visit all formulas
	for name := range fp.formulas {
		if err := visit(name); err != nil {
			return err
		}
	}

	fp.order = order
	return nil
}

// ResolveDerivedStats resolves all derived stats using the pipeline
func (fp *FormulaPipeline) ResolveDerivedStats(pc *core.PrimaryCore, previous *core.DerivedStats) (*core.DerivedStats, error) {
	// Clone previous snapshot or create fresh
	var ds *core.DerivedStats
	if previous != nil {
		ds = previous.Clone()
	} else {
		ds = core.NewDerivedStats()
	}

	// Evaluate each stat in order
	for _, statName := range fp.order {
		formula, exists := fp.formulas[statName]
		if !exists {
			continue
		}

		// Run through all steps for this stat
		value := 0.0
		for _, step := range formula.Steps {
			stepValue, err := step.Apply(ds, pc)
			if err != nil {
				return nil, fmt.Errorf("error in step %s for stat %s: %w", step.GetType(), statName, err)
			}

			switch step.GetType() {
			case "flat":
				value += stepValue
			case "mult":
				value *= stepValue
			case "clamp":
				value = math.Max(stepValue, math.Min(value, stepValue))
			}
		}

		// Set the final value
		if err := ds.SetStat(statName, value); err != nil {
			return nil, fmt.Errorf("error setting stat %s: %w", statName, err)
		}
	}

	// Bump version and update timestamp
	ds.Version = ds.Version + 1
	ds.UpdatedAt = time.Now().Unix()

	return ds, nil
}

// GetOrder returns the evaluation order
func (fp *FormulaPipeline) GetOrder() []string {
	return fp.order
}

// GetFormula returns a formula by name
func (fp *FormulaPipeline) GetFormula(name string) (*FormulaDefinition, bool) {
	formula, exists := fp.formulas[name]
	return formula, exists
}
