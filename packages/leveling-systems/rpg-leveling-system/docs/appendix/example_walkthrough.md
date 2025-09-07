# Example Walkthrough

- Level 10, exponential (base 100, growth 1.12) → `XP_to_next ≈ 311`.
- Kill elite L12 → after decorators → gain ~156 XP.
- Level up: +5 points (from [3..7]); allocate +3 strength, +2 endurance.
- Contributions:
  - Primary: `strength +3`, `endurance +2` (FLAT).
  - Derived (baseline): `hp_max += 5*2 + 2*3 = 16`.
  - Caps: `lifespan_years ADDITIVE max += 0.5`.
- Resistances: endurance 12 ⇒ `physical_resist +0.12` (clamp 0.85).
