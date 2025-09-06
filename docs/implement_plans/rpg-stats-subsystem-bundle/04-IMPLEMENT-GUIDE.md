# Implement Guide (Step-by-Step)

> Use this guide with `00-COLLECTION-CURSOR.md` to drive Cursor.

## Phase 1 — Types
- Implement `model/stats.types.ts`.

## Phase 2 — Stacking Rules
- Implement deterministic order and helpers in `rules/stackingRules.ts`.
- Cover edge cases: multiple overrides, competing caps, group stacks.

## Phase 3 — Registry & Curves
- Implement `registry/index.ts` and `baseCurves.ts` using `08-STAT-REGISTRY-EXAMPLE.md`.

## Phase 4 — Resolver
- Implement `resolver/StatResolver.ts` per algorithm.
- Add `buildHash` via `util/hashing.ts` (stable stringify + SHA-256).

## Phase 5 — Integration Services
- `integration/SnapshotProvider.ts`: DB adapters + resolver + breakdown toggle.
- `integration/ProgressionService.ts`: grantXp, allocatePoints, respec; validation rules.

## Phase 6 — MongoDB Adapters
- Implement adapters for each collection using official Node driver.
- Add index creation from `db/indexes.js` on startup.

## Phase 7 — Testing
- Write unit tests per `05-TEST-GUIDE.md`.
- Add a demo script that prints a resolved snapshot for a sample build.

## Phase 8 — Integration with Core
- Provide example glue: `actor.applyStatSnapshot(snapshot)`.
- Ensure read-only snapshot access inside core.
