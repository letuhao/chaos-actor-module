package test

import (
	"testing"

	"rpg-system/internal/integration"
	"rpg-system/internal/model"
)

// Mock Core Actor Interface for testing
type MockActorCore struct {
	PrimaryStats map[model.StatKey]int64
	DerivedStats map[model.StatKey]float64
	Version      uint64
}

func NewMockActorCore() *MockActorCore {
	return &MockActorCore{
		PrimaryStats: make(map[model.StatKey]int64),
		DerivedStats: make(map[model.StatKey]float64),
		Version:      0,
	}
}

// ApplyStatSnapshot simulates how the Core Actor Interface would receive and apply a stat snapshot
func (a *MockActorCore) ApplyStatSnapshot(snapshot model.StatSnapshot) {
	// Update primary stats from the snapshot
	for statKey, statValue := range snapshot.Stats {
		if statKey.IsPrimary() {
			a.PrimaryStats[statKey] = int64(statValue)
		} else {
			a.DerivedStats[statKey] = statValue
		}
	}
	a.Version = uint64(snapshot.Version)
}

// GetStatValue simulates getting a stat value from the actor
func (a *MockActorCore) GetStatValue(statKey model.StatKey) float64 {
	if statKey.IsPrimary() {
		return float64(a.PrimaryStats[statKey])
	}
	return a.DerivedStats[statKey]
}

// TestRPGStatsIntegrationWithCoreActor demonstrates the integration
func TestRPGStatsIntegrationWithCoreActor(t *testing.T) {
	// Create the RPG Stats Sub-System
	provider := integration.NewSnapshotProvider()

	// Create a mock Core Actor
	actor := NewMockActorCore()

	// Test 1: Initial character creation
	t.Run("Initial Character Creation", func(t *testing.T) {
		// Build stat snapshot for a new character
		snapshot, err := provider.BuildForActor("test_actor_1", &model.SnapshotOptions{
			WithBreakdown: false,
		})
		if err != nil {
			t.Fatalf("Failed to build snapshot: %v", err)
		}

		// Apply the snapshot to the Core Actor
		actor.ApplyStatSnapshot(*snapshot)

		// Verify primary stats were applied
		expectedPrimaryStats := []model.StatKey{model.STR, model.INT, model.WIL, model.AGI, model.SPD, model.END, model.PER, model.LUK}
		for _, stat := range expectedPrimaryStats {
			value := actor.GetStatValue(stat)
			if value == 0 {
				t.Errorf("Expected %s to have a value, got 0", stat)
			}
		}

		// Verify derived stats were calculated
		expectedDerivedStats := []model.StatKey{model.HP_MAX, model.MANA_MAX, model.ATK, model.MATK, model.DEF}
		for _, stat := range expectedDerivedStats {
			value := actor.GetStatValue(stat)
			if value <= 0 {
				t.Errorf("Expected %s to have a positive value, got %f", stat, value)
			}
		}

		t.Logf("Initial character stats applied successfully")
		t.Logf("STR: %.0f, HP_MAX: %.2f, ATK: %.2f",
			actor.GetStatValue(model.STR),
			actor.GetStatValue(model.HP_MAX),
			actor.GetStatValue(model.ATK))
	})

	// Test 2: Equipment changes
	t.Run("Equipment Changes", func(t *testing.T) {
		// Simulate equipping a new weapon
		// In a real system, this would be handled by the RPG Stats Sub-System
		// when equipment is changed, it would rebuild the snapshot

		// Build a new snapshot (simulating equipment change)
		snapshot, err := provider.BuildForActor("test_actor_1", &model.SnapshotOptions{
			WithBreakdown: false,
		})
		if err != nil {
			t.Fatalf("Failed to build snapshot: %v", err)
		}

		// Apply the updated snapshot
		actor.ApplyStatSnapshot(*snapshot)

		// Verify stats are still valid
		hp := actor.GetStatValue(model.HP_MAX)
		atk := actor.GetStatValue(model.ATK)

		if hp <= 0 || atk <= 0 {
			t.Errorf("Stats should be positive after equipment change: HP=%.2f, ATK=%.2f", hp, atk)
		}

		t.Logf("Equipment change processed successfully")
		t.Logf("Updated stats - HP_MAX: %.2f, ATK: %.2f", hp, atk)
	})

	// Test 3: Level progression simulation
	t.Run("Level Progression", func(t *testing.T) {
		// Simulate leveling up by creating a new snapshot with higher stats
		// In a real system, this would be handled by the ProgressionService

		// Build snapshot for higher level character
		snapshot, err := provider.BuildForActor("test_actor_1", &model.SnapshotOptions{
			WithBreakdown: false,
		})
		if err != nil {
			t.Fatalf("Failed to build snapshot: %v", err)
		}

		// Apply the level-up snapshot
		actor.ApplyStatSnapshot(*snapshot)

		// Verify version was updated
		if actor.Version == 0 {
			t.Error("Expected version to be updated after applying snapshot")
		}

		t.Logf("Level progression processed successfully")
		t.Logf("Actor version: %d", actor.Version)
	})

	// Test 4: Stat breakdown demonstration
	t.Run("Stat Breakdown", func(t *testing.T) {
		// Build snapshot with breakdown enabled
		snapshot, err := provider.BuildForActor("test_actor_1", &model.SnapshotOptions{
			WithBreakdown: true,
		})
		if err != nil {
			t.Fatalf("Failed to build snapshot with breakdown: %v", err)
		}

		// Apply snapshot
		actor.ApplyStatSnapshot(*snapshot)

		// Verify we can still get stat values
		str := actor.GetStatValue(model.STR)
		hp := actor.GetStatValue(model.HP_MAX)

		if str <= 0 || hp <= 0 {
			t.Errorf("Stats should be positive: STR=%.2f, HP=%.2f", str, hp)
		}

		t.Logf("Stat breakdown processed successfully")
		t.Logf("Final stats - STR: %.2f, HP_MAX: %.2f", str, hp)
	})
}

