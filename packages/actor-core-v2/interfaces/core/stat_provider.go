package core

import "context"

// StatProvider defines the interface for providing stats
type StatProvider interface {
	// GetPrimaryStats returns all primary stats
	GetPrimaryStats() map[string]int64

	// GetDerivedStats returns all derived stats
	GetDerivedStats() map[string]float64

	// GetCustomPrimaryStats returns custom primary stats
	GetCustomPrimaryStats() map[string]int64

	// GetCustomDerivedStats returns custom derived stats
	GetCustomDerivedStats() map[string]float64

	// GetSubSystemStats returns subsystem-specific stats
	GetSubSystemStats(systemName string) map[string]float64

	// GetAllStats returns a complete stat snapshot
	GetAllStats() *StatSnapshot

	// HasStat checks if a stat exists
	HasStat(statName string) bool

	// GetStatValue gets a specific stat value
	GetStatValue(statName string) (interface{}, error)

	// GetStatValueWithContext gets a specific stat value with context
	GetStatValueWithContext(ctx context.Context, statName string) (interface{}, error)
}

// StatSnapshot represents a complete snapshot of all stats
type StatSnapshot struct {
	PrimaryStats       map[string]int64              `json:"primary_stats"`
	DerivedStats       map[string]float64            `json:"derived_stats"`
	CustomPrimaryStats map[string]int64              `json:"custom_primary_stats"`
	CustomDerivedStats map[string]float64            `json:"custom_derived_stats"`
	SubSystemStats     map[string]map[string]float64 `json:"subsystem_stats"`
	Timestamp          int64                         `json:"timestamp"`
	Version            int64                         `json:"version"`
	Hash               string                        `json:"hash"`
}
