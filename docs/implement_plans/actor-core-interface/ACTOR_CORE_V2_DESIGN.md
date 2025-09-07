# Actor Core v2.0 Design Specification

## Tổng Quan

Actor Core v2.0 là hệ thống cốt lõi mới được thiết kế để hỗ trợ Universal Cultivation Stats và tích hợp với các sub-systems như RPG System, Kim Đan System, v.v. Hệ thống này tập trung vào việc định nghĩa các chỉ số cơ bản và universal, không bao gồm level progression.

## Kiến Trúc Tổng Thể

```
Actor Core v2.0
├── Core Interfaces
│   ├── StatProvider (cho sub-systems)
│   ├── StatConsumer (cho game systems)
│   └── StatResolver (cho calculations)
├── PrimaryCore (Universal Stats)
│   ├── Basic Stats (10)
│   ├── Physical Stats (3)
│   └── Universal Cultivation Stats (5)
├── Derived (Calculated Stats)
│   ├── Core Derived Stats (20+)
│   ├── Talent Amplifiers (6)
│   └── Sub-System Derived Stats (RPG, Kim Đan, etc.)
└── Integration Layer
    ├── Stat Mapping
    ├── Conversion Utilities
    └── Sub-System Adapters
```

## PrimaryCore Stats

### 1. Basic Stats (10 stats)
```go
type BasicStats struct {
    // Core attributes
    Vitality     int64  // Sức sống, ảnh hưởng HP, regen
    Endurance    int64  // Sức chịu đựng, ảnh hưởng stamina, HP
    Constitution int64  // Thể chất, ảnh hưởng HP, resistances
    Intelligence int64  // Trí tuệ, ảnh hưởng mana, learning
    Wisdom       int64  // Trí tuệ, ảnh hưởng mana efficiency, decision
    Charisma     int64  // Sức hút, ảnh hưởng social interactions
    Willpower    int64  // Ý chí, ảnh hưởng mental resistance, mana
    Luck         int64  // May mắn, ảnh hưởng critical, random events
    Fate         int64  // Số phận, ảnh hưởng destiny, karma
    Karma        int64  // Nghiệp lực, ảnh hưởng moral alignment
}
```

### 2. Physical Stats (3 stats)
```go
type PhysicalStats struct {
    Strength   int64  // Sức mạnh vật lý, ảnh hưởng damage, carry weight
    Agility    int64  // Nhanh nhẹn, ảnh hưởng speed, evasion
    Personality int64  // Tính cách, ảnh hưởng social, merchant prices
}
```

### 3. Universal Cultivation Stats (5 stats)
```go
type UniversalCultivationStats struct {
    SpiritualEnergy  int64  // Năng lượng tinh thần, cho tu luyện
    PhysicalEnergy   int64  // Năng lượng vật lý, cho luyện thể
    MentalEnergy     int64  // Năng lượng tinh thần, cho trí tuệ
    CultivationLevel int64  // Cấp độ tu luyện tổng thể
    BreakthroughPoints int64  // Điểm đột phá, cho breakthrough
}
```

### 4. Complete PrimaryCore Structure
```go
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
    
    // Flexible Stats (for custom stats shared across systems)
    FlexibleStats *FlexibleStats
    
    // Flexible Speed System (for different types of speed)
    SpeedSystem *FlexibleSpeedSystem
    
    // Flexible Administrative Division System (for different classifications)
    AdministrativeSystem *FlexibleAdministrativeSystem
    
    // Flexible Karma System (for different worlds/regions/sects)
    KarmaSystem *FlexibleKarmaSystem
    
    // Proficiency System
    Proficiency *ProficiencySystem
    
    // Universal Skills
    Skills *UniversalSkillSystem
}
```

## Derived Stats

### 1. Complete Derived Structure
```go
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
    
    // Flexible Stats (for custom derived stats shared across systems)
    FlexibleStats *FlexibleStats
    
    // Flexible Speed System (for different types of speed)
    SpeedSystem *FlexibleSpeedSystem
}

### 2. Core Derived Stats (Detailed)
```go
type CoreDerivedStats struct {
    // Health & Resources
    HPMax           float64  // Max health points
    Stamina         float64  // Max stamina
    Speed           float64  // Movement speed
    
    // Combat Stats
    Accuracy        float64  // Hit chance
    Penetration     float64  // Armor penetration
    Lethality       float64  // Critical damage
    Brutality       float64  // Damage multiplier
    
    // Defense Stats
    ArmorClass      float64  // Armor rating
    Evasion         float64  // Dodge chance
    BlockChance     float64  // Block chance
    ParryChance     float64  // Parry chance
    DodgeChance     float64  // Dodge chance
    
    // Energy Stats
    EnergyEfficiency float64  // Energy usage efficiency
    EnergyCapacity   float64  // Max energy capacity
    EnergyDrain      float64  // Energy drain rate
    ResourceRegen    float64  // Resource regeneration
    ResourceDecay    float64  // Resource decay rate
    
    // Resistance Stats
    StatusResistance float64  // Status effect resistance
    Immunity         float64  // Immunity chance
    WeatherResist    float64  // Weather resistance
    TerrainMastery   float64  // Terrain mastery
    ClimateAdapt     float64  // Climate adaptation
    
    // Learning Stats
    LearningRate     float64  // Learning speed
    Adaptation       float64  // Adaptation speed
    Memory           float64  // Memory capacity
    Experience       float64  // Experience gain
    
    // Social Stats
    Leadership       float64  // Leadership ability
    Diplomacy        float64  // Diplomacy skill
    Intimidation     float64  // Intimidation skill
    Empathy          float64  // Empathy skill
    Deception        float64  // Deception skill
    Performance      float64  // Performance skill
    
    // Mystical Stats
    ManaEfficiency   float64  // Mana usage efficiency
    SpellPower       float64  // Spell power
    MysticResonance  float64  // Mystical resonance
    RealityBend      float64  // Reality bending
    TimeSense        float64  // Time perception
    SpaceSense       float64  // Space perception
    
    // Movement Stats (Basic)
    JumpHeight       float64  // Jump height
    ClimbSpeed       float64  // Climbing speed
    SwimSpeed        float64  // Swimming speed
    FlightSpeed      float64  // Flight speed
    TeleportRange    float64  // Teleportation range
    Stealth          float64  // Stealth ability
    
    // Aura Stats
    AuraRadius       float64  // Aura range
    AuraStrength     float64  // Aura power
    Presence         float64  // Presence
    Awe              float64  // Awe factor
    
    // Proficiency Stats
    WeaponMastery    float64  // Weapon proficiency
    SkillLevel       float64  // General skill level
    LifeSteal        float64  // Life steal
    CastSpeed        float64  // Casting speed
    WeightCapacity   float64  // Carry weight
    Persuasion       float64  // Persuasion skill
    MerchantPriceModifier float64  // Merchant price modifier
    FactionReputationGain float64  // Faction reputation gain
}
```

### 2. Talent Amplifiers
```go
type TalentAmplifiers struct {
    CultivationSpeed    float64  // Tu luyện tốc độ
    EnergyEfficiency    float64  // Hiệu suất năng lượng
    BreakthroughSuccess float64  // Thành công đột phá
    SkillLearning       float64  // Học kỹ năng
    CombatEffectiveness float64  // Hiệu quả chiến đấu
    ResourceGathering   float64  // Thu thập tài nguyên
}
```

### 3. Flexible Stat Maps
```go
// Flexible maps for custom stats that can be shared across systems
type FlexibleStats struct {
    // Custom Primary Stats (can be shared across systems)
    CustomPrimary map[string]int64  // e.g., "FireMastery", "SwordProficiency", "AlchemyLevel"
    
    // Custom Derived Stats (can be shared across systems)  
    CustomDerived map[string]float64  // e.g., "FireDamage", "SwordSpeed", "PotionQuality"
    
    // Sub-System Specific Stats (independent per system)
    SubSystemStats map[string]map[string]float64  // systemName -> statName -> value
}

