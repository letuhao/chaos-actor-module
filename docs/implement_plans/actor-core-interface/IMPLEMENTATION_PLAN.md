# Actor Core Interface v2.0 Implementation Plan

## Overview
This document outlines the implementation plan for the updated Actor Core Interface v2.0, which introduces flexible damage types, defence types, energy systems, and enhanced stat management.

## Phase 1: Core Type Updates

### 1.1 Update PrimaryCore
```go
type PrimaryCore struct {
    HPMax    int64
    LifeSpan int64
    Age      int64  // NEW: How long the actor has lived
    Stamina  int64  // NEW: Physical endurance
    Speed    int64
}
```

### 1.2 Update Derived
```go
type Derived struct {
    HPMax           float64
    Energies        map[string]float64                    // NEW: MP/Qi/Lust/Wrath/etc.
    Haste           float64
    CritChance      float64
    CritChanceResist float64                             // NEW: Resistance to critical hits
    CritMulti       float64
    CritMultiResist float64                              // NEW: Resistance to critical damage
    MoveSpeed       float64
    RegenHP         float64
    RegenEnergies   map[string]float64                    // NEW: Energy regeneration
    Damages         map[DamageType, float64]              // NEW: Damage by type
    Defences        map[DefenceType, DefenceValue]        // NEW: Defence against damage types
    Amplifiers      map[AmplifierType, AmplifierValue]    // NEW: Damage amplification
    Version         uint64
}
```

## Phase 2: Interface Definitions

### 2.1 DamageType Interface
```go
type DamageType interface {
    Name() string
    AffectedStats() []StatAffection
    IsElemental() bool
    IsPhysical() bool
    IsMagical() bool
    Priority() int // Higher priority = applied first
}

type StatAffection struct {
    StatKey    string
    Multiplier float64
    IsDrain    bool // true if drains the stat instead of damaging
}
```

### 2.2 DefenceType Interface
```go
type DefenceType interface {
    Name() string
    DamageTypes() []DamageType
    ResistCap() float64
    DrainCap() float64
    ReflectCap() float64
    Priority() int // Higher priority = applied first
}

type DefenceValue struct {
    Resist  float64 // reduces damage (0.0 to 1.0)
    Drain   float64 // converts damage to other resource (0.0 to 1.0)
    Reflect float64 // returns damage to attacker (0.0 to 1.0)
}
```

### 2.3 AmplifierType Interface
```go
type AmplifierType interface {
    Name() string
    DamageTypes() []DamageType
    IsMultiplicative() bool
    IsPiercing() bool
    Priority() int // Higher priority = applied first
}

type AmplifierValue struct {
    Multiplier float64 // damage multiplier
    Piercing   float64 // pierces through defences (0.0 to 1.0)
}
```

## Phase 3: Registry System

### 3.1 Type Registry
```go
type TypeRegistry struct {
    damageTypes   map[string]DamageType
    defenceTypes  map[string]DefenceType
    amplifierTypes map[string]AmplifierType
    mutex         sync.RWMutex
}

func (tr *TypeRegistry) RegisterDamageType(dt DamageType) error
func (tr *TypeRegistry) RegisterDefenceType(dt DefenceType) error
func (tr *TypeRegistry) RegisterAmplifierType(at AmplifierType) error
func (tr *TypeRegistry) GetDamageType(name string) (DamageType, bool)
func (tr *TypeRegistry) GetDefenceType(name string) (DefenceType, bool)
func (tr *TypeRegistry) GetAmplifierType(name string) (AmplifierType, bool)
```

## Phase 4: Predefined Types

### 4.1 Physical Damage Types
```go
type PhysicalDamage struct{}
type PiercingDamage struct{}
type SlashingDamage struct{}
type BluntDamage struct{}
```

### 4.2 Elemental Damage Types
```go
type FireDamage struct{}
type WaterDamage struct{}
type EarthDamage struct{}
type MetalDamage struct{}
type WoodDamage struct{}
type WindDamage struct{}
type LightningDamage struct{}
```

### 4.3 Magical Damage Types
```go
type LightDamage struct{}
type DarkDamage struct{}
type ArcaneDamage struct{}
```

### 4.4 Special Damage Types
```go
type TimeDamage struct{}
type CurseDamage struct{}
type NecrosisDamage struct{}
type BloodDamage struct{}
type PoisonDamage struct{}
type DiseaseDamage struct{}
```

## Phase 5: Damage Calculation Engine

### 5.1 Damage Calculator
```go
type DamageCalculator struct {
    registry *TypeRegistry
}

func (dc *DamageCalculator) CalculateDamage(
    attacker *Derived,
    defender *Derived,
    damageType DamageType,
    baseDamage float64,
) *DamageResult

type DamageResult struct {
    FinalDamage     float64
    DrainedResource map[string]float64
    ReflectedDamage float64
    IsCritical      bool
    CritMultiplier  float64
}
```

### 5.2 Critical Hit Calculator
```go
func (dc *DamageCalculator) CalculateCriticalHit(
    attacker *Derived,
    defender *Derived,
) (bool, float64)
```

## Phase 6: Migration Strategy

### 6.1 Backward Compatibility
- Keep old types in deprecated state
- Provide migration functions
- Gradual transition period

### 6.2 Data Migration
```go
func MigrateDerivedV1ToV2(old *DerivedV1) *DerivedV2
func MigratePrimaryCoreV1ToV2(old *PrimaryCoreV1) *PrimaryCoreV2
```

## Phase 7: Testing Strategy

### 7.1 Unit Tests
- Type registration tests
- Damage calculation tests
- Critical hit calculation tests
- Defence calculation tests
- Amplifier calculation tests

### 7.2 Integration Tests
- End-to-end damage flow
- Multiple damage type interactions
- Defence stacking tests
- Amplifier combination tests

### 7.3 Performance Tests
- Large number of damage types
- Complex damage calculations
- Memory usage optimization

## Phase 8: Documentation

### 8.1 API Documentation
- Complete API reference
- Usage examples
- Migration guide

### 8.2 Game Design Documentation
- Damage type guidelines
- Defence type guidelines
- Amplifier type guidelines
- Balance considerations

## Implementation Timeline

### Week 1-2: Core Type Updates
- Update PrimaryCore and Derived structs
- Implement basic interfaces
- Create type registry

### Week 3-4: Predefined Types
- Implement all predefined damage types
- Implement all predefined defence types
- Implement all predefined amplifier types

### Week 5-6: Calculation Engine
- Implement damage calculation engine
- Implement critical hit calculator
- Add comprehensive tests

### Week 7-8: Migration and Testing
- Implement migration functions
- Comprehensive testing
- Performance optimization

### Week 9-10: Documentation and Polish
- Complete documentation
- Final testing
- Release preparation

## Risk Mitigation

### 1. Performance Concerns
- Use efficient data structures
- Implement caching where appropriate
- Profile and optimize hot paths

### 2. Complexity Management
- Clear separation of concerns
- Comprehensive documentation
- Extensive testing

### 3. Backward Compatibility
- Gradual migration path
- Deprecation warnings
- Clear migration guide

## Success Criteria

1. All new types are properly implemented
2. Damage calculation is accurate and efficient
3. Type system is flexible and extensible
4. Backward compatibility is maintained
5. Performance meets requirements
6. Comprehensive test coverage
7. Complete documentation
