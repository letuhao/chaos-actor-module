package monitoring

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// PerformanceMonitor implements the PerformanceMonitor interface
type PerformanceMonitor struct {
	mu            sync.RWMutex
	metrics       map[string]*Metric
	alerts        map[string]*Alert
	thresholds    map[string]float64
	calculations  map[string]*CalculationTimer
	alertHandlers map[string]AlertHandler
	enabled       bool
	alertEnabled  bool
	version       int64
	createdAt     int64
	updatedAt     int64
	// Context-driven fields
	contextMetrics map[string]map[string]*Metric           // contextID -> metricName -> Metric
	contextTimers  map[string]map[string]*CalculationTimer // contextID -> operation -> Timer
	// Prometheus-style counters
	counters   map[string]*Counter
	histograms map[string]*Histogram
	gauges     map[string]*Gauge
}

// Metric represents a performance metric
type Metric struct {
	Name        string    `json:"name"`
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Timestamp   int64     `json:"timestamp"`
	History     []float64 `json:"history"`
	MaxHistory  int       `json:"max_history"`
	// New fields for interface compatibility
	Tags  map[string]string `json:"tags"`
	Count int64             `json:"count"`
	Min   float64           `json:"min"`
	Max   float64           `json:"max"`
	Avg   float64           `json:"avg"`
	Sum   float64           `json:"sum"`
}

// Alert represents a performance alert
type Alert struct {
	ID          string  `json:"id"`
	MetricName  string  `json:"metric_name"`
	Threshold   float64 `json:"threshold"`
	Operator    string  `json:"operator"` // ">", "<", ">=", "<=", "==", "!="
	Message     string  `json:"message"`
	Severity    string  `json:"severity"` // "low", "medium", "high", "critical"
	Enabled     bool    `json:"enabled"`
	Triggered   bool    `json:"triggered"`
	LastTrigger int64   `json:"last_trigger"`
	CreatedAt   int64   `json:"created_at"`
	// New fields for interface compatibility
	Metric    string            `json:"metric"`
	Value     float64           `json:"value"`
	Tags      map[string]string `json:"tags"`
	Timestamp int64             `json:"timestamp"`
	IsActive  bool              `json:"is_active"`
}

// PerformanceStats represents overall performance statistics
type PerformanceStats struct {
	TotalMetrics   int     `json:"total_metrics"`
	ActiveAlerts   int     `json:"active_alerts"`
	AverageLatency float64 `json:"average_latency"`
	PeakMemory     float64 `json:"peak_memory"`
	CPUUsage       float64 `json:"cpu_usage"`
	Throughput     float64 `json:"throughput"`
	ErrorRate      float64 `json:"error_rate"`
	Uptime         int64   `json:"uptime"`
	LastUpdated    int64   `json:"last_updated"`
}

// Counter represents a Prometheus-style counter
type Counter struct {
	Name        string            `json:"name"`
	Value       float64           `json:"value"`
	Tags        map[string]string `json:"tags"`
	Description string            `json:"description"`
	Timestamp   int64             `json:"timestamp"`
}

// Histogram represents a Prometheus-style histogram
type Histogram struct {
	Name        string            `json:"name"`
	Buckets     map[string]int64  `json:"buckets"` // bucket -> count
	Sum         float64           `json:"sum"`
	Count       int64             `json:"count"`
	Tags        map[string]string `json:"tags"`
	Description string            `json:"description"`
	Timestamp   int64             `json:"timestamp"`
}

// Gauge represents a Prometheus-style gauge
type Gauge struct {
	Name        string            `json:"name"`
	Value       float64           `json:"value"`
	Tags        map[string]string `json:"tags"`
	Description string            `json:"description"`
	Timestamp   int64             `json:"timestamp"`
}

// CalculationTimer represents a timer for calculations
type CalculationTimer struct {
	Operation string            `json:"operation"`
	StartTime int64             `json:"start_time"`
	EndTime   int64             `json:"end_time"`
	Duration  int64             `json:"duration"`
	Tags      map[string]string `json:"tags"`
}

// PerformanceReport represents a performance report
type PerformanceReport struct {
	Timestamp    int64              `json:"timestamp"`
	Duration     int64              `json:"duration"`
	Metrics      map[string]Metric  `json:"metrics"`
	Calculations []CalculationTimer `json:"calculations"`
	Alerts       []Alert            `json:"alerts"`
	Summary      PerformanceSummary `json:"summary"`
}

// PerformanceSummary represents a performance summary
type PerformanceSummary struct {
	TotalCalculations  int64   `json:"total_calculations"`
	AvgCalculationTime float64 `json:"avg_calculation_time"`
	MaxCalculationTime int64   `json:"max_calculation_time"`
	MinCalculationTime int64   `json:"min_calculation_time"`
	TotalMetrics       int64   `json:"total_metrics"`
	ActiveAlerts       int64   `json:"active_alerts"`
	MemoryUsage        int64   `json:"memory_usage"`
	CPUUsage           float64 `json:"cpu_usage"`
}