// Example usage:
// SubSystemStats["RPG"]["LifeSteal"] = 0.05
// SubSystemStats["KimDan"]["QiCapacity"] = 1000.0
// SubSystemStats["Magic"]["FireMastery"] = 150.0
// CustomPrimary["FireMastery"] = 150  // Shared between Magic and Alchemy systems
```

### 4. Flexible Speed System
```go
// Flexible speed system for different types of speed
type FlexibleSpeedSystem struct {
    // Speed Categories
    MovementSpeeds    map[string]float64  // Basic movement speeds
    CastingSpeeds     map[string]float64  // Magic/cultivation casting speeds
    CraftingSpeeds    map[string]float64  // Crafting/production speeds
    LearningSpeeds    map[string]float64  // Learning/studying speeds
    CombatSpeeds      map[string]float64  // Combat action speeds
    SocialSpeeds      map[string]float64  // Social interaction speeds
    AdministrativeSpeeds map[string]float64  // Administrative action speeds
}
```

### 5. Flexible Administrative Division System
```go
// Flexible administrative division system for different classifications
type FlexibleAdministrativeSystem struct {
    // Administrative Divisions (flexible hierarchy)
    Divisions map[string]map[string]AdministrativeDivision  // divisionType -> divisionName -> Division
    
    // Division Types
    DivisionTypes map[string]DivisionType  // divisionType -> DivisionType definition
    
    // Division Hierarchy
    Hierarchy map[string][]string  // parentDivision -> list of childDivisions
    
    // Cross-Division Relationships
    Relationships map[string]map[string]DivisionRelationship  // divisionA -> divisionB -> Relationship
}

// Administrative Division Definition
type AdministrativeDivision struct {
    ID          string   // Unique identifier
    Name        string   // Display name
    Type        string   // Division type: "World", "Continent", "Realm", "Nation", "Region", "District", "Race", "Sect", etc.
    Level       int64    // Hierarchy level (0 = highest)
    Parent      string   // Parent division ID
    Children    []string // Child division IDs
    Attributes  map[string]interface{}  // Flexible attributes
    KarmaTypes  []string // Available karma types
    Multiplier  float64  // Karma multiplier
    CreatedAt   int64    // Creation timestamp
    UpdatedAt   int64    // Last update timestamp
}

// Division Type Definition
type DivisionType struct {
    Name        string   // Type name
    Description string   // Description
    Level       int64    // Default hierarchy level
    Attributes  []string // Required attributes
    KarmaTypes  []string // Default karma types
    Multiplier  float64  // Default multiplier
    IsActive    bool     // Is this type active
}

// Division Relationship
type DivisionRelationship struct {
    Type        string  // Relationship type: "Contains", "BelongsTo", "Allied", "Enemy", "Neutral", "Trading", "Cultural", "Religious", etc.
    Strength    float64 // Relationship strength (0.0 - 1.0)
    Description string  // Relationship description
    CreatedAt   int64   // Creation timestamp
    UpdatedAt   int64   // Last update timestamp
}

// Division Types Constants
const (
    DIVISION_WORLD      = "World"      // Thế giới
    DIVISION_CONTINENT  = "Continent"  // Lục địa
    DIVISION_REALM      = "Realm"      // Vương quốc/Realm
    DIVISION_NATION     = "Nation"     // Quốc gia
    DIVISION_REGION     = "Region"     // Khu vực
    DIVISION_DISTRICT   = "District"   // Quận/Huyện
    DIVISION_CITY       = "City"       // Thành phố
    DIVISION_VILLAGE    = "Village"    // Làng
    DIVISION_RACE       = "Race"       // Chủng tộc
    DIVISION_TRIBE      = "Tribe"      // Bộ tộc
    DIVISION_CLAN       = "Clan"       // Gia tộc
    DIVISION_SECT       = "Sect"       // Tông môn
    DIVISION_ORDER      = "Order"      // Hội/Order
    DIVISION_GUILD      = "Guild"      // Hội nghề nghiệp
    DIVISION_EMPIRE     = "Empire"     // Đế quốc
    DIVISION_KINGDOM    = "Kingdom"    // Vương quốc
    DIVISION_REPUBLIC   = "Republic"   // Cộng hòa
    DIVISION_FEDERATION = "Federation" // Liên bang
    DIVISION_ALLIANCE   = "Alliance"   // Liên minh
    DIVISION_CONFEDERACY = "Confederacy" // Liên minh lỏng lẻo
    DIVISION_PROVINCE   = "Province"   // Tỉnh
    DIVISION_STATE      = "State"      // Bang
    DIVISION_COUNTY     = "County"     // Hạt
    DIVISION_MUNICIPALITY = "Municipality" // Thành phố tự trị
    DIVISION_TERRITORY  = "Territory"  // Lãnh thổ
    DIVISION_DOMAIN     = "Domain"     // Lãnh địa
    DIVISION_FIEF       = "Fief"       // Phong địa
    DIVISION_MARCH      = "March"      // Biên giới
    DIVISION_FRONTIER   = "Frontier"   // Vùng biên
    DIVISION_WILDLAND   = "Wildland"   // Vùng hoang dã
    DIVISION_WASTELAND  = "Wasteland"  // Vùng hoang tàn
    DIVISION_SANCTUARY  = "Sanctuary"  // Thánh địa
    DIVISION_SHRINE     = "Shrine"     // Đền thờ
    DIVISION_TEMPLE     = "Temple"     // Chùa/Đền
    DIVISION_MONASTERY  = "Monastery"  // Tu viện
    DIVISION_ACADEMY    = "Academy"    // Học viện
    DIVISION_UNIVERSITY = "University" // Đại học
    DIVISION_LIBRARY    = "Library"    // Thư viện
    DIVISION_LABORATORY = "Laboratory" // Phòng thí nghiệm
    DIVISION_WORKSHOP   = "Workshop"   // Xưởng
    DIVISION_FACTORY    = "Factory"    // Nhà máy
    DIVISION_MARKET     = "Market"     // Chợ
    DIVISION_PORT       = "Port"       // Cảng
    DIVISION_HARBOR     = "Harbor"     // Bến cảng
    DIVISION_AIRPORT    = "Airport"    // Sân bay
    DIVISION_STATION    = "Station"    // Ga
    DIVISION_DEPOT      = "Depot"      // Kho
    DIVISION_WAREHOUSE  = "Warehouse"  // Nhà kho
    DIVISION_VAULT      = "Vault"      // Kho báu
    DIVISION_TREASURY   = "Treasury"   // Kho bạc
    DIVISION_ARMORY     = "Armory"     // Kho vũ khí
    DIVISION_BARRACKS   = "Barracks"   // Doanh trại
    DIVISION_FORTRESS   = "Fortress"   // Pháo đài
    DIVISION_CASTLE     = "Castle"     // Lâu đài
    DIVISION_PALACE     = "Palace"     // Cung điện
    DIVISION_MANSION    = "Mansion"    // Biệt thự
    DIVISION_ESTATE     = "Estate"     // Điền trang
    DIVISION_PLANTATION = "Plantation" // Đồn điền
    DIVISION_FARM       = "Farm"       // Trang trại
    DIVISION_RANCH      = "Ranch"      // Trại chăn nuôi
    DIVISION_MINE       = "Mine"       // Mỏ
    DIVISION_QUARRY     = "Quarry"     // Mỏ đá
    DIVISION_FOREST     = "Forest"     // Rừng
    DIVISION_JUNGLE     = "Jungle"     // Rừng rậm
    DIVISION_DESERT     = "Desert"     // Sa mạc
    DIVISION_TUNDRA     = "Tundra"     // Đài nguyên
    DIVISION_GLACIER    = "Glacier"    // Sông băng
    DIVISION_VOLCANO    = "Volcano"    // Núi lửa
    DIVISION_MOUNTAIN   = "Mountain"   // Núi
    DIVISION_HILL       = "Hill"       // Đồi
    DIVISION_VALLEY     = "Valley"     // Thung lũng
    DIVISION_PLAIN      = "Plain"      // Đồng bằng
    DIVISION_PLATEAU    = "Plateau"    // Cao nguyên
    DIVISION_CANYON     = "Canyon"     // Hẻm núi
    DIVISION_CAVE       = "Cave"       // Hang động
    DIVISION_UNDERGROUND = "Underground" // Dưới lòng đất
    DIVISION_SKY        = "Sky"        // Bầu trời
    DIVISION_CLOUD      = "Cloud"      // Mây
    DIVISION_ASTRAL     = "Astral"     // Cõi tinh thần
    DIVISION_ETHEREAL   = "Ethereal"   // Cõi ether
    DIVISION_SHADOW     = "Shadow"     // Cõi bóng tối
    DIVISION_LIGHT      = "Light"      // Cõi ánh sáng
    DIVISION_VOID       = "Void"       // Cõi hư vô
    DIVISION_CHAOS      = "Chaos"      // Cõi hỗn loạn
    DIVISION_ORDER      = "Order"      // Cõi trật tự
    DIVISION_TIME       = "Time"       // Cõi thời gian
    DIVISION_SPACE      = "Space"      // Cõi không gian
    DIVISION_DIMENSION  = "Dimension"  // Chiều không gian
    DIVISION_PLANE      = "Plane"      // Mặt phẳng
    DIVISION_REALM      = "Realm"      // Cõi giới
    DIVISION_DOMAIN     = "Domain"     // Lãnh địa
    DIVISION_TERRITORY  = "Territory"  // Lãnh thổ
    DIVISION_ZONE       = "Zone"       // Vùng
    DIVISION_AREA       = "Area"       // Khu vực
    DIVISION_SECTOR     = "Sector"     // Khu vực
    DIVISION_QUADRANT   = "Quadrant"   // Góc phần tư
    DIVISION_HEMISPHERE = "Hemisphere" // Bán cầu
    DIVISION_SPHERE     = "Sphere"     // Cầu
    DIVISION_CIRCLE     = "Circle"     // Vòng tròn
    DIVISION_RING       = "Ring"       // Vòng
    DIVISION_BELT       = "Belt"       // Vành đai
)
```

### 6. Flexible Karma System
```go
// Flexible karma system for different worlds/regions/sects
type FlexibleKarmaSystem struct {
    // Global Karma (tổng của tất cả thế giới)
    GlobalKarma map[string]int64  // karmaType -> totalValue
    
    // Division-specific Karma (theo từng đơn vị hành chính)
    DivisionKarma map[string]map[string]int64  // divisionType -> divisionName -> karmaType -> value
    
    // Karma Categories and Types
    KarmaCategories map[string][]string  // category -> list of karma types
    KarmaTypes      map[string]KarmaType  // karmaType -> KarmaType definition
}

