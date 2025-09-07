# Performance Guide - Actor Core v2.0

## Tổng Quan

Performance Guide cho Actor Core v2.0 cung cấp hướng dẫn chi tiết về tối ưu hóa hiệu suất, bao gồm best practices, optimization strategies, và performance monitoring.

## 1. Performance Targets

### 1.1 Calculation Performance
| Operation | Target | Acceptable | Critical |
|-----------|--------|------------|----------|
| PrimaryCore Calculation | < 1ms | < 2ms | > 5ms |
| Derived Stats Calculation | < 5ms | < 10ms | > 20ms |
| Flexible Systems Calculation | < 10ms | < 20ms | > 50ms |
| Multi-System Calculation | < 20ms | < 50ms | > 100ms |
| Configuration Loading | < 100ms | < 200ms | > 500ms |
| Hot Reload | < 50ms | < 100ms | > 200ms |

### 1.2 Memory Performance
| Metric | Target | Acceptable | Critical |
|--------|--------|------------|----------|
| Memory per Actor | < 1KB | < 2KB | > 5KB |
| Memory for 1000 Actors | < 100MB | < 200MB | > 500MB |
| GC Time | < 5% | < 10% | > 20% |
| Memory Leaks | 0 | < 1MB/hour | > 10MB/hour |

### 1.3 Throughput Performance
| Operation | Target | Acceptable | Critical |
|-----------|--------|------------|----------|
| Calculations per Second | > 10,000 | > 5,000 | < 1,000 |
| Concurrent Calculations | > 1,000 | > 500 | < 100 |
| Configuration Updates/sec | > 100 | > 50 | < 10 |

## 2. Optimization Strategies

### 2.1 Formula Engine Optimization

#### 2.1.1 Pre-compilation
```go
// ✅ GOOD: Pre-compile formulas at startup
type OptimizedFormulaEngine struct {
    compiledFormulas map[string]*CompiledFormula
    statCache        map[string]float64
    dependencyGraph  *DependencyGraph
    calculationOrder []string
}

func (engine *FormulaEngine) CompileFormulas() error {
    for id, formula := range engine.formulas {
        compiled := &CompiledFormula{
            Expression:   formula.Expression,
            Dependencies: formula.Dependencies,
            CompiledCode: engine.compileExpression(formula.Expression), // Pre-compile
            CacheKey:     engine.generateCacheKey(formula),
            LastUpdate:   time.Now().Unix(),
        }
        engine.compiledFormulas[id] = compiled
    }
    return nil
}

// ❌ BAD: Compile every time
func (engine *FormulaEngine) CalculateStat(statID string, stats map[string]float64) float64 {
    formula := engine.formulas[statID]
    return engine.parseAndExecute(formula.Expression, stats) // Slow!
}
```

#### 2.1.2 Caching Strategy
```go
// ✅ GOOD: Smart caching with invalidation
type SmartCache struct {
    cache        map[string]float64
    dependencies map[string][]string
    lastUpdate   map[string]int64
    ttl          int64
}

func (sc *SmartCache) Get(statID string, stats map[string]float64) (float64, bool) {
    // Check if cache is valid
    if sc.isCacheValid(statID, stats) {
        return sc.cache[statID], true
    }
    return 0, false
}

func (sc *SmartCache) Set(statID string, value float64, stats map[string]float64) {
    sc.cache[statID] = value
    sc.lastUpdate[statID] = time.Now().Unix()
}

func (sc *SmartCache) isCacheValid(statID string, stats map[string]float64) bool {
    if value, exists := sc.cache[statID]; !exists {
        return false
    }
    
    // Check TTL
    if time.Now().Unix()-sc.lastUpdate[statID] > sc.ttl {
        return false
    }
    
    // Check dependencies
    for _, dep := range sc.dependencies[statID] {
        if _, exists := stats[dep]; !exists {
            return false
        }
    }
    
    return true
}
```

