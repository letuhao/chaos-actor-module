# RPG System API Documentation

## Core Components

### StatResolver

The `StatResolver` is the core calculation engine that computes final stat values from base allocations and modifiers.

```go
type StatResolver struct {
    registry *registry.StatRegistry
}

func NewStatResolver() *StatResolver
func (sr *StatResolver) ComputeSnapshot(input ComputeInput) *StatSnapshot
func (sr *StatResolver) CollectAllModifiers(input ComputeInput) []StatModifier
func (sr *StatResolver) FilterModifiersForStat(modifiers []StatModifier, statKey StatKey) []StatModifier
```

### StatRegistry

Manages stat definitions, level curves, and derived formulas.

```go
type StatRegistry struct {
    statDefinitions map[StatKey]*StatDef
    levelCurves     map[StatKey]*LevelCurve
    derivedFormulas map[StatKey]*DerivedFormula
}

func NewStatRegistry() *StatRegistry
func (r *StatRegistry) GetStatDefinition(statKey StatKey) (*StatDef, bool)
func (r *StatRegistry) GetAllPrimaryStats() []*StatDef
func (r *StatRegistry) GetAllDerivedStats() []*StatDef
func (r *StatRegistry) CalculateLevelValue(statKey StatKey, level int) float64
func (r *StatRegistry) GetDerivedFormula(statKey StatKey) (*DerivedFormula, bool)
```

### ProgressionService

Handles character progression including XP, leveling, and stat allocation.

```go
type ProgressionService struct {
    registry       *registry.StatRegistry
    playerProgress map[string]*PlayerProgress
}

func NewProgressionService() *ProgressionService
func (ps *ProgressionService) GrantXP(actorID string, xpAmount int64) (*ProgressionResult, error)
func (ps *ProgressionService) AllocatePoints(actorID string, statKey StatKey, points int) error
func (ps *ProgressionService) Respec(actorID string) error
func (ps *ProgressionService) GetPlayerProgress(actorID string) (*PlayerProgress, error)
func (ps *ProgressionService) GetAvailablePoints(actorID string) (int, error)
func (ps *ProgressionService) GetLevelProgression(level int) (*LevelProgression, error)
func (ps *ProgressionService) GetXPToNextLevel(actorID string) (int64, error)
func (ps *ProgressionService) GetLevelProgress(actorID string) (float64, error)
```

### CoreActorIntegration

Provides integration between the RPG System and Core Actor.

```go
type CoreActorIntegration struct {
    resolver         *resolver.StatResolver
    adapter          DatabaseAdapter
    progressionService *ProgressionService
}

func NewCoreActorIntegration(adapter DatabaseAdapter) *CoreActorIntegration
func (cai *CoreActorIntegration) BuildCoreActorContribution(ctx context.Context, actorID string) (*CoreActorContribution, error)
func (cai *CoreActorIntegration) UpdatePlayerStats(ctx context.Context, actorID string, updates *PlayerStatsUpdates) (*CoreActorContribution, error)
func (cai *CoreActorIntegration) GetStatBreakdown(ctx context.Context, actorID string, statKey StatKey) (*StatBreakdown, error)
```

## Data Types

### StatSnapshot

Represents the final calculated stats for a character.

```go
type StatSnapshot struct {
    ActorID   string
    Stats     map[StatKey]float64
    Breakdown map[StatKey]*StatBreakdown
    Version   int
    Ts        int64
    Hash      string
}
```

### StatBreakdown

Provides detailed information about how a stat was calculated.

```go
type StatBreakdown struct {
    Base            float64
    AdditiveFlat    float64
    AdditivePct     float64
    Multiplicative  float64
    Override        *float64
    CappedTo        *float64
    Final           float64
    Modifiers       []ModifierContribution
}
```

### PlayerProgress

Tracks character progression data.

```go
type PlayerProgress struct {
    ActorID     string
    Level       int
    XP          int64
    Allocations map[StatKey]int
    LastUpdated int64
}
```

### StatModifier

Represents a modifier that affects stat calculations.

