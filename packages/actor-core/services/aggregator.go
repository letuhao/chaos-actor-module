package services

import (
	"chaos-actor-module/packages/actor-core/interfaces"
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
)

// AggregatorImpl implements the Aggregator interface
type AggregatorImpl struct {
	combinerRegistry interfaces.CombinerRegistry
	capsProvider     interfaces.CapsProvider
	pluginRegistry   interfaces.PluginRegistry
	cache            interfaces.Cache
	mu               sync.RWMutex
}

// NewAggregator creates a new aggregator
func NewAggregator(
	combinerRegistry interfaces.CombinerRegistry,
	capsProvider interfaces.CapsProvider,
	pluginRegistry interfaces.PluginRegistry,
	cache interfaces.Cache,
) interfaces.Aggregator {
	return &AggregatorImpl{
		combinerRegistry: combinerRegistry,
		capsProvider:     capsProvider,
		pluginRegistry:   pluginRegistry,
		cache:            cache,
	}
}

// Resolve resolves actor stats using the aggregation pipeline
func (a *AggregatorImpl) Resolve(ctx context.Context, actor *interfaces.Actor) (*interfaces.Snapshot, error) {
	return a.ResolveWithContext(ctx, actor, nil)
}

// ResolveWithContext resolves actor stats with additional context
func (a *AggregatorImpl) ResolveWithContext(ctx context.Context, actor *interfaces.Actor, context map[string]interface{}) (*interfaces.Snapshot, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if actor == nil {
		return nil, fmt.Errorf("actor cannot be nil")
	}

	// Check cache first
	if a.cache != nil {
		if cached, exists := a.cache.Get(actor.ID); exists {
			if snapshot, ok := cached.(*interfaces.Snapshot); ok {
				return snapshot, nil
			}
		}
	}

	// Get all subsystems
	subsystems := a.pluginRegistry.GetByPriority()

	// Collect subsystem outputs
	outputs := make([]*interfaces.SubsystemOutput, 0, len(subsystems))

	for _, subsystem := range subsystems {
		output, err := subsystem.Contribute(ctx, actor)
		if err != nil {
			// Log error but continue with other subsystems
			continue
		}

		if output != nil {
			outputs = append(outputs, output)
		}
	}

	// Calculate effective caps
	effectiveCaps, err := a.capsProvider.EffectiveCapsAcrossLayers(ctx, actor, outputs)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate effective caps: %w", err)
	}

	// Aggregate primary stats
	primaryStats, err := a.aggregatePrimaryStats(outputs, effectiveCaps)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate primary stats: %w", err)
	}

	// Aggregate derived stats
	derivedStats, err := a.aggregateDerivedStats(outputs, primaryStats, effectiveCaps)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate derived stats: %w", err)
	}

	// Create snapshot
	snapshot := &interfaces.Snapshot{
		ActorID:   actor.ID,
		Primary:   primaryStats,
		Derived:   derivedStats,
		CapsUsed:  effectiveCaps,
		Version:   actor.Version,
		CreatedAt: time.Now(),
	}

	// Cache the result
	if a.cache != nil {
		a.cache.Set(actor.ID, snapshot, "1h")
	}

	return snapshot, nil
}

// ResolveBatch resolves multiple actors
func (a *AggregatorImpl) ResolveBatch(ctx context.Context, actors []*interfaces.Actor) ([]*interfaces.Snapshot, error) {
	if actors == nil {
		return nil, fmt.Errorf("actors cannot be nil")
	}

	snapshots := make([]*interfaces.Snapshot, 0, len(actors))

	for _, actor := range actors {
		if actor == nil {
			continue
		}

		snapshot, err := a.Resolve(ctx, actor)
		if err != nil {
			// Log error but continue with other actors
			continue
		}

		snapshots = append(snapshots, snapshot)
	}

	return snapshots, nil
}

// GetCachedSnapshot returns a cached snapshot
func (a *AggregatorImpl) GetCachedSnapshot(actorID string) (*interfaces.Snapshot, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.cache == nil {
		return nil, false
	}

	if cached, exists := a.cache.Get(actorID); exists {
		if snapshot, ok := cached.(*interfaces.Snapshot); ok {
			return snapshot, true
		}
	}

	return nil, false
}

