package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// ConfigManager implements the ConfigManager interface
type ConfigManager struct {
	mu        sync.RWMutex
	configs   map[string]interface{}
	filePath  string
	version   int64
	createdAt int64
	updatedAt int64
}

// ConfigData represents a configuration data structure
type ConfigData struct {
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
	Type      string      `json:"type"`
	Category  string      `json:"category"`
	UpdatedAt int64       `json:"updated_at"`
}

// ConfigFile represents the structure of a configuration file
type ConfigFile struct {
	Version   int64                 `json:"version"`
	CreatedAt int64                 `json:"created_at"`
	UpdatedAt int64                 `json:"updated_at"`
	Configs   map[string]ConfigData `json:"configs"`
}

// NewConfigManager creates a new ConfigManager instance
func NewConfigManager() *ConfigManager {
	now := time.Now().Unix()
	return &ConfigManager{
		configs:   make(map[string]interface{}),
		version:   1,
		createdAt: now,
		updatedAt: now,
	}
}

// NewConfigManagerWithFile creates a new ConfigManager with a specific file path
func NewConfigManagerWithFile(filePath string) *ConfigManager {
	cm := NewConfigManager()
	cm.filePath = filePath
	return cm
}

// SetConfig sets a configuration value
func (cm *ConfigManager) SetConfig(key string, value interface{}) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if key == "" {
		return fmt.Errorf("config key cannot be empty")
	}

	cm.configs[key] = value
	cm.updatedAt = time.Now().Unix()
	cm.version++

	return nil
}

// GetConfig gets a configuration value
func (cm *ConfigManager) GetConfig(key string) (interface{}, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	value, exists := cm.configs[key]
	if !exists {
		return nil, fmt.Errorf("config key '%s' not found", key)
	}

	return value, nil
}

// GetConfigString gets a configuration value as string
func (cm *ConfigManager) GetConfigString(key string) (string, error) {
	value, err := cm.GetConfig(key)
	if err != nil {
		return "", err
	}

	str, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("config key '%s' is not a string", key)
	}

	return str, nil
}

// GetConfigInt gets a configuration value as int
func (cm *ConfigManager) GetConfigInt(key string) (int, error) {
	value, err := cm.GetConfig(key)
	if err != nil {
		return 0, err
	}

	switch v := value.(type) {
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case float64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("config key '%s' is not a number", key)
	}
}

// GetConfigInt64 gets a configuration value as int64
func (cm *ConfigManager) GetConfigInt64(key string) (int64, error) {
	value, err := cm.GetConfig(key)
	if err != nil {
		return 0, err
	}

	switch v := value.(type) {
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case float64:
		return int64(v), nil
	default:
		return 0, fmt.Errorf("config key '%s' is not a number", key)
	}
}

// GetConfigFloat64 gets a configuration value as float64
func (cm *ConfigManager) GetConfigFloat64(key string) (float64, error) {
	value, err := cm.GetConfig(key)
	if err != nil {
		return 0, err
	}

	switch v := value.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	default:
		return 0, fmt.Errorf("config key '%s' is not a number", key)
	}
}

// GetConfigBool gets a configuration value as bool
func (cm *ConfigManager) GetConfigBool(key string) (bool, error) {
	value, err := cm.GetConfig(key)
	if err != nil {
		return false, err
	}

	boolVal, ok := value.(bool)
	if !ok {
		return false, fmt.Errorf("config key '%s' is not a boolean", key)
	}

	return boolVal, nil
}

// GetConfigMap gets a configuration value as map
func (cm *ConfigManager) GetConfigMap(key string) (map[string]interface{}, error) {
	value, err := cm.GetConfig(key)
	if err != nil {
		return nil, err
	}

	mapVal, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("config key '%s' is not a map", key)
	}

	return mapVal, nil
}

// GetConfigSlice gets a configuration value as slice
func (cm *ConfigManager) GetConfigSlice(key string) ([]interface{}, error) {
	value, err := cm.GetConfig(key)
	if err != nil {
		return nil, err
	}

	sliceVal, ok := value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("config key '%s' is not a slice", key)
	}

	return sliceVal, nil
}

// HasConfig checks if a configuration key exists
func (cm *ConfigManager) HasConfig(key string) bool {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	_, exists := cm.configs[key]
	return exists
}

// RemoveConfig removes a configuration key
func (cm *ConfigManager) RemoveConfig(key string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if _, exists := cm.configs[key]; !exists {
		return fmt.Errorf("config key '%s' not found", key)
	}

	delete(cm.configs, key)
	cm.updatedAt = time.Now().Unix()
	cm.version++

	return nil
}

// GetAllConfigs returns all configuration keys and values
func (cm *ConfigManager) GetAllConfigs() map[string]interface{} {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	result := make(map[string]interface{})
	for k, v := range cm.configs {
		result[k] = v
	}
	return result
}

// GetConfigKeys returns all configuration keys
func (cm *ConfigManager) GetConfigKeys() []string {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	keys := make([]string, 0, len(cm.configs))
	for key := range cm.configs {
		keys = append(keys, key)
	}
	return keys
}

// GetConfigCount returns the number of configuration keys
func (cm *ConfigManager) GetConfigCount() int {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	return len(cm.configs)
}

// ClearConfigs clears all configurations
func (cm *ConfigManager) ClearConfigs() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if len(cm.configs) > 0 {
		cm.configs = make(map[string]interface{})
		cm.updatedAt = time.Now().Unix()
		cm.version++
	}
}

