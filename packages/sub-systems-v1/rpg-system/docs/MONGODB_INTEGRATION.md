# MongoDB Integration for RPG Stats Sub-System

This document describes how to use the RPG Stats Sub-System with MongoDB for persistent storage.

## Overview

The RPG Stats Sub-System now supports MongoDB integration for storing player progress, effects, equipment, titles, and stat registry data. The integration is designed to work with MongoDB running on `localhost:27017` by default.

## Features

- **Player Progress Storage**: Store character levels, XP, and stat allocations
- **Effects Management**: Track active buffs, debuffs, and temporary effects with expiration
- **Equipment System**: Manage equipped items with slot-based organization
- **Titles System**: Store and manage character titles and their stat bonuses
- **Stat Registry**: Persist stat definitions and formulas
- **Core Actor Integration**: Seamless integration with the Core Actor system

## Setup

### Prerequisites

1. **MongoDB Server**: Install and run MongoDB on your system
2. **Go Dependencies**: The MongoDB driver is included in `go.mod`

### Database Setup

1. **Start MongoDB**:
   ```bash
   # On Windows
   mongod --dbpath C:\data\db
   
   # On Linux/Mac
   sudo systemctl start mongod
   ```

2. **Create Database and Collections**:
   ```bash
   # Connect to MongoDB
   mongosh
   
   # Switch to RPG Stats database
   use rpg_stats
   
   # Run the schema creation script
   load("db/schemas.js")
   
   # Create indexes for performance
   load("db/indexes.js")
   ```

### Go Build Tags

The MongoDB integration uses build tags to conditionally compile MongoDB-specific code:

```bash
# Build with MongoDB support
go build -tags mongodb ./cmd/mongodb_example

# Run tests with MongoDB support
go test -tags mongodb ./test

# Build without MongoDB (uses mock adapter)
go build ./cmd/demo
```

## Usage

### Basic Setup

```go
package main

import (
    "context"
    "log"
    
    "rpg-stats/internal/integration"
)

func main() {
    // Connect to MongoDB
    adapter, err := integration.NewMongoAdapterLocal("rpg_stats_demo")
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    defer adapter.Close()
    
    // Create integration
    integration := integration.NewCoreActorIntegration(adapter)
    ctx := context.Background()
    
    // Use the integration...
}
```

### Player Progress

```go
// Create a new character
progress := &model.PlayerProgress{
    ActorID: "player_001",
    Level:   1,
    XP:      0,
    Allocations: map[model.StatKey]int{
        model.STR: 15,
        model.INT: 12,
        model.END: 16,
    },
    LastUpdated: time.Now().Unix(),
}

err := adapter.SavePlayerProgress(ctx, progress)
if err != nil {
    log.Fatalf("Failed to save progress: %v", err)
}

// Retrieve player progress
retrieved, err := adapter.GetPlayerProgress(ctx, "player_001")
if err != nil {
    log.Fatalf("Failed to retrieve progress: %v", err)
}
```

### Effects Management

```go
// Add a temporary effect
effect := model.StatModifier{
    Key:   model.STR,
    Op:    model.ADD_FLAT,
    Value: 5.0,
    Source: model.ModifierSourceRef{
        Kind:  "buff",
        ID:    "strength_potion",
        Label: "Strength Potion",
    },
    Priority: 1,
    Conditions: &model.ModifierConditions{
        DurationMs: 300000, // 5 minutes
    },
}

err := adapter.AddEffect(ctx, "player_001", effect, 5*time.Minute)
if err != nil {
    log.Fatalf("Failed to add effect: %v", err)
}

// Get active effects
effects, err := adapter.GetActiveEffects(ctx, "player_001")
if err != nil {
    log.Fatalf("Failed to get effects: %v", err)
}

// Remove effect
err = adapter.RemoveEffect(ctx, "player_001", "strength_potion")
if err != nil {
    log.Fatalf("Failed to remove effect: %v", err)
}
```

### Equipment System

```go
// Equip an item
item := model.StatModifier{
    Key:   model.ATK,
    Op:    model.ADD_FLAT,
    Value: 10.0,
    Source: model.ModifierSourceRef{
        Kind:  "item",
        ID:    "iron_sword",
        Label: "Iron Sword",
    },
    Priority: 1,
}

err := adapter.EquipItem(ctx, "player_001", "weapon", item)
if err != nil {
    log.Fatalf("Failed to equip item: %v", err)
}

// Get equipped items
equipment, err := adapter.GetEquippedItems(ctx, "player_001")
if err != nil {
    log.Fatalf("Failed to get equipment: %v", err)
}

// Unequip item
err = adapter.UnequipItem(ctx, "player_001", "weapon")
if err != nil {
    log.Fatalf("Failed to unequip item: %v", err)
}
```

### Titles System

```go
// Grant a title
title := model.StatModifier{
    Key:   model.PER,
    Op:    model.ADD_FLAT,
    Value: 2.0,
    Source: model.ModifierSourceRef{
        Kind:  "title",
        ID:    "noble_warrior",
        Label: "Noble Warrior",
    },
    Priority: 1,
}

err := adapter.GrantTitle(ctx, "player_001", "noble_warrior", title)
if err != nil {
    log.Fatalf("Failed to grant title: %v", err)
}

// Get owned titles
titles, err := adapter.GetOwnedTitles(ctx, "player_001")
if err != nil {
    log.Fatalf("Failed to get titles: %v", err)
}
```

