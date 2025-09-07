# 03 — File-by-File Steps (Cursor-ready)


## constants/derived_stats.go


## constants/error_codes.go


## constants/flexible_systems.go


## constants/formula_constants.go


## constants/primary_stats.go


## constants/system_config.go


## docs/actor-core-interface/packages/actor-core/go/actorcore/actorcore.go


## enums/clamp/clamp_type.go

### `func (ct ClampType)  String() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ct ClampType)  IsValid() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ct ClampType)  GetDisplayName() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ct ClampType)  GetDescription() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ct ClampType)  RequiresMinValue() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ct ClampType)  RequiresMaxValue() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ct ClampType)  RequiresSoftCapValue() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ct ClampType)  RequiresSoftCapRate() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetAllClampTypes() []ClampType`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetClampTypeFromString(s string) (ClampType, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## enums/formula/formula_type.go

### `func (ft FormulaType)  String() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ft FormulaType)  IsValid() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ft FormulaType)  GetDisplayName() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ft FormulaType)  GetDescription() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ft FormulaType)  RequiresDependencies() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ft FormulaType)  SupportsCaching() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetAllFormulaTypes() []FormulaType`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetFormulaTypeFromString(s string) (FormulaType, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## enums/stat/data_type.go

### `func (dt DataType)  String() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (dt DataType)  IsValid() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (dt DataType)  GetDisplayName() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (dt DataType)  GetDescription() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (dt DataType)  GetGoType() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (dt DataType)  GetDefaultValue() interface`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetAllDataTypes() []DataType`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetDataTypeFromString(s string) (DataType, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (dt DataType)  IsNumeric() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (dt DataType)  IsComparable() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (dt DataType)  SupportsArithmetic() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## enums/stat/stat_category.go

### `func (sc StatCategory)  String() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sc StatCategory)  IsValid() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sc StatCategory)  GetDisplayName() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sc StatCategory)  GetDescription() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sc StatCategory)  IsPrimaryCategory() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sc StatCategory)  IsDerivedCategory() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sc StatCategory)  IsFlexibleSystemCategory() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sc StatCategory)  GetParentCategory() *StatCategory`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetAllStatCategories() []StatCategory`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetStatCategoryFromString(s string) (StatCategory, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## enums/stat/stat_type.go

### `func (st StatType)  String() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st StatType)  IsValid() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st StatType)  GetDisplayName() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st StatType)  GetDescription() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetAllStatTypes() []StatType`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetStatTypeFromString(s string) (StatType, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## enums/type/type_category.go

### `func (tc TypeCategory)  String() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (tc TypeCategory)  IsValid() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (tc TypeCategory)  GetDisplayName() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (tc TypeCategory)  GetDescription() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (tc TypeCategory)  HasResistance() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (tc TypeCategory)  HasWeakness() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (tc TypeCategory)  SupportsInteraction() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetAllTypeCategories() []TypeCategory`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetTypeCategoryFromString(s string) (TypeCategory, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## enums/validation/validation_severity.go

### `func (vs ValidationSeverity)  String() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vs ValidationSeverity)  IsValid() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vs ValidationSeverity)  GetDisplayName() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vs ValidationSeverity)  GetDescription() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vs ValidationSeverity)  GetPriority() int`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vs ValidationSeverity)  IsHigherThan(other ValidationSeverity) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vs ValidationSeverity)  IsLowerThan(other ValidationSeverity) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vs ValidationSeverity)  IsEqualOrHigherThan(other ValidationSeverity) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vs ValidationSeverity)  IsEqualOrLowerThan(other ValidationSeverity) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vs ValidationSeverity)  ShouldStopExecution() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vs ValidationSeverity)  ShouldLog() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vs ValidationSeverity)  ShouldAlert() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetAllValidationSeverities() []ValidationSeverity`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetValidationSeverityFromString(s string) (ValidationSeverity, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## enums/validation/validation_type.go

### `func (vt ValidationType)  String() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vt ValidationType)  IsValid() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vt ValidationType)  GetDisplayName() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vt ValidationType)  GetDescription() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vt ValidationType)  RequiresCondition() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vt ValidationType)  SupportsAsync() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (vt ValidationType)  SupportsCaching() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetAllValidationTypes() []ValidationType`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func GetValidationTypeFromString(s string) (ValidationType, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## interfaces/cache/cache_interface.go


