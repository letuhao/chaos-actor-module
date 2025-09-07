# Actor Core v2.0 Implementation Plan

## Tổng Quan

Kế hoạch implementation cho Actor Core v2.0 với Universal Cultivation Stats và tích hợp RPG System. Plan này chia thành 4 phases chính với timeline rõ ràng và deliverables cụ thể.

## Timeline Overview

```
Phase 1: Core Foundation (2-3 weeks)
├── Week 1: Actor Core v2.0 structure
├── Week 2: Basic interfaces & calculations
└── Week 3: Unit tests & validation

Phase 2: Integration Layer (2-3 weeks)
├── Week 1: Integration interfaces
├── Week 2: Stat mapping & conversion
└── Week 3: Integration tests

Phase 3: Sub-System Adapters (3-4 weeks)
├── Week 1: RPG System adapter
├── Week 2: Kim Đan System adapter
├── Week 3: Other sub-system adapters
└── Week 4: End-to-end testing

Phase 4: Optimization & Production (2-3 weeks)
├── Week 1: Performance optimization
├── Week 2: Caching & memory optimization
└── Week 3: Load testing & production readiness
```

## Phase 1: Core Foundation (2-3 weeks)

### Week 1: Actor Core v2.0 Structure

#### Day 1-2: Project Setup
```bash
# Tạo project structure
mkdir -p chaos-actor-module/packages/actor-core-v2
cd chaos-actor-module/packages/actor-core-v2

# Initialize Go module
go mod init actor-core-v2

# Tạo directory structure
mkdir -p internal/{core,interfaces,calculations,utils}
mkdir -p pkg/{statprovider,statconsumer,statresolver}
mkdir -p test/{unit,integration,performance}
mkdir -p docs/{api,examples,guides}
```

#### Day 3-4: Core Data Structures
```go
// internal/core/primarycore.go
type PrimaryCore struct {
    // Basic Stats
    Vitality     int64
    Endurance    int64
    Constitution int64
    Intelligence int64
    Wisdom       int64
    Charisma     int64
    Willpower    int64
    Luck         int64
    Fate         int64
    Karma        int64
    
    // Physical Stats
    Strength    int64
    Agility     int64
    Personality int64
    
    // Universal Cultivation Stats
    SpiritualEnergy   int64
    PhysicalEnergy    int64
    MentalEnergy      int64
    CultivationLevel  int64
    BreakthroughPoints int64
}

// internal/core/derived.go
type Derived struct {
    // Core Derived Stats
    HPMax           float64
    Stamina         float64
    Speed           float64
    Accuracy        float64
    Penetration     float64
    Lethality       float64
    Brutality       float64
    ArmorClass      float64
    Evasion         float64
    BlockChance     float64
    ParryChance     float64
    DodgeChance     float64
    EnergyEfficiency float64
    EnergyCapacity   float64
    EnergyDrain      float64
    ResourceRegen    float64
    ResourceDecay    float64
    StatusResistance float64
    Immunity         float64
    WeatherResist    float64
    TerrainMastery   float64
    ClimateAdapt     float64
    LearningRate     float64
    Adaptation       float64
    Memory           float64
    Experience       float64
    Leadership       float64
    Diplomacy        float64
    Intimidation     float64
    Empathy          float64
    Deception        float64
    Performance      float64
    ManaEfficiency   float64
    SpellPower       float64
    MysticResonance  float64
    RealityBend      float64
    TimeSense        float64
    SpaceSense       float64
    JumpHeight       float64
    ClimbSpeed       float64
    SwimSpeed        float64
    FlightSpeed      float64
    TeleportRange    float64
    Stealth          float64
    AuraRadius       float64
    AuraStrength     float64
    Presence         float64
    Awe              float64
    WeaponMastery    float64
    SkillLevel       float64
    LifeSteal        float64
    CastSpeed        float64
    WeightCapacity   float64
    Persuasion       float64
    MerchantPriceModifier float64
    FactionReputationGain float64
    
    // Talent Amplifiers
    CultivationSpeed    float64
    EnergyEfficiency    float64
    BreakthroughSuccess float64
    SkillLearning       float64
    CombatEffectiveness float64
    ResourceGathering   float64
    
    // Sub-System Derived Stats
    RPG     map[string]float64
    KimDan  map[string]float64
    Magic   map[string]float64
    Succubus map[string]float64
    LuyenThe map[string]float64
    VoHiep  map[string]float64
    NguThuSu map[string]float64
}
```