### Core Actor Integration

```go
// Build Core Actor contribution
contribution, err := integration.BuildCoreActorContribution(ctx, "player_001")
if err != nil {
    log.Fatalf("Failed to build contribution: %v", err)
}

// Use the contribution with Core Actor
fmt.Printf("Primary Stats: HP=%d, Attack=%d, Defense=%d\n", 
    contribution.Primary.HPMax, 
    contribution.Primary.Attack, 
    contribution.Primary.Defense)

fmt.Printf("Flat Modifiers: ATK=%.2f, DEF=%.2f\n", 
    contribution.Flat["ATK"], 
    contribution.Flat["DEF"])

fmt.Printf("Tags: %v\n", contribution.Tags)
```

### Stat Registry

```go
// Save stat definitions
registry := []model.StatDef{
    {
        Key:          model.STR,
        Category:     "primary",
        DisplayName:  "Strength",
        Description:  "Physical power",
        IsPrimary:    true,
        MinValue:     1,
        MaxValue:     100,
        DefaultValue: 10,
    },
    // ... more stat definitions
}

err := adapter.SaveStatRegistry(ctx, registry)
if err != nil {
    log.Fatalf("Failed to save registry: %v", err)
}

// Retrieve stat registry
retrieved, err := adapter.GetStatRegistry(ctx)
if err != nil {
    log.Fatalf("Failed to get registry: %v", err)
}
```

## Database Schema

### Collections

1. **player_progress**: Character progression data
2. **player_effects_active**: Active effects with expiration
3. **player_equipment**: Equipped items by slot
4. **titles_owned**: Character titles and bonuses
5. **content_stat_registry**: Stat definitions and formulas

### Indexes

The system automatically creates optimized indexes for:
- Player lookups by actor_id
- Effect expiration cleanup
- Equipment slot management
- Title ownership tracking
- Stat registry queries

## Examples

### Complete Example

See `cmd/mongodb_example/main.go` for a comprehensive example that demonstrates:
- Character creation
- Level progression
- Equipment management
- Effect application
- Title granting
- Stat breakdown analysis

### Running Examples

```bash
# Run MongoDB example (requires MongoDB running)
go run -tags mongodb cmd/mongodb_example/main.go

# Run basic demo (uses mock adapter)
go run cmd/demo/main.go
```

## Testing

### Unit Tests

```bash
# Run all tests (uses mock adapter)
go test ./test -v

# Run MongoDB integration tests (requires MongoDB)
go test -tags mongodb ./test -v
```

### Integration Tests

The system includes comprehensive integration tests that verify:
- Database operations
- Core Actor integration
- Stat calculation consistency
- Effect management
- Equipment system

## Performance Considerations

### Indexing

The system creates optimized indexes for common query patterns:
- Player lookups by actor_id
- Effect expiration queries
- Equipment slot management
- Title ownership

### Caching

The RPG Stats system includes built-in caching through:
- Stat snapshot hashing
- Deterministic calculation results
- Efficient modifier processing

### Cleanup

Use the cleanup methods to maintain database performance:

```go
// Clean up expired effects
err := adapter.CleanupExpiredEffects(ctx)
if err != nil {
    log.Printf("Failed to cleanup effects: %v", err)
}
```

## Error Handling

The MongoDB integration includes comprehensive error handling:

```go
// Check for specific errors
if err != nil {
    if strings.Contains(err.Error(), "player not found") {
        // Handle missing player
    } else if strings.Contains(err.Error(), "connection") {
        // Handle connection issues
    } else {
        // Handle other errors
    }
}
```

## Security Considerations

1. **Connection Security**: Use MongoDB authentication and TLS in production
2. **Input Validation**: Validate all input data before database operations
3. **Access Control**: Implement proper user permissions for database access
4. **Data Sanitization**: Sanitize user input to prevent injection attacks

## Troubleshooting

### Common Issues

1. **Connection Failed**: Ensure MongoDB is running on localhost:27017
2. **Build Errors**: Use `-tags mongodb` for MongoDB-specific builds
3. **Test Failures**: Check MongoDB connection and database setup
4. **Performance Issues**: Verify indexes are created and consider cleanup

### Debug Mode

Enable debug logging to troubleshoot issues:

```go
// Set log level for debugging
log.SetLevel(log.DebugLevel)
```

## Migration

### From Mock to MongoDB

To migrate from mock adapter to MongoDB:

1. Replace `integration.NewMockMongoAdapter()` with `integration.NewMongoAdapterLocal()`
2. Add `-tags mongodb` to build commands
3. Ensure MongoDB is running and accessible
4. Run database setup scripts

### Data Migration

For existing data migration:

1. Export data from current system
2. Transform to RPG Stats format
3. Import using the MongoDB adapter
4. Verify data integrity

## Contributing

When contributing to the MongoDB integration:

1. Follow the existing code patterns
2. Add comprehensive tests
3. Update documentation
4. Consider performance implications
5. Test with both mock and real MongoDB

## License

This MongoDB integration is part of the RPG Stats Sub-System and follows the same license terms.