// AlertHandler defines the interface for alert handlers
type AlertHandler interface {
	HandleAlert(alert Alert)
	GetHandlerID() string
	GetSeverity() string
}

// NewPerformanceMonitor creates a new PerformanceMonitor instance
func NewPerformanceMonitor() *PerformanceMonitor {
	now := time.Now().Unix()
	return &PerformanceMonitor{
		metrics:        make(map[string]*Metric),
		alerts:         make(map[string]*Alert),
		thresholds:     make(map[string]float64),
		calculations:   make(map[string]*CalculationTimer),
		alertHandlers:  make(map[string]AlertHandler),
		contextMetrics: make(map[string]map[string]*Metric),
		contextTimers:  make(map[string]map[string]*CalculationTimer),
		counters:       make(map[string]*Counter),
		histograms:     make(map[string]*Histogram),
		gauges:         make(map[string]*Gauge),
		enabled:        true,
		alertEnabled:   true,
		version:        1,
		createdAt:      now,
		updatedAt:      now,
	}
}

// SetMetric sets a performance metric
func (pm *PerformanceMonitor) SetMetric(name string, value float64, unit string, category string, description string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if name == "" {
		return fmt.Errorf("metric name cannot be empty")
	}

	now := time.Now().Unix()

	// Create or update metric
	if metric, exists := pm.metrics[name]; exists {
		// Update existing metric
		metric.Value = value
		metric.Timestamp = now

		// Add to history
		metric.History = append(metric.History, value)
		if len(metric.History) > metric.MaxHistory {
			metric.History = metric.History[1:]
		}
	} else {
		// Create new metric
		pm.metrics[name] = &Metric{
			Name:        name,
			Value:       value,
			Unit:        unit,
			Category:    category,
			Description: description,
			Timestamp:   now,
			History:     []float64{value},
			MaxHistory:  100, // Default max history
			Tags:        make(map[string]string),
			Count:       1,
			Min:         value,
			Max:         value,
			Avg:         value,
			Sum:         value,
		}
	}

	pm.updatedAt = now
	pm.version++

	// Check alerts
	if pm.alertEnabled {
		pm.checkAlerts(name, value)
	}

	return nil
}

// GetMetricValue gets a metric value
func (pm *PerformanceMonitor) GetMetricValue(name string) (float64, error) {
	metric, err := pm.GetMetric(name)
	if err != nil {
		return 0, err
	}
	return metric.Value, nil
}

// GetMetricHistory gets metric history
func (pm *PerformanceMonitor) GetMetricHistory(name string) ([]float64, error) {
	metric, err := pm.GetMetric(name)
	if err != nil {
		return nil, err
	}
	return metric.History, nil
}

// GetAllMetrics returns all metrics
func (pm *PerformanceMonitor) GetAllMetrics() map[string]*Metric {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	result := make(map[string]*Metric)
	for k, v := range pm.metrics {
		result[k] = v
	}
	return result
}

// GetMetricsByCategory returns metrics by category
func (pm *PerformanceMonitor) GetMetricsByCategory(category string) map[string]*Metric {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	result := make(map[string]*Metric)
	for k, v := range pm.metrics {
		if v.Category == category {
			result[k] = v
		}
	}
	return result
}

// RemoveMetric removes a metric
func (pm *PerformanceMonitor) RemoveMetric(name string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if _, exists := pm.metrics[name]; !exists {
		return fmt.Errorf("metric '%s' not found", name)
	}

	delete(pm.metrics, name)
	pm.updatedAt = time.Now().Unix()
	pm.version++

	return nil
}

// ClearMetrics clears all metrics
func (pm *PerformanceMonitor) ClearMetrics() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if len(pm.metrics) > 0 {
		pm.metrics = make(map[string]*Metric)
		pm.updatedAt = time.Now().Unix()
		pm.version++
	}
}

// SetThreshold sets a threshold for a metric
func (pm *PerformanceMonitor) SetThreshold(metricName string, threshold float64) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.thresholds[metricName] = threshold
	pm.updatedAt = time.Now().Unix()
	pm.version++
}

// GetThreshold gets a threshold for a metric
func (pm *PerformanceMonitor) GetThreshold(metricName string) (float64, bool) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	threshold, exists := pm.thresholds[metricName]
	return threshold, exists
}

// CreateAlert creates a new alert
func (pm *PerformanceMonitor) CreateAlert(id, metricName string, threshold float64, operator, message, severity string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if id == "" {
		return fmt.Errorf("alert ID cannot be empty")
	}

	if _, exists := pm.alerts[id]; exists {
		return fmt.Errorf("alert with ID '%s' already exists", id)
	}

	now := time.Now().Unix()
	pm.alerts[id] = &Alert{
		ID:          id,
		MetricName:  metricName,
		Threshold:   threshold,
		Operator:    operator,
		Message:     message,
		Severity:    severity,
		Enabled:     true,
		Triggered:   false,
		LastTrigger: 0,
		CreatedAt:   now,
	}

	pm.updatedAt = now
	pm.version++

	return nil
}

