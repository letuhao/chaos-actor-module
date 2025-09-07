# 03 — Dao Management Functions (Các Hàm Quản Lý Đạo)

**Generated:** 2025-01-27  
**Based on:** Map-based approach for flexible Dao management

## Tổng quan

Các hàm quản lý Đạo cho phép thêm, xóa, chuyển đổi và quản lý trạng thái của các loại Đạo một cách linh hoạt.

## 🚀 Khởi Tạo Hệ Thống

### **InitializeDaoSystem - Khởi Tạo Hệ Thống Đạo**

```go
// Initialize Dao System (初始化道系统) - Khởi tạo hệ thống Đạo
func InitializeDaoSystem() DaoSystem {
    return DaoSystem{
        DaoPaths: make(map[string]DaoPath),
        PrimaryDaoType: "",
        DaoComprehension: DaoComprehension{},
        DaoTechniques: make(map[string][]DaoTechnique),
        DaoCompatibility: DaoCompatibility{
            CompatibilityMatrix: make(map[string]map[string]float64),
            ResonanceEffects: make(map[string]float64),
            SuppressionEffects: make(map[string]float64),
        },
    }
}
```

## ➕ Thêm & Xóa Đạo

### **AddDaoPath - Thêm Con Đường Đạo**

```go
// Add Dao Path (添加道径) - Thêm con đường Đạo
func (ds *DaoSystem) AddDaoPath(daoType string, daoPath DaoPath) {
    daoPath.DaoType = daoType
    daoPath.IsActive = true
    daoPath.UnlockedAt = time.Now().Unix()
    
    ds.DaoPaths[daoType] = daoPath
    
    // Set as primary if it's the first one (如果是第一个则设为主道) - Nếu là cái đầu tiên thì đặt làm chính
    if ds.PrimaryDaoType == "" {
        ds.PrimaryDaoType = daoType
        daoPath.IsPrimary = true
        ds.DaoPaths[daoType] = daoPath
    }
}
```

### **RemoveDaoPath - Xóa Con Đường Đạo**

```go
// Remove Dao Path (移除道径) - Xóa con đường Đạo
func (ds *DaoSystem) RemoveDaoPath(daoType string) {
    delete(ds.DaoPaths, daoType)
    
    // If removing primary Dao, set a new primary (如果移除主道则设置新的主道) - Nếu xóa Đạo chính thì đặt Đạo chính mới
    if ds.PrimaryDaoType == daoType {
        ds.PrimaryDaoType = ""
        for daoType, dao := range ds.DaoPaths {
            if dao.IsActive {
                ds.PrimaryDaoType = daoType
                dao.IsPrimary = true
                ds.DaoPaths[daoType] = dao
                break
            }
        }
    }
}
```

## 🎯 Quản Lý Đạo Chính

### **SetPrimaryDao - Đặt Đạo Chính**

```go
// Set Primary Dao (设置主道) - Đặt Đạo chính
func (ds *DaoSystem) SetPrimaryDao(daoType string) bool {
    if dao, exists := ds.DaoPaths[daoType]; exists && dao.IsActive {
        // Remove primary flag from current primary (移除当前主道的主道标志) - Xóa cờ chính khỏi Đạo chính hiện tại
        if ds.PrimaryDaoType != "" {
            currentPrimary := ds.DaoPaths[ds.PrimaryDaoType]
            currentPrimary.IsPrimary = false
            ds.DaoPaths[ds.PrimaryDaoType] = currentPrimary
        }
        
        // Set new primary (设置新主道) - Đặt Đạo chính mới
        ds.PrimaryDaoType = daoType
        dao.IsPrimary = true
        ds.DaoPaths[daoType] = dao
        return true
    }
    return false
}
```

## 🔄 Quản Lý Trạng Thái

### **ToggleDaoActive - Chuyển Đổi Trạng Thái Hoạt Động**

