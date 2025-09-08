# 02 — Event Hub & Chain System (Hệ Thống Hub & Chuỗi Event)

**Generated:** 2025-01-27  
**Status:** Hub Design  
**Based on:** Event-driven architecture & causality patterns

## Tổng quan

Event Hub & Chain System là trung tâm liên kết các event với nhau, tạo ra hệ thống nhân quả và hiệu ứng cánh bướm trong game. Hệ thống này cho phép các event tương tác với nhau và tạo ra chuỗi sự kiện phức tạp.

## 🏗️ Event Hub Architecture (Kiến Trúc Event Hub)

### **EventHub - Hub Trung Tâm**

```go
// EventHub - Central event hub (Hub event trung tâm)
type EventHub struct {
    // Core components (Thành phần cốt lõi)
    registry        *EventRegistry        // Event registry (Đăng ký event)
    scheduler       *EventScheduler       // Event scheduler (Lịch trình event)
    chainManager    *EventChainManager    // Chain manager (Quản lý chuỗi)
    causalityEngine *CausalityEngine      // Causality engine (Động cơ nhân quả)
    
    // Event processing (Xử lý event)
    eventQueue      chan IEvent           // Event queue (Hàng đợi event)
    eventProcessor  *EventProcessor       // Event processor (Bộ xử lý event)
    eventValidator  *EventValidator       // Event validator (Bộ xác thực event)
    
    // Event storage (Lưu trữ event)
    eventStore      EventStore            // Event store (Kho lưu trữ event)
    eventLogger     EventLogger           // Event logger (Ghi log event)
    
    // Configuration (Cấu hình)
    config          *EventHubConfig       // Hub configuration (Cấu hình hub)
    metrics         *EventMetrics         // Event metrics (Chỉ số event)
}

// EventHubConfig - Hub configuration (Cấu hình hub)
type EventHubConfig struct {
    // Processing settings (Cài đặt xử lý)
    MaxConcurrentEvents int               `json:"max_concurrent_events"`
    EventTimeout        time.Duration     `json:"event_timeout"`
    RetryAttempts       int               `json:"retry_attempts"`
    RetryDelay          time.Duration     `json:"retry_delay"`
    
    // Chain settings (Cài đặt chuỗi)
    MaxChainDepth       int               `json:"max_chain_depth"`
    ChainTimeout        time.Duration     `json:"chain_timeout"`
    ParallelExecution   bool              `json:"parallel_execution"`
    
    // Causality settings (Cài đặt nhân quả)
    CausalityThreshold  float64           `json:"causality_threshold"`
    MaxCausalityDepth   int               `json:"max_causality_depth"`
    CausalityTimeout    time.Duration     `json:"causality_timeout"`
    
    // Storage settings (Cài đặt lưu trữ)
    EnableEventLogging  bool              `json:"enable_event_logging"`
    EventRetentionDays  int               `json:"event_retention_days"`
    EnableCompression   bool              `json:"enable_compression"`
}
```

### **EventRegistry - Đăng Ký Event**

```go
// EventRegistry - Event type registry (Đăng ký loại event)
type EventRegistry struct {
    // Event types (Loại event)
    eventTypes      map[string]EventTypeInfo    // Event type information (Thông tin loại event)
    eventHandlers   map[string][]EventHandler   // Event handlers (Bộ xử lý event)
    eventValidators map[string][]EventValidator // Event validators (Bộ xác thực event)
    
    // Event categories (Danh mục event)
    categories      map[string][]string         // Event categories (Danh mục event)
    priorities      map[string]int              // Event priorities (Độ ưu tiên event)
    
    // Event relationships (Mối quan hệ event)
    dependencies    map[string][]string         // Event dependencies (Phụ thuộc event)
    causalities     map[string][]string         // Event causalities (Nhân quả event)
    
    // Configuration (Cấu hình)
    config          *RegistryConfig             // Registry configuration (Cấu hình đăng ký)
}

// EventTypeInfo - Event type information (Thông tin loại event)
type EventTypeInfo struct {
    Type            string                      `json:"type"`
    Category        string                      `json:"category"`
    Priority        int                         `json:"priority"`
    Weight          float64                     `json:"weight"`
    Timeout         time.Duration               `json:"timeout"`
    Retryable       bool                        `json:"retryable"`
    Dependencies    []string                    `json:"dependencies"`
    Causalities     []string                    `json:"causalities"`
    DataSchema      map[string]interface{}      `json:"data_schema"`
    MetadataSchema  map[string]interface{}      `json:"metadata_schema"`
}
```

