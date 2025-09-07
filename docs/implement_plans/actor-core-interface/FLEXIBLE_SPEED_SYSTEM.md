# Flexible Speed System - Actor Core v2.0

## Tổng Quan

Flexible Speed System được thiết kế để hỗ trợ nhiều loại tốc độ khác nhau trong game, từ movement cơ bản đến casting ma pháp, crafting, learning, và administrative tasks. Hệ thống này tích hợp với talent system để cung cấp bonuses dựa trên tư chất của nhân vật.

## Kiến Trúc

```
Flexible Speed System
├── Speed Categories (7 categories)
│   ├── Movement Speeds (8 types)
│   ├── Casting Speeds (7 types)
│   ├── Crafting Speeds (10 types)
│   ├── Learning Speeds (7 types)
│   ├── Combat Speeds (8 types)
│   ├── Social Speeds (8 types)
│   └── Administrative Speeds (7 types)
├── Talent Bonuses
│   ├── General Speed Talents (7 types)
│   └── Specific Speed Talents (6 types)
└── Speed Calculation Engine
    ├── Base Speed Calculation
    ├── Category Multipliers
    ├── Talent Bonus Application
    └── Final Speed Calculation
```

## Speed Categories

### 1. Movement Speeds (8 types)
```go
const (
    SPEED_WALKING     = "Walking"      // Đi bộ
    SPEED_RUNNING     = "Running"      // Chạy
    SPEED_SWIMMING    = "Swimming"     // Bơi
    SPEED_CLIMBING    = "Climbing"     // Leo trèo
    SPEED_FLYING      = "Flying"       // Bay
    SPEED_TELEPORT    = "Teleport"     // Dịch chuyển tức thời
    SPEED_DASH        = "Dash"         // Xông tới
    SPEED_BLINK       = "Blink"        // Nhấp nháy
)
```

**Formulas:**
- Walking = (Agility × 0.8 + Endurance × 0.2) × (1 + CultivationLevel/100)
- Running = (Agility × 1.0 + Endurance × 0.4) × (1 + CultivationLevel/120)
- Swimming = (Endurance × 0.6 + Agility × 0.4) × (1 + CultivationLevel/140)
- Climbing = (Strength × 0.4 + Agility × 0.6) × (1 + CultivationLevel/130)
- Flying = (SpiritualEnergy × 0.3 + MentalEnergy × 0.2 + CultivationLevel × 0.1) × (1 + CultivationLevel/100)
- Teleport = (SpiritualEnergy × 0.2 + MentalEnergy × 0.1 + CultivationLevel × 0.05) × (1 + CultivationLevel/150)
- Dash = (Agility × 1.2 + Strength × 0.3) × (1 + CultivationLevel/110)
- Blink = (SpiritualEnergy × 0.4 + MentalEnergy × 0.2) × (1 + CultivationLevel/120)

### 2. Casting Speeds (7 types)
```go
const (
    SPEED_SPELL_CASTING     = "SpellCasting"      // Thi triển ma pháp
    SPEED_MAGIC_FORMATION   = "MagicFormation"    // Bày trận pháp
    SPEED_IMMORTAL_TECHNIQUE = "ImmortalTechnique" // Tiên thuật
    SPEED_QI_CIRCULATION    = "QiCirculation"     // Vận khí
    SPEED_MEDITATION        = "Meditation"        // Thiền định
    SPEED_BREAKTHROUGH      = "Breakthrough"      // Đột phá
    SPEED_CULTIVATION       = "Cultivation"       // Tu luyện
)
```

**Formulas:**
- SpellCasting = (Intelligence × 0.8 + Wisdom × 0.6 + SpiritualEnergy × 0.1) × (1 + CultivationLevel/100)
- MagicFormation = (Intelligence × 0.6 + Wisdom × 0.8 + SpiritualEnergy × 0.2) × (1 + CultivationLevel/120)
- ImmortalTechnique = (Wisdom × 1.0 + SpiritualEnergy × 0.3 + CultivationLevel × 0.2) × (1 + CultivationLevel/100)
- QiCirculation = (Willpower × 0.8 + SpiritualEnergy × 0.4 + CultivationLevel × 0.1) × (1 + CultivationLevel/110)
- Meditation = (Wisdom × 0.6 + Willpower × 0.4 + SpiritualEnergy × 0.2) × (1 + CultivationLevel/130)
- Breakthrough = (Willpower × 1.0 + SpiritualEnergy × 0.5 + CultivationLevel × 0.3) × (1 + CultivationLevel/100)
- Cultivation = (Wisdom × 0.8 + Willpower × 0.6 + SpiritualEnergy × 0.3) × (1 + CultivationLevel/120)

