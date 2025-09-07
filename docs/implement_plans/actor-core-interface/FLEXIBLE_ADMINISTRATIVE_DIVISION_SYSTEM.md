# Flexible Administrative Division System - Actor Core v2.0

## Tổng Quan

Flexible Administrative Division System được thiết kế để hỗ trợ nhiều cách phân chia đơn vị hành chính khác nhau, từ thế giới, lục địa, quốc gia đến chủng tộc, tông môn, và các đơn vị đặc biệt khác. Hệ thống này cho phép tạo ra các phân cấp hành chính linh hoạt và tích hợp với karma system.

## Kiến Trúc

```
Flexible Administrative Division System
├── Administrative Divisions (flexible hierarchy)
│   ├── World Divisions (thế giới)
│   ├── Geographic Divisions (địa lý)
│   ├── Political Divisions (chính trị)
│   ├── Social Divisions (xã hội)
│   ├── Religious Divisions (tôn giáo)
│   ├── Economic Divisions (kinh tế)
│   ├── Military Divisions (quân sự)
│   ├── Cultural Divisions (văn hóa)
│   ├── Natural Divisions (tự nhiên)
│   └── Mystical Divisions (thần bí)
├── Division Types (100+ types)
├── Division Hierarchy (flexible levels)
├── Cross-Division Relationships
└── Karma Integration
    ├── Division-specific Karma
    ├── Karma Multipliers
    └── Karma Types per Division
```

## Division Types

### 1. Geographic Divisions (20 types)
```go
const (
    DIVISION_WORLD      = "World"      // Thế giới
    DIVISION_CONTINENT  = "Continent"  // Lục địa
    DIVISION_REALM      = "Realm"      // Vương quốc/Realm
    DIVISION_NATION     = "Nation"     // Quốc gia
    DIVISION_REGION     = "Region"     // Khu vực
    DIVISION_DISTRICT   = "District"   // Quận/Huyện
    DIVISION_CITY       = "City"       // Thành phố
    DIVISION_VILLAGE    = "Village"    // Làng
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
)
```

### 2. Political Divisions (15 types)
```go
const (
    DIVISION_EMPIRE     = "Empire"     // Đế quốc
    DIVISION_KINGDOM    = "Kingdom"    // Vương quốc
    DIVISION_REPUBLIC   = "Republic"   // Cộng hòa
    DIVISION_FEDERATION = "Federation" // Liên bang
    DIVISION_ALLIANCE   = "Alliance"   // Liên minh
    DIVISION_CONFEDERACY = "Confederacy" // Liên minh lỏng lẻo
    DIVISION_DEMOCRACY  = "Democracy"  // Dân chủ
    DIVISION_MONARCHY   = "Monarchy"   // Quân chủ
    DIVISION_DICTATORSHIP = "Dictatorship" // Độc tài
    DIVISION_THEOCRACY  = "Theocracy"  // Thần quyền
    DIVISION_OLIGARCHY  = "Oligarchy"  // Tài phiệt
    DIVISION_ANARCHY    = "Anarchy"    // Vô chính phủ
    DIVISION_TRIBAL     = "Tribal"     // Bộ lạc
    DIVISION_CLAN       = "Clan"       // Gia tộc
    DIVISION_FAMILY     = "Family"     // Gia đình
)
```

### 3. Social Divisions (15 types)
```go
const (
    DIVISION_RACE       = "Race"       // Chủng tộc
    DIVISION_TRIBE      = "Tribe"      // Bộ tộc
    DIVISION_CLAN       = "Clan"       // Gia tộc
    DIVISION_FAMILY     = "Family"     // Gia đình
    DIVISION_GUILD      = "Guild"      // Hội nghề nghiệp
    DIVISION_ORDER      = "Order"      // Hội/Order
    DIVISION_SOCIETY    = "Society"    // Xã hội
    DIVISION_COMMUNITY  = "Community"  // Cộng đồng
    DIVISION_GROUP      = "Group"      // Nhóm
    DIVISION_ORGANIZATION = "Organization" // Tổ chức
    DIVISION_ASSOCIATION = "Association" // Hiệp hội
    DIVISION_UNION      = "Union"      // Liên đoàn
    DIVISION_COALITION  = "Coalition"  // Liên minh
    DIVISION_FRATERNITY = "Fraternity" // Hội huynh đệ
    DIVISION_SORORITY   = "Sorority"   // Hội chị em
)
```

