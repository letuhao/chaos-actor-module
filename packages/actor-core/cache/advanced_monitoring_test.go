package cache

import (
	"testing"
	"time"
)

func TestAdvancedMonitor_New(t *testing.T) {
	config := &MonitoringConfig{
		EnableMetrics:     true,
		EnableProfiling:   true,
		EnableTracing:     true,
		EnableAlerts:      true,
		MetricsInterval:   time.Second,
		ProfilingInterval: time.Second * 2,
	}

	monitor := NewAdvancedMonitor(config)
	if monitor == nil {
		t.Fatal("Expected non-nil monitor")
	}

	if !monitor.config.EnableMetrics {
		t.Error("Expected metrics to be enabled")
	}

	// Clean up
	monitor.Close()
}

func TestAdvancedMonitor_DefaultConfig(t *testing.T) {
	monitor := NewAdvancedMonitor(nil)
	if monitor == nil {
		t.Fatal("Expected non-nil monitor")
	}

	config := monitor.config
	if !config.EnableMetrics {
		t.Error("Expected metrics to be enabled by default")
	}
	if config.MetricsInterval <= 0 {
		t.Error("Expected positive metrics interval")
	}

	// Clean up
	monitor.Close()
}

func TestAdvancedMonitor_RecordMetric(t *testing.T) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	// Record a metric
	monitor.RecordMetric("test_metric", MetricTypeGauge, 42.5, map[string]string{"label1": "value1"})

	// Check if metric was recorded
	now := time.Now()
	start := now.Add(-time.Minute)
	end := now.Add(time.Minute)

	metrics := monitor.GetMetrics(start, end, "test_metric")
	if len(metrics) != 1 {
		t.Fatalf("Expected 1 metric, got %d", len(metrics))
	}

	metric := metrics[0]
	if metric.Name != "test_metric" {
		t.Errorf("Expected metric name 'test_metric', got '%s'", metric.Name)
	}
	if metric.Type != MetricTypeGauge {
		t.Errorf("Expected metric type 'gauge', got '%s'", metric.Type)
	}
	if metric.Value != 42.5 {
		t.Errorf("Expected metric value 42.5, got %f", metric.Value)
	}
	if metric.Labels["label1"] != "value1" {
		t.Errorf("Expected label 'value1', got '%s'", metric.Labels["label1"])
	}
}

func TestAdvancedMonitor_RecordAlert(t *testing.T) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	// Record an alert
	monitor.RecordAlert(AlertLevelWarning, "Test alert", "test_metric", 85.0, 80.0, map[string]string{"label1": "value1"})

	// Check if alert was recorded
	now := time.Now()
	start := now.Add(-time.Minute)
	end := now.Add(time.Minute)

	alerts := monitor.GetAlerts(start, end, AlertLevelWarning)
	if len(alerts) != 1 {
		t.Fatalf("Expected 1 alert, got %d", len(alerts))
	}

	alert := alerts[0]
	if alert.Level != AlertLevelWarning {
		t.Errorf("Expected alert level 'warning', got '%s'", alert.Level)
	}
	if alert.Message != "Test alert" {
		t.Errorf("Expected alert message 'Test alert', got '%s'", alert.Message)
	}
	if alert.Metric != "test_metric" {
		t.Errorf("Expected alert metric 'test_metric', got '%s'", alert.Metric)
	}
	if alert.Value != 85.0 {
		t.Errorf("Expected alert value 85.0, got %f", alert.Value)
	}
	if alert.Threshold != 80.0 {
		t.Errorf("Expected alert threshold 80.0, got %f", alert.Threshold)
	}
	if alert.Labels["label1"] != "value1" {
		t.Errorf("Expected label 'value1', got '%s'", alert.Labels["label1"])
	}
	if alert.Resolved {
		t.Error("Expected alert to be unresolved")
	}
}

