package configuration

import "context"

// ConfigManager defines the interface for configuration management
type ConfigManager interface {
	// LoadFromFile loads configuration from a file
	LoadFromFile(filename string) error

	// LoadFromFileWithContext loads configuration from a file with context
	LoadFromFileWithContext(ctx context.Context, filename string) error

	// SaveToFile saves configuration to a file
	SaveToFile(filename string) error

	// SaveToFileWithContext saves configuration to a file with context
	SaveToFileWithContext(ctx context.Context, filename string) error

	// Validate validates the configuration
	Validate() []ValidationError

	// ValidateWithContext validates the configuration with context
	ValidateWithContext(ctx context.Context) []ValidationError

	// GetStat gets a stat definition
	GetStat(id string) (*StatDefinition, error)

	// GetStatWithContext gets a stat definition with context
	GetStatWithContext(ctx context.Context, id string) (*StatDefinition, error)

	// GetFormula gets a formula definition
	GetFormula(id string) (*FormulaDefinition, error)

	// GetFormulaWithContext gets a formula definition with context
	GetFormulaWithContext(ctx context.Context, id string) (*FormulaDefinition, error)

	// GetType gets a type definition
	GetType(id string) (*TypeDefinition, error)

	// GetTypeWithContext gets a type definition with context
	GetTypeWithContext(ctx context.Context, id string) (*TypeDefinition, error)

	// GetClamp gets a clamp definition
	GetClamp(id string) (*ClampDefinition, error)

	// GetClampWithContext gets a clamp definition with context
	GetClampWithContext(ctx context.Context, id string) (*ClampDefinition, error)

	// GetCategory gets a category definition
	GetCategory(id string) (*CategoryDefinition, error)

	// GetCategoryWithContext gets a category definition with context
	GetCategoryWithContext(ctx context.Context, id string) (*CategoryDefinition, error)

	// AddStat adds a stat definition
	AddStat(stat *StatDefinition) error

	// AddStatWithContext adds a stat definition with context
	AddStatWithContext(ctx context.Context, stat *StatDefinition) error

	// UpdateStat updates a stat definition
	UpdateStat(stat *StatDefinition) error

	// UpdateStatWithContext updates a stat definition with context
	UpdateStatWithContext(ctx context.Context, stat *StatDefinition) error

	// DeleteStat deletes a stat definition
	DeleteStat(id string) error

	// DeleteStatWithContext deletes a stat definition with context
	DeleteStatWithContext(ctx context.Context, id string) error

	// GetAllStats returns all stat definitions
	GetAllStats() map[string]*StatDefinition

	// GetAllFormulas returns all formula definitions
	GetAllFormulas() map[string]*FormulaDefinition

	// GetAllTypes returns all type definitions
	GetAllTypes() map[string]*TypeDefinition

	// GetAllClamps returns all clamp definitions
	GetAllClamps() map[string]*ClampDefinition

	// GetAllCategories returns all category definitions
	GetAllCategories() map[string]*CategoryDefinition

	// Reload reloads the configuration
	Reload() error

	// ReloadWithContext reloads the configuration with context
	ReloadWithContext(ctx context.Context) error

	// WatchFile watches a configuration file for changes
	WatchFile(filename string) error

	// StopWatching stops watching configuration files
	StopWatching() error

	// RegisterCallback registers a callback for configuration changes
	RegisterCallback(callback ConfigUpdateCallback) error

	// UnregisterCallback unregisters a configuration change callback
	UnregisterCallback(callbackID string) error
}

// StatDefinition represents a stat definition
type StatDefinition struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Category     string      `json:"category"`
	DataType     string      `json:"data_type"`
	DefaultValue interface{} `json:"default_value"`
	MinValue     interface{} `json:"min_value"`
	MaxValue     interface{} `json:"max_value"`
	Description  string      `json:"description"`
	IsActive     bool        `json:"is_active"`
	Dependencies []string    `json:"dependencies"`
	Tags         []string    `json:"tags"`
	CreatedAt    int64       `json:"created_at"`
	UpdatedAt    int64       `json:"updated_at"`
}

// FormulaDefinition represents a formula definition
type FormulaDefinition struct {
	ID           string             `json:"id"`
	StatID       string             `json:"stat_id"`
	Name         string             `json:"name"`
	Type         string             `json:"type"`
	Expression   string             `json:"expression"`
	Dependencies []string           `json:"dependencies"`
	Conditions   []FormulaCondition `json:"conditions"`
	Priority     int                `json:"priority"`
	IsActive     bool               `json:"is_active"`
	Description  string             `json:"description"`
	CreatedAt    int64              `json:"created_at"`
	UpdatedAt    int64              `json:"updated_at"`
}

// FormulaCondition represents a formula condition
type FormulaCondition struct {
	Condition string `json:"condition"`
	Formula   string `json:"formula"`
	Priority  int    `json:"priority"`
}

// TypeDefinition represents a type definition
type TypeDefinition struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Type         string                 `json:"type"`
	Category     string                 `json:"category"`
	Attributes   map[string]interface{} `json:"attributes"`
	Dependencies []string               `json:"dependencies"`
	IsActive     bool                   `json:"is_active"`
	Description  string                 `json:"description"`
	CreatedAt    int64                  `json:"created_at"`
	UpdatedAt    int64                  `json:"updated_at"`
}

// ClampDefinition represents a clamp definition
type ClampDefinition struct {
	ID          string      `json:"id"`
	StatID      string      `json:"stat_id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Value       interface{} `json:"value"`
	SoftCap     interface{} `json:"soft_cap"`
	SoftCapRate float64     `json:"soft_cap_rate"`
	IsActive    bool        `json:"is_active"`
	Description string      `json:"description"`
	CreatedAt   int64       `json:"created_at"`
	UpdatedAt   int64       `json:"updated_at"`
}

// CategoryDefinition represents a category definition
type CategoryDefinition struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Parent      string   `json:"parent"`
	Children    []string `json:"children"`
	Stats       []string `json:"stats"`
	Formulas    []string `json:"formulas"`
	Types       []string `json:"types"`
	Clamps      []string `json:"clamps"`
	IsActive    bool     `json:"is_active"`
	Description string   `json:"description"`
	CreatedAt   int64    `json:"created_at"`
	UpdatedAt   int64    `json:"updated_at"`
}

// ValidationError represents a validation error
type ValidationError struct {
	Field     string `json:"field"`
	Message   string `json:"message"`
	Severity  string `json:"severity"`
	Timestamp int64  `json:"timestamp"`
}

// ConfigUpdateCallback represents a callback for configuration updates
type ConfigUpdateCallback interface {
	OnConfigUpdated(updateType string, id string, data interface{})
	GetCallbackID() string
}