// GetAlert gets an alert by ID
func (pm *PerformanceMonitor) GetAlert(id string) (*Alert, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	alert, exists := pm.alerts[id]
	if !exists {
		return nil, fmt.Errorf("alert '%s' not found", id)
	}

	return alert, nil
}

// GetAllAlerts returns all alerts
func (pm *PerformanceMonitor) GetAllAlerts() map[string]*Alert {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	result := make(map[string]*Alert)
	for k, v := range pm.alerts {
		result[k] = v
	}
	return result
}

// GetActiveAlerts returns active alerts
func (pm *PerformanceMonitor) GetActiveAlerts() map[string]*Alert {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	result := make(map[string]*Alert)
	for k, v := range pm.alerts {
		if v.Enabled && v.Triggered {
			result[k] = v
		}
	}
	return result
}

// EnableAlert enables an alert
func (pm *PerformanceMonitor) EnableAlert(id string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	alert, exists := pm.alerts[id]
	if !exists {
		return fmt.Errorf("alert '%s' not found", id)
	}

	alert.Enabled = true
	pm.updatedAt = time.Now().Unix()
	pm.version++

	return nil
}

// DisableAlert disables an alert
func (pm *PerformanceMonitor) DisableAlert(id string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	alert, exists := pm.alerts[id]
	if !exists {
		return fmt.Errorf("alert '%s' not found", id)
	}

	alert.Enabled = false
	pm.updatedAt = time.Now().Unix()
	pm.version++

	return nil
}

// RemoveAlert removes an alert
func (pm *PerformanceMonitor) RemoveAlert(id string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if _, exists := pm.alerts[id]; !exists {
		return fmt.Errorf("alert '%s' not found", id)
	}

	delete(pm.alerts, id)
	pm.updatedAt = time.Now().Unix()
	pm.version++

	return nil
}

// ClearAlerts clears all alerts
func (pm *PerformanceMonitor) ClearAlerts() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if len(pm.alerts) > 0 {
		pm.alerts = make(map[string]*Alert)
		pm.updatedAt = time.Now().Unix()
		pm.version++
	}
}

// checkAlerts checks if any alerts should be triggered
func (pm *PerformanceMonitor) checkAlerts(metricName string, value float64) {
	now := time.Now().Unix()

	for _, alert := range pm.alerts {
		if !alert.Enabled || alert.MetricName != metricName {
			continue
		}

		shouldTrigger := false
		switch alert.Operator {
		case ">":
			shouldTrigger = value > alert.Threshold
		case "<":
			shouldTrigger = value < alert.Threshold
		case ">=":
			shouldTrigger = value >= alert.Threshold
		case "<=":
			shouldTrigger = value <= alert.Threshold
		case "==":
			shouldTrigger = value == alert.Threshold
		case "!=":
			shouldTrigger = value != alert.Threshold
		}

		if shouldTrigger && !alert.Triggered {
			alert.Triggered = true
			alert.LastTrigger = now
		} else if !shouldTrigger && alert.Triggered {
			alert.Triggered = false
		}
	}
}

// GetPerformanceStats returns overall performance statistics
func (pm *PerformanceMonitor) GetPerformanceStats() *PerformanceStats {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	stats := &PerformanceStats{
		TotalMetrics: len(pm.metrics),
		LastUpdated:  pm.updatedAt,
	}

	// Count active alerts
	for _, alert := range pm.alerts {
		if alert.Enabled && alert.Triggered {
			stats.ActiveAlerts++
		}
	}

	// Calculate average latency
	if latencyMetric, exists := pm.metrics["latency"]; exists {
		if len(latencyMetric.History) > 0 {
			sum := 0.0
			for _, v := range latencyMetric.History {
				sum += v
			}
			stats.AverageLatency = sum / float64(len(latencyMetric.History))
		}
	}

	// Get peak memory
	if memoryMetric, exists := pm.metrics["memory_usage"]; exists {
		stats.PeakMemory = memoryMetric.Value
	}

	// Get CPU usage
	if cpuMetric, exists := pm.metrics["cpu_usage"]; exists {
		stats.CPUUsage = cpuMetric.Value
	}

	// Get throughput
	if throughputMetric, exists := pm.metrics["throughput"]; exists {
		stats.Throughput = throughputMetric.Value
	}

	// Get error rate
	if errorMetric, exists := pm.metrics["error_rate"]; exists {
		stats.ErrorRate = errorMetric.Value
	}

	// Calculate uptime
	stats.Uptime = pm.updatedAt - pm.createdAt

	return stats
}

// SetEnabled sets the monitor enabled state
func (pm *PerformanceMonitor) SetEnabled(enabled bool) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.enabled = enabled
	pm.updatedAt = time.Now().Unix()
	pm.version++
}

