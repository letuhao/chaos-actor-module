package registry

import (
	"chaos-actor-module/packages/actor-core/registry"
	"testing"
)

func TestCapLayerRegistryImpl_GetLayerOrder(t *testing.T) {
	clr := registry.NewCapLayerRegistry()

	order := clr.GetLayerOrder()
	if len(order) == 0 {
		t.Error("GetLayerOrder() should return non-empty order")
	}

	// Check that it returns a copy
	order[0] = "MODIFIED"
	originalOrder := clr.GetLayerOrder()
	if originalOrder[0] == "MODIFIED" {
		t.Error("GetLayerOrder() should return a copy")
	}
}

func TestCapLayerRegistryImpl_GetAcrossLayerPolicy(t *testing.T) {
	clr := registry.NewCapLayerRegistry()

	policy := clr.GetAcrossLayerPolicy()
	if policy == "" {
		t.Error("GetAcrossLayerPolicy() should return non-empty policy")
	}
}

func TestCapLayerRegistryImpl_SetLayerOrder(t *testing.T) {
	clr := registry.NewCapLayerRegistry()

	// Test valid order
	validOrder := []string{"REALM", "WORLD", "EVENT", "GUILD", "TOTAL"}
	err := clr.SetLayerOrder(validOrder)
	if err != nil {
		t.Errorf("SetLayerOrder() error = %v", err)
	}

	// Verify the order was set
	order := clr.GetLayerOrder()
	if len(order) != len(validOrder) {
		t.Errorf("GetLayerOrder() length = %v, want %v", len(order), len(validOrder))
	}

	for i, layer := range order {
		if layer != validOrder[i] {
			t.Errorf("GetLayerOrder()[%d] = %v, want %v", i, layer, validOrder[i])
		}
	}
}

func TestCapLayerRegistryImpl_SetLayerOrder_Invalid(t *testing.T) {
	clr := registry.NewCapLayerRegistry()

	// Test empty order
	err := clr.SetLayerOrder([]string{})
	if err == nil {
		t.Error("SetLayerOrder() should return error for empty order")
	}

	// Test invalid layer
	invalidOrder := []string{"INVALID_LAYER"}
	err = clr.SetLayerOrder(invalidOrder)
	if err == nil {
		t.Error("SetLayerOrder() should return error for invalid layer")
	}

	// Test duplicate layers
	duplicateOrder := []string{"REALM", "REALM"}
	err = clr.SetLayerOrder(duplicateOrder)
	if err == nil {
		t.Error("SetLayerOrder() should return error for duplicate layers")
	}
}

func TestCapLayerRegistryImpl_SetAcrossLayerPolicy(t *testing.T) {
	clr := registry.NewCapLayerRegistry()

	// Test valid policy
	err := clr.SetAcrossLayerPolicy("intersect")
	if err != nil {
		t.Errorf("SetAcrossLayerPolicy() error = %v", err)
	}

	// Verify the policy was set
	policy := clr.GetAcrossLayerPolicy()
	if policy != "intersect" {
		t.Errorf("GetAcrossLayerPolicy() = %v, want intersect", policy)
	}
}

func TestCapLayerRegistryImpl_SetAcrossLayerPolicy_Invalid(t *testing.T) {
	clr := registry.NewCapLayerRegistry()

	// Test invalid policy
	err := clr.SetAcrossLayerPolicy("invalid_policy")
	if err == nil {
		t.Error("SetAcrossLayerPolicy() should return error for invalid policy")
	}
}

func TestCapLayerRegistryImpl_Validate(t *testing.T) {
	clr := registry.NewCapLayerRegistry()

	// Default registry should be valid
	err := clr.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v", err)
	}
}

func TestCapLayerRegistryImpl_GetLayerIndex(t *testing.T) {
	clr := registry.NewCapLayerRegistry()

	// Test existing layer
	index, exists := clr.GetLayerIndex("REALM")
	if !exists {
		t.Error("GetLayerIndex() should return true for existing layer")
	}

	if index < 0 {
		t.Error("GetLayerIndex() should return non-negative index")
	}

	// Test non-existing layer
	_, exists = clr.GetLayerIndex("INVALID_LAYER")
	if exists {
		t.Error("GetLayerIndex() should return false for non-existing layer")
	}
}

func TestCapLayerRegistryImpl_IsLayerInOrder(t *testing.T) {
	clr := registry.NewCapLayerRegistry()

	// Test existing layer
	if !clr.IsLayerInOrder("REALM") {
		t.Error("IsLayerInOrder() should return true for existing layer")
	}

	// Test non-existing layer
	if clr.IsLayerInOrder("INVALID_LAYER") {
		t.Error("IsLayerInOrder() should return false for non-existing layer")
	}
}

func TestCapLayerRegistryImpl_GetLayerCount(t *testing.T) {
	clr := registry.NewCapLayerRegistry()

	count := clr.GetLayerCount()
	if count <= 0 {
		t.Error("GetLayerCount() should return positive count")
	}
}

func TestCapLayerRegistryImpl_Reset(t *testing.T) {
	clr := registry.NewCapLayerRegistry()

	// Modify the registry
	err := clr.SetAcrossLayerPolicy("union")
	if err != nil {
		t.Errorf("SetAcrossLayerPolicy() error = %v", err)
	}

	// Reset
	clr.Reset()

	// Should be back to default
	policy := clr.GetAcrossLayerPolicy()
	if policy != "intersect" {
		t.Errorf("GetAcrossLayerPolicy() after reset = %v, want intersect", policy)
	}
}
