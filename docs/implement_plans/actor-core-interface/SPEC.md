# Actor Core Interface — SPEC (v2.0.0)

## Scope
This is the **highest-level contract** for the Actor Core v2.0. Implementations must keep all functions **pure** (no IO, no globals).
Combat and other systems depend on these signatures and invariants.

## Core Types (language-agnostic)

### PrimaryCore
- LifeSpan: int64
- Age: int64 (how long the actor has lived)
- Vitality: int64 (overall health and resilience)
- Endurance: int64 (physical stamina and recovery)
- Constitution: int64 (body's natural resistance to damage)
- Intelligence: int64 (problem solving, memory, learning)
- Wisdom: int64 (insight, perception, decision making)
- Charisma: int64 (social influence, leadership, presence)
- Willpower: int64 (mental resistance, focus, determination)
- Luck: int64 (general luck and fortune)
- Fate: int64 (destiny and fate manipulation)
- Karma: int64 (good/evil alignment influence)
- Strength: int64 (physical power and damage)
- Agility: int64 (speed and precision)
- Personality: int64 (social influence and charisma)

#### Universal Cultivation Stats (Compatible with all subsystems)
- SpiritualEnergy: int64 (universal spiritual energy for all cultivation systems)
- PhysicalEnergy: int64 (universal physical energy for body cultivation)
- MentalEnergy: int64 (universal mental energy for magic and mental cultivation)
- CultivationLevel: int64 (overall cultivation level across all systems)
- BreakthroughPoints: int64 (universal breakthrough points for advancement)

#### Flexible Stats (for custom stats shared across systems)
- FlexibleStats: FlexibleStats (custom primary/derived stats)
- SpeedSystem: FlexibleSpeedSystem (flexible speed system)
- AdministrativeSystem: FlexibleAdministrativeSystem (flexible administrative divisions)
- KarmaSystem: FlexibleKarmaSystem (flexible karma system)
- Proficiency: ProficiencySystem (proficiency tracking)
- Skills: UniversalSkillSystem (universal skill system)

### Derived
- HPMax: float64 (calculated from Vitality, Endurance, Constitution)
- Stamina: float64 (calculated from Endurance, Vitality)
- Speed: float64 (calculated from Agility, Intelligence, Luck)
- Energies: map[string, float64] (MP/Qi/Lust/Wrath/etc. - defined per sub-system)
- Haste: float64
- CritChance: float64
- CritChanceResist: float64 (resistance to critical hits)
- CritMulti: float64
- CritMultiResist: float64 (resistance to critical damage multiplier)
- MoveSpeed: float64
- RegenHP: float64
- RegenEnergies: map[string, float64] (regeneration for each energy type)
- Damages: map[DamageType, float64] (damage output by type)
- Defences: map[DefenceType, DefenceValue] (defence against damage types)
- Amplifiers: map[AmplifierType, AmplifierValue] (damage amplification)
- Version: uint64 (bump on recompute)

#### Flexible Stats (for custom derived stats shared across systems)
- FlexibleStats: FlexibleStats (custom derived stats)
- SpeedSystem: FlexibleSpeedSystem (flexible speed system)
- AdministrativeSystem: FlexibleAdministrativeSystem (flexible administrative divisions)
- KarmaSystem: FlexibleKarmaSystem (flexible karma system)

#### Combat Stats
- Accuracy: float64 (hit chance modifier)
- Penetration: float64 (armor penetration)
- Lethality: float64 (critical hit severity)
- Brutality: float64 (damage over time effectiveness)

#### Defense Stats
- ArmorClass: float64 (base armor class for hit avoidance)
- Evasion: float64 (chance to completely avoid attacks)
- BlockChance: float64 (chance to block incoming attacks)
- ParryChance: float64 (chance to parry melee attacks)
- DodgeChance: float64 (chance to dodge ranged attacks)

#### Energy Management
- EnergyEfficiency: map[string, float64] (how efficiently each energy is used)
- EnergyCapacity: map[string, float64] (maximum capacity for each energy type)
- EnergyDrain: map[string, float64] (passive energy consumption rates)
- ResourceRegen: map[string, float64] (regeneration of various resources)
- ResourceDecay: map[string, float64] (natural decay of resources over time)

#### Status & Resistance
- StatusResistance: map[string, float64] (resistance to specific status effects)
- Immunity: map[string, bool] (complete immunity to certain effects)
- WeatherResist: map[string, float64] (resistance to weather effects)
- TerrainMastery: map[string, float64] (effectiveness in different terrains)
- ClimateAdapt: map[string, float64] (adaptation to different climates)

#### Learning & Adaptation
- LearningRate: float64 (how fast the actor learns)
- Adaptation: float64 (how well they adapt to new situations)
- Memory: float64 (information retention)
- Experience: float64 (overall experience level)

#### Social Stats
- Leadership: float64 (ability to lead others)
- Diplomacy: float64 (negotiation and persuasion)
- Intimidation: float64 (scaring or threatening others)
- Empathy: float64 (understanding others' emotions)
- Deception: float64 (lying and misdirection)
- Performance: float64 (acting, singing, dancing, etc.)

#### Mystical Stats
- ManaEfficiency: float64 (how efficiently mana is used)
- SpellPower: float64 (power of magical effects)
- MysticResonance: float64 (connection to mystical forces)
- RealityBend: float64 (ability to bend reality)
- TimeSense: float64 (perception of time)
- SpaceSense: float64 (perception of space)

#### Movement & Mobility
- JumpHeight: float64 (how high the actor can jump)
- ClimbSpeed: float64 (speed of climbing)
- SwimSpeed: float64 (speed of swimming)
- FlightSpeed: float64 (speed of flying, if applicable)
- TeleportRange: float64 (range of teleportation abilities)
- Stealth: float64 (ability to move undetected)

#### Aura & Presence
- AuraRadius: float64 (area of effect for auras)
- AuraStrength: float64 (power of auras)
- Presence: float64 (intimidation and influence)
- Awe: float64 (inspiring fear or respect)

#### Weapon & Skill Proficiency
- WeaponMastery: map[string, float64] (proficiency with different weapon types)
- SkillLevel: map[string, float64] (level in various skills)

#### Talent Amplifiers (Tư Chất Khuếch Đại)
- CultivationSpeed: float64 (tốc độ tu luyện - multiplier for all cultivation systems)
- EnergyEfficiency: float64 (hiệu suất sử dụng năng lượng - reduces energy costs)
- BreakthroughSuccess: float64 (tỷ lệ thành công đột phá - increases breakthrough chance)
- SkillLearning: float64 (tốc độ học kỹ năng - faster skill acquisition)
- CombatEffectiveness: float64 (hiệu quả chiến đấu - damage and defense multipliers)
- ResourceGathering: float64 (hiệu quả thu thập tài nguyên - better resource yields)

#### RPG Derived Stats (from RPG Stats Sub-System)
- LifeSteal: float64 (Life steal percentage)
- CastSpeed: float64 (Spell casting speed modifier)
- WeightCapacity: float64 (Maximum carry weight)
- Persuasion: float64 (Social persuasion ability)
- MerchantPriceModifier: float64 (Merchant price modification)
- FactionReputationGain: float64 (Faction reputation gain rate)

### DamageType Interface
```go
type DamageType interface {
    Name() string
    AffectedStats() []StatAffection
    IsElemental() bool
    IsPhysical() bool
    IsMagical() bool
}

type StatAffection struct {
    StatKey string
    Multiplier float64
    IsDrain bool // true if drains the stat instead of damaging
}
```

### DefenceType Interface
```go
type DefenceType interface {
    Name() string
    DamageTypes() []DamageType
    ResistCap() float64
    DrainCap() float64
    ReflectCap() float64
}

type DefenceValue struct {
    Resist float64  // reduces damage (0.0 to 1.0)
    Drain float64   // converts damage to other resource (0.0 to 1.0)
    Reflect float64 // returns damage to attacker (0.0 to 1.0)
}
```

### AmplifierType Interface
```go
type AmplifierType interface {
    Name() string
    DamageTypes() []DamageType
    IsMultiplicative() bool
    IsPiercing() bool
}

type AmplifierValue struct {
    Multiplier float64 // damage multiplier
    Piercing float64   // pierces through defences (0.0 to 1.0)
}
```

### CoreContribution
- Primary?: PrimaryCore (additive)
- Flat?: map[string, float64] (e.g., "HPMax": +80, "energies.MP": +50, "damages.physical": +25)
- Mult?: map[string, float64] (e.g., "HPMax": 1.10, "amplifiers.physical.multiplier": 1.15)
- Tags?: string[] (capabilities)

### ActorResources
- map[string, ResourceState{current, max, regen, epoch}]

## Stat Categories & Purposes

### Core Survival Stats
- **LifeSpan, Age**: Basic life and longevity
- **Vitality, Endurance, Constitution**: Physical resilience and recovery
- **HPMax, Stamina**: Calculated health and endurance (derived from primary stats)

### Physical & Combat Stats
- **Strength**: Physical power and damage
- **Agility**: Speed and precision
- **Constitution**: Body's natural resistance to damage

### Mental & Social Stats
- **Intelligence, Wisdom**: Mental capacity and insight
- **Charisma, Willpower**: Social influence and mental strength
- **Personality**: Social influence and charisma
- **Leadership, Diplomacy, Intimidation, Empathy, Deception, Performance**: Social interaction

### Mystical & Supernatural Stats
- **Luck, Fate, Karma**: Supernatural influence on events
- **MysticResonance, RealityBend, TimeSense, SpaceSense**: Reality manipulation
- **ManaEfficiency, SpellPower**: Magical effectiveness

### Universal Cultivation Stats
- **SpiritualEnergy, PhysicalEnergy, MentalEnergy**: Universal energy types for all cultivation systems
- **CultivationLevel**: Overall advancement level across all systems
- **BreakthroughPoints**: Universal advancement currency

### Talent Amplifiers (Tư Chất Khuếch Đại)
- **CultivationSpeed, EnergyEfficiency, BreakthroughSuccess**: Universal cultivation multipliers
- **SkillLearning, CombatEffectiveness, ResourceGathering**: Universal effectiveness multipliers

### Combat Effectiveness Stats
- **Accuracy, Penetration, Lethality, Brutality**: Offensive capabilities
- **ArmorClass, Evasion, BlockChance, ParryChance, DodgeChance**: Defensive capabilities
- **WeaponMastery, SkillLevel**: Proficiency with tools and techniques

### Environmental Adaptation Stats
- **WeatherResist, TerrainMastery, ClimateAdapt**: Environmental survival
- **LearningRate, Adaptation, Memory, Experience**: Growth and adaptation
- **Speed**: Base movement capability (derived from primary stats)
- **JumpHeight, ClimbSpeed, SwimSpeed, FlightSpeed, TeleportRange, Stealth**: Movement capabilities

### Aura & Presence Stats
- **AuraRadius, AuraStrength, Presence, Awe**: Area of effect and influence

### RPG Combat Stats
- **LifeSteal**: Combat mechanics
- **CastSpeed**: Magical combat timing

### RPG Social & Economic Stats
- **Persuasion, MerchantPriceModifier, FactionReputationGain**: Social interactions
- **WeightCapacity**: Inventory management

## Predefined Damage Types

### Physical Damage Types
- **Physical**: Affects HP directly
- **Piercing**: Affects HP with armor penetration
- **Slashing**: Affects HP and Stamina
- **Blunt**: Affects HP and Stamina

### Elemental Damage Types
- **Fire**: Affects HP, can spread
- **Water**: Affects HP and Stamina
- **Earth**: Affects HP and Stamina
- **Metal**: Affects HP and Stamina
- **Wood**: Affects HP and Stamina
- **Wind**: Affects HP and Stamina
- **Lightning**: Affects HP and Energies

### Magical Damage Types
- **Light**: Affects HP and Energies
- **Dark**: Affects HP and Energies
- **Arcane**: Affects HP and Energies

### Special Damage Types
- **Time**: Affects LifeSpan and Age
- **Curse**: Affects multiple stats based on curse type
- **Necrosis**: Affects HP and Stamina over time
- **Blood**: Affects HP and Energies
- **Poison**: Affects HP and Stamina over time
- **Disease**: Affects HP, Stamina, and Energies over time

## Predefined Defence Types

### Physical Defences
- **Armor**: Physical, Piercing, Slashing, Blunt
- **Shield**: Physical, Piercing, Slashing, Blunt

### Elemental Defences
- **FireResist**: Fire, Lightning
- **WaterResist**: Water, Ice
- **EarthResist**: Earth, Metal
- **AirResist**: Wind, Lightning

### Magical Defences
- **MagicResist**: Light, Dark, Arcane
- **SpiritResist**: Curse, Necrosis, Blood

### Special Defences
- **TimeResist**: Time
- **DiseaseResist**: Poison, Disease
- **VoidResist**: All damage types (very rare)

## Predefined Amplifier Types

### Physical Amplifiers
- **PhysicalAmp**: Physical, Piercing, Slashing, Blunt
- **WeaponMastery**: All physical damage types

### Elemental Amplifiers
- **ElementalAmp**: All elemental damage types
- **FireMastery**: Fire, Lightning
- **WaterMastery**: Water, Ice
- **EarthMastery**: Earth, Metal
- **AirMastery**: Wind, Lightning

### Magical Amplifiers
- **MagicAmp**: Light, Dark, Arcane
- **SpiritAmp**: Curse, Necrosis, Blood

### Special Amplifiers
- **TimeAmp**: Time
- **VoidAmp**: All damage types (very rare)

## Composition Law (frozen)
```
Derived_final = clamp( ( baseFromPrimary( Σ PrimaryCore )
                       + Σ flat )
                       × Π mult )
```
- Merge buckets in **lexicographic order of their keys** for determinism.
- Multiplicative maps default to **1.0** when absent.

## Clamps (initial)

### Core Stats
- hpMax ≥ 1
- stamina ≥ 0
- speed ≥ 0
- energies.* ≥ 0
- damages.* ≥ 0
- haste ∈ [0.5, 2.0]

### Universal Cultivation Stats
- spiritualEnergy ≥ 0
- physicalEnergy ≥ 0
- mentalEnergy ≥ 0
- cultivationLevel ∈ [0, 100]
- breakthroughPoints ≥ 0

### Talent Amplifiers
- cultivationSpeed ≥ 1.0
- energyEfficiency ≥ 1.0
- breakthroughSuccess ∈ [0.1, 1.0]
- skillLearning ≥ 1.0
- combatEffectiveness ≥ 1.0
- resourceGathering ≥ 1.0

### Critical Hit Stats
- critChance ∈ [0, 1]
- critChanceResist ∈ [0, 1]
- critMulti ∈ [1.0, 5.0]
- critMultiResist ∈ [0, 0.8]

### Defence Stats
- defences.*.resist ∈ [0, 1]
- defences.*.drain ∈ [0, 1]
- defences.*.reflect ∈ [0, 1]
- armourClass ≥ 0
- evasion ∈ [0, 1]
- blockChance ∈ [0, 1]
- parryChance ∈ [0, 1]
- dodgeChance ∈ [0, 1]

### Combat Stats
- accuracy ≥ 0
- penetration ≥ 0
- lethality ≥ 0
- brutality ≥ 0

### Amplifier Stats
- amplifiers.*.multiplier ≥ 0
- amplifiers.*.piercing ∈ [0, 1]

### Energy Management
- energyEfficiency.* ≥ 0
- energyCapacity.* ≥ 0
- energyDrain.* ≥ 0
- resourceRegen.* ≥ 0
- resourceDecay.* ≥ 0

### Status & Resistance
- statusResistance.* ∈ [0, 1]
- weatherResist.* ∈ [0, 1]
- terrainMastery.* ≥ 0
- climateAdapt.* ≥ 0

### Learning & Adaptation
- learningRate ≥ 0
- adaptation ≥ 0
- memory ≥ 0
- experience ≥ 0

### Social Stats
- leadership ≥ 0
- diplomacy ≥ 0
- intimidation ≥ 0
- empathy ≥ 0
- deception ≥ 0
- performance ≥ 0

### Mystical Stats
- manaEfficiency ≥ 0
- spellPower ≥ 0
- mysticResonance ≥ 0
- realityBend ≥ 0
- timeSense ≥ 0
- spaceSense ≥ 0

### Movement & Mobility
- jumpHeight ≥ 0
- climbSpeed ≥ 0
- swimSpeed ≥ 0
- flightSpeed ≥ 0
- teleportRange ≥ 0
- stealth ≥ 0

### Aura & Presence
- auraRadius ≥ 0
- auraStrength ≥ 0
- presence ≥ 0
- awe ≥ 0

### Weapon & Skill Proficiency
- weaponMastery.* ≥ 0
- skillLevel.* ≥ 0

### RPG Stats
- lifeSteal ∈ [0, 1]
- castSpeed ∈ [50, 220]
- weightCapacity ≥ 0
- persuasion ≥ 0
- merchantPriceModifier ≥ 0
- factionReputationGain ≥ 0

## Function Contracts
- ComposeCore(buckets: map[string, CoreContribution]) -> CoreContribution
  - Deterministic; stable sort by key; add Primary & Flat; multiply Mult (identity 1.0).

- BaseFromPrimary(p: PrimaryCore, level: int64) -> Derived
  - Pure; defines balance curve; may read **constants** from module scope but not external IO.

- FinalizeDerived(base: Derived, flat: map[string, float64], mult: map[string, float64]) -> Derived
  - Applies flats then mults; calls ClampDerived; bumps Version (+1).

- ClampDerived(d: Derived) -> Derived
  - Applies bounds above; never produces NaN/Inf.

- RegisterDamageType(dt: DamageType) -> error
  - Registers a new damage type for the system.

- RegisterDefenceType(dt: DefenceType) -> error
  - Registers a new defence type for the system.

- RegisterAmplifierType(at: AmplifierType) -> error
  - Registers a new amplifier type for the system.

## Damage Calculation
```
finalDamage = (baseDamage * amplifier.multiplier) * (1 - defence.resist * (1 - amplifier.piercing))
drainedResource = finalDamage * defence.drain
reflectedDamage = finalDamage * defence.reflect
```

## Critical Hit Calculation
```
critChance = max(0, attacker.critChance - defender.critChanceResist)
critMulti = max(1.0, attacker.critMulti * (1 - defender.critMultiResist))
```

## RPG Stat Stacking Rules (from RPG Stats Sub-System)

### Deterministic Stacking Order
1. **Base** (species/class/level + allocations)
2. **Additive Flat** (items → titles → passives)
3. **Additive Percent** (sum, then apply once)
4. **Multiplicative** (Buffs → Debuffs → Aura/Env)
5. **Override** (priority desc, value desc)
6. **Caps** (min/max, soft DR)
7. **Rounding**

### Group Stacking (`stackId`)
- Same `stackId` stacks up to `maxStacks` (default unlimited)
- **ADD_FLAT/PCT**: sum; **MULTIPLY**: multiply each; **OVERRIDE**: highest priority wins
- **CAP_MAX**: pick smallest; **CAP_MIN**: pick largest

### Examples
- Two items each +10 ATK flat → total +20 before %
- Buff +15% ATK and Debuff -10% ATK → `(base+flat)*(1.15)*(0.90)`
- Override MOVE_SPEED=100 (prio 10) beats prio 5 for duration

## Stat Relationships & Interactions

### Primary Stat Influences
- **Vitality** → HPMax calculation, HP regeneration, disease resistance, poison resistance
- **Endurance** → HPMax calculation, Stamina calculation, stamina regeneration, fatigue resistance, physical recovery
- **Constitution** → HPMax calculation, natural damage resistance, status effect resistance
- **Strength** → Physical attack power, carry weight, melee damage, weapon effectiveness
- **Agility** → Speed calculation, evasion, critical hit chance, precision, reaction time
- **Intelligence** → Speed calculation, energy efficiency, learning rate, spell power, problem solving
- **Wisdom** → Status resistance, mystical resonance, decision making, insight
- **Charisma** → Social influence, leadership effectiveness, aura strength, persuasion
- **Willpower** → Mental resistance, focus, energy control, determination
- **Personality** → Social interactions, merchant prices, faction relations, charisma
- **Luck** → Speed calculation, critical hit chance, random event outcomes, loot quality
- **Fate** → Destiny manipulation, probability bending, reality influence
- **Karma** → Alignment-based bonuses/penalties, moral event outcomes

### Universal Cultivation Stat Influences
- **SpiritualEnergy** → All cultivation systems, breakthrough requirements, energy capacity
- **PhysicalEnergy** → Body cultivation, physical enhancement, stamina regeneration
- **MentalEnergy** → Magic systems, mental cultivation, spell casting, learning rate
- **CultivationLevel** → Overall power scaling, system access, advancement requirements
- **BreakthroughPoints** → System advancement, breakthrough attempts, cultivation progress

### Talent Amplifier Influences
- **CultivationSpeed** → All cultivation systems, learning rate, advancement speed
- **EnergyEfficiency** → All energy types, reduced costs, better resource utilization
- **BreakthroughSuccess** → All breakthrough attempts, success rate, advancement probability
- **SkillLearning** → All skill acquisition, technique mastery, knowledge gain
- **CombatEffectiveness** → All combat systems, damage output, defense effectiveness
- **ResourceGathering** → All resource collection, material acquisition, treasure hunting

### Derived Stat Dependencies
- **HPMax** = f(Vitality, Endurance, Constitution, Level)
- **Stamina** = f(Endurance, Vitality, Level)
- **Speed** = f(Agility, Intelligence, Luck, Level)
- **Accuracy** = f(Speed, Intelligence, Luck, WeaponMastery)
- **Evasion** = f(Speed, Wisdom, Luck, TerrainMastery)
- **LearningRate** = f(Intelligence, Memory, Experience)
- **EnergyEfficiency** = f(Intelligence, Willpower, ManaEfficiency)
- **AuraRadius** = f(Charisma, MysticResonance, Level)
- **SpellPower** = f(Intelligence, MysticResonance, ManaEfficiency)

### Combat Stat Interactions
- **Hit Chance** = Accuracy - Evasion + Luck modifier
- **Damage** = BaseDamage * (1 + Penetration) * (1 + Lethality) * CritMultiplier
- **Defence** = ArmorClass + Evasion + BlockChance + ParryChance
- **Status Resistance** = Constitution + Willpower + StatusResistance

### Energy System Interactions
- **Energy Cost** = BaseCost / EnergyEfficiency
- **Energy Regeneration** = BaseRegen * (1 + Vitality/100) * EnergyEfficiency
- **Energy Capacity** = BaseCapacity * (1 + Intelligence/100) * EnergyCapacity

### Social Stat Interactions
- **Leadership Effectiveness** = Leadership + Charisma + AuraStrength
- **Diplomacy Success** = Diplomacy + Empathy + Charisma
- **Intimidation Power** = Intimidation + Presence + Awe
- **Deception Success** = Deception + Performance + Luck

### Core Derived Stat Formulas
- **HPMax** = (Vitality × 15 + Endurance × 10 + Constitution × 5) × (1 + Level/100) + curve.baseHp(level, classId)
- **Stamina** = (Endurance × 20 + Vitality × 5) × (1 + Level/100) + curve.baseStamina(level, classId)
- **Speed** = (Agility × 2 + Intelligence × 0.5 + Luck × 0.3) × (1 + Level/100) + curve.baseSpeed(level, classId)

### Universal Cultivation Stat Formulas
- **SpiritualEnergy** = (Wisdom × 10 + Willpower × 8 + Intelligence × 6) × (1 + CultivationLevel/100) + curve.baseSpiritualEnergy(level, classId)
- **PhysicalEnergy** = (Strength × 12 + Endurance × 10 + Constitution × 8) × (1 + CultivationLevel/100) + curve.basePhysicalEnergy(level, classId)
- **MentalEnergy** = (Intelligence × 12 + Willpower × 10 + Wisdom × 8) × (1 + CultivationLevel/100) + curve.baseMentalEnergy(level, classId)
- **CultivationLevel** = min(100, (SpiritualEnergy + PhysicalEnergy + MentalEnergy) / 1000 + baseCultivationLevel)
- **BreakthroughPoints** = (CultivationLevel × 10 + Luck × 2) × (1 + Age/1000) + curve.baseBreakthroughPoints(level, classId)

### Talent Amplifier Formulas
- **CultivationSpeed** = 1.0 + (CultivationLevel × 0.01 + Luck × 0.005 + Intelligence × 0.003) × (1 + Age/1000)
- **EnergyEfficiency** = 1.0 + (Wisdom × 0.01 + Willpower × 0.008 + CultivationLevel × 0.005) × (1 + Age/1000)
- **BreakthroughSuccess** = 0.1 + (CultivationLevel × 0.002 + Luck × 0.001 + Willpower × 0.0008) × (1 + Age/1000)
- **SkillLearning** = 1.0 + (Intelligence × 0.01 + Wisdom × 0.008 + CultivationLevel × 0.005) × (1 + Age/1000)
- **CombatEffectiveness** = 1.0 + (Strength × 0.01 + Agility × 0.008 + CultivationLevel × 0.005) × (1 + Age/1000)
- **ResourceGathering** = 1.0 + (Luck × 0.01 + Wisdom × 0.008 + CultivationLevel × 0.005) × (1 + Age/1000)

### RPG Derived Stat Formulas
- **CritChance** = min(1.00, 0.01 + Luck × 0.003 + Agility × 0.001)
- **CritDamage** = 1.5 + Luck × 0.002
- **MoveSpeed** = clamp(100 + Speed × 0.8 + Agility × 0.2, 50, 220)
- **CastSpeed** = clamp(100 + Willpower × 0.6 + Intelligence × 0.2, 50, 220)
- **Persuasion** = Personality × 2 + Luck × 0.5
- **WeightCapacity** = Strength × 10 + Endurance × 5 + baseWeight
- **LifeSteal** = Luck × 0.001 + Willpower × 0.0005

## Sub-System Compatibility

### Universal Cultivation Stats (Compatible with all subsystems)
- **SpiritualEnergy**: Used by Kim Dan, Magic, Succubus, Võ Hiệp systems
- **PhysicalEnergy**: Used by Luyện Thể, Võ Hiệp, Kim Dan systems
- **MentalEnergy**: Used by Magic, Succubus, Kim Dan systems
- **CultivationLevel**: Overall advancement level across all systems
- **BreakthroughPoints**: Universal currency for system advancement

### Sub-System Integration Examples
- **Kim Dan System**: Uses SpiritualEnergy + PhysicalEnergy + MentalEnergy
- **Magic System**: Uses MentalEnergy + SpiritualEnergy + Intelligence
- **Succubus System**: Uses MentalEnergy + SpiritualEnergy + Charisma
- **Luyện Thể System**: Uses PhysicalEnergy + SpiritualEnergy + Strength
- **Võ Hiệp System**: Uses PhysicalEnergy + SpiritualEnergy + MentalEnergy

### Cross-System Bonuses
- **Multi-System Cultivation**: Bonus to CultivationLevel when training multiple systems
- **Energy Synergy**: SpiritualEnergy + PhysicalEnergy + MentalEnergy = Synergy bonus
- **Breakthrough Efficiency**: Reduced BreakthroughPoints cost for multi-system practitioners

## Flexible Systems (Actor Core v2.0)

### Flexible Stats System
- **CustomPrimary**: map[string]int64 (custom primary stats shared across systems)
- **CustomDerived**: map[string]float64 (custom derived stats shared across systems)
- **SubSystemStats**: map[string]map[string]float64 (system-specific stats)

### Flexible Speed System
- **MovementSpeeds**: map[string]float64 (Walking, Running, Swimming, Climbing, Flying, Teleport, Dash, Blink)
- **CastingSpeeds**: map[string]float64 (SpellCasting, MagicFormation, ImmortalTechnique, QiCirculation, Meditation, Breakthrough, Cultivation)
- **CraftingSpeeds**: map[string]float64 (Alchemy, Refining, Forging, Enchanting, ArrayFormation, FormationSetup, PillRefining, WeaponCrafting, ArmorCrafting, JewelryCrafting)
- **LearningSpeeds**: map[string]float64 (Reading, Studying, Comprehension, SkillLearning, TechniqueMastery, Memorization, Research)
- **CombatSpeeds**: map[string]float64 (Attack, Defense, Dodge, Block, Parry, CounterAttack, WeaponSwitch, StanceChange)
- **SocialSpeeds**: map[string]float64 (Conversation, Negotiation, Persuasion, Intimidation, Diplomacy, Leadership, Trading, Networking)
- **AdministrativeSpeeds**: map[string]float64 (Planning, Organization, DecisionMaking, ResourceManagement, StrategyFormulation, TacticalAnalysis, RiskAssessment)

### Flexible Administrative Division System
- **Divisions**: map[string]map[string]AdministrativeDivision (divisionType -> divisionName -> Division)
- **DivisionTypes**: map[string]DivisionType (100+ division types)
- **Hierarchy**: map[string][]string (parentDivision -> childDivisions)
- **Relationships**: map[string]map[string]DivisionRelationship (divisionA -> divisionB -> Relationship)

### Flexible Karma System
- **GlobalKarma**: map[string]int64 (karmaType -> totalValue)
- **DivisionKarma**: map[string]map[string]int64 (divisionType -> divisionName -> karmaType -> value)
- **KarmaCategories**: map[string][]string (category -> karma types)
- **KarmaTypes**: map[string]KarmaType (karma type definitions)

### Proficiency System
- **Proficiencies**: map[string]*Proficiency (skillName -> Proficiency)
- **Categories**: map[string][]string (category -> skill names)
- **MaxSkills**: int (maximum number of skills)

### Universal Skill System
- **Skills**: map[string]*UniversalSkill (skillName -> Skill)
- **Categories**: map[string][]string (category -> skill names)
- **SkillTrees**: map[string]*SkillTree (skillTreeName -> SkillTree)
- **MaxSkills**: int (maximum number of skills)

## Talent System (Hệ Thống Tư Chất)

### Talent Tiers (Cấp Độ Tư Chất)
- **0-20**: Phế Vật (Waste) - 0.5x multiplier
- **21-40**: Phàm Nhân (Mortal) - 1.0x multiplier
- **41-60**: Tài Năng (Talent) - 1.5x multiplier
- **61-80**: Thiên Tài (Genius) - 2.0x multiplier
- **81-100**: Thánh Tài (Saint) - 3.0x multiplier

### Talent Amplifier Effects
- **CultivationSpeed**: Multiplies all cultivation progress rates
- **EnergyEfficiency**: Reduces energy costs for all actions
- **BreakthroughSuccess**: Increases breakthrough success probability
- **SkillLearning**: Multiplies skill acquisition rates
- **CombatEffectiveness**: Multiplies damage and defense values
- **ResourceGathering**: Multiplies resource collection yields

### Cross-System Talent Synergies
- **Dual System Mastery**: 2+ systems with high CultivationLevel → +10% overall effectiveness
- **Triple System Mastery**: 3+ systems with high CultivationLevel → +20% overall effectiveness
- **Perfect Harmony**: All systems balanced → +50% overall effectiveness

## Invariants / Laws
- Determinism: same inputs → same Derived; independent of map iteration order.
- Idempotence: composing with unchanged buckets does not change results.
- Monotonicity: increasing any PrimaryCore field never reduces derived stat outputs from BaseFromPrimary.
- Purity: no time/IO/randomness.
- Damage Type Flexibility: new damage types can be registered at runtime.
- Defence Type Flexibility: new defence types can be registered at runtime.
- Amplifier Type Flexibility: new amplifier types can be registered at runtime.
- Sub-System Compatibility: Universal Cultivation Stats work with all cultivation systems.

## Versioning
Embed `schemaVersion`, `balanceVersion`, `seedVersion` in outer snapshots (out of scope for this bundle).

## Implementation Priorities

### Phase 1: Core Stats (Essential)
- **PrimaryCore**: LifeSpan, Age, Vitality, Endurance, Constitution, Intelligence, Willpower, Luck, Strength, Agility, Personality
- **Universal Cultivation**: SpiritualEnergy, PhysicalEnergy, MentalEnergy, CultivationLevel, BreakthroughPoints
- **Talent Amplifiers**: CultivationSpeed, EnergyEfficiency, BreakthroughSuccess, SkillLearning, CombatEffectiveness, ResourceGathering
- **Derived**: HPMax, Stamina, Speed, Energies, Haste, CritChance, CritChanceResist, CritMulti, CritMultiResist
- **Combat**: Accuracy, Penetration, ArmorClass, Evasion, BlockChance, ParryChance
- **Energy**: EnergyEfficiency, EnergyCapacity, RegenEnergies

### Phase 2: Flexible Systems (Actor Core v2.0)
- **Flexible Stats**: CustomPrimary, CustomDerived, SubSystemStats
- **Flexible Speed**: MovementSpeeds, CastingSpeeds, CraftingSpeeds, LearningSpeeds, CombatSpeeds, SocialSpeeds, AdministrativeSpeeds
- **Flexible Administrative**: Divisions, DivisionTypes, Hierarchy, Relationships
- **Flexible Karma**: GlobalKarma, DivisionKarma, KarmaCategories, KarmaTypes

### Phase 3: Combat Enhancement
- **RPG Combat**: LifeSteal, CastSpeed
- **Defence**: StatusResistance
- **Offence**: Lethality, Brutality, WeaponMastery
- **Mystical**: ManaEfficiency, SpellPower, MysticResonance

### Phase 4: Social & Mental
- **RPG Social**: Persuasion, MerchantPriceModifier, FactionReputationGain, WeightCapacity
- **PrimaryCore**: Wisdom, Charisma, Fate, Karma
- **Social**: Leadership, Diplomacy, Intimidation, Empathy, Deception, Performance
- **Mental**: LearningRate, Memory, Adaptation

### Phase 5: Advanced Features
- **Supernatural**: RealityBend, TimeSense, SpaceSense
- **Environmental**: WeatherResist, TerrainMastery, ClimateAdapt
- **Movement**: JumpHeight, ClimbSpeed, SwimSpeed, FlightSpeed, TeleportRange, Stealth
- **Aura**: AuraRadius, AuraStrength, Presence, Awe

### Phase 6: Specialized Systems
- **Resource Management**: ResourceRegen, ResourceDecay, EnergyDrain
- **Immunity System**: Immunity map
- **Proficiency System**: Proficiencies, Categories, MaxSkills
- **Universal Skill System**: Skills, Categories, SkillTrees, MaxSkills
- **Advanced Interactions**: Complex stat relationships

## Testing Guidance (for implementers)
- Property tests for commutativity (with stable sort), idempotence, monotonicity, clamp ranges.
- Golden test: fixed buckets produce the same Derived across runs.
- Damage type registration tests.
- Defence type registration tests.
- Amplifier type registration tests.
- Damage calculation tests with various combinations.
- Critical hit calculation tests.
- Stat relationship tests.
- Energy system tests.
- Social interaction tests.
- Environmental adaptation tests.
- RPG stat formula tests.
- RPG stacking rule tests.
- RPG resistance calculation tests.
- RPG social interaction tests.

### Flexible Systems Testing (Actor Core v2.0)
- **Flexible Stats**: CustomPrimary/CustomDerived/SubSystemStats integration tests
- **Flexible Speed**: Speed calculation tests across all categories
- **Flexible Administrative**: Division hierarchy and relationship tests
- **Flexible Karma**: Karma calculation and influence tests
- **Proficiency System**: Skill mastery and progression tests
- **Universal Skill System**: Skill learning and advancement tests
- **Multi-System Integration**: Cross-system compatibility tests
- **Performance Tests**: Large-scale flexible system performance
- **Scalability Tests**: System expansion and growth tests