#### Day 5: Core Interfaces
```go
// internal/interfaces/statprovider.go
type StatProvider interface {
    GetPrimaryStats(actorID string) (*PrimaryCore, error)
    GetDerivedStats(actorID string) (*Derived, error)
    GetStat(actorID string, statKey string) (float64, error)
    GetStatBreakdown(actorID string, statKey string) (*StatBreakdown, error)
    UpdateStats(actorID string, updates *StatUpdates) error
}

// internal/interfaces/statconsumer.go
type StatConsumer interface {
    ConsumeStats(actorID string, stats *ActorCoreStats) error
    GetStatRequirements() []string
    ValidateStats(stats *ActorCoreStats) error
}

// internal/interfaces/statresolver.go
type StatResolver interface {
    ResolveStats(actorID string) (*ActorCoreStats, error)
    ResolveStat(actorID string, statKey string) (float64, error)
    GetStatDependencies(statKey string) []string
    CalculateDerivedStats(primary *PrimaryCore) (*Derived, error)
}
```

### Week 2: Basic Interfaces & Calculations

#### Day 1-2: Stat Calculation Engine
```go
// internal/calculations/calculator.go
type StatCalculator struct {
    formulas map[string]Formula
    cache    map[string]float64
}

type Formula interface {
    Calculate(primary *PrimaryCore) float64
    GetDependencies() []string
    GetDescription() string
}

// Implement specific formulas
type HPMaxFormula struct{}
func (f *HPMaxFormula) Calculate(primary *PrimaryCore) float64 {
    return (float64(primary.Vitality)*10 + 
            float64(primary.Endurance)*5 + 
            float64(primary.Constitution)*3) * 
           (1 + float64(primary.CultivationLevel)/100)
}

type SpeedFormula struct{}
func (f *SpeedFormula) Calculate(primary *PrimaryCore) float64 {
    return (float64(primary.Agility)*2 + 
            float64(primary.Intelligence)*1 + 
            float64(primary.Luck)*0.5) * 
           (1 + float64(primary.CultivationLevel)/200)
}
```

#### Day 3-4: Stat Resolver Implementation
```go
// internal/calculations/resolver.go
type StatResolverImpl struct {
    calculator *StatCalculator
    cache      *Cache
    mutex      sync.RWMutex
}

func (sr *StatResolverImpl) ResolveStats(actorID string) (*ActorCoreStats, error) {
    // Check cache first
    if cached, exists := sr.cache.Get(actorID); exists {
        return cached, nil
    }
    
    // Get primary stats
    primary, err := sr.getPrimaryStats(actorID)
    if err != nil {
        return nil, err
    }
    
    // Calculate derived stats
    derived, err := sr.calculator.CalculateAllDerived(primary)
    if err != nil {
        return nil, err
    }
    
    // Create result
    result := &ActorCoreStats{
        Primary: primary,
        Derived: derived,
        ActorID: actorID,
        Version: 1,
        Timestamp: time.Now().Unix(),
    }
    
    // Cache result
    sr.cache.Set(actorID, result)
    
    return result, nil
}
```

#### Day 5: Basic Unit Tests
```go
// test/unit/calculator_test.go
func TestHPMaxFormula(t *testing.T) {
    formula := &HPMaxFormula{}
    
    primary := &PrimaryCore{
        Vitality: 100,
        Endurance: 80,
        Constitution: 90,
        CultivationLevel: 5,
    }
    
    expected := (100.0*10 + 80.0*5 + 90.0*3) * (1 + 5.0/100)
    actual := formula.Calculate(primary)
    
    assert.Equal(t, expected, actual)
}

func TestSpeedFormula(t *testing.T) {
    formula := &SpeedFormula{}
    
    primary := &PrimaryCore{
        Agility: 90,
        Intelligence: 120,
        Luck: 75,
        CultivationLevel: 10,
    }
    
    expected := (90.0*2 + 120.0*1 + 75.0*0.5) * (1 + 10.0/200)
    actual := formula.Calculate(primary)
    
    assert.Equal(t, expected, actual)
}
```

### Week 3: Unit Tests & Validation

#### Day 1-2: Comprehensive Unit Tests
```go
// test/unit/primarycore_test.go
func TestPrimaryCoreValidation(t *testing.T) {
    tests := []struct {
        name    string
        primary PrimaryCore
        valid   bool
    }{
        {
            name: "Valid primary core",
            primary: PrimaryCore{
                Vitality: 100,
                Endurance: 80,
                // ... other stats
            },
            valid: true,
        },
        {
            name: "Invalid negative stats",
            primary: PrimaryCore{
                Vitality: -10,
                // ... other stats
            },
            valid: false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.primary.Validate()
            if tt.valid {
                assert.NoError(t, err)
            } else {
                assert.Error(t, err)
            }
        })
    }
}
```

