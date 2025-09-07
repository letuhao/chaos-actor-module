# 04 — Profound Meaning System (Hệ Thống Áo Nghĩa)

**Generated:** 2025-01-27  
**Based on:** Map-based approach for scalable Profound Meaning management

## Tổng quan

Hệ thống **Áo Nghĩa (Profound Meaning)** là sự hiểu biết sâu sắc về bản chất của Đạo, ảnh hưởng trực tiếp đến combat effectiveness và khả năng đặc biệt.

## 🎯 Mục tiêu

- **Deep Understanding**: Sự hiểu biết sâu sắc về Đạo
- **Combat Enhancement**: Tăng cường hiệu quả chiến đấu
- **Special Abilities**: Mở khóa khả năng đặc biệt
- **Turn-Based Integration**: Tích hợp với combat theo lượt

## 📚 Các Loại Áo Nghĩa

### **Vũ Khí & Võ Thuật:**
- **Kiếm Ý (Sword Intent)** - Ý niệm kiếm thuật
- **Đao Ý (Blade Intent)** - Ý niệm đao thuật
- **Quyền Ý (Fist Intent)** - Ý niệm quyền pháp

### **Pháp Thuật & Trận Pháp:**
- **Pháp Ý (Spell Intent)** - Ý niệm pháp thuật
- **Trận Ý (Formation Intent)** - Ý niệm trận pháp

### **Luyện Khí & Tâm Cảnh:**
- **Luyện Ý (Refinement Intent)** - Ý niệm luyện khí
- **Tâm Ý (Heart Intent)** - Ý niệm tâm cảnh

### **Tam Tài Ý:**
- **Thiên Ý (Heaven Intent)** - Ý niệm thiên đạo
- **Địa Ý (Earth Intent)** - Ý niệm địa đạo
- **Nhân Ý (Human Intent)** - Ý niệm nhân đạo

## 🏆 Cấp Độ Áo Nghĩa

### **1. Mơ Hồ (Vague)**
- **Tầng**: 1-2
- **Mô tả**: Hiểu biết mơ hồ về Áo Nghĩa
- **Đặc điểm**: Khả năng cơ bản, hiệu quả thấp

### **2. Rõ Ràng (Clear)**
- **Tầng**: 3-4
- **Mô tả**: Hiểu biết rõ ràng về Áo Nghĩa
- **Đặc điểm**: Khả năng cải thiện, hiệu quả trung bình

### **3. Sâu Sắc (Profound)**
- **Tầng**: 5-6
- **Mô tả**: Hiểu biết sâu sắc về Áo Nghĩa
- **Đặc điểm**: Khả năng mạnh, hiệu quả cao

### **4. Tinh Thông (Mastery)**
- **Tầng**: 7-8
- **Mô tả**: Thành thạo Áo Nghĩa
- **Đặc điểm**: Khả năng đặc biệt, hiệu quả rất cao

### **5. Siêu Phàm (Transcendent)**
- **Tầng**: 9+
- **Mô tả**: Vượt trội trong Áo Nghĩa
- **Đặc điểm**: Khả năng siêu phàm, hiệu quả tối đa

## 🏗️ Cấu Trúc Dữ Liệu

### **ProfoundMeaningSystem - Hệ Thống Áo Nghĩa Chính**

```go
// Profound Meaning System (奥义系统) - Hệ thống Áo Nghĩa
type ProfoundMeaningSystem struct {
    // Profound Meanings Map (奥义映射) - Map các Áo Nghĩa
    // Key: ProfoundMeaningType string, Value: ProfoundMeaning
    ProfoundMeanings map[string]ProfoundMeaning `json:"profound_meanings"`
    
    // Primary Profound Meaning Type (主奥义类型) - Loại Áo Nghĩa chính
    PrimaryProfoundMeaningType string `json:"primary_profound_meaning_type"`
    
    // Profound Understanding (奥义理解) - Sự hiểu biết Áo Nghĩa
    ProfoundUnderstanding ProfoundUnderstanding `json:"profound_understanding"`
    
    // Profound Abilities Map (奥义能力映射) - Map khả năng Áo Nghĩa
    // Key: ProfoundMeaningType string, Value: []ProfoundAbility
    ProfoundAbilities map[string][]ProfoundAbility `json:"profound_abilities"`
    
    // Profound Resonance (奥义共振) - Cộng hưởng Áo Nghĩa
    ProfoundResonance ProfoundResonance `json:"profound_resonance"`
}
```

