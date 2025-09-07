# 02 — Dao Data Structures (Cấu Trúc Dữ Liệu Đạo)

**Generated:** 2025-01-27  
**Based on:** Map-based approach for scalability

## Tổng quan

Hệ thống Đạo sử dụng **map-based approach** để quản lý nhiều loại Đạo một cách linh hoạt và dễ mở rộng.

## 🏗️ Cấu Trúc Chính

### **DaoSystem - Hệ Thống Đạo Chính**

```go
// Dao System (道系统) - Hệ thống Đạo
type DaoSystem struct {
    // Dao Paths Map (道径映射) - Map các con đường Đạo
    // Key: DaoType string, Value: DaoPath
    DaoPaths map[string]DaoPath `json:"dao_paths"`
    
    // Primary Dao Type (主道类型) - Loại Đạo chính
    PrimaryDaoType string `json:"primary_dao_type"`
    
    // Dao Comprehension (道理解) - Sự hiểu biết về Đạo
    DaoComprehension DaoComprehension `json:"dao_comprehension"`
    
    // Dao Techniques Map (道术映射) - Map kỹ thuật Đạo
    // Key: DaoType string, Value: []DaoTechnique
    DaoTechniques map[string][]DaoTechnique `json:"dao_techniques"`
    
    // Dao Compatibility (道兼容性) - Tương thích Đạo
    DaoCompatibility DaoCompatibility `json:"dao_compatibility"`
}
```

### **DaoPath - Con Đường Đạo**

```go
// Dao Path (道径) - Con đường Đạo
type DaoPath struct {
    // Basic info (基本信息) - Thông tin cơ bản
    DaoType        string  `json:"dao_type"`        // Type of Dao (道类型) - Loại Đạo
    DaoLevel       int     `json:"dao_level"`       // Dao level (道等级) - Cấp độ Đạo
    DaoExperience  float64 `json:"dao_experience"`  // Dao experience (道经验) - Kinh nghiệm Đạo
    DaoMastery     float64 `json:"dao_mastery"`     // Dao mastery (道掌握) - Mức độ thành thạo
    
    // Advanced properties (高级属性) - Thuộc tính nâng cao
    DaoInsight     float64 `json:"dao_insight"`     // Dao insight (道洞察) - Sự thấu hiểu Đạo
    DaoWisdom      float64 `json:"dao_wisdom"`      // Dao wisdom (道智慧) - Trí tuệ Đạo
    DaoHarmony     float64 `json:"dao_harmony"`     // Dao harmony (道和谐) - Sự hài hòa Đạo
    DaoTranscendence float64 `json:"dao_transcendence"` // Dao transcendence (道超脱) - Siêu thoát Đạo
    
    // Status flags (状态标志) - Cờ trạng thái
    IsActive       bool    `json:"is_active"`       // Is this Dao active (是否激活) - Đạo này có đang hoạt động không
    IsPrimary      bool    `json:"is_primary"`      // Is this the primary Dao (是否主道) - Có phải Đạo chính không
    UnlockedAt     int64   `json:"unlocked_at"`     // When this Dao was unlocked (解锁时间) - Thời điểm mở khóa Đạo này
}
```

## 🏷️ Định Nghĩa Loại Đạo

### **Dao Type Constants**

```go
// Dao Type Constants (道类型常量) - Hằng số loại Đạo
const (
    DaoTypeSword     = "sword_dao"     // Sword Dao (剑道) - Kiếm Đạo
    DaoTypeBlade     = "blade_dao"     // Blade Dao (刀道) - Đao Đạo
    DaoTypeFist      = "fist_dao"      // Fist Dao (拳道) - Quyền Đạo
    DaoTypeSpell     = "spell_dao"     // Spell Dao (法道) - Pháp Đạo
    DaoTypeFormation = "formation_dao" // Formation Dao (阵道) - Trận Đạo
    DaoTypeRefinement = "refinement_dao" // Refinement Dao (炼道) - Luyện Đạo
    DaoTypeHeart     = "heart_dao"     // Heart Dao (心道) - Tâm Đạo
    DaoTypeHeaven    = "heaven_dao"    // Heaven Dao (天道) - Thiên Đạo
    DaoTypeEarth     = "earth_dao"     // Earth Dao (地道) - Địa Đạo
    DaoTypeHuman     = "human_dao"     // Human Dao (人道) - Nhân Đạo
)

// All Dao Types (所有道类型) - Tất cả loại Đạo
var AllDaoTypes = []string{
    DaoTypeSword, DaoTypeBlade, DaoTypeFist, DaoTypeSpell, DaoTypeFormation,
    DaoTypeRefinement, DaoTypeHeart, DaoTypeHeaven, DaoTypeEarth, DaoTypeHuman,
}
```

## 🧠 Dao Comprehension - Sự Hiểu Biết Đạo

```go
// Dao Comprehension (道理解) - Sự hiểu biết về Đạo
type DaoComprehension struct {
    // Basic comprehension (基础理解) - Hiểu biết cơ bản
    BasicUnderstanding float64 `json:"basic_understanding"` // Basic understanding (基础理解) - Hiểu biết cơ bản
    IntermediateUnderstanding float64 `json:"intermediate_understanding"` // Intermediate understanding (中级理解) - Hiểu biết trung cấp
    AdvancedUnderstanding float64 `json:"advanced_understanding"` // Advanced understanding (高级理解) - Hiểu biết nâng cao
    
    // Profound comprehension (深刻理解) - Hiểu biết sâu sắc
    ProfoundInsight    float64 `json:"profound_insight"`    // Profound insight (深刻洞察) - Sự thấu hiểu sâu sắc
    TranscendentWisdom float64 `json:"transcendent_wisdom"` // Transcendent wisdom (超脱智慧) - Trí tuệ siêu thoát
    DaoEssence        float64 `json:"dao_essence"`          // Dao essence (道本质) - Bản chất Đạo
}
```

