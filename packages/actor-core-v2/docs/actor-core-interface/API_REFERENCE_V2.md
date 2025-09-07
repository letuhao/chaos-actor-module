# API Reference - Actor Core v2.0

## Tổng Quan

API Reference cho Actor Core v2.0 cung cấp tài liệu chi tiết về tất cả interfaces, structs, methods, và functions trong hệ thống.

## 1. Core Interfaces

### 1.1 StatProvider Interface
```go
type StatProvider interface {
    // Get primary stats
    GetPrimaryStats() map[string]int64
    
    // Get derived stats
    GetDerivedStats() map[string]float64
    
    // Get custom primary stats
    GetCustomPrimaryStats() map[string]int64
    
    // Get custom derived stats
    GetCustomDerivedStats() map[string]float64
    
    // Get subsystem stats
    GetSubSystemStats(systemName string) map[string]float64
    
    // Get all stats
    GetAllStats() *StatSnapshot
    
    // Check if stat exists
    HasStat(statName string) bool
    
    // Get stat value
    GetStatValue(statName string) (interface{}, error)
}
```

### 1.2 StatConsumer Interface
```go
type StatConsumer interface {
    // Consume stat changes
    OnStatChanged(statName string, oldValue, newValue interface{})
    
    // Consume batch stat changes
    OnStatsChanged(changes map[string]StatChange)
    
    // Consume stat snapshot
    OnStatSnapshot(snapshot *StatSnapshot)
    
    // Check if consumer is active
    IsActive() bool
    
    // Get consumer priority
    GetPriority() int
}
```

### 1.3 StatResolver Interface
```go
type StatResolver interface {
    // Resolve all stats
    ResolveStats(primaryStats map[string]int64) (*StatSnapshot, error)
    
    // Resolve specific stat
    ResolveStat(statName string, primaryStats map[string]int64) (interface{}, error)
    
    // Resolve derived stats
    ResolveDerivedStats(primaryStats map[string]int64) (map[string]float64, error)
    
    // Resolve flexible stats
    ResolveFlexibleStats(primaryStats map[string]int64) (*FlexibleStats, error)
    
    // Check dependencies
    CheckDependencies(statName string) ([]string, error)
    
    // Get calculation order
    GetCalculationOrder() []string
}
```

## 2. Core Structs

### 2.1 PrimaryCore Struct
```go
type PrimaryCore struct {
    // Basic Stats
    Vitality     int64  // Overall health and resilience
    Endurance    int64  // Physical stamina and recovery
    Constitution int64  // Body's natural resistance to damage
    Intelligence int64  // Problem solving, memory, learning
    Wisdom       int64  // Insight, perception, decision making
    Charisma     int64  // Social influence, leadership, presence
    Willpower    int64  // Mental resistance, focus, determination
    Luck         int64  // General luck and fortune
    Fate         int64  // Destiny and fate manipulation
    Karma        int64  // Good/evil alignment influence
    
    // Physical Stats
    Strength    int64  // Physical power and damage
    Agility     int64  // Speed and precision
    Personality int64  // Social influence and charisma
    
    // Universal Cultivation Stats
    SpiritualEnergy   int64  // Universal spiritual energy for all cultivation systems
    PhysicalEnergy    int64  // Universal physical energy for body cultivation
    MentalEnergy      int64  // Universal mental energy for magic and mental cultivation
    CultivationLevel  int64  // Overall cultivation level across all systems
    BreakthroughPoints int64  // Universal breakthrough points for advancement
    
    // Flexible Stats
    FlexibleStats *FlexibleStats
    
    // Flexible Systems
    SpeedSystem        *FlexibleSpeedSystem
    AdministrativeSystem *FlexibleAdministrativeSystem
    KarmaSystem        *FlexibleKarmaSystem
    Proficiency        *ProficiencySystem
    Skills             *UniversalSkillSystem
}
```

