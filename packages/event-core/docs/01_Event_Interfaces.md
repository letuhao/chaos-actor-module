# 01 — Event Interfaces (Giao Diện Event)

**Generated:** 2025-01-27  
**Status:** Interface Design  
**Based on:** Event-driven architecture patterns

## Tổng quan

Event Interfaces định nghĩa các giao diện trừu tượng cho tất cả các loại event trong game, cho phép các hệ thống cụ thể kế thừa và implement theo chuẩn chung.

## 🏗️ Core Event Interface (Giao Diện Event Cốt Lõi)

### **IEvent - Interface Cơ Bản**

```go
// IEvent - Base event interface (Giao diện event cơ bản)
type IEvent interface {
    // Basic properties (Thuộc tính cơ bản)
    GetEventID() string                    // Unique event identifier (ID duy nhất)
    GetEventType() string                  // Event type (Loại event)
    GetEventCategory() string              // Event category (Danh mục event)
    GetTimestamp() time.Time               // Event timestamp (Thời gian event)
    GetActorID() string                    // Actor who triggered event (Actor kích hoạt)
    GetTargetID() string                   // Target of event (Mục tiêu event)
    GetWorldID() string                    // World where event occurred (Thế giới xảy ra)
    
    // Event data (Dữ liệu event)
    GetData() map[string]interface{}       // Event-specific data (Dữ liệu cụ thể)
    GetMetadata() map[string]interface{}   // Event metadata (Siêu dữ liệu)
    
    // Event lifecycle (Vòng đời event)
    Validate() error                       // Validate event data (Xác thực dữ liệu)
    Process() error                        // Process event (Xử lý event)
    Complete() error                       // Complete event (Hoàn thành event)
    Rollback() error                       // Rollback event (Hoàn tác event)
    
    // Event relationships (Mối quan hệ event)
    GetParentEventID() string              // Parent event ID (ID event cha)
    GetChildEventIDs() []string            // Child event IDs (ID event con)
    GetDependencies() []string             // Event dependencies (Phụ thuộc event)
    
    // Event priority (Độ ưu tiên event)
    GetPriority() int                      // Event priority (Độ ưu tiên)
    GetWeight() float64                    // Event weight (Trọng số event)
}
```

### **IEventChain - Interface Chuỗi Event**

```go
// IEventChain - Event chain interface (Giao diện chuỗi event)
type IEventChain interface {
    // Chain properties (Thuộc tính chuỗi)
    GetChainID() string                    // Chain identifier (ID chuỗi)
    GetChainType() string                  // Chain type (Loại chuỗi)
    GetChainOrder() int                    // Chain order (Thứ tự chuỗi)
    
    // Chain management (Quản lý chuỗi)
    AddEvent(event IEvent) error           // Add event to chain (Thêm event vào chuỗi)
    RemoveEvent(eventID string) error      // Remove event from chain (Xóa event khỏi chuỗi)
    GetEvents() []IEvent                   // Get all events in chain (Lấy tất cả event)
    
    // Chain execution (Thực thi chuỗi)
    Execute() error                        // Execute entire chain (Thực thi toàn bộ chuỗi)
    Pause() error                          // Pause chain execution (Tạm dừng chuỗi)
    Resume() error                         // Resume chain execution (Tiếp tục chuỗi)
    Cancel() error                         // Cancel chain execution (Hủy chuỗi)
    
    // Chain status (Trạng thái chuỗi)
    GetStatus() ChainStatus                // Get chain status (Lấy trạng thái chuỗi)
    IsComplete() bool                      // Check if chain is complete (Kiểm tra hoàn thành)
    IsFailed() bool                        // Check if chain failed (Kiểm tra thất bại)
}
```

### **IEventCausality - Interface Nhân Quả Event**

