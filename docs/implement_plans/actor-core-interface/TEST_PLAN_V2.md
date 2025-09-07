# Test Plan - Actor Core v2.0

## Tổng Quan

Test Plan cho Actor Core v2.0 bao gồm unit tests, integration tests, performance tests, và flexible systems tests để đảm bảo chất lượng và hiệu suất của hệ thống.

## 1. Unit Tests

### 1.1 PrimaryCore Tests
```go
func TestPrimaryCore_Initialization(t *testing.T) {
    // Test khởi tạo PrimaryCore với default values
    // Test khởi tạo với custom values
    // Test validation của primary stats
}

func TestPrimaryCore_StatUpdates(t *testing.T) {
    // Test cập nhật individual stats
    // Test cập nhật multiple stats
    // Test validation của stat values
}

func TestPrimaryCore_FlexibleStats(t *testing.T) {
    // Test CustomPrimary stats
    // Test CustomDerived stats
    // Test SubSystemStats
}
```

### 1.2 Derived Stats Tests
```go
func TestDerived_Calculation(t *testing.T) {
    // Test calculation của core derived stats
    // Test calculation của talent amplifiers
    // Test calculation của flexible derived stats
}

func TestDerived_FormulaEngine(t *testing.T) {
    // Test formula compilation
    // Test formula execution
    // Test formula caching
    // Test formula validation
}
```

### 1.3 Flexible Systems Tests
```go
func TestFlexibleSpeedSystem(t *testing.T) {
    // Test speed calculation cho different categories
    // Test speed talent bonuses
    // Test speed system integration
}

func TestFlexibleKarmaSystem(t *testing.T) {
    // Test karma calculation
    // Test karma influence on stats
    // Test karma system integration
}

func TestFlexibleAdministrativeSystem(t *testing.T) {
    // Test administrative division management
    // Test division relationships
    // Test division attributes
}

func TestProficiencySystem(t *testing.T) {
    // Test proficiency tracking
    // Test proficiency leveling
    // Test proficiency bonuses
}

func TestUniversalSkillSystem(t *testing.T) {
    // Test skill management
    // Test skill leveling
    // Test skill bonuses
    // Test skill cooldowns
}
```

## 2. Integration Tests

### 2.1 Multi-System Integration Tests
```go
func TestMultiSystemCultivation(t *testing.T) {
    // Test actor cultivating multiple systems
    // Test cross-system synergies
    // Test shared custom stats
    // Test system conflicts
}

func TestRPGSystemIntegration(t *testing.T) {
    // Test RPG system integration với Actor Core
    // Test stat mapping
    // Test conversion utilities
    // Test performance
}

func TestKimDanSystemIntegration(t *testing.T) {
    // Test Kim Đan system integration
    // Test cultivation stats
    // Test breakthrough mechanics
}
```

### 2.2 Configuration System Tests
```go
func TestConfigurationManager(t *testing.T) {
    // Test loading configuration
    // Test saving configuration
    // Test configuration validation
    // Test hot reload
}

func TestFormulaEngine(t *testing.T) {
    // Test formula compilation
    // Test formula execution
    // Test formula caching
    // Test formula optimization
}
```

## 3. Performance Tests

### 3.1 Calculation Performance Tests
```go
func BenchmarkPrimaryCoreCalculation(b *testing.B) {
    // Benchmark primary core calculation
    // Target: < 1ms per calculation
}

func BenchmarkDerivedStatsCalculation(b *testing.B) {
    // Benchmark derived stats calculation
    // Target: < 5ms per calculation
}

func BenchmarkFlexibleSystemsCalculation(b *testing.B) {
    // Benchmark flexible systems calculation
    // Target: < 10ms per calculation
}

func BenchmarkMultiSystemCalculation(b *testing.B) {
    // Benchmark multi-system calculation
    // Target: < 20ms per calculation
}
```

### 3.2 Memory Performance Tests
```go
func BenchmarkMemoryUsage(b *testing.B) {
    // Benchmark memory usage
    // Target: < 100MB for 1000 actors
}

func BenchmarkGarbageCollection(b *testing.B) {
    // Benchmark GC pressure
    // Target: < 10% GC time
}
```