### 2.2 Derived Struct
```go
type Derived struct {
    // Core Derived Stats
    HPMax              float64  // Maximum health points
    Stamina            float64  // Physical stamina
    Speed              float64  // Movement speed
    Energies           map[string]float64  // Energy pools (MP, Qi, Lust, Wrath, etc.)
    Haste              float64  // Action speed multiplier
    CritChance         float64  // Critical hit chance
    CritMulti          float64  // Critical hit damage multiplier
    MoveSpeed          float64  // Movement speed
    RegenHP            float64  // Health regeneration rate
    RegenEnergies      map[string]float64  // Energy regeneration rates
    
    // Combat Stats
    Accuracy           float64  // Hit chance
    Penetration        float64  // Armor penetration
    Lethality          float64  // Lethal damage multiplier
    Brutality          float64  // Brutal damage multiplier
    ArmorClass         float64  // Armor class
    Evasion            float64  // Evasion chance
    BlockChance        float64  // Block chance
    ParryChance        float64  // Parry chance
    DodgeChance        float64  // Dodge chance
    
    // Energy Stats
    EnergyEfficiency   float64  // Energy efficiency multiplier
    EnergyCapacity     float64  // Energy capacity multiplier
    EnergyDrain        float64  // Energy drain rate
    ResourceRegen      float64  // Resource regeneration rate
    ResourceDecay      float64  // Resource decay rate
    
    // Learning Stats
    LearningRate       float64  // Learning rate multiplier
    Adaptation         float64  // Adaptation rate
    Memory             float64  // Memory capacity
    Experience         float64  // Experience gain multiplier
    
    // Social Stats
    Leadership         float64  // Leadership effectiveness
    Diplomacy          float64  // Diplomacy success rate
    Intimidation       float64  // Intimidation power
    Empathy            float64  // Empathy level
    Deception          float64  // Deception success rate
    Performance        float64  // Performance ability
    
    // Mystical Stats
    ManaEfficiency     float64  // Mana efficiency
    SpellPower         float64  // Spell power
    MysticResonance    float64  // Mystic resonance
    RealityBend        float64  // Reality bending ability
    TimeSense          float64  // Time sense
    SpaceSense         float64  // Space sense
    
    // Movement Stats
    JumpHeight         float64  // Jump height
    ClimbSpeed         float64  // Climbing speed
    SwimSpeed          float64  // Swimming speed
    FlightSpeed        float64  // Flight speed
    TeleportRange      float64  // Teleportation range
    Stealth            float64  // Stealth ability
    
    // Aura Stats
    AuraRadius         float64  // Aura radius
    AuraStrength       float64  // Aura strength
    Presence           float64  // Presence level
    Awe                float64  // Awe factor
    
    // Proficiency Stats
    WeaponMastery      float64  // Weapon mastery
    SkillLevel         float64  // Skill level
    LifeSteal          float64  // Life steal
    CastSpeed          float64  // Casting speed
    WeightCapacity     float64  // Weight capacity
    Persuasion         float64  // Persuasion ability
    MerchantPriceModifier float64  // Merchant price modifier
    FactionReputationGain float64  // Faction reputation gain
    
    // Talent Amplifiers
    CultivationSpeed   float64  // Cultivation speed multiplier
    EnergyEfficiency   float64  // Energy efficiency multiplier
    BreakthroughSuccess float64  // Breakthrough success rate
    SkillLearning      float64  // Skill learning rate
    CombatEffectiveness float64  // Combat effectiveness multiplier
    ResourceGathering  float64  // Resource gathering multiplier
    
    // Flexible Stats
    FlexibleStats *FlexibleStats
    SpeedSystem   *FlexibleSpeedSystem
}
```

### 2.3 FlexibleStats Struct
```go
type FlexibleStats struct {
    // Custom Primary Stats (shared across systems)
    CustomPrimary map[string]int64
    
    // Custom Derived Stats (shared across systems)
    CustomDerived map[string]float64
    
    // Subsystem-specific Stats
    SubSystemStats map[string]map[string]float64  // systemName -> statName -> value
}
```

## 3. Flexible Systems

### 3.1 FlexibleSpeedSystem
```go
type FlexibleSpeedSystem struct {
    // Speed Categories
    MovementSpeeds      map[string]float64  // walking, running, swimming, etc.
    CastingSpeeds       map[string]float64  // spell casting, technique casting, etc.
    CraftingSpeeds      map[string]float64  // alchemy, forging, enchanting, etc.
    LearningSpeeds      map[string]float64  // reading, studying, practicing, etc.
    CombatSpeeds        map[string]float64  // attack speed, defense speed, etc.
    SocialSpeeds        map[string]float64  // conversation, negotiation, etc.
    AdministrativeSpeeds map[string]float64  // administration, management, etc.
    
    // Speed Talent Bonuses
    SpeedTalentBonuses map[string]map[string]float64  // talent -> speedCategory -> bonus
}

// Methods
func (fss *FlexibleSpeedSystem) CalculateSpeed(category, speedType string, primaryStats map[string]int64) float64
func (fss *FlexibleSpeedSystem) GetSpeedCategories() []string
func (fss *FlexibleSpeedSystem) GetSpeedTypes(category string) []string
func (fss *FlexibleSpeedSystem) AddSpeedType(category, speedType string, formula string) error
func (fss *FlexibleSpeedSystem) RemoveSpeedType(category, speedType string) error
```