```go
// IEventCausality - Event causality interface (Giao diện nhân quả event)
type IEventCausality interface {
    // Causality properties (Thuộc tính nhân quả)
    GetCauseEventID() string               // Cause event ID (ID event nguyên nhân)
    GetEffectEventID() string              // Effect event ID (ID event hậu quả)
    GetCausalityType() string              // Causality type (Loại nhân quả)
    GetStrength() float64                  // Causality strength (Sức mạnh nhân quả)
    
    // Causality analysis (Phân tích nhân quả)
    Analyze() error                        // Analyze causality (Phân tích nhân quả)
    Predict() []IEvent                     // Predict future events (Dự đoán event tương lai)
    GetImpact() float64                    // Get impact score (Lấy điểm tác động)
    
    // Causality management (Quản lý nhân quả)
    AddCausality(cause, effect string, strength float64) error
    RemoveCausality(cause, effect string) error
    GetCausalityChain() []IEvent           // Get causality chain (Lấy chuỗi nhân quả)
}
```

## 🎮 Game-Specific Event Interfaces (Giao Diện Event Cụ Thể)

### **ICombatEvent - Interface Event Chiến Đấu**

```go
// ICombatEvent - Combat event interface (Giao diện event chiến đấu)
type ICombatEvent interface {
    IEvent                                 // Inherit base event (Kế thừa event cơ bản)
    
    // Combat properties (Thuộc tính chiến đấu)
    GetCombatType() string                 // Combat type (Loại chiến đấu)
    GetDamage() float64                    // Damage dealt (Sát thương gây ra)
    GetHealing() float64                   // Healing done (Hồi phục thực hiện)
    GetSkillUsed() string                  // Skill used (Kỹ năng sử dụng)
    GetWeaponUsed() string                 // Weapon used (Vũ khí sử dụng)
    
    // Combat state (Trạng thái chiến đấu)
    GetCombatState() CombatState           // Current combat state (Trạng thái hiện tại)
    GetCombatPhase() CombatPhase           // Current combat phase (Giai đoạn hiện tại)
    GetCombatDuration() time.Duration      // Combat duration (Thời gian chiến đấu)
    
    // Combat results (Kết quả chiến đấu)
    GetVictory() bool                      // Victory status (Trạng thái thắng)
    GetDefeat() bool                       // Defeat status (Trạng thái thua)
    GetExperienceGained() float64          // Experience gained (Kinh nghiệm nhận được)
    GetLootDropped() []string              // Loot dropped (Đồ rơi)
}
```

### **IItemEvent - Interface Event Vật Phẩm**

```go
// IItemEvent - Item event interface (Giao diện event vật phẩm)
type IItemEvent interface {
    IEvent                                 // Inherit base event (Kế thừa event cơ bản)
    
    // Item properties (Thuộc tính vật phẩm)
    GetItemID() string                     // Item identifier (ID vật phẩm)
    GetItemType() string                   // Item type (Loại vật phẩm)
    GetItemRarity() string                 // Item rarity (Độ hiếm)
    GetItemQuantity() int                  // Item quantity (Số lượng)
    GetItemValue() float64                 // Item value (Giá trị)
    
    // Item operations (Thao tác vật phẩm)
    GetOperation() ItemOperation           // Operation performed (Thao tác thực hiện)
    GetSourceLocation() string             // Source location (Vị trí nguồn)
    GetDestinationLocation() string        // Destination location (Vị trí đích)
    
    // Item enhancement (Cường hóa vật phẩm)
    GetEnhancementLevel() int              // Enhancement level (Cấp độ cường hóa)
    GetEnhancementSuccess() bool           // Enhancement success (Thành công cường hóa)
    GetEnhancementCost() float64           // Enhancement cost (Chi phí cường hóa)
}
```

### **ICultivationEvent - Interface Event Tu Luyện**