### 4. Religious Divisions (10 types)
```go
const (
    DIVISION_SECT       = "Sect"       // Tông môn
    DIVISION_ORDER      = "Order"      // Hội/Order
    DIVISION_CHURCH     = "Church"     // Giáo hội
    DIVISION_TEMPLE     = "Temple"     // Chùa/Đền
    DIVISION_MONASTERY  = "Monastery"  // Tu viện
    DIVISION_SHRINE     = "Shrine"     // Đền thờ
    DIVISION_CATHEDRAL  = "Cathedral"  // Nhà thờ lớn
    DIVISION_MOSQUE     = "Mosque"     // Thánh đường
    DIVISION_SYNAGOGUE  = "Synagogue"  // Giáo đường
    DIVISION_PAGODA     = "Pagoda"     // Chùa
)
```

### 5. Economic Divisions (10 types)
```go
const (
    DIVISION_MARKET     = "Market"     // Chợ
    DIVISION_PORT       = "Port"       // Cảng
    DIVISION_HARBOR     = "Harbor"     // Bến cảng
    DIVISION_AIRPORT    = "Airport"    // Sân bay
    DIVISION_STATION    = "Station"    // Ga
    DIVISION_DEPOT      = "Depot"      // Kho
    DIVISION_WAREHOUSE  = "Warehouse"  // Nhà kho
    DIVISION_VAULT      = "Vault"      // Kho báu
    DIVISION_TREASURY   = "Treasury"   // Kho bạc
    DIVISION_BANK       = "Bank"       // Ngân hàng
)
```

### 6. Military Divisions (10 types)
```go
const (
    DIVISION_ARMORY     = "Armory"     // Kho vũ khí
    DIVISION_BARRACKS   = "Barracks"   // Doanh trại
    DIVISION_FORTRESS   = "Fortress"   // Pháo đài
    DIVISION_CASTLE     = "Castle"     // Lâu đài
    DIVISION_PALACE     = "Palace"     // Cung điện
    DIVISION_MANSION    = "Mansion"    // Biệt thự
    DIVISION_ESTATE     = "Estate"     // Điền trang
    DIVISION_GARRISON   = "Garrison"   // Đồn trú
    DIVISION_OUTPOST    = "Outpost"    // Tiền đồn
    DIVISION_WATCHTOWER = "Watchtower" // Tháp canh
)
```

### 7. Cultural Divisions (10 types)
```go
const (
    DIVISION_ACADEMY    = "Academy"    // Học viện
    DIVISION_UNIVERSITY = "University" // Đại học
    DIVISION_LIBRARY    = "Library"    // Thư viện
    DIVISION_LABORATORY = "Laboratory" // Phòng thí nghiệm
    DIVISION_WORKSHOP   = "Workshop"   // Xưởng
    DIVISION_FACTORY    = "Factory"    // Nhà máy
    DIVISION_MUSEUM     = "Museum"     // Bảo tàng
    DIVISION_THEATER    = "Theater"    // Nhà hát
    DIVISION_STADIUM    = "Stadium"    // Sân vận động
    DIVISION_ARENA      = "Arena"      // Đấu trường
)
```

### 8. Natural Divisions (15 types)
```go
const (
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
)
```

### 9. Mystical Divisions (10 types)
```go
const (
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
)
```

## Administrative Division Structure

### 1. Division Definition
```go
type AdministrativeDivision struct {
    ID          string   // Unique identifier
    Name        string   // Display name
    Type        string   // Division type
    Level       int64    // Hierarchy level (0 = highest)
    Parent      string   // Parent division ID
    Children    []string // Child division IDs
    Attributes  map[string]interface{}  // Flexible attributes
    KarmaTypes  []string // Available karma types
    Multiplier  float64  // Karma multiplier
    CreatedAt   int64    // Creation timestamp
    UpdatedAt   int64    // Last update timestamp
}
```

### 2. Division Type Definition
```go
type DivisionType struct {
    Name        string   // Type name
    Description string   // Description
    Level       int64    // Default hierarchy level
    Attributes  []string // Required attributes
    KarmaTypes  []string // Default karma types
    Multiplier  float64  // Default multiplier
    IsActive    bool     // Is this type active
}
```

### 3. Division Relationship
```go
type DivisionRelationship struct {
    Type        string  // Relationship type
    Strength    float64 // Relationship strength (0.0 - 1.0)
    Description string  // Relationship description
    CreatedAt   int64   // Creation timestamp
    UpdatedAt   int64   // Last update timestamp
}
```

## Division Hierarchy Examples

