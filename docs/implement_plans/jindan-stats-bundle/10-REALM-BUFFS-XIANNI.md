# REALM-BUFFS-XIANNI — Thiết kế cảnh giới theo **Tiên Nghịch (Xian Ni)**
Generated: 2025-09-06T22:03:24.656130Z

> Ghi chú: Thứ tự cảnh giới và nhị cảnh **Âm Hư → Dương Thực** theo tài liệu cộng đồng về *Tiên Nghịch*.
> Các **giá trị min/max** dưới đây là **đề xuất thiết kế game** (không phải số liệu nguyên tác), để Cursor AI có thể triển khai code.
> Stage (tiểu cấp): **Early, Mid, Late, Peak** với **Multiplier**: 1.00 / 1.20 / 1.40 / 1.60. (Có thể chỉnh trong config).

## Thứ tự cảnh giới (First Step + Quá độ + Second Step)
**Nhất bộ / First Step**  
1) Ngưng Khí (Qi Refining)  
2) Trúc Cơ (Foundation Establishment)  
3) Kết Đan (Core/Golden Core)  
4) Nguyên Anh (Nascent Soul)  
5) Hóa Thần (Spirit Severing)  
6) Anh Biến (Infant/Soul Transformation)  
7) Vấn Đỉnh (Ascendant / Wen Ding)  

**Quá độ giữa Nhất bộ → Nhị bộ:**  
8) Âm Hư (Yin Xu / Illusory Yin)  
9) Dương Thực (Yang Shi / Corporeal Yang)

**Nhị bộ / Second Step (Toái Niết Tam Cảnh):**  
10) Khuy Niết (Kui Nie / Nirvana Scryer)  
11) Tịnh Niết (Jing Nie / Nirvana Cleanser)  
12) Toái Niết (Sui Nie / Nirvana Shatterer)

---

## Primary Stats (Kim Đan) ảnh hưởng khi đột phá
- `VitalEssence` — Nguyên khí/nguyên thần
- `QiControl` — Khống khí/điều tức
- `MeridianStrength` — Kinh mạch
- `BodyConstitution` — Thể chất/luyện thể
- `SoulConsciousness` — Thần thức
- `DaoComprehension` — Ngộ đạo
- `Fortune` — Khí vận

> Công thức cộng: **delta_final = random(min..max) × StageMultiplier**  
> RNG seed = `hash(actorId + realm + stage + contentVersion)` để deterministic.

---

## Bảng đề xuất **min/max** theo từng cảnh giới
> Đơn vị: **điểm primary**; dùng để cộng trực tiếp vào total primary trước khi tính derived.

### 1) Ngưng Khí
- VitalEssence: **+1..+3**
- QiControl: **+1..+2**
- BodyConstitution: **+0..+1**

### 2) Trúc Cơ
- MeridianStrength: **+3..+6**
- BodyConstitution: **+2..+4**
- QiControl: **+2..+3**

### 3) Kết Đan
- VitalEssence: **+3..+6**
- QiControl: **+3..+5**
- SoulConsciousness: **+1..+2**
- DaoComprehension: **+1..+2**

### 4) Nguyên Anh
- SoulConsciousness: **+4..+8**
- QiControl: **+3..+5**
- VitalEssence: **+2..+4**
- DaoComprehension: **+2..+4**

### 5) Hóa Thần
- DaoComprehension: **+5..+10**
- QiControl: **+4..+8**
- VitalEssence: **+3..+6**
- Fortune: **+1..+3**

### 6) Anh Biến
- BodyConstitution: **+6..+12**
- MeridianStrength: **+5..+10**
- SoulConsciousness: **+3..+6**

### 7) Vấn Đỉnh
- VitalEssence: **+10..+20**
- SoulConsciousness: **+8..+16**
- DaoComprehension: **+8..+16**
- QiControl: **+6..+12**
- Fortune: **+3..+7**

### 8) Âm Hư (Quá độ)
- SoulConsciousness: **+10..+18**
- VitalEssence: **+6..+12**
- QiControl: **+4..+8**

### 9) Dương Thực (Quá độ)
- QiControl: **+10..+18**
- MeridianStrength: **+8..+14**
- BodyConstitution: **+6..+12**
- DaoComprehension: **+6..+10**

### 10) Khuy Niết
- DaoComprehension: **+10..+20**
- SoulConsciousness: **+8..+16**
- VitalEssence: **+5..+10**

### 11) Tịnh Niết
- DaoComprehension: **+18..+30**
- SoulConsciousness: **+12..+20**
- QiControl: **+10..+18**
- Fortune: **+5..+10**

### 12) Toái Niết
- DaoComprehension: **+25..+40**
- SoulConsciousness: **+18..+32**
- VitalEssence: **+15..+25**
- Fortune: **+8..+15**

---

## Gợi ý tích hợp
- **Registry**: thêm `realmBuffRanges[realm][stage][primary] = {min,max}`.
- **ProgressionService.TryBreakthrough**: khi thành công, gọi `RollRealmBuff` và commit vào allocations tổng.
- **UI**: tooltip hiển thị roll cụ thể và stage multiplier.
- **Balancing**: chỉnh min/max + multipliers ở YAML/JSON content, không hard-code.



## Bổ sung Step 4 — Heaven Trampling
- Thêm các khóa realm: HeavenTrampling1..HeavenTrampling9.
- Min/Max đề xuất xem `12-REALM-STEP4-HEAVEN-TRAMPLING.md`.
