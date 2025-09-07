# Flexible Configuration System - Actor Core v2.0

## Tổng Quan

Flexible Configuration System được thiết kế để làm cho Actor Core v2.0 hoàn toàn flexible, cho phép tùy chỉnh tất cả stats, formulas, types, và bounds thông qua configuration files thay vì hardcode.

## Kiến Trúc

```
Flexible Configuration System
├── Stat Definitions (configurable primary/derived stats)
├── Formula Engine (configurable calculation formulas)
├── Type Registry (configurable damage/defence/amplifier types)
├── Clamp System (configurable bounds and limits)
├── Category System (configurable stat categories)
└── Validation Engine (configurable validation rules)
```

## 1. Flexible Stat Definitions

### Stat Definition Structure
```go
type StatDefinition struct {
    ID          string                 // Unique identifier
    Name        string                 // Display name
    Type        StatType               // PRIMARY, DERIVED, CUSTOM
    Category    string                 // Category (Basic, Physical, Cultivation, etc.)
    DataType    DataType               // INT64, FLOAT64, BOOLEAN, STRING, MAP
    DefaultValue interface{}           // Default value
    MinValue    interface{}           // Minimum value
    MaxValue    interface{}           // Maximum value
    Description string                 // Description
    IsActive    bool                  // Is this stat active
    Dependencies []string             // Dependencies on other stats
    Tags        []string              // Tags for grouping
    CreatedAt   int64                 // Creation timestamp
    UpdatedAt   int64                 // Last update timestamp
}

type StatType int
const (
    STAT_PRIMARY StatType = iota
    STAT_DERIVED
    STAT_CUSTOM
)

type DataType int
const (
    DATA_INT64 DataType = iota
    DATA_FLOAT64
    DATA_BOOLEAN
    DATA_STRING
    DATA_MAP
)
```

### Stat Configuration Example
```json
{
  "stats": {
    "vitality": {
      "id": "vitality",
      "name": "Vitality",
      "type": "PRIMARY",
      "category": "Basic",
      "dataType": "INT64",
      "defaultValue": 10,
      "minValue": 1,
      "maxValue": 999999,
      "description": "Overall health and resilience",
      "isActive": true,
      "dependencies": [],
      "tags": ["health", "survival", "core"]
    },
    "hpMax": {
      "id": "hpMax",
      "name": "Maximum Health",
      "type": "DERIVED",
      "category": "Combat",
      "dataType": "FLOAT64",
      "defaultValue": 100.0,
      "minValue": 1.0,
      "maxValue": 999999.0,
      "description": "Maximum health points",
      "isActive": true,
      "dependencies": ["vitality", "endurance", "constitution"],
      "tags": ["health", "combat", "derived"]
    }
  }
}
```

## 2. Flexible Formula Engine

### Formula Definition Structure
```go
type FormulaDefinition struct {
    ID          string                 // Unique identifier
    StatID      string                 // Target stat ID
    Name        string                 // Formula name
    Type        FormulaType            // CALCULATION, CONDITIONAL, LOOKUP
    Expression  string                 // Formula expression
    Dependencies []string              // Required stats
    Conditions  []FormulaCondition     // Conditional logic
    Priority    int                    // Calculation priority
    IsActive    bool                  // Is this formula active
    Description string                 // Description
    CreatedAt   int64                 // Creation timestamp
    UpdatedAt   int64                 // Last update timestamp
}

type FormulaType int
const (
    FORMULA_CALCULATION FormulaType = iota
    FORMULA_CONDITIONAL
    FORMULA_LOOKUP
)

type FormulaCondition struct {
    Condition string  // Condition expression
    Formula   string  // Formula to use if condition is true
    Priority  int     // Priority within conditions
}
```

