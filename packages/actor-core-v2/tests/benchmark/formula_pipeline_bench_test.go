package benchmark

import (
	"testing"

	"actor-core/constants"
	"actor-core/models/core"
	coreServices "actor-core/services/core"
)

func BenchmarkFormulaPipeline_AddFormula(b *testing.B) {
	fp := coreServices.NewFormulaPipeline()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		steps := []coreServices.Step{
			&coreServices.MultStep{Formula: "2.0"},
		}
		fp.AddFormula("test_stat", []string{constants.Stat_VITALITY}, steps)
	}
}

func BenchmarkFormulaPipeline_BuildOrder(b *testing.B) {
	fp := coreServices.NewFormulaPipeline()

	// Pre-populate with formulas
	for i := 0; i < 100; i++ {
		steps := []coreServices.Step{
			&coreServices.MultStep{Formula: "2.0"},
		}
		fp.AddFormula("test_stat", []string{constants.Stat_VITALITY}, steps)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := fp.BuildOrder()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFormulaPipeline_ResolveDerivedStats(b *testing.B) {
	fp := coreServices.NewFormulaPipeline()
	pc := core.NewPrimaryCore()
	pc.Vitality = 100
	pc.Strength = 80
	pc.Endurance = 90
	pc.Intelligence = 120
	pc.Wisdom = 110
	pc.Charisma = 95
	pc.Luck = 85
	pc.SpiritualEnergy = 150

	// Add some formulas
	hpSteps := []coreServices.Step{
		&coreServices.MultStep{Formula: "10.0"},
	}
	fp.AddFormula(constants.Stat_HP_MAX, []string{constants.Stat_VITALITY}, hpSteps)

	mpSteps := []coreServices.Step{
		&coreServices.MultStep{Formula: "8.0"},
	}
	fp.AddFormula(constants.Stat_MANA_MAX, []string{constants.Stat_INTELLIGENCE}, mpSteps)

	staminaSteps := []coreServices.Step{
		&coreServices.MultStep{Formula: "5.0"},
	}
	fp.AddFormula(constants.Stat_STAMINA_MAX, []string{constants.Stat_ENDURANCE}, staminaSteps)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := fp.ResolveDerivedStats(pc, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFormulaPipeline_ResolveDerivedStats_Complex(b *testing.B) {
	fp := coreServices.NewFormulaPipeline()
	pc := core.NewPrimaryCore()
	pc.Vitality = 1000
	pc.Strength = 800
	pc.Endurance = 900
	pc.Intelligence = 1200
	pc.Wisdom = 1100
	pc.Charisma = 950
	pc.Luck = 850
	pc.SpiritualEnergy = 1500

	// Add complex formulas with multiple dependencies
	hpSteps := []coreServices.Step{
		&coreServices.MultStep{Formula: "10.0"},
		&coreServices.FlatStep{Formula: "100.0"},
	}
	fp.AddFormula(constants.Stat_HP_MAX, []string{constants.Stat_VITALITY, constants.Stat_STRENGTH}, hpSteps)

	mpSteps := []coreServices.Step{
		&coreServices.MultStep{Formula: "8.0"},
		&coreServices.FlatStep{Formula: "50.0"},
	}
	fp.AddFormula(constants.Stat_MANA_MAX, []string{constants.Stat_INTELLIGENCE, constants.Stat_WISDOM}, mpSteps)

	staminaSteps := []coreServices.Step{
		&coreServices.MultStep{Formula: "5.0"},
		&coreServices.FlatStep{Formula: "25.0"},
	}
	fp.AddFormula(constants.Stat_STAMINA_MAX, []string{constants.Stat_ENDURANCE, constants.Stat_VITALITY}, staminaSteps)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := fp.ResolveDerivedStats(pc, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFormulaPipeline_ConcurrentAccess(b *testing.B) {
	fp := coreServices.NewFormulaPipeline()
	pc := core.NewPrimaryCore()
	pc.Vitality = 100

	// Pre-populate with formulas
	for i := 0; i < 50; i++ {
		steps := []coreServices.Step{
			&coreServices.MultStep{Formula: "2.0"},
		}
		fp.AddFormula("test_stat", []string{constants.Stat_VITALITY}, steps)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%2 == 0 {
				fp.BuildOrder()
			} else {
				fp.ResolveDerivedStats(pc, nil)
			}
			i++
		}
	})
}
