# Actor Core v2.0 Design Summary

## Tổng Quan

Actor Core v2.0 được thiết kế lại để hỗ trợ **Multi-System Cultivation** với các tính năng mới:

1. **Flexible Stats Maps** - Cho phép custom stats được chia sẻ giữa các hệ thống
2. **Proficiency System** - Theo dõi độ thành thạo của các hành động
3. **Universal Skills** - Hệ thống kỹ năng universal cho combat/magic/profession
4. **Multi-System Support** - Actor có thể tu luyện nhiều hệ thống cùng lúc

## Kiến Trúc Mới

```
Actor Core v2.0
├── PrimaryCore
│   ├── Basic Stats (10) - Vitality, Endurance, Constitution, etc.
│   ├── Physical Stats (3) - Strength, Agility, Personality
│   ├── Universal Cultivation Stats (5) - SpiritualEnergy, PhysicalEnergy, etc.
│   ├── FlexibleStats - Custom stats shared across systems
│   ├── ProficiencySystem - Skill mastery tracking
│   └── UniversalSkillSystem - Universal skills
├── Derived
│   ├── Core Derived Stats (50+) - HPMax, Speed, Accuracy, etc.
│   ├── Talent Amplifiers (6) - CultivationSpeed, EnergyEfficiency, etc.
│   └── FlexibleStats - Custom derived stats shared across systems
└── Multi-System Support
    ├── ActiveSystems - Track which systems are active
    ├── SystemLevels - Level in each system
    ├── SystemProgress - Progress in each system
    └── CrossSystemSynergies - Bonuses between systems
```

## Tính Năng Mới

### 1. Flexible Stats Maps
```go
type FlexibleStats struct {
    // Custom Primary Stats (shared across systems)
    CustomPrimary map[string]int64  // e.g., "FireMastery", "SwordProficiency"
    
    // Custom Derived Stats (shared across systems)  
    CustomDerived map[string]float64  // e.g., "FireDamage", "SwordSpeed"
    
    // Sub-System Specific Stats (independent per system)
    SubSystemStats map[string]map[string]float64  // systemName -> statName -> value
}
```

**Lợi ích:**
- **Shared Stats**: Fire Mastery có thể được chia sẻ giữa Magic và Alchemy
- **System Independence**: Mỗi hệ thống có stats riêng
- **Easy Expansion**: Dễ dàng thêm stats mới

### 2. Proficiency System
```go
type Proficiency struct {
    SkillName     string    // Name of the skill/action
    Category      string    // Category: "Combat", "Magic", "Crafting", etc.
    Level         int64     // Proficiency level (0-100)
    Experience    int64     // Experience points in this skill
    MaxLevel      int64     // Maximum possible level
    Multiplier    float64   // Experience multiplier based on talent
    LastUsed      int64     // Timestamp of last use
    TotalUses     int64     // Total number of uses
}
```

**Lợi ích:**
- **Skill Mastery**: Theo dõi độ thành thạo của từng hành động
- **Experience Tracking**: Tăng kinh nghiệm qua sử dụng
- **Talent Integration**: Tích hợp với talent system

### 3. Universal Skills
```go
type UniversalSkill struct {
    Name          string    // Skill name
    Category      string    // "Combat", "Magic", "Crafting", "Social", "Movement"
    SubCategory   string    // "Melee", "Ranged", "Fire", "Water", etc.
    Level         int64     // Skill level
    Experience    int64     // Experience points
    MaxLevel      int64     // Maximum level
    Requirements  []string  // Required stats or other skills
    Bonuses       map[string]float64  // Bonuses provided by this skill
    Cooldown      int64     // Cooldown in seconds
    ManaCost      float64   // Mana cost
    StaminaCost   float64   // Stamina cost
}
```

**Categories:**
- **Combat**: MeleeCombat, RangedCombat, Defense, Evasion, etc.
- **Magic**: FireMagic, WaterMagic, EarthMagic, AirMagic, etc.
- **Crafting**: Blacksmithing, Alchemy, Enchanting, Cooking, etc.
- **Social**: Persuasion, Intimidation, Diplomacy, Leadership, etc.
- **Movement**: Acrobatics, Stealth, Climbing, Swimming, etc.
- **Survival**: Tracking, Hunting, Foraging, AnimalHandling, etc.

### 4. Multi-System Support
```go
type MultiSystemActor struct {
    ActorID        string
    PrimaryCore    *PrimaryCore
    Derived        *Derived
    ActiveSystems  map[string]bool  // systemName -> isActive
    SystemLevels   map[string]int64 // systemName -> level
    SystemProgress map[string]int64 // systemName -> progress
}
```

**Ví dụ Multi-System Actor:**
```go
actor := &MultiSystemActor{
    ActorID: "player_001",
    ActiveSystems: map[string]bool{
        "RPG": true,
        "KimDan": true,
        "Magic": true,
    },
    SystemLevels: map[string]int64{
        "RPG": 25,
        "KimDan": 3,  // Golden Core
        "Magic": 15,
    },
    SystemProgress: map[string]int64{
        "RPG": 15000,
        "KimDan": 5000,
        "Magic": 8000,
    },
}
```

## Cross-System Synergies

### 1. Shared Custom Stats
```go
// Fire Mastery được chia sẻ giữa Magic và Alchemy
CustomPrimary["FireMastery"] = 150

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

### 2. Cross-System Bonuses
```go
type CrossSystemSynergy struct {
    Systems     []string  // Các hệ thống liên quan
    BonusType   string    // Loại bonus
    BonusValue  float64   // Giá trị bonus
    Requirements map[string]int64  // Yêu cầu tối thiểu cho mỗi hệ thống
}

