package monitoring

import "context"

// PerformanceMonitor defines the interface for performance monitoring
type PerformanceMonitor interface {
	// StartCalculation starts timing a calculation
	StartCalculation(operation string) *CalculationTimer

	// EndCalculation ends timing a calculation
	EndCalculation(timer *CalculationTimer)

	// RecordMetric records a performance metric
	RecordMetric(metric string, value float64, tags map[string]string)

	// RecordMetricWithContext records a performance metric with context
	RecordMetricWithContext(ctx context.Context, metric string, value float64, tags map[string]string)

	// GetMetrics returns all recorded metrics
	GetMetrics() map[string]Metric

	// GetMetric returns a specific metric
	GetMetric(name string) (Metric, error)

	// GetMetricWithContext returns a specific metric with context
	GetMetricWithContext(ctx context.Context, name string) (Metric, error)

	// GetMetricsByTag returns metrics filtered by tag
	GetMetricsByTag(tag string, value string) map[string]Metric

	// GetMetricsByTagWithContext returns metrics filtered by tag with context
	GetMetricsByTagWithContext(ctx context.Context, tag string, value string) map[string]Metric

	// ClearMetrics clears all metrics
	ClearMetrics()

	// ClearMetricsWithContext clears all metrics with context
	ClearMetricsWithContext(ctx context.Context)

	// ExportMetrics exports metrics in a specific format
	ExportMetrics(format string) ([]byte, error)

	// ExportMetricsWithContext exports metrics with context
	ExportMetricsWithContext(ctx context.Context, format string) ([]byte, error)

	// GetPerformanceReport returns a performance report
	GetPerformanceReport() *PerformanceReport

	// GetPerformanceReportWithContext returns a performance report with context
	GetPerformanceReportWithContext(ctx context.Context) *PerformanceReport

	// SetAlertThreshold sets an alert threshold for a metric
	SetAlertThreshold(metric string, threshold float64, operator string) error

	// SetAlertThresholdWithContext sets an alert threshold with context
	SetAlertThresholdWithContext(ctx context.Context, metric string, threshold float64, operator string) error

	// CheckAlerts checks for alert conditions
	CheckAlerts() []Alert

	// CheckAlertsWithContext checks for alert conditions with context
	CheckAlertsWithContext(ctx context.Context) []Alert

	// RegisterAlertHandler registers an alert handler
	RegisterAlertHandler(handler AlertHandler) error

	// UnregisterAlertHandler unregisters an alert handler
	UnregisterAlertHandler(handlerID string) error
}

// CalculationTimer represents a timer for calculations
type CalculationTimer struct {
	Operation string            `json:"operation"`
	StartTime int64             `json:"start_time"`
	EndTime   int64             `json:"end_time"`
	Duration  int64             `json:"duration"`
	Tags      map[string]string `json:"tags"`
}

// Metric represents a performance metric
type Metric struct {
	Name      string            `json:"name"`
	Value     float64           `json:"value"`
	Tags      map[string]string `json:"tags"`
	Timestamp int64             `json:"timestamp"`
	Count     int64             `json:"count"`
	Min       float64           `json:"min"`
	Max       float64           `json:"max"`
	Avg       float64           `json:"avg"`
	Sum       float64           `json:"sum"`
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

// Alert represents a performance alert
type Alert struct {
	ID        string            `json:"id"`
	Metric    string            `json:"metric"`
	Threshold float64           `json:"threshold"`
	Value     float64           `json:"value"`
	Operator  string            `json:"operator"`
	Severity  string            `json:"severity"`
	Message   string            `json:"message"`
	Tags      map[string]string `json:"tags"`
	Timestamp int64             `json:"timestamp"`
	IsActive  bool              `json:"is_active"`
}

// AlertHandler defines the interface for alert handlers
type AlertHandler interface {
	HandleAlert(alert Alert)
	GetHandlerID() string
	GetSeverity() string
}