## ⚔️ Dao Techniques - Kỹ Thuật Đạo

```go
// Dao Technique (道术) - Kỹ thuật Đạo
type DaoTechnique struct {
    // Basic info (基本信息) - Thông tin cơ bản
    TechniqueName    string  `json:"technique_name"`    // Technique name (术名称) - Tên kỹ thuật
    TechniqueType    string  `json:"technique_type"`    // Technique type (术类型) - Loại kỹ thuật
    TechniqueLevel   int     `json:"technique_level"`   // Technique level (术等级) - Cấp độ kỹ thuật
    TechniquePower   float64 `json:"technique_power"`   // Technique power (术威力) - Sức mạnh kỹ thuật
    
    // Advanced properties (高级属性) - Thuộc tính nâng cao
    TechniqueCost    float64 `json:"technique_cost"`    // Technique cost (术消耗) - Chi phí kỹ thuật
    TechniqueCooldown float64 `json:"technique_cooldown"` // Technique cooldown (术冷却) - Thời gian hồi chiêu
    TechniqueRange   float64 `json:"technique_range"`   // Technique range (术范围) - Phạm vi kỹ thuật
    TechniqueAccuracy float64 `json:"technique_accuracy"` // Technique accuracy (术精度) - Độ chính xác kỹ thuật
}
```

## 🔗 Dao Compatibility - Tương Thích Đạo

```go
// Dao Compatibility (道兼容性) - Tương thích Đạo
type DaoCompatibility struct {
    // Compatibility matrix (兼容性矩阵) - Ma trận tương thích
    CompatibilityMatrix map[string]map[string]float64 `json:"compatibility_matrix"`
    
    // Resonance effects (共振效果) - Hiệu ứng cộng hưởng
    ResonanceEffects map[string]float64 `json:"resonance_effects"`
    
    // Suppression effects (压制效果) - Hiệu ứng áp chế
    SuppressionEffects map[string]float64 `json:"suppression_effects"`
}
```

## 📊 Primary Stats từ Đạo

```go
// Dao-related primary stats (道相关主要属性) - Các primary stats từ Đạo
var DaoPrimaryStats = []string{
    // Core Dao stats (核心道属性) - Stats cốt lõi Đạo
    "dao_level",              // Dao level (道等级) - Cấp độ Đạo
    "dao_experience",         // Dao experience (道经验) - Kinh nghiệm Đạo
    "dao_mastery",           // Dao mastery (道掌握) - Mức độ thành thạo Đạo
    "dao_insight",           // Dao insight (道洞察) - Sự thấu hiểu Đạo
    "dao_wisdom",            // Dao wisdom (道智慧) - Trí tuệ Đạo
    
    // Advanced Dao stats (高级道属性) - Stats nâng cao Đạo
    "dao_harmony",           // Dao harmony (道和谐) - Sự hài hòa Đạo
    "dao_transcendence",     // Dao transcendence (道超脱) - Siêu thoát Đạo
    "dao_compatibility",     // Dao compatibility (道兼容性) - Tương thích Đạo
    "dao_suppression",       // Dao suppression (道压制) - Áp chế Đạo
    "dao_resonance",         // Dao resonance (道共振) - Cộng hưởng Đạo
    
    // Multi-Dao stats (多道属性) - Stats đa Đạo
    "total_dao_level",       // Total Dao level (总道等级) - Tổng cấp độ Đạo
    "total_dao_experience",  // Total Dao experience (总道经验) - Tổng kinh nghiệm Đạo
    "total_dao_mastery",     // Total Dao mastery (总道掌握) - Tổng mức độ thành thạo Đạo
    "active_dao_count",      // Active Dao count (激活道数量) - Số lượng Đạo đang hoạt động
    "average_dao_level",     // Average Dao level (平均道等级) - Cấp độ Đạo trung bình
    
    // Dao technique stats (道术属性) - Stats kỹ thuật Đạo
    "dao_technique_power",   // Dao technique power (道术威力) - Sức mạnh kỹ thuật Đạo
    "dao_technique_cost",    // Dao technique cost (道术消耗) - Chi phí kỹ thuật Đạo
    "dao_technique_cooldown", // Dao technique cooldown (道术冷却) - Thời gian hồi chiêu kỹ thuật Đạo
    "dao_technique_range",   // Dao technique range (道术范围) - Phạm vi kỹ thuật Đạo
    "dao_technique_accuracy", // Dao technique accuracy (道术精度) - Độ chính xác kỹ thuật Đạo
}
```

## 💡 Đặc Điểm Nổi Bật

### **Map-Based Design:**
- **Scalable**: Dễ dàng thêm loại Đạo mới
- **Flexible**: Quản lý nhiều Đạo cùng lúc
- **Performance**: Map lookup O(1)
- **Maintainable**: Code dễ bảo trì

### **Status Management:**
- **Active/Inactive**: Bật/tắt Đạo
- **Primary/Secondary**: Đạo chính/phụ
- **Unlock Tracking**: Theo dõi thời điểm mở khóa

### **Combat Integration:**
- **Technique Management**: Quản lý kỹ thuật theo Đạo
- **Compatibility Matrix**: Ma trận tương thích
- **Resonance Effects**: Hiệu ứng cộng hưởng