## 🔗 Event Chain System (Hệ Thống Chuỗi Event)

### **EventChainManager - Quản Lý Chuỗi Event**

```go
// EventChainManager - Event chain manager (Quản lý chuỗi event)
type EventChainManager struct {
    // Chain storage (Lưu trữ chuỗi)
    chains          map[string]*EventChain       // Active chains (Chuỗi hoạt động)
    chainHistory    map[string]*EventChain       // Chain history (Lịch sử chuỗi)
    
    // Chain execution (Thực thi chuỗi)
    executor        *ChainExecutor               // Chain executor (Thực thi chuỗi)
    scheduler       *ChainScheduler              // Chain scheduler (Lịch trình chuỗi)
    
    // Chain monitoring (Giám sát chuỗi)
    monitor         *ChainMonitor                // Chain monitor (Giám sát chuỗi)
    metrics         *ChainMetrics                // Chain metrics (Chỉ số chuỗi)
    
    // Configuration (Cấu hình)
    config          *ChainManagerConfig          // Manager configuration (Cấu hình quản lý)
}

// EventChain - Event chain structure (Cấu trúc chuỗi event)
type EventChain struct {
    // Chain info (Thông tin chuỗi)
    ChainID         string                      `json:"chain_id"`
    ChainType       string                      `json:"chain_type"`
    ChainOrder      int                         `json:"chain_order"`
    Status          ChainStatus                 `json:"status"`
    
    // Chain events (Sự kiện chuỗi)
    Events          []IEvent                    `json:"events"`
    EventOrder      []string                    `json:"event_order"`
    
    // Chain execution (Thực thi chuỗi)
    StartTime       time.Time                   `json:"start_time"`
    EndTime         time.Time                   `json:"end_time"`
    Duration        time.Duration               `json:"duration"`
    CurrentEvent    int                         `json:"current_event"`
    
    // Chain dependencies (Phụ thuộc chuỗi)
    Dependencies    []string                    `json:"dependencies"`
    Dependents      []string                    `json:"dependents"`
    
    // Chain configuration (Cấu hình chuỗi)
    Config          *ChainConfig                `json:"config"`
    Metadata        map[string]interface{}      `json:"metadata"`
}
```

### **ChainExecutor - Thực Thi Chuỗi**

```go
// ChainExecutor - Chain execution engine (Động cơ thực thi chuỗi)
type ChainExecutor struct {
    // Execution state (Trạng thái thực thi)
    activeChains    map[string]*EventChain      // Active chains (Chuỗi hoạt động)
    pausedChains    map[string]*EventChain      // Paused chains (Chuỗi tạm dừng)
    failedChains    map[string]*EventChain      // Failed chains (Chuỗi thất bại)
    
    // Execution control (Điều khiển thực thi)
    executionQueue  chan *EventChain            // Execution queue (Hàng đợi thực thi)
    executionPool   *WorkerPool                 // Execution pool (Pool thực thi)
    
    // Execution monitoring (Giám sát thực thi)
    monitor         *ExecutionMonitor           // Execution monitor (Giám sát thực thi)
    metrics         *ExecutionMetrics           // Execution metrics (Chỉ số thực thi)
    
    // Configuration (Cấu hình)
    config          *ExecutorConfig             // Executor configuration (Cấu hình thực thi)
}

// ChainConfig - Chain configuration (Cấu hình chuỗi)
type ChainConfig struct {
    // Execution settings (Cài đặt thực thi)
    ParallelExecution   bool                    `json:"parallel_execution"`
    MaxConcurrency      int                     `json:"max_concurrency"`
    ExecutionTimeout    time.Duration           `json:"execution_timeout"`
    RetryOnFailure      bool                    `json:"retry_on_failure"`
    MaxRetries          int                     `json:"max_retries"`
    
    // Chain settings (Cài đặt chuỗi)
    MaxChainDepth       int                     `json:"max_chain_depth"`
    ChainTimeout        time.Duration           `json:"chain_timeout"`
    EventTimeout        time.Duration           `json:"event_timeout"`
    
    // Dependencies (Phụ thuộc)
    WaitForDependencies bool                    `json:"wait_for_dependencies"`
    DependencyTimeout   time.Duration           `json:"dependency_timeout"`
}
```

## 🌊 Causality Engine (Động Cơ Nhân Quả)

### **CausalityEngine - Động Cơ Nhân Quả**

