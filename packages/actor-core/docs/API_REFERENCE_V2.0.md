# Actor Core v2.0 - API Reference

## 📚 API Reference

### **Cache Interfaces**

#### **CacheInterface**
```go
type CacheInterface interface {
    Get(key string) (interface{}, error)
    Set(key string, value interface{}, ttl time.Duration) error
    Delete(key string) error
    Clear() error
    Exists(key string) bool
    GetTTL(key string) (time.Duration, error)
    SetTTL(key string, ttl time.Duration) error
}
```

#### **StateTrackerInterface**
```go
type StateTrackerInterface interface {
    TrackChange(change *StateChange) error
    GetChanges(entityType, entityID string) ([]*StateChange, error)
    GetChangesByTimeRange(start, end int64) ([]*StateChange, error)
    RollbackChange(changeID string) error
    ClearChanges(entityType, entityID string) error
}
```

### **Core Interfaces**

#### **StatProvider Interface**
```go
type StatProvider interface {
    GetPrimaryStats() *PrimaryCore
    GetDerivedStats() *DerivedStats
    GetFlexibleStats() *FlexibleStats
    GetVersion() int64
    GetUpdatedAt() int64
}
```

#### **StatConsumer Interface**
```go
type StatConsumer interface {
    UpdateStats(primary *PrimaryCore, derived *DerivedStats, flexible *FlexibleStats) error
    ValidateStats(stats interface{}) error
    GetStatsHash() string
}
```

#### **StatResolver Interface**
```go
type StatResolver interface {
    ResolveStats(primary *PrimaryCore) (*DerivedStats, error)
    AddFormula(formula Formula) error
    RemoveFormula(name string) error
    GetFormula(name string) (Formula, error)
    ValidateStats(stats map[string]float64) error
}
```

### **Cache Implementations**

#### **MemCache**
```go
type MemCache struct {
    // In-memory cache implementation
    data    map[string]*CacheEntry
    mutex   sync.RWMutex
    config  MemCacheConfig
    stats   *CacheStats
}

type CacheEntry struct {
    Value     interface{}
    ExpiresAt int64
    CreatedAt int64
    AccessCount int64
    LastAccess int64
}
```

#### **RedisCache**
```go
type RedisCache struct {
    client  redis.Client
    config  RedisConfig
    stats   *CacheStats
    mutex   sync.RWMutex
}
```

#### **StateTracker**
```go
type StateTracker struct {
    changes map[string]*StateChange
    mutex   sync.RWMutex
    config  StateTrackerConfig
}

type StateChange struct {
    ID          string      `json:"id"`
    EntityType  string      `json:"entity_type"`
    EntityID    string      `json:"entity_id"`
    ChangeType  string      `json:"change_type"`
    OldValue    interface{} `json:"old_value"`
    NewValue    interface{} `json:"new_value"`
    Timestamp   int64       `json:"timestamp"`
    UserID      string      `json:"user_id"`
    SystemID    string      `json:"system_id"`
}
```

### **Core Classes**

#### **PrimaryCore Class**
```go
type PrimaryCore struct {
    // Life Stats
    LifeSpan int64 `json:"life_span"`
    Age      int64 `json:"age"`
    
    // Physical Stats
    Vitality     int64 `json:"vitality"`
    Endurance    int64 `json:"endurance"`
    Constitution int64 `json:"constitution"`
    Intelligence int64 `json:"intelligence"`
    Wisdom       int64 `json:"wisdom"`
    Charisma     int64 `json:"charisma"`
    Willpower    int64 `json:"willpower"`
    Luck         int64 `json:"luck"`
    Fate         int64 `json:"fate"`
    Karma        int64 `json:"karma"`
    
    // Physical Attributes
    Strength   int64 `json:"strength"`
    Agility    int64 `json:"agility"`
    Personality int64 `json:"personality"`
    
    // Universal Cultivation Stats
    SpiritualEnergy    int64 `json:"spiritual_energy"`
    PhysicalEnergy     int64 `json:"physical_energy"`
    MentalEnergy       int64 `json:"mental_energy"`
    CultivationLevel   int64 `json:"cultivation_level"`
    BreakthroughPoints int64 `json:"breakthrough_points"`
    
    // Flexible Stats
    CustomPrimary map[string]int64 `json:"custom_primary"`
    
    // Metadata
    Version   int64 `json:"version"`
    CreatedAt int64 `json:"created_at"`
    UpdatedAt int64 `json:"updated_at"`
}
```

