package stat

// StatCategory represents the category of a stat
type StatCategory int

const (
	// Primary Stat Categories
	CategoryBasicStats StatCategory = iota
	CategoryPhysicalStats
	CategoryCultivationStats
	CategoryLifeStats

	// Derived Stat Categories
	CategoryCoreDerived
	CategoryCombatStats
	CategoryEnergyStats
	CategoryLearningStats
	CategorySocialStats
	CategoryMysticalStats
	CategoryMovementStats
	CategoryAuraStats
	CategoryProficiencyStats
	CategoryTalentAmplifiers

	// Flexible System Categories
	CategorySpeedSystem
	CategoryKarmaSystem
	CategoryAdministrativeSystem
	CategoryProficiencySystem
	CategorySkillSystem
	CategoryEnergySystem
	CategoryDamageSystem
	CategoryDefenceSystem
	CategoryAmplifierSystem
)

// String returns the string representation of StatCategory
func (sc StatCategory) String() string {
	switch sc {
	// Primary Stat Categories
	case CategoryBasicStats:
		return "basic"
	case CategoryPhysicalStats:
		return "physical"
	case CategoryCultivationStats:
		return "cultivation"
	case CategoryLifeStats:
		return "life"

	// Derived Stat Categories
	case CategoryCoreDerived:
		return "core_derived"
	case CategoryCombatStats:
		return "combat"
	case CategoryEnergyStats:
		return "energy"
	case CategoryLearningStats:
		return "learning"
	case CategorySocialStats:
		return "social"
	case CategoryMysticalStats:
		return "mystical"
	case CategoryMovementStats:
		return "movement"
	case CategoryAuraStats:
		return "aura"
	case CategoryProficiencyStats:
		return "proficiency"
	case CategoryTalentAmplifiers:
		return "talent_amplifiers"

	// Flexible System Categories
	case CategorySpeedSystem:
		return "speed_system"
	case CategoryKarmaSystem:
		return "karma_system"
	case CategoryAdministrativeSystem:
		return "administrative_system"
	case CategoryProficiencySystem:
		return "proficiency_system"
	case CategorySkillSystem:
		return "skill_system"
	case CategoryEnergySystem:
		return "energy_system"
	case CategoryDamageSystem:
		return "damage_system"
	case CategoryDefenceSystem:
		return "defence_system"
	case CategoryAmplifierSystem:
		return "amplifier_system"

	default:
		return "unknown"
	}
}

// IsValid checks if the StatCategory is valid
func (sc StatCategory) IsValid() bool {
	return sc >= CategoryBasicStats && sc <= CategoryAmplifierSystem
}

// GetDisplayName returns the display name of StatCategory
func (sc StatCategory) GetDisplayName() string {
	switch sc {
	// Primary Stat Categories
	case CategoryBasicStats:
		return "Basic Stats"
	case CategoryPhysicalStats:
		return "Physical Stats"
	case CategoryCultivationStats:
		return "Cultivation Stats"
	case CategoryLifeStats:
		return "Life Stats"

	// Derived Stat Categories
	case CategoryCoreDerived:
		return "Core Derived Stats"
	case CategoryCombatStats:
		return "Combat Stats"
	case CategoryEnergyStats:
		return "Energy Stats"
	case CategoryLearningStats:
		return "Learning Stats"
	case CategorySocialStats:
		return "Social Stats"
	case CategoryMysticalStats:
		return "Mystical Stats"
	case CategoryMovementStats:
		return "Movement Stats"
	case CategoryAuraStats:
		return "Aura Stats"
	case CategoryProficiencyStats:
		return "Proficiency Stats"
	case CategoryTalentAmplifiers:
		return "Talent Amplifiers"

	// Flexible System Categories
	case CategorySpeedSystem:
		return "Speed System"
	case CategoryKarmaSystem:
		return "Karma System"
	case CategoryAdministrativeSystem:
		return "Administrative System"
	case CategoryProficiencySystem:
		return "Proficiency System"
	case CategorySkillSystem:
		return "Skill System"
	case CategoryEnergySystem:
		return "Energy System"
	case CategoryDamageSystem:
		return "Damage System"
	case CategoryDefenceSystem:
		return "Defence System"
	case CategoryAmplifierSystem:
		return "Amplifier System"

	default:
		return "Unknown Category"
	}
}

