# Actor Core Interface v2.0 Examples

## Basic Usage Examples

### 1. Creating a Character with New Stats

```go
// Create a character with the new PrimaryCore
character := &PrimaryCore{
    HPMax:    1000,
    LifeSpan: 100,  // 100 years to live
    Age:      25,   // Currently 25 years old
    Stamina:  500,  // Physical endurance
    Speed:    50,   // Movement speed
}

// Calculate derived stats
derived := actorCore.BaseFromPrimary(character, 10) // Level 10

// Access new energy system
fmt.Printf("MP: %.2f\n", derived.Energies["MP"])
fmt.Printf("Qi: %.2f\n", derived.Energies["Qi"])
fmt.Printf("Stamina: %.2f\n", derived.Stamina)
```

### 2. Damage Type System

```go
// Define custom damage types
type FireDamage struct{}

func (fd FireDamage) Name() string { return "fire" }
func (fd FireDamage) AffectedStats() []StatAffection {
    return []StatAffection{
        {StatKey: "HP", Multiplier: 1.0, IsDrain: false},
        {StatKey: "Stamina", Multiplier: 0.5, IsDrain: false},
    }
}
func (fd FireDamage) IsElemental() bool { return true }
func (fd FireDamage) IsPhysical() bool { return false }
func (fd FireDamage) IsMagical() bool { return false }
func (fd FireDamage) Priority() int { return 100 }

// Register damage type
registry.RegisterDamageType(FireDamage{})

// Use in combat
damage := derived.Damages[FireDamage{}] // Get fire damage
```

### 3. Defence System

```go
// Define defence types
type FireResistance struct{}

func (fr FireResistance) Name() string { return "fire_resist" }
func (fr FireResistance) DamageTypes() []DamageType {
    return []DamageType{FireDamage{}, LightningDamage{}}
}
func (fr FireResistance) ResistCap() float64 { return 0.8 }
func (fr FireResistance) DrainCap() float64 { return 0.3 }
func (fr FireResistance) ReflectCap() float64 { return 0.1 }
func (fr FireResistance) Priority() int { return 200 }

// Set up defences
derived.Defences[FireResistance{}] = DefenceValue{
    Resist:  0.5,  // 50% fire damage reduction
    Drain:   0.1,  // 10% of fire damage converted to MP
    Reflect: 0.05, // 5% of fire damage reflected back
}
```

### 4. Amplifier System

```go
// Define amplifier types
type FireMastery struct{}

func (fm FireMastery) Name() string { return "fire_mastery" }
func (fm FireMastery) DamageTypes() []DamageType {
    return []DamageType{FireDamage{}, LightningDamage{}}
}
func (fm FireMastery) IsMultiplicative() bool { return true }
func (fm FireMastery) IsPiercing() bool { return false }
func (fm FireMastery) Priority() int { return 150 }

// Set up amplifiers
derived.Amplifiers[FireMastery{}] = AmplifierValue{
    Multiplier: 1.5, // 50% more fire damage
    Piercing:   0.2, // 20% pierces through fire resistance
}
```

## Combat Examples

### 1. Basic Damage Calculation

```go
func CalculateDamage(attacker, defender *Derived, damageType DamageType, baseDamage float64) *DamageResult {
    // Get attacker's damage multiplier
    damage := baseDamage
    if amp, exists := attacker.Amplifiers[damageType]; exists {
        damage *= amp.Multiplier
    }
    
    // Apply defender's resistance
    if def, exists := defender.Defences[damageType]; exists {
        // Calculate effective resistance (considering piercing)
        effectiveResist := def.Resist * (1 - getPiercing(attacker, damageType))
        damage *= (1 - effectiveResist)
        
        // Calculate drained resource
        drained := damage * def.Drain
        
        // Calculate reflected damage
        reflected := damage * def.Reflect
        
        return &DamageResult{
            FinalDamage:     damage,
            DrainedResource: map[string]float64{"MP": drained},
            ReflectedDamage: reflected,
        }
    }
    
    return &DamageResult{FinalDamage: damage}
}
```

### 2. Critical Hit Calculation

```go
func CalculateCriticalHit(attacker, defender *Derived) (bool, float64) {
    // Calculate effective crit chance
    critChance := attacker.CritChance - defender.CritChanceResist
    if critChance < 0 {
        critChance = 0
    }
    
    // Roll for critical hit
    isCrit := rand.Float64() < critChance
    
    if isCrit {
        // Calculate effective crit multiplier
        critMulti := attacker.CritMulti * (1 - defender.CritMultiResist)
        if critMulti < 1.0 {
            critMulti = 1.0
        }
        return true, critMulti
    }
    
    return false, 1.0
}
```

### 3. Complex Damage Interaction

```go
func ComplexDamageCalculation(attacker, defender *Derived, damageType DamageType, baseDamage float64) *DamageResult {
    result := &DamageResult{
        DrainedResource: make(map[string]float64),
    }
    
    // Apply all amplifiers for this damage type
    for ampType, ampValue := range attacker.Amplifiers {
        if contains(ampType.DamageTypes(), damageType) {
            baseDamage *= ampValue.Multiplier
        }
    }
    
    // Apply all defences for this damage type
    for defType, defValue := range defender.Defences {
        if contains(defType.DamageTypes(), damageType) {
            // Calculate effective resistance
            piercing := getTotalPiercing(attacker, damageType)
            effectiveResist := defValue.Resist * (1 - piercing)
            
            // Apply resistance
            baseDamage *= (1 - effectiveResist)
            
            // Calculate drained resources
            for _, statAffection := range damageType.AffectedStats() {
                if statAffection.IsDrain {
                    drained := baseDamage * defValue.Drain * statAffection.Multiplier
                    result.DrainedResource[statAffection.StatKey] += drained
                }
            }
            
            // Calculate reflected damage
            result.ReflectedDamage += baseDamage * defValue.Reflect
        }
    }
    
    result.FinalDamage = baseDamage
    return result
}
```

