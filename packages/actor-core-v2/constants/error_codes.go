package constants

// Error Codes
const (
	// General Errors
	ErrInvalidInput       = "INVALID_INPUT"
	ErrNotFound           = "NOT_FOUND"
	ErrAlreadyExists      = "ALREADY_EXISTS"
	ErrPermissionDenied   = "PERMISSION_DENIED"
	ErrTimeout            = "TIMEOUT"
	ErrInternalError      = "INTERNAL_ERROR"
	ErrConfigurationError = "CONFIGURATION_ERROR"
	ErrValidationError    = "VALIDATION_ERROR"

	// Stat Errors
	ErrStatNotFound          = "STAT_NOT_FOUND"
	ErrStatCalculationFailed = "STAT_CALCULATION_FAILED"
	ErrStatValueOutOfBounds  = "STAT_VALUE_OUT_OF_BOUNDS"
	ErrStatDependencyMissing = "STAT_DEPENDENCY_MISSING"
	ErrStatFormulaInvalid    = "STAT_FORMULA_INVALID"
	ErrStatCacheMiss         = "STAT_CACHE_MISS"
	ErrStatCacheCorrupted    = "STAT_CACHE_CORRUPTED"
	ErrStatNameInvalid       = "STAT_NAME_INVALID"
	ErrStatTypeInvalid       = "STAT_TYPE_INVALID"
	ErrStatCategoryInvalid   = "STAT_CATEGORY_INVALID"

	// Flexible System Errors
	ErrFlexibleSystemNotFound      = "FLEXIBLE_SYSTEM_NOT_FOUND"
	ErrFlexibleSystemDisabled      = "FLEXIBLE_SYSTEM_DISABLED"
	ErrFlexibleSystemConfigInvalid = "FLEXIBLE_SYSTEM_CONFIG_INVALID"
	ErrFlexibleSystemLimitExceeded = "FLEXIBLE_SYSTEM_LIMIT_EXCEEDED"
	ErrFlexibleSystemConflict      = "FLEXIBLE_SYSTEM_CONFLICT"
	ErrFlexibleSystemNotSupported  = "FLEXIBLE_SYSTEM_NOT_SUPPORTED"

	// Configuration Errors
	ErrConfigNotFound         = "CONFIG_NOT_FOUND"
	ErrConfigInvalid          = "CONFIG_INVALID"
	ErrConfigLoadFailed       = "CONFIG_LOAD_FAILED"
	ErrConfigSaveFailed       = "CONFIG_SAVE_FAILED"
	ErrConfigValidationFailed = "CONFIG_VALIDATION_FAILED"
	ErrConfigConflict         = "CONFIG_CONFLICT"
	ErrConfigHotReloadFailed  = "CONFIG_HOT_RELOAD_FAILED"

	// Formula Errors
	ErrFormulaNotFound           = "FORMULA_NOT_FOUND"
	ErrFormulaInvalid            = "FORMULA_INVALID"
	ErrFormulaCompilationFailed  = "FORMULA_COMPILATION_FAILED"
	ErrFormulaExecutionFailed    = "FORMULA_EXECUTION_FAILED"
	ErrFormulaDependencyMissing  = "FORMULA_DEPENDENCY_MISSING"
	ErrFormulaCircularDependency = "FORMULA_CIRCULAR_DEPENDENCY"

	// Database Errors
	ErrDatabaseConnectionFailed        = "DATABASE_CONNECTION_FAILED"
	ErrDatabaseConnectionLost          = "DATABASE_CONNECTION_LOST"
	ErrDatabaseConnectionPoolExhausted = "DATABASE_CONNECTION_POOL_EXHAUSTED"
	ErrDatabaseTimeout                 = "DATABASE_TIMEOUT"
	ErrDatabaseQueryFailed             = "DATABASE_QUERY_FAILED"
	ErrDatabaseTransactionFailed       = "DATABASE_TRANSACTION_FAILED"

	// Data Errors
	ErrDataNotFound              = "DATA_NOT_FOUND"
	ErrDataCorrupted             = "DATA_CORRUPTED"
	ErrDataValidationFailed      = "DATA_VALIDATION_FAILED"
	ErrDataSerializationFailed   = "DATA_SERIALIZATION_FAILED"
	ErrDataDeserializationFailed = "DATA_DESERIALIZATION_FAILED"
	ErrDataIntegrityViolation    = "DATA_INTEGRITY_VIOLATION"

	// Cache Errors
	ErrCacheNotFound        = "CACHE_NOT_FOUND"
	ErrCacheExpired         = "CACHE_EXPIRED"
	ErrCacheCorrupted       = "CACHE_CORRUPTED"
	ErrCacheLimitExceeded   = "CACHE_LIMIT_EXCEEDED"
	ErrCacheOperationFailed = "CACHE_OPERATION_FAILED"

	// Performance Errors
	ErrPerformanceTimeout       = "PERFORMANCE_TIMEOUT"
	ErrMemoryLimitExceeded      = "MEMORY_LIMIT_EXCEEDED"
	ErrCPUUsageExceeded         = "CPU_USAGE_EXCEEDED"
	ErrConcurrencyLimitExceeded = "CONCURRENCY_LIMIT_EXCEEDED"

	// Security Errors
	ErrSecurityViolation    = "SECURITY_VIOLATION"
	ErrAuthenticationFailed = "AUTHENTICATION_FAILED"
	ErrAuthorizationFailed  = "AUTHORIZATION_FAILED"
	ErrAccessDenied         = "ACCESS_DENIED"
	ErrInvalidToken         = "INVALID_TOKEN"
	ErrTokenExpired         = "TOKEN_EXPIRED"
	ErrEncryptionFailed     = "ENCRYPTION_FAILED"
	ErrDecryptionFailed     = "DECRYPTION_FAILED"

	// Validation Errors
	ErrValidationFailed     = "VALIDATION_FAILED"
	ErrRequiredFieldMissing = "REQUIRED_FIELD_MISSING"
	ErrInvalidValue         = "INVALID_VALUE"
	ErrValueOutOfRange      = "VALUE_OUT_OF_RANGE"
	ErrInvalidFormat        = "INVALID_FORMAT"
	ErrConstraintViolation  = "CONSTRAINT_VIOLATION"

	// System Errors
	ErrSystemOverload    = "SYSTEM_OVERLOAD"
	ErrSystemMaintenance = "SYSTEM_MAINTENANCE"
	ErrSystemUnavailable = "SYSTEM_UNAVAILABLE"
	ErrSystemShutdown    = "SYSTEM_SHUTDOWN"
	ErrSystemRestart     = "SYSTEM_RESTART"
)