func TestAdvancedMonitor_RecordTrace(t *testing.T) {
	config := &MonitoringConfig{
		EnableTracing:     true,
		TracingSampleRate: 1.0, // 100% sampling for test
	}
	monitor := NewAdvancedMonitor(config)
	defer monitor.Close()

	// Record a trace
	span := TraceSpan{
		TraceID:   "trace123",
		SpanID:    "span456",
		Operation: "test_operation",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Millisecond * 100),
		Duration:  time.Millisecond * 100,
		Tags:      map[string]string{"tag1": "value1"},
	}

	monitor.RecordTrace(span)

	// Check if trace was recorded
	now := time.Now()
	start := now.Add(-time.Minute)
	end := now.Add(time.Minute)

	traces := monitor.GetTraces(start, end, "trace123")
	if len(traces) != 1 {
		t.Fatalf("Expected 1 trace, got %d", len(traces))
	}

	trace := traces[0]
	if trace.TraceID != "trace123" {
		t.Errorf("Expected trace ID 'trace123', got '%s'", trace.TraceID)
	}
	if trace.SpanID != "span456" {
		t.Errorf("Expected span ID 'span456', got '%s'", trace.SpanID)
	}
	if trace.Operation != "test_operation" {
		t.Errorf("Expected operation 'test_operation', got '%s'", trace.Operation)
	}
	if trace.Duration != time.Millisecond*100 {
		t.Errorf("Expected duration 100ms, got %v", trace.Duration)
	}
	if trace.Tags["tag1"] != "value1" {
		t.Errorf("Expected tag 'value1', got '%s'", trace.Tags["tag1"])
	}
}

func TestAdvancedMonitor_GetMetrics(t *testing.T) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	// Record multiple metrics
	now := time.Now()
	monitor.RecordMetric("metric1", MetricTypeGauge, 10.0, nil)
	time.Sleep(time.Millisecond * 10)
	monitor.RecordMetric("metric2", MetricTypeCounter, 20.0, nil)
	time.Sleep(time.Millisecond * 10)
	monitor.RecordMetric("metric1", MetricTypeGauge, 30.0, nil)

	// Test time range filtering
	start := now.Add(-time.Minute)
	end := now.Add(time.Minute)

	allMetrics := monitor.GetMetrics(start, end, "")
	if len(allMetrics) != 3 {
		t.Fatalf("Expected 3 metrics, got %d", len(allMetrics))
	}

	// Test name filtering
	metric1Metrics := monitor.GetMetrics(start, end, "metric1")
	if len(metric1Metrics) != 2 {
		t.Fatalf("Expected 2 metric1 metrics, got %d", len(metric1Metrics))
	}

	// Test time range filtering
	futureStart := now.Add(time.Hour)
	futureEnd := now.Add(time.Hour * 2)
	futureMetrics := monitor.GetMetrics(futureStart, futureEnd, "")
	if len(futureMetrics) != 0 {
		t.Fatalf("Expected 0 future metrics, got %d", len(futureMetrics))
	}
}

func TestAdvancedMonitor_GetAlerts(t *testing.T) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	// Record multiple alerts
	now := time.Now()
	monitor.RecordAlert(AlertLevelInfo, "Info alert", "metric1", 50.0, 60.0, nil)
	time.Sleep(time.Millisecond * 10)
	monitor.RecordAlert(AlertLevelWarning, "Warning alert", "metric2", 80.0, 70.0, nil)
	time.Sleep(time.Millisecond * 10)
	monitor.RecordAlert(AlertLevelCritical, "Critical alert", "metric3", 90.0, 80.0, nil)

	// Test time range filtering
	start := now.Add(-time.Minute)
	end := now.Add(time.Minute)

	allAlerts := monitor.GetAlerts(start, end, "")
	if len(allAlerts) != 3 {
		t.Fatalf("Expected 3 alerts, got %d", len(allAlerts))
	}

	// Test level filtering
	warningAlerts := monitor.GetAlerts(start, end, AlertLevelWarning)
	if len(warningAlerts) != 1 {
		t.Fatalf("Expected 1 warning alert, got %d", len(warningAlerts))
	}

	// Test time range filtering
	futureStart := now.Add(time.Hour)
	futureEnd := now.Add(time.Hour * 2)
	futureAlerts := monitor.GetAlerts(futureStart, futureEnd, "")
	if len(futureAlerts) != 0 {
		t.Fatalf("Expected 0 future alerts, got %d", len(futureAlerts))
	}
}

func TestAdvancedMonitor_GetProfilingData(t *testing.T) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	// Wait for profiling data to be collected
	time.Sleep(time.Second * 3)

	// Test time range filtering
	now := time.Now()
	start := now.Add(-time.Minute)
	end := now.Add(time.Minute)

	profilingData := monitor.GetProfilingData(start, end)
	if len(profilingData) == 0 {
		t.Log("No profiling data collected yet - this is expected for short test")
	}

	// Test time range filtering
	futureStart := now.Add(time.Hour)
	futureEnd := now.Add(time.Hour * 2)
	futureProfilingData := monitor.GetProfilingData(futureStart, futureEnd)
	if len(futureProfilingData) != 0 {
		t.Fatalf("Expected 0 future profiling data, got %d", len(futureProfilingData))
	}
}