// Karma Type Definition
type KarmaType struct {
    Name        string  // Tên loại karma
    Category    string  // Danh mục: "Fortune", "Karma", "Merit", "Contribution", etc.
    Description string  // Mô tả
    MinValue    int64   // Giá trị tối thiểu
    MaxValue    int64   // Giá trị tối đa
    DefaultValue int64  // Giá trị mặc định
    IsPositive  bool    // Có phải là karma tích cực không
    Weight      float64 // Trọng số trong tính toán tổng
}

// Karma Categories Constants
const (
    KARMA_FORTUNE      = "Fortune"      // Khí vận
    KARMA_KARMA        = "Karma"        // Nghiệp lực
    KARMA_MERIT        = "Merit"        // Công đức
    KARMA_CONTRIBUTION = "Contribution" // Cống hiến
    KARMA_REPUTATION   = "Reputation"   // Danh tiếng
    KARMA_HONOR        = "Honor"        // Danh dự
    KARMA_GLORY        = "Glory"        // Vinh quang
    KARMA_WISDOM       = "Wisdom"       // Trí tuệ
    KARMA_COMPASSION   = "Compassion"   // Từ bi
    KARMA_JUSTICE      = "Justice"      // Công lý
    KARMA_VALOR        = "Valor"        // Dũng khí
    KARMA_LOYALTY      = "Loyalty"      // Trung thành
    KARMA_FAITH        = "Faith"        // Đức tin
    KARMA_HOPE         = "Hope"         // Hy vọng
    KARMA_LOVE         = "Love"         // Tình yêu
    KARMA_PEACE        = "Peace"        // Hòa bình
    KARMA_HARMONY      = "Harmony"      // Hòa hợp
    KARMA_BALANCE      = "Balance"      // Cân bằng
    KARMA_TRUTH        = "Truth"        // Chân lý
    KARMA_FREEDOM      = "Freedom"      // Tự do
)

// World/Region/Sect Definitions
type WorldDefinition struct {
    Name        string   // Tên thế giới
    Type        string   // Loại: "Mortal", "Immortal", "Divine", "Demon", etc.
    Level       int64    // Cấp độ thế giới
    KarmaTypes  []string // Các loại karma có trong thế giới này
    Multiplier  float64  // Hệ số nhân cho karma
}

type RegionDefinition struct {
    Name        string   // Tên khu vực
    World       string   // Thuộc thế giới nào
    Type        string   // Loại: "City", "Wilderness", "Dungeon", "Temple", etc.
    Level       int64    // Cấp độ khu vực
    KarmaTypes  []string // Các loại karma có trong khu vực này
    Multiplier  float64  // Hệ số nhân cho karma
}

type SectDefinition struct {
    Name        string   // Tên tông môn
    World       string   // Thuộc thế giới nào
    Type        string   // Loại: "Righteous", "Demonic", "Neutral", "Heretic", etc.
    Level       int64    // Cấp độ tông môn
    KarmaTypes  []string // Các loại karma có trong tông môn này
    Multiplier  float64  // Hệ số nhân cho karma
}

type NationDefinition struct {
    Name        string   // Tên quốc gia
    World       string   // Thuộc thế giới nào
    Type        string   // Loại: "Empire", "Kingdom", "Republic", "Tribe", etc.
    Level       int64    // Cấp độ quốc gia
    KarmaTypes  []string // Các loại karma có trong quốc gia này
    Multiplier  float64  // Hệ số nhân cho karma
}

// Karma Calculation Engine
type KarmaCalculationEngine struct {
    karmaSystem *FlexibleKarmaSystem
    worldDefs   map[string]*WorldDefinition
    regionDefs  map[string]*RegionDefinition
    sectDefs    map[string]*SectDefinition
    nationDefs  map[string]*NationDefinition
}

// Calculate total karma for a specific type
func (kce *KarmaCalculationEngine) CalculateTotalKarma(karmaType string) int64 {
    total := int64(0)
    
    // Add global karma
    if globalValue, exists := kce.karmaSystem.GlobalKarma[karmaType]; exists {
        total += globalValue
    }
    
    // Add world-specific karma
    for _, worldKarma := range kce.karmaSystem.WorldKarma {
        if worldValue, exists := worldKarma[karmaType]; exists {
            total += worldValue
        }
    }
    
    // Add region-specific karma
    for _, regionKarma := range kce.karmaSystem.RegionKarma {
        if regionValue, exists := regionKarma[karmaType]; exists {
            total += regionValue
        }
    }
    
    // Add sect-specific karma
    for _, sectKarma := range kce.karmaSystem.SectKarma {
        if sectValue, exists := sectKarma[karmaType]; exists {
            total += sectValue
        }
    }
    
    // Add nation-specific karma
    for _, nationKarma := range kce.karmaSystem.NationKarma {
        if nationValue, exists := nationKarma[karmaType]; exists {
            total += nationValue
        }
    }
    
    return total
}

