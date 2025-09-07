package integration

import (
	"errors"
	"rpg-system/internal/model"
	"rpg-system/internal/registry"
	"time"
)

// ProgressionService handles character progression including XP, leveling, and stat allocation
type ProgressionService struct {
	registry *registry.StatRegistry
	// In a real implementation, this would be a database interface
	playerProgress map[string]*model.PlayerProgress
}

// NewProgressionService creates a new progression service
func NewProgressionService() *ProgressionService {
	return &ProgressionService{
		registry:       registry.NewStatRegistry(),
		playerProgress: make(map[string]*model.PlayerProgress),
	}
}

// GrantXP grants experience points to a player and handles leveling up
func (ps *ProgressionService) GrantXP(actorID string, xpAmount int64) (*model.ProgressionResult, error) {
	if xpAmount <= 0 {
		return nil, errors.New("XP amount must be positive")
	}

	// Get or create player progress
	progress, exists := ps.playerProgress[actorID]
	if !exists {
		progress = &model.PlayerProgress{
			ActorID:     actorID,
			Level:       1,
			XP:          0,
			Allocations: make(map[model.StatKey]int64),
			LastUpdated: time.Now().Unix(),
		}
		ps.playerProgress[actorID] = progress
	}

	// Add XP
	progress.XP += xpAmount
	progress.LastUpdated = time.Now().Unix()

	result := &model.ProgressionResult{}

	// Check for level up
	oldLevel := progress.Level
	newLevel := ps.calculateLevelFromXP(progress.XP)

	if newLevel > oldLevel {
		progress.Level = newLevel
		pointsGranted := ps.calculateStatPointsForLevels(oldLevel, newLevel)
		result.NewLevel = &newLevel
		result.PointsGranted = &pointsGranted

		// In a real implementation, you might want to notify the player or trigger events
	}

	return result, nil
}

// AllocatePoints allocates stat points to a specific stat
func (ps *ProgressionService) AllocatePoints(actorID string, statKey model.StatKey, points int64) error {
	if points <= 0 {
		return errors.New("points must be positive")
	}

	if !statKey.IsPrimary() {
		return errors.New("can only allocate points to primary stats")
	}

	progress, exists := ps.playerProgress[actorID]
	if !exists {
		return errors.New("player not found")
	}

	// Check if player has enough unallocated points
	availablePoints := ps.calculateAvailablePoints(progress)
	if availablePoints < points {
		return errors.New("insufficient points available")
	}

	// Allocate points
	progress.Allocations[statKey] += points
	progress.LastUpdated = time.Now().Unix()

	return nil
}

// Respec allows a player to reset their stat allocations
func (ps *ProgressionService) Respec(actorID string) error {
	progress, exists := ps.playerProgress[actorID]
	if !exists {
		return errors.New("player not found")
	}

	// Reset all allocations
	progress.Allocations = make(map[model.StatKey]int64)
	progress.LastUpdated = time.Now().Unix()

	return nil
}

// GetPlayerProgress returns the current progress for a player
func (ps *ProgressionService) GetPlayerProgress(actorID string) (*model.PlayerProgress, error) {
	progress, exists := ps.playerProgress[actorID]
	if !exists {
		return nil, errors.New("player not found")
	}

	// Return a copy to prevent external modification
	progressCopy := *progress
	progressCopy.Allocations = make(map[model.StatKey]int64)
	for k, v := range progress.Allocations {
		progressCopy.Allocations[k] = v
	}

	return &progressCopy, nil
}

// GetAvailablePoints returns the number of unallocated stat points for a player
func (ps *ProgressionService) GetAvailablePoints(actorID string) (int64, error) {
	progress, exists := ps.playerProgress[actorID]
	if !exists {
		return 0, errors.New("player not found")
	}

	return ps.calculateAvailablePoints(progress), nil
}

// GetLevelProgression returns the progression curve for a specific level
func (ps *ProgressionService) GetLevelProgression(level int64) *model.LevelProgression {
	return &model.LevelProgression{
		Level:         level,
		XPRequired:    ps.calculateXPForLevel(level),
		PointsGranted: ps.calculateStatPointsForLevel(level),
	}
}

// calculateLevelFromXP calculates the level based on total XP
func (ps *ProgressionService) calculateLevelFromXP(totalXP int64) int64 {
	level := int64(1)
	for {
		xpForNextLevel := ps.calculateXPForLevel(level + 1)
		if totalXP < xpForNextLevel {
			break
		}
		level++
		if level >= 100 { // Max level cap
			break
		}
	}
	return level
}

// calculateXPForLevel calculates the XP required to reach a specific level
func (ps *ProgressionService) calculateXPForLevel(level int64) int64 {
	if level <= 1 {
		return 0
	}

	// Simple quadratic progression: (level-1)^2 * 100
	// Level 2: 100 XP, Level 3: 400 XP, Level 4: 900 XP, etc.
	return int64((level - 1) * (level - 1) * 100)
}

// calculateStatPointsForLevel calculates stat points granted at a specific level
func (ps *ProgressionService) calculateStatPointsForLevel(level int64) int64 {
	if level <= 1 {
		return 0
	}

	// Grant 2 points per level, with bonus points every 5 levels
	points := int64(2)
	if level%5 == 0 {
		points += 1 // Bonus point every 5 levels
	}
	return points
}

// calculateStatPointsForLevels calculates total stat points for a level range
func (ps *ProgressionService) calculateStatPointsForLevels(fromLevel, toLevel int64) int64 {
	totalPoints := int64(0)
	for level := fromLevel + 1; level <= toLevel; level++ {
		totalPoints += ps.calculateStatPointsForLevel(level)
	}
	return totalPoints
}

// calculateAvailablePoints calculates unallocated points for a player
func (ps *ProgressionService) calculateAvailablePoints(progress *model.PlayerProgress) int64 {
	totalEarned := int64(0)
	for level := int64(2); level <= progress.Level; level++ {
		totalEarned += ps.calculateStatPointsForLevel(level)
	}

	totalAllocated := int64(0)
	for _, points := range progress.Allocations {
		totalAllocated += points
	}

	return totalEarned - totalAllocated
}

// GetXPToNextLevel returns the XP needed to reach the next level
func (ps *ProgressionService) GetXPToNextLevel(actorID string) (int64, error) {
	progress, exists := ps.playerProgress[actorID]
	if !exists {
		return 0, errors.New("player not found")
	}

	nextLevel := progress.Level + 1
	if nextLevel > 100 {
		return 0, nil // Max level reached
	}

	xpForNextLevel := ps.calculateXPForLevel(nextLevel)
	return xpForNextLevel - progress.XP, nil
}

// GetLevelProgress returns the progress towards the next level (0.0 to 1.0)
func (ps *ProgressionService) GetLevelProgress(actorID string) (float64, error) {
	progress, exists := ps.playerProgress[actorID]
	if !exists {
		return 0, errors.New("player not found")
	}

	nextLevel := progress.Level + 1
	if nextLevel > 100 {
		return 1.0, nil // Max level reached
	}

	currentLevelXP := ps.calculateXPForLevel(progress.Level)
	nextLevelXP := ps.calculateXPForLevel(nextLevel)

	if nextLevelXP == currentLevelXP {
		return 1.0, nil
	}

	progressXP := progress.XP - currentLevelXP
	requiredXP := nextLevelXP - currentLevelXP

	return float64(progressXP) / float64(requiredXP), nil
}
