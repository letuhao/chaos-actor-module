# REALM SCALING MODEL — Base Factors & Exponential Gap
Generated: 2025-09-06T22:07:45.336928Z


> Mục tiêu: phản ánh **độ chênh lệch khủng** giữa các cảnh giới (có thể đến hàng tỷ lần),
> nhưng vẫn giữ được tính **deterministic** và **dễ cân bằng** qua content.

---

## 1) Khái niệm
- **BaseFactor[Realm]**: hệ số nền tảng cho cảnh giới, phản ánh **bậc** sức mạnh.
- **StageMultiplier[Early, Mid, Late, Peak]**: hệ số tiểu cấp trong cùng cảnh giới.
- **Min/Max Delta Table**: cộng bồi thêm (random có seed) khi đột phá, để mỗi nhân vật có khác biệt.
- **RealmRank**: mapping mỗi cảnh giới → chỉ số bậc `r = 0..N` theo thứ tự Tiên Nghịch (Ngưng Khí=0, …, Toái Niết=N).

---

## 2) Công thức tổng quát
Giả sử `S_prev` là tổng primary trước đột phá cảnh giới `R` ở tiểu cấp `T`:
```
Gap(R)           = A * (B ^ RealmRank[R])
StageMul(T)      ∈ {1.00, 1.20, 1.40, 1.60}  // ví dụ
DeltaRoll(R, T)  = Σ_i random(min_i[R], max_i[R]) * StageMul(T)  // cho từng primary i
S_base(R, T)     = max(S_prev, Baseline[R]) * Gap(R) * StageMul(T)
S_new            = S_base(R, T) + DeltaRoll(R, T)
```
- `A` (RealmBaseA) và `B` (RealmBaseB) là **tham số cân bằng** (ví dụ `A=1.0`, `B=8.0` hoặc `B=10.0`).
- `Baseline[R]` giúp tránh cảnh “leo cảnh giới nhưng stat giảm”.

> **Chú ý**: This models **exponential realm gap** (B^r), nên chỉ cần chỉnh `B` để thay đổi **độ vọt** giữa các cảnh giới (đến hàng tỷ lần khi r lớn).

### Mở rộng (ràng buộc & mềm hóa)
- **SoftCap**: Sau khi tính `S_new`, áp dụng `SoftCapFn` để tránh overflow số học.
  - Ví dụ: `SoftCapFn(x) = x * (1 - exp(-x / K))` với K rất lớn, hoặc dùng **logistic**.
- **Piecewise**: Dùng mốc **Transcendence** (ví dụ từ Vấn Đỉnh → Âm Hư) để tăng **B** hoặc thêm **ExponentBoost**:
  - `Gap(R) = A * (B ^ r) * (TBoost ^ t)`, với `t` là số mốc đã vượt.

---

## 3) Tham số gợi ý theo Tiên Nghịch
- Realm order (r tăng): Ngưng Khí (0) → Trúc Cơ (1) → Kết Đan (2) → Nguyên Anh (3) → Hóa Thần (4) → Anh Biến (5) → Vấn Đỉnh (6) → Âm Hư (7) → Dương Thực (8) → Khuy Niết (9) → Tịnh Niết (10) → Toái Niết (11)
- **A = 1.0**
- **B (First Step)**: 6 → 10 (tuỳ game).  
- **ExponentBoost tại mốc**: sau **Vấn Đỉnh** (r ≥ 7) dùng `B2 = B * 2` hoặc `Gap(R) = A * (B ^ r) * (2 ^ (r - 6))`.
- **StageMultiplier**: Early=1.00, Mid=1.25, Late=1.50, Peak=1.75 (có thể điều chỉnh).
- **Baseline[R]**: tăng dần theo r (ví dụ Baseline[Nguyên Anh]=1000, Hóa Thần=3500, …).

---

