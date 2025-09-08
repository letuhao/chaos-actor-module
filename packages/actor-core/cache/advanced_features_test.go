package cache

import (
	"context"
	"testing"
	"time"
)

func TestDistributedCache_New(t *testing.T) {
	config := &DistributedCacheConfig{
		EnableDistributed: true,
		ReplicationFactor: 3,
		PartitionCount:    16,
	}

	dc, err := NewDistributedCache(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if dc == nil {
		t.Fatal("Expected non-nil distributed cache")
	}

	if !dc.config.EnableDistributed {
		t.Error("Expected distributed cache to be enabled")
	}
}

func TestDistributedCache_DefaultConfig(t *testing.T) {
	dc, err := NewDistributedCache(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if dc == nil {
		t.Fatal("Expected non-nil distributed cache")
	}

	config := dc.config
	if !config.EnableDistributed {
		t.Error("Expected distributed cache to be enabled by default")
	}
	if config.ReplicationFactor <= 0 {
		t.Error("Expected positive replication factor")
	}
	if config.PartitionCount <= 0 {
		t.Error("Expected positive partition count")
	}
}

func TestDistributedCache_SetAndGet(t *testing.T) {
	dc, err := NewDistributedCache(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ctx := context.Background()
	key := "test_key"
	value := "test_value"
	ttl := time.Hour

	// Test Set
	err = dc.Set(ctx, key, value, ttl)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test Get
	retrievedValue, exists, err := dc.Get(ctx, key)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !exists {
		t.Error("Expected key to exist")
	}

	if retrievedValue != value {
		t.Errorf("Expected value %v, got %v", value, retrievedValue)
	}
}

func TestDistributedCache_Delete(t *testing.T) {
	dc, err := NewDistributedCache(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ctx := context.Background()
	key := "test_key"
	value := "test_value"
	ttl := time.Hour

	// Set value
	err = dc.Set(ctx, key, value, ttl)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Delete value
	err = dc.Delete(ctx, key)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify deletion
	_, exists, err := dc.Get(ctx, key)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if exists {
		t.Error("Expected key to be deleted")
	}
}

func TestDistributedCache_GetStats(t *testing.T) {
	dc, err := NewDistributedCache(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	stats := dc.GetStats()
	if stats == nil {
		t.Fatal("Expected non-nil stats")
	}

	// Check that stats have been initialized
	if stats.TotalOperations < 0 {
		t.Error("Expected non-negative total operations")
	}
}

func TestDistributedCache_GetClusterStatus(t *testing.T) {
	dc, err := NewDistributedCache(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	status := dc.GetClusterStatus()
	if status == nil {
		t.Fatal("Expected non-nil cluster status")
	}

	// Check required fields
	requiredFields := []string{"nodes", "leader", "shards", "replication", "failover", "statistics"}
	for _, field := range requiredFields {
		if _, exists := status[field]; !exists {
			t.Errorf("Missing required field: %s", field)
		}
	}
}

func TestEventDrivenSystem_New(t *testing.T) {
	config := &EventDrivenConfig{
		EnableEventSourcing: true,
		EnableCQRS:          true,
		EnableEventBus:      true,
	}

	eds, err := NewEventDrivenSystem(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if eds == nil {
		t.Fatal("Expected non-nil event driven system")
	}

	if !eds.config.EnableEventSourcing {
		t.Error("Expected event sourcing to be enabled")
	}
}

func TestEventDrivenSystem_DefaultConfig(t *testing.T) {
	eds, err := NewEventDrivenSystem(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if eds == nil {
		t.Fatal("Expected non-nil event driven system")
	}

	config := eds.config
	if !config.EnableEventSourcing {
		t.Error("Expected event sourcing to be enabled by default")
	}
	if !config.EnableCQRS {
		t.Error("Expected CQRS to be enabled by default")
	}
	if !config.EnableEventBus {
		t.Error("Expected event bus to be enabled by default")
	}
}

func TestEventDrivenSystem_PublishEvent(t *testing.T) {
	eds, err := NewEventDrivenSystem(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ctx := context.Background()
	event := &Event{
		ID:            "test_event_1",
		Type:          "test_event",
		AggregateID:   "test_aggregate_1",
		AggregateType: "test_aggregate",
		Version:       1,
		Data:          map[string]interface{}{"key": "value"},
		Metadata:      map[string]interface{}{"source": "test"},
		Timestamp:     time.Now(),
		CorrelationID: "test_correlation_1",
		CausationID:   "test_causation_1",
	}

	err = eds.PublishEvent(ctx, event)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check that event was recorded
	stats := eds.GetStats()
	if stats.TotalEvents == 0 {
		t.Error("Expected total events to be greater than 0")
	}
}

func TestEventDrivenSystem_SendCommand(t *testing.T) {
	eds, err := NewEventDrivenSystem(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Register a test command handler
	handler := &TestCommandHandler{}
	eds.RegisterCommandHandler("test_command", handler)

	ctx := context.Background()
	command := &Command{
		ID:            "test_command_1",
		Type:          "test_command",
		AggregateID:   "test_aggregate_1",
		AggregateType: "test_aggregate",
		Data:          map[string]interface{}{"key": "value"},
		Metadata:      map[string]interface{}{"source": "test"},
		Timestamp:     time.Now(),
		CorrelationID: "test_correlation_1",
		CausationID:   "test_causation_1",
	}

	event, err := eds.SendCommand(ctx, command)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if event == nil {
		t.Error("Expected non-nil event")
	}

	// Check that command was processed
	stats := eds.GetStats()
	if stats.TotalCommands == 0 {
		t.Error("Expected total commands to be greater than 0")
	}
}

func TestEventDrivenSystem_SendQuery(t *testing.T) {
	eds, err := NewEventDrivenSystem(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Register a test query handler
	handler := &TestQueryHandler{}
	eds.RegisterQueryHandler("test_query", handler)

	ctx := context.Background()
	query := &Query{
		ID:            "test_query_1",
		Type:          "test_query",
		Data:          map[string]interface{}{"key": "value"},
		Metadata:      map[string]interface{}{"source": "test"},
		Timestamp:     time.Now(),
		CorrelationID: "test_correlation_1",
	}

	result, err := eds.SendQuery(ctx, query)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Error("Expected non-nil result")
	}

	// Check that query was processed
	stats := eds.GetStats()
	if stats.TotalQueries == 0 {
		t.Error("Expected total queries to be greater than 0")
	}
}

func TestRealTimeAnalytics_New(t *testing.T) {
	config := &RealTimeAnalyticsConfig{
		EnableStreaming:  true,
		EnableDashboards: true,
		EnableAlerts:     true,
		EnableML:         true,
	}

	rta, err := NewRealTimeAnalytics(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if rta == nil {
		t.Fatal("Expected non-nil real-time analytics")
	}

	if !rta.config.EnableStreaming {
		t.Error("Expected streaming to be enabled")
	}
}

func TestRealTimeAnalytics_DefaultConfig(t *testing.T) {
	rta, err := NewRealTimeAnalytics(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if rta == nil {
		t.Fatal("Expected non-nil real-time analytics")
	}

	config := rta.config
	if !config.EnableStreaming {
		t.Error("Expected streaming to be enabled by default")
	}
	if !config.EnableDashboards {
		t.Error("Expected dashboards to be enabled by default")
	}
	if !config.EnableAlerts {
		t.Error("Expected alerts to be enabled by default")
	}
	if !config.EnableML {
		t.Error("Expected ML to be enabled by default")
	}
}

func TestRealTimeAnalytics_RecordMetric(t *testing.T) {
	rta, err := NewRealTimeAnalytics(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ctx := context.Background()
	metric := &Metric{
		Name:        "test_metric",
		Value:       42.0,
		Timestamp:   time.Now(),
		Tags:        map[string]string{"test": "value"},
		Type:        AnalyticsMetricTypeGauge,
		Unit:        "count",
		Description: "Test metric",
	}

	err = rta.RecordMetric(ctx, metric)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check that metric was recorded
	stats := rta.GetStats()
	if stats.TotalMetrics == 0 {
		t.Error("Expected total metrics to be greater than 0")
	}
}

func TestRealTimeAnalytics_CreateDashboard(t *testing.T) {
	rta, err := NewRealTimeAnalytics(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ctx := context.Background()
	dashboard := &Dashboard{
		ID:              "test_dashboard_1",
		Name:            "Test Dashboard",
		Description:     "A test dashboard",
		Widgets:         []*Widget{},
		Layout:          &Layout{Rows: []*Row{}, Columns: 12, Spacing: 16},
		RefreshInterval: time.Second * 30,
		LastUpdated:     time.Now(),
		Public:          false,
		Tags:            []string{"test"},
	}

	err = rta.CreateDashboard(ctx, dashboard)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify dashboard was created
	retrievedDashboard, err := rta.GetDashboard(ctx, dashboard.ID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if retrievedDashboard == nil {
		t.Fatal("Expected non-nil dashboard")
	}

	if retrievedDashboard.ID != dashboard.ID {
		t.Errorf("Expected dashboard ID %s, got %s", dashboard.ID, retrievedDashboard.ID)
	}
}

func TestRealTimeAnalytics_CreateAlertRule(t *testing.T) {
	rta, err := NewRealTimeAnalytics(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ctx := context.Background()
	rule := &AlertRule{
		ID:                   "test_rule_1",
		Name:                 "Test Alert Rule",
		Description:          "A test alert rule",
		Metric:               "test_metric",
		Condition:            ">",
		Threshold:            100.0,
		Severity:             AlertSeverityWarning,
		Enabled:              true,
		CooldownPeriod:       time.Minute * 5,
		NotificationChannels: []string{"email", "slack"},
	}

	err = rta.CreateAlertRule(ctx, rule)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify alert rule was created
	alerts, err := rta.GetAlerts(ctx, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check that alerts were created (simplified test)
	if alerts == nil {
		t.Error("Expected non-nil alerts")
	}
}

func TestRealTimeAnalytics_GetMLPrediction(t *testing.T) {
	rta, err := NewRealTimeAnalytics(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ctx := context.Background()
	modelName := "test_model"
	input := map[string]interface{}{"feature1": 1.0, "feature2": 2.0}

	// This will fail because no model is registered, but we can test the interface
	_, err = rta.GetMLPrediction(ctx, modelName, input)
	if err == nil {
		t.Error("Expected error for non-existent model")
	}
}

func TestRealTimeAnalytics_GetABTestResult(t *testing.T) {
	rta, err := NewRealTimeAnalytics(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ctx := context.Background()
	testID := "test_ab_test_1"

	result, err := rta.GetABTestResult(ctx, testID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Error("Expected non-nil test result")
	}
}

// Test helpers

type TestCommandHandler struct{}

func (h *TestCommandHandler) Handle(ctx context.Context, command *Command) (*Event, error) {
	return &Event{
		ID:            "test_event_from_command",
		Type:          "test_event",
		AggregateID:   command.AggregateID,
		AggregateType: command.AggregateType,
		Version:       1,
		Data:          command.Data,
		Metadata:      command.Metadata,
		Timestamp:     time.Now(),
		CorrelationID: command.CorrelationID,
		CausationID:   command.CausationID,
	}, nil
}

type TestQueryHandler struct{}

func (h *TestQueryHandler) Handle(ctx context.Context, query *Query) (interface{}, error) {
	return map[string]interface{}{
		"result": "test_query_result",
		"data":   query.Data,
	}, nil
}

func BenchmarkDistributedCache_SetAndGet(b *testing.B) {
	dc, err := NewDistributedCache(nil)
	if err != nil {
		b.Fatalf("Expected no error, got %v", err)
	}

	ctx := context.Background()
	key := "benchmark_key"
	value := "benchmark_value"
	ttl := time.Hour

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dc.Set(ctx, key, value, ttl)
		dc.Get(ctx, key)
	}
}

func BenchmarkEventDrivenSystem_PublishEvent(b *testing.B) {
	eds, err := NewEventDrivenSystem(nil)
	if err != nil {
		b.Fatalf("Expected no error, got %v", err)
	}

	ctx := context.Background()
	event := &Event{
		ID:            "benchmark_event",
		Type:          "benchmark_event",
		AggregateID:   "benchmark_aggregate",
		AggregateType: "benchmark_aggregate",
		Version:       1,
		Data:          map[string]interface{}{"key": "value"},
		Metadata:      map[string]interface{}{"source": "benchmark"},
		Timestamp:     time.Now(),
		CorrelationID: "benchmark_correlation",
		CausationID:   "benchmark_causation",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		eds.PublishEvent(ctx, event)
	}
}

func BenchmarkRealTimeAnalytics_RecordMetric(b *testing.B) {
	rta, err := NewRealTimeAnalytics(nil)
	if err != nil {
		b.Fatalf("Expected no error, got %v", err)
	}

	ctx := context.Background()
	metric := &Metric{
		Name:        "benchmark_metric",
		Value:       42.0,
		Timestamp:   time.Now(),
		Tags:        map[string]string{"test": "benchmark"},
		Type:        AnalyticsMetricTypeGauge,
		Unit:        "count",
		Description: "Benchmark metric",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rta.RecordMetric(ctx, metric)
	}
}
