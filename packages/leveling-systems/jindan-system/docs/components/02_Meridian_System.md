# 02 — Meridian System (Hệ Thống Kinh Mạch)

**Generated:** 2025-01-27  
**Based on:** Tinh-Khí-Thần system with Meridians as Qi circulation system

## Tổng quan

**Kinh Mạch (Meridians)** là hệ thống vận chuyển khí và tinh, ảnh hưởng đến các primary stats liên quan đến **Tinh** và lưu thông khí trong hệ thống Tinh-Khí-Thần.

## 🎯 Mục tiêu

- **Qi Circulation**: Hệ thống vận chuyển khí
- **Jing Management**: Quản lý Tinh (Essence)
- **Speed Enhancement**: Tăng cường tốc độ
- **Cultivation Support**: Hỗ trợ tu luyện

## 🏗️ Cấu Trúc Dữ Liệu

### **MeridianSystem - Hệ Thống Kinh Mạch**

```go
// Meridian System (经脉系统) - Hệ thống kinh mạch
type MeridianSystem struct {
    // Meridian quality (经脉质量) - Chất lượng kinh mạch
    MeridianQuality MeridianQuality `json:"meridian_quality"`
    
    // Meridian capacity (经脉容量) - Dung lượng kinh mạch
    MeridianCapacity MeridianCapacity `json:"meridian_capacity"`
    
    // Jing reserve (精储备) - Dự trữ Tinh
    JingReserve JingReserve `json:"jing_reserve"`
    
    // Meridian channels (经脉通道) - Các kênh kinh mạch
    MeridianChannels MeridianChannels `json:"meridian_channels"`
    
    // Meridian cultivation (经脉修炼) - Tu luyện kinh mạch
    MeridianCultivation MeridianCultivation `json:"meridian_cultivation"`
}
```

### **MeridianQuality - Chất Lượng Kinh Mạch**

```go
// Meridian Quality (经脉质量) - Chất lượng kinh mạch
type MeridianQuality struct {
    Conductivity      float64 `json:"conductivity"`       // 0.0 - 1.0 (传导性) - Độ dẫn khí
    Toughness         float64 `json:"toughness"`          // 0.0 - 1.0 (韧性) - Sức chịu áp lực khí
    PurityResistance  float64 `json:"purity_resistance"`  // 0.0 - 1.0 (净抗性) - Chống tạp chất bám lưu
    Flexibility       float64 `json:"flexibility"`        // 0.0 - 1.0 (柔韧性) - Độ linh hoạt kinh mạch
    Resilience        float64 `json:"resilience"`         // 0.0 - 1.0 (恢复力) - Khả năng phục hồi kinh mạch
}
```

### **MeridianCapacity - Dung Lượng Kinh Mạch**

```go
// Meridian Capacity (经脉容量) - Dung lượng kinh mạch
type MeridianCapacity struct {
    MaxFlowRate       float64 `json:"max_flow_rate"`       // Maximum flow rate (最大流速) - Tốc độ dòng chảy tối đa
    MaxPressure       float64 `json:"max_pressure"`        // Maximum pressure (最大压力) - Áp lực tối đa
    TotalCapacity     float64 `json:"total_capacity"`      // Total capacity (总容量) - Dung lượng tổng
    CurrentFlow       float64 `json:"current_flow"`        // Current flow (当前流量) - Dòng chảy hiện tại
    FlowEfficiency    float64 `json:"flow_efficiency"`     // Flow efficiency (流动效率) - Hiệu suất dòng chảy
    PressureTolerance float64 `json:"pressure_tolerance"`  // Pressure tolerance (压力耐受) - Khả năng chịu áp lực
}
```

### **JingReserve - Dự Trữ Tinh**

```go
// Jing Reserve (精储备) - Dự trữ Tinh
type JingReserve struct {
    InnateJing        float64 `json:"innate_jing"`         // Innate Jing (先天精) - Tinh tiên thiên
    AcquiredJing      float64 `json:"acquired_jing"`       // Acquired Jing (后天精) - Tinh hậu thiên
    TotalJing         float64 `json:"total_jing"`          // Total Jing (总精) - Tổng tinh
    JingRegeneration  float64 `json:"jing_regeneration"`   // Jing regeneration (精恢复) - Tốc độ hồi tinh
    JingConsumption   float64 `json:"jing_consumption"`    // Jing consumption (精消耗) - Tốc độ tiêu hao tinh
    JingEfficiency    float64 `json:"jing_efficiency"`     // Jing efficiency (精效率) - Hiệu suất sử dụng tinh
}
```

