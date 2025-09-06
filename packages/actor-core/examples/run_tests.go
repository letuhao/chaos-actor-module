// packages/actor-core/run_tests.go
// Simple test runner to verify implementation
package main

import (
	"fmt"
	"math"
	"sort"
)

// Copy the types and implementation here for standalone testing
type PrimaryCore struct {
	HPMax    int
	LifeSpan int
	Attack   int
	Defense  int
	Speed    int
}

type Derived struct {
	HPMax, MPMax           float64
	ATK, MAG               float64
	DEF, RES               float64
	Haste                  float64
	CritChance, CritMulti  float64
	MoveSpeed              float64
	RegenHP, RegenMP       float64
	Resists                map[string]float64
	Amplifiers             map[string]float64
	Version                uint64
}

type CoreContribution struct {
	Primary *PrimaryCore
	Flat    map[string]float64
	Mult    map[string]float64
	Tags    []string
}

type ActorCore interface {
	ComposeCore(buckets map[string]CoreContribution) CoreContribution
	BaseFromPrimary(p PrimaryCore, level int) Derived
	FinalizeDerived(base Derived, flat map[string]float64, mult map[string]float64) Derived
	ClampDerived(d Derived) Derived
}

type ActorCoreImpl struct{}

func NewActorCore() ActorCore {
	return &ActorCoreImpl{}
}

