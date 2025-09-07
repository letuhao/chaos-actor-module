# 01 — Dantian System (Hệ Thống Đan Điền)

**Generated:** 2025-01-27  
**Based on:** Tinh-Khí-Thần system with Dantian as Qi storage center

## Tổng quan

**Đan Điền (Dantian)** là trung tâm lưu trữ và xử lý khí, ảnh hưởng đến các primary stats liên quan đến **Khí** trong hệ thống Tinh-Khí-Thần.

## 🎯 Mục tiêu

- **Qi Storage**: Trung tâm lưu trữ khí
- **Qi Processing**: Xử lý và chuyển hóa khí
- **Cultivation Center**: Trung tâm tu luyện
- **Energy Management**: Quản lý năng lượng

## 🏗️ Cấu Trúc Dữ Liệu

### **DantianInfo - Thông Tin Đan Điền**

```go
// Dantian Info (丹田信息) - Thông tin đan điền
type DantianInfo struct {
    // Core dantian (核心丹田) - Đan điền cốt lõi
    Core DantianCore `json:"core"`
    
    // Storage system (存储系统) - Hệ thống lưu trữ
    Storage DantianStorage `json:"storage"`
    
    // Quality metrics (质量指标) - Chỉ số chất lượng
    Quality DantianQuality `json:"quality"`
}
```

### **DantianCore - Lõi Đan Điền**

```go
// Dantian Core (丹田核心) - Lõi đan điền
type DantianCore struct {
    // Basic capacity (基础容量) - Dung lượng cơ bản
    BaseCapacity    float64 `json:"base_capacity"`    // Base capacity (基础容量) - Dung lượng cơ bản
    CurrentCapacity float64 `json:"current_capacity"` // Current capacity (当前容量) - Dung lượng hiện tại
    MaxCapacity     float64 `json:"max_capacity"`     // Maximum capacity (最大容量) - Dung lượng tối đa
    
    // Compression ratio (压缩比) - Tỷ lệ nén
    CompressionRatio float64 `json:"compression_ratio"` // Compression ratio (压缩比) - Tỷ lệ nén
    CompressionLevel int     `json:"compression_level"` // Compression level (压缩等级) - Cấp độ nén
    
    // Stability metrics (稳定性指标) - Chỉ số ổn định
    Stability       float64 `json:"stability"`        // Stability (稳定性) - Độ ổn định
    PurityLevel     float64 `json:"purity_level"`     // Purity level (纯度等级) - Cấp độ tinh khiết
    Efficiency      float64 `json:"efficiency"`       // Efficiency (效率) - Hiệu suất
    
    // Cultivation progress (修炼进度) - Tiến độ tu luyện
    CultivationLevel int     `json:"cultivation_level"` // Cultivation level (修炼等级) - Cấp độ tu luyện
    BreakthroughProgress float64 `json:"breakthrough_progress"` // Breakthrough progress (突破进度) - Tiến độ đột phá
}
```

### **DantianStorage - Hệ Thống Lưu Trữ**

```go
// Dantian Storage (丹田存储) - Hệ thống lưu trữ đan điền
type DantianStorage struct {
    // Qi storage (气存储) - Lưu trữ khí
    QiCapacity      float64 `json:"qi_capacity"`      // Qi capacity (气容量) - Dung lượng khí
    QiCurrent       float64 `json:"qi_current"`       // Current Qi (当前气) - Khí hiện tại
    QiRegeneration  float64 `json:"qi_regeneration"`  // Qi regeneration (气恢复) - Tốc độ hồi khí
    QiConsumption   float64 `json:"qi_consumption"`   // Qi consumption (气消耗) - Tốc độ tiêu hao khí
    
    // Tu Vi storage (修为存储) - Lưu trữ tu vi
    TuViCapacity    float64 `json:"tu_vi_capacity"`   // Tu Vi capacity (修为容量) - Dung lượng tu vi
    TuViCurrent     float64 `json:"tu_vi_current"`    // Current Tu Vi (当前修为) - Tu vi hiện tại
    TuViRegeneration float64 `json:"tu_vi_regeneration"` // Tu Vi regeneration (修为恢复) - Tốc độ hồi tu vi
    
    // Energy types (能量类型) - Các loại năng lượng
    LingQiCapacity  float64 `json:"ling_qi_capacity"` // Ling Qi capacity (灵气容量) - Dung lượng linh khí
    XianQiCapacity  float64 `json:"xian_qi_capacity"` // Xian Qi capacity (仙气容量) - Dung lượng tiên khí
    ShenQiCapacity  float64 `json:"shen_qi_capacity"` // Shen Qi capacity (神气容量) - Dung lượng thần khí
}
```