```go
// ICultivationEvent - Cultivation event interface (Giao diện event tu luyện)
type ICultivationEvent interface {
    IEvent                                 // Inherit base event (Kế thừa event cơ bản)
    
    // Cultivation properties (Thuộc tính tu luyện)
    GetCultivationType() string            // Cultivation type (Loại tu luyện)
    GetRealm() string                      // Current realm (Cảnh giới hiện tại)
    GetSubstage() string                   // Current substage (Tiểu cảnh giới hiện tại)
    GetCultivationMethod() string          // Cultivation method (Phương pháp tu luyện)
    
    // Cultivation progress (Tiến độ tu luyện)
    GetProgressGained() float64            // Progress gained (Tiến độ nhận được)
    GetExperienceGained() float64          // Experience gained (Kinh nghiệm nhận được)
    GetEnergyConsumed() float64            // Energy consumed (Năng lượng tiêu hao)
    GetTimeSpent() time.Duration           // Time spent (Thời gian bỏ ra)
    
    // Cultivation results (Kết quả tu luyện)
    GetBreakthrough() bool                 // Breakthrough achieved (Đạt đột phá)
    GetNewRealm() string                   // New realm reached (Cảnh giới mới đạt)
    GetNewSubstage() string                // New substage reached (Tiểu cảnh giới mới đạt)
    GetSkillsLearned() []string            // Skills learned (Kỹ năng học được)
}
```

### **IQuestEvent - Interface Event Nhiệm Vụ**

```go
// IQuestEvent - Quest event interface (Giao diện event nhiệm vụ)
type IQuestEvent interface {
    IEvent                                 // Inherit base event (Kế thừa event cơ bản)
    
    // Quest properties (Thuộc tính nhiệm vụ)
    GetQuestID() string                    // Quest identifier (ID nhiệm vụ)
    GetQuestType() string                  // Quest type (Loại nhiệm vụ)
    GetQuestDifficulty() string            // Quest difficulty (Độ khó nhiệm vụ)
    GetQuestLevel() int                    // Quest level (Cấp độ nhiệm vụ)
    
    // Quest progress (Tiến độ nhiệm vụ)
    GetProgress() float64                  // Current progress (Tiến độ hiện tại)
    GetMaxProgress() float64               // Maximum progress (Tiến độ tối đa)
    GetObjectives() []string               // Quest objectives (Mục tiêu nhiệm vụ)
    GetCompletedObjectives() []string      // Completed objectives (Mục tiêu hoàn thành)
    
    // Quest rewards (Phần thưởng nhiệm vụ)
    GetRewards() map[string]interface{}    // Quest rewards (Phần thưởng nhiệm vụ)
    GetExperienceReward() float64          // Experience reward (Phần thưởng kinh nghiệm)
    GetItemRewards() []string              // Item rewards (Phần thưởng vật phẩm)
    GetCurrencyReward() float64            // Currency reward (Phần thưởng tiền tệ)
}
```

### **IWorldCreationEvent - Interface Event Sáng Tạo Thế Giới**

```go
// IWorldCreationEvent - World creation event interface (Giao diện event sáng tạo thế giới)
type IWorldCreationEvent interface {
    IEvent                                 // Inherit base event (Kế thừa event cơ bản)
    
    // World properties (Thuộc tính thế giới)
    GetWorldID() string                    // World identifier (ID thế giới)
    GetWorldType() string                  // World type (Loại thế giới)
    GetWorldSize() WorldSize               // World size (Kích thước thế giới)
    GetWorldBiome() string                 // World biome (Sinh quyển thế giới)
    
    // Creation details (Chi tiết sáng tạo)
    GetCreationMethod() string             // Creation method (Phương pháp sáng tạo)
    GetCreationCost() float64              // Creation cost (Chi phí sáng tạo)
    GetCreationTime() time.Duration        // Creation time (Thời gian sáng tạo)
    GetCreationResources() []string        // Resources used (Tài nguyên sử dụng)
    
    // World features (Đặc điểm thế giới)
    GetTerrainFeatures() []string          // Terrain features (Đặc điểm địa hình)
    GetResourceNodes() []string            // Resource nodes (Nút tài nguyên)
    GetSpawnPoints() []string              // Spawn points (Điểm sinh)
    GetSpecialAreas() []string             // Special areas (Khu vực đặc biệt)
}
```

## 🔧 Event Status & States (Trạng Thái Event)

### **EventStatus - Trạng Thái Event**

