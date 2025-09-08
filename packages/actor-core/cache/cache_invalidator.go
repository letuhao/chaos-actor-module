package cache

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// CacheInvalidator provides intelligent cache invalidation
type CacheInvalidator struct {
	// Configuration
	config *CacheInvalidatorConfig

	// Cache layers
	l1Cache *LockFreeL1Cache
	l2Cache *MemoryMappedL2Cache
	l3Cache *PersistentL3Cache

	// Invalidation strategies
	strategies map[string]InvalidationStrategy

	// Dependency tracking
	dependencies map[string][]string // key -> list of dependent keys
	dependents   map[string][]string // key -> list of keys that depend on it
	depsMutex    sync.RWMutex

	// TTL management
	ttlManager *TTLManager

	// Background workers
	workers    []*InvalidationWorker
	workerPool sync.Pool

	// Statistics
	stats *CacheInvalidatorStats

	// State
	closed int32
	ctx    context.Context
	cancel context.CancelFunc
}

// CacheInvalidatorConfig holds configuration for cache invalidation
type CacheInvalidatorConfig struct {
	// Invalidation settings
	EnableInvalidation     bool
	InvalidationInterval   time.Duration
	MaxInvalidationWorkers int

	// TTL settings
	EnableTTL        bool
	TTLCheckInterval time.Duration
	DefaultTTL       time.Duration
	MaxTTL           time.Duration

	// Dependency settings
	EnableDependencies bool
	MaxDependencyDepth int
	DependencyTimeout  time.Duration

	// Performance settings
	InvalidationPriority   int
	BatchSize              int
	EnableLazyInvalidation bool
	LazyThreshold          int64
}

// InvalidationStrategy defines how to invalidate cache entries
type InvalidationStrategy interface {
	ShouldInvalidate(key string, reason InvalidationReason) bool
	GetInvalidationKeys(key string, reason InvalidationReason) []string
	GetPriority() int
	GetName() string
}

// InvalidationReason represents why a cache entry is being invalidated
type InvalidationReason int

const (
	InvalidationReasonTTL InvalidationReason = iota
	InvalidationReasonExplicit
	InvalidationReasonDependency
	InvalidationReasonUpdate
	InvalidationReasonDelete
	InvalidationReasonMemory
	InvalidationReasonError
)

// TTLManager manages TTL for cache entries
type TTLManager struct {
	// TTL tracking
	ttlEntries map[string]*TTLEntry
	ttlMutex   sync.RWMutex

	// Configuration
	checkInterval time.Duration
	defaultTTL    time.Duration
	maxTTL        time.Duration

	// Statistics
	stats *TTLStats
}

// TTLEntry represents a TTL entry
type TTLEntry struct {
	Key         string
	ExpiresAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AccessCount int64
	LastAccess  time.Time
}

// TTLStats holds TTL statistics
type TTLStats struct {
	TotalEntries   int64
	ExpiredEntries int64
	ActiveEntries  int64
	AverageTTL     time.Duration
	ExpirationRate float64
	LastCleanup    time.Time
}

// InvalidationWorker handles cache invalidation tasks
type InvalidationWorker struct {
	id          int
	invalidator *CacheInvalidator
	workChan    chan *InvalidationTask
	ctx         context.Context
	cancel      context.CancelFunc
}

// InvalidationTask represents a cache invalidation task
type InvalidationTask struct {
	Key       string
	Reason    InvalidationReason
	Priority  int
	Source    string
	CreatedAt time.Time
	Metadata  map[string]interface{}
}

// CacheInvalidatorStats holds statistics for cache invalidation
type CacheInvalidatorStats struct {
	// Invalidation stats
	TotalInvalidations      int64
	SuccessfulInvalidations int64
	FailedInvalidations     int64
	ActiveWorkers           int64

	// TTL stats
	TTLExpirations int64
	TTLCleanups    int64
	AverageTTL     time.Duration

	// Dependency stats
	DependencyInvalidations int64
	DependencyChains        int64
	MaxDependencyDepth      int

	// Performance stats
	AverageInvalidationTime time.Duration
	InvalidationThroughput  float64
	MemoryFreed             int64

	// Resource stats
	MemoryUsage          int64
	CPUUsage             float64
	LastInvalidationTime time.Time
}

