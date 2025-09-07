# RPG System Integration với Actor Core v2.0

## Tổng Quan

Tài liệu này mô tả cách tích hợp RPG System vào Actor Core v2.0, sử dụng Universal Cultivation Stats và loại bỏ level progression khỏi Actor Core.

## Kiến Trúc Tích Hợp

### 1. Actor Core v2.0 (Universal Foundation)
```
Actor Core v2.0
├── PrimaryCore (Universal Stats)
│   ├── Basic Stats: Vitality, Endurance, Constitution, Intelligence, Wisdom, Charisma, Willpower, Luck, Fate, Karma
│   ├── Physical Stats: Strength, Agility, Personality
│   └── Universal Cultivation Stats: SpiritualEnergy, PhysicalEnergy, MentalEnergy, CultivationLevel, BreakthroughPoints
├── Derived (Calculated Stats)
│   ├── Core Derived: HPMax, Stamina, Speed, Accuracy, Penetration, etc.
│   ├── Talent Amplifiers: CultivationSpeed, EnergyEfficiency, BreakthroughSuccess, etc.
│   └── RPG Derived: LifeSteal, CastSpeed, WeightCapacity, Persuasion, etc.
└── Interfaces
    ├── StatProvider (cho sub-systems)
    ├── StatConsumer (cho game systems)
    └── StatResolver (cho calculations)
```

### 2. RPG System (Sub-System)
```
RPG System
├── Level Progression (XP, Level, Stat Allocation)
├── Stat Registry (Primary/Derived definitions)
├── Modifier System (Equipment, Effects, Buffs)
├── Resolver Engine (Stat calculations)
└── Integration Layer (Actor Core communication)
```

## Mapping Stats

### Primary Stats Mapping
```go
// RPG Primary Stats → Actor Core PrimaryCore
STR  → Strength      (Actor Core PrimaryCore)
INT  → Intelligence  (Actor Core PrimaryCore)
WIL  → Willpower     (Actor Core PrimaryCore)
AGI  → Agility       (Actor Core PrimaryCore)
SPD  → Speed         (Derived from Agility + Intelligence + Luck)
END  → Endurance     (Actor Core PrimaryCore)
PER  → Personality   (Actor Core PrimaryCore)
LUK  → Luck          (Actor Core PrimaryCore)
```

### Derived Stats Mapping
```go
// RPG Derived Stats → Actor Core Derived
HP_MAX      → HPMax           (Actor Core Derived)
MANA_MAX    → Energies["MP"]  (Actor Core Energies map)
ATK         → Damages["Physical"] (Actor Core Damages map)
MATK        → Damages["Magical"]  (Actor Core Damages map)
DEF         → Defences["Physical"] (Actor Core Defences map)
EVASION     → Evasion         (Actor Core Derived)
MOVE_SPEED  → MoveSpeed       (Actor Core Derived)
CRIT_CHANCE → CritChance      (Actor Core Derived)
CRIT_DAMAGE → CritMulti       (Actor Core Derived)
```

### Flexible Stats Mapping
```go
// RPG System có thể sử dụng Flexible Stats
type RPGFlexibleStats struct {
    // Custom Primary Stats (shared with other systems)
    CustomPrimary map[string]int64{
        "SwordProficiency": 150,    // Shared with Combat system
        "MagicAffinity": 120,       // Shared with Magic system
        "TradingSkill": 80,         // Shared with Social system
    }
    
    // Custom Derived Stats (shared with other systems)
    CustomDerived map[string]float64{
        "SwordDamage": 1.5,         // From SwordProficiency
        "MagicEfficiency": 1.2,     // From MagicAffinity
        "MerchantDiscount": 0.1,    // From TradingSkill
    }
    
    // RPG System Specific Stats (independent)
    SubSystemStats map[string]map[string]float64{
        "RPG": {
            "LifeSteal": 0.05,
            "CastSpeed": 1.0,
            "WeightCapacity": 100.0,
            "Persuasion": 0.8,
            "MerchantPriceModifier": 0.9,
            "FactionReputationGain": 1.2,
        },
    }
}
```

## Multi-System Support

