package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TestAdvancedFeaturesIntegration tests the integration of all advanced features
func TestAdvancedFeaturesIntegration(t *testing.T) {
	// Create all advanced feature systems
	distributedConfig := DefaultDistributedCacheConfig()
	distributedCache, err := NewDistributedCache(distributedConfig)
	if err != nil {
		t.Fatalf("Failed to create distributed cache: %v", err)
	}

	eventDrivenConfig := DefaultEventDrivenConfig()
	eventDrivenSystem, err := NewEventDrivenSystem(eventDrivenConfig)
	if err != nil {
		t.Fatalf("Failed to create event driven system: %v", err)
	}

	analyticsConfig := DefaultRealTimeAnalyticsConfig()
	analyticsSystem, err := NewRealTimeAnalytics(analyticsConfig)
	if err != nil {
		t.Fatalf("Failed to create analytics system: %v", err)
	}

	// Test data
	testData := map[string]interface{}{
		"user_id":    "user_123",
		"action":     "login",
		"timestamp":  time.Now(),
		"ip_address": "192.168.1.1",
		"user_agent": "Mozilla/5.0",
	}

	ctx := context.Background()

	// Test distributed caching
	t.Run("Distributed_Caching", func(t *testing.T) {
		key := "user_session_123"
		value := testData
		ttl := time.Hour

		// Test Set
		err := distributedCache.Set(ctx, key, value, ttl)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Test Get
		retrievedValue, exists, err := distributedCache.Get(ctx, key)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if !exists {
			t.Error("Expected key to exist")
		}

		if retrievedValue == nil {
			t.Error("Expected non-nil value")
		}

		// Test cluster status
		status := distributedCache.GetClusterStatus()
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
	})

	// Test event-driven architecture
	t.Run("Event_Driven_Architecture", func(t *testing.T) {
		// Test event publishing
		event := &Event{
			ID:            "user_login_event_1",
			Type:          "user_login",
			AggregateID:   "user_123",
			AggregateType: "user",
			Version:       1,
			Data:          testData,
			Metadata:      map[string]interface{}{"source": "web", "version": "1.0.0"},
			Timestamp:     time.Now(),
			CorrelationID: "correlation_123",
			CausationID:   "causation_123",
		}

		err := eventDrivenSystem.PublishEvent(ctx, event)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Test command processing
		commandHandler := &TestUserCommandHandler{}
		eventDrivenSystem.RegisterCommandHandler("create_user", commandHandler)

		command := &Command{
			ID:            "create_user_command_1",
			Type:          "create_user",
			AggregateID:   "user_123",
			AggregateType: "user",
			Data:          testData,
			Metadata:      map[string]interface{}{"source": "api"},
			Timestamp:     time.Now(),
			CorrelationID: "correlation_123",
			CausationID:   "causation_123",
		}

		event, err = eventDrivenSystem.SendCommand(ctx, command)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if event == nil {
			t.Error("Expected non-nil event")
		}

		// Test query processing
		queryHandler := &TestUserQueryHandler{}
		eventDrivenSystem.RegisterQueryHandler("get_user", queryHandler)

		query := &Query{
			ID:            "get_user_query_1",
			Type:          "get_user",
			Data:          map[string]interface{}{"user_id": "user_123"},
			Metadata:      map[string]interface{}{"source": "api"},
			Timestamp:     time.Now(),
			CorrelationID: "correlation_123",
		}

		result, err := eventDrivenSystem.SendQuery(ctx, query)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if result == nil {
			t.Error("Expected non-nil result")
		}

		// Test event retrieval
		events, err := eventDrivenSystem.GetEvents(ctx, "user_123", 0)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Events might be empty in simplified implementation
		if events == nil {
			t.Error("Expected non-nil events")
		}
	})

	// Test real-time analytics
	t.Run("Real_Time_Analytics", func(t *testing.T) {
		// Test metric recording
		metric := &Metric{
			Name:        "user_login_count",
			Value:       1.0,
			Timestamp:   time.Now(),
			Tags:        map[string]string{"user_id": "user_123", "action": "login"},
			Type:        AnalyticsMetricTypeCounter,
			Unit:        "count",
			Description: "Number of user logins",
		}

		err := analyticsSystem.RecordMetric(ctx, metric)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Test dashboard creation
		dashboard := &Dashboard{
			ID:          "user_analytics_dashboard",
			Name:        "User Analytics Dashboard",
			Description: "Dashboard for user analytics",
			Widgets: []*Widget{
				{
					ID:              "login_count_widget",
					Type:            WidgetTypeCounter,
					Title:           "Login Count",
					Description:     "Total number of logins",
					Config:          map[string]interface{}{"metric": "user_login_count"},
					Data:            nil,
					LastUpdated:     time.Now(),
					RefreshInterval: time.Second * 30,
				},
			},
			Layout:          &Layout{Rows: []*Row{{Widgets: []*Widget{}, Height: 200}}, Columns: 12, Spacing: 16},
			RefreshInterval: time.Second * 30,
			LastUpdated:     time.Now(),
			Public:          false,
			Tags:            []string{"user", "analytics"},
		}

		err = analyticsSystem.CreateDashboard(ctx, dashboard)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Verify dashboard was created
		retrievedDashboard, err := analyticsSystem.GetDashboard(ctx, dashboard.ID)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if retrievedDashboard == nil {
			t.Fatal("Expected non-nil dashboard")
		}

		if retrievedDashboard.ID != dashboard.ID {
			t.Errorf("Expected dashboard ID %s, got %s", dashboard.ID, retrievedDashboard.ID)
		}

		// Test alert rule creation
		alertRule := &AlertRule{
			ID:                   "high_login_rate_alert",
			Name:                 "High Login Rate Alert",
			Description:          "Alert when login rate is too high",
			Metric:               "user_login_count",
			Condition:            ">",
			Threshold:            100.0,
			Severity:             AlertSeverityWarning,
			Enabled:              true,
			CooldownPeriod:       time.Minute * 5,
			NotificationChannels: []string{"email", "slack"},
		}

		err = analyticsSystem.CreateAlertRule(ctx, alertRule)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Test ML prediction
		modelName := "user_behavior_model"
		input := map[string]interface{}{
			"user_id":    "user_123",
			"login_time": time.Now(),
			"ip_address": "192.168.1.1",
			"user_agent": "Mozilla/5.0",
		}

		// This will fail because no model is registered, but we can test the interface
		_, err = analyticsSystem.GetMLPrediction(ctx, modelName, input)
		if err == nil {
			t.Error("Expected error for non-existent model")
		}

		// Test A/B test result
		testID := "user_interface_test"
		result, err := analyticsSystem.GetABTestResult(ctx, testID)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if result == nil {
			t.Error("Expected non-nil test result")
		}
	})

	// Test combined workflow
	t.Run("Combined_Workflow", func(t *testing.T) {
		// Simulate a complete user workflow
		userID := "user_456"

		// 1. User login event
		loginEvent := &Event{
			ID:            "user_login_event_2",
			Type:          "user_login",
			AggregateID:   userID,
			AggregateType: "user",
			Version:       1,
			Data:          map[string]interface{}{"user_id": userID, "action": "login", "timestamp": time.Now()},
			Metadata:      map[string]interface{}{"source": "mobile", "version": "2.0.0"},
			Timestamp:     time.Now(),
			CorrelationID: "correlation_456",
			CausationID:   "causation_456",
		}

		err := eventDrivenSystem.PublishEvent(ctx, loginEvent)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// 2. Cache user session
		sessionKey := fmt.Sprintf("user_session_%s", userID)
		sessionData := map[string]interface{}{
			"user_id":    userID,
			"login_time": time.Now(),
			"session_id": "session_123",
		}

		err = distributedCache.Set(ctx, sessionKey, sessionData, time.Hour)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// 3. Record analytics metrics
		loginMetric := &Metric{
			Name:        "user_login_count",
			Value:       1.0,
			Timestamp:   time.Now(),
			Tags:        map[string]string{"user_id": userID, "action": "login", "source": "mobile"},
			Type:        AnalyticsMetricTypeCounter,
			Unit:        "count",
			Description: "Number of user logins",
		}

		err = analyticsSystem.RecordMetric(ctx, loginMetric)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// 4. Verify session was cached
		cachedSession, exists, err := distributedCache.Get(ctx, sessionKey)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if !exists {
			t.Error("Expected session to be cached")
		}

		if cachedSession == nil {
			t.Error("Expected non-nil cached session")
		}

		// 5. Check analytics stats
		analyticsStats := analyticsSystem.GetStats()
		if analyticsStats.TotalMetrics == 0 {
			t.Error("Expected total metrics to be greater than 0")
		}

		// 6. Check event-driven stats
		eventStats := eventDrivenSystem.GetStats()
		if eventStats.TotalEvents == 0 {
			t.Error("Expected total events to be greater than 0")
		}

		// 7. Check distributed cache stats
		distributedStats := distributedCache.GetStats()
		if distributedStats.TotalOperations == 0 {
			t.Error("Expected total operations to be greater than 0")
		}
	})
}

