package core

import (
	"time"
)

// PrimaryCore represents the primary core stats of an actor
type PrimaryCore struct {
	// Basic Stats
	Vitality     int64 `json:"vitality"`
	Endurance    int64 `json:"endurance"`
	Constitution int64 `json:"constitution"`
	Intelligence int64 `json:"intelligence"`
	Wisdom       int64 `json:"wisdom"`
	Charisma     int64 `json:"charisma"`
	Willpower    int64 `json:"willpower"`
	Luck         int64 `json:"luck"`
	Fate         int64 `json:"fate"`
	Karma        int64 `json:"karma"`

	// Physical Stats
	Strength    int64 `json:"strength"`
	Agility     int64 `json:"agility"`
	Personality int64 `json:"personality"`

	// Universal Cultivation Stats
	SpiritualEnergy    int64 `json:"spiritual_energy"`
	PhysicalEnergy     int64 `json:"physical_energy"`
	MentalEnergy       int64 `json:"mental_energy"`
	CultivationLevel   int64 `json:"cultivation_level"`
	BreakthroughPoints int64 `json:"breakthrough_points"`

	// Life Stats
	LifeSpan int64 `json:"life_span"`
	Age      int64 `json:"age"`

	// Metadata
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	Version   int64 `json:"version"`
}

// NewPrimaryCore creates a new PrimaryCore with default values
func NewPrimaryCore() *PrimaryCore {
	now := time.Now().Unix()
	return &PrimaryCore{
		// Basic Stats - default to 10
		Vitality:     10,
		Endurance:    10,
		Constitution: 10,
		Intelligence: 10,
		Wisdom:       10,
		Charisma:     10,
		Willpower:    10,
		Luck:         10,
		Fate:         10,
		Karma:        0, // Karma starts at 0

		// Physical Stats - default to 10
		Strength:    10,
		Agility:     10,
		Personality: 10,

		// Universal Cultivation Stats - start at 0
		SpiritualEnergy:    0,
		PhysicalEnergy:     0,
		MentalEnergy:       0,
		CultivationLevel:   0,
		BreakthroughPoints: 0,

		// Life Stats
		LifeSpan: 100, // Default lifespan
		Age:      0,   // Start at age 0

		// Metadata
		CreatedAt: now,
		UpdatedAt: now,
		Version:   1,
	}
}

// NewPrimaryCoreWithValues creates a new PrimaryCore with specific values
func NewPrimaryCoreWithValues(values map[string]int64) *PrimaryCore {
	pc := NewPrimaryCore()

	if vitality, exists := values["vitality"]; exists {
		pc.Vitality = vitality
	}
	if endurance, exists := values["endurance"]; exists {
		pc.Endurance = endurance
	}
	if constitution, exists := values["constitution"]; exists {
		pc.Constitution = constitution
	}
	if intelligence, exists := values["intelligence"]; exists {
		pc.Intelligence = intelligence
	}
	if wisdom, exists := values["wisdom"]; exists {
		pc.Wisdom = wisdom
	}
	if charisma, exists := values["charisma"]; exists {
		pc.Charisma = charisma
	}
	if willpower, exists := values["willpower"]; exists {
		pc.Willpower = willpower
	}
	if luck, exists := values["luck"]; exists {
		pc.Luck = luck
	}
	if fate, exists := values["fate"]; exists {
		pc.Fate = fate
	}
	if karma, exists := values["karma"]; exists {
		pc.Karma = karma
	}
	if strength, exists := values["strength"]; exists {
		pc.Strength = strength
	}
	if agility, exists := values["agility"]; exists {
		pc.Agility = agility
	}
	if personality, exists := values["personality"]; exists {
		pc.Personality = personality
	}
	if spiritualEnergy, exists := values["spiritual_energy"]; exists {
		pc.SpiritualEnergy = spiritualEnergy
	}
	if physicalEnergy, exists := values["physical_energy"]; exists {
		pc.PhysicalEnergy = physicalEnergy
	}
	if mentalEnergy, exists := values["mental_energy"]; exists {
		pc.MentalEnergy = mentalEnergy
	}
	if cultivationLevel, exists := values["cultivation_level"]; exists {
		pc.CultivationLevel = cultivationLevel
	}
	if breakthroughPoints, exists := values["breakthrough_points"]; exists {
		pc.BreakthroughPoints = breakthroughPoints
	}
	if lifeSpan, exists := values["life_span"]; exists {
		pc.LifeSpan = lifeSpan
	}
	if age, exists := values["age"]; exists {
		pc.Age = age
	}

	pc.UpdatedAt = time.Now().Unix()
	return pc
}

