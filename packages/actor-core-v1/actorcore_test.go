package actorcore

import (
"math"
"testing"
)

func TestAmplifiersMultiply(t *testing.T) {
actorCore := NewActorCore()

base := Derived{
ATK:        100,
Amplifiers: map[string]float64{"internal": 1.0},
}

flat := map[string]float64{}
mult := map[string]float64{"amplifiers.internal": 1.10}

result := actorCore.FinalizeDerived(base, flat, mult)

if math.Abs(result.Amplifiers["internal"]-1.10) > 1e-9 {
t.Fatalf("expected 1.10, got %v", result.Amplifiers["internal"])
}
}

func TestBaseFromPrimaryLifeSpanAffectsHP(t *testing.T) {
actorCore := NewActorCore()

p0 := PrimaryCore{HPMax: 0, LifeSpan: 0, Attack: 0, Defense: 0, Speed: 0}
p1 := PrimaryCore{HPMax: 0, LifeSpan: 10, Attack: 0, Defense: 0, Speed: 0}

d0 := actorCore.BaseFromPrimary(p0, 1).HPMax
d1 := actorCore.BaseFromPrimary(p1, 1).HPMax

if !(d1 > d0) {
t.Fatalf("LifeSpan should increase HPMax: %v !> %v", d1, d0)
}
}

func TestFinalizeDerived_VersionBump(t *testing.T) {
actorCore := NewActorCore()

base := Derived{
HPMax:   100,
MPMax:   80,
ATK:     50,
Version: 5,
}

flat := map[string]float64{"HPMax": 20}
mult := map[string]float64{"ATK": 1.2}

result := actorCore.FinalizeDerived(base, flat, mult)

if result.Version != base.Version+1 {
t.Errorf("Version not incremented properly: expected %d, got %d", base.Version+1, result.Version)
}
}