// Error Messages
var ErrorMessages = map[string]string{
	// General Errors
	ErrInvalidInput:       "Invalid input provided",
	ErrNotFound:           "Resource not found",
	ErrAlreadyExists:      "Resource already exists",
	ErrPermissionDenied:   "Permission denied",
	ErrTimeout:            "Operation timed out",
	ErrInternalError:      "Internal server error",
	ErrConfigurationError: "Configuration error",
	ErrValidationError:    "Validation error",

	// Stat Errors
	ErrStatNotFound:          "Stat not found: %s",
	ErrStatCalculationFailed: "Stat calculation failed: %s",
	ErrStatValueOutOfBounds:  "Stat value out of bounds: %s (value: %v, min: %v, max: %v)",
	ErrStatDependencyMissing: "Stat dependency missing: %s",
	ErrStatFormulaInvalid:    "Stat formula invalid: %s",
	ErrStatCacheMiss:         "Stat cache miss: %s",
	ErrStatCacheCorrupted:    "Stat cache corrupted: %s",
	ErrStatNameInvalid:       "Stat name invalid: %s",
	ErrStatTypeInvalid:       "Stat type invalid: %s",
	ErrStatCategoryInvalid:   "Stat category invalid: %s",

	// Flexible System Errors
	ErrFlexibleSystemNotFound:      "Flexible system not found: %s",
	ErrFlexibleSystemDisabled:      "Flexible system disabled: %s",
	ErrFlexibleSystemConfigInvalid: "Flexible system config invalid: %s",
	ErrFlexibleSystemLimitExceeded: "Flexible system limit exceeded: %s (current: %d, max: %d)",
	ErrFlexibleSystemConflict:      "Flexible system conflict: %s",
	ErrFlexibleSystemNotSupported:  "Flexible system not supported: %s",

	// Configuration Errors
	ErrConfigNotFound:         "Configuration not found: %s",
	ErrConfigInvalid:          "Configuration invalid: %s",
	ErrConfigLoadFailed:       "Configuration load failed: %s",
	ErrConfigSaveFailed:       "Configuration save failed: %s",
	ErrConfigValidationFailed: "Configuration validation failed: %s",
	ErrConfigConflict:         "Configuration conflict: %s",
	ErrConfigHotReloadFailed:  "Configuration hot reload failed: %s",

	// Formula Errors
	ErrFormulaNotFound:           "Formula not found: %s",
	ErrFormulaInvalid:            "Formula invalid: %s",
	ErrFormulaCompilationFailed:  "Formula compilation failed: %s",
	ErrFormulaExecutionFailed:    "Formula execution failed: %s",
	ErrFormulaDependencyMissing:  "Formula dependency missing: %s",
	ErrFormulaCircularDependency: "Formula circular dependency: %s",

	// Database Errors
	ErrDatabaseConnectionFailed:        "Database connection failed: %s",
	ErrDatabaseConnectionLost:          "Database connection lost: %s",
	ErrDatabaseConnectionPoolExhausted: "Database connection pool exhausted",
	ErrDatabaseTimeout:                 "Database timeout: %s",
	ErrDatabaseQueryFailed:             "Database query failed: %s",
	ErrDatabaseTransactionFailed:       "Database transaction failed: %s",

	// Data Errors
	ErrDataNotFound:              "Data not found: %s",
	ErrDataCorrupted:             "Data corrupted: %s",
	ErrDataValidationFailed:      "Data validation failed: %s",
	ErrDataSerializationFailed:   "Data serialization failed: %s",
	ErrDataDeserializationFailed: "Data deserialization failed: %s",
	ErrDataIntegrityViolation:    "Data integrity violation: %s",

	// Cache Errors
	ErrCacheNotFound:        "Cache not found: %s",
	ErrCacheExpired:         "Cache expired: %s",
	ErrCacheCorrupted:       "Cache corrupted: %s",
	ErrCacheLimitExceeded:   "Cache limit exceeded: %s",
	ErrCacheOperationFailed: "Cache operation failed: %s",

	// Performance Errors
	ErrPerformanceTimeout:       "Performance timeout: %s",
	ErrMemoryLimitExceeded:      "Memory limit exceeded: %s",
	ErrCPUUsageExceeded:         "CPU usage exceeded: %s",
	ErrConcurrencyLimitExceeded: "Concurrency limit exceeded: %s",

	// Security Errors
	ErrSecurityViolation:    "Security violation: %s",
	ErrAuthenticationFailed: "Authentication failed: %s",
	ErrAuthorizationFailed:  "Authorization failed: %s",
	ErrAccessDenied:         "Access denied: %s",
	ErrInvalidToken:         "Invalid token: %s",
	ErrTokenExpired:         "Token expired: %s",
	ErrEncryptionFailed:     "Encryption failed: %s",
	ErrDecryptionFailed:     "Decryption failed: %s",

	// Validation Errors
	ErrValidationFailed:     "Validation failed: %s",
	ErrRequiredFieldMissing: "Required field missing: %s",
	ErrInvalidValue:         "Invalid value: %s",
	ErrValueOutOfRange:      "Value out of range: %s",
	ErrInvalidFormat:        "Invalid format: %s",
	ErrConstraintViolation:  "Constraint violation: %s",

	// System Errors
	ErrSystemOverload:    "System overload: %s",
	ErrSystemMaintenance: "System maintenance: %s",
	ErrSystemUnavailable: "System unavailable: %s",
	ErrSystemShutdown:    "System shutdown: %s",
	ErrSystemRestart:     "System restart: %s",
}

