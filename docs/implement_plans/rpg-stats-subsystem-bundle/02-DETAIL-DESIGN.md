# Detail Design

## Types
See `src/model/stats.types.ts`:
- `StatKey` (now includes STR, INT, WIL, AGI, SPD, END, PER, LUK), `StatValue`, `StatBreakdown`, `StatSnapshot`
- `ModifierStacking`, `StatModifier`, `ModifierSourceRef`
- `StatCategory`, `StatDef`, `ResolveContext`
- `LevelProgression`, `PlayerProgress`

## Resolver Algorithm
1) Seed base via registry (species/class curves) + player allocations.  
2) Apply **ADD_FLAT** (equipment → titles → passives).  
3) Apply **ADD_PCT** (sum separately, then apply once).  
4) Apply **MULTIPLY** in tiers: Buffs → Debuffs → Aura/Env.  
5) Apply **OVERRIDE** by priority → value.  
6) Apply **CAPS** and soft DR functions.  
7) Apply **rounding** rules.  
8) Build `StatBreakdown` if requested.

## Extensibility
- Add new `StatKey` in registry & formulas.
- Add new source kinds (e.g., seasonal aura) without core changes.
- Pluggable curves for classes/species and balance patches via content versioning.

## Determinism
- Stable stacking order.
- Seeded RNG where needed (avoid in resolver).
- Hash input to cache snapshot by `buildHash` (equipment ids/rolls, effects, titles, level, allocations, env key).
