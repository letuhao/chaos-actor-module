# Actor Core Interface v2.0 Design

## Overview
The Actor Core Interface v2.0 represents a major evolution of the actor system, introducing flexible damage types, energy systems, and enhanced stat management to support complex RPG mechanics.

## Key Features

### 1. Enhanced Primary Stats
- **Age**: Tracks how long the actor has lived
- **Stamina**: Physical endurance for actions
- **Removed**: Attack/Defense (replaced by flexible damage/defence system)

### 2. Flexible Energy System
- **Energies Map**: Support for MP, Qi, Lust, Wrath, and custom energy types
- **Per-Sub-system**: Each sub-system can define its own energy types
- **Regeneration**: Individual regeneration rates for each energy type

### 3. Damage Type System
- **Interface-based**: Extensible damage type system
- **Stat Affection**: Each damage type can affect multiple stats
- **Categories**: Physical, Elemental, Magical, and Special damage types
- **Priority System**: Damage types can have priority for processing order

### 4. Defence Type System
- **Three Layers**: Resist (reduce), Drain (convert), Reflect (return)
- **Damage Type Mapping**: Each defence type can protect against multiple damage types
- **Caps**: Configurable limits for each defence layer

### 5. Amplifier System
- **Multiplier**: Damage amplification
- **Piercing**: Ability to bypass defences
- **Type-specific**: Amplifiers can target specific damage types

### 6. Critical Hit System
- **Resistance**: Separate resistance to critical hits and critical damage
- **Balanced**: Prevents critical hit stacking issues

## Design Principles

### 1. Flexibility
- Easy to add new damage types, defence types, and amplifier types
- Runtime registration of new types
- Sub-system specific implementations

### 2. Performance
- Efficient lookups using maps
- Cached calculations where appropriate
- Object pooling for frequently allocated objects

### 3. Extensibility
- Interface-based design allows for custom implementations
- Plugin architecture for sub-systems
- Mod support through type registration

### 4. Backward Compatibility
- Migration functions for existing code
- Gradual transition support
- Deprecation warnings for old APIs

## File Structure

```
docs/implement_plans/actor-core-interface/
├── README.md                 # This file
├── SPEC.md                   # Complete specification
├── IMPLEMENTATION_PLAN.md    # Detailed implementation plan
├── EXAMPLES.md              # Usage examples
└── MIGRATION_GUIDE.md       # Migration from v1.0 to v2.0
```

## Quick Start

### 1. Basic Usage
```go
// Create character with new stats
character := &PrimaryCore{
    HPMax:    1000,
    LifeSpan: 100,
    Age:      25,
    Stamina:  500,
    Speed:    50,
}

// Calculate derived stats
derived := actorCore.BaseFromPrimary(character, 10)
```

### 2. Energy System
```go
// Set up energies
derived.Energies["MP"] = 100
derived.Energies["Qi"] = 80
derived.RegenEnergies["MP"] = 5
derived.RegenEnergies["Qi"] = 3
```

### 3. Damage System
```go
// Register damage types
registry.RegisterDamageType(FireDamage{})
registry.RegisterDefenceType(FireResistance{})
registry.RegisterAmplifierType(FireMastery{})

// Use in combat
damage := derived.Damages[FireDamage{}]
```

## Implementation Status

- [x] Design specification
- [x] Implementation plan
- [x] Examples and documentation
- [x] Migration guide
- [ ] Core implementation
- [ ] Type registry system
- [ ] Damage calculation engine
- [ ] Testing suite
- [ ] Performance optimization

## Next Steps

1. **Review Design**: Review the specification and provide feedback
2. **Implementation**: Begin implementing the core types and interfaces
3. **Testing**: Create comprehensive test suite
4. **Migration**: Implement migration functions
5. **Documentation**: Complete API documentation
6. **Performance**: Optimize for production use

## Contributing

When contributing to this design:

1. Follow the established patterns
2. Maintain backward compatibility
3. Add comprehensive tests
4. Update documentation
5. Consider performance implications

## Support

For questions or issues:

1. Check the examples in `EXAMPLES.md`
2. Review the migration guide in `MIGRATION_GUIDE.md`
3. Consult the implementation plan in `IMPLEMENTATION_PLAN.md`
4. Create an issue for bugs or feature requests

## Version History

- **v1.0**: Original design with basic stats
- **v2.0**: Current design with flexible damage types and energy system

## License

This design is part of the Chaos Actor Module project and follows the same license terms.