#### 2.1.3 Dependency Optimization
```go
// ✅ GOOD: Calculate in dependency order
func (engine *FormulaEngine) CalculateAllStats(stats map[string]float64) map[string]float64 {
    results := make(map[string]float64)
    
    // Calculate in dependency order to avoid redundant calculations
    for _, statID := range engine.calculationOrder {
        if compiled, exists := engine.compiledFormulas[statID]; exists {
            results[statID] = engine.executeCompiledCode(compiled.CompiledCode, stats)
        }
    }
    
    return results
}

// ❌ BAD: Calculate in random order
func (engine *FormulaEngine) CalculateAllStats(stats map[string]float64) map[string]float64 {
    results := make(map[string]float64)
    
    // Random order - may cause redundant calculations
    for statID := range engine.formulas {
        results[statID] = engine.CalculateStat(statID, stats)
    }
    
    return results
}
```

### 2.2 Memory Optimization

#### 2.2.1 Object Pooling
```go
// ✅ GOOD: Object pooling to reduce GC pressure
type ObjectPool struct {
    statMaps    sync.Pool
    formulaMaps sync.Pool
    actorCores  sync.Pool
}

func (pool *ObjectPool) GetStatMap() map[string]float64 {
    if v := pool.statMaps.Get(); v != nil {
        return v.(map[string]float64)
    }
    return make(map[string]float64)
}

func (pool *ObjectPool) PutStatMap(m map[string]float64) {
    // Clear map but keep capacity
    for k := range m {
        delete(m, k)
    }
    pool.statMaps.Put(m)
}

func (pool *ObjectPool) GetActorCore() *ActorCore {
    if v := pool.actorCores.Get(); v != nil {
        return v.(*ActorCore)
    }
    return &ActorCore{}
}

func (pool *ObjectPool) PutActorCore(ac *ActorCore) {
    // Reset but keep capacity
    ac.Reset()
    pool.actorCores.Put(ac)
}
```

#### 2.2.2 Memory Pre-allocation
```go
// ✅ GOOD: Pre-allocate with known capacity
type PreAllocatedActorCore struct {
    PrimaryStats    map[string]int64
    DerivedStats    map[string]float64
    FlexibleStats   *FlexibleStats
    SpeedSystem     *FlexibleSpeedSystem
    // ... other fields
}

func NewPreAllocatedActorCore() *PreAllocatedActorCore {
    return &PreAllocatedActorCore{
        PrimaryStats:  make(map[string]int64, 20),   // Pre-allocate for 20 stats
        DerivedStats:  make(map[string]float64, 50), // Pre-allocate for 50 stats
        FlexibleStats: &FlexibleStats{
            CustomPrimary:    make(map[string]int64, 10),
            CustomDerived:    make(map[string]float64, 20),
            SubSystemStats:   make(map[string]map[string]float64, 5),
        },
        // ... other pre-allocations
    }
}

// ❌ BAD: Allocate on demand
type NaiveActorCore struct {
    PrimaryStats map[string]int64
    DerivedStats map[string]float64
    // ... other fields
}

func NewNaiveActorCore() *NaiveActorCore {
    return &NaiveActorCore{
        PrimaryStats: make(map[string]int64), // Will grow dynamically
        DerivedStats: make(map[string]float64), // Will grow dynamically
    }
}
```

#### 2.2.3 String Interning
```go
// ✅ GOOD: String interning to reduce memory usage
type StringInterner struct {
    strings map[string]string
    mutex   sync.RWMutex
}

func (si *StringInterner) Intern(s string) string {
    si.mutex.RLock()
    if interned, exists := si.strings[s]; exists {
        si.mutex.RUnlock()
        return interned
    }
    si.mutex.RUnlock()
    
    si.mutex.Lock()
    defer si.mutex.Unlock()
    
    if interned, exists := si.strings[s]; exists {
        return interned
    }
    
    si.strings[s] = s
    return s
}

// Use interned strings for stat names
func (ac *ActorCore) SetStat(statName string, value int64) {
    ac.PrimaryStats[ac.stringInterner.Intern(statName)] = value
}
```

