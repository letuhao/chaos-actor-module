package unit

import (
	"context"
	"testing"
	"time"

	"actor-core/services/monitoring"
)

func TestPerformanceMonitorContextDriven(t *testing.T) {
	pm := monitoring.NewPerformanceMonitor()

	// Test context-driven metric recording
	ctx := context.Background()
	tags := map[string]string{"service": "test", "version": "1.0"}

	pm.RecordMetricWithContext(ctx, "test_metric", 100.0, tags)

	// Test getting metric with context
	metric, err := pm.GetMetricWithContext(ctx, "test_metric")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if metric.Name != "test_metric" {
		t.Errorf("Expected metric name 'test_metric', got '%s'", metric.Name)
	}

	if metric.Value != 100.0 {
		t.Errorf("Expected metric value 100.0, got %f", metric.Value)
	}

	// Test getting metrics by tag with context
	metricsByTag := pm.GetMetricsByTagWithContext(ctx, "service", "test")
	if len(metricsByTag) != 1 {
		t.Errorf("Expected 1 metric with service=test, got %d", len(metricsByTag))
	}

	// Test performance report with context
	report := pm.GetPerformanceReportWithContext(ctx)
	if report == nil {
		t.Fatal("Expected performance report, got nil")
	}

	if len(report.Metrics) != 1 {
		t.Errorf("Expected 1 metric in report, got %d", len(report.Metrics))
	}

	// Test clearing metrics with context
	pm.ClearMetricsWithContext(ctx)
	// Note: ClearMetricsWithContext only clears context-specific metrics,
	// global metrics are still available
	metricsAfterClear := pm.GetMetricsByTagWithContext(ctx, "service", "test")
	if len(metricsAfterClear) != 1 {
		t.Errorf("Expected 1 metric after context clear (global metric still available), got %d", len(metricsAfterClear))
	}
}

func TestPerformanceMonitorCalculationTiming(t *testing.T) {
	pm := monitoring.NewPerformanceMonitor()

	// Test calculation timing
	timer := pm.StartCalculation("test_operation")
	if timer == nil {
		t.Fatal("Expected timer, got nil")
	}

	if timer.Operation != "test_operation" {
		t.Errorf("Expected operation 'test_operation', got '%s'", timer.Operation)
	}

	// Simulate some work
	time.Sleep(10 * time.Millisecond)

	pm.EndCalculation(timer)

	if timer.Duration == 0 {
		t.Error("Expected non-zero duration")
	}

	// Test performance report includes calculations
	report := pm.GetPerformanceReport()
	if len(report.Calculations) != 1 {
		t.Errorf("Expected 1 calculation in report, got %d", len(report.Calculations))
	}

	if report.Calculations[0].Operation != "test_operation" {
		t.Errorf("Expected calculation operation 'test_operation', got '%s'", report.Calculations[0].Operation)
	}
}

func TestPerformanceMonitorPrometheusCounters(t *testing.T) {
	pm := monitoring.NewPerformanceMonitor()

	// Test counter increment
	tags := map[string]string{"service": "test"}
	pm.IncrementCounter("test_counter", tags)

	counter, err := pm.GetCounter("test_counter")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if counter.Value != 1.0 {
		t.Errorf("Expected counter value 1.0, got %f", counter.Value)
	}

	// Test counter addition
	pm.AddToCounter("test_counter", 5.0, tags)

	counter, err = pm.GetCounter("test_counter")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if counter.Value != 6.0 {
		t.Errorf("Expected counter value 6.0, got %f", counter.Value)
	}
}

func TestPerformanceMonitorPrometheusGauges(t *testing.T) {
	pm := monitoring.NewPerformanceMonitor()

	// Test gauge setting
	tags := map[string]string{"service": "test"}
	pm.SetGauge("test_gauge", 42.0, tags)

	gauge, err := pm.GetGauge("test_gauge")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if gauge.Value != 42.0 {
		t.Errorf("Expected gauge value 42.0, got %f", gauge.Value)
	}

	// Test gauge update
	pm.SetGauge("test_gauge", 84.0, tags)

	gauge, err = pm.GetGauge("test_gauge")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if gauge.Value != 84.0 {
		t.Errorf("Expected gauge value 84.0, got %f", gauge.Value)
	}
}