### **MeridianChannels - Các Kênh Kinh Mạch**

```go
// Meridian Channels (经脉通道) - Các kênh kinh mạch
type MeridianChannels struct {
    // Main channels (主脉) - Kênh chính
    MainChannels MainChannels `json:"main_channels"`
    
    // Secondary channels (支脉) - Kênh phụ
    SecondaryChannels SecondaryChannels `json:"secondary_channels"`
    
    // Micro channels (微脉) - Kênh vi mô
    MicroChannels MicroChannels `json:"micro_channels"`
}

// Main Channels (主脉) - Kênh chính
type MainChannels struct {
    HandTaiyinLung     MeridianChannel `json:"hand_taiyin_lung"`     // Hand Taiyin Lung (手太阴肺经) - Thủ Thái Âm Phế Kinh
    HandYangmingLarge  MeridianChannel `json:"hand_yangming_large"`  // Hand Yangming Large (手阳明大肠经) - Thủ Dương Minh Đại Tràng Kinh
    FootYangmingStomach MeridianChannel `json:"foot_yangming_stomach"` // Foot Yangming Stomach (足阳明胃经) - Túc Dương Minh Vị Kinh
    FootTaiyinSpleen   MeridianChannel `json:"foot_taiyin_spleen"`   // Foot Taiyin Spleen (足太阴脾经) - Túc Thái Âm Tỳ Kinh
    HandShaoyinHeart   MeridianChannel `json:"hand_shaoyin_heart"`   // Hand Shaoyin Heart (手少阴心经) - Thủ Thiếu Âm Tâm Kinh
    HandTaiyangSmall   MeridianChannel `json:"hand_taiyang_small"`   // Hand Taiyang Small (手太阳小肠经) - Thủ Thái Dương Tiểu Tràng Kinh
    FootTaiyangBladder MeridianChannel `json:"foot_taiyang_bladder"` // Foot Taiyang Bladder (足太阳膀胱经) - Túc Thái Dương Bàng Quang Kinh
    FootShaoyinKidney  MeridianChannel `json:"foot_shaoyin_kidney"`  // Foot Shaoyin Kidney (足少阴肾经) - Túc Thiếu Âm Thận Kinh
    HandJueyinPericardium MeridianChannel `json:"hand_jueyin_pericardium"` // Hand Jueyin Pericardium (手厥阴心包经) - Thủ Quyết Âm Tâm Bào Kinh
    HandShaoyangTriple MeridianChannel `json:"hand_shaoyang_triple"` // Hand Shaoyang Triple (手少阳三焦经) - Thủ Thiếu Dương Tam Tiêu Kinh
    FootShaoyangGallbladder MeridianChannel `json:"foot_shaoyang_gallbladder"` // Foot Shaoyang Gallbladder (足少阳胆经) - Túc Thiếu Dương Đởm Kinh
    FootJueyinLiver    MeridianChannel `json:"foot_jueyin_liver"`    // Foot Jueyin Liver (足厥阴肝经) - Túc Quyết Âm Can Kinh
}

// Secondary Channels (支脉) - Kênh phụ
type SecondaryChannels struct {
    CollateralChannels []MeridianChannel `json:"collateral_channels"` // Collateral channels (络脉) - Lạc mạch
    DivergentChannels  []MeridianChannel `json:"divergent_channels"`  // Divergent channels (别脉) - Biệt mạch
    MuscleChannels     []MeridianChannel `json:"muscle_channels"`     // Muscle channels (经筋) - Kinh cân
    SkinChannels       []MeridianChannel `json:"skin_channels"`       // Skin channels (皮部) - Bì bộ
}

// Micro Channels (微脉) - Kênh vi mô
type MicroChannels struct {
    CapillaryChannels  []MeridianChannel `json:"capillary_channels"`  // Capillary channels (毛细血管) - Mao tế huyết quản
    NerveChannels      []MeridianChannel `json:"nerve_channels"`      // Nerve channels (神经通道) - Thần kinh thông đạo
    EnergyChannels     []MeridianChannel `json:"energy_channels"`     // Energy channels (能量通道) - Năng lượng thông đạo
}

// Meridian Channel (经脉通道) - Kênh kinh mạch
type MeridianChannel struct {
    ChannelName        string  `json:"channel_name"`        // Channel name (通道名称) - Tên kênh
    ChannelType        string  `json:"channel_type"`        // Channel type (通道类型) - Loại kênh
    ChannelHealth      float64 `json:"channel_health"`      // Channel health (通道健康度) - Sức khỏe kênh
    ChannelFlow        float64 `json:"channel_flow"`        // Channel flow (通道流量) - Dòng chảy kênh
    ChannelPressure    float64 `json:"channel_pressure"`    // Channel pressure (通道压力) - Áp lực kênh
    ChannelResistance  float64 `json:"channel_resistance"`  // Channel resistance (通道阻力) - Điện trở kênh
    ChannelEfficiency  float64 `json:"channel_efficiency"`  // Channel efficiency (通道效率) - Hiệu suất kênh
}
```

