package test

import (
	"testing"

	"rpg-system/internal/integration"
	"rpg-system/internal/model"
)

func TestProgressionService_GrantXP(t *testing.T) {
	service := integration.NewProgressionService()

	tests := []struct {
		name           string
		actorID        string
		xpAmount       int64
		expectedLevel  int64
		expectedPoints int64
		expectError    bool
	}{
		{
			name:           "New player with 1000 XP",
			actorID:        "player1",
			xpAmount:       1000,
			expectedLevel:  4, // Level 4 requires 900 XP
			expectedPoints: 6, // 2 points per level for levels 2,3,4
			expectError:    false,
		},
		{
			name:           "Existing player with 500 XP",
			actorID:        "player1",
			xpAmount:       500,
			expectedLevel:  0, // No level up expected
			expectedPoints: 0, // No new level up
			expectError:    false,
		},
		{
			name:           "Level up to 5",
			actorID:        "player1",
			xpAmount:       600, // Total 2100 XP, level 5 requires 1600 XP
			expectedLevel:  5,
			expectedPoints: 3, // 2 points for level 5 + 1 bonus
			expectError:    false,
		},
		{
			name:        "Negative XP",
			actorID:     "player2",
			xpAmount:    -100,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GrantXP(tt.actorID, tt.xpAmount)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Check if level up occurred
			if tt.expectedLevel > 0 {
				if result.NewLevel == nil {
					t.Errorf("Expected level up to level %d, but no level up occurred", tt.expectedLevel)
				} else if *result.NewLevel != tt.expectedLevel {
					t.Errorf("Expected level %d, got %d", tt.expectedLevel, *result.NewLevel)
				}
			}

			// Check points granted
			if tt.expectedPoints > 0 {
				if result.PointsGranted == nil {
					t.Errorf("Expected %d points granted, but no points were granted", tt.expectedPoints)
				} else if *result.PointsGranted != tt.expectedPoints {
					t.Errorf("Expected %d points granted, got %d", tt.expectedPoints, *result.PointsGranted)
				}
			}
		})
	}
}