### Formula Configuration Example
```json
{
  "formulas": {
    "hpMax_basic": {
      "id": "hpMax_basic",
      "statId": "hpMax",
      "name": "Basic HP Calculation",
      "type": "CALCULATION",
      "expression": "(vitality * 15 + endurance * 10 + constitution * 5) * (1 + level / 100)",
      "dependencies": ["vitality", "endurance", "constitution", "level"],
      "conditions": [],
      "priority": 1,
      "isActive": true,
      "description": "Basic HP calculation formula"
    },
    "hpMax_cultivation": {
      "id": "hpMax_cultivation",
      "statId": "hpMax",
      "name": "Cultivation HP Bonus",
      "type": "CONDITIONAL",
      "expression": "cultivationLevel > 0 ? hpMax * (1 + cultivationLevel / 100) : hpMax",
      "dependencies": ["hpMax", "cultivationLevel"],
      "conditions": [
        {
          "condition": "cultivationLevel > 0",
          "formula": "hpMax * (1 + cultivationLevel / 100)",
          "priority": 1
        }
      ],
      "priority": 2,
      "isActive": true,
      "description": "Cultivation level HP bonus"
    }
  }
}
```

## 3. Flexible Type Registry

### Type Definition Structure
```go
type TypeDefinition struct {
    ID          string                 // Unique identifier
    Name        string                 // Display name
    Type        TypeCategory           // DAMAGE, DEFENCE, AMPLIFIER, ENERGY, STATUS
    Category    string                 // Category (Physical, Elemental, Magical, etc.)
    Attributes  map[string]interface{} // Type-specific attributes
    Dependencies []string              // Required stats
    IsActive    bool                  // Is this type active
    Description string                 // Description
    CreatedAt   int64                 // Creation timestamp
    UpdatedAt   int64                 // Last update timestamp
}

type TypeCategory int
const (
    TYPE_DAMAGE TypeCategory = iota
    TYPE_DEFENCE
    TYPE_AMPLIFIER
    TYPE_ENERGY
    TYPE_STATUS
)
```

### Type Configuration Example
```json
{
  "types": {
    "damage_physical": {
      "id": "damage_physical",
      "name": "Physical Damage",
      "type": "DAMAGE",
      "category": "Physical",
      "attributes": {
        "isElemental": false,
        "isPhysical": true,
        "isMagical": false,
        "affectsHP": true,
        "affectsStamina": false,
        "canSpread": false
      },
      "dependencies": ["strength", "agility"],
      "isActive": true,
      "description": "Physical damage type"
    },
    "defence_armor": {
      "id": "defence_armor",
      "name": "Armor Defence",
      "type": "DEFENCE",
      "category": "Physical",
      "attributes": {
        "damageTypes": ["damage_physical", "damage_piercing", "damage_slashing", "damage_blunt"],
        "resistCap": 0.95,
        "drainCap": 0.5,
        "reflectCap": 0.3
      },
      "dependencies": ["constitution", "endurance"],
      "isActive": true,
      "description": "Armor-based physical defence"
    }
  }
}
```

## 4. Flexible Clamp System

### Clamp Definition Structure
```go
type ClampDefinition struct {
    ID          string                 // Unique identifier
    StatID      string                 // Target stat ID
    Name        string                 // Clamp name
    Type        ClampType              // MIN, MAX, RANGE, SOFT_CAP
    Value       interface{}            // Clamp value
    SoftCap     interface{}            // Soft cap value (for SOFT_CAP type)
    SoftCapRate float64                // Soft cap rate (for SOFT_CAP type)
    IsActive    bool                  // Is this clamp active
    Description string                 // Description
    CreatedAt   int64                 // Creation timestamp
    UpdatedAt   int64                 // Last update timestamp
}

type ClampType int
const (
    CLAMP_MIN ClampType = iota
    CLAMP_MAX
    CLAMP_RANGE
    CLAMP_SOFT_CAP
)
```

