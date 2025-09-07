

# --- 00_README.md ---

# RPG Leveling System — Design Docs (Subsystem for Actor Core v3)
**Generated:** 2025-09-07 18:50

This bundle specifies a **combat leveling system** (đánh quái thăng cấp) that plugs into **Actor Core v3** as an independent **Subsystem**.
- Uses **8 primary attributes (Morrowind)**.
- No skills/spells inside this system. It focuses on: **XP → Level → Stat Points → Derived Resources**.
- Provides **decorators** so other systems (Magic, Cultivation/Thể, Items, World Rules) can plug XP logic, resource formulas, and damage types.
- Supports **unbounded levels** and small **lifespan** gains per level.


# --- 01_Product_Scope.md ---

# 01 — Phạm vi & Mục tiêu

## Mục tiêu
- Hệ thống **Leveling** độc lập: nhận XP (giết quái/PvP/nhiệm vụ/khám phá) → tăng level → cấp **điểm thuộc tính** (stat points).
- Không có skill/spell nội bộ; **decorator** cho hệ thống khác tính **mana/stamina**, **damage types**, và **XP modifiers**.
- **Cấp bậc vô hạn** (practically unbounded), tăng **thọ nguyên** rất nhỏ mỗi cấp.

## Không nằm trong phạm vi
- Không chứa công thức chiến đấu chi tiết (crit, hit, v.v.): chỉ cung cấp **primary stats** và **decorator hooks**.
- Không điều khiển loot/economy; không làm UI.


# --- 02_Primary_Stats.md ---

# 02 — Primary Stats (Morrowind 8)

We adopt the classic 8 attributes:
1. **strength**
2. **intelligence**
3. **willpower**
4. **agility**
5. **speed**
6. **endurance**
7. **personality**
8. **luck**

## Derived (from decorators)
- **hp_max** (baseline mapping in this system; tunable)
- **mana_max** (via decorators; default mapping optional)
- **stamina_max** (via decorators; default mapping optional)
- **resistances** per damage type (baseline from primary stats; decorators may extend)


# --- 03_Architecture_and_Source_Structure.md ---

# 03 — Kiến trúc & Cấu trúc mã cho Subsystem

This **RPG Leveling** module is a **Subsystem** for Actor Core v3.

```
Actor (metadata) ──▶ [RPGLevelingSubsystem] ─┬─▶ Contributions (Primary/Derived)
                                             ├─▶ CapContributions (lifespan_years ADDITIVE max small per level)
                                             └─▶ Context Modifiers (for damage/exp hooks, via decorators)
```

Suggested repo layout for this subsystem:
```
/src/subsystems/rpg_leveling/
  ├─ constants/         # min/max points per level, XP curve defaults, resist caps
  ├─ config/            # YAML/JSON configs (curves, tables)
  ├─ decorators/        # interfaces & registration
  ├─ xp/                # xp source handlers (kill/quest/explore), anti-exploit
  ├─ damage/            # damage type registry & resistance mapping
  ├─ resources/         # mana/stamina mapping (via decorators)
  ├─ model/             # data contracts
  └─ subsystem.ts|go    # entrypoint implements Subsystem
```


# --- 04_Constants_Enums_Interfaces.md ---

# 04 — Constants / Enums / Interfaces

## Constants (configurable in data)
- `POINTS_PER_LEVEL_MIN` / `POINTS_PER_LEVEL_MAX`: số điểm mỗi lần thăng cấp (ví dụ: 3..10).
- `XP_BASE`, `XP_GROWTH`, `XP_EXP`: tham số đường cong XP lên cấp (hỗ trợ nhiều dạng).
- `LIFESPAN_PER_LEVEL` (years): +0.25~1.0 tuỳ thiết kế (rất nhỏ).
- `RESIST_CAP_MAX`: ví dụ 0.85 (85%); `RESIST_CAP_MIN`: 0.
- `DAMAGE_TYPES_BASELINE`: danh sách mặc định damage types.
- `LEVEL_SOFTCAPS`: mốc mềm để điều chỉnh tăng trưởng (tuỳ chọn).

