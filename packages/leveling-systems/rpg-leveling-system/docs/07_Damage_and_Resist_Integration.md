# 07 — Damage & Resistance Integration

## Dimensions
- `*_resist` (e.g., `physical_resist`, `fire_resist`, ...). Range `[0..RESIST_CAP_MAX]`.

## Mapping gợi ý
- `physical_resist` ← `endurance` (+), `agility` (+small)
- `magic_resist` ← `willpower` (+), `intelligence` (+small)
- `poison_resist` ← `endurance` (+), `willpower` (+small)

## Context hooks
- Dùng channels: `damage_out`, `damage_in` với `{ additive_percent, multipliers[], post_add }`.
- Resist áp vào **damage_in** multiplier: `mult *= (1 - physical_resist)` sau khi combine decorators.
