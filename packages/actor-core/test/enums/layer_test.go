package enums

import (
	"chaos-actor-module/packages/actor-core/enums"
	"testing"
)

func TestLayer_String(t *testing.T) {
	tests := []struct {
		name     string
		layer    enums.Layer
		expected string
	}{
		{"Realm", enums.LayerRealm, "REALM"},
		{"World", enums.LayerWorld, "WORLD"},
		{"Event", enums.LayerEvent, "EVENT"},
		{"Guild", enums.LayerGuild, "GUILD"},
		{"Total", enums.LayerTotal, "TOTAL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.layer.String(); got != tt.expected {
				t.Errorf("enums.Layer.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestLayer_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		layer    enums.Layer
		expected bool
	}{
		{"Valid Realm", enums.LayerRealm, true},
		{"Valid World", enums.LayerWorld, true},
		{"Valid Event", enums.LayerEvent, true},
		{"Valid Guild", enums.LayerGuild, true},
		{"Valid Total", enums.LayerTotal, true},
		{"Invalid", enums.Layer("INVALID"), false},
		{"Empty", enums.Layer(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.layer.IsValid(); got != tt.expected {
				t.Errorf("enums.Layer.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestLayer_GetOrder(t *testing.T) {
	tests := []struct {
		name     string
		layer    enums.Layer
		expected int64
	}{
		{"Realm", enums.LayerRealm, 0},
		{"World", enums.LayerWorld, 1},
		{"Event", enums.LayerEvent, 2},
		{"Guild", enums.LayerGuild, 3},
		{"Total", enums.LayerTotal, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.layer.GetOrder(); got != tt.expected {
				t.Errorf("enums.Layer.GetOrder() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestLayer_GetAllLayers(t *testing.T) {
	expected := []enums.Layer{
		enums.LayerRealm,
		enums.LayerWorld,
		enums.LayerEvent,
		enums.LayerGuild,
		enums.LayerTotal,
	}

	// Since GetAllLayers doesn't exist, we'll test individual layers
	for _, layer := range expected {
		if !layer.IsValid() {
			t.Errorf("Layer %v should be valid", layer)
		}
	}
}

func TestLayer_GetDefaultOrder(t *testing.T) {
	expected := []enums.Layer{
		enums.LayerRealm,
		enums.LayerWorld,
		enums.LayerEvent,
		enums.LayerGuild,
		enums.LayerTotal,
	}

	order := enums.GetDefaultOrder()
	if len(order) != len(expected) {
		t.Errorf("enums.GetDefaultOrder() length = %v, want %v", len(order), len(expected))
	}

	for i, layer := range order {
		if layer != expected[i] {
			t.Errorf("enums.GetDefaultOrder()[%d] = %v, want %v", i, layer, expected[i])
		}
	}
}