// LoadFromFile loads configurations from a file
func (cm *ConfigManager) LoadFromFile(filePath string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("config file '%s' does not exist", filePath)
	}

	// Read file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse JSON
	var configFile ConfigFile
	if err := json.Unmarshal(data, &configFile); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	// Load configurations
	cm.configs = make(map[string]interface{})
	for key, configData := range configFile.Configs {
		cm.configs[key] = configData.Value
	}

	cm.filePath = filePath
	cm.version = configFile.Version
	cm.createdAt = configFile.CreatedAt
	cm.updatedAt = time.Now().Unix()

	return nil
}

// SaveToFile saves configurations to a file
func (cm *ConfigManager) SaveToFile(filePath string) error {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	// Create directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Prepare config file data
	configFile := ConfigFile{
		Version:   cm.version,
		CreatedAt: cm.createdAt,
		UpdatedAt: cm.updatedAt,
		Configs:   make(map[string]ConfigData),
	}

	// Convert configs to ConfigData
	for key, value := range cm.configs {
		configFile.Configs[key] = ConfigData{
			Key:       key,
			Value:     value,
			Type:      fmt.Sprintf("%T", value),
			Category:  "general",
			UpdatedAt: cm.updatedAt,
		}
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(configFile, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Write to file
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	cm.filePath = filePath
	return nil
}

// ReloadFromFile reloads configurations from the current file
func (cm *ConfigManager) ReloadFromFile() error {
	if cm.filePath == "" {
		return fmt.Errorf("no file path set")
	}
	return cm.LoadFromFile(cm.filePath)
}

// SaveToCurrentFile saves configurations to the current file
func (cm *ConfigManager) SaveToCurrentFile() error {
	if cm.filePath == "" {
		return fmt.Errorf("no file path set")
	}
	return cm.SaveToFile(cm.filePath)
}

// SetFilePath sets the file path for the config manager
func (cm *ConfigManager) SetFilePath(filePath string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.filePath = filePath
}

// GetFilePath returns the current file path
func (cm *ConfigManager) GetFilePath() string {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	return cm.filePath
}

// GetVersion returns the current version
func (cm *ConfigManager) GetVersion() int64 {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	return cm.version
}

// GetUpdatedAt returns the last update timestamp
func (cm *ConfigManager) GetUpdatedAt() int64 {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	return cm.updatedAt
}

// GetCreatedAt returns the creation timestamp
func (cm *ConfigManager) GetCreatedAt() int64 {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	return cm.createdAt
}

// ValidateConfig validates a configuration value
func (cm *ConfigManager) ValidateConfig(key string, value interface{}) error {
	if key == "" {
		return fmt.Errorf("config key cannot be empty")
	}

	// Add custom validation rules here
	switch key {
	case "max_connections":
		if intVal, ok := value.(int); ok {
			if intVal <= 0 {
				return fmt.Errorf("max_connections must be positive")
			}
		}
	case "timeout":
		if floatVal, ok := value.(float64); ok {
			if floatVal < 0 {
				return fmt.Errorf("timeout must be non-negative")
			}
		}
	case "debug":
		if _, ok := value.(bool); !ok {
			return fmt.Errorf("debug must be a boolean")
		}
	}

	return nil
}

// SetConfigWithValidation sets a configuration value with validation
func (cm *ConfigManager) SetConfigWithValidation(key string, value interface{}) error {
	if err := cm.ValidateConfig(key, value); err != nil {
		return err
	}
	return cm.SetConfig(key, value)
}

// Clone creates a deep copy of the ConfigManager
func (cm *ConfigManager) Clone() *ConfigManager {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	clone := &ConfigManager{
		configs:   make(map[string]interface{}),
		filePath:  cm.filePath,
		version:   cm.version,
		createdAt: cm.createdAt,
		updatedAt: cm.updatedAt,
	}

	// Deep copy configs
	for k, v := range cm.configs {
		clone.configs[k] = v
	}

	return clone
}

// Merge merges another ConfigManager into this one
func (cm *ConfigManager) Merge(other *ConfigManager) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	other.mu.RLock()
	defer other.mu.RUnlock()

	// Merge configurations
	for key, value := range other.configs {
		cm.configs[key] = value
	}

	cm.updatedAt = time.Now().Unix()
	cm.version++
}

// ToJSON converts ConfigManager to JSON
func (cm *ConfigManager) ToJSON() ([]byte, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	configFile := ConfigFile{
		Version:   cm.version,
		CreatedAt: cm.createdAt,
		UpdatedAt: cm.updatedAt,
		Configs:   make(map[string]ConfigData),
	}

	for key, value := range cm.configs {
		configFile.Configs[key] = ConfigData{
			Key:       key,
			Value:     value,
			Type:      fmt.Sprintf("%T", value),
			Category:  "general",
			UpdatedAt: cm.updatedAt,
		}
	}

	return json.MarshalIndent(configFile, "", "  ")
}

// FromJSON creates ConfigManager from JSON
func (cm *ConfigManager) FromJSON(data []byte) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	var configFile ConfigFile
	if err := json.Unmarshal(data, &configFile); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	cm.configs = make(map[string]interface{})
	for key, configData := range configFile.Configs {
		cm.configs[key] = configData.Value
	}

	cm.version = configFile.Version
	cm.createdAt = configFile.CreatedAt
	cm.updatedAt = configFile.UpdatedAt

	return nil
}
