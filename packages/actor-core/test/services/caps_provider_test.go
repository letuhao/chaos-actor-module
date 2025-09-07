package services

import (
	"chaos-actor-module/packages/actor-core/interfaces"
	"chaos-actor-module/packages/actor-core/registry"
	"chaos-actor-module/packages/actor-core/services"
	"context"
	"testing"
)

func TestCapsProviderImpl_EffectiveCapsWithinLayer(t *testing.T) {
	// Create layer registry
	layerRegistry := registry.NewCapLayerRegistry()

	// Create caps provider
	cp := services.NewCapsProvider(layerRegistry)

	// Create test actor
	actor := &interfaces.Actor{
		ID:      "test_actor",
		Version: 1,
	}

	// Create test outputs
	outputs := []*interfaces.SubsystemOutput{
		{
			Caps: []interfaces.CapContribution{
				{
					System:    "test_system",
					Dimension: "strength",
					Mode:      "BASELINE",
					Value:     100.0,
					Priority:  1000,
					Scope:     "REALM",
					Realm:     "test_realm",
					Tags:      []string{"test"},
				},
			},
		},
	}

	// Test getting caps within layer
	effectiveCaps, err := cp.EffectiveCapsWithinLayer(context.Background(), actor, outputs, "REALM")
	if err != nil {
		t.Errorf("EffectiveCapsWithinLayer() error = %v", err)
	}

	if len(effectiveCaps) != 1 {
		t.Errorf("EffectiveCapsWithinLayer() length = %v, want 1", len(effectiveCaps))
	}

	if cap, exists := effectiveCaps["strength"]; exists {
		if cap.Min != 100.0 || cap.Max != 100.0 {
			t.Errorf("EffectiveCapsWithinLayer() cap = %v, want {Min: 100.0, Max: 100.0}", cap)
		}
	} else {
		t.Error("EffectiveCapsWithinLayer() should contain strength cap")
	}
}

func TestCapsProviderImpl_EffectiveCapsWithinLayer_InvalidInput(t *testing.T) {
	layerRegistry := registry.NewCapLayerRegistry()
	cp := services.NewCapsProvider(layerRegistry)

	// Test with nil actor
	_, err := cp.EffectiveCapsWithinLayer(context.Background(), nil, []*interfaces.SubsystemOutput{}, "REALM")
	if err == nil {
		t.Error("EffectiveCapsWithinLayer() should return error for nil actor")
	}

	// Test with nil outputs
	actor := &interfaces.Actor{ID: "test", Version: 1}
	_, err = cp.EffectiveCapsWithinLayer(context.Background(), actor, nil, "REALM")
	if err == nil {
		t.Error("EffectiveCapsWithinLayer() should return error for nil outputs")
	}

	// Test with empty layer
	_, err = cp.EffectiveCapsWithinLayer(context.Background(), actor, []*interfaces.SubsystemOutput{}, "")
	if err == nil {
		t.Error("EffectiveCapsWithinLayer() should return error for empty layer")
	}
}

func TestCapsProviderImpl_EffectiveCapsAcrossLayers(t *testing.T) {
	layerRegistry := registry.NewCapLayerRegistry()
	cp := services.NewCapsProvider(layerRegistry)

	actor := &interfaces.Actor{
		ID:      "test_actor",
		Version: 1,
	}

	outputs := []*interfaces.SubsystemOutput{
		{
			Caps: []interfaces.CapContribution{
				{
					System:    "test_system",
					Dimension: "strength",
					Mode:      "BASELINE",
					Value:     100.0,
					Priority:  1000,
					Scope:     "REALM",
					Realm:     "test_realm",
					Tags:      []string{"test"},
				},
			},
		},
	}

	effectiveCaps, err := cp.EffectiveCapsAcrossLayers(context.Background(), actor, outputs)
	if err != nil {
		t.Errorf("EffectiveCapsAcrossLayers() error = %v", err)
	}

	if len(effectiveCaps) != 1 {
		t.Errorf("EffectiveCapsAcrossLayers() length = %v, want 1", len(effectiveCaps))
	}
}

func TestCapsProviderImpl_GetLayerOrder(t *testing.T) {
	layerRegistry := registry.NewCapLayerRegistry()
	cp := services.NewCapsProvider(layerRegistry)

	order := cp.GetLayerOrder()
	if len(order) == 0 {
		t.Error("GetLayerOrder() should return non-empty order")
	}
}

