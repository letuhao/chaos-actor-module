# Cursor Fix Bundle — Runbook (v1.0.0)
Generated: 2025-09-06T20:45:48.450735Z

This bundle tells Cursor exactly how to fix the Actor Core implementation issues you saw:
1) **Bump `Derived.Version`** in `FinalizeDerived`.
2) **Multiply** (not add) for `amplifiers.*` in the multiplicative phase.
3) Include **`LifeSpan`** in `BaseFromPrimary` HP derivation.
4) **Remove `...` ellipses** in source/tests (they break builds).

> Assumption: Go implementation at `packages/actor-core/src/actorcore.go` and tests in `packages/actor-core/tests/`.

## Steps (Cursor executes in order)

### Step 0 — Read the spec
- Read `SPEC.md` from your Actor Core Interface (contracts and invariants).

### Step 1 — Patch `FinalizeDerived`
- Ensure multiplicative phase treats `amplifiers.*` like a product (default 1.0).
- Ensure `result.Version++` occurs **after clamping** and before return.

### Step 2 — Patch `BaseFromPrimary`
- Add `LifeSpan` contribution to HP.
- Keep constants tidy and comment with `// balanceVersion=<x>` if you have versioning.

Example (you can tune constants):
```go
levelFactor := math.Max(1, float64(level))
hpMax := float64(p.HPMax)*levelFactor + 10.0*float64(p.LifeSpan) + 8.0*float64(p.Defense)
// mpMax can scale from hpMax or independently
mpMax := 0.8 * hpMax
```

### Step 3 — Remove ellipses
- Find and delete any literal `...` in `.go` sources/tests that were placeholders.
- Files known to contain them in your zip:
  - `packages/actor-core/src/actorcore.go`
  - `packages/actor-core/tests/actorcore_test.go`

### Step 4 — Tests
- Add/ensure property tests:
  - **Commutativity**: merging buckets is order-independent.
  - **Idempotence**: unchanged buckets → no output change.
  - **Monotonicity**: increasing PrimaryCore never reduces HP/ATK/DEF/Speed outputs.
  - **Clamp bounds**: resists ≤ 0.8, critChance ∈ [0,1], haste ∈ [0.5,2.0].
- Add a **golden test** with fixed buckets + expected snapshot.

### Step 5 — Run
- `go test ./... -race` at repo root.
- If any test fails, Cursor adjusts only within the rules above (no API changes).

---

## Acceptance Criteria (Cursor must satisfy)
- All tests pass with `-race`.
- No `...` remains in sources/tests.
- Multiplicative `amplifiers.*` verified by a unit test (see test template).
- `Derived.Version` increments by exactly 1 each `FinalizeDerived` call.
- `BaseFromPrimary` includes a non-zero `LifeSpan` contribution to `HPMax`.