```go
// Toggle Dao Active Status (切换道激活状态) - Chuyển đổi trạng thái hoạt động của Đạo
func (ds *DaoSystem) ToggleDaoActive(daoType string) bool {
    if dao, exists := ds.DaoPaths[daoType]; exists {
        dao.IsActive = !dao.IsActive
        ds.DaoPaths[daoType] = dao
        
        // If deactivating primary Dao, set a new primary (如果停用主道则设置新的主道) - Nếu tắt Đạo chính thì đặt Đạo chính mới
        if !dao.IsActive && ds.PrimaryDaoType == daoType {
            ds.PrimaryDaoType = ""
            for daoType, dao := range ds.DaoPaths {
                if dao.IsActive {
                    ds.PrimaryDaoType = daoType
                    dao.IsPrimary = true
                    ds.DaoPaths[daoType] = dao
                    break
                }
            }
        }
        return true
    }
    return false
}
```

## 📊 Truy Vấn Thông Tin

### **GetActiveDaos - Lấy Các Đạo Đang Hoạt Động**

```go
// Get Active Daos (获取激活的道) - Lấy các Đạo đang hoạt động
func (ds *DaoSystem) GetActiveDaos() []DaoPath {
    var activeDaos []DaoPath
    for _, dao := range ds.DaoPaths {
        if dao.IsActive {
            activeDaos = append(activeDaos, dao)
        }
    }
    return activeDaos
}
```

### **GetDaoByType - Lấy Đạo Theo Loại**

```go
// Get Dao by Type (根据类型获取道) - Lấy Đạo theo loại
func (ds *DaoSystem) GetDaoByType(daoType string) (DaoPath, bool) {
    dao, exists := ds.DaoPaths[daoType]
    return dao, exists
}
```

## 🧮 Tính Toán Stats

### **CalculateDaoStats - Tính Toán Stats Đạo**

```go
// Calculate Dao stats (计算道属性) - Tính toán các stats của Đạo
func CalculateDaoStats(daoSystem DaoSystem) map[string]float64 {
    stats := make(map[string]float64)
    
    // Get primary Dao (获取主道) - Lấy Đạo chính
    primaryDao, exists := daoSystem.DaoPaths[daoSystem.PrimaryDaoType]
    if !exists {
        return stats // Return empty if no primary Dao (如果没有主道则返回空) - Nếu không có Đạo chính thì trả về rỗng
    }
    
    // Core Dao stats (核心道属性) - Stats cốt lõi Đạo
    stats["dao_level"] = float64(primaryDao.DaoLevel)
    stats["dao_experience"] = primaryDao.DaoExperience
    stats["dao_mastery"] = primaryDao.DaoMastery
    stats["dao_insight"] = primaryDao.DaoInsight
    stats["dao_wisdom"] = primaryDao.DaoWisdom
    
    // Advanced Dao stats (高级道属性) - Stats nâng cao Đạo
    stats["dao_harmony"] = primaryDao.DaoHarmony
    stats["dao_transcendence"] = primaryDao.DaoTranscendence
    
    // Calculate total stats across all active Daos (计算所有激活道的总属性) - Tính toán tổng stats từ tất cả Đạo đang hoạt động
    totalLevel := 0.0
    totalExperience := 0.0
    totalMastery := 0.0
    activeDaoCount := 0
    
    for _, dao := range daoSystem.DaoPaths {
        if dao.IsActive {
            totalLevel += float64(dao.DaoLevel)
            totalExperience += dao.DaoExperience
            totalMastery += dao.DaoMastery
            activeDaoCount++
        }
    }
    
    // Average stats across active Daos (激活道的平均属性) - Stats trung bình của các Đạo đang hoạt động
    if activeDaoCount > 0 {
        stats["total_dao_level"] = totalLevel
        stats["total_dao_experience"] = totalExperience
        stats["total_dao_mastery"] = totalMastery
        stats["active_dao_count"] = float64(activeDaoCount)
        stats["average_dao_level"] = totalLevel / float64(activeDaoCount)
    }
    
    // Calculate compatibility and resonance (计算兼容性和共振) - Tính toán tương thích và cộng hưởng
    stats["dao_compatibility"] = calculateDaoCompatibility(daoSystem.DaoCompatibility)
    stats["dao_resonance"] = calculateDaoResonance(daoSystem.DaoCompatibility)
    stats["dao_suppression"] = calculateDaoSuppression(daoSystem.DaoCompatibility)
    
    // Calculate technique stats (计算道术属性) - Tính toán stats kỹ thuật Đạo
    techniqueStats := calculateDaoTechniqueStats(daoSystem.DaoTechniques)
    for key, value := range techniqueStats {
        stats[key] = value
    }
    
    return stats
}
```