func TestCapsProviderImpl_GetAcrossLayerPolicy(t *testing.T) {
	layerRegistry := registry.NewCapLayerRegistry()
	cp := services.NewCapsProvider(layerRegistry)

	policy := cp.GetAcrossLayerPolicy()
	if policy == "" {
		t.Error("GetAcrossLayerPolicy() should return non-empty policy")
	}
}

func TestCapsProviderImpl_GetCapsForDimension(t *testing.T) {
	layerRegistry := registry.NewCapLayerRegistry()
	cp := services.NewCapsProvider(layerRegistry)

	caps, err := cp.GetCapsForDimension("strength")
	if err != nil {
		t.Errorf("GetCapsForDimension() error = %v", err)
	}

	if caps.Min < 0 || caps.Max <= caps.Min {
		t.Errorf("GetCapsForDimension() returned invalid caps: %v", caps)
	}
}

func TestCapsProviderImpl_GetCapsForDimension_Invalid(t *testing.T) {
	layerRegistry := registry.NewCapLayerRegistry()
	cp := services.NewCapsProvider(layerRegistry)

	_, err := cp.GetCapsForDimension("")
	if err == nil {
		t.Error("GetCapsForDimension() should return error for empty dimension")
	}
}

func TestCapsProviderImpl_ValidateCaps(t *testing.T) {
	layerRegistry := registry.NewCapLayerRegistry()
	cp := services.NewCapsProvider(layerRegistry)

	// Test valid caps
	validCaps := interfaces.EffectiveCaps{
		"strength": interfaces.Caps{Min: 0.0, Max: 100.0},
		"vitality": interfaces.Caps{Min: 10.0, Max: 200.0},
	}

	err := cp.ValidateCaps(validCaps)
	if err != nil {
		t.Errorf("ValidateCaps() error = %v", err)
	}

	// Test invalid caps (min > max)
	invalidCaps := interfaces.EffectiveCaps{
		"strength": interfaces.Caps{Min: 100.0, Max: 50.0},
	}

	err = cp.ValidateCaps(invalidCaps)
	if err == nil {
		t.Error("ValidateCaps() should return error for invalid caps")
	}

	// Test nil caps
	err = cp.ValidateCaps(nil)
	if err == nil {
		t.Error("ValidateCaps() should return error for nil caps")
	}
}

func TestCapsProviderImpl_GetSupportedDimensions(t *testing.T) {
	layerRegistry := registry.NewCapLayerRegistry()
	cp := services.NewCapsProvider(layerRegistry)

	dimensions := cp.GetSupportedDimensions()
	if len(dimensions) == 0 {
		t.Error("GetSupportedDimensions() should return non-empty list")
	}

	// Check for some expected dimensions
	expectedDimensions := []string{"strength", "vitality", "hp_max", "attack_power"}
	for _, expected := range expectedDimensions {
		found := false
		for _, dim := range dimensions {
			if dim == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("GetSupportedDimensions() should contain %s", expected)
		}
	}
}

func TestCapsProviderImpl_GetCapStatistics(t *testing.T) {
	layerRegistry := registry.NewCapLayerRegistry()
	cp := services.NewCapsProvider(layerRegistry)

	actor := &interfaces.Actor{
		ID:      "test_actor",
		Version: 1,
	}

	outputs := []*interfaces.SubsystemOutput{
		{
			Caps: []interfaces.CapContribution{
				{
					System:    "test_system",
					Dimension: "strength",
					Mode:      "BASELINE",
					Value:     100.0,
					Priority:  1000,
					Scope:     "REALM",
					Realm:     "test_realm",
					Tags:      []string{"test"},
				},
			},
		},
	}

	stats, err := cp.GetCapStatistics(context.Background(), actor, outputs)
	if err != nil {
		t.Errorf("GetCapStatistics() error = %v", err)
	}

	if stats == nil {
		t.Error("GetCapStatistics() should return non-nil stats")
	}
}

func TestCapsProviderImpl_Validate(t *testing.T) {
	layerRegistry := registry.NewCapLayerRegistry()
	cp := services.NewCapsProvider(layerRegistry)

	err := cp.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v", err)
	}
}