// Ví dụ: Dual System Mastery
dualSystemBonus := &CrossSystemSynergy{
    Systems: []string{"RPG", "Magic"},
    BonusType: "SpellDamage",
    BonusValue: 0.2,  // +20% spell damage
    Requirements: map[string]int64{
        "RPG": 20,
        "Magic": 15,
    },
}
```

## Database Schema

### Actor Core Stats Table
```sql
CREATE TABLE actor_core_stats (
    actor_id VARCHAR(255) PRIMARY KEY,
    
    -- Basic Stats
    vitality INT64 NOT NULL DEFAULT 100,
    endurance INT64 NOT NULL DEFAULT 100,
    constitution INT64 NOT NULL DEFAULT 100,
    intelligence INT64 NOT NULL DEFAULT 100,
    wisdom INT64 NOT NULL DEFAULT 100,
    charisma INT64 NOT NULL DEFAULT 100,
    willpower INT64 NOT NULL DEFAULT 100,
    luck INT64 NOT NULL DEFAULT 100,
    fate INT64 NOT NULL DEFAULT 100,
    karma INT64 NOT NULL DEFAULT 100,
    
    -- Physical Stats
    strength INT64 NOT NULL DEFAULT 100,
    agility INT64 NOT NULL DEFAULT 100,
    personality INT64 NOT NULL DEFAULT 100,
    
    -- Universal Cultivation Stats
    spiritual_energy INT64 NOT NULL DEFAULT 1000,
    physical_energy INT64 NOT NULL DEFAULT 1000,
    mental_energy INT64 NOT NULL DEFAULT 1000,
    cultivation_level INT64 NOT NULL DEFAULT 1,
    breakthrough_points INT64 NOT NULL DEFAULT 0,
    
    -- Flexible Stats (JSON)
    custom_primary JSON,
    custom_derived JSON,
    sub_system_stats JSON,
    
    -- Multi-System Support
    active_systems JSON,
    system_levels JSON,
    system_progress JSON,
    
    -- Metadata
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    version INT64 NOT NULL DEFAULT 1
);
```

### Proficiency Table
```sql
CREATE TABLE actor_proficiencies (
    actor_id VARCHAR(255),
    skill_name VARCHAR(255),
    category VARCHAR(100),
    level INT64 NOT NULL DEFAULT 0,
    experience INT64 NOT NULL DEFAULT 0,
    max_level INT64 NOT NULL DEFAULT 100,
    multiplier FLOAT64 NOT NULL DEFAULT 1.0,
    last_used INT64 NOT NULL DEFAULT 0,
    total_uses INT64 NOT NULL DEFAULT 0,
    
    PRIMARY KEY (actor_id, skill_name),
    FOREIGN KEY (actor_id) REFERENCES actor_core_stats(actor_id)
);
```

### Universal Skills Table
```sql
CREATE TABLE actor_skills (
    actor_id VARCHAR(255),
    skill_name VARCHAR(255),
    category VARCHAR(100),
    sub_category VARCHAR(100),
    level INT64 NOT NULL DEFAULT 0,
    experience INT64 NOT NULL DEFAULT 0,
    max_level INT64 NOT NULL DEFAULT 100,
    requirements JSON,
    bonuses JSON,
    cooldown INT64 NOT NULL DEFAULT 0,
    mana_cost FLOAT64 NOT NULL DEFAULT 0.0,
    stamina_cost FLOAT64 NOT NULL DEFAULT 0.0,
    
    PRIMARY KEY (actor_id, skill_name),
    FOREIGN KEY (actor_id) REFERENCES actor_core_stats(actor_id)
);
```

## Benefits

### 1. Flexibility
- **Custom Stats**: Dễ dàng thêm stats mới cho từng hệ thống
- **Shared Stats**: Stats có thể được chia sẻ giữa các hệ thống
- **System Independence**: Mỗi hệ thống hoạt động độc lập

### 2. Scalability
- **Multi-System**: Actor có thể tu luyện nhiều hệ thống
- **Easy Integration**: Dễ dàng tích hợp hệ thống mới
- **Cross-System Synergies**: Bonuses giữa các hệ thống

### 3. Realism
- **Proficiency**: Độ thành thạo tăng qua sử dụng
- **Universal Skills**: Kỹ năng có thể dùng ở nhiều hệ thống
- **Shared Mastery**: Mastery được chia sẻ giữa các hệ thống

### 4. Performance
- **Efficient Storage**: JSON storage cho flexible stats
- **Caching**: Cache cho frequently accessed stats
- **Optimized Queries**: Optimized database queries

## Implementation Priority

### Phase 1: Core Foundation
1. **PrimaryCore** với FlexibleStats
2. **Derived** với FlexibleStats
3. **Basic interfaces** (StatProvider, StatConsumer, StatResolver)

### Phase 2: Proficiency & Skills
1. **ProficiencySystem** implementation
2. **UniversalSkillSystem** implementation
3. **Database schema** cho proficiency và skills

### Phase 3: Multi-System Support
1. **MultiSystemActor** implementation
2. **Cross-System Synergies** implementation
3. **Integration** với existing sub-systems

### Phase 4: Optimization
1. **Performance optimization**
2. **Caching system**
3. **Load testing**

## Next Steps

1. **Review và approve** design này
2. **Tạo Actor Core v2.0** implementation
3. **Implement Proficiency System**
4. **Implement Universal Skills**
5. **Test Multi-System Support**

---

*Tài liệu này tóm tắt thiết kế mới của Actor Core v2.0 với Multi-System Cultivation support.*