func TestProgressionService_AllocatePoints(t *testing.T) {
	service := integration.NewProgressionService()

	// Set up a player with some points
	_, err := service.GrantXP("player1", 1000) // Should give level 4 with 6 points
	if err != nil {
		t.Fatalf("Failed to grant XP: %v", err)
	}

	tests := []struct {
		name        string
		actorID     string
		statKey     model.StatKey
		points      int64
		expectError bool
	}{
		{
			name:        "Allocate 3 points to STR",
			actorID:     "player1",
			statKey:     model.STR,
			points:      3,
			expectError: false,
		},
		{
			name:        "Allocate 2 more points to STR",
			actorID:     "player1",
			statKey:     model.STR,
			points:      2,
			expectError: false,
		},
		{
			name:        "Try to allocate too many points",
			actorID:     "player1",
			statKey:     model.STR,
			points:      10, // Only 1 point left
			expectError: true,
		},
		{
			name:        "Allocate to derived stat (should fail)",
			actorID:     "player1",
			statKey:     model.HP_MAX,
			points:      1,
			expectError: true,
		},
		{
			name:        "Negative points",
			actorID:     "player1",
			statKey:     model.STR,
			points:      -1,
			expectError: true,
		},
		{
			name:        "Non-existent player",
			actorID:     "nonexistent",
			statKey:     model.STR,
			points:      1,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.AllocatePoints(tt.actorID, tt.statKey, tt.points)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestProgressionService_Respec(t *testing.T) {
	service := integration.NewProgressionService()

	// Set up a player with allocated points
	_, err := service.GrantXP("player1", 1000) // Level 4 with 6 points
	if err != nil {
		t.Fatalf("Failed to grant XP: %v", err)
	}

	err = service.AllocatePoints("player1", model.STR, 3)
	if err != nil {
		t.Fatalf("Failed to allocate points: %v", err)
	}

	err = service.AllocatePoints("player1", model.INT, 2)
	if err != nil {
		t.Fatalf("Failed to allocate points: %v", err)
	}

	// Test respec
	err = service.Respec("player1")
	if err != nil {
		t.Errorf("Unexpected error during respec: %v", err)
	}

	// Verify points are available again
	available, err := service.GetAvailablePoints("player1")
	if err != nil {
		t.Errorf("Failed to get available points: %v", err)
	}

	if available != 6 {
		t.Errorf("Expected 6 available points after respec, got %d", available)
	}
}

func TestProgressionService_GetPlayerProgress(t *testing.T) {
	service := integration.NewProgressionService()

	// Test non-existent player
	_, err := service.GetPlayerProgress("nonexistent")
	if err == nil {
		t.Errorf("Expected error for non-existent player")
	}

	// Create a player and test
	_, err = service.GrantXP("player1", 1000)
	if err != nil {
		t.Fatalf("Failed to grant XP: %v", err)
	}

	err = service.AllocatePoints("player1", model.STR, 3)
	if err != nil {
		t.Fatalf("Failed to allocate points: %v", err)
	}

	progress, err := service.GetPlayerProgress("player1")
	if err != nil {
		t.Errorf("Failed to get player progress: %v", err)
	}

	if progress.ActorID != "player1" {
		t.Errorf("Expected actor ID 'player1', got '%s'", progress.ActorID)
	}

	if progress.Level != 4 {
		t.Errorf("Expected level 4, got %d", progress.Level)
	}

	if progress.Allocations[model.STR] != 3 {
		t.Errorf("Expected 3 STR allocations, got %d", progress.Allocations[model.STR])
	}
}

func TestProgressionService_GetLevelProgression(t *testing.T) {
	service := integration.NewProgressionService()

	tests := []struct {
		level          int64
		expectedXP     int64
		expectedPoints int64
	}{
		{1, 0, 0},
		{2, 100, 2},
		{3, 400, 2},
		{4, 900, 2},
		{5, 1600, 3},  // Bonus point every 5 levels
		{10, 8100, 3}, // Another bonus point
	}

	for _, tt := range tests {
		t.Run("Level "+string(rune(tt.level)), func(t *testing.T) {
			progression := service.GetLevelProgression(tt.level)

			if progression.Level != tt.level {
				t.Errorf("Expected level %d, got %d", tt.level, progression.Level)
			}

			if progression.XPRequired != tt.expectedXP {
				t.Errorf("Expected %d XP for level %d, got %d", tt.expectedXP, tt.level, progression.XPRequired)
			}

			if progression.PointsGranted != tt.expectedPoints {
				t.Errorf("Expected %d points for level %d, got %d", tt.expectedPoints, tt.level, progression.PointsGranted)
			}
		})
	}
}

func TestProgressionService_GetXPToNextLevel(t *testing.T) {
	service := integration.NewProgressionService()

	// Test non-existent player
	_, err := service.GetXPToNextLevel("nonexistent")
	if err == nil {
		t.Errorf("Expected error for non-existent player")
	}

	// Test level 1 player
	_, err = service.GrantXP("player1", 1) // Level 1 with 1 XP
	if err != nil {
		t.Fatalf("Failed to grant XP: %v", err)
	}

	xpToNext, err := service.GetXPToNextLevel("player1")
	if err != nil {
		t.Errorf("Failed to get XP to next level: %v", err)
	}

	if xpToNext != 99 { // Level 2 requires 100 XP, player has 1 XP
		t.Errorf("Expected 99 XP to next level, got %d", xpToNext)
	}
}

func TestProgressionService_GetLevelProgress(t *testing.T) {
	service := integration.NewProgressionService()

	// Test non-existent player
	_, err := service.GetLevelProgress("nonexistent")
	if err == nil {
		t.Errorf("Expected error for non-existent player")
	}

	// Test level 1 player with some XP
	_, err = service.GrantXP("player1", 50) // Halfway to level 2
	if err != nil {
		t.Fatalf("Failed to grant XP: %v", err)
	}

	progress, err := service.GetLevelProgress("player1")
	if err != nil {
		t.Errorf("Failed to get level progress: %v", err)
	}

	if progress < 0.49 || progress > 0.51 {
		t.Errorf("Expected progress around 0.5, got %f", progress)
	}
}
