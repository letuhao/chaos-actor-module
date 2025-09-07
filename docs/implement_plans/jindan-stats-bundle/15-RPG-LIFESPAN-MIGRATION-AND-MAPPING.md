# 15 — RPG Lifespan Migration & Cross-System Mapping (Kim Đan ⇄ RPG)
Generated: 2025-09-06T22:31:19.603831Z

> Comment tiếng Việt; **class/method/property code dùng tiếng Anh**.  
> Tài liệu này hướng dẫn:
> 1) **Migrate chỉ số Lifespan** cho **RPG hệ thống** (tuyến tính).  
> 2) **Thêm "review mapping"** giữa **Kim Đan** và **RPG** vào **source code RPG** (in ra bảng so sánh theo XP).

---

## A) Migrate Lifespan vào RPG System

### A.1 Mục tiêu
- RPG (linear) hiện **chưa có** hoặc **khác cách** quản lý thọ nguyên. Ta thêm **Lifespan** (năm) tương tự Core Actor, nhưng gắn quy tắc **tuyến tính** theo level và milestone.
- Đảm bảo **tương thích ngược**: nếu record cũ không có Lifespan → **backfill** theo Level.

### A.2 MongoDB schema change (gợi ý)
Collection `rpg_player_progress` (hoặc tương đương):
- Thêm trường:
  - `lifespan` (number) — thọ nguyên tối đa hiện tại (năm).
  - `age` (number, optional) — tuổi hiện tại.
  - `lifespanVersion` (number) — version để track migration.

**Index**: không bắt buộc thêm index cho lifespan.

### A.3 Backfill chiến lược
- Công thức gợi ý (linear theo Level):
  ```
  LifespanBase(Lv) = L0 + k * Lv
  Milestone bonus: +m tại các mốc (Lv20, 40, 60, 80, 100, ...)
  ```
  Ví dụ: `L0=80`, `k=2`, mốc 20/40/60/80/100: `+20/+30/+40/+50/+60`.

- Pseudocode (Go):
  ```go
  // Convert level -> lifespan (linear)
  func ComputeRpgLifespanFromLevel(level int, cfg RpgLifespanConfig) int {
      base := cfg.L0 + cfg.K*level
      bonus := 0
      for _, ms := range cfg.Milestones {
          if level >= ms.Level {
              bonus += ms.Bonus
          }
      }
      return base + bonus
  }
  ```

- **Backfill job**:
  1) Quét tất cả người chơi, nếu `lifespan` thiếu → set bằng `ComputeRpgLifespanFromLevel(level)`.
  2) Set `lifespanVersion = 1`.
  3) Ghi log số lượng record cập nhật; chạy **idempotent**.

### A.4 API thay đổi (RPG service)
- Thêm field Lifespan vào DTO trả ra UI: `PlayerProfile.Lifespan`.
- Khi level-up: cập nhật Lifespan theo hàm tuyến tính (nếu chọn auto-sync) hoặc tách cấu hình tùy game.
- Nếu dùng Core Actor chung: đồng bộ `ActorCore.Lifespan` = Lifespan của hệ đang kích hoạt (RPG hoặc Kim Đan).

---

## B) "Review Mapping" giữa Kim Đan & RPG trong mã nguồn RPG

### B.1 Mục tiêu
- Viết **module tiện ích** trong RPG để **review** (so sánh) sức mạnh giữa **Kim Đan** (exponential realm) và **RPG** (linear) theo **cùng tổng XP**.
- Module này phục vụ **balancing**: in ra bảng so sánh theo dải XP do designer nhập.

### B.2 Public API (Go, trong RPG repo)
```go
// Ước lượng Level & Power ở RPG theo tổng XP
type RpgXpConfig struct {
    L0 float64     // hệ số cơ sở cho đường cong XP
    Beta float64   // bậc tăng XP theo level (1.2..1.6)
    P0 float64     // power offset
    K  float64     // power gain per level (linear)
}

func ProjectRpgFromXp(totalXP float64, cfg RpgXpConfig) (level int, approxPower float64)

// Ước lượng Realm/Stage & Power ở Kim Đan theo tổng XP (đọc từ content/bundle)
type KdXpConfig struct {
    LevelL0 float64
    Alpha float64 // bậc tăng XP cho tiểu cấp KD
    RealmR0 float64
    Gamma float64 // bậc tăng XP cho breakthrough (6..12)
}

type KdScalingConfig struct {
    A float64  // base factor
    B float64  // realm exponential base (6..12)
    StageMultiplier map[string]float64
}

func ProjectKdFromXp(totalXP float64, xpCfg KdXpConfig, scaleCfg KdScalingConfig) (realm string, stage string, approxPower float64)
```

