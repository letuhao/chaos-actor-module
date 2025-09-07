package enums

// Priority represents the processing priority
type Priority int64

const (
	// PriorityLowest represents the lowest priority
	PriorityLowest Priority = 0

	// PriorityLow represents low priority
	PriorityLow Priority = 25

	// PriorityNormal represents normal priority
	PriorityNormal Priority = 50

	// PriorityHigh represents high priority
	PriorityHigh Priority = 75

	// PriorityHighest represents the highest priority
	PriorityHighest Priority = 100

	// PriorityCritical represents critical priority
	PriorityCritical Priority = 125

	// PrioritySystem represents system priority (highest)
	PrioritySystem Priority = 150
)

// IsValid checks if the priority is valid
func (p Priority) IsValid() bool {
	return p >= PriorityLowest && p <= PrioritySystem
}

// Int64 returns the int64 value of the priority
func (p Priority) Int64() int64 {
	return int64(p)
}

// String returns the string representation of the priority
func (p Priority) String() string {
	switch {
	case p <= PriorityLowest:
		return "lowest"
	case p <= PriorityLow:
		return "low"
	case p <= PriorityNormal:
		return "normal"
	case p <= PriorityHigh:
		return "high"
	case p <= PriorityHighest:
		return "highest"
	case p <= PriorityCritical:
		return "critical"
	case p <= PrioritySystem:
		return "system"
	default:
		return "unknown"
	}
}

// IsHigherThan checks if this priority is higher than another
func (p Priority) IsHigherThan(other Priority) bool {
	return p.Int64() > other.Int64()
}

// IsLowerThan checks if this priority is lower than another
func (p Priority) IsLowerThan(other Priority) bool {
	return p.Int64() < other.Int64()
}

// IsEqual checks if this priority is equal to another
func (p Priority) IsEqual(other Priority) bool {
	return p.Int64() == other.Int64()
}

// GetDefaultPriority returns the default priority for subsystems
func GetDefaultPriority() Priority {
	return PriorityNormal
}

// GetSystemPriority returns the system priority
func GetSystemPriority() Priority {
	return PrioritySystem
}

// GetUserPriority returns the user priority
func GetUserPriority() Priority {
	return PriorityNormal
}
