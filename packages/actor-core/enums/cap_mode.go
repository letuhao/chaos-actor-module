package enums

// CapMode represents the mode of cap contribution
type CapMode string

const (
	// CapModeBaseline represents baseline cap mode
	CapModeBaseline CapMode = "BASELINE"

	// CapModeAdditive represents additive cap mode
	CapModeAdditive CapMode = "ADDITIVE"

	// CapModeHardMax represents hard maximum cap mode
	CapModeHardMax CapMode = "HARD_MAX"

	// CapModeHardMin represents hard minimum cap mode
	CapModeHardMin CapMode = "HARD_MIN"

	// CapModeOverride represents override cap mode
	CapModeOverride CapMode = "OVERRIDE"
)

// IsValid checks if the cap mode is valid
func (cm CapMode) IsValid() bool {
	switch cm {
	case CapModeBaseline, CapModeAdditive, CapModeHardMax, CapModeHardMin, CapModeOverride:
		return true
	default:
		return false
	}
}

// String returns the string representation of the cap mode
func (cm CapMode) String() string {
	return string(cm)
}

// IsBaseline checks if the cap mode is baseline
func (cm CapMode) IsBaseline() bool {
	return cm == CapModeBaseline
}

// IsAdditive checks if the cap mode is additive
func (cm CapMode) IsAdditive() bool {
	return cm == CapModeAdditive
}

// IsHardMax checks if the cap mode is hard maximum
func (cm CapMode) IsHardMax() bool {
	return cm == CapModeHardMax
}

// IsHardMin checks if the cap mode is hard minimum
func (cm CapMode) IsHardMin() bool {
	return cm == CapModeHardMin
}

// IsOverride checks if the cap mode is override
func (cm CapMode) IsOverride() bool {
	return cm == CapModeOverride
}

// IsHardCap checks if the cap mode is a hard cap (max or min)
func (cm CapMode) IsHardCap() bool {
	return cm == CapModeHardMax || cm == CapModeHardMin
}

// GetPriority returns the priority for cap mode processing
// Lower numbers are processed first
func (cm CapMode) GetPriority() int64 {
	switch cm {
	case CapModeBaseline:
		return 1
	case CapModeAdditive:
		return 2
	case CapModeHardMax, CapModeHardMin:
		return 3
	case CapModeOverride:
		return 4
	default:
		return 5
	}
}
