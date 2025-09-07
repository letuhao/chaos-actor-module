# 14-CROSS-SYSTEM-MAPPING — Mapping Kim Đan vs RPG tuyến tính
Generated: 2025-09-06T22:20:20.193280Z

Mục tiêu: so sánh **hai hệ thống khác nhau** dùng **cùng lượng XP tích lũy**:
- **Kim Đan (Xianxia)**: **exponential realm scaling** (Gap(R) = A·B^r · multipliers), buff min/max theo realm.
- **RPG tuyến tính**: Primary/Derived tăng **tuyến tính theo level** (ví dụ: mỗi level +k điểm).

Kết luận mong muốn: với **cùng XP**, nhân vật Kim Đan đạt **sức mạnh vượt trội** do bản chất tăng **cấp số mũ**; nhân vật RPG cần **level rất cao** mới tiệm cận.

---

## 1) Giả định XP (mỗi hệ thống có đường cong riêng)
- **Kim Đan XP**: chia làm **LevelXP** (tiểu cấp) + **RealmXP** (đột phá). Ví dụ:
  - Level n trong cùng realm: `XP_KD_Level(n) = L0 · n^α` (α≈1.5)
  - Breakthrough realm r: `XP_KD_Realm(r) = R0 · γ^r` (γ≈6..12)

- **RPG XP**: `XP_RPG_Level(n) = L0_RPG · n^β` (β≈1.2..1.6)

> Hai hệ XP **không cần trùng**; ta chỉ **so sánh cùng tổng XP tích lũy**.

---

## 2) Hàm sức mạnh (Power) mô hình hóa
- **Kim Đan**: `Power_KD ≈ Base · (B^r) · StageMul · (1 + ΣΔ_i)`  
  (B≈6..12; r là RealmRank; StageMul∈[1.0..1.75]; Δ_i là buff min/max)
- **RPG**: `Power_RPG ≈ P0 + k · Level` (tuyến tính) hoặc `≈ P0 · (1 + c · Level)` (tăng tuyến tính theo % nhỏ)

---

## 3) Bảng so sánh (ví dụ minh họa, không phải số liệu cuối)
Giả sử:
- `B = 8`, `StageMul = 1.5`, `Base = 1`
- RPG: `Power_RPG = 10 + 5·Level`

| Tổng XP (giả lập) | Kim Đan đạt | r (rank) | Power_KD (xấp xỉ)         | RPG Level | Power_RPG |
|-------------------|-------------|----------|----------------------------|-----------|-----------|
| 1e4               | Foundation  | 1        | ~ 8^1 × 1.5  = **12**      | 10        | 60        |
| 1e5               | GoldenCore  | 2        | ~ 8^2 × 1.5  = **96**      | 30        | 160       |
| 1e6               | NascentSoul | 3        | ~ 8^3 × 1.5  = **768**     | 70        | 360       |
| 1e7               | SpiritSever | 4        | ~ 8^4 × 1.5  = **6144**    | 120       | 610       |
| 1e8               | SuiNie      | 11       | ~ 8^11× 1.5 = **~3.6e10**  | 200       | 1010      |

> Thấy rõ: với **cùng XP** cao, Kim Đan **leo bậc realm** nên **Power** nhảy vọt **cấp số mũ**; RPG tuyến tính tăng chậm.

---

## 4) Mapping đề xuất (thống nhất nội bộ để UI/UX so sánh)
- **XP normalize**: chuyển tổng XP của mỗi hệ sang **tỉ lệ phần trăm tiến độ** (`progress ∈ [0,1]`) để so sánh công bằng.
- **UI**: hiển thị song song: *“Với XP hiện có, nhân vật KD ~ Realm X (Power≈...), còn RPG ~ Lv Y (Power≈...)”*.
- **API**: cung cấp 2 hàm:
  - `ProjectKdFromXp(totalXp) -> (realm, stage, approxPower)`
  - `ProjectRpgFromXp(totalXp) -> (level, approxPower)`

---

## 5) Pseudocode (Go) — convert theo XP
```go
// Trả về ước lượng realm, stage và power cho Kim Đan từ tổng XP
func ProjectKdFromXp(totalXP float64, cfg XpConfigKD, scale ScalingConfig) (Realm, RealmStage, float64) {
    // 1) Trừ dần XP theo các breakpoint RealmXP[r]
    // 2) Xác định stage theo LevelXP trong realm hiện tại
    // 3) Tính power xấp xỉ qua Gap(R) = A * B^r * StageMul + optional Δ
    // (Ở bước 3 có thể bỏ Δ để nhanh)
    return realm, stage, approxPower
}

// Trả về ước lượng level và power cho RPG từ tổng XP
func ProjectRpgFromXp(totalXP float64, cfg XpConfigRPG) (int, float64) {
    // Lặp tăng level đến khi tích lũy XP vượt totalXP
    // Power_RPG = P0 + k*level (hoặc biến thể)
    return level, approxPower
}
```

---

## 6) Bảng convert mẫu (để UI tooltips)
| KD Realm → RPG Level (cùng XP) | RPG Level → KD Realm (cùng XP) |
|--------------------------------|--------------------------------|
| GoldenCore ~ Lv30              | Lv50 ~ NascentSoul Early       |
| NascentSoul ~ Lv70             | Lv100 ~ SpiritSever Early      |
| SpiritSever ~ Lv120            | Lv180 ~ JingNie Early          |

> Thực tế sẽ phụ thuộc tham số XP từng hệ; doc này đưa **cách làm & API** để Cursor AI implement rồi team bạn cân chỉnh content.