// NewCacheInvalidator creates a new cache invalidator
func NewCacheInvalidator(config *CacheInvalidatorConfig, l1Cache *LockFreeL1Cache, l2Cache *MemoryMappedL2Cache, l3Cache *PersistentL3Cache) *CacheInvalidator {
	if config == nil {
		config = &CacheInvalidatorConfig{
			EnableInvalidation:     true,
			InvalidationInterval:   time.Minute * 2,
			MaxInvalidationWorkers: 4,
			EnableTTL:              true,
			TTLCheckInterval:       time.Minute,
			DefaultTTL:             time.Hour,
			MaxTTL:                 time.Hour * 24,
			EnableDependencies:     true,
			MaxDependencyDepth:     10,
			DependencyTimeout:      time.Second * 30,
			InvalidationPriority:   1,
			BatchSize:              100,
			EnableLazyInvalidation: true,
			LazyThreshold:          1000,
		}
	}

	ctx, cancel := context.WithCancel(context.Background())

	invalidator := &CacheInvalidator{
		config:       config,
		l1Cache:      l1Cache,
		l2Cache:      l2Cache,
		l3Cache:      l3Cache,
		strategies:   make(map[string]InvalidationStrategy),
		dependencies: make(map[string][]string),
		dependents:   make(map[string][]string),
		ttlManager:   NewTTLManager(config),
		stats:        &CacheInvalidatorStats{},
		ctx:          ctx,
		cancel:       cancel,
	}

	// Initialize default strategies
	invalidator.initializeStrategies()

	// Initialize workers
	invalidator.initializeWorkers()

	// Start background processes
	if config.EnableInvalidation {
		go invalidator.startBackgroundInvalidation()
	}

	if config.EnableTTL {
		go invalidator.startTTLManagement()
	}

	return invalidator
}

// NewTTLManager creates a new TTL manager
func NewTTLManager(config *CacheInvalidatorConfig) *TTLManager {
	return &TTLManager{
		ttlEntries:    make(map[string]*TTLEntry),
		checkInterval: config.TTLCheckInterval,
		defaultTTL:    config.DefaultTTL,
		maxTTL:        config.MaxTTL,
		stats:         &TTLStats{},
	}
}

// initializeStrategies initializes default invalidation strategies
func (i *CacheInvalidator) initializeStrategies() {
	// TTL-based invalidation
	i.strategies["ttl"] = &TTLInvalidationStrategy{
		ttlManager: i.ttlManager,
		priority:   1,
	}

	// Dependency-based invalidation
	i.strategies["dependency"] = &DependencyInvalidationStrategy{
		invalidator: i,
		priority:    2,
	}

	// Memory-based invalidation
	i.strategies["memory"] = &MemoryInvalidationStrategy{
		threshold: i.config.LazyThreshold,
		priority:  3,
	}

	// Error-based invalidation
	i.strategies["error"] = &ErrorInvalidationStrategy{
		priority: 4,
	}
}

// TTLInvalidationStrategy implements TTL-based invalidation
type TTLInvalidationStrategy struct {
	ttlManager *TTLManager
	priority   int
}

func (s *TTLInvalidationStrategy) ShouldInvalidate(key string, reason InvalidationReason) bool {
	return reason == InvalidationReasonTTL
}

func (s *TTLInvalidationStrategy) GetInvalidationKeys(key string, reason InvalidationReason) []string {
	if reason == InvalidationReasonTTL {
		return []string{key}
	}
	return nil
}

func (s *TTLInvalidationStrategy) GetPriority() int {
	return s.priority
}

func (s *TTLInvalidationStrategy) GetName() string {
	return "ttl"
}

