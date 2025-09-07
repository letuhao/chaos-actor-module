# 02 — Derived Stats & Công thức gợi ý

> Dùng Actor Core pipeline: FLAT/MULT/POST_ADD, sau đó **clamp theo cap**.

- `qi_max`         ← `dantian_capacity * (1 + dantian_compression)`
- `qi_regen`       ← `meridian_conductivity * f1(qi_purity)` (ví dụ `f1(p)=1+p`)
- `spell_power`    ← `qi_max^α * qi_purity^β * eff_qi_to_shen^γ` (α,β,γ cấu hình)
- `channel_speed`  ← `meridian_conductivity * g1(shen_control)` (tốc độ dẫn pháp)
- `divine_sense`   ← `shen_depth * h1(shen_clarity)` (tầm thần thức)
- `mental_resist`  ← `shen_clarity + k1*shen_depth`
- `tribulation_resist` ← `dantian_stability + meridian_toughness + c1*qi_purity`
- `alchemy_success`   ← `dantian_stability + qi_purity + dao_comprehension` (nếu có chế đan)

Các hệ số `(α..)` đưa vào registry theo **dimension** để dễ tinh chỉnh.
