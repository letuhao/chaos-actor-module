package enums

import "chaos-actor-module/packages/actor-core/constants"

// ErrorType represents the type of error
type ErrorType string

const (
	// ErrorTypeValidation represents validation errors
	ErrorTypeValidation ErrorType = constants.ErrorTypeValidation

	// ErrorTypeSystem represents system errors
	ErrorTypeSystem ErrorType = constants.ErrorTypeSystem

	// ErrorTypePerformance represents performance errors
	ErrorTypePerformance ErrorType = constants.ErrorTypePerformance
)

// IsValid checks if the error type is valid
func (et ErrorType) IsValid() bool {
	switch et {
	case ErrorTypeValidation, ErrorTypeSystem, ErrorTypePerformance:
		return true
	default:
		return false
	}
}

// String returns the string representation of the error type
func (et ErrorType) String() string {
	return string(et)
}

// IsValidation checks if the error type is validation
func (et ErrorType) IsValidation() bool {
	return et == ErrorTypeValidation
}

// IsSystem checks if the error type is system
func (et ErrorType) IsSystem() bool {
	return et == ErrorTypeSystem
}

// IsPerformance checks if the error type is performance
func (et ErrorType) IsPerformance() bool {
	return et == ErrorTypePerformance
}

// GetSeverity returns the severity level for the error type
func (et ErrorType) GetSeverity() ErrorSeverity {
	switch et {
	case ErrorTypeValidation:
		return ErrorSeverityWarning
	case ErrorTypeSystem:
		return ErrorSeverityError
	case ErrorTypePerformance:
		return ErrorSeverityWarning
	default:
		return ErrorSeverityUnknown
	}
}

// ErrorSeverity represents the severity level of an error
type ErrorSeverity string

const (
	// ErrorSeverityInfo represents informational severity
	ErrorSeverityInfo ErrorSeverity = "info"

	// ErrorSeverityWarning represents warning severity
	ErrorSeverityWarning ErrorSeverity = "warning"

	// ErrorSeverityError represents error severity
	ErrorSeverityError ErrorSeverity = "error"

	// ErrorSeverityCritical represents critical severity
	ErrorSeverityCritical ErrorSeverity = "critical"

	// ErrorSeverityUnknown represents unknown severity
	ErrorSeverityUnknown ErrorSeverity = "unknown"
)

// IsValid checks if the error severity is valid
func (es ErrorSeverity) IsValid() bool {
	switch es {
	case ErrorSeverityInfo, ErrorSeverityWarning, ErrorSeverityError, ErrorSeverityCritical, ErrorSeverityUnknown:
		return true
	default:
		return false
	}
}

// String returns the string representation of the error severity
func (es ErrorSeverity) String() string {
	return string(es)
}

// IsInfo checks if the severity is info
func (es ErrorSeverity) IsInfo() bool {
	return es == ErrorSeverityInfo
}

// IsWarning checks if the severity is warning
func (es ErrorSeverity) IsWarning() bool {
	return es == ErrorSeverityWarning
}

// IsError checks if the severity is error
func (es ErrorSeverity) IsError() bool {
	return es == ErrorSeverityError
}

// IsCritical checks if the severity is critical
func (es ErrorSeverity) IsCritical() bool {
	return es == ErrorSeverityCritical
}

// IsUnknown checks if the severity is unknown
func (es ErrorSeverity) IsUnknown() bool {
	return es == ErrorSeverityUnknown
}

// GetLevel returns the numeric level for the severity
func (es ErrorSeverity) GetLevel() int64 {
	switch es {
	case ErrorSeverityInfo:
		return 1
	case ErrorSeverityWarning:
		return 2
	case ErrorSeverityError:
		return 3
	case ErrorSeverityCritical:
		return 4
	case ErrorSeverityUnknown:
		return 0
	default:
		return 0
	}
}

// IsHigherThan checks if this severity is higher than another
func (es ErrorSeverity) IsHigherThan(other ErrorSeverity) bool {
	return es.GetLevel() > other.GetLevel()
}

// IsLowerThan checks if this severity is lower than another
func (es ErrorSeverity) IsLowerThan(other ErrorSeverity) bool {
	return es.GetLevel() < other.GetLevel()
}

// IsEqual checks if this severity is equal to another
func (es ErrorSeverity) IsEqual(other ErrorSeverity) bool {
	return es.GetLevel() == other.GetLevel()
}
