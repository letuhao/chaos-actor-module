//go:build !race
// +build !race

package core_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type Inputs struct {
	Strength     int `json:"strength"`
	Agility      int `json:"agility"`
	Vitality     int `json:"vitality"`
	Intelligence int `json:"intelligence"`
	Spirit       int `json:"spirit"`
	Luck         int `json:"luck"`

	// mastery placeholders (would normally come from gear/skills)
	MasteryHP      float64 `json:"mastery_hp"`
	MasteryMP      float64 `json:"mastery_mp"`
	MasteryStamina float64 `json:"mastery_stamina"`
	MasteryAttack  float64 `json:"mastery_attack"`
	MasteryMagic   float64 `json:"mastery_magic"`
	MasteryDefense float64 `json:"mastery_defense"`
	MasteryResist  float64 `json:"mastery_resist"`
	MasterySpeed   float64 `json:"mastery_speed"`
}

func clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func calc(v Inputs) map[string]float64 {
	m := map[string]float64{}
	m["hp_max"] = clamp(float64(v.Vitality*20+v.Strength*5)*(1.0+v.MasteryHP*0.02), 1, 2000000)
	m["mp_max"] = clamp(float64(v.Intelligence*18+v.Spirit*12)*(1.0+v.MasteryMP*0.02), 0, 2000000)
	m["stamina_max"] = clamp(float64(v.Vitality*5+v.Agility*8)*(1.0+v.MasteryStamina*0.02), 0, 100000)
	m["attack_power"] = clamp(float64(v.Strength*3+v.Agility*1)*(1.0+v.MasteryAttack*0.03), 1, 1000000)
	m["magic_power"] = clamp(float64(v.Intelligence*3+v.Spirit*1)*(1.0+v.MasteryMagic*0.03), 0, 1000000)
	m["defense"] = clamp(float64(v.Vitality*2+v.Strength*1)*(1.0+v.MasteryDefense*0.02), 0, 500000)
	m["resist"] = clamp(float64(v.Spirit*2+v.Intelligence*1)*(1.0+v.MasteryResist*0.02), 0, 500000)
	m["crit_rate"] = clamp(float64(v.Luck)*0.001+float64(v.Agility)*0.0005, 0, 1)
	m["crit_damage"] = clamp(1.5+float64(v.Luck)*0.001, 1, 3)
	m["accuracy"] = clamp(0.7+float64(v.Agility)*0.0008+float64(v.Luck)*0.0004, 0, 1)
	m["evasion"] = clamp(float64(v.Agility)*0.0009+float64(v.Luck)*0.0003, 0, 0.8)
	m["block_chance"] = clamp(float64(v.Strength)*0.0006+float64(v.Vitality)*0.0004, 0, 0.6)
	m["move_speed"] = clamp((5.0+float64(v.Agility)*0.01)*(1.0+v.MasterySpeed*0.01), 0, 12)
	m["hp_regen"] = clamp(float64(v.Vitality)*0.1+float64(v.Spirit)*0.05, 0, 10000)
	m["mp_regen"] = clamp(float64(v.Spirit)*0.12+float64(v.Intelligence)*0.08, 0, 10000)
	m["stamina_regen"] = clamp(float64(v.Vitality)*0.1+float64(v.Agility)*0.1, 0, 10000)
	m["cooldown_reduction"] = clamp(float64(v.Intelligence)*0.0005+float64(v.Spirit)*0.0003, 0, 0.5)
	m["lifesteal"] = clamp(float64(v.Luck)*0.0004+float64(v.Spirit)*0.0002, 0, 0.3)
	return m
}

func Test_Derived_All_Golden(t *testing.T) {
	vectors := []Inputs{
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{10, 10, 10, 10, 10, 10, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1},
		{50, 50, 50, 50, 50, 50, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2},
		{100, 100, 100, 100, 100, 100, 0.3, 0.3, 0.3, 0.3, 0.3, 0.3, 0.3, 0.3},
	}
	got := make([]map[string]float64, 0, len(vectors))
	for _, v := range vectors {
		got = append(got, calc(v))
	}
	b, _ := json.MarshalIndent(got, "", "  ")
	gold := filepath.Join("testdata", "derived", "all.golden.json")
	want, err := os.ReadFile(gold)
	if err != nil {
		// first run creates golden
		_ = os.MkdirAll(filepath.Dir(gold), 0o755)
		_ = os.WriteFile(gold, b, 0o644)
		return
	}
	if string(want) != string(b) {
		t.Errorf("golden mismatch\nwant:\n%s\n----\ngot:\n%s", string(want), string(b))
	}
}
