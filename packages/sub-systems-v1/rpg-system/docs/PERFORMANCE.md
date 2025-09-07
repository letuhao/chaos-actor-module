# RPG System Performance Guide

## Overview

This guide covers performance optimization techniques for the RPG System, including caching strategies, memory management, and database optimization.

## Caching Strategies

### Snapshot Caching

The system uses deterministic hashing to enable efficient caching of stat calculations.

```go
// Enable caching by checking snapshot hash
snapshot := resolver.ComputeSnapshot(input)
if snapshot.Hash != "" {
    // Cache the result using the hash as key
    cache.Set(snapshot.Hash, snapshot)
}

// Check cache before computation
if cached, exists := cache.Get(inputHash); exists {
    return cached.(*StatSnapshot)
}
```

### Cache Invalidation

```go
// Invalidate cache when character data changes
func (cai *CoreActorIntegration) UpdatePlayerStats(ctx context.Context, actorID string, updates *PlayerStatsUpdates) (*CoreActorContribution, error) {
    // Update data
    // ...
    
    // Invalidate cache
    cache.Delete(actorID)
    
    // Rebuild contribution
    return cai.BuildCoreActorContribution(ctx, actorID)
}
```

### Memory-Efficient Caching

```go
// Use weak references for large objects
type CacheEntry struct {
    Snapshot *StatSnapshot
    LastUsed time.Time
    RefCount int
}

// Implement LRU eviction
func (c *Cache) evictOldEntries() {
    now := time.Now()
    for key, entry := range c.entries {
        if now.Sub(entry.LastUsed) > c.maxAge {
            delete(c.entries, key)
        }
    }
}
```

## Memory Management

### Object Pooling

```go
// Pool for frequently allocated objects
var snapshotPool = sync.Pool{
    New: func() interface{} {
        return &StatSnapshot{
            Stats:     make(map[StatKey]float64),
            Breakdown: make(map[StatKey]*StatBreakdown),
        }
    },
}

// Use pooled objects
func (sr *StatResolver) ComputeSnapshot(input ComputeInput) *StatSnapshot {
    snapshot := snapshotPool.Get().(*StatSnapshot)
    defer snapshotPool.Put(snapshot)
    
    // Reset and reuse
    snapshot.ActorID = input.ActorID
    snapshot.Version = 1
    snapshot.Ts = time.Now().Unix()
    
    // Clear maps
    for k := range snapshot.Stats {
        delete(snapshot.Stats, k)
    }
    for k := range snapshot.Breakdown {
        delete(snapshot.Breakdown, k)
    }
    
    // ... computation logic
    
    return snapshot
}
```

### Slice Reuse

```go
// Reuse slices to reduce allocations
type Resolver struct {
    modifierBuffer []StatModifier
    statBuffer     []StatKey
}

func (r *Resolver) ComputeSnapshot(input ComputeInput) *StatSnapshot {
    // Reuse modifier buffer
    modifiers := r.modifierBuffer[:0]
    modifiers = r.CollectAllModifiers(input, modifiers)
    
    // Reuse stat buffer
    stats := r.statBuffer[:0]
    for statKey := range input.BaseAllocations {
        stats = append(stats, statKey)
    }
    
    // ... computation logic
}
```

### String Interning

```go
// Intern frequently used strings
var stringIntern = make(map[string]string)

func internString(s string) string {
    if interned, exists := stringIntern[s]; exists {
        return interned
    }
    stringIntern[s] = s
    return s
}

// Use interned strings for stat keys
func (r *Registry) GetStatDefinition(statKey StatKey) (*StatDef, bool) {
    internedKey := internString(string(statKey))
    return r.statDefinitions[StatKey(internedKey)], true
}
```

## Database Optimization

### Connection Pooling

```go
// Configure MongoDB connection pool
clientOptions := options.Client().
    ApplyURI("mongodb://localhost:27017").
    SetMaxPoolSize(100).           // Maximum connections
    SetMinPoolSize(10).            // Minimum connections
    SetMaxConnIdleTime(30 * time.Second).
    SetConnectTimeout(10 * time.Second).
    SetSocketTimeout(30 * time.Second)
```

### Batch Operations

