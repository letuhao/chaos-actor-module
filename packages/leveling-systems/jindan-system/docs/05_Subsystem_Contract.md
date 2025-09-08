# 05 — Hợp đồng dữ liệu (SubsystemOutput)

## Primary (emit theo phân bổ/tu luyện)
- Tất cả chỉ số trong **components/** và **dao-system/** (FLAT/MULT tuỳ pháp môn).

## Derived
- `qi_max`, `qi_regen`, `spell_power`, `divine_sense`, `channel_speed`,
  `mental_resist`, `tribulation_resist`, `alchemy_success`… (từ các file components).

## Caps
- Theo **REALM** cho primary & derived cốt lõi (`qi_purity`, `dantian_*`, `meridian_*`, `shen_*`).  
- **EVENT** cho độ kiếp (`tribulation_resist HARD_MIN`, …).  
- **WORLD** khi cần ràng buộc thế giới (pháp cấm, bí cảnh).  
- **TOTAL** khi muốn khóa trần toàn hệ.

## Meta/Context
- `meta.system = "kim_dan"`  
- `context.kim_dan`: có thể xuất `modifier` dùng cho damage/heal/formation nếu hệ khác cần.