### **DantianQuality - Chất Lượng Đan Điền**

```go
// Dantian Quality (丹田质量) - Chất lượng đan điền
type DantianQuality struct {
    // Purity metrics (纯度指标) - Chỉ số tinh khiết
    QiPurity        float64 `json:"qi_purity"`        // Qi purity (气纯度) - Độ tinh khiết khí
    DantianPurity   float64 `json:"dantian_purity"`   // Dantian purity (丹田纯度) - Độ tinh khiết đan điền
    EnergyPurity    float64 `json:"energy_purity"`    // Energy purity (能量纯度) - Độ tinh khiết năng lượng
    
    // Efficiency metrics (效率指标) - Chỉ số hiệu suất
    ConversionEfficiency float64 `json:"conversion_efficiency"` // Conversion efficiency (转换效率) - Hiệu suất chuyển đổi
    StorageEfficiency   float64 `json:"storage_efficiency"`   // Storage efficiency (存储效率) - Hiệu suất lưu trữ
    ProcessingEfficiency float64 `json:"processing_efficiency"` // Processing efficiency (处理效率) - Hiệu suất xử lý
    
    // Stability metrics (稳定性指标) - Chỉ số ổn định
    CoreStability   float64 `json:"core_stability"`   // Core stability (核心稳定性) - Độ ổn định lõi
    EnergyStability float64 `json:"energy_stability"` // Energy stability (能量稳定性) - Độ ổn định năng lượng
    FlowStability   float64 `json:"flow_stability"`   // Flow stability (流动稳定性) - Độ ổn định dòng chảy
}
```

## 📊 Primary Stats từ Đan Điền

```go
// Dantian-related primary stats (丹田相关主要属性) - Các primary stats từ đan điền
var DantianPrimaryStats = []string{
    "dantian_capacity",        // Dantian capacity (丹田容量) - Dung lượng đan điền
    "dantian_compression",     // Dantian compression (丹田压缩) - Suất nén khí trong đan điền
    "dantian_stability",       // Dantian stability (丹田稳定性) - Độ ổn định đan điền
    "dantian_purity",          // Dantian purity (丹田纯度) - Độ tinh khiết đan điền
    "qi_purity",              // Qi purity (气纯度) - Độ tinh khiết khí
    "eff_jing_to_qi",         // Jing to Qi efficiency (精转气效率) - Hiệu suất chuyển Tinh → Khí
    "eff_qi_to_shen",         // Qi to Shen efficiency (气转神效率) - Hiệu suất chuyển Khí → Thần
    "cultivation_exp",         // Main cultivation experience (主要修炼经验) - Tu vi chính
    "aux_cultivation_exp",     // Auxiliary cultivation experience (辅助修炼经验) - Tu vi phụ
    "total_cultivation_exp",   // Total cultivation experience (总修炼经验) - Tu vi tổng hợp
}
```

## 🧮 Công Thức Tính Toán

### **CalculateDantianCapacity - Tính Dung Lượng Đan Điền**

```go
// Calculate dantian capacity (计算丹田容量) - Tính dung lượng đan điền
func CalculateDantianCapacity(dantian DantianInfo) float64 {
    baseCapacity := dantian.Core.BaseCapacity
    compressionMultiplier := 1.0 + (float64(dantian.Core.CompressionLevel) * 0.1)
    qualityMultiplier := 1.0 + dantian.Quality.StorageEfficiency
    
    return baseCapacity * compressionMultiplier * qualityMultiplier
}
```

### **CalculateDantianLevel - Tính Cấp Độ Đan Điền**

```go
// Calculate dantian level (计算丹田等级) - Tính cấp độ đan điền
func CalculateDantianLevel(dantian DantianInfo) int {
    capacity := CalculateDantianCapacity(dantian)
    
    // Level calculation based on capacity (基于容量的等级计算) - Tính cấp độ dựa trên dung lượng
    if capacity < 100 {
        return 1
    } else if capacity < 500 {
        return 2
    } else if capacity < 1000 {
        return 3
    } else if capacity < 5000 {
        return 4
    } else {
        return 5
    }
}
```

### **DetermineSubstage - Xác Định Tiểu Cảnh Giới**