// Error Severity Levels
const (
	SeverityInfo     = "info"
	SeverityWarning  = "warning"
	SeverityError    = "error"
	SeverityCritical = "critical"
)

// Error Categories
const (
	CategoryGeneral       = "general"
	CategoryStat          = "stat"
	CategoryFlexible      = "flexible"
	CategoryConfiguration = "configuration"
	CategoryFormula       = "formula"
	CategoryDatabase      = "database"
	CategoryData          = "data"
	CategoryCache         = "cache"
	CategoryPerformance   = "performance"
	CategorySecurity      = "security"
	CategoryValidation    = "validation"
	CategorySystem        = "system"
)

// Error Categories by Code
var ErrorCategories = map[string]string{
	// General Errors
	ErrInvalidInput:       CategoryGeneral,
	ErrNotFound:           CategoryGeneral,
	ErrAlreadyExists:      CategoryGeneral,
	ErrPermissionDenied:   CategoryGeneral,
	ErrTimeout:            CategoryGeneral,
	ErrInternalError:      CategoryGeneral,
	ErrConfigurationError: CategoryGeneral,
	ErrValidationError:    CategoryGeneral,

	// Stat Errors
	ErrStatNotFound:          CategoryStat,
	ErrStatCalculationFailed: CategoryStat,
	ErrStatValueOutOfBounds:  CategoryStat,
	ErrStatDependencyMissing: CategoryStat,
	ErrStatFormulaInvalid:    CategoryStat,
	ErrStatCacheMiss:         CategoryStat,
	ErrStatCacheCorrupted:    CategoryStat,
	ErrStatNameInvalid:       CategoryStat,
	ErrStatTypeInvalid:       CategoryStat,
	ErrStatCategoryInvalid:   CategoryStat,

	// Flexible System Errors
	ErrFlexibleSystemNotFound:      CategoryFlexible,
	ErrFlexibleSystemDisabled:      CategoryFlexible,
	ErrFlexibleSystemConfigInvalid: CategoryFlexible,
	ErrFlexibleSystemLimitExceeded: CategoryFlexible,
	ErrFlexibleSystemConflict:      CategoryFlexible,
	ErrFlexibleSystemNotSupported:  CategoryFlexible,

	// Configuration Errors
	ErrConfigNotFound:         CategoryConfiguration,
	ErrConfigInvalid:          CategoryConfiguration,
	ErrConfigLoadFailed:       CategoryConfiguration,
	ErrConfigSaveFailed:       CategoryConfiguration,
	ErrConfigValidationFailed: CategoryConfiguration,
	ErrConfigConflict:         CategoryConfiguration,
	ErrConfigHotReloadFailed:  CategoryConfiguration,

	// Formula Errors
	ErrFormulaNotFound:           CategoryFormula,
	ErrFormulaInvalid:            CategoryFormula,
	ErrFormulaCompilationFailed:  CategoryFormula,
	ErrFormulaExecutionFailed:    CategoryFormula,
	ErrFormulaDependencyMissing:  CategoryFormula,
	ErrFormulaCircularDependency: CategoryFormula,

	// Database Errors
	ErrDatabaseConnectionFailed:        CategoryDatabase,
	ErrDatabaseConnectionLost:          CategoryDatabase,
	ErrDatabaseConnectionPoolExhausted: CategoryDatabase,
	ErrDatabaseTimeout:                 CategoryDatabase,
	ErrDatabaseQueryFailed:             CategoryDatabase,
	ErrDatabaseTransactionFailed:       CategoryDatabase,

	// Data Errors
	ErrDataNotFound:              CategoryData,
	ErrDataCorrupted:             CategoryData,
	ErrDataValidationFailed:      CategoryData,
	ErrDataSerializationFailed:   CategoryData,
	ErrDataDeserializationFailed: CategoryData,
	ErrDataIntegrityViolation:    CategoryData,

	// Cache Errors
	ErrCacheNotFound:        CategoryCache,
	ErrCacheExpired:         CategoryCache,
	ErrCacheCorrupted:       CategoryCache,
	ErrCacheLimitExceeded:   CategoryCache,
	ErrCacheOperationFailed: CategoryCache,

	// Performance Errors
	ErrPerformanceTimeout:       CategoryPerformance,
	ErrMemoryLimitExceeded:      CategoryPerformance,
	ErrCPUUsageExceeded:         CategoryPerformance,
	ErrConcurrencyLimitExceeded: CategoryPerformance,

	// Security Errors
	ErrSecurityViolation:    CategorySecurity,
	ErrAuthenticationFailed: CategorySecurity,
	ErrAuthorizationFailed:  CategorySecurity,
	ErrAccessDenied:         CategorySecurity,
	ErrInvalidToken:         CategorySecurity,
	ErrTokenExpired:         CategorySecurity,
	ErrEncryptionFailed:     CategorySecurity,
	ErrDecryptionFailed:     CategorySecurity,

	// Validation Errors
	ErrValidationFailed:     CategoryValidation,
	ErrRequiredFieldMissing: CategoryValidation,
	ErrInvalidValue:         CategoryValidation,
	ErrValueOutOfRange:      CategoryValidation,
	ErrInvalidFormat:        CategoryValidation,
	ErrConstraintViolation:  CategoryValidation,

	// System Errors
	ErrSystemOverload:    CategorySystem,
	ErrSystemMaintenance: CategorySystem,
	ErrSystemUnavailable: CategorySystem,
	ErrSystemShutdown:    CategorySystem,
	ErrSystemRestart:     CategorySystem,
}
