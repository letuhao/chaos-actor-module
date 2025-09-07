package interfaces

import (
	"time"
)

// Monitor represents a monitoring interface
type Monitor interface {
	// Start starts monitoring
	Start() error

	// Stop stops monitoring
	Stop() error

	// GetMetrics returns current metrics
	GetMetrics() *Metrics

	// GetHealth returns health status
	GetHealth() *HealthStatus

	// RecordMetric records a metric
	RecordMetric(name string, value float64, tags map[string]string)

	// IncrementCounter increments a counter
	IncrementCounter(name string, tags map[string]string)

	// SetGauge sets a gauge value
	SetGauge(name string, value float64, tags map[string]string)

	// RecordHistogram records a histogram value
	RecordHistogram(name string, value float64, tags map[string]string)
}

// Metrics represents system metrics
type Metrics struct {
	// Counters are counter metrics
	Counters map[string]CounterMetric `json:"counters"`

	// Gauges are gauge metrics
	Gauges map[string]GaugeMetric `json:"gauges"`

	// Histograms are histogram metrics
	Histograms map[string]HistogramMetric `json:"histograms"`

	// Timestamp is when the metrics were collected
	Timestamp time.Time `json:"timestamp"`
}

// CounterMetric represents a counter metric
type CounterMetric struct {
	// Value is the counter value
	Value int64 `json:"value"`

	// Tags are metric tags
	Tags map[string]string `json:"tags"`

	// LastUpdated is when the metric was last updated
	LastUpdated time.Time `json:"last_updated"`
}

// GaugeMetric represents a gauge metric
type GaugeMetric struct {
	// Value is the gauge value
	Value float64 `json:"value"`

	// Tags are metric tags
	Tags map[string]string `json:"tags"`

	// LastUpdated is when the metric was last updated
	LastUpdated time.Time `json:"last_updated"`
}

// HistogramMetric represents a histogram metric
type HistogramMetric struct {
	// Count is the number of observations
	Count int64 `json:"count"`

	// Sum is the sum of all observations
	Sum float64 `json:"sum"`

	// Min is the minimum observation
	Min float64 `json:"min"`

	// Max is the maximum observation
	Max float64 `json:"max"`

	// Mean is the mean of observations
	Mean float64 `json:"mean"`

	// Percentiles are percentile values
	Percentiles map[string]float64 `json:"percentiles"`

	// Tags are metric tags
	Tags map[string]string `json:"tags"`

	// LastUpdated is when the metric was last updated
	LastUpdated time.Time `json:"last_updated"`
}

// HealthStatus represents system health status
type HealthStatus struct {
	// Overall is the overall health status
	Overall bool `json:"overall"`

	// Checks are individual health checks
	Checks map[string]bool `json:"checks"`

	// Message is the health message
	Message string `json:"message,omitempty"`

	// Timestamp is when the health was checked
	Timestamp time.Time `json:"timestamp"`
}

// IsHealthy checks if the system is healthy
func (hs *HealthStatus) IsHealthy() bool {
	return hs.Overall
}

// IsUnhealthy checks if the system is unhealthy
func (hs *HealthStatus) IsUnhealthy() bool {
	return !hs.Overall
}

// GetCheckStatus returns the status of a specific check
func (hs *HealthStatus) GetCheckStatus(name string) (bool, bool) {
	status, exists := hs.Checks[name]
	return status, exists
}

// SetCheckStatus sets the status of a specific check
func (hs *HealthStatus) SetCheckStatus(name string, status bool) {
	if hs.Checks == nil {
		hs.Checks = make(map[string]bool)
	}
	hs.Checks[name] = status
}

// UpdateOverall updates the overall health status
func (hs *HealthStatus) UpdateOverall() {
	hs.Overall = true
	for _, status := range hs.Checks {
		if !status {
			hs.Overall = false
			break
		}
	}
}

// Alert represents an alert
type Alert struct {
	// Name is the alert name
	Name string `json:"name"`

	// Severity is the alert severity
	Severity string `json:"severity"`

	// Message is the alert message
	Message string `json:"message"`

	// Tags are alert tags
	Tags map[string]string `json:"tags"`

	// Timestamp is when the alert was triggered
	Timestamp time.Time `json:"timestamp"`

	// Resolved indicates if the alert is resolved
	Resolved bool `json:"resolved"`
}

// IsResolved checks if the alert is resolved
func (a *Alert) IsResolved() bool {
	return a.Resolved
}

// Resolve resolves the alert
func (a *Alert) Resolve() {
	a.Resolved = true
}

// AlertHandler represents an alert handler
type AlertHandler interface {
	// Handle handles an alert
	Handle(alert *Alert) error

	// CanHandle checks if this handler can handle the alert
	CanHandle(alert *Alert) bool

	// GetName returns the handler name
	GetName() string
}

// PerformanceMonitor represents a performance monitor
type PerformanceMonitor interface {
	// StartTimer starts a performance timer
	StartTimer(name string) *Timer

	// RecordDuration records a duration
	RecordDuration(name string, duration time.Duration, tags map[string]string)

	// RecordMemoryUsage records memory usage
	RecordMemoryUsage(usage int64, tags map[string]string)

	// RecordGoroutineCount records goroutine count
	RecordGoroutineCount(count int, tags map[string]string)

	// GetPerformanceMetrics returns performance metrics
	GetPerformanceMetrics() *PerformanceMetrics
}

// Timer represents a performance timer
type Timer struct {
	// Name is the timer name
	Name string

	// StartTime is when the timer started
	StartTime time.Time

	// Tags are timer tags
	Tags map[string]string

	// Monitor is the performance monitor
	Monitor PerformanceMonitor
}

// Stop stops the timer and records the duration
func (t *Timer) Stop() {
	if t.Monitor != nil {
		duration := time.Since(t.StartTime)
		t.Monitor.RecordDuration(t.Name, duration, t.Tags)
	}
}

// PerformanceMetrics represents performance metrics
type PerformanceMetrics struct {
	// Timers are timer metrics
	Timers map[string]HistogramMetric `json:"timers"`

	// MemoryUsage is memory usage metrics
	MemoryUsage GaugeMetric `json:"memory_usage"`

	// GoroutineCount is goroutine count metrics
	GoroutineCount GaugeMetric `json:"goroutine_count"`

	// Timestamp is when the metrics were collected
	Timestamp time.Time `json:"timestamp"`
}

// GetTimerMetrics returns timer metrics for a specific timer
func (pm *PerformanceMetrics) GetTimerMetrics(name string) (HistogramMetric, bool) {
	metrics, exists := pm.Timers[name]
	return metrics, exists
}

// GetMemoryUsage returns memory usage metrics
func (pm *PerformanceMetrics) GetMemoryUsage() GaugeMetric {
	return pm.MemoryUsage
}

// GetGoroutineCount returns goroutine count metrics
func (pm *PerformanceMetrics) GetGoroutineCount() GaugeMetric {
	return pm.GoroutineCount
}
