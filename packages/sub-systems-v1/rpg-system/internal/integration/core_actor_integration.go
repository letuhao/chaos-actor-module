package integration

import (
	"context"
	"rpg-system/internal/model"
	"rpg-system/internal/resolver"
	"time"
)

// DatabaseAdapter interface for database operations
type DatabaseAdapter interface {
	GetPlayerStatsSummary(ctx context.Context, actorID string) (*PlayerStatsSummary, error)
	GetPlayerProgress(ctx context.Context, actorID string) (*model.PlayerProgress, error)
	SavePlayerProgress(ctx context.Context, progress *model.PlayerProgress) error
	AddEffect(ctx context.Context, actorID string, effect model.StatModifier, duration time.Duration) error
	RemoveEffect(ctx context.Context, actorID, effectID string) error
	EquipItem(ctx context.Context, actorID, slot string, item model.StatModifier) error
	UnequipItem(ctx context.Context, actorID, slot string) error
	GrantTitle(ctx context.Context, actorID, titleID string, title model.StatModifier) error
	Close() error
}

// CoreActorIntegration provides integration between RPG Stats and Core Actor
type CoreActorIntegration struct {
	resolver           *resolver.StatResolver
	adapter            DatabaseAdapter
	progressionService *ProgressionService
}

// NewCoreActorIntegration creates a new integration instance
func NewCoreActorIntegration(adapter DatabaseAdapter) *CoreActorIntegration {
	return &CoreActorIntegration{
		resolver:           resolver.NewStatResolver(),
		adapter:            adapter,
		progressionService: NewProgressionService(),
	}
}

// CoreActorContribution represents a contribution to the Core Actor system
type CoreActorContribution struct {
	Primary *PrimaryCoreStats
	Flat    map[string]float64
	Mult    map[string]float64
	Tags    []string
}

// PrimaryCoreStats maps RPG primary stats to Core Actor primary stats
type PrimaryCoreStats struct {
	HPMax    int64 // Mapped from END + STR
	LifeSpan int64 // Mapped from END
	Attack   int64 // Mapped from STR + AGI
	Defense  int64 // Mapped from END + STR
	Speed    int64 // Mapped from SPD + AGI
}

// BuildCoreActorContribution builds a Core Actor contribution from RPG Stats
func (cai *CoreActorIntegration) BuildCoreActorContribution(ctx context.Context, actorID string) (*CoreActorContribution, error) {
	// Get player data from database
	summary, err := cai.adapter.GetPlayerStatsSummary(ctx, actorID)
	if err != nil {
		return nil, err
	}

	// Get stat registry from the progression service
	primaryStats := cai.progressionService.registry.GetAllPrimaryStats()
	derivedStats := cai.progressionService.registry.GetAllDerivedStats()

	// Combine primary and derived stats
	registry := make([]model.StatDef, 0, len(primaryStats)+len(derivedStats))
	for _, stat := range primaryStats {
		registry = append(registry, *stat)
	}
	for _, stat := range derivedStats {
		registry = append(registry, *stat)
	}

	// Build compute input
	input := model.ComputeInput{
		ActorID:         actorID,
		Level:           summary.Progress.Level,
		BaseAllocations: summary.Progress.Allocations,
		Registry:        registry,
		Items:           summary.Equipment,
		Titles:          summary.Titles,
		Passives:        []model.StatModifier{}, // Could be loaded from database
		Buffs:           summary.Effects,
		Debuffs:         []model.StatModifier{}, // Could be separated from effects
		Auras:           []model.StatModifier{}, // Could be loaded from environment
		Environment:     []model.StatModifier{}, // Could be loaded from environment
		WithBreakdown:   false,
	}

	// Compute stat snapshot
	snapshot := cai.resolver.ComputeSnapshot(input)

	// Convert to Core Actor contribution
	contribution := cai.convertToCoreActorContribution(snapshot)

	return contribution, nil
}