```go
// EventStatus - Event status enum (Enum trạng thái event)
type EventStatus int

const (
    EventStatusPending    EventStatus = iota // Pending (Chờ xử lý)
    EventStatusProcessing                     // Processing (Đang xử lý)
    EventStatusCompleted                      // Completed (Hoàn thành)
    EventStatusFailed                         // Failed (Thất bại)
    EventStatusCancelled                      // Cancelled (Hủy bỏ)
    EventStatusRolledBack                     // Rolled back (Hoàn tác)
)
```

### **ChainStatus - Trạng Thái Chuỗi**

```go
// ChainStatus - Chain status enum (Enum trạng thái chuỗi)
type ChainStatus int

const (
    ChainStatusPending    ChainStatus = iota // Pending (Chờ xử lý)
    ChainStatusExecuting                      // Executing (Đang thực thi)
    ChainStatusPaused                         // Paused (Tạm dừng)
    ChainStatusCompleted                      // Completed (Hoàn thành)
    ChainStatusFailed                         // Failed (Thất bại)
    ChainStatusCancelled                      // Cancelled (Hủy bỏ)
)
```

### **CombatState - Trạng Thái Chiến Đấu**

```go
// CombatState - Combat state enum (Enum trạng thái chiến đấu)
type CombatState int

const (
    CombatStateIdle        CombatState = iota // Idle (Nhàn rỗi)
    CombatStateEngaging                       // Engaging (Tham gia)
    CombatStateFighting                       // Fighting (Chiến đấu)
    CombatStateRetreating                     // Retreating (Rút lui)
    CombatStateVictory                        // Victory (Thắng)
    CombatStateDefeat                         // Defeat (Thua)
)
```

### **CombatPhase - Giai Đoạn Chiến Đấu**

```go
// CombatPhase - Combat phase enum (Enum giai đoạn chiến đấu)
type CombatPhase int

const (
    CombatPhasePreparation CombatPhase = iota // Preparation (Chuẩn bị)
    CombatPhaseEngagement                      // Engagement (Tham gia)
    CombatPhaseClash                           // Clash (Đụng độ)
    CombatPhaseResolution                      // Resolution (Giải quyết)
    CombatPhaseCleanup                         // Cleanup (Dọn dẹp)
)
```

### **ItemOperation - Thao Tác Vật Phẩm**

```go
// ItemOperation - Item operation enum (Enum thao tác vật phẩm)
type ItemOperation int

const (
    ItemOperationCreate     ItemOperation = iota // Create (Tạo)
    ItemOperationDestroy                          // Destroy (Phá hủy)
    ItemOperationEnhance                          // Enhance (Cường hóa)
    ItemOperationTrade                            // Trade (Giao dịch)
    ItemOperationDrop                             // Drop (Rơi)
    ItemOperationPickup                           // Pickup (Nhặt)
    ItemOperationUse                              // Use (Sử dụng)
    ItemOperationRepair                           // Repair (Sửa chữa)
)
```

## 📊 Event Data Structures (Cấu Trúc Dữ Liệu Event)

### **EventData - Dữ Liệu Event**

```go
// EventData - Event data structure (Cấu trúc dữ liệu event)
type EventData struct {
    // Basic info (Thông tin cơ bản)
    EventID      string                 `json:"event_id"`
    EventType    string                 `json:"event_type"`
    Category     string                 `json:"category"`
    Timestamp    time.Time              `json:"timestamp"`
    ActorID      string                 `json:"actor_id"`
    TargetID     string                 `json:"target_id"`
    WorldID      string                 `json:"world_id"`
    
    // Event-specific data (Dữ liệu cụ thể)
    Data         map[string]interface{} `json:"data"`
    Metadata     map[string]interface{} `json:"metadata"`
    
    // Relationships (Mối quan hệ)
    ParentEventID string                `json:"parent_event_id"`
    ChildEventIDs []string              `json:"child_event_ids"`
    Dependencies  []string              `json:"dependencies"`
    
    // Priority (Độ ưu tiên)
    Priority     int                    `json:"priority"`
    Weight       float64                `json:"weight"`
    
    // Status (Trạng thái)
    Status       EventStatus            `json:"status"`
    CreatedAt    time.Time              `json:"created_at"`
    UpdatedAt    time.Time              `json:"updated_at"`
}
```