### 3.2 FlexibleKarmaSystem
```go
type FlexibleKarmaSystem struct {
    // Karma at different levels
    DivisionKarma map[string]map[string]int64  // divisionType -> divisionId -> karmaValue
    
    // Karma Categories
    KarmaCategories map[string][]string  // category -> karmaTypes
    
    // Karma Types
    KarmaTypes map[string]KarmaType  // karmaType -> definition
    
    // Karma Influence
    KarmaInfluence map[string]float64  // statName -> influenceMultiplier
}

type KarmaType struct {
    Name        string
    Category    string
    Description string
    Influence   map[string]float64  // statName -> influence
    DecayRate   float64
    MaxValue    int64
    MinValue    int64
}

// Methods
func (fks *FlexibleKarmaSystem) CalculateTotalKarma(divisionType, divisionId string) int64
func (fks *FlexibleKarmaSystem) CalculateWeightedKarmaScore(divisionType, divisionId string) float64
func (fks *FlexibleKarmaSystem) CalculateKarmaInfluence(statName string) float64
func (fks *FlexibleKarmaSystem) AddKarma(divisionType, divisionId, karmaType string, amount int64) error
func (fks *FlexibleKarmaSystem) RemoveKarma(divisionType, divisionId, karmaType string, amount int64) error
```

### 3.3 FlexibleAdministrativeSystem
```go
type FlexibleAdministrativeSystem struct {
    // Administrative Divisions
    Divisions map[string]map[string]AdministrativeDivision  // divisionType -> divisionId -> division
    
    // Division Types
    DivisionTypes map[string]DivisionType  // divisionType -> definition
    
    // Hierarchy
    Hierarchy map[string][]string  // parentDivision -> childDivisions
    
    // Relationships
    Relationships map[string]map[string]DivisionRelationship  // divisionId -> relatedDivision -> relationship
}

type AdministrativeDivision struct {
    ID          string
    Name        string
    Type        string
    Parent      string
    Children    []string
    Attributes  map[string]interface{}
    Level       int
    IsActive    bool
    CreatedAt   int64
    UpdatedAt   int64
}

type DivisionType struct {
    Name        string
    Category    string
    Description string
    Attributes  map[string]interface{}
    MaxLevel    int
    MinLevel    int
}

type DivisionRelationship struct {
    Type        string
    Strength    float64
    Description string
    IsActive    bool
}

// Methods
func (fas *FlexibleAdministrativeSystem) GetDivision(divisionType, divisionId string) (*AdministrativeDivision, error)
func (fas *FlexibleAdministrativeSystem) GetDivisionsByType(divisionType string) map[string]AdministrativeDivision
func (fas *FlexibleAdministrativeSystem) GetDivisionsByLevel(level int) []AdministrativeDivision
func (fas *FlexibleAdministrativeSystem) AddDivision(division *AdministrativeDivision) error
func (fas *FlexibleAdministrativeSystem) UpdateDivision(division *AdministrativeDivision) error
func (fas *FlexibleAdministrativeSystem) RemoveDivision(divisionType, divisionId string) error
```

## 4. Proficiency System

### 4.1 ProficiencySystem
```go
type ProficiencySystem struct {
    Proficiencies map[string]*Proficiency  // skillName -> proficiency
    Categories    map[string][]string      // category -> skillNames
    MaxSkills     int
}

type Proficiency struct {
    SkillName     string
    Category      string
    Level         int64
    Experience    int64
    MaxLevel      int64
    Multiplier    float64
    LastUsed      int64
    TotalUses     int64
}

// Methods
func (ps *ProficiencySystem) GetProficiency(skillName string) (*Proficiency, error)
func (ps *ProficiencySystem) AddExperience(skillName string, amount int64) error
func (ps *ProficiencySystem) LevelUp(skillName string) error
func (ps *ProficiencySystem) GetMultiplier(skillName string) float64
func (ps *ProficiencySystem) GetSkillsByCategory(category string) []string
func (ps *ProficiencySystem) AddSkill(skillName, category string) error
func (ps *ProficiencySystem) RemoveSkill(skillName string) error
```

## 5. Universal Skill System

