package validation

// ValidationSeverity represents the severity level of a validation error
type ValidationSeverity int

const (
	// ValidationSeverityInfo represents an informational validation message
	ValidationSeverityInfo ValidationSeverity = iota

	// ValidationSeverityWarning represents a warning validation message
	ValidationSeverityWarning

	// ValidationSeverityError represents an error validation message
	ValidationSeverityError

	// ValidationSeverityCritical represents a critical validation message
	ValidationSeverityCritical
)

// String returns the string representation of ValidationSeverity
func (vs ValidationSeverity) String() string {
	switch vs {
	case ValidationSeverityInfo:
		return "info"
	case ValidationSeverityWarning:
		return "warning"
	case ValidationSeverityError:
		return "error"
	case ValidationSeverityCritical:
		return "critical"
	default:
		return "unknown"
	}
}

// IsValid checks if the ValidationSeverity is valid
func (vs ValidationSeverity) IsValid() bool {
	return vs >= ValidationSeverityInfo && vs <= ValidationSeverityCritical
}

// GetDisplayName returns the display name of ValidationSeverity
func (vs ValidationSeverity) GetDisplayName() string {
	switch vs {
	case ValidationSeverityInfo:
		return "Information"
	case ValidationSeverityWarning:
		return "Warning"
	case ValidationSeverityError:
		return "Error"
	case ValidationSeverityCritical:
		return "Critical"
	default:
		return "Unknown Severity"
	}
}

// GetDescription returns the description of ValidationSeverity
func (vs ValidationSeverity) GetDescription() string {
	switch vs {
	case ValidationSeverityInfo:
		return "Informational message that provides additional context"
	case ValidationSeverityWarning:
		return "Warning message that indicates a potential issue"
	case ValidationSeverityError:
		return "Error message that indicates a validation failure"
	case ValidationSeverityCritical:
		return "Critical message that indicates a severe validation failure"
	default:
		return "Unknown severity level"
	}
}

// GetPriority returns the priority level of ValidationSeverity (higher number = higher priority)
func (vs ValidationSeverity) GetPriority() int {
	switch vs {
	case ValidationSeverityInfo:
		return 1
	case ValidationSeverityWarning:
		return 2
	case ValidationSeverityError:
		return 3
	case ValidationSeverityCritical:
		return 4
	default:
		return 0
	}
}

// IsHigherThan checks if this severity is higher than another
func (vs ValidationSeverity) IsHigherThan(other ValidationSeverity) bool {
	return vs.GetPriority() > other.GetPriority()
}

// IsLowerThan checks if this severity is lower than another
func (vs ValidationSeverity) IsLowerThan(other ValidationSeverity) bool {
	return vs.GetPriority() < other.GetPriority()
}

// IsEqualOrHigherThan checks if this severity is equal or higher than another
func (vs ValidationSeverity) IsEqualOrHigherThan(other ValidationSeverity) bool {
	return vs.GetPriority() >= other.GetPriority()
}

// IsEqualOrLowerThan checks if this severity is equal or lower than another
func (vs ValidationSeverity) IsEqualOrLowerThan(other ValidationSeverity) bool {
	return vs.GetPriority() <= other.GetPriority()
}

// ShouldStopExecution checks if this severity should stop execution
func (vs ValidationSeverity) ShouldStopExecution() bool {
	return vs == ValidationSeverityError || vs == ValidationSeverityCritical
}

// ShouldLog checks if this severity should be logged
func (vs ValidationSeverity) ShouldLog() bool {
	return vs != ValidationSeverityInfo
}

// ShouldAlert checks if this severity should trigger an alert
func (vs ValidationSeverity) ShouldAlert() bool {
	return vs == ValidationSeverityCritical
}

// GetAllValidationSeverities returns all valid ValidationSeverity values
func GetAllValidationSeverities() []ValidationSeverity {
	return []ValidationSeverity{
		ValidationSeverityInfo,
		ValidationSeverityWarning,
		ValidationSeverityError,
		ValidationSeverityCritical,
	}
}

// GetValidationSeverityFromString converts a string to ValidationSeverity
func GetValidationSeverityFromString(s string) (ValidationSeverity, bool) {
	switch s {
	case "info":
		return ValidationSeverityInfo, true
	case "warning":
		return ValidationSeverityWarning, true
	case "error":
		return ValidationSeverityError, true
	case "critical":
		return ValidationSeverityCritical, true
	default:
		return ValidationSeverityInfo, false
	}
}