// GetStat gets a stat value by name
func (pc *PrimaryCore) GetStat(statName string) (int64, error) {
	switch statName {
	case "vitality":
		return pc.Vitality, nil
	case "endurance":
		return pc.Endurance, nil
	case "constitution":
		return pc.Constitution, nil
	case "intelligence":
		return pc.Intelligence, nil
	case "wisdom":
		return pc.Wisdom, nil
	case "charisma":
		return pc.Charisma, nil
	case "willpower":
		return pc.Willpower, nil
	case "luck":
		return pc.Luck, nil
	case "fate":
		return pc.Fate, nil
	case "karma":
		return pc.Karma, nil
	case "strength":
		return pc.Strength, nil
	case "agility":
		return pc.Agility, nil
	case "personality":
		return pc.Personality, nil
	case "spiritual_energy":
		return pc.SpiritualEnergy, nil
	case "physical_energy":
		return pc.PhysicalEnergy, nil
	case "mental_energy":
		return pc.MentalEnergy, nil
	case "cultivation_level":
		return pc.CultivationLevel, nil
	case "breakthrough_points":
		return pc.BreakthroughPoints, nil
	case "life_span":
		return pc.LifeSpan, nil
	case "age":
		return pc.Age, nil
	default:
		return 0, ErrStatNotFound
	}
}

// SetStat sets a stat value by name
func (pc *PrimaryCore) SetStat(statName string, value int64) error {
	switch statName {
	case "vitality":
		pc.Vitality = value
	case "endurance":
		pc.Endurance = value
	case "constitution":
		pc.Constitution = value
	case "intelligence":
		pc.Intelligence = value
	case "wisdom":
		pc.Wisdom = value
	case "charisma":
		pc.Charisma = value
	case "willpower":
		pc.Willpower = value
	case "luck":
		pc.Luck = value
	case "fate":
		pc.Fate = value
	case "karma":
		pc.Karma = value
	case "strength":
		pc.Strength = value
	case "agility":
		pc.Agility = value
	case "personality":
		pc.Personality = value
	case "spiritual_energy":
		pc.SpiritualEnergy = value
	case "physical_energy":
		pc.PhysicalEnergy = value
	case "mental_energy":
		pc.MentalEnergy = value
	case "cultivation_level":
		pc.CultivationLevel = value
	case "breakthrough_points":
		pc.BreakthroughPoints = value
	case "life_span":
		pc.LifeSpan = value
	case "age":
		pc.Age = value
	default:
		return ErrStatNotFound
	}

	pc.UpdatedAt = time.Now().Unix()
	pc.Version++
	return nil
}

// GetAllStats returns all primary stats as a map
func (pc *PrimaryCore) GetAllStats() map[string]int64 {
	return map[string]int64{
		"vitality":            pc.Vitality,
		"endurance":           pc.Endurance,
		"constitution":        pc.Constitution,
		"intelligence":        pc.Intelligence,
		"wisdom":              pc.Wisdom,
		"charisma":            pc.Charisma,
		"willpower":           pc.Willpower,
		"luck":                pc.Luck,
		"fate":                pc.Fate,
		"karma":               pc.Karma,
		"strength":            pc.Strength,
		"agility":             pc.Agility,
		"personality":         pc.Personality,
		"spiritual_energy":    pc.SpiritualEnergy,
		"physical_energy":     pc.PhysicalEnergy,
		"mental_energy":       pc.MentalEnergy,
		"cultivation_level":   pc.CultivationLevel,
		"breakthrough_points": pc.BreakthroughPoints,
		"life_span":           pc.LifeSpan,
		"age":                 pc.Age,
	}
}