```go
// CausalityEngine - Causality analysis engine (Động cơ phân tích nhân quả)
type CausalityEngine struct {
    // Causality storage (Lưu trữ nhân quả)
    causalities     map[string]*CausalityNode   // Causality graph (Đồ thị nhân quả)
    causalityIndex  map[string][]string         // Causality index (Chỉ mục nhân quả)
    
    // Causality analysis (Phân tích nhân quả)
    analyzer        *CausalityAnalyzer          // Causality analyzer (Bộ phân tích nhân quả)
    predictor       *EventPredictor             // Event predictor (Bộ dự đoán event)
    
    // Causality monitoring (Giám sát nhân quả)
    monitor         *CausalityMonitor           // Causality monitor (Giám sát nhân quả)
    metrics         *CausalityMetrics           // Causality metrics (Chỉ số nhân quả)
    
    // Configuration (Cấu hình)
    config          *CausalityConfig            // Causality configuration (Cấu hình nhân quả)
}

// CausalityNode - Causality graph node (Nút đồ thị nhân quả)
type CausalityNode struct {
    // Node info (Thông tin nút)
    EventID         string                      `json:"event_id"`
    EventType       string                      `json:"event_type"`
    Timestamp       time.Time                   `json:"timestamp"`
    
    // Causality relationships (Mối quan hệ nhân quả)
    Causes          []*CausalityEdge            `json:"causes"`
    Effects         []*CausalityEdge            `json:"effects"`
    
    // Causality metrics (Chỉ số nhân quả)
    Impact          float64                     `json:"impact"`
    Influence       float64                     `json:"influence"`
    Centrality      float64                     `json:"centrality"`
    
    // Node metadata (Siêu dữ liệu nút)
    Metadata        map[string]interface{}      `json:"metadata"`
}

// CausalityEdge - Causality relationship edge (Cạnh mối quan hệ nhân quả)
type CausalityEdge struct {
    // Edge info (Thông tin cạnh)
    FromEventID     string                      `json:"from_event_id"`
    ToEventID       string                      `json:"to_event_id"`
    CausalityType   string                      `json:"causality_type"`
    Strength        float64                     `json:"strength"`
    
    // Edge metrics (Chỉ số cạnh)
    Confidence      float64                     `json:"confidence"`
    Latency         time.Duration               `json:"latency"`
    Frequency       float64                     `json:"frequency"`
    
    // Edge metadata (Siêu dữ liệu cạnh)
    Metadata        map[string]interface{}      `json:"metadata"`
}
```

### **CausalityAnalyzer - Bộ Phân Tích Nhân Quả**

```go
// CausalityAnalyzer - Causality analysis engine (Động cơ phân tích nhân quả)
type CausalityAnalyzer struct {
    // Analysis algorithms (Thuật toán phân tích)
    algorithms      map[string]CausalityAlgorithm // Analysis algorithms (Thuật toán phân tích)
    
    // Analysis cache (Cache phân tích)
    cache           *CausalityCache              // Analysis cache (Cache phân tích)
    
    // Analysis configuration (Cấu hình phân tích)
    config          *AnalyzerConfig              // Analyzer configuration (Cấu hình bộ phân tích)
}

// CausalityAlgorithm - Causality analysis algorithm (Thuật toán phân tích nhân quả)
type CausalityAlgorithm interface {
    // Analysis methods (Phương pháp phân tích)
    Analyze(event1, event2 IEvent) (*CausalityEdge, error)
    Predict(event IEvent) ([]IEvent, error)
    GetImpact(event IEvent) (float64, error)
    GetInfluence(event IEvent) (float64, error)
    
    // Algorithm info (Thông tin thuật toán)
    GetName() string
    GetDescription() string
    GetComplexity() int
    GetAccuracy() float64
}

// Common causality algorithms (Thuật toán nhân quả phổ biến)
type TemporalCausalityAlgorithm struct {
    // Temporal analysis (Phân tích thời gian)
    timeWindow      time.Duration               // Time window (Cửa sổ thời gian)
    minCorrelation  float64                     // Minimum correlation (Tương quan tối thiểu)
}

type SpatialCausalityAlgorithm struct {
    // Spatial analysis (Phân tích không gian)
    maxDistance     float64                     // Maximum distance (Khoảng cách tối đa)
    spatialWeight   float64                     // Spatial weight (Trọng số không gian)
}

type LogicalCausalityAlgorithm struct {
    // Logical analysis (Phân tích logic)
    ruleEngine      *RuleEngine                 // Rule engine (Động cơ quy tắc)
    logicWeight     float64                     // Logic weight (Trọng số logic)
}
```