### B.3 Bảng in "review mapping" (CLI nhỏ trong RPG)
```go
// cmd/review-mapping/main.go
func main() {
    xpSamples := []float64{1e4, 5e4, 1e5, 5e5, 1e6, 1e7, 1e8}
    rpgCfg := RpgXpConfig{L0: 100, Beta: 1.3, P0: 10, K: 5}
    kdXp := KdXpConfig{LevelL0: 50, Alpha: 1.5, RealmR0: 1000, Gamma: 8}
    kdScale := KdScalingConfig{A: 1, B: 8, StageMultiplier: map[string]float64{"Early":1.0,"Mid":1.25,"Late":1.5,"Peak":1.75}}

    fmt.Println("| TotalXP | KD (Realm/Stage, Power) | RPG (Lv, Power) |")
    fmt.Println("|---------|--------------------------|-----------------|")
    for _, xp := range xpSamples {
        kdRealm, kdStage, kdPower := ProjectKdFromXp(xp, kdXp, kdScale)
        rpgLevel, rpgPower := ProjectRpgFromXp(xp, rpgCfg)
        fmt.Printf("| %.0f | %s/%s, ~%.2f | Lv%d, ~%.2f |\n", xp, kdRealm, kdStage, kdPower, rpgLevel, rpgPower)
    }
}
```

> **Lưu ý**: Trong repo RPG, đặt module này dưới `tools/review-mapping/` hoặc `cmd/review-mapping/`.  
> Kết quả in bảng dùng để **điều chỉnh** tham số `B, Gamma` (Kim Đan) và `K, Beta` (RPG) sao cho **cùng XP** nhưng **chênh lệch sức mạnh đáng kể** (Kim Đan > RPG).

### B.4 Phần Lifespan trong review
- Có thể in thêm cột Lifespan ước lượng:
  - RPG: `ComputeRpgLifespanFromLevel(level)`
  - Kim Đan: `ApplyRealmLifespan` (dùng base & delta theo realm) — hoặc ước lượng bằng bảng ở `13-REALM-LIFESPAN.md`.

---

## C) Testing & CI

### C.1 Unit tests
- `ProjectRpgFromXp`: kiểm tra monotonicity — XP tăng → Level tăng, Power tăng tuyến tính.
- `ProjectKdFromXp`: kiểm tra phân bổ XP vào realm & stage; Power tăng **siêu tuyến tính**.
- Lifespan backfill: case level biên (Lv1, Lv100, Lv200…).

### C.2 Regression/Golden tests
- Fix sẵn một bộ tham số (rpgCfg/kdCfg) và một list XP → snapshot bảng kết quả (golden).  
- Khi thay đổi tham số, test sẽ báo diff để designer duyệt.

### C.3 DevEx
- CLI `review-mapping` trả về **exit code 0**, viết ra **CSV** để nhập Excel:
  - `--out csv` → `mapping_report.csv`.

---

## D) Checklist triển khai (RPG repo)
1) **Schema**: thêm `lifespan`, `age`, `lifespanVersion`.  
2) **Backfill**: viết job dùng `ComputeRpgLifespanFromLevel`.  
3) **Service/API**: expose Lifespan ra UI.  
4) **Mapping Tool**: thêm `ProjectRpgFromXp`, `ProjectKdFromXp` (SDK hoặc copy stub), và CLI `review-mapping`.  
5) **Test**: unit + golden; thêm CI step chạy CLI `--out csv` và upload artifact.  
6) **Docs**: liên kết tới `13-REALM-LIFESPAN.md` và `14-CROSS-SYSTEM-MAPPING.md` để cân bằng tham số.