### 3. Crafting Speeds (10 types)
```go
const (
    SPEED_ALCHEMY          = "Alchemy"           // Luyện đan
    SPEED_REFINING         = "Refining"          // Luyện khí
    SPEED_FORGING          = "Forging"           // Rèn
    SPEED_ENCHANTING       = "Enchanting"        // Phù chú
    SPEED_ARRAY_FORMATION  = "ArrayFormation"    // Bày trận
    SPEED_FORMATION_SETUP  = "FormationSetup"    // Bày cấm chế
    SPEED_PILL_REFINING    = "PillRefining"      // Luyện đan
    SPEED_WEAPON_CRAFTING  = "WeaponCrafting"    // Chế tạo vũ khí
    SPEED_ARMOR_CRAFTING   = "ArmorCrafting"     // Chế tạo giáp
    SPEED_JEWELRY_CRAFTING = "JewelryCrafting"   // Chế tạo trang sức
)
```

**Formulas:**
- Alchemy = (Intelligence × 0.8 + Wisdom × 0.6 + SpiritualEnergy × 0.1) × (1 + CultivationLevel/100)
- Refining = (Intelligence × 0.6 + Wisdom × 0.8 + SpiritualEnergy × 0.2) × (1 + CultivationLevel/120)
- Forging = (Strength × 0.6 + Intelligence × 0.4 + PhysicalEnergy × 0.2) × (1 + CultivationLevel/110)
- Enchanting = (Intelligence × 0.8 + Wisdom × 0.6 + SpiritualEnergy × 0.3) × (1 + CultivationLevel/130)
- ArrayFormation = (Intelligence × 0.7 + Wisdom × 0.8 + SpiritualEnergy × 0.4) × (1 + CultivationLevel/140)
- FormationSetup = (Intelligence × 0.6 + Wisdom × 0.7 + SpiritualEnergy × 0.3) × (1 + CultivationLevel/130)
- PillRefining = (Intelligence × 0.8 + Wisdom × 0.7 + SpiritualEnergy × 0.2) × (1 + CultivationLevel/120)
- WeaponCrafting = (Strength × 0.7 + Intelligence × 0.5 + PhysicalEnergy × 0.2) × (1 + CultivationLevel/110)
- ArmorCrafting = (Strength × 0.6 + Intelligence × 0.6 + PhysicalEnergy × 0.3) × (1 + CultivationLevel/120)
- JewelryCrafting = (Intelligence × 0.8 + Wisdom × 0.6 + SpiritualEnergy × 0.3) × (1 + CultivationLevel/130)

### 4. Learning Speeds (7 types)
```go
const (
    SPEED_READING          = "Reading"           // Đọc sách
    SPEED_STUDYING         = "Studying"          // Học tập
    SPEED_COMPREHENSION    = "Comprehension"     // Hiểu biết
    SPEED_SKILL_LEARNING   = "SkillLearning"     // Học kỹ năng
    SPEED_TECHNIQUE_MASTERY = "TechniqueMastery" // Thành thạo kỹ thuật
    SPEED_MEMORIZATION     = "Memorization"      // Ghi nhớ
    SPEED_RESEARCH         = "Research"          // Nghiên cứu
)
```

**Formulas:**
- Reading = (Intelligence × 0.8 + Wisdom × 0.6 + MentalEnergy × 0.1) × (1 + CultivationLevel/100)
- Studying = (Intelligence × 0.7 + Wisdom × 0.8 + MentalEnergy × 0.2) × (1 + CultivationLevel/120)
- Comprehension = (Intelligence × 0.6 + Wisdom × 0.9 + MentalEnergy × 0.3) × (1 + CultivationLevel/130)
- SkillLearning = (Intelligence × 0.8 + Wisdom × 0.7 + MentalEnergy × 0.2) × (1 + CultivationLevel/110)
- TechniqueMastery = (Intelligence × 0.7 + Wisdom × 0.8 + MentalEnergy × 0.3) × (1 + CultivationLevel/140)
- Memorization = (Intelligence × 0.9 + Wisdom × 0.5 + MentalEnergy × 0.2) × (1 + CultivationLevel/100)
- Research = (Intelligence × 0.8 + Wisdom × 0.8 + MentalEnergy × 0.3) × (1 + CultivationLevel/130)

