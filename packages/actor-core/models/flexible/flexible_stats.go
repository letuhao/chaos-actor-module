package flexible

import (
	"encoding/json"
	"time"
)

// FlexibleStats represents flexible stats that can be shared across systems
type FlexibleStats struct {
	// Custom Primary Stats (int64)
	CustomPrimary map[string]int64 `json:"custom_primary"`

	// Custom Derived Stats (float64)
	CustomDerived map[string]float64 `json:"custom_derived"`

	// Sub-System Stats (systemName -> statName -> value)
	SubSystemStats map[string]map[string]float64 `json:"sub_system_stats"`

	// Metadata
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	Version   int64 `json:"version"`
}

// NewFlexibleStats creates a new FlexibleStats instance
func NewFlexibleStats() *FlexibleStats {
	now := time.Now().Unix()
	return &FlexibleStats{
		CustomPrimary:  make(map[string]int64),
		CustomDerived:  make(map[string]float64),
		SubSystemStats: make(map[string]map[string]float64),
		CreatedAt:      now,
		UpdatedAt:      now,
		Version:        1,
	}
}

// SetCustomPrimary sets a custom primary stat
func (fs *FlexibleStats) SetCustomPrimary(statName string, value int64) {
	fs.CustomPrimary[statName] = value
	fs.UpdatedAt = time.Now().Unix()
	fs.Version++
}

// GetCustomPrimary gets a custom primary stat
func (fs *FlexibleStats) GetCustomPrimary(statName string) (int64, bool) {
	value, exists := fs.CustomPrimary[statName]
	return value, exists
}

// SetCustomDerived sets a custom derived stat
func (fs *FlexibleStats) SetCustomDerived(statName string, value float64) {
	fs.CustomDerived[statName] = value
	fs.UpdatedAt = time.Now().Unix()
	fs.Version++
}

// GetCustomDerived gets a custom derived stat
func (fs *FlexibleStats) GetCustomDerived(statName string) (float64, bool) {
	value, exists := fs.CustomDerived[statName]
	return value, exists
}

// SetSubSystemStat sets a sub-system stat
func (fs *FlexibleStats) SetSubSystemStat(systemName, statName string, value float64) {
	if fs.SubSystemStats[systemName] == nil {
		fs.SubSystemStats[systemName] = make(map[string]float64)
	}
	fs.SubSystemStats[systemName][statName] = value
	fs.UpdatedAt = time.Now().Unix()
	fs.Version++
}

// GetSubSystemStat gets a sub-system stat
func (fs *FlexibleStats) GetSubSystemStat(systemName, statName string) (float64, bool) {
	if systemStats, exists := fs.SubSystemStats[systemName]; exists {
		value, exists := systemStats[statName]
		return value, exists
	}
	return 0.0, false
}

// GetAllSubSystemStats gets all stats for a specific sub-system
func (fs *FlexibleStats) GetAllSubSystemStats(systemName string) (map[string]float64, bool) {
	stats, exists := fs.SubSystemStats[systemName]
	if !exists {
		return nil, false
	}

	// Return a copy to prevent external modification
	result := make(map[string]float64)
	for k, v := range stats {
		result[k] = v
	}
	return result, true
}

// RemoveCustomPrimary removes a custom primary stat
func (fs *FlexibleStats) RemoveCustomPrimary(statName string) {
	if _, exists := fs.CustomPrimary[statName]; exists {
		delete(fs.CustomPrimary, statName)
		fs.UpdatedAt = time.Now().Unix()
		fs.Version++
	}
}

// RemoveCustomDerived removes a custom derived stat
func (fs *FlexibleStats) RemoveCustomDerived(statName string) {
	if _, exists := fs.CustomDerived[statName]; exists {
		delete(fs.CustomDerived, statName)
		fs.UpdatedAt = time.Now().Unix()
		fs.Version++
	}
}