## interfaces/configuration/config_manager.go


## interfaces/core/stat_consumer.go


## interfaces/core/stat_provider.go


## interfaces/core/stat_resolver.go


## interfaces/effects/effect_manager.go

### `func DefaultEffectConfig() *EffectConfig`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## interfaces/flexible/karma_interface.go


## interfaces/flexible/speed_interface.go


## interfaces/monitoring/performance_monitor.go


## models/core/derived_stats.go

### `func NewDerivedStats() *DerivedStats`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (ds *DerivedStats)  CalculateFromPrimary(pc *PrimaryCore)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Accept **context.Context** (if not already) and **propagate** it.
- 2) Enforce pipeline **Flat → Mult → Clamp**, and **bump version** once at finalize.
- 3) Use **topological order** by dependencies; detect **cycles** (return error).
- 4) Read/write cache under **RWMutex**; include **FormulaRegistryVersion** in key.
- 5) Do **not** mutate inputs; return **new snapshot** objects.

### `func (ds *DerivedStats)  GetStat(statName string) (float64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ds *DerivedStats)  SetStat(statName string, value float64) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (ds *DerivedStats)  GetAllStats() map[string]float64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ds *DerivedStats)  Clone() *DerivedStats`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Deep-copy all nested structs/slices/maps; keep timestamps & version.

### `func (ds *DerivedStats)  GetVersion() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ds *DerivedStats)  GetUpdatedAt() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ds *DerivedStats)  GetCreatedAt() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## models/core/errors.go


## models/core/primary_core.go

### `func NewPrimaryCore() *PrimaryCore`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func NewPrimaryCoreWithValues(values map[string]int64) *PrimaryCore`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (pc *PrimaryCore)  GetStat(statName string) (int64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  SetStat(statName string, value int64) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (pc *PrimaryCore)  GetAllStats() map[string]int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  UpdateStats(stats map[string]int64) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  Clone() *PrimaryCore`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Deep-copy all nested structs/slices/maps; keep timestamps & version.

### `func (pc *PrimaryCore)  Reset()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  Validate() []ValidationError`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Validate bounds; use `clamp.Spec` where available; return `[]ValidationError`.

### `func (pc *PrimaryCore)  GetBasicStats() map[string]int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  GetPhysicalStats() map[string]int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  GetCultivationStats() map[string]int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  GetLifeStats() map[string]int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  IsAlive() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  GetRemainingLife() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  AgeUp()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  AgeUpBy(amount int64)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  GetVersion() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  GetUpdatedAt() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pc *PrimaryCore)  GetCreatedAt() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## models/effects/combat_effect.go

> ⚠ Found placeholders or stubs here. Finish them first.

### `func NewCombatEffect(id, name string, effectType NonCombatEffectType, category EffectCategory) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (ce *CombatEffect)  SetCombatInfo(combatID, attackerID, targetID string, combatStartTime int64) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (ce *CombatEffect)  SetCombatEndTime(combatEndTime int64) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (ce *CombatEffect)  SetPersistAfterCombat(persist bool) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (ce *CombatEffect)  SetCombatOnly(combatOnly bool) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (ce *CombatEffect)  SetCombatIntensity(intensity float64) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (ce *CombatEffect)  AddCombatModifier(modifier CombatEffectModifier) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (ce *CombatEffect)  AddNonCombatModifier(modifier NonCombatEffectModifier) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (ce *CombatEffect)  IsInCombat() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ce *CombatEffect)  IsCombatEnded() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ce *CombatEffect)  GetEffectiveIntensity() float64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ce *CombatEffect)  GetEffectiveModifiers() []NonCombatEffectModifier`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ce *CombatEffect)  ShouldPersistAfterCombat() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ce *CombatEffect)  ConvertToNonCombatEffect() *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ce *CombatEffect)  Clone() *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Deep-copy all nested structs/slices/maps; keep timestamps & version.


## models/effects/combat_effect_examples.go