### **ProfoundMeaning - Áo Nghĩa**

```go
// Profound Meaning (奥义) - Áo Nghĩa
type ProfoundMeaning struct {
    // Basic info (基本信息) - Thông tin cơ bản
    MeaningType      string  `json:"meaning_type"`      // Type of Profound Meaning (奥义类型) - Loại Áo Nghĩa
    MeaningLevel     int     `json:"meaning_level"`     // Profound Meaning level (奥义等级) - Cấp độ Áo Nghĩa
    MeaningExperience float64 `json:"meaning_experience"` // Profound Meaning experience (奥义经验) - Kinh nghiệm Áo Nghĩa
    MeaningMastery   float64 `json:"meaning_mastery"`    // Profound Meaning mastery (奥义掌握) - Mức độ thành thạo
    
    // Advanced properties (高级属性) - Thuộc tính nâng cao
    MeaningInsight   float64 `json:"meaning_insight"`   // Profound Meaning insight (奥义洞察) - Sự thấu hiểu Áo Nghĩa
    MeaningWisdom    float64 `json:"meaning_wisdom"`    // Profound Meaning wisdom (奥义智慧) - Trí tuệ Áo Nghĩa
    MeaningHarmony   float64 `json:"meaning_harmony"`   // Profound Meaning harmony (奥义和谐) - Sự hài hòa Áo Nghĩa
    MeaningTranscendence float64 `json:"meaning_transcendence"` // Profound Meaning transcendence (奥义超脱) - Siêu thoát Áo Nghĩa
    
    // Status flags (状态标志) - Cờ trạng thái
    IsActive         bool    `json:"is_active"`         // Is this Profound Meaning active (是否激活) - Áo Nghĩa này có đang hoạt động không
    IsPrimary        bool    `json:"is_primary"`        // Is this the primary Profound Meaning (是否主奥义) - Có phải Áo Nghĩa chính không
    UnlockedAt       int64   `json:"unlocked_at"`       // When this Profound Meaning was unlocked (解锁时间) - Thời điểm mở khóa Áo Nghĩa này
}
```

## 🏷️ Định Nghĩa Loại Áo Nghĩa

### **Profound Meaning Type Constants**

```go
// Profound Meaning Type Constants (奥义类型常量) - Hằng số loại Áo Nghĩa
const (
    ProfoundMeaningTypeSwordIntent     = "sword_intent"     // Sword Intent (剑意) - Kiếm Ý
    ProfoundMeaningTypeBladeIntent     = "blade_intent"     // Blade Intent (刀意) - Đao Ý
    ProfoundMeaningTypeFistIntent      = "fist_intent"      // Fist Intent (拳意) - Quyền Ý
    ProfoundMeaningTypeSpellIntent     = "spell_intent"     // Spell Intent (法意) - Pháp Ý
    ProfoundMeaningTypeFormationIntent = "formation_intent" // Formation Intent (阵意) - Trận Ý
    ProfoundMeaningTypeRefinementIntent = "refinement_intent" // Refinement Intent (炼意) - Luyện Ý
    ProfoundMeaningTypeHeartIntent     = "heart_intent"     // Heart Intent (心意) - Tâm Ý
    ProfoundMeaningTypeHeavenIntent    = "heaven_intent"    // Heaven Intent (天意) - Thiên Ý
    ProfoundMeaningTypeEarthIntent     = "earth_intent"     // Earth Intent (地意) - Địa Ý
    ProfoundMeaningTypeHumanIntent     = "human_intent"     // Human Intent (人意) - Nhân Ý
)

// All Profound Meaning Types (所有奥义类型) - Tất cả loại Áo Nghĩa
var AllProfoundMeaningTypes = []string{
    ProfoundMeaningTypeSwordIntent, ProfoundMeaningTypeBladeIntent, ProfoundMeaningTypeFistIntent,
    ProfoundMeaningTypeSpellIntent, ProfoundMeaningTypeFormationIntent, ProfoundMeaningTypeRefinementIntent,
    ProfoundMeaningTypeHeartIntent, ProfoundMeaningTypeHeavenIntent, ProfoundMeaningTypeEarthIntent,
    ProfoundMeaningTypeHumanIntent,
}
```

