# Combat System Design v2.0

## ** Overview**

The Combat System is designed to be a flexible, high-performance system that supports complex damage calculations, status effects, talent integration, and fake real-time combat. It integrates seamlessly with Actor Core v2.0 and supports multiple cultivation systems.

## ** Key Design Principles**

1. **Flexibility First** - Easy to add new damage types, status effects, and combat mechanics
2. **Performance Optimized** - Concrete classes for common operations, caching for complex calculations
3. **Fake Real-time** - 60 FPS turn-based combat that feels real-time
4. **Unified System** - No distinction between PvP and PvE
5. **Actor Core Integration** - Leverages all Actor Core stats and systems
6. **Decorator Pattern** - Equipment and environment effects as decorators

## **🏗️ System Architecture**

### **Core Components**

```
Combat System
├── Damage Engine
│   ├── Damage Formula Engine (Flexible + Conditional)
│   ├── Damage Type System
│   └── Damage Calculation Engine
├── Status Effect System
│   ├── Status Effect Manager
│   ├── Elemental Interactions
│   └── Status Effect Stacking
├── Talent System Integration
│   ├── Talent → Amplifier Mapping
│   ├── Talent Conditions
│   └── Talent Leveling
├── Combat Engine
│   ├── Fake Real-time Loop (60 FPS)
│   ├── Action Queue System
│   └── Combat State Management
├── Equipment System
│   ├── Decorator Pattern
│   ├── Equipment Stats
│   └── Enchantment System
├── Environment System
│   ├── Terrain Effects
│   ├── Weather Effects
│   └── Time of Day Effects
└── Actor Core Integration
    ├── Combat Stats Mapping
    ├── Flexible Stats Support
    └── Cultivation System Integration
```

## **⚔️ Damage System**

### **Flexible Damage Formula System**

```go
type DamageFormula struct {
    ID          string                 `json:"id"`
    Name        string                 `json:"name"`
    BaseFormula string                 `json:"base_formula"`
    Conditions  []FormulaCondition     `json:"conditions"`
    Variables   map[string]FormulaVar  `json:"variables"`
    Modifiers   []FormulaModifier      `json:"modifiers"`
    CacheKey    string                 `json:"cache_key"`
}

type FormulaCondition struct {
    ID          string            `json:"id"`
    Type        ConditionType     `json:"type"`
    Variable    string            `json:"variable"`
    Operator    string            `json:"operator"`
    Value       interface{}       `json:"value"`
    Formula     string            `json:"formula"`
    ElseFormula string            `json:"else_formula"`
}

type ConditionType string
const (
    HealthCondition     ConditionType = "health"
    ManaCondition       ConditionType = "mana"
    StatusCondition     ConditionType = "status"
    EquipmentCondition  ConditionType = "equipment"
    EnvironmentCondition ConditionType = "environment"
    TimeCondition       ConditionType = "time"
    RandomCondition     ConditionType = "random"
)
```

### **Damage Types & Categories**

```go
type DamageType struct {
    ID          string         `json:"id"`
    Name        string         `json:"name"`
    Category    DamageCategory `json:"category"`
    Element     ElementType    `json:"element,omitempty"`
    Properties  []string       `json:"properties"`
    Formula     DamageFormula  `json:"formula"`
    Interactions []ElementInteraction `json:"interactions"`
}

type DamageCategory string
const (
    PhysicalCategory DamageCategory = "physical"
    MagicalCategory  DamageCategory = "magical"
    ElementalCategory DamageCategory = "elemental"
    TrueCategory     DamageCategory = "true"
    HealingCategory  DamageCategory = "healing"
    PoisonCategory   DamageCategory = "poison"
    CurseCategory    DamageCategory = "curse"
    DivineCategory   DamageCategory = "divine"
    DemonicCategory  DamageCategory = "demonic"
)
```

### **Performance-Optimized Damage Handlers**