// UpdateStats updates multiple stats at once
func (pc *PrimaryCore) UpdateStats(stats map[string]int64) error {
	for statName, value := range stats {
		if err := pc.SetStat(statName, value); err != nil {
			return err
		}
	}
	return nil
}

// Clone creates a deep copy of the PrimaryCore
func (pc *PrimaryCore) Clone() *PrimaryCore {
	return &PrimaryCore{
		Vitality:           pc.Vitality,
		Endurance:          pc.Endurance,
		Constitution:       pc.Constitution,
		Intelligence:       pc.Intelligence,
		Wisdom:             pc.Wisdom,
		Charisma:           pc.Charisma,
		Willpower:          pc.Willpower,
		Luck:               pc.Luck,
		Fate:               pc.Fate,
		Karma:              pc.Karma,
		Strength:           pc.Strength,
		Agility:            pc.Agility,
		Personality:        pc.Personality,
		SpiritualEnergy:    pc.SpiritualEnergy,
		PhysicalEnergy:     pc.PhysicalEnergy,
		MentalEnergy:       pc.MentalEnergy,
		CultivationLevel:   pc.CultivationLevel,
		BreakthroughPoints: pc.BreakthroughPoints,
		LifeSpan:           pc.LifeSpan,
		Age:                pc.Age,
		CreatedAt:          pc.CreatedAt,
		UpdatedAt:          pc.UpdatedAt,
		Version:            pc.Version,
	}
}

// Reset resets all stats to default values
func (pc *PrimaryCore) Reset() {
	now := time.Now().Unix()

	// Reset to default values
	pc.Vitality = 10
	pc.Endurance = 10
	pc.Constitution = 10
	pc.Intelligence = 10
	pc.Wisdom = 10
	pc.Charisma = 10
	pc.Willpower = 10
	pc.Luck = 10
	pc.Fate = 10
	pc.Karma = 0
	pc.Strength = 10
	pc.Agility = 10
	pc.Personality = 10
	pc.SpiritualEnergy = 0
	pc.PhysicalEnergy = 0
	pc.MentalEnergy = 0
	pc.CultivationLevel = 0
	pc.BreakthroughPoints = 0
	pc.LifeSpan = 100
	pc.Age = 0

	pc.UpdatedAt = now
	pc.Version++
}

// Validate validates the PrimaryCore stats
func (pc *PrimaryCore) Validate() []ValidationError {
	var errors []ValidationError

	// Validate basic stats (should be >= 0)
	basicStats := map[string]int64{
		"vitality":     pc.Vitality,
		"endurance":    pc.Endurance,
		"constitution": pc.Constitution,
		"intelligence": pc.Intelligence,
		"wisdom":       pc.Wisdom,
		"charisma":     pc.Charisma,
		"willpower":    pc.Willpower,
		"luck":         pc.Luck,
		"fate":         pc.Fate,
	}

	for statName, value := range basicStats {
		if value < 0 {
			errors = append(errors, ValidationError{
				Field:     statName,
				Message:   "Value cannot be negative",
				Severity:  "error",
				Timestamp: time.Now().Unix(),
			})
		}
	}

	// Validate physical stats (should be >= 0)
	physicalStats := map[string]int64{
		"strength":    pc.Strength,
		"agility":     pc.Agility,
		"personality": pc.Personality,
	}

	for statName, value := range physicalStats {
		if value < 0 {
			errors = append(errors, ValidationError{
				Field:     statName,
				Message:   "Value cannot be negative",
				Severity:  "error",
				Timestamp: time.Now().Unix(),
			})
		}
	}

	// Validate cultivation stats (should be >= 0)
	cultivationStats := map[string]int64{
		"spiritual_energy":    pc.SpiritualEnergy,
		"physical_energy":     pc.PhysicalEnergy,
		"mental_energy":       pc.MentalEnergy,
		"cultivation_level":   pc.CultivationLevel,
		"breakthrough_points": pc.BreakthroughPoints,
	}

	for statName, value := range cultivationStats {
		if value < 0 {
			errors = append(errors, ValidationError{
				Field:     statName,
				Message:   "Value cannot be negative",
				Severity:  "error",
				Timestamp: time.Now().Unix(),
			})
		}
	}

	// Validate life stats
	if pc.LifeSpan <= 0 {
		errors = append(errors, ValidationError{
			Field:     "life_span",
			Message:   "Life span must be positive",
			Severity:  "error",
			Timestamp: time.Now().Unix(),
		})
	}

	if pc.Age < 0 {
		errors = append(errors, ValidationError{
			Field:     "age",
			Message:   "Age cannot be negative",
			Severity:  "error",
			Timestamp: time.Now().Unix(),
		})
	}

	if pc.Age > pc.LifeSpan {
		errors = append(errors, ValidationError{
			Field:     "age",
			Message:   "Age cannot exceed life span",
			Severity:  "warning",
			Timestamp: time.Now().Unix(),
		})
	}

	return errors
}