#### Day 3-4: Integration Tests
```go
// test/integration/statresolver_test.go
func TestStatResolverIntegration(t *testing.T) {
    resolver := NewStatResolver()
    
    // Test complete stat resolution
    stats, err := resolver.ResolveStats("test_actor")
    assert.NoError(t, err)
    assert.NotNil(t, stats)
    assert.NotNil(t, stats.Primary)
    assert.NotNil(t, stats.Derived)
    
    // Test specific stat resolution
    hpMax, err := resolver.ResolveStat("test_actor", "HPMax")
    assert.NoError(t, err)
    assert.Greater(t, hpMax, 0.0)
}
```

#### Day 5: Performance Tests
```go
// test/performance/calculator_test.go
func BenchmarkStatCalculation(b *testing.B) {
    calculator := NewStatCalculator()
    primary := &PrimaryCore{
        // ... test data
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        calculator.CalculateAllDerived(primary)
    }
}

func BenchmarkStatResolver(b *testing.B) {
    resolver := NewStatResolver()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        resolver.ResolveStats("test_actor")
    }
}
```

## Phase 2: Integration Layer (2-3 weeks)

### Week 1: Integration Interfaces

#### Day 1-2: Sub-System Interface
```go
// internal/interfaces/subsystem.go
type SubSystem interface {
    GetName() string
    GetVersion() string
    GetStatMappings() map[string]string
    ConvertToActorCore(stats interface{}) (*ActorCoreContribution, error)
    ConvertFromActorCore(contribution *ActorCoreContribution) (interface{}, error)
    ValidateStats(stats interface{}) error
}

// internal/interfaces/statmapping.go
type StatMapping struct {
    SourceStat string
    TargetStat string
    Conversion Formula
    Priority   int
}

type StatMapper struct {
    mappings map[string][]StatMapping
    mutex    sync.RWMutex
}

func (sm *StatMapper) MapStats(sourceSystem string, sourceStats interface{}) (*ActorCoreContribution, error) {
    // Implementation
}
```

#### Day 3-4: Conversion Utilities
```go
// internal/utils/converter.go
type StatConverter struct {
    mappers map[string]StatMapper
}

func (sc *StatConverter) ConvertRPGToActorCore(rpgStats *RPGStats) (*ActorCoreContribution, error) {
    return &ActorCoreContribution{
        Primary: &PrimaryCore{
            Strength:    rpgStats.STR,
            Intelligence: rpgStats.INT,
            Willpower:   rpgStats.WIL,
            Agility:     rpgStats.AGI,
            Endurance:   rpgStats.END,
            Personality: rpgStats.PER,
            Luck:        rpgStats.LUK,
        },
        Derived: &Derived{
            HPMax: rpgStats.HP_MAX,
            // ... other mappings
        },
    }, nil
}

func (sc *StatConverter) ConvertKimDanToActorCore(kimDanStats *KimDanStats) (*ActorCoreContribution, error) {
    return &ActorCoreContribution{
        Primary: &PrimaryCore{
            Vitality:      kimDanStats.VitalEssence,
            Willpower:     kimDanStats.QiControl,
            Constitution:  kimDanStats.MeridianStrength,
            Endurance:     kimDanStats.BodyConstitution,
            Intelligence:  kimDanStats.SoulConsciousness,
            Wisdom:        kimDanStats.DaoComprehension,
            Luck:          kimDanStats.Fortune,
        },
        Derived: &Derived{
            // ... mappings
        },
    }, nil
}
```

#### Day 5: Integration Tests
```go
// test/integration/converter_test.go
func TestRPGToActorCoreConversion(t *testing.T) {
    converter := NewStatConverter()
    
    rpgStats := &RPGStats{
        STR: 100,
        INT: 120,
        WIL: 80,
        // ... other stats
    }
    
    contribution, err := converter.ConvertRPGToActorCore(rpgStats)
    assert.NoError(t, err)
    assert.NotNil(t, contribution)
    assert.Equal(t, int64(100), contribution.Primary.Strength)
    assert.Equal(t, int64(120), contribution.Primary.Intelligence)
}
```

### Week 2: Stat Mapping & Conversion

#### Day 1-2: Advanced Stat Mapping
```go
// internal/calculations/mapping.go
type StatMappingEngine struct {
    mappings map[string]map[string]StatMapping
    formulas map[string]Formula
}

func (sme *StatMappingEngine) RegisterMapping(sourceSystem, targetSystem string, mapping StatMapping) {
    if sme.mappings[sourceSystem] == nil {
        sme.mappings[sourceSystem] = make(map[string]StatMapping)
    }
    sme.mappings[sourceSystem][targetSystem] = mapping
}

func (sme *StatMappingEngine) MapStats(sourceSystem, targetSystem string, sourceStats interface{}) (interface{}, error) {
    mapping, exists := sme.mappings[sourceSystem][targetSystem]
    if !exists {
        return nil, fmt.Errorf("no mapping found from %s to %s", sourceSystem, targetSystem)
    }
    
    return mapping.Conversion.Convert(sourceStats), nil
}
```

