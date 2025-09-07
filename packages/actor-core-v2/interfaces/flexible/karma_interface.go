package flexible

import "context"

// KarmaInterface defines the interface for karma calculations
type KarmaInterface interface {
	// CalculateTotalKarma calculates total karma for a division
	CalculateTotalKarma(divisionType, divisionId string) int64

	// CalculateTotalKarmaWithContext calculates total karma with context
	CalculateTotalKarmaWithContext(ctx context.Context, divisionType, divisionId string) int64

	// CalculateWeightedKarmaScore calculates weighted karma score
	CalculateWeightedKarmaScore(divisionType, divisionId string) float64

	// CalculateWeightedKarmaScoreWithContext calculates weighted karma score with context
	CalculateWeightedKarmaScoreWithContext(ctx context.Context, divisionType, divisionId string) float64

	// CalculateKarmaInfluence calculates karma influence on a stat
	CalculateKarmaInfluence(statName string) float64

	// CalculateKarmaInfluenceWithContext calculates karma influence with context
	CalculateKarmaInfluenceWithContext(ctx context.Context, statName string) float64

	// AddKarma adds karma to a division
	AddKarma(divisionType, divisionId, karmaType string, amount int64) error

	// AddKarmaWithContext adds karma with context
	AddKarmaWithContext(ctx context.Context, divisionType, divisionId, karmaType string, amount int64) error

	// RemoveKarma removes karma from a division
	RemoveKarma(divisionType, divisionId, karmaType string, amount int64) error

	// RemoveKarmaWithContext removes karma with context
	RemoveKarmaWithContext(ctx context.Context, divisionType, divisionId, karmaType string, amount int64) error

	// GetKarmaTypes returns all available karma types
	GetKarmaTypes() []string

	// GetKarmaCategories returns all available karma categories
	GetKarmaCategories() []string

	// GetKarmaTypesByCategory returns karma types for a category
	GetKarmaTypesByCategory(category string) []string

	// AddKarmaType adds a new karma type
	AddKarmaType(karmaType *KarmaType) error

	// RemoveKarmaType removes a karma type
	RemoveKarmaType(karmaType string) error

	// UpdateKarmaType updates a karma type
	UpdateKarmaType(karmaType *KarmaType) error

	// GetKarmaType gets a karma type
	GetKarmaType(karmaType string) (*KarmaType, error)

	// GetKarmaInfluence returns karma influence on stats
	GetKarmaInfluence() map[string]float64

	// SetKarmaInfluence sets karma influence on a stat
	SetKarmaInfluence(statName string, influence float64) error

	// GetDivisionKarma returns karma for a division
	GetDivisionKarma(divisionType, divisionId string) map[string]int64

	// SetDivisionKarma sets karma for a division
	SetDivisionKarma(divisionType, divisionId string, karma map[string]int64) error
}

// KarmaType represents a type of karma
type KarmaType struct {
	Name        string             `json:"name"`
	Category    string             `json:"category"`
	Description string             `json:"description"`
	Influence   map[string]float64 `json:"influence"`
	DecayRate   float64            `json:"decay_rate"`
	MaxValue    int64              `json:"max_value"`
	MinValue    int64              `json:"min_value"`
}

// KarmaCalculationResult represents the result of a karma calculation
type KarmaCalculationResult struct {
	DivisionType   string             `json:"division_type"`
	DivisionId     string             `json:"division_id"`
	TotalKarma     int64              `json:"total_karma"`
	WeightedScore  float64            `json:"weighted_score"`
	KarmaBreakdown map[string]int64   `json:"karma_breakdown"`
	Influence      map[string]float64 `json:"influence"`
	Timestamp      int64              `json:"timestamp"`
}

// KarmaChange represents a change in karma
type KarmaChange struct {
	DivisionType string `json:"division_type"`
	DivisionId   string `json:"division_id"`
	KarmaType    string `json:"karma_type"`
	Amount       int64  `json:"amount"`
	OldValue     int64  `json:"old_value"`
	NewValue     int64  `json:"new_value"`
	Timestamp    int64  `json:"timestamp"`
	Source       string `json:"source"`
}
