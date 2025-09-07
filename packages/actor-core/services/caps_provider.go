package services

import (
	"chaos-actor-module/packages/actor-core/interfaces"
	"context"
	"fmt"
	"sort"
	"sync"
)

// CapsProviderImpl implements the CapsProvider interface
type CapsProviderImpl struct {
	layerRegistry interfaces.CapLayerRegistry
	mu            sync.RWMutex
}

// NewCapsProvider creates a new caps provider
func NewCapsProvider(layerRegistry interfaces.CapLayerRegistry) interfaces.CapsProvider {
	return &CapsProviderImpl{
		layerRegistry: layerRegistry,
	}
}

// EffectiveCapsWithinLayer returns effective caps within a specific layer
func (cp *CapsProviderImpl) EffectiveCapsWithinLayer(ctx context.Context, actor *interfaces.Actor, outputs []*interfaces.SubsystemOutput, layer string) (interfaces.EffectiveCaps, error) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	if actor == nil {
		return nil, fmt.Errorf("actor cannot be nil")
	}

	if outputs == nil {
		return nil, fmt.Errorf("outputs cannot be nil")
	}

	if layer == "" {
		return nil, fmt.Errorf("layer cannot be empty")
	}

	// Get layer order
	layerOrder := cp.layerRegistry.GetLayerOrder()
	layerIndex := -1
	for i, l := range layerOrder {
		if l == layer {
			layerIndex = i
			break
		}
	}

	if layerIndex == -1 {
		return nil, fmt.Errorf("invalid layer: %s", layer)
	}

	// Collect caps for this layer
	layerCaps := make(map[string][]interfaces.CapContribution)

	for _, output := range outputs {
		if output == nil {
			continue
		}

		for _, cap := range output.Caps {
			if cap.Scope == layer {
				layerCaps[cap.Dimension] = append(layerCaps[cap.Dimension], cap)
			}
		}
	}

	// Calculate effective caps for each dimension
	effectiveCaps := make(interfaces.EffectiveCaps)

	for dimension, caps := range layerCaps {
		if len(caps) == 0 {
			continue
		}

		effectiveCap, err := cp.calculateEffectiveCap(caps, dimension)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate effective cap for dimension %s: %w", dimension, err)
		}

		effectiveCaps[dimension] = effectiveCap
	}

	return effectiveCaps, nil
}

// EffectiveCapsAcrossLayers returns effective caps across all layers
func (cp *CapsProviderImpl) EffectiveCapsAcrossLayers(ctx context.Context, actor *interfaces.Actor, outputs []*interfaces.SubsystemOutput) (interfaces.EffectiveCaps, error) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	if actor == nil {
		return nil, fmt.Errorf("actor cannot be nil")
	}

	if outputs == nil {
		return nil, fmt.Errorf("outputs cannot be nil")
	}

	// Get layer order
	layerOrder := cp.layerRegistry.GetLayerOrder()
	acrossPolicy := cp.layerRegistry.GetAcrossLayerPolicy()

	// Collect caps by layer
	layerCaps := make(map[string]interfaces.EffectiveCaps)

	for _, layer := range layerOrder {
		effectiveCaps, err := cp.EffectiveCapsWithinLayer(ctx, actor, outputs, layer)
		if err != nil {
			return nil, fmt.Errorf("failed to get effective caps for layer %s: %w", layer, err)
		}

		layerCaps[layer] = effectiveCaps
	}

	// Combine caps across layers
	effectiveCaps := make(interfaces.EffectiveCaps)

	// Get all dimensions
	allDimensions := make(map[string]bool)
	for _, caps := range layerCaps {
		for dimension := range caps {
			allDimensions[dimension] = true
		}
	}

	// Calculate effective caps for each dimension
	for dimension := range allDimensions {
		effectiveCap, err := cp.combineCapsAcrossLayers(dimension, layerCaps, acrossPolicy)
		if err != nil {
			return nil, fmt.Errorf("failed to combine caps for dimension %s: %w", dimension, err)
		}

		effectiveCaps[dimension] = effectiveCap
	}

	return effectiveCaps, nil
}

