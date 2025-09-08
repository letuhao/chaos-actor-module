package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// MonitoringConfig holds configuration for advanced monitoring
type MonitoringConfig struct {
	EnableMetrics     bool
	EnableProfiling   bool
	EnableTracing     bool
	EnableAlerts      bool
	MetricsInterval   time.Duration
	ProfilingInterval time.Duration
	TracingSampleRate float64
	AlertThresholds   map[string]float64
	RetentionPeriod   time.Duration
	MaxMetricsPoints  int
	EnableRealTime    bool
	EnableHistorical  bool
	EnablePredictive  bool
}

// DefaultMonitoringConfig returns default monitoring configuration
func DefaultMonitoringConfig() *MonitoringConfig {
	return &MonitoringConfig{
		EnableMetrics:     true,
		EnableProfiling:   true,
		EnableTracing:     true,
		EnableAlerts:      true,
		MetricsInterval:   time.Second * 5,
		ProfilingInterval: time.Minute,
		TracingSampleRate: 0.1, // 10% sampling
		AlertThresholds: map[string]float64{
			"cpu_usage":       80.0,
			"memory_usage":    85.0,
			"latency_p99":     100.0, // ms
			"error_rate":      5.0,   // %
			"cache_hit_rate":  70.0,  // %
			"throughput_drop": 50.0,  // %
		},
		RetentionPeriod:  time.Hour * 24,
		MaxMetricsPoints: 10000,
		EnableRealTime:   true,
		EnableHistorical: true,
		EnablePredictive: true,
	}
}

// MetricType represents different types of metrics
type MetricType string

const (
	MetricTypeCounter   MetricType = "counter"
	MetricTypeGauge     MetricType = "gauge"
	MetricTypeHistogram MetricType = "histogram"
	MetricTypeSummary   MetricType = "summary"
)

// MetricPoint represents a single metric data point
type MetricPoint struct {
	Timestamp time.Time              `json:"timestamp"`
	Name      string                 `json:"name"`
	Type      MetricType             `json:"type"`
	Value     float64                `json:"value"`
	Labels    map[string]string      `json:"labels"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// AlertLevel represents the severity of an alert
type AlertLevel string

const (
	AlertLevelInfo      AlertLevel = "info"
	AlertLevelWarning   AlertLevel = "warning"
	AlertLevelCritical  AlertLevel = "critical"
	AlertLevelEmergency AlertLevel = "emergency"
)

// Alert represents a monitoring alert
type Alert struct {
	ID         string                 `json:"id"`
	Timestamp  time.Time              `json:"timestamp"`
	Level      AlertLevel             `json:"level"`
	Message    string                 `json:"message"`
	Metric     string                 `json:"metric"`
	Value      float64                `json:"value"`
	Threshold  float64                `json:"threshold"`
	Labels     map[string]string      `json:"labels"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Resolved   bool                   `json:"resolved"`
	ResolvedAt *time.Time             `json:"resolved_at,omitempty"`
}