### 2.3 Concurrent Processing

#### 2.3.1 Parallel Calculations
```go
// ✅ GOOD: Parallel calculation for independent stats
func (engine *FormulaEngine) CalculateAllStatsParallel(stats map[string]float64) map[string]float64 {
    results := make(map[string]float64)
    var wg sync.WaitGroup
    var mutex sync.Mutex
    
    // Group stats by dependency level
    dependencyLevels := engine.groupStatsByDependencyLevel()
    
    for level, statIDs := range dependencyLevels {
        wg.Add(len(statIDs))
        
        for _, statID := range statIDs {
            go func(id string) {
                defer wg.Done()
                
                if compiled, exists := engine.compiledFormulas[id]; exists {
                    value := engine.executeCompiledCode(compiled.CompiledCode, stats)
                    
                    mutex.Lock()
                    results[id] = value
                    mutex.Unlock()
                }
            }(statID)
        }
        
        wg.Wait() // Wait for current level to complete before next level
    }
    
    return results
}
```

#### 2.3.2 Worker Pool
```go
// ✅ GOOD: Worker pool for high-throughput calculations
type CalculationWorkerPool struct {
    workers    int
    jobQueue   chan CalculationJob
    resultQueue chan CalculationResult
    workers    []*CalculationWorker
}

type CalculationJob struct {
    ID       string
    Stats    map[string]float64
    StatID   string
}

type CalculationResult struct {
    ID    string
    Value float64
    Error error
}

func (pool *CalculationWorkerPool) Start() {
    for i := 0; i < pool.workers; i++ {
        worker := &CalculationWorker{
            id:         i,
            jobQueue:   pool.jobQueue,
            resultQueue: pool.resultQueue,
            engine:     pool.engine,
        }
        go worker.Start()
        pool.workers = append(pool.workers, worker)
    }
}

func (pool *CalculationWorkerPool) SubmitJob(job CalculationJob) {
    pool.jobQueue <- job
}

func (pool *CalculationWorkerPool) GetResult() CalculationResult {
    return <-pool.resultQueue
}
```

### 2.4 Configuration Optimization

#### 2.4.1 Lazy Loading
```go
// ✅ GOOD: Lazy loading for configuration
type LazyConfigurationManager struct {
    configs        map[string]*ConfigDefinition
    loadedConfigs  map[string]bool
    mutex          sync.RWMutex
}

func (lcm *LazyConfigurationManager) GetConfig(id string) (*ConfigDefinition, error) {
    lcm.mutex.RLock()
    if loaded, exists := lcm.loadedConfigs[id]; exists && loaded {
        config := lcm.configs[id]
        lcm.mutex.RUnlock()
        return config, nil
    }
    lcm.mutex.RUnlock()
    
    lcm.mutex.Lock()
    defer lcm.mutex.Unlock()
    
    // Double-check after acquiring write lock
    if loaded, exists := lcm.loadedConfigs[id]; exists && loaded {
        return lcm.configs[id], nil
    }
    
    // Load configuration
    config, err := lcm.loadConfigFromFile(id)
    if err != nil {
        return nil, err
    }
    
    lcm.configs[id] = config
    lcm.loadedConfigs[id] = true
    
    return config, nil
}
```

#### 2.4.2 Configuration Caching
```go
// ✅ GOOD: Configuration caching with TTL
type CachedConfigurationManager struct {
    configs        map[string]*ConfigDefinition
    cache          map[string]*CachedConfig
    ttl            time.Duration
    mutex          sync.RWMutex
}

type CachedConfig struct {
    Config    *ConfigDefinition
    Timestamp time.Time
}

func (ccm *CachedConfigurationManager) GetConfig(id string) (*ConfigDefinition, error) {
    ccm.mutex.RLock()
    if cached, exists := ccm.cache[id]; exists {
        if time.Since(cached.Timestamp) < ccm.ttl {
            ccm.mutex.RUnlock()
            return cached.Config, nil
        }
    }
    ccm.mutex.RUnlock()
    
    // Load and cache configuration
    ccm.mutex.Lock()
    defer ccm.mutex.Unlock()
    
    config, err := ccm.loadConfigFromFile(id)
    if err != nil {
        return nil, err
    }
    
    ccm.cache[id] = &CachedConfig{
        Config:    config,
        Timestamp: time.Now(),
    }
    
    return config, nil
}
```

