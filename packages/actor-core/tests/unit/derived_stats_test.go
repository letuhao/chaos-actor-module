package unit

import (
	"testing"
	"time"

	"actor-core-v2/models/core"
)

func TestNewDerivedStats(t *testing.T) {
	ds := core.NewDerivedStats()

	// Test default values
	if ds.HPMax != 100.0 {
		t.Errorf("Expected HPMax to be 100.0, got %f", ds.HPMax)
	}

	if ds.Stamina != 100.0 {
		t.Errorf("Expected Stamina to be 100.0, got %f", ds.Stamina)
	}

	if ds.Speed != 1.0 {
		t.Errorf("Expected Speed to be 1.0, got %f", ds.Speed)
	}

	if ds.Haste != 1.0 {
		t.Errorf("Expected Haste to be 1.0, got %f", ds.Haste)
	}

	if ds.CritChance != 0.05 {
		t.Errorf("Expected CritChance to be 0.05, got %f", ds.CritChance)
	}

	if ds.CritMulti != 1.5 {
		t.Errorf("Expected CritMulti to be 1.5, got %f", ds.CritMulti)
	}

	if ds.MoveSpeed != 1.0 {
		t.Errorf("Expected MoveSpeed to be 1.0, got %f", ds.MoveSpeed)
	}

	if ds.RegenHP != 0.1 {
		t.Errorf("Expected RegenHP to be 0.1, got %f", ds.RegenHP)
	}

	if ds.Version != 1 {
		t.Errorf("Expected Version to be 1, got %d", ds.Version)
	}

	if ds.CreatedAt == 0 {
		t.Error("Expected CreatedAt to be set")
	}

	if ds.UpdatedAt == 0 {
		t.Error("Expected UpdatedAt to be set")
	}
}