// TestAdvancedFeaturesPerformance tests the performance of advanced features
func TestAdvancedFeaturesPerformance(t *testing.T) {
	// Create all systems
	distributedCache, err := NewDistributedCache(nil)
	if err != nil {
		t.Fatalf("Failed to create distributed cache: %v", err)
	}

	eventDrivenSystem, err := NewEventDrivenSystem(nil)
	if err != nil {
		t.Fatalf("Failed to create event driven system: %v", err)
	}

	analyticsSystem, err := NewRealTimeAnalytics(nil)
	if err != nil {
		t.Fatalf("Failed to create analytics system: %v", err)
	}

	ctx := context.Background()

	// Performance test parameters
	numOperations := 1000
	numUsers := 100

	// Test distributed cache performance
	t.Run("Distributed_Cache_Performance", func(t *testing.T) {
		start := time.Now()

		for i := 0; i < numOperations; i++ {
			key := fmt.Sprintf("perf_key_%d", i)
			value := map[string]interface{}{"index": i, "timestamp": time.Now()}
			ttl := time.Hour

			err := distributedCache.Set(ctx, key, value, ttl)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			_, _, err = distributedCache.Get(ctx, key)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
		}

		duration := time.Since(start)
		t.Logf("Distributed cache operations completed in %v", duration)

		// Verify performance is reasonable (less than 1 second for 1000 operations)
		if duration > time.Second {
			t.Errorf("Performance too slow: %v", duration)
		}
	})

	// Test event-driven system performance
	t.Run("Event_Driven_System_Performance", func(t *testing.T) {
		start := time.Now()

		for i := 0; i < numOperations; i++ {
			event := &Event{
				ID:            fmt.Sprintf("perf_event_%d", i),
				Type:          "performance_test",
				AggregateID:   fmt.Sprintf("user_%d", i%numUsers),
				AggregateType: "user",
				Version:       1,
				Data:          map[string]interface{}{"index": i, "timestamp": time.Now()},
				Metadata:      map[string]interface{}{"source": "performance_test"},
				Timestamp:     time.Now(),
				CorrelationID: fmt.Sprintf("correlation_%d", i),
				CausationID:   fmt.Sprintf("causation_%d", i),
			}

			err := eventDrivenSystem.PublishEvent(ctx, event)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
		}

		duration := time.Since(start)
		t.Logf("Event-driven system operations completed in %v", duration)

		// Verify performance is reasonable (less than 1 second for 1000 operations)
		if duration > time.Second {
			t.Errorf("Performance too slow: %v", duration)
		}
	})

	// Test analytics system performance
	t.Run("Analytics_System_Performance", func(t *testing.T) {
		start := time.Now()

		for i := 0; i < numOperations; i++ {
			metric := &Metric{
				Name:        "performance_metric",
				Value:       float64(i),
				Timestamp:   time.Now(),
				Tags:        map[string]string{"index": fmt.Sprintf("%d", i), "user_id": fmt.Sprintf("user_%d", i%numUsers)},
				Type:        AnalyticsMetricTypeCounter,
				Unit:        "count",
				Description: "Performance test metric",
			}

			err := analyticsSystem.RecordMetric(ctx, metric)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
		}

		duration := time.Since(start)
		t.Logf("Analytics system operations completed in %v", duration)

		// Verify performance is reasonable (less than 1 second for 1000 operations)
		if duration > time.Second {
			t.Errorf("Performance too slow: %v", duration)
		}
	})

	// Test combined performance
	t.Run("Combined_Performance", func(t *testing.T) {
		start := time.Now()

		for i := 0; i < numOperations; i++ {
			userID := fmt.Sprintf("user_%d", i%numUsers)

			// 1. Cache user data
			key := fmt.Sprintf("user_data_%s", userID)
			value := map[string]interface{}{"user_id": userID, "index": i, "timestamp": time.Now()}
			ttl := time.Hour

			err := distributedCache.Set(ctx, key, value, ttl)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			// 2. Publish event
			event := &Event{
				ID:            fmt.Sprintf("combined_event_%d", i),
				Type:          "user_action",
				AggregateID:   userID,
				AggregateType: "user",
				Version:       1,
				Data:          value,
				Metadata:      map[string]interface{}{"source": "combined_test"},
				Timestamp:     time.Now(),
				CorrelationID: fmt.Sprintf("combined_correlation_%d", i),
				CausationID:   fmt.Sprintf("combined_causation_%d", i),
			}

			err = eventDrivenSystem.PublishEvent(ctx, event)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			// 3. Record metric
			metric := &Metric{
				Name:        "combined_metric",
				Value:       float64(i),
				Timestamp:   time.Now(),
				Tags:        map[string]string{"user_id": userID, "index": fmt.Sprintf("%d", i)},
				Type:        AnalyticsMetricTypeCounter,
				Unit:        "count",
				Description: "Combined performance test metric",
			}

			err = analyticsSystem.RecordMetric(ctx, metric)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
		}

		duration := time.Since(start)
		t.Logf("Combined operations completed in %v", duration)

		// Verify performance is reasonable (less than 2 seconds for 1000 operations)
		if duration > 2*time.Second {
			t.Errorf("Performance too slow: %v", duration)
		}
	})
}