// IsEnabled returns the monitor enabled state
func (pm *PerformanceMonitor) IsEnabled() bool {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.enabled
}

// SetAlertEnabled sets the alert enabled state
func (pm *PerformanceMonitor) SetAlertEnabled(enabled bool) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.alertEnabled = enabled
	pm.updatedAt = time.Now().Unix()
	pm.version++
}

// IsAlertEnabled returns the alert enabled state
func (pm *PerformanceMonitor) IsAlertEnabled() bool {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.alertEnabled
}

// GetVersion returns the current version
func (pm *PerformanceMonitor) GetVersion() int64 {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.version
}

// GetUpdatedAt returns the last update timestamp
func (pm *PerformanceMonitor) GetUpdatedAt() int64 {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.updatedAt
}

// GetCreatedAt returns the creation timestamp
func (pm *PerformanceMonitor) GetCreatedAt() int64 {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.createdAt
}

// Clone creates a deep copy of the PerformanceMonitor
func (pm *PerformanceMonitor) Clone() *PerformanceMonitor {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	clone := &PerformanceMonitor{
		metrics:      make(map[string]*Metric),
		alerts:       make(map[string]*Alert),
		thresholds:   make(map[string]float64),
		enabled:      pm.enabled,
		alertEnabled: pm.alertEnabled,
		version:      pm.version,
		createdAt:    pm.createdAt,
		updatedAt:    pm.updatedAt,
	}

	// Deep copy metrics
	for k, v := range pm.metrics {
		clone.metrics[k] = &Metric{
			Name:        v.Name,
			Value:       v.Value,
			Unit:        v.Unit,
			Category:    v.Category,
			Description: v.Description,
			Timestamp:   v.Timestamp,
			History:     append([]float64{}, v.History...),
			MaxHistory:  v.MaxHistory,
		}
	}

	// Deep copy alerts
	for k, v := range pm.alerts {
		clone.alerts[k] = &Alert{
			ID:          v.ID,
			MetricName:  v.MetricName,
			Threshold:   v.Threshold,
			Operator:    v.Operator,
			Message:     v.Message,
			Severity:    v.Severity,
			Enabled:     v.Enabled,
			Triggered:   v.Triggered,
			LastTrigger: v.LastTrigger,
			CreatedAt:   v.CreatedAt,
		}
	}

	// Copy thresholds
	for k, v := range pm.thresholds {
		clone.thresholds[k] = v
	}

	return clone
}

// Reset resets the performance monitor
func (pm *PerformanceMonitor) Reset() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.metrics = make(map[string]*Metric)
	pm.alerts = make(map[string]*Alert)
	pm.thresholds = make(map[string]float64)
	pm.updatedAt = time.Now().Unix()
	pm.version++
}

// StartMonitoring starts monitoring with a context
func (pm *PerformanceMonitor) StartMonitoring(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if pm.IsEnabled() {
					pm.collectSystemMetrics()
				}
			}
		}
	}()
}

// collectSystemMetrics collects system-level metrics
func (pm *PerformanceMonitor) collectSystemMetrics() {
	// This is a placeholder for system metrics collection
	// In a real implementation, you would collect actual system metrics
	// like CPU usage, memory usage, disk I/O, network I/O, etc.

	// Example: Collect timestamp as a metric
	_ = pm.SetMetric("system_timestamp", float64(time.Now().Unix()), "seconds", "system", "Current system timestamp")
}

// ===== CONTEXT-DRIVEN METHODS =====

// StartCalculation starts timing a calculation
func (pm *PerformanceMonitor) StartCalculation(operation string) *CalculationTimer {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	timer := &CalculationTimer{
		Operation: operation,
		StartTime: time.Now().UnixNano(),
		Tags:      make(map[string]string),
	}

	pm.calculations[operation] = timer
	pm.updatedAt = time.Now().Unix()
	pm.version++

	return timer
}

// EndCalculation ends timing a calculation
func (pm *PerformanceMonitor) EndCalculation(timer *CalculationTimer) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if timer != nil {
		timer.EndTime = time.Now().UnixNano()
		timer.Duration = timer.EndTime - timer.StartTime
		pm.updatedAt = time.Now().Unix()
		pm.version++
	}
}

// RecordMetric records a performance metric
func (pm *PerformanceMonitor) RecordMetric(metric string, value float64, tags map[string]string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	now := time.Now().Unix()
	pm.metrics[metric] = &Metric{
		Name:      metric,
		Value:     value,
		Tags:      tags,
		Timestamp: now,
		Count:     1,
		Min:       value,
		Max:       value,
		Avg:       value,
		Sum:       value,
	}

	pm.updatedAt = now
	pm.version++
}

