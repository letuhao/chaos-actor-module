package core

import "context"

// StatConsumer defines the interface for consuming stat changes
type StatConsumer interface {
	// OnStatChanged is called when a stat value changes
	OnStatChanged(statName string, oldValue, newValue interface{})

	// OnStatsChanged is called when multiple stats change
	OnStatsChanged(changes map[string]StatChange)

	// OnStatSnapshot is called when a complete stat snapshot is available
	OnStatSnapshot(snapshot *StatSnapshot)

	// IsActive checks if the consumer is active
	IsActive() bool

	// GetPriority returns the consumer priority (higher number = higher priority)
	GetPriority() int

	// GetConsumerID returns the unique consumer ID
	GetConsumerID() string

	// OnStatChangedWithContext is called when a stat value changes with context
	OnStatChangedWithContext(ctx context.Context, statName string, oldValue, newValue interface{})

	// OnStatsChangedWithContext is called when multiple stats change with context
	OnStatsChangedWithContext(ctx context.Context, changes map[string]StatChange)

	// OnStatSnapshotWithContext is called when a complete stat snapshot is available with context
	OnStatSnapshotWithContext(ctx context.Context, snapshot *StatSnapshot)
}

// StatChange represents a change in a stat value
type StatChange struct {
	StatName  string      `json:"stat_name"`
	OldValue  interface{} `json:"old_value"`
	NewValue  interface{} `json:"new_value"`
	Timestamp int64       `json:"timestamp"`
	Source    string      `json:"source"`
}

// StatChangeBatch represents a batch of stat changes
type StatChangeBatch struct {
	Changes   []StatChange `json:"changes"`
	Timestamp int64        `json:"timestamp"`
	Source    string       `json:"source"`
}