// DependencyInvalidationStrategy implements dependency-based invalidation
type DependencyInvalidationStrategy struct {
	invalidator *CacheInvalidator
	priority    int
}

func (s *DependencyInvalidationStrategy) ShouldInvalidate(key string, reason InvalidationReason) bool {
	return reason == InvalidationReasonDependency || reason == InvalidationReasonUpdate
}

func (s *DependencyInvalidationStrategy) GetInvalidationKeys(key string, reason InvalidationReason) []string {
	if reason == InvalidationReasonDependency || reason == InvalidationReasonUpdate {
		return s.invalidator.getDependentKeys(key)
	}
	return nil
}

func (s *DependencyInvalidationStrategy) GetPriority() int {
	return s.priority
}

func (s *DependencyInvalidationStrategy) GetName() string {
	return "dependency"
}

// MemoryInvalidationStrategy implements memory-based invalidation
type MemoryInvalidationStrategy struct {
	threshold int64
	priority  int
}

func (s *MemoryInvalidationStrategy) ShouldInvalidate(key string, reason InvalidationReason) bool {
	return reason == InvalidationReasonMemory
}

func (s *MemoryInvalidationStrategy) GetInvalidationKeys(key string, reason InvalidationReason) []string {
	if reason == InvalidationReasonMemory {
		return []string{key}
	}
	return nil
}

func (s *MemoryInvalidationStrategy) GetPriority() int {
	return s.priority
}

func (s *MemoryInvalidationStrategy) GetName() string {
	return "memory"
}

// ErrorInvalidationStrategy implements error-based invalidation
type ErrorInvalidationStrategy struct {
	priority int
}

func (s *ErrorInvalidationStrategy) ShouldInvalidate(key string, reason InvalidationReason) bool {
	return reason == InvalidationReasonError
}

func (s *ErrorInvalidationStrategy) GetInvalidationKeys(key string, reason InvalidationReason) []string {
	if reason == InvalidationReasonError {
		return []string{key}
	}
	return nil
}

func (s *ErrorInvalidationStrategy) GetPriority() int {
	return s.priority
}

func (s *ErrorInvalidationStrategy) GetName() string {
	return "error"
}

// Invalidate invalidates cache entries based on the given key and reason
func (i *CacheInvalidator) Invalidate(key string, reason InvalidationReason) error {
	if atomic.LoadInt32(&i.closed) == 1 {
		return fmt.Errorf("cache invalidator is closed")
	}

	// Get invalidation keys from strategies
	keysToInvalidate := i.getInvalidationKeys(key, reason)

	// Create invalidation tasks
	tasks := make([]*InvalidationTask, len(keysToInvalidate))
	for j, invKey := range keysToInvalidate {
		tasks[j] = &InvalidationTask{
			Key:       invKey,
			Reason:    reason,
			Priority:  i.config.InvalidationPriority,
			Source:    "explicit",
			CreatedAt: time.Now(),
			Metadata:  make(map[string]interface{}),
		}
	}

	// Process tasks
	return i.processInvalidationTasks(tasks)
}

// getInvalidationKeys gets keys to invalidate based on strategies
func (i *CacheInvalidator) getInvalidationKeys(key string, reason InvalidationReason) []string {
	var allKeys []string

	// Always include the key itself for explicit invalidation
	if reason == InvalidationReasonExplicit || reason == InvalidationReasonUpdate || reason == InvalidationReasonDelete {
		allKeys = append(allKeys, key)
	}

	for _, strategy := range i.strategies {
		if strategy.ShouldInvalidate(key, reason) {
			keys := strategy.GetInvalidationKeys(key, reason)
			allKeys = append(allKeys, keys...)
		}
	}

	// Remove duplicates
	keyMap := make(map[string]bool)
	for _, k := range allKeys {
		keyMap[k] = true
	}

	var uniqueKeys []string
	for k := range keyMap {
		uniqueKeys = append(uniqueKeys, k)
	}

	return uniqueKeys
}

