# Basic Design

## Goals
- Keep **Core Actor** pure and DB-agnostic.
- RPG Sub-System manages **level/XP**, **allocatable points**, **equipment/titles/effects**, and **stat resolution**.
- Emit a deterministic **StatSnapshot** for combat and UI.
- Make stats **extensible** and **stacking predictable**.

## Scope
- In: Stat registry, progression, modifier ingestion, resolver, MongoDB persistence, unit tests.
- Out: Rendering, combat math itself (handled by Core using the snapshot), authentication.

## Architecture (modules)
- **model/**: TS types.
- **registry/**: Stat definitions & derived formulas.
- **rules/**: Stacking & caps.
- **resolver/**: Snapshot computation.
- **integration/**: SnapshotProvider + ProgressionService + DB adapters.
- **util/**: hashing/time helpers.

## Key Flows
1. Load player data + content.
2. Aggregate modifiers by source (items, titles, buffs/debuffs, auras, env).
3. Resolve stats → snapshot.
4. Push snapshot to Core.
5. On changes, recompute and push again.


**Primary Stats (8):** STR, INT, WIL, AGI, SPD, END, PER, LUK.
