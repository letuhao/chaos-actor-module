package unit

import (
	"testing"
	"time"

	"actor-core-v2/models/core"
)

func TestNewPrimaryCore(t *testing.T) {
	pc := core.NewPrimaryCore()

	// Test default values
	if pc.Vitality != 10 {
		t.Errorf("Expected Vitality to be 10, got %d", pc.Vitality)
	}

	if pc.Endurance != 10 {
		t.Errorf("Expected Endurance to be 10, got %d", pc.Endurance)
	}

	if pc.Constitution != 10 {
		t.Errorf("Expected Constitution to be 10, got %d", pc.Constitution)
	}

	if pc.Intelligence != 10 {
		t.Errorf("Expected Intelligence to be 10, got %d", pc.Intelligence)
	}

	if pc.Wisdom != 10 {
		t.Errorf("Expected Wisdom to be 10, got %d", pc.Wisdom)
	}

	if pc.Charisma != 10 {
		t.Errorf("Expected Charisma to be 10, got %d", pc.Charisma)
	}

	if pc.Willpower != 10 {
		t.Errorf("Expected Willpower to be 10, got %d", pc.Willpower)
	}

	if pc.Luck != 10 {
		t.Errorf("Expected Luck to be 10, got %d", pc.Luck)
	}

	if pc.Fate != 10 {
		t.Errorf("Expected Fate to be 10, got %d", pc.Fate)
	}

	if pc.Karma != 0 {
		t.Errorf("Expected Karma to be 0, got %d", pc.Karma)
	}

	if pc.Strength != 10 {
		t.Errorf("Expected Strength to be 10, got %d", pc.Strength)
	}

	if pc.Agility != 10 {
		t.Errorf("Expected Agility to be 10, got %d", pc.Agility)
	}

	if pc.Personality != 10 {
		t.Errorf("Expected Personality to be 10, got %d", pc.Personality)
	}

	if pc.SpiritualEnergy != 0 {
		t.Errorf("Expected SpiritualEnergy to be 0, got %d", pc.SpiritualEnergy)
	}

	if pc.PhysicalEnergy != 0 {
		t.Errorf("Expected PhysicalEnergy to be 0, got %d", pc.PhysicalEnergy)
	}

	if pc.MentalEnergy != 0 {
		t.Errorf("Expected MentalEnergy to be 0, got %d", pc.MentalEnergy)
	}

	if pc.CultivationLevel != 0 {
		t.Errorf("Expected CultivationLevel to be 0, got %d", pc.CultivationLevel)
	}

	if pc.BreakthroughPoints != 0 {
		t.Errorf("Expected BreakthroughPoints to be 0, got %d", pc.BreakthroughPoints)
	}

	if pc.LifeSpan != 100 {
		t.Errorf("Expected LifeSpan to be 100, got %d", pc.LifeSpan)
	}

	if pc.Age != 0 {
		t.Errorf("Expected Age to be 0, got %d", pc.Age)
	}

	if pc.Version != 1 {
		t.Errorf("Expected Version to be 1, got %d", pc.Version)
	}

	if pc.CreatedAt == 0 {
		t.Error("Expected CreatedAt to be set")
	}

	if pc.UpdatedAt == 0 {
		t.Error("Expected UpdatedAt to be set")
	}
}

