package registry

import (
	"chaos-actor-module/packages/actor-core/constants"
	"chaos-actor-module/packages/actor-core/interfaces"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// CombinerRegistryImpl implements the CombinerRegistry interface
type CombinerRegistryImpl struct {
	rules    map[string]*interfaces.MergeRule
	mu       sync.RWMutex
	filePath string
}

// NewCombinerRegistry creates a new combiner registry
func NewCombinerRegistry() interfaces.CombinerRegistry {
	return &CombinerRegistryImpl{
		rules: make(map[string]*interfaces.MergeRule),
	}
}

// NewCombinerRegistryFromFile creates a new combiner registry from a file
func NewCombinerRegistryFromFile(filePath string) (interfaces.CombinerRegistry, error) {
	registry := &CombinerRegistryImpl{
		rules:    make(map[string]*interfaces.MergeRule),
		filePath: filePath,
	}
	
	if err := registry.LoadFromFile(filePath); err != nil {
		return nil, fmt.Errorf("failed to load combiner registry from file: %w", err)
	}
	
	return registry, nil
}

// GetRule returns the merge rule for the given dimension
func (cr *CombinerRegistryImpl) GetRule(dimension string) (*interfaces.MergeRule, error) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	
	rule, exists := cr.rules[dimension]
	if !exists {
		// Return default rule if not found
		return cr.getDefaultRule(dimension), nil
	}
	
	return rule, nil
}

// SetRule sets the merge rule for the given dimension
func (cr *CombinerRegistryImpl) SetRule(dimension string, rule *interfaces.MergeRule) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	
	if rule == nil {
		return fmt.Errorf("rule cannot be nil")
	}
	
	if !rule.IsValid() {
		return fmt.Errorf("invalid rule for dimension %s", dimension)
	}
	
	cr.rules[dimension] = rule
	return nil
}

// GetDimensions returns all dimensions with rules
func (cr *CombinerRegistryImpl) GetDimensions() []string {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	
	dimensions := make([]string, 0, len(cr.rules))
	for dimension := range cr.rules {
		dimensions = append(dimensions, dimension)
	}
	
	return dimensions
}

// LoadFromConfig loads rules from configuration
func (cr *CombinerRegistryImpl) LoadFromConfig(config map[string]interface{}) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	
	rules, ok := config["rules"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid config: rules not found or not a map")
	}
	
	for dimension, ruleData := range rules {
		rule, err := cr.parseRule(ruleData)
		if err != nil {
			return fmt.Errorf("failed to parse rule for dimension %s: %w", dimension, err)
		}
		
		cr.rules[dimension] = rule
	}
	
	return nil
}

// LoadFromFile loads rules from a file
func (cr *CombinerRegistryImpl) LoadFromFile(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	
	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	
	cr.filePath = filePath
	return cr.LoadFromConfig(config)
}

// SaveToFile saves rules to a file
func (cr *CombinerRegistryImpl) SaveToFile(filePath string) error {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	
	config := map[string]interface{}{
		"rules": cr.rules,
	}
	
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	
	// Ensure directory exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}
	
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	
	cr.filePath = filePath
	return nil
}

// Validate validates all rules
func (cr *CombinerRegistryImpl) Validate() error {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	
	for dimension, rule := range cr.rules {
		if rule == nil {
			return fmt.Errorf("rule for dimension %s is nil", dimension)
		}
		
		if !rule.IsValid() {
			return fmt.Errorf("invalid rule for dimension %s", dimension)
		}
	}
	
	return nil
}

// Clear clears all rules
func (cr *CombinerRegistryImpl) Clear() {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	
	cr.rules = make(map[string]*interfaces.MergeRule)
}

// Count returns the number of rules
func (cr *CombinerRegistryImpl) Count() int64 {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	
	return int64(len(cr.rules))
}

// HasRule checks if a rule exists for the given dimension
func (cr *CombinerRegistryImpl) HasRule(dimension string) bool {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	
	_, exists := cr.rules[dimension]
	return exists
}

// RemoveRule removes a rule for the given dimension
func (cr *CombinerRegistryImpl) RemoveRule(dimension string) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	
	delete(cr.rules, dimension)
}

// GetDefaultRule returns the default rule for a dimension
func (cr *CombinerRegistryImpl) getDefaultRule(dimension string) *interfaces.MergeRule {
	// Get default clamp range based on dimension type
	min, max := cr.getDefaultClampRange(dimension)
	
	return &interfaces.MergeRule{
		UsePipeline: true,
		ClampDefault: interfaces.Caps{
			Min: min,
			Max: max,
		},
	}
}

