# Test Templates (Go)

Add these into `packages/actor-core/tests/actorcore_test.go` or split files.

## 1) Multiplicative amplifiers test
```go
func TestAmplifiersMultiply(t *testing.T) {
    base := Derived{ ATK: 100, Amplifiers: map[string]float64{"internal": 1.0} }
    flat := map[string]float64{}
    mult := map[string]float64{"amplifiers.internal": 1.10}
    ac := NewActorCore() // or your constructor
    got := ac.FinalizeDerived(base, flat, mult)
    if math.Abs(got.Amplifiers["internal"]-1.10) > 1e-9 {
        t.Fatalf("expected 1.10, got %v", got.Amplifiers["internal"])
    }
}
```

## 2) Version bump test
```go
func TestFinalizeBumpsVersion(t *testing.T) {
    base := Derived{ Version: 7 }
    ac := NewActorCore()
    got := ac.FinalizeDerived(base, nil, nil)
    if got.Version != 8 {
        t.Fatalf("expected version 8, got %d", got.Version)
    }
}
```

## 3) LifeSpan influence test
```go
func TestBaseFromPrimaryLifeSpanAffectsHP(t *testing.T) {
    ac := NewActorCore()
    p0 := PrimaryCore{HPMax: 0, LifeSpan: 0, Attack:0, Defense:0, Speed:0}
    p1 := PrimaryCore{HPMax: 0, LifeSpan: 10, Attack:0, Defense:0, Speed:0}
    d0 := ac.BaseFromPrimary(p0, 1).HPMax
    d1 := ac.BaseFromPrimary(p1, 1).HPMax
    if !(d1 > d0) {
        t.Fatalf("LifeSpan should increase HPMax: %v !> %v", d1, d0)
    }
}
```

## 4) Commutativity (stable lexicographic merge)
```go
func TestComposeCoreOrderIndependent(t *testing.T) {
    b1 := CoreContribution{ Flat: map[string]float64{"ATK": 10}, Mult: map[string]float64{"ATK": 1.2} }
    b2 := CoreContribution{ Flat: map[string]float64{"ATK": 5},  Mult: map[string]float64{"ATK": 1.1} }
    ac := NewActorCore()
    x := ac.ComposeCore(map[string]CoreContribution{"a": b1, "b": b2})
    y := ac.ComposeCore(map[string]CoreContribution{"b": b2, "a": b1})
    if !reflect.DeepEqual(x, y) {
        t.Fatalf("compose should be order independent")
    }
}
```

## 5) Clamp bounds
```go
func TestClamps(t *testing.T) {
    ac := NewActorCore()
    d := Derived{ CritChance: 2.0, Haste: 10.0, Resists: map[string]float64{"fire": 1.2} }
    d2 := ac.ClampDerived(d)
    if d2.CritChance > 1.0 || d2.CritChance < 0.0 { t.Fatal("crit clamp failed") }
    if d2.Haste < 0.5 || d2.Haste > 2.0 { t.Fatal("haste clamp failed") }
    if d2.Resists["fire"] > 0.8 { t.Fatal("resist clamp failed") }
}
```