// Calculate karma for a specific world
func (kce *KarmaCalculationEngine) CalculateWorldKarma(worldName string, karmaType string) int64 {
    if worldKarma, exists := kce.karmaSystem.WorldKarma[worldName]; exists {
        if value, exists := worldKarma[karmaType]; exists {
            return value
        }
    }
    return 0
}

// Calculate karma for a specific region
func (kce *KarmaCalculationEngine) CalculateRegionKarma(regionName string, karmaType string) int64 {
    if regionKarma, exists := kce.karmaSystem.RegionKarma[regionName]; exists {
        if value, exists := regionKarma[karmaType]; exists {
            return value
        }
    }
    return 0
}

// Calculate karma for a specific sect
func (kce *KarmaCalculationEngine) CalculateSectKarma(sectName string, karmaType string) int64 {
    if sectKarma, exists := kce.karmaSystem.SectKarma[sectName]; exists {
        if value, exists := sectKarma[karmaType]; exists {
            return value
        }
    }
    return 0
}

// Calculate karma for a specific nation
func (kce *KarmaCalculationEngine) CalculateNationKarma(nationName string, karmaType string) int64 {
    if nationKarma, exists := kce.karmaSystem.NationKarma[nationName]; exists {
        if value, exists := nationKarma[karmaType]; exists {
            return value
        }
    }
    return 0
}

// Calculate weighted karma score
func (kce *KarmaCalculationEngine) CalculateWeightedKarmaScore(karmaType string) float64 {
    totalKarma := kce.CalculateTotalKarma(karmaType)
    
    // Get karma type definition
    if karmaDef, exists := kce.karmaSystem.KarmaTypes[karmaType]; exists {
        // Apply weight
        return float64(totalKarma) * karmaDef.Weight
    }
    
    return float64(totalKarma)
}

// Calculate karma influence on stats
func (kce *KarmaCalculationEngine) CalculateKarmaInfluence(karmaType string, baseStat float64) float64 {
    weightedKarma := kce.CalculateWeightedKarmaScore(karmaType)
    
    // Karma influence formula: baseStat * (1 + karmaInfluence)
    karmaInfluence := weightedKarma / 10000.0  // Scale karma to reasonable influence
    
    return baseStat * (1.0 + karmaInfluence)
}

// Speed Categories Constants
const (
    // Movement Speeds
    SPEED_WALKING     = "Walking"
    SPEED_RUNNING     = "Running"
    SPEED_SWIMMING    = "Swimming"
    SPEED_CLIMBING    = "Climbing"
    SPEED_FLYING      = "Flying"
    SPEED_TELEPORT    = "Teleport"
    SPEED_DASH        = "Dash"
    SPEED_BLINK       = "Blink"
    
    // Casting Speeds (Magic/Cultivation)
    SPEED_SPELL_CASTING     = "SpellCasting"
    SPEED_MAGIC_FORMATION   = "MagicFormation"
    SPEED_IMMORTAL_TECHNIQUE = "ImmortalTechnique"
    SPEED_QI_CIRCULATION    = "QiCirculation"
    SPEED_MEDITATION        = "Meditation"
    SPEED_BREAKTHROUGH      = "Breakthrough"
    SPEED_CULTIVATION       = "Cultivation"
    
    // Crafting Speeds
    SPEED_ALCHEMY          = "Alchemy"
    SPEED_REFINING         = "Refining"
    SPEED_FORGING          = "Forging"
    SPEED_ENCHANTING       = "Enchanting"
    SPEED_ARRAY_FORMATION  = "ArrayFormation"
    SPEED_FORMATION_SETUP  = "FormationSetup"
    SPEED_PILL_REFINING    = "PillRefining"
    SPEED_WEAPON_CRAFTING  = "WeaponCrafting"
    SPEED_ARMOR_CRAFTING   = "ArmorCrafting"
    SPEED_JEWELRY_CRAFTING = "JewelryCrafting"
    
    // Learning Speeds
    SPEED_READING          = "Reading"
    SPEED_STUDYING         = "Studying"
    SPEED_COMPREHENSION    = "Comprehension"
    SPEED_SKILL_LEARNING   = "SkillLearning"
    SPEED_TECHNIQUE_MASTERY = "TechniqueMastery"
    SPEED_MEMORIZATION     = "Memorization"
    SPEED_RESEARCH         = "Research"
    
    // Combat Speeds
    SPEED_ATTACK           = "Attack"
    SPEED_DEFENSE          = "Defense"
    SPEED_DODGE            = "Dodge"
    SPEED_BLOCK            = "Block"
    SPEED_PARRY            = "Parry"
    SPEED_COUNTER_ATTACK   = "CounterAttack"
    SPEED_WEAPON_SWITCH    = "WeaponSwitch"
    SPEED_STANCE_CHANGE    = "StanceChange"
    
    // Social Speeds
    SPEED_CONVERSATION     = "Conversation"
    SPEED_NEGOTIATION      = "Negotiation"
    SPEED_PERSUASION       = "Persuasion"
    SPEED_INTIMIDATION     = "Intimidation"
    SPEED_DIPLOMACY        = "Diplomacy"
    SPEED_LEADERSHIP       = "Leadership"
    SPEED_TRADING          = "Trading"
    SPEED_NETWORKING       = "Networking"
    
    // Administrative Speeds
    SPEED_PLANNING         = "Planning"
    SPEED_ORGANIZATION     = "Organization"
    SPEED_DECISION_MAKING  = "DecisionMaking"
    SPEED_RESOURCE_MANAGEMENT = "ResourceManagement"
    SPEED_STRATEGY_FORMULATION = "StrategyFormulation"
    SPEED_TACTICAL_ANALYSIS = "TacticalAnalysis"
    SPEED_RISK_ASSESSMENT  = "RiskAssessment"
)

// Speed calculation with talent bonuses
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

// Talent bonuses for different speed types
type SpeedTalentBonuses struct {
    // General Speed Talents
    SpeedMastery     float64  // General speed bonus
    QuickLearner     float64  // Learning speed bonus
    FastCaster       float64  // Casting speed bonus
    SwiftCrafter     float64  // Crafting speed bonus
    AgileFighter     float64  // Combat speed bonus
    SmoothTalker     float64  // Social speed bonus
    EfficientPlanner float64  // Administrative speed bonus
    
    // Specific Speed Talents
    MagicAffinity    float64  // Magic casting speed
    CultivationGift  float64  // Cultivation speed
    CraftingGenius   float64  // Crafting speed
    CombatInstinct   float64  // Combat speed
    SocialCharm      float64  // Social speed
    StrategicMind    float64  // Planning speed
}
```

### 4. Proficiency System
```go
// Proficiency tracks mastery of specific actions/skills
type Proficiency struct {
    SkillName     string    // Name of the skill/action
    Category      string    // Category: "Combat", "Magic", "Crafting", "Social", etc.
    Level         int64     // Proficiency level (0-100)
    Experience    int64     // Experience points in this skill
    MaxLevel      int64     // Maximum possible level
    Multiplier    float64   // Experience multiplier based on talent
    LastUsed      int64     // Timestamp of last use
    TotalUses     int64     // Total number of uses
}

type ProficiencySystem struct {
    Proficiencies map[string]*Proficiency  // skillName -> Proficiency
    Categories    map[string][]string      // category -> list of skills
    MaxSkills     int                      // Maximum number of skills
}

// Proficiency calculation
func (ps *ProficiencySystem) CalculateProficiency(skillName string) float64 {
    prof, exists := ps.Proficiencies[skillName]
    if !exists {
        return 0.0
    }
    
    // Base proficiency from level
    baseProficiency := float64(prof.Level) / 100.0
    
    // Bonus from experience
    experienceBonus := float64(prof.Experience) / 10000.0
    
    // Talent multiplier
    talentMultiplier := prof.Multiplier
    
    return (baseProficiency + experienceBonus) * talentMultiplier
}
```

### 5. Universal Skill System
```go
// Universal skills that can be used across multiple systems
type UniversalSkill struct {
    Name          string    // Skill name
    Category      string    // "Combat", "Magic", "Crafting", "Social", "Movement"
    SubCategory   string    // "Melee", "Ranged", "Fire", "Water", "Blacksmithing", etc.
    Level         int64     // Skill level
    Experience    int64     // Experience points
    MaxLevel      int64     // Maximum level
    Requirements  []string  // Required stats or other skills
    Bonuses       map[string]float64  // Bonuses provided by this skill
    Cooldown      int64     // Cooldown in seconds
    ManaCost      float64   // Mana cost
    StaminaCost   float64   // Stamina cost
}

