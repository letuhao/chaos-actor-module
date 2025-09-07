package interfaces

// CombinerRegistry represents a registry for merge rules
type CombinerRegistry interface {
	// GetRule returns the merge rule for the given dimension
	GetRule(dimension string) (*MergeRule, error)

	// SetRule sets the merge rule for the given dimension
	SetRule(dimension string, rule *MergeRule) error

	// GetDimensions returns all dimensions with rules
	GetDimensions() []string

	// LoadFromConfig loads rules from configuration
	LoadFromConfig(config map[string]interface{}) error

	// Validate validates all rules
	Validate() error

	// Clear clears all rules
	Clear()

	// HasRule checks if a rule exists for the given dimension
	HasRule(dimension string) bool

	// RemoveRule removes a rule for the given dimension
	RemoveRule(dimension string)

	// Count returns the number of rules
	Count() int64
}

// CapLayerRegistry represents a registry for cap layer configuration
type CapLayerRegistry interface {
	// GetLayerOrder returns the processing order for layers
	GetLayerOrder() []string

	// GetAcrossLayerPolicy returns the across-layer policy
	GetAcrossLayerPolicy() string

	// SetLayerOrder sets the processing order for layers
	SetLayerOrder(order []string) error

	// SetAcrossLayerPolicy sets the across-layer policy
	SetAcrossLayerPolicy(policy string) error

	// LoadFromConfig loads configuration from config
	LoadFromConfig(config map[string]interface{}) error

	// Validate validates the configuration
	Validate() error

	// GetLayerIndex returns the index of a layer in the order
	GetLayerIndex(layer string) (int64, bool)

	// IsLayerInOrder checks if a layer is in the order
	IsLayerInOrder(layer string) bool

	// GetLayerCount returns the number of layers
	GetLayerCount() int64

	// Reset resets to default configuration
	Reset()
}

// MergeRule represents a rule for merging contributions
type MergeRule struct {
	// UsePipeline indicates whether to use the pipeline method
	UsePipeline bool `json:"use_pipeline"`

	// Operator is the aggregation operator (SUM, MAX, MIN, etc.)
	Operator string `json:"operator,omitempty"`

	// ClampDefault is the default clamp range
	ClampDefault Caps `json:"clamp_default"`
}

// IsValid checks if the merge rule is valid
func (mr *MergeRule) IsValid() bool {
	// Note: ClampDefault methods will be available when using actual types
	// For now, just check basic validation
	if !mr.UsePipeline && mr.Operator == "" {
		return false
	}

	return true
}

// GetDefaultClampRange returns the default clamp range
func (mr *MergeRule) GetDefaultClampRange() Caps {
	return mr.ClampDefault
}

// ShouldUsePipeline returns whether to use the pipeline method
func (mr *MergeRule) ShouldUsePipeline() bool {
	return mr.UsePipeline
}

// GetOperator returns the aggregation operator
func (mr *MergeRule) GetOperator() string {
	return mr.Operator
}

// PluginRegistry represents a registry for subsystems
type PluginRegistry interface {
	// Register registers a subsystem
	Register(subsystem Subsystem) error

	// Unregister unregisters a subsystem
	Unregister(systemID string) error

	// Get returns a subsystem by ID
	Get(systemID string) (Subsystem, bool)

	// GetAll returns all registered subsystems
	GetAll() []Subsystem

	// GetByPriority returns subsystems ordered by priority
	GetByPriority() []Subsystem

	// Clear clears all registered subsystems
	Clear()

	// Count returns the number of registered subsystems
	Count() int

	// HasSubsystem checks if a subsystem is registered
	HasSubsystem(systemID string) bool

	// GetSystemIDs returns all registered system IDs
	GetSystemIDs() []string

	// IsEmpty checks if the registry is empty
	IsEmpty() bool
}

// ConfigLoader represents a configuration loader
type ConfigLoader interface {
	// Load loads configuration from a file
	Load(filename string) (map[string]interface{}, error)

	// LoadFromBytes loads configuration from bytes
	LoadFromBytes(data []byte) (map[string]interface{}, error)

	// Save saves configuration to a file
	Save(filename string, config map[string]interface{}) error

	// Validate validates the configuration
	Validate(config map[string]interface{}) error
}

// Cache represents a cache interface
type Cache interface {
	// Get gets a value from the cache
	Get(key string) (interface{}, bool)

	// Set sets a value in the cache
	Set(key string, value interface{}, ttl string) error

	// Delete deletes a value from the cache
	Delete(key string) error

	// Clear clears all values from the cache
	Clear() error

	// GetStats returns cache statistics
	GetStats() *CacheStats
}

// CacheStats represents cache statistics
type CacheStats struct {
	// Hits is the number of cache hits
	Hits int64

	// Misses is the number of cache misses
	Misses int64

	// Size is the current cache size
	Size int64

	// MaxSize is the maximum cache size
	MaxSize int64

	// MemoryUsage is the current memory usage in bytes
	MemoryUsage int64
}

// GetHitRate returns the cache hit rate
func (cs *CacheStats) GetHitRate() float64 {
	total := cs.Hits + cs.Misses
	if total == 0 {
		return 0.0
	}
	return float64(cs.Hits) / float64(total)
}

// GetUsagePercentage returns the cache usage percentage
func (cs *CacheStats) GetUsagePercentage() float64 {
	if cs.MaxSize == 0 {
		return 0.0
	}
	return float64(cs.Size) / float64(cs.MaxSize) * 100.0
}