// convertToCoreActorContribution converts RPG Stats snapshot to Core Actor contribution
func (cai *CoreActorIntegration) convertToCoreActorContribution(snapshot *model.StatSnapshot) *CoreActorContribution {
	// Map primary stats to Core Actor primary stats
	primary := &PrimaryCoreStats{
		HPMax:    int64(snapshot.Stats[model.END]*10 + snapshot.Stats[model.STR]*2),
		LifeSpan: int64(snapshot.Stats[model.END]),
		Attack:   int64(snapshot.Stats[model.STR]*2 + snapshot.Stats[model.AGI]),
		Defense:  int64(snapshot.Stats[model.END]*1.5 + snapshot.Stats[model.STR]*0.5),
		Speed:    int64(snapshot.Stats[model.SPD] + snapshot.Stats[model.AGI]*0.5),
	}

	// Map derived stats to Core Actor flat modifiers
	flat := make(map[string]float64)
	mult := make(map[string]float64)

	// Direct mappings
	flat["HPMax"] = snapshot.Stats[model.HP_MAX]
	flat["MPMax"] = snapshot.Stats[model.MANA_MAX]
	flat["ATK"] = snapshot.Stats[model.ATK]
	flat["MAG"] = snapshot.Stats[model.MATK]
	flat["DEF"] = snapshot.Stats[model.DEF]
	flat["RES"] = snapshot.Stats[model.DEF] * 0.8 // Magic resistance
	flat["Haste"] = 1.0 + (snapshot.Stats[model.SPD]-50)/100.0
	flat["CritChance"] = snapshot.Stats[model.CRIT_CHANCE]
	flat["CritMulti"] = snapshot.Stats[model.CRIT_DAMAGE]
	flat["MoveSpeed"] = snapshot.Stats[model.MOVE_SPEED]

	// Regeneration based on max stats
	flat["RegenHP"] = snapshot.Stats[model.HP_MAX] * 0.01
	flat["RegenMP"] = snapshot.Stats[model.MANA_MAX] * 0.01

	// Evasion as a resistance
	flat["resists.evasion"] = snapshot.Stats[model.EVASION] / 100.0

	// Personality affects social interactions (could be amplifiers)
	flat["amplifiers.social"] = 1.0 + snapshot.Stats[model.PER]/100.0

	// Luck affects critical chance and other random events
	mult["CritChance"] = 1.0 + snapshot.Stats[model.LUK]/1000.0

	// Tags based on stat levels
	tags := cai.generateTags(snapshot)

	return &CoreActorContribution{
		Primary: primary,
		Flat:    flat,
		Mult:    mult,
		Tags:    tags,
	}
}

// generateTags generates tags based on stat levels
func (cai *CoreActorIntegration) generateTags(snapshot *model.StatSnapshot) []string {
	var tags []string

	// Strength-based tags
	if snapshot.Stats[model.STR] > 20 {
		tags = append(tags, "strong")
	}
	if snapshot.Stats[model.STR] > 40 {
		tags = append(tags, "very_strong")
	}

	// Intelligence-based tags
	if snapshot.Stats[model.INT] > 20 {
		tags = append(tags, "intelligent")
	}
	if snapshot.Stats[model.INT] > 40 {
		tags = append(tags, "very_intelligent")
	}

	// Agility-based tags
	if snapshot.Stats[model.AGI] > 20 {
		tags = append(tags, "agile")
	}
	if snapshot.Stats[model.AGI] > 40 {
		tags = append(tags, "very_agile")
	}

	// Endurance-based tags
	if snapshot.Stats[model.END] > 20 {
		tags = append(tags, "tough")
	}
	if snapshot.Stats[model.END] > 40 {
		tags = append(tags, "very_tough")
	}

	// Speed-based tags
	if snapshot.Stats[model.SPD] > 20 {
		tags = append(tags, "fast")
	}
	if snapshot.Stats[model.SPD] > 40 {
		tags = append(tags, "very_fast")
	}

	// Personality-based tags
	if snapshot.Stats[model.PER] > 20 {
		tags = append(tags, "charismatic")
	}
	if snapshot.Stats[model.PER] > 40 {
		tags = append(tags, "very_charismatic")
	}

	// Luck-based tags
	if snapshot.Stats[model.LUK] > 20 {
		tags = append(tags, "lucky")
	}
	if snapshot.Stats[model.LUK] > 40 {
		tags = append(tags, "very_lucky")
	}

	// Willpower-based tags
	if snapshot.Stats[model.WIL] > 20 {
		tags = append(tags, "strong_willed")
	}
	if snapshot.Stats[model.WIL] > 40 {
		tags = append(tags, "very_strong_willed")
	}

	return tags
}

