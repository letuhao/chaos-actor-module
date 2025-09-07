# Detail Design
- PrimaryStats: VitalEssence, QiControl, MeridianStrength, BodyConstitution, SoulConsciousness, DaoComprehension, Fortune
- Realm: QiRefining, FoundationEstablishment, GoldenCore, NascentSoul, DivineSoul, Ascension
- Stage: Early, Mid, Late, Peak
- Resolver pipeline: Base → Flat → Pct → Multiply → Override → Caps → Round
- Derived: QiCapacity, QiRegen, SpiritualPower, BodyRefinement, MentalDefense, FlightSpeed, TribulationResist, HP_MAX


**Realm scaling**: xem `11-REALM-SCALING-MODEL.md` (A/B base factor, StageMultiplier, Baseline, SoftCap, deterministic RNG) và `10-REALM-BUFFS-XIANNI.md`.
