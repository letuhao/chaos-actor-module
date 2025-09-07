# PR Checklist — RPG Leveling Subsystem

- [ ] Config validated against `RPG_Config.schema.json`.
- [ ] XP pipeline phases implemented; decorators sorted deterministically.
- [ ] Points/level within `[MIN..MAX]`; allocation exported as Contributions.
- [ ] Resource decorators invoked; baseline mapping toggle respected.
- [ ] Damage type registry & resist mapping; resist caps emitted as HARD_MAX.
- [ ] Lifespan cap ADDITIVE per level.
- [ ] SubsystemOutput conforms (Primary/Derived/Caps/Context/Meta).
- [ ] Golden tests cover kill/quest/explore; anti-exploit; decorators stacking.