## 3. Performance Monitoring

### 3.1 Metrics Collection
```go
// ✅ GOOD: Comprehensive performance metrics
type PerformanceMetrics struct {
    CalculationTime    time.Duration
    MemoryUsage        uint64
    CacheHitRate      float64
    CacheMissRate     float64
    GCTime            time.Duration
    GCPauseTime       time.Duration
    GoroutineCount    int
    MutexWaitTime     time.Duration
    ChannelWaitTime   time.Duration
}

type PerformanceMonitor struct {
    metrics    *PerformanceMetrics
    startTime  time.Time
    mutex      sync.RWMutex
}

func (pm *PerformanceMonitor) StartCalculation() {
    pm.mutex.Lock()
    pm.startTime = time.Now()
    pm.mutex.Unlock()
}

func (pm *PerformanceMonitor) EndCalculation() {
    pm.mutex.Lock()
    pm.metrics.CalculationTime = time.Since(pm.startTime)
    pm.mutex.Unlock()
}

func (pm *PerformanceMonitor) RecordCacheHit() {
    pm.mutex.Lock()
    pm.metrics.CacheHitRate++
    pm.mutex.Unlock()
}

func (pm *PerformanceMonitor) RecordCacheMiss() {
    pm.mutex.Lock()
    pm.metrics.CacheMissRate++
    pm.mutex.Unlock()
}
```

### 3.2 Performance Profiling
```go
// ✅ GOOD: Built-in performance profiling
func (engine *FormulaEngine) ProfileCalculation(statID string, stats map[string]float64) (float64, *PerformanceProfile) {
    profile := &PerformanceProfile{
        StatID:      statID,
        StartTime:   time.Now(),
        MemoryStart: getMemoryUsage(),
    }
    
    // Enable CPU profiling
    if engine.profilingEnabled {
        engine.cpuProfiler.Start()
    }
    
    // Perform calculation
    result := engine.CalculateStat(statID, stats)
    
    // Stop profiling
    if engine.profilingEnabled {
        engine.cpuProfiler.Stop()
    }
    
    // Record metrics
    profile.EndTime = time.Now()
    profile.Duration = profile.EndTime.Sub(profile.StartTime)
    profile.MemoryEnd = getMemoryUsage()
    profile.MemoryDelta = profile.MemoryEnd - profile.MemoryStart
    
    return result, profile
}
```

### 3.3 Performance Alerts
```go
// ✅ GOOD: Performance alerting system
type PerformanceAlert struct {
    Type        string
    Message     string
    Severity    string
    Timestamp   time.Time
    Metrics     *PerformanceMetrics
}

type PerformanceAlerter struct {
    thresholds  map[string]float64
    alerters    []AlertHandler
    mutex       sync.RWMutex
}

func (pa *PerformanceAlerter) CheckPerformance(metrics *PerformanceMetrics) {
    pa.mutex.RLock()
    defer pa.mutex.RUnlock()
    
    // Check calculation time
    if metrics.CalculationTime > time.Duration(pa.thresholds["calculation_time"])*time.Millisecond {
        pa.alert(PerformanceAlert{
            Type:      "calculation_time",
            Message:   fmt.Sprintf("Calculation time exceeded threshold: %v", metrics.CalculationTime),
            Severity:  "warning",
            Timestamp: time.Now(),
            Metrics:   metrics,
        })
    }
    
    // Check memory usage
    if metrics.MemoryUsage > uint64(pa.thresholds["memory_usage"]) {
        pa.alert(PerformanceAlert{
            Type:      "memory_usage",
            Message:   fmt.Sprintf("Memory usage exceeded threshold: %d bytes", metrics.MemoryUsage),
            Severity:  "critical",
            Timestamp: time.Now(),
            Metrics:   metrics,
        })
    }
    
    // Check cache hit rate
    if metrics.CacheHitRate < pa.thresholds["cache_hit_rate"] {
        pa.alert(PerformanceAlert{
            Type:      "cache_hit_rate",
            Message:   fmt.Sprintf("Cache hit rate below threshold: %.2f%%", metrics.CacheHitRate),
            Severity:  "warning",
            Timestamp: time.Now(),
            Metrics:   metrics,
        })
    }
}
```