// GetLayerOrder returns the processing order for layers
func (cp *CapsProviderImpl) GetLayerOrder() []string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	return cp.layerRegistry.GetLayerOrder()
}

// GetAcrossLayerPolicy returns the across-layer policy
func (cp *CapsProviderImpl) GetAcrossLayerPolicy() string {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	return cp.layerRegistry.GetAcrossLayerPolicy()
}

// GetCapsForDimension returns caps for a specific dimension
func (cp *CapsProviderImpl) GetCapsForDimension(dimension string) (interfaces.Caps, error) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	if dimension == "" {
		return interfaces.Caps{}, fmt.Errorf("dimension cannot be empty")
	}

	// Return default caps for the dimension
	// In a real implementation, this would query the combiner registry
	return interfaces.Caps{
		Min: 0.0,
		Max: 1000000.0,
	}, nil
}

// calculateEffectiveCap calculates effective cap for a dimension within a layer
func (cp *CapsProviderImpl) calculateEffectiveCap(caps []interfaces.CapContribution, dimension string) (interfaces.Caps, error) {
	if len(caps) == 0 {
		return interfaces.Caps{}, fmt.Errorf("no caps provided")
	}

	// Sort caps by priority (higher priority first)
	sort.Slice(caps, func(i, j int) bool {
		return caps[i].Priority > caps[j].Priority
	})

	// Apply caps based on mode
	var effectiveCap interfaces.Caps
	first := true

	for _, cap := range caps {
		if first {
			effectiveCap = interfaces.Caps{
				Min: cap.Value,
				Max: cap.Value,
			}
			first = false
		} else {
			switch cap.Mode {
			case "BASELINE":
				// Baseline sets the base value
				effectiveCap.Min = cap.Value
				effectiveCap.Max = cap.Value
			case "ADDITIVE":
				// Additive adds to the current range
				effectiveCap.Min += cap.Value
				effectiveCap.Max += cap.Value
			case "HARD_MAX":
				// Hard max sets the maximum
				effectiveCap.Max = cap.Value
			case "HARD_MIN":
				// Hard min sets the minimum
				effectiveCap.Min = cap.Value
			case "OVERRIDE":
				// Override replaces the current range
				effectiveCap.Min = cap.Value
				effectiveCap.Max = cap.Value
			}
		}
	}

	// Ensure min <= max
	if effectiveCap.Min > effectiveCap.Max {
		effectiveCap.Min = effectiveCap.Max
	}

	return effectiveCap, nil
}

// combineCapsAcrossLayers combines caps across layers
func (cp *CapsProviderImpl) combineCapsAcrossLayers(dimension string, layerCaps map[string]interfaces.EffectiveCaps, policy string) (interfaces.Caps, error) {
	// Collect caps from all layers
	var caps []interfaces.Caps

	for _, effectiveCaps := range layerCaps {
		if cap, exists := effectiveCaps[dimension]; exists {
			caps = append(caps, cap)
		}
	}

	if len(caps) == 0 {
		return interfaces.Caps{}, fmt.Errorf("no caps found for dimension %s", dimension)
	}

	// Combine based on policy
	switch policy {
	case "intersect":
		return cp.intersectCaps(caps)
	case "union":
		return cp.unionCaps(caps)
	default:
		return cp.intersectCaps(caps) // Default to intersect
	}
}

// intersectCaps intersects multiple caps
func (cp *CapsProviderImpl) intersectCaps(caps []interfaces.Caps) (interfaces.Caps, error) {
	if len(caps) == 0 {
		return interfaces.Caps{}, fmt.Errorf("no caps provided")
	}

	result := caps[0]

	for i := 1; i < len(caps); i++ {
		result = interfaces.Caps{
			Min: max(result.Min, caps[i].Min),
			Max: min(result.Max, caps[i].Max),
		}
	}

	// Ensure min <= max
	if result.Min > result.Max {
		result.Min = result.Max
	}

	return result, nil
}