// GetBasicStats returns only the basic stats
func (pc *PrimaryCore) GetBasicStats() map[string]int64 {
	return map[string]int64{
		"vitality":     pc.Vitality,
		"endurance":    pc.Endurance,
		"constitution": pc.Constitution,
		"intelligence": pc.Intelligence,
		"wisdom":       pc.Wisdom,
		"charisma":     pc.Charisma,
		"willpower":    pc.Willpower,
		"luck":         pc.Luck,
		"fate":         pc.Fate,
		"karma":        pc.Karma,
	}
}

// GetPhysicalStats returns only the physical stats
func (pc *PrimaryCore) GetPhysicalStats() map[string]int64 {
	return map[string]int64{
		"strength":    pc.Strength,
		"agility":     pc.Agility,
		"personality": pc.Personality,
	}
}

// GetCultivationStats returns only the cultivation stats
func (pc *PrimaryCore) GetCultivationStats() map[string]int64 {
	return map[string]int64{
		"spiritual_energy":    pc.SpiritualEnergy,
		"physical_energy":     pc.PhysicalEnergy,
		"mental_energy":       pc.MentalEnergy,
		"cultivation_level":   pc.CultivationLevel,
		"breakthrough_points": pc.BreakthroughPoints,
	}
}

// GetLifeStats returns only the life stats
func (pc *PrimaryCore) GetLifeStats() map[string]int64 {
	return map[string]int64{
		"life_span": pc.LifeSpan,
		"age":       pc.Age,
	}
}

// IsAlive checks if the actor is still alive
func (pc *PrimaryCore) IsAlive() bool {
	return pc.Age < pc.LifeSpan
}

// GetRemainingLife returns the remaining life
func (pc *PrimaryCore) GetRemainingLife() int64 {
	if pc.Age >= pc.LifeSpan {
		return 0
	}
	return pc.LifeSpan - pc.Age
}

// AgeUp increases the age by 1
func (pc *PrimaryCore) AgeUp() {
	pc.Age++
	pc.UpdatedAt = time.Now().Unix()
	pc.Version++
}

// AgeUpBy increases the age by the specified amount
func (pc *PrimaryCore) AgeUpBy(amount int64) {
	pc.Age += amount
	pc.UpdatedAt = time.Now().Unix()
	pc.Version++
}

// GetVersion returns the current version
func (pc *PrimaryCore) GetVersion() int64 {
	return pc.Version
}

// GetUpdatedAt returns the last update timestamp
func (pc *PrimaryCore) GetUpdatedAt() int64 {
	return pc.UpdatedAt
}

// GetCreatedAt returns the creation timestamp
func (pc *PrimaryCore) GetCreatedAt() int64 {
	return pc.CreatedAt
}
