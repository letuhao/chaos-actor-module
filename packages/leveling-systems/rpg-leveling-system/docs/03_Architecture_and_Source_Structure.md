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
