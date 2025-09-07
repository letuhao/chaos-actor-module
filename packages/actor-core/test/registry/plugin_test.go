package registry

import (
	"chaos-actor-module/packages/actor-core/interfaces"
	"chaos-actor-module/packages/actor-core/registry"
	"context"
	"testing"
)

// MockSubsystem for testing
type MockSubsystem struct {
	systemID string
	priority int64
}

func (m *MockSubsystem) SystemID() string {
	return m.systemID
}

func (m *MockSubsystem) Priority() int64 {
	return m.priority
}

func (m *MockSubsystem) Contribute(ctx context.Context, actor *interfaces.Actor) (*interfaces.SubsystemOutput, error) {
	return &interfaces.SubsystemOutput{
		Primary: []interfaces.Contribution{},
		Derived: []interfaces.Contribution{},
		Caps:    []interfaces.CapContribution{},
		Context: make(map[string]interfaces.ModifierPack),
		Meta:    interfaces.SubsystemMeta{},
	}, nil
}

func TestPluginRegistryImpl_Register(t *testing.T) {
	pr := registry.NewPluginRegistry()

	// Test valid registration
	subsystem := &MockSubsystem{
		systemID: "test_system",
		priority: 100,
	}

	err := pr.Register(subsystem)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	// Verify it was registered
	if !pr.HasSubsystem("test_system") {
		t.Error("HasSubsystem() should return true for registered subsystem")
	}
}

func TestPluginRegistryImpl_Register_Invalid(t *testing.T) {
	pr := registry.NewPluginRegistry()

	// Test nil subsystem
	err := pr.Register(nil)
	if err == nil {
		t.Error("Register() should return error for nil subsystem")
	}

	// Test empty system ID
	subsystem := &MockSubsystem{
		systemID: "",
		priority: 100,
	}

	err = pr.Register(subsystem)
	if err == nil {
		t.Error("Register() should return error for empty system ID")
	}

	// Test duplicate registration
	subsystem1 := &MockSubsystem{
		systemID: "test_system",
		priority: 100,
	}

	subsystem2 := &MockSubsystem{
		systemID: "test_system",
		priority: 200,
	}

	err = pr.Register(subsystem1)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	err = pr.Register(subsystem2)
	if err == nil {
		t.Error("Register() should return error for duplicate system ID")
	}
}

func TestPluginRegistryImpl_Unregister(t *testing.T) {
	pr := registry.NewPluginRegistry()

	// Register a subsystem
	subsystem := &MockSubsystem{
		systemID: "test_system",
		priority: 100,
	}

	err := pr.Register(subsystem)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	// Unregister it
	err = pr.Unregister("test_system")
	if err != nil {
		t.Errorf("Unregister() error = %v", err)
	}

	// Verify it's gone
	if pr.HasSubsystem("test_system") {
		t.Error("HasSubsystem() should return false after unregistering")
	}
}

func TestPluginRegistryImpl_Unregister_Invalid(t *testing.T) {
	pr := registry.NewPluginRegistry()

	// Test unregistering non-existent subsystem
	err := pr.Unregister("non_existent")
	if err == nil {
		t.Error("Unregister() should return error for non-existent subsystem")
	}

	// Test empty system ID
	err = pr.Unregister("")
	if err == nil {
		t.Error("Unregister() should return error for empty system ID")
	}
}

func TestPluginRegistryImpl_Get(t *testing.T) {
	pr := registry.NewPluginRegistry()

	// Register a subsystem
	subsystem := &MockSubsystem{
		systemID: "test_system",
		priority: 100,
	}

	err := pr.Register(subsystem)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	// Get it
	retrieved, exists := pr.Get("test_system")
	if !exists {
		t.Error("Get() should return true for existing subsystem")
	}

	if retrieved == nil {
		t.Error("Get() should return non-nil subsystem")
	}

	// Test non-existent subsystem
	_, exists = pr.Get("non_existent")
	if exists {
		t.Error("Get() should return false for non-existent subsystem")
	}
}

