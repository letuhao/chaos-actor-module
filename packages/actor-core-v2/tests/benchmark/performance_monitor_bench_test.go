package benchmark

import (
	"context"
	"testing"

	"actor-core/services/monitoring"
)

func BenchmarkPerformanceMonitor_SetMetric(b *testing.B) {
	pm := monitoring.NewPerformanceMonitor()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := pm.SetMetric("test_metric", float64(i), "ms", "performance", "Test metric")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPerformanceMonitor_GetMetric(b *testing.B) {
	pm := monitoring.NewPerformanceMonitor()

	// Pre-populate
	for i := 0; i < 1000; i++ {
		pm.SetMetric("test_metric", float64(i), "ms", "performance", "Test metric")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := pm.GetMetric("test_metric")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPerformanceMonitor_RecordMetric(b *testing.B) {
	pm := monitoring.NewPerformanceMonitor()
	tags := map[string]string{"service": "test", "version": "1.0"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pm.RecordMetric("test_metric", float64(i), tags)
	}
}

func BenchmarkPerformanceMonitor_RecordMetricWithContext(b *testing.B) {
	pm := monitoring.NewPerformanceMonitor()
	ctx := context.Background()
	tags := map[string]string{"service": "test", "version": "1.0"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pm.RecordMetricWithContext(ctx, "test_metric", float64(i), tags)
	}
}

func BenchmarkPerformanceMonitor_StartCalculation(b *testing.B) {
	pm := monitoring.NewPerformanceMonitor()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		timer := pm.StartCalculation("test_operation")
		pm.EndCalculation(timer)
	}
}

func BenchmarkPerformanceMonitor_IncrementCounter(b *testing.B) {
	pm := monitoring.NewPerformanceMonitor()
	tags := map[string]string{"service": "test"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pm.IncrementCounter("test_counter", tags)
	}
}

func BenchmarkPerformanceMonitor_SetGauge(b *testing.B) {
	pm := monitoring.NewPerformanceMonitor()
	tags := map[string]string{"service": "test"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pm.SetGauge("test_gauge", float64(i), tags)
	}
}

func BenchmarkPerformanceMonitor_ObserveHistogram(b *testing.B) {
	pm := monitoring.NewPerformanceMonitor()
	tags := map[string]string{"service": "test"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pm.ObserveHistogram("test_histogram", float64(i), tags)
	}
}

func BenchmarkPerformanceMonitor_ExportMetrics(b *testing.B) {
	pm := monitoring.NewPerformanceMonitor()

	// Pre-populate with metrics
	for i := 0; i < 1000; i++ {
		pm.SetMetric("test_metric", float64(i), "ms", "performance", "Test metric")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := pm.ExportMetrics("json")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPerformanceMonitor_ConcurrentAccess(b *testing.B) {
	pm := monitoring.NewPerformanceMonitor()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%4 == 0 {
				pm.SetMetric("test_metric", float64(i), "ms", "performance", "Test metric")
			} else if i%4 == 1 {
				pm.GetMetric("test_metric")
			} else if i%4 == 2 {
				pm.IncrementCounter("test_counter", map[string]string{"service": "test"})
			} else {
				pm.SetGauge("test_gauge", float64(i), map[string]string{"service": "test"})
			}
			i++
		}
	})
}