#### **DerivedStats Class**
```go
type DerivedStats struct {
    // Core Derived Stats
    HPMax       float64 `json:"hp_max"`
    Stamina     float64 `json:"stamina"`
    Speed       float64 `json:"speed"`
    Haste       float64 `json:"haste"`
    CritChance  float64 `json:"crit_chance"`
    CritMulti   float64 `json:"crit_multi"`
    MoveSpeed   float64 `json:"move_speed"`
    RegenHP     float64 `json:"regen_hp"`
    
    // Combat Stats
    Accuracy    float64 `json:"accuracy"`
    Penetration float64 `json:"penetration"`
    Lethality   float64 `json:"lethality"`
    Brutality   float64 `json:"brutality"`
    ArmorClass  float64 `json:"armor_class"`
    Evasion     float64 `json:"evasion"`
    BlockChance float64 `json:"block_chance"`
    ParryChance float64 `json:"parry_chance"`
    DodgeChance float64 `json:"dodge_chance"`
    
    // Energy Stats
    EnergyEfficiency float64            `json:"energy_efficiency"`
    EnergyCapacity   float64            `json:"energy_capacity"`
    EnergyDrain      float64            `json:"energy_drain"`
    RegenEnergies    map[string]float64 `json:"regen_energies"`
    
    // Damage and Defense
    Damages  map[string]float64    `json:"damages"`
    Defences map[string]DefenceValue `json:"defences"`
    
    // Amplifiers
    Amplifiers map[string]float64 `json:"amplifiers"`
    
    // Flexible Stats
    CustomDerived  map[string]float64          `json:"custom_derived"`
    SubSystemStats map[string]map[string]float64 `json:"sub_system_stats"`
    
    // Metadata
    Version   int64 `json:"version"`
    CreatedAt int64 `json:"created_at"`
    UpdatedAt int64 `json:"updated_at"`
}
```

#### **FlexibleStats Class**
```go
type FlexibleStats struct {
    CustomPrimary  map[string]int64            `json:"custom_primary"`
    CustomDerived  map[string]float64          `json:"custom_derived"`
    SubSystemStats map[string]map[string]float64 `json:"sub_system_stats"`
    
    Version   int64 `json:"version"`
    CreatedAt int64 `json:"created_at"`
    UpdatedAt int64 `json:"updated_at"`
}
```

#### **ConfigManager Class**
```go
type ConfigManager struct {
    configs   map[string]interface{}
    filePath  string
    version   int64
    createdAt int64
    updatedAt int64
    mutex     sync.RWMutex
}
```

#### **PerformanceMonitor Class**
```go
type PerformanceMonitor struct {
    metrics      map[string]*Metric
    alerts       map[string]*Alert
    thresholds   map[string]float64
    enabled      bool
    alertEnabled bool
    version      int64
    createdAt    int64
    updatedAt    int64
    mutex        sync.RWMutex
}
```

### **Cache Configuration**

#### **MemCacheConfig**
```go
type MemCacheConfig struct {
    MaxSize         string        `json:"max_size"`
    DefaultTTL      time.Duration `json:"default_ttl"`
    EvictionPolicy  string        `json:"eviction_policy"`
    EnableStats     bool          `json:"enable_statistics"`
    CleanupInterval time.Duration `json:"cleanup_interval"`
}
```

#### **RedisConfig**
```go
type RedisConfig struct {
    Host         string        `json:"host"`
    Port         int           `json:"port"`
    Password     string        `json:"password"`
    DB           int           `json:"db"`
    MaxRetries   int           `json:"max_retries"`
    DialTimeout  time.Duration `json:"dial_timeout"`
    ReadTimeout  time.Duration `json:"read_timeout"`
    WriteTimeout time.Duration `json:"write_timeout"`
    PoolSize     int           `json:"pool_size"`
    MinIdleConns int           `json:"min_idle_conns"`
}
```

#### **CacheStrategies**
```go
type CacheStrategies struct {
    PrimaryCore   CacheStrategy `json:"primary_core"`
    DerivedStats  CacheStrategy `json:"derived_stats"`
    FlexibleStats CacheStrategy `json:"flexible_stats"`
    ConfigManager CacheStrategy `json:"config_manager"`
}

type CacheStrategy struct {
    Strategy     string        `json:"strategy"`
    TTL          time.Duration `json:"ttl"`
    RefreshAhead bool          `json:"refresh_ahead"`
}
```

### **Usage Examples**