## 4. Best Practices

### 4.1 Code Organization
```go
// ✅ GOOD: Organized code structure
type ActorCore struct {
    // Core stats (most frequently accessed)
    PrimaryStats map[string]int64
    DerivedStats map[string]float64
    
    // Flexible systems (less frequently accessed)
    FlexibleStats *FlexibleStats
    SpeedSystem   *FlexibleSpeedSystem
    
    // Configuration (rarely accessed)
    Config *Configuration
}

// ✅ GOOD: Separate hot and cold paths
func (ac *ActorCore) GetStat(statName string) (interface{}, error) {
    // Hot path: check primary stats first
    if value, exists := ac.PrimaryStats[statName]; exists {
        return value, nil
    }
    
    // Hot path: check derived stats
    if value, exists := ac.DerivedStats[statName]; exists {
        return value, nil
    }
    
    // Cold path: check flexible stats
    return ac.getFlexibleStat(statName)
}
```

### 4.2 Error Handling
```go
// ✅ GOOD: Fast error handling
func (engine *FormulaEngine) CalculateStat(statID string, stats map[string]float64) (float64, error) {
    // Fast path: check cache first
    if value, exists := engine.cache[statID]; exists {
        return value, nil
    }
    
    // Fast path: check if formula exists
    compiled, exists := engine.compiledFormulas[statID]
    if !exists {
        return 0, ErrFormulaNotFound
    }
    
    // Fast path: check dependencies
    if !engine.checkDependencies(compiled.Dependencies, stats) {
        return 0, ErrMissingDependencies
    }
    
    // Calculate and cache
    value := engine.executeCompiledCode(compiled.CompiledCode, stats)
    engine.cache[statID] = value
    
    return value, nil
}
```

### 4.3 Memory Management
```go
// ✅ GOOD: Proper memory management
func (ac *ActorCore) Reset() {
    // Clear maps but keep capacity
    for k := range ac.PrimaryStats {
        delete(ac.PrimaryStats, k)
    }
    for k := range ac.DerivedStats {
        delete(ac.DerivedStats, k)
    }
    
    // Reset flexible systems
    if ac.FlexibleStats != nil {
        ac.FlexibleStats.Reset()
    }
    if ac.SpeedSystem != nil {
        ac.SpeedSystem.Reset()
    }
}

func (ac *ActorCore) Clone() *ActorCore {
    clone := &ActorCore{
        PrimaryStats: make(map[string]int64, len(ac.PrimaryStats)),
        DerivedStats: make(map[string]float64, len(ac.DerivedStats)),
    }
    
    // Copy primary stats
    for k, v := range ac.PrimaryStats {
        clone.PrimaryStats[k] = v
    }
    
    // Copy derived stats
    for k, v := range ac.DerivedStats {
        clone.DerivedStats[k] = v
    }
    
    // Clone flexible systems
    if ac.FlexibleStats != nil {
        clone.FlexibleStats = ac.FlexibleStats.Clone()
    }
    if ac.SpeedSystem != nil {
        clone.SpeedSystem = ac.SpeedSystem.Clone()
    }
    
    return clone
}
```

## 5. Performance Testing

