# LANGUAGE-GUIDE (Polyglot)

This subsystem is specified in **language-agnostic** terms. Below are reference mappings for **Go (default)** and **TypeScript**.

---

## Core Concepts → Go

- `StatKey`: `type StatKey string` with `const` keys: `STR, INT, WIL, AGI, SPD, END, PER, LUK`, etc.
- `ModifierStacking`: `type ModifierStacking string` with consts: `ADD_FLAT`, `ADD_PCT`, `MULTIPLY`, `OVERRIDE`, `CAP_MAX`, `CAP_MIN`.
- `StatModifier`: struct with fields (`Key StatKey`, `Op ModifierStacking`, `Value float64`, `Source ModifierSourceRef`, `Conditions *ModifierConditions`, `Priority int`).
- `StatSnapshot`: struct: `ActorID string`, `Stats map[StatKey]StatValue`, `Version int`, `Ts int64`.
- `Resolver` interface:
  ```go
  type Resolver interface {
      ComputeSnapshot(in ComputeInput) StatSnapshot
  }
  ```
- `SnapshotProvider`: loads Mongo docs, builds inputs, calls `Resolver`.
- `ProgressionService`: `GrantXP`, `AllocatePoints`, `Respec`.

### Go Project Layout
```
rpg-stats/
  go.mod
  internal/
    model/               // types
    rules/               // stacking rules
    registry/            // stat defs & curves
    resolver/            // Resolve implementation
    integration/         // Mongo adapters, provider, progression
    util/                // hashing/time
  cmd/demo/              // small demo to print snapshot
  test/                  // table-driven tests
```

### Go Tech Choices
- **MongoDB**: `go.mongodb.org/mongo-driver/mongo`
- **Tests**: `testing` package (table-driven), `t.Run` subtests.
- **Hashing**: `crypto/sha256` + stable JSON (e.g., `encoding/json` with sorted keys via struct ordering / custom canonicalizer).

---

## Core Concepts → TypeScript (Node)
[Same as before; see existing TS guidance in the bundle.]

---

## Portability Notes
- Keep resolver **pure** (no I/O).
- Time & RNG should be injected or mocked.
- Avoid language-specific floating rounding quirks; centralize rounding helpers.