func TestNewPrimaryCoreWithValues(t *testing.T) {
	values := map[string]int64{
		"vitality":            15,
		"endurance":           20,
		"constitution":        25,
		"intelligence":        30,
		"wisdom":              35,
		"charisma":            40,
		"willpower":           45,
		"luck":                50,
		"fate":                55,
		"karma":               60,
		"strength":            65,
		"agility":             70,
		"personality":         75,
		"spiritual_energy":    80,
		"physical_energy":     85,
		"mental_energy":       90,
		"cultivation_level":   95,
		"breakthrough_points": 100,
		"life_span":           200,
		"age":                 25,
	}

	pc := core.NewPrimaryCoreWithValues(values)

	// Test that values were set correctly
	if pc.Vitality != 15 {
		t.Errorf("Expected Vitality to be 15, got %d", pc.Vitality)
	}

	if pc.Endurance != 20 {
		t.Errorf("Expected Endurance to be 20, got %d", pc.Endurance)
	}

	if pc.Constitution != 25 {
		t.Errorf("Expected Constitution to be 25, got %d", pc.Constitution)
	}

	if pc.Intelligence != 30 {
		t.Errorf("Expected Intelligence to be 30, got %d", pc.Intelligence)
	}

	if pc.Wisdom != 35 {
		t.Errorf("Expected Wisdom to be 35, got %d", pc.Wisdom)
	}

	if pc.Charisma != 40 {
		t.Errorf("Expected Charisma to be 40, got %d", pc.Charisma)
	}

	if pc.Willpower != 45 {
		t.Errorf("Expected Willpower to be 45, got %d", pc.Willpower)
	}

	if pc.Luck != 50 {
		t.Errorf("Expected Luck to be 50, got %d", pc.Luck)
	}

	if pc.Fate != 55 {
		t.Errorf("Expected Fate to be 55, got %d", pc.Fate)
	}

	if pc.Karma != 60 {
		t.Errorf("Expected Karma to be 60, got %d", pc.Karma)
	}

	if pc.Strength != 65 {
		t.Errorf("Expected Strength to be 65, got %d", pc.Strength)
	}

	if pc.Agility != 70 {
		t.Errorf("Expected Agility to be 70, got %d", pc.Agility)
	}

	if pc.Personality != 75 {
		t.Errorf("Expected Personality to be 75, got %d", pc.Personality)
	}

	if pc.SpiritualEnergy != 80 {
		t.Errorf("Expected SpiritualEnergy to be 80, got %d", pc.SpiritualEnergy)
	}

	if pc.PhysicalEnergy != 85 {
		t.Errorf("Expected PhysicalEnergy to be 85, got %d", pc.PhysicalEnergy)
	}

	if pc.MentalEnergy != 90 {
		t.Errorf("Expected MentalEnergy to be 90, got %d", pc.MentalEnergy)
	}

	if pc.CultivationLevel != 95 {
		t.Errorf("Expected CultivationLevel to be 95, got %d", pc.CultivationLevel)
	}

	if pc.BreakthroughPoints != 100 {
		t.Errorf("Expected BreakthroughPoints to be 100, got %d", pc.BreakthroughPoints)
	}

	if pc.LifeSpan != 200 {
		t.Errorf("Expected LifeSpan to be 200, got %d", pc.LifeSpan)
	}

	if pc.Age != 25 {
		t.Errorf("Expected Age to be 25, got %d", pc.Age)
	}
}

func TestGetStat(t *testing.T) {
	pc := core.NewPrimaryCore()

	// Test valid stat names
	stat, err := pc.GetStat("vitality")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if stat != 10 {
		t.Errorf("Expected vitality to be 10, got %d", stat)
	}

	stat, err = pc.GetStat("karma")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if stat != 0 {
		t.Errorf("Expected karma to be 0, got %d", stat)
	}

	// Test invalid stat name
	_, err = pc.GetStat("invalid_stat")
	if err == nil {
		t.Error("Expected error for invalid stat name")
	}
}

func TestSetStat(t *testing.T) {
	pc := core.NewPrimaryCore()

	// Test setting a valid stat
	err := pc.SetStat("vitality", 20)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if pc.Vitality != 20 {
		t.Errorf("Expected vitality to be 20, got %d", pc.Vitality)
	}

	// Test setting karma
	err = pc.SetStat("karma", 50)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if pc.Karma != 50 {
		t.Errorf("Expected karma to be 50, got %d", pc.Karma)
	}

	// Test setting invalid stat
	err = pc.SetStat("invalid_stat", 10)
	if err == nil {
		t.Error("Expected error for invalid stat name")
	}

	// Test that version was incremented (2 calls to SetStat = 2 increments)
	if pc.Version != 3 {
		t.Errorf("Expected version to be 3, got %d", pc.Version)
	}
}

func TestGetAllStats(t *testing.T) {
	pc := core.NewPrimaryCore()
	stats := pc.GetAllStats()

	// Test that all stats are present
	expectedStats := []string{
		"vitality", "endurance", "constitution", "intelligence", "wisdom",
		"charisma", "willpower", "luck", "fate", "karma",
		"strength", "agility", "personality",
		"spiritual_energy", "physical_energy", "mental_energy",
		"cultivation_level", "breakthrough_points",
		"life_span", "age",
	}

	for _, statName := range expectedStats {
		if _, exists := stats[statName]; !exists {
			t.Errorf("Expected stat %s to be present", statName)
		}
	}

	// Test that values are correct
	if stats["vitality"] != 10 {
		t.Errorf("Expected vitality to be 10, got %d", stats["vitality"])
	}

	if stats["karma"] != 0 {
		t.Errorf("Expected karma to be 0, got %d", stats["karma"])
	}
}