#### **Basic Cache Usage**
```go
// Initialize cache
memCache := cache.NewMemCache(memConfig)
redisCache := cache.NewRedisCache(redisConfig)

// Get from cache
value, err := memCache.Get("actor:123:primary_core")
if err != nil {
    // Cache miss, load from database
    value = loadFromDatabase("actor:123")
    memCache.Set("actor:123:primary_core", value, 5*time.Minute)
}

// Set cache with TTL
err := memCache.Set("actor:123:derived_stats", stats, 10*time.Minute)

// Delete from cache
err := memCache.Delete("actor:123:flexible_stats")
```

#### **State Change Tracking**
```go
// Track state change
change := &StateChange{
    EntityType: "primary_core",
    EntityID:   "actor:123",
    ChangeType: "update",
    OldValue:   oldStats,
    NewValue:   newStats,
    Timestamp:  time.Now().Unix(),
    UserID:     "user:456",
}

err := stateTracker.TrackChange(change)
```

#### **Cache Integration with Actor Core**
```go
// Initialize actor core with cache
actorCore := core.NewActorCore()
actorCore.SetCache(memCache, redisCache)
actorCore.SetStateTracker(stateTracker)

// Operations will automatically use cache
err := actorCore.UpdatePrimaryStats(newStats)
err := actorCore.UpdateDerivedStats(derivedStats)
err := actorCore.UpdateFlexibleStats(flexibleStats)
```

### **Error Handling**

#### **Cache Errors**
```go
type CacheError struct {
    Type    string `json:"type"`
    Message string `json:"message"`
    Key     string `json:"key"`
    Err     error  `json:"error"`
}

func (e *CacheError) Error() string {
    return fmt.Sprintf("cache error [%s]: %s (key: %s)", e.Type, e.Message, e.Key)
}
```

#### **State Tracking Errors**
```go
type StateTrackingError struct {
    Type      string `json:"type"`
    Message   string `json:"message"`
    ChangeID  string `json:"change_id"`
    EntityID  string `json:"entity_id"`
    Err       error  `json:"error"`
}

func (e *StateTrackingError) Error() string {
    return fmt.Sprintf("state tracking error [%s]: %s (change: %s)", e.Type, e.Message, e.ChangeID)
}
```

### **Performance Metrics**

#### **Cache Statistics**
```go
type CacheStats struct {
    Hits       int64   `json:"hits"`
    Misses     int64   `json:"misses"`
    HitRatio   float64 `json:"hit_ratio"`
    Size       int64   `json:"size"`
    MaxSize    int64   `json:"max_size"`
    Evictions  int64   `json:"evictions"`
    Errors     int64   `json:"errors"`
    LastReset  int64   `json:"last_reset"`
}
```

#### **State Tracking Statistics**
```go
type StateTrackingStats struct {
    TotalChanges    int64   `json:"total_changes"`
    ChangesPerHour  float64 `json:"changes_per_hour"`
    Rollbacks       int64   `json:"rollbacks"`
    Conflicts       int64   `json:"conflicts"`
    LastChange      int64   `json:"last_change"`
}
```

### **Configuration Examples**

#### **YAML Configuration**
```yaml
cache:
  mem_cache:
    max_size: "100MB"
    default_ttl: "300s"
    eviction_policy: "lru"
    enable_statistics: true
    cleanup_interval: "60s"
  
  redis:
    host: "localhost"
    port: 6379
    password: ""
    db: 0
    max_retries: 3
    dial_timeout: "5s"
    read_timeout: "3s"
    write_timeout: "3s"
    pool_size: 10
    min_idle_conns: 5
  
  strategies:
    primary_core:
      strategy: "write_through"
      ttl: "300s"
      refresh_ahead: true
    derived_stats:
      strategy: "write_behind"
      ttl: "600s"
      refresh_ahead: false
    flexible_stats:
      strategy: "write_around"
      ttl: "1800s"
      refresh_ahead: true
```

#### **JSON Configuration**
```json
{
  "cache": {
    "mem_cache": {
      "max_size": "100MB",
      "default_ttl": "300s",
      "eviction_policy": "lru",
      "enable_statistics": true,
      "cleanup_interval": "60s"
    },
    "redis": {
      "host": "localhost",
      "port": 6379,
      "password": "",
      "db": 0,
      "max_retries": 3,
      "dial_timeout": "5s",
      "read_timeout": "3s",
      "write_timeout": "3s",
      "pool_size": 10,
      "min_idle_conns": 5
    },
    "strategies": {
      "primary_core": {
        "strategy": "write_through",
        "ttl": "300s",
        "refresh_ahead": true
      },
      "derived_stats": {
        "strategy": "write_behind",
        "ttl": "600s",
        "refresh_ahead": false
      },
      "flexible_stats": {
        "strategy": "write_around",
        "ttl": "1800s",
        "refresh_ahead": true
      }
    }
  }
}
```
