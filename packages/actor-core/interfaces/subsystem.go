package interfaces

import (
	"context"
	"time"
)

// Subsystem represents a subsystem that can contribute to actor stats
type Subsystem interface {
	// SystemID returns the unique identifier for this subsystem
	SystemID() string

	// Priority returns the processing priority for this subsystem
	Priority() int64

	// Contribute calculates and returns contributions for the given actor
	Contribute(ctx context.Context, actor *Actor) (*SubsystemOutput, error)
}

// ConfigurableSubsystem represents a subsystem that can be configured
type ConfigurableSubsystem interface {
	// Configure configures the subsystem with the given configuration
	Configure(config map[string]interface{}) error
}

// ValidatingSubsystem represents a subsystem that can validate actors
type ValidatingSubsystem interface {
	// Validate validates the given actor for this subsystem
	Validate(actor *Actor) error
}

// CachingSubsystem represents a subsystem that supports caching
type CachingSubsystem interface {
	// GetCacheKey returns the cache key for the given actor
	GetCacheKey(actor *Actor) string

	// ShouldCache returns whether this subsystem should use caching
	ShouldCache() bool
}

// LifecycleSubsystem represents a subsystem with lifecycle management
type LifecycleSubsystem interface {
	// Initialize initializes the subsystem
	Initialize() error

	// Shutdown shuts down the subsystem
	Shutdown() error
}

// EventDrivenSubsystem represents a subsystem that can handle events
type EventDrivenSubsystem interface {
	// RegisterHandler registers an event handler
	RegisterHandler(eventType string, handler func(*Actor, interface{}) error)

	// HandleEvent handles an event for the given actor
	HandleEvent(actor *Actor, eventType string, data interface{}) error
}

// StatefulSubsystem represents a subsystem that maintains state
type StatefulSubsystem interface {
	// GetState returns the state value for the given key
	GetState(key string) (interface{}, bool)

	// SetState sets the state value for the given key
	SetState(key string, value interface{})

	// ClearState clears all state
	ClearState()
}

// ConditionalSubsystem represents a subsystem that can be conditionally active
type ConditionalSubsystem interface {
	// IsActive returns whether the subsystem is active for the given actor
	IsActive(actor *Actor) bool
}

// PerformanceSubsystem represents a subsystem that provides performance metrics
type PerformanceSubsystem interface {
	// GetMetrics returns performance metrics for this subsystem
	GetMetrics() *SubsystemMetrics
}

// SubsystemMetrics represents performance metrics for a subsystem
type SubsystemMetrics struct {
	// ProcessingTime is the total processing time
	ProcessingTime time.Duration

	// CacheHits is the number of cache hits
	CacheHits int64

	// CacheMisses is the number of cache misses
	CacheMisses int64

	// Errors is the number of errors
	Errors int64

	// LastProcessed is the timestamp of the last processing
	LastProcessed time.Time

	// MemoryUsage is the current memory usage in bytes
	MemoryUsage int64
}

// GetCacheHitRate returns the cache hit rate
func (sm *SubsystemMetrics) GetCacheHitRate() float64 {
	total := sm.CacheHits + sm.CacheMisses
	if total == 0 {
		return 0.0
	}
	return float64(sm.CacheHits) / float64(total)
}

// GetErrorRate returns the error rate
func (sm *SubsystemMetrics) GetErrorRate() float64 {
	total := sm.CacheHits + sm.CacheMisses
	if total == 0 {
		return 0.0
	}
	return float64(sm.Errors) / float64(total)
}
