# Troubleshooting Guide - Actor Core v2.0

## Tổng Quan

Troubleshooting Guide cho Actor Core v2.0 cung cấp hướng dẫn chi tiết về việc chẩn đoán và khắc phục sự cố, bao gồm common issues, error codes, debugging techniques, và recovery procedures.

## 1. Common Issues & Solutions

### 1.1 Performance Issues

#### 1.1.1 Slow Stat Calculations
**Symptoms:**
- Stat calculations take > 20ms
- High CPU usage during calculations
- Timeout errors

**Diagnosis:**
```bash
# Check calculation performance
curl -s http://localhost:9090/metrics | grep actor_core_stat_calculation_duration

# Check CPU usage
top -bn1 | grep "Cpu(s)"

# Check memory usage
free -h
```

**Solutions:**
```go
// 1. Enable formula caching
func (engine *FormulaEngine) EnableCaching() {
    engine.cacheEnabled = true
    engine.cacheTTL = 5 * time.Minute
}

// 2. Optimize formula compilation
func (engine *FormulaEngine) OptimizeFormulas() {
    for id, formula := range engine.formulas {
        // Pre-compile formulas
        compiled := engine.compileFormula(formula)
        engine.compiledFormulas[id] = compiled
    }
}

// 3. Use parallel calculations
func (engine *FormulaEngine) CalculateStatsParallel(stats map[string]float64) map[string]float64 {
    results := make(map[string]float64)
    var wg sync.WaitGroup
    
    for statID := range engine.formulas {
        wg.Add(1)
        go func(id string) {
            defer wg.Done()
            results[id] = engine.CalculateStat(id, stats)
        }(statID)
    }
    
    wg.Wait()
    return results
}
```

#### 1.1.2 High Memory Usage
**Symptoms:**
- Memory usage > 1GB for 1000 actors
- Frequent garbage collection
- Out of memory errors

**Diagnosis:**
```bash
# Check memory usage
free -h
ps aux | grep actor-core

# Check garbage collection
curl -s http://localhost:9090/metrics | grep go_gc_duration_seconds

# Check memory leaks
go tool pprof http://localhost:8080/debug/pprof/heap
```

**Solutions:**
```go
// 1. Implement object pooling
type ObjectPool struct {
    statMaps sync.Pool
    actorCores sync.Pool
}

func (pool *ObjectPool) GetStatMap() map[string]float64 {
    if v := pool.statMaps.Get(); v != nil {
        return v.(map[string]float64)
    }
    return make(map[string]float64, 50) // Pre-allocate capacity
}

func (pool *ObjectPool) PutStatMap(m map[string]float64) {
    // Clear map but keep capacity
    for k := range m {
        delete(m, k)
    }
    pool.statMaps.Put(m)
}

// 2. Use string interning
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

// 3. Implement memory limits
type MemoryManager struct {
    maxMemory    uint64
    currentMemory uint64
    mutex        sync.RWMutex
}

func (mm *MemoryManager) Allocate(size uint64) error {
    mm.mutex.Lock()
    defer mm.mutex.Unlock()
    
    if mm.currentMemory+size > mm.maxMemory {
        return ErrMemoryLimitExceeded
    }
    
    mm.currentMemory += size
    return nil
}
```

### 1.2 Database Issues

#### 1.2.1 Connection Pool Exhaustion
**Symptoms:**
- "connection pool exhausted" errors
- High connection count
- Slow database queries

**Diagnosis:**
```bash
# Check connection pool status
mongosh --eval "db.serverStatus().connections"

# Check active connections
mongosh --eval "db.currentOp()"

# Check connection pool metrics
curl -s http://localhost:9090/metrics | grep mongodb_connections
```

**Solutions:**
```go
// 1. Optimize connection pool settings
type DatabaseConfig struct {
    MaxPoolSize    int           `yaml:"max_pool_size"`
    MinPoolSize    int           `yaml:"min_pool_size"`
    MaxIdleTime    time.Duration `yaml:"max_idle_time"`
    ConnectTimeout time.Duration `yaml:"connect_timeout"`
}

func NewOptimizedDatabaseConfig() *DatabaseConfig {
    return &DatabaseConfig{
        MaxPoolSize:    100,
        MinPoolSize:    10,
        MaxIdleTime:    30 * time.Minute,
        ConnectTimeout: 10 * time.Second,
    }
}

// 2. Implement connection retry logic
type DatabaseManager struct {
    client    *mongo.Client
    retryCount int
    retryDelay time.Duration
}

func (dm *DatabaseManager) ExecuteWithRetry(operation func() error) error {
    for i := 0; i < dm.retryCount; i++ {
        if err := operation(); err != nil {
            if i == dm.retryCount-1 {
                return err
            }
            time.Sleep(dm.retryDelay)
            continue
        }
        return nil
    }
    return ErrMaxRetriesExceeded
}

// 3. Implement connection health checks
func (dm *DatabaseManager) HealthCheck() error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    return dm.client.Ping(ctx, nil)
}
```