```go
type DamageHandler interface {
    CalculateDamage(attack Attack, target Target) (*DamageResult, error)
    GetHandlerType() DamageType
}

// Concrete implementations for performance
type PhysicalDamageHandler struct {
    formula DamageFormula
    cache   *FormulaCache
}

type MagicalDamageHandler struct {
    formula DamageFormula
    cache   *FormulaCache
}

type ElementalDamageHandler struct {
    formula DamageFormula
    cache   *FormulaCache
    element ElementType
}
```

## ** Status Effect System**

### **Status Effect Definition**

```go
type StatusEffect struct {
    ID          string                `json:"id"`
    Name        string                `json:"name"`
    Type        StatusEffectType      `json:"type"`
    Element     ElementType           `json:"element"`
    Duration    int64                 `json:"duration"`
    Intensity   float64               `json:"intensity"`
    Stackable   bool                  `json:"stackable"`
    MaxStacks   int                   `json:"max_stacks"`
    Effects     []StatusEffectModifier `json:"effects"`
    Interactions []ElementInteraction  `json:"interactions"`
    Conditions  []StatusCondition     `json:"conditions"`
}

type StatusEffectType string
const (
    BuffEffect     StatusEffectType = "buff"
    DebuffEffect   StatusEffectType = "debuff"
    DoTEffect      StatusEffectType = "dot"
    HoTEffect      StatusEffectType = "hot"
    ControlEffect  StatusEffectType = "control"
    TransformEffect StatusEffectType = "transform"
)
```

### **Elemental Interactions**

```go
type ElementType string
const (
    FireElement     ElementType = "fire"
    WaterElement    ElementType = "water"
    EarthElement    ElementType = "earth"
    AirElement      ElementType = "air"
    LightningElement ElementType = "lightning"
    IceElement      ElementType = "ice"
    PoisonElement   ElementType = "poison"
    DarkElement     ElementType = "dark"
    LightElement    ElementType = "light"
    VoidElement     ElementType = "void"
)

type ElementInteraction struct {
    TargetElement ElementType `json:"target_element"`
    Result        string      `json:"result"`
    NewElement    ElementType `json:"new_element"`
    DamageMulti   float64     `json:"damage_multiplier"`
    DurationMulti float64     `json:"duration_multiplier"`
    Description   string      `json:"description"`
}
```

### **Status Effect Manager**

```go
type StatusManager struct {
    effects      map[string]*StatusEffect
    participants map[string][]*ActiveStatusEffect
    interactions map[string][]ElementInteraction
    mutex        sync.RWMutex
}

type ActiveStatusEffect struct {
    Effect      *StatusEffect `json:"effect"`
    StartTime   int64         `json:"start_time"`
    EndTime     int64         `json:"end_time"`
    Stacks      int           `json:"stacks"`
    Intensity   float64       `json:"intensity"`
    Source      string        `json:"source"`
    Target      string        `json:"target"`
}
```

## **🎯 Talent System Integration**

### **Talent → Amplifier Mapping**

```go
type Talent struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Tier        TalentTier        `json:"tier"`
    Amplifiers  []TalentAmplifier `json:"amplifiers"`
    Conditions  []TalentCondition `json:"conditions"`
    Prerequisites []string        `json:"prerequisites"`
    MaxLevel    int               `json:"max_level"`
    CurrentLevel int              `json:"current_level"`
}

type TalentAmplifier struct {
    ID          string            `json:"id"`
    Type        AmplifierType     `json:"type"`
    Value       float64           `json:"value"`
    PerLevel    float64           `json:"per_level"`
    Conditions  []AmplifierCondition `json:"conditions"`
    Stackable   bool              `json:"stackable"`
    MaxStacks   int               `json:"max_stacks"`
}

type AmplifierType string
const (
    MultiplierAmplifier AmplifierType = "multiplier"
    AdditionAmplifier   AmplifierType = "addition"
    PenetrationAmplifier AmplifierType = "penetration"
    ResistanceAmplifier AmplifierType = "resistance"
    CriticalAmplifier   AmplifierType = "critical"
    SpeedAmplifier      AmplifierType = "speed"
)
```