func TestPluginRegistryImpl_GetAll(t *testing.T) {
	pr := registry.NewPluginRegistry()

	// Initially should be empty
	all := pr.GetAll()
	if len(all) != 0 {
		t.Errorf("GetAll() length = %v, want 0", len(all))
	}

	// Register some subsystems
	subsystem1 := &MockSubsystem{
		systemID: "system1",
		priority: 100,
	}

	subsystem2 := &MockSubsystem{
		systemID: "system2",
		priority: 200,
	}

	err := pr.Register(subsystem1)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	err = pr.Register(subsystem2)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	// Now should have 2
	all = pr.GetAll()
	if len(all) != 2 {
		t.Errorf("GetAll() length = %v, want 2", len(all))
	}
}

func TestPluginRegistryImpl_GetByPriority(t *testing.T) {
	pr := registry.NewPluginRegistry()

	// Register subsystems with different priorities
	subsystem1 := &MockSubsystem{
		systemID: "low_priority",
		priority: 100,
	}

	subsystem2 := &MockSubsystem{
		systemID: "high_priority",
		priority: 300,
	}

	subsystem3 := &MockSubsystem{
		systemID: "medium_priority",
		priority: 200,
	}

	err := pr.Register(subsystem1)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	err = pr.Register(subsystem2)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	err = pr.Register(subsystem3)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	// Get by priority (should be sorted by priority descending)
	byPriority := pr.GetByPriority()
	if len(byPriority) != 3 {
		t.Errorf("GetByPriority() length = %v, want 3", len(byPriority))
	}

	// Check order (highest priority first)
	if byPriority[0].Priority() != 300 {
		t.Errorf("GetByPriority()[0].Priority() = %v, want 300", byPriority[0].Priority())
	}

	if byPriority[1].Priority() != 200 {
		t.Errorf("GetByPriority()[1].Priority() = %v, want 200", byPriority[1].Priority())
	}

	if byPriority[2].Priority() != 100 {
		t.Errorf("GetByPriority()[2].Priority() = %v, want 100", byPriority[2].Priority())
	}
}

func TestPluginRegistryImpl_Clear(t *testing.T) {
	pr := registry.NewPluginRegistry()

	// Register some subsystems
	subsystem := &MockSubsystem{
		systemID: "test_system",
		priority: 100,
	}

	err := pr.Register(subsystem)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	// Verify it exists
	if !pr.HasSubsystem("test_system") {
		t.Error("HasSubsystem() should return true for registered subsystem")
	}

	// Clear
	pr.Clear()

	// Verify it's gone
	if pr.HasSubsystem("test_system") {
		t.Error("HasSubsystem() should return false after clearing")
	}
}

func TestPluginRegistryImpl_Count(t *testing.T) {
	pr := registry.NewPluginRegistry()

	// Initially should be 0
	count := pr.Count()
	if count != 0 {
		t.Errorf("Count() = %v, want 0", count)
	}

	// Register a subsystem
	subsystem := &MockSubsystem{
		systemID: "test_system",
		priority: 100,
	}

	err := pr.Register(subsystem)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	// Now should be 1
	count = pr.Count()
	if count != 1 {
		t.Errorf("Count() = %v, want 1", count)
	}
}

func TestPluginRegistryImpl_GetSystemIDs(t *testing.T) {
	pr := registry.NewPluginRegistry()

	// Initially should be empty
	ids := pr.GetSystemIDs()
	if len(ids) != 0 {
		t.Errorf("GetSystemIDs() length = %v, want 0", len(ids))
	}

	// Register some subsystems
	subsystem1 := &MockSubsystem{
		systemID: "system1",
		priority: 100,
	}

	subsystem2 := &MockSubsystem{
		systemID: "system2",
		priority: 200,
	}

	err := pr.Register(subsystem1)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	err = pr.Register(subsystem2)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	// Now should have 2 IDs
	ids = pr.GetSystemIDs()
	if len(ids) != 2 {
		t.Errorf("GetSystemIDs() length = %v, want 2", len(ids))
	}
}

func TestPluginRegistryImpl_IsEmpty(t *testing.T) {
	pr := registry.NewPluginRegistry()

	// Initially should be empty
	if !pr.IsEmpty() {
		t.Error("IsEmpty() should return true for empty registry")
	}

	// Register a subsystem
	subsystem := &MockSubsystem{
		systemID: "test_system",
		priority: 100,
	}

	err := pr.Register(subsystem)
	if err != nil {
		t.Errorf("Register() error = %v", err)
	}

	// Now should not be empty
	if pr.IsEmpty() {
		t.Error("IsEmpty() should return false for non-empty registry")
	}
}