#### 1.2.2 Slow Database Queries
**Symptoms:**
- Database queries take > 1 second
- High database CPU usage
- Timeout errors

**Diagnosis:**
```bash
# Check slow queries
mongosh --eval "db.setProfilingLevel(2, {slowms: 100})"
mongosh --eval "db.system.profile.find().sort({ts: -1}).limit(10)"

# Check database performance
mongosh --eval "db.stats()"

# Check indexes
mongosh --eval "db.actors.getIndexes()"
```

**Solutions:**
```go
// 1. Optimize database queries
func (dm *DatabaseManager) GetActorOptimized(actorID string) (*Actor, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Use projection to limit fields
    opts := options.FindOne().SetProjection(bson.M{
        "primary_stats": 1,
        "derived_stats": 1,
        "updated_at": 1,
    })
    
    var actor Actor
    err := dm.collection.FindOne(ctx, bson.M{"_id": actorID}, opts).Decode(&actor)
    return &actor, err
}

// 2. Implement query caching
type QueryCache struct {
    cache map[string]*CacheEntry
    mutex sync.RWMutex
    ttl   time.Duration
}

type CacheEntry struct {
    Data      interface{}
    Timestamp time.Time
}

func (qc *QueryCache) Get(key string) (interface{}, bool) {
    qc.mutex.RLock()
    defer qc.mutex.RUnlock()
    
    entry, exists := qc.cache[key]
    if !exists {
        return nil, false
    }
    
    if time.Since(entry.Timestamp) > qc.ttl {
        return nil, false
    }
    
    return entry.Data, true
}

// 3. Create proper indexes
func (dm *DatabaseManager) CreateIndexes() error {
    ctx := context.Background()
    
    // Create compound index for common queries
    indexModel := mongo.IndexModel{
        Keys: bson.D{
            {Key: "primary_stats.vitality", Value: 1},
            {Key: "primary_stats.endurance", Value: 1},
            {Key: "updated_at", Value: -1},
        },
    }
    
    _, err := dm.collection.Indexes().CreateOne(ctx, indexModel)
    return err
}
```

### 1.3 Configuration Issues

#### 1.3.1 Invalid Configuration
**Symptoms:**
- Application fails to start
- Configuration validation errors
- Unexpected behavior

**Diagnosis:**
```bash
# Check configuration validation
./actor-core --validate-config

# Check configuration syntax
yaml-lint config/application.yml

# Check environment variables
env | grep ACTOR_CORE
```

**Solutions:**
```go
// 1. Implement configuration validation
type ConfigurationValidator struct {
    rules map[string]ValidationRule
}

type ValidationRule struct {
    Required bool
    Type     string
    Min      interface{}
    Max      interface{}
    Pattern  string
}

func (cv *ConfigurationValidator) Validate(config *Configuration) error {
    for field, rule := range cv.rules {
        value := cv.getFieldValue(config, field)
        
        if rule.Required && value == nil {
            return fmt.Errorf("required field %s is missing", field)
        }
        
        if value != nil {
            if err := cv.validateValue(field, value, rule); err != nil {
                return err
            }
        }
    }
    
    return nil
}

// 2. Implement configuration hot reload
type ConfigurationManager struct {
    config     *Configuration
    watchers   map[string]*fsnotify.Watcher
    callbacks  []ConfigUpdateCallback
    mutex      sync.RWMutex
}

func (cm *ConfigurationManager) WatchFile(filename string) error {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return err
    }
    
    if err := watcher.Add(filename); err != nil {
        return err
    }
    
    go cm.watchFile(watcher, filename)
    cm.watchers[filename] = watcher
    
    return nil
}

func (cm *ConfigurationManager) watchFile(watcher *fsnotify.Watcher, filename string) {
    for {
        select {
        case event := <-watcher.Events:
            if event.Op&fsnotify.Write == fsnotify.Write {
                cm.reloadConfig(filename)
            }
        case err := <-watcher.Errors:
            log.Printf("Config watcher error: %v", err)
        }
    }
}
```