## **⚡ Fake Real-time Combat Engine**

### **Combat Engine Architecture**

```go
type CombatEngine struct {
    participants    map[string]*CombatParticipant
    actionQueue     *ActionQueue
    statusManager   *StatusManager
    damageHandler   *DamageHandlerFactory
    frameRate       int64 // 60 FPS
    frameDuration   time.Duration
    lastFrame       time.Time
    combatState     *CombatState
    eventBus        *EventBus
}

type ActionQueue struct {
    actions    []CombatAction
    mutex      sync.RWMutex
    maxSize    int
    processing bool
}

type CombatAction struct {
    ID          string            `json:"id"`
    ActorID     string            `json:"actor_id"`
    ActionType  ActionType        `json:"action_type"`
    TargetIDs   []string          `json:"target_ids"`
    DamageType  DamageType        `json:"damage_type"`
    BaseDamage  float64           `json:"base_damage"`
    Modifiers   map[string]float64 `json:"modifiers"`
    Priority    int               `json:"priority"`
    Timestamp   int64             `json:"timestamp"`
    Frame       int64             `json:"frame"`
    Cooldown    int64             `json:"cooldown"`
    ManaCost    float64           `json:"mana_cost"`
    StaminaCost float64           `json:"stamina_cost"`
}
```

### **60 FPS Combat Loop**

```go
func (ce *CombatEngine) StartCombat() {
    ce.frameDuration = time.Second / time.Duration(ce.frameRate)
    ce.lastFrame = time.Now()
    
    go func() {
        ticker := time.NewTicker(ce.frameDuration)
        defer ticker.Stop()
        
        for {
            select {
            case <-ticker.C:
                ce.processFrame()
            case <-ce.combatState.Done:
                return
            }
        }
    }()
}

func (ce *CombatEngine) processFrame() {
    currentTime := time.Now()
    frameDelta := currentTime.Sub(ce.lastFrame)
    
    if frameDelta < ce.frameDuration {
        return // Skip frame if too early
    }
    
    ce.lastFrame = currentTime
    ce.combatState.CurrentFrame++
    
    // Process actions in this frame
    ce.processActions()
    
    // Process status effects
    ce.statusManager.ProcessEffects(ce.combatState.CurrentFrame)
    
    // Check combat end conditions
    ce.checkCombatEnd()
}
```

## **🛡️ Unified PvP/PvE System**

### **Combat Participants**

```go
type CombatParticipant struct {
    ID          string                     `json:"id"`
    Name        string                     `json:"name"`
    Type        ParticipantType            `json:"type"`
    PrimaryStats *actorcore.PrimaryCore    `json:"primary_stats"`
    DerivedStats *actorcore.DerivedStats   `json:"derived_stats"`
    FlexibleStats *actorcore.FlexibleStats `json:"flexible_stats"`
    Equipment   *EquipmentSystem           `json:"equipment"`
    StatusEffects []StatusEffect           `json:"status_effects"`
    Talents     []Talent                   `json:"talents"`
    Amplifiers  []Amplifier                `json:"amplifiers"`
    Position    Position                   `json:"position"`
    Health      float64                    `json:"health"`
    Mana        float64                    `json:"mana"`
    Stamina     float64                    `json:"stamina"`
    IsAlive     bool                       `json:"is_alive"`
    IsPlayer    bool                       `json:"is_player"`
    AI          *AIController              `json:"ai,omitempty"`
}

type ParticipantType string
const (
    PlayerType    ParticipantType = "player"
    MonsterType   ParticipantType = "monster"
    NPCType       ParticipantType = "npc"
    BossType      ParticipantType = "boss"
    SummonType    ParticipantType = "summon"
)
```

## **⚔️ Equipment System (Decorator Pattern)**

### **Equipment Decorators**

