package cache

import (
	"context"
	"fmt"
	"math"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// CacheWarmer provides intelligent cache warming and preloading
type CacheWarmer struct {
	// Configuration
	config *CacheWarmerConfig

	// Cache layers
	l1Cache *LockFreeL1Cache
	l2Cache *MemoryMappedL2Cache
	l3Cache *PersistentL3Cache

	// Access pattern tracking
	accessPatterns map[string]*AccessPattern
	patternMutex   sync.RWMutex

	// ML-based prediction
	predictor *AccessPredictor

	// Background workers
	workers    []*WarmerWorker
	workerPool sync.Pool

	// Statistics
	stats *CacheWarmerStats

	// State
	closed int32
	ctx    context.Context
	cancel context.CancelFunc
}

// CacheWarmerConfig holds configuration for cache warming
type CacheWarmerConfig struct {
	// Warming settings
	EnableWarming     bool
	WarmingInterval   time.Duration
	MaxWarmingWorkers int

	// Prediction settings
	EnablePrediction    bool
	PredictionWindow    time.Duration
	MinAccessCount      int64
	ConfidenceThreshold float64

	// Preloading settings
	PreloadBatchSize int
	PreloadTimeout   time.Duration
	MaxPreloadSize   int64

	// Performance settings
	WarmingPriority  int
	EnableHotPath    bool
	HotPathThreshold float64
}

// AccessPattern tracks access patterns for a key
type AccessPattern struct {
	Key             string
	AccessCount     int64
	LastAccess      time.Time
	FirstAccess     time.Time
	AccessTimes     []time.Time
	AccessIntervals []time.Duration
	Frequency       float64
	Recency         float64
	Popularity      float64
	Trend           float64
}

// AccessPredictor provides ML-based access prediction
type AccessPredictor struct {
	// Pattern analysis
	patterns map[string]*AccessPattern

	// Prediction models
	frequencyModel *FrequencyModel
	recencyModel   *RecencyModel
	trendModel     *TrendModel

	// Configuration
	windowSize          time.Duration
	minAccessCount      int64
	confidenceThreshold float64

	// Statistics
	stats *PredictorStats
}

// FrequencyModel predicts access based on frequency
type FrequencyModel struct {
	weights map[string]float64
	mu      sync.RWMutex
}

// RecencyModel predicts access based on recency
type RecencyModel struct {
	decayFactor float64
	mu          sync.RWMutex
}

// TrendModel predicts access based on trend analysis
type TrendModel struct {
	trends map[string]float64
	mu     sync.RWMutex
}

// PredictorStats holds statistics for the predictor
type PredictorStats struct {
	TotalPredictions   int64
	CorrectPredictions int64
	Accuracy           float64
	LastPrediction     time.Time
}

// WarmerWorker handles cache warming tasks
type WarmerWorker struct {
	id       int
	warmer   *CacheWarmer
	workChan chan *WarmingTask
	ctx      context.Context
	cancel   context.CancelFunc
}

// WarmingTask represents a cache warming task
type WarmingTask struct {
	Key       string
	Priority  int
	Source    string
	CreatedAt time.Time
}

// CacheWarmerStats holds statistics for cache warming
type CacheWarmerStats struct {
	// Warming stats
	TotalWarmingTasks int64
	CompletedTasks    int64
	FailedTasks       int64
	ActiveWorkers     int64

	// Prediction stats
	TotalPredictions   int64
	CorrectPredictions int64
	PredictionAccuracy float64

	// Performance stats
	AverageWarmingTime  time.Duration
	WarmingThroughput   float64
	CacheHitImprovement float64

	// Resource stats
	MemoryUsage     int64
	CPUUsage        float64
	LastWarmingTime time.Time
}

// NewCacheWarmer creates a new cache warmer
func NewCacheWarmer(config *CacheWarmerConfig, l1Cache *LockFreeL1Cache, l2Cache *MemoryMappedL2Cache, l3Cache *PersistentL3Cache) *CacheWarmer {
	if config == nil {
		config = &CacheWarmerConfig{
			EnableWarming:       true,
			WarmingInterval:     time.Minute * 5,
			MaxWarmingWorkers:   4,
			EnablePrediction:    true,
			PredictionWindow:    time.Hour,
			MinAccessCount:      5,
			ConfidenceThreshold: 0.7,
			PreloadBatchSize:    100,
			PreloadTimeout:      time.Second * 30,
			MaxPreloadSize:      1024 * 1024, // 1MB
			WarmingPriority:     1,
			EnableHotPath:       true,
			HotPathThreshold:    0.8,
		}
	}

	ctx, cancel := context.WithCancel(context.Background())

	warmer := &CacheWarmer{
		config:         config,
		l1Cache:        l1Cache,
		l2Cache:        l2Cache,
		l3Cache:        l3Cache,
		accessPatterns: make(map[string]*AccessPattern),
		predictor:      NewAccessPredictor(config),
		stats:          &CacheWarmerStats{},
		ctx:            ctx,
		cancel:         cancel,
	}

	// Initialize workers
	warmer.initializeWorkers()

	// Start background warming if enabled
	if config.EnableWarming {
		go warmer.startBackgroundWarming()
	}

	return warmer
}

// NewAccessPredictor creates a new access predictor
func NewAccessPredictor(config *CacheWarmerConfig) *AccessPredictor {
	return &AccessPredictor{
		patterns:            make(map[string]*AccessPattern),
		frequencyModel:      &FrequencyModel{weights: make(map[string]float64)},
		recencyModel:        &RecencyModel{decayFactor: 0.9},
		trendModel:          &TrendModel{trends: make(map[string]float64)},
		windowSize:          config.PredictionWindow,
		minAccessCount:      config.MinAccessCount,
		confidenceThreshold: config.ConfidenceThreshold,
		stats:               &PredictorStats{},
	}
}

// RecordAccess records an access pattern for a key
func (w *CacheWarmer) RecordAccess(key string) {
	if atomic.LoadInt32(&w.closed) == 1 {
		return
	}

	now := time.Now()

	w.patternMutex.Lock()
	defer w.patternMutex.Unlock()

	pattern, exists := w.accessPatterns[key]
	if !exists {
		pattern = &AccessPattern{
			Key:         key,
			AccessCount: 0,
			FirstAccess: now,
			AccessTimes: make([]time.Time, 0, 100),
		}
		w.accessPatterns[key] = pattern
	}

	// Update access pattern
	pattern.AccessCount++
	pattern.LastAccess = now
	pattern.AccessTimes = append(pattern.AccessTimes, now)

	// Keep only recent access times (sliding window)
	if len(pattern.AccessTimes) > 100 {
		pattern.AccessTimes = pattern.AccessTimes[1:]
	}

	// Update intervals
	if len(pattern.AccessTimes) > 1 {
		interval := now.Sub(pattern.AccessTimes[len(pattern.AccessTimes)-2])
		pattern.AccessIntervals = append(pattern.AccessIntervals, interval)
		if len(pattern.AccessIntervals) > 50 {
			pattern.AccessIntervals = pattern.AccessIntervals[1:]
		}
	}

	// Update metrics
	w.updatePatternMetrics(pattern)
}

// updatePatternMetrics updates the calculated metrics for a pattern
func (w *CacheWarmer) updatePatternMetrics(pattern *AccessPattern) {
	now := time.Now()

	// Calculate frequency (accesses per minute)
	if len(pattern.AccessTimes) > 1 {
		timeSpan := now.Sub(pattern.FirstAccess).Minutes()
		if timeSpan > 0 {
			pattern.Frequency = float64(pattern.AccessCount) / timeSpan
		}
	}

	// Calculate recency (inverse of time since last access)
	timeSinceLastAccess := now.Sub(pattern.LastAccess).Seconds()
	pattern.Recency = 1.0 / (1.0 + timeSinceLastAccess/60.0) // Decay over minutes

	// Calculate popularity (normalized access count)
	maxAccessCount := w.getMaxAccessCount()
	if maxAccessCount > 0 {
		pattern.Popularity = float64(pattern.AccessCount) / float64(maxAccessCount)
	}

	// Calculate trend (change in access rate over time)
	if len(pattern.AccessIntervals) > 1 {
		recentIntervals := pattern.AccessIntervals[len(pattern.AccessIntervals)/2:]
		olderIntervals := pattern.AccessIntervals[:len(pattern.AccessIntervals)/2]

		recentAvg := w.calculateAverageInterval(recentIntervals)
		olderAvg := w.calculateAverageInterval(olderIntervals)

		if olderAvg > 0 {
			pattern.Trend = (olderAvg - recentAvg) / olderAvg
		}
	}
}

// getMaxAccessCount returns the maximum access count across all patterns
func (w *CacheWarmer) getMaxAccessCount() int64 {
	maxCount := int64(0)
	for _, pattern := range w.accessPatterns {
		if pattern.AccessCount > maxCount {
			maxCount = pattern.AccessCount
		}
	}
	return maxCount
}

// calculateAverageInterval calculates the average interval between accesses
func (w *CacheWarmer) calculateAverageInterval(intervals []time.Duration) float64 {
	if len(intervals) == 0 {
		return 0
	}

	sum := float64(0)
	for _, interval := range intervals {
		sum += interval.Seconds()
	}
	return sum / float64(len(intervals))
}

// PredictNextAccess predicts the next access time for a key
func (w *CacheWarmer) PredictNextAccess(key string) (time.Time, float64) {
	if atomic.LoadInt32(&w.closed) == 1 {
		return time.Time{}, 0
	}

	w.patternMutex.RLock()
	pattern, exists := w.accessPatterns[key]
	w.patternMutex.RUnlock()

	if !exists {
		return time.Time{}, 0
	}

	return w.predictor.PredictNextAccess(key, pattern)
}

// PredictNextAccess predicts the next access time for a key
func (p *AccessPredictor) PredictNextAccess(key string, pattern *AccessPattern) (time.Time, float64) {
	if pattern == nil {
		return time.Time{}, 0
	}

	if pattern.AccessCount < p.minAccessCount {
		return time.Time{}, 0
	}

	// Calculate prediction confidence
	confidence := p.calculateConfidence(pattern)
	if confidence < p.confidenceThreshold {
		return time.Time{}, confidence
	}

	// Predict next access time based on frequency and intervals
	nextAccess := p.predictNextAccessTime(pattern)

	atomic.AddInt64(&p.stats.TotalPredictions, 1)
	p.stats.LastPrediction = time.Now()

	return nextAccess, confidence
}

// calculateConfidence calculates the confidence for a prediction
func (p *AccessPredictor) calculateConfidence(pattern *AccessPattern) float64 {
	// Base confidence on access count and recency
	accessConfidence := math.Min(float64(pattern.AccessCount)/10.0, 1.0)
	recencyConfidence := pattern.Recency

	// Combine confidences
	confidence := (accessConfidence + recencyConfidence) / 2.0

	// Boost confidence for high-frequency patterns
	if pattern.Frequency > 1.0 {
		confidence = math.Min(confidence*1.2, 1.0)
	}

	return confidence
}

// predictNextAccessTime predicts the next access time
func (p *AccessPredictor) predictNextAccessTime(pattern *AccessPattern) time.Time {
	now := time.Now()

	// If we have intervals, use them for prediction
	if len(pattern.AccessIntervals) > 0 {
		avgInterval := p.calculateAverageInterval(pattern.AccessIntervals)
		return now.Add(time.Duration(avgInterval) * time.Second)
	}

	// Fallback to frequency-based prediction
	if pattern.Frequency > 0 {
		intervalMinutes := 1.0 / pattern.Frequency
		return now.Add(time.Duration(intervalMinutes) * time.Minute)
	}

	// Default prediction (1 hour)
	return now.Add(time.Hour)
}

// calculateAverageInterval calculates the average interval
func (p *AccessPredictor) calculateAverageInterval(intervals []time.Duration) float64 {
	if len(intervals) == 0 {
		return 0
	}

	sum := float64(0)
	for _, interval := range intervals {
		sum += interval.Seconds()
	}
	return sum / float64(len(intervals))
}

// GetWarmingCandidates returns keys that should be warmed
func (w *CacheWarmer) GetWarmingCandidates(limit int) []string {
	if atomic.LoadInt32(&w.closed) == 1 {
		return nil
	}

	w.patternMutex.RLock()
	defer w.patternMutex.RUnlock()

	// Score all patterns
	type ScoredPattern struct {
		Key   string
		Score float64
	}

	var scoredPatterns []ScoredPattern
	for key, pattern := range w.accessPatterns {
		score := w.calculateWarmingScore(pattern)
		scoredPatterns = append(scoredPatterns, ScoredPattern{Key: key, Score: score})
	}

	// Sort by score (descending)
	sort.Slice(scoredPatterns, func(i, j int) bool {
		return scoredPatterns[i].Score > scoredPatterns[j].Score
	})

	// Return top candidates
	candidates := make([]string, 0, limit)
	for i, scored := range scoredPatterns {
		if i >= limit {
			break
		}
		candidates = append(candidates, scored.Key)
	}

	return candidates
}

// calculateWarmingScore calculates the warming score for a pattern
func (w *CacheWarmer) calculateWarmingScore(pattern *AccessPattern) float64 {
	// Weighted combination of metrics
	frequencyWeight := 0.4
	recencyWeight := 0.3
	popularityWeight := 0.2
	trendWeight := 0.1

	score := pattern.Frequency*frequencyWeight +
		pattern.Recency*recencyWeight +
		pattern.Popularity*popularityWeight +
		pattern.Trend*trendWeight

	// Boost score for hot path items
	if w.config.EnableHotPath && pattern.Popularity > w.config.HotPathThreshold {
		score *= 1.5
	}

	return score
}

// WarmCache warms the cache with predicted data
func (w *CacheWarmer) WarmCache(candidates []string) error {
	if atomic.LoadInt32(&w.closed) == 1 {
		return fmt.Errorf("cache warmer is closed")
	}

	start := time.Now()
	atomic.AddInt64(&w.stats.TotalWarmingTasks, int64(len(candidates)))

	// Process candidates in batches
	batchSize := w.config.PreloadBatchSize
	for i := 0; i < len(candidates); i += batchSize {
		end := i + batchSize
		if end > len(candidates) {
			end = len(candidates)
		}

		batch := candidates[i:end]
		if err := w.warmBatch(batch); err != nil {
			atomic.AddInt64(&w.stats.FailedTasks, int64(len(batch)))
			continue
		}

		atomic.AddInt64(&w.stats.CompletedTasks, int64(len(batch)))
	}

	// Update statistics
	duration := time.Since(start)
	w.stats.AverageWarmingTime = duration / time.Duration(len(candidates))
	w.stats.WarmingThroughput = float64(len(candidates)) / duration.Seconds()
	w.stats.LastWarmingTime = time.Now()

	return nil
}

// warmBatch warms a batch of keys
func (w *CacheWarmer) warmBatch(keys []string) error {
	// Create warming tasks
	tasks := make([]*WarmingTask, len(keys))
	for i, key := range keys {
		tasks[i] = &WarmingTask{
			Key:       key,
			Priority:  w.config.WarmingPriority,
			Source:    "prediction",
			CreatedAt: time.Now(),
		}
	}

	// Distribute tasks to workers
	for _, task := range tasks {
		select {
		case w.workers[0].workChan <- task:
		default:
			// If worker is busy, try next worker
			for i := 1; i < len(w.workers); i++ {
				select {
				case w.workers[i].workChan <- task:
					break
				default:
					continue
				}
			}
		}
	}

	return nil
}

// initializeWorkers initializes the worker pool
func (w *CacheWarmer) initializeWorkers() {
	w.workers = make([]*WarmerWorker, w.config.MaxWarmingWorkers)

	for i := 0; i < w.config.MaxWarmingWorkers; i++ {
		ctx, cancel := context.WithCancel(w.ctx)
		worker := &WarmerWorker{
			id:       i,
			warmer:   w,
			workChan: make(chan *WarmingTask, 100),
			ctx:      ctx,
			cancel:   cancel,
		}

		w.workers[i] = worker
		go worker.start()
	}
}

// start starts the worker
func (w *WarmerWorker) start() {
	for {
		select {
		case task := <-w.workChan:
			w.processTask(task)
		case <-w.ctx.Done():
			return
		}
	}
}

// processTask processes a warming task
func (w *WarmerWorker) processTask(task *WarmingTask) {
	if task == nil {
		return
	}

	// Try to get the value from L3 cache
	value, found := w.warmer.l3Cache.Get(task.Key)
	if !found {
		return
	}

	// Promote to L2 and L1 caches
	w.warmer.l2Cache.Set(task.Key, value, time.Hour)
	w.warmer.l1Cache.Set(task.Key, value, time.Hour)
}

// startBackgroundWarming starts the background warming process
func (w *CacheWarmer) startBackgroundWarming() {
	ticker := time.NewTicker(w.config.WarmingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if atomic.LoadInt32(&w.closed) == 1 {
				return
			}

			// Get warming candidates
			candidates := w.GetWarmingCandidates(100)
			if len(candidates) > 0 {
				w.WarmCache(candidates)
			}

		case <-w.ctx.Done():
			return
		}
	}
}