// GetDescription returns the description of StatCategory
func (sc StatCategory) GetDescription() string {
	switch sc {
	// Primary Stat Categories
	case CategoryBasicStats:
		return "Fundamental character attributes that form the core of the character"
	case CategoryPhysicalStats:
		return "Physical attributes related to strength, agility, and physical capabilities"
	case CategoryCultivationStats:
		return "Universal cultivation attributes that work across all cultivation systems"
	case CategoryLifeStats:
		return "Life-related attributes such as lifespan and age"

	// Derived Stat Categories
	case CategoryCoreDerived:
		return "Core derived stats calculated from primary stats"
	case CategoryCombatStats:
		return "Combat-related derived stats for fighting and combat effectiveness"
	case CategoryEnergyStats:
		return "Energy management and efficiency derived stats"
	case CategoryLearningStats:
		return "Learning and adaptation derived stats"
	case CategorySocialStats:
		return "Social interaction and influence derived stats"
	case CategoryMysticalStats:
		return "Mystical and magical derived stats"
	case CategoryMovementStats:
		return "Movement and mobility derived stats"
	case CategoryAuraStats:
		return "Aura and presence derived stats"
	case CategoryProficiencyStats:
		return "Proficiency and mastery derived stats"
	case CategoryTalentAmplifiers:
		return "Talent-based amplification derived stats"

	// Flexible System Categories
	case CategorySpeedSystem:
		return "Flexible speed system for different types of speed"
	case CategoryKarmaSystem:
		return "Flexible karma system for different types of karma"
	case CategoryAdministrativeSystem:
		return "Flexible administrative division system"
	case CategoryProficiencySystem:
		return "Flexible proficiency tracking system"
	case CategorySkillSystem:
		return "Flexible universal skill system"
	case CategoryEnergySystem:
		return "Flexible energy type system"
	case CategoryDamageSystem:
		return "Flexible damage type system"
	case CategoryDefenceSystem:
		return "Flexible defence type system"
	case CategoryAmplifierSystem:
		return "Flexible amplifier type system"

	default:
		return "Unknown category"
	}
}

// IsPrimaryCategory checks if the category is for primary stats
func (sc StatCategory) IsPrimaryCategory() bool {
	return sc >= CategoryBasicStats && sc <= CategoryLifeStats
}

// IsDerivedCategory checks if the category is for derived stats
func (sc StatCategory) IsDerivedCategory() bool {
	return sc >= CategoryCoreDerived && sc <= CategoryTalentAmplifiers
}

// IsFlexibleSystemCategory checks if the category is for flexible systems
func (sc StatCategory) IsFlexibleSystemCategory() bool {
	return sc >= CategorySpeedSystem && sc <= CategoryAmplifierSystem
}

// GetParentCategory returns the parent category if applicable
func (sc StatCategory) GetParentCategory() *StatCategory {
	switch sc {
	case CategoryBasicStats, CategoryPhysicalStats, CategoryCultivationStats, CategoryLifeStats:
		return nil // These are top-level primary categories
	case CategoryCoreDerived, CategoryCombatStats, CategoryEnergyStats, CategoryLearningStats,
		CategorySocialStats, CategoryMysticalStats, CategoryMovementStats, CategoryAuraStats,
		CategoryProficiencyStats, CategoryTalentAmplifiers:
		return nil // These are top-level derived categories
	case CategorySpeedSystem, CategoryKarmaSystem, CategoryAdministrativeSystem,
		CategoryProficiencySystem, CategorySkillSystem, CategoryEnergySystem,
		CategoryDamageSystem, CategoryDefenceSystem, CategoryAmplifierSystem:
		return nil // These are top-level flexible system categories
	default:
		return nil
	}
}

// GetAllStatCategories returns all valid StatCategory values
func GetAllStatCategories() []StatCategory {
	return []StatCategory{
		// Primary Stat Categories
		CategoryBasicStats,
		CategoryPhysicalStats,
		CategoryCultivationStats,
		CategoryLifeStats,

		// Derived Stat Categories
		CategoryCoreDerived,
		CategoryCombatStats,
		CategoryEnergyStats,
		CategoryLearningStats,
		CategorySocialStats,
		CategoryMysticalStats,
		CategoryMovementStats,
		CategoryAuraStats,
		CategoryProficiencyStats,
		CategoryTalentAmplifiers,

		// Flexible System Categories
		CategorySpeedSystem,
		CategoryKarmaSystem,
		CategoryAdministrativeSystem,
		CategoryProficiencySystem,
		CategorySkillSystem,
		CategoryEnergySystem,
		CategoryDamageSystem,
		CategoryDefenceSystem,
		CategoryAmplifierSystem,
	}
}

// GetStatCategoryFromString converts a string to StatCategory
func GetStatCategoryFromString(s string) (StatCategory, bool) {
	switch s {
	// Primary Stat Categories
	case "basic":
		return CategoryBasicStats, true
	case "physical":
		return CategoryPhysicalStats, true
	case "cultivation":
		return CategoryCultivationStats, true
	case "life":
		return CategoryLifeStats, true

	// Derived Stat Categories
	case "core_derived":
		return CategoryCoreDerived, true
	case "combat":
		return CategoryCombatStats, true
	case "energy":
		return CategoryEnergyStats, true
	case "learning":
		return CategoryLearningStats, true
	case "social":
		return CategorySocialStats, true
	case "mystical":
		return CategoryMysticalStats, true
	case "movement":
		return CategoryMovementStats, true
	case "aura":
		return CategoryAuraStats, true
	case "proficiency":
		return CategoryProficiencyStats, true
	case "talent_amplifiers":
		return CategoryTalentAmplifiers, true

	// Flexible System Categories
	case "speed_system":
		return CategorySpeedSystem, true
	case "karma_system":
		return CategoryKarmaSystem, true
	case "administrative_system":
		return CategoryAdministrativeSystem, true
	case "proficiency_system":
		return CategoryProficiencySystem, true
	case "skill_system":
		return CategorySkillSystem, true
	case "energy_system":
		return CategoryEnergySystem, true
	case "damage_system":
		return CategoryDamageSystem, true
	case "defence_system":
		return CategoryDefenceSystem, true
	case "amplifier_system":
		return CategoryAmplifierSystem, true

	default:
		return CategoryBasicStats, false
	}
}