// getDefaultClampRange returns the default clamp range for a dimension
func (cr *CombinerRegistryImpl) getDefaultClampRange(dimension string) (float64, float64) {
	switch dimension {
	// Primary dimensions
	case constants.DimensionStrength, constants.DimensionVitality, constants.DimensionDexterity,
		 constants.DimensionIntelligence, constants.DimensionSpirit, constants.DimensionCharisma:
		return constants.MinStrength, constants.MaxStrength
	
	// Health & resources
	case constants.DimensionHPMax:
		return constants.MinHPMax, constants.MaxHPMax
	case constants.DimensionMPMax:
		return constants.MinMPMax, constants.MaxMPMax
	case constants.DimensionStaminaMax:
		return constants.MinStaminaMax, constants.MaxStaminaMax
	
	// Combat attributes
	case constants.DimensionAttackPower, constants.DimensionDefense,
		 constants.DimensionMagicPower, constants.DimensionMagicResistance:
		return constants.MinAttackPower, constants.MaxAttackPower
	
	// Critical & accuracy
	case constants.DimensionCritRate, constants.DimensionAccuracy:
		return constants.MinCritRate, constants.MaxCritRate
	case constants.DimensionCritDamage:
		return constants.MinCritDamage, constants.MaxCritDamage
	
	// Speed & movement
	case constants.DimensionMoveSpeed:
		return constants.MinMoveSpeed, constants.MaxMoveSpeed
	case constants.DimensionAttackSpeed:
		return constants.MinAttackSpeed, constants.MaxAttackSpeed
	case constants.DimensionCastSpeed:
		return constants.MinCastSpeed, constants.MaxCastSpeed
	
	// Resource management
	case constants.DimensionCooldownReduction, constants.DimensionManaEfficiency,
		 constants.DimensionEnergyEfficiency:
		return constants.MinCooldownReduction, constants.MaxCooldownReduction
	
	// Learning & progression
	case constants.DimensionLearningRate:
		return constants.MinLearningRate, constants.MaxLearningRate
	case constants.DimensionCultivationSpeed:
		return constants.MinCultivationSpeed, constants.MaxCultivationSpeed
	case constants.DimensionBreakthroughSuccess:
		return constants.MinBreakthroughSuccess, constants.MaxBreakthroughSuccess
	
	// Meta/World
	case constants.DimensionLifespanYears:
		return constants.MinLifespanYears, constants.MaxLifespanYears
	case constants.DimensionPoiseRank:
		return constants.MinPoiseRank, constants.MaxPoiseRank
	case constants.DimensionStealth, constants.DimensionPerception:
		return constants.MinStealth, constants.MaxStealth
	case constants.DimensionLuck:
		return constants.MinLuck, constants.MaxLuck
	
	default:
		// Default range for unknown dimensions
		return 0.0, 1000000.0
	}
}

// parseRule parses a rule from interface{} data
func (cr *CombinerRegistryImpl) parseRule(data interface{}) (*interfaces.MergeRule, error) {
	ruleMap, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("rule data is not a map")
	}
	
	rule := &interfaces.MergeRule{}
	
	// Parse UsePipeline
	if usePipeline, exists := ruleMap["use_pipeline"]; exists {
		if bp, ok := usePipeline.(bool); ok {
			rule.UsePipeline = bp
		} else {
			return nil, fmt.Errorf("use_pipeline must be a boolean")
		}
	} else {
		rule.UsePipeline = true // Default
	}
	
	// Parse Operator
	if operator, exists := ruleMap["operator"]; exists {
		if op, ok := operator.(string); ok {
			rule.Operator = op
		} else {
			return nil, fmt.Errorf("operator must be a string")
		}
	}
	
	// Parse ClampDefault
	if clampDefault, exists := ruleMap["clamp_default"]; exists {
		clampMap, ok := clampDefault.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("clamp_default must be a map")
		}
		
		var min, max float64
		
		if minVal, exists := clampMap["min"]; exists {
			if m, ok := minVal.(float64); ok {
				min = m
			} else {
				return nil, fmt.Errorf("clamp_default.min must be a number")
			}
		}
		
		if maxVal, exists := clampMap["max"]; exists {
			if m, ok := maxVal.(float64); ok {
				max = m
			} else {
				return nil, fmt.Errorf("clamp_default.max must be a number")
			}
		}
		
		rule.ClampDefault = interfaces.Caps{
			Min: min,
			Max: max,
		}
	} else {
		// Use default clamp range
		rule.ClampDefault = interfaces.Caps{
			Min: 0.0,
			Max: 1000000.0,
		}
	}
	
	return rule, nil
}

// GetFilepath returns the current file path
func (cr *CombinerRegistryImpl) GetFilepath() string {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	
	return cr.filePath
}

// SetFilepath sets the file path
func (cr *CombinerRegistryImpl) SetFilepath(filePath string) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	
	cr.filePath = filePath
}

// Reload reloads rules from the current file
func (cr *CombinerRegistryImpl) Reload() error {
	if cr.filePath == "" {
		return fmt.Errorf("no file path set")
	}
	
	return cr.LoadFromFile(cr.filePath)
}

// Save saves rules to the current file
func (cr *CombinerRegistryImpl) Save() error {
	if cr.filePath == "" {
		return fmt.Errorf("no file path set")
	}
	
	return cr.SaveToFile(cr.filePath)
}
