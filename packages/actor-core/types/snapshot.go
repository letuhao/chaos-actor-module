package types

import (
	"time"
)

// Snapshot represents the final aggregated snapshot of an actor's stats
type Snapshot struct {
	// Primary contains primary dimension values
	Primary map[string]float64 `json:"primary"`

	// Derived contains derived dimension values
	Derived map[string]float64 `json:"derived"`

	// CapsUsed contains the caps that were applied
	CapsUsed map[string]Caps `json:"caps_used"`

	// Version is the actor version when this snapshot was created
	Version int64 `json:"version"`

	// Timestamp is when the snapshot was created
	Timestamp time.Time `json:"timestamp"`

	// SubsystemsProcessed contains the subsystems that were processed
	SubsystemsProcessed []string `json:"subsystems_processed,omitempty"`

	// ProcessingTime is the time taken to process this snapshot
	ProcessingTime time.Duration `json:"processing_time,omitempty"`

	// Metadata contains additional metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// IsValid checks if the snapshot is valid
func (s *Snapshot) IsValid() bool {
	if s.Version < 0 {
		return false
	}
	if s.Primary == nil {
		return false
	}
	if s.Derived == nil {
		return false
	}
	if s.CapsUsed == nil {
		return false
	}
	return true
}

// GetPrimary returns the primary dimension values
func (s *Snapshot) GetPrimary() map[string]float64 {
	return s.Primary
}

// GetDerived returns the derived dimension values
func (s *Snapshot) GetDerived() map[string]float64 {
	return s.Derived
}

// GetCapsUsed returns the caps that were applied
func (s *Snapshot) GetCapsUsed() map[string]Caps {
	return s.CapsUsed
}

// GetVersion returns the actor version
func (s *Snapshot) GetVersion() int64 {
	return s.Version
}

// GetTimestamp returns the timestamp
func (s *Snapshot) GetTimestamp() time.Time {
	return s.Timestamp
}

// GetSubsystemsProcessed returns the subsystems that were processed
func (s *Snapshot) GetSubsystemsProcessed() []string {
	return s.SubsystemsProcessed
}

// GetProcessingTime returns the processing time
func (s *Snapshot) GetProcessingTime() time.Duration {
	return s.ProcessingTime
}

// GetMetadata returns the metadata
func (s *Snapshot) GetMetadata() map[string]interface{} {
	return s.Metadata
}

// SetPrimary sets the primary dimension values
func (s *Snapshot) SetPrimary(primary map[string]float64) {
	s.Primary = primary
}

// SetDerived sets the derived dimension values
func (s *Snapshot) SetDerived(derived map[string]float64) {
	s.Derived = derived
}

// SetCapsUsed sets the caps that were applied
func (s *Snapshot) SetCapsUsed(capsUsed map[string]Caps) {
	s.CapsUsed = capsUsed
}

// SetVersion sets the actor version
func (s *Snapshot) SetVersion(version int64) {
	s.Version = version
}

// SetTimestamp sets the timestamp
func (s *Snapshot) SetTimestamp(timestamp time.Time) {
	s.Timestamp = timestamp
}

// SetSubsystemsProcessed sets the subsystems that were processed
func (s *Snapshot) SetSubsystemsProcessed(subsystems []string) {
	s.SubsystemsProcessed = subsystems
}

// SetProcessingTime sets the processing time
func (s *Snapshot) SetProcessingTime(processingTime time.Duration) {
	s.ProcessingTime = processingTime
}

// SetMetadata sets the metadata
func (s *Snapshot) SetMetadata(metadata map[string]interface{}) {
	s.Metadata = metadata
}

// GetPrimaryValue returns a primary dimension value
func (s *Snapshot) GetPrimaryValue(dimension string) (float64, bool) {
	value, exists := s.Primary[dimension]
	return value, exists
}

// GetDerivedValue returns a derived dimension value
func (s *Snapshot) GetDerivedValue(dimension string) (float64, bool) {
	value, exists := s.Derived[dimension]
	return value, exists
}

// GetCapsForDimension returns caps for a specific dimension
func (s *Snapshot) GetCapsForDimension(dimension string) (Caps, bool) {
	caps, exists := s.CapsUsed[dimension]
	return caps, exists
}

// SetPrimaryValue sets a primary dimension value
func (s *Snapshot) SetPrimaryValue(dimension string, value float64) {
	if s.Primary == nil {
		s.Primary = make(map[string]float64)
	}
	s.Primary[dimension] = value
}

// SetDerivedValue sets a derived dimension value
func (s *Snapshot) SetDerivedValue(dimension string, value float64) {
	if s.Derived == nil {
		s.Derived = make(map[string]float64)
	}
	s.Derived[dimension] = value
}

// SetCapsForDimension sets caps for a specific dimension
func (s *Snapshot) SetCapsForDimension(dimension string, caps Caps) {
	if s.CapsUsed == nil {
		s.CapsUsed = make(map[string]Caps)
	}
	s.CapsUsed[dimension] = caps
}

// AddSubsystemProcessed adds a subsystem to the processed list
func (s *Snapshot) AddSubsystemProcessed(subsystem string) {
	if s.SubsystemsProcessed == nil {
		s.SubsystemsProcessed = make([]string, 0)
	}
	s.SubsystemsProcessed = append(s.SubsystemsProcessed, subsystem)
}

// AddMetadata adds metadata
func (s *Snapshot) AddMetadata(key string, value interface{}) {
	if s.Metadata == nil {
		s.Metadata = make(map[string]interface{})
	}
	s.Metadata[key] = value
}

// GetMetadataValue returns a metadata value
func (s *Snapshot) GetMetadataValue(key string) (interface{}, bool) {
	if s.Metadata == nil {
		return nil, false
	}
	value, exists := s.Metadata[key]
	return value, exists
}

// GetDimensions returns all dimensions in the snapshot
func (s *Snapshot) GetDimensions() []string {
	dimensions := make([]string, 0)

	// Add primary dimensions
	for dimension := range s.Primary {
		dimensions = append(dimensions, dimension)
	}

	// Add derived dimensions
	for dimension := range s.Derived {
		dimensions = append(dimensions, dimension)
	}

	return dimensions
}

// GetPrimaryDimensions returns primary dimensions
func (s *Snapshot) GetPrimaryDimensions() []string {
	dimensions := make([]string, 0, len(s.Primary))
	for dimension := range s.Primary {
		dimensions = append(dimensions, dimension)
	}
	return dimensions
}

// GetDerivedDimensions returns derived dimensions
func (s *Snapshot) GetDerivedDimensions() []string {
	dimensions := make([]string, 0, len(s.Derived))
	for dimension := range s.Derived {
		dimensions = append(dimensions, dimension)
	}
	return dimensions
}

// GetCappedDimensions returns dimensions that have caps applied
func (s *Snapshot) GetCappedDimensions() []string {
	dimensions := make([]string, 0, len(s.CapsUsed))
	for dimension := range s.CapsUsed {
		dimensions = append(dimensions, dimension)
	}
	return dimensions
}

// IsEmpty checks if the snapshot is empty
func (s *Snapshot) IsEmpty() bool {
	return len(s.Primary) == 0 && len(s.Derived) == 0
}

// GetTotalDimensions returns the total number of dimensions
func (s *Snapshot) GetTotalDimensions() int64 {
	return int64(len(s.Primary) + len(s.Derived))
}

// GetTotalCaps returns the total number of caps applied
func (s *Snapshot) GetTotalCaps() int64 {
	return int64(len(s.CapsUsed))
}

// Clone creates a deep copy of the snapshot
func (s *Snapshot) Clone() *Snapshot {
	clone := &Snapshot{
		Version:             s.Version,
		Timestamp:           s.Timestamp,
		ProcessingTime:      s.ProcessingTime,
		SubsystemsProcessed: make([]string, len(s.SubsystemsProcessed)),
	}

	// Copy primary dimensions
	clone.Primary = make(map[string]float64)
	for k, v := range s.Primary {
		clone.Primary[k] = v
	}

	// Copy derived dimensions
	clone.Derived = make(map[string]float64)
	for k, v := range s.Derived {
		clone.Derived[k] = v
	}

	// Copy caps used
	clone.CapsUsed = make(map[string]Caps)
	for k, v := range s.CapsUsed {
		clone.CapsUsed[k] = v
	}

	// Copy subsystems processed
	copy(clone.SubsystemsProcessed, s.SubsystemsProcessed)

	// Copy metadata
	if s.Metadata != nil {
		clone.Metadata = make(map[string]interface{})
		for k, v := range s.Metadata {
			clone.Metadata[k] = v
		}
	}

	return clone
}