#### 1.3.2 Configuration Conflicts
**Symptoms:**
- Conflicting configuration values
- Unexpected behavior
- Validation errors

**Diagnosis:**
```bash
# Check configuration conflicts
./actor-core --check-conflicts

# Check configuration hierarchy
./actor-core --show-config-hierarchy

# Check environment variable overrides
./actor-core --show-env-overrides
```

**Solutions:**
```go
// 1. Implement configuration hierarchy
type ConfigurationHierarchy struct {
    Defaults    *Configuration
    FileConfig  *Configuration
    EnvConfig   *Configuration
    FinalConfig *Configuration
}

func (ch *ConfigurationHierarchy) Resolve() *Configuration {
    // Start with defaults
    config := ch.Defaults.Clone()
    
    // Override with file config
    if ch.FileConfig != nil {
        config.Merge(ch.FileConfig)
    }
    
    // Override with environment config
    if ch.EnvConfig != nil {
        config.Merge(ch.EnvConfig)
    }
    
    ch.FinalConfig = config
    return config
}

// 2. Implement configuration conflict detection
func (ch *ConfigurationHierarchy) DetectConflicts() []Conflict {
    var conflicts []Conflict
    
    // Check for conflicting values
    if ch.FileConfig != nil && ch.EnvConfig != nil {
        conflicts = append(conflicts, ch.compareConfigs(ch.FileConfig, ch.EnvConfig)...)
    }
    
    return conflicts
}

type Conflict struct {
    Field     string
    FileValue interface{}
    EnvValue  interface{}
    Message   string
}
```

## 2. Error Codes & Messages

### 2.1 Actor Core Errors

#### 2.1.1 Stat Calculation Errors
```go
// Error codes for stat calculations
const (
    ErrStatNotFound           = "STAT_NOT_FOUND"
    ErrStatCalculationFailed  = "STAT_CALCULATION_FAILED"
    ErrStatValueOutOfBounds   = "STAT_VALUE_OUT_OF_BOUNDS"
    ErrStatDependencyMissing  = "STAT_DEPENDENCY_MISSING"
    ErrStatFormulaInvalid     = "STAT_FORMULA_INVALID"
    ErrStatCacheMiss          = "STAT_CACHE_MISS"
    ErrStatCacheCorrupted     = "STAT_CACHE_CORRUPTED"
)

// Error messages
var ErrorMessages = map[string]string{
    ErrStatNotFound:           "Stat not found: %s",
    ErrStatCalculationFailed:  "Stat calculation failed: %s",
    ErrStatValueOutOfBounds:   "Stat value out of bounds: %s (value: %v, min: %v, max: %v)",
    ErrStatDependencyMissing:  "Stat dependency missing: %s",
    ErrStatFormulaInvalid:     "Stat formula invalid: %s",
    ErrStatCacheMiss:          "Stat cache miss: %s",
    ErrStatCacheCorrupted:     "Stat cache corrupted: %s",
}
```

#### 2.1.2 Flexible System Errors
```go
// Error codes for flexible systems
const (
    ErrFlexibleSystemNotFound     = "FLEXIBLE_SYSTEM_NOT_FOUND"
    ErrFlexibleSystemDisabled     = "FLEXIBLE_SYSTEM_DISABLED"
    ErrFlexibleSystemConfigInvalid = "FLEXIBLE_SYSTEM_CONFIG_INVALID"
    ErrFlexibleSystemLimitExceeded = "FLEXIBLE_SYSTEM_LIMIT_EXCEEDED"
    ErrFlexibleSystemConflict     = "FLEXIBLE_SYSTEM_CONFLICT"
)

// Error messages
var FlexibleSystemErrorMessages = map[string]string{
    ErrFlexibleSystemNotFound:      "Flexible system not found: %s",
    ErrFlexibleSystemDisabled:      "Flexible system disabled: %s",
    ErrFlexibleSystemConfigInvalid: "Flexible system config invalid: %s",
    ErrFlexibleSystemLimitExceeded: "Flexible system limit exceeded: %s (current: %d, max: %d)",
    ErrFlexibleSystemConflict:      "Flexible system conflict: %s",
}
```

### 2.2 Database Errors

