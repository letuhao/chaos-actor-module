# Actor Core Interface v2.0 Migration Guide

## Overview
This guide helps you migrate from Actor Core Interface v1.0 to v2.0. The new version introduces significant changes to support flexible damage types, energy systems, and enhanced stat management.

## Breaking Changes

### 1. PrimaryCore Changes

#### v1.0 (Old)
```go
type PrimaryCore struct {
    HPMax    int
    LifeSpan int
    Attack   int    // REMOVED
    Defense  int    // REMOVED
    Speed    int
}
```

#### v2.0 (New)
```go
type PrimaryCore struct {
    HPMax    int64
    LifeSpan int64
    Age      int64  // NEW: How long the actor has lived
    Stamina  int64  // NEW: Physical endurance
    Speed    int64
}
```

**Migration Steps:**
1. Update all `int` to `int64`
2. Remove `Attack` and `Defense` fields
3. Add `Age` field (initialize to 0 for new characters)
4. Add `Stamina` field (initialize based on character class/race)

### 2. Derived Changes

#### v1.0 (Old)
```go
type Derived struct {
    HPMax, MPMax           float64
    ATK, MAG               float64  // REMOVED
    DEF, RES               float64  // REMOVED
    Haste                  float64
    CritChance, CritMulti  float64
    MoveSpeed              float64
    RegenHP, RegenMP       float64
    Resists                map[string]float64
    Amplifiers             map[string]float64
    Version                uint64
}
```

#### v2.0 (New)
```go
type Derived struct {
    HPMax           float64
    Energies        map[string]float64                    // NEW: Replaces MPMax
    Haste           float64
    CritChance      float64
    CritChanceResist float64                             // NEW
    CritMulti       float64
    CritMultiResist float64                              // NEW
    MoveSpeed       float64
    RegenHP         float64
    RegenEnergies   map[string]float64                    // NEW: Replaces RegenMP
    Damages         map[DamageType, float64]              // NEW: Replaces ATK/MAG
    Defences        map[DefenceType, DefenceValue]        // NEW: Replaces DEF/RES
    Amplifiers      map[AmplifierType, AmplifierValue]    // NEW: Enhanced version
    Version         uint64
}
```

**Migration Steps:**
1. Remove `MPMax`, `ATK`, `MAG`, `DEF`, `RES`, `RegenMP`
2. Add `Energies` map for energy types (MP, Qi, etc.)
3. Add `CritChanceResist` and `CritMultiResist`
4. Add `RegenEnergies` map
5. Add `Damages` map for damage types
6. Add `Defences` map for defence types
7. Update `Amplifiers` to use new `AmplifierValue` type

## Migration Functions

### 1. PrimaryCore Migration

```go
func MigratePrimaryCoreV1ToV2(old *PrimaryCoreV1) *PrimaryCoreV2 {
    return &PrimaryCoreV2{
        HPMax:    int64(old.HPMax),
        LifeSpan: int64(old.LifeSpan),
        Age:      0, // Initialize to 0 for new characters
        Stamina:  int64(old.Speed * 2), // Estimate based on speed
        Speed:    int64(old.Speed),
    }
}
```

### 2. Derived Migration

```go
func MigrateDerivedV1ToV2(old *DerivedV1, registry *TypeRegistry) *DerivedV2 {
    // Initialize new derived
    new := &DerivedV2{
        HPMax:           old.HPMax,
        Energies:        make(map[string]float64),
        Haste:           old.Haste,
        CritChance:      old.CritChance,
        CritChanceResist: 0, // Default to 0
        CritMulti:       old.CritMulti,
        CritMultiResist: 0,  // Default to 0
        MoveSpeed:       old.MoveSpeed,
        RegenHP:         old.RegenHP,
        RegenEnergies:   make(map[string]float64),
        Damages:         make(map[DamageType, float64]),
        Defences:        make(map[DefenceType, DefenceValue]),
        Amplifiers:      make(map[AmplifierType, AmplifierValue]),
        Version:         old.Version,
    }
    
    // Migrate MPMax to Energies
    if old.MPMax > 0 {
        new.Energies["MP"] = old.MPMax
        new.RegenEnergies["MP"] = old.RegenMP
    }
    
    // Migrate ATK/MAG to Damages
    if old.ATK > 0 {
        if physicalType, exists := registry.GetDamageType("physical"); exists {
            new.Damages[physicalType] = old.ATK
        }
    }
    if old.MAG > 0 {
        if magicalType, exists := registry.GetDamageType("magical"); exists {
            new.Damages[magicalType] = old.MAG
        }
    }
    
    // Migrate DEF/RES to Defences
    if old.DEF > 0 {
        if physicalDef, exists := registry.GetDefenceType("physical"); exists {
            new.Defences[physicalDef] = DefenceValue{
                Resist:  old.DEF / 100.0, // Convert to percentage
                Drain:   0,
                Reflect: 0,
            }
        }
    }
    if old.RES > 0 {
        if magicalDef, exists := registry.GetDefenceType("magical"); exists {
            new.Defences[magicalDef] = DefenceValue{
                Resist:  old.RES / 100.0, // Convert to percentage
                Drain:   0,
                Reflect: 0,
            }
        }
    }
    
    // Migrate old Amplifiers
    for key, value := range old.Amplifiers {
        if ampType, exists := registry.GetAmplifierType(key); exists {
            new.Amplifiers[ampType] = AmplifierValue{
                Multiplier: value,
                Piercing:   0,
            }
        }
    }
    
    // Migrate old Resists
    for key, value := range old.Resists {
        if defType, exists := registry.GetDefenceType(key); exists {
            new.Defences[defType] = DefenceValue{
                Resist:  value,
                Drain:   0,
                Reflect: 0,
            }
        }
    }
    
    return new
}
```