### **MeridianCultivation - Tu Luyện Kinh Mạch**

```go
// Meridian Cultivation (经脉修炼) - Tu luyện kinh mạch
type MeridianCultivation struct {
    CultivationLevel   int     `json:"cultivation_level"`   // Cultivation level (修炼等级) - Cấp độ tu luyện
    BreakthroughProgress float64 `json:"breakthrough_progress"` // Breakthrough progress (突破进度) - Tiến độ đột phá
    TribulationResistance float64 `json:"tribulation_resistance"` // Tribulation resistance (劫难抗性) - Khả năng chống kiếp
    ChannelSynchronization float64 `json:"channel_synchronization"` // Channel synchronization (通道同步) - Đồng bộ kênh
    FlowOptimization   float64 `json:"flow_optimization"`   // Flow optimization (流动优化) - Tối ưu dòng chảy
}
```

## 📊 Primary Stats từ Kinh Mạch

```go
// Meridian-related primary stats (经脉相关主要属性) - Các primary stats từ kinh mạch
var MeridianPrimaryStats = []string{
    // Core meridian stats (核心经脉属性) - Stats cốt lõi kinh mạch
    "meridian_conductivity",      // Meridian conductivity (经脉传导性) - Độ dẫn khí trong kinh mạch
    "meridian_toughness",         // Meridian toughness (经脉韧性) - Sức chịu áp lực khí
    "meridian_purity_resistance", // Meridian purity resistance (经脉净抗性) - Chống tạp chất bám lưu
    "meridian_flexibility",       // Meridian flexibility (经脉柔韧性) - Độ linh hoạt kinh mạch
    "meridian_resilience",        // Meridian resilience (经脉恢复力) - Khả năng phục hồi kinh mạch
    
    // Capacity and flow stats (容量和流动属性) - Stats dung lượng và dòng chảy
    "meridian_capacity",          // Meridian capacity (经脉容量) - Dung lượng kinh mạch
    "meridian_flow_rate",         // Meridian flow rate (经脉流速) - Tốc độ dòng chảy
    "meridian_pressure",          // Meridian pressure (经脉压力) - Áp lực kinh mạch
    "meridian_efficiency",        // Meridian efficiency (经脉效率) - Hiệu suất vận chuyển khí
    "meridian_flow_efficiency",   // Meridian flow efficiency (经脉流动效率) - Hiệu suất dòng chảy
    
    // Jing-related stats (精相关属性) - Stats liên quan đến Tinh
    "jing_reserve",              // Jing reserve (精储备) - Dự trữ Tinh (nguyên khí)
    "jing_regeneration",         // Jing regeneration (精恢复) - Tốc độ hồi tinh
    "jing_consumption",          // Jing consumption (精消耗) - Tốc độ tiêu hao tinh
    "jing_efficiency",           // Jing efficiency (精效率) - Hiệu suất sử dụng tinh
    
    // Channel-specific stats (通道特定属性) - Stats đặc thù kênh
    "main_channel_health",       // Main channel health (主脉健康度) - Sức khỏe kênh chính
    "secondary_channel_health",  // Secondary channel health (支脉健康度) - Sức khỏe kênh phụ
    "micro_channel_health",      // Micro channel health (微脉健康度) - Sức khỏe kênh vi mô
    "channel_balance",           // Channel balance (通道平衡) - Cân bằng kênh
    
    // Cultivation stats (修炼属性) - Stats tu luyện
    "meridian_cultivation_level", // Meridian cultivation level (经脉修炼等级) - Cấp độ tu luyện kinh mạch
    "meridian_breakthrough_progress", // Meridian breakthrough progress (经脉突破进度) - Tiến độ đột phá kinh mạch
    "meridian_tribulation_resistance", // Meridian tribulation resistance (经脉劫难抗性) - Khả năng chống kiếp kinh mạch
    "channel_synchronization",   // Channel synchronization (通道同步) - Đồng bộ kênh
    "meridian_stability",        // Meridian stability (经脉稳定性) - Độ ổn định kinh mạch
    "meridian_cultivation_progress", // Meridian cultivation progress (经脉修炼进度) - Tiến độ tu luyện kinh mạch
}
```