#### 2.2.1 Connection Errors
```go
// Error codes for database connections
const (
    ErrDatabaseConnectionFailed = "DATABASE_CONNECTION_FAILED"
    ErrDatabaseConnectionLost   = "DATABASE_CONNECTION_LOST"
    ErrDatabaseConnectionPoolExhausted = "DATABASE_CONNECTION_POOL_EXHAUSTED"
    ErrDatabaseTimeout          = "DATABASE_TIMEOUT"
    ErrDatabaseQueryFailed      = "DATABASE_QUERY_FAILED"
)

// Error messages
var DatabaseErrorMessages = map[string]string{
    ErrDatabaseConnectionFailed:     "Database connection failed: %s",
    ErrDatabaseConnectionLost:       "Database connection lost: %s",
    ErrDatabaseConnectionPoolExhausted: "Database connection pool exhausted",
    ErrDatabaseTimeout:              "Database timeout: %s",
    ErrDatabaseQueryFailed:          "Database query failed: %s",
}
```

#### 2.2.2 Data Errors
```go
// Error codes for data operations
const (
    ErrDataNotFound        = "DATA_NOT_FOUND"
    ErrDataCorrupted       = "DATA_CORRUPTED"
    ErrDataValidationFailed = "DATA_VALIDATION_FAILED"
    ErrDataSerializationFailed = "DATA_SERIALIZATION_FAILED"
    ErrDataDeserializationFailed = "DATA_DESERIALIZATION_FAILED"
)

// Error messages
var DataErrorMessages = map[string]string{
    ErrDataNotFound:            "Data not found: %s",
    ErrDataCorrupted:           "Data corrupted: %s",
    ErrDataValidationFailed:    "Data validation failed: %s",
    ErrDataSerializationFailed: "Data serialization failed: %s",
    ErrDataDeserializationFailed: "Data deserialization failed: %s",
}
```

## 3. Debugging Techniques

### 3.1 Logging & Debugging

#### 3.1.1 Structured Logging
```go
// Structured logging for debugging
type DebugLogger struct {
    logger *logrus.Logger
    level  logrus.Level
}

func (dl *DebugLogger) LogStatCalculation(statID string, input map[string]float64, output float64, duration time.Duration) {
    dl.logger.WithFields(logrus.Fields{
        "stat_id":   statID,
        "input":     input,
        "output":    output,
        "duration":  duration,
        "timestamp": time.Now(),
    }).Debug("Stat calculation")
}

func (dl *DebugLogger) LogFlexibleSystemUsage(systemName string, component string, usage float64) {
    dl.logger.WithFields(logrus.Fields{
        "system_name": systemName,
        "component":   component,
        "usage":       usage,
        "timestamp":   time.Now(),
    }).Info("Flexible system usage")
}

func (dl *DebugLogger) LogError(err error, context map[string]interface{}) {
    dl.logger.WithFields(logrus.Fields{
        "error":     err.Error(),
        "context":   context,
        "timestamp": time.Now(),
    }).Error("Error occurred")
}
```

#### 3.1.2 Debug Endpoints
```go
// Debug endpoints for troubleshooting
func (ac *ActorCore) RegisterDebugEndpoints() {
    http.HandleFunc("/debug/stats", ac.debugStats)
    http.HandleFunc("/debug/performance", ac.debugPerformance)
    http.HandleFunc("/debug/memory", ac.debugMemory)
    http.HandleFunc("/debug/cache", ac.debugCache)
    http.HandleFunc("/debug/flexible-systems", ac.debugFlexibleSystems)
}

func (ac *ActorCore) debugStats(w http.ResponseWriter, r *http.Request) {
    stats := map[string]interface{}{
        "total_actors":     ac.GetActorCount(),
        "active_actors":    ac.GetActiveActorCount(),
        "stat_calculations": ac.GetStatCalculationCount(),
        "cache_hit_rate":   ac.GetCacheHitRate(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}

func (ac *ActorCore) debugPerformance(w http.ResponseWriter, r *http.Request) {
    performance := map[string]interface{}{
        "avg_calculation_time": ac.GetAverageCalculationTime(),
        "max_calculation_time": ac.GetMaxCalculationTime(),
        "memory_usage":        ac.GetMemoryUsage(),
        "goroutine_count":     runtime.NumGoroutine(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(performance)
}
```

### 3.2 Profiling & Monitoring

