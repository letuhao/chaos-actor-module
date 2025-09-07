package stat

// StatType represents the type of a stat
type StatType int

const (
	// StatTypePrimary represents primary stats (basic attributes)
	StatTypePrimary StatType = iota

	// StatTypeDerived represents derived stats (calculated from primary stats)
	StatTypeDerived

	// StatTypeCustom represents custom stats (user-defined)
	StatTypeCustom

	// StatTypeSubSystem represents subsystem-specific stats
	StatTypeSubSystem
)

// String returns the string representation of StatType
func (st StatType) String() string {
	switch st {
	case StatTypePrimary:
		return "primary"
	case StatTypeDerived:
		return "derived"
	case StatTypeCustom:
		return "custom"
	case StatTypeSubSystem:
		return "subsystem"
	default:
		return "unknown"
	}
}

// IsValid checks if the StatType is valid
func (st StatType) IsValid() bool {
	return st >= StatTypePrimary && st <= StatTypeSubSystem
}

// GetDisplayName returns the display name of StatType
func (st StatType) GetDisplayName() string {
	switch st {
	case StatTypePrimary:
		return "Primary Stat"
	case StatTypeDerived:
		return "Derived Stat"
	case StatTypeCustom:
		return "Custom Stat"
	case StatTypeSubSystem:
		return "Subsystem Stat"
	default:
		return "Unknown Stat Type"
	}
}

// GetDescription returns the description of StatType
func (st StatType) GetDescription() string {
	switch st {
	case StatTypePrimary:
		return "Basic character attributes that form the foundation of all other stats"
	case StatTypeDerived:
		return "Calculated stats that are derived from primary stats and other factors"
	case StatTypeCustom:
		return "User-defined stats that can be customized for specific needs"
	case StatTypeSubSystem:
		return "Stats specific to a particular subsystem or cultivation system"
	default:
		return "Unknown stat type"
	}
}

// GetAllStatTypes returns all valid StatType values
func GetAllStatTypes() []StatType {
	return []StatType{
		StatTypePrimary,
		StatTypeDerived,
		StatTypeCustom,
		StatTypeSubSystem,
	}
}

// GetStatTypeFromString converts a string to StatType
func GetStatTypeFromString(s string) (StatType, bool) {
	switch s {
	case "primary":
		return StatTypePrimary, true
	case "derived":
		return StatTypeDerived, true
	case "custom":
		return StatTypeCustom, true
	case "subsystem":
		return StatTypeSubSystem, true
	default:
		return StatTypePrimary, false
	}
}