// unionCaps unions multiple caps
func (cp *CapsProviderImpl) unionCaps(caps []interfaces.Caps) (interfaces.Caps, error) {
	if len(caps) == 0 {
		return interfaces.Caps{}, fmt.Errorf("no caps provided")
	}

	result := caps[0]

	for i := 1; i < len(caps); i++ {
		result = interfaces.Caps{
			Min: min(result.Min, caps[i].Min),
			Max: max(result.Max, caps[i].Max),
		}
	}

	return result, nil
}

// SetLayerRegistry sets the layer registry
func (cp *CapsProviderImpl) SetLayerRegistry(registry interfaces.CapLayerRegistry) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	cp.layerRegistry = registry
}

// GetLayerRegistry returns the layer registry
func (cp *CapsProviderImpl) GetLayerRegistry() interfaces.CapLayerRegistry {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	return cp.layerRegistry
}

// Validate validates the caps provider
func (cp *CapsProviderImpl) Validate() error {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	if cp.layerRegistry == nil {
		return fmt.Errorf("layer registry is nil")
	}

	return nil
}

// GetSupportedDimensions returns all supported dimensions
func (cp *CapsProviderImpl) GetSupportedDimensions() []string {
	// In a real implementation, this would query the combiner registry
	return []string{
		"strength", "vitality", "dexterity", "intelligence", "spirit", "charisma",
		"hp_max", "mp_max", "stamina_max",
		"attack_power", "defense", "magic_power", "magic_resistance",
		"crit_rate", "crit_damage", "accuracy",
		"move_speed", "attack_speed", "cast_speed",
		"cooldown_reduction", "mana_efficiency", "energy_efficiency",
		"learning_rate", "cultivation_speed", "breakthrough_success",
		"lifespan_years", "poise_rank", "stealth", "perception", "luck",
	}
}

// ValidateCaps validates effective caps
func (cp *CapsProviderImpl) ValidateCaps(caps interfaces.EffectiveCaps) error {
	if caps == nil {
		return fmt.Errorf("caps cannot be nil")
	}

	for dimension, cap := range caps {
		if dimension == "" {
			return fmt.Errorf("dimension cannot be empty")
		}

		if cap.Min > cap.Max {
			return fmt.Errorf("min cap (%f) cannot be greater than max cap (%f) for dimension %s", cap.Min, cap.Max, dimension)
		}

		if cap.Min < 0 {
			return fmt.Errorf("min cap (%f) cannot be negative for dimension %s", cap.Min, dimension)
		}
	}

	return nil
}

// GetCapStatistics returns statistics about caps
func (cp *CapsProviderImpl) GetCapStatistics(ctx context.Context, actor *interfaces.Actor, outputs []*interfaces.SubsystemOutput) (map[string]interface{}, error) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	if actor == nil {
		return nil, fmt.Errorf("actor cannot be nil")
	}

	if outputs == nil {
		return nil, fmt.Errorf("outputs cannot be nil")
	}

	stats := make(map[string]interface{})

	// Get caps within each layer
	layerOrder := cp.layerRegistry.GetLayerOrder()
	layerStats := make(map[string]int)

	for _, layer := range layerOrder {
		effectiveCaps, err := cp.EffectiveCapsWithinLayer(ctx, actor, outputs, layer)
		if err != nil {
			continue
		}

		layerStats[layer] = len(effectiveCaps)
	}

	stats["layer_caps"] = layerStats

	// Get total caps across all layers
	totalCaps, err := cp.EffectiveCapsAcrossLayers(ctx, actor, outputs)
	if err != nil {
		return nil, fmt.Errorf("failed to get total caps: %w", err)
	}

	stats["total_caps"] = len(totalCaps)
	stats["across_policy"] = cp.layerRegistry.GetAcrossLayerPolicy()

	return stats, nil
}

// Helper functions
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
