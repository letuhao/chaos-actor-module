package types

import (
	"time"
)

// Actor represents an actor in the system
type Actor struct {
	// ID is the unique identifier
	ID string `json:"id"`

	// Name is the actor's name
	Name string `json:"name"`

	// Race is the actor's race
	Race string `json:"race"`

	// LifeSpan is the actor's lifespan in years
	LifeSpan int64 `json:"lifespan"`

	// Age is the actor's current age in years
	Age int64 `json:"age"`

	// CreatedAt is when the actor was created
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt is when the actor was last updated
	UpdatedAt time.Time `json:"updated_at"`

	// Version is the actor's version
	Version int64 `json:"version"`

	// Subsystems are the actor's subsystems
	Subsystems []Subsystem `json:"subsystems"`

	// Data contains additional actor data
	Data map[string]interface{} `json:"data,omitempty"`
}

// Subsystem represents a subsystem reference
type Subsystem struct {
	// SystemID is the subsystem's system ID
	SystemID string `json:"system_id"`

	// Priority is the subsystem's priority
	Priority int64 `json:"priority"`

	// Enabled indicates if the subsystem is enabled
	Enabled bool `json:"enabled"`

	// Config contains subsystem-specific configuration
	Config map[string]interface{} `json:"config,omitempty"`

	// Data contains subsystem-specific data
	Data map[string]interface{} `json:"data,omitempty"`
}

// GetSystemID returns the system ID
func (s Subsystem) GetSystemID() string {
	return s.SystemID
}

// GetPriority returns the priority
func (s Subsystem) GetPriority() int64 {
	return s.Priority
}

// IsEnabled checks if the subsystem is enabled
func (s Subsystem) IsEnabled() bool {
	return s.Enabled
}

// GetConfig returns the configuration
func (s Subsystem) GetConfig() map[string]interface{} {
	return s.Config
}

// GetData returns the data
func (s Subsystem) GetData() map[string]interface{} {
	return s.Data
}

// SetConfig sets the configuration
func (s *Subsystem) SetConfig(config map[string]interface{}) {
	s.Config = config
}

// SetData sets the data
func (s *Subsystem) SetData(data map[string]interface{}) {
	s.Data = data
}

// SetEnabled sets the enabled status
func (s *Subsystem) SetEnabled(enabled bool) {
	s.Enabled = enabled
}

// GetID returns the actor's ID
func (a *Actor) GetID() string {
	return a.ID
}

// GetName returns the actor's name
func (a *Actor) GetName() string {
	return a.Name
}

// GetRace returns the actor's race
func (a *Actor) GetRace() string {
	return a.Race
}

// GetLifeSpan returns the actor's lifespan
func (a *Actor) GetLifeSpan() int64 {
	return a.LifeSpan
}

// GetAge returns the actor's age
func (a *Actor) GetAge() int64 {
	return a.Age
}

// GetCreatedAt returns the creation timestamp
func (a *Actor) GetCreatedAt() time.Time {
	return a.CreatedAt
}

// GetUpdatedAt returns the last update timestamp
func (a *Actor) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

// GetVersion returns the actor's version
func (a *Actor) GetVersion() int64 {
	return a.Version
}

// GetSubsystems returns the actor's subsystems
func (a *Actor) GetSubsystems() []Subsystem {
	return a.Subsystems
}

// GetData returns the actor's data
func (a *Actor) GetData() map[string]interface{} {
	return a.Data
}

// SetName sets the actor's name
func (a *Actor) SetName(name string) {
	a.Name = name
}

// SetRace sets the actor's race
func (a *Actor) SetRace(race string) {
	a.Race = race
}

// SetLifeSpan sets the actor's lifespan
func (a *Actor) SetLifeSpan(lifespan int64) {
	a.LifeSpan = lifespan
}

// SetAge sets the actor's age
func (a *Actor) SetAge(age int64) {
	a.Age = age
}

// SetUpdatedAt sets the last update timestamp
func (a *Actor) SetUpdatedAt(updatedAt time.Time) {
	a.UpdatedAt = updatedAt
}