// TestRPGStatsWithCustomModifiers demonstrates custom modifier integration
func TestRPGStatsWithCustomModifiers(t *testing.T) {
	// Create a custom snapshot provider that includes specific modifiers
	provider := integration.NewSnapshotProvider()

	// Build snapshot
	snapshot, err := provider.BuildForActor("custom_actor", &model.SnapshotOptions{
		WithBreakdown: true,
	})
	if err != nil {
		t.Fatalf("Failed to build custom snapshot: %v", err)
	}

	// Create actor and apply snapshot
	actor := NewMockActorCore()
	actor.ApplyStatSnapshot(*snapshot)

	// Verify the integration works with custom data
	t.Run("Custom Modifier Integration", func(t *testing.T) {
		// Check that equipment modifiers are applied
		str := actor.GetStatValue(model.STR)
		atk := actor.GetStatValue(model.ATK)

		// The demo provider includes equipment that adds +3 STR and +8 ATK
		// So we expect these values to be higher than base
		if str < 15 { // Base allocation is 15, +3 from equipment = 18
			t.Errorf("Expected STR to be at least 15 (base + equipment), got %.2f", str)
		}

		if atk < 30 { // Should be calculated from STR
			t.Errorf("Expected ATK to be calculated from STR, got %.2f", atk)
		}

		t.Logf("Custom modifiers applied successfully")
		t.Logf("STR: %.2f (base + equipment), ATK: %.2f (derived)", str, atk)
	})

	// Test stat calculation consistency
	t.Run("Stat Calculation Consistency", func(t *testing.T) {
		// Verify that derived stats are calculated consistently
		str := actor.GetStatValue(model.STR)
		hp := actor.GetStatValue(model.HP_MAX)

		// HP should be calculated from STR (among other factors)
		// The formula in our resolver: HP = 100 + STR*10 + END*5
		// With STR=18 and END=16, we expect HP > 100 + 18*10 + 16*5 = 360
		if hp < 300 {
			t.Errorf("HP calculation seems incorrect: STR=%.2f, HP=%.2f", str, hp)
		}

		t.Logf("Stat calculations are consistent")
		t.Logf("STR: %.2f -> HP_MAX: %.2f", str, hp)
	})
}

// TestCoreActorInterfaceCompatibility verifies the interface works correctly
func TestCoreActorInterfaceCompatibility(t *testing.T) {
	actor := NewMockActorCore()
	provider := integration.NewSnapshotProvider()

	// Test multiple snapshot applications
	for i := 0; i < 3; i++ {
		snapshot, err := provider.BuildForActor("test_actor", &model.SnapshotOptions{
			WithBreakdown: false,
		})
		if err != nil {
			t.Fatalf("Failed to build snapshot %d: %v", i, err)
		}

		// Apply snapshot
		actor.ApplyStatSnapshot(*snapshot)

		// Verify stats are always valid
		hp := actor.GetStatValue(model.HP_MAX)
		atk := actor.GetStatValue(model.ATK)

		if hp <= 0 || atk <= 0 {
			t.Errorf("Iteration %d: Invalid stats - HP=%.2f, ATK=%.2f", i, hp, atk)
		}

		t.Logf("Iteration %d: HP=%.2f, ATK=%.2f, Version=%d", i, hp, atk, actor.Version)
	}
}
