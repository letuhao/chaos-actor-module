# RPG System Configuration Guide

## Overview

The RPG System is highly configurable to support different game mechanics and balance requirements. This guide covers all configuration options and how to customize the system for your specific needs.

## Stat Definitions

### Primary Stats

Primary stats are the core attributes that define a character's base capabilities.

```go
// Example: Strength stat definition
{
    Key:          model.STR,
    Category:     "primary",
    DisplayName:  "Strength",
    Description:  "Physical power and melee damage",
    IsPrimary:    true,
    MinValue:     1,
    MaxValue:     100,
    DefaultValue: 10,
}
```

**Configuration Options:**
- `Key`: Unique identifier for the stat
- `Category`: Stat category (primary, derived, special)
- `DisplayName`: Human-readable name
- `Description`: Tooltip or help text
- `IsPrimary`: Whether this is a primary stat
- `MinValue`: Minimum allowed value
- `MaxValue`: Maximum allowed value
- `DefaultValue`: Starting value for new characters

### Derived Stats

Derived stats are calculated from primary stats using formulas.

```go
// Example: Health Points derived from END and STR
{
    Key:          model.HP_MAX,
    Category:     "derived",
    DisplayName:  "Health Points",
    Description:  "Maximum health points",
    IsPrimary:    false,
    MinValue:     1,
    MaxValue:     10000,
    DefaultValue: 100,
}
```

## Level Curves

Level curves define how stats scale with character level.

### Basic Level Curve

```go
{
    BaseValue:    10.0,    // Starting value at level 1
    PerLevel:     2.0,     // Points gained per level
    MaxLevel:     100,     // Maximum level
    SoftCapLevel: 50,      // Level where soft cap begins
    SoftCapValue: 0.5,     // Multiplier after soft cap
}
```

### Soft Cap Calculation

The soft cap reduces stat gains after a certain level:

```go
if level <= softCapLevel {
    value = baseValue + (level - 1) * perLevel
} else {
    baseGain = baseValue + (softCapLevel - 1) * perLevel
    softCapGain = (level - softCapLevel) * perLevel * softCapValue
    value = baseGain + softCapGain
}
```

### Example Configurations

**Linear Growth (No Soft Cap):**
```go
{
    BaseValue:    10.0,
    PerLevel:     2.0,
    MaxLevel:     100,
    SoftCapLevel: 100,  // No soft cap
    SoftCapValue: 1.0,
}
```

**Exponential Decay:**
```go
{
    BaseValue:    10.0,
    PerLevel:     3.0,
    MaxLevel:     100,
    SoftCapLevel: 25,   // Early soft cap
    SoftCapValue: 0.3,  // Strong reduction
}
```

## Derived Formulas

Derived stats use mathematical formulas to calculate values from primary stats.

### Formula Syntax

Formulas support basic mathematical operations:
- `+` Addition
- `-` Subtraction
- `*` Multiplication
- `/` Division
- `()` Parentheses for grouping
- Stat names as variables

### Example Formulas

**Health Points:**
```go
{
    StatKey:    model.HP_MAX,
    Formula:    "END * 20 + STR * 5",
    Components: []model.StatKey{model.END, model.STR},
    BaseValue:  100.0,
}
```

**Attack Power:**
```go
{
    StatKey:    model.ATK,
    Formula:    "STR * 2 + AGI * 0.5",
    Components: []model.StatKey{model.STR, model.AGI},
    BaseValue:  10.0,
}
```

**Critical Chance:**
```go
{
    StatKey:    model.CRIT_CHANCE,
    Formula:    "LUK * 0.01 + AGI * 0.005",
    Components: []model.StatKey{model.LUK, model.AGI},
    BaseValue:  0.05,  // 5% base crit chance
}
```

## Modifier Stacking

### Stacking Rules

Modifiers are applied in a specific order:

1. **ADD_FLAT**: Add flat values
2. **ADD_PCT**: Add percentage bonuses
3. **MULTIPLY**: Multiply by values
4. **OVERRIDE**: Override with specific values
5. **CAPS**: Apply min/max caps
6. **ROUNDING**: Round to appropriate precision

### Priority System

Modifiers can have priorities to control application order:

```go
type StatModifier struct {
    Key         StatKey
    Op          ModifierOp
    Value       float64
    Source      ModifierSourceRef
    Priority    int  // Higher priority = applied first
    Stack       int  // Stack group for limits
    Conditions  *ModifierConditions
}
```

### Stack Limits

Limit the number of modifiers in a stack:

```go
// Example: Only 3 equipment modifiers can affect STR
stackLimits := map[int]int{
    1: 3,  // Stack 1 (equipment) limited to 3 modifiers
    2: 1,  // Stack 2 (titles) limited to 1 modifier
}
```

## Progression Configuration

### XP Requirements

Configure how much XP is needed for each level:

