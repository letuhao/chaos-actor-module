# Magic Numbers Refactoring

## Tổng Quan

Đã thực hiện refactor tất cả magic numbers trong Actor Core v2.0 thành constants để dễ quản lý và bảo trì.

## Files Được Refactor

### 1. `constants/formula_constants.go`
- **Mục đích**: Chứa tất cả constants cho formula calculations
- **Nội dung**: 50+ constants cho các phép tính stat
- **Categories**:
  - Core Stats Multipliers (HPMax, Stamina, Speed, Haste)
  - Critical Hit Multipliers (CritChance, CritMulti)
  - Movement Multipliers (MoveSpeed, RegenHP)
  - Combat Stats Multipliers (Accuracy, Penetration, Lethality, etc.)
  - Energy Stats Multipliers (EnergyEfficiency, EnergyCapacity, etc.)
  - Learning Stats Multipliers (LearningRate, Adaptation, Memory, etc.)
  - Social Stats Multipliers (Leadership, Diplomacy, Intimidation, etc.)
  - Mystical Stats Multipliers (ManaEfficiency, SpellPower, etc.)
  - Movement Stats Multipliers (JumpHeight, ClimbSpeed, SwimSpeed, etc.)
  - Aura Stats Multipliers (AuraRadius, AuraStrength, Presence, etc.)
  - Proficiency Stats Multipliers (WeaponMastery, SkillLevel, etc.)
  - Talent Amplifiers Multipliers (CultivationSpeed, EnergyEfficiencyAmp, etc.)
  - Validation Constants (MinHasteValue, MaxCritChanceValue, etc.)

### 2. `services/core/stat_resolver.go`
- **Thay đổi**: Tất cả magic numbers trong formulas được thay thế bằng constants
- **Ví dụ**:
  ```go
  // Trước
  return float64(primary.Vitality*10 + primary.Constitution*5)
  
  // Sau
  return float64(primary.Vitality)*constants.HPMaxVitalityMultiplier + float64(primary.Constitution)*constants.HPMaxConstitutionMultiplier
  ```

### 3. `models/core/derived_stats.go`
- **Thay đổi**: Tất cả magic numbers trong CalculateFromPrimary method được thay thế bằng constants
- **Ví dụ**:
  ```go
  // Trước
  ds.HPMax = float64(pc.Vitality * 10 + pc.Constitution * 5)
  
  // Sau
  ds.HPMax = float64(pc.Vitality)*constants.HPMaxVitalityMultiplier + float64(pc.Constitution)*constants.HPMaxConstitutionMultiplier
  ```

## Lợi Ích

### 1. **Dễ Quản Lý**
- Tất cả magic numbers được tập trung trong một file
- Dễ dàng tìm kiếm và thay đổi values
- Tránh duplicate values

### 2. **Dễ Bảo Trì**
- Thay đổi một constant sẽ update tất cả nơi sử dụng
- Giảm lỗi khi copy-paste values
- Code dễ đọc và hiểu hơn

### 3. **Dễ Mở Rộng**
- Thêm constants mới dễ dàng
- Có thể tạo constants cho các hệ thống mới
- Hỗ trợ configuration từ file external

### 4. **Type Safety**
- Constants có type rõ ràng
- Compiler sẽ catch lỗi type mismatch
- IDE support tốt hơn

## Constants Categories

### Core Stats
- `HPMaxVitalityMultiplier = 10.0`
- `HPMaxConstitutionMultiplier = 5.0`
- `StaminaEnduranceMultiplier = 10.0`
- `StaminaConstitutionMultiplier = 3.0`

### Speed & Movement
- `SpeedAgilityMultiplier = 0.1`
- `HasteAgilityMultiplier = 0.01`
- `MoveSpeedAgilityMultiplier = 0.1`

### Critical Hit
- `CritChanceBaseValue = 0.05`
- `CritChanceLuckMultiplier = 0.001`
- `CritMultiBaseValue = 1.5`
- `CritMultiLuckMultiplier = 0.01`

### Combat Stats
- `AccuracyBaseValue = 0.8`
- `AccuracyIntelligenceMultiplier = 0.01`
- `PenetrationStrengthMultiplier = 0.01`
- `LethalityStrengthAgilityMultiplier = 0.005`

### Validation
- `MinHasteValue = 0.1`
- `MaxCritChanceValue = 1.0`
- `MinCritMultiValue = 1.0`
- `MinStatValue = 0.0`

## Testing

Tất cả unit tests vẫn PASS sau khi refactor:
- ✅ StatResolver tests: 16/16 PASS
- ✅ DerivedStats tests: 9/9 PASS
- ✅ PrimaryCore tests: 19/19 PASS

## Kết Luận

Việc refactor magic numbers thành constants đã hoàn thành thành công, giúp code dễ quản lý và bảo trì hơn mà không ảnh hưởng đến functionality.