// GetStats returns cache warmer statistics
func (w *CacheWarmer) GetStats() *CacheWarmerStats {
	w.patternMutex.RLock()
	defer w.patternMutex.RUnlock()

	stats := &CacheWarmerStats{
		TotalWarmingTasks:   atomic.LoadInt64(&w.stats.TotalWarmingTasks),
		CompletedTasks:      atomic.LoadInt64(&w.stats.CompletedTasks),
		FailedTasks:         atomic.LoadInt64(&w.stats.FailedTasks),
		ActiveWorkers:       int64(len(w.workers)),
		TotalPredictions:    atomic.LoadInt64(&w.predictor.stats.TotalPredictions),
		CorrectPredictions:  atomic.LoadInt64(&w.predictor.stats.CorrectPredictions),
		PredictionAccuracy:  w.predictor.stats.Accuracy,
		AverageWarmingTime:  w.stats.AverageWarmingTime,
		WarmingThroughput:   w.stats.WarmingThroughput,
		CacheHitImprovement: w.stats.CacheHitImprovement,
		MemoryUsage:         w.stats.MemoryUsage,
		CPUUsage:            w.stats.CPUUsage,
		LastWarmingTime:     w.stats.LastWarmingTime,
	}

	return stats
}

// Close closes the cache warmer
func (w *CacheWarmer) Close() error {
	atomic.StoreInt32(&w.closed, 1)

	// Cancel context
	w.cancel()

	// Close workers
	for _, worker := range w.workers {
		worker.cancel()
		close(worker.workChan)
	}

	return nil
}

// GetAccessPatterns returns access patterns for analysis
func (w *CacheWarmer) GetAccessPatterns() map[string]*AccessPattern {
	w.patternMutex.RLock()
	defer w.patternMutex.RUnlock()

	patterns := make(map[string]*AccessPattern)
	for key, pattern := range w.accessPatterns {
		patterns[key] = pattern
	}

	return patterns
}

// ClearPatterns clears all access patterns
func (w *CacheWarmer) ClearPatterns() {
	w.patternMutex.Lock()
	defer w.patternMutex.Unlock()

	w.accessPatterns = make(map[string]*AccessPattern)
}

// SetConfig updates the configuration
func (w *CacheWarmer) SetConfig(config *CacheWarmerConfig) error {
	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	w.config = config
	return nil
}

// GetConfig returns the current configuration
func (w *CacheWarmer) GetConfig() *CacheWarmerConfig {
	return w.config
}