// UpdatePlayerStats updates player stats and returns the new Core Actor contribution
func (cai *CoreActorIntegration) UpdatePlayerStats(ctx context.Context, actorID string, updates *PlayerStatsUpdates) (*CoreActorContribution, error) {
	// Get current progress
	progress, err := cai.adapter.GetPlayerProgress(ctx, actorID)
	if err != nil {
		return nil, err
	}

	// Apply XP and level updates through ProgressionService
	if updates.XPGranted > 0 {
		// Set the current progress in the service
		cai.progressionService.playerProgress[actorID] = progress

		// Grant XP (this will handle level-ups properly)
		_, err = cai.progressionService.GrantXP(actorID, updates.XPGranted)
		if err != nil {
			return nil, err
		}

		// Get the updated progress from the progression service
		updatedProgress, err := cai.progressionService.GetPlayerProgress(actorID)
		if err != nil {
			return nil, err
		}

		// Update the progress with the new values
		progress.Level = updatedProgress.Level
		progress.XP = updatedProgress.XP
		progress.LastUpdated = time.Now().Unix()
	}

	// Apply stat allocations
	if updates.StatAllocations != nil {
		// Set the current progress in the service
		cai.progressionService.playerProgress[actorID] = progress

		for stat, points := range updates.StatAllocations {
			if points > 0 {
				err = cai.progressionService.AllocatePoints(actorID, stat, points)
				if err != nil {
					return nil, err
				}
			}
		}

		// Get the updated progress
		updatedProgress, err := cai.progressionService.GetPlayerProgress(actorID)
		if err != nil {
			return nil, err
		}

		// Update the progress with new allocations
		progress.Allocations = updatedProgress.Allocations
		progress.LastUpdated = time.Now().Unix()
	}

	// Save updated progress
	err = cai.adapter.SavePlayerProgress(ctx, progress)
	if err != nil {
		return nil, err
	}

	// Add/remove effects
	for _, effect := range updates.EffectsAdded {
		duration := 5 * time.Minute // Default duration
		if effect.Conditions != nil && effect.Conditions.DurationMs > 0 {
			duration = time.Duration(effect.Conditions.DurationMs) * time.Millisecond
		}
		err = cai.adapter.AddEffect(ctx, actorID, effect, duration)
		if err != nil {
			return nil, err
		}
	}

	for _, effectID := range updates.EffectsRemoved {
		err = cai.adapter.RemoveEffect(ctx, actorID, effectID)
		if err != nil {
			return nil, err
		}
	}

	// Equip/unequip items
	for slot, item := range updates.ItemsEquipped {
		err = cai.adapter.EquipItem(ctx, actorID, slot, item)
		if err != nil {
			return nil, err
		}
	}

	for _, slot := range updates.ItemsUnequipped {
		err = cai.adapter.UnequipItem(ctx, actorID, slot)
		if err != nil {
			return nil, err
		}
	}

	// Grant/revoke titles
	for titleID, title := range updates.TitlesGranted {
		err = cai.adapter.GrantTitle(ctx, actorID, titleID, title)
		if err != nil {
			return nil, err
		}
	}

	// Build new contribution
	return cai.BuildCoreActorContribution(ctx, actorID)
}

// PlayerStatsUpdates represents updates to player stats
type PlayerStatsUpdates struct {
	XPGranted       int64
	LevelIncrease   int64
	StatAllocations map[model.StatKey]int64
	EffectsAdded    []model.StatModifier
	EffectsRemoved  []string
	ItemsEquipped   map[string]model.StatModifier
	ItemsUnequipped []string
	TitlesGranted   map[string]model.StatModifier
	TitlesRevoked   []string
}

// GetStatBreakdown returns detailed breakdown for a specific stat
func (cai *CoreActorIntegration) GetStatBreakdown(ctx context.Context, actorID string, statKey model.StatKey) (*model.StatBreakdown, error) {
	// Get player data
	summary, err := cai.adapter.GetPlayerStatsSummary(ctx, actorID)
	if err != nil {
		return nil, err
	}

	// Get stat registry from the progression service
	primaryStats := cai.progressionService.registry.GetAllPrimaryStats()
	derivedStats := cai.progressionService.registry.GetAllDerivedStats()

	// Combine primary and derived stats
	registry := make([]model.StatDef, 0, len(primaryStats)+len(derivedStats))
	for _, stat := range primaryStats {
		registry = append(registry, *stat)
	}
	for _, stat := range derivedStats {
		registry = append(registry, *stat)
	}

	// Build compute input with breakdown enabled
	input := model.ComputeInput{
		ActorID:         actorID,
		Level:           summary.Progress.Level,
		BaseAllocations: summary.Progress.Allocations,
		Registry:        registry,
		Items:           summary.Equipment,
		Titles:          summary.Titles,
		Passives:        []model.StatModifier{},
		Buffs:           summary.Effects,
		Debuffs:         []model.StatModifier{},
		Auras:           []model.StatModifier{},
		Environment:     []model.StatModifier{},
		WithBreakdown:   true,
	}

	// Compute snapshot with breakdown
	snapshot := cai.resolver.ComputeSnapshot(input)

	// Return breakdown for the requested stat
	if breakdown, exists := snapshot.Breakdown[statKey]; exists {
		return breakdown, nil
	}

	return nil, nil
}