### 5. Combat Speeds (8 types)
```go
const (
    SPEED_ATTACK           = "Attack"            // Tấn công
    SPEED_DEFENSE          = "Defense"           // Phòng thủ
    SPEED_DODGE            = "Dodge"             // Né tránh
    SPEED_BLOCK            = "Block"             // Chặn
    SPEED_PARRY            = "Parry"             // Đỡ
    SPEED_COUNTER_ATTACK   = "CounterAttack"     // Phản công
    SPEED_WEAPON_SWITCH    = "WeaponSwitch"      // Đổi vũ khí
    SPEED_STANCE_CHANGE    = "StanceChange"      // Đổi tư thế
)
```

**Formulas:**
- Attack = (Agility × 0.8 + Strength × 0.4 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/100)
- Defense = (Agility × 0.6 + Constitution × 0.6 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/110)
- Dodge = (Agility × 1.0 + Luck × 0.3 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/120)
- Block = (Strength × 0.6 + Constitution × 0.8 + PhysicalEnergy × 0.2) × (1 + CultivationLevel/110)
- Parry = (Agility × 0.8 + Intelligence × 0.4 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/120)
- CounterAttack = (Agility × 0.7 + Strength × 0.5 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/110)
- WeaponSwitch = (Agility × 0.6 + Intelligence × 0.4 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/100)
- StanceChange = (Agility × 0.5 + Intelligence × 0.6 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/100)

### 6. Social Speeds (8 types)
```go
const (
    SPEED_CONVERSATION     = "Conversation"      // Trò chuyện
    SPEED_NEGOTIATION      = "Negotiation"       // Đàm phán
    SPEED_PERSUASION       = "Persuasion"        // Thuyết phục
    SPEED_INTIMIDATION     = "Intimidation"      // Uy hiếp
    SPEED_DIPLOMACY        = "Diplomacy"         // Ngoại giao
    SPEED_LEADERSHIP       = "Leadership"        // Lãnh đạo
    SPEED_TRADING          = "Trading"           // Buôn bán
    SPEED_NETWORKING       = "Networking"        // Kết nối
)
```

**Formulas:**
- Conversation = (Charisma × 0.8 + Intelligence × 0.4 + Personality × 0.3) × (1 + CultivationLevel/100)
- Negotiation = (Charisma × 0.6 + Intelligence × 0.8 + Personality × 0.4) × (1 + CultivationLevel/120)
- Persuasion = (Charisma × 0.8 + Intelligence × 0.5 + Personality × 0.5) × (1 + CultivationLevel/110)
- Intimidation = (Strength × 0.4 + Charisma × 0.6 + Personality × 0.4) × (1 + CultivationLevel/100)
- Diplomacy = (Charisma × 0.7 + Intelligence × 0.7 + Personality × 0.6) × (1 + CultivationLevel/130)
- Leadership = (Charisma × 0.8 + Intelligence × 0.6 + Personality × 0.7) × (1 + CultivationLevel/120)
- Trading = (Intelligence × 0.6 + Charisma × 0.5 + Personality × 0.6) × (1 + CultivationLevel/100)
- Networking = (Charisma × 0.7 + Intelligence × 0.5 + Personality × 0.8) × (1 + CultivationLevel/110)

### 7. Administrative Speeds (7 types)
```go
const (
    SPEED_PLANNING         = "Planning"          // Lập kế hoạch
    SPEED_ORGANIZATION     = "Organization"      // Tổ chức
    SPEED_DECISION_MAKING  = "DecisionMaking"    // Ra quyết định
    SPEED_RESOURCE_MANAGEMENT = "ResourceManagement" // Quản lý tài nguyên
    SPEED_STRATEGY_FORMULATION = "StrategyFormulation" // Lập chiến lược
    SPEED_TACTICAL_ANALYSIS = "TacticalAnalysis" // Phân tích chiến thuật
    SPEED_RISK_ASSESSMENT  = "RiskAssessment"    // Đánh giá rủi ro
)
```