// processInvalidationTasks processes invalidation tasks
func (i *CacheInvalidator) processInvalidationTasks(tasks []*InvalidationTask) error {
	start := time.Now()
	atomic.AddInt64(&i.stats.TotalInvalidations, int64(len(tasks)))

	// Process tasks in batches
	batchSize := i.config.BatchSize
	for j := 0; j < len(tasks); j += batchSize {
		end := j + batchSize
		if end > len(tasks) {
			end = len(tasks)
		}

		batch := tasks[j:end]
		if err := i.processBatch(batch); err != nil {
			atomic.AddInt64(&i.stats.FailedInvalidations, int64(len(batch)))
			continue
		}

		atomic.AddInt64(&i.stats.SuccessfulInvalidations, int64(len(batch)))
	}

	// Update statistics
	duration := time.Since(start)
	if len(tasks) > 0 {
		i.stats.AverageInvalidationTime = duration / time.Duration(len(tasks))
		i.stats.InvalidationThroughput = float64(len(tasks)) / duration.Seconds()
	}
	i.stats.LastInvalidationTime = time.Now()

	return nil
}

// processBatch processes a batch of invalidation tasks
func (i *CacheInvalidator) processBatch(tasks []*InvalidationTask) error {
	// Distribute tasks to workers
	for _, task := range tasks {
		select {
		case i.workers[0].workChan <- task:
		default:
			// If worker is busy, try next worker
			for k := 1; k < len(i.workers); k++ {
				select {
				case i.workers[k].workChan <- task:
					break
				default:
					continue
				}
			}
		}
	}

	return nil
}

// AddDependency adds a dependency relationship
func (i *CacheInvalidator) AddDependency(key string, dependentKey string) {
	if !i.config.EnableDependencies {
		return
	}

	i.depsMutex.Lock()
	defer i.depsMutex.Unlock()

	// Add to dependencies
	if deps, exists := i.dependencies[key]; exists {
		i.dependencies[key] = append(deps, dependentKey)
	} else {
		i.dependencies[key] = []string{dependentKey}
	}

	// Add to dependents
	if deps, exists := i.dependents[dependentKey]; exists {
		i.dependents[dependentKey] = append(deps, key)
	} else {
		i.dependents[dependentKey] = []string{key}
	}
}

// RemoveDependency removes a dependency relationship
func (i *CacheInvalidator) RemoveDependency(key string, dependentKey string) {
	if !i.config.EnableDependencies {
		return
	}

	i.depsMutex.Lock()
	defer i.depsMutex.Unlock()

	// Remove from dependencies
	if deps, exists := i.dependencies[key]; exists {
		for j, dep := range deps {
			if dep == dependentKey {
				i.dependencies[key] = append(deps[:j], deps[j+1:]...)
				break
			}
		}
	}

	// Remove from dependents
	if deps, exists := i.dependents[dependentKey]; exists {
		for j, dep := range deps {
			if dep == key {
				i.dependents[dependentKey] = append(deps[:j], deps[j+1:]...)
				break
			}
		}
	}
}

// getDependentKeys gets all keys that depend on the given key
func (i *CacheInvalidator) getDependentKeys(key string) []string {
	i.depsMutex.RLock()
	defer i.depsMutex.RUnlock()

	var allDeps []string
	visited := make(map[string]bool)

	i.collectDependents(key, &allDeps, visited, 0)

	return allDeps
}

// collectDependents recursively collects dependent keys
func (i *CacheInvalidator) collectDependents(key string, deps *[]string, visited map[string]bool, depth int) {
	if depth >= i.config.MaxDependencyDepth {
		return
	}

	if visited[key] {
		return
	}

	visited[key] = true

	if dependents, exists := i.dependents[key]; exists {
		for _, dep := range dependents {
			*deps = append(*deps, dep)
			i.collectDependents(dep, deps, visited, depth+1)
		}
	}
}

