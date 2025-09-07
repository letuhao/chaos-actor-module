package unit

import (
	"testing"
	"time"

	"actor-core-v2/models/flexible"
)

func TestNewFlexibleStats(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	if fs == nil {
		t.Error("Expected FlexibleStats to be created")
	}

	if fs.CustomPrimary == nil {
		t.Error("Expected CustomPrimary to be initialized")
	}

	if fs.CustomDerived == nil {
		t.Error("Expected CustomDerived to be initialized")
	}

	if fs.SubSystemStats == nil {
		t.Error("Expected SubSystemStats to be initialized")
	}

	if fs.Version != 1 {
		t.Errorf("Expected version to be 1, got %d", fs.Version)
	}

	if fs.CreatedAt == 0 {
		t.Error("Expected CreatedAt to be set")
	}

	if fs.UpdatedAt == 0 {
		t.Error("Expected UpdatedAt to be set")
	}
}

func TestSetCustomPrimary(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test setting a custom primary stat
	fs.SetCustomPrimary("strength", 100)

	value, exists := fs.GetCustomPrimary("strength")
	if !exists {
		t.Error("Expected stat to exist")
	}

	if value != 100 {
		t.Errorf("Expected value to be 100, got %d", value)
	}

	// Test that version was incremented
	if fs.Version != 2 {
		t.Errorf("Expected version to be 2, got %d", fs.Version)
	}
}

func TestGetCustomPrimary(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test getting non-existent stat
	_, exists := fs.GetCustomPrimary("non_existent")
	if exists {
		t.Error("Expected stat to not exist")
	}

	// Test getting existing stat
	fs.SetCustomPrimary("strength", 100)
	value, exists := fs.GetCustomPrimary("strength")
	if !exists {
		t.Error("Expected stat to exist")
	}

	if value != 100 {
		t.Errorf("Expected value to be 100, got %d", value)
	}
}

func TestSetCustomDerived(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test setting a custom derived stat
	fs.SetCustomDerived("damage", 150.5)

	value, exists := fs.GetCustomDerived("damage")
	if !exists {
		t.Error("Expected stat to exist")
	}

	if value != 150.5 {
		t.Errorf("Expected value to be 150.5, got %f", value)
	}

	// Test that version was incremented
	if fs.Version != 2 {
		t.Errorf("Expected version to be 2, got %d", fs.Version)
	}
}

func TestGetCustomDerived(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test getting non-existent stat
	_, exists := fs.GetCustomDerived("non_existent")
	if exists {
		t.Error("Expected stat to not exist")
	}

	// Test getting existing stat
	fs.SetCustomDerived("damage", 150.5)
	value, exists := fs.GetCustomDerived("damage")
	if !exists {
		t.Error("Expected stat to exist")
	}

	if value != 150.5 {
		t.Errorf("Expected value to be 150.5, got %f", value)
	}
}

func TestSetSubSystemStat(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test setting a sub-system stat
	fs.SetSubSystemStat("rpg", "level", 10.0)

	value, exists := fs.GetSubSystemStat("rpg", "level")
	if !exists {
		t.Error("Expected stat to exist")
	}

	if value != 10.0 {
		t.Errorf("Expected value to be 10.0, got %f", value)
	}

	// Test that version was incremented
	if fs.Version != 2 {
		t.Errorf("Expected version to be 2, got %d", fs.Version)
	}
}

func TestGetSubSystemStat(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test getting non-existent system stat
	_, exists := fs.GetSubSystemStat("non_existent", "level")
	if exists {
		t.Error("Expected stat to not exist")
	}

	// Test getting existing stat
	fs.SetSubSystemStat("rpg", "level", 10.0)
	value, exists := fs.GetSubSystemStat("rpg", "level")
	if !exists {
		t.Error("Expected stat to exist")
	}

	if value != 10.0 {
		t.Errorf("Expected value to be 10.0, got %f", value)
	}
}

func TestGetAllSubSystemStats(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test getting non-existent system
	_, exists := fs.GetAllSubSystemStats("non_existent")
	if exists {
		t.Error("Expected system to not exist")
	}

	// Test getting existing system
	fs.SetSubSystemStat("rpg", "level", 10.0)
	fs.SetSubSystemStat("rpg", "exp", 1000.0)

	stats, exists := fs.GetAllSubSystemStats("rpg")
	if !exists {
		t.Error("Expected system to exist")
	}

	if len(stats) != 2 {
		t.Errorf("Expected 2 stats, got %d", len(stats))
	}

	if stats["level"] != 10.0 {
		t.Errorf("Expected level to be 10.0, got %f", stats["level"])
	}

	if stats["exp"] != 1000.0 {
		t.Errorf("Expected exp to be 1000.0, got %f", stats["exp"])
	}
}

