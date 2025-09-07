// packages/actor-core/go/actorcore/actorcore.go
// Interfaces only; pure function contracts for the Actor Core.
package actorcore

type PrimaryCore struct {
	HPMax    int
	LifeSpan int
	Attack   int
	Defense  int
	Speed    int
}

type Derived struct {
	HPMax, MPMax          float64
	ATK, MAG              float64
	DEF, RES              float64
	Haste                 float64
	CritChance, CritMulti float64
	MoveSpeed             float64
	RegenHP, RegenMP      float64
	Resists               map[string]float64
	Amplifiers            map[string]float64
	Version               uint64
}

type CoreContribution struct {
	Primary *PrimaryCore
	Flat    map[string]float64
	Mult    map[string]float64
	Tags    []string
}

type ResourceState struct {
	Current, Max, Regen float64
	Epoch               uint64
}
type ActorResources map[string]*ResourceState

type ActorCore interface {
	ComposeCore(buckets map[string]CoreContribution) CoreContribution
	BaseFromPrimary(p PrimaryCore, level int) Derived
	FinalizeDerived(base Derived, flat map[string]float64, mult map[string]float64) Derived
	ClampDerived(d Derived) Derived
}

// Invariants:
// - ComposeCore merges keys in lexicographic order; Flat sums; Mult multiplies (default 1.0).
// - FinalizeDerived applies Flat, then Mult, then ClampDerived, and bumps Version by 1.
