package formula

// FormulaType represents the type of a formula
type FormulaType int

const (
	// FormulaTypeCalculation represents a calculation formula
	FormulaTypeCalculation FormulaType = iota

	// FormulaTypeConditional represents a conditional formula
	FormulaTypeConditional

	// FormulaTypeLookup represents a lookup formula
	FormulaTypeLookup

	// FormulaTypeCustom represents a custom formula
	FormulaTypeCustom
)

// String returns the string representation of FormulaType
func (ft FormulaType) String() string {
	switch ft {
	case FormulaTypeCalculation:
		return "calculation"
	case FormulaTypeConditional:
		return "conditional"
	case FormulaTypeLookup:
		return "lookup"
	case FormulaTypeCustom:
		return "custom"
	default:
		return "unknown"
	}
}

// IsValid checks if the FormulaType is valid
func (ft FormulaType) IsValid() bool {
	return ft >= FormulaTypeCalculation && ft <= FormulaTypeCustom
}

// GetDisplayName returns the display name of FormulaType
func (ft FormulaType) GetDisplayName() string {
	switch ft {
	case FormulaTypeCalculation:
		return "Calculation Formula"
	case FormulaTypeConditional:
		return "Conditional Formula"
	case FormulaTypeLookup:
		return "Lookup Formula"
	case FormulaTypeCustom:
		return "Custom Formula"
	default:
		return "Unknown Formula Type"
	}
}

// GetDescription returns the description of FormulaType
func (ft FormulaType) GetDescription() string {
	switch ft {
	case FormulaTypeCalculation:
		return "Mathematical calculation formula that computes a value from input parameters"
	case FormulaTypeConditional:
		return "Conditional formula that returns different values based on conditions"
	case FormulaTypeLookup:
		return "Lookup formula that retrieves values from tables or maps"
	case FormulaTypeCustom:
		return "Custom formula that can be defined by users or subsystems"
	default:
		return "Unknown formula type"
	}
}

// RequiresDependencies checks if the FormulaType requires dependencies
func (ft FormulaType) RequiresDependencies() bool {
	switch ft {
	case FormulaTypeCalculation:
		return true
	case FormulaTypeConditional:
		return true
	case FormulaTypeLookup:
		return false
	case FormulaTypeCustom:
		return true
	default:
		return false
	}
}

// SupportsCaching checks if the FormulaType supports caching
func (ft FormulaType) SupportsCaching() bool {
	switch ft {
	case FormulaTypeCalculation:
		return true
	case FormulaTypeConditional:
		return true
	case FormulaTypeLookup:
		return true
	case FormulaTypeCustom:
		return true
	default:
		return false
	}
}

// GetAllFormulaTypes returns all valid FormulaType values
func GetAllFormulaTypes() []FormulaType {
	return []FormulaType{
		FormulaTypeCalculation,
		FormulaTypeConditional,
		FormulaTypeLookup,
		FormulaTypeCustom,
	}
}

// GetFormulaTypeFromString converts a string to FormulaType
func GetFormulaTypeFromString(s string) (FormulaType, bool) {
	switch s {
	case "calculation":
		return FormulaTypeCalculation, true
	case "conditional":
		return FormulaTypeConditional, true
	case "lookup":
		return FormulaTypeLookup, true
	case "custom":
		return FormulaTypeCustom, true
	default:
		return FormulaTypeCalculation, false
	}
}
