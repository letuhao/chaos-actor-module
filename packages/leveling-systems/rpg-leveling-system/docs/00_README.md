# RPG Leveling System — Design Docs (Subsystem for Actor Core v3)
**Generated:** 2025-09-07 18:50

This bundle specifies a **combat leveling system** (đánh quái thăng cấp) that plugs into **Actor Core v3** as an independent **Subsystem**.
- Uses **8 primary attributes (Morrowind)**.
- No skills/spells inside this system. It focuses on: **XP → Level → Stat Points → Derived Resources**.
- Provides **decorators** so other systems (Magic, Cultivation/Thể, Items, World Rules) can plug XP logic, resource formulas, and damage types.
- Supports **unbounded levels** and small **lifespan** gains per level.