type UniversalSkillSystem struct {
    Skills        map[string]*UniversalSkill  // skillName -> Skill
    Categories    map[string][]string         // category -> list of skills
    SkillTrees    map[string]*SkillTree       // skillTreeName -> SkillTree
    MaxSkills     int                         // Maximum number of skills
}

// Skill categories
const (
    SKILL_COMBAT    = "Combat"
    SKILL_MAGIC     = "Magic"
    SKILL_CRAFTING  = "Crafting"
    SKILL_SOCIAL    = "Social"
    SKILL_MOVEMENT  = "Movement"
    SKILL_SURVIVAL  = "Survival"
)

// Combat skills
const (
    SKILL_MELEE_COMBAT    = "MeleeCombat"
    SKILL_RANGED_COMBAT   = "RangedCombat"
    SKILL_UNARMED_COMBAT  = "UnarmedCombat"
    SKILL_DEFENSE         = "Defense"
    SKILL_EVASION         = "Evasion"
    SKILL_BLOCKING        = "Blocking"
    SKILL_PARRYING        = "Parrying"
)

// Magic skills
const (
    SKILL_FIRE_MAGIC      = "FireMagic"
    SKILL_WATER_MAGIC     = "WaterMagic"
    SKILL_EARTH_MAGIC     = "EarthMagic"
    SKILL_AIR_MAGIC       = "AirMagic"
    SKILL_LIGHT_MAGIC     = "LightMagic"
    SKILL_DARK_MAGIC      = "DarkMagic"
    SKILL_HEALING_MAGIC   = "HealingMagic"
    SKILL_SUMMONING       = "Summoning"
    SKILL_ENCHANTING      = "Enchanting"
)

// Crafting skills
const (
    SKILL_BLACKSMITHING   = "Blacksmithing"
    SKILL_ALCHEMY         = "Alchemy"
    SKILL_ENCHANTING      = "Enchanting"
    SKILL_COOKING         = "Cooking"
    SKILL_TAILORING       = "Tailoring"
    SKILL_CARPENTRY       = "Carpentry"
    SKILL_JEWELCRAFTING   = "Jewelcrafting"
)

// Social skills
const (
    SKILL_PERSUASION      = "Persuasion"
    SKILL_INTIMIDATION    = "Intimidation"
    SKILL_DIPLOMACY       = "Diplomacy"
    SKILL_LEADERSHIP      = "Leadership"
    SKILL_TRADING         = "Trading"
    SKILL_NEGOTIATION     = "Negotiation"
)

// Movement skills
const (
    SKILL_ACROBATICS      = "Acrobatics"
    SKILL_STEALTH         = "Stealth"
    SKILL_CLIMBING        = "Climbing"
    SKILL_SWIMMING        = "Swimming"
    SKILL_RIDING          = "Riding"
    SKILL_FLYING          = "Flying"
)

// Survival skills
const (
    SKILL_TRACKING        = "Tracking"
    SKILL_HUNTING         = "Hunting"
    SKILL_FORAGING        = "Foraging"
    SKILL_SURVIVAL        = "Survival"
    SKILL_ANIMAL_HANDLING = "AnimalHandling"
    SKILL_NATURE_LORE     = "NatureLore"
)
```

## Multi-System Cultivation Support

### 1. Multi-System Actor Support
```go
// Actor có thể tu luyện nhiều hệ thống cùng lúc
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

### 2. Shared Custom Stats
```go
// Ví dụ: Fire Mastery được chia sẻ giữa Magic và Alchemy
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

## Interfaces

### 1. StatProvider Interface
```go
type StatProvider interface {
    // Get primary stats
    GetPrimaryStats(actorID string) (*PrimaryCore, error)
    
    // Get derived stats
    GetDerivedStats(actorID string) (*Derived, error)
    
    // Get specific stat
    GetStat(actorID string, statKey string) (float64, error)
    
    // Get custom stat (from flexible maps)
    GetCustomStat(actorID string, statKey string) (interface{}, error)
    
    // Get sub-system stat
    GetSubSystemStat(actorID string, systemName string, statKey string) (float64, error)
    
    // Get stat breakdown
    GetStatBreakdown(actorID string, statKey string) (*StatBreakdown, error)
    
    // Update stats
    UpdateStats(actorID string, updates *StatUpdates) error
    
    // Update custom stats
    UpdateCustomStats(actorID string, updates *CustomStatUpdates) error
    
    // Update sub-system stats
    UpdateSubSystemStats(actorID string, systemName string, updates *SubSystemStatUpdates) error
}
```

### 2. StatConsumer Interface
```go
type StatConsumer interface {
    // Consume stats for calculations
    ConsumeStats(actorID string, stats *ActorCoreStats) error
    
    // Get stat requirements
    GetStatRequirements() []string
    
    // Get custom stat requirements
    GetCustomStatRequirements() []string
    
    // Get sub-system stat requirements
    GetSubSystemStatRequirements(systemName string) []string
    
    // Validate stat compatibility
    ValidateStats(stats *ActorCoreStats) error
    
    // Validate multi-system compatibility
    ValidateMultiSystemStats(actor *MultiSystemActor) error
}
```

### 3. StatResolver Interface
```go
type StatResolver interface {
    // Resolve all stats
    ResolveStats(actorID string) (*ActorCoreStats, error)
    
    // Resolve specific stat
    ResolveStat(actorID string, statKey string) (float64, error)
    
    // Resolve custom stat
    ResolveCustomStat(actorID string, statKey string) (interface{}, error)
    
    // Resolve sub-system stat
    ResolveSubSystemStat(actorID string, systemName string, statKey string) (float64, error)
    
    // Get stat dependencies
    GetStatDependencies(statKey string) []string
    
    // Get custom stat dependencies
    GetCustomStatDependencies(statKey string) []string
    
    // Calculate derived stats
    CalculateDerivedStats(primary *PrimaryCore) (*Derived, error)
    
    // Calculate custom derived stats
    CalculateCustomDerivedStats(primary *PrimaryCore, customPrimary map[string]int64) (map[string]float64, error)
    
    // Calculate cross-system synergies
    CalculateCrossSystemSynergies(actor *MultiSystemActor) (map[string]float64, error)
}
```

### 4. Proficiency Interface
```go
type ProficiencyInterface interface {
    // Get proficiency level
    GetProficiency(actorID string, skillName string) (*Proficiency, error)
    
    // Update proficiency
    UpdateProficiency(actorID string, skillName string, experience int64) error
    
    // Calculate proficiency bonus
    CalculateProficiencyBonus(actorID string, skillName string) (float64, error)
    
    // Get all proficiencies
    GetAllProficiencies(actorID string) (map[string]*Proficiency, error)
    
    // Get proficiencies by category
    GetProficienciesByCategory(actorID string, category string) (map[string]*Proficiency, error)
}
```

### 5. Skill Interface
```go
type SkillInterface interface {
    // Get skill level
    GetSkill(actorID string, skillName string) (*UniversalSkill, error)
    
    // Update skill
    UpdateSkill(actorID string, skillName string, experience int64) error
    
    // Calculate skill bonus
    CalculateSkillBonus(actorID string, skillName string) (float64, error)
    
    // Get all skills
    GetAllSkills(actorID string) (map[string]*UniversalSkill, error)
    
    // Get skills by category
    GetSkillsByCategory(actorID string, category string) (map[string]*UniversalSkill, error)
    
    // Check skill requirements
    CheckSkillRequirements(actorID string, skillName string) (bool, error)
    
    // Learn new skill
    LearnSkill(actorID string, skillName string) error
}
```

### 6. Speed Interface
```go
type SpeedInterface interface {
    // Get speed for specific category and type
    GetSpeed(actorID string, category string, speedType string) (float64, error)
    
    // Get all speeds for a category
    GetSpeedsByCategory(actorID string, category string) (map[string]float64, error)
    
    // Get all speeds
    GetAllSpeeds(actorID string) (*FlexibleSpeedSystem, error)
    
    // Calculate speed with talent bonuses
    CalculateSpeed(actorID string, category string, speedType string, baseSpeed float64) (float64, error)
    
    // Update speed
    UpdateSpeed(actorID string, category string, speedType string, value float64) error
    
    // Get speed talent bonuses
    GetSpeedTalentBonuses(actorID string) (*SpeedTalentBonuses, error)
    
    // Update speed talent bonuses
    UpdateSpeedTalentBonuses(actorID string, bonuses *SpeedTalentBonuses) error
    
    // Get speed breakdown
    GetSpeedBreakdown(actorID string, category string, speedType string) (*SpeedBreakdown, error)
}

