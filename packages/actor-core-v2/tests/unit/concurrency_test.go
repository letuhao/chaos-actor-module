package unit

import (
	"strings"
	"sync"
	"testing"
	"time"

	"actor-core/constants"
	"actor-core/models/core"
	coreServices "actor-core/services/core"
)

func TestStatResolverConcurrency(t *testing.T) {
	sr := coreServices.NewStatResolver()
	pc := core.NewPrimaryCore()
	pc.Vitality = 50
	pc.Constitution = 30

	// Test concurrent reads
	var wg sync.WaitGroup
	numGoroutines := 10
	numIterations := 100

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				_, err := sr.ResolveStat(constants.Stat_HP_MAX, pc)
				if err != nil {
					t.Errorf("ResolveStat failed: %v", err)
				}
			}
		}()
	}
	wg.Wait()
}

func TestStatResolverConcurrentReadWrite(t *testing.T) {
	sr := coreServices.NewStatResolver()
	pc := core.NewPrimaryCore()
	pc.Vitality = 50

	var wg sync.WaitGroup

	// Concurrent reads
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				_, err := sr.ResolveStat(constants.Stat_HP_MAX, pc)
				if err != nil {
					t.Errorf("ResolveStat failed: %v", err)
				}
			}
		}()
	}

	// Concurrent cache operations
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			sr.ClearCache()
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			size := sr.GetCacheSize()
			_ = size
			time.Sleep(time.Millisecond)
		}
	}()

	wg.Wait()
}

func TestStatResolverConcurrentFormulaOperations(t *testing.T) {
	sr := coreServices.NewStatResolver()
	pc := core.NewPrimaryCore()
	pc.Vitality = 50

	var wg sync.WaitGroup

	// Concurrent formula additions
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			formula := &coreServices.BasicFormula{
				Name:         "test_stat_1",
				Type:         constants.FormulaTypeCalculation,
				Dependencies: []string{constants.Stat_VITALITY},
				Calculator: func(primary *core.PrimaryCore) float64 {
					return float64(primary.Vitality) * 2.0
				},
			}
			sr.AddFormula(formula)
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			formula := &coreServices.BasicFormula{
				Name:         "test_stat_2",
				Type:         constants.FormulaTypeCalculation,
				Dependencies: []string{constants.Stat_VITALITY},
				Calculator: func(primary *core.PrimaryCore) float64 {
					return float64(primary.Vitality) * 3.0
				},
			}
			sr.AddFormula(formula)
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			_, err := sr.GetFormula("test_stat_1")
			if err != nil {
				// Formula might not exist yet, that's ok
			}
			time.Sleep(time.Millisecond)
		}
	}()

	wg.Wait()
}

func TestDAGBuilderConcurrency(t *testing.T) {
	db := coreServices.NewDAGBuilder()

	var wg sync.WaitGroup

	// Concurrent node additions
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			db.AddNode("node1", []string{"node2"})
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			db.AddNode("node2", []string{})
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			order, err := db.BuildOrder()
			if err != nil {
				t.Errorf("BuildOrder failed: %v", err)
			}
			_ = order
			time.Sleep(time.Millisecond)
		}
	}()

	wg.Wait()
}

func TestStatResolverVersionIncrement(t *testing.T) {
	sr := coreServices.NewStatResolver()
	initialVersion := sr.GetVersion()

	// Add formula should increment version
	formula := &coreServices.BasicFormula{
		Name:         "test_stat",
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_VITALITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Vitality) * 2.0
		},
	}

	err := sr.AddFormula(formula)
	if err != nil {
		t.Errorf("AddFormula failed: %v", err)
	}

	if sr.GetVersion() != initialVersion+1 {
		t.Errorf("Expected version %d, got %d", initialVersion+1, sr.GetVersion())
	}

	// Remove formula should increment version
	err = sr.RemoveFormula("test_stat")
	if err != nil {
		t.Errorf("RemoveFormula failed: %v", err)
	}

	if sr.GetVersion() != initialVersion+2 {
		t.Errorf("Expected version %d, got %d", initialVersion+2, sr.GetVersion())
	}
}

func TestStatResolverCacheVersioning(t *testing.T) {
	sr := coreServices.NewStatResolver()
	pc := core.NewPrimaryCore()
	pc.Vitality = 50

	// Resolve stat - should cache with version 1
	value1, err := sr.ResolveStat(constants.Stat_HP_MAX, pc)
	if err != nil {
		t.Errorf("ResolveStat failed: %v", err)
	}

	// Add formula - should increment version
	formula := &coreServices.BasicFormula{
		Name:         "test_stat",
		Type:         constants.FormulaTypeCalculation,
		Dependencies: []string{constants.Stat_VITALITY},
		Calculator: func(primary *core.PrimaryCore) float64 {
			return float64(primary.Vitality) * 2.0
		},
	}

	err = sr.AddFormula(formula)
	if err != nil {
		t.Errorf("AddFormula failed: %v", err)
	}

	// Resolve same stat - should not use cache due to version change
	value2, err := sr.ResolveStat(constants.Stat_HP_MAX, pc)
	if err != nil {
		t.Errorf("ResolveStat failed: %v", err)
	}

	// Values should be the same (same formula, same input)
	if value1 != value2 {
		t.Errorf("Expected same values, got %f and %f", value1, value2)
	}
}

func TestDAGBuilderCircularDependency(t *testing.T) {
	db := coreServices.NewDAGBuilder()

	// Create circular dependency: A -> B -> C -> A
	db.AddNode("A", []string{"C"})
	db.AddNode("B", []string{"A"})
	db.AddNode("C", []string{"B"})

	_, err := db.BuildOrder()
	if err == nil {
		t.Error("Expected error for circular dependency")
	}

	if !strings.Contains(err.Error(), "circular dependency detected") {
		t.Errorf("Expected circular dependency error, got: %v", err)
	}
}

func TestDAGBuilderValidOrder(t *testing.T) {
	db := coreServices.NewDAGBuilder()

	// Create valid DAG: A -> B -> C, D -> C
	db.AddNode("A", []string{})
	db.AddNode("B", []string{"A"})
	db.AddNode("C", []string{"B"})
	db.AddNode("D", []string{})
	db.AddNode("E", []string{"C", "D"})

	order, err := db.BuildOrder()
	if err != nil {
		t.Errorf("BuildOrder failed: %v", err)
	}

	// Check that dependencies come before dependents
	positions := make(map[string]int)
	for i, node := range order {
		positions[node] = i
	}

	// A should come before B
	if positions["A"] >= positions["B"] {
		t.Error("A should come before B")
	}

	// B should come before C
	if positions["B"] >= positions["C"] {
		t.Error("B should come before C")
	}

	// C and D should come before E
	if positions["C"] >= positions["E"] {
		t.Error("C should come before E")
	}
	if positions["D"] >= positions["E"] {
		t.Error("D should come before E")
	}
}
