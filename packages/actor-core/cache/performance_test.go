package cache

import (
	"testing"
	"time"
)

func TestPerformanceTest(t *testing.T) {
	config := &BenchmarkConfig{
		KeyCount:          1000,
		ValueSize:         1024,
		Duration:          time.Second * 10,
		Concurrency:       10,
		L1MaxSize:         10000,
		L2MaxSize:         100000,
		L3MaxSize:         1000000,
		TargetLatency:     time.Microsecond * 100,
		TargetThroughput:  10000,
		TargetMemoryUsage: 100 * 1024 * 1024, // 100MB
	}

	results, err := RunPerformanceTest(config)
	if err != nil {
		t.Fatalf("Performance test failed: %v", err)
	}

	// Print results
	PrintBenchmarkResults(results)

	// Validate performance targets
	if results.Latency > config.TargetLatency {
		t.Errorf("Latency %v exceeds target %v", results.Latency, config.TargetLatency)
	}

	if results.Throughput < float64(config.TargetThroughput) {
		t.Errorf("Throughput %.2f below target %d", results.Throughput, config.TargetThroughput)
	}

	if results.MemoryUsage > config.TargetMemoryUsage {
		t.Errorf("Memory usage %d exceeds target %d", results.MemoryUsage, config.TargetMemoryUsage)
	}

	// Validate cache performance
	if results.HitRate < 50.0 {
		t.Errorf("Hit rate %.2f%% is too low", results.HitRate)
	}

	if results.GoroutineCount > 100 {
		t.Errorf("Too many goroutines: %d", results.GoroutineCount)
	}
}

func TestPerformanceTestHighLoad(t *testing.T) {
	config := &BenchmarkConfig{
		KeyCount:          10000,
		ValueSize:         4096,
		Duration:          time.Second * 30,
		Concurrency:       100,
		L1MaxSize:         100000,
		L2MaxSize:         1000000,
		L3MaxSize:         10000000,
		TargetLatency:     time.Millisecond * 1,
		TargetThroughput:  5000,
		TargetMemoryUsage: 500 * 1024 * 1024, // 500MB
	}

	results, err := RunPerformanceTest(config)
	if err != nil {
		t.Fatalf("High load performance test failed: %v", err)
	}

	// Print results
	PrintBenchmarkResults(results)

	// Validate performance targets
	if results.Latency > config.TargetLatency {
		t.Errorf("Latency %v exceeds target %v", results.Latency, config.TargetLatency)
	}

	if results.Throughput < float64(config.TargetThroughput) {
		t.Errorf("Throughput %.2f below target %d", results.Throughput, config.TargetThroughput)
	}

	if results.MemoryUsage > config.TargetMemoryUsage {
		t.Errorf("Memory usage %d exceeds target %d", results.MemoryUsage, config.TargetMemoryUsage)
	}
}

func TestPerformanceTestMemoryPressure(t *testing.T) {
	config := &BenchmarkConfig{
		KeyCount:          50000,
		ValueSize:         8192,
		Duration:          time.Second * 60,
		Concurrency:       50,
		L1MaxSize:         1000,  // Small L1 to force evictions
		L2MaxSize:         10000, // Small L2 to force evictions
		L3MaxSize:         100000,
		TargetLatency:     time.Millisecond * 5,
		TargetThroughput:  1000,
		TargetMemoryUsage: 1000 * 1024 * 1024, // 1GB
	}

	results, err := RunPerformanceTest(config)
	if err != nil {
		t.Fatalf("Memory pressure performance test failed: %v", err)
	}

	// Print results
	PrintBenchmarkResults(results)

	// Validate performance targets
	if results.Latency > config.TargetLatency {
		t.Errorf("Latency %v exceeds target %v", results.Latency, config.TargetLatency)
	}

	if results.Throughput < float64(config.TargetThroughput) {
		t.Errorf("Throughput %.2f below target %d", results.Throughput, config.TargetThroughput)
	}

	if results.MemoryUsage > config.TargetMemoryUsage {
		t.Errorf("Memory usage %d exceeds target %d", results.MemoryUsage, config.TargetMemoryUsage)
	}

	// Validate eviction behavior
	if results.EvictionRate < 10.0 {
		t.Errorf("Eviction rate %.2f%% is too low for memory pressure test", results.EvictionRate)
	}
}