### **EventChainData - Dữ Liệu Chuỗi Event**

```go
// EventChainData - Event chain data structure (Cấu trúc dữ liệu chuỗi event)
type EventChainData struct {
    // Chain info (Thông tin chuỗi)
    ChainID      string                 `json:"chain_id"`
    ChainType    string                 `json:"chain_type"`
    ChainOrder   int                    `json:"chain_order"`
    
    // Events (Sự kiện)
    Events       []EventData            `json:"events"`
    
    // Status (Trạng thái)
    Status       ChainStatus            `json:"status"`
    CreatedAt    time.Time              `json:"created_at"`
    UpdatedAt    time.Time              `json:"updated_at"`
}
```

### **EventCausalityData - Dữ Liệu Nhân Quả Event**

```go
// EventCausalityData - Event causality data structure (Cấu trúc dữ liệu nhân quả event)
type EventCausalityData struct {
    // Causality info (Thông tin nhân quả)
    CauseEventID  string                `json:"cause_event_id"`
    EffectEventID string                `json:"effect_event_id"`
    CausalityType string                `json:"causality_type"`
    Strength      float64               `json:"strength"`
    
    // Analysis (Phân tích)
    Impact        float64               `json:"impact"`
    Confidence    float64               `json:"confidence"`
    
    // Timestamps (Thời gian)
    CreatedAt     time.Time             `json:"created_at"`
    UpdatedAt     time.Time             `json:"updated_at"`
}
```

## 🚀 Implementation Examples (Ví Dụ Triển Khai)

### **Combat Event Implementation**

```go
// CombatEvent - Combat event implementation (Triển khai event chiến đấu)
type CombatEvent struct {
    EventData
    CombatType      string
    Damage          float64
    Healing         float64
    SkillUsed       string
    WeaponUsed      string
    CombatState     CombatState
    CombatPhase     CombatPhase
    Duration        time.Duration
    Victory         bool
    ExperienceGained float64
    LootDropped     []string
}

// Implement ICombatEvent interface
func (ce *CombatEvent) GetCombatType() string {
    return ce.CombatType
}

func (ce *CombatEvent) GetDamage() float64 {
    return ce.Damage
}

func (ce *CombatEvent) GetHealing() float64 {
    return ce.Healing
}

// ... other interface methods
```

### **Item Event Implementation**

```go
// ItemEvent - Item event implementation (Triển khai event vật phẩm)
type ItemEvent struct {
    EventData
    ItemID              string
    ItemType            string
    ItemRarity          string
    Quantity            int
    Value               float64
    Operation           ItemOperation
    SourceLocation      string
    DestinationLocation string
    EnhancementLevel    int
    EnhancementSuccess  bool
    EnhancementCost     float64
}

// Implement IItemEvent interface
func (ie *ItemEvent) GetItemID() string {
    return ie.ItemID
}

func (ie *ItemEvent) GetItemType() string {
    return ie.ItemType
}

// ... other interface methods
```

## 💡 Best Practices (Thực Hành Tốt Nhất)

### **Event Design (Thiết Kế Event)**
1. **Single Responsibility** - Mỗi event chỉ có một trách nhiệm
2. **Immutable Data** - Dữ liệu event không thay đổi sau khi tạo
3. **Clear Naming** - Đặt tên rõ ràng và có ý nghĩa
4. **Consistent Structure** - Cấu trúc nhất quán

### **Interface Design (Thiết Kế Interface)**
1. **Minimal Interface** - Interface tối thiểu và cần thiết
2. **Clear Contracts** - Hợp đồng rõ ràng
3. **Extensible** - Có thể mở rộng
4. **Testable** - Có thể kiểm thử

### **Performance Considerations (Xem Xét Hiệu Suất)**
1. **Lazy Loading** - Tải dữ liệu khi cần
2. **Caching** - Cache dữ liệu thường dùng
3. **Batch Processing** - Xử lý theo lô
4. **Async Processing** - Xử lý bất đồng bộ