### 1. Geographic Hierarchy
```
World (Level 0)
├── Continent (Level 1)
│   ├── Nation (Level 2)
│   │   ├── Region (Level 3)
│   │   │   ├── District (Level 4)
│   │   │   │   ├── City (Level 5)
│   │   │   │   │   ├── Village (Level 6)
│   │   │   │   │   └── Neighborhood (Level 7)
│   │   │   │   └── County (Level 5)
│   │   │   └── Province (Level 4)
│   │   └── State (Level 3)
│   └── Territory (Level 2)
└── Realm (Level 1)
    └── Domain (Level 2)
        └── Fief (Level 3)
```

### 2. Political Hierarchy
```
Empire (Level 0)
├── Kingdom (Level 1)
│   ├── Republic (Level 2)
│   │   ├── State (Level 3)
│   │   │   ├── County (Level 4)
│   │   │   └── Municipality (Level 4)
│   │   └── Province (Level 3)
│   └── Federation (Level 2)
│       └── Alliance (Level 3)
└── Confederacy (Level 1)
    └── Tribal (Level 2)
        └── Clan (Level 3)
```

### 3. Social Hierarchy
```
Race (Level 0)
├── Tribe (Level 1)
│   ├── Clan (Level 2)
│   │   ├── Family (Level 3)
│   │   │   ├── Extended Family (Level 4)
│   │   │   └── Nuclear Family (Level 4)
│   │   └── Guild (Level 3)
│   └── Society (Level 2)
│       └── Community (Level 3)
└── Organization (Level 1)
    └── Association (Level 2)
        └── Union (Level 3)
```

### 4. Religious Hierarchy
```
Sect (Level 0)
├── Order (Level 1)
│   ├── Church (Level 2)
│   │   ├── Cathedral (Level 3)
│   │   │   ├── Chapel (Level 4)
│   │   │   └── Shrine (Level 4)
│   │   └── Temple (Level 3)
│   └── Monastery (Level 2)
│       └── Hermitage (Level 3)
└── Guild (Level 1)
    └── Society (Level 2)
        └── Community (Level 3)
```

## Division Relationships

### 1. Relationship Types
```go
const (
    RELATIONSHIP_CONTAINS    = "Contains"    // Chứa đựng
    RELATIONSHIP_BELONGS_TO  = "BelongsTo"   // Thuộc về
    RELATIONSHIP_ALLIED      = "Allied"      // Đồng minh
    RELATIONSHIP_ENEMY       = "Enemy"       // Kẻ thù
    RELATIONSHIP_NEUTRAL     = "Neutral"     // Trung lập
    RELATIONSHIP_TRADING     = "Trading"     // Thương mại
    RELATIONSHIP_CULTURAL    = "Cultural"    // Văn hóa
    RELATIONSHIP_RELIGIOUS   = "Religious"   // Tôn giáo
    RELATIONSHIP_ECONOMIC    = "Economic"    // Kinh tế
    RELATIONSHIP_MILITARY    = "Military"    // Quân sự
    RELATIONSHIP_DIPLOMATIC  = "Diplomatic"  // Ngoại giao
    RELATIONSHIP_EDUCATIONAL = "Educational" // Giáo dục
    RELATIONSHIP_SCIENTIFIC  = "Scientific"  // Khoa học
    RELATIONSHIP_ARTISTIC    = "Artistic"    // Nghệ thuật
    RELATIONSHIP_SPORTING    = "Sporting"    // Thể thao
    RELATIONSHIP_MEDICAL     = "Medical"     // Y tế
    RELATIONSHIP_LEGAL       = "Legal"       // Pháp lý
    RELATIONSHIP_TECHNICAL   = "Technical"   // Kỹ thuật
    RELATIONSHIP_ACADEMIC    = "Academic"    // Học thuật
    RELATIONSHIP_PROFESSIONAL = "Professional" // Nghề nghiệp
)
```

### 2. Relationship Examples
```go
// Geographic relationships
worldToContinent := &DivisionRelationship{
    Type: "Contains",
    Strength: 1.0,
    Description: "World contains continents",
}

continentToNation := &DivisionRelationship{
    Type: "Contains",
    Strength: 0.9,
    Description: "Continent contains nations",
}

// Political relationships
empireToKingdom := &DivisionRelationship{
    Type: "Contains",
    Strength: 1.0,
    Description: "Empire contains kingdoms",
}

kingdomToRepublic := &DivisionRelationship{
    Type: "Allied",
    Strength: 0.8,
    Description: "Kingdom allied with republic",
}

// Social relationships
raceToTribe := &DivisionRelationship{
    Type: "Contains",
    Strength: 1.0,
    Description: "Race contains tribes",
}

tribeToClan := &DivisionRelationship{
    Type: "Contains",
    Strength: 0.9,
    Description: "Tribe contains clans",
}

// Religious relationships
sectToOrder := &DivisionRelationship{
    Type: "Contains",
    Strength: 1.0,
    Description: "Sect contains orders",
}

orderToChurch := &DivisionRelationship{
    Type: "Contains",
    Strength: 0.9,
    Description: "Order contains churches",
}
```