## 🦋 Butterfly Effect System (Hệ Thống Hiệu Ứng Cánh Bướm)

### **ButterflyEffectEngine - Động Cơ Hiệu Ứng Cánh Bướm**

```go
// ButterflyEffectEngine - Butterfly effect engine (Động cơ hiệu ứng cánh bướm)
type ButterflyEffectEngine struct {
    // Effect tracking (Theo dõi hiệu ứng)
    effectChains    map[string]*EffectChain     // Effect chains (Chuỗi hiệu ứng)
    effectIndex     map[string][]string         // Effect index (Chỉ mục hiệu ứng)
    
    // Effect analysis (Phân tích hiệu ứng)
    analyzer        *EffectAnalyzer             // Effect analyzer (Bộ phân tích hiệu ứng)
    predictor       *EffectPredictor            // Effect predictor (Bộ dự đoán hiệu ứng)
    
    // Effect monitoring (Giám sát hiệu ứng)
    monitor         *EffectMonitor              // Effect monitor (Giám sát hiệu ứng)
    metrics         *EffectMetrics              // Effect metrics (Chỉ số hiệu ứng)
    
    // Configuration (Cấu hình)
    config          *ButterflyEffectConfig      // Effect configuration (Cấu hình hiệu ứng)
}

// EffectChain - Butterfly effect chain (Chuỗi hiệu ứng cánh bướm)
type EffectChain struct {
    // Chain info (Thông tin chuỗi)
    ChainID         string                      `json:"chain_id"`
    TriggerEvent    IEvent                      `json:"trigger_event"`
    EffectEvents    []IEvent                    `json:"effect_events"`
    ChainDepth      int                         `json:"chain_depth"`
    
    // Effect propagation (Lan truyền hiệu ứng)
    PropagationPath []string                    `json:"propagation_path"`
    PropagationTime time.Duration               `json:"propagation_time"`
    Amplification   float64                     `json:"amplification"`
    
    // Effect analysis (Phân tích hiệu ứng)
    Impact          float64                     `json:"impact"`
    Reach           int                         `json:"reach"`
    Duration        time.Duration               `json:"duration"`
    
    // Chain metadata (Siêu dữ liệu chuỗi)
    Metadata        map[string]interface{}      `json:"metadata"`
}

// EffectAnalyzer - Butterfly effect analyzer (Bộ phân tích hiệu ứng cánh bướm)
type EffectAnalyzer struct {
    // Analysis algorithms (Thuật toán phân tích)
    algorithms      map[string]EffectAlgorithm  // Effect algorithms (Thuật toán hiệu ứng)
    
    // Analysis cache (Cache phân tích)
    cache           *EffectCache                // Effect cache (Cache hiệu ứng)
    
    // Analysis configuration (Cấu hình phân tích)
    config          *EffectAnalyzerConfig       // Analyzer configuration (Cấu hình bộ phân tích)
}

// EffectAlgorithm - Butterfly effect algorithm (Thuật toán hiệu ứng cánh bướm)
type EffectAlgorithm interface {
    // Analysis methods (Phương pháp phân tích)
    Analyze(trigger IEvent) (*EffectChain, error)
    Predict(trigger IEvent) ([]IEvent, error)
    GetAmplification(chain *EffectChain) (float64, error)
    GetReach(chain *EffectChain) (int, error)
    
    // Algorithm info (Thông tin thuật toán)
    GetName() string
    GetDescription() string
    GetComplexity() int
    GetAccuracy() float64
}
```

## 📊 Event Chain Examples (Ví Dụ Chuỗi Event)

### **Combat Chain Example (Ví Dụ Chuỗi Chiến Đấu)**

```go
// CombatChain - Combat event chain example (Ví dụ chuỗi event chiến đấu)
type CombatChain struct {
    // Chain sequence (Trình tự chuỗi)
    Events []IEvent
    
    // Chain flow (Luồng chuỗi)
    Flow []string
}

// Example combat chain (Ví dụ chuỗi chiến đấu)
func CreateCombatChain() *CombatChain {
    return &CombatChain{
        Events: []IEvent{
            &CombatEvent{
                EventData: EventData{
                    EventType: "combat_start",
                    Category:  "combat",
                },
                CombatType: "player_vs_monster",
            },
            &CombatEvent{
                EventData: EventData{
                    EventType: "attack",
                    Category:  "combat",
                },
                SkillUsed: "fireball",
                Damage:    100.0,
            },
            &CombatEvent{
                EventData: EventData{
                    EventType: "damage_dealt",
                    Category:  "combat",
                },
                Damage: 100.0,
            },
            &CombatEvent{
                EventData: EventData{
                    EventType: "monster_death",
                    Category:  "combat",
                },
                Victory: true,
            },
            &ItemEvent{
                EventData: EventData{
                    EventType: "item_drop",
                    Category:  "item",
                },
                ItemID: "sword_001",
                Operation: ItemOperationDrop,
            },
            &CultivationEvent{
                EventData: EventData{
                    EventType: "experience_gain",
                    Category:  "cultivation",
                },
                ExperienceGained: 50.0,
            },
        },
        Flow: []string{
            "combat_start",
            "attack",
            "damage_dealt",
            "monster_death",
            "item_drop",
            "experience_gain",
        },
    }
}
```

