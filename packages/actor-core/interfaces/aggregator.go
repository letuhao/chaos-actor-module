package interfaces

import (
	"context"
	"time"
)

// Aggregator represents the main aggregation engine
type Aggregator interface {
	// Resolve resolves the actor's stats by aggregating all subsystem contributions
	Resolve(ctx context.Context, actor *Actor) (*Snapshot, error)

	// ResolveWithContext resolves with additional context information
	ResolveWithContext(ctx context.Context, actor *Actor, context map[string]interface{}) (*Snapshot, error)

	// ResolveBatch resolves multiple actors in batch
	ResolveBatch(ctx context.Context, actors []*Actor) ([]*Snapshot, error)

	// GetCachedSnapshot returns a cached snapshot if available
	GetCachedSnapshot(actorID string) (*Snapshot, bool)

	// InvalidateCache invalidates the cache for the given actor
	InvalidateCache(actorID string)

	// ClearCache clears all cached snapshots
	ClearCache()

	// GetMetrics returns performance metrics
	GetMetrics() *AggregatorMetrics
}

// AggregatorMetrics represents performance metrics for the aggregator
type AggregatorMetrics struct {
	// TotalRequests is the total number of requests processed
	TotalRequests int64

	// TotalErrors is the total number of errors
	TotalErrors int64

	// AverageProcessingTime is the average processing time
	AverageProcessingTime time.Duration

	// CacheHits is the number of cache hits
	CacheHits int64

	// CacheMisses is the number of cache misses
	CacheMisses int64

	// ActiveActors is the number of active actors
	ActiveActors int64

	// SubsystemsProcessed is the total number of subsystems processed
	SubsystemsProcessed int64

	// LastProcessed is the timestamp of the last processing
	LastProcessed time.Time

	// MemoryUsage is the current memory usage in bytes
	MemoryUsage int64
}

// GetCacheHitRate returns the cache hit rate
func (am *AggregatorMetrics) GetCacheHitRate() float64 {
	total := am.CacheHits + am.CacheMisses
	if total == 0 {
		return 0.0
	}
	return float64(am.CacheHits) / float64(total)
}

// GetErrorRate returns the error rate
func (am *AggregatorMetrics) GetErrorRate() float64 {
	if am.TotalRequests == 0 {
		return 0.0
	}
	return float64(am.TotalErrors) / float64(am.TotalRequests)
}

// GetThroughput returns the requests per second
func (am *AggregatorMetrics) GetThroughput() float64 {
	if am.AverageProcessingTime == 0 {
		return 0.0
	}
	return float64(time.Second) / float64(am.AverageProcessingTime)
}
