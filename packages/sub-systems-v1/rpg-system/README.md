# RPG System

A comprehensive, modular RPG statistics and progression system designed for integration with the Chaos Actor Module.

## Overview

The RPG System provides a complete solution for managing character statistics, progression, equipment, effects, and Core Actor integration. It features a deterministic stat calculation engine, flexible modifier system, and seamless database integration.

## Features

### Core Statistics
- **8 Primary Stats**: STR, INT, WIL, AGI, SPD, END, PER, LUK
- **9 Derived Stats**: HP_MAX, MANA_MAX, ATK, MATK, DEF, EVASION, MOVE_SPEED, CRIT_CHANCE, CRIT_DAMAGE
- **Level Scaling**: Soft-capped progression curves for balanced growth
- **Stat Allocation**: Point-based character customization

### Modifier System
- **Deterministic Stacking**: ADD_FLAT → ADD_PCT → MULTIPLY → OVERRIDE → CAPS → Rounding
- **Priority System**: Configurable modifier priority for complex interactions
- **Stack Limits**: Per-stack modifier limits for balance
- **Breakdown Analysis**: Detailed stat calculation transparency

### Progression System
- **XP Management**: Flexible experience point system
- **Level Calculation**: Configurable level curves with soft caps
- **Stat Points**: Automatic stat point allocation per level
- **Respec System**: Character stat redistribution

### Integration
- **Core Actor Compatible**: Seamless integration with Core Actor system
- **Database Agnostic**: Mock and MongoDB adapters available
- **Caching Support**: Deterministic hashing for performance
- **Breakdown System**: Detailed stat calculation analysis

## Architecture

```
rpg-system/
├── cmd/                    # Command-line applications
│   ├── demo/              # Basic demonstration
│   ├── integration_example/ # Core Actor integration example
│   └── mongodb_example/   # MongoDB integration example
├── internal/              # Core system components
│   ├── model/            # Data structures and types
│   ├── registry/         # Stat definitions and formulas
│   ├── resolver/         # Stat calculation engine
│   ├── rules/            # Modifier stacking rules
│   ├── util/             # Utility functions
│   └── integration/      # External system integration
├── test/                 # Comprehensive test suite
├── db/                   # Database schemas and indexes
└── docs/                 # Documentation
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "rpg-system/internal/integration"
    "rpg-system/internal/model"
)

func main() {
    // Create integration with mock adapter
    adapter := integration.NewMockMongoAdapter()
    integration := integration.NewCoreActorIntegration(adapter)
    
    // Create a character
    ctx := context.Background()
    actorID := "player_001"
    
    // Build Core Actor contribution
    contribution, err := integration.BuildCoreActorContribution(ctx, actorID)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Character Stats: HP=%d, ATK=%d, DEF=%d\n",
        contribution.Primary.HPMax,
        contribution.Primary.Attack,
        contribution.Primary.Defense)
}
```

### Stat Calculation

```go
// Create compute input
input := model.ComputeInput{
    ActorID: "player_001",
    Level: 5,
    BaseAllocations: map[model.StatKey]int{
        model.STR: 15,
        model.INT: 12,
        model.END: 16,
    },
    Items: []model.StatModifier{
        {
            Key: model.STR,
            Op: model.ADD_FLAT,
            Value: 5.0,
            Source: model.ModifierSourceRef{
                Kind: "item",
                ID: "iron_sword",
                Label: "Iron Sword",
            },
        },
    },
    WithBreakdown: true,
}

// Compute stat snapshot
resolver := resolver.NewStatResolver()
snapshot := resolver.ComputeSnapshot(input)

// Access calculated stats
strValue := snapshot.Stats[model.STR]
hpMax := snapshot.Stats[model.HP_MAX]

// Access breakdown for transparency
if breakdown, exists := snapshot.Breakdown[model.STR]; exists {
    fmt.Printf("STR Breakdown: Base=%.2f, Flat=%.2f, Final=%.2f\n",
        breakdown.Base,
        breakdown.AdditiveFlat,
        breakdown.Final)
}
```

### Progression Management

```go
// Create progression service
progression := integration.NewProgressionService()

// Grant XP and handle leveling
result, err := progression.GrantXP("player_001", 1000)
if err != nil {
    log.Fatal(err)
}

// Allocate stat points
err = progression.AllocatePoints("player_001", model.STR, 3)
if err != nil {
    log.Fatal(err)
}

// Get player progress
progress, err := progression.GetPlayerProgress("player_001")
if err != nil {
    log.Fatal(err)
}
```

## Database Integration

### Mock Adapter (Development)

```go
adapter := integration.NewMockMongoAdapter()
```

### MongoDB Adapter (Production)

```go
// Build with MongoDB support
go build -tags mongodb ./cmd/mongodb_example

// Create MongoDB adapter
adapter, err := integration.NewMongoAdapterLocal("rpg_database")
if err != nil {
    log.Fatal(err)
}
defer adapter.Close()
```

## Testing

```bash
# Run all tests
go test ./test -v

# Run with MongoDB support
go test -tags mongodb ./test -v

# Run specific test
go test ./test -v -run TestStatResolver
```

## Performance

- **Deterministic Calculations**: Consistent results across runs
- **Efficient Caching**: Hash-based snapshot caching
- **Optimized Algorithms**: O(n) modifier processing
- **Memory Efficient**: Minimal allocations during calculations

## Configuration

### Stat Definitions

Stats are defined in `internal/registry/registry.go` with configurable:
- Base values and level scaling
- Min/max caps and soft caps
- Derived stat formulas
- Display names and descriptions

### Modifier Stacking

Stacking rules are defined in `internal/rules/stacking.go`:
1. ADD_FLAT modifiers
2. ADD_PCT modifiers  
3. MULTIPLY modifiers
4. OVERRIDE modifiers
5. CAPS application
6. Rounding

## Contributing

1. Follow Go best practices
2. Add comprehensive tests
3. Update documentation
4. Ensure all tests pass
5. Consider performance implications

## License

This project is part of the Chaos Actor Module and follows the same license terms.