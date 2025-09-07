package cache

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// StateTracker implements state change tracking
type StateTracker struct {
	changes map[string]*StateChange
	mutex   sync.RWMutex
	config  StateTrackerConfig
	stats   *StateTrackingStats
	created int64
}

// NewStateTracker creates a new StateTracker instance
func NewStateTracker(config StateTrackerConfig) *StateTracker {
	now := time.Now().Unix()
	return &StateTracker{
		changes: make(map[string]*StateChange),
		config:  config,
		stats:   &StateTrackingStats{LastChange: now},
		created: now,
	}
}

// TrackChange tracks a state change
func (st *StateTracker) TrackChange(change *StateChange) error {
	if !st.config.EnableTracking {
		return nil
	}

	st.mutex.Lock()
	defer st.mutex.Unlock()

	// Validate change
	if err := st.validateChange(change); err != nil {
		st.stats.ErrorRate++
		return fmt.Errorf("invalid change: %w", err)
	}

	// Set timestamps if not set
	if change.Timestamp == 0 {
		change.Timestamp = time.Now().Unix()
	}

	// Generate ID if not set
	if change.ID == "" {
		change.ID = st.generateChangeID(change)
	}

	// Check for conflicts if enabled
	if st.config.EnableConflictDetection {
		if err := st.checkConflict(change); err != nil {
			st.stats.Conflicts++
			return fmt.Errorf("conflict detected: %w", err)
		}
	}

	// Store change
	st.changes[change.ID] = change
	st.stats.TotalChanges++
	st.stats.LastChange = change.Timestamp
	st.stats.ActiveChanges++

	// Cleanup old changes if needed
	if st.config.MaxHistorySize > 0 && int64(len(st.changes)) > st.config.MaxHistorySize {
		st.cleanupOldChanges()
	}

	return nil
}

// GetChanges retrieves changes for a specific entity
func (st *StateTracker) GetChanges(entityType, entityID string) ([]*StateChange, error) {
	st.mutex.RLock()
	defer st.mutex.RUnlock()

	var changes []*StateChange
	for _, change := range st.changes {
		if change.EntityType == entityType && change.EntityID == entityID {
			changes = append(changes, change)
		}
	}

	return changes, nil
}

// GetChangesByTimeRange retrieves changes within a time range
func (st *StateTracker) GetChangesByTimeRange(start, end int64) ([]*StateChange, error) {
	st.mutex.RLock()
	defer st.mutex.RUnlock()

	var changes []*StateChange
	for _, change := range st.changes {
		if change.Timestamp >= start && change.Timestamp <= end {
			changes = append(changes, change)
		}
	}

	return changes, nil
}

// GetChangesByUser retrieves changes by user
func (st *StateTracker) GetChangesByUser(userID string) ([]*StateChange, error) {
	st.mutex.RLock()
	defer st.mutex.RUnlock()

	var changes []*StateChange
	for _, change := range st.changes {
		if change.UserID == userID {
			changes = append(changes, change)
		}
	}

	return changes, nil
}

// RollbackChange rolls back a specific change
func (st *StateTracker) RollbackChange(changeID string) error {
	if !st.config.EnableRollback {
		return fmt.Errorf("rollback is disabled")
	}

	st.mutex.Lock()
	defer st.mutex.Unlock()

	change, exists := st.changes[changeID]
	if !exists {
		return fmt.Errorf("change not found: %s", changeID)
	}

	// Create rollback change
	rollbackChange := &StateChange{
		ID:          st.generateChangeID(change),
		EntityType:  change.EntityType,
		EntityID:    change.EntityID,
		ChangeType:  "rollback",
		OldValue:    change.NewValue,
		NewValue:    change.OldValue,
		Timestamp:   time.Now().Unix(),
		UserID:      change.UserID,
		SystemID:    change.SystemID,
		Description: fmt.Sprintf("Rollback of change %s", changeID),
		Metadata: map[string]interface{}{
			"original_change_id": changeID,
			"rollback_reason":    "manual_rollback",
		},
	}

	// Store rollback change
	st.changes[rollbackChange.ID] = rollbackChange
	st.stats.TotalChanges++
	st.stats.Rollbacks++
	st.stats.LastChange = rollbackChange.Timestamp

	return nil
}

// ClearChanges removes changes for a specific entity
func (st *StateTracker) ClearChanges(entityType, entityID string) error {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	var toDelete []string
	for id, change := range st.changes {
		if change.EntityType == entityType && change.EntityID == entityID {
			toDelete = append(toDelete, id)
		}
	}

	for _, id := range toDelete {
		delete(st.changes, id)
		st.stats.ActiveChanges--
	}

	return nil
}

// ClearChangesByTimeRange removes changes within a time range
func (st *StateTracker) ClearChangesByTimeRange(start, end int64) error {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	var toDelete []string
	for id, change := range st.changes {
		if change.Timestamp >= start && change.Timestamp <= end {
			toDelete = append(toDelete, id)
		}
	}

	for _, id := range toDelete {
		delete(st.changes, id)
		st.stats.ActiveChanges--
	}

	return nil
}

// GetStats returns state tracking statistics
func (st *StateTracker) GetStats() *StateTrackingStats {
	st.mutex.RLock()
	defer st.mutex.RUnlock()

	// Calculate changes per hour
	now := time.Now().Unix()
	hoursElapsed := float64(now-st.created) / 3600.0
	if hoursElapsed > 0 {
		st.stats.ChangesPerHour = float64(st.stats.TotalChanges) / hoursElapsed
	}

	// Calculate error rate
	totalOperations := st.stats.TotalChanges + st.stats.Conflicts
	if totalOperations > 0 {
		st.stats.ErrorRate = float64(st.stats.Conflicts) / float64(totalOperations)
	}

	// Create a copy to avoid race conditions
	stats := *st.stats
	return &stats
}

