// Actor Core v3 — Interfaces (skeleton)
package core

import "context"

type Bucket string
const (
  BucketFlat Bucket = "FLAT"
  BucketMult Bucket = "MULT"
  BucketPostAdd Bucket = "POST_ADD"
  BucketOverride Bucket = "OVERRIDE"
)
type CapMode string
const (
  CapBaseline CapMode = "BASELINE"
  CapAdd CapMode = "ADDITIVE"
  CapOverride CapMode = "OVERRIDE"
  CapHardMax CapMode = "HARD_MAX"
  CapHardMin CapMode = "HARD_MIN"
)

type Contribution struct { Dimension string; Bucket Bucket; Value float64; System string; Priority int }
type CapContribution struct { System, Dimension string; Mode CapMode; Kind string; Value float64; Priority int; Scope string; Realm string; Tags []string }
type SubsystemMeta struct { System string; Stage string; Version int64 }
type ModifierPack struct{ AdditivePercent float64; Multipliers []float64; PostAdd float64 }

type SubsystemOutput struct {
  Primary []Contribution
  Derived []Contribution
  Caps    []CapContribution
  Context map[string]ModifierPack
  Meta    SubsystemMeta
}
type ClampSpec struct{ Min, Max float64 }
type EffectiveCaps map[string]ClampSpec
type MergeRule struct{ UsePipeline bool; Operator string; ClampDefault ClampSpec }
type Snapshot struct{ Primary, Derived map[string]float64; CapsUsed EffectiveCaps; Version int64 }

type Subsystem interface {
  SystemID() string
  Priority() int
  Contribute(ctx context.Context, actor *Actor) (SubsystemOutput, error)
}
type CombinerRegistry interface{ RuleFor(dim string) MergeRule }
type CapsProvider interface{
  EffectiveCapsWithinLayer(ctx context.Context, actor *Actor, outs []SubsystemOutput, layer string) (EffectiveCaps, error)
  ReduceAcrossLayers(ctx context.Context, layerCaps map[string]EffectiveCaps, combiner CombinerRegistry, order []string, acrossPolicy string) (EffectiveCaps, error)
}

type Actor struct {
  GUID, Name string
  Race interface{ ID()string; Name()string }
  LifeSpan, Age float64
  Version int64
  Subsystems []Subsystem
}
type Aggregator interface{
  Resolve(ctx context.Context, actor *Actor, combiner CombinerRegistry, caps CapsProvider, order []string, acrossPolicy string) (Snapshot, error)
}
