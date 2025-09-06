# Stacking Rules (Deterministic)

Order per stat:
1) Base (species/class/level + allocations)
2) Additive Flat (items → titles → passives)
3) Additive Percent (sum, then apply once)
4) Multiplicative (Buffs → Debuffs → Aura/Env)
5) Override (priority desc, value desc)
6) Caps (min/max, soft DR)
7) Rounding

## Group Stacking (`stackId`)
- Same `stackId` stacks up to `maxStacks` (default unlimited).
- ADD_FLAT/PCT: sum; MULTIPLY: multiply each; OVERRIDE: highest priority wins.
- CAP_MAX: pick smallest; CAP_MIN: pick largest.

## Examples
- Two items each +10 ATK flat → total +20 before %.
- Buff +15% ATK and Debuff -10% ATK → `(base+flat)*(1.15)*(0.90)`.
- Override MOVE_SPEED=100 (prio 10) beats prio 5 for duration.