### **World Creation Chain Example (Ví Dụ Chuỗi Sáng Tạo Thế Giới)**

```go
// WorldCreationChain - World creation chain example (Ví dụ chuỗi sáng tạo thế giới)
type WorldCreationChain struct {
    // Chain sequence (Trình tự chuỗi)
    Events []IEvent
    
    // Chain flow (Luồng chuỗi)
    Flow []string
}

// Example world creation chain (Ví dụ chuỗi sáng tạo thế giới)
func CreateWorldCreationChain() *WorldCreationChain {
    return &WorldCreationChain{
        Events: []IEvent{
            &WorldCreationEvent{
                EventData: EventData{
                    EventType: "world_creation_start",
                    Category:  "world_creation",
                },
                WorldType: "dungeon",
                CreationMethod: "player_triggered",
            },
            &WorldCreationEvent{
                EventData: EventData{
                    EventType: "terrain_generation",
                    Category:  "world_creation",
                },
                TerrainFeatures: []string{"mountains", "rivers", "forests"},
            },
            &WorldCreationEvent{
                EventData: EventData{
                    EventType: "resource_spawning",
                    Category:  "world_creation",
                },
                ResourceNodes: []string{"iron_ore", "gold_ore", "magic_crystal"},
            },
            &WorldCreationEvent{
                EventData: EventData{
                    EventType: "monster_spawning",
                    Category:  "world_creation",
                },
                SpawnPoints: []string{"cave_entrance", "forest_clearing", "mountain_peak"},
            },
            &WorldCreationEvent{
                EventData: EventData{
                    EventType: "world_creation_complete",
                    Category:  "world_creation",
                },
                WorldID: "dungeon_001",
            },
        },
        Flow: []string{
            "world_creation_start",
            "terrain_generation",
            "resource_spawning",
            "monster_spawning",
            "world_creation_complete",
        },
    }
}
```

## 🔧 Implementation Strategy (Chiến Lược Triển Khai)

### **Phase 1: Basic Hub (Giai Đoạn 1: Hub Cơ Bản)**
1. Implement EventHub core
2. Create EventRegistry
3. Implement basic event processing
4. Add event logging

### **Phase 2: Chain System (Giai Đoạn 2: Hệ Thống Chuỗi)**
1. Implement EventChainManager
2. Create ChainExecutor
3. Add chain monitoring
4. Implement chain dependencies

### **Phase 3: Causality Engine (Giai Đoạn 3: Động Cơ Nhân Quả)**
1. Implement CausalityEngine
2. Create CausalityAnalyzer
3. Add causality algorithms
4. Implement event prediction

### **Phase 4: Butterfly Effect (Giai Đoạn 4: Hiệu Ứng Cánh Bướm)**
1. Implement ButterflyEffectEngine
2. Create EffectAnalyzer
3. Add effect algorithms
4. Implement effect prediction

## 💡 Best Practices (Thực Hành Tốt Nhất)

### **Event Chain Design (Thiết Kế Chuỗi Event)**
1. **Clear Dependencies** - Phụ thuộc rõ ràng
2. **Fail-Safe Design** - Thiết kế an toàn
3. **Timeout Handling** - Xử lý timeout
4. **Rollback Support** - Hỗ trợ rollback

### **Causality Analysis (Phân Tích Nhân Quả)**
1. **Statistical Analysis** - Phân tích thống kê
2. **Temporal Analysis** - Phân tích thời gian
3. **Spatial Analysis** - Phân tích không gian
4. **Logical Analysis** - Phân tích logic

### **Performance Optimization (Tối Ưu Hiệu Suất)**
1. **Async Processing** - Xử lý bất đồng bộ
2. **Caching** - Cache dữ liệu
3. **Batch Processing** - Xử lý theo lô
4. **Load Balancing** - Cân bằng tải