#### Day 3-4: Cross-System Validation
```go
// internal/validation/crosssystem.go
type CrossSystemValidator struct {
    validators map[string]Validator
}

func (csv *CrossSystemValidator) ValidateCrossSystem(stats *ActorCoreStats, systems []string) error {
    for _, system := range systems {
        validator, exists := csv.validators[system]
        if !exists {
            continue
        }
        
        if err := validator.Validate(stats); err != nil {
            return fmt.Errorf("validation failed for system %s: %w", system, err)
        }
    }
    
    return nil
}
```

#### Day 5: Integration Tests
```go
// test/integration/mapping_test.go
func TestCrossSystemMapping(t *testing.T) {
    engine := NewStatMappingEngine()
    
    // Register mappings
    engine.RegisterMapping("RPG", "ActorCore", StatMapping{
        SourceStat: "STR",
        TargetStat: "Strength",
        Conversion: &DirectConversion{},
    })
    
    // Test mapping
    rpgStats := &RPGStats{STR: 100}
    actorCoreStats, err := engine.MapStats("RPG", "ActorCore", rpgStats)
    assert.NoError(t, err)
    assert.Equal(t, int64(100), actorCoreStats.Strength)
}
```

### Week 3: Integration Tests

#### Day 1-2: End-to-End Integration Tests
```go
// test/integration/e2e_test.go
func TestEndToEndIntegration(t *testing.T) {
    // Setup
    actorCore := NewActorCoreV2()
    rpgSystem := NewRPGSystem()
    kimDanSystem := NewKimDanSystem()
    
    // Test data
    actorID := "test_actor"
    
    // Test RPG System integration
    rpgStats, err := rpgSystem.GetStats(actorID)
    assert.NoError(t, err)
    
    contribution, err := actorCore.IntegrateRPGStats(rpgStats)
    assert.NoError(t, err)
    
    // Test Kim Dan System integration
    kimDanStats, err := kimDanSystem.GetStats(actorID)
    assert.NoError(t, err)
    
    contribution, err = actorCore.IntegrateKimDanStats(kimDanStats)
    assert.NoError(t, err)
    
    // Test combined stats
    finalStats, err := actorCore.ResolveStats(actorID)
    assert.NoError(t, err)
    assert.NotNil(t, finalStats)
}
```

#### Day 3-4: Performance Integration Tests
```go
// test/performance/integration_test.go
func BenchmarkIntegrationPerformance(b *testing.B) {
    actorCore := NewActorCoreV2()
    rpgSystem := NewRPGSystem()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        rpgStats, _ := rpgSystem.GetStats("test_actor")
        actorCore.IntegrateRPGStats(rpgStats)
    }
}
```

#### Day 5: Documentation & Examples
```go
// docs/examples/integration_example.go
func ExampleRPGIntegration() {
    // Create Actor Core v2.0
    actorCore := NewActorCoreV2()
    
    // Create RPG System
    rpgSystem := NewRPGSystem()
    
    // Get RPG stats
    rpgStats, err := rpgSystem.GetStats("player_001")
    if err != nil {
        log.Fatal(err)
    }
    
    // Integrate with Actor Core
    contribution, err := actorCore.IntegrateRPGStats(rpgStats)
    if err != nil {
        log.Fatal(err)
    }
    
    // Use integrated stats
    fmt.Printf("HP Max: %.2f\n", contribution.Derived.HPMax)
    fmt.Printf("Speed: %.2f\n", contribution.Derived.Speed)
}
```

## Phase 3: Sub-System Adapters (3-4 weeks)

### Week 1: RPG System Adapter

#### Day 1-2: RPG System Adapter Implementation
```go
// pkg/adapters/rpg_adapter.go
type RPGSystemAdapter struct {
    rpgSystem    *rpg_system.CoreActorIntegration
    actorCore    *ActorCoreV2
    converter    *StatConverter
}

func (rpg *RPGSystemAdapter) GetName() string {
    return "RPG System"
}

func (rpg *RPGSystemAdapter) GetVersion() string {
    return "1.0.0"
}

func (rpg *RPGSystemAdapter) ConvertToActorCore(stats interface{}) (*ActorCoreContribution, error) {
    rpgStats, ok := stats.(*rpg_system.RPGStats)
    if !ok {
        return nil, fmt.Errorf("invalid stats type")
    }
    
    return rpg.converter.ConvertRPGToActorCore(rpgStats)
}

func (rpg *RPGSystemAdapter) ConvertFromActorCore(contribution *ActorCoreContribution) (interface{}, error) {
    return rpg.converter.ConvertActorCoreToRPG(contribution)
}
```

