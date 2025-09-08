# Kim Đan (Golden Core) — **Luyện Khí** Subsystem Design (for Actor Core v3)
**Generated:** 2025-01-27  
**Status:** Restructured & Optimized

Hệ thống **Kim Đan** dựa cảm hứng từ _Phong Thần Diễn Nghĩa_ và đạo gia **nội đan** (internal alchemy).
Đây là **một Subsystem luyện khí** (không phải luyện thể), chịu trách nhiệm mô hình hoá **Tinh–Khí–Thần**,
các **khí quan** (đan điền, kinh mạch, thức hải), **cảnh giới/tiểu cảnh giới**, và phát ra
**Contributions** / **CapContributions** cho Actor Core.

> Core giữ vai trò Aggregator — toàn bộ công thức nằm ở Subsystem.

## 📁 Cấu Trúc Tài Liệu

### **🏗️ Components (Thành Phần Tinh-Khí-Thần)**
- `components/01_Dantian_System.md` - Hệ thống Đan Điền (Khí)
- `components/02_Meridian_System.md` - Hệ thống Kinh Mạch (Tinh)
- `components/03_Spirit_Sea_System.md` - Hệ thống Thức Hải (Thần)

### **🌅 Dao System (Hệ Thống Đạo & Áo Nghĩa)**
- `dao-system/01_Dao_System_Overview.md` - Tổng quan hệ thống Đạo
- `dao-system/02_Dao_Data_Structures.md` - Cấu trúc dữ liệu Đạo
- `dao-system/03_Dao_Management_Functions.md` - Các hàm quản lý Đạo
- `dao-system/04_Profound_Meaning_System.md` - Hệ thống Áo Nghĩa

### **🏛️ Realms (Cảnh Giới)**
- `realms/01_Realm_Definitions.md` - Định nghĩa cảnh giới từ Luyện Khí đến Thánh Nhân

### **⚙️ Core System (Hệ Thống Cốt Lõi)**
- `04_Caps_and_Layers_For_KimDan.md` - Caps và layers
- `05_Subsystem_Contract.md` - Subsystem contract
- `06_Implementation_Guide.md` - Hướng dẫn implement
- `08_Int64_Power_Scale_Analysis.md` - Phân tích int64
- `09_Float64_Performance_Impact_Analysis.md` - Phân tích float64

## 🚀 Tính Năng Mới

### **Map-Based Design:**
- **Dao System** - `map[string]DaoPath` cho flexibility
- **Profound Meaning System** - `map[string]ProfoundMeaning`
- **Scalable** - Dễ dàng thêm loại mới
- **Flexible** - Quản lý nhiều Đạo/Áo Nghĩa cùng lúc

### **Enhanced Components:**
- **Dantian System** - Hệ thống Đan Điền với flexible energy
- **Meridian System** - Hệ thống Kinh Mạch với 12 kênh chính
- **Spirit Sea System** - Hệ thống Thức Hải với Thần Thức

### **Turn-Based Combat Integration:**
- **Dao Techniques** - Kỹ thuật đặc biệt theo Đạo
- **Profound Abilities** - Khả năng đặc biệt theo Áo Nghĩa
- **Combat Multipliers** - Hệ số chiến đấu
