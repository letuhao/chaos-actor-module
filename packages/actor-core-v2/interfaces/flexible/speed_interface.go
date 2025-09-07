package flexible

import "context"

// SpeedInterface defines the interface for speed calculations
type SpeedInterface interface {
	// CalculateSpeed calculates speed for a specific category and type
	CalculateSpeed(category, speedType string, primaryStats map[string]int64) float64

	// CalculateSpeedWithContext calculates speed with context
	CalculateSpeedWithContext(ctx context.Context, category, speedType string, primaryStats map[string]int64) float64

	// GetSpeedCategories returns all available speed categories
	GetSpeedCategories() []string

	// GetSpeedTypes returns all speed types for a category
	GetSpeedTypes(category string) []string

	// AddSpeedType adds a new speed type
	AddSpeedType(category, speedType string, formula string) error

	// RemoveSpeedType removes a speed type
	RemoveSpeedType(category, speedType string) error

	// UpdateSpeedFormula updates the formula for a speed type
	UpdateSpeedFormula(category, speedType string, formula string) error

	// GetSpeedFormula returns the formula for a speed type
	GetSpeedFormula(category, speedType string) (string, error)

	// CalculateAllSpeeds calculates all speeds for given primary stats
	CalculateAllSpeeds(primaryStats map[string]int64) map[string]map[string]float64

	// CalculateAllSpeedsWithContext calculates all speeds with context
	CalculateAllSpeedsWithContext(ctx context.Context, primaryStats map[string]int64) map[string]map[string]float64

	// GetSpeedTalentBonuses returns talent bonuses for speed calculations
	GetSpeedTalentBonuses() map[string]map[string]float64

	// SetSpeedTalentBonus sets a talent bonus for speed calculation
	SetSpeedTalentBonus(talent, speedCategory string, bonus float64) error

	// GetSpeedTalentBonus gets a talent bonus for speed calculation
	GetSpeedTalentBonus(talent, speedCategory string) (float64, error)
}

// SpeedCalculationResult represents the result of a speed calculation
type SpeedCalculationResult struct {
	Category    string  `json:"category"`
	SpeedType   string  `json:"speed_type"`
	Value       float64 `json:"value"`
	BaseValue   float64 `json:"base_value"`
	TalentBonus float64 `json:"talent_bonus"`
	FinalValue  float64 `json:"final_value"`
	Formula     string  `json:"formula"`
	Timestamp   int64   `json:"timestamp"`
}

// SpeedTalentBonus represents a talent bonus for speed
type SpeedTalentBonus struct {
	Talent        string  `json:"talent"`
	SpeedCategory string  `json:"speed_category"`
	Bonus         float64 `json:"bonus"`
	Multiplier    float64 `json:"multiplier"`
	IsActive      bool    `json:"is_active"`
}