#### Day 3-4: RPG System Integration
```go
// pkg/adapters/rpg_integration.go
type RPGIntegration struct {
    adapter   *RPGSystemAdapter
    actorCore *ActorCoreV2
}

func (rpg *RPGIntegration) SyncStats(actorID string) error {
    // Get RPG stats
    rpgStats, err := rpg.adapter.rpgSystem.GetStats(actorID)
    if err != nil {
        return err
    }
    
    // Convert to Actor Core
    contribution, err := rpg.adapter.ConvertToActorCore(rpgStats)
    if err != nil {
        return err
    }
    
    // Update Actor Core
    return rpg.actorCore.UpdateStats(actorID, contribution)
}

func (rpg *RPGIntegration) SyncFromActorCore(actorID string) error {
    // Get Actor Core stats
    stats, err := rpg.actorCore.GetStats(actorID)
    if err != nil {
        return err
    }
    
    // Convert to RPG
    rpgStats, err := rpg.adapter.ConvertFromActorCore(stats)
    if err != nil {
        return err
    }
    
    // Update RPG System
    return rpg.adapter.rpgSystem.UpdateStats(actorID, rpgStats)
}
```

#### Day 5: RPG System Tests
```go
// test/adapters/rpg_adapter_test.go
func TestRPGAdapter(t *testing.T) {
    adapter := NewRPGSystemAdapter()
    
    rpgStats := &rpg_system.RPGStats{
        STR: 100,
        INT: 120,
        WIL: 80,
        // ... other stats
    }
    
    contribution, err := adapter.ConvertToActorCore(rpgStats)
    assert.NoError(t, err)
    assert.NotNil(t, contribution)
    assert.Equal(t, int64(100), contribution.Primary.Strength)
}
```

### Week 2: Kim Đan System Adapter

#### Day 1-2: Kim Đan System Adapter
```go
// pkg/adapters/kimdan_adapter.go
type KimDanSystemAdapter struct {
    kimDanSystem *kimdan_system.CoreActorIntegration
    actorCore    *ActorCoreV2
    converter    *StatConverter
}

func (kd *KimDanSystemAdapter) GetName() string {
    return "Kim Dan System"
}

func (kd *KimDanSystemAdapter) ConvertToActorCore(stats interface{}) (*ActorCoreContribution, error) {
    kimDanStats, ok := stats.(*kimdan_system.KimDanStats)
    if !ok {
        return nil, fmt.Errorf("invalid stats type")
    }
    
    return kd.converter.ConvertKimDanToActorCore(kimDanStats)
}
```

#### Day 3-4: Kim Đan System Integration
```go
// pkg/adapters/kimdan_integration.go
type KimDanIntegration struct {
    adapter   *KimDanSystemAdapter
    actorCore *ActorCoreV2
}

func (kd *KimDanIntegration) SyncStats(actorID string) error {
    // Get Kim Dan stats
    kimDanStats, err := kd.adapter.kimDanSystem.GetStats(actorID)
    if err != nil {
        return err
    }
    
    // Convert to Actor Core
    contribution, err := kd.adapter.ConvertToActorCore(kimDanStats)
    if err != nil {
        return err
    }
    
    // Update Actor Core
    return kd.actorCore.UpdateStats(actorID, contribution)
}
```

#### Day 5: Kim Đan System Tests
```go
// test/adapters/kimdan_adapter_test.go
func TestKimDanAdapter(t *testing.T) {
    adapter := NewKimDanSystemAdapter()
    
    kimDanStats := &kimdan_system.KimDanStats{
        VitalEssence: 1000,
        QiControl: 800,
        MeridianStrength: 900,
        // ... other stats
    }
    
    contribution, err := adapter.ConvertToActorCore(kimDanStats)
    assert.NoError(t, err)
    assert.NotNil(t, contribution)
    assert.Equal(t, int64(1000), contribution.Primary.Vitality)
}
```

### Week 3: Other Sub-System Adapters

#### Day 1-2: Magic System Adapter
```go
// pkg/adapters/magic_adapter.go
type MagicSystemAdapter struct {
    magicSystem *magic_system.CoreActorIntegration
    actorCore   *ActorCoreV2
    converter   *StatConverter
}

func (magic *MagicSystemAdapter) GetName() string {
    return "Magic System"
}

func (magic *MagicSystemAdapter) ConvertToActorCore(stats interface{}) (*ActorCoreContribution, error) {
    magicStats, ok := stats.(*magic_system.MagicStats)
    if !ok {
        return nil, fmt.Errorf("invalid stats type")
    }
    
    return magic.converter.ConvertMagicToActorCore(magicStats)
}
```

