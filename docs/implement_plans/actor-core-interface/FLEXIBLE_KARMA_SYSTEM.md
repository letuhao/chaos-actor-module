# Flexible Karma System - Actor Core v2.0

## Tổng Quan

Flexible Karma System được thiết kế để hỗ trợ nhiều loại khí vận/nghiệp lực/công đức/cống hiến khác nhau trong các thế giới, khu vực, tông môn, và quốc gia khác nhau. Hệ thống này cho phép tracking karma theo nhiều cấp độ và tích hợp với stat system để ảnh hưởng đến hiệu suất của nhân vật.

## Kiến Trúc

```
Flexible Karma System
├── Global Karma (tổng của tất cả thế giới)
├── World-specific Karma (theo từng thế giới)
├── Region-specific Karma (theo từng khu vực)
├── Sect-specific Karma (theo từng tông môn)
├── Nation-specific Karma (theo từng quốc gia)
├── Karma Categories (20+ categories)
├── Karma Types (flexible definitions)
└── Karma Calculation Engine
    ├── Total Karma Calculation
    ├── World/Region/Sect/Nation Karma
    ├── Weighted Score Calculation
    └── Stat Influence Calculation
```

## Karma Categories

### 1. Core Karma Categories (20 types)
```go
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
```

### 2. Karma Type Definition
```go
type KarmaType struct {
    Name        string  // Tên loại karma
    Category    string  // Danh mục: "Fortune", "Karma", "Merit", etc.
    Description string  // Mô tả
    MinValue    int64   // Giá trị tối thiểu
    MaxValue    int64   // Giá trị tối đa
    DefaultValue int64  // Giá trị mặc định
    IsPositive  bool    // Có phải là karma tích cực không
    Weight      float64 // Trọng số trong tính toán tổng
}
```

## World/Region/Sect/Nation Definitions

### 1. World Definition
```go
type WorldDefinition struct {
    Name        string   // Tên thế giới
    Type        string   // Loại: "Mortal", "Immortal", "Divine", "Demon", etc.
    Level       int64    // Cấp độ thế giới
    KarmaTypes  []string // Các loại karma có trong thế giới này
    Multiplier  float64  // Hệ số nhân cho karma
}
```

**Ví dụ Worlds:**
```go
// Mortal World
mortalWorld := &WorldDefinition{
    Name: "Mortal World",
    Type: "Mortal",
    Level: 1,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor"},
    Multiplier: 1.0,
}

// Immortal World
immortalWorld := &WorldDefinition{
    Name: "Immortal World",
    Type: "Immortal",
    Level: 5,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory", "Wisdom"},
    Multiplier: 2.0,
}

// Divine World
divineWorld := &WorldDefinition{
    Name: "Divine World",
    Type: "Divine",
    Level: 10,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory", "Wisdom", "Compassion", "Justice"},
    Multiplier: 5.0,
}
```

### 2. Region Definition
```go
type RegionDefinition struct {
    Name        string   // Tên khu vực
    World       string   // Thuộc thế giới nào
    Type        string   // Loại: "City", "Wilderness", "Dungeon", "Temple", etc.
    Level       int64    // Cấp độ khu vực
    KarmaTypes  []string // Các loại karma có trong khu vực này
    Multiplier  float64  // Hệ số nhân cho karma
}
```

**Ví dụ Regions:**
```go
// City Region
cityRegion := &RegionDefinition{
    Name: "Capital City",
    World: "Mortal World",
    Type: "City",
    Level: 3,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor"},
    Multiplier: 1.2,
}

// Temple Region
templeRegion := &RegionDefinition{
    Name: "Sacred Temple",
    World: "Immortal World",
    Type: "Temple",
    Level: 8,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory", "Wisdom", "Compassion"},
    Multiplier: 2.5,
}

// Dungeon Region
dungeonRegion := &RegionDefinition{
    Name: "Ancient Dungeon",
    World: "Mortal World",
    Type: "Dungeon",
    Level: 5,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Valor", "Honor"},
    Multiplier: 1.5,
}
```

### 3. Sect Definition
```go
type SectDefinition struct {
    Name        string   // Tên tông môn
    World       string   // Thuộc thế giới nào
    Type        string   // Loại: "Righteous", "Demonic", "Neutral", "Heretic", etc.
    Level       int64    // Cấp độ tông môn
    KarmaTypes  []string // Các loại karma có trong tông môn này
    Multiplier  float64  // Hệ số nhân cho karma
}
```

**Ví dụ Sects:**
```go
// Righteous Sect
righteousSect := &SectDefinition{
    Name: "Heavenly Sword Sect",
    World: "Immortal World",
    Type: "Righteous",
    Level: 7,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory", "Wisdom", "Justice", "Valor"},
    Multiplier: 2.0,
}

// Demonic Sect
demonicSect := &SectDefinition{
    Name: "Blood Demon Sect",
    World: "Immortal World",
    Type: "Demonic",
    Level: 6,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory", "Valor"},
    Multiplier: 1.8,
}

// Neutral Sect
neutralSect := &SectDefinition{
    Name: "Mystic Arts Sect",
    World: "Immortal World",
    Type: "Neutral",
    Level: 5,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Wisdom", "Balance"},
    Multiplier: 1.5,
}
```