## 🧠 Profound Understanding - Sự Hiểu Biết Áo Nghĩa

```go
// Profound Understanding (奥义理解) - Sự hiểu biết Áo Nghĩa
type ProfoundUnderstanding struct {
    // Basic understanding (基础理解) - Hiểu biết cơ bản
    BasicUnderstanding float64 `json:"basic_understanding"` // Basic understanding (基础理解) - Hiểu biết cơ bản
    IntermediateUnderstanding float64 `json:"intermediate_understanding"` // Intermediate understanding (中级理解) - Hiểu biết trung cấp
    AdvancedUnderstanding float64 `json:"advanced_understanding"` // Advanced understanding (高级理解) - Hiểu biết nâng cao
    
    // Profound understanding (深刻理解) - Hiểu biết sâu sắc
    ProfoundInsight    float64 `json:"profound_insight"`    // Profound insight (深刻洞察) - Sự thấu hiểu sâu sắc
    TranscendentWisdom float64 `json:"transcendent_wisdom"` // Transcendent wisdom (超脱智慧) - Trí tuệ siêu thoát
    MeaningEssence     float64 `json:"meaning_essence"`     // Profound Meaning essence (奥义本质) - Bản chất Áo Nghĩa
}
```

## ⚔️ Profound Abilities - Khả Năng Áo Nghĩa

```go
// Profound Ability (奥义能力) - Khả năng Áo Nghĩa
type ProfoundAbility struct {
    // Basic info (基本信息) - Thông tin cơ bản
    AbilityName    string  `json:"ability_name"`    // Ability name (能力名称) - Tên khả năng
    AbilityType    string  `json:"ability_type"`    // Ability type (能力类型) - Loại khả năng
    AbilityLevel   int     `json:"ability_level"`   // Ability level (能力等级) - Cấp độ khả năng
    AbilityPower   float64 `json:"ability_power"`   // Ability power (能力威力) - Sức mạnh khả năng
    
    // Advanced properties (高级属性) - Thuộc tính nâng cao
    AbilityCost    float64 `json:"ability_cost"`    // Ability cost (能力消耗) - Chi phí khả năng
    AbilityCooldown float64 `json:"ability_cooldown"` // Ability cooldown (能力冷却) - Thời gian hồi chiêu
    AbilityRange   float64 `json:"ability_range"`   // Ability range (能力范围) - Phạm vi khả năng
    AbilityAccuracy float64 `json:"ability_accuracy"` // Ability accuracy (能力精度) - Độ chính xác khả năng
}
```

## 🔗 Profound Resonance - Cộng Hưởng Áo Nghĩa

```go
// Profound Resonance (奥义共振) - Cộng hưởng Áo Nghĩa
type ProfoundResonance struct {
    // Resonance matrix (共振矩阵) - Ma trận cộng hưởng
    ResonanceMatrix map[string]map[string]float64 `json:"resonance_matrix"`
    
    // Resonance effects (共振效果) - Hiệu ứng cộng hưởng
    ResonanceEffects map[string]float64 `json:"resonance_effects"`
    
    // Overwhelm effects (压倒效果) - Hiệu ứng áp đảo
    OverwhelmEffects map[string]float64 `json:"overwhelm_effects"`
}
```

## 📊 Primary Stats từ Áo Nghĩa