### 1. Multi-System Actor Example
```go
// Actor có thể tu luyện cả RPG và Kim Đan
type MultiSystemActor struct {
    ActorID        string
    PrimaryCore    *PrimaryCore
    Derived        *Derived
    ActiveSystems  map[string]bool  // systemName -> isActive
    SystemLevels   map[string]int64 // systemName -> level
    SystemProgress map[string]int64 // systemName -> progress
}

// Ví dụ: Actor tu luyện cả RPG và Kim Đan
func (msa *MultiSystemActor) CultivateMultipleSystems() {
    // RPG System
    msa.ActiveSystems["RPG"] = true
    msa.SystemLevels["RPG"] = 25
    msa.SystemProgress["RPG"] = 15000
    
    // Kim Dan System  
    msa.ActiveSystems["KimDan"] = true
    msa.SystemLevels["KimDan"] = 3  // Golden Core
    msa.SystemProgress["KimDan"] = 5000
    
    // Magic System
    msa.ActiveSystems["Magic"] = true
    msa.SystemLevels["Magic"] = 15
    msa.SystemProgress["Magic"] = 8000
}
```

### 2. Shared Custom Stats Example
```go
// Fire Mastery được chia sẻ giữa Magic và Alchemy
type SharedCustomStats struct {
    // Fire Mastery - shared between Magic and Alchemy systems
    FireMastery int64  // CustomPrimary["FireMastery"]
    
    // Derived from Fire Mastery
    FireDamage      float64  // CustomDerived["FireDamage"]
    FireResistance  float64  // CustomDerived["FireResistance"]
    AlchemyFireBonus float64 // CustomDerived["AlchemyFireBonus"]
}

// Magic System sử dụng Fire Mastery
func (magic *MagicSystem) CalculateFireSpellDamage(actor *MultiSystemActor) float64 {
    fireMastery := actor.PrimaryCore.FlexibleStats.CustomPrimary["FireMastery"]
    baseDamage := magic.GetBaseFireDamage()
    return baseDamage * (1.0 + float64(fireMastery)/100.0)
}

// Alchemy System cũng sử dụng Fire Mastery
func (alchemy *AlchemySystem) CalculateFirePotionQuality(actor *MultiSystemActor) float64 {
    fireMastery := actor.PrimaryCore.FlexibleStats.CustomPrimary["FireMastery"]
    baseQuality := alchemy.GetBasePotionQuality()
    return baseQuality * (1.0 + float64(fireMastery)/200.0)
}
```

### 3. Cross-System Synergies
```go
// Synergies giữa các hệ thống
type CrossSystemSynergy struct {
    Systems     []string  // Các hệ thống liên quan
    BonusType   string    // Loại bonus
    BonusValue  float64   // Giá trị bonus
    Requirements map[string]int64  // Yêu cầu tối thiểu cho mỗi hệ thống
}

// Ví dụ: Dual System Mastery
func (css *CrossSystemSynergy) CalculateDualSystemBonus(actor *MultiSystemActor) float64 {
    // Kiểm tra yêu cầu
    for system, minLevel := range css.Requirements {
        if actor.SystemLevels[system] < minLevel {
            return 0.0
        }
    }
    
    // Tính bonus dựa trên level của các hệ thống
    totalLevel := int64(0)
    for _, system := range css.Systems {
        totalLevel += actor.SystemLevels[system]
    }
    
    return css.BonusValue * float64(totalLevel) / 100.0
}
```

## Integration Design

### 1. Stat Provider Interface
```go
type StatProvider interface {
    // Get primary stats for Actor Core
    GetPrimaryStats(actorID string) (*PrimaryCoreStats, error)
    
    // Get derived stats for Actor Core
    GetDerivedStats(actorID string) (*DerivedStats, error)
    
    // Get stat breakdown for debugging
    GetStatBreakdown(actorID string, statKey string) (*StatBreakdown, error)
    
    // Update stats from RPG System
    UpdateStats(actorID string, updates *RPGStatUpdates) error
}
```

### 2. RPG System Integration
```go
type RPGSystemIntegration struct {
    // RPG System components
    registry    *registry.StatRegistry
    resolver    *resolver.StatResolver
    progression *progression.ProgressionService
    
    // Actor Core integration
    statProvider StatProvider
    actorCore    ActorCore
}

// Convert RPG stats to Actor Core format
func (rsi *RPGSystemIntegration) ConvertToActorCore(rpgStats *RPGStatSnapshot) *ActorCoreContribution {
    return &ActorCoreContribution{
        Primary: rsi.mapPrimaryStats(rpgStats),
        Derived: rsi.mapDerivedStats(rpgStats),
        Energies: rsi.mapEnergies(rpgStats),
        Damages: rsi.mapDamages(rpgStats),
        Defences: rsi.mapDefences(rpgStats),
        Amplifiers: rsi.mapAmplifiers(rpgStats),
    }
}
```

