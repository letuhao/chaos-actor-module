package registry

import (
	"chaos-actor-module/packages/actor-core/interfaces"
	"chaos-actor-module/packages/actor-core/registry"
	"testing"
)

func TestCombinerRegistryImpl_GetRule(t *testing.T) {
	cr := registry.NewCombinerRegistry()

	// Test getting rule for non-existent dimension
	rule, err := cr.GetRule("strength")
	if err != nil {
		t.Errorf("GetRule() error = %v", err)
	}

	if rule == nil {
		t.Error("GetRule() returned nil rule")
	}

	if !rule.UsePipeline {
		t.Error("GetRule() should return rule with UsePipeline = true")
	}
}

func TestCombinerRegistryImpl_SetRule(t *testing.T) {
	cr := registry.NewCombinerRegistry()

	rule := &interfaces.MergeRule{
		UsePipeline: true,
		ClampDefault: interfaces.Caps{
			Min: 0.0,
			Max: 100.0,
		},
	}

	err := cr.SetRule("strength", rule)
	if err != nil {
		t.Errorf("SetRule() error = %v", err)
	}

	// Verify the rule was set
	retrievedRule, err := cr.GetRule("strength")
	if err != nil {
		t.Errorf("GetRule() error = %v", err)
	}

	if retrievedRule.UsePipeline != rule.UsePipeline {
		t.Errorf("GetRule() UsePipeline = %v, want %v", retrievedRule.UsePipeline, rule.UsePipeline)
	}
}

func TestCombinerRegistryImpl_SetRule_InvalidRule(t *testing.T) {
	cr := registry.NewCombinerRegistry()

	// Test with nil rule
	err := cr.SetRule("strength", nil)
	if err == nil {
		t.Error("SetRule() should return error for nil rule")
	}
}

func TestCombinerRegistryImpl_GetDimensions(t *testing.T) {
	cr := registry.NewCombinerRegistry()

	// Initially should be empty
	dimensions := cr.GetDimensions()
	if len(dimensions) != 0 {
		t.Errorf("GetDimensions() length = %v, want 0", len(dimensions))
	}

	// Add a rule
	rule := &interfaces.MergeRule{
		UsePipeline: true,
		ClampDefault: interfaces.Caps{
			Min: 0.0,
			Max: 100.0,
		},
	}

	err := cr.SetRule("strength", rule)
	if err != nil {
		t.Errorf("SetRule() error = %v", err)
	}

	// Now should have one dimension
	dimensions = cr.GetDimensions()
	if len(dimensions) != 1 {
		t.Errorf("GetDimensions() length = %v, want 1", len(dimensions))
	}

	if dimensions[0] != "strength" {
		t.Errorf("GetDimensions()[0] = %v, want strength", dimensions[0])
	}
}

func TestCombinerRegistryImpl_Validate(t *testing.T) {
	cr := registry.NewCombinerRegistry()

	// Empty registry should be valid
	err := cr.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v", err)
	}

	// Add a valid rule
	rule := &interfaces.MergeRule{
		UsePipeline: true,
		ClampDefault: interfaces.Caps{
			Min: 0.0,
			Max: 100.0,
		},
	}

	err = cr.SetRule("strength", rule)
	if err != nil {
		t.Errorf("SetRule() error = %v", err)
	}

	// Should still be valid
	err = cr.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v", err)
	}
}

func TestCombinerRegistryImpl_Clear(t *testing.T) {
	cr := registry.NewCombinerRegistry()

	// Add a rule
	rule := &interfaces.MergeRule{
		UsePipeline: true,
		ClampDefault: interfaces.Caps{
			Min: 0.0,
			Max: 100.0,
		},
	}

	err := cr.SetRule("strength", rule)
	if err != nil {
		t.Errorf("SetRule() error = %v", err)
	}

	// Verify rule exists
	dimensions := cr.GetDimensions()
	if len(dimensions) != 1 {
		t.Errorf("GetDimensions() length = %v, want 1", len(dimensions))
	}

	// Clear registry
	cr.Clear()

	// Verify it's empty
	dimensions = cr.GetDimensions()
	if len(dimensions) != 0 {
		t.Errorf("GetDimensions() length = %v, want 0", len(dimensions))
	}
}

func TestCombinerRegistryImpl_HasRule(t *testing.T) {
	cr := registry.NewCombinerRegistry()

	// Initially should not have any rules
	if cr.HasRule("strength") {
		t.Error("HasRule() should return false for non-existent rule")
	}

	// Add a rule
	rule := &interfaces.MergeRule{
		UsePipeline: true,
		ClampDefault: interfaces.Caps{
			Min: 0.0,
			Max: 100.0,
		},
	}

	err := cr.SetRule("strength", rule)
	if err != nil {
		t.Errorf("SetRule() error = %v", err)
	}

	// Now should have the rule
	if !cr.HasRule("strength") {
		t.Error("HasRule() should return true for existing rule")
	}
}

func TestCombinerRegistryImpl_RemoveRule(t *testing.T) {
	cr := registry.NewCombinerRegistry()

	// Add a rule
	rule := &interfaces.MergeRule{
		UsePipeline: true,
		ClampDefault: interfaces.Caps{
			Min: 0.0,
			Max: 100.0,
		},
	}

	err := cr.SetRule("strength", rule)
	if err != nil {
		t.Errorf("SetRule() error = %v", err)
	}

	// Verify rule exists
	if !cr.HasRule("strength") {
		t.Error("HasRule() should return true for existing rule")
	}

	// Remove the rule
	cr.RemoveRule("strength")

	// Verify rule is gone
	if cr.HasRule("strength") {
		t.Error("HasRule() should return false after removing rule")
	}
}

func TestCombinerRegistryImpl_Count(t *testing.T) {
	cr := registry.NewCombinerRegistry()

	// Initially should be 0
	count := cr.Count()
	if count != 0 {
		t.Errorf("Count() = %v, want 0", count)
	}

	// Add a rule
	rule := &interfaces.MergeRule{
		UsePipeline: true,
		ClampDefault: interfaces.Caps{
			Min: 0.0,
			Max: 100.0,
		},
	}

	err := cr.SetRule("strength", rule)
	if err != nil {
		t.Errorf("SetRule() error = %v", err)
	}

	// Now should be 1
	count = cr.Count()
	if count != 1 {
		t.Errorf("Count() = %v, want 1", count)
	}
}
