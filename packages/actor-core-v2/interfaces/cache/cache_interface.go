package cache

import "time"

// CacheInterface defines the interface for cache operations
type CacheInterface interface {
	// Basic operations
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, ttl time.Duration) error
	Delete(key string) error
	Clear() error
	Exists(key string) bool

	// TTL operations
	GetTTL(key string) (time.Duration, error)
	SetTTL(key string, ttl time.Duration) error

	// Statistics
	GetStats() *CacheStats
	ResetStats() error

	// Health check
	Health() error
}

// StateTrackerInterface defines the interface for state change tracking
type StateTrackerInterface interface {
	// Change tracking
	TrackChange(change *StateChange) error
	GetChanges(entityType, entityID string) ([]*StateChange, error)
	GetChangesByTimeRange(start, end int64) ([]*StateChange, error)
	GetChangesByUser(userID string) ([]*StateChange, error)

	// Change management
	RollbackChange(changeID string) error
	ClearChanges(entityType, entityID string) error
	ClearChangesByTimeRange(start, end int64) error

	// Statistics
	GetStats() *StateTrackingStats
	ResetStats() error
}

// CacheStats represents cache performance statistics
type CacheStats struct {
	Hits      int64   `json:"hits"`
	Misses    int64   `json:"misses"`
	HitRatio  float64 `json:"hit_ratio"`
	Size      int64   `json:"size"`
	MaxSize   int64   `json:"max_size"`
	Evictions int64   `json:"evictions"`
	Errors    int64   `json:"errors"`
	LastReset int64   `json:"last_reset"`
}

// StateChange represents a state change event
type StateChange struct {
	ID          string                 `json:"id"`
	EntityType  string                 `json:"entity_type"`
	EntityID    string                 `json:"entity_id"`
	ChangeType  string                 `json:"change_type"` // create, update, delete
	OldValue    interface{}            `json:"old_value"`
	NewValue    interface{}            `json:"new_value"`
	Timestamp   int64                  `json:"timestamp"`
	UserID      string                 `json:"user_id"`
	SystemID    string                 `json:"system_id"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// StateTrackingStats represents state tracking statistics
type StateTrackingStats struct {
	TotalChanges   int64   `json:"total_changes"`
	ChangesPerHour float64 `json:"changes_per_hour"`
	Rollbacks      int64   `json:"rollbacks"`
	Conflicts      int64   `json:"conflicts"`
	LastChange     int64   `json:"last_change"`
	ActiveChanges  int64   `json:"active_changes"`
	ErrorRate      float64 `json:"error_rate"`
}

// CacheEntry represents a cache entry
type CacheEntry struct {
	Value       interface{} `json:"value"`
	ExpiresAt   int64       `json:"expires_at"`
	CreatedAt   int64       `json:"created_at"`
	AccessCount int64       `json:"access_count"`
	LastAccess  int64       `json:"last_access"`
	Size        int64       `json:"size"`
}

// CacheStrategy defines cache strategy configuration
type CacheStrategy struct {
	Strategy       string        `json:"strategy"` // write_through, write_behind, write_around
	TTL            time.Duration `json:"ttl"`
	RefreshAhead   bool          `json:"refresh_ahead"`
	MaxSize        int64         `json:"max_size"`
	EvictionPolicy string        `json:"eviction_policy"` // lru, lfu, ttl
}

// CacheConfig represents cache configuration
type CacheConfig struct {
	MemCache     MemCacheConfig     `json:"mem_cache"`
	Redis        RedisConfig        `json:"redis"`
	Strategies   CacheStrategies    `json:"strategies"`
	Invalidation InvalidationConfig `json:"invalidation"`
}

// MemCacheConfig represents in-memory cache configuration
type MemCacheConfig struct {
	MaxSize         string        `json:"max_size"`
	DefaultTTL      time.Duration `json:"default_ttl"`
	EvictionPolicy  string        `json:"eviction_policy"`
	EnableStats     bool          `json:"enable_statistics"`
	CleanupInterval time.Duration `json:"cleanup_interval"`
	MaxEntries      int64         `json:"max_entries"`
}

// RedisConfig represents Redis cache configuration
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
	EnableTLS    bool          `json:"enable_tls"`
}

// CacheStrategies represents cache strategies for different components
type CacheStrategies struct {
	PrimaryCore   CacheStrategy `json:"primary_core"`
	DerivedStats  CacheStrategy `json:"derived_stats"`
	FlexibleStats CacheStrategy `json:"flexible_stats"`
	ConfigManager CacheStrategy `json:"config_manager"`
	Performance   CacheStrategy `json:"performance"`
}

// InvalidationConfig represents cache invalidation configuration
type InvalidationConfig struct {
	EnableTimeBased       bool          `json:"enable_time_based"`
	EnableEventBased      bool          `json:"enable_event_based"`
	EnableDependencyBased bool          `json:"enable_dependency_based"`
	DefaultTTL            time.Duration `json:"default_ttl"`
	CleanupInterval       time.Duration `json:"cleanup_interval"`
	MaxHistorySize        int64         `json:"max_history_size"`
}

// StateTrackerConfig represents state tracker configuration
type StateTrackerConfig struct {
	EnableTracking          bool          `json:"enable_tracking"`
	MaxHistorySize          int64         `json:"max_history_size"`
	RetentionPeriod         time.Duration `json:"retention_period"`
	EnableRollback          bool          `json:"enable_rollback"`
	EnableConflictDetection bool          `json:"enable_conflict_detection"`
	CleanupInterval         time.Duration `json:"cleanup_interval"`
}