```go
type EquipmentSystem struct {
    Weapon      *WeaponDecorator      `json:"weapon"`
    Armor       *ArmorDecorator       `json:"armor"`
    Accessories []*AccessoryDecorator `json:"accessories"`
    Enchantments []*EnchantmentDecorator `json:"enchantments"`
}

type EquipmentDecorator interface {
    GetStats() map[string]float64
    GetModifiers() map[string]float64
    GetEffects() []StatusEffect
    GetDamageTypes() []DamageType
    GetResistances() map[string]float64
    GetAmplifiers() []Amplifier
    GetConditions() []EquipmentCondition
}

type WeaponDecorator struct {
    base        *BaseWeapon
    enchantments []*EnchantmentDecorator
    modifications []*ModificationDecorator
}

func (w *WeaponDecorator) GetStats() map[string]float64 {
    stats := w.base.GetStats()
    
    // Apply enchantments
    for _, enchant := range w.enchantments {
        enchantStats := enchant.GetStats()
        for key, value := range enchantStats {
            stats[key] += value
        }
    }
    
    // Apply modifications
    for _, mod := range w.modifications {
        modStats := mod.GetStats()
        for key, value := range modStats {
            stats[key] += value
        }
    }
    
    return stats
}
```

## **🌍 Environment System**

### **Environment Decorators**

```go
type EnvironmentDecorator struct {
    terrain     TerrainType
    weather     WeatherType
    timeOfDay   TimeOfDay
    obstacles   []Obstacle
    modifiers   map[string]float64
}

type TerrainType string
const (
    DesertTerrain  TerrainType = "desert"
    ForestTerrain  TerrainType = "forest"
    MountainTerrain TerrainType = "mountain"
    WaterTerrain   TerrainType = "water"
    UrbanTerrain   TerrainType = "urban"
    DungeonTerrain TerrainType = "dungeon"
)

type WeatherType string
const (
    ClearWeather   WeatherType = "clear"
    RainWeather    WeatherType = "rain"
    SnowWeather    WeatherType = "snow"
    StormWeather   WeatherType = "storm"
    FogWeather     WeatherType = "fog"
    SandstormWeather WeatherType = "sandstorm"
)
```

## **🔗 Actor Core Integration**

### **Combat Stats Mapping**

```go
type CombatStats struct {
    // Từ PrimaryCore
    Strength      int64   `json:"strength"`
    Agility       int64   `json:"agility"`
    Intelligence  int64   `json:"intelligence"`
    Constitution  int64   `json:"constitution"`
    Vitality      int64   `json:"vitality"`
    Endurance     int64   `json:"endurance"`
    Willpower     int64   `json:"willpower"`
    Luck          int64   `json:"luck"`
    Fate          int64   `json:"fate"`
    Karma         int64   `json:"karma"`
    
    // Từ DerivedStats
    HPMax         float64 `json:"hp_max"`
    Stamina       float64 `json:"stamina"`
    Speed         float64 `json:"speed"`
    CritChance    float64 `json:"crit_chance"`
    CritMulti     float64 `json:"crit_multi"`
    Accuracy      float64 `json:"accuracy"`
    Penetration   float64 `json:"penetration"`
    Lethality     float64 `json:"lethality"`
    Brutality     float64 `json:"brutality"`
    ArmorClass    float64 `json:"armor_class"`
    Evasion       float64 `json:"evasion"`
    BlockChance   float64 `json:"block_chance"`
    ParryChance   float64 `json:"parry_chance"`
    DodgeChance   float64 `json:"dodge_chance"`
    
    // Từ FlexibleStats - Custom combat stats
    CustomCombat  map[string]float64 `json:"custom_combat"`
    
    // Từ SubSystemStats - Cultivation system stats
    CultivationStats map[string]map[string]float64 `json:"cultivation_stats"`
}
```

## **📊 Performance Considerations**

### **Caching Strategy**