// SetVersion sets the actor's version
func (a *Actor) SetVersion(version int64) {
	a.Version = version
}

// SetSubsystems sets the actor's subsystems
func (a *Actor) SetSubsystems(subsystems []Subsystem) {
	a.Subsystems = subsystems
}

// SetData sets the actor's data
func (a *Actor) SetData(data map[string]interface{}) {
	a.Data = data
}

// AddSubsystem adds a subsystem to the actor
func (a *Actor) AddSubsystem(subsystem Subsystem) {
	a.Subsystems = append(a.Subsystems, subsystem)
}

// RemoveSubsystem removes a subsystem from the actor
func (a *Actor) RemoveSubsystem(systemID string) {
	for i, subsystem := range a.Subsystems {
		if subsystem.SystemID == systemID {
			a.Subsystems = append(a.Subsystems[:i], a.Subsystems[i+1:]...)
			break
		}
	}
}

// GetSubsystem returns a subsystem by system ID
func (a *Actor) GetSubsystem(systemID string) (Subsystem, bool) {
	for _, subsystem := range a.Subsystems {
		if subsystem.SystemID == systemID {
			return subsystem, true
		}
	}
	return Subsystem{}, false
}

// HasSubsystem checks if the actor has a subsystem
func (a *Actor) HasSubsystem(systemID string) bool {
	_, exists := a.GetSubsystem(systemID)
	return exists
}

// IsValid checks if the actor is valid
func (a *Actor) IsValid() bool {
	if a.ID == "" {
		return false
	}
	if a.Name == "" {
		return false
	}
	if a.Version < 0 {
		return false
	}
	return true
}

// UpdateVersion increments the actor's version
func (a *Actor) UpdateVersion() {
	a.Version++
	a.UpdatedAt = time.Now()
}

// GetGuildID returns the guild ID from the actor's data
func (a *Actor) GetGuildID() string {
	if guildID, exists := a.Data["guild_id"]; exists {
		if id, ok := guildID.(string); ok {
			return id
		}
	}
	return ""
}

// SetGuildID sets the guild ID in the actor's data
func (a *Actor) SetGuildID(guildID string) {
	if a.Data == nil {
		a.Data = make(map[string]interface{})
	}
	a.Data["guild_id"] = guildID
}

// IsInCombat checks if the actor is in combat
func (a *Actor) IsInCombat() bool {
	if inCombat, exists := a.Data["in_combat"]; exists {
		if combat, ok := inCombat.(bool); ok {
			return combat
		}
	}
	return false
}

// SetInCombat sets the combat status
func (a *Actor) SetInCombat(inCombat bool) {
	if a.Data == nil {
		a.Data = make(map[string]interface{})
	}
	a.Data["in_combat"] = inCombat
}

// HasBuff checks if the actor has a specific buff
func (a *Actor) HasBuff(buffName string) bool {
	if buffs, exists := a.Data["buffs"]; exists {
		if buffList, ok := buffs.([]string); ok {
			for _, buff := range buffList {
				if buff == buffName {
					return true
				}
			}
		}
	}
	return false
}

// AddBuff adds a buff to the actor
func (a *Actor) AddBuff(buffName string) {
	if a.Data == nil {
		a.Data = make(map[string]interface{})
	}

	var buffs []string
	if existingBuffs, exists := a.Data["buffs"]; exists {
		if buffList, ok := existingBuffs.([]string); ok {
			buffs = buffList
		}
	}

	// Check if buff already exists
	for _, buff := range buffs {
		if buff == buffName {
			return
		}
	}

	buffs = append(buffs, buffName)
	a.Data["buffs"] = buffs
}

// RemoveBuff removes a buff from the actor
func (a *Actor) RemoveBuff(buffName string) {
	if buffs, exists := a.Data["buffs"]; exists {
		if buffList, ok := buffs.([]string); ok {
			for i, buff := range buffList {
				if buff == buffName {
					a.Data["buffs"] = append(buffList[:i], buffList[i+1:]...)
					break
				}
			}
		}
	}
}