#### Day 3-4: Succubus System Adapter
```go
// pkg/adapters/succubus_adapter.go
type SuccubusSystemAdapter struct {
    succubusSystem *succubus_system.CoreActorIntegration
    actorCore      *ActorCoreV2
    converter      *StatConverter
}

func (succubus *SuccubusSystemAdapter) GetName() string {
    return "Succubus System"
}

func (succubus *SuccubusSystemAdapter) ConvertToActorCore(stats interface{}) (*ActorCoreContribution, error) {
    succubusStats, ok := stats.(*succubus_system.SuccubusStats)
    if !ok {
        return nil, fmt.Errorf("invalid stats type")
    }
    
    return succubus.converter.ConvertSuccubusToActorCore(succubusStats)
}
```

#### Day 5: Other System Tests
```go
// test/adapters/all_adapters_test.go
func TestAllAdapters(t *testing.T) {
    adapters := []SubSystem{
        NewRPGSystemAdapter(),
        NewKimDanSystemAdapter(),
        NewMagicSystemAdapter(),
        NewSuccubusSystemAdapter(),
    }
    
    for _, adapter := range adapters {
        t.Run(adapter.GetName(), func(t *testing.T) {
            assert.NotEmpty(t, adapter.GetName())
            assert.NotEmpty(t, adapter.GetVersion())
        })
    }
}
```

### Week 4: End-to-End Testing

#### Day 1-2: Complete Integration Tests
```go
// test/integration/complete_integration_test.go
func TestCompleteIntegration(t *testing.T) {
    // Setup all systems
    actorCore := NewActorCoreV2()
    rpgSystem := NewRPGSystem()
    kimDanSystem := NewKimDanSystem()
    magicSystem := NewMagicSystem()
    
    // Test actor
    actorID := "test_actor"
    
    // Test all system integrations
    systems := []SubSystem{
        NewRPGSystemAdapter(),
        NewKimDanSystemAdapter(),
        NewMagicSystemAdapter(),
    }
    
    for _, system := range systems {
        t.Run(system.GetName(), func(t *testing.T) {
            // Get stats from system
            stats, err := system.GetStats(actorID)
            assert.NoError(t, err)
            
            // Convert to Actor Core
            contribution, err := system.ConvertToActorCore(stats)
            assert.NoError(t, err)
            
            // Update Actor Core
            err = actorCore.UpdateStats(actorID, contribution)
            assert.NoError(t, err)
        })
    }
    
    // Test final stats resolution
    finalStats, err := actorCore.ResolveStats(actorID)
    assert.NoError(t, err)
    assert.NotNil(t, finalStats)
}
```

#### Day 3-4: Performance End-to-End Tests
```go
// test/performance/e2e_performance_test.go
func BenchmarkCompleteIntegration(b *testing.B) {
    actorCore := NewActorCoreV2()
    rpgSystem := NewRPGSystem()
    kimDanSystem := NewKimDanSystem()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        // Test complete workflow
        rpgStats, _ := rpgSystem.GetStats("test_actor")
        kimDanStats, _ := kimDanSystem.GetStats("test_actor")
        
        actorCore.IntegrateRPGStats(rpgStats)
        actorCore.IntegrateKimDanStats(kimDanStats)
        
        actorCore.ResolveStats("test_actor")
    }
}
```

#### Day 5: Documentation & Examples
```go
// docs/examples/complete_integration_example.go
func ExampleCompleteIntegration() {
    // Create Actor Core v2.0
    actorCore := NewActorCoreV2()
    
    // Create all sub-systems
    rpgSystem := NewRPGSystem()
    kimDanSystem := NewKimDanSystem()
    magicSystem := NewMagicSystem()
    
    // Integrate all systems
    actorID := "player_001"
    
    // RPG System
    rpgStats, _ := rpgSystem.GetStats(actorID)
    actorCore.IntegrateRPGStats(rpgStats)
    
    // Kim Dan System
    kimDanStats, _ := kimDanSystem.GetStats(actorID)
    actorCore.IntegrateKimDanStats(kimDanStats)
    
    // Magic System
    magicStats, _ := magicSystem.GetStats(actorID)
    actorCore.IntegrateMagicStats(magicStats)
    
    // Get final integrated stats
    finalStats, _ := actorCore.ResolveStats(actorID)
    
    fmt.Printf("Final HP Max: %.2f\n", finalStats.Derived.HPMax)
    fmt.Printf("Final Speed: %.2f\n", finalStats.Derived.Speed)
    fmt.Printf("Cultivation Level: %d\n", finalStats.Primary.CultivationLevel)
}
```

## Phase 4: Optimization & Production (2-3 weeks)

### Week 1: Performance Optimization