```go
// Profound Meaning primary stats (奥义主要属性) - Các primary stats từ Áo Nghĩa
var ProfoundMeaningPrimaryStats = []string{
    // Core Profound Meaning stats (核心奥义属性) - Stats cốt lõi Áo Nghĩa
    "profound_meaning_level",     // Profound Meaning level (奥义等级) - Cấp độ Áo Nghĩa
    "profound_meaning_experience", // Profound Meaning experience (奥义经验) - Kinh nghiệm Áo Nghĩa
    "profound_meaning_mastery",   // Profound Meaning mastery (奥义掌握) - Mức độ thành thạo Áo Nghĩa
    "profound_meaning_insight",   // Profound Meaning insight (奥义洞察) - Sự thấu hiểu Áo Nghĩa
    "profound_meaning_wisdom",    // Profound Meaning wisdom (奥义智慧) - Trí tuệ Áo Nghĩa
    
    // Advanced Profound Meaning stats (高级奥义属性) - Stats nâng cao Áo Nghĩa
    "profound_meaning_harmony",   // Profound Meaning harmony (奥义和谐) - Sự hài hòa Áo Nghĩa
    "profound_meaning_transcendence", // Profound Meaning transcendence (奥义超脱) - Siêu thoát Áo Nghĩa
    "profound_meaning_resonance", // Profound Meaning resonance (奥义共振) - Cộng hưởng Áo Nghĩa
    "profound_meaning_overwhelm", // Profound Meaning overwhelm (奥义压倒) - Áp đảo Áo Nghĩa
    "profound_meaning_essence",   // Profound Meaning essence (奥义本质) - Bản chất Áo Nghĩa
    
    // Multi-Profound Meaning stats (多奥义属性) - Stats đa Áo Nghĩa
    "total_profound_meaning_level",     // Total Profound Meaning level (总奥义等级) - Tổng cấp độ Áo Nghĩa
    "total_profound_meaning_experience", // Total Profound Meaning experience (总奥义经验) - Tổng kinh nghiệm Áo Nghĩa
    "total_profound_meaning_mastery",   // Total Profound Meaning mastery (总奥义掌握) - Tổng mức độ thành thạo Áo Nghĩa
    "active_profound_meaning_count",    // Active Profound Meaning count (激活奥义数量) - Số lượng Áo Nghĩa đang hoạt động
    "average_profound_meaning_level",   // Average Profound Meaning level (平均奥义等级) - Cấp độ Áo Nghĩa trung bình
    
    // Profound Meaning ability stats (奥义能力属性) - Stats khả năng Áo Nghĩa
    "profound_ability_power",     // Profound Meaning ability power (奥义能力威力) - Sức mạnh khả năng Áo Nghĩa
    "profound_ability_cost",      // Profound Meaning ability cost (奥义能力消耗) - Chi phí khả năng Áo Nghĩa
    "profound_ability_cooldown",  // Profound Meaning ability cooldown (奥义能力冷却) - Thời gian hồi chiêu khả năng Áo Nghĩa
    "profound_ability_range",     // Profound Meaning ability range (奥义能力范围) - Phạm vi khả năng Áo Nghĩa
    "profound_ability_accuracy",  // Profound Meaning ability accuracy (奥义能力精度) - Độ chính xác khả năng Áo Nghĩa
}
```

## 💡 Đặc Điểm Nổi Bật

### **Map-Based Design:**
- **Scalable**: Dễ dàng thêm loại Áo Nghĩa mới
- **Flexible**: Quản lý nhiều Áo Nghĩa cùng lúc
- **Performance**: Map lookup O(1)
- **Maintainable**: Code dễ bảo trì

### **Combat Integration:**
- **Ability Management**: Quản lý khả năng theo Áo Nghĩa
- **Resonance Matrix**: Ma trận cộng hưởng
- **Overwhelm Effects**: Hiệu ứng áp đảo

### **Status Management:**
- **Active/Inactive**: Bật/tắt Áo Nghĩa
- **Primary/Secondary**: Áo Nghĩa chính/phụ
- **Unlock Tracking**: Theo dõi thời điểm mở khóa