### 4. Nation Definition
```go
type NationDefinition struct {
    Name        string   // Tên quốc gia
    World       string   // Thuộc thế giới nào
    Type        string   // Loại: "Empire", "Kingdom", "Republic", "Tribe", etc.
    Level       int64    // Cấp độ quốc gia
    KarmaTypes  []string // Các loại karma có trong quốc gia này
    Multiplier  float64  // Hệ số nhân cho karma
}
```

**Ví dụ Nations:**
```go
// Empire
empire := &NationDefinition{
    Name: "Great Empire",
    World: "Mortal World",
    Type: "Empire",
    Level: 8,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory", "Justice", "Loyalty"},
    Multiplier: 2.0,
}

// Kingdom
kingdom := &NationDefinition{
    Name: "Peaceful Kingdom",
    World: "Mortal World",
    Type: "Kingdom",
    Level: 6,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Peace", "Harmony"},
    Multiplier: 1.5,
}

// Republic
republic := &NationDefinition{
    Name: "Free Republic",
    World: "Mortal World",
    Type: "Republic",
    Level: 4,
    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Justice", "Freedom", "Truth"},
    Multiplier: 1.2,
}
```

## Karma Calculation Engine

### 1. Total Karma Calculation
```go
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
```

### 2. World/Region/Sect/Nation Karma
```go
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
```

### 3. Weighted Karma Score
```go
func (kce *KarmaCalculationEngine) CalculateWeightedKarmaScore(karmaType string) float64 {
    totalKarma := kce.CalculateTotalKarma(karmaType)
    
    // Get karma type definition
    if karmaDef, exists := kce.karmaSystem.KarmaTypes[karmaType]; exists {
        // Apply weight
        return float64(totalKarma) * karmaDef.Weight
    }
    
    return float64(totalKarma)
}
```

### 4. Karma Influence on Stats
```go
func (kce *KarmaCalculationEngine) CalculateKarmaInfluence(karmaType string, baseStat float64) float64 {
    weightedKarma := kce.CalculateWeightedKarmaScore(karmaType)
    
    // Karma influence formula: baseStat * (1 + karmaInfluence)
    karmaInfluence := weightedKarma / 10000.0  // Scale karma to reasonable influence
    
    return baseStat * (1.0 + karmaInfluence)
}
```

## Database Schema

### Karma System Table
```sql
CREATE TABLE actor_karma_system (
    actor_id VARCHAR(255) PRIMARY KEY,
    
    -- Global Karma
    global_karma JSON,
    
    -- World-specific Karma
    world_karma JSON,
    
    -- Region-specific Karma
    region_karma JSON,
    
    -- Sect-specific Karma
    sect_karma JSON,
    
    -- Nation-specific Karma
    nation_karma JSON,
    
    -- Metadata
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    version INT64 NOT NULL DEFAULT 1,
    
    FOREIGN KEY (actor_id) REFERENCES actor_core_stats(actor_id)
);
```

### World/Region/Sect/Nation Definitions Table
```sql
CREATE TABLE world_definitions (
    name VARCHAR(255) PRIMARY KEY,
    type VARCHAR(100) NOT NULL,
    level INT64 NOT NULL,
    karma_types JSON,
    multiplier FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE region_definitions (
    name VARCHAR(255) PRIMARY KEY,
    world VARCHAR(255) NOT NULL,
    type VARCHAR(100) NOT NULL,
    level INT64 NOT NULL,
    karma_types JSON,
    multiplier FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (world) REFERENCES world_definitions(name)
);

CREATE TABLE sect_definitions (
    name VARCHAR(255) PRIMARY KEY,
    world VARCHAR(255) NOT NULL,
    type VARCHAR(100) NOT NULL,
    level INT64 NOT NULL,
    karma_types JSON,
    multiplier FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (world) REFERENCES world_definitions(name)
);

CREATE TABLE nation_definitions (
    name VARCHAR(255) PRIMARY KEY,
    world VARCHAR(255) NOT NULL,
    type VARCHAR(100) NOT NULL,
    level INT64 NOT NULL,
    karma_types JSON,
    multiplier FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (world) REFERENCES world_definitions(name)
);
```

## Usage Examples

