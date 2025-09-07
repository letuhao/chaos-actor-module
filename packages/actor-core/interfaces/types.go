package interfaces

import "time"

// Forward declarations to avoid circular imports
// These will be implemented in the types package

type Actor struct {
	ID      string
	Version int64
}

type SubsystemOutput struct {
	Primary []Contribution
	Derived []Contribution
	Caps    []CapContribution
	Context map[string]ModifierPack
	Meta    SubsystemMeta
}

type Snapshot struct {
	ActorID   string
	Primary   map[string]float64
	Derived   map[string]float64
	CapsUsed  map[string]Caps
	Version   int64
	CreatedAt time.Time
}

// Caps is defined in caps_provider.go
type Contribution struct {
	Dimension string
	Bucket    string
	Value     float64
	System    string
	Priority  int64
}

type CapContribution struct {
	System    string
	Dimension string
	Mode      string
	Kind      string
	Value     float64
	Priority  int64
	Scope     string
	Realm     string
	Tags      []string
}

type SubsystemMeta struct {
	System     string
	Version    int64
	APILevel   int64
	Compatible bool
}

type ModifierPack struct {
	AdditivePercent float64
	Multipliers     []float64
	PostAdd         float64
}

// EffectiveCaps will be defined in caps_provider.go
