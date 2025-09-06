# RPG Stats Sub-System

A modular RPG Stats Sub-System that plugs into the Core Actor Interface. This system handles stat progression, equipment modifiers, buffs/debuffs, and emits deterministic stat snapshots for combat and UI.

## Features

- **8 Primary Stats**: STR, INT, WIL, AGI, SPD, END, PER, LUK (Morrowind-style)
- **Derived Stats**: HP, Mana, Attack, Defense, Critical Chance, Movement Speed, Resistances, etc.
- **Deterministic Stacking**: Predictable modifier application order
- **Extensible**: Easy to add new stats and modifier sources
- **Database Agnostic**: Core logic is independent of persistence layer
- **Comprehensive Testing**: Unit tests for all components

## Architecture

```
internal/
├── model/          # Core types and interfaces
├── rules/          # Stat stacking rules and modifiers
├── registry/       # Stat definitions and base curves
├── resolver/       # Stat resolution algorithm
├── integration/    # SnapshotProvider and ProgressionService
└── util/           # Utility functions

cmd/demo/           # Demo application
test/               # Unit tests
```

## Quick Start

### Running the Demo

```bash
go run cmd/demo/main.go
```

### Using the API

```go
package main

import (
    "fmt"
    "rpg-stats/internal/integration"
    "rpg-stats/internal/model"
)

func main() {
    // Create services
    provider := integration.NewSnapshotProvider()
    progression := integration.NewProgressionService()

    // Build a stat snapshot
    snapshot, err := provider.BuildForActor("player1", &model.SnapshotOptions{
        WithBreakdown: true,
    })
    if err != nil {
        panic(err)
    }

    // Display stats
    fmt.Printf("STR: %.1f\n", snapshot.Stats[model.STR].Value)
    fmt.Printf("HP: %.1f\n", snapshot.Stats[model.HP_MAX].Value)
    fmt.Printf("ATK: %.1f\n", snapshot.Stats[model.ATK].Value)

    // Grant XP and level up
    result, err := progression.GrantXP("player1", 1000)
    if err != nil {
        panic(err)
    }

    if result.NewLevel != nil {
        fmt.Printf("Level up! New level: %d\n", *result.NewLevel)
    }
}
```

## Stat System

### Primary Stats

- **STR** (Strength): Physical power and melee damage
- **INT** (Intelligence): Magical power and mana
- **WIL** (Willpower): Mental fortitude and spell resistance
- **AGI** (Agility): Speed and evasion
- **SPD** (Speed): Movement and action speed
- **END** (Endurance): Health and stamina
- **PER** (Personality): Social influence and merchant prices
- **LUK** (Luck): Critical hits and random events

### Derived Stats

Stats are calculated from primary stats using formulas:

- **HP_MAX** = Base + END × 20 + STR × 2
- **MANA_MAX** = Base + INT × 15 + WIL × 5
- **ATK** = STR × 2 + AGI × 0.3
- **MATK** = INT × 2 + WIL × 0.4
- **CRIT_CHANCE** = min(1.0, 0.01 + LUK × 0.003 + AGI × 0.001)
- And many more...

### Modifier Stacking

Modifiers are applied in deterministic order:

1. **Base** values from primary stats
2. **ADD_FLAT** modifiers (equipment, titles, passives)
3. **ADD_PCT** modifiers (percentage bonuses)
4. **MULTIPLY** modifiers (buffs, debuffs, auras, environment)
5. **OVERRIDE** modifiers (highest priority/value wins)
6. **CAPS** and soft DR functions
7. **Rounding** rules

## Testing

Run the test suite:

```bash
go test ./...
```

Run with verbose output:

```bash
go test -v ./...
```

## Integration with Core Actor

The RPG Stats Sub-System is designed to work with the Core Actor Interface:

1. Load player data and content
2. Aggregate modifiers by source
3. Resolve stats to create a snapshot
4. Feed the snapshot into the Core Actor
5. On changes, recompute and push again

## Extensibility

### Adding New Stats

1. Add the stat key to `model.StatKey`
2. Add the stat definition to the registry
3. Implement the calculation formula in `ResolveDerivedStat`

### Adding New Modifier Sources

1. Add the source kind to `ModifierSourceRef`
2. Update the resolver to handle the new source
3. Add database adapters if needed

## Database Design

The system is designed to work with MongoDB collections:

- `player_progress`: Character level, XP, stat allocations
- `player_effects_active`: Active buffs, debuffs, auras
- `player_equipment`: Equipped items and their modifiers
- `titles_owned`: Owned titles and their effects
- `content_stat_registry`: Stat definitions and formulas

## License

This project is part of the Chaos Actor Module system.
