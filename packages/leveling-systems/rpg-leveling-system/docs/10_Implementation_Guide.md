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
