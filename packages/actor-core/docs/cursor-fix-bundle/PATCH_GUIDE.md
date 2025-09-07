# Patch Guide (Go) — Exact Patterns

> If your paths differ, adjust them, but keep the logic.

## 1) FinalizeDerived: Amplifiers multiply, not add

Find the multiplicative phase in `packages/actor-core/src/actorcore.go`. Replace the `amplifiers.*` handling from `+=` to `*=` with identity default 1.0.

**Before (buggy):**
```go
} else if strings.HasPrefix(key, "amplifiers.") {
    at := strings.TrimPrefix(key, "amplifiers.")
    if result.Amplifiers == nil { result.Amplifiers = map[string]float64{} }
    result.Amplifiers[at] += value
}
```

**After (fixed):**
```go
} else if strings.HasPrefix(key, "amplifiers.") {
    at := strings.TrimPrefix(key, "amplifiers.")
    if result.Amplifiers == nil { result.Amplifiers = map[string]float64{} }
    if existing, ok := result.Amplifiers[at]; ok {
        result.Amplifiers[at] = existing * value
    } else {
        result.Amplifiers[at] = value // identity(1.0) * value
    }
}
```

## 2) FinalizeDerived: bump Version

At the end of `FinalizeDerived`, after clamping:

```go
result = a.ClampDerived(result)
result.Version++
return result
```

## 3) BaseFromPrimary: include LifeSpan

In `BaseFromPrimary`:

```go
levelFactor := math.Max(1, float64(level))
hpMax := float64(p.HPMax)*levelFactor + 10.0*float64(p.LifeSpan) + 8.0*float64(p.Defense)
mpMax := 0.8 * hpMax
// ... keep your ATK/DEF/MAG/RES formulas consistent
```

> Import `"math"` & `"strings"` at top if not present.

## 4) Remove ellipses

Search for literal `...` and delete them:
- `packages/actor-core/src/actorcore.go`
- `packages/actor-core/tests/actorcore_test.go`