### `func NewCombatEffectExamples() *CombatEffectExamples`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (cee *CombatEffectExamples)  CreateCombatInjuryEffect(severity string, duration int64, combatID, attackerID, targetID string) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cee *CombatEffectExamples)  CreatePoisonEffect(poisonType string, duration int64, combatID, attackerID, targetID string) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cee *CombatEffectExamples)  CreateCombatCurseEffect(curseType string, duration int64, combatID, attackerID, targetID string) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cee *CombatEffectExamples)  CreateFearEffect(fearType string, duration int64, combatID, attackerID, targetID string) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cee *CombatEffectExamples)  CreateCombatFatigueEffect(fatigueType string, duration int64, combatID, attackerID, targetID string) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cee *CombatEffectExamples)  CreateCombatOnlyEffect(effectType string, duration int64, combatID, attackerID, targetID string) *CombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## models/effects/effect_examples.go

### `func NewEffectExamples() *EffectExamples`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (ee *EffectExamples)  CreateInjuryEffect(severity string, duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ee *EffectExamples)  CreateBrokenBoneEffect(bone string, duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ee *EffectExamples)  CreateQiDeviationEffect(severity string, duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ee *EffectExamples)  CreateMeridianBlockageEffect(meridian string, duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ee *EffectExamples)  CreateCurseEffect(curseType string, duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ee *EffectExamples)  CreateDivineCurseEffect(curseType string, duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ee *EffectExamples)  CreateCultivationBoostEffect(boostType string, duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ee *EffectExamples)  CreateLearningBoostEffect(boostType string, duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ee *EffectExamples)  CreateWeatherEffect(weatherType string, duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ee *EffectExamples)  CreateCombatFatigueEffect(severity string, duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (ee *EffectExamples)  CreateSpiritualExhaustionEffect(severity string, duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## models/effects/non_combat_effect.go

### `func NewNonCombatEffect(id, name string, effectType NonCombatEffectType, category EffectCategory) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (e *NonCombatEffect)  SetDuration(duration int64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (e *NonCombatEffect)  SetIntensity(intensity float64) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (e *NonCombatEffect)  SetStackable(stackable bool, maxStacks int) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (e *NonCombatEffect)  AddModifier(modifier NonCombatEffectModifier) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (e *NonCombatEffect)  AddCondition(condition EffectCondition) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (e *NonCombatEffect)  SetSource(source string) *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (e *NonCombatEffect)  Activate() *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (e *NonCombatEffect)  Deactivate() *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (e *NonCombatEffect)  IsExpired() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (e *NonCombatEffect)  GetRemainingDuration() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (e *NonCombatEffect)  Clone() *NonCombatEffect`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Deep-copy all nested structs/slices/maps; keep timestamps & version.


## models/flexible/flexible_stats.go

### `func NewFlexibleStats() *FlexibleStats`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (fs *FlexibleStats)  SetCustomPrimary(statName string, value int64)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (fs *FlexibleStats)  GetCustomPrimary(statName string) (int64, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  SetCustomDerived(statName string, value float64)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (fs *FlexibleStats)  GetCustomDerived(statName string) (float64, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  SetSubSystemStat(systemName, statName string, value float64)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (fs *FlexibleStats)  GetSubSystemStat(systemName, statName string) (float64, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  GetAllSubSystemStats(systemName string) (map[string]float64, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  RemoveCustomPrimary(statName string)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (fs *FlexibleStats)  RemoveCustomDerived(statName string)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (fs *FlexibleStats)  RemoveSubSystemStat(systemName, statName string)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (fs *FlexibleStats)  RemoveSubSystem(systemName string)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (fs *FlexibleStats)  GetAllCustomPrimary() map[string]int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  GetAllCustomDerived() map[string]float64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  GetAllSubSystems() []string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  GetStatsCount() int`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  GetCustomPrimaryCount() int`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  GetCustomDerivedCount() int`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  GetSubSystemCount() int`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  GetSubSystemStatsCount(systemName string) int`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  HasCustomPrimary(statName string) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  HasCustomDerived(statName string) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  HasSubSystemStat(systemName, statName string) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  HasSubSystem(systemName string) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  ClearCustomPrimary()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (fs *FlexibleStats)  ClearCustomDerived()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (fs *FlexibleStats)  ClearSubSystemStats()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (fs *FlexibleStats)  ClearAll()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (fs *FlexibleStats)  Clone() *FlexibleStats`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Deep-copy all nested structs/slices/maps; keep timestamps & version.

### `func (fs *FlexibleStats)  Merge(other *FlexibleStats)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  ToJSON() ([]byte, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  FromJSON(data []byte) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  GetVersion() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  GetUpdatedAt() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  GetCreatedAt() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (fs *FlexibleStats)  Validate() error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Validate bounds; use `clamp.Spec` where available; return `[]ValidationError`.