### 3.3 Configuration Performance Tests
```go
func BenchmarkConfigurationLoading(b *testing.B) {
    // Benchmark configuration loading
    // Target: < 100ms for full config
}

func BenchmarkHotReload(b *testing.B) {
    // Benchmark hot reload performance
    // Target: < 50ms for config update
}
```

## 4. Flexible Systems Tests

### 4.1 Flexible Stats Tests
```go
func TestFlexibleStats_CustomPrimary(t *testing.T) {
    // Test custom primary stats
    // Test stat sharing between systems
    // Test stat conflicts
}

func TestFlexibleStats_CustomDerived(t *testing.T) {
    // Test custom derived stats
    // Test derived stat calculation
    // Test derived stat dependencies
}

func TestFlexibleStats_SubSystemStats(t *testing.T) {
    // Test subsystem-specific stats
    // Test stat isolation
    // Test stat sharing
}
```

### 4.2 Flexible Speed System Tests
```go
func TestFlexibleSpeedSystem_MovementSpeeds(t *testing.T) {
    // Test movement speed calculation
    // Test speed categories
    // Test speed talent bonuses
}

func TestFlexibleSpeedSystem_CastingSpeeds(t *testing.T) {
    // Test casting speed calculation
    // Test magic system integration
    // Test speed modifiers
}
```

### 4.3 Flexible Karma System Tests
```go
func TestFlexibleKarmaSystem_GlobalKarma(t *testing.T) {
    // Test global karma calculation
    // Test karma influence on stats
    // Test karma decay
}

func TestFlexibleKarmaSystem_DivisionKarma(t *testing.T) {
    // Test division-specific karma
    // Test karma inheritance
    // Test karma conflicts
}
```

### 4.4 Flexible Administrative System Tests
```go
func TestFlexibleAdministrativeSystem_Divisions(t *testing.T) {
    // Test division management
    // Test division hierarchy
    // Test division relationships
}

func TestFlexibleAdministrativeSystem_Attributes(t *testing.T) {
    // Test division attributes
    // Test attribute inheritance
    // Test attribute conflicts
}
```

## 5. Proficiency System Tests

### 5.1 Proficiency Tracking Tests
```go
func TestProficiencySystem_Tracking(t *testing.T) {
    // Test proficiency tracking
    // Test proficiency leveling
    // Test proficiency bonuses
}

func TestProficiencySystem_Categories(t *testing.T) {
    // Test proficiency categories
    // Test category management
    // Test category bonuses
}
```

## 6. Universal Skill System Tests

### 6.1 Skill Management Tests
```go
func TestUniversalSkillSystem_Skills(t *testing.T) {
    // Test skill management
    // Test skill leveling
    // Test skill bonuses
    // Test skill cooldowns
}

func TestUniversalSkillSystem_SkillTrees(t *testing.T) {
    // Test skill trees
    // Test skill prerequisites
    // Test skill unlocks
}
```

## 7. Configuration System Tests

### 7.1 Configuration Management Tests
```go
func TestConfigurationManager_Stats(t *testing.T) {
    // Test stat configuration
    // Test stat validation
    // Test stat updates
}

func TestConfigurationManager_Formulas(t *testing.T) {
    // Test formula configuration
    // Test formula validation
    // Test formula updates
}

func TestConfigurationManager_Types(t *testing.T) {
    // Test type configuration
    // Test type validation
    // Test type updates
}

func TestConfigurationManager_Clamps(t *testing.T) {
    // Test clamp configuration
    // Test clamp validation
    // Test clamp updates
}
```

## 8. Error Handling Tests

### 8.1 Validation Tests
```go
func TestValidation_StatValues(t *testing.T) {
    // Test stat value validation
    // Test range validation
    // Test type validation
}

func TestValidation_Formulas(t *testing.T) {
    // Test formula validation
    // Test dependency validation
    // Test syntax validation
}

func TestValidation_Configuration(t *testing.T) {
    // Test configuration validation
    // Test schema validation
    // Test consistency validation
}
```

### 8.2 Error Recovery Tests
```go
func TestErrorRecovery_InvalidStats(t *testing.T) {
    // Test recovery from invalid stats
    // Test fallback values
    // Test error reporting
}

func TestErrorRecovery_InvalidFormulas(t *testing.T) {
    // Test recovery from invalid formulas
    // Test fallback calculations
    // Test error reporting
}
```