#### Day 1-2: Caching System
```go
// internal/cache/stat_cache.go
type StatCache struct {
    cache    map[string]*CachedStats
    mutex    sync.RWMutex
    ttl      time.Duration
    maxSize  int
}

type CachedStats struct {
    Stats     *ActorCoreStats
    Timestamp time.Time
    TTL       time.Duration
}

func (sc *StatCache) Get(actorID string) (*ActorCoreStats, bool) {
    sc.mutex.RLock()
    defer sc.mutex.RUnlock()
    
    cached, exists := sc.cache[actorID]
    if !exists {
        return nil, false
    }
    
    if time.Since(cached.Timestamp) > cached.TTL {
        delete(sc.cache, actorID)
        return nil, false
    }
    
    return cached.Stats, true
}

func (sc *StatCache) Set(actorID string, stats *ActorCoreStats) {
    sc.mutex.Lock()
    defer sc.mutex.Unlock()
    
    // Check cache size
    if len(sc.cache) >= sc.maxSize {
        sc.evictOldest()
    }
    
    sc.cache[actorID] = &CachedStats{
        Stats:     stats,
        Timestamp: time.Now(),
        TTL:       sc.ttl,
    }
}
```

#### Day 3-4: Memory Optimization
```go
// internal/optimization/memory.go
type MemoryOptimizer struct {
    pool sync.Pool
}

func (mo *MemoryOptimizer) GetPrimaryCore() *PrimaryCore {
    if v := mo.pool.Get(); v != nil {
        return v.(*PrimaryCore)
    }
    return &PrimaryCore{}
}

func (mo *MemoryOptimizer) PutPrimaryCore(pc *PrimaryCore) {
    // Reset fields
    *pc = PrimaryCore{}
    mo.pool.Put(pc)
}

// internal/optimization/stat_pool.go
type StatPool struct {
    primaryPool sync.Pool
    derivedPool sync.Pool
}

func (sp *StatPool) GetPrimary() *PrimaryCore {
    if v := sp.primaryPool.Get(); v != nil {
        return v.(*PrimaryCore)
    }
    return &PrimaryCore{}
}

func (sp *StatPool) GetDerived() *Derived {
    if v := sp.derivedPool.Get(); v != nil {
        return v.(*Derived)
    }
    return &Derived{}
}
```

#### Day 5: Performance Tests
```go
// test/performance/optimization_test.go
func BenchmarkCachedStats(b *testing.B) {
    cache := NewStatCache()
    resolver := NewStatResolver()
    
    // Warm up cache
    resolver.ResolveStats("test_actor")
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        resolver.ResolveStats("test_actor")
    }
}

func BenchmarkMemoryOptimization(b *testing.B) {
    optimizer := NewMemoryOptimizer()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        pc := optimizer.GetPrimaryCore()
        // Use pc
        optimizer.PutPrimaryCore(pc)
    }
}
```

### Week 2: Caching & Memory Optimization

#### Day 1-2: Advanced Caching
```go
// internal/cache/advanced_cache.go
type AdvancedCache struct {
    lru    *lru.Cache
    ttl    time.Duration
    mutex  sync.RWMutex
}

func (ac *AdvancedCache) Get(actorID string) (*ActorCoreStats, bool) {
    ac.mutex.RLock()
    defer ac.mutex.RUnlock()
    
    if cached, exists := ac.lru.Get(actorID); exists {
        stats := cached.(*CachedStats)
        if time.Since(stats.Timestamp) <= stats.TTL {
            return stats.Stats, true
        }
        ac.lru.Remove(actorID)
    }
    
    return nil, false
}

func (ac *AdvancedCache) Set(actorID string, stats *ActorCoreStats) {
    ac.mutex.Lock()
    defer ac.mutex.Unlock()
    
    ac.lru.Add(actorID, &CachedStats{
        Stats:     stats,
        Timestamp: time.Now(),
        TTL:       ac.ttl,
    })
}
```

#### Day 3-4: Database Optimization
```go
// internal/database/optimized_db.go
type OptimizedDB struct {
    db        *sql.DB
    cache     *AdvancedCache
    batchSize int
}

func (odb *OptimizedDB) BatchUpdateStats(updates []StatUpdate) error {
    if len(updates) == 0 {
        return nil
    }
    
    // Batch update
    tx, err := odb.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    stmt, err := tx.Prepare("UPDATE actor_core_stats SET ? = ? WHERE actor_id = ?")
    if err != nil {
        return err
    }
    defer stmt.Close()
    
    for _, update := range updates {
        _, err = stmt.Exec(update.StatKey, update.Value, update.ActorID)
        if err != nil {
            return err
        }
    }
    
    return tx.Commit()
}
```