// ProfilingData represents profiling information
type ProfilingData struct {
	Timestamp     time.Time              `json:"timestamp"`
	CPUProfile    *CPUProfile            `json:"cpu_profile,omitempty"`
	MemoryProfile *MemoryProfile         `json:"memory_profile,omitempty"`
	GoroutineInfo *GoroutineInfo         `json:"goroutine_info,omitempty"`
	GCStats       *GCStats               `json:"gc_stats,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

// CPUProfile represents CPU profiling data
type CPUProfile struct {
	UsagePercent    float64 `json:"usage_percent"`
	UserTime        float64 `json:"user_time"`
	SystemTime      float64 `json:"system_time"`
	IdleTime        float64 `json:"idle_time"`
	IOWaitTime      float64 `json:"io_wait_time"`
	ContextSwitches int64   `json:"context_switches"`
	Interrupts      int64   `json:"interrupts"`
}

// MemoryProfile represents memory profiling data
type MemoryProfile struct {
	Alloc         uint64  `json:"alloc"`
	TotalAlloc    uint64  `json:"total_alloc"`
	Sys           uint64  `json:"sys"`
	Lookups       uint64  `json:"lookups"`
	Mallocs       uint64  `json:"mallocs"`
	Frees         uint64  `json:"frees"`
	HeapAlloc     uint64  `json:"heap_alloc"`
	HeapSys       uint64  `json:"heap_sys"`
	HeapIdle      uint64  `json:"heap_idle"`
	HeapInuse     uint64  `json:"heap_inuse"`
	HeapReleased  uint64  `json:"heap_released"`
	HeapObjects   uint64  `json:"heap_objects"`
	StackInuse    uint64  `json:"stack_inuse"`
	StackSys      uint64  `json:"stack_sys"`
	MSpanInuse    uint64  `json:"mspan_inuse"`
	MSpanSys      uint64  `json:"mspan_sys"`
	MCacheInuse   uint64  `json:"mcache_inuse"`
	MCacheSys     uint64  `json:"mcache_sys"`
	BuckHashSys   uint64  `json:"buck_hash_sys"`
	GCSys         uint64  `json:"gc_sys"`
	OtherSys      uint64  `json:"other_sys"`
	NextGC        uint64  `json:"next_gc"`
	LastGC        uint64  `json:"last_gc"`
	PauseTotalNs  uint64  `json:"pause_total_ns"`
	NumGC         int32   `json:"num_gc"`
	NumForcedGC   int32   `json:"num_forced_gc"`
	GCCPUFraction float64 `json:"gc_cpu_fraction"`
}

// GoroutineInfo represents goroutine information
type GoroutineInfo struct {
	Count         int     `json:"count"`
	BlockedCount  int     `json:"blocked_count"`
	RunnableCount int     `json:"runnable_count"`
	RunningCount  int     `json:"running_count"`
	WaitingCount  int     `json:"waiting_count"`
	AverageStack  float64 `json:"average_stack"`
	MaxStack      float64 `json:"max_stack"`
	MinStack      float64 `json:"min_stack"`
}

// GCStats represents garbage collection statistics
type GCStats struct {
	NumGC         int32   `json:"num_gc"`
	PauseTotal    int64   `json:"pause_total"`
	PauseAverage  float64 `json:"pause_average"`
	PauseMax      int64   `json:"pause_max"`
	PauseMin      int64   `json:"pause_min"`
	PauseP50      int64   `json:"pause_p50"`
	PauseP90      int64   `json:"pause_p90"`
	PauseP99      int64   `json:"pause_p99"`
	PauseP999     int64   `json:"pause_p999"`
	PauseP9999    int64   `json:"pause_p9999"`
	PauseP99999   int64   `json:"pause_p99999"`
	PauseP999999  int64   `json:"pause_p999999"`
	PauseP9999999 int64   `json:"pause_p9999999"`
}

// TraceSpan represents a tracing span
type TraceSpan struct {
	TraceID      string                 `json:"trace_id"`
	SpanID       string                 `json:"span_id"`
	ParentSpanID string                 `json:"parent_span_id,omitempty"`
	Operation    string                 `json:"operation"`
	StartTime    time.Time              `json:"start_time"`
	EndTime      time.Time              `json:"end_time"`
	Duration     time.Duration          `json:"duration"`
	Tags         map[string]string      `json:"tags"`
	Logs         []TraceLog             `json:"logs,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

// TraceLog represents a log entry within a trace span
type TraceLog struct {
	Timestamp time.Time              `json:"timestamp"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

// AdvancedMonitor provides comprehensive monitoring capabilities
type AdvancedMonitor struct {
	config        *MonitoringConfig
	metrics       []MetricPoint
	alerts        []Alert
	profilingData []ProfilingData
	traces        []TraceSpan

	// Internal state
	mu           sync.RWMutex
	ctx          context.Context
	cancel       context.CancelFunc
	lastCPUUsage float64
	lastMemStats runtime.MemStats
	lastGCStats  runtime.MemStats

	// Callbacks
	alertCallback   func(Alert)
	metricCallback  func(MetricPoint)
	profileCallback func(ProfilingData)
	traceCallback   func(TraceSpan)
}

// NewAdvancedMonitor creates a new advanced monitor
func NewAdvancedMonitor(config *MonitoringConfig) *AdvancedMonitor {
	if config == nil {
		config = DefaultMonitoringConfig()
	}

	ctx, cancel := context.WithCancel(context.Background())

	monitor := &AdvancedMonitor{
		config: config,
		ctx:    ctx,
		cancel: cancel,
	}

	// Start background monitoring
	if config.EnableMetrics {
		go monitor.metricsCollector()
	}

	if config.EnableProfiling {
		go monitor.profilingCollector()
	}

	return monitor
}

// RecordMetric records a metric point
func (m *AdvancedMonitor) RecordMetric(name string, metricType MetricType, value float64, labels map[string]string) {
	if !m.config.EnableMetrics {
		return
	}

	point := MetricPoint{
		Timestamp: time.Now(),
		Name:      name,
		Type:      metricType,
		Value:     value,
		Labels:    labels,
	}

	m.mu.Lock()
	m.metrics = append(m.metrics, point)

	// Trim old metrics if needed
	if len(m.metrics) > m.config.MaxMetricsPoints {
		m.metrics = m.metrics[len(m.metrics)-m.config.MaxMetricsPoints:]
	}
	m.mu.Unlock()

	// Check for alerts
	m.checkAlerts(point)

	// Call callback if set
	if m.metricCallback != nil {
		m.metricCallback(point)
	}
}

// RecordAlert records an alert
func (m *AdvancedMonitor) RecordAlert(level AlertLevel, message, metric string, value, threshold float64, labels map[string]string) {
	alert := Alert{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Metric:    metric,
		Value:     value,
		Threshold: threshold,
		Labels:    labels,
		Resolved:  false,
	}

	m.mu.Lock()
	m.alerts = append(m.alerts, alert)
	m.mu.Unlock()

	// Call callback if set
	if m.alertCallback != nil {
		m.alertCallback(alert)
	}
}

// RecordTrace records a trace span
func (m *AdvancedMonitor) RecordTrace(span TraceSpan) {
	if !m.config.EnableTracing {
		return
	}

	// Apply sampling
	if m.config.TracingSampleRate < 1.0 {
		if float64(time.Now().UnixNano()%1000000)/1000000.0 > m.config.TracingSampleRate {
			return
		}
	}

	// Calculate duration if not set
	if span.Duration == 0 && !span.EndTime.IsZero() && !span.StartTime.IsZero() {
		span.Duration = span.EndTime.Sub(span.StartTime)
	}

	m.mu.Lock()
	m.traces = append(m.traces, span)
	m.mu.Unlock()

	// Call callback if set
	if m.traceCallback != nil {
		m.traceCallback(span)
	}
}

// GetMetrics returns metrics within a time range
func (m *AdvancedMonitor) GetMetrics(start, end time.Time, name string) []MetricPoint {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var result []MetricPoint
	for _, point := range m.metrics {
		if point.Timestamp.After(start) && point.Timestamp.Before(end) {
			if name == "" || point.Name == name {
				result = append(result, point)
			}
		}
	}

	return result
}

// GetAlerts returns alerts within a time range
func (m *AdvancedMonitor) GetAlerts(start, end time.Time, level AlertLevel) []Alert {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var result []Alert
	for _, alert := range m.alerts {
		if alert.Timestamp.After(start) && alert.Timestamp.Before(end) {
			if level == "" || alert.Level == level {
				result = append(result, alert)
			}
		}
	}

	return result
}

// GetProfilingData returns profiling data within a time range
func (m *AdvancedMonitor) GetProfilingData(start, end time.Time) []ProfilingData {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var result []ProfilingData
	for _, data := range m.profilingData {
		if data.Timestamp.After(start) && data.Timestamp.Before(end) {
			result = append(result, data)
		}
	}

	return result
}

// GetTraces returns traces within a time range
func (m *AdvancedMonitor) GetTraces(start, end time.Time, traceID string) []TraceSpan {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var result []TraceSpan
	for _, trace := range m.traces {
		if trace.StartTime.After(start) && trace.StartTime.Before(end) {
			if traceID == "" || trace.TraceID == traceID {
				result = append(result, trace)
			}
		}
	}

	return result
}

// GetDashboardData returns data for monitoring dashboard
func (m *AdvancedMonitor) GetDashboardData() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()

	now := time.Now()
	last5Min := now.Add(-5 * time.Minute)
	lastHour := now.Add(-1 * time.Hour)

	// Get recent metrics
	recentMetrics := m.getMetricsInRange(last5Min, now, "")
	hourlyMetrics := m.getMetricsInRange(lastHour, now, "")

	// Get recent alerts
	recentAlerts := m.getAlertsInRange(last5Min, now, "")

	// Get recent profiling data
	recentProfiling := m.getProfilingDataInRange(last5Min, now)

	// Calculate statistics
	stats := m.calculateStatistics(recentMetrics, hourlyMetrics)

	return map[string]interface{}{
		"timestamp":         now,
		"metrics":           recentMetrics,
		"alerts":            recentAlerts,
		"profiling":         recentProfiling,
		"statistics":        stats,
		"system_info":       m.getSystemInfo(),
		"cache_performance": m.getCachePerformance(),
	}
}

// SetAlertCallback sets the alert callback function
func (m *AdvancedMonitor) SetAlertCallback(callback func(Alert)) {
	m.alertCallback = callback
}

// SetMetricCallback sets the metric callback function
func (m *AdvancedMonitor) SetMetricCallback(callback func(MetricPoint)) {
	m.metricCallback = callback
}

// SetProfileCallback sets the profiling callback function
func (m *AdvancedMonitor) SetProfileCallback(callback func(ProfilingData)) {
	m.profileCallback = callback
}

// SetTraceCallback sets the trace callback function
func (m *AdvancedMonitor) SetTraceCallback(callback func(TraceSpan)) {
	m.traceCallback = callback
}

// Close stops the monitor
func (m *AdvancedMonitor) Close() {
	m.cancel()
}

// Private methods

func (m *AdvancedMonitor) metricsCollector() {
	ticker := time.NewTicker(m.config.MetricsInterval)
	defer ticker.Stop()

	for {
		select {
		case <-m.ctx.Done():
			return
		case <-ticker.C:
			m.collectSystemMetrics()
		}
	}
}

func (m *AdvancedMonitor) profilingCollector() {
	ticker := time.NewTicker(m.config.ProfilingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-m.ctx.Done():
			return
		case <-ticker.C:
			m.collectProfilingData()
		}
	}
}

func (m *AdvancedMonitor) collectSystemMetrics() {
	// CPU usage
	cpuUsage := m.getCPUUsage()
	m.RecordMetric("cpu_usage", MetricTypeGauge, cpuUsage, map[string]string{"type": "system"})

	// Memory usage
	memStats := m.getMemoryStats()
	m.RecordMetric("memory_usage", MetricTypeGauge, float64(memStats.Alloc)/float64(memStats.Sys)*100, map[string]string{"type": "system"})
	m.RecordMetric("memory_alloc", MetricTypeGauge, float64(memStats.Alloc), map[string]string{"type": "system"})
	m.RecordMetric("memory_sys", MetricTypeGauge, float64(memStats.Sys), map[string]string{"type": "system"})

	// GC stats
	gcStats := m.getGCStats()
	m.RecordMetric("gc_pause_total", MetricTypeCounter, float64(gcStats.PauseTotal), map[string]string{"type": "system"})
	m.RecordMetric("gc_num", MetricTypeCounter, float64(gcStats.NumGC), map[string]string{"type": "system"})

	// Goroutine count
	goroutineCount := runtime.NumGoroutine()
	m.RecordMetric("goroutine_count", MetricTypeGauge, float64(goroutineCount), map[string]string{"type": "system"})
}

func (m *AdvancedMonitor) collectProfilingData() {
	profilingData := ProfilingData{
		Timestamp:     time.Now(),
		CPUProfile:    m.getCPUProfile(),
		MemoryProfile: m.getMemoryProfile(),
		GoroutineInfo: m.getGoroutineInfo(),
		GCStats:       m.getGCStats(),
	}

	m.mu.Lock()
	m.profilingData = append(m.profilingData, profilingData)

	// Trim old data if needed
	if len(m.profilingData) > m.config.MaxMetricsPoints {
		m.profilingData = m.profilingData[len(m.profilingData)-m.config.MaxMetricsPoints:]
	}
	m.mu.Unlock()

	// Call callback if set
	if m.profileCallback != nil {
		m.profileCallback(profilingData)
	}
}

func (m *AdvancedMonitor) checkAlerts(point MetricPoint) {
	if !m.config.EnableAlerts {
		return
	}

	threshold, exists := m.config.AlertThresholds[point.Name]
	if !exists {
		return
	}

	var level AlertLevel
	var message string

	if point.Value > threshold {
		if point.Value > threshold*1.5 {
			level = AlertLevelCritical
			message = fmt.Sprintf("Critical: %s is %.2f (threshold: %.2f)", point.Name, point.Value, threshold)
		} else {
			level = AlertLevelWarning
			message = fmt.Sprintf("Warning: %s is %.2f (threshold: %.2f)", point.Name, point.Value, threshold)
		}

		// Create alert directly to avoid recursion
		alert := Alert{
			ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
			Timestamp: time.Now(),
			Level:     level,
			Message:   message,
			Metric:    point.Name,
			Value:     point.Value,
			Threshold: threshold,
			Labels:    point.Labels,
			Resolved:  false,
		}

		m.mu.Lock()
		m.alerts = append(m.alerts, alert)
		m.mu.Unlock()

		// Call callback if set
		if m.alertCallback != nil {
			m.alertCallback(alert)
		}
	}
}

func (m *AdvancedMonitor) getCPUUsage() float64 {
	// Simplified CPU usage calculation
	// In a real implementation, you'd use more sophisticated methods
	return float64(runtime.NumGoroutine()) * 0.1
}

func (m *AdvancedMonitor) getMemoryStats() *MemoryProfile {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	return &MemoryProfile{
		Alloc:         memStats.Alloc,
		TotalAlloc:    memStats.TotalAlloc,
		Sys:           memStats.Sys,
		Lookups:       memStats.Lookups,
		Mallocs:       memStats.Mallocs,
		Frees:         memStats.Frees,
		HeapAlloc:     memStats.HeapAlloc,
		HeapSys:       memStats.HeapSys,
		HeapIdle:      memStats.HeapIdle,
		HeapInuse:     memStats.HeapInuse,
		HeapReleased:  memStats.HeapReleased,
		HeapObjects:   memStats.HeapObjects,
		StackInuse:    memStats.StackInuse,
		StackSys:      memStats.StackSys,
		MSpanInuse:    memStats.MSpanInuse,
		MSpanSys:      memStats.MSpanSys,
		MCacheInuse:   memStats.MCacheInuse,
		MCacheSys:     memStats.MCacheSys,
		BuckHashSys:   memStats.BuckHashSys,
		GCSys:         memStats.GCSys,
		OtherSys:      memStats.OtherSys,
		NextGC:        memStats.NextGC,
		LastGC:        memStats.LastGC,
		PauseTotalNs:  memStats.PauseTotalNs,
		NumGC:         int32(memStats.NumGC),
		NumForcedGC:   int32(memStats.NumForcedGC),
		GCCPUFraction: memStats.GCCPUFraction,
	}
}

func (m *AdvancedMonitor) getCPUProfile() *CPUProfile {
	// Simplified CPU profile
	return &CPUProfile{
		UsagePercent:    m.getCPUUsage(),
		UserTime:        0,
		SystemTime:      0,
		IdleTime:        0,
		IOWaitTime:      0,
		ContextSwitches: 0,
		Interrupts:      0,
	}
}

func (m *AdvancedMonitor) getMemoryProfile() *MemoryProfile {
	return m.getMemoryStats()
}

func (m *AdvancedMonitor) getGoroutineInfo() *GoroutineInfo {
	count := runtime.NumGoroutine()
	return &GoroutineInfo{
		Count:         count,
		BlockedCount:  0,
		RunnableCount: 0,
		RunningCount:  0,
		WaitingCount:  0,
		AverageStack:  0,
		MaxStack:      0,
		MinStack:      0,
	}
}

func (m *AdvancedMonitor) getGCStats() *GCStats {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	return &GCStats{
		NumGC:         int32(memStats.NumGC),
		PauseTotal:    int64(memStats.PauseTotalNs),
		PauseAverage:  float64(memStats.PauseTotalNs) / float64(memStats.NumGC),
		PauseMax:      int64(memStats.PauseNs[(memStats.NumGC+255)%256]),
		PauseMin:      int64(memStats.PauseNs[(memStats.NumGC+255)%256]),
		PauseP50:      0,
		PauseP90:      0,
		PauseP99:      0,
		PauseP999:     0,
		PauseP9999:    0,
		PauseP99999:   0,
		PauseP999999:  0,
		PauseP9999999: 0,
	}
}

func (m *AdvancedMonitor) getMetricsInRange(start, end time.Time, name string) []MetricPoint {
	var result []MetricPoint
	for _, point := range m.metrics {
		if point.Timestamp.After(start) && point.Timestamp.Before(end) {
			if name == "" || point.Name == name {
				result = append(result, point)
			}
		}
	}
	return result
}

func (m *AdvancedMonitor) getAlertsInRange(start, end time.Time, level AlertLevel) []Alert {
	var result []Alert
	for _, alert := range m.alerts {
		if alert.Timestamp.After(start) && alert.Timestamp.Before(end) {
			if level == "" || alert.Level == level {
				result = append(result, alert)
			}
		}
	}
	return result
}

func (m *AdvancedMonitor) getProfilingDataInRange(start, end time.Time) []ProfilingData {
	var result []ProfilingData
	for _, data := range m.profilingData {
		if data.Timestamp.After(start) && data.Timestamp.Before(end) {
			result = append(result, data)
		}
	}
	return result
}

func (m *AdvancedMonitor) calculateStatistics(recent, hourly []MetricPoint) map[string]interface{} {
	stats := make(map[string]interface{})

	// Calculate averages for recent metrics
	metricAverages := make(map[string]float64)
	metricCounts := make(map[string]int)

	for _, point := range recent {
		metricAverages[point.Name] += point.Value
		metricCounts[point.Name]++
	}

	for name, total := range metricAverages {
		if count := metricCounts[name]; count > 0 {
			metricAverages[name] = total / float64(count)
		}
	}

	stats["recent_averages"] = metricAverages
	stats["recent_count"] = len(recent)
	stats["hourly_count"] = len(hourly)

	return stats
}

func (m *AdvancedMonitor) getSystemInfo() map[string]interface{} {
	return map[string]interface{}{
		"go_version": runtime.Version(),
		"goos":       runtime.GOOS,
		"goarch":     runtime.GOARCH,
		"num_cpu":    runtime.NumCPU(),
		"goroutines": runtime.NumGoroutine(),
		"timestamp":  time.Now(),
	}
}

func (m *AdvancedMonitor) getCachePerformance() map[string]interface{} {
	// This would integrate with the actual cache systems
	return map[string]interface{}{
		"l1_hit_rate":    0.0,
		"l2_hit_rate":    0.0,
		"l3_hit_rate":    0.0,
		"total_hit_rate": 0.0,
		"throughput":     0.0,
		"latency_p50":    0.0,
		"latency_p90":    0.0,
		"latency_p99":    0.0,
	}
}

// ExportMetrics exports metrics in various formats
func (m *AdvancedMonitor) ExportMetrics(format string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	switch format {
	case "json":
		return json.Marshal(m.metrics)
	case "prometheus":
		return m.exportPrometheusFormat(), nil
	case "influxdb":
		return m.exportInfluxDBFormat(), nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}

func (m *AdvancedMonitor) exportPrometheusFormat() []byte {
	// Simplified Prometheus format export
	var result []byte
	for _, point := range m.metrics {
		line := fmt.Sprintf("%s{", point.Name)
		for k, v := range point.Labels {
			line += fmt.Sprintf("%s=\"%s\",", k, v)
		}
		line = line[:len(line)-1] // Remove trailing comma
		line += fmt.Sprintf("} %f %d\n", point.Value, point.Timestamp.UnixMilli())
		result = append(result, []byte(line)...)
	}
	return result
}

func (m *AdvancedMonitor) exportInfluxDBFormat() []byte {
	// Simplified InfluxDB format export
	var result []byte
	for _, point := range m.metrics {
		line := fmt.Sprintf("%s", point.Name)
		for k, v := range point.Labels {
			line += fmt.Sprintf(",%s=%s", k, v)
		}
		line += fmt.Sprintf(" value=%f %d\n", point.Value, point.Timestamp.UnixNano())
		result = append(result, []byte(line)...)
	}
	return result
}
