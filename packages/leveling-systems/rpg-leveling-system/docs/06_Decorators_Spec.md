# 06 — Decorators Spec & Pipeline

## XP Pipeline
1) **PRE** → 2) **BASE** → 3) **POST** → 4) **FINALIZE**
- Sort deterministic: `(phase, -priority, id)`.

## Resource Decorators
- Magic: tính `mana_max` từ `intelligence`, `willpower`.
- Luyện Thể: tính `stamina_max` từ `endurance`, `strength`.
- Baseline (optional): `hp_max`, `mana_max`, `stamina_max` linear từ stats (bật/tắt bằng config).

## Damage Type Decorators
- `registerTypes()` thêm kiểu damage.
- `resistFromStats` sinh **resist_*** contributions (clamp bằng `HARD_MAX`).
- `damageCompute` thêm biến thiên damage vào `damage_out/damage_in` context.
