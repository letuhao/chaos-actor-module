package integration

import (
	"context"
	"errors"
	"time"

	"rpg-system/internal/model"
)

// MockMongoAdapter is a mock implementation for testing and development
type MockMongoAdapter struct {
	playerProgress  map[string]*model.PlayerProgress
	playerEffects   map[string][]model.StatModifier
	playerEquipment map[string]map[string]model.StatModifier // actorID -> slot -> item
	titlesOwned     map[string][]model.StatModifier
	contentRegistry []model.StatDef
}

// NewMockMongoAdapter creates a new mock adapter
func NewMockMongoAdapter() *MockMongoAdapter {
	return &MockMongoAdapter{
		playerProgress:  make(map[string]*model.PlayerProgress),
		playerEffects:   make(map[string][]model.StatModifier),
		playerEquipment: make(map[string]map[string]model.StatModifier),
		titlesOwned:     make(map[string][]model.StatModifier),
		contentRegistry: []model.StatDef{},
	}
}

// Close closes the mock adapter
func (m *MockMongoAdapter) Close() error {
	return nil
}

// GetPlayerProgress retrieves player progression data
func (m *MockMongoAdapter) GetPlayerProgress(ctx context.Context, actorID string) (*model.PlayerProgress, error) {
	progress, exists := m.playerProgress[actorID]
	if !exists {
		return nil, errors.New("player not found")
	}
	return progress, nil
}

// SavePlayerProgress saves player progression data
func (m *MockMongoAdapter) SavePlayerProgress(ctx context.Context, progress *model.PlayerProgress) error {
	m.playerProgress[progress.ActorID] = progress
	return nil
}

// GetActiveEffects retrieves active effects for a player
func (m *MockMongoAdapter) GetActiveEffects(ctx context.Context, actorID string) ([]model.StatModifier, error) {
	effects, exists := m.playerEffects[actorID]
	if !exists {
		return []model.StatModifier{}, nil
	}
	return effects, nil
}

// AddEffect adds a new effect to a player
func (m *MockMongoAdapter) AddEffect(ctx context.Context, actorID string, effect model.StatModifier, duration time.Duration) error {
	if m.playerEffects[actorID] == nil {
		m.playerEffects[actorID] = []model.StatModifier{}
	}
	m.playerEffects[actorID] = append(m.playerEffects[actorID], effect)
	return nil
}

// RemoveEffect removes an effect from a player
func (m *MockMongoAdapter) RemoveEffect(ctx context.Context, actorID, effectID string) error {
	effects := m.playerEffects[actorID]
	for i, effect := range effects {
		if effect.Source.ID == effectID {
			m.playerEffects[actorID] = append(effects[:i], effects[i+1:]...)
			break
		}
	}
	return nil
}

// GetEquippedItems retrieves equipped items for a player
func (m *MockMongoAdapter) GetEquippedItems(ctx context.Context, actorID string) ([]model.StatModifier, error) {
	equipment, exists := m.playerEquipment[actorID]
	if !exists {
		return []model.StatModifier{}, nil
	}

	var items []model.StatModifier
	for _, item := range equipment {
		items = append(items, item)
	}
	return items, nil
}

// EquipItem equips an item to a player
func (m *MockMongoAdapter) EquipItem(ctx context.Context, actorID, slot string, item model.StatModifier) error {
	if m.playerEquipment[actorID] == nil {
		m.playerEquipment[actorID] = make(map[string]model.StatModifier)
	}
	m.playerEquipment[actorID][slot] = item
	return nil
}

// UnequipItem unequips an item from a player
func (m *MockMongoAdapter) UnequipItem(ctx context.Context, actorID, slot string) error {
	if m.playerEquipment[actorID] != nil {
		delete(m.playerEquipment[actorID], slot)
	}
	return nil
}

// GetOwnedTitles retrieves owned titles for a player
func (m *MockMongoAdapter) GetOwnedTitles(ctx context.Context, actorID string) ([]model.StatModifier, error) {
	titles, exists := m.titlesOwned[actorID]
	if !exists {
		return []model.StatModifier{}, nil
	}
	return titles, nil
}

// GrantTitle grants a title to a player
func (m *MockMongoAdapter) GrantTitle(ctx context.Context, actorID, titleID string, title model.StatModifier) error {
	if m.titlesOwned[actorID] == nil {
		m.titlesOwned[actorID] = []model.StatModifier{}
	}
	m.titlesOwned[actorID] = append(m.titlesOwned[actorID], title)
	return nil
}

// GetStatRegistry retrieves the stat registry from the database
func (m *MockMongoAdapter) GetStatRegistry(ctx context.Context) ([]model.StatDef, error) {
	return m.contentRegistry, nil
}

// SaveStatRegistry saves stat definitions to the database
func (m *MockMongoAdapter) SaveStatRegistry(ctx context.Context, registry []model.StatDef) error {
	m.contentRegistry = registry
	return nil
}

// CleanupExpiredEffects removes expired effects from the database
func (m *MockMongoAdapter) CleanupExpiredEffects(ctx context.Context) error {
	// Mock implementation - just return success
	return nil
}

// GetPlayerStatsSummary returns a summary of all player stats from the database
func (m *MockMongoAdapter) GetPlayerStatsSummary(ctx context.Context, actorID string) (*PlayerStatsSummary, error) {
	progress, err := m.GetPlayerProgress(ctx, actorID)
	if err != nil {
		return nil, err
	}

	effects, _ := m.GetActiveEffects(ctx, actorID)
	equipment, _ := m.GetEquippedItems(ctx, actorID)
	titles, _ := m.GetOwnedTitles(ctx, actorID)

	return &PlayerStatsSummary{
		Progress:  progress,
		Effects:   effects,
		Equipment: equipment,
		Titles:    titles,
	}, nil
}

// PlayerStatsSummary contains all player data needed for stat calculation
type PlayerStatsSummary struct {
	Progress  *model.PlayerProgress
	Effects   []model.StatModifier
	Equipment []model.StatModifier
	Titles    []model.StatModifier
}