### 5.1 UniversalSkillSystem
```go
type UniversalSkillSystem struct {
    Skills        map[string]*UniversalSkill  // skillName -> skill
    Categories    map[string][]string         // category -> skillNames
    SkillTrees    map[string]*SkillTree       // skillTreeName -> skillTree
    MaxSkills     int
}

type UniversalSkill struct {
    Name          string
    Category      string
    SubCategory   string
    Level         int64
    Experience    int64
    MaxLevel      int64
    Requirements  []string
    Bonuses       map[string]float64
    Cooldown      int64
    ManaCost      float64
    StaminaCost   float64
}

type SkillTree struct {
    Name        string
    Category    string
    Skills      map[string]*UniversalSkill
    Prerequisites map[string][]string  // skillName -> prerequisites
    Unlocks     map[string][]string    // skillName -> unlocks
}

// Methods
func (uss *UniversalSkillSystem) GetSkill(skillName string) (*UniversalSkill, error)
func (uss *UniversalSkillSystem) AddExperience(skillName string, amount int64) error
func (uss *UniversalSkillSystem) LevelUp(skillName string) error
func (uss *UniversalSkillSystem) GetSkillsByCategory(category string) []string
func (uss *UniversalSkillSystem) GetSkillTree(skillTreeName string) (*SkillTree, error)
func (uss *UniversalSkillSystem) AddSkill(skill *UniversalSkill) error
func (uss *UniversalSkillSystem) RemoveSkill(skillName string) error
```

## 6. Configuration System

### 6.1 ConfigurationManager
```go
type ConfigurationManager struct {
    stats       map[string]*StatDefinition
    formulas    map[string]*FormulaDefinition
    types       map[string]*TypeDefinition
    clamps      map[string]*ClampDefinition
    categories  map[string]*CategoryDefinition
    validations map[string]*ValidationRule
}

// Methods
func (cm *ConfigurationManager) LoadFromFile(filename string) error
func (cm *ConfigurationManager) SaveToFile(filename string) error
func (cm *ConfigurationManager) Validate() []ValidationError
func (cm *ConfigurationManager) GetStat(id string) (*StatDefinition, error)
func (cm *ConfigurationManager) GetFormula(id string) (*FormulaDefinition, error)
func (cm *ConfigurationManager) GetType(id string) (*TypeDefinition, error)
func (cm *ConfigurationManager) GetClamp(id string) (*ClampDefinition, error)
func (cm *ConfigurationManager) GetCategory(id string) (*CategoryDefinition, error)
func (cm *ConfigurationManager) AddStat(stat *StatDefinition) error
func (cm *ConfigurationManager) UpdateStat(stat *StatDefinition) error
func (cm *ConfigurationManager) DeleteStat(id string) error
```

### 6.2 FormulaEngine
```go
type FormulaEngine struct {
    compiledFormulas map[string]*CompiledFormula
    statCache        map[string]float64
    dependencyGraph  *DependencyGraph
    calculationOrder []string
}

type CompiledFormula struct {
    Expression    string
    Dependencies  []string
    CompiledCode  *CompiledCode
    CacheKey      string
    LastUpdate    int64
}

// Methods
func (fe *FormulaEngine) CompileFormulas() error
func (fe *FormulaEngine) CalculateStat(statID string, stats map[string]float64) float64
func (fe *FormulaEngine) CalculateAllStats(stats map[string]float64) map[string]float64
func (fe *FormulaEngine) InvalidateCache(statID string)
func (fe *FormulaEngine) ClearCache()
```

## 7. Multi-System Support

### 7.1 MultiSystemActor
```go
type MultiSystemActor struct {
    ActorCore        *ActorCore
    ActiveSystems    []string
    SystemLevels     map[string]int64
    SystemProgress   map[string]float64
    CrossSystemSynergies map[string]map[string]float64  // system1 -> system2 -> synergy
}

// Methods
func (msa *MultiSystemActor) ActivateSystem(systemName string) error
func (msa *MultiSystemActor) DeactivateSystem(systemName string) error
func (msa *MultiSystemActor) GetSystemLevel(systemName string) int64
func (msa *MultiSystemActor) SetSystemLevel(systemName string, level int64) error
func (msa *MultiSystemActor) GetSystemProgress(systemName string) float64
func (msa *MultiSystemActor) SetSystemProgress(systemName string, progress float64) error
func (msa *MultiSystemActor) GetCrossSystemSynergy(system1, system2 string) float64
func (msa *MultiSystemActor) SetCrossSystemSynergy(system1, system2 string, synergy float64) error
```

