package monitoring

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// PerformanceMonitor implements the PerformanceMonitor interface
type PerformanceMonitor struct {
	mu           sync.RWMutex
	metrics      map[string]*Metric
	alerts       map[string]*Alert
	thresholds   map[string]float64
	enabled      bool
	alertEnabled bool
	version      int64
	createdAt    int64
	updatedAt    int64
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

// NewPerformanceMonitor creates a new PerformanceMonitor instance
func NewPerformanceMonitor() *PerformanceMonitor {
	now := time.Now().Unix()
	return &PerformanceMonitor{
		metrics:      make(map[string]*Metric),
		alerts:       make(map[string]*Alert),
		thresholds:   make(map[string]float64),
		enabled:      true,
		alertEnabled: true,
		version:      1,
		createdAt:    now,
		updatedAt:    now,
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

// GetMetric gets a performance metric
func (pm *PerformanceMonitor) GetMetric(name string) (*Metric, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	metric, exists := pm.metrics[name]
	if !exists {
		return nil, fmt.Errorf("metric '%s' not found", name)
	}

	return metric, nil
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
	pm.SetMetric("system_timestamp", float64(time.Now().Unix()), "seconds", "system", "Current system timestamp")
}