**Formulas:**
- Planning = (Intelligence × 0.8 + Wisdom × 0.7 + MentalEnergy × 0.2) × (1 + CultivationLevel/120)
- Organization = (Intelligence × 0.7 + Wisdom × 0.8 + MentalEnergy × 0.2) × (1 + CultivationLevel/130)
- DecisionMaking = (Intelligence × 0.6 + Wisdom × 0.9 + MentalEnergy × 0.3) × (1 + CultivationLevel/140)
- ResourceManagement = (Intelligence × 0.8 + Wisdom × 0.6 + MentalEnergy × 0.2) × (1 + CultivationLevel/110)
- StrategyFormulation = (Intelligence × 0.7 + Wisdom × 0.8 + MentalEnergy × 0.3) × (1 + CultivationLevel/130)
- TacticalAnalysis = (Intelligence × 0.8 + Wisdom × 0.7 + MentalEnergy × 0.2) × (1 + CultivationLevel/120)
- RiskAssessment = (Intelligence × 0.6 + Wisdom × 0.8 + MentalEnergy × 0.3) × (1 + CultivationLevel/130)

## Talent Bonuses

### 1. General Speed Talents
```go
type SpeedTalentBonuses struct {
    SpeedMastery     float64  // General speed bonus (+10-50%)
    QuickLearner     float64  // Learning speed bonus (+15-60%)
    FastCaster       float64  // Casting speed bonus (+20-70%)
    SwiftCrafter     float64  // Crafting speed bonus (+15-55%)
    AgileFighter     float64  // Combat speed bonus (+25-80%)
    SmoothTalker     float64  // Social speed bonus (+10-40%)
    EfficientPlanner float64  // Administrative speed bonus (+12-45%)
}
```

### 2. Specific Speed Talents
```go
type SpeedTalentBonuses struct {
    MagicAffinity    float64  // Magic casting speed (+30-100%)
    CultivationGift  float64  // Cultivation speed (+25-90%)
    CraftingGenius   float64  // Crafting speed (+35-120%)
    CombatInstinct   float64  // Combat speed (+40-150%)
    SocialCharm      float64  // Social speed (+20-70%)
    StrategicMind    float64  // Planning speed (+25-85%)
}
```

## Speed Calculation Engine

### 1. Base Speed Calculation
```go
func (fss *FlexibleSpeedSystem) CalculateSpeed(category string, speedType string, baseSpeed float64, talentBonuses map[string]float64) float64 {
    // Get base speed
    speed := baseSpeed
    
    // Apply category multiplier
    if categorySpeed, exists := fss.getCategorySpeed(category, speedType); exists {
        speed *= categorySpeed
    }
    
    // Apply talent bonuses
    for talent, bonus := range talentBonuses {
        if fss.isTalentApplicable(talent, speedType) {
            speed *= (1.0 + bonus)
        }
    }
    
    return speed
}
```

### 2. Talent Applicability
```go
func (fss *FlexibleSpeedSystem) isTalentApplicable(talent string, speedType string) bool {
    switch talent {
    case "SpeedMastery":
        return true  // Applicable to all speed types
    case "QuickLearner":
        return fss.isLearningSpeed(speedType)
    case "FastCaster":
        return fss.isCastingSpeed(speedType)
    case "SwiftCrafter":
        return fss.isCraftingSpeed(speedType)
    case "AgileFighter":
        return fss.isCombatSpeed(speedType)
    case "SmoothTalker":
        return fss.isSocialSpeed(speedType)
    case "EfficientPlanner":
        return fss.isAdministrativeSpeed(speedType)
    case "MagicAffinity":
        return fss.isMagicCastingSpeed(speedType)
    case "CultivationGift":
        return fss.isCultivationSpeed(speedType)
    case "CraftingGenius":
        return fss.isCraftingSpeed(speedType)
    case "CombatInstinct":
        return fss.isCombatSpeed(speedType)
    case "SocialCharm":
        return fss.isSocialSpeed(speedType)
    case "StrategicMind":
        return fss.isAdministrativeSpeed(speedType)
    default:
        return false
    }
}
```

## Database Schema

