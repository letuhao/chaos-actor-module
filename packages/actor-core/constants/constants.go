package constants

// System IDs
const (
	SystemIDCombat        = "combat"
	SystemIDMagic         = "magic"
	SystemIDCultivation   = "cultivation"
	SystemIDExperience    = "experience"
	SystemIDReputation    = "reputation"
	SystemIDGuild         = "guild"
	SystemIDTrading       = "trading"
	SystemIDWeather       = "weather"
	SystemIDLocation      = "location"
	SystemIDTime          = "time"
	SystemIDStealth       = "stealth"
	SystemIDPerception    = "perception"
	SystemIDLuck          = "luck"
)

// Dimension Names
const (
	// Primary Dimensions
	DimensionStrength     = "strength"
	DimensionVitality     = "vitality"
	DimensionDexterity    = "dexterity"
	DimensionIntelligence = "intelligence"
	DimensionSpirit       = "spirit"
	DimensionCharisma     = "charisma"

	// Derived Dimensions - Health & Resources
	DimensionHPMax        = "hp_max"
	DimensionMPMax        = "mp_max"
	DimensionStaminaMax   = "stamina_max"

	// Derived Dimensions - Combat
	DimensionAttackPower     = "attack_power"
	DimensionDefense         = "defense"
	DimensionMagicPower      = "magic_power"
	DimensionMagicResistance = "magic_resistance"

	// Derived Dimensions - Critical & Accuracy
	DimensionCritRate    = "crit_rate"
	DimensionCritDamage  = "crit_damage"
	DimensionAccuracy    = "accuracy"

	// Derived Dimensions - Speed & Movement
	DimensionMoveSpeed   = "move_speed"
	DimensionAttackSpeed = "attack_speed"
	DimensionCastSpeed   = "cast_speed"

	// Derived Dimensions - Resource Management
	DimensionCooldownReduction = "cooldown_reduction"
	DimensionManaEfficiency    = "mana_efficiency"
	DimensionEnergyEfficiency  = "energy_efficiency"

	// Derived Dimensions - Learning & Progression
	DimensionLearningRate      = "learning_rate"
	DimensionCultivationSpeed  = "cultivation_speed"
	DimensionBreakthroughSuccess = "breakthrough_success"

	// Meta/World Dimensions
	DimensionLifespanYears = "lifespan_years"
	DimensionPoiseRank     = "poise_rank"
	DimensionStealth       = "stealth"
	DimensionPerception    = "perception"
	DimensionLuck          = "luck"
)

// Context Types
const (
	ContextDamage      = "damage"
	ContextHealing     = "healing"
	ContextExperience  = "experience"
	ContextDropRate    = "drop_rate"
	ContextGold        = "gold"
	ContextReputation  = "reputation"
	ContextCultivation = "cultivation_speed"
	ContextBreakthrough = "breakthrough_chance"
	ContextSkillLearning = "skill_learning"
	ContextCraftingQuality = "crafting_quality"
)

// Error Codes
const (
	// Validation Errors
	ErrorCodeInvalidBucket        = "V001"
	ErrorCodeInvalidCapMode       = "V002"
	ErrorCodeInvalidLayerScope    = "V003"
	ErrorCodeMissingRequiredField = "V004"
	ErrorCodeValueOutOfRange      = "V005"
	ErrorCodeSchemaValidationFailed = "V006"

	// System Errors
	ErrorCodeSubsystemContributionFailed = "S001"
	ErrorCodeRegistryNotFound            = "S002"
	ErrorCodeInvalidRegistryFormat       = "S003"
	ErrorCodeCapsConflictDetected        = "S004"
	ErrorCodeLayerOrderInvalid           = "S005"
	ErrorCodeAcrossLayerPolicyInvalid    = "S006"

	// Performance Errors
	ErrorCodeOperationTimeout        = "P001"
	ErrorCodeMemoryLimitExceeded     = "P002"
	ErrorCodeCacheOperationFailed    = "P003"
	ErrorCodeTooManyConcurrentOps    = "P004"
	ErrorCodeResourceUnavailable     = "P005"
)

// Error Types
const (
	ErrorTypeValidation = "validation"
	ErrorTypeSystem     = "system"
	ErrorTypePerformance = "performance"
)

// Default Values
const (
	DefaultPriority           = 100
	DefaultVersion            = 1
	DefaultMaxConnections     = 100
	DefaultMaxIdleConnections = 10
	DefaultConnectionMaxLifetime = "1h"
	DefaultCacheTTL          = "1h"
	DefaultMaxMemory         = "2gb"
	DefaultPoolSize          = 100
	DefaultMinIdleConns      = 10
)