```go
type StatModifier struct {
    Key         StatKey
    Op          ModifierOp
    Value       float64
    Source      ModifierSourceRef
    Priority    int
    Stack       int
    Conditions  *ModifierConditions
}
```

## Database Adapters

### DatabaseAdapter Interface

```go
type DatabaseAdapter interface {
    GetPlayerStatsSummary(ctx context.Context, actorID string) (*PlayerStatsSummary, error)
    GetPlayerProgress(ctx context.Context, actorID string) (*PlayerProgress, error)
    SavePlayerProgress(ctx context.Context, progress *PlayerProgress) error
    AddEffect(ctx context.Context, actorID string, effect StatModifier, duration time.Duration) error
    RemoveEffect(ctx context.Context, actorID, effectID string) error
    EquipItem(ctx context.Context, actorID, slot string, item StatModifier) error
    UnequipItem(ctx context.Context, actorID, slot string) error
    GrantTitle(ctx context.Context, actorID, titleID string, title StatModifier) error
    Close() error
}
```

### MockMongoAdapter

In-memory implementation for testing and development.

```go
func NewMockMongoAdapter() *MockMongoAdapter
```

### MongoAdapter (with build tag)

MongoDB implementation for production use.

```go
func NewMongoAdapter(uri, databaseName string) (*MongoAdapter, error)
func NewMongoAdapterLocal(databaseName string) (*MongoAdapter, error)
```

## Usage Examples

### Basic Stat Calculation

```go
// Create resolver
resolver := resolver.NewStatResolver()

// Define input
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

// Compute snapshot
snapshot := resolver.ComputeSnapshot(input)

// Access results
strValue := snapshot.Stats[model.STR]
hpMax := snapshot.Stats[model.HP_MAX]
```

### Character Progression

```go
// Create progression service
progression := integration.NewProgressionService()

// Grant XP
result, err := progression.GrantXP("player_001", 1000)
if err != nil {
    log.Fatal(err)
}

// Allocate stat points
err = progression.AllocatePoints("player_001", model.STR, 3)
if err != nil {
    log.Fatal(err)
}

// Get progress
progress, err := progression.GetPlayerProgress("player_001")
if err != nil {
    log.Fatal(err)
}
```

### Core Actor Integration

```go
// Create integration
adapter := integration.NewMockMongoAdapter()
integration := integration.NewCoreActorIntegration(adapter)

// Build contribution
contribution, err := integration.BuildCoreActorContribution(ctx, "player_001")
if err != nil {
    log.Fatal(err)
}

// Use with Core Actor
fmt.Printf("Primary Stats: HP=%d, Attack=%d, Defense=%d\n",
    contribution.Primary.HPMax,
    contribution.Primary.Attack,
    contribution.Primary.Defense)
```

## Configuration

### Stat Definitions

Stats are configured in `internal/registry/registry.go`:

```go
// Primary stats
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

// Derived stats
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

### Level Curves

Level scaling is configured with soft caps:

```go
{
    BaseValue:    10.0,
    PerLevel:     2.0,
    MaxLevel:     100,
    SoftCapLevel: 50,
    SoftCapValue: 0.5,
}
```

### Derived Formulas

Derived stats use configurable formulas:

```go
{
    StatKey:    model.HP_MAX,
    Formula:    "END * 20 + STR * 5",
    Components: []model.StatKey{model.END, model.STR},
    BaseValue:  100.0,
}
```

## Error Handling

The system provides comprehensive error handling:

```go
// Check for specific errors
if err != nil {
    if strings.Contains(err.Error(), "player not found") {
        // Handle missing player
    } else if strings.Contains(err.Error(), "insufficient points") {
        // Handle allocation errors
    } else {
        // Handle other errors
    }
}
```

## Performance Considerations

- **Caching**: Use snapshot hashing for efficient caching
- **Batch Operations**: Process multiple modifiers together
- **Memory Management**: Reuse objects where possible
- **Database Optimization**: Use appropriate indexes and queries

## Testing

The system includes comprehensive tests:

```bash
# Run all tests
go test ./test -v

# Run specific test suite
go test ./test -v -run TestStatResolver

# Run with coverage
go test ./test -v -cover
```