// Speed breakdown for transparency
type SpeedBreakdown struct {
    BaseSpeed      float64  // Base speed from stats
    CategoryMultiplier float64  // Category multiplier
    TalentBonuses  map[string]float64  // Talent bonuses applied
    FinalSpeed     float64  // Final calculated speed
    Notes          []string  // Additional notes
}
```

### 7. Administrative Division Interface
```go
type AdministrativeDivisionInterface interface {
    // Get division by ID
    GetDivision(actorID string, divisionID string) (*AdministrativeDivision, error)
    
    // Get divisions by type
    GetDivisionsByType(actorID string, divisionType string) (map[string]*AdministrativeDivision, error)
    
    // Get divisions by level
    GetDivisionsByLevel(actorID string, level int64) (map[string]*AdministrativeDivision, error)
    
    // Get parent division
    GetParentDivision(actorID string, divisionID string) (*AdministrativeDivision, error)
    
    // Get child divisions
    GetChildDivisions(actorID string, divisionID string) (map[string]*AdministrativeDivision, error)
    
    // Get division hierarchy
    GetDivisionHierarchy(actorID string, divisionID string) ([]*AdministrativeDivision, error)
    
    // Get division relationships
    GetDivisionRelationships(actorID string, divisionID string) (map[string]*DivisionRelationship, error)
    
    // Create division
    CreateDivision(actorID string, division *AdministrativeDivision) error
    
    // Update division
    UpdateDivision(actorID string, divisionID string, division *AdministrativeDivision) error
    
    // Delete division
    DeleteDivision(actorID string, divisionID string) error
    
    // Add division relationship
    AddDivisionRelationship(actorID string, divisionA string, divisionB string, relationship *DivisionRelationship) error
    
    // Remove division relationship
    RemoveDivisionRelationship(actorID string, divisionA string, divisionB string) error
    
    // Get division types
    GetDivisionTypes() (map[string]DivisionType, error)
    
    // Get division attributes
    GetDivisionAttributes(actorID string, divisionID string) (map[string]interface{}, error)
    
    // Update division attributes
    UpdateDivisionAttributes(actorID string, divisionID string, attributes map[string]interface{}) error
    
    // Get division karma types
    GetDivisionKarmaTypes(actorID string, divisionID string) ([]string, error)
    
    // Update division karma types
    UpdateDivisionKarmaTypes(actorID string, divisionID string, karmaTypes []string) error
    
    // Get division multiplier
    GetDivisionMultiplier(actorID string, divisionID string) (float64, error)
    
    // Update division multiplier
    UpdateDivisionMultiplier(actorID string, divisionID string, multiplier float64) error
}

// Division breakdown for transparency
type DivisionBreakdown struct {
    Division      *AdministrativeDivision           // Division details
    Parent        *AdministrativeDivision           // Parent division
    Children      map[string]*AdministrativeDivision // Child divisions
    Hierarchy     []*AdministrativeDivision         // Full hierarchy
    Relationships map[string]*DivisionRelationship  // Relationships
    Attributes    map[string]interface{}            // Division attributes
    KarmaTypes    []string                          // Available karma types
    Multiplier    float64                           // Karma multiplier
    Notes         []string                          // Additional notes
}
```

### 8. Karma Interface
```go
type KarmaInterface interface {
    // Get total karma for a specific type
    GetTotalKarma(actorID string, karmaType string) (int64, error)
    
    // Get karma for a specific division
    GetDivisionKarma(actorID string, divisionType string, divisionName string, karmaType string) (int64, error)
    
    // Get all karma for a specific division
    GetDivisionKarmaAll(actorID string, divisionType string, divisionName string) (map[string]int64, error)
    
    // Update karma
    UpdateKarma(actorID string, karmaType string, value int64) error
    
    // Update division karma
    UpdateDivisionKarma(actorID string, divisionType string, divisionName string, karmaType string, value int64) error
    
    // Calculate karma influence on stats
    CalculateKarmaInfluence(actorID string, karmaType string, baseStat float64) (float64, error)
    
    // Get karma breakdown
    GetKarmaBreakdown(actorID string, karmaType string) (*KarmaBreakdown, error)
    
    // Get karma categories
    GetKarmaCategories() (map[string][]string, error)
    
    // Get karma types
    GetKarmaTypes() (map[string]KarmaType, error)
}

// Karma breakdown for transparency
type KarmaBreakdown struct {
    TotalKarma    int64                    // Total karma across all divisions
    DivisionKarma map[string]map[string]int64  // Karma by division type and name
    WeightedScore float64                  // Weighted karma score
    Influence     float64                  // Influence on stats
    Notes         []string                 // Additional notes
}
```

## Stat Formulas

### 1. Core Derived Stat Formulas
```go
// Health & Resources
HPMax = (Vitality × 10 + Endurance × 5 + Constitution × 3) × (1 + CultivationLevel/100)
Stamina = (Endurance × 8 + Vitality × 2) × (1 + PhysicalEnergy/1000)
Speed = (Agility × 2 + Intelligence × 1 + Luck × 0.5) × (1 + CultivationLevel/200)

// Combat Stats
Accuracy = 0.5 + (Intelligence × 0.01 + Agility × 0.008 + Luck × 0.005)
Penetration = (Strength × 0.02 + Agility × 0.01) × (1 + CultivationLevel/100)
Lethality = 1.0 + (Luck × 0.01 + Strength × 0.005) × (1 + CultivationLevel/200)
Brutality = 1.0 + (Strength × 0.01 + Willpower × 0.005) × (1 + CultivationLevel/150)

// Defense Stats
ArmorClass = (Constitution × 1.5 + Endurance × 1.0) × (1 + PhysicalEnergy/1000)
Evasion = (Agility × 0.8 + Luck × 0.2) × (1 + CultivationLevel/300)
BlockChance = (Constitution × 0.01 + Strength × 0.005) × (1 + PhysicalEnergy/2000)
ParryChance = (Agility × 0.01 + Intelligence × 0.005) × (1 + MentalEnergy/2000)
DodgeChance = (Agility × 0.01 + Luck × 0.008) × (1 + CultivationLevel/400)

// Energy Stats
EnergyEfficiency = 1.0 + (Wisdom × 0.01 + Willpower × 0.008 + CultivationLevel × 0.005)
EnergyCapacity = (SpiritualEnergy + PhysicalEnergy + MentalEnergy) × 0.1
EnergyDrain = (CultivationLevel × 0.01 + Willpower × 0.005) × (1 + Age/1000)
ResourceRegen = (Endurance × 0.01 + Vitality × 0.008) × (1 + CultivationLevel/100)
ResourceDecay = (CultivationLevel × 0.005 + Age × 0.001) × (1 + Karma/1000)

