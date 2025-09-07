package validation

// ValidationType represents the type of a validation
type ValidationType int

const (
	// ValidationTypeRange represents a range validation
	ValidationTypeRange ValidationType = iota

	// ValidationTypeDependency represents a dependency validation
	ValidationTypeDependency

	// ValidationTypeFormula represents a formula validation
	ValidationTypeFormula

	// ValidationTypeCustom represents a custom validation
	ValidationTypeCustom
)

// String returns the string representation of ValidationType
func (vt ValidationType) String() string {
	switch vt {
	case ValidationTypeRange:
		return "range"
	case ValidationTypeDependency:
		return "dependency"
	case ValidationTypeFormula:
		return "formula"
	case ValidationTypeCustom:
		return "custom"
	default:
		return "unknown"
	}
}

// IsValid checks if the ValidationType is valid
func (vt ValidationType) IsValid() bool {
	return vt >= ValidationTypeRange && vt <= ValidationTypeCustom
}

// GetDisplayName returns the display name of ValidationType
func (vt ValidationType) GetDisplayName() string {
	switch vt {
	case ValidationTypeRange:
		return "Range Validation"
	case ValidationTypeDependency:
		return "Dependency Validation"
	case ValidationTypeFormula:
		return "Formula Validation"
	case ValidationTypeCustom:
		return "Custom Validation"
	default:
		return "Unknown Validation Type"
	}
}

// GetDescription returns the description of ValidationType
func (vt ValidationType) GetDescription() string {
	switch vt {
	case ValidationTypeRange:
		return "Validates that a value is within a specified range"
	case ValidationTypeDependency:
		return "Validates that all required dependencies are present"
	case ValidationTypeFormula:
		return "Validates that a formula is syntactically correct and executable"
	case ValidationTypeCustom:
		return "Validates using custom validation logic"
	default:
		return "Unknown validation type"
	}
}

// RequiresCondition checks if the ValidationType requires a condition
func (vt ValidationType) RequiresCondition() bool {
	switch vt {
	case ValidationTypeRange:
		return true
	case ValidationTypeDependency:
		return true
	case ValidationTypeFormula:
		return true
	case ValidationTypeCustom:
		return true
	default:
		return false
	}
}

// SupportsAsync checks if the ValidationType supports asynchronous validation
func (vt ValidationType) SupportsAsync() bool {
	switch vt {
	case ValidationTypeRange:
		return false
	case ValidationTypeDependency:
		return true
	case ValidationTypeFormula:
		return true
	case ValidationTypeCustom:
		return true
	default:
		return false
	}
}

// SupportsCaching checks if the ValidationType supports caching
func (vt ValidationType) SupportsCaching() bool {
	switch vt {
	case ValidationTypeRange:
		return true
	case ValidationTypeDependency:
		return true
	case ValidationTypeFormula:
		return true
	case ValidationTypeCustom:
		return false
	default:
		return false
	}
}

// GetAllValidationTypes returns all valid ValidationType values
func GetAllValidationTypes() []ValidationType {
	return []ValidationType{
		ValidationTypeRange,
		ValidationTypeDependency,
		ValidationTypeFormula,
		ValidationTypeCustom,
	}
}

// GetValidationTypeFromString converts a string to ValidationType
func GetValidationTypeFromString(s string) (ValidationType, bool) {
	switch s {
	case "range":
		return ValidationTypeRange, true
	case "dependency":
		return ValidationTypeDependency, true
	case "formula":
		return ValidationTypeFormula, true
	case "custom":
		return ValidationTypeCustom, true
	default:
		return ValidationTypeRange, false
	}
}