#### 3.2.1 CPU Profiling
```go
// CPU profiling for performance analysis
func (ac *ActorCore) EnableCPUProfiling() {
    go func() {
        log.Println("CPU profiling enabled on :6060")
        log.Println(http.ListenAndServe(":6060", nil))
    }()
    
    http.HandleFunc("/debug/pprof/profile", pprof.Profile)
    http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
    http.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

// Usage: go tool pprof http://localhost:6060/debug/pprof/profile
```

#### 3.2.2 Memory Profiling
```go
// Memory profiling for memory analysis
func (ac *ActorCore) EnableMemoryProfiling() {
    go func() {
        log.Println("Memory profiling enabled on :6060")
        log.Println(http.ListenAndServe(":6060", nil))
    }()
    
    http.HandleFunc("/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
    http.HandleFunc("/debug/pprof/allocs", pprof.Handler("allocs").ServeHTTP)
}

// Usage: go tool pprof http://localhost:6060/debug/pprof/heap
```

#### 3.2.3 Goroutine Profiling
```go
// Goroutine profiling for concurrency analysis
func (ac *ActorCore) EnableGoroutineProfiling() {
    go func() {
        log.Println("Goroutine profiling enabled on :6060")
        log.Println(http.ListenAndServe(":6060", nil))
    }()
    
    http.HandleFunc("/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
}

// Usage: go tool pprof http://localhost:6060/debug/pprof/goroutine
```

## 4. Recovery Procedures

### 4.1 Data Recovery

#### 4.1.1 Database Recovery
```bash
#!/bin/bash
# database-recovery.sh

echo "=== Database Recovery ==="

# 1. Check database status
echo "1. Checking database status..."
mongosh --eval "db.adminCommand('ping')"

# 2. Check for corruption
echo "2. Checking for corruption..."
mongosh --eval "db.runCommand({dbStats: 1})"

# 3. Repair database if needed
echo "3. Repairing database..."
mongod --repair --dbpath /var/lib/mongodb

# 4. Restore from backup if needed
echo "4. Restoring from backup..."
LATEST_BACKUP=$(ls -t /backups/mongodb/*.tar.gz | head -n1)
tar -xzf $LATEST_BACKUP -C /tmp/
mongorestore --db actor_core /tmp/actor_core/

# 5. Verify recovery
echo "5. Verifying recovery..."
mongosh --eval "db.actors.count()"

echo "Database recovery completed"
```

#### 4.1.2 Configuration Recovery
```bash
#!/bin/bash
# config-recovery.sh

echo "=== Configuration Recovery ==="

# 1. Check configuration files
echo "1. Checking configuration files..."
ls -la /app/config/

# 2. Validate configuration
echo "2. Validating configuration..."
./actor-core --validate-config

# 3. Restore from backup if needed
echo "3. Restoring from backup..."
LATEST_BACKUP=$(ls -t /backups/config/*.tar.gz | head -n1)
tar -xzf $LATEST_BACKUP -C /app/config/

# 4. Verify recovery
echo "4. Verifying recovery..."
./actor-core --validate-config

echo "Configuration recovery completed"
```

### 4.2 Service Recovery

#### 4.2.1 Application Recovery
```bash
#!/bin/bash
# application-recovery.sh

echo "=== Application Recovery ==="

# 1. Check application status
echo "1. Checking application status..."
kubectl get pods -l app=actor-core

# 2. Check application logs
echo "2. Checking application logs..."
kubectl logs -l app=actor-core --tail=100

# 3. Restart application if needed
echo "3. Restarting application..."
kubectl rollout restart deployment/actor-core

# 4. Wait for rollout
echo "4. Waiting for rollout..."
kubectl rollout status deployment/actor-core

# 5. Verify recovery
echo "5. Verifying recovery..."
curl -f http://localhost:8080/health

echo "Application recovery completed"
```

#### 4.2.2 Full System Recovery
```bash
#!/bin/bash
# full-system-recovery.sh

echo "=== Full System Recovery ==="

# 1. Stop all services
echo "1. Stopping all services..."
kubectl scale deployment actor-core --replicas=0
kubectl scale deployment mongodb --replicas=0
kubectl scale deployment redis --replicas=0

# 2. Wait for shutdown
echo "2. Waiting for shutdown..."
sleep 30

# 3. Start database
echo "3. Starting database..."
kubectl scale deployment mongodb --replicas=1
kubectl wait --for=condition=ready pod -l app=mongodb --timeout=300s

# 4. Start cache
echo "4. Starting cache..."
kubectl scale deployment redis --replicas=1
kubectl wait --for=condition=ready pod -l app=redis --timeout=300s

# 5. Start application
echo "5. Starting application..."
kubectl scale deployment actor-core --replicas=3
kubectl wait --for=condition=ready pod -l app=actor-core --timeout=300s

# 6. Verify recovery
echo "6. Verifying recovery..."
curl -f http://localhost:8080/health

echo "Full system recovery completed"
```