```go
func calculateXPForLevel(level int) int64 {
    if level <= 1 {
        return 0
    }
    
    // Linear progression
    return int64((level - 1) * 100)
    
    // Or exponential progression
    // return int64(math.Pow(1.5, float64(level-1)) * 100)
}
```

### Stat Points Per Level

Configure how many stat points characters get per level:

```go
func getStatPointsForLevel(level int) int {
    // 2 points per level
    return level * 2
    
    // Or diminishing returns
    // return int(math.Sqrt(float64(level)) * 2)
}
```

## Database Configuration

### MongoDB Collections

Configure collection names and indexes:

```go
// Collection names
playerProgressCol     = "player_progress"
playerEffectsCol      = "player_effects_active"
playerEquipmentCol    = "player_equipment"
titlesOwnedCol        = "titles_owned"
contentRegistryCol    = "content_stat_registry"

// Indexes for performance
indexes := []mongo.IndexModel{
    {
        Keys: bson.D{{Key: "actor_id", Value: 1}},
        Options: options.Index().SetUnique(true),
    },
    {
        Keys: bson.D{{Key: "expires_at", Value: 1}},
    },
}
```

### Connection Settings

```go
// MongoDB connection options
clientOptions := options.Client().
    ApplyURI("mongodb://localhost:27017").
    SetConnectTimeout(10 * time.Second).
    SetSocketTimeout(30 * time.Second).
    SetMaxPoolSize(100)
```

## Performance Tuning

### Caching Configuration

```go
// Enable snapshot caching
input := model.ComputeInput{
    // ... other fields
    WithBreakdown: false,  // Disable for better performance
}

// Use deterministic hashing
snapshot := resolver.ComputeSnapshot(input)
if snapshot.Hash != "" {
    // Cache the result using the hash
}
```

### Memory Optimization

```go
// Reuse objects to reduce allocations
var input model.ComputeInput
var snapshot *model.StatSnapshot

// Process multiple characters
for _, actorID := range actorIDs {
    input.ActorID = actorID
    snapshot = resolver.ComputeSnapshot(input)
    // Process snapshot...
}
```

## Custom Stat Types

### Adding New Primary Stats

1. Add the stat key to the enum:
```go
const (
    // ... existing stats
    VITALITY StatKey = "VITALITY"
)
```

2. Add the stat definition:
```go
{
    Key:          model.VITALITY,
    Category:     "primary",
    DisplayName:  "Vitality",
    Description:  "Life force and regeneration",
    IsPrimary:    true,
    MinValue:     1,
    MaxValue:     100,
    DefaultValue: 10,
}
```

3. Add level curve:
```go
{
    BaseValue:    10.0,
    PerLevel:     2.0,
    MaxLevel:     100,
    SoftCapLevel: 50,
    SoftCapValue: 0.5,
}
```

### Adding New Derived Stats

1. Add the stat key:
```go
const (
    // ... existing stats
    REGEN_RATE StatKey = "REGEN_RATE"
)
```

2. Add the stat definition:
```go
{
    Key:          model.REGEN_RATE,
    Category:     "derived",
    DisplayName:  "Regeneration Rate",
    Description:  "Health points regenerated per second",
    IsPrimary:    false,
    MinValue:     0,
    MaxValue:     100,
    DefaultValue: 1.0,
}
```

3. Add the formula:
```go
{
    StatKey:    model.REGEN_RATE,
    Formula:    "VITALITY * 0.1 + END * 0.05",
    Components: []model.StatKey{model.VITALITY, model.END},
    BaseValue:  1.0,
}
```

## Balance Testing

### Automated Testing

```go
func TestStatBalance(t *testing.T) {
    // Test stat progression
    for level := 1; level <= 100; level++ {
        value := registry.CalculateLevelValue(model.STR, level)
        
        // Ensure reasonable progression
        if value < 10 || value > 200 {
            t.Errorf("STR value %f at level %d is out of range", value, level)
        }
    }
}
```

### Performance Testing

```go
func BenchmarkStatCalculation(b *testing.B) {
    resolver := resolver.NewStatResolver()
    input := createTestInput()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        resolver.ComputeSnapshot(input)
    }
}
```

## Migration Guide

### Updating Stat Definitions

1. **Additive Changes**: New stats can be added safely
2. **Modifying Formulas**: Test thoroughly as this affects all characters
3. **Changing Level Curves**: Consider migration for existing characters

### Database Migration

```go
// Example: Add new stat to existing characters
func migrateAddVitality() error {
    // Update all existing player progress records
    // Add VITALITY allocation with default value
}
```

## Best Practices

1. **Start Simple**: Begin with basic stats and add complexity gradually
2. **Test Thoroughly**: Use automated tests for balance validation
3. **Document Changes**: Keep track of configuration changes
4. **Version Control**: Store configurations in version control
5. **Performance Monitor**: Monitor performance impact of changes
6. **Backup Data**: Always backup before major configuration changes