### `func (e *ValidationError)  Error() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## services/cache/mem_cache.go

### `func NewMemCache(config MemCacheConfig) *MemCache`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (mc *MemCache)  Get(key string) (interface{}, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  Set(key string, value interface{}, ttl time.Duration) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (mc *MemCache)  Delete(key string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  Clear() error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (mc *MemCache)  Exists(key string) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  GetTTL(key string) (time.Duration, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  SetTTL(key string, ttl time.Duration) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (mc *MemCache)  GetStats() *CacheStats`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  ResetStats() error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  Health() error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  updateHitRatio()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  calculateSize(value interface{}) int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Accept **context.Context** (if not already) and **propagate** it.
- 2) Enforce pipeline **Flat → Mult → Clamp**, and **bump version** once at finalize.
- 3) Use **topological order** by dependencies; detect **cycles** (return error).
- 4) Read/write cache under **RWMutex**; include **FormulaRegistryVersion** in key.
- 5) Do **not** mutate inputs; return **new snapshot** objects.

### `func (mc *MemCache)  needsEviction(additionalSize int64) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  getMaxSize() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  evictEntries(requiredSize int64)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  evictLRU(requiredSize int64)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  evictLFU(requiredSize int64)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  evictTTL(requiredSize int64)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  StartCleanup()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Start goroutines with context; use `select` on ctx.Done(); close channels cleanly.

### `func (mc *MemCache)  CleanupExpired()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  GetSize() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  GetEntryCount() int`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (mc *MemCache)  GetKeys() []string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## services/cache/redis_cache.go

> ⚠ Found placeholders or stubs here. Finish them first.

### `func NewRedisCache(config RedisConfig) (*RedisCache, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (rc *RedisCache)  Get(key string) (interface{}, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  Set(key string, value interface{}, ttl time.Duration) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (rc *RedisCache)  Delete(key string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  Clear() error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (rc *RedisCache)  Exists(key string) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  GetTTL(key string) (time.Duration, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  SetTTL(key string, ttl time.Duration) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (rc *RedisCache)  GetStats() *CacheStats`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  ResetStats() error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  Health() error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  updateHitRatio()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  parseMemoryInfo(info string) int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  GetClient() *redis.Client`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  GetContext() context.Context`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  Close() error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  GetKeys(pattern string) ([]string, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  GetSize() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  GetMemoryUsage() (int64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  SetWithTags(key string, value interface{}, ttl time.Duration, tags []string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (rc *RedisCache)  InvalidateByTag(tag string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Validate bounds; use `clamp.Spec` where available; return `[]ValidationError`.

### `func (rc *RedisCache)  Publish(channel string, message interface{}) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (rc *RedisCache)  Subscribe(channels ...string) *redis.PubSub`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## services/cache/state_tracker.go

### `func NewStateTracker(config StateTrackerConfig) *StateTracker`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (st *StateTracker)  TrackChange(change *StateChange) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  GetChanges(entityType, entityID string) ([]*StateChange, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  GetChangesByTimeRange(start, end int64) ([]*StateChange, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  GetChangesByUser(userID string) ([]*StateChange, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  RollbackChange(changeID string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  ClearChanges(entityType, entityID string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (st *StateTracker)  ClearChangesByTimeRange(start, end int64) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (st *StateTracker)  GetStats() *StateTrackingStats`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  ResetStats() error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  validateChange(change *StateChange) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Validate bounds; use `clamp.Spec` where available; return `[]ValidationError`.

### `func (st *StateTracker)  generateChangeID(change *StateChange) string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  checkConflict(change *StateChange) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  cleanupOldChanges()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  StartCleanup()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Start goroutines with context; use `select` on ctx.Done(); close channels cleanly.

### `func (st *StateTracker)  GetChange(changeID string) (*StateChange, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  GetChangesByType(changeType string) ([]*StateChange, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  GetChangesBySystem(systemID string) ([]*StateChange, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  ExportChanges() ([]byte, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  ImportChanges(data []byte) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  GetChangeCount() int`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (st *StateTracker)  GetActiveChangeCount() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## services/cache/types.go


## services/configuration/config_manager.go