func TestRemoveCustomPrimary(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test removing non-existent stat
	fs.RemoveCustomPrimary("non_existent")

	// Test removing existing stat
	fs.SetCustomPrimary("strength", 100)
	fs.RemoveCustomPrimary("strength")

	_, exists := fs.GetCustomPrimary("strength")
	if exists {
		t.Error("Expected stat to be removed")
	}

	// Test that version was incremented (1 call to SetCustomPrimary + 1 call to RemoveCustomPrimary = 2 increments)
	if fs.Version != 3 {
		t.Errorf("Expected version to be 3, got %d", fs.Version)
	}
}

func TestRemoveCustomDerived(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test removing non-existent stat
	fs.RemoveCustomDerived("non_existent")

	// Test removing existing stat
	fs.SetCustomDerived("damage", 150.5)
	fs.RemoveCustomDerived("damage")

	_, exists := fs.GetCustomDerived("damage")
	if exists {
		t.Error("Expected stat to be removed")
	}

	// Test that version was incremented (1 call to SetCustomDerived + 1 call to RemoveCustomDerived = 2 increments)
	if fs.Version != 3 {
		t.Errorf("Expected version to be 3, got %d", fs.Version)
	}
}

func TestRemoveSubSystemStat(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test removing non-existent stat
	fs.RemoveSubSystemStat("non_existent", "level")

	// Test removing existing stat
	fs.SetSubSystemStat("rpg", "level", 10.0)
	fs.RemoveSubSystemStat("rpg", "level")

	_, exists := fs.GetSubSystemStat("rpg", "level")
	if exists {
		t.Error("Expected stat to be removed")
	}

	// Test that version was incremented
	if fs.Version != 3 {
		t.Errorf("Expected version to be 3, got %d", fs.Version)
	}
}

func TestRemoveSubSystem(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test removing non-existent system
	fs.RemoveSubSystem("non_existent")

	// Test removing existing system
	fs.SetSubSystemStat("rpg", "level", 10.0)
	fs.SetSubSystemStat("rpg", "exp", 1000.0)
	fs.RemoveSubSystem("rpg")

	_, exists := fs.GetAllSubSystemStats("rpg")
	if exists {
		t.Error("Expected system to be removed")
	}

	// Test that version was incremented (2 calls to SetSubSystemStat + 1 call to RemoveSubSystem = 3 increments)
	if fs.Version != 4 {
		t.Errorf("Expected version to be 4, got %d", fs.Version)
	}
}

func TestGetAllCustomPrimary(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test empty stats
	stats := fs.GetAllCustomPrimary()
	if len(stats) != 0 {
		t.Errorf("Expected 0 stats, got %d", len(stats))
	}

	// Test with stats
	fs.SetCustomPrimary("strength", 100)
	fs.SetCustomPrimary("agility", 80)

	stats = fs.GetAllCustomPrimary()
	if len(stats) != 2 {
		t.Errorf("Expected 2 stats, got %d", len(stats))
	}

	if stats["strength"] != 100 {
		t.Errorf("Expected strength to be 100, got %d", stats["strength"])
	}

	if stats["agility"] != 80 {
		t.Errorf("Expected agility to be 80, got %d", stats["agility"])
	}
}

func TestGetAllCustomDerived(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test empty stats
	stats := fs.GetAllCustomDerived()
	if len(stats) != 0 {
		t.Errorf("Expected 0 stats, got %d", len(stats))
	}

	// Test with stats
	fs.SetCustomDerived("damage", 150.5)
	fs.SetCustomDerived("speed", 120.0)

	stats = fs.GetAllCustomDerived()
	if len(stats) != 2 {
		t.Errorf("Expected 2 stats, got %d", len(stats))
	}

	if stats["damage"] != 150.5 {
		t.Errorf("Expected damage to be 150.5, got %f", stats["damage"])
	}

	if stats["speed"] != 120.0 {
		t.Errorf("Expected speed to be 120.0, got %f", stats["speed"])
	}
}

