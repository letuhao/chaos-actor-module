package unit

import (
	"testing"
	"time"

	"actor-core/services/cache"
)

func TestNewStateTracker(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	if st == nil {
		t.Error("Expected StateTracker to be created")
	}

	stats := st.GetStats()
	if stats == nil {
		t.Error("Expected stats to be initialized")
	}

	if stats.TotalChanges != 0 {
		t.Errorf("Expected total changes to be 0, got %d", stats.TotalChanges)
	}
}

func TestStateTrackerTrackChange(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Test tracking a valid change
	change := &cache.StateChange{
		EntityType:  "primary_core",
		EntityID:    "actor:123",
		ChangeType:  "update",
		OldValue:    map[string]interface{}{"vitality": 100},
		NewValue:    map[string]interface{}{"vitality": 150},
		UserID:      "user:456",
		SystemID:    "system:789",
		Description: "Updated vitality stat",
	}

	err := st.TrackChange(change)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify change was tracked
	stats := st.GetStats()
	if stats.TotalChanges != 1 {
		t.Errorf("Expected total changes to be 1, got %d", stats.TotalChanges)
	}

	if stats.ActiveChanges != 1 {
		t.Errorf("Expected active changes to be 1, got %d", stats.ActiveChanges)
	}
}

func TestStateTrackerTrackChangeValidation(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Test tracking invalid change (nil)
	err := st.TrackChange(nil)
	if err == nil {
		t.Error("Expected error for nil change")
	}

	// Test tracking invalid change (empty entity type)
	change := &cache.StateChange{
		EntityType: "",
		EntityID:   "actor:123",
		ChangeType: "update",
	}

	err = st.TrackChange(change)
	if err == nil {
		t.Error("Expected error for empty entity type")
	}

	// Test tracking invalid change (empty entity ID)
	change = &cache.StateChange{
		EntityType: "primary_core",
		EntityID:   "",
		ChangeType: "update",
	}

	err = st.TrackChange(change)
	if err == nil {
		t.Error("Expected error for empty entity ID")
	}

	// Test tracking invalid change (invalid change type)
	change = &cache.StateChange{
		EntityType: "primary_core",
		EntityID:   "actor:123",
		ChangeType: "invalid",
	}

	err = st.TrackChange(change)
	if err == nil {
		t.Error("Expected error for invalid change type")
	}
}

func TestStateTrackerGetChanges(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Track multiple changes
	changes := []*cache.StateChange{
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "create",
			UserID:     "user:456",
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "update",
			UserID:     "user:456",
		},
		{
			EntityType: "derived_stats",
			EntityID:   "actor:123",
			ChangeType: "update",
			UserID:     "user:456",
		},
	}

	for _, change := range changes {
		st.TrackChange(change)
	}

	// Get changes for specific entity
	entityChanges, err := st.GetChanges("primary_core", "actor:123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(entityChanges) != 2 {
		t.Errorf("Expected 2 changes for primary_core:actor:123, got %d", len(entityChanges))
	}

	// Get changes for different entity
	entityChanges, err = st.GetChanges("derived_stats", "actor:123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(entityChanges) != 1 {
		t.Errorf("Expected 1 change for derived_stats:actor:123, got %d", len(entityChanges))
	}
}

func TestStateTrackerGetChangesByTimeRange(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	now := time.Now().Unix()

	// Track changes with different timestamps
	changes := []*cache.StateChange{
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "create",
			Timestamp:  now - 3600, // 1 hour ago
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:124", // Different entity ID
			ChangeType: "update",
			Timestamp:  now - 1800, // 30 minutes ago
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:125", // Different entity ID
			ChangeType: "update",
			Timestamp:  now - 900, // 15 minutes ago
		},
	}

	for _, change := range changes {
		st.TrackChange(change)
	}

	// Get changes in last 30 minutes
	recentChanges, err := st.GetChangesByTimeRange(now-1800, now)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(recentChanges) != 2 {
		t.Errorf("Expected 2 changes in last 30 minutes, got %d", len(recentChanges))
	}
}

func TestStateTrackerGetChangesByUser(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Track changes by different users
	changes := []*cache.StateChange{
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "create",
			UserID:     "user:456",
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:124",
			ChangeType: "create",
			UserID:     "user:456",
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:125",
			ChangeType: "create",
			UserID:     "user:789",
		},
	}

	for _, change := range changes {
		st.TrackChange(change)
	}

	// Get changes by user
	userChanges, err := st.GetChangesByUser("user:456")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(userChanges) != 2 {
		t.Errorf("Expected 2 changes for user:456, got %d", len(userChanges))
	}
}