## RPG System Integration Examples

### 1. Energy System Integration

```go
// Define energy types for different RPG systems
type EnergySystem struct {
    Name        string
    BaseValue   float64
    RegenRate   float64
    MaxValue    float64
}

// Fantasy RPG energies
fantasyEnergies := map[string]EnergySystem{
    "MP":    {Name: "Mana", BaseValue: 100, RegenRate: 5, MaxValue: 500},
    "Qi":    {Name: "Qi", BaseValue: 80, RegenRate: 3, MaxValue: 400},
    "Rage":  {Name: "Rage", BaseValue: 0, RegenRate: 2, MaxValue: 100},
}

// Sci-fi RPG energies
scifiEnergies := map[string]EnergySystem{
    "Energy": {Name: "Energy", BaseValue: 200, RegenRate: 10, MaxValue: 1000},
    "Shield": {Name: "Shield", BaseValue: 150, RegenRate: 8, MaxValue: 750},
    "Heat":   {Name: "Heat", BaseValue: 0, RegenRate: 5, MaxValue: 200},
}
```

### 2. Sub-system Specific Damage Types

```go
// Fantasy RPG damage types
type FantasyDamageTypes struct{}

func (fdt FantasyDamageTypes) RegisterAll(registry *TypeRegistry) {
    registry.RegisterDamageType(FireDamage{})
    registry.RegisterDamageType(IceDamage{})
    registry.RegisterDamageType(LightningDamage{})
    registry.RegisterDamageType(DarkDamage{})
    registry.RegisterDamageType(HolyDamage{})
    registry.RegisterDamageType(PoisonDamage{})
    registry.RegisterDamageType(CurseDamage{})
}

// Sci-fi RPG damage types
type SciFiDamageTypes struct{}

func (sdt SciFiDamageTypes) RegisterAll(registry *TypeRegistry) {
    registry.RegisterDamageType(PlasmaDamage{})
    registry.RegisterDamageType(LaserDamage{})
    registry.RegisterDamageType(KineticDamage{})
    registry.RegisterDamageType(EMPDamage{})
    registry.RegisterDamageType(RadiationDamage{})
    registry.RegisterDamageType(GraviticDamage{})
}
```

### 3. Dynamic Type Registration

```go
// Register types at runtime based on game configuration
func InitializeDamageTypes(config GameConfig) {
    registry := GetTypeRegistry()
    
    // Register based on game mode
    switch config.GameMode {
    case "fantasy":
        FantasyDamageTypes{}.RegisterAll(registry)
    case "scifi":
        SciFiDamageTypes{}.RegisterAll(registry)
    case "mixed":
        FantasyDamageTypes{}.RegisterAll(registry)
        SciFiDamageTypes{}.RegisterAll(registry)
    }
    
    // Register custom types from mods
    for _, mod := range config.Mods {
        mod.RegisterDamageTypes(registry)
    }
}
```

## Performance Considerations

### 1. Efficient Type Lookups

```go
// Use maps for O(1) lookups
type TypeRegistry struct {
    damageTypes   map[string]DamageType
    defenceTypes  map[string]DefenceType
    amplifierTypes map[string]AmplifierType
}

// Cache frequently used combinations
type DamageCalculationCache struct {
    cache map[string]*DamageResult
    mutex sync.RWMutex
}
```

### 2. Memory Optimization

```go
// Use pointers for large structs
type Derived struct {
    Energies      map[string]float64
    Damages       map[DamageType, float64]
    Defences      map[DefenceType, DefenceValue]
    Amplifiers    map[AmplifierType, AmplifierValue]
    // ... other fields
}

// Pool frequently allocated objects
var damageResultPool = sync.Pool{
    New: func() interface{} {
        return &DamageResult{
            DrainedResource: make(map[string]float64),
        }
    },
}
```

## Testing Examples

### 1. Unit Tests

```go
func TestDamageCalculation(t *testing.T) {
    registry := NewTypeRegistry()
    registry.RegisterDamageType(FireDamage{})
    registry.RegisterDefenceType(FireResistance{})
    
    attacker := &Derived{
        Damages: map[DamageType, float64]{
            FireDamage{}: 100,
        },
    }
    
    defender := &Derived{
        Defences: map[DefenceType, DefenceValue]{
            FireResistance{}: {Resist: 0.5},
        },
    }
    
    result := CalculateDamage(attacker, defender, FireDamage{}, 100)
    
    assert.Equal(t, 50.0, result.FinalDamage) // 50% resistance
}
```

### 2. Integration Tests

```go
func TestComplexCombatScenario(t *testing.T) {
    // Set up complex combat scenario
    // Test multiple damage types
    // Test multiple defence types
    // Test amplifier interactions
    // Verify all calculations are correct
}
```

This new system provides the flexibility and extensibility needed for complex RPG systems while maintaining performance and ease of use.