## 5. Preventive Measures

### 5.1 Health Checks

#### 5.1.1 Comprehensive Health Checks
```go
// Comprehensive health checks
type HealthChecker struct {
    checks map[string]HealthCheck
}

func (hc *HealthChecker) CheckAll() *HealthStatus {
    status := &HealthStatus{
        Status:    "healthy",
        Timestamp: time.Now(),
        Checks:    make(map[string]string),
    }
    
    for name, check := range hc.checks {
        if err := check.Check(); err != nil {
            status.Status = "unhealthy"
            status.Checks[name] = err.Error()
        } else {
            status.Checks[name] = "ok"
        }
    }
    
    return status
}

// Database health check
type DatabaseHealthCheck struct {
    client *mongo.Client
}

func (dhc *DatabaseHealthCheck) Check() error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    return dhc.client.Ping(ctx, nil)
}

// Cache health check
type CacheHealthCheck struct {
    client *redis.Client
}

func (chc *CacheHealthCheck) Check() error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    return chc.client.Ping(ctx).Err()
}
```

#### 5.1.2 Automated Health Monitoring
```yaml
# kubernetes/health-monitor.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: health-monitor
spec:
  replicas: 1
  selector:
    matchLabels:
      app: health-monitor
  template:
    metadata:
      labels:
        app: health-monitor
    spec:
      containers:
      - name: health-monitor
        image: health-monitor:latest
        command:
        - /bin/bash
        - -c
        - |
          while true; do
            # Check application health
            if ! curl -f http://actor-core:8080/health; then
              echo "Application health check failed"
              # Send alert
            fi
            
            # Check database health
            if ! mongosh --eval "db.adminCommand('ping')"; then
              echo "Database health check failed"
              # Send alert
            fi
            
            # Check cache health
            if ! redis-cli ping; then
              echo "Cache health check failed"
              # Send alert
            fi
            
            sleep 30
          done
```

### 5.2 Monitoring & Alerting

#### 5.2.1 Prometheus Alerts
```yaml
# monitoring/alerts.yml
groups:
- name: actor-core
  rules:
  - alert: HighCPUUsage
    expr: cpu_usage_percent > 80
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "High CPU usage detected"
      description: "CPU usage is above 80% for more than 5 minutes"
      
  - alert: HighMemoryUsage
    expr: memory_usage_percent > 90
    for: 5m
    labels:
      severity: critical
    annotations:
      summary: "High memory usage detected"
      description: "Memory usage is above 90% for more than 5 minutes"
      
  - alert: DatabaseConnectionFailed
    expr: mongodb_connections_failed > 0
    for: 1m
    labels:
      severity: critical
    annotations:
      summary: "Database connection failed"
      description: "Database connection failures detected"
      
  - alert: StatCalculationSlow
    expr: actor_core_stat_calculation_duration_seconds > 0.1
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "Slow stat calculations"
      description: "Stat calculations are taking longer than 100ms"
```

#### 5.2.2 Grafana Dashboards
```json
{
  "dashboard": {
    "title": "Actor Core v2.0 Monitoring",
    "panels": [
      {
        "title": "CPU Usage",
        "type": "graph",
        "targets": [
          {
            "expr": "cpu_usage_percent",
            "legendFormat": "CPU Usage %"
          }
        ]
      },
      {
        "title": "Memory Usage",
        "type": "graph",
        "targets": [
          {
            "expr": "memory_usage_bytes",
            "legendFormat": "Memory Usage"
          }
        ]
      },
      {
        "title": "Stat Calculations",
        "type": "graph",
        "targets": [
          {
            "expr": "actor_core_stat_calculation_duration_seconds",
            "legendFormat": "Calculation Duration"
          }
        ]
      },
      {
        "title": "Database Connections",
        "type": "graph",
        "targets": [
          {
            "expr": "mongodb_connections_active",
            "legendFormat": "Active Connections"
          }
        ]
      }
    ]
  }
}
```

---

*Tài liệu này cung cấp hướng dẫn chi tiết về troubleshooting cho Actor Core v2.0, giúp khắc phục sự cố một cách nhanh chóng và hiệu quả.*