func TestUpdateStats(t *testing.T) {
	pc := core.NewPrimaryCore()

	updates := map[string]int64{
		"vitality":     20,
		"endurance":    25,
		"constitution": 30,
		"karma":        50,
	}

	err := pc.UpdateStats(updates)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test that values were updated
	if pc.Vitality != 20 {
		t.Errorf("Expected vitality to be 20, got %d", pc.Vitality)
	}

	if pc.Endurance != 25 {
		t.Errorf("Expected endurance to be 25, got %d", pc.Endurance)
	}

	if pc.Constitution != 30 {
		t.Errorf("Expected constitution to be 30, got %d", pc.Constitution)
	}

	if pc.Karma != 50 {
		t.Errorf("Expected karma to be 50, got %d", pc.Karma)
	}

	// Test that version was incremented (4 calls to SetStat = 4 increments)
	if pc.Version != 5 {
		t.Errorf("Expected version to be 5, got %d", pc.Version)
	}
}

func TestClone(t *testing.T) {
	pc := core.NewPrimaryCore()
	pc.Vitality = 20
	pc.Karma = 50
	pc.Version = 5

	cloned := pc.Clone()

	// Test that values are copied
	if cloned.Vitality != 20 {
		t.Errorf("Expected cloned vitality to be 20, got %d", cloned.Vitality)
	}

	if cloned.Karma != 50 {
		t.Errorf("Expected cloned karma to be 50, got %d", cloned.Karma)
	}

	if cloned.Version != 5 {
		t.Errorf("Expected cloned version to be 5, got %d", cloned.Version)
	}

	// Test that modifying clone doesn't affect original
	cloned.Vitality = 30
	if pc.Vitality != 20 {
		t.Error("Modifying clone should not affect original")
	}
}

func TestReset(t *testing.T) {
	pc := core.NewPrimaryCore()
	pc.Vitality = 20
	pc.Karma = 50
	pc.Version = 5

	pc.Reset()

	// Test that values are reset to defaults
	if pc.Vitality != 10 {
		t.Errorf("Expected vitality to be reset to 10, got %d", pc.Vitality)
	}

	if pc.Karma != 0 {
		t.Errorf("Expected karma to be reset to 0, got %d", pc.Karma)
	}

	if pc.Version != 6 {
		t.Errorf("Expected version to be 6, got %d", pc.Version)
	}
}

func TestValidate(t *testing.T) {
	pc := core.NewPrimaryCore()

	// Test valid stats
	errors := pc.Validate()
	if len(errors) > 0 {
		t.Errorf("Expected no validation errors, got %d", len(errors))
	}

	// Test invalid stats
	pc.Vitality = -1
	pc.Age = 150 // exceeds life span
	pc.LifeSpan = 100

	errors = pc.Validate()
	if len(errors) == 0 {
		t.Error("Expected validation errors for invalid stats")
	}

	// Check specific errors
	hasVitalityError := false
	hasAgeError := false

	for _, err := range errors {
		if err.Field == "vitality" && err.Message == "Value cannot be negative" {
			hasVitalityError = true
		}
		if err.Field == "age" && err.Message == "Age cannot exceed life span" {
			hasAgeError = true
		}
	}

	if !hasVitalityError {
		t.Error("Expected vitality validation error")
	}

	if !hasAgeError {
		t.Error("Expected age validation error")
	}
}

func TestGetBasicStats(t *testing.T) {
	pc := core.NewPrimaryCore()
	basicStats := pc.GetBasicStats()

	expectedBasicStats := []string{
		"vitality", "endurance", "constitution", "intelligence", "wisdom",
		"charisma", "willpower", "luck", "fate", "karma",
	}

	for _, statName := range expectedBasicStats {
		if _, exists := basicStats[statName]; !exists {
			t.Errorf("Expected basic stat %s to be present", statName)
		}
	}

	// Test that physical stats are not included
	if _, exists := basicStats["strength"]; exists {
		t.Error("Physical stats should not be in basic stats")
	}
}

func TestGetPhysicalStats(t *testing.T) {
	pc := core.NewPrimaryCore()
	physicalStats := pc.GetPhysicalStats()

	expectedPhysicalStats := []string{
		"strength", "agility", "personality",
	}

	for _, statName := range expectedPhysicalStats {
		if _, exists := physicalStats[statName]; !exists {
			t.Errorf("Expected physical stat %s to be present", statName)
		}
	}

	// Test that basic stats are not included
	if _, exists := physicalStats["vitality"]; exists {
		t.Error("Basic stats should not be in physical stats")
	}
}