// ResetStats resets state tracking statistics
func (st *StateTracker) ResetStats() error {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	now := time.Now().Unix()
	st.stats = &StateTrackingStats{
		LastChange: now,
	}
	return nil
}

// Private methods

// validateChange validates a state change
func (st *StateTracker) validateChange(change *StateChange) error {
	if change == nil {
		return fmt.Errorf("change cannot be nil")
	}

	if change.EntityType == "" {
		return fmt.Errorf("entity type cannot be empty")
	}

	if change.EntityID == "" {
		return fmt.Errorf("entity ID cannot be empty")
	}

	if change.ChangeType == "" {
		return fmt.Errorf("change type cannot be empty")
	}

	validChangeTypes := map[string]bool{
		"create":   true,
		"update":   true,
		"delete":   true,
		"rollback": true,
	}

	if !validChangeTypes[change.ChangeType] {
		return fmt.Errorf("invalid change type: %s", change.ChangeType)
	}

	return nil
}

// generateChangeID generates a unique change ID
func (st *StateTracker) generateChangeID(change *StateChange) string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s:%s:%s:%d", change.EntityType, change.EntityID, change.ChangeType, timestamp)
}

// checkConflict checks for conflicts with existing changes
func (st *StateTracker) checkConflict(change *StateChange) error {
	// Check for recent changes to the same entity
	for _, existingChange := range st.changes {
		if existingChange.EntityType == change.EntityType &&
			existingChange.EntityID == change.EntityID &&
			existingChange.ChangeType != "rollback" {

			// Check if changes are too close in time (within 1 second)
			timeDiff := change.Timestamp - existingChange.Timestamp
			if timeDiff >= 0 && timeDiff < 1 {
				return fmt.Errorf("conflict with change %s (time diff: %d)", existingChange.ID, timeDiff)
			}
		}
	}

	return nil
}

// cleanupOldChanges removes old changes based on retention policy
func (st *StateTracker) cleanupOldChanges() {
	if st.config.RetentionPeriod <= 0 {
		return
	}

	cutoffTime := time.Now().Unix() - int64(st.config.RetentionPeriod.Seconds())
	var toDelete []string

	for id, change := range st.changes {
		if change.Timestamp < cutoffTime {
			toDelete = append(toDelete, id)
		}
	}

	for _, id := range toDelete {
		delete(st.changes, id)
		st.stats.ActiveChanges--
	}
}

// StartCleanup starts the cleanup goroutine
func (st *StateTracker) StartCleanup() {
	if st.config.CleanupInterval <= 0 {
		return
	}

	go func() {
		ticker := time.NewTicker(st.config.CleanupInterval)
		defer ticker.Stop()

		for range ticker.C {
			st.cleanupOldChanges()
		}
	}()
}

// GetChange retrieves a specific change by ID
func (st *StateTracker) GetChange(changeID string) (*StateChange, error) {
	st.mutex.RLock()
	defer st.mutex.RUnlock()

	change, exists := st.changes[changeID]
	if !exists {
		return nil, fmt.Errorf("change not found: %s", changeID)
	}

	return change, nil
}

// GetChangesByType retrieves changes by change type
func (st *StateTracker) GetChangesByType(changeType string) ([]*StateChange, error) {
	st.mutex.RLock()
	defer st.mutex.RUnlock()

	var changes []*StateChange
	for _, change := range st.changes {
		if change.ChangeType == changeType {
			changes = append(changes, change)
		}
	}

	return changes, nil
}

// GetChangesBySystem retrieves changes by system ID
func (st *StateTracker) GetChangesBySystem(systemID string) ([]*StateChange, error) {
	st.mutex.RLock()
	defer st.mutex.RUnlock()

	var changes []*StateChange
	for _, change := range st.changes {
		if change.SystemID == systemID {
			changes = append(changes, change)
		}
	}

	return changes, nil
}

// ExportChanges exports changes to JSON
func (st *StateTracker) ExportChanges() ([]byte, error) {
	st.mutex.RLock()
	defer st.mutex.RUnlock()

	changes := make([]*StateChange, 0, len(st.changes))
	for _, change := range st.changes {
		changes = append(changes, change)
	}

	return json.MarshalIndent(changes, "", "  ")
}

// ImportChanges imports changes from JSON
func (st *StateTracker) ImportChanges(data []byte) error {
	var changes []*StateChange
	if err := json.Unmarshal(data, &changes); err != nil {
		return fmt.Errorf("failed to unmarshal changes: %w", err)
	}

	st.mutex.Lock()
	defer st.mutex.Unlock()

	for _, change := range changes {
		if err := st.validateChange(change); err != nil {
			continue // Skip invalid changes
		}
		st.changes[change.ID] = change
	}

	return nil
}

// GetChangeCount returns the total number of changes
func (st *StateTracker) GetChangeCount() int {
	st.mutex.RLock()
	defer st.mutex.RUnlock()
	return len(st.changes)
}

// GetActiveChangeCount returns the number of active changes
func (st *StateTracker) GetActiveChangeCount() int64 {
	st.mutex.RLock()
	defer st.mutex.RUnlock()
	return st.stats.ActiveChanges
}