func TestPerformanceTestConcurrentAccess(t *testing.T) {
	config := &BenchmarkConfig{
		KeyCount:          1000,
		ValueSize:         1024,
		Duration:          time.Second * 20,
		Concurrency:       200, // High concurrency
		L1MaxSize:         10000,
		L2MaxSize:         100000,
		L3MaxSize:         1000000,
		TargetLatency:     time.Millisecond * 2,
		TargetThroughput:  2000,
		TargetMemoryUsage: 200 * 1024 * 1024, // 200MB
	}

	results, err := RunPerformanceTest(config)
	if err != nil {
		t.Fatalf("Concurrent access performance test failed: %v", err)
	}

	// Print results
	PrintBenchmarkResults(results)

	// Validate performance targets
	if results.Latency > config.TargetLatency {
		t.Errorf("Latency %v exceeds target %v", results.Latency, config.TargetLatency)
	}

	if results.Throughput < float64(config.TargetThroughput) {
		t.Errorf("Throughput %.2f below target %d", results.Throughput, config.TargetThroughput)
	}

	if results.MemoryUsage > config.TargetMemoryUsage {
		t.Errorf("Memory usage %d exceeds target %d", results.MemoryUsage, config.TargetMemoryUsage)
	}

	// Validate concurrency handling
	if results.GoroutineCount > 500 {
		t.Errorf("Too many goroutines: %d", results.GoroutineCount)
	}
}

func TestPerformanceTestWarmingEffectiveness(t *testing.T) {
	config := &BenchmarkConfig{
		KeyCount:          5000,
		ValueSize:         2048,
		Duration:          time.Second * 40,
		Concurrency:       20,
		L1MaxSize:         1000, // Small L1 to test warming
		L2MaxSize:         5000, // Small L2 to test warming
		L3MaxSize:         50000,
		TargetLatency:     time.Millisecond * 3,
		TargetThroughput:  1500,
		TargetMemoryUsage: 300 * 1024 * 1024, // 300MB
	}

	results, err := RunPerformanceTest(config)
	if err != nil {
		t.Fatalf("Warming effectiveness performance test failed: %v", err)
	}

	// Print results
	PrintBenchmarkResults(results)

	// Validate performance targets
	if results.Latency > config.TargetLatency {
		t.Errorf("Latency %v exceeds target %v", results.Latency, config.TargetLatency)
	}

	if results.Throughput < float64(config.TargetThroughput) {
		t.Errorf("Throughput %.2f below target %d", results.Throughput, config.TargetThroughput)
	}

	if results.MemoryUsage > config.TargetMemoryUsage {
		t.Errorf("Memory usage %d exceeds target %d", results.MemoryUsage, config.TargetMemoryUsage)
	}

	// Validate warming effectiveness
	if results.HitRate < 30.0 {
		t.Errorf("Hit rate %.2f%% is too low for warming test", results.HitRate)
	}
}

func TestPerformanceTestInvalidationOverhead(t *testing.T) {
	config := &BenchmarkConfig{
		KeyCount:          2000,
		ValueSize:         1024,
		Duration:          time.Second * 30,
		Concurrency:       30,
		L1MaxSize:         5000,
		L2MaxSize:         20000,
		L3MaxSize:         100000,
		TargetLatency:     time.Millisecond * 2,
		TargetThroughput:  3000,
		TargetMemoryUsage: 150 * 1024 * 1024, // 150MB
	}

	results, err := RunPerformanceTest(config)
	if err != nil {
		t.Fatalf("Invalidation overhead performance test failed: %v", err)
	}

	// Print results
	PrintBenchmarkResults(results)

	// Validate performance targets
	if results.Latency > config.TargetLatency {
		t.Errorf("Latency %v exceeds target %v", results.Latency, config.TargetLatency)
	}

	if results.Throughput < float64(config.TargetThroughput) {
		t.Errorf("Throughput %.2f below target %d", results.Throughput, config.TargetThroughput)
	}

	if results.MemoryUsage > config.TargetMemoryUsage {
		t.Errorf("Memory usage %d exceeds target %d", results.MemoryUsage, config.TargetMemoryUsage)
	}
}

func TestPerformanceTestStress(t *testing.T) {
	config := &BenchmarkConfig{
		KeyCount:          100000,
		ValueSize:         16384,
		Duration:          time.Minute * 2,
		Concurrency:       500,
		L1MaxSize:         10000,
		L2MaxSize:         100000,
		L3MaxSize:         1000000,
		TargetLatency:     time.Millisecond * 10,
		TargetThroughput:  500,
		TargetMemoryUsage: 2000 * 1024 * 1024, // 2GB
	}

	results, err := RunPerformanceTest(config)
	if err != nil {
		t.Fatalf("Stress performance test failed: %v", err)
	}

	// Print results
	PrintBenchmarkResults(results)

	// Validate performance targets
	if results.Latency > config.TargetLatency {
		t.Errorf("Latency %v exceeds target %v", results.Latency, config.TargetLatency)
	}

	if results.Throughput < float64(config.TargetThroughput) {
		t.Errorf("Throughput %.2f below target %d", results.Throughput, config.TargetThroughput)
	}

	if results.MemoryUsage > config.TargetMemoryUsage {
		t.Errorf("Memory usage %d exceeds target %d", results.MemoryUsage, config.TargetMemoryUsage)
	}

	// Validate system stability
	if results.GoroutineCount > 1000 {
		t.Errorf("Too many goroutines: %d", results.GoroutineCount)
	}

	if results.GCStats.NumGC > 1000 {
		t.Errorf("Too many GC cycles: %d", results.GCStats.NumGC)
	}
}
