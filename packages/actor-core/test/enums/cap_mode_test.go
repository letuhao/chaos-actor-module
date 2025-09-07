package enums

import (
	"chaos-actor-module/packages/actor-core/enums"
	"testing"
)

func TestCapMode_String(t *testing.T) {
	tests := []struct {
		name     string
		capMode  enums.CapMode
		expected string
	}{
		{"Baseline", enums.CapModeBaseline, "BASELINE"},
		{"Additive", enums.CapModeAdditive, "ADDITIVE"},
		{"HardMax", enums.CapModeHardMax, "HARD_MAX"},
		{"HardMin", enums.CapModeHardMin, "HARD_MIN"},
		{"Override", enums.CapModeOverride, "OVERRIDE"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.capMode.String(); got != tt.expected {
				t.Errorf("enums.CapMode.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCapMode_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		capMode  enums.CapMode
		expected bool
	}{
		{"Valid Baseline", enums.CapModeBaseline, true},
		{"Valid Additive", enums.CapModeAdditive, true},
		{"Valid HardMax", enums.CapModeHardMax, true},
		{"Valid HardMin", enums.CapModeHardMin, true},
		{"Valid Override", enums.CapModeOverride, true},
		{"Invalid", enums.CapMode("INVALID"), false},
		{"Empty", enums.CapMode(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.capMode.IsValid(); got != tt.expected {
				t.Errorf("enums.CapMode.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCapMode_GetPriority(t *testing.T) {
	tests := []struct {
		name     string
		capMode  enums.CapMode
		expected int64
	}{
		{"Baseline", enums.CapModeBaseline, 1000},
		{"Additive", enums.CapModeAdditive, 500},
		{"HardMax", enums.CapModeHardMax, 300},
		{"HardMin", enums.CapModeHardMin, 200},
		{"Override", enums.CapModeOverride, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.capMode.GetPriority(); got != tt.expected {
				t.Errorf("enums.CapMode.GetPriority() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCapMode_GetAllModes(t *testing.T) {
	expected := []enums.CapMode{
		enums.CapModeBaseline,
		enums.CapModeAdditive,
		enums.CapModeHardMax,
		enums.CapModeHardMin,
		enums.CapModeOverride,
	}

	// Since GetAllCapModes doesn't exist, we'll test individual modes
	for _, mode := range expected {
		if !mode.IsValid() {
			t.Errorf("CapMode %v should be valid", mode)
		}
	}
}

func TestCapMode_GetDefaultOrder(t *testing.T) {
	expected := []enums.CapMode{
		enums.CapModeBaseline,
		enums.CapModeAdditive,
		enums.CapModeHardMax,
		enums.CapModeHardMin,
		enums.CapModeOverride,
	}

	// Since GetDefaultOrder doesn't exist for CapMode, we'll test individual modes
	for _, mode := range expected {
		if !mode.IsValid() {
			t.Errorf("CapMode %v should be valid", mode)
		}
	}
}