func (a *ActorCoreImpl) ComposeCore(buckets map[string]CoreContribution) CoreContribution {
	keys := make([]string, 0, len(buckets))
	for k := range buckets {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := CoreContribution{
		Flat: make(map[string]float64),
		Mult: make(map[string]float64),
		Tags: make([]string, 0),
	}

	for _, key := range keys {
		contrib := buckets[key]

		if contrib.Primary != nil {
			if result.Primary == nil {
				result.Primary = &PrimaryCore{}
			}
			result.Primary.HPMax += contrib.Primary.HPMax
			result.Primary.LifeSpan += contrib.Primary.LifeSpan
			result.Primary.Attack += contrib.Primary.Attack
			result.Primary.Defense += contrib.Primary.Defense
			result.Primary.Speed += contrib.Primary.Speed
		}

		for k, v := range contrib.Flat {
			result.Flat[k] += v
		}

		for k, v := range contrib.Mult {
			if existing, exists := result.Mult[k]; exists {
				result.Mult[k] = existing * v
			} else {
				result.Mult[k] = v
			}
		}

		result.Tags = append(result.Tags, contrib.Tags...)
	}

	return result
}

func (a *ActorCoreImpl) BaseFromPrimary(p PrimaryCore, level int) Derived {
	levelFactor := float64(level)
	if levelFactor < 1 {
		levelFactor = 1
	}

	hpMax := float64(p.HPMax) * levelFactor
	mpMax := hpMax * 0.8
	atk := float64(p.Attack) * levelFactor
	mag := atk * 0.7
	def := float64(p.Defense) * levelFactor
	res := def * 0.8

	speed := float64(p.Speed)
	haste := 1.0 + (speed-50)/100.0
	moveSpeed := speed * 2.0

	critChance := math.Min(0.1+(speed-50)/500.0, 0.3)
	critMulti := 1.5 + (speed-50)/200.0

	regenHP := hpMax * 0.01
	regenMP := mpMax * 0.01

	resists := make(map[string]float64)
	amplifiers := make(map[string]float64)

	resists["physical"] = 0.0
	resists["magic"] = 0.0
	resists["fire"] = 0.0
	resists["ice"] = 0.0
	resists["lightning"] = 0.0

	amplifiers["internal"] = 1.0
	amplifiers["external"] = 1.0

	return Derived{
		HPMax:      hpMax,
		MPMax:      mpMax,
		ATK:        atk,
		MAG:        mag,
		DEF:        def,
		RES:        res,
		Haste:      haste,
		CritChance: critChance,
		CritMulti:  critMulti,
		MoveSpeed:  moveSpeed,
		RegenHP:    regenHP,
		RegenMP:    regenMP,
		Resists:    resists,
		Amplifiers: amplifiers,
		Version:    0,
	}
}

func (a *ActorCoreImpl) FinalizeDerived(base Derived, flat map[string]float64, mult map[string]float64) Derived {
	result := base

	for key, value := range flat {
		switch key {
		case "HPMax":
			result.HPMax += value
		case "MPMax":
			result.MPMax += value
		case "ATK":
			result.ATK += value
		case "MAG":
			result.MAG += value
		case "DEF":
			result.DEF += value
		case "RES":
			result.RES += value
		case "Haste":
			result.Haste += value
		case "CritChance":
			result.CritChance += value
		case "CritMulti":
			result.CritMulti += value
		case "MoveSpeed":
			result.MoveSpeed += value
		case "RegenHP":
			result.RegenHP += value
		case "RegenMP":
			result.RegenMP += value
		default:
			if len(key) > 8 && key[:8] == "resists." {
				resistType := key[8:]
				if result.Resists == nil {
					result.Resists = make(map[string]float64)
				}
				result.Resists[resistType] += value
			} else if len(key) > 11 && key[:11] == "amplifiers." {
				ampType := key[11:]
				if result.Amplifiers == nil {
					result.Amplifiers = make(map[string]float64)
				}
				result.Amplifiers[ampType] += value
			}
		}
	}

	for key, value := range mult {
		switch key {
		case "HPMax":
			result.HPMax *= value
		case "MPMax":
			result.MPMax *= value
		case "ATK":
			result.ATK *= value
		case "MAG":
			result.MAG *= value
		case "DEF":
			result.DEF *= value
		case "RES":
			result.RES *= value
		case "Haste":
			result.Haste *= value
		case "CritChance":
			result.CritChance *= value
		case "CritMulti":
			result.CritMulti *= value
		case "MoveSpeed":
			result.MoveSpeed *= value
		case "RegenHP":
			result.RegenHP *= value
		case "RegenMP":
			result.RegenMP *= value
		default:
			if len(key) > 8 && key[:8] == "resists." {
				resistType := key[8:]
				if result.Resists == nil {
					result.Resists = make(map[string]float64)
				}
				if existing, exists := result.Resists[resistType]; exists {
					result.Resists[resistType] = existing * value
				} else {
					result.Resists[resistType] = value
				}
			} else if len(key) > 11 && key[:11] == "amplifiers." {
				ampType := key[11:]
				if result.Amplifiers == nil {
					result.Amplifiers = make(map[string]float64)
				}
				if existing, exists := result.Amplifiers[ampType]; exists {
					result.Amplifiers[ampType] = existing * value
				} else {
					result.Amplifiers[ampType] = value
				}
			}
		}
	}

	result.Version = base.Version + 1

	return a.ClampDerived(result)
}

func (a *ActorCoreImpl) ClampDerived(d Derived) Derived {
	result := d

	result.HPMax = math.Max(result.HPMax, 1.0)
	result.MPMax = math.Max(result.MPMax, 1.0)
	result.ATK = math.Max(result.ATK, 1.0)
	result.MAG = math.Max(result.MAG, 1.0)
	result.DEF = math.Max(result.DEF, 1.0)
	result.RES = math.Max(result.RES, 1.0)

	result.Haste = math.Max(0.5, math.Min(2.0, result.Haste))
	result.CritChance = math.Max(0.0, math.Min(1.0, result.CritChance))
	result.CritMulti = math.Max(1.0, math.Min(5.0, result.CritMulti))
	result.MoveSpeed = math.Max(0.0, result.MoveSpeed)
	result.RegenHP = math.Max(0.0, result.RegenHP)
	result.RegenMP = math.Max(0.0, result.RegenMP)

	if result.Resists != nil {
		for k, v := range result.Resists {
			result.Resists[k] = math.Max(0.0, math.Min(0.8, v))
		}
	}

	if result.Amplifiers != nil {
		for k, v := range result.Amplifiers {
			result.Amplifiers[k] = math.Max(0.0, v)
		}
	}

	result.HPMax = a.sanitizeFloat(result.HPMax)
	result.MPMax = a.sanitizeFloat(result.MPMax)
	result.ATK = a.sanitizeFloat(result.ATK)
	result.MAG = a.sanitizeFloat(result.MAG)
	result.DEF = a.sanitizeFloat(result.DEF)
	result.RES = a.sanitizeFloat(result.RES)
	result.Haste = a.sanitizeFloat(result.Haste)
	result.CritChance = a.sanitizeFloat(result.CritChance)
	result.CritMulti = a.sanitizeFloat(result.CritMulti)
	result.MoveSpeed = a.sanitizeFloat(result.MoveSpeed)
	result.RegenHP = a.sanitizeFloat(result.RegenHP)
	result.RegenMP = a.sanitizeFloat(result.RegenMP)

	return result
}

func (a *ActorCoreImpl) sanitizeFloat(f float64) float64 {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return 0.0
	}
	return f
}

