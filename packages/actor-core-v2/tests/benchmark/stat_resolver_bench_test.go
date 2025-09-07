package benchmark

import (
	"testing"

	"actor-core/constants"
	"actor-core/models/core"
	coreServices "actor-core/services/core"
)

func BenchmarkStatResolver_ResolveStats(b *testing.B) {
	sr := coreServices.NewStatResolver()
	pc := core.NewPrimaryCore()
	pc.Vitality = 100
	pc.Strength = 80
	pc.Endurance = 90
	pc.Intelligence = 120
	pc.Wisdom = 110
	pc.Charisma = 95
	pc.Luck = 85
	pc.SpiritualEnergy = 150

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := sr.ResolveStats(pc)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStatResolver_ResolveStats_HighValues(b *testing.B) {
	sr := coreServices.NewStatResolver()
	pc := core.NewPrimaryCore()
	pc.Vitality = 1000
	pc.Strength = 800
	pc.Endurance = 900
	pc.Intelligence = 1200
	pc.Wisdom = 1100
	pc.Charisma = 950
	pc.Luck = 850
	pc.SpiritualEnergy = 1500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := sr.ResolveStats(pc)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStatResolver_ResolveStat(b *testing.B) {
	sr := coreServices.NewStatResolver()
	pc := core.NewPrimaryCore()
	pc.Vitality = 100
	pc.Strength = 80
	pc.Endurance = 90
	pc.Intelligence = 120
	pc.Wisdom = 110
	pc.Charisma = 95
	pc.Luck = 85
	pc.SpiritualEnergy = 150

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := sr.ResolveStat(constants.Stat_HP_MAX, pc)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStatResolver_ResolveStats_Parallel(b *testing.B) {
	sr := coreServices.NewStatResolver()
	pc := core.NewPrimaryCore()
	pc.Vitality = 100
	pc.Strength = 80
	pc.Endurance = 90
	pc.Intelligence = 120
	pc.Wisdom = 110
	pc.Charisma = 95
	pc.Luck = 85
	pc.SpiritualEnergy = 150

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := sr.ResolveStats(pc)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkStatResolver_AddFormula(b *testing.B) {
	sr := coreServices.NewStatResolver()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formula := &coreServices.BasicFormula{
			Name:         "test_stat",
			Type:         "calculation",
			Dependencies: []string{constants.Stat_VITALITY},
			Calculator: func(primary *core.PrimaryCore) float64 {
				return float64(primary.Vitality) * 2.0
			},
		}
		err := sr.AddFormula(formula)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStatResolver_GetCalculationOrder(b *testing.B) {
	sr := coreServices.NewStatResolver()
	pc := core.NewPrimaryCore()
	pc.Vitality = 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sr.GetCalculationOrder()
	}
}