## Karma Integration

### 1. Division-specific Karma
```go
type FlexibleKarmaSystem struct {
    // Global Karma (tổng của tất cả thế giới)
    GlobalKarma map[string]int64  // karmaType -> totalValue
    
    // Division-specific Karma (theo từng đơn vị hành chính)
    DivisionKarma map[string]map[string]int64  // divisionType -> divisionName -> karmaType -> value
}
```

### 2. Karma Multipliers by Division Type
```go
// Division type karma multipliers
var divisionKarmaMultipliers = map[string]float64{
    "World":      5.0,  // Highest multiplier
    "Continent":  4.0,
    "Realm":      3.5,
    "Nation":     3.0,
    "Region":     2.5,
    "District":   2.0,
    "City":       1.5,
    "Village":    1.0,
    "Race":       4.0,
    "Sect":       3.5,
    "Guild":      2.0,
    "Family":     1.0,
}
```

### 3. Karma Types by Division Type
```go
// Division type karma types
var divisionKarmaTypes = map[string][]string{
    "World":      {"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory", "Wisdom", "Compassion", "Justice"},
    "Continent":  {"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory"},
    "Nation":     {"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Justice", "Loyalty"},
    "Region":     {"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor"},
    "City":       {"Fortune", "Karma", "Merit", "Contribution", "Reputation"},
    "Village":    {"Fortune", "Karma", "Merit", "Contribution"},
    "Race":       {"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory", "Wisdom"},
    "Sect":       {"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory", "Wisdom", "Compassion", "Justice", "Valor"},
    "Guild":      {"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor"},
    "Family":     {"Fortune", "Karma", "Merit", "Contribution", "Love", "Loyalty"},
}
```

## Database Schema

### Administrative Divisions Table
```sql
CREATE TABLE administrative_divisions (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(100) NOT NULL,
    level INT64 NOT NULL,
    parent_id VARCHAR(255),
    attributes JSON,
    karma_types JSON,
    multiplier FLOAT NOT NULL DEFAULT 1.0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES administrative_divisions(id)
);

CREATE TABLE division_relationships (
    id VARCHAR(255) PRIMARY KEY,
    division_a VARCHAR(255) NOT NULL,
    division_b VARCHAR(255) NOT NULL,
    relationship_type VARCHAR(100) NOT NULL,
    strength FLOAT NOT NULL DEFAULT 0.0,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (division_a) REFERENCES administrative_divisions(id),
    FOREIGN KEY (division_b) REFERENCES administrative_divisions(id)
);

CREATE TABLE division_types (
    name VARCHAR(100) PRIMARY KEY,
    description TEXT,
    level INT64 NOT NULL,
    attributes JSON,
    karma_types JSON,
    multiplier FLOAT NOT NULL DEFAULT 1.0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## Usage Examples

### 1. Basic Division Usage
```go
// Get division by ID
adminSystem := actor.PrimaryCore.AdministrativeSystem
division := adminSystem.GetDivision(actorID, "world_001")

// Get divisions by type
worlds := adminSystem.GetDivisionsByType(actorID, "World")
nations := adminSystem.GetDivisionsByType(actorID, "Nation")
sects := adminSystem.GetDivisionsByType(actorID, "Sect")

// Get divisions by level
level0Divisions := adminSystem.GetDivisionsByLevel(actorID, 0) // Worlds
level1Divisions := adminSystem.GetDivisionsByLevel(actorID, 1) // Continents
level2Divisions := adminSystem.GetDivisionsByLevel(actorID, 2) // Nations
```

### 2. Division Hierarchy
```go
// Get parent division
parent := adminSystem.GetParentDivision(actorID, "city_001")

// Get child divisions
children := adminSystem.GetChildDivisions(actorID, "nation_001")

// Get full hierarchy
hierarchy := adminSystem.GetDivisionHierarchy(actorID, "village_001")
// Result: [World, Continent, Nation, Region, District, City, Village]
```

### 3. Division Relationships
```go
// Get division relationships
relationships := adminSystem.GetDivisionRelationships(actorID, "nation_001")