// RecordMetricWithContext records a performance metric with context
func (pm *PerformanceMonitor) RecordMetricWithContext(ctx context.Context, metric string, value float64, tags map[string]string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	now := time.Now().Unix()
	contextID := pm.getContextID(ctx)

	// Store in global metrics
	pm.metrics[metric] = &Metric{
		Name:      metric,
		Value:     value,
		Tags:      tags,
		Timestamp: now,
		Count:     1,
		Min:       value,
		Max:       value,
		Avg:       value,
		Sum:       value,
	}

	// Store in context-specific metrics
	if pm.contextMetrics[contextID] == nil {
		pm.contextMetrics[contextID] = make(map[string]*Metric)
	}
	pm.contextMetrics[contextID][metric] = &Metric{
		Name:      metric,
		Value:     value,
		Tags:      tags,
		Timestamp: now,
		Count:     1,
		Min:       value,
		Max:       value,
		Avg:       value,
		Sum:       value,
	}

	pm.updatedAt = now
	pm.version++
}

// GetMetrics returns all recorded metrics
func (pm *PerformanceMonitor) GetMetrics() map[string]Metric {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	result := make(map[string]Metric)
	for k, v := range pm.metrics {
		result[k] = Metric{
			Name:      v.Name,
			Value:     v.Value,
			Tags:      v.Tags,
			Timestamp: v.Timestamp,
			Count:     v.Count,
			Min:       v.Min,
			Max:       v.Max,
			Avg:       v.Avg,
			Sum:       v.Sum,
		}
	}
	return result
}

// GetMetric returns a specific metric
func (pm *PerformanceMonitor) GetMetric(name string) (*Metric, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	metric, exists := pm.metrics[name]
	if !exists {
		return nil, fmt.Errorf("metric '%s' not found", name)
	}

	return metric, nil
}

// GetMetricWithContext returns a specific metric with context
func (pm *PerformanceMonitor) GetMetricWithContext(ctx context.Context, name string) (Metric, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	contextID := pm.getContextID(ctx)

	// Try context-specific first
	if contextMetrics, exists := pm.contextMetrics[contextID]; exists {
		if metric, exists := contextMetrics[name]; exists {
			return Metric{
				Name:      metric.Name,
				Value:     metric.Value,
				Tags:      metric.Tags,
				Timestamp: metric.Timestamp,
				Count:     metric.Count,
				Min:       metric.Min,
				Max:       metric.Max,
				Avg:       metric.Avg,
				Sum:       metric.Sum,
			}, nil
		}
	}

	// Fall back to global metrics
	metric, exists := pm.metrics[name]
	if !exists {
		return Metric{}, fmt.Errorf("metric '%s' not found", name)
	}

	return Metric{
		Name:      metric.Name,
		Value:     metric.Value,
		Tags:      metric.Tags,
		Timestamp: metric.Timestamp,
		Count:     metric.Count,
		Min:       metric.Min,
		Max:       metric.Max,
		Avg:       metric.Avg,
		Sum:       metric.Sum,
	}, nil
}

// GetMetricsByTag returns metrics filtered by tag
func (pm *PerformanceMonitor) GetMetricsByTag(tag string, value string) map[string]Metric {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	result := make(map[string]Metric)
	for k, v := range pm.metrics {
		if tagValue, exists := v.Tags[tag]; exists && tagValue == value {
			result[k] = Metric{
				Name:      v.Name,
				Value:     v.Value,
				Tags:      v.Tags,
				Timestamp: v.Timestamp,
				Count:     v.Count,
				Min:       v.Min,
				Max:       v.Max,
				Avg:       v.Avg,
				Sum:       v.Sum,
			}
		}
	}
	return result
}

// GetMetricsByTagWithContext returns metrics filtered by tag with context
func (pm *PerformanceMonitor) GetMetricsByTagWithContext(ctx context.Context, tag string, value string) map[string]Metric {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	contextID := pm.getContextID(ctx)
	result := make(map[string]Metric)

	// Check context-specific metrics first
	if contextMetrics, exists := pm.contextMetrics[contextID]; exists {
		for k, v := range contextMetrics {
			if tagValue, exists := v.Tags[tag]; exists && tagValue == value {
				result[k] = Metric{
					Name:      v.Name,
					Value:     v.Value,
					Tags:      v.Tags,
					Timestamp: v.Timestamp,
					Count:     v.Count,
					Min:       v.Min,
					Max:       v.Max,
					Avg:       v.Avg,
					Sum:       v.Sum,
				}
			}
		}
	}

	// Also include global metrics
	for k, v := range pm.metrics {
		if tagValue, exists := v.Tags[tag]; exists && tagValue == value {
			result[k] = Metric{
				Name:      v.Name,
				Value:     v.Value,
				Tags:      v.Tags,
				Timestamp: v.Timestamp,
				Count:     v.Count,
				Min:       v.Min,
				Max:       v.Max,
				Avg:       v.Avg,
				Sum:       v.Sum,
			}
		}
	}

	return result
}

// ClearMetricsWithContext clears all metrics with context
func (pm *PerformanceMonitor) ClearMetricsWithContext(ctx context.Context) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	contextID := pm.getContextID(ctx)
	delete(pm.contextMetrics, contextID)
	pm.updatedAt = time.Now().Unix()
	pm.version++
}