func TestCalculateFromPrimary(t *testing.T) {
	pc := core.NewPrimaryCore()
	pc.Vitality = 20
	pc.Constitution = 15
	pc.Agility = 25
	pc.Intelligence = 30
	pc.Charisma = 35
	pc.Luck = 40
	pc.SpiritualEnergy = 50
	pc.PhysicalEnergy = 60
	pc.MentalEnergy = 70
	pc.Strength = 80
	pc.Willpower = 90
	pc.Wisdom = 100

	ds := core.NewDerivedStats()
	ds.CalculateFromPrimary(pc)

	// Test HPMax calculation: Vitality * 10 + Constitution * 5
	expectedHPMax := float64(20*10 + 15*5)
	if ds.HPMax != expectedHPMax {
		t.Errorf("Expected HPMax to be %f, got %f", expectedHPMax, ds.HPMax)
	}

	// Test Stamina calculation: Endurance * 10 + Constitution * 3
	expectedStamina := float64(10*10 + 15*3) // Endurance is 10 by default
	if ds.Stamina != expectedStamina {
		t.Errorf("Expected Stamina to be %f, got %f", expectedStamina, ds.Stamina)
	}

	// Test Speed calculation: Agility * 0.1
	expectedSpeed := float64(25) * 0.1
	if ds.Speed != expectedSpeed {
		t.Errorf("Expected Speed to be %f, got %f", expectedSpeed, ds.Speed)
	}

	// Test Haste calculation: 1.0 + Agility * 0.01
	expectedHaste := 1.0 + float64(25)*0.01
	if ds.Haste != expectedHaste {
		t.Errorf("Expected Haste to be %f, got %f", expectedHaste, ds.Haste)
	}

	// Test CritChance calculation: 0.05 + Luck * 0.001
	expectedCritChance := 0.05 + float64(40)*0.001
	if ds.CritChance != expectedCritChance {
		t.Errorf("Expected CritChance to be %f, got %f", expectedCritChance, ds.CritChance)
	}

	// Test CritMulti calculation: 1.5 + Luck * 0.01
	expectedCritMulti := 1.5 + float64(40)*0.01
	if ds.CritMulti != expectedCritMulti {
		t.Errorf("Expected CritMulti to be %f, got %f", expectedCritMulti, ds.CritMulti)
	}

	// Test MoveSpeed calculation: Agility * 0.1
	expectedMoveSpeed := float64(25) * 0.1
	if ds.MoveSpeed != expectedMoveSpeed {
		t.Errorf("Expected MoveSpeed to be %f, got %f", expectedMoveSpeed, ds.MoveSpeed)
	}

	// Test RegenHP calculation: Vitality * 0.01
	expectedRegenHP := float64(20) * 0.01
	if ds.RegenHP != expectedRegenHP {
		t.Errorf("Expected RegenHP to be %f, got %f", expectedRegenHP, ds.RegenHP)
	}

	// Test Accuracy calculation: 0.8 + Intelligence * 0.01
	expectedAccuracy := 0.8 + float64(30)*0.01
	if ds.Accuracy != expectedAccuracy {
		t.Errorf("Expected Accuracy to be %f, got %f", expectedAccuracy, ds.Accuracy)
	}

	// Test Penetration calculation: Strength * 0.01
	expectedPenetration := float64(80) * 0.01
	if ds.Penetration != expectedPenetration {
		t.Errorf("Expected Penetration to be %f, got %f", expectedPenetration, ds.Penetration)
	}

	// Test Lethality calculation: (Strength + Agility) * 0.005
	expectedLethality := float64(80+25) * 0.005
	if ds.Lethality != expectedLethality {
		t.Errorf("Expected Lethality to be %f, got %f", expectedLethality, ds.Lethality)
	}

	// Test Brutality calculation: Strength * 0.01
	expectedBrutality := float64(80) * 0.01
	if ds.Brutality != expectedBrutality {
		t.Errorf("Expected Brutality to be %f, got %f", expectedBrutality, ds.Brutality)
	}

	// Test ArmorClass calculation: 10.0 + Constitution * 0.5
	expectedArmorClass := 10.0 + float64(15)*0.5
	if ds.ArmorClass != expectedArmorClass {
		t.Errorf("Expected ArmorClass to be %f, got %f", expectedArmorClass, ds.ArmorClass)
	}

	// Test Evasion calculation: Agility * 0.01
	expectedEvasion := float64(25) * 0.01
	if ds.Evasion != expectedEvasion {
		t.Errorf("Expected Evasion to be %f, got %f", expectedEvasion, ds.Evasion)
	}

	// Test BlockChance calculation: Constitution * 0.005
	expectedBlockChance := float64(15) * 0.005
	if ds.BlockChance != expectedBlockChance {
		t.Errorf("Expected BlockChance to be %f, got %f", expectedBlockChance, ds.BlockChance)
	}

	// Test ParryChance calculation: Agility * 0.005
	expectedParryChance := float64(25) * 0.005
	if ds.ParryChance != expectedParryChance {
		t.Errorf("Expected ParryChance to be %f, got %f", expectedParryChance, ds.ParryChance)
	}

	// Test DodgeChance calculation: Agility * 0.01
	expectedDodgeChance := float64(25) * 0.01
	if ds.DodgeChance != expectedDodgeChance {
		t.Errorf("Expected DodgeChance to be %f, got %f", expectedDodgeChance, ds.DodgeChance)
	}

	// Test EnergyEfficiency calculation: 1.0 + Intelligence * 0.01
	expectedEnergyEfficiency := 1.0 + float64(30)*0.01
	if ds.EnergyEfficiency != expectedEnergyEfficiency {
		t.Errorf("Expected EnergyEfficiency to be %f, got %f", expectedEnergyEfficiency, ds.EnergyEfficiency)
	}

	// Test EnergyCapacity calculation: SpiritualEnergy + PhysicalEnergy + MentalEnergy
	expectedEnergyCapacity := float64(50 + 60 + 70)
	if ds.EnergyCapacity != expectedEnergyCapacity {
		t.Errorf("Expected EnergyCapacity to be %f, got %f", expectedEnergyCapacity, ds.EnergyCapacity)
	}

	// Test EnergyDrain calculation: Willpower * 0.01
	expectedEnergyDrain := float64(90) * 0.01
	if ds.EnergyDrain != expectedEnergyDrain {
		t.Errorf("Expected EnergyDrain to be %f, got %f", expectedEnergyDrain, ds.EnergyDrain)
	}

	// Test ResourceRegen calculation: Vitality * 0.01
	expectedResourceRegen := float64(20) * 0.01
	if ds.ResourceRegen != expectedResourceRegen {
		t.Errorf("Expected ResourceRegen to be %f, got %f", expectedResourceRegen, ds.ResourceRegen)
	}

	// Test LearningRate calculation: 1.0 + Intelligence * 0.01
	expectedLearningRate := 1.0 + float64(30)*0.01
	if ds.LearningRate != expectedLearningRate {
		t.Errorf("Expected LearningRate to be %f, got %f", expectedLearningRate, ds.LearningRate)
	}

	// Test Adaptation calculation: 1.0 + Wisdom * 0.01
	expectedAdaptation := 1.0 + float64(100)*0.01
	if ds.Adaptation != expectedAdaptation {
		t.Errorf("Expected Adaptation to be %f, got %f", expectedAdaptation, ds.Adaptation)
	}

	// Test Memory calculation: 1.0 + Intelligence * 0.01
	expectedMemory := 1.0 + float64(30)*0.01
	if ds.Memory != expectedMemory {
		t.Errorf("Expected Memory to be %f, got %f", expectedMemory, ds.Memory)
	}

	// Test Experience calculation: 1.0 + Wisdom * 0.01
	expectedExperience := 1.0 + float64(100)*0.01
	if ds.Experience != expectedExperience {
		t.Errorf("Expected Experience to be %f, got %f", expectedExperience, ds.Experience)
	}

	// Test Leadership calculation: 1.0 + Charisma * 0.01
	expectedLeadership := 1.0 + float64(35)*0.01
	if ds.Leadership != expectedLeadership {
		t.Errorf("Expected Leadership to be %f, got %f", expectedLeadership, ds.Leadership)
	}

	// Test Diplomacy calculation: 1.0 + (Charisma + Intelligence) * 0.005
	expectedDiplomacy := 1.0 + float64(35+30)*0.005
	if ds.Diplomacy != expectedDiplomacy {
		t.Errorf("Expected Diplomacy to be %f, got %f", expectedDiplomacy, ds.Diplomacy)
	}

	// Test Intimidation calculation: 1.0 + (Strength + Charisma) * 0.005
	expectedIntimidation := 1.0 + float64(80+35)*0.005
	if ds.Intimidation != expectedIntimidation {
		t.Errorf("Expected Intimidation to be %f, got %f", expectedIntimidation, ds.Intimidation)
	}

	// Test Empathy calculation: 1.0 + (Wisdom + Charisma) * 0.005
	expectedEmpathy := 1.0 + float64(100+35)*0.005
	if ds.Empathy != expectedEmpathy {
		t.Errorf("Expected Empathy to be %f, got %f", expectedEmpathy, ds.Empathy)
	}

	// Test Deception calculation: 1.0 + (Intelligence + Charisma) * 0.005
	expectedDeception := 1.0 + float64(30+35)*0.005
	if ds.Deception != expectedDeception {
		t.Errorf("Expected Deception to be %f, got %f", expectedDeception, ds.Deception)
	}

	// Test Performance calculation: 1.0 + Charisma * 0.01
	expectedPerformance := 1.0 + float64(35)*0.01
	if ds.Performance != expectedPerformance {
		t.Errorf("Expected Performance to be %f, got %f", expectedPerformance, ds.Performance)
	}

	// Test ManaEfficiency calculation: 1.0 + Intelligence * 0.01
	expectedManaEfficiency := 1.0 + float64(30)*0.01
	if ds.ManaEfficiency != expectedManaEfficiency {
		t.Errorf("Expected ManaEfficiency to be %f, got %f", expectedManaEfficiency, ds.ManaEfficiency)
	}

	// Test SpellPower calculation: 1.0 + SpiritualEnergy * 0.01
	expectedSpellPower := 1.0 + float64(50)*0.01
	if ds.SpellPower != expectedSpellPower {
		t.Errorf("Expected SpellPower to be %f, got %f", expectedSpellPower, ds.SpellPower)
	}

	// Test MysticResonance calculation: 1.0 + (SpiritualEnergy + MentalEnergy) * 0.005
	expectedMysticResonance := 1.0 + float64(50+70)*0.005
	if ds.MysticResonance != expectedMysticResonance {
		t.Errorf("Expected MysticResonance to be %f, got %f", expectedMysticResonance, ds.MysticResonance)
	}

	// Test RealityBend calculation: 1.0 + MentalEnergy * 0.01
	expectedRealityBend := 1.0 + float64(70)*0.01
	if ds.RealityBend != expectedRealityBend {
		t.Errorf("Expected RealityBend to be %f, got %f", expectedRealityBend, ds.RealityBend)
	}

	// Test TimeSense calculation: 1.0 + MentalEnergy * 0.01
	expectedTimeSense := 1.0 + float64(70)*0.01
	if ds.TimeSense != expectedTimeSense {
		t.Errorf("Expected TimeSense to be %f, got %f", expectedTimeSense, ds.TimeSense)
	}

	// Test SpaceSense calculation: 1.0 + MentalEnergy * 0.01
	expectedSpaceSense := 1.0 + float64(70)*0.01
	if ds.SpaceSense != expectedSpaceSense {
		t.Errorf("Expected SpaceSense to be %f, got %f", expectedSpaceSense, ds.SpaceSense)
	}

	// Test JumpHeight calculation: 1.0 + (Strength + Agility) * 0.005
	expectedJumpHeight := 1.0 + float64(80+25)*0.005
	if ds.JumpHeight != expectedJumpHeight {
		t.Errorf("Expected JumpHeight to be %f, got %f", expectedJumpHeight, ds.JumpHeight)
	}

	// Test ClimbSpeed calculation: 1.0 + (Strength + Agility) * 0.005
	expectedClimbSpeed := 1.0 + float64(80+25)*0.005
	if ds.ClimbSpeed != expectedClimbSpeed {
		t.Errorf("Expected ClimbSpeed to be %f, got %f", expectedClimbSpeed, ds.ClimbSpeed)
	}

	// Test SwimSpeed calculation: 1.0 + (Strength + Agility) * 0.005
	expectedSwimSpeed := 1.0 + float64(80+25)*0.005
	if ds.SwimSpeed != expectedSwimSpeed {
		t.Errorf("Expected SwimSpeed to be %f, got %f", expectedSwimSpeed, ds.SwimSpeed)
	}

	// Test FlightSpeed calculation: 1.0 + SpiritualEnergy * 0.01
	expectedFlightSpeed := 1.0 + float64(50)*0.01
	if ds.FlightSpeed != expectedFlightSpeed {
		t.Errorf("Expected FlightSpeed to be %f, got %f", expectedFlightSpeed, ds.FlightSpeed)
	}

	// Test TeleportRange calculation: 1.0 + MentalEnergy * 0.01
	expectedTeleportRange := 1.0 + float64(70)*0.01
	if ds.TeleportRange != expectedTeleportRange {
		t.Errorf("Expected TeleportRange to be %f, got %f", expectedTeleportRange, ds.TeleportRange)
	}

	// Test Stealth calculation: 1.0 + Agility * 0.01
	expectedStealth := 1.0 + float64(25)*0.01
	if ds.Stealth != expectedStealth {
		t.Errorf("Expected Stealth to be %f, got %f", expectedStealth, ds.Stealth)
	}

	// Test AuraRadius calculation: 1.0 + SpiritualEnergy * 0.01
	expectedAuraRadius := 1.0 + float64(50)*0.01
	if ds.AuraRadius != expectedAuraRadius {
		t.Errorf("Expected AuraRadius to be %f, got %f", expectedAuraRadius, ds.AuraRadius)
	}

	// Test AuraStrength calculation: 1.0 + SpiritualEnergy * 0.01
	expectedAuraStrength := 1.0 + float64(50)*0.01
	if ds.AuraStrength != expectedAuraStrength {
		t.Errorf("Expected AuraStrength to be %f, got %f", expectedAuraStrength, ds.AuraStrength)
	}

	// Test Presence calculation: 1.0 + (Charisma + SpiritualEnergy) * 0.005
	expectedPresence := 1.0 + float64(35+50)*0.005
	if ds.Presence != expectedPresence {
		t.Errorf("Expected Presence to be %f, got %f", expectedPresence, ds.Presence)
	}

	// Test Awe calculation: 1.0 + (Charisma + SpiritualEnergy) * 0.005
	expectedAwe := 1.0 + float64(35+50)*0.005
	if ds.Awe != expectedAwe {
		t.Errorf("Expected Awe to be %f, got %f", expectedAwe, ds.Awe)
	}

	// Test WeaponMastery calculation: 1.0 + (Strength + Agility) * 0.005
	expectedWeaponMastery := 1.0 + float64(80+25)*0.005
	if ds.WeaponMastery != expectedWeaponMastery {
		t.Errorf("Expected WeaponMastery to be %f, got %f", expectedWeaponMastery, ds.WeaponMastery)
	}

	// Test SkillLevel calculation: 1.0 + Intelligence * 0.01
	expectedSkillLevel := 1.0 + float64(30)*0.01
	if ds.SkillLevel != expectedSkillLevel {
		t.Errorf("Expected SkillLevel to be %f, got %f", expectedSkillLevel, ds.SkillLevel)
	}

	// Test LifeSteal calculation: 0.0 (default)
	if ds.LifeSteal != 0.0 {
		t.Errorf("Expected LifeSteal to be 0.0, got %f", ds.LifeSteal)
	}

	// Test CastSpeed calculation: 1.0 + Intelligence * 0.01
	expectedCastSpeed := 1.0 + float64(30)*0.01
	if ds.CastSpeed != expectedCastSpeed {
		t.Errorf("Expected CastSpeed to be %f, got %f", expectedCastSpeed, ds.CastSpeed)
	}

	// Test WeightCapacity calculation: Strength * 10.0
	expectedWeightCapacity := float64(80) * 10.0
	if ds.WeightCapacity != expectedWeightCapacity {
		t.Errorf("Expected WeightCapacity to be %f, got %f", expectedWeightCapacity, ds.WeightCapacity)
	}

	// Test Persuasion calculation: 1.0 + Charisma * 0.01
	expectedPersuasion := 1.0 + float64(35)*0.01
	if ds.Persuasion != expectedPersuasion {
		t.Errorf("Expected Persuasion to be %f, got %f", expectedPersuasion, ds.Persuasion)
	}

	// Test MerchantPriceModifier calculation: 1.0 + Charisma * 0.01
	expectedMerchantPriceModifier := 1.0 + float64(35)*0.01
	if ds.MerchantPriceModifier != expectedMerchantPriceModifier {
		t.Errorf("Expected MerchantPriceModifier to be %f, got %f", expectedMerchantPriceModifier, ds.MerchantPriceModifier)
	}

	// Test FactionReputationGain calculation: 1.0 + Charisma * 0.01
	expectedFactionReputationGain := 1.0 + float64(35)*0.01
	if ds.FactionReputationGain != expectedFactionReputationGain {
		t.Errorf("Expected FactionReputationGain to be %f, got %f", expectedFactionReputationGain, ds.FactionReputationGain)
	}

	// Test CultivationSpeed calculation: 1.0 + (SpiritualEnergy + PhysicalEnergy + MentalEnergy) * 0.001
	expectedCultivationSpeed := 1.0 + float64(50+60+70)*0.001
	if ds.CultivationSpeed != expectedCultivationSpeed {
		t.Errorf("Expected CultivationSpeed to be %f, got %f", expectedCultivationSpeed, ds.CultivationSpeed)
	}

	// Test EnergyEfficiencyAmp calculation: 1.0 + Intelligence * 0.01
	expectedEnergyEfficiencyAmp := 1.0 + float64(30)*0.01
	if ds.EnergyEfficiencyAmp != expectedEnergyEfficiencyAmp {
		t.Errorf("Expected EnergyEfficiencyAmp to be %f, got %f", expectedEnergyEfficiencyAmp, ds.EnergyEfficiencyAmp)
	}

	// Test BreakthroughSuccess calculation: 1.0 + (Willpower + Luck) * 0.005
	expectedBreakthroughSuccess := 1.0 + float64(90+40)*0.005
	if ds.BreakthroughSuccess != expectedBreakthroughSuccess {
		t.Errorf("Expected BreakthroughSuccess to be %f, got %f", expectedBreakthroughSuccess, ds.BreakthroughSuccess)
	}

	// Test SkillLearning calculation: 1.0 + (Intelligence + Wisdom) * 0.005
	expectedSkillLearning := 1.0 + float64(30+100)*0.005
	if ds.SkillLearning != expectedSkillLearning {
		t.Errorf("Expected SkillLearning to be %f, got %f", expectedSkillLearning, ds.SkillLearning)
	}

	// Test CombatEffectiveness calculation: 1.0 + (Strength + Agility + Intelligence) * 0.003
	expectedCombatEffectiveness := 1.0 + float64(80+25+30)*0.003
	if ds.CombatEffectiveness != expectedCombatEffectiveness {
		t.Errorf("Expected CombatEffectiveness to be %f, got %f", expectedCombatEffectiveness, ds.CombatEffectiveness)
	}

	// Test ResourceGathering calculation: 1.0 + (Luck + Wisdom) * 0.005
	expectedResourceGathering := 1.0 + float64(40+100)*0.005
	if ds.ResourceGathering != expectedResourceGathering {
		t.Errorf("Expected ResourceGathering to be %f, got %f", expectedResourceGathering, ds.ResourceGathering)
	}

	// Test that version was incremented
	if ds.Version != 2 {
		t.Errorf("Expected version to be 2, got %d", ds.Version)
	}
}