// InvalidateCache invalidates cache for an actor
func (a *AggregatorImpl) InvalidateCache(actorID string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.cache != nil {
		a.cache.Delete(actorID)
	}
}

// ClearCache clears all cache
func (a *AggregatorImpl) ClearCache() {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.cache != nil {
		a.cache.Clear()
	}
}

// aggregatePrimaryStats aggregates primary stats from subsystem outputs
func (a *AggregatorImpl) aggregatePrimaryStats(outputs []*interfaces.SubsystemOutput, effectiveCaps interfaces.EffectiveCaps) (map[string]float64, error) {
	// Collect all primary contributions
	contributions := make(map[string][]interfaces.Contribution)

	for _, output := range outputs {
		if output == nil || output.Primary == nil {
			continue
		}

		for _, contribution := range output.Primary {
			contributions[contribution.Dimension] = append(contributions[contribution.Dimension], contribution)
		}
	}

	// Aggregate each dimension
	primaryStats := make(map[string]float64)

	for dimension, contribs := range contributions {
		if len(contribs) == 0 {
			continue
		}

		// Get merge rule for this dimension
		rule, err := a.combinerRegistry.GetRule(dimension)
		if err != nil {
			return nil, fmt.Errorf("failed to get rule for dimension %s: %w", dimension, err)
		}

		// Aggregate contributions
		value, err := a.aggregateContributions(contribs, rule)
		if err != nil {
			return nil, fmt.Errorf("failed to aggregate contributions for dimension %s: %w", dimension, err)
		}

		// Apply caps
		if caps, exists := effectiveCaps[dimension]; exists {
			value = a.applyCaps(value, caps)
		}

		primaryStats[dimension] = value
	}

	return primaryStats, nil
}

// aggregateDerivedStats aggregates derived stats from subsystem outputs
func (a *AggregatorImpl) aggregateDerivedStats(outputs []*interfaces.SubsystemOutput, primaryStats map[string]float64, effectiveCaps interfaces.EffectiveCaps) (map[string]float64, error) {
	// Collect all derived contributions
	contributions := make(map[string][]interfaces.Contribution)

	for _, output := range outputs {
		if output == nil || output.Derived == nil {
			continue
		}

		for _, contribution := range output.Derived {
			contributions[contribution.Dimension] = append(contributions[contribution.Dimension], contribution)
		}
	}

	// Aggregate each dimension
	derivedStats := make(map[string]float64)

	for dimension, contribs := range contributions {
		if len(contribs) == 0 {
			continue
		}

		// Get merge rule for this dimension
		rule, err := a.combinerRegistry.GetRule(dimension)
		if err != nil {
			return nil, fmt.Errorf("failed to get rule for dimension %s: %w", dimension, err)
		}

		// Aggregate contributions
		value, err := a.aggregateContributions(contribs, rule)
		if err != nil {
			return nil, fmt.Errorf("failed to aggregate contributions for dimension %s: %w", dimension, err)
		}

		// Apply caps
		if caps, exists := effectiveCaps[dimension]; exists {
			value = a.applyCaps(value, caps)
		}

		derivedStats[dimension] = value
	}

	return derivedStats, nil
}

