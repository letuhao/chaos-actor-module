# Stat Registry — Example (8 Primary Stats)

## Primary (Morrowind-style)
- STR — Strength
- INT — Intelligence
- WIL — Willpower
- AGI — Agility
- SPD — Speed
- END — Endurance
- PER — Personality
- LUK — Luck

## Derived & Secondary (examples)
- HP_MAX, MANA_MAX, ATK, MATK, DEF, MDEF
- CRIT_CHANCE, CRIT_DAMAGE, HASTE, ACCURACY, EVASION
- LIFESTEAL, BLOCK_CHANCE, PARRY_CHANCE
- MOVE_SPEED, CAST_SPEED
- FIRE_RES, ICE_RES, LIGHTNING_RES, POISON_RES
- STAMINA_MAX, ENERGY_REGEN, WEIGHT_CAP
- PERSUASION, MERCHANT_PRICE_MOD, FACTION_REPUTATION_GAIN

## Example Formulas (pseudo)
- `HP_MAX = END * 20 + STR * 2 + curve.baseHp(level, classId)`
- `MANA_MAX = INT * 15 + WIL * 5 + curve.baseMana(level, classId)`
- `ATK = STR * 2 + AGI * 0.3 + weaponBase`
- `MATK = INT * 2 + WIL * 0.4 + staffBase`
- `CRIT_CHANCE = min(1.00, 0.01 + LUK * 0.003 + AGI * 0.001)`
- `CRIT_DAMAGE = 1.5 + LUK * 0.002`
- `MOVE_SPEED = clamp(100 + SPD * 0.8 + AGI * 0.2, 50, 220)` (% scale)
- `CAST_SPEED = clamp(100 + WIL * 0.6 + INT * 0.2, 50, 220)`
- `PERSUASION = PER * 2 + LUK * 0.5`
- Resistances capped at 0–0.75 (75%).

## Rounding
- Resources/caps can be integers; percentages kept as floats then formatted in UI.