// RemoveSubSystemStat removes a sub-system stat
func (fs *FlexibleStats) RemoveSubSystemStat(systemName, statName string) {
	if systemStats, exists := fs.SubSystemStats[systemName]; exists {
		delete(systemStats, statName)
		if len(systemStats) == 0 {
			delete(fs.SubSystemStats, systemName)
		}
		fs.UpdatedAt = time.Now().Unix()
		fs.Version++
	}
}

// RemoveSubSystem removes an entire sub-system
func (fs *FlexibleStats) RemoveSubSystem(systemName string) {
	if _, exists := fs.SubSystemStats[systemName]; exists {
		delete(fs.SubSystemStats, systemName)
		fs.UpdatedAt = time.Now().Unix()
		fs.Version++
	}
}

// GetAllCustomPrimary returns all custom primary stats
func (fs *FlexibleStats) GetAllCustomPrimary() map[string]int64 {
	result := make(map[string]int64)
	for k, v := range fs.CustomPrimary {
		result[k] = v
	}
	return result
}

// GetAllCustomDerived returns all custom derived stats
func (fs *FlexibleStats) GetAllCustomDerived() map[string]float64 {
	result := make(map[string]float64)
	for k, v := range fs.CustomDerived {
		result[k] = v
	}
	return result
}

// GetAllSubSystems returns all sub-system names
func (fs *FlexibleStats) GetAllSubSystems() []string {
	systems := make([]string, 0, len(fs.SubSystemStats))
	for systemName := range fs.SubSystemStats {
		systems = append(systems, systemName)
	}
	return systems
}

// GetStatsCount returns the total number of stats
func (fs *FlexibleStats) GetStatsCount() int {
	count := len(fs.CustomPrimary) + len(fs.CustomDerived)
	for _, systemStats := range fs.SubSystemStats {
		count += len(systemStats)
	}
	return count
}

// GetCustomPrimaryCount returns the number of custom primary stats
func (fs *FlexibleStats) GetCustomPrimaryCount() int {
	return len(fs.CustomPrimary)
}

// GetCustomDerivedCount returns the number of custom derived stats
func (fs *FlexibleStats) GetCustomDerivedCount() int {
	return len(fs.CustomDerived)
}

// GetSubSystemCount returns the number of sub-systems
func (fs *FlexibleStats) GetSubSystemCount() int {
	return len(fs.SubSystemStats)
}

// GetSubSystemStatsCount returns the number of stats in a specific sub-system
func (fs *FlexibleStats) GetSubSystemStatsCount(systemName string) int {
	if systemStats, exists := fs.SubSystemStats[systemName]; exists {
		return len(systemStats)
	}
	return 0
}

// HasCustomPrimary checks if a custom primary stat exists
func (fs *FlexibleStats) HasCustomPrimary(statName string) bool {
	_, exists := fs.CustomPrimary[statName]
	return exists
}

// HasCustomDerived checks if a custom derived stat exists
func (fs *FlexibleStats) HasCustomDerived(statName string) bool {
	_, exists := fs.CustomDerived[statName]
	return exists
}

// HasSubSystemStat checks if a sub-system stat exists
func (fs *FlexibleStats) HasSubSystemStat(systemName, statName string) bool {
	if systemStats, exists := fs.SubSystemStats[systemName]; exists {
		_, exists := systemStats[statName]
		return exists
	}
	return false
}

// HasSubSystem checks if a sub-system exists
func (fs *FlexibleStats) HasSubSystem(systemName string) bool {
	_, exists := fs.SubSystemStats[systemName]
	return exists
}

// ClearCustomPrimary clears all custom primary stats
func (fs *FlexibleStats) ClearCustomPrimary() {
	if len(fs.CustomPrimary) > 0 {
		fs.CustomPrimary = make(map[string]int64)
		fs.UpdatedAt = time.Now().Unix()
		fs.Version++
	}
}

// ClearCustomDerived clears all custom derived stats
func (fs *FlexibleStats) ClearCustomDerived() {
	if len(fs.CustomDerived) > 0 {
		fs.CustomDerived = make(map[string]float64)
		fs.UpdatedAt = time.Now().Unix()
		fs.Version++
	}
}