### Speed System Table
```sql
CREATE TABLE actor_speed_system (
    actor_id VARCHAR(255) PRIMARY KEY,
    
    -- Movement Speeds
    movement_speeds JSON,
    
    -- Casting Speeds
    casting_speeds JSON,
    
    -- Crafting Speeds
    crafting_speeds JSON,
    
    -- Learning Speeds
    learning_speeds JSON,
    
    -- Combat Speeds
    combat_speeds JSON,
    
    -- Social Speeds
    social_speeds JSON,
    
    -- Administrative Speeds
    administrative_speeds JSON,
    
    -- Talent Bonuses
    talent_bonuses JSON,
    
    -- Metadata
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    version INT64 NOT NULL DEFAULT 1,
    
    FOREIGN KEY (actor_id) REFERENCES actor_core_stats(actor_id)
);
```

## Usage Examples

### 1. Basic Speed Usage
```go
// Get movement speed
speedSystem := actor.PrimaryCore.SpeedSystem
walkingSpeed := speedSystem.CalculateSpeed("MovementSpeeds", "Walking", baseSpeed, talentBonuses)

// Get casting speed
spellCastingSpeed := speedSystem.CalculateSpeed("CastingSpeeds", "SpellCasting", baseSpeed, talentBonuses)

// Get crafting speed
alchemySpeed := speedSystem.CalculateSpeed("CraftingSpeeds", "Alchemy", baseSpeed, talentBonuses)
```

### 2. Talent Bonus Application
```go
// Create talent bonuses
talentBonuses := &SpeedTalentBonuses{
    SpeedMastery: 0.2,      // +20% general speed
    FastCaster: 0.3,        // +30% casting speed
    MagicAffinity: 0.5,     // +50% magic casting speed
    SwiftCrafter: 0.25,     // +25% crafting speed
}

// Apply to specific speed
spellCastingSpeed := speedSystem.CalculateSpeed("CastingSpeeds", "SpellCasting", baseSpeed, talentBonuses)
// Result: baseSpeed * (1 + 0.2) * (1 + 0.3) * (1 + 0.5) = baseSpeed * 2.34
```

### 3. Speed Breakdown
```go
// Get detailed speed breakdown
breakdown := speedSystem.GetSpeedBreakdown(actorID, "CastingSpeeds", "SpellCasting")
fmt.Printf("Base Speed: %.2f\n", breakdown.BaseSpeed)
fmt.Printf("Category Multiplier: %.2f\n", breakdown.CategoryMultiplier)
fmt.Printf("Talent Bonuses: %v\n", breakdown.TalentBonuses)
fmt.Printf("Final Speed: %.2f\n", breakdown.FinalSpeed)
```

## Benefits

### 1. Flexibility
- **Multiple Speed Types**: Hỗ trợ nhiều loại tốc độ khác nhau
- **Category Organization**: Tổ chức tốc độ theo danh mục
- **Easy Extension**: Dễ dàng thêm loại tốc độ mới

### 2. Realism
- **Different Requirements**: Mỗi loại tốc độ có yêu cầu khác nhau
- **Talent Integration**: Tích hợp với hệ thống tư chất
- **Cultivation Level**: Tốc độ tăng theo cấp độ tu luyện

### 3. Performance
- **Efficient Calculation**: Tính toán hiệu quả
- **Caching Support**: Hỗ trợ caching
- **Breakdown Transparency**: Minh bạch trong tính toán

### 4. Scalability
- **Modular Design**: Thiết kế modular
- **Easy Integration**: Dễ dàng tích hợp với các hệ thống khác
- **Database Optimization**: Tối ưu hóa database

## Implementation Priority

### Phase 1: Core Speed System
1. **FlexibleSpeedSystem** struct
2. **Speed calculation engine**
3. **Basic speed categories**

### Phase 2: Talent Integration
1. **SpeedTalentBonuses** struct
2. **Talent applicability logic**
3. **Bonus calculation**

### Phase 3: Advanced Features
1. **Speed breakdown system**
2. **Caching optimization**
3. **Database integration**

### Phase 4: Testing & Optimization
1. **Unit tests**
2. **Performance tests**
3. **Integration tests**

---

*Tài liệu này mô tả Flexible Speed System cho Actor Core v2.0, hỗ trợ nhiều loại tốc độ khác nhau với tích hợp talent system.*
