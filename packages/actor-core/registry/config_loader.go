package registry

import (
	"chaos-actor-module/packages/actor-core/interfaces"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ConfigLoaderImpl implements the ConfigLoader interface
type ConfigLoaderImpl struct {
	supportedFormats map[string]bool
}

// NewConfigLoader creates a new config loader
func NewConfigLoader() interfaces.ConfigLoader {
	return &ConfigLoaderImpl{
		supportedFormats: map[string]bool{
			"json": true,
			"yaml": true,
			"yml":  true,
		},
	}
}

// Load loads configuration from a file
func (cl *ConfigLoaderImpl) Load(filename string) (map[string]interface{}, error) {
	if filename == "" {
		return nil, fmt.Errorf("filename cannot be empty")
	}

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", filename)
	}

	// Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	// Determine format from file extension
	ext := strings.ToLower(filepath.Ext(filename))
	if ext != "" {
		ext = ext[1:] // Remove the dot
	}

	return cl.LoadFromBytes(data)
}

// LoadFromBytes loads configuration from bytes
func (cl *ConfigLoaderImpl) LoadFromBytes(data []byte) (map[string]interface{}, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("data cannot be empty")
	}

	// Default to JSON format
	format := "json"

	format = strings.ToLower(format)
	if !cl.supportedFormats[format] {
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	var config map[string]interface{}

	switch format {
	case "json":
		if err := json.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %w", err)
		}
	case "yaml", "yml":
		// For now, we'll treat YAML as JSON since we don't have YAML parser
		// In a real implementation, you would use gopkg.in/yaml.v3
		if err := json.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("failed to parse YAML (using JSON parser): %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	return config, nil
}

// Save saves configuration to a file
func (cl *ConfigLoaderImpl) Save(filename string, config map[string]interface{}) error {
	if filename == "" {
		return fmt.Errorf("filename cannot be empty")
	}

	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	// Determine format from file extension
	ext := strings.ToLower(filepath.Ext(filename))
	if ext != "" {
		ext = ext[1:] // Remove the dot
	}

	if ext == "" {
		ext = "json" // Default format
	}

	format := strings.ToLower(ext)
	if !cl.supportedFormats[format] {
		return fmt.Errorf("unsupported format: %s", format)
	}

	var data []byte
	var err error

	switch format {
	case "json":
		data, err = json.MarshalIndent(config, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
	case "yaml", "yml":
		// For now, we'll save as JSON since we don't have YAML marshaler
		// In a real implementation, you would use gopkg.in/yaml.v3
		data, err = json.MarshalIndent(config, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal YAML (using JSON marshaler): %w", err)
		}
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}

	// Ensure directory exists
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write file
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}

	return nil
}

// Validate validates the configuration
func (cl *ConfigLoaderImpl) Validate(config map[string]interface{}) error {
	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	// Basic validation - check if config is a map
	if len(config) == 0 {
		return fmt.Errorf("config cannot be empty")
	}

	// You can add more specific validation rules here
	// For example, check for required fields, validate types, etc.

	return nil
}

// GetSupportedFormats returns the supported formats
func (cl *ConfigLoaderImpl) GetSupportedFormats() []string {
	formats := make([]string, 0, len(cl.supportedFormats))
	for format := range cl.supportedFormats {
		formats = append(formats, format)
	}
	return formats
}

// IsFormatSupported checks if a format is supported
func (cl *ConfigLoaderImpl) IsFormatSupported(format string) bool {
	return cl.supportedFormats[strings.ToLower(format)]
}

// LoadFromFileWithFormat loads configuration from a file with a specific format
func (cl *ConfigLoaderImpl) LoadFromFileWithFormat(filename, format string) (map[string]interface{}, error) {
	if filename == "" {
		return nil, fmt.Errorf("filename cannot be empty")
	}

	if format == "" {
		return nil, fmt.Errorf("format cannot be empty")
	}

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", filename)
	}

	// Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	return cl.LoadFromBytes(data)
}

// SaveToFileWithFormat saves configuration to a file with a specific format
func (cl *ConfigLoaderImpl) SaveToFileWithFormat(filename string, config map[string]interface{}, format string) error {
	if filename == "" {
		return fmt.Errorf("filename cannot be empty")
	}

	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	if format == "" {
		return fmt.Errorf("format cannot be empty")
	}

	// Determine format from file extension if not provided
	ext := strings.ToLower(filepath.Ext(filename))
	if ext != "" {
		ext = ext[1:] // Remove the dot
	}

	if ext == "" {
		ext = format
	}

	format = strings.ToLower(format)
	if !cl.supportedFormats[format] {
		return fmt.Errorf("unsupported format: %s", format)
	}

	var data []byte
	var err error

	switch format {
	case "json":
		data, err = json.MarshalIndent(config, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
	case "yaml", "yml":
		// For now, we'll save as JSON since we don't have YAML marshaler
		data, err = json.MarshalIndent(config, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal YAML (using JSON marshaler): %w", err)
		}
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}

	// Ensure directory exists
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write file
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}

	return nil
}

// ValidateFile validates a configuration file
func (cl *ConfigLoaderImpl) ValidateFile(filename string) error {
	config, err := cl.Load(filename)
	if err != nil {
		return err
	}

	return cl.Validate(config)
}

// GetFileFormat returns the format of a file based on its extension
func (cl *ConfigLoaderImpl) GetFileFormat(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	if ext != "" {
		ext = ext[1:] // Remove the dot
	}

	if ext == "" {
		return "json" // Default format
	}

	return ext
}

// IsValidConfigFile checks if a file is a valid configuration file
func (cl *ConfigLoaderImpl) IsValidConfigFile(filename string) bool {
	_, err := cl.Load(filename)
	return err == nil
}