func TestGetCultivationStats(t *testing.T) {
	pc := core.NewPrimaryCore()
	cultivationStats := pc.GetCultivationStats()

	expectedCultivationStats := []string{
		"spiritual_energy", "physical_energy", "mental_energy",
		"cultivation_level", "breakthrough_points",
	}

	for _, statName := range expectedCultivationStats {
		if _, exists := cultivationStats[statName]; !exists {
			t.Errorf("Expected cultivation stat %s to be present", statName)
		}
	}

	// Test that basic stats are not included
	if _, exists := cultivationStats["vitality"]; exists {
		t.Error("Basic stats should not be in cultivation stats")
	}
}

func TestGetLifeStats(t *testing.T) {
	pc := core.NewPrimaryCore()
	lifeStats := pc.GetLifeStats()

	expectedLifeStats := []string{
		"life_span", "age",
	}

	for _, statName := range expectedLifeStats {
		if _, exists := lifeStats[statName]; !exists {
			t.Errorf("Expected life stat %s to be present", statName)
		}
	}

	// Test that basic stats are not included
	if _, exists := lifeStats["vitality"]; exists {
		t.Error("Basic stats should not be in life stats")
	}
}

func TestIsAlive(t *testing.T) {
	pc := core.NewPrimaryCore()

	// Test alive actor
	if !pc.IsAlive() {
		t.Error("Expected actor to be alive")
	}

	// Test dead actor
	pc.Age = 100
	if pc.IsAlive() {
		t.Error("Expected actor to be dead")
	}

	// Test actor at life span
	pc.Age = 99
	if !pc.IsAlive() {
		t.Error("Expected actor to be alive at age 99")
	}
}

func TestGetRemainingLife(t *testing.T) {
	pc := core.NewPrimaryCore()

	// Test normal case
	remaining := pc.GetRemainingLife()
	if remaining != 100 {
		t.Errorf("Expected remaining life to be 100, got %d", remaining)
	}

	// Test actor at half life
	pc.Age = 50
	remaining = pc.GetRemainingLife()
	if remaining != 50 {
		t.Errorf("Expected remaining life to be 50, got %d", remaining)
	}

	// Test dead actor
	pc.Age = 100
	remaining = pc.GetRemainingLife()
	if remaining != 0 {
		t.Errorf("Expected remaining life to be 0, got %d", remaining)
	}
}

func TestAgeUp(t *testing.T) {
	pc := core.NewPrimaryCore()
	originalAge := pc.Age
	originalVersion := pc.Version

	pc.AgeUp()

	if pc.Age != originalAge+1 {
		t.Errorf("Expected age to be %d, got %d", originalAge+1, pc.Age)
	}

	if pc.Version != originalVersion+1 {
		t.Errorf("Expected version to be %d, got %d", originalVersion+1, pc.Version)
	}
}

func TestAgeUpBy(t *testing.T) {
	pc := core.NewPrimaryCore()
	originalAge := pc.Age
	originalVersion := pc.Version

	pc.AgeUpBy(5)

	if pc.Age != originalAge+5 {
		t.Errorf("Expected age to be %d, got %d", originalAge+5, pc.Age)
	}

	if pc.Version != originalVersion+1 {
		t.Errorf("Expected version to be %d, got %d", originalVersion+1, pc.Version)
	}
}

func TestGetVersion(t *testing.T) {
	pc := core.NewPrimaryCore()

	if pc.GetVersion() != 1 {
		t.Errorf("Expected version to be 1, got %d", pc.GetVersion())
	}

	pc.SetStat("vitality", 20)

	if pc.GetVersion() != 2 {
		t.Errorf("Expected version to be 2, got %d", pc.GetVersion())
	}
}

func TestGetUpdatedAt(t *testing.T) {
	pc := core.NewPrimaryCore()
	originalUpdatedAt := pc.GetUpdatedAt()

	// Wait a bit to ensure timestamp changes (Unix returns seconds, so we need at least 1 second)
	time.Sleep(1 * time.Second)

	pc.SetStat("vitality", 20)

	if pc.GetUpdatedAt() <= originalUpdatedAt {
		t.Errorf("Expected UpdatedAt to be updated. Original: %d, New: %d", originalUpdatedAt, pc.GetUpdatedAt())
	}
}

func TestGetCreatedAt(t *testing.T) {
	pc := core.NewPrimaryCore()

	if pc.GetCreatedAt() == 0 {
		t.Error("Expected CreatedAt to be set")
	}
}