### Clamp Configuration Example
```json
{
  "clamps": {
    "hpMax_min": {
      "id": "hpMax_min",
      "statId": "hpMax",
      "name": "HP Minimum",
      "type": "MIN",
      "value": 1.0,
      "softCap": null,
      "softCapRate": 0.0,
      "isActive": true,
      "description": "Minimum HP is 1"
    },
    "haste_range": {
      "id": "haste_range",
      "statId": "haste",
      "name": "Haste Range",
      "type": "RANGE",
      "value": [0.5, 2.0],
      "softCap": null,
      "softCapRate": 0.0,
      "isActive": true,
      "description": "Haste must be between 0.5 and 2.0"
    },
    "cultivationLevel_soft_cap": {
      "id": "cultivationLevel_soft_cap",
      "statId": "cultivationLevel",
      "name": "Cultivation Soft Cap",
      "type": "SOFT_CAP",
      "value": 100,
      "softCap": 80,
      "softCapRate": 0.1,
      "isActive": true,
      "description": "Cultivation level soft cap at 80 with 10% reduction"
    }
  }
}
```

## 5. Flexible Category System

### Category Definition Structure
```go
type CategoryDefinition struct {
    ID          string                 // Unique identifier
    Name        string                 // Display name
    Type        CategoryType           // STAT, FORMULA, TYPE, CLAMP
    Parent      string                 // Parent category ID
    Children    []string               // Child category IDs
    Stats       []string               // Stats in this category
    Formulas    []string               // Formulas in this category
    Types       []string               // Types in this category
    Clamps      []string               // Clamps in this category
    IsActive    bool                  // Is this category active
    Description string                 // Description
    CreatedAt   int64                 // Creation timestamp
    UpdatedAt   int64                 // Last update timestamp
}

type CategoryType int
const (
    CATEGORY_STAT CategoryType = iota
    CATEGORY_FORMULA
    CATEGORY_TYPE
    CATEGORY_CLAMP
)
```

### Category Configuration Example
```json
{
  "categories": {
    "basic_stats": {
      "id": "basic_stats",
      "name": "Basic Stats",
      "type": "STAT",
      "parent": null,
      "children": ["physical_stats", "mental_stats"],
      "stats": ["vitality", "endurance", "constitution", "intelligence", "wisdom", "charisma", "willpower", "luck", "fate", "karma"],
      "formulas": [],
      "types": [],
      "clamps": [],
      "isActive": true,
      "description": "Basic character attributes"
    },
    "combat_stats": {
      "id": "combat_stats",
      "name": "Combat Stats",
      "type": "STAT",
      "parent": null,
      "children": ["offensive_stats", "defensive_stats"],
      "stats": ["hpMax", "stamina", "accuracy", "penetration", "armorClass", "evasion"],
      "formulas": ["hpMax_basic", "stamina_basic", "accuracy_basic"],
      "types": ["damage_physical", "defence_armor"],
      "clamps": ["hpMax_min", "stamina_min"],
      "isActive": true,
      "description": "Combat-related statistics"
    }
  }
}
```

## 6. Flexible Validation Engine

### Validation Rule Structure
```go
type ValidationRule struct {
    ID          string                 // Unique identifier
    Name        string                 // Rule name
    Type        ValidationType         // RANGE, DEPENDENCY, FORMULA, CUSTOM
    Target      string                 // Target stat/type/formula ID
    Condition   string                 // Validation condition
    Message     string                 // Error message
    Severity    ValidationSeverity     // ERROR, WARNING, INFO
    IsActive    bool                  // Is this rule active
    Description string                 // Description
    CreatedAt   int64                 // Creation timestamp
    UpdatedAt   int64                 // Last update timestamp
}

type ValidationType int
const (
    VALIDATION_RANGE ValidationType = iota
    VALIDATION_DEPENDENCY
    VALIDATION_FORMULA
    VALIDATION_CUSTOM
)

type ValidationSeverity int
const (
    SEVERITY_ERROR ValidationSeverity = iota
    SEVERITY_WARNING
    SEVERITY_INFO
)
```

### Validation Configuration Example
```json
{
  "validations": {
    "hpMax_range": {
      "id": "hpMax_range",
      "name": "HP Range Validation",
      "type": "RANGE",
      "target": "hpMax",
      "condition": "value >= 1.0 && value <= 999999.0",
      "message": "HP must be between 1 and 999999",
      "severity": "ERROR",
      "isActive": true,
      "description": "Validates HP range"
    },
    "formula_dependency": {
      "id": "formula_dependency",
      "name": "Formula Dependency Validation",
      "type": "DEPENDENCY",
      "target": "hpMax_basic",
      "condition": "all(dependencies, dep => hasStat(dep))",
      "message": "All formula dependencies must exist",
      "severity": "ERROR",
      "isActive": true,
      "description": "Validates formula dependencies"
    }
  }
}
```