```go
// Batch multiple operations
func (ma *MongoAdapter) BatchUpdatePlayerStats(updates []PlayerStatsUpdate) error {
    var operations []mongo.WriteModel
    
    for _, update := range updates {
        operation := mongo.NewUpdateOneModel().
            SetFilter(bson.M{"actor_id": update.ActorID}).
            SetUpdate(bson.M{"$set": update.Progress}).
            SetUpsert(true)
        operations = append(operations, operation)
    }
    
    _, err := ma.playerProgressCol.BulkWrite(ctx, operations)
    return err
}
```

### Index Optimization

```go
// Create compound indexes for common queries
indexes := []mongo.IndexModel{
    // Player progress by actor_id
    {
        Keys: bson.D{{Key: "actor_id", Value: 1}},
        Options: options.Index().SetUnique(true),
    },
    // Effects by actor_id and expiration
    {
        Keys: bson.D{
            {Key: "actor_id", Value: 1},
            {Key: "expires_at", Value: 1},
        },
    },
    // Equipment by actor_id and slot
    {
        Keys: bson.D{
            {Key: "actor_id", Value: 1},
            {Key: "slot", Value: 1},
        },
        Options: options.Index().SetUnique(true),
    },
}
```

### Query Optimization

```go
// Use projection to limit returned data
func (ma *MongoAdapter) GetPlayerProgress(ctx context.Context, actorID string) (*PlayerProgress, error) {
    var progress PlayerProgress
    err := ma.playerProgressCol.FindOne(
        ctx,
        bson.M{"actor_id": actorID},
        options.FindOne().SetProjection(bson.M{
            "actor_id": 1,
            "level": 1,
            "xp": 1,
            "allocations": 1,
            "last_updated": 1,
        }),
    ).Decode(&progress)
    
    return &progress, err
}
```

## Algorithm Optimization

### Modifier Processing

```go
// Optimize modifier collection
func (sr *StatResolver) CollectAllModifiers(input ComputeInput) []StatModifier {
    // Pre-allocate with estimated capacity
    modifiers := make([]StatModifier, 0, 64)
    
    // Collect from all sources
    modifiers = append(modifiers, input.Items...)
    modifiers = append(modifiers, input.Titles...)
    modifiers = append(modifiers, input.Passives...)
    modifiers = append(modifiers, input.Buffs...)
    modifiers = append(modifiers, input.Debuffs...)
    modifiers = append(modifiers, input.Auras...)
    modifiers = append(modifiers, input.Environment...)
    
    return modifiers
}
```

### Stat Calculation

```go
// Optimize stat calculation loop
func (sr *StatResolver) calculateStats(input ComputeInput, modifiers []StatModifier) map[StatKey]float64 {
    stats := make(map[StatKey]float64, len(input.BaseAllocations))
    
    // Calculate base values first
    for statKey, allocation := range input.BaseAllocations {
        baseValue := sr.registry.CalculateLevelValue(statKey, input.Level)
        stats[statKey] = baseValue + float64(allocation)
    }
    
    // Apply modifiers in batches by stat
    statModifiers := make(map[StatKey][]StatModifier)
    for _, modifier := range modifiers {
        statModifiers[modifier.Key] = append(statModifiers[modifier.Key], modifier)
    }
    
    // Process each stat's modifiers
    for statKey, statMods := range statModifiers {
        if len(statMods) > 0 {
            stats[statKey] = sr.applyModifiers(stats[statKey], statMods)
        }
    }
    
    return stats
}
```

### Parallel Processing

```go
// Process multiple characters in parallel
func (cai *CoreActorIntegration) BatchBuildContributions(ctx context.Context, actorIDs []string) ([]*CoreActorContribution, error) {
    results := make([]*CoreActorContribution, len(actorIDs))
    errors := make([]error, len(actorIDs))
    
    var wg sync.WaitGroup
    for i, actorID := range actorIDs {
        wg.Add(1)
        go func(index int, id string) {
            defer wg.Done()
            contribution, err := cai.BuildCoreActorContribution(ctx, id)
            results[index] = contribution
            errors[index] = err
        }(i, actorID)
    }
    
    wg.Wait()
    
    // Check for errors
    for i, err := range errors {
        if err != nil {
            return nil, fmt.Errorf("error processing actor %s: %w", actorIDs[i], err)
        }
    }
    
    return results, nil
}
```

## Profiling and Monitoring

### CPU Profiling

```go
import _ "net/http/pprof"

func main() {
    // Enable profiling
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // Your application code
}
```

### Memory Profiling