// Test helpers

type TestUserCommandHandler struct{}

func (h *TestUserCommandHandler) Handle(ctx context.Context, command *Command) (*Event, error) {
	return &Event{
		ID:            "user_created_event",
		Type:          "user_created",
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

type TestUserQueryHandler struct{}

func (h *TestUserQueryHandler) Handle(ctx context.Context, query *Query) (interface{}, error) {
	return map[string]interface{}{
		"user_id":    query.Data["user_id"],
		"name":       "Test User",
		"email":      "test@example.com",
		"created_at": time.Now(),
	}, nil
}

// BenchmarkAdvancedFeatures benchmarks the advanced features
func BenchmarkAdvancedFeatures(b *testing.B) {
	// Create all systems
	distributedCache, err := NewDistributedCache(nil)
	if err != nil {
		b.Fatalf("Failed to create distributed cache: %v", err)
	}

	eventDrivenSystem, err := NewEventDrivenSystem(nil)
	if err != nil {
		b.Fatalf("Failed to create event driven system: %v", err)
	}

	analyticsSystem, err := NewRealTimeAnalytics(nil)
	if err != nil {
		b.Fatalf("Failed to create analytics system: %v", err)
	}

	ctx := context.Background()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		userID := fmt.Sprintf("user_%d", i%100)

		// 1. Cache user data
		key := fmt.Sprintf("user_data_%s", userID)
		value := map[string]interface{}{"user_id": userID, "index": i, "timestamp": time.Now()}
		ttl := time.Hour

		distributedCache.Set(ctx, key, value, ttl)

		// 2. Publish event
		event := &Event{
			ID:            fmt.Sprintf("benchmark_event_%d", i),
			Type:          "user_action",
			AggregateID:   userID,
			AggregateType: "user",
			Version:       1,
			Data:          value,
			Metadata:      map[string]interface{}{"source": "benchmark"},
			Timestamp:     time.Now(),
			CorrelationID: fmt.Sprintf("benchmark_correlation_%d", i),
			CausationID:   fmt.Sprintf("benchmark_causation_%d", i),
		}

		eventDrivenSystem.PublishEvent(ctx, event)

		// 3. Record metric
		metric := &Metric{
			Name:        "benchmark_metric",
			Value:       float64(i),
			Timestamp:   time.Now(),
			Tags:        map[string]string{"user_id": userID, "index": fmt.Sprintf("%d", i)},
			Type:        AnalyticsMetricTypeCounter,
			Unit:        "count",
			Description: "Benchmark test metric",
		}

		analyticsSystem.RecordMetric(ctx, metric)
	}
}