// Learning Stats
LearningRate = 1.0 + (Intelligence × 0.01 + Wisdom × 0.008 + CultivationLevel × 0.005)
Adaptation = 1.0 + (Intelligence × 0.008 + Luck × 0.005 + CultivationLevel × 0.003)
Memory = (Intelligence × 10 + Wisdom × 5 + MentalEnergy × 0.1) × (1 + CultivationLevel/200)
Experience = 1.0 + (Intelligence × 0.005 + Luck × 0.003 + CultivationLevel × 0.002)

// Social Stats
Leadership = (Charisma × 0.8 + Willpower × 0.5 + Intelligence × 0.3) × (1 + CultivationLevel/100)
Diplomacy = (Charisma × 0.6 + Intelligence × 0.4 + Wisdom × 0.3) × (1 + CultivationLevel/150)
Intimidation = (Strength × 0.5 + Willpower × 0.4 + Charisma × 0.3) × (1 + CultivationLevel/120)
Empathy = (Charisma × 0.4 + Intelligence × 0.3 + Wisdom × 0.5) × (1 + CultivationLevel/200)
Deception = (Intelligence × 0.5 + Charisma × 0.4 + Luck × 0.3) × (1 + CultivationLevel/180)
Performance = (Charisma × 0.6 + Agility × 0.3 + Personality × 0.4) × (1 + CultivationLevel/160)

// Mystical Stats
ManaEfficiency = 1.0 + (Wisdom × 0.01 + Willpower × 0.008 + SpiritualEnergy × 0.001)
SpellPower = (Intelligence × 0.8 + Wisdom × 0.6 + SpiritualEnergy × 0.1) × (1 + CultivationLevel/100)
MysticResonance = (Wisdom × 0.5 + SpiritualEnergy × 0.2 + MentalEnergy × 0.1) × (1 + CultivationLevel/150)
RealityBend = (Wisdom × 0.3 + SpiritualEnergy × 0.1 + MentalEnergy × 0.05) × (1 + CultivationLevel/200)
TimeSense = (Intelligence × 0.4 + Wisdom × 0.3 + MentalEnergy × 0.1) × (1 + CultivationLevel/250)
SpaceSense = (Intelligence × 0.3 + Wisdom × 0.4 + MentalEnergy × 0.1) × (1 + CultivationLevel/250)

// Movement Stats (Basic)
JumpHeight = (Strength × 0.5 + Agility × 0.3 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/100)
ClimbSpeed = (Strength × 0.3 + Agility × 0.5 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/120)
SwimSpeed = (Endurance × 0.4 + Agility × 0.3 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/140)
FlightSpeed = (SpiritualEnergy × 0.2 + MentalEnergy × 0.1 + CultivationLevel × 0.05) × (1 + CultivationLevel/100)
TeleportRange = (SpiritualEnergy × 0.1 + MentalEnergy × 0.05 + CultivationLevel × 0.02) × (1 + CultivationLevel/150)
Stealth = (Agility × 0.6 + Intelligence × 0.3 + Luck × 0.2) × (1 + CultivationLevel/200)

// Flexible Speed System Formulas
// Movement Speeds
Walking = (Agility × 0.8 + Endurance × 0.2) × (1 + CultivationLevel/100)
Running = (Agility × 1.0 + Endurance × 0.4) × (1 + CultivationLevel/120)
Swimming = (Endurance × 0.6 + Agility × 0.4) × (1 + CultivationLevel/140)
Climbing = (Strength × 0.4 + Agility × 0.6) × (1 + CultivationLevel/130)
Flying = (SpiritualEnergy × 0.3 + MentalEnergy × 0.2 + CultivationLevel × 0.1) × (1 + CultivationLevel/100)
Teleport = (SpiritualEnergy × 0.2 + MentalEnergy × 0.1 + CultivationLevel × 0.05) × (1 + CultivationLevel/150)
Dash = (Agility × 1.2 + Strength × 0.3) × (1 + CultivationLevel/110)
Blink = (SpiritualEnergy × 0.4 + MentalEnergy × 0.2) × (1 + CultivationLevel/120)

// Casting Speeds (Magic/Cultivation)
SpellCasting = (Intelligence × 0.8 + Wisdom × 0.6 + SpiritualEnergy × 0.1) × (1 + CultivationLevel/100)
MagicFormation = (Intelligence × 0.6 + Wisdom × 0.8 + SpiritualEnergy × 0.2) × (1 + CultivationLevel/120)
ImmortalTechnique = (Wisdom × 1.0 + SpiritualEnergy × 0.3 + CultivationLevel × 0.2) × (1 + CultivationLevel/100)
QiCirculation = (Willpower × 0.8 + SpiritualEnergy × 0.4 + CultivationLevel × 0.1) × (1 + CultivationLevel/110)
Meditation = (Wisdom × 0.6 + Willpower × 0.4 + SpiritualEnergy × 0.2) × (1 + CultivationLevel/130)
Breakthrough = (Willpower × 1.0 + SpiritualEnergy × 0.5 + CultivationLevel × 0.3) × (1 + CultivationLevel/100)
Cultivation = (Wisdom × 0.8 + Willpower × 0.6 + SpiritualEnergy × 0.3) × (1 + CultivationLevel/120)

// Crafting Speeds
Alchemy = (Intelligence × 0.8 + Wisdom × 0.6 + SpiritualEnergy × 0.1) × (1 + CultivationLevel/100)
Refining = (Intelligence × 0.6 + Wisdom × 0.8 + SpiritualEnergy × 0.2) × (1 + CultivationLevel/120)
Forging = (Strength × 0.6 + Intelligence × 0.4 + PhysicalEnergy × 0.2) × (1 + CultivationLevel/110)
Enchanting = (Intelligence × 0.8 + Wisdom × 0.6 + SpiritualEnergy × 0.3) × (1 + CultivationLevel/130)
ArrayFormation = (Intelligence × 0.7 + Wisdom × 0.8 + SpiritualEnergy × 0.4) × (1 + CultivationLevel/140)
FormationSetup = (Intelligence × 0.6 + Wisdom × 0.7 + SpiritualEnergy × 0.3) × (1 + CultivationLevel/130)
PillRefining = (Intelligence × 0.8 + Wisdom × 0.7 + SpiritualEnergy × 0.2) × (1 + CultivationLevel/120)
WeaponCrafting = (Strength × 0.7 + Intelligence × 0.5 + PhysicalEnergy × 0.2) × (1 + CultivationLevel/110)
ArmorCrafting = (Strength × 0.6 + Intelligence × 0.6 + PhysicalEnergy × 0.3) × (1 + CultivationLevel/120)
JewelryCrafting = (Intelligence × 0.8 + Wisdom × 0.6 + SpiritualEnergy × 0.3) × (1 + CultivationLevel/130)

// Learning Speeds
Reading = (Intelligence × 0.8 + Wisdom × 0.6 + MentalEnergy × 0.1) × (1 + CultivationLevel/100)
Studying = (Intelligence × 0.7 + Wisdom × 0.8 + MentalEnergy × 0.2) × (1 + CultivationLevel/120)
Comprehension = (Intelligence × 0.6 + Wisdom × 0.9 + MentalEnergy × 0.3) × (1 + CultivationLevel/130)
SkillLearning = (Intelligence × 0.8 + Wisdom × 0.7 + MentalEnergy × 0.2) × (1 + CultivationLevel/110)
TechniqueMastery = (Intelligence × 0.7 + Wisdom × 0.8 + MentalEnergy × 0.3) × (1 + CultivationLevel/140)
Memorization = (Intelligence × 0.9 + Wisdom × 0.5 + MentalEnergy × 0.2) × (1 + CultivationLevel/100)
Research = (Intelligence × 0.8 + Wisdom × 0.8 + MentalEnergy × 0.3) × (1 + CultivationLevel/130)

