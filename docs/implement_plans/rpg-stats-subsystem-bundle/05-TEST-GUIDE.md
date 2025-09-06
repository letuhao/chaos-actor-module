# Test Guide

## Framework
- Use Vitest or Jest.
- 100% coverage on `rules` and `resolver`; >=80% overall.

## Unit Tests
1. **Add Flat**: stacking sums, order independence within the same tier.
2. **Add %**: sum before applying; verify against base+flat.
3. **Multiply Tiers**: Buffs then Debuffs then Aura/Env (multiplicative chain).
4. **Override**: higher priority wins; tie → greater value.
5. **Caps**: min/max; soft DR function applied after overrides.
6. **Rounding**: per-stat rounding applied last.
7. **Derived**: changing STR alters ATK and HP as per formulas.
8. **Stacks**: `stackId` + `maxStacks` behavior.

## Integration Tests
- Full snapshot build with: allocations + equipment + title + buff + debuff + aura + env.
- Expiring effects (simulate `expiresAt`).
- Respec path: allocations reset; snapshot updates.
- Large registry: ensure performance and determinism.

## Determinism
- Use fixed fixtures and no RNG in resolver.
- Assert `buildHash` changes only when inputs change.


**Tip:** Include test fixtures that exercise all 8 primary stats influencing derived values.