// ExportMetrics exports metrics in a specific format
func (pm *PerformanceMonitor) ExportMetrics(format string) ([]byte, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	switch format {
	case "json":
		return json.Marshal(pm.metrics)
	case "prometheus":
		return pm.exportPrometheusFormat(), nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}

// ExportMetricsWithContext exports metrics with context
func (pm *PerformanceMonitor) ExportMetricsWithContext(ctx context.Context, format string) ([]byte, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	contextID := pm.getContextID(ctx)
	contextMetrics := pm.contextMetrics[contextID]

	switch format {
	case "json":
		return json.Marshal(contextMetrics)
	case "prometheus":
		return pm.exportPrometheusFormatWithContext(contextMetrics), nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}

// GetPerformanceReport returns a performance report
func (pm *PerformanceMonitor) GetPerformanceReport() *PerformanceReport {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	now := time.Now().Unix()
	calculations := make([]CalculationTimer, 0, len(pm.calculations))
	for _, timer := range pm.calculations {
		calculations = append(calculations, *timer)
	}

	alerts := make([]Alert, 0, len(pm.alerts))
	for _, alert := range pm.alerts {
		alerts = append(alerts, Alert{
			ID:        alert.ID,
			Metric:    alert.MetricName,
			Threshold: alert.Threshold,
			Value:     0, // Will be filled by caller
			Operator:  alert.Operator,
			Severity:  alert.Severity,
			Message:   alert.Message,
			Tags:      make(map[string]string),
			Timestamp: alert.CreatedAt,
			IsActive:  alert.Triggered,
		})
	}

	metrics := make(map[string]Metric)
	for k, v := range pm.metrics {
		metrics[k] = Metric{
			Name:      v.Name,
			Value:     v.Value,
			Tags:      v.Tags,
			Timestamp: v.Timestamp,
			Count:     v.Count,
			Min:       v.Min,
			Max:       v.Max,
			Avg:       v.Avg,
			Sum:       v.Sum,
		}
	}

	return &PerformanceReport{
		Timestamp:    now,
		Duration:     now - pm.createdAt,
		Metrics:      metrics,
		Calculations: calculations,
		Alerts:       alerts,
		Summary: PerformanceSummary{
			TotalCalculations: int64(len(calculations)),
			TotalMetrics:      int64(len(metrics)),
			ActiveAlerts:      int64(len(alerts)),
		},
	}
}

// GetPerformanceReportWithContext returns a performance report with context
func (pm *PerformanceMonitor) GetPerformanceReportWithContext(ctx context.Context) *PerformanceReport {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	contextID := pm.getContextID(ctx)
	now := time.Now().Unix()

	// Get context-specific calculations
	calculations := make([]CalculationTimer, 0)
	if contextTimers, exists := pm.contextTimers[contextID]; exists {
		for _, timer := range contextTimers {
			calculations = append(calculations, *timer)
		}
	}

	// Get context-specific metrics
	metrics := make(map[string]Metric)
	if contextMetrics, exists := pm.contextMetrics[contextID]; exists {
		for k, v := range contextMetrics {
			metrics[k] = Metric{
				Name:      v.Name,
				Value:     v.Value,
				Tags:      v.Tags,
				Timestamp: v.Timestamp,
				Count:     v.Count,
				Min:       v.Min,
				Max:       v.Max,
				Avg:       v.Avg,
				Sum:       v.Sum,
			}
		}
	}

	// Get alerts (global)
	alerts := make([]Alert, 0, len(pm.alerts))
	for _, alert := range pm.alerts {
		alerts = append(alerts, Alert{
			ID:        alert.ID,
			Metric:    alert.MetricName,
			Threshold: alert.Threshold,
			Value:     0, // Will be filled by caller
			Operator:  alert.Operator,
			Severity:  alert.Severity,
			Message:   alert.Message,
			Tags:      make(map[string]string),
			Timestamp: alert.CreatedAt,
			IsActive:  alert.Triggered,
		})
	}

	return &PerformanceReport{
		Timestamp:    now,
		Duration:     now - pm.createdAt,
		Metrics:      metrics,
		Calculations: calculations,
		Alerts:       alerts,
		Summary: PerformanceSummary{
			TotalCalculations: int64(len(calculations)),
			TotalMetrics:      int64(len(metrics)),
			ActiveAlerts:      int64(len(alerts)),
		},
	}
}

// SetAlertThreshold sets an alert threshold for a metric
func (pm *PerformanceMonitor) SetAlertThreshold(metric string, threshold float64, operator string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.thresholds[metric] = threshold
	pm.updatedAt = time.Now().Unix()
	pm.version++

	return nil
}

// SetAlertThresholdWithContext sets an alert threshold with context
func (pm *PerformanceMonitor) SetAlertThresholdWithContext(ctx context.Context, metric string, threshold float64, operator string) error {
	return pm.SetAlertThreshold(metric, threshold, operator)
}

// CheckAlerts checks for alert conditions
func (pm *PerformanceMonitor) CheckAlerts() []Alert {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	alerts := make([]Alert, 0)
	for _, alert := range pm.alerts {
		if !alert.Enabled {
			continue
		}

		// Get current metric value
		metric, exists := pm.metrics[alert.MetricName]
		if !exists {
			continue
		}

		shouldTrigger := false
		switch alert.Operator {
		case ">":
			shouldTrigger = metric.Value > alert.Threshold
		case "<":
			shouldTrigger = metric.Value < alert.Threshold
		case ">=":
			shouldTrigger = metric.Value >= alert.Threshold
		case "<=":
			shouldTrigger = metric.Value <= alert.Threshold
		case "==":
			shouldTrigger = metric.Value == alert.Threshold
		case "!=":
			shouldTrigger = metric.Value != alert.Threshold
		}

		if shouldTrigger {
			alerts = append(alerts, Alert{
				ID:        alert.ID,
				Metric:    alert.MetricName,
				Threshold: alert.Threshold,
				Value:     metric.Value,
				Operator:  alert.Operator,
				Severity:  alert.Severity,
				Message:   alert.Message,
				Tags:      make(map[string]string),
				Timestamp: time.Now().Unix(),
				IsActive:  true,
			})
		}
	}

	return alerts
}

// CheckAlertsWithContext checks for alert conditions with context
func (pm *PerformanceMonitor) CheckAlertsWithContext(ctx context.Context) []Alert {
	return pm.CheckAlerts()
}

// RegisterAlertHandler registers an alert handler
func (pm *PerformanceMonitor) RegisterAlertHandler(handler AlertHandler) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if handler == nil {
		return fmt.Errorf("alert handler cannot be nil")
	}

	handlerID := handler.GetHandlerID()
	if handlerID == "" {
		return fmt.Errorf("alert handler ID cannot be empty")
	}

	pm.alertHandlers[handlerID] = handler
	pm.updatedAt = time.Now().Unix()
	pm.version++

	return nil
}

// UnregisterAlertHandler unregisters an alert handler
func (pm *PerformanceMonitor) UnregisterAlertHandler(handlerID string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if _, exists := pm.alertHandlers[handlerID]; !exists {
		return fmt.Errorf("alert handler '%s' not found", handlerID)
	}

	delete(pm.alertHandlers, handlerID)
	pm.updatedAt = time.Now().Unix()
	pm.version++

	return nil
}

// ===== PROMETHEUS-STYLE COUNTERS =====

// IncrementCounter increments a counter
func (pm *PerformanceMonitor) IncrementCounter(name string, tags map[string]string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	now := time.Now().Unix()
	if counter, exists := pm.counters[name]; exists {
		counter.Value++
		counter.Timestamp = now
	} else {
		pm.counters[name] = &Counter{
			Name:        name,
			Value:       1,
			Tags:        tags,
			Description: fmt.Sprintf("Counter for %s", name),
			Timestamp:   now,
		}
	}

	pm.updatedAt = now
	pm.version++
}

// AddToCounter adds a value to a counter
func (pm *PerformanceMonitor) AddToCounter(name string, value float64, tags map[string]string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	now := time.Now().Unix()
	if counter, exists := pm.counters[name]; exists {
		counter.Value += value
		counter.Timestamp = now
	} else {
		pm.counters[name] = &Counter{
			Name:        name,
			Value:       value,
			Tags:        tags,
			Description: fmt.Sprintf("Counter for %s", name),
			Timestamp:   now,
		}
	}

	pm.updatedAt = now
	pm.version++
}

// SetGauge sets a gauge value
func (pm *PerformanceMonitor) SetGauge(name string, value float64, tags map[string]string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	now := time.Now().Unix()
	pm.gauges[name] = &Gauge{
		Name:        name,
		Value:       value,
		Tags:        tags,
		Description: fmt.Sprintf("Gauge for %s", name),
		Timestamp:   now,
	}

	pm.updatedAt = now
	pm.version++
}

// ObserveHistogram observes a value in a histogram
func (pm *PerformanceMonitor) ObserveHistogram(name string, value float64, tags map[string]string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	now := time.Now().Unix()
	if histogram, exists := pm.histograms[name]; exists {
		histogram.Count++
		histogram.Sum += value
		histogram.Timestamp = now

		// Add to appropriate bucket
		bucket := pm.getBucket(value)
		histogram.Buckets[bucket]++
	} else {
		pm.histograms[name] = &Histogram{
			Name:        name,
			Buckets:     make(map[string]int64),
			Sum:         value,
			Count:       1,
			Tags:        tags,
			Description: fmt.Sprintf("Histogram for %s", name),
			Timestamp:   now,
		}
		bucket := pm.getBucket(value)
		pm.histograms[name].Buckets[bucket] = 1
	}

	pm.updatedAt = now
	pm.version++
}

// GetCounter returns a counter
func (pm *PerformanceMonitor) GetCounter(name string) (*Counter, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	counter, exists := pm.counters[name]
	if !exists {
		return nil, fmt.Errorf("counter '%s' not found", name)
	}

	return counter, nil
}

// GetGauge returns a gauge
func (pm *PerformanceMonitor) GetGauge(name string) (*Gauge, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	gauge, exists := pm.gauges[name]
	if !exists {
		return nil, fmt.Errorf("gauge '%s' not found", name)
	}

	return gauge, nil
}

// GetHistogram returns a histogram
func (pm *PerformanceMonitor) GetHistogram(name string) (*Histogram, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	histogram, exists := pm.histograms[name]
	if !exists {
		return nil, fmt.Errorf("histogram '%s' not found", name)
	}

	return histogram, nil
}

// ===== HELPER METHODS =====

// getContextID extracts context ID from context
func (pm *PerformanceMonitor) getContextID(ctx context.Context) string {
	// Simple implementation - in real world, you might use context values
	return fmt.Sprintf("ctx_%d", time.Now().UnixNano())
}

// getBucket determines which bucket a value belongs to
func (pm *PerformanceMonitor) getBucket(value float64) string {
	// Simple bucket implementation
	if value < 0.001 {
		return "0.001"
	} else if value < 0.01 {
		return "0.01"
	} else if value < 0.1 {
		return "0.1"
	} else if value < 1.0 {
		return "1.0"
	} else if value < 10.0 {
		return "10.0"
	} else if value < 100.0 {
		return "100.0"
	} else {
		return "+Inf"
	}
}

// exportPrometheusFormat exports metrics in Prometheus format
func (pm *PerformanceMonitor) exportPrometheusFormat() []byte {
	var result string

	// Export counters
	for _, counter := range pm.counters {
		result += fmt.Sprintf("# TYPE %s counter\n", counter.Name)
		result += fmt.Sprintf("%s{", counter.Name)
		first := true
		for k, v := range counter.Tags {
			if !first {
				result += ","
			}
			result += fmt.Sprintf("%s=\"%s\"", k, v)
			first = false
		}
		result += fmt.Sprintf("} %f %d\n", counter.Value, counter.Timestamp*1000)
	}

	// Export gauges
	for _, gauge := range pm.gauges {
		result += fmt.Sprintf("# TYPE %s gauge\n", gauge.Name)
		result += fmt.Sprintf("%s{", gauge.Name)
		first := true
		for k, v := range gauge.Tags {
			if !first {
				result += ","
			}
			result += fmt.Sprintf("%s=\"%s\"", k, v)
			first = false
		}
		result += fmt.Sprintf("} %f %d\n", gauge.Value, gauge.Timestamp*1000)
	}

	// Export histograms
	for _, histogram := range pm.histograms {
		result += fmt.Sprintf("# TYPE %s histogram\n", histogram.Name)
		for bucket, count := range histogram.Buckets {
			result += fmt.Sprintf("%s_bucket{le=\"%s\"", histogram.Name, bucket)
			first := true
			for k, v := range histogram.Tags {
				if !first {
					result += ","
				}
				result += fmt.Sprintf("%s=\"%s\"", k, v)
				first = false
			}
			result += fmt.Sprintf("} %d %d\n", count, histogram.Timestamp*1000)
		}
		result += fmt.Sprintf("%s_sum{", histogram.Name)
		first := true
		for k, v := range histogram.Tags {
			if !first {
				result += ","
			}
			result += fmt.Sprintf("%s=\"%s\"", k, v)
			first = false
		}
		result += fmt.Sprintf("} %f %d\n", histogram.Sum, histogram.Timestamp*1000)
		result += fmt.Sprintf("%s_count{", histogram.Name)
		first = true
		for k, v := range histogram.Tags {
			if !first {
				result += ","
			}
			result += fmt.Sprintf("%s=\"%s\"", k, v)
			first = false
		}
		result += fmt.Sprintf("} %d %d\n", histogram.Count, histogram.Timestamp*1000)
	}

	return []byte(result)
}

// exportPrometheusFormatWithContext exports context-specific metrics in Prometheus format
func (pm *PerformanceMonitor) exportPrometheusFormatWithContext(contextMetrics map[string]*Metric) []byte {
	var result string

	// Export context-specific metrics as gauges
	for _, metric := range contextMetrics {
		result += fmt.Sprintf("# TYPE %s gauge\n", metric.Name)
		result += fmt.Sprintf("%s{", metric.Name)
		first := true
		for k, v := range metric.Tags {
			if !first {
				result += ","
			}
			result += fmt.Sprintf("%s=\"%s\"", k, v)
			first = false
		}
		result += fmt.Sprintf("} %f %d\n", metric.Value, metric.Timestamp*1000)
	}

	return []byte(result)
}