## 8. Utility Functions

### 8.1 Stat Calculation Utilities
```go
// Calculate derived stats from primary stats
func CalculateDerivedStats(primaryStats map[string]int64, formulas map[string]string) (map[string]float64, error)

// Calculate flexible stats
func CalculateFlexibleStats(primaryStats map[string]int64, flexibleStats *FlexibleStats) (*FlexibleStats, error)

// Calculate speed stats
func CalculateSpeedStats(primaryStats map[string]int64, speedSystem *FlexibleSpeedSystem) (*FlexibleSpeedSystem, error)

// Calculate karma stats
func CalculateKarmaStats(primaryStats map[string]int64, karmaSystem *FlexibleKarmaSystem) (*FlexibleKarmaSystem, error)
```

### 8.2 Validation Utilities
```go
// Validate stat values
func ValidateStatValue(statName string, value interface{}) error

// Validate formula
func ValidateFormula(formula string, dependencies []string) error

// Validate configuration
func ValidateConfiguration(config *Configuration) []ValidationError

// Validate actor core
func ValidateActorCore(actorCore *ActorCore) []ValidationError
```

### 8.3 Conversion Utilities
```go
// Convert between different stat formats
func ConvertStatFormat(stats map[string]interface{}, fromFormat, toFormat string) (map[string]interface{}, error)

// Convert between different systems
func ConvertBetweenSystems(stats map[string]interface{}, fromSystem, toSystem string) (map[string]interface{}, error)

// Convert between different versions
func ConvertBetweenVersions(stats map[string]interface{}, fromVersion, toVersion string) (map[string]interface{}, error)
```

## 9. Error Types

### 9.1 Custom Error Types
```go
type StatError struct {
    StatName string
    Message  string
    Cause    error
}

func (se *StatError) Error() string {
    return fmt.Sprintf("Stat error for %s: %s", se.StatName, se.Message)
}

type FormulaError struct {
    FormulaName string
    Message     string
    Cause       error
}

func (fe *FormulaError) Error() string {
    return fmt.Sprintf("Formula error for %s: %s", fe.FormulaName, fe.Message)
}

type ConfigurationError struct {
    ConfigName string
    Message    string
    Cause      error
}

func (ce *ConfigurationError) Error() string {
    return fmt.Sprintf("Configuration error for %s: %s", ce.ConfigName, ce.Message)
}
```

## 10. Constants