func TestStateTrackerRollbackChange(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Track a change
	change := &cache.StateChange{
		EntityType: "primary_core",
		EntityID:   "actor:123",
		ChangeType: "update",
		OldValue:   map[string]interface{}{"vitality": 100},
		NewValue:   map[string]interface{}{"vitality": 150},
		UserID:     "user:456",
	}

	err := st.TrackChange(change)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Rollback the change
	err = st.RollbackChange(change.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify rollback was tracked
	stats := st.GetStats()
	if stats.Rollbacks != 1 {
		t.Errorf("Expected rollbacks to be 1, got %d", stats.Rollbacks)
	}
}

func TestStateTrackerRollbackDisabled(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          false, // Disabled
		EnableConflictDetection: true,
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Track a change
	change := &cache.StateChange{
		EntityType: "primary_core",
		EntityID:   "actor:123",
		ChangeType: "update",
		UserID:     "user:456",
	}

	st.TrackChange(change)

	// Try to rollback (should fail)
	err := st.RollbackChange(change.ID)
	if err == nil {
		t.Error("Expected error for rollback when disabled")
	}
}

func TestStateTrackerClearChanges(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Track multiple changes
	changes := []*cache.StateChange{
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "create",
			UserID:     "user:456",
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "update",
			UserID:     "user:456",
		},
		{
			EntityType: "derived_stats",
			EntityID:   "actor:123",
			ChangeType: "update",
			UserID:     "user:456",
		},
	}

	for _, change := range changes {
		st.TrackChange(change)
	}

	// Clear changes for specific entity
	err := st.ClearChanges("primary_core", "actor:123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify changes were cleared
	entityChanges, err := st.GetChanges("primary_core", "actor:123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(entityChanges) != 0 {
		t.Errorf("Expected 0 changes for primary_core:actor:123, got %d", len(entityChanges))
	}

	// Verify other entity changes still exist
	entityChanges, err = st.GetChanges("derived_stats", "actor:123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(entityChanges) != 1 {
		t.Errorf("Expected 1 change for derived_stats:actor:123, got %d", len(entityChanges))
	}
}

func TestStateTrackerClearChangesByTimeRange(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	now := time.Now().Unix()

	// Track changes with different timestamps
	changes := []*cache.StateChange{
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "create",
			Timestamp:  now - 3600, // 1 hour ago
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "update",
			Timestamp:  now - 1800, // 30 minutes ago
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "update",
			Timestamp:  now - 900, // 15 minutes ago
		},
	}

	for _, change := range changes {
		st.TrackChange(change)
	}

	// Clear changes in last 30 minutes
	err := st.ClearChangesByTimeRange(now-1800, now)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify recent changes were cleared
	recentChanges, err := st.GetChangesByTimeRange(now-1800, now)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(recentChanges) != 0 {
		t.Errorf("Expected 0 changes in last 30 minutes, got %d", len(recentChanges))
	}

	// Verify old changes still exist
	oldChanges, err := st.GetChangesByTimeRange(now-3600, now-1800)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(oldChanges) != 1 {
		t.Errorf("Expected 1 change in old range, got %d", len(oldChanges))
	}
}

func TestStateTrackerGetStats(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Track some changes
	changes := []*cache.StateChange{
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "create",
			UserID:     "user:456",
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "update",
			UserID:     "user:456",
		},
	}

	for _, change := range changes {
		st.TrackChange(change)
	}

	// Get stats
	stats := st.GetStats()
	if stats.TotalChanges != 2 {
		t.Errorf("Expected total changes to be 2, got %d", stats.TotalChanges)
	}

	if stats.ActiveChanges != 2 {
		t.Errorf("Expected active changes to be 2, got %d", stats.ActiveChanges)
	}
}

func TestStateTrackerResetStats(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Track some changes
	change := &cache.StateChange{
		EntityType: "primary_core",
		EntityID:   "actor:123",
		ChangeType: "create",
		UserID:     "user:456",
	}

	st.TrackChange(change)

	// Reset stats
	err := st.ResetStats()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify stats are reset
	stats := st.GetStats()
	if stats.TotalChanges != 0 {
		t.Errorf("Expected total changes to be 0, got %d", stats.TotalChanges)
	}
}

func TestStateTrackerGetChange(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Track a change
	change := &cache.StateChange{
		EntityType: "primary_core",
		EntityID:   "actor:123",
		ChangeType: "create",
		UserID:     "user:456",
	}

	st.TrackChange(change)

	// Get the change
	retrievedChange, err := st.GetChange(change.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if retrievedChange.EntityType != change.EntityType {
		t.Errorf("Expected entity type to be %s, got %s", change.EntityType, retrievedChange.EntityType)
	}

	// Try to get non-existent change
	_, err = st.GetChange("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent change")
	}
}

func TestStateTrackerGetChangesByType(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Track changes of different types
	changes := []*cache.StateChange{
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "create",
			UserID:     "user:456",
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:124",
			ChangeType: "create",
			UserID:     "user:456",
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "update",
			UserID:     "user:456",
		},
	}

	for _, change := range changes {
		st.TrackChange(change)
	}

	// Get changes by type
	createChanges, err := st.GetChangesByType("create")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(createChanges) != 2 {
		t.Errorf("Expected 2 create changes, got %d", len(createChanges))
	}

	updateChanges, err := st.GetChangesByType("update")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(updateChanges) != 1 {
		t.Errorf("Expected 1 update change, got %d", len(updateChanges))
	}
}

func TestStateTrackerGetChangeCount(t *testing.T) {
	config := cache.StateTrackerConfig{
		EnableTracking:          true,
		MaxHistorySize:          1000,
		RetentionPeriod:         24 * time.Hour,
		EnableRollback:          true,
		EnableConflictDetection: false, // Disable for testing
		CleanupInterval:         1 * time.Hour,
	}

	st := cache.NewStateTracker(config)

	// Initial count should be 0
	if st.GetChangeCount() != 0 {
		t.Errorf("Expected initial change count to be 0, got %d", st.GetChangeCount())
	}

	// Track some changes
	changes := []*cache.StateChange{
		{
			EntityType: "primary_core",
			EntityID:   "actor:123",
			ChangeType: "create",
			UserID:     "user:456",
		},
		{
			EntityType: "primary_core",
			EntityID:   "actor:124",
			ChangeType: "create",
			UserID:     "user:456",
		},
	}

	for _, change := range changes {
		st.TrackChange(change)
	}

	// Count should be 2
	if st.GetChangeCount() != 2 {
		t.Errorf("Expected change count to be 2, got %d", st.GetChangeCount())
	}
}
