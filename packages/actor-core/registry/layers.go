package registry

import (
	"chaos-actor-module/packages/actor-core/enums"
	"chaos-actor-module/packages/actor-core/interfaces"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// CapLayerRegistryImpl implements the CapLayerRegistry interface
type CapLayerRegistryImpl struct {
	layerOrder      []string
	acrossPolicy    string
	mu              sync.RWMutex
	filePath        string
}

// NewCapLayerRegistry creates a new cap layer registry with default configuration
func NewCapLayerRegistry() interfaces.CapLayerRegistry {
	return &CapLayerRegistryImpl{
		layerOrder:   []string{enums.LayerRealm.String(), enums.LayerWorld.String(), enums.LayerEvent.String(), enums.LayerGuild.String(), enums.LayerTotal.String()},
		acrossPolicy: "intersect",
	}
}

// NewCapLayerRegistryFromFile creates a new cap layer registry from a file
func NewCapLayerRegistryFromFile(filePath string) (interfaces.CapLayerRegistry, error) {
	registry := &CapLayerRegistryImpl{
		layerOrder:   []string{enums.LayerRealm.String(), enums.LayerWorld.String(), enums.LayerEvent.String(), enums.LayerGuild.String(), enums.LayerTotal.String()},
		acrossPolicy: "intersect",
		filePath:     filePath,
	}
	
	if err := registry.LoadFromFile(filePath); err != nil {
		return nil, fmt.Errorf("failed to load cap layer registry from file: %w", err)
	}
	
	return registry, nil
}

// GetLayerOrder returns the processing order for layers
func (clr *CapLayerRegistryImpl) GetLayerOrder() []string {
	clr.mu.RLock()
	defer clr.mu.RUnlock()
	
	// Return a copy to prevent external modification
	order := make([]string, len(clr.layerOrder))
	copy(order, clr.layerOrder)
	return order
}

// GetAcrossLayerPolicy returns the across-layer policy
func (clr *CapLayerRegistryImpl) GetAcrossLayerPolicy() string {
	clr.mu.RLock()
	defer clr.mu.RUnlock()
	
	return clr.acrossPolicy
}

// SetLayerOrder sets the processing order for layers
func (clr *CapLayerRegistryImpl) SetLayerOrder(order []string) error {
	clr.mu.Lock()
	defer clr.mu.Unlock()
	
	if len(order) == 0 {
		return fmt.Errorf("layer order cannot be empty")
	}
	
	// Validate layer names
	validLayers := map[string]bool{
		enums.LayerRealm.String(): true,
		enums.LayerWorld.String(): true,
		enums.LayerEvent.String(): true,
		enums.LayerGuild.String(): true,
		enums.LayerTotal.String(): true,
	}
	
	for _, layer := range order {
		if !validLayers[layer] {
			return fmt.Errorf("invalid layer: %s", layer)
		}
	}
	
	// Check for duplicates
	seen := make(map[string]bool)
	for _, layer := range order {
		if seen[layer] {
			return fmt.Errorf("duplicate layer in order: %s", layer)
		}
		seen[layer] = true
	}
	
	clr.layerOrder = make([]string, len(order))
	copy(clr.layerOrder, order)
	return nil
}

// SetAcrossLayerPolicy sets the across-layer policy
func (clr *CapLayerRegistryImpl) SetAcrossLayerPolicy(policy string) error {
	clr.mu.Lock()
	defer clr.mu.Unlock()
	
	validPolicies := map[string]bool{
		"intersect": true,
		"union":     true,
	}
	
	if !validPolicies[policy] {
		return fmt.Errorf("invalid across-layer policy: %s", policy)
	}
	
	clr.acrossPolicy = policy
	return nil
}

// LoadFromConfig loads configuration from config
func (clr *CapLayerRegistryImpl) LoadFromConfig(config map[string]interface{}) error {
	clr.mu.Lock()
	defer clr.mu.Unlock()
	
	// Load layer order
	if orderData, exists := config["order"]; exists {
		orderSlice, ok := orderData.([]interface{})
		if !ok {
			return fmt.Errorf("order must be an array")
		}
		
		order := make([]string, len(orderSlice))
		for i, layer := range orderSlice {
			if layerStr, ok := layer.(string); ok {
				order[i] = layerStr
			} else {
				return fmt.Errorf("order[%d] must be a string", i)
			}
		}
		
		if err := clr.setLayerOrderUnsafe(order); err != nil {
			return err
		}
	}
	
	// Load across-layer policy
	if policyData, exists := config["across_policy"]; exists {
		if policy, ok := policyData.(string); ok {
			if err := clr.setAcrossLayerPolicyUnsafe(policy); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("across_policy must be a string")
		}
	}
	
	return nil
}

// LoadFromFile loads configuration from a file
func (clr *CapLayerRegistryImpl) LoadFromFile(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	
	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	
	clr.filePath = filePath
	return clr.LoadFromConfig(config)
}

// SaveToFile saves configuration to a file
func (clr *CapLayerRegistryImpl) SaveToFile(filePath string) error {
	clr.mu.RLock()
	defer clr.mu.RUnlock()
	
	config := map[string]interface{}{
		"order":         clr.layerOrder,
		"across_policy": clr.acrossPolicy,
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
	
	clr.filePath = filePath
	return nil
}

// Validate validates the configuration
func (clr *CapLayerRegistryImpl) Validate() error {
	clr.mu.RLock()
	defer clr.mu.RUnlock()
	
	// Validate layer order
	if len(clr.layerOrder) == 0 {
		return fmt.Errorf("layer order cannot be empty")
	}
	
	// Validate layer names
	validLayers := map[string]bool{
		enums.LayerRealm.String(): true,
		enums.LayerWorld.String(): true,
		enums.LayerEvent.String(): true,
		enums.LayerGuild.String(): true,
		enums.LayerTotal.String(): true,
	}
	
	for _, layer := range clr.layerOrder {
		if !validLayers[layer] {
			return fmt.Errorf("invalid layer: %s", layer)
		}
	}
	
	// Check for duplicates
	seen := make(map[string]bool)
	for _, layer := range clr.layerOrder {
		if seen[layer] {
			return fmt.Errorf("duplicate layer in order: %s", layer)
		}
		seen[layer] = true
	}
	
	// Validate across-layer policy
	validPolicies := map[string]bool{
		"intersect": true,
		"union":     true,
	}
	
	if !validPolicies[clr.acrossPolicy] {
		return fmt.Errorf("invalid across-layer policy: %s", clr.acrossPolicy)
	}
	
	return nil
}

// GetLayerIndex returns the index of a layer in the order
func (clr *CapLayerRegistryImpl) GetLayerIndex(layer string) (int64, bool) {
	clr.mu.RLock()
	defer clr.mu.RUnlock()
	
	for i, l := range clr.layerOrder {
		if l == layer {
			return int64(i), true
		}
	}
	
	return -1, false
}

// IsLayerInOrder checks if a layer is in the order
func (clr *CapLayerRegistryImpl) IsLayerInOrder(layer string) bool {
	clr.mu.RLock()
	defer clr.mu.RUnlock()
	
	for _, l := range clr.layerOrder {
		if l == layer {
			return true
		}
	}
	
	return false
}

// GetLayerCount returns the number of layers
func (clr *CapLayerRegistryImpl) GetLayerCount() int64 {
	clr.mu.RLock()
	defer clr.mu.RUnlock()
	
	return int64(len(clr.layerOrder))
}

// GetFilepath returns the current file path
func (clr *CapLayerRegistryImpl) GetFilepath() string {
	clr.mu.RLock()
	defer clr.mu.RUnlock()
	
	return clr.filePath
}

// SetFilepath sets the file path
func (clr *CapLayerRegistryImpl) SetFilepath(filePath string) {
	clr.mu.Lock()
	defer clr.mu.Unlock()
	
	clr.filePath = filePath
}

// Reload reloads configuration from the current file
func (clr *CapLayerRegistryImpl) Reload() error {
	if clr.filePath == "" {
		return fmt.Errorf("no file path set")
	}
	
	return clr.LoadFromFile(clr.filePath)
}

// Save saves configuration to the current file
func (clr *CapLayerRegistryImpl) Save() error {
	if clr.filePath == "" {
		return fmt.Errorf("no file path set")
	}
	
	return clr.SaveToFile(clr.filePath)
}

// Reset resets to default configuration
func (clr *CapLayerRegistryImpl) Reset() {
	clr.mu.Lock()
	defer clr.mu.Unlock()
	
	clr.layerOrder = []string{enums.LayerRealm.String(), enums.LayerWorld.String(), enums.LayerEvent.String(), enums.LayerGuild.String(), enums.LayerTotal.String()}
	clr.acrossPolicy = "intersect"
}

// setLayerOrderUnsafe sets layer order without locking (internal use)
func (clr *CapLayerRegistryImpl) setLayerOrderUnsafe(order []string) error {
	if len(order) == 0 {
		return fmt.Errorf("layer order cannot be empty")
	}
	
	// Validate layer names
	validLayers := map[string]bool{
		enums.LayerRealm.String(): true,
		enums.LayerWorld.String(): true,
		enums.LayerEvent.String(): true,
		enums.LayerGuild.String(): true,
		enums.LayerTotal.String(): true,
	}
	
	for _, layer := range order {
		if !validLayers[layer] {
			return fmt.Errorf("invalid layer: %s", layer)
		}
	}
	
	// Check for duplicates
	seen := make(map[string]bool)
	for _, layer := range order {
		if seen[layer] {
			return fmt.Errorf("duplicate layer in order: %s", layer)
		}
		seen[layer] = true
	}
	
	clr.layerOrder = make([]string, len(order))
	copy(clr.layerOrder, order)
	return nil
}

// setAcrossLayerPolicyUnsafe sets across-layer policy without locking (internal use)
func (clr *CapLayerRegistryImpl) setAcrossLayerPolicyUnsafe(policy string) error {
	validPolicies := map[string]bool{
		"intersect": true,
		"union":     true,
	}
	
	if !validPolicies[policy] {
		return fmt.Errorf("invalid across-layer policy: %s", policy)
	}
	
	clr.acrossPolicy = policy
	return nil
}