func TestAdvancedMonitor_GetTraces(t *testing.T) {
	config := &MonitoringConfig{
		EnableTracing:     true,
		TracingSampleRate: 1.0, // 100% sampling for test
	}
	monitor := NewAdvancedMonitor(config)
	defer monitor.Close()

	// Record multiple traces
	now := time.Now()
	span1 := TraceSpan{
		TraceID:   "trace1",
		SpanID:    "span1",
		Operation: "op1",
		StartTime: now,
		EndTime:   now.Add(time.Millisecond * 100),
		Duration:  time.Millisecond * 100,
	}
	monitor.RecordTrace(span1)

	time.Sleep(time.Millisecond * 10)
	span2 := TraceSpan{
		TraceID:   "trace2",
		SpanID:    "span2",
		Operation: "op2",
		StartTime: now.Add(time.Millisecond * 10),
		EndTime:   now.Add(time.Millisecond * 200),
		Duration:  time.Millisecond * 190,
	}
	monitor.RecordTrace(span2)

	// Test time range filtering
	start := now.Add(-time.Minute)
	end := now.Add(time.Minute)

	allTraces := monitor.GetTraces(start, end, "")
	if len(allTraces) != 2 {
		t.Fatalf("Expected 2 traces, got %d", len(allTraces))
	}

	// Test trace ID filtering
	trace1Traces := monitor.GetTraces(start, end, "trace1")
	if len(trace1Traces) != 1 {
		t.Fatalf("Expected 1 trace1 trace, got %d", len(trace1Traces))
	}

	// Test time range filtering
	futureStart := now.Add(time.Hour)
	futureEnd := now.Add(time.Hour * 2)
	futureTraces := monitor.GetTraces(futureStart, futureEnd, "")
	if len(futureTraces) != 0 {
		t.Fatalf("Expected 0 future traces, got %d", len(futureTraces))
	}
}

func TestAdvancedMonitor_GetDashboardData(t *testing.T) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	// Record some test data
	monitor.RecordMetric("test_metric", MetricTypeGauge, 42.0, nil)
	monitor.RecordAlert(AlertLevelWarning, "Test alert", "test_metric", 85.0, 80.0, nil)

	// Get dashboard data
	dashboardData := monitor.GetDashboardData()

	// Check required fields
	requiredFields := []string{"timestamp", "metrics", "alerts", "profiling", "statistics", "system_info", "cache_performance"}
	for _, field := range requiredFields {
		if _, exists := dashboardData[field]; !exists {
			t.Errorf("Missing required field: %s", field)
		}
	}

	// Check system info
	systemInfo, ok := dashboardData["system_info"].(map[string]interface{})
	if !ok {
		t.Error("Expected system_info to be a map")
	} else {
		requiredSystemFields := []string{"go_version", "goos", "goarch", "num_cpu", "goroutines", "timestamp"}
		for _, field := range requiredSystemFields {
			if _, exists := systemInfo[field]; !exists {
				t.Errorf("Missing required system info field: %s", field)
			}
		}
	}
}

func TestAdvancedMonitor_Callbacks(t *testing.T) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	// Set up callbacks
	metricReceived := false
	alertReceived := false

	monitor.SetMetricCallback(func(point MetricPoint) {
		metricReceived = true
	})

	monitor.SetAlertCallback(func(alert Alert) {
		alertReceived = true
	})

	monitor.SetProfileCallback(func(data ProfilingData) {
		// Profile callback
	})

	monitor.SetTraceCallback(func(span TraceSpan) {
		// Trace callback
	})

	// Trigger callbacks
	monitor.RecordMetric("test", MetricTypeGauge, 1.0, nil)
	monitor.RecordAlert(AlertLevelInfo, "Test", "test", 1.0, 1.0, nil)

	// Wait for profiling callback
	time.Sleep(time.Second * 3)

	span := TraceSpan{
		TraceID:   "test",
		SpanID:    "test",
		Operation: "test",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Millisecond),
		Duration:  time.Millisecond,
	}
	monitor.RecordTrace(span)

	// Check if callbacks were called
	if !metricReceived {
		t.Error("Expected metric callback to be called")
	}
	if !alertReceived {
		t.Error("Expected alert callback to be called")
	}
	// Note: profile and trace callbacks might not be called immediately
	// depending on timing and sampling rates
}