```go
type FormulaCache struct {
    cache    map[string]*CachedFormula
    mutex    sync.RWMutex
    maxSize  int
    ttl      time.Duration
}

type CachedFormula struct {
    Formula    *DamageFormula
    Result     float64
    Timestamp  int64
    ExpiresAt  int64
}

type CombatCache struct {
    damageCache    *FormulaCache
    statusCache    *StatusCache
    equipmentCache *EquipmentCache
    talentCache    *TalentCache
}
```

### **Batch Processing**

```go
type CombatBatch struct {
    Actions    []CombatAction
    StatusUpdates []StatusUpdate
    DamageResults []DamageResult
    Frame      int64
    Timestamp  int64
}

func (ce *CombatEngine) ProcessBatch(batch *CombatBatch) error {
    // Process all actions in batch
    for _, action := range batch.Actions {
        if err := ce.processAction(action); err != nil {
            return err
        }
    }
    
    // Process status updates
    for _, update := range batch.StatusUpdates {
        if err := ce.statusManager.ApplyUpdate(update); err != nil {
            return err
        }
    }
    
    return nil
}
```

## **🧪 Testing Strategy**

### **Unit Tests**

- **Damage Formula Tests** - Test formula calculations and conditions
- **Status Effect Tests** - Test status application and interactions
- **Talent System Tests** - Test talent → amplifier mapping
- **Combat Engine Tests** - Test frame processing and action queue
- **Equipment Tests** - Test decorator pattern functionality

### **Integration Tests**

- **Actor Core Integration** - Test stats mapping and conversion
- **Multi-System Tests** - Test multiple cultivation systems
- **Performance Tests** - Test 60 FPS performance under load
- **Combat Flow Tests** - Test complete combat scenarios

### **Load Tests**

- **High Participant Count** - Test with 100+ participants
- **Complex Status Effects** - Test with many active status effects
- **Heavy Calculations** - Test with complex damage formulas
- **Memory Usage** - Test memory consumption over time

## **📈 Implementation Phases**

### **Phase 1: Core System (Week 1-2)**
1. **Damage Formula Engine** - Flexible formula system với conditional support
2. **Status Effect System** - Elemental interactions và status management
3. **Combat Engine** - Fake real-time combat với 60 FPS
4. **Actor Core Integration** - Convert stats từ Actor Core

### **Phase 2: Advanced Features (Week 3-4)**
1. **Talent System** - Talent → Amplifier integration
2. **Equipment System** - Decorator pattern cho equipment
3. **Environment System** - Terrain, weather, time effects
4. **AI System** - AI cho monsters và NPCs

### **Phase 3: Polish & Optimization (Week 5-6)**
1. **Performance Optimization** - Caching, batch processing
2. **Advanced Interactions** - Complex status effect interactions
3. **Combat Analytics** - Detailed combat statistics
4. **Testing & Documentation** - Comprehensive test suite

## **❓ Questions for Discussion**

1. **Damage Formula Complexity** - How complex should conditional formulas be? Should we support nested conditions?
2. **Status Effect Interactions** - How many elemental interactions should we support? Should we have a visual representation?
3. **Talent System Integration** - Should talents be part of Actor Core or separate? How do we handle talent progression?
4. **Performance vs Flexibility** - What's the acceptable performance impact for flexibility? Should we have different performance tiers?
5. **Real-time vs Turn-based** - How do we handle network latency in fake real-time? Should we have client-side prediction?
6. **PvP vs PvE** - Are there any special considerations for PvP balance? Should we have different damage formulas?
7. **Equipment Integration** - How complex should equipment effects be? Should we support equipment sets?
8. **Cultivation System Integration** - How do we handle different cultivation systems in combat? Should we have system-specific combat mechanics?

## **🎯 Next Steps**

1. **Update Actor Core** - Add status effect support to Actor Core
2. **Implement Status Module** - Create status effect system
3. **Design Damage Formulas** - Create flexible damage formula system
4. **Implement Combat Engine** - Create fake real-time combat system
5. **Integration Testing** - Test with Actor Core v2.0

---

*This document will be updated as the system evolves and new requirements are identified.*