## Enums
- `XPSource`: `KILL`, `PVP_KILL`, `QUEST`, `EXPLORE`, `EVENT`.
- `DamageType`: dynamic registry; default include: `physical`, `slash`, `pierce`, `blunt`, `fire`, `frost`, `shock`, `poison`, `acid`, `holy`, `shadow`, `spirit`, `psychic`.
- `DecoratorPhase` (XP pipeline): `PRE`, `BASE`, `POST`, `FINALIZE`.

## Interfaces (language-agnostic)
### Subsystem (as in Actor Core v3)
- `systemId(): "rpg_leveling"`
- `priority(): number`
- `contribute(actor, ctx): SubsystemOutput`

### Decorators
#### 1) XP Decorator
```ts
interface XPDecorator {
  id(): string;
  phase(): 'PRE'|'BASE'|'POST'|'FINALIZE';
  onXPCompute(input: {
    actorGUID: string;
    source: XPSource;
    baseXP: number;
    level: number;
    targetLevel?: number;
    tags?: string[];
  }): { deltaFlat?: number; deltaPercent?: number; capMax?: number }
}
```
- Kết quả gộp: flat cộng thẳng, percent gộp delta vào multiplier, capMax (nếu có) kẹp kết quả.

#### 2) Resource Decorator (mana/stamina)
```ts
interface ResourceDecorator {
  id(): string;
  computeResources(input: {
    level: number;
    strength: number; intelligence: number; willpower: number; agility: number;
    speed: number; endurance: number; personality: number; luck: number;
  }): { mana_max?: number; stamina_max?: number; hp_max?: number }
}
```
- Trả về **contributions** (FLAT/MULT/POST_ADD) qua SubsystemOutput.

#### 3) Damage Type Decorator
```ts
interface DamageTypeDecorator {
  id(): string;
  registerTypes(): Array<{ key: string; inherits?: string }>;
  resistFromStats?(ps: PrimaryStats): Array<{ dimension: string; value: number; bucket: 'FLAT'|'MULT'|'POST_ADD' }>;
  damageCompute?(input: {
    attackerLevel: number; targetLevel: number;
    damageType: string; baseDamage: number; stats: PrimaryStats;
  }): { deltaFlat?: number; deltaPercent?: number }
}
```
- `resistFromStats` biến primary stats thành **resist_*** contributions.
- `damageCompute` cho phép các hệ (ma pháp/tu tiên) chỉnh damage theo stats.

### Data Contracts (RPG Leveling → Actor Core)
- Primary: 8 stats contributions (từ điểm người chơi phân phối).
- Derived: resources/resistances contributions (từ decorators hoặc baseline).
- Caps: `lifespan_years` via `ADDITIVE max` (nhỏ mỗi level), resist caps `HARD_MAX`.
- Context: `xp_gain`, `damage_out`, `damage_in` modifiers (tuỳ chọn).


# --- 05_XP_and_Leveling.md ---

# 05 — XP & Leveling

## Sources
- **Kill** (`KILL`, `PVP_KILL`): XP dựa trên `targetLevel`, `rarity`, `elite/boss`.
- **Quest** (`QUEST`): XP từ nhiệm vụ; table dữ liệu.
- **Explore** (`EXPLORE`): XP khi khám phá mốc bản đồ.
- **Event** (`EVENT`): XP theo sự kiện thời gian thực.

## XP Formula (configurable)
1) **Exponential**: `XP_to_next(L) = XP_BASE * (XP_GROWTH ^ L)`
2) **Polynomial**: `XP_to_next(L) = XP_BASE * (L + 1) ^ XP_EXP`
3) **Hybrid**: poly đến L* rồi exponent (softcaps).

### Kill XP
```
XP = BaseByTier(targetTier) * DisparityModifier(level, targetLevel) * (1 + ΣdecoratorPerc) + ΣdecoratorFlat
XP_final = min(XP, capMax_from_decorators? else +∞)
```
- Anti‑exploit: DiminishingReturn(sameTargetCount), DailyCap, PartyShare, Rested bonus (qua decorators).