### `func NewConfigManager() *ConfigManager`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func NewConfigManagerWithFile(filePath string) *ConfigManager`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (cm *ConfigManager)  SetConfig(key string, value interface{}) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (cm *ConfigManager)  GetConfig(key string) (interface{}, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetConfigString(key string) (string, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetConfigInt(key string) (int, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetConfigInt64(key string) (int64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetConfigFloat64(key string) (float64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetConfigBool(key string) (bool, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetConfigMap(key string) (map[string]interface{}, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetConfigSlice(key string) ([]interface{}, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  HasConfig(key string) bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  RemoveConfig(key string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (cm *ConfigManager)  GetAllConfigs() map[string]interface`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetConfigKeys() []string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetConfigCount() int`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  ClearConfigs()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (cm *ConfigManager)  LoadFromFile(filePath string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  SaveToFile(filePath string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  ReloadFromFile() error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  SaveToCurrentFile() error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  SetFilePath(filePath string)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (cm *ConfigManager)  GetFilePath() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetVersion() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetUpdatedAt() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  GetCreatedAt() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  ValidateConfig(key string, value interface{}) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Validate bounds; use `clamp.Spec` where available; return `[]ValidationError`.

### `func (cm *ConfigManager)  SetConfigWithValidation(key string, value interface{}) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (cm *ConfigManager)  Clone() *ConfigManager`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Deep-copy all nested structs/slices/maps; keep timestamps & version.

### `func (cm *ConfigManager)  Merge(other *ConfigManager)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  ToJSON() ([]byte, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cm *ConfigManager)  FromJSON(data []byte) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## services/core/stat_resolver.go

### `func (f *BasicFormula)  Calculate(primary *core.PrimaryCore) float64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Accept **context.Context** (if not already) and **propagate** it.
- 2) Enforce pipeline **Flat → Mult → Clamp**, and **bump version** once at finalize.
- 3) Use **topological order** by dependencies; detect **cycles** (return error).
- 4) Read/write cache under **RWMutex**; include **FormulaRegistryVersion** in key.
- 5) Do **not** mutate inputs; return **new snapshot** objects.

### `func (f *BasicFormula)  GetDependencies() []string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (f *BasicFormula)  GetName() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (f *BasicFormula)  GetType() string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func NewStatResolver() *StatResolver`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.
- 1) Accept **context.Context** (if not already) and **propagate** it.
- 2) Enforce pipeline **Flat → Mult → Clamp**, and **bump version** once at finalize.
- 3) Use **topological order** by dependencies; detect **cycles** (return error).
- 4) Read/write cache under **RWMutex**; include **FormulaRegistryVersion** in key.
- 5) Do **not** mutate inputs; return **new snapshot** objects.

### `func (sr *StatResolver)  initializeDefaultFormulas()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sr *StatResolver)  ResolveStats(primaryStats *core.PrimaryCore) (*core.DerivedStats, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Accept **context.Context** (if not already) and **propagate** it.
- 2) Enforce pipeline **Flat → Mult → Clamp**, and **bump version** once at finalize.
- 3) Use **topological order** by dependencies; detect **cycles** (return error).
- 4) Read/write cache under **RWMutex**; include **FormulaRegistryVersion** in key.
- 5) Do **not** mutate inputs; return **new snapshot** objects.

### `func (sr *StatResolver)  ResolveStatsWithContext(ctx context.Context, primaryStats *core.PrimaryCore) (*core.DerivedStats, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Accept **context.Context** (if not already) and **propagate** it.
- 2) Enforce pipeline **Flat → Mult → Clamp**, and **bump version** once at finalize.
- 3) Use **topological order** by dependencies; detect **cycles** (return error).
- 4) Read/write cache under **RWMutex**; include **FormulaRegistryVersion** in key.
- 5) Do **not** mutate inputs; return **new snapshot** objects.

### `func (sr *StatResolver)  ResolveStat(statName string, primaryStats *core.PrimaryCore) (float64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Accept **context.Context** (if not already) and **propagate** it.
- 2) Enforce pipeline **Flat → Mult → Clamp**, and **bump version** once at finalize.
- 3) Use **topological order** by dependencies; detect **cycles** (return error).
- 4) Read/write cache under **RWMutex**; include **FormulaRegistryVersion** in key.
- 5) Do **not** mutate inputs; return **new snapshot** objects.