### 5.1 Benchmark Tests
```go
// ✅ GOOD: Comprehensive benchmark tests
func BenchmarkPrimaryCoreCalculation(b *testing.B) {
    actorCore := NewActorCore()
    stats := generateTestStats()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        actorCore.CalculateDerivedStats(stats)
    }
}

func BenchmarkFlexibleSystemsCalculation(b *testing.B) {
    actorCore := NewActorCore()
    stats := generateTestStats()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        actorCore.CalculateFlexibleStats(stats)
    }
}

func BenchmarkMultiSystemCalculation(b *testing.B) {
    actorCore := NewActorCore()
    stats := generateTestStats()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        actorCore.CalculateAllStats(stats)
    }
}
```

### 5.2 Memory Tests
```go
// ✅ GOOD: Memory usage tests
func TestMemoryUsage(t *testing.T) {
    var m1, m2 runtime.MemStats
    runtime.GC()
    runtime.ReadMemStats(&m1)
    
    // Create 1000 actors
    actors := make([]*ActorCore, 1000)
    for i := 0; i < 1000; i++ {
        actors[i] = NewActorCore()
    }
    
    runtime.GC()
    runtime.ReadMemStats(&m2)
    
    memoryUsed := m2.Alloc - m1.Alloc
    memoryPerActor := memoryUsed / 1000
    
    // Should be less than 1KB per actor
    if memoryPerActor > 1024 {
        t.Errorf("Memory usage per actor too high: %d bytes", memoryPerActor)
    }
}
```

### 5.3 Concurrent Tests
```go
// ✅ GOOD: Concurrent performance tests
func TestConcurrentCalculation(t *testing.T) {
    actorCore := NewActorCore()
    stats := generateTestStats()
    
    var wg sync.WaitGroup
    numGoroutines := 100
    calculationsPerGoroutine := 1000
    
    start := time.Now()
    
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < calculationsPerGoroutine; j++ {
                actorCore.CalculateDerivedStats(stats)
            }
        }()
    }
    
    wg.Wait()
    duration := time.Since(start)
    
    totalCalculations := numGoroutines * calculationsPerGoroutine
    calculationsPerSecond := float64(totalCalculations) / duration.Seconds()
    
    // Should handle at least 10,000 calculations per second
    if calculationsPerSecond < 10000 {
        t.Errorf("Concurrent performance too low: %.2f calculations/second", calculationsPerSecond)
    }
}
```

## 6. Performance Tuning

### 6.1 Configuration Tuning
```go
// ✅ GOOD: Performance-tuned configuration
type PerformanceConfig struct {
    // Cache settings
    CacheSize        int
    CacheTTL         time.Duration
    CacheEvictionPolicy string
    
    // Worker pool settings
    WorkerPoolSize   int
    JobQueueSize     int
    
    // Memory settings
    PreAllocateMaps  bool
    MapInitialSize   int
    
    // Profiling settings
    EnableProfiling  bool
    ProfileInterval  time.Duration
}

func NewPerformanceTunedActorCore(config *PerformanceConfig) *ActorCore {
    return &ActorCore{
        formulaEngine: NewOptimizedFormulaEngine(config),
        cache:        NewOptimizedCache(config),
        workerPool:   NewWorkerPool(config),
        // ... other optimized components
    }
}
```

### 6.2 Runtime Tuning
```go
// ✅ GOOD: Runtime performance tuning
func (engine *FormulaEngine) TunePerformance() {
    // Adjust cache size based on usage
    if engine.cacheHitRate < 0.8 {
        engine.cacheSize *= 2
    }
    
    // Adjust worker pool size based on load
    if engine.avgCalculationTime > 10*time.Millisecond {
        engine.workerPoolSize++
    }
    
    // Adjust TTL based on stat change frequency
    if engine.statChangeFrequency > 0.1 {
        engine.cacheTTL /= 2
    }
}
```

---

*Tài liệu này cung cấp hướng dẫn chi tiết về tối ưu hóa hiệu suất cho Actor Core v2.0, đảm bảo hệ thống hoạt động hiệu quả và ổn định.*