func TestPerformanceMonitorPrometheusHistograms(t *testing.T) {
	pm := monitoring.NewPerformanceMonitor()

	// Test histogram observation
	tags := map[string]string{"service": "test"}
	pm.ObserveHistogram("test_histogram", 0.5, tags)
	pm.ObserveHistogram("test_histogram", 1.5, tags)
	pm.ObserveHistogram("test_histogram", 2.5, tags)

	histogram, err := pm.GetHistogram("test_histogram")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if histogram.Count != 3 {
		t.Errorf("Expected histogram count 3, got %d", histogram.Count)
	}

	if histogram.Sum != 4.5 {
		t.Errorf("Expected histogram sum 4.5, got %f", histogram.Sum)
	}

	// Check buckets
	if len(histogram.Buckets) == 0 {
		t.Error("Expected histogram to have buckets")
	}
}

func TestPerformanceMonitorExportFormats(t *testing.T) {
	pm := monitoring.NewPerformanceMonitor()

	// Add some metrics
	pm.SetGauge("test_gauge", 42.0, map[string]string{"service": "test"})
	pm.IncrementCounter("test_counter", map[string]string{"service": "test"})
	pm.ObserveHistogram("test_histogram", 1.0, map[string]string{"service": "test"})

	// Test JSON export
	jsonData, err := pm.ExportMetrics("json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(jsonData) == 0 {
		t.Error("Expected non-empty JSON data")
	}

	// Test Prometheus export
	promData, err := pm.ExportMetrics("prometheus")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(promData) == 0 {
		t.Error("Expected non-empty Prometheus data")
	}

	// Test unsupported format
	_, err = pm.ExportMetrics("unsupported")
	if err == nil {
		t.Error("Expected error for unsupported format")
	}
}

func TestPerformanceMonitorExportWithContext(t *testing.T) {
	pm := monitoring.NewPerformanceMonitor()

	ctx := context.Background()
	tags := map[string]string{"service": "test"}

	// Add context-specific metrics
	pm.RecordMetricWithContext(ctx, "context_metric", 100.0, tags)

	// Test JSON export with context
	jsonData, err := pm.ExportMetricsWithContext(ctx, "json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(jsonData) == 0 {
		t.Error("Expected non-empty JSON data")
	}

	// Test Prometheus export with context
	promData, err := pm.ExportMetricsWithContext(ctx, "prometheus")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(promData) == 0 {
		t.Error("Expected non-empty Prometheus data")
	}
}

func TestPerformanceMonitorAlertHandlers(t *testing.T) {
	pm := monitoring.NewPerformanceMonitor()

	// Create a mock alert handler
	handler := &mockAlertHandler{
		id:       "test_handler",
		severity: "high",
	}

	// Test registering alert handler
	err := pm.RegisterAlertHandler(handler)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test unregistering alert handler
	err = pm.UnregisterAlertHandler("test_handler")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test unregistering non-existent handler
	err = pm.UnregisterAlertHandler("non_existent")
	if err == nil {
		t.Error("Expected error for non-existent handler")
	}
}

func TestPerformanceMonitorAlertThresholds(t *testing.T) {
	pm := monitoring.NewPerformanceMonitor()

	// Test setting alert threshold
	err := pm.SetAlertThreshold("test_metric", 100.0, ">")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test setting alert threshold with context
	ctx := context.Background()
	err = pm.SetAlertThresholdWithContext(ctx, "test_metric_ctx", 200.0, "<")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestPerformanceMonitorCheckAlerts(t *testing.T) {
	pm := monitoring.NewPerformanceMonitor()

	// Create an alert
	err := pm.CreateAlert("test_alert", "test_metric", 100.0, ">", "Test alert", "high")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Record a metric that should trigger the alert
	pm.RecordMetric("test_metric", 150.0, map[string]string{"service": "test"})

	// Check alerts
	alerts := pm.CheckAlerts()
	if len(alerts) != 1 {
		t.Errorf("Expected 1 alert, got %d", len(alerts))
	}

	if alerts[0].Metric != "test_metric" {
		t.Errorf("Expected alert metric 'test_metric', got '%s'", alerts[0].Metric)
	}

	if alerts[0].Value != 150.0 {
		t.Errorf("Expected alert value 150.0, got %f", alerts[0].Value)
	}

	// Test checking alerts with context
	ctx := context.Background()
	alertsWithContext := pm.CheckAlertsWithContext(ctx)
	if len(alertsWithContext) != 1 {
		t.Errorf("Expected 1 alert with context, got %d", len(alertsWithContext))
	}
}

// Mock alert handler for testing
type mockAlertHandler struct {
	id       string
	severity string
}

func (h *mockAlertHandler) HandleAlert(alert monitoring.Alert) {
	// Mock implementation
}

func (h *mockAlertHandler) GetHandlerID() string {
	return h.id
}

func (h *mockAlertHandler) GetSeverity() string {
	return h.severity
}