func TestGetAllSubSystems(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test empty systems
	systems := fs.GetAllSubSystems()
	if len(systems) != 0 {
		t.Errorf("Expected 0 systems, got %d", len(systems))
	}

	// Test with systems
	fs.SetSubSystemStat("rpg", "level", 10.0)
	fs.SetSubSystemStat("magic", "mana", 100.0)

	systems = fs.GetAllSubSystems()
	if len(systems) != 2 {
		t.Errorf("Expected 2 systems, got %d", len(systems))
	}

	// Check that both systems are present
	systemMap := make(map[string]bool)
	for _, system := range systems {
		systemMap[system] = true
	}

	if !systemMap["rpg"] {
		t.Error("Expected rpg system to be present")
	}

	if !systemMap["magic"] {
		t.Error("Expected magic system to be present")
	}
}

func TestGetStatsCount(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test empty stats
	count := fs.GetStatsCount()
	if count != 0 {
		t.Errorf("Expected 0 stats, got %d", count)
	}

	// Test with stats
	fs.SetCustomPrimary("strength", 100)
	fs.SetCustomDerived("damage", 150.5)
	fs.SetSubSystemStat("rpg", "level", 10.0)
	fs.SetSubSystemStat("rpg", "exp", 1000.0)
	fs.SetSubSystemStat("magic", "mana", 100.0)

	count = fs.GetStatsCount()
	if count != 5 {
		t.Errorf("Expected 5 stats, got %d", count)
	}
}

func TestGetCustomPrimaryCount(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test empty stats
	count := fs.GetCustomPrimaryCount()
	if count != 0 {
		t.Errorf("Expected 0 stats, got %d", count)
	}

	// Test with stats
	fs.SetCustomPrimary("strength", 100)
	fs.SetCustomPrimary("agility", 80)

	count = fs.GetCustomPrimaryCount()
	if count != 2 {
		t.Errorf("Expected 2 stats, got %d", count)
	}
}

func TestGetCustomDerivedCount(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test empty stats
	count := fs.GetCustomDerivedCount()
	if count != 0 {
		t.Errorf("Expected 0 stats, got %d", count)
	}

	// Test with stats
	fs.SetCustomDerived("damage", 150.5)
	fs.SetCustomDerived("speed", 120.0)

	count = fs.GetCustomDerivedCount()
	if count != 2 {
		t.Errorf("Expected 2 stats, got %d", count)
	}
}

func TestGetSubSystemCount(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test empty systems
	count := fs.GetSubSystemCount()
	if count != 0 {
		t.Errorf("Expected 0 systems, got %d", count)
	}

	// Test with systems
	fs.SetSubSystemStat("rpg", "level", 10.0)
	fs.SetSubSystemStat("magic", "mana", 100.0)

	count = fs.GetSubSystemCount()
	if count != 2 {
		t.Errorf("Expected 2 systems, got %d", count)
	}
}

func TestGetSubSystemStatsCount(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test non-existent system
	count := fs.GetSubSystemStatsCount("non_existent")
	if count != 0 {
		t.Errorf("Expected 0 stats, got %d", count)
	}

	// Test with stats
	fs.SetSubSystemStat("rpg", "level", 10.0)
	fs.SetSubSystemStat("rpg", "exp", 1000.0)

	count = fs.GetSubSystemStatsCount("rpg")
	if count != 2 {
		t.Errorf("Expected 2 stats, got %d", count)
	}
}

func TestHasCustomPrimary(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test non-existent stat
	if fs.HasCustomPrimary("non_existent") {
		t.Error("Expected stat to not exist")
	}

	// Test existing stat
	fs.SetCustomPrimary("strength", 100)
	if !fs.HasCustomPrimary("strength") {
		t.Error("Expected stat to exist")
	}
}

func TestHasCustomDerived(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test non-existent stat
	if fs.HasCustomDerived("non_existent") {
		t.Error("Expected stat to not exist")
	}

	// Test existing stat
	fs.SetCustomDerived("damage", 150.5)
	if !fs.HasCustomDerived("damage") {
		t.Error("Expected stat to exist")
	}
}

func TestHasSubSystemStat(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test non-existent stat
	if fs.HasSubSystemStat("non_existent", "level") {
		t.Error("Expected stat to not exist")
	}

	// Test existing stat
	fs.SetSubSystemStat("rpg", "level", 10.0)
	if !fs.HasSubSystemStat("rpg", "level") {
		t.Error("Expected stat to exist")
	}
}

