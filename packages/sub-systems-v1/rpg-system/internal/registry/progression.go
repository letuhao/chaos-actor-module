package registry

import (
	"math"
	"rpg-system/internal/model"
)

// ProgressionManager handles level progression and XP calculations
type ProgressionManager struct {
	registry *StatRegistry
}

// NewProgressionManager creates a new progression manager
func NewProgressionManager(registry *StatRegistry) *ProgressionManager {
	return &ProgressionManager{
		registry: registry,
	}
}

// XPTableEntry represents an XP requirement for a level
type XPTableEntry struct {
	Level         int64
	XPRequired    int64
	PointsGranted int64
}

// GetXPRequired calculates XP required for a given level
func (pm *ProgressionManager) GetXPRequired(level int64) int64 {
	if level <= 1 {
		return 0
	}

	// Exponential growth formula: XP = base * (level^exponent)
	// This creates a curve where higher levels require significantly more XP
	base := 1000.0
	exponent := 1.5

	xp := base * math.Pow(float64(level-1), exponent)
	return int64(xp)
}

// GetPointsGranted calculates stat points granted at a given level
func (pm *ProgressionManager) GetPointsGranted(level int64) int64 {
	if level <= 1 {
		return 0
	}

	// Grant 2 points per level, with bonus points at certain milestones
	points := int64(2)

	// Bonus points at level milestones
	if level%10 == 0 {
		points += 2 // +2 bonus every 10 levels
	}
	if level%25 == 0 {
		points += 3 // +3 bonus every 25 levels
	}
	if level%50 == 0 {
		points += 5 // +5 bonus every 50 levels
	}

	return points
}

// GetLevelFromXP calculates the level for a given XP amount
func (pm *ProgressionManager) GetLevelFromXP(xp int64) int64 {
	if xp <= 0 {
		return 1
	}

	// Reverse the XP formula to find level
	base := 1000.0
	exponent := 1.5

	level := math.Pow(float64(xp)/base, 1.0/exponent) + 1
	return int64(math.Floor(level))
}

// GetXPToNextLevel calculates XP needed to reach the next level
func (pm *ProgressionManager) GetXPToNextLevel(currentLevel int64, currentXP int64) int64 {
	nextLevel := currentLevel + 1
	nextLevelXP := pm.GetXPRequired(nextLevel)
	return nextLevelXP - currentXP
}

// GetXPProgress calculates progress towards next level (0.0 to 1.0)
func (pm *ProgressionManager) GetXPProgress(currentLevel int64, currentXP int64) float64 {
	if currentLevel <= 1 {
		return 0.0
	}

	currentLevelXP := pm.GetXPRequired(currentLevel)
	nextLevelXP := pm.GetXPRequired(currentLevel + 1)

	if nextLevelXP <= currentLevelXP {
		return 1.0
	}

	progress := float64(currentXP-currentLevelXP) / float64(nextLevelXP-currentLevelXP)
	if progress < 0 {
		return 0.0
	}
	if progress > 1 {
		return 1.0
	}

	return progress
}

// GenerateXPTable generates an XP table for levels 1 to maxLevel
func (pm *ProgressionManager) GenerateXPTable(maxLevel int64) []XPTableEntry {
	var table []XPTableEntry

	for level := int64(1); level <= maxLevel; level++ {
		entry := XPTableEntry{
			Level:         level,
			XPRequired:    pm.GetXPRequired(level),
			PointsGranted: pm.GetPointsGranted(level),
		}
		table = append(table, entry)
	}

	return table
}

// CalculateLevelUpResult calculates what happens when a player levels up
func (pm *ProgressionManager) CalculateLevelUpResult(currentLevel int64, currentXP int64, xpGained int64) *model.ProgressionResult {
	newXP := currentXP + xpGained
	newLevel := pm.GetLevelFromXP(newXP)

	if newLevel <= currentLevel {
		return &model.ProgressionResult{
			NewLevel:      nil,
			PointsGranted: nil,
		}
	}

	// Calculate total points granted for all levels gained
	totalPoints := int64(0)
	for level := currentLevel + 1; level <= newLevel; level++ {
		totalPoints += pm.GetPointsGranted(level)
	}

	return &model.ProgressionResult{
		NewLevel:      &newLevel,
		PointsGranted: &totalPoints,
	}
}

// GetRecommendedAllocation suggests stat allocation based on level and class
func (pm *ProgressionManager) GetRecommendedAllocation(level int64, class string) map[model.StatKey]int64 {
	// Get base points for this level
	points := pm.GetPointsGranted(level)

	// Class-based recommendations
	recommendations := make(map[model.StatKey]int64)

	switch class {
	case "warrior":
		// Focus on STR, END, some AGI
		recommendations[model.STR] = int64(float64(points) * 0.5)
		recommendations[model.END] = int64(float64(points) * 0.3)
		recommendations[model.AGI] = int64(float64(points) * 0.2)

	case "mage":
		// Focus on INT, WIL, some END
		recommendations[model.INT] = int64(float64(points) * 0.5)
		recommendations[model.WIL] = int64(float64(points) * 0.3)
		recommendations[model.END] = int64(float64(points) * 0.2)

	case "rogue":
		// Focus on AGI, SPD, some STR
		recommendations[model.AGI] = int64(float64(points) * 0.4)
		recommendations[model.SPD] = int64(float64(points) * 0.4)
		recommendations[model.STR] = int64(float64(points) * 0.2)

	case "paladin":
		// Balanced STR, END, WIL
		recommendations[model.STR] = int64(float64(points) * 0.4)
		recommendations[model.END] = int64(float64(points) * 0.4)
		recommendations[model.WIL] = int64(float64(points) * 0.2)

	default:
		// Balanced distribution
		pointsPerStat := points / 8
		remaining := points % 8

		for _, stat := range model.PrimaryStats() {
			recommendations[stat] = pointsPerStat
		}

		// Distribute remaining points
		for i, stat := range model.PrimaryStats() {
			if int64(i) < remaining {
				recommendations[stat]++
			}
		}
	}

	return recommendations
}

// ValidateAllocation checks if a stat allocation is valid
func (pm *ProgressionManager) ValidateAllocation(allocations map[model.StatKey]int64, availablePoints int64) bool {
	totalUsed := int64(0)
	for _, points := range allocations {
		if points < 0 {
			return false
		}
		totalUsed += points
	}

	return totalUsed <= availablePoints
}

// GetStatCap returns the maximum value a stat can reach at a given level
func (pm *ProgressionManager) GetStatCap(statKey model.StatKey, level int64) float64 {
	curve, exists := pm.registry.GetLevelCurve(statKey)
	if !exists {
		return 100.0 // Default cap
	}

	// Calculate max value based on level curve
	maxValue := curve.BaseValue + float64(level-1)*curve.PerLevel

	// Apply soft cap if applicable
	if curve.SoftCapLevel > 0 && level > curve.SoftCapLevel {
		softCapValue := curve.BaseValue + float64(curve.SoftCapLevel-1)*curve.PerLevel
		excessLevels := level - curve.SoftCapLevel
		maxValue = softCapValue + float64(excessLevels)*curve.SoftCapValue
	}

	// Apply hard cap
	if curve.MaxLevel > 0 && level > curve.MaxLevel {
		maxValue = curve.BaseValue + float64(curve.MaxLevel-1)*curve.PerLevel
	}

	return maxValue
}