## 🧮 Công Thức Tính Toán

### **CalculateMeridianStats - Tính Stats Kinh Mạch**

```go
// Calculate meridian stats (计算经脉属性) - Tính toán các stats của kinh mạch
func CalculateMeridianStats(meridian MeridianSystem) map[string]float64 {
    stats := make(map[string]float64)
    
    // Core meridian stats (核心经脉属性) - Stats cốt lõi kinh mạch
    stats["meridian_conductivity"] = meridian.MeridianQuality.Conductivity
    stats["meridian_toughness"] = meridian.MeridianQuality.Toughness
    stats["meridian_purity_resistance"] = meridian.MeridianQuality.PurityResistance
    stats["meridian_flexibility"] = meridian.MeridianQuality.Flexibility
    stats["meridian_resilience"] = meridian.MeridianQuality.Resilience
    
    // Capacity and flow stats (容量和流动属性) - Stats dung lượng và dòng chảy
    stats["meridian_capacity"] = meridian.MeridianCapacity.TotalCapacity
    stats["meridian_flow_rate"] = meridian.MeridianCapacity.MaxFlowRate
    stats["meridian_pressure"] = meridian.MeridianCapacity.MaxPressure
    stats["meridian_efficiency"] = meridian.MeridianCapacity.FlowEfficiency
    stats["meridian_flow_efficiency"] = meridian.MeridianCapacity.FlowEfficiency
    
    // Jing-related stats (精相关属性) - Stats liên quan đến Tinh
    stats["jing_reserve"] = meridian.JingReserve.TotalJing
    stats["jing_regeneration"] = meridian.JingReserve.JingRegeneration
    stats["jing_consumption"] = meridian.JingReserve.JingConsumption
    stats["jing_efficiency"] = meridian.JingReserve.JingEfficiency
    
    // Channel-specific stats (通道特定属性) - Stats đặc thù kênh
    stats["main_channel_health"] = CalculateMainChannelHealth(meridian.MeridianChannels.MainChannels)
    stats["secondary_channel_health"] = CalculateSecondaryChannelHealth(meridian.MeridianChannels.SecondaryChannels)
    stats["micro_channel_health"] = CalculateMicroChannelHealth(meridian.MeridianChannels.MicroChannels)
    stats["channel_balance"] = CalculateChannelBalance(meridian.MeridianChannels)
    
    // Cultivation stats (修炼属性) - Stats tu luyện
    stats["meridian_cultivation_level"] = float64(meridian.MeridianCultivation.CultivationLevel)
    stats["meridian_breakthrough_progress"] = meridian.MeridianCultivation.BreakthroughProgress
    stats["meridian_tribulation_resistance"] = meridian.MeridianCultivation.TribulationResistance
    stats["channel_synchronization"] = meridian.MeridianCultivation.ChannelSynchronization
    stats["meridian_stability"] = CalculateMeridianStability(meridian)
    stats["meridian_cultivation_progress"] = CalculateMeridianCultivationProgress(meridian)
    
    return stats
}
```