## 4) Min/Max Delta Table (kết hợp từ 10-REALM-BUFFS-XIANNI.md)
Thay vì cộng thẳng, ta tính delta cho từng **Primary** i:
```
Delta_i(R,T) = random(min_i[R], max_i[R]) * StageMul(T)
S_new_i = max(S_prev_i, Baseline_i[R]) * Gap(R) * StageMul(T) + Delta_i(R,T)
```
- Bảng `min_i[R], max_i[R]` giữ “hương vị” từng cảnh giới (ví dụ Hóa Thần ưu tiên QiControl & DaoComprehension).

---

## 5) Deterministic RNG
- Seed = `hash(actorId + realm + stage + contentVersion)`
- Trả về cùng giá trị cho cùng (actorId, realm, stage) → đảm bảo tái lập.

---

## 6) Pseudocode (Go, comment tiếng Việt)
```go
// Tính stat mới sau khi đột phá cảnh giới
func BreakthroughApply(prev map[string]float64, realm Realm, stage RealmStage, primaries []string, cfg Config, seed string) map[string]float64 { 
    // Gap theo cảnh giới (mô phỏng chênh lệch bậc)
    r := cfg.RealmRank[realm]                   // chỉ số bậc
    gap := cfg.A * math.Pow(cfg.BFor(realm), float64(r)) 
    stageMul := cfg.StageMultiplier[stage]      // ví dụ 1.0/1.25/1.5/1.75

    rng := NewDeterministicRNG(seed)            // hash(actorId+realm+stage+version)

    out := map[string]float64{}
    for _, k := range primaries {
        base := math.Max(prev[k], cfg.Baseline[k][realm])
        minmax := cfg.MinMax[k][realm]          // {Min, Max}
        delta := rng.Range(minmax.Min, minmax.Max) * stageMul

        // Áp dụng công thức
        v := base * gap * stageMul + delta

        // SoftCap (tuỳ chọn)
        if fn := cfg.SoftCap[k]; fn != nil {
            v = fn(v)
        }
        out[k] = v
    }
    return out
}
```

---

## 7) Config đề xuất (JSON/YAML)
```json
{
  "A": 1.0,
  "B": 8.0,
  "StageMultiplier": { "Early": 1.0, "Mid": 1.25, "Late": 1.5, "Peak": 1.75 },
  "RealmRank": { "QiRefining":0, "FoundationEstablishment":1, "GoldenCore":2, "NascentSoul":3, "SpiritSevering":4, "SoulTransformation":5, "WenDing":6, "YinXu":7, "YangShi":8, "KuiNie":9, "JingNie":10, "SuiNie":11 },
  "Baseline": {
    "VitalEssence": { "NascentSoul": 1000, "SpiritSevering": 3500, "SuiNie": 10000000 },
    "QiControl":    { "NascentSoul": 800,  "SpiritSevering": 3000, "SuiNie": 8000000 }
  },
  "MinMax": {
    "VitalEssence": { "GoldenCore": {"Min":5,"Max":10}, "NascentSoul":{"Min":2,"Max":4}, "SuiNie":{"Min":25,"Max":40} },
    "QiControl":    { "GoldenCore": {"Min":4,"Max":8},  "SpiritSevering":{"Min":4,"Max":8} }
  }
}
```
---

## 8) Gợi ý cân bằng
- B nhỏ (5–6) → nhịp tăng mượt nhưng chưa bộc phát.
- B lớn (9–12) → chênh lệch khủng giữa các realm; dùng SoftCap và Baseline để tránh số tràn/đảo dấu.
- Tại mốc **Âm Hư/Dương Thực** nên tăng **B** hoặc nhân thêm hệ số **ExponentBoost** để thể hiện “chuyển bậc” sức mạnh.


---

## 9) Step 4 — Heaven Trampling Extensions
- RealmRank tiếp tục: `HeavenTrampling1..9` → r = 12..20
- Thêm **HTBoost** theo "độ": `Gap(R) = A * (B ^ r) * (HTBoost ^ d)`
- Có thể đặt `HTBoost` lớn (3–6) để phản ánh mức tăng **siêu cấp số mũ**.
- Gợi ý: chuyển từ `SoftCap` dạng logistic sang **piecewise-log** cho tài nguyên (HP/QiCapacity) để ổn định số lớn.
