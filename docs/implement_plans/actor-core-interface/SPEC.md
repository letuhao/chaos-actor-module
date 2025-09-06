# Actor Core Interface — SPEC (v1.0.0)

## Scope
This is the **highest-level contract** for the Actor Core. Implementations must keep all functions **pure** (no IO, no globals).
Combat and other systems depend on these signatures and invariants.

## Core Types (language-agnostic)
**PrimaryCore**
- HPMax: int
- LifeSpan: int
- Attack: int
- Defense: int
- Speed: int

**Derived**
- HPMax, MPMax, ATK, MAG, DEF, RES: float
- Haste, CritChance, CritMulti, MoveSpeed, RegenHP, RegenMP: float
- Resists: map<string,float>
- Amplifiers: map<string,float>
- Version: uint64 (bump on recompute)

**CoreContribution**
- Primary?: PrimaryCore (additive)
- Flat?: map<string,float>  (e.g., "HPMax": +80, "resists.fire": +0.05)
- Mult?: map<string,float>  (e.g., "ATK": 1.10, "amplifiers.internal": 1.15)
- Tags?: string[] (capabilities)

**ActorResources**
- map<string, ResourceState{current,max,regen,epoch}>

## Composition Law (frozen)
```
Derived_final = clamp( ( baseFromPrimary( Σ PrimaryCore )
                       + Σ flat )
                       × Π mult )
```
- Merge buckets in **lexicographic order of their keys** for determinism.
- Multiplicative maps default to **1.0** when absent.

## Clamps (initial)
- resists.* ∈ [0, 0.8]
- critChance ∈ [0, 1]
- haste ∈ [0.5, 2.0]
- hp/mp ≥ 1

## Function Contracts
- ComposeCore(buckets: map<string, CoreContribution>) -> CoreContribution
  - Deterministic; stable sort by key; add Primary & Flat; multiply Mult (identity 1.0).

- BaseFromPrimary(p: PrimaryCore, level: int) -> Derived
  - Pure; defines balance curve; may read **constants** from module scope but not external IO.

- FinalizeDerived(base: Derived, flat: map<string,float>, mult: map<string,float>) -> Derived
  - Applies flats then mults; calls ClampDerived; bumps Version (+1).

- ClampDerived(d: Derived) -> Derived
  - Applies bounds above; never produces NaN/Inf.

## Invariants / Laws
- Determinism: same inputs → same Derived; independent of map iteration order.
- Idempotence: composing with unchanged buckets does not change results.
- Monotonicity: increasing any PrimaryCore field never reduces HPMax/ATK/DEF/Speed outputs from BaseFromPrimary.
- Purity: no time/IO/randomness.

## Versioning
Embed `schemaVersion`, `balanceVersion`, `seedVersion` in outer snapshots (out of scope for this bundle).

## Testing Guidance (for implementers)
- Property tests for commutativity (with stable sort), idempotence, monotonicity, clamp ranges.
- Golden test: fixed buckets produce the same Derived across runs.