func TestHasSubSystem(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test non-existent system
	if fs.HasSubSystem("non_existent") {
		t.Error("Expected system to not exist")
	}

	// Test existing system
	fs.SetSubSystemStat("rpg", "level", 10.0)
	if !fs.HasSubSystem("rpg") {
		t.Error("Expected system to exist")
	}
}

func TestClearCustomPrimary(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test clearing empty stats
	fs.ClearCustomPrimary()

	// Test clearing with stats
	fs.SetCustomPrimary("strength", 100)
	fs.SetCustomPrimary("agility", 80)
	fs.ClearCustomPrimary()

	if len(fs.CustomPrimary) != 0 {
		t.Error("Expected stats to be cleared")
	}

	// Test that version was incremented (2 calls to SetCustomPrimary + 1 call to ClearCustomPrimary = 3 increments)
	if fs.Version != 4 {
		t.Errorf("Expected version to be 4, got %d", fs.Version)
	}
}

func TestClearCustomDerived(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test clearing empty stats
	fs.ClearCustomDerived()

	// Test clearing with stats
	fs.SetCustomDerived("damage", 150.5)
	fs.SetCustomDerived("speed", 120.0)
	fs.ClearCustomDerived()

	if len(fs.CustomDerived) != 0 {
		t.Error("Expected stats to be cleared")
	}

	// Test that version was incremented (2 calls to SetCustomDerived + 1 call to ClearCustomDerived = 3 increments)
	if fs.Version != 4 {
		t.Errorf("Expected version to be 4, got %d", fs.Version)
	}
}

func TestClearSubSystemStats(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test clearing empty stats
	fs.ClearSubSystemStats()

	// Test clearing with stats
	fs.SetSubSystemStat("rpg", "level", 10.0)
	fs.SetSubSystemStat("magic", "mana", 100.0)
	fs.ClearSubSystemStats()

	if len(fs.SubSystemStats) != 0 {
		t.Error("Expected stats to be cleared")
	}

	// Test that version was incremented (2 calls to SetSubSystemStat + 1 call to ClearSubSystemStats = 3 increments)
	if fs.Version != 4 {
		t.Errorf("Expected version to be 4, got %d", fs.Version)
	}
}

func TestClearAll(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test clearing with stats
	fs.SetCustomPrimary("strength", 100)
	fs.SetCustomDerived("damage", 150.5)
	fs.SetSubSystemStat("rpg", "level", 10.0)
	fs.ClearAll()

	if len(fs.CustomPrimary) != 0 {
		t.Error("Expected custom primary stats to be cleared")
	}

	if len(fs.CustomDerived) != 0 {
		t.Error("Expected custom derived stats to be cleared")
	}

	if len(fs.SubSystemStats) != 0 {
		t.Error("Expected sub-system stats to be cleared")
	}

	// Test that version was incremented (3 calls to Set methods + 1 call to ClearAll = 4 increments)
	if fs.Version != 5 {
		t.Errorf("Expected version to be 5, got %d", fs.Version)
	}
}

func TestClone(t *testing.T) {
	fs := flexible.NewFlexibleStats()
	fs.SetCustomPrimary("strength", 100)
	fs.SetCustomDerived("damage", 150.5)
	fs.SetSubSystemStat("rpg", "level", 10.0)

	clone := fs.Clone()

	// Test that values are copied
	if clone.CustomPrimary["strength"] != 100 {
		t.Errorf("Expected strength to be 100, got %d", clone.CustomPrimary["strength"])
	}

	if clone.CustomDerived["damage"] != 150.5 {
		t.Errorf("Expected damage to be 150.5, got %f", clone.CustomDerived["damage"])
	}

	if clone.SubSystemStats["rpg"]["level"] != 10.0 {
		t.Errorf("Expected level to be 10.0, got %f", clone.SubSystemStats["rpg"]["level"])
	}

	// Test that modifying clone doesn't affect original
	clone.SetCustomPrimary("strength", 200)
	if fs.CustomPrimary["strength"] != 100 {
		t.Error("Modifying clone should not affect original")
	}
}