## Step-by-Step Migration Process

### Step 1: Update Dependencies
```bash
# Update to new version
go get github.com/your-org/actor-core@v2.0.0
```

### Step 2: Update Type Definitions
```go
// Update your structs to use new types
type MyCharacter struct {
    Primary *PrimaryCoreV2  // Updated type
    Derived *DerivedV2      // Updated type
    // ... other fields
}
```

### Step 3: Initialize Type Registry
```go
// Set up type registry
registry := NewTypeRegistry()

// Register basic damage types
registry.RegisterDamageType(PhysicalDamage{})
registry.RegisterDamageType(MagicalDamage{})
registry.RegisterDamageType(FireDamage{})
// ... register other types

// Register basic defence types
registry.RegisterDefenceType(PhysicalDefence{})
registry.RegisterDefenceType(MagicalDefence{})
registry.RegisterDefenceType(FireResistance{})
// ... register other types

// Register basic amplifier types
registry.RegisterAmplifierType(PhysicalAmplifier{})
registry.RegisterAmplifierType(MagicalAmplifier{})
registry.RegisterAmplifierType(FireMastery{})
// ... register other types
```

### Step 4: Migrate Existing Data
```go
// Migrate existing characters
func MigrateCharacter(oldChar *CharacterV1) *CharacterV2 {
    return &CharacterV2{
        ID:       oldChar.ID,
        Primary:  MigratePrimaryCoreV1ToV2(oldChar.Primary),
        Derived:  MigrateDerivedV1ToV2(oldChar.Derived, registry),
        // ... other fields
    }
}
```

### Step 5: Update Combat Logic
```go
// Old combat logic
func OldCombat(attacker, defender *DerivedV1) {
    damage := attacker.ATK - defender.DEF
    if damage < 0 {
        damage = 0
    }
    defender.HPMax -= damage
}

// New combat logic
func NewCombat(attacker, defender *DerivedV2, damageType DamageType) {
    result := CalculateDamage(attacker, defender, damageType, baseDamage)
    defender.HPMax -= result.FinalDamage
    
    // Apply drained resources
    for resource, amount := range result.DrainedResource {
        defender.Energies[resource] += amount
    }
    
    // Apply reflected damage
    if result.ReflectedDamage > 0 {
        attacker.HPMax -= result.ReflectedDamage
    }
}
```

### Step 6: Update Energy System
```go
// Define energy types for your game
func InitializeEnergySystem(character *DerivedV2) {
    // Fantasy RPG
    character.Energies["MP"] = 100
    character.Energies["Qi"] = 80
    character.Energies["Rage"] = 0
    
    character.RegenEnergies["MP"] = 5
    character.RegenEnergies["Qi"] = 3
    character.RegenEnergies["Rage"] = 2
}
```

## Common Migration Patterns

### 1. Converting Old Stats to New System

```go
// Convert old ATK to physical damage
if old.ATK > 0 {
    physicalType, _ := registry.GetDamageType("physical")
    new.Damages[physicalType] = float64(old.ATK)
}

// Convert old MAG to magical damage
if old.MAG > 0 {
    magicalType, _ := registry.GetDamageType("magical")
    new.Damages[magicalType] = float64(old.MAG)
}
```

### 2. Converting Old Defences

```go
// Convert old DEF to physical defence
if old.DEF > 0 {
    physicalDef, _ := registry.GetDefenceType("physical")
    new.Defences[physicalDef] = DefenceValue{
        Resist:  float64(old.DEF) / 100.0,
        Drain:   0,
        Reflect: 0,
    }
}
```

### 3. Converting Old Amplifiers

```go
// Convert old amplifiers
for key, value := range old.Amplifiers {
    if ampType, exists := registry.GetAmplifierType(key); exists {
        new.Amplifiers[ampType] = AmplifierValue{
            Multiplier: value,
            Piercing:   0,
        }
    }
}
```

## Testing Migration

### 1. Unit Tests
```go
func TestMigration(t *testing.T) {
    // Test PrimaryCore migration
    oldPrimary := &PrimaryCoreV1{HPMax: 100, LifeSpan: 50, Speed: 25}
    newPrimary := MigratePrimaryCoreV1ToV2(oldPrimary)
    
    assert.Equal(t, int64(100), newPrimary.HPMax)
    assert.Equal(t, int64(50), newPrimary.LifeSpan)
    assert.Equal(t, int64(0), newPrimary.Age)
    assert.Equal(t, int64(50), newPrimary.Stamina)
    assert.Equal(t, int64(25), newPrimary.Speed)
}
```

### 2. Integration Tests
```go
func TestCombatMigration(t *testing.T) {
    // Test that old combat logic produces similar results to new
    // (within acceptable tolerance)
}
```

## Rollback Plan

If you need to rollback:

1. Keep old types in deprecated state
2. Provide rollback migration functions
3. Maintain backward compatibility for one version
4. Document rollback procedure

```go
func RollbackDerivedV2ToV1(new *DerivedV2) *DerivedV1 {
    // Convert back to old format
    // This should be used only for emergency rollback
}
```

## Performance Considerations

1. **Memory Usage**: New system uses more memory due to maps
2. **Lookup Performance**: Use efficient data structures
3. **Cache Frequently Used Types**: Cache type lookups
4. **Pool Objects**: Use object pooling for frequently allocated objects

## Support

For migration support:
- Check the examples in `EXAMPLES.md`
- Review the implementation plan in `IMPLEMENTATION_PLAN.md`
- Test thoroughly before deploying
- Consider gradual migration for large systems