// Clamp Ranges
const (
	// Primary attributes
	MinStrength     = 0
	MaxStrength     = 999999
	MinVitality     = 0
	MaxVitality     = 999999
	MinDexterity    = 0
	MaxDexterity    = 999999
	MinIntelligence = 0
	MaxIntelligence = 999999
	MinSpirit       = 0
	MaxSpirit       = 999999
	MinCharisma     = 0
	MaxCharisma     = 999999

	// Health & resources
	MinHPMax      = 1
	MaxHPMax      = 2000000
	MinMPMax      = 1
	MaxMPMax      = 1000000
	MinStaminaMax = 1
	MaxStaminaMax = 500000

	// Combat attributes
	MinAttackPower     = 0
	MaxAttackPower     = 999999
	MinDefense         = 0
	MaxDefense         = 999999
	MinMagicPower      = 0
	MaxMagicPower      = 999999
	MinMagicResistance = 0
	MaxMagicResistance = 999999

	// Critical & accuracy
	MinCritRate   = 0.0
	MaxCritRate   = 1.0
	MinCritDamage = 1.0
	MaxCritDamage = 10.0
	MinAccuracy   = 0.0
	MaxAccuracy   = 1.0

	// Speed & movement
	MinMoveSpeed   = 0.0
	MaxMoveSpeed   = 12.0
	MinAttackSpeed = 0.1
	MaxAttackSpeed = 10.0
	MinCastSpeed   = 0.1
	MaxCastSpeed   = 10.0

	// Resource management
	MinCooldownReduction = 0.0
	MaxCooldownReduction = 0.5
	MinManaEfficiency    = 0.0
	MaxManaEfficiency    = 0.8
	MinEnergyEfficiency  = 0.0
	MaxEnergyEfficiency  = 0.8

	// Learning & progression
	MinLearningRate       = 0.1
	MaxLearningRate       = 5.0
	MinCultivationSpeed   = 0.1
	MaxCultivationSpeed   = 10.0
	MinBreakthroughSuccess = 0.0
	MaxBreakthroughSuccess = 1.0

	// Meta/World
	MinLifespanYears = 1
	MaxLifespanYears = 10000
	MinPoiseRank     = 0
	MaxPoiseRank     = 10
	MinStealth       = 0.0
	MaxStealth       = 1.0
	MinPerception    = 0.0
	MaxPerception    = 1.0
	MinLuck          = 0
	MaxLuck          = 999999
)

// Timeouts and Intervals
const (
	DefaultReadTimeout    = "30s"
	DefaultWriteTimeout   = "30s"
	DefaultIdleTimeout    = "120s"
	DefaultDialTimeout    = "5s"
	DefaultRequestTimeout = "10s"
	DefaultHealthCheckInterval = "30s"
	DefaultHealthCheckTimeout  = "10s"
	DefaultHealthCheckRetries  = 3
)

// Cache Keys
const (
	CacheKeyActorPrefix    = "actor:"
	CacheKeySnapshotPrefix = "snapshot:"
	CacheKeyCapsPrefix     = "caps:"
	CacheKeyRegistryPrefix = "registry:"
)

// Log Levels
const (
	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
	LogLevelFatal = "fatal"
)

// Log Formats
const (
	LogFormatJSON = "json"
	LogFormatText = "text"
)

// Monitoring
const (
	MetricsPathPrometheus = "/metrics"
	HealthPathHealth      = "/health"
	HealthPathReady       = "/ready"
	TracePathJaeger       = "/api/traces"
)

// Database
const (
	DatabaseSSLModeDisable = "disable"
	DatabaseSSLModeRequire = "require"
	DatabaseSSLModeVerify  = "verify-full"
)

// Cache Policies
const (
	CachePolicyAllKeysLRU     = "allkeys-lru"
	CachePolicyAllKeysLFU     = "allkeys-lfu"
	CachePolicyVolatileLRU    = "volatile-lru"
	CachePolicyVolatileLFU    = "volatile-lfu"
	CachePolicyVolatileTTL    = "volatile-ttl"
	CachePolicyNoEviction     = "noeviction"
)

// Deployment Strategies
const (
	DeploymentStrategyBlueGreen = "blue-green"
	DeploymentStrategyRolling  = "rolling"
	DeploymentStrategyCanary   = "canary"
)

// Health Status
const (
	HealthStatusHealthy   = "healthy"
	HealthStatusUnhealthy = "unhealthy"
	HealthStatusDegraded  = "degraded"
)

// API Levels
const (
	APILegacy    = 1
	APIV1        = 2
	APIV2        = 3
	APICurrent   = 4
	APIMinimum   = 2
)

// Version Information
const (
	VersionMajor = 3
	VersionMinor = 0
	VersionPatch = 0
	VersionBuild = "dev"
)