// ClearSubSystemStats clears all sub-system stats
func (fs *FlexibleStats) ClearSubSystemStats() {
	if len(fs.SubSystemStats) > 0 {
		fs.SubSystemStats = make(map[string]map[string]float64)
		fs.UpdatedAt = time.Now().Unix()
		fs.Version++
	}
}

// ClearAll clears all stats
func (fs *FlexibleStats) ClearAll() {
	fs.CustomPrimary = make(map[string]int64)
	fs.CustomDerived = make(map[string]float64)
	fs.SubSystemStats = make(map[string]map[string]float64)
	fs.UpdatedAt = time.Now().Unix()
	fs.Version++
}

// Clone creates a deep copy of FlexibleStats
func (fs *FlexibleStats) Clone() *FlexibleStats {
	clone := &FlexibleStats{
		CustomPrimary:  make(map[string]int64),
		CustomDerived:  make(map[string]float64),
		SubSystemStats: make(map[string]map[string]float64),
		CreatedAt:      fs.CreatedAt,
		UpdatedAt:      fs.UpdatedAt,
		Version:        fs.Version,
	}

	// Copy custom primary stats
	for k, v := range fs.CustomPrimary {
		clone.CustomPrimary[k] = v
	}

	// Copy custom derived stats
	for k, v := range fs.CustomDerived {
		clone.CustomDerived[k] = v
	}

	// Copy sub-system stats
	for systemName, systemStats := range fs.SubSystemStats {
		clone.SubSystemStats[systemName] = make(map[string]float64)
		for statName, value := range systemStats {
			clone.SubSystemStats[systemName][statName] = value
		}
	}

	return clone
}

// Merge merges another FlexibleStats into this one
func (fs *FlexibleStats) Merge(other *FlexibleStats) {
	// Merge custom primary stats
	for k, v := range other.CustomPrimary {
		fs.CustomPrimary[k] = v
	}

	// Merge custom derived stats
	for k, v := range other.CustomDerived {
		fs.CustomDerived[k] = v
	}

	// Merge sub-system stats
	for systemName, systemStats := range other.SubSystemStats {
		if fs.SubSystemStats[systemName] == nil {
			fs.SubSystemStats[systemName] = make(map[string]float64)
		}
		for statName, value := range systemStats {
			fs.SubSystemStats[systemName][statName] = value
		}
	}

	fs.UpdatedAt = time.Now().Unix()
	fs.Version++
}

// ToJSON converts FlexibleStats to JSON
func (fs *FlexibleStats) ToJSON() ([]byte, error) {
	return json.Marshal(fs)
}

// FromJSON creates FlexibleStats from JSON
func (fs *FlexibleStats) FromJSON(data []byte) error {
	return json.Unmarshal(data, fs)
}

// GetVersion returns the current version
func (fs *FlexibleStats) GetVersion() int64 {
	return fs.Version
}

// GetUpdatedAt returns the last update timestamp
func (fs *FlexibleStats) GetUpdatedAt() int64 {
	return fs.UpdatedAt
}

// GetCreatedAt returns the creation timestamp
func (fs *FlexibleStats) GetCreatedAt() int64 {
	return fs.CreatedAt
}

// Validate validates the FlexibleStats
func (fs *FlexibleStats) Validate() error {
	// Check for empty stat names
	for statName := range fs.CustomPrimary {
		if statName == "" {
			return &ValidationError{
				Field:   "custom_primary",
				Value:   statName,
				Message: "stat name cannot be empty",
			}
		}
	}

	for statName := range fs.CustomDerived {
		if statName == "" {
			return &ValidationError{
				Field:   "custom_derived",
				Value:   statName,
				Message: "stat name cannot be empty",
			}
		}
	}

	for systemName, systemStats := range fs.SubSystemStats {
		if systemName == "" {
			return &ValidationError{
				Field:   "sub_system_stats",
				Value:   systemName,
				Message: "system name cannot be empty",
			}
		}
		for statName := range systemStats {
			if statName == "" {
				return &ValidationError{
					Field:   "sub_system_stats",
					Value:   statName,
					Message: "stat name cannot be empty",
				}
			}
		}
	}

	return nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