### `func (sr *StatResolver)  ResolveStatWithContext(ctx context.Context, statName string, primaryStats *core.PrimaryCore) (float64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Accept **context.Context** (if not already) and **propagate** it.
- 2) Enforce pipeline **Flat → Mult → Clamp**, and **bump version** once at finalize.
- 3) Use **topological order** by dependencies; detect **cycles** (return error).
- 4) Read/write cache under **RWMutex**; include **FormulaRegistryVersion** in key.
- 5) Do **not** mutate inputs; return **new snapshot** objects.

### `func (sr *StatResolver)  ResolveDerivedStats(primaryStats *core.PrimaryCore) (map[string]float64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Accept **context.Context** (if not already) and **propagate** it.
- 2) Enforce pipeline **Flat → Mult → Clamp**, and **bump version** once at finalize.
- 3) Use **topological order** by dependencies; detect **cycles** (return error).
- 4) Read/write cache under **RWMutex**; include **FormulaRegistryVersion** in key.
- 5) Do **not** mutate inputs; return **new snapshot** objects.

### `func (sr *StatResolver)  ResolveDerivedStatsWithContext(ctx context.Context, primaryStats *core.PrimaryCore) (map[string]float64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Accept **context.Context** (if not already) and **propagate** it.
- 2) Enforce pipeline **Flat → Mult → Clamp**, and **bump version** once at finalize.
- 3) Use **topological order** by dependencies; detect **cycles** (return error).
- 4) Read/write cache under **RWMutex**; include **FormulaRegistryVersion** in key.
- 5) Do **not** mutate inputs; return **new snapshot** objects.

### `func (sr *StatResolver)  CheckDependencies(statName string) ([]string, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sr *StatResolver)  CheckDependenciesWithContext(ctx context.Context, statName string) ([]string, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sr *StatResolver)  GetCalculationOrder() []string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sr *StatResolver)  GetCalculationOrderWithContext(ctx context.Context) []string`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sr *StatResolver)  ValidateStats(stats map[string]float64) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Validate bounds; use `clamp.Spec` where available; return `[]ValidationError`.

### `func (sr *StatResolver)  ValidateStatsWithContext(ctx context.Context, stats map[string]float64) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Validate bounds; use `clamp.Spec` where available; return `[]ValidationError`.

### `func (sr *StatResolver)  AddFormula(formula Formula) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (sr *StatResolver)  RemoveFormula(statName string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (sr *StatResolver)  GetFormula(statName string) (Formula, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sr *StatResolver)  GetAllFormulas() map[string]Formula`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sr *StatResolver)  ClearCache()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (sr *StatResolver)  GetCacheSize() int`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Read under `RLock`; avoid exposing internal map (copy out scalar).

### `func (sr *StatResolver)  GetVersion() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (sr *StatResolver)  GetStatsCount() int`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## services/effects/combat_effect_manager.go

### `func NewCombatEffectManager(effectManager *EffectManagerImpl) *CombatEffectManager`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (cem *CombatEffectManager)  StartCombat(ctx context.Context, combatID string, participants []string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Start goroutines with context; use `select` on ctx.Done(); close channels cleanly.

### `func (cem *CombatEffectManager)  EndCombat(ctx context.Context, combatID string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cem *CombatEffectManager)  ApplyCombatEffect(ctx context.Context, combatEffect *CombatEffect) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cem *CombatEffectManager)  RemoveCombatEffect(ctx context.Context, combatID, effectID string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (cem *CombatEffectManager)  GetCombatEffects(ctx context.Context, combatID string) ([]*CombatEffect, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cem *CombatEffectManager)  GetCombatEffect(ctx context.Context, effectID string) (*CombatEffect, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cem *CombatEffectManager)  GetCombatState(ctx context.Context, combatID string) (*CombatState, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cem *CombatEffectManager)  ProcessCombatEffects(ctx context.Context, combatID string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cem *CombatEffectManager)  ProcessAllCombatEffects(ctx context.Context) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cem *CombatEffectManager)  GetCombatStats(ctx context.Context) (*CombatStats, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (cem *CombatEffectManager)  ClearCombatEffects(ctx context.Context, combatID string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).


## services/effects/effect_manager.go

### `func NewEffectManager(config *EffectConfig) *EffectManagerImpl`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.