### 10.1 Stat Constants
```go
const (
    // Primary Stats
    STAT_VITALITY     = "vitality"
    STAT_ENDURANCE    = "endurance"
    STAT_CONSTITUTION = "constitution"
    STAT_INTELLIGENCE = "intelligence"
    STAT_WISDOM       = "wisdom"
    STAT_CHARISMA     = "charisma"
    STAT_WILLPOWER    = "willpower"
    STAT_LUCK         = "luck"
    STAT_FATE         = "fate"
    STAT_KARMA        = "karma"
    STAT_STRENGTH     = "strength"
    STAT_AGILITY      = "agility"
    STAT_PERSONALITY  = "personality"
    
    // Universal Cultivation Stats
    STAT_SPIRITUAL_ENERGY   = "spiritualEnergy"
    STAT_PHYSICAL_ENERGY    = "physicalEnergy"
    STAT_MENTAL_ENERGY      = "mentalEnergy"
    STAT_CULTIVATION_LEVEL  = "cultivationLevel"
    STAT_BREAKTHROUGH_POINTS = "breakthroughPoints"
)

// Derived Stats
const (
    STAT_HP_MAX              = "hpMax"
    STAT_STAMINA             = "stamina"
    STAT_SPEED               = "speed"
    STAT_HASTE               = "haste"
    STAT_CRIT_CHANCE         = "critChance"
    STAT_CRIT_MULTI          = "critMulti"
    STAT_ACCURACY            = "accuracy"
    STAT_PENETRATION         = "penetration"
    STAT_LETHALITY           = "lethality"
    STAT_BRUTALITY           = "brutality"
    STAT_ARMOR_CLASS         = "armorClass"
    STAT_EVASION             = "evasion"
    STAT_BLOCK_CHANCE        = "blockChance"
    STAT_PARRY_CHANCE        = "parryChance"
    STAT_DODGE_CHANCE        = "dodgeChance"
    STAT_ENERGY_EFFICIENCY   = "energyEfficiency"
    STAT_ENERGY_CAPACITY     = "energyCapacity"
    STAT_ENERGY_DRAIN        = "energyDrain"
    STAT_RESOURCE_REGEN      = "resourceRegen"
    STAT_RESOURCE_DECAY      = "resourceDecay"
    STAT_LEARNING_RATE       = "learningRate"
    STAT_ADAPTATION          = "adaptation"
    STAT_MEMORY              = "memory"
    STAT_EXPERIENCE          = "experience"
    STAT_LEADERSHIP          = "leadership"
    STAT_DIPLOMACY           = "diplomacy"
    STAT_INTIMIDATION        = "intimidation"
    STAT_EMPATHY             = "empathy"
    STAT_DECEPTION           = "deception"
    STAT_PERFORMANCE         = "performance"
    STAT_MANA_EFFICIENCY     = "manaEfficiency"
    STAT_SPELL_POWER         = "spellPower"
    STAT_MYSTIC_RESONANCE    = "mysticResonance"
    STAT_REALITY_BEND        = "realityBend"
    STAT_TIME_SENSE          = "timeSense"
    STAT_SPACE_SENSE         = "spaceSense"
    STAT_JUMP_HEIGHT         = "jumpHeight"
    STAT_CLIMB_SPEED         = "climbSpeed"
    STAT_SWIM_SPEED          = "swimSpeed"
    STAT_FLIGHT_SPEED        = "flightSpeed"
    STAT_TELEPORT_RANGE      = "teleportRange"
    STAT_STEALTH             = "stealth"
    STAT_AURA_RADIUS         = "auraRadius"
    STAT_AURA_STRENGTH       = "auraStrength"
    STAT_PRESENCE            = "presence"
    STAT_AWE                 = "awe"
    STAT_WEAPON_MASTERY      = "weaponMastery"
    STAT_SKILL_LEVEL         = "skillLevel"
    STAT_LIFE_STEAL          = "lifeSteal"
    STAT_CAST_SPEED          = "castSpeed"
    STAT_WEIGHT_CAPACITY     = "weightCapacity"
    STAT_PERSUASION          = "persuasion"
    STAT_MERCHANT_PRICE_MODIFIER = "merchantPriceModifier"
    STAT_FACTION_REPUTATION_GAIN = "factionReputationGain"
    
    // Talent Amplifiers
    STAT_CULTIVATION_SPEED   = "cultivationSpeed"
    STAT_ENERGY_EFFICIENCY   = "energyEfficiency"
    STAT_BREAKTHROUGH_SUCCESS = "breakthroughSuccess"
    STAT_SKILL_LEARNING      = "skillLearning"
    STAT_COMBAT_EFFECTIVENESS = "combatEffectiveness"
    STAT_RESOURCE_GATHERING  = "resourceGathering"
)
```

### 10.2 System Constants
```go
const (
    // System Names
    SYSTEM_RPG        = "rpg"
    SYSTEM_KIM_DAN    = "kim_dan"
    SYSTEM_MAGIC      = "magic"
    SYSTEM_SUCCUBUS   = "succubus"
    SYSTEM_BODY_CULTIVATION = "body_cultivation"
    SYSTEM_MARTIAL_ARTS = "martial_arts"
    SYSTEM_BEAST_TAMER = "beast_tamer"
    
    // Speed Categories
    SPEED_MOVEMENT      = "movement"
    SPEED_CASTING       = "casting"
    SPEED_CRAFTING      = "crafting"
    SPEED_LEARNING      = "learning"
    SPEED_COMBAT        = "combat"
    SPEED_SOCIAL        = "social"
    SPEED_ADMINISTRATIVE = "administrative"
    
    // Karma Categories
    KARMA_FORTUNE       = "fortune"
    KARMA_KARMA         = "karma"
    KARMA_MERIT         = "merit"
    KARMA_CONTRIBUTION  = "contribution"
    
    // Administrative Division Types
    DIVISION_WORLD      = "world"
    DIVISION_CONTINENT  = "continent"
    DIVISION_NATION     = "nation"
    DIVISION_RACE       = "race"
    DIVISION_SECT       = "sect"
    DIVISION_REALM      = "realm"
    DIVISION_ZONE       = "zone"
    DIVISION_CITY       = "city"
    DIVISION_VILLAGE    = "village"
)
```

---

*Tài liệu này cung cấp API Reference chi tiết cho Actor Core v2.0, bao gồm tất cả interfaces, structs, methods, và functions.*