// SetTTL sets TTL for a key
func (i *CacheInvalidator) SetTTL(key string, ttl time.Duration) {
	if !i.config.EnableTTL {
		return
	}

	i.ttlManager.SetTTL(key, ttl)
}

// GetTTL gets TTL for a key
func (i *CacheInvalidator) GetTTL(key string) (time.Duration, bool) {
	if !i.config.EnableTTL {
		return 0, false
	}

	return i.ttlManager.GetTTL(key)
}

// SetTTL sets TTL for a key
func (tm *TTLManager) SetTTL(key string, ttl time.Duration) {
	tm.ttlMutex.Lock()
	defer tm.ttlMutex.Unlock()

	now := time.Now()
	expiresAt := now.Add(ttl)

	// Clamp TTL to max
	if ttl > tm.maxTTL {
		ttl = tm.maxTTL
		expiresAt = now.Add(ttl)
	}

	tm.ttlEntries[key] = &TTLEntry{
		Key:         key,
		ExpiresAt:   expiresAt,
		CreatedAt:   now,
		UpdatedAt:   now,
		AccessCount: 0,
		LastAccess:  now,
	}

	atomic.AddInt64(&tm.stats.TotalEntries, 1)
}

// GetTTL gets TTL for a key
func (tm *TTLManager) GetTTL(key string) (time.Duration, bool) {
	tm.ttlMutex.RLock()
	defer tm.ttlMutex.RUnlock()

	entry, exists := tm.ttlEntries[key]
	if !exists {
		return 0, false
	}

	remaining := time.Until(entry.ExpiresAt)
	return remaining, true
}

// CheckExpired checks for expired entries
func (tm *TTLManager) CheckExpired() []string {
	tm.ttlMutex.Lock()
	defer tm.ttlMutex.Unlock()

	var expiredKeys []string
	now := time.Now()

	for key, entry := range tm.ttlEntries {
		if now.After(entry.ExpiresAt) {
			expiredKeys = append(expiredKeys, key)
			delete(tm.ttlEntries, key)
			atomic.AddInt64(&tm.stats.ExpiredEntries, 1)
		}
	}

	atomic.AddInt64(&tm.stats.ActiveEntries, int64(len(tm.ttlEntries)))
	tm.stats.LastCleanup = now

	return expiredKeys
}

// initializeWorkers initializes the worker pool
func (i *CacheInvalidator) initializeWorkers() {
	i.workers = make([]*InvalidationWorker, i.config.MaxInvalidationWorkers)

	for j := 0; j < i.config.MaxInvalidationWorkers; j++ {
		ctx, cancel := context.WithCancel(i.ctx)
		worker := &InvalidationWorker{
			id:          j,
			invalidator: i,
			workChan:    make(chan *InvalidationTask, 100),
			ctx:         ctx,
			cancel:      cancel,
		}

		i.workers[j] = worker
		go worker.start()
	}
}

// start starts the worker
func (w *InvalidationWorker) start() {
	for {
		select {
		case task := <-w.workChan:
			w.processTask(task)
		case <-w.ctx.Done():
			return
		}
	}
}

// processTask processes an invalidation task
func (w *InvalidationWorker) processTask(task *InvalidationTask) {
	if task == nil {
		return
	}

	// Invalidate from all cache layers
	w.invalidator.l1Cache.Delete(task.Key)
	w.invalidator.l2Cache.Delete(task.Key)
	w.invalidator.l3Cache.Delete(task.Key)

	// Remove TTL entry if exists
	if w.invalidator.config.EnableTTL {
		w.invalidator.ttlManager.RemoveTTL(task.Key)
	}
}

// RemoveTTL removes TTL for a key
func (tm *TTLManager) RemoveTTL(key string) {
	tm.ttlMutex.Lock()
	defer tm.ttlMutex.Unlock()

	if _, exists := tm.ttlEntries[key]; exists {
		delete(tm.ttlEntries, key)
		atomic.AddInt64(&tm.stats.ExpiredEntries, 1)
	}
}