### 1. Basic Karma Usage
```go
// Get total karma
karmaSystem := actor.PrimaryCore.KarmaSystem
totalFortune := karmaSystem.CalculateTotalKarma("Fortune")

// Get world-specific karma
mortalWorldFortune := karmaSystem.CalculateWorldKarma("Mortal World", "Fortune")

// Get region-specific karma
cityFortune := karmaSystem.CalculateRegionKarma("Capital City", "Fortune")

// Get sect-specific karma
sectFortune := karmaSystem.CalculateSectKarma("Heavenly Sword Sect", "Fortune")

// Get nation-specific karma
nationFortune := karmaSystem.CalculateNationKarma("Great Empire", "Fortune")
```

### 2. Karma Influence on Stats
```go
// Calculate karma influence on stats
baseAttack := 1000.0
fortuneInfluence := karmaSystem.CalculateKarmaInfluence("Fortune", baseAttack)
// Result: baseAttack * (1 + fortuneInfluence)

baseCastingSpeed := 500.0
karmaInfluence := karmaSystem.CalculateKarmaInfluence("Karma", baseCastingSpeed)
// Result: baseCastingSpeed * (1 + karmaInfluence)
```

### 3. Karma Breakdown
```go
// Get detailed karma breakdown
breakdown := karmaSystem.GetKarmaBreakdown(actorID, "Fortune")
fmt.Printf("Total Fortune: %d\n", breakdown.TotalKarma)
fmt.Printf("World Fortune: %v\n", breakdown.WorldKarma)
fmt.Printf("Region Fortune: %v\n", breakdown.RegionKarma)
fmt.Printf("Sect Fortune: %v\n", breakdown.SectKarma)
fmt.Printf("Nation Fortune: %v\n", breakdown.NationKarma)
fmt.Printf("Weighted Score: %.2f\n", breakdown.WeightedScore)
fmt.Printf("Influence: %.2f\n", breakdown.Influence)
```

### 4. Multi-World Karma Tracking
```go
// Track karma across multiple worlds
actor := &MultiSystemActor{
    ID: "player123",
    ActiveSystems: []string{"RPG", "KimDan", "Magic"},
    KarmaSystem: &FlexibleKarmaSystem{
        GlobalKarma: map[string]int64{
            "Fortune": 1000,
            "Karma": 500,
            "Merit": 800,
        },
        WorldKarma: map[string]map[string]int64{
            "Mortal World": {
                "Fortune": 300,
                "Karma": 200,
                "Merit": 400,
            },
            "Immortal World": {
                "Fortune": 500,
                "Karma": 200,
                "Merit": 300,
                "Glory": 100,
            },
            "Divine World": {
                "Fortune": 200,
                "Karma": 100,
                "Merit": 100,
                "Glory": 50,
                "Wisdom": 75,
            },
        },
        RegionKarma: map[string]map[string]int64{
            "Capital City": {
                "Fortune": 100,
                "Karma": 50,
                "Merit": 200,
            },
            "Sacred Temple": {
                "Fortune": 150,
                "Karma": 75,
                "Merit": 100,
                "Glory": 25,
            },
        },
        SectKarma: map[string]map[string]int64{
            "Heavenly Sword Sect": {
                "Fortune": 200,
                "Karma": 100,
                "Merit": 150,
                "Glory": 50,
                "Justice": 75,
            },
        },
        NationKarma: map[string]map[string]int64{
            "Great Empire": {
                "Fortune": 100,
                "Karma": 50,
                "Merit": 100,
                "Justice": 25,
                "Loyalty": 50,
            },
        },
    },
}
```

## Benefits

### 1. Flexibility
- **Multi-World Support**: Hỗ trợ nhiều thế giới khác nhau
- **Multi-Level Tracking**: Tracking theo world/region/sect/nation
- **Customizable Karma Types**: Có thể tùy chỉnh loại karma

### 2. Realism
- **World-Specific Karma**: Mỗi thế giới có loại karma riêng
- **Sect/Nation Influence**: Tông môn và quốc gia ảnh hưởng đến karma
- **Weighted Calculation**: Tính toán có trọng số

### 3. Performance
- **Efficient Calculation**: Tính toán hiệu quả
- **Caching Support**: Hỗ trợ caching
- **Breakdown Transparency**: Minh bạch trong tính toán

### 4. Scalability
- **Modular Design**: Thiết kế modular
- **Easy Extension**: Dễ dàng mở rộng
- **Database Optimization**: Tối ưu hóa database

## Implementation Priority

### Phase 1: Core Karma System
1. **FlexibleKarmaSystem** struct
2. **Karma calculation engine**
3. **Basic karma categories**

### Phase 2: World/Region/Sect/Nation Integration
1. **World/Region/Sect/Nation definitions**
2. **Multi-level karma tracking**
3. **Weighted calculation**

### Phase 3: Advanced Features
1. **Karma influence on stats**
2. **Karma breakdown system**
3. **Caching optimization**

### Phase 4: Testing & Optimization
1. **Unit tests**
2. **Performance tests**
3. **Integration tests**

---

*Tài liệu này mô tả Flexible Karma System cho Actor Core v2.0, hỗ trợ tracking karma theo nhiều cấp độ và thế giới khác nhau.*
