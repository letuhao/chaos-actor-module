package typeenum

// TypeCategory represents the category of a type
type TypeCategory int

const (
	// TypeCategoryDamage represents damage types
	TypeCategoryDamage TypeCategory = iota

	// TypeCategoryDefence represents defence types
	TypeCategoryDefence

	// TypeCategoryAmplifier represents amplifier types
	TypeCategoryAmplifier

	// TypeCategoryEnergy represents energy types
	TypeCategoryEnergy

	// TypeCategoryStatus represents status types
	TypeCategoryStatus

	// TypeCategoryElement represents element types
	TypeCategoryElement

	// TypeCategorySchool represents school types
	TypeCategorySchool

	// TypeCategoryRealm represents realm types
	TypeCategoryRealm
)

// String returns the string representation of TypeCategory
func (tc TypeCategory) String() string {
	switch tc {
	case TypeCategoryDamage:
		return "damage"
	case TypeCategoryDefence:
		return "defence"
	case TypeCategoryAmplifier:
		return "amplifier"
	case TypeCategoryEnergy:
		return "energy"
	case TypeCategoryStatus:
		return "status"
	case TypeCategoryElement:
		return "element"
	case TypeCategorySchool:
		return "school"
	case TypeCategoryRealm:
		return "realm"
	default:
		return "unknown"
	}
}

// IsValid checks if the TypeCategory is valid
func (tc TypeCategory) IsValid() bool {
	return tc >= TypeCategoryDamage && tc <= TypeCategoryRealm
}

// GetDisplayName returns the display name of TypeCategory
func (tc TypeCategory) GetDisplayName() string {
	switch tc {
	case TypeCategoryDamage:
		return "Damage Types"
	case TypeCategoryDefence:
		return "Defence Types"
	case TypeCategoryAmplifier:
		return "Amplifier Types"
	case TypeCategoryEnergy:
		return "Energy Types"
	case TypeCategoryStatus:
		return "Status Types"
	case TypeCategoryElement:
		return "Element Types"
	case TypeCategorySchool:
		return "School Types"
	case TypeCategoryRealm:
		return "Realm Types"
	default:
		return "Unknown Type Category"
	}
}

// GetDescription returns the description of TypeCategory
func (tc TypeCategory) GetDescription() string {
	switch tc {
	case TypeCategoryDamage:
		return "Types of damage that can be dealt to targets"
	case TypeCategoryDefence:
		return "Types of defence that can protect against damage"
	case TypeCategoryAmplifier:
		return "Types of amplifiers that can modify damage or effects"
	case TypeCategoryEnergy:
		return "Types of energy that can be used for various purposes"
	case TypeCategoryStatus:
		return "Types of status effects that can be applied to targets"
	case TypeCategoryElement:
		return "Types of elements that can be used in magic or cultivation"
	case TypeCategorySchool:
		return "Types of schools or disciplines in magic or cultivation"
	case TypeCategoryRealm:
		return "Types of realms or dimensions in the world"
	default:
		return "Unknown type category"
	}
}

// HasResistance checks if the TypeCategory has resistance mechanics
func (tc TypeCategory) HasResistance() bool {
	switch tc {
	case TypeCategoryDamage:
		return true
	case TypeCategoryDefence:
		return true
	case TypeCategoryAmplifier:
		return false
	case TypeCategoryEnergy:
		return true
	case TypeCategoryStatus:
		return true
	case TypeCategoryElement:
		return true
	case TypeCategorySchool:
		return false
	case TypeCategoryRealm:
		return false
	default:
		return false
	}
}

// HasWeakness checks if the TypeCategory has weakness mechanics
func (tc TypeCategory) HasWeakness() bool {
	switch tc {
	case TypeCategoryDamage:
		return true
	case TypeCategoryDefence:
		return false
	case TypeCategoryAmplifier:
		return false
	case TypeCategoryEnergy:
		return true
	case TypeCategoryStatus:
		return true
	case TypeCategoryElement:
		return true
	case TypeCategorySchool:
		return false
	case TypeCategoryRealm:
		return false
	default:
		return false
	}
}

// SupportsInteraction checks if the TypeCategory supports type interactions
func (tc TypeCategory) SupportsInteraction() bool {
	switch tc {
	case TypeCategoryDamage:
		return true
	case TypeCategoryDefence:
		return true
	case TypeCategoryAmplifier:
		return true
	case TypeCategoryEnergy:
		return true
	case TypeCategoryStatus:
		return true
	case TypeCategoryElement:
		return true
	case TypeCategorySchool:
		return false
	case TypeCategoryRealm:
		return false
	default:
		return false
	}
}

// GetAllTypeCategories returns all valid TypeCategory values
func GetAllTypeCategories() []TypeCategory {
	return []TypeCategory{
		TypeCategoryDamage,
		TypeCategoryDefence,
		TypeCategoryAmplifier,
		TypeCategoryEnergy,
		TypeCategoryStatus,
		TypeCategoryElement,
		TypeCategorySchool,
		TypeCategoryRealm,
	}
}

// GetTypeCategoryFromString converts a string to TypeCategory
func GetTypeCategoryFromString(s string) (TypeCategory, bool) {
	switch s {
	case "damage":
		return TypeCategoryDamage, true
	case "defence":
		return TypeCategoryDefence, true
	case "amplifier":
		return TypeCategoryAmplifier, true
	case "energy":
		return TypeCategoryEnergy, true
	case "status":
		return TypeCategoryStatus, true
	case "element":
		return TypeCategoryElement, true
	case "school":
		return TypeCategorySchool, true
	case "realm":
		return TypeCategoryRealm, true
	default:
		return TypeCategoryDamage, false
	}
}
