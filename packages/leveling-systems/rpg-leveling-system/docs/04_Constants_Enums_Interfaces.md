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
