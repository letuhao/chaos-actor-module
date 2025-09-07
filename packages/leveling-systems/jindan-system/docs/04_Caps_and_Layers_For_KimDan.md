# 04 — Caps & Layers (Kim Đan ↔ Actor Core)

- **scope=REALM**: tất cả **caps theo cảnh giới/tiểu cảnh giới**.
- **scope=WORLD**: thiên địa pháp tắc (ví dụ giới hạn `divine_sense` trong “Phong Thần cấm vực”).
- **scope=EVENT**: độ kiếp, bí cảnh, trận pháp (tạm thời).
- **scope=TOTAL**: hạn trần của hệ **Luyện Khí** (để không vượt game thiết kế).

**Across-layer policy** mặc định: `INTERSECT`.  
Ví dụ: REALM cho `qi_purity ≤ 0.85`, WORLD bắt `≤ 0.80` ⇒ **cuối cùng 0.80**.

**CapContribution mẫu**:
```json
{ "system":"kim_dan", "dimension":"qi_purity", "mode":"HARD_MAX", "kind":"max",
  "value":0.85, "scope":"REALM", "realm":"Kim Đan: Trung Phẩm" }
```