```go
// Determine substage based on dantian level (根据丹田等级确定子阶段) - Xác định tiểu cảnh giới dựa trên cấp độ đan điền
func DetermineSubstage(dantian DantianInfo, realm string) string {
    level := CalculateDantianLevel(dantian)
    progress := dantian.Core.BreakthroughProgress
    
    switch realm {
    case "Trúc Cơ":
        if progress < 0.25 {
            return "Sơ kỳ"
        } else if progress < 0.5 {
            return "Trung kỳ"
        } else if progress < 0.75 {
            return "Hậu kỳ"
        } else {
            return "Viên mãn"
        }
    case "Kim Đan":
        if progress < 0.25 {
            return "Sơ kỳ"
        } else if progress < 0.5 {
            return "Trung kỳ"
        } else if progress < 0.75 {
            return "Hậu kỳ"
        } else {
            return "Viên mãn"
        }
    default:
        return "Unknown"
    }
}
```

## 🔗 Tích Hợp Với Actor Core v3

### **ConvertToSubsystemOutput - Chuyển Đổi Thành SubsystemOutput**

```go
// Convert dantian to SubsystemOutput (转换丹田到子系统输出) - Chuyển đổi đan điền thành SubsystemOutput
func ConvertDantianToSubsystemOutput(dantian DantianInfo) SubsystemOutput {
    // Calculate primary stats (计算主要属性) - Tính toán primary stats
    primaryStats := make(map[string]float64)
    
    primaryStats["dantian_capacity"] = CalculateDantianCapacity(dantian)
    primaryStats["dantian_compression"] = dantian.Core.CompressionRatio
    primaryStats["dantian_stability"] = dantian.Core.Stability
    primaryStats["dantian_purity"] = dantian.Quality.DantianPurity
    primaryStats["qi_purity"] = dantian.Quality.QiPurity
    primaryStats["eff_jing_to_qi"] = dantian.Quality.ConversionEfficiency
    primaryStats["eff_qi_to_shen"] = dantian.Quality.ProcessingEfficiency
    primaryStats["cultivation_exp"] = dantian.Storage.TuViCurrent
    primaryStats["aux_cultivation_exp"] = dantian.Storage.TuViCurrent * 0.1
    primaryStats["total_cultivation_exp"] = dantian.Storage.TuViCurrent * 1.1
    
    // Calculate derived stats (计算派生属性) - Tính toán derived stats
    derivedStats := make(map[string]float64)
    
    // Energy calculations (能量计算) - Tính toán năng lượng
    derivedStats["qi_max"] = dantian.Storage.QiCapacity
    derivedStats["qi_current"] = dantian.Storage.QiCurrent
    derivedStats["qi_regen"] = dantian.Storage.QiRegeneration
    
    derivedStats["tu_vi_max"] = dantian.Storage.TuViCapacity
    derivedStats["tu_vi_current"] = dantian.Storage.TuViCurrent
    derivedStats["tu_vi_regen"] = dantian.Storage.TuViRegeneration
    
    // Create contributions (创建贡献) - Tạo contributions
    var contributions []Contribution
    
    // Primary stats contributions (主要属性贡献) - Contributions primary stats
    for stat, value := range primaryStats {
        contributions = append(contributions, Contribution{
            Dimension: stat,
            Bucket:    "Primary",
            Value:     value,
            System:    "DantianSystem",
            Priority:  1,
        })
    }
    
    // Derived stats contributions (派生属性贡献) - Contributions derived stats
    for stat, value := range derivedStats {
        contributions = append(contributions, Contribution{
            Dimension: stat,
            Bucket:    "Derived",
            Value:     value,
            System:    "DantianSystem",
            Priority:  2,
        })
    }
    
    return SubsystemOutput{
        Primary:   primaryStats,
        Derived:   derivedStats,
        Caps:      make(map[string]float64),
        Context:   map[string]interface{}{"dantian": dantian},
        Contributions: contributions,
    }
}
```

## 💡 Đặc Điểm Nổi Bật

### **Qi Storage Center:**
- **Central Hub**: Trung tâm lưu trữ khí
- **Energy Management**: Quản lý năng lượng
- **Cultivation Focus**: Tập trung tu luyện

### **Flexible Energy System:**
- **Multiple Energy Types**: Nhiều loại năng lượng
- **Dynamic Capacity**: Dung lượng động
- **Efficiency Metrics**: Chỉ số hiệu suất

### **Quality Management:**
- **Purity Tracking**: Theo dõi độ tinh khiết
- **Stability Monitoring**: Giám sát độ ổn định
- **Efficiency Optimization**: Tối ưu hiệu suất