### 3. Level Progression (RPG System Only)
```go
type RPGProgressionService struct {
    // XP and level management
    xpTable      map[int64]int64
    levelCurves  map[string]*LevelCurve
    
    // Stat allocation
    statPoints   map[string]int64
    allocations  map[string]map[StatKey]int64
    
    // Database integration
    adapter      DatabaseAdapter
}

// Level progression methods
func (rps *RPGProgressionService) GrantXP(actorID string, xp int64) (*ProgressionResult, error)
func (rps *RPGProgressionService) AllocatePoints(actorID string, stat StatKey, points int64) error
func (rps *RPGProgressionService) Respec(actorID string) error
```

## Implementation Plan

### Phase 1: Actor Core v2.0 Foundation
1. **Tạo Actor Core v2.0 mới** với Universal Cultivation Stats
2. **Định nghĩa interfaces** cho StatProvider, StatConsumer, StatResolver
3. **Implement basic stat calculations** cho PrimaryCore và Derived
4. **Tạo test suite** cho Actor Core v2.0

### Phase 2: RPG System Integration
1. **Refactor RPG System** để loại bỏ level progression
2. **Tạo integration layer** giữa RPG System và Actor Core
3. **Implement stat mapping** từ RPG sang Actor Core
4. **Tạo conversion utilities** cho stat formats

### Phase 3: Level Progression System
1. **Tạo RPG Progression Service** độc lập
2. **Implement XP và level management**
3. **Tạo stat allocation system**
4. **Integrate với database**

### Phase 4: Testing & Validation
1. **Unit tests** cho từng component
2. **Integration tests** giữa RPG System và Actor Core
3. **Performance tests** cho stat calculations
4. **End-to-end tests** cho complete workflow

## Database Schema

### Actor Core Stats
```json
{
  "actorId": "player_001",
  "primaryCore": {
    "vitality": 100,
    "endurance": 80,
    "constitution": 90,
    "intelligence": 120,
    "wisdom": 110,
    "charisma": 85,
    "willpower": 95,
    "luck": 75,
    "fate": 60,
    "karma": 50,
    "strength": 85,
    "agility": 90,
    "personality": 80,
    "spiritualEnergy": 1000,
    "physicalEnergy": 800,
    "mentalEnergy": 1200,
    "cultivationLevel": 5,
    "breakthroughPoints": 50
  },
  "derived": {
    "hpMax": 1500,
    "stamina": 800,
    "speed": 120,
    "cultivationSpeed": 1.2,
    "energyEfficiency": 1.1,
    "lifeSteal": 0.05,
    "castSpeed": 1.0
  }
}
```

### RPG System Progress
```json
{
  "actorId": "player_001",
  "level": 25,
  "xp": 15000,
  "allocations": {
    "STR": 15,
    "INT": 20,
    "WIL": 12,
    "AGI": 18,
    "SPD": 16,
    "END": 14,
    "PER": 10,
    "LUK": 8
  },
  "equipment": [...],
  "effects": [...],
  "titles": [...]
}
```

## Benefits

### 1. Separation of Concerns
- **Actor Core**: Universal stats, không có level progression
- **RPG System**: Level progression, equipment, effects
- **Clear boundaries** giữa các systems

### 2. Scalability
- **Universal Cultivation Stats** hỗ trợ nhiều hệ thống tu luyện
- **Modular design** cho phép thêm sub-systems mới
- **Flexible stat system** cho customization

### 3. Performance
- **Deterministic calculations** cho consistency
- **Efficient caching** với hash-based snapshots
- **Optimized algorithms** cho stat resolution

### 4. Maintainability
- **Clear interfaces** giữa các components
- **Comprehensive testing** cho reliability
- **Documentation** cho development

## Next Steps

1. **Review và approve** design này
2. **Tạo Actor Core v2.0** implementation
3. **Refactor RPG System** để tích hợp
4. **Implement level progression** system
5. **Testing và validation**

---

*Tài liệu này sẽ được cập nhật khi có thay đổi trong design hoặc implementation.*