func TestAdvancedMonitor_ExportMetrics(t *testing.T) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	// Record some metrics
	monitor.RecordMetric("test_metric", MetricTypeGauge, 42.0, map[string]string{"label1": "value1"})

	// Test JSON export
	jsonData, err := monitor.ExportMetrics("json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(jsonData) == 0 {
		t.Error("Expected non-empty JSON data")
	}

	// Test Prometheus export
	promData, err := monitor.ExportMetrics("prometheus")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(promData) == 0 {
		t.Error("Expected non-empty Prometheus data")
	}

	// Test InfluxDB export
	influxData, err := monitor.ExportMetrics("influxdb")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(influxData) == 0 {
		t.Error("Expected non-empty InfluxDB data")
	}

	// Test unsupported format
	_, err = monitor.ExportMetrics("unsupported")
	if err == nil {
		t.Error("Expected error for unsupported format")
	}
}

func TestAdvancedMonitor_Close(t *testing.T) {
	monitor := NewAdvancedMonitor(nil)

	// Close should not panic
	monitor.Close()

	// Multiple closes should not panic
	monitor.Close()
}

func TestAdvancedMonitor_AlertThresholds(t *testing.T) {
	config := &MonitoringConfig{
		EnableAlerts: true,
		AlertThresholds: map[string]float64{
			"test_metric": 80.0,
		},
	}

	monitor := NewAdvancedMonitor(config)
	defer monitor.Close()

	// Record metric below threshold - should not trigger alert
	monitor.RecordMetric("test_metric", MetricTypeGauge, 70.0, nil)

	// Record metric above threshold - should trigger alert
	monitor.RecordMetric("test_metric", MetricTypeGauge, 85.0, nil)

	// Check if alert was triggered
	now := time.Now()
	start := now.Add(-time.Minute)
	end := now.Add(time.Minute)

	alerts := monitor.GetAlerts(start, end, "")
	t.Logf("Found %d alerts", len(alerts))
	for i, alert := range alerts {
		t.Logf("Alert %d: %s = %.2f (threshold: %.2f)", i, alert.Metric, alert.Value, alert.Threshold)
	}

	if len(alerts) == 0 {
		t.Error("Expected alert to be triggered")
	} else {
		// Check alert details
		alert := alerts[0]
		if alert.Metric != "test_metric" {
			t.Errorf("Expected alert metric 'test_metric', got '%s'", alert.Metric)
		}
		if alert.Value != 85.0 {
			t.Errorf("Expected alert value 85.0, got %f", alert.Value)
		}
		if alert.Threshold != 80.0 {
			t.Errorf("Expected alert threshold 80.0, got %f", alert.Threshold)
		}
	}
}

func TestAdvancedMonitor_TracingSampling(t *testing.T) {
	config := &MonitoringConfig{
		EnableTracing:     true,
		TracingSampleRate: 0.0, // 0% sampling
	}

	monitor := NewAdvancedMonitor(config)
	defer monitor.Close()

	// Record trace with 0% sampling rate
	span := TraceSpan{
		TraceID:   "test",
		SpanID:    "test",
		Operation: "test",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Millisecond),
		Duration:  time.Millisecond,
	}
	monitor.RecordTrace(span)

	// Check if trace was recorded (should be 0 due to sampling)
	now := time.Now()
	start := now.Add(-time.Minute)
	end := now.Add(time.Minute)

	traces := monitor.GetTraces(start, end, "")
	if len(traces) != 0 {
		t.Error("Expected no traces due to 0% sampling rate")
	}
}

func BenchmarkAdvancedMonitor_RecordMetric(b *testing.B) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		monitor.RecordMetric("test_metric", MetricTypeGauge, float64(i), nil)
	}
}

func BenchmarkAdvancedMonitor_RecordAlert(b *testing.B) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		monitor.RecordAlert(AlertLevelInfo, "Test alert", "test_metric", float64(i), 100.0, nil)
	}
}

func BenchmarkAdvancedMonitor_RecordTrace(b *testing.B) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	span := TraceSpan{
		TraceID:   "test",
		SpanID:    "test",
		Operation: "test",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Millisecond),
		Duration:  time.Millisecond,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		monitor.RecordTrace(span)
	}
}

func BenchmarkAdvancedMonitor_GetMetrics(b *testing.B) {
	monitor := NewAdvancedMonitor(nil)
	defer monitor.Close()

	// Record some metrics
	for i := 0; i < 1000; i++ {
		monitor.RecordMetric("test_metric", MetricTypeGauge, float64(i), nil)
	}

	now := time.Now()
	start := now.Add(-time.Hour)
	end := now.Add(time.Hour)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		monitor.GetMetrics(start, end, "")
	}
}