func TestMerge(t *testing.T) {
	fs1 := flexible.NewFlexibleStats()
	fs1.SetCustomPrimary("strength", 100)
	fs1.SetCustomDerived("damage", 150.5)
	fs1.SetSubSystemStat("rpg", "level", 10.0)

	fs2 := flexible.NewFlexibleStats()
	fs2.SetCustomPrimary("agility", 80)
	fs2.SetCustomDerived("speed", 120.0)
	fs2.SetSubSystemStat("magic", "mana", 100.0)

	fs1.Merge(fs2)

	// Test that all stats are merged
	if fs1.CustomPrimary["strength"] != 100 {
		t.Errorf("Expected strength to be 100, got %d", fs1.CustomPrimary["strength"])
	}

	if fs1.CustomPrimary["agility"] != 80 {
		t.Errorf("Expected agility to be 80, got %d", fs1.CustomPrimary["agility"])
	}

	if fs1.CustomDerived["damage"] != 150.5 {
		t.Errorf("Expected damage to be 150.5, got %f", fs1.CustomDerived["damage"])
	}

	if fs1.CustomDerived["speed"] != 120.0 {
		t.Errorf("Expected speed to be 120.0, got %f", fs1.CustomDerived["speed"])
	}

	if fs1.SubSystemStats["rpg"]["level"] != 10.0 {
		t.Errorf("Expected level to be 10.0, got %f", fs1.SubSystemStats["rpg"]["level"])
	}

	if fs1.SubSystemStats["magic"]["mana"] != 100.0 {
		t.Errorf("Expected mana to be 100.0, got %f", fs1.SubSystemStats["magic"]["mana"])
	}

	// Test that version was incremented (3 calls to Set methods + 1 call to Merge = 4 increments)
	if fs1.Version != 5 {
		t.Errorf("Expected version to be 5, got %d", fs1.Version)
	}
}

func TestToJSON(t *testing.T) {
	fs := flexible.NewFlexibleStats()
	fs.SetCustomPrimary("strength", 100)
	fs.SetCustomDerived("damage", 150.5)
	fs.SetSubSystemStat("rpg", "level", 10.0)

	jsonData, err := fs.ToJSON()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(jsonData) == 0 {
		t.Error("Expected JSON data to be non-empty")
	}
}

func TestFromJSON(t *testing.T) {
	fs := flexible.NewFlexibleStats()
	fs.SetCustomPrimary("strength", 100)
	fs.SetCustomDerived("damage", 150.5)
	fs.SetSubSystemStat("rpg", "level", 10.0)

	jsonData, err := fs.ToJSON()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	newFs := flexible.NewFlexibleStats()
	err = newFs.FromJSON(jsonData)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test that values are restored
	if newFs.CustomPrimary["strength"] != 100 {
		t.Errorf("Expected strength to be 100, got %d", newFs.CustomPrimary["strength"])
	}

	if newFs.CustomDerived["damage"] != 150.5 {
		t.Errorf("Expected damage to be 150.5, got %f", newFs.CustomDerived["damage"])
	}

	if newFs.SubSystemStats["rpg"]["level"] != 10.0 {
		t.Errorf("Expected level to be 10.0, got %f", newFs.SubSystemStats["rpg"]["level"])
	}
}

func TestGetVersion(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	if fs.GetVersion() != 1 {
		t.Errorf("Expected version to be 1, got %d", fs.GetVersion())
	}

	fs.SetCustomPrimary("strength", 100)

	if fs.GetVersion() != 2 {
		t.Errorf("Expected version to be 2, got %d", fs.GetVersion())
	}
}

func TestGetUpdatedAt(t *testing.T) {
	fs := flexible.NewFlexibleStats()
	originalUpdatedAt := fs.GetUpdatedAt()

	// Wait a bit to ensure timestamp changes
	time.Sleep(1 * time.Second)

	fs.SetCustomPrimary("strength", 100)

	if fs.GetUpdatedAt() <= originalUpdatedAt {
		t.Errorf("Expected UpdatedAt to be updated. Original: %d, New: %d", originalUpdatedAt, fs.GetUpdatedAt())
	}
}

func TestGetCreatedAt(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	if fs.GetCreatedAt() == 0 {
		t.Error("Expected CreatedAt to be set")
	}
}

func TestValidate(t *testing.T) {
	fs := flexible.NewFlexibleStats()

	// Test valid stats
	fs.SetCustomPrimary("strength", 100)
	fs.SetCustomDerived("damage", 150.5)
	fs.SetSubSystemStat("rpg", "level", 10.0)

	err := fs.Validate()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test invalid stat name (empty)
	fs.CustomPrimary[""] = 100
	err = fs.Validate()
	if err == nil {
		t.Error("Expected error for empty stat name")
	}

	// Test invalid system name (empty)
	fs.SubSystemStats[""] = make(map[string]float64)
	err = fs.Validate()
	if err == nil {
		t.Error("Expected error for empty system name")
	}
}
