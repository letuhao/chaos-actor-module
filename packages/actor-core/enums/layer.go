package enums

// Layer represents the cap layer scope
type Layer string

const (
	// LayerRealm represents realm-level caps
	LayerRealm Layer = "REALM"

	// LayerWorld represents world-level caps
	LayerWorld Layer = "WORLD"

	// LayerEvent represents event-level caps
	LayerEvent Layer = "EVENT"

	// LayerGuild represents guild-level caps
	LayerGuild Layer = "GUILD"

	// LayerTotal represents total-level caps
	LayerTotal Layer = "TOTAL"
)

// IsValid checks if the layer is valid
func (l Layer) IsValid() bool {
	switch l {
	case LayerRealm, LayerWorld, LayerEvent, LayerGuild, LayerTotal:
		return true
	default:
		return false
	}
}

// String returns the string representation of the layer
func (l Layer) String() string {
	return string(l)
}

// GetOrder returns the processing order for the layer
// Lower numbers are processed first
func (l Layer) GetOrder() int64 {
	switch l {
	case LayerRealm:
		return 1
	case LayerWorld:
		return 2
	case LayerEvent:
		return 3
	case LayerGuild:
		return 4
	case LayerTotal:
		return 5
	default:
		return 999
	}
}

// IsRealmLevel checks if the layer is realm level
func (l Layer) IsRealmLevel() bool {
	return l == LayerRealm
}

// IsWorldLevel checks if the layer is world level
func (l Layer) IsWorldLevel() bool {
	return l == LayerWorld
}

// IsEventLevel checks if the layer is event level
func (l Layer) IsEventLevel() bool {
	return l == LayerEvent
}

// IsGuildLevel checks if the layer is guild level
func (l Layer) IsGuildLevel() bool {
	return l == LayerGuild
}

// IsTotalLevel checks if the layer is total level
func (l Layer) IsTotalLevel() bool {
	return l == LayerTotal
}

// GetDefaultOrder returns the default layer processing order
func GetDefaultOrder() []Layer {
	return []Layer{
		LayerRealm,
		LayerWorld,
		LayerEvent,
		LayerTotal,
	}
}

// GetExtendedOrder returns the extended layer processing order (including guild)
func GetExtendedOrder() []Layer {
	return []Layer{
		LayerRealm,
		LayerWorld,
		LayerEvent,
		LayerGuild,
		LayerTotal,
	}
}