#### Day 5: Memory & Database Tests
```go
// test/performance/memory_db_test.go
func BenchmarkAdvancedCache(b *testing.B) {
    cache := NewAdvancedCache()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        cache.Get("test_actor")
    }
}

func BenchmarkBatchUpdate(b *testing.B) {
    db := NewOptimizedDB()
    updates := make([]StatUpdate, 1000)
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        db.BatchUpdateStats(updates)
    }
}
```

### Week 3: Load Testing & Production Readiness

#### Day 1-2: Load Testing
```go
// test/load/load_test.go
func TestLoadPerformance(t *testing.T) {
    actorCore := NewActorCoreV2()
    
    // Test with multiple actors
    numActors := 10000
    actors := make([]string, numActors)
    for i := 0; i < numActors; i++ {
        actors[i] = fmt.Sprintf("actor_%d", i)
    }
    
    // Concurrent stat resolution
    var wg sync.WaitGroup
    start := time.Now()
    
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for _, actorID := range actors {
                actorCore.ResolveStats(actorID)
            }
        }()
    }
    
    wg.Wait()
    duration := time.Since(start)
    
    t.Logf("Processed %d actors in %v", numActors*100, duration)
    t.Logf("Average time per actor: %v", duration/time.Duration(numActors*100))
}
```

#### Day 3-4: Production Configuration
```go
// internal/config/production.go
type ProductionConfig struct {
    CacheSize     int           `yaml:"cache_size"`
    CacheTTL      time.Duration `yaml:"cache_ttl"`
    BatchSize     int           `yaml:"batch_size"`
    MaxWorkers    int           `yaml:"max_workers"`
    DBConnections int           `yaml:"db_connections"`
}

func LoadProductionConfig() (*ProductionConfig, error) {
    config := &ProductionConfig{
        CacheSize:     10000,
        CacheTTL:      5 * time.Minute,
        BatchSize:     1000,
        MaxWorkers:    100,
        DBConnections: 50,
    }
    
    // Load from file or environment
    return config, nil
}
```

#### Day 5: Final Documentation
```go
// docs/production/README.md
# Actor Core v2.0 Production Guide

## Configuration
- Cache Size: 10,000 entries
- Cache TTL: 5 minutes
- Batch Size: 1,000 updates
- Max Workers: 100
- DB Connections: 50

## Performance Metrics
- Stat Resolution: < 1ms average
- Cache Hit Rate: > 95%
- Memory Usage: < 100MB
- Throughput: > 10,000 requests/second

## Monitoring
- Use Prometheus metrics
- Monitor cache hit rate
- Track memory usage
- Alert on performance degradation
```

## Deliverables

### Phase 1 Deliverables
- [ ] Actor Core v2.0 core structure
- [ ] Basic interfaces (StatProvider, StatConsumer, StatResolver)
- [ ] Stat calculation engine
- [ ] Unit tests (90%+ coverage)
- [ ] Basic documentation

### Phase 2 Deliverables
- [ ] Integration interfaces
- [ ] Stat mapping & conversion utilities
- [ ] Cross-system validation
- [ ] Integration tests
- [ ] Conversion examples

### Phase 3 Deliverables
- [ ] RPG System adapter
- [ ] Kim Đan System adapter
- [ ] Other sub-system adapters
- [ ] End-to-end integration tests
- [ ] Complete integration examples

### Phase 4 Deliverables
- [ ] Performance optimization
- [ ] Caching system
- [ ] Memory optimization
- [ ] Load testing results
- [ ] Production configuration
- [ ] Production documentation

## Risk Mitigation

### Technical Risks
1. **Performance Issues**: Implement caching and optimization early
2. **Memory Leaks**: Use object pooling and proper cleanup
3. **Integration Complexity**: Start with simple mappings, iterate
4. **Database Bottlenecks**: Implement batching and connection pooling

### Timeline Risks
1. **Scope Creep**: Stick to defined phases, defer nice-to-haves
2. **Integration Delays**: Start integration work early
3. **Testing Overhead**: Automate testing from day 1
4. **Documentation Debt**: Document as you go

### Quality Risks
1. **Code Quality**: Use linters and code review
2. **Test Coverage**: Maintain 90%+ coverage
3. **Performance Regression**: Continuous performance testing
4. **Integration Issues**: Comprehensive integration tests

## Success Criteria

### Technical Success
- [ ] All stats calculate correctly
- [ ] Performance meets requirements (< 1ms average)
- [ ] Memory usage is reasonable (< 100MB)
- [ ] Integration works with all sub-systems
- [ ] 90%+ test coverage

### Business Success
- [ ] Easy to integrate new sub-systems
- [ ] Clear separation of concerns
- [ ] Scalable architecture
- [ ] Maintainable codebase
- [ ] Comprehensive documentation

---

*Implementation plan này sẽ được cập nhật khi có thay đổi trong requirements hoặc timeline.*