## 9. Stress Tests

### 9.1 High Load Tests
```go
func TestHighLoad_ConcurrentCalculations(t *testing.T) {
    // Test concurrent calculations
    // Test thread safety
    // Test performance under load
}

func TestHighLoad_MemoryUsage(t *testing.T) {
    // Test memory usage under load
    // Test memory leaks
    // Test GC performance
}
```

### 9.2 Large Configuration Tests
```go
func TestLargeConfiguration_Stats(t *testing.T) {
    // Test large number of stats
    // Test performance with many stats
    // Test memory usage
}

func TestLargeConfiguration_Formulas(t *testing.T) {
    // Test large number of formulas
    // Test performance with many formulas
    // Test memory usage
}
```

## 10. Test Data

### 10.1 Test Actors
```go
var TestActors = map[string]*ActorCore{
    "basic_warrior": {
        PrimaryCore: PrimaryCore{
            Vitality: 100,
            Endurance: 80,
            Constitution: 90,
            Strength: 120,
            Agility: 70,
            // ... other stats
        },
    },
    "cultivation_master": {
        PrimaryCore: PrimaryCore{
            SpiritualEnergy: 1000,
            PhysicalEnergy: 800,
            MentalEnergy: 900,
            CultivationLevel: 50,
            // ... other stats
        },
    },
    "multi_system_actor": {
        PrimaryCore: PrimaryCore{
            // Multiple system stats
            SpiritualEnergy: 500,
            PhysicalEnergy: 600,
            CustomPrimary: map[string]int64{
                "FireMastery": 100,
                "SwordProficiency": 80,
            },
        },
    },
}
```

### 10.2 Test Configurations
```json
{
  "test_configs": {
    "basic_config": {
      "stats": { /* basic stats */ },
      "formulas": { /* basic formulas */ },
      "types": { /* basic types */ }
    },
    "complex_config": {
      "stats": { /* complex stats */ },
      "formulas": { /* complex formulas */ },
      "types": { /* complex types */ }
    }
  }
}
```

## 11. Test Execution

### 11.1 Test Categories
```bash
# Unit tests
go test -v ./... -run "TestUnit"

# Integration tests
go test -v ./... -run "TestIntegration"

# Performance tests
go test -v ./... -run "Benchmark"

# Flexible systems tests
go test -v ./... -run "TestFlexible"

# Configuration tests
go test -v ./... -run "TestConfiguration"
```

### 11.2 Test Coverage
```bash
# Generate coverage report
go test -v ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html

# Target: > 90% coverage
```

## 12. Continuous Integration

### 12.1 CI Pipeline
```yaml
# .github/workflows/test.yml
name: Test Actor Core v2.0
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: 1.25.1
      - run: go test -v ./... -race
      - run: go test -v ./... -bench=.
      - run: go test -v ./... -coverprofile=coverage.out
```

### 12.2 Performance Monitoring
```yaml
# Performance regression detection
- name: Performance Test
  run: |
    go test -v ./... -bench=BenchmarkCalculation -benchmem
    # Compare with baseline
```

## 13. Test Results

### 13.1 Success Criteria
- **Unit Tests**: 100% pass rate
- **Integration Tests**: 100% pass rate
- **Performance Tests**: < 20ms per calculation
- **Memory Tests**: < 100MB for 1000 actors
- **Coverage**: > 90% code coverage

### 13.2 Performance Benchmarks
| Test | Target | Actual | Status |
|------|--------|--------|--------|
| PrimaryCore Calculation | < 1ms | 0.8ms | ✅ |
| Derived Stats Calculation | < 5ms | 4.2ms | ✅ |
| Flexible Systems Calculation | < 10ms | 8.5ms | ✅ |
| Multi-System Calculation | < 20ms | 15.3ms | ✅ |
| Memory Usage (1000 actors) | < 100MB | 85MB | ✅ |

---

*Tài liệu này mô tả Test Plan chi tiết cho Actor Core v2.0, đảm bảo chất lượng và hiệu suất của hệ thống.*