## 7. Configuration Management

### Configuration Manager
```go
type ConfigurationManager struct {
    stats       map[string]*StatDefinition
    formulas    map[string]*FormulaDefinition
    types       map[string]*TypeDefinition
    clamps      map[string]*ClampDefinition
    categories  map[string]*CategoryDefinition
    validations map[string]*ValidationRule
}

// Load configuration from file
func (cm *ConfigurationManager) LoadFromFile(filename string) error

// Save configuration to file
func (cm *ConfigurationManager) SaveToFile(filename string) error

// Validate configuration
func (cm *ConfigurationManager) Validate() []ValidationError

// Get stat definition
func (cm *ConfigurationManager) GetStat(id string) (*StatDefinition, error)

// Get formula definition
func (cm *ConfigurationManager) GetFormula(id string) (*FormulaDefinition, error)

// Get type definition
func (cm *ConfigurationManager) GetType(id string) (*TypeDefinition, error)

// Get clamp definition
func (cm *ConfigurationManager) GetClamp(id string) (*ClampDefinition, error)

// Get category definition
func (cm *ConfigurationManager) GetCategory(id string) (*CategoryDefinition, error)

// Add/Update/Delete definitions
func (cm *ConfigurationManager) AddStat(stat *StatDefinition) error
func (cm *ConfigurationManager) UpdateStat(stat *StatDefinition) error
func (cm *ConfigurationManager) DeleteStat(id string) error
```

## 8. Runtime Configuration Updates

### Hot Reload System
```go
type HotReloadManager struct {
    configManager *ConfigurationManager
    watchers      map[string]*fsnotify.Watcher
    callbacks     []ConfigUpdateCallback
}

type ConfigUpdateCallback func(updateType string, id string, data interface{})

// Watch configuration files for changes
func (hrm *HotReloadManager) WatchFile(filename string) error

// Register callback for configuration updates
func (hrm *HotReloadManager) RegisterCallback(callback ConfigUpdateCallback)

// Handle configuration updates
func (hrm *HotReloadManager) HandleUpdate(updateType string, id string, data interface{})
```

## 9. Benefits

### 1. Complete Flexibility
- **No Hardcoded Stats**: All stats can be configured
- **No Hardcoded Formulas**: All formulas can be customized
- **No Hardcoded Types**: All types can be defined
- **No Hardcoded Clamps**: All bounds can be adjusted

### 2. Runtime Customization
- **Hot Reload**: Update configuration without restart
- **A/B Testing**: Test different configurations
- **Mod Support**: Easy modding through configuration
- **Balance Updates**: Quick balance adjustments

### 3. Validation & Safety
- **Configuration Validation**: Ensure configuration is valid
- **Dependency Checking**: Verify all dependencies exist
- **Range Validation**: Ensure values are within bounds
- **Formula Validation**: Verify formulas are correct

### 4. Performance
- **Cached Calculations**: Cache formula results
- **Lazy Loading**: Load configurations on demand
- **Incremental Updates**: Only update changed parts
- **Memory Efficient**: Efficient memory usage

## 10. Implementation Priority

### Phase 1: Core Configuration System
1. **ConfigurationManager** struct
2. **Stat/Formula/Type/Clamp definitions**
3. **Basic validation engine**

### Phase 2: Formula Engine
1. **Expression parser**
2. **Formula calculator**
3. **Dependency resolver**

### Phase 3: Hot Reload System
1. **File watcher**
2. **Update callbacks**
3. **Runtime updates**

### Phase 4: Advanced Features
1. **Category system**
2. **Validation rules**
3. **Performance optimization**

---

*Tài liệu này mô tả Flexible Configuration System cho Actor Core v2.0, làm cho toàn bộ hệ thống hoàn toàn flexible và có thể tùy chỉnh.*