// startBackgroundInvalidation starts the background invalidation process
func (i *CacheInvalidator) startBackgroundInvalidation() {
	ticker := time.NewTicker(i.config.InvalidationInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if atomic.LoadInt32(&i.closed) == 1 {
				return
			}

			// Check for expired TTL entries
			if i.config.EnableTTL {
				expiredKeys := i.ttlManager.CheckExpired()
				if len(expiredKeys) > 0 {
					for _, key := range expiredKeys {
						i.Invalidate(key, InvalidationReasonTTL)
					}
				}
			}

		case <-i.ctx.Done():
			return
		}
	}
}

// startTTLManagement starts the TTL management process
func (i *CacheInvalidator) startTTLManagement() {
	ticker := time.NewTicker(i.config.TTLCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if atomic.LoadInt32(&i.closed) == 1 {
				return
			}

			// Check for expired entries
			expiredKeys := i.ttlManager.CheckExpired()
			if len(expiredKeys) > 0 {
				atomic.AddInt64(&i.stats.TTLExpirations, int64(len(expiredKeys)))
				atomic.AddInt64(&i.stats.TTLCleanups, 1)
			}

		case <-i.ctx.Done():
			return
		}
	}
}

// GetStats returns cache invalidator statistics
func (i *CacheInvalidator) GetStats() *CacheInvalidatorStats {
	i.depsMutex.RLock()
	defer i.depsMutex.RUnlock()

	stats := &CacheInvalidatorStats{
		TotalInvalidations:      atomic.LoadInt64(&i.stats.TotalInvalidations),
		SuccessfulInvalidations: atomic.LoadInt64(&i.stats.SuccessfulInvalidations),
		FailedInvalidations:     atomic.LoadInt64(&i.stats.FailedInvalidations),
		ActiveWorkers:           int64(len(i.workers)),
		TTLExpirations:          atomic.LoadInt64(&i.stats.TTLExpirations),
		TTLCleanups:             atomic.LoadInt64(&i.stats.TTLCleanups),
		AverageTTL:              i.stats.AverageTTL,
		DependencyInvalidations: atomic.LoadInt64(&i.stats.DependencyInvalidations),
		DependencyChains:        atomic.LoadInt64(&i.stats.DependencyChains),
		MaxDependencyDepth:      i.stats.MaxDependencyDepth,
		AverageInvalidationTime: i.stats.AverageInvalidationTime,
		InvalidationThroughput:  i.stats.InvalidationThroughput,
		MemoryFreed:             i.stats.MemoryFreed,
		MemoryUsage:             i.stats.MemoryUsage,
		CPUUsage:                i.stats.CPUUsage,
		LastInvalidationTime:    i.stats.LastInvalidationTime,
	}

	return stats
}

// Close closes the cache invalidator
func (i *CacheInvalidator) Close() error {
	atomic.StoreInt32(&i.closed, 1)

	// Cancel context
	i.cancel()

	// Close workers
	for _, worker := range i.workers {
		worker.cancel()
		close(worker.workChan)
	}

	return nil
}

// GetDependencies returns dependency information
func (i *CacheInvalidator) GetDependencies() map[string][]string {
	i.depsMutex.RLock()
	defer i.depsMutex.RUnlock()

	deps := make(map[string][]string)
	for key, dependents := range i.dependencies {
		deps[key] = make([]string, len(dependents))
		copy(deps[key], dependents)
	}

	return deps
}

// ClearDependencies clears all dependencies
func (i *CacheInvalidator) ClearDependencies() {
	i.depsMutex.Lock()
	defer i.depsMutex.Unlock()

	i.dependencies = make(map[string][]string)
	i.dependents = make(map[string][]string)
}

// SetConfig updates the configuration
func (i *CacheInvalidator) SetConfig(config *CacheInvalidatorConfig) error {
	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	i.config = config
	return nil
}

// GetConfig returns the current configuration
func (i *CacheInvalidator) GetConfig() *CacheInvalidatorConfig {
	return i.config
}