// Add division relationship
relationship := &DivisionRelationship{
    Type: "Allied",
    Strength: 0.8,
    Description: "Allied nations",
}
adminSystem.AddDivisionRelationship(actorID, "nation_001", "nation_002", relationship)
```

### 4. Division Attributes
```go
// Get division attributes
attributes := adminSystem.GetDivisionAttributes(actorID, "city_001")
// Result: map[string]interface{}{
//     "Population": 1000000,
//     "Area": 500.0,
//     "Climate": "Temperate",
//     "Language": "Common",
//     "Currency": "Gold",
// }

// Update division attributes
newAttributes := map[string]interface{}{
    "Population": 1200000,
    "Area": 550.0,
    "Climate": "Temperate",
    "Language": "Common",
    "Currency": "Gold",
    "Technology": "Advanced",
}
adminSystem.UpdateDivisionAttributes(actorID, "city_001", newAttributes)
```

### 5. Division Karma
```go
// Get division karma
karmaSystem := actor.PrimaryCore.KarmaSystem
worldFortune := karmaSystem.GetDivisionKarma(actorID, "World", "Mortal World", "Fortune")
nationKarma := karmaSystem.GetDivisionKarma(actorID, "Nation", "Great Empire", "Karma")
sectMerit := karmaSystem.GetDivisionKarma(actorID, "Sect", "Heavenly Sword Sect", "Merit")

// Update division karma
karmaSystem.UpdateDivisionKarma(actorID, "World", "Mortal World", "Fortune", 1000)
karmaSystem.UpdateDivisionKarma(actorID, "Nation", "Great Empire", "Karma", 500)
karmaSystem.UpdateDivisionKarma(actorID, "Sect", "Heavenly Sword Sect", "Merit", 800)
```

### 6. Multi-Division Actor
```go
// Actor with multiple divisions
actor := &MultiSystemActor{
    ID: "player123",
    AdministrativeSystem: &FlexibleAdministrativeSystem{
        Divisions: map[string]map[string]AdministrativeDivision{
            "World": {
                "Mortal World": {
                    ID: "world_001",
                    Name: "Mortal World",
                    Type: "World",
                    Level: 0,
                    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory", "Wisdom", "Compassion", "Justice"},
                    Multiplier: 5.0,
                },
            },
            "Nation": {
                "Great Empire": {
                    ID: "nation_001",
                    Name: "Great Empire",
                    Type: "Nation",
                    Level: 2,
                    Parent: "continent_001",
                    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Justice", "Loyalty"},
                    Multiplier: 3.0,
                },
            },
            "Sect": {
                "Heavenly Sword Sect": {
                    ID: "sect_001",
                    Name: "Heavenly Sword Sect",
                    Type: "Sect",
                    Level: 2,
                    Parent: "world_001",
                    KarmaTypes: []string{"Fortune", "Karma", "Merit", "Contribution", "Reputation", "Honor", "Glory", "Wisdom", "Compassion", "Justice", "Valor"},
                    Multiplier: 3.5,
                },
            },
        },
    },
}
```

## Benefits

### 1. Flexibility
- **Multiple Division Types**: Hỗ trợ 100+ loại đơn vị hành chính
- **Flexible Hierarchy**: Phân cấp linh hoạt
- **Custom Attributes**: Thuộc tính tùy chỉnh

### 2. Realism
- **Real-world Divisions**: Dựa trên thực tế
- **Complex Relationships**: Mối quan hệ phức tạp
- **Cultural Diversity**: Đa dạng văn hóa

### 3. Performance
- **Efficient Queries**: Truy vấn hiệu quả
- **Caching Support**: Hỗ trợ caching
- **Scalable Design**: Thiết kế có thể mở rộng

### 4. Integration
- **Karma Integration**: Tích hợp với karma system
- **Stat Influence**: Ảnh hưởng đến stats
- **Multi-System Support**: Hỗ trợ nhiều hệ thống

## Implementation Priority

### Phase 1: Core Division System
1. **FlexibleAdministrativeSystem** struct
2. **Division management** functions
3. **Basic division types**

### Phase 2: Hierarchy and Relationships
1. **Division hierarchy** system
2. **Division relationships** system
3. **Division attributes** system

### Phase 3: Karma Integration
1. **Division-specific karma** tracking
2. **Karma multipliers** by division type
3. **Karma types** per division

### Phase 4: Testing & Optimization
1. **Unit tests**
2. **Performance tests**
3. **Integration tests**

---

*Tài liệu này mô tả Flexible Administrative Division System cho Actor Core v2.0, hỗ trợ nhiều cách phân chia đơn vị hành chính khác nhau.*