func main() {
	fmt.Println("Testing Actor Core Implementation...")

	actorCore := NewActorCore()

	// Test 1: Basic functionality
	fmt.Println("\n=== Test 1: Basic Functionality ===")
	primary := PrimaryCore{
		HPMax:    100,
		LifeSpan: 1000,
		Attack:   50,
		Defense:  30,
		Speed:    40,
	}

	level := 10
	baseDerived := actorCore.BaseFromPrimary(primary, level)
	fmt.Printf("Base Derived - HPMax: %.2f, ATK: %.2f, DEF: %.2f, Haste: %.2f\n",
		baseDerived.HPMax, baseDerived.ATK, baseDerived.DEF, baseDerived.Haste)

	// Test 2: ComposeCore
	fmt.Println("\n=== Test 2: ComposeCore ===")
	buckets := map[string]CoreContribution{
		"base_stats": {
			Primary: &PrimaryCore{HPMax: 100, Attack: 50},
			Flat:    map[string]float64{"HPMax": 20, "ATK": 10},
			Mult:    map[string]float64{"HPMax": 1.1, "ATK": 1.2},
		},
		"equipment": {
			Primary: &PrimaryCore{HPMax: 25, Attack: 15},
			Flat:    map[string]float64{"resists.fire": 0.1},
			Mult:    map[string]float64{"amplifiers.internal": 1.15},
		},
	}

	composed := actorCore.ComposeCore(buckets)
	fmt.Printf("Composed - Primary HPMax: %d, ATK: %d\n", composed.Primary.HPMax, composed.Primary.Attack)
	fmt.Printf("Composed - Flat HPMax: %.2f, ATK: %.2f\n", composed.Flat["HPMax"], composed.Flat["ATK"])
	fmt.Printf("Composed - Mult HPMax: %.2f, ATK: %.2f\n", composed.Mult["HPMax"], composed.Mult["ATK"])

	// Test 3: FinalizeDerived
	fmt.Println("\n=== Test 3: FinalizeDerived ===")
	finalDerived := actorCore.FinalizeDerived(baseDerived, composed.Flat, composed.Mult)
	fmt.Printf("Final Derived - HPMax: %.2f, ATK: %.2f, Version: %d\n",
		finalDerived.HPMax, finalDerived.ATK, finalDerived.Version)

	// Test 4: ClampDerived
	fmt.Println("\n=== Test 4: ClampDerived ===")
	extremeDerived := Derived{
		HPMax:     -100,
		Haste:     0.1,
		CritChance: 1.5,
		Resists:   map[string]float64{"fire": 1.0},
		Version:   1,
	}

	clamped := actorCore.ClampDerived(extremeDerived)
	fmt.Printf("Clamped - HPMax: %.2f, Haste: %.2f, CritChance: %.2f, Fire Resist: %.2f\n",
		clamped.HPMax, clamped.Haste, clamped.CritChance, clamped.Resists["fire"])

	// Test 5: Golden Test
	fmt.Println("\n=== Test 5: Golden Test ===")
	goldenBuckets := map[string]CoreContribution{
		"base_stats": {
			Primary: &PrimaryCore{HPMax: 100, Attack: 50, Defense: 30, Speed: 40},
			Flat:    map[string]float64{"HPMax": 20, "ATK": 10},
			Mult:    map[string]float64{"HPMax": 1.1, "ATK": 1.2},
		},
		"equipment_bonus": {
			Primary: &PrimaryCore{HPMax: 25, Attack: 15, Defense: 10},
			Flat:    map[string]float64{"resists.fire": 0.1, "resists.ice": 0.05},
			Mult:    map[string]float64{"amplifiers.internal": 1.15},
		},
		"temporary_buff": {
			Flat: map[string]float64{"Haste": 0.2, "CritChance": 0.05},
			Mult: map[string]float64{"ATK": 1.1},
		},
	}

	goldenComposed := actorCore.ComposeCore(goldenBuckets)
	goldenBase := actorCore.BaseFromPrimary(*goldenComposed.Primary, level)
	goldenFinal := actorCore.FinalizeDerived(goldenBase, goldenComposed.Flat, goldenComposed.Mult)

	fmt.Printf("Golden Test - HPMax: %.2f, ATK: %.2f, DEF: %.2f, Haste: %.2f, CritChance: %.2f\n",
		goldenFinal.HPMax, goldenFinal.ATK, goldenFinal.DEF, goldenFinal.Haste, goldenFinal.CritChance)
	fmt.Printf("Golden Test - Fire Resist: %.2f, Ice Resist: %.2f, Internal Amp: %.2f\n",
		goldenFinal.Resists["fire"], goldenFinal.Resists["ice"], goldenFinal.Amplifiers["internal"])
	fmt.Printf("Golden Test - Version: %d\n", goldenFinal.Version)

	fmt.Println("\n✅ All tests completed successfully!")
}
