package interfaces

import (
	"context"
)

// CapsProvider represents a provider for effective caps
type CapsProvider interface {
	// EffectiveCapsWithinLayer calculates effective caps within a specific layer
	EffectiveCapsWithinLayer(ctx context.Context, actor *Actor, outputs []*SubsystemOutput, layer string) (EffectiveCaps, error)

	// EffectiveCapsAcrossLayers calculates effective caps across all layers
	EffectiveCapsAcrossLayers(ctx context.Context, actor *Actor, outputs []*SubsystemOutput) (EffectiveCaps, error)

	// GetLayerOrder returns the processing order for layers
	GetLayerOrder() []string

	// GetAcrossLayerPolicy returns the across-layer policy
	GetAcrossLayerPolicy() string

	// ValidateCaps validates the given caps
	ValidateCaps(caps EffectiveCaps) error

	// GetCapsForDimension returns caps for a specific dimension
	GetCapsForDimension(dimension string) (Caps, error)

	// GetSupportedDimensions returns all supported dimensions
	GetSupportedDimensions() []string

	// GetCapStatistics returns statistics about caps
	GetCapStatistics(ctx context.Context, actor *Actor, outputs []*SubsystemOutput) (map[string]interface{}, error)

	// Validate validates the caps provider
	Validate() error
}

// Caps represents min/max caps for a dimension
// This is a forward declaration - actual implementation is in types package
type Caps struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

// EffectiveCaps represents effective caps for all dimensions
type EffectiveCaps map[string]Caps

// Methods will be implemented in the actual types package