## Level Up
- Khi XP ≥ `XP_to_next(L)`: tăng L, trừ XP.
- **Points per level**: random `[MIN..MAX]` hoặc theo bảng; **player tự phân phối**.
- **Lifespan bonus**: mỗi level emit `CapContribution`: `{ dimension:'lifespan_years', mode:'ADDITIVE', kind:'max', value: LIFESPAN_PER_LEVEL, scope:'TOTAL' }`.


# --- 06_Decorators_Spec.md ---

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


# --- 07_Damage_and_Resist_Integration.md ---

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


# --- 08_Caps_and_Realm_Integration.md ---

# 08 — Caps & Realm Integration

- **Lifespan**: emit `ADDITIVE max` mỗi level (layer `TOTAL`).
- **Resistances**: emit `HARD_MAX` = `RESIST_CAP_MAX` (layer `WORLD` hoặc `TOTAL`).
- **Primary caps** (tuỳ chọn) theo mốc level mềm.
- Tương thích **CapLayerRegistry**: chỉ emit caps cho **layers đang active**.


# --- 09_Data_Schemas.md ---

# 09 — Data Schemas
- `RPG_Config.schema.json` — config hệ.
- `XP_Event.schema.json` — sự kiện XP.
- `Damage_Type_Registry.schema.json` — đăng ký damage types.


# --- 10_Implementation_Guide.md ---

# 10 — Implementation Guide (Subsystem)

1) Load config.
2) Quản lý state: `level`, `xp_current`, `points_unspent`, `allocations`.
3) XP intake qua **XPDecorators** (phases).
4) Level up → thêm points (MIN..MAX).
5) Player allocate → phát **Primary Contributions**.
6) Gọi **ResourceDecorators** (+baseline nếu bật) → **Derived Contributions**.
7) Gọi **DamageTypeDecorators.resistFromStats** → **resist_* Contributions** + caps.
8) Emit **lifespan cap** ADDITIVE mỗi level.
9) Bổ sung **Context** nếu cần (`xp_gain`, `damage_out/in`).
10) Trả **SubsystemOutput** cho Aggregator.


# --- 11_Testing_Strategy.md ---

# 11 — Testing Strategy

- Unit: XP curve, anti-exploit, points allocation.
- Property: sort invariance trong cùng phase; resist clamp.
- Golden: (config + allocations + xp_events) → SubsystemOutput expected.
- Integration: chạy qua Actor Core aggregator.


# --- appendix/example_walkthrough.md ---

# Example Walkthrough

- Level 10, exponential (base 100, growth 1.12) → `XP_to_next ≈ 311`.
- Kill elite L12 → after decorators → gain ~156 XP.
- Level up: +5 points (from [3..7]); allocate +3 strength, +2 endurance.
- Contributions:
  - Primary: `strength +3`, `endurance +2` (FLAT).
  - Derived (baseline): `hp_max += 5*2 + 2*3 = 16`.
  - Caps: `lifespan_years ADDITIVE max += 0.5`.
- Resistances: endurance 12 ⇒ `physical_resist +0.12` (clamp 0.85).


# --- PR_Checklist.md ---

# PR Checklist — RPG Leveling Subsystem

- [ ] Config validated against `RPG_Config.schema.json`.
- [ ] XP pipeline phases implemented; decorators sorted deterministically.
- [ ] Points/level within `[MIN..MAX]`; allocation exported as Contributions.
- [ ] Resource decorators invoked; baseline mapping toggle respected.
- [ ] Damage type registry & resist mapping; resist caps emitted as HARD_MAX.
- [ ] Lifespan cap ADDITIVE per level.
- [ ] SubsystemOutput conforms (Primary/Derived/Caps/Context/Meta).
- [ ] Golden tests cover kill/quest/explore; anti-exploit; decorators stacking.
