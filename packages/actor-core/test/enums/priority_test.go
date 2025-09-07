package enums

import (
	"chaos-actor-module/packages/actor-core/enums"
	"testing"
)

func TestPriority_String(t *testing.T) {
	tests := []struct {
		name     string
		priority enums.Priority
		expected string
	}{
		{"Highest", enums.PriorityHighest, "1000"},
		{"High", enums.PriorityHigh, "500"},
		{"Normal", enums.PriorityNormal, "0"},
		{"Low", enums.PriorityLow, "-500"},
		{"Lowest", enums.PriorityLowest, "-1000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.priority.String(); got != tt.expected {
				t.Errorf("enums.Priority.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPriority_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		priority enums.Priority
		expected bool
	}{
		{"Valid Highest", enums.PriorityHighest, true},
		{"Valid High", enums.PriorityHigh, true},
		{"Valid Normal", enums.PriorityNormal, true},
		{"Valid Low", enums.PriorityLow, true},
		{"Valid Lowest", enums.PriorityLowest, true},
		{"Valid Custom", enums.Priority(100), true},
		{"Valid Negative", enums.Priority(-100), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.priority.IsValid(); got != tt.expected {
				t.Errorf("enums.Priority.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPriority_Int64(t *testing.T) {
	tests := []struct {
		name     string
		priority enums.Priority
		expected int64
	}{
		{"Highest", enums.PriorityHighest, 1000},
		{"High", enums.PriorityHigh, 500},
		{"Normal", enums.PriorityNormal, 0},
		{"Low", enums.PriorityLow, -500},
		{"Lowest", enums.PriorityLowest, -1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.priority.Int64(); got != tt.expected {
				t.Errorf("enums.Priority.Int64() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPriority_IsHigherThan(t *testing.T) {
	tests := []struct {
		name     string
		priority enums.Priority
		other    enums.Priority
		expected bool
	}{
		{"Highest > High", enums.PriorityHighest, enums.PriorityHigh, true},
		{"High > Normal", enums.PriorityHigh, enums.PriorityNormal, true},
		{"Normal > Low", enums.PriorityNormal, enums.PriorityLow, true},
		{"Low > Lowest", enums.PriorityLow, enums.PriorityLowest, true},
		{"High < Highest", enums.PriorityHigh, enums.PriorityHighest, false},
		{"Normal < High", enums.PriorityNormal, enums.PriorityHigh, false},
		{"Low < Normal", enums.PriorityLow, enums.PriorityNormal, false},
		{"Lowest < Low", enums.PriorityLowest, enums.PriorityLow, false},
		{"Equal priorities", enums.PriorityNormal, enums.PriorityNormal, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.priority.IsHigherThan(tt.other); got != tt.expected {
				t.Errorf("enums.Priority.IsHigherThan() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPriority_IsLowerThan(t *testing.T) {
	tests := []struct {
		name     string
		priority enums.Priority
		other    enums.Priority
		expected bool
	}{
		{"Highest < High", enums.PriorityHighest, enums.PriorityHigh, false},
		{"High < Normal", enums.PriorityHigh, enums.PriorityNormal, false},
		{"Normal < Low", enums.PriorityNormal, enums.PriorityLow, false},
		{"Low < Lowest", enums.PriorityLow, enums.PriorityLowest, false},
		{"High > Highest", enums.PriorityHigh, enums.PriorityHighest, true},
		{"Normal > High", enums.PriorityNormal, enums.PriorityHigh, true},
		{"Low > Normal", enums.PriorityLow, enums.PriorityNormal, true},
		{"Lowest > Low", enums.PriorityLowest, enums.PriorityLow, true},
		{"Equal priorities", enums.PriorityNormal, enums.PriorityNormal, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.priority.IsLowerThan(tt.other); got != tt.expected {
				t.Errorf("enums.Priority.IsLowerThan() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPriority_IsEqual(t *testing.T) {
	tests := []struct {
		name     string
		priority enums.Priority
		other    enums.Priority
		expected bool
	}{
		{"Equal priorities", enums.PriorityNormal, enums.PriorityNormal, true},
		{"Different priorities", enums.PriorityHigh, enums.PriorityLow, false},
		{"Same value", enums.Priority(100), enums.Priority(100), true},
		{"Different values", enums.Priority(100), enums.Priority(200), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.priority.IsEqual(tt.other); got != tt.expected {
				t.Errorf("enums.Priority.IsEqual() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPriority_GetDefaultPriority(t *testing.T) {
	expected := enums.PriorityNormal
	if got := enums.GetDefaultPriority(); got != expected {
		t.Errorf("GetDefaultPriority() = %v, want %v", got, expected)
	}
}
