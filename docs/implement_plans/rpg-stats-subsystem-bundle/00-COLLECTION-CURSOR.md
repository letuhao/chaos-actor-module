# Collection — Cursor Control Script

> Set `IMPLEMENT_LANG` to choose language. Default is `go`.
## 1) Select implementation language
- Tell Cursor: `IMPLEMENT_LANG=go` (default) or `IMPLEMENT_LANG=ts`.
- Cursor should remember this choice for all subsequent steps.


> **Purpose**: Drive Cursor to read docs, generate code file-by-file, and run tests for the RPG Stats Sub-System.
> You can paste commands from each step directly into Cursor's chat/terminal.

---

## 0) Read the docs (ONE TIME)
Read these files in the repo (use Cursor's "Read file" or paste paths):
- `01-BASIC-DESIGN.md`
- `02-DETAIL-DESIGN.md`
- `03-DATABASE-DESIGN-MONGODB.md`
- `06-API-CONTRACTS.md`
- `07-STACKING-RULES.md`
- `08-STAT-REGISTRY-EXAMPLE.md`

Tell Cursor: *“Load and memorize the above docs as the source of truth for the RPG Stats Sub-System.”*

---

## 1) Create package skeleton
**Command to Cursor:**

Create a new workspace package (monorepo-friendly):

```
packages/rpg-stats/
  src/
    model/
      stats.types.ts
    registry/
      index.ts
      baseCurves.ts
    rules/
      stackingRules.ts
    resolver/
      StatResolver.ts
    integration/
      SnapshotProvider.ts
      ProgressionService.ts
    util/
      hashing.ts
      time.ts
  tests/
    resolver.spec.ts
    stacking.spec.ts
    registry.spec.ts
  package.json
  tsconfig.json
  README.md
```

- Configure `tsconfig` for strict mode.
- Target Node 18+/ES2020.
- Export `computeSnapshot`, `StatSnapshot`, `StatDef`, `StatModifier`, and services.

---

## 2) Implement model types
**Command to Cursor:** Create `src/model/stats.types.ts` per `02-DETAIL-DESIGN.md` and `06-API-CONTRACTS.md`:
- `StatKey` (string union) — start with the 8 primaries in `08-STAT-REGISTRY-EXAMPLE.md` (STR, INT, WIL, AGI, SPD, END, PER, LUK).
- `StatValue`, `StatBreakdown`, `StatSnapshot`.
- `ModifierStacking`, `StatModifier`, `ModifierSourceRef`.
- `StatCategory`, `StatDef`, `ResolveContext`.
- `LevelProgression`, `PlayerProgress` (interface for service return values).

---

## 3) Implement stacking rules
**Command to Cursor:** Create `src/rules/stackingRules.ts`:
- Deterministic order: Base → ADD_FLAT → ADD_PCT → MULTIPLY (Buffs → Debuffs → Aura/Env) → OVERRIDE → Caps → Rounding.
- Group stacking by `stackId` and `maxStacks`.
- Tie-break for OVERRIDE: higher `priority`, then higher `value`.
- Helpers: `applyAddFlat`, `applyAddPct`, `applyMultiplyTiers`, `applyOverrides`, `applyCaps`, `applyRounding`.

Write unit tests in `tests/stacking.spec.ts` using cases from `05-TEST-GUIDE.md`.

---

## 4) Implement registry & base curves
**Command to Cursor:** Create `src/registry/index.ts` and `src/registry/baseCurves.ts`:
- Registry loader holding `StatDef[]` and `BaseCurve` fns.
- Implement formulas from `08-STAT-REGISTRY-EXAMPLE.md`.
- Provide typed helpers: `getStatDef`, `resolveDerivedStat`.

Write unit tests in `tests/registry.spec.ts` with fixtures.

---

## 5) Implement resolver
**Command to Cursor:** Create `src/resolver/StatResolver.ts`:
- `computeSnapshot(input)` implementing the algorithm in `02-DETAIL-DESIGN.md`.
- Optional `breakdown` building for tooltips.
- Deterministic hashing of build inputs for cache key (use `util/hashing.ts`).

Write unit tests in `tests/resolver.spec.ts`.

---

## 6) Implement integration services
**Command to Cursor:** Create `src/integration/SnapshotProvider.ts` and `src/integration/ProgressionService.ts`:
- SnapshotProvider loads player progress, equipment, titles, effects, registry + curves (from DB/content service), then calls resolver.
- ProgressionService: `grantXp`, `allocatePoints`, `respec` (in-memory first; DB adapters added later).

---

## 7) Wire MongoDB adapters
**Command to Cursor:** Create DB adapters per `03-DATABASE-DESIGN-MONGODB.md`:
- Collections: `player_progress`, `player_effects_active`, `player_equipment`, `titles_owned`, `content_stat_registry`.
- Indexes from `db/indexes.js`.
- Implement CRUD in adapters; keep a clean interface boundary so the core remains DB-agnostic.

---

## 8) Testing
**Command to Cursor:** Run unit tests and add any missing edge cases from `05-TEST-GUIDE.md`.
- Use Jest/Vitest.
- Mock time and use seeded RNG for deterministic tests.

---

## 9) Integration with Core Actor
**Command to Cursor:** Provide a sample integration snippet:
- Given an `actorId`, call `SnapshotProvider.buildForActor(actorId)`.
- Feed the `StatSnapshot` into your Core (`actor.applyStatSnapshot(snapshot)`).
- Demonstrate reacting to changes (equip, buff, level up) → new snapshot.

---

## 10) Deliverables check
- All tests green.
- Lint passes (ESLint + Prettier).
- Public API exported in `packages/rpg-stats`.
- Small demo script that prints a resolved `StatSnapshot` for a sample actor.

## 2A) Create package skeleton (Go)
**Command to Cursor (Go):**

Create the following structure and files:

```
packages/rpg-stats-go/
  go.mod (module github.com/yourorg/rpg-stats-go; go 1.22)
  internal/model/types.go
  internal/rules/stacking.go
  internal/registry/registry.go
  internal/registry/curves.go
  internal/resolver/resolver.go
  internal/integration/provider.go
  internal/integration/progression.go
  internal/integration/mongo_adapters.go
  internal/util/hash.go
  cmd/demo/main.go
  test/stacking_test.go
  test/resolver_test.go
  test/registry_test.go
  README.md
```

Fill files per specs in `02-DETAIL-DESIGN.md`, `06-API-CONTRACTS.md`, and `09-LANGUAGE-GUIDE.md`.

### Go Type Stubs (paste into files)
- `internal/model/types.go` should define:
  - `type StatKey string` and the 8 primary keys.
  - `type ModifierStacking string` and ops.
  - `type ModifierSourceRef struct { Kind string; ID string; Label string }`
  - `type ModifierConditions struct { RequiresTagsAll, RequiresTagsAny, ForbidTags []string; DurationMs int64; StackID string; MaxStacks int }`
  - `type StatModifier struct { Key StatKey; Op ModifierStacking; Value float64; Source ModifierSourceRef; Conditions *ModifierConditions; Priority int }`
  - `type StatBreakdown struct { Base, AdditiveFlat, AdditivePct, Multiplicative float64; Overrides []OverrideEntry; CappedTo *float64; Notes []string }`
  - `type StatValue struct { Key StatKey; Value float64; Breakdown *StatBreakdown }`
  - `type StatSnapshot struct { ActorID string; Stats map[StatKey]StatValue; Version int; Ts int64 }`
  - `type StatDef struct { Key StatKey; Category string; DisplayName string; Description string; Rounding string }`
  - `type ComputeInput struct { ActorID string; Level int; BaseAllocations map[StatKey]int; Registry []StatDef; Items, Titles, Passives, Buffs, Debuffs, Auras, Environment []StatModifier; WithBreakdown bool }`

- `internal/rules/stacking.go`: implement deterministic order and helpers.

- `internal/resolver/resolver.go`: implement `type Resolver interface { ComputeSnapshot(ComputeInput) StatSnapshot }` and the algorithm.

- `internal/integration/mongo_adapters.go`: collections from `03-DATABASE-DESIGN-MONGODB.md` using `mongo-driver`.

- `cmd/demo/main.go`: load fixtures (JSON), call resolver, print snapshot.

**Tests:** Use table-driven tests in `test/*.go` per `05-TEST-GUIDE.md`.

---

## 2B) Create package skeleton (TypeScript)
(Identical to the earlier TS scaffold; use `packages/rpg-stats/` with the files listed previously in this document.)

---