```go
// Monitor memory usage
func (sr *StatResolver) ComputeSnapshot(input ComputeInput) *StatSnapshot {
    var m1, m2 runtime.MemStats
    runtime.ReadMemStats(&m1)
    
    // ... computation logic
    
    runtime.ReadMemStats(&m2)
    if m2.Alloc > m1.Alloc {
        log.Printf("Memory allocated: %d bytes", m2.Alloc-m1.Alloc)
    }
    
    return snapshot
}
```

### Performance Metrics

```go
// Track performance metrics
type PerformanceMetrics struct {
    ComputeTime    time.Duration
    CacheHits      int64
    CacheMisses    int64
    ModifierCount  int
    StatCount      int
}

func (sr *StatResolver) ComputeSnapshotWithMetrics(input ComputeInput) (*StatSnapshot, *PerformanceMetrics) {
    start := time.Now()
    metrics := &PerformanceMetrics{}
    
    // Check cache
    if cached := sr.cache.Get(inputHash); cached != nil {
        metrics.CacheHits++
        return cached.(*StatSnapshot), metrics
    }
    metrics.CacheMisses++
    
    // Count modifiers and stats
    modifiers := sr.CollectAllModifiers(input)
    metrics.ModifierCount = len(modifiers)
    metrics.StatCount = len(input.BaseAllocations)
    
    // Compute snapshot
    snapshot := sr.computeSnapshotInternal(input, modifiers)
    
    // Store in cache
    sr.cache.Set(inputHash, snapshot)
    
    metrics.ComputeTime = time.Since(start)
    return snapshot, metrics
}
```

## Benchmarking

### Benchmark Tests

```go
func BenchmarkStatResolver_ComputeSnapshot(b *testing.B) {
    resolver := resolver.NewStatResolver()
    input := createTestInput()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        resolver.ComputeSnapshot(input)
    }
}

func BenchmarkStatResolver_WithCache(b *testing.B) {
    resolver := resolver.NewStatResolver()
    resolver.EnableCaching()
    input := createTestInput()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        resolver.ComputeSnapshot(input)
    }
}

func BenchmarkBatchProcessing(b *testing.B) {
    integration := createTestIntegration()
    actorIDs := generateTestActorIDs(100)
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        integration.BatchBuildContributions(context.Background(), actorIDs)
    }
}
```

### Load Testing

```go
func TestLoad(t *testing.T) {
    integration := createTestIntegration()
    
    // Simulate concurrent requests
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            integration.BuildCoreActorContribution(context.Background(), "test_actor")
        }()
    }
    
    wg.Wait()
}
```

## Configuration Tuning

### Cache Configuration

```go
type CacheConfig struct {
    MaxSize      int
    MaxAge       time.Duration
    CleanupInterval time.Duration
}

func NewCache(config CacheConfig) *Cache {
    cache := &Cache{
        entries: make(map[string]*CacheEntry),
        config:  config,
    }
    
    // Start cleanup goroutine
    go cache.cleanupLoop()
    
    return cache
}
```

### Database Configuration

```go
type DatabaseConfig struct {
    MaxPoolSize     int
    MinPoolSize     int
    MaxConnIdleTime time.Duration
    ConnectTimeout  time.Duration
    SocketTimeout   time.Duration
}

func NewMongoAdapter(config DatabaseConfig) (*MongoAdapter, error) {
    clientOptions := options.Client().
        ApplyURI("mongodb://localhost:27017").
        SetMaxPoolSize(config.MaxPoolSize).
        SetMinPoolSize(config.MinPoolSize).
        SetMaxConnIdleTime(config.MaxConnIdleTime).
        SetConnectTimeout(config.ConnectTimeout).
        SetSocketTimeout(config.SocketTimeout)
    
    // ... rest of implementation
}
```

## Best Practices

1. **Profile First**: Always profile before optimizing
2. **Measure Impact**: Measure the impact of each optimization
3. **Cache Wisely**: Cache expensive computations, not cheap ones
4. **Batch Operations**: Group database operations when possible
5. **Monitor Memory**: Keep track of memory usage and leaks
6. **Test Under Load**: Test with realistic load patterns
7. **Use Pools**: Reuse objects to reduce allocations
8. **Optimize Hot Paths**: Focus on frequently executed code
9. **Consider Trade-offs**: Balance memory usage vs. computation time
10. **Document Changes**: Keep track of performance optimizations