func TestGetStat(t *testing.T) {
	ds := core.NewDerivedStats()

	// Test valid stat names
	stat, err := ds.GetStat("hp_max")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if stat != 100.0 {
		t.Errorf("Expected hp_max to be 100.0, got %f", stat)
	}

	stat, err = ds.GetStat("stamina")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if stat != 100.0 {
		t.Errorf("Expected stamina to be 100.0, got %f", stat)
	}

	// Test invalid stat name
	_, err = ds.GetStat("invalid_stat")
	if err == nil {
		t.Error("Expected error for invalid stat name")
	}
}

func TestSetStat(t *testing.T) {
	ds := core.NewDerivedStats()

	// Test setting a valid stat
	err := ds.SetStat("hp_max", 200.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if ds.HPMax != 200.0 {
		t.Errorf("Expected hp_max to be 200.0, got %f", ds.HPMax)
	}

	// Test setting stamina
	err = ds.SetStat("stamina", 150.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if ds.Stamina != 150.0 {
		t.Errorf("Expected stamina to be 150.0, got %f", ds.Stamina)
	}

	// Test setting invalid stat
	err = ds.SetStat("invalid_stat", 10.0)
	if err == nil {
		t.Error("Expected error for invalid stat name")
	}

	// Test that version was incremented
	if ds.Version != 3 {
		t.Errorf("Expected version to be 3, got %d", ds.Version)
	}
}

func TestGetAllStats(t *testing.T) {
	ds := core.NewDerivedStats()
	stats := ds.GetAllStats()

	// Test that all stats are present
	expectedStats := []string{
		"hp_max", "stamina", "speed", "haste", "crit_chance", "crit_multi",
		"move_speed", "regen_hp", "accuracy", "penetration", "lethality",
		"brutality", "armor_class", "evasion", "block_chance", "parry_chance",
		"dodge_chance", "energy_efficiency", "energy_capacity", "energy_drain",
		"resource_regen", "resource_decay", "learning_rate", "adaptation",
		"memory", "experience", "leadership", "diplomacy", "intimidation",
		"empathy", "deception", "performance", "mana_efficiency", "spell_power",
		"mystic_resonance", "reality_bend", "time_sense", "space_sense",
		"jump_height", "climb_speed", "swim_speed", "flight_speed",
		"teleport_range", "stealth", "aura_radius", "aura_strength",
		"presence", "awe", "weapon_mastery", "skill_level", "life_steal",
		"cast_speed", "weight_capacity", "persuasion", "merchant_price_modifier",
		"faction_reputation_gain", "cultivation_speed", "energy_efficiency_amp",
		"breakthrough_success", "skill_learning", "combat_effectiveness",
		"resource_gathering",
	}

	for _, statName := range expectedStats {
		if _, exists := stats[statName]; !exists {
			t.Errorf("Expected stat %s to be present", statName)
		}
	}

	// Test that values are correct
	if stats["hp_max"] != 100.0 {
		t.Errorf("Expected hp_max to be 100.0, got %f", stats["hp_max"])
	}

	if stats["stamina"] != 100.0 {
		t.Errorf("Expected stamina to be 100.0, got %f", stats["stamina"])
	}
}

func TestClone(t *testing.T) {
	ds := core.NewDerivedStats()
	ds.HPMax = 200.0
	ds.Stamina = 150.0
	ds.Version = 5

	cloned := ds.Clone()

	// Test that values are copied
	if cloned.HPMax != 200.0 {
		t.Errorf("Expected cloned hp_max to be 200.0, got %f", cloned.HPMax)
	}

	if cloned.Stamina != 150.0 {
		t.Errorf("Expected cloned stamina to be 150.0, got %f", cloned.Stamina)
	}

	if cloned.Version != 5 {
		t.Errorf("Expected cloned version to be 5, got %d", cloned.Version)
	}

	// Test that modifying clone doesn't affect original
	cloned.HPMax = 300.0
	if ds.HPMax != 200.0 {
		t.Error("Modifying clone should not affect original")
	}
}

func TestGetVersion(t *testing.T) {
	ds := core.NewDerivedStats()

	if ds.GetVersion() != 1 {
		t.Errorf("Expected version to be 1, got %d", ds.GetVersion())
	}

	ds.SetStat("hp_max", 200.0)

	if ds.GetVersion() != 2 {
		t.Errorf("Expected version to be 2, got %d", ds.GetVersion())
	}
}

func TestGetUpdatedAt(t *testing.T) {
	ds := core.NewDerivedStats()
	originalUpdatedAt := ds.GetUpdatedAt()

	// Wait a bit to ensure timestamp changes
	time.Sleep(1 * time.Second)

	ds.SetStat("hp_max", 200.0)

	if ds.GetUpdatedAt() <= originalUpdatedAt {
		t.Errorf("Expected UpdatedAt to be updated. Original: %d, New: %d", originalUpdatedAt, ds.GetUpdatedAt())
	}
}

func TestGetCreatedAt(t *testing.T) {
	ds := core.NewDerivedStats()

	if ds.GetCreatedAt() == 0 {
		t.Error("Expected CreatedAt to be set")
	}
}
