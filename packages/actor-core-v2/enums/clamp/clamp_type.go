package clamp

// ClampType represents the type of a clamp
type ClampType int

const (
	// ClampTypeMin represents a minimum value clamp
	ClampTypeMin ClampType = iota

	// ClampTypeMax represents a maximum value clamp
	ClampTypeMax

	// ClampTypeRange represents a range clamp (min and max)
	ClampTypeRange

	// ClampTypeSoftCap represents a soft cap clamp
	ClampTypeSoftCap
)

// String returns the string representation of ClampType
func (ct ClampType) String() string {
	switch ct {
	case ClampTypeMin:
		return "min"
	case ClampTypeMax:
		return "max"
	case ClampTypeRange:
		return "range"
	case ClampTypeSoftCap:
		return "soft_cap"
	default:
		return "unknown"
	}
}

// IsValid checks if the ClampType is valid
func (ct ClampType) IsValid() bool {
	return ct >= ClampTypeMin && ct <= ClampTypeSoftCap
}

// GetDisplayName returns the display name of ClampType
func (ct ClampType) GetDisplayName() string {
	switch ct {
	case ClampTypeMin:
		return "Minimum Value Clamp"
	case ClampTypeMax:
		return "Maximum Value Clamp"
	case ClampTypeRange:
		return "Range Clamp"
	case ClampTypeSoftCap:
		return "Soft Cap Clamp"
	default:
		return "Unknown Clamp Type"
	}
}

// GetDescription returns the description of ClampType
func (ct ClampType) GetDescription() string {
	switch ct {
	case ClampTypeMin:
		return "Ensures a value never goes below a specified minimum"
	case ClampTypeMax:
		return "Ensures a value never goes above a specified maximum"
	case ClampTypeRange:
		return "Ensures a value stays within a specified range (min and max)"
	case ClampTypeSoftCap:
		return "Applies diminishing returns after a certain threshold"
	default:
		return "Unknown clamp type"
	}
}

// RequiresMinValue checks if the ClampType requires a minimum value
func (ct ClampType) RequiresMinValue() bool {
	switch ct {
	case ClampTypeMin:
		return true
	case ClampTypeMax:
		return false
	case ClampTypeRange:
		return true
	case ClampTypeSoftCap:
		return true
	default:
		return false
	}
}

// RequiresMaxValue checks if the ClampType requires a maximum value
func (ct ClampType) RequiresMaxValue() bool {
	switch ct {
	case ClampTypeMin:
		return false
	case ClampTypeMax:
		return true
	case ClampTypeRange:
		return true
	case ClampTypeSoftCap:
		return false
	default:
		return false
	}
}

// RequiresSoftCapValue checks if the ClampType requires a soft cap value
func (ct ClampType) RequiresSoftCapValue() bool {
	switch ct {
	case ClampTypeMin:
		return false
	case ClampTypeMax:
		return false
	case ClampTypeRange:
		return false
	case ClampTypeSoftCap:
		return true
	default:
		return false
	}
}

// RequiresSoftCapRate checks if the ClampType requires a soft cap rate
func (ct ClampType) RequiresSoftCapRate() bool {
	switch ct {
	case ClampTypeMin:
		return false
	case ClampTypeMax:
		return false
	case ClampTypeRange:
		return false
	case ClampTypeSoftCap:
		return true
	default:
		return false
	}
}

// GetAllClampTypes returns all valid ClampType values
func GetAllClampTypes() []ClampType {
	return []ClampType{
		ClampTypeMin,
		ClampTypeMax,
		ClampTypeRange,
		ClampTypeSoftCap,
	}
}

// GetClampTypeFromString converts a string to ClampType
func GetClampTypeFromString(s string) (ClampType, bool) {
	switch s {
	case "min":
		return ClampTypeMin, true
	case "max":
		return ClampTypeMax, true
	case "range":
		return ClampTypeRange, true
	case "soft_cap":
		return ClampTypeSoftCap, true
	default:
		return ClampTypeMin, false
	}
}