// aggregateContributions aggregates contributions for a dimension
func (a *AggregatorImpl) aggregateContributions(contribs []interfaces.Contribution, rule *interfaces.MergeRule) (float64, error) {
	if len(contribs) == 0 {
		return 0.0, fmt.Errorf("no contributions provided")
	}

	// Sort by priority (higher priority first)
	sort.Slice(contribs, func(i, j int) bool {
		return contribs[i].Priority > contribs[j].Priority
	})

	// Group by bucket
	buckets := make(map[string][]interfaces.Contribution)

	for _, contrib := range contribs {
		buckets[contrib.Bucket] = append(buckets[contrib.Bucket], contrib)
	}

	// Process buckets in order
	var result float64

	// FLAT bucket (additive)
	if flatContribs, exists := buckets["FLAT"]; exists {
		for _, contrib := range flatContribs {
			result += contrib.Value
		}
	}

	// MULT bucket (multiplicative)
	if multContribs, exists := buckets["MULT"]; exists {
		for _, contrib := range multContribs {
			result *= contrib.Value
		}
	}

	// POST_ADD bucket (post-additive)
	if postAddContribs, exists := buckets["POST_ADD"]; exists {
		for _, contrib := range postAddContribs {
			result += contrib.Value
		}
	}

	// OVERRIDE bucket (overrides all previous)
	if overrideContribs, exists := buckets["OVERRIDE"]; exists {
		if len(overrideContribs) > 0 {
			// Use the highest priority override
			result = overrideContribs[0].Value
		}
	}

	// EXPONENTIAL bucket (exponential)
	if expContribs, exists := buckets["EXPONENTIAL"]; exists {
		for _, contrib := range expContribs {
			result = result * (1.0 + contrib.Value)
		}
	}

	// LOGARITHMIC bucket (logarithmic)
	if logContribs, exists := buckets["LOGARITHMIC"]; exists {
		for _, contrib := range logContribs {
			result = result * (1.0 + contrib.Value/100.0)
		}
	}

	// CONDITIONAL bucket (conditional)
	if condContribs, exists := buckets["CONDITIONAL"]; exists {
		for _, contrib := range condContribs {
			// For now, just add the value
			// In a real implementation, this would evaluate conditions
			result += contrib.Value
		}
	}

	return result, nil
}

// applyCaps applies caps to a value
func (a *AggregatorImpl) applyCaps(value float64, caps interfaces.Caps) float64 {
	if value < caps.Min {
		return caps.Min
	}

	if value > caps.Max {
		return caps.Max
	}

	return value
}

// SetCombinerRegistry sets the combiner registry
func (a *AggregatorImpl) SetCombinerRegistry(registry interfaces.CombinerRegistry) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.combinerRegistry = registry
}

// SetCapsProvider sets the caps provider
func (a *AggregatorImpl) SetCapsProvider(provider interfaces.CapsProvider) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.capsProvider = provider
}

// SetPluginRegistry sets the plugin registry
func (a *AggregatorImpl) SetPluginRegistry(registry interfaces.PluginRegistry) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.pluginRegistry = registry
}

// SetCache sets the cache
func (a *AggregatorImpl) SetCache(cache interfaces.Cache) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.cache = cache
}

// GetCombinerRegistry returns the combiner registry
func (a *AggregatorImpl) GetCombinerRegistry() interfaces.CombinerRegistry {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.combinerRegistry
}

// GetCapsProvider returns the caps provider
func (a *AggregatorImpl) GetCapsProvider() interfaces.CapsProvider {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.capsProvider
}

// GetPluginRegistry returns the plugin registry
func (a *AggregatorImpl) GetPluginRegistry() interfaces.PluginRegistry {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.pluginRegistry
}

// GetCache returns the cache
func (a *AggregatorImpl) GetCache() interfaces.Cache {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.cache
}

// Validate validates the aggregator
func (a *AggregatorImpl) Validate() error {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.combinerRegistry == nil {
		return fmt.Errorf("combiner registry is nil")
	}

	if a.capsProvider == nil {
		return fmt.Errorf("caps provider is nil")
	}

	if a.pluginRegistry == nil {
		return fmt.Errorf("plugin registry is nil")
	}

	return nil
}

// GetMetrics returns aggregator metrics
func (a *AggregatorImpl) GetMetrics() *interfaces.AggregatorMetrics {
	a.mu.RLock()
	defer a.mu.RUnlock()

	metrics := &interfaces.AggregatorMetrics{
		TotalRequests:         0, // This would be tracked in a real implementation
		TotalErrors:           0,
		AverageProcessingTime: 0,
		CacheHits:             0,
		CacheMisses:           0,
		ActiveActors:          0,
		MemoryUsage:           0,
	}

	if a.cache != nil {
		cacheStats := a.cache.GetStats()
		metrics.CacheHits = cacheStats.Hits
		metrics.CacheMisses = cacheStats.Misses
		metrics.MemoryUsage = cacheStats.MemoryUsage
	}

	return metrics
}