// Combat Speeds
Attack = (Agility × 0.8 + Strength × 0.4 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/100)
Defense = (Agility × 0.6 + Constitution × 0.6 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/110)
Dodge = (Agility × 1.0 + Luck × 0.3 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/120)
Block = (Strength × 0.6 + Constitution × 0.8 + PhysicalEnergy × 0.2) × (1 + CultivationLevel/110)
Parry = (Agility × 0.8 + Intelligence × 0.4 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/120)
CounterAttack = (Agility × 0.7 + Strength × 0.5 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/110)
WeaponSwitch = (Agility × 0.6 + Intelligence × 0.4 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/100)
StanceChange = (Agility × 0.5 + Intelligence × 0.6 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/100)

// Social Speeds
Conversation = (Charisma × 0.8 + Intelligence × 0.4 + Personality × 0.3) × (1 + CultivationLevel/100)
Negotiation = (Charisma × 0.6 + Intelligence × 0.8 + Personality × 0.4) × (1 + CultivationLevel/120)
Persuasion = (Charisma × 0.8 + Intelligence × 0.5 + Personality × 0.5) × (1 + CultivationLevel/110)
Intimidation = (Strength × 0.4 + Charisma × 0.6 + Personality × 0.4) × (1 + CultivationLevel/100)
Diplomacy = (Charisma × 0.7 + Intelligence × 0.7 + Personality × 0.6) × (1 + CultivationLevel/130)
Leadership = (Charisma × 0.8 + Intelligence × 0.6 + Personality × 0.7) × (1 + CultivationLevel/120)
Trading = (Intelligence × 0.6 + Charisma × 0.5 + Personality × 0.6) × (1 + CultivationLevel/100)
Networking = (Charisma × 0.7 + Intelligence × 0.5 + Personality × 0.8) × (1 + CultivationLevel/110)

// Administrative Speeds
Planning = (Intelligence × 0.8 + Wisdom × 0.7 + MentalEnergy × 0.2) × (1 + CultivationLevel/120)
Organization = (Intelligence × 0.7 + Wisdom × 0.8 + MentalEnergy × 0.2) × (1 + CultivationLevel/130)
DecisionMaking = (Intelligence × 0.6 + Wisdom × 0.9 + MentalEnergy × 0.3) × (1 + CultivationLevel/140)
ResourceManagement = (Intelligence × 0.8 + Wisdom × 0.6 + MentalEnergy × 0.2) × (1 + CultivationLevel/110)
StrategyFormulation = (Intelligence × 0.7 + Wisdom × 0.8 + MentalEnergy × 0.3) × (1 + CultivationLevel/130)
TacticalAnalysis = (Intelligence × 0.8 + Wisdom × 0.7 + MentalEnergy × 0.2) × (1 + CultivationLevel/120)
RiskAssessment = (Intelligence × 0.6 + Wisdom × 0.8 + MentalEnergy × 0.3) × (1 + CultivationLevel/130)

// Aura Stats
AuraRadius = (Charisma × 0.5 + SpiritualEnergy × 0.2 + CultivationLevel × 0.1) × (1 + CultivationLevel/100)
AuraStrength = (Willpower × 0.6 + SpiritualEnergy × 0.3 + CultivationLevel × 0.2) × (1 + CultivationLevel/120)
Presence = (Charisma × 0.8 + Willpower × 0.4 + SpiritualEnergy × 0.2) × (1 + CultivationLevel/150)
Awe = (Charisma × 0.6 + SpiritualEnergy × 0.4 + CultivationLevel × 0.3) × (1 + CultivationLevel/180)

// Proficiency Stats
WeaponMastery = (Strength × 0.4 + Agility × 0.4 + Intelligence × 0.2) × (1 + CultivationLevel/100)
SkillLevel = (Intelligence × 0.5 + Wisdom × 0.3 + CultivationLevel × 0.2) × (1 + CultivationLevel/120)
LifeSteal = (Luck × 0.01 + Strength × 0.005 + CultivationLevel × 0.003) × (1 + CultivationLevel/200)
CastSpeed = 1.0 + (Intelligence × 0.01 + Wisdom × 0.008 + SpiritualEnergy × 0.001)
WeightCapacity = (Strength × 10 + Endurance × 5 + PhysicalEnergy × 0.1) × (1 + CultivationLevel/100)
Persuasion = (Charisma × 0.6 + Intelligence × 0.3 + Personality × 0.4) × (1 + CultivationLevel/150)
MerchantPriceModifier = 1.0 + (Personality × 0.01 + Charisma × 0.005 + Luck × 0.003)
FactionReputationGain = 1.0 + (Charisma × 0.01 + Personality × 0.008 + CultivationLevel × 0.005)
```

### 2. Talent Amplifier Formulas
```go
// Talent Amplifiers
CultivationSpeed = 1.0 + (CultivationLevel × 0.01 + Luck × 0.005 + Intelligence × 0.003) × (1 + Age/1000)
EnergyEfficiency = 1.0 + (Wisdom × 0.01 + Willpower × 0.008 + CultivationLevel × 0.005) × (1 + Age/1000)
BreakthroughSuccess = 0.1 + (CultivationLevel × 0.002 + Luck × 0.001 + Willpower × 0.0008) × (1 + Age/1000)
SkillLearning = 1.0 + (Intelligence × 0.01 + Wisdom × 0.008 + CultivationLevel × 0.005) × (1 + Age/1000)
CombatEffectiveness = 1.0 + (Strength × 0.01 + Agility × 0.008 + CultivationLevel × 0.005) × (1 + Age/1000)
ResourceGathering = 1.0 + (Luck × 0.01 + Wisdom × 0.008 + CultivationLevel × 0.005) × (1 + Age/1000)
```

## Implementation Strategy

### Phase 1: Core Foundation
1. **Tạo Actor Core v2.0** với PrimaryCore và Derived stats
2. **Implement basic interfaces** (StatProvider, StatConsumer, StatResolver)
3. **Tạo stat calculation engine** cho derived stats
4. **Unit tests** cho core functionality

### Phase 2: Integration Layer
1. **Tạo integration interfaces** cho sub-systems
2. **Implement stat mapping** từ sub-systems sang Actor Core
3. **Tạo conversion utilities** cho different stat formats
4. **Integration tests** cho sub-system communication

### Phase 3: Sub-System Adapters
1. **RPG System adapter** cho level progression
2. **Kim Đan System adapter** cho cultivation stats
3. **Other sub-system adapters** as needed
4. **End-to-end tests** cho complete workflow

### Phase 4: Optimization & Performance
1. **Performance optimization** cho stat calculations
2. **Caching system** cho frequently accessed stats
3. **Memory optimization** cho large stat sets
4. **Load testing** cho production readiness

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
    
    -- Metadata
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    version INT64 NOT NULL DEFAULT 1
);
```

### Derived Stats Cache Table
```sql
CREATE TABLE derived_stats_cache (
    actor_id VARCHAR(255) PRIMARY KEY,
    stats_json JSON NOT NULL,
    calculated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    version INT64 NOT NULL DEFAULT 1,
    
    FOREIGN KEY (actor_id) REFERENCES actor_core_stats(actor_id)
);
```

## Benefits

### 1. Universal Foundation
- **Single source of truth** cho core stats
- **Consistent calculations** across all systems
- **Easy integration** với new sub-systems

### 2. Scalability
- **Modular design** cho easy expansion
- **Flexible stat system** cho customization
- **Performance optimized** cho large scale

### 3. Maintainability
- **Clear separation** of concerns
- **Comprehensive testing** cho reliability
- **Documentation** cho development

### 4. Flexibility
- **Sub-system independence** cho level progression
- **Universal cultivation stats** cho multiple systems
- **Easy stat mapping** between systems

---

*Tài liệu này sẽ được cập nhật khi có thay đổi trong design hoặc implementation.*