## 🔧 Hàm Trợ Giúp

### **calculateDaoCompatibility - Tính Tương Thích Đạo**

```go
// Helper functions (辅助函数) - Các hàm trợ giúp
func calculateDaoCompatibility(compatibility DaoCompatibility) float64 {
    // Calculate overall compatibility (计算总体兼容性) - Tính toán tương thích tổng thể
    totalCompatibility := 0.0
    count := 0
    
    for _, innerMap := range compatibility.CompatibilityMatrix {
        for _, value := range innerMap {
            totalCompatibility += value
            count++
        }
    }
    
    if count > 0 {
        return totalCompatibility / float64(count)
    }
    return 0.0
}
```

### **calculateDaoResonance - Tính Cộng Hưởng Đạo**

```go
func calculateDaoResonance(compatibility DaoCompatibility) float64 {
    // Calculate overall resonance (计算总体共振) - Tính toán cộng hưởng tổng thể
    totalResonance := 0.0
    count := 0
    
    for _, value := range compatibility.ResonanceEffects {
        totalResonance += value
        count++
    }
    
    if count > 0 {
        return totalResonance / float64(count)
    }
    return 0.0
}
```

### **calculateDaoSuppression - Tính Áp Chế Đạo**

```go
func calculateDaoSuppression(compatibility DaoCompatibility) float64 {
    // Calculate overall suppression (计算总体压制) - Tính toán áp chế tổng thể
    totalSuppression := 0.0
    count := 0
    
    for _, value := range compatibility.SuppressionEffects {
        totalSuppression += value
        count++
    }
    
    if count > 0 {
        return totalSuppression / float64(count)
    }
    return 0.0
}
```

### **calculateDaoTechniqueStats - Tính Stats Kỹ Thuật Đạo**

```go
func calculateDaoTechniqueStats(techniques map[string][]DaoTechnique) map[string]float64 {
    stats := make(map[string]float64)
    
    totalPower := 0.0
    totalCost := 0.0
    totalCooldown := 0.0
    totalRange := 0.0
    totalAccuracy := 0.0
    totalTechniques := 0
    
    for _, techniqueList := range techniques {
        for _, technique := range techniqueList {
            totalPower += technique.TechniquePower
            totalCost += technique.TechniqueCost
            totalCooldown += technique.TechniqueCooldown
            totalRange += technique.TechniqueRange
            totalAccuracy += technique.TechniqueAccuracy
            totalTechniques++
        }
    }
    
    if totalTechniques > 0 {
        stats["dao_technique_power"] = totalPower / float64(totalTechniques)
        stats["dao_technique_cost"] = totalCost / float64(totalTechniques)
        stats["dao_technique_cooldown"] = totalCooldown / float64(totalTechniques)
        stats["dao_technique_range"] = totalRange / float64(totalTechniques)
        stats["dao_technique_accuracy"] = totalAccuracy / float64(totalTechniques)
    }
    
    return stats
}
```

## 💡 Đặc Điểm Nổi Bật

### **Map-Based Management:**
- **Flexible**: Quản lý nhiều Đạo cùng lúc
- **Scalable**: Dễ dàng thêm loại Đạo mới
- **Efficient**: Truy cập O(1) cho mọi thao tác

### **Automatic Primary Management:**
- **Auto-assign**: Tự động gán Đạo chính khi thêm đầu tiên
- **Auto-switch**: Tự động chuyển Đạo chính khi xóa/tắt
- **Consistency**: Đảm bảo luôn có Đạo chính hợp lệ

### **Status Tracking:**
- **Active/Inactive**: Theo dõi trạng thái hoạt động
- **Unlock Time**: Ghi nhận thời điểm mở khóa
- **Primary Flag**: Quản lý cờ Đạo chính