### `func (em *EffectManagerImpl)  ApplyEffect(ctx context.Context, actorID string, effect *effects.NonCombatEffect) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (em *EffectManagerImpl)  RemoveEffect(ctx context.Context, actorID string, effectID string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (em *EffectManagerImpl)  RemoveEffectByType(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (em *EffectManagerImpl)  GetActiveEffects(ctx context.Context, actorID string) ([]*effects.NonCombatEffect, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (em *EffectManagerImpl)  GetEffectByID(ctx context.Context, actorID string, effectID string) (*effects.NonCombatEffect, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (em *EffectManagerImpl)  GetEffectsByType(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) ([]*effects.NonCombatEffect, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (em *EffectManagerImpl)  GetEffectsByCategory(ctx context.Context, actorID string, category effects.EffectCategory) ([]*effects.NonCombatEffect, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (em *EffectManagerImpl)  ProcessEffects(ctx context.Context, actorID string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (em *EffectManagerImpl)  ProcessAllEffects(ctx context.Context) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (em *EffectManagerImpl)  SetResistance(ctx context.Context, actorID string, effectType effects.NonCombatEffectType, resistance float64) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (em *EffectManagerImpl)  GetResistance(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) (float64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (em *EffectManagerImpl)  SetImmunity(ctx context.Context, actorID string, effectType effects.NonCombatEffectType, immune bool) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (em *EffectManagerImpl)  GetImmunity(ctx context.Context, actorID string, effectType effects.NonCombatEffectType) (bool, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (em *EffectManagerImpl)  ClearAllEffects(ctx context.Context, actorID string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (em *EffectManagerImpl)  GetEffectStats(ctx context.Context, actorID string) (*EffectStats, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (em *EffectManagerImpl)  getOrCreateActorEffects(actorID string) *ActorEffects`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (em *EffectManagerImpl)  processEffectTick(effect *effects.NonCombatEffect)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.


## services/monitoring/performance_monitor.go

> ⚠ Found placeholders or stubs here. Finish them first.

### `func NewPerformanceMonitor() *PerformanceMonitor`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Initialize **timestamps** (`CreatedAt`,`UpdatedAt`) and **Version`=1`.
- 2) Apply **sane defaults**; avoid magic numbers (move to `constants/`).
- 3) Ensure **zero allocations** beyond what's necessary; avoid slices/maps unless needed.
- 1) Start goroutines with context; use `select` on ctx.Done(); close channels cleanly.

### `func (pm *PerformanceMonitor)  SetMetric(name string, value float64, unit string, category string, description string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (pm *PerformanceMonitor)  GetMetric(name string) (*Metric, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  GetMetricValue(name string) (float64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  GetMetricHistory(name string) ([]float64, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  GetAllMetrics() map[string]*Metric`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  GetMetricsByCategory(category string) map[string]*Metric`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  RemoveMetric(name string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (pm *PerformanceMonitor)  ClearMetrics()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (pm *PerformanceMonitor)  SetThreshold(metricName string, threshold float64)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (pm *PerformanceMonitor)  GetThreshold(metricName string) (float64, bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  CreateAlert(id, metricName string, threshold float64, operator, message, severity string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  GetAlert(id string) (*Alert, error)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  GetAllAlerts() map[string]*Alert`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  GetActiveAlerts() map[string]*Alert`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  EnableAlert(id string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  DisableAlert(id string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  RemoveAlert(id string) error`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (pm *PerformanceMonitor)  ClearAlerts()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (pm *PerformanceMonitor)  checkAlerts(metricName string, value float64)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  GetPerformanceStats() *PerformanceStats`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  SetEnabled(enabled bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (pm *PerformanceMonitor)  IsEnabled() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  SetAlertEnabled(enabled bool)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Protect mutation with `Lock` and defer `Unlock()`.
- 2) **Invalidate cache** appropriately (matching tags/keys).

### `func (pm *PerformanceMonitor)  IsAlertEnabled() bool`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  GetVersion() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  GetUpdatedAt() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  GetCreatedAt() int64`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  Clone() *PerformanceMonitor`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Deep-copy all nested structs/slices/maps; keep timestamps & version.

### `func (pm *PerformanceMonitor)  Reset()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.

### `func (pm *PerformanceMonitor)  StartMonitoring(ctx context.Context)`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
- 1) Start goroutines with context; use `select` on ctx.Done(); close channels cleanly.

### `func (pm *PerformanceMonitor)  collectSystemMetrics()`
- ✅ Keep function **pure/deterministic** (no hidden global state).
- ✅ **Return errors** instead of panicking; include context in error messages.
