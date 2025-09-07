//go:build !race
// +build !race

package core_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type Vector struct {
	Vitality     int `json:"vitality"`
	Constitution int `json:"constitution"`
	Agility      int `json:"agility"`
}

func TestDerived_HPMax_Golden(t *testing.T) {
	vectors := []Vector{{1, 1, 1}, {10, 10, 10}, {50, 50, 50}, {100, 100, 100}}
	got := make([]float64, 0, len(vectors))
	for _, v := range vectors {
		// TODO: wire resolver from your package
		// hp := resolver.ResolveHPMax(v)
		hp := float64(v.Vitality*10 + v.Constitution*5) // placeholder mirrors schema example
		got = append(got, hp)
	}
	b, _ := json.MarshalIndent(got, "", "  ")
	gold := filepath.Join("testdata", "derived", "hp_max.golden.json")
	want, err := os.ReadFile(gold)
	if err != nil {
		// First run: write golden
		os.MkdirAll(filepath.Dir(gold), 0o755)
		_ = os.WriteFile(gold, b, 0o644)
		return
	}
	if string(want) != string(b) {
		t.Errorf("golden mismatch\nwant:\n%s\n----\ngot:\n%s", string(want), string(b))
	}
}