### **CalculateCombatSpeed - Tính Tốc Độ Combat**

```go
// Calculate combat speed from meridians (根据经脉计算战斗速度) - Tính tốc độ combat từ kinh mạch
func CalculateCombatSpeed(meridian MeridianSystem) float64 {
    baseSpeed := 100.0
    flowMultiplier := 1.0 + (meridian.MeridianCapacity.FlowEfficiency * 0.5)
    conductivityMultiplier := 1.0 + (meridian.MeridianQuality.Conductivity * 0.3)
    flexibilityMultiplier := 1.0 + (meridian.MeridianQuality.Flexibility * 0.2)
    
    return baseSpeed * flowMultiplier * conductivityMultiplier * flexibilityMultiplier
}
```

### **CalculateCultivationSpeed - Tính Tốc Độ Tu Luyện**

```go
// Calculate cultivation speed from meridians (根据经脉计算修炼速度) - Tính tốc độ tu luyện từ kinh mạch
func CalculateCultivationSpeed(meridian MeridianSystem) float64 {
    baseSpeed := 1.0
    jingEfficiency := meridian.JingReserve.JingEfficiency
    channelSync := meridian.MeridianCultivation.ChannelSynchronization
    flowOpt := meridian.MeridianCultivation.FlowOptimization
    
    return baseSpeed * jingEfficiency * channelSync * flowOpt
}
```

## 🔗 Tích Hợp Với Actor Core v3

### **ConvertToSubsystemOutput - Chuyển Đổi Thành SubsystemOutput**

```go
// Convert meridian to SubsystemOutput (转换经脉到子系统输出) - Chuyển đổi kinh mạch thành SubsystemOutput
func ConvertMeridianToSubsystemOutput(meridian MeridianSystem) SubsystemOutput {
    // Calculate primary stats (计算主要属性) - Tính toán primary stats
    primaryStats := CalculateMeridianStats(meridian)
    
    // Calculate derived stats (计算派生属性) - Tính toán derived stats
    derivedStats := make(map[string]float64)
    
    // Speed calculations (速度计算) - Tính toán tốc độ
    derivedStats["combat_speed"] = CalculateCombatSpeed(meridian)
    derivedStats["spell_casting_speed"] = CalculateSpellCastingSpeed(meridian)
    derivedStats["movement_speed"] = CalculateMovementSpeed(meridian)
    derivedStats["cultivation_speed"] = CalculateCultivationSpeed(meridian)
    
    // Create contributions (创建贡献) - Tạo contributions
    var contributions []Contribution
    
    // Primary stats contributions (主要属性贡献) - Contributions primary stats
    for stat, value := range primaryStats {
        contributions = append(contributions, Contribution{
            Dimension: stat,
            Bucket:    "Primary",
            Value:     value,
            System:    "MeridianSystem",
            Priority:  1,
        })
    }
    
    // Derived stats contributions (派生属性贡献) - Contributions derived stats
    for stat, value := range derivedStats {
        contributions = append(contributions, Contribution{
            Dimension: stat,
            Bucket:    "Derived",
            Value:     value,
            System:    "MeridianSystem",
            Priority:  2,
        })
    }
    
    return SubsystemOutput{
        Primary:   primaryStats,
        Derived:   derivedStats,
        Caps:      make(map[string]float64),
        Context:   map[string]interface{}{"meridian": meridian},
        Contributions: contributions,
    }
}
```

## 💡 Đặc Điểm Nổi Bật

### **Qi Circulation System:**
- **Main Channels**: 12 kênh chính theo y học cổ truyền
- **Secondary Channels**: Kênh phụ và lạc mạch
- **Micro Channels**: Kênh vi mô cho chi tiết

### **Jing Management:**
- **Innate/Acquired Jing**: Tinh tiên thiên và hậu thiên
- **Efficiency Tracking**: Theo dõi hiệu suất sử dụng
- **Regeneration System**: Hệ thống hồi phục

### **Speed Enhancement:**
- **Combat Speed**: Tốc độ chiến đấu
- **Cultivation Speed**: Tốc độ tu luyện
- **Spell Casting**: Tốc độ thi triển pháp thuật
