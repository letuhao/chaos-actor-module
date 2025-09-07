package types

import (
	"time"
)

// SubsystemOutput represents the output from a subsystem
type SubsystemOutput struct {
	// Primary are primary contributions
	Primary []Contribution `json:"primary"`

	// Derived are derived contributions
	Derived []Contribution `json:"derived"`

	// Caps are cap contributions
	Caps []CapContribution `json:"caps"`

	// Context are context modifiers
	Context map[string]ModifierPack `json:"context,omitempty"`

	// Meta contains subsystem metadata
	Meta SubsystemMeta `json:"meta"`
}

// IsValid checks if the subsystem output is valid
func (so *SubsystemOutput) IsValid() bool {
	// Validate primary contributions
	for _, contribution := range so.Primary {
		if !contribution.IsValid() {
			return false
		}
	}

	// Validate derived contributions
	for _, contribution := range so.Derived {
		if !contribution.IsValid() {
			return false
		}
	}

	// Validate cap contributions
	for _, cap := range so.Caps {
		if !cap.IsValid() {
			return false
		}
	}

	// Validate meta
	if !so.Meta.IsValid() {
		return false
	}

	return true
}

// GetPrimary returns the primary contributions
func (so *SubsystemOutput) GetPrimary() []Contribution {
	return so.Primary
}

// GetDerived returns the derived contributions
func (so *SubsystemOutput) GetDerived() []Contribution {
	return so.Derived
}

// GetCaps returns the cap contributions
func (so *SubsystemOutput) GetCaps() []CapContribution {
	return so.Caps
}

// GetContext returns the context modifiers
func (so *SubsystemOutput) GetContext() map[string]ModifierPack {
	return so.Context
}

// GetMeta returns the metadata
func (so *SubsystemOutput) GetMeta() SubsystemMeta {
	return so.Meta
}

// SetPrimary sets the primary contributions
func (so *SubsystemOutput) SetPrimary(primary []Contribution) {
	so.Primary = primary
}

// SetDerived sets the derived contributions
func (so *SubsystemOutput) SetDerived(derived []Contribution) {
	so.Derived = derived
}

// SetCaps sets the cap contributions
func (so *SubsystemOutput) SetCaps(caps []CapContribution) {
	so.Caps = caps
}

// SetContext sets the context modifiers
func (so *SubsystemOutput) SetContext(context map[string]ModifierPack) {
	so.Context = context
}

// SetMeta sets the metadata
func (so *SubsystemOutput) SetMeta(meta SubsystemMeta) {
	so.Meta = meta
}

// AddPrimary adds a primary contribution
func (so *SubsystemOutput) AddPrimary(contribution Contribution) {
	so.Primary = append(so.Primary, contribution)
}

// AddDerived adds a derived contribution
func (so *SubsystemOutput) AddDerived(contribution Contribution) {
	so.Derived = append(so.Derived, contribution)
}

// AddCap adds a cap contribution
func (so *SubsystemOutput) AddCap(cap CapContribution) {
	so.Caps = append(so.Caps, cap)
}

// AddContext adds a context modifier
func (so *SubsystemOutput) AddContext(contextType string, modifier ModifierPack) {
	if so.Context == nil {
		so.Context = make(map[string]ModifierPack)
	}
	so.Context[contextType] = modifier
}

// GetContextModifier returns a context modifier
func (so *SubsystemOutput) GetContextModifier(contextType string) (ModifierPack, bool) {
	if so.Context == nil {
		return ModifierPack{}, false
	}
	modifier, exists := so.Context[contextType]
	return modifier, exists
}

// SubsystemMeta represents subsystem metadata
type SubsystemMeta struct {
	// System is the system ID
	System string `json:"system"`

	// Version is the subsystem version
	Version int64 `json:"version"`

	// APILevel is the API level
	APILevel int64 `json:"api_level,omitempty"`

	// Compatible indicates if the subsystem is compatible
	Compatible bool `json:"compatible,omitempty"`

	// Timestamp is when the output was generated
	Timestamp time.Time `json:"timestamp,omitempty"`

	// Tags are additional tags
	Tags map[string]string `json:"tags,omitempty"`
}

// IsValid checks if the subsystem meta is valid
func (sm *SubsystemMeta) IsValid() bool {
	if sm.System == "" {
		return false
	}
	if sm.Version < 0 {
		return false
	}
	if sm.APILevel < 0 {
		return false
	}
	return true
}

// GetSystem returns the system
func (sm *SubsystemMeta) GetSystem() string {
	return sm.System
}

// GetVersion returns the version
func (sm *SubsystemMeta) GetVersion() int64 {
	return sm.Version
}

// GetAPILevel returns the API level
func (sm *SubsystemMeta) GetAPILevel() int64 {
	return sm.APILevel
}

// IsCompatible checks if the subsystem is compatible
func (sm *SubsystemMeta) IsCompatible() bool {
	return sm.Compatible
}

// GetTimestamp returns the timestamp
func (sm *SubsystemMeta) GetTimestamp() time.Time {
	return sm.Timestamp
}

// GetTags returns the tags
func (sm *SubsystemMeta) GetTags() map[string]string {
	return sm.Tags
}

// SetSystem sets the system
func (sm *SubsystemMeta) SetSystem(system string) {
	sm.System = system
}

// SetVersion sets the version
func (sm *SubsystemMeta) SetVersion(version int64) {
	sm.Version = version
}

// SetAPILevel sets the API level
func (sm *SubsystemMeta) SetAPILevel(apiLevel int64) {
	sm.APILevel = apiLevel
}

// SetCompatible sets the compatible flag
func (sm *SubsystemMeta) SetCompatible(compatible bool) {
	sm.Compatible = compatible
}

// SetTimestamp sets the timestamp
func (sm *SubsystemMeta) SetTimestamp(timestamp time.Time) {
	sm.Timestamp = timestamp
}

// SetTags sets the tags
func (sm *SubsystemMeta) SetTags(tags map[string]string) {
	sm.Tags = tags
}

// AddTag adds a tag
func (sm *SubsystemMeta) AddTag(key, value string) {
	if sm.Tags == nil {
		sm.Tags = make(map[string]string)
	}
	sm.Tags[key] = value
}

// GetTag returns a tag value
func (sm *SubsystemMeta) GetTag(key string) (string, bool) {
	if sm.Tags == nil {
		return "", false
	}
	value, exists := sm.Tags[key]
	return value, exists
}

// ModifierPack represents a context modifier pack
type ModifierPack struct {
	// AdditivePercent is the additive percentage
	AdditivePercent float64 `json:"additive_percent"`

	// Multipliers are the multipliers
	Multipliers []float64 `json:"multipliers"`

	// PostAdd is the post-addition value
	PostAdd float64 `json:"post_add"`
}

// IsValid checks if the modifier pack is valid
func (mp *ModifierPack) IsValid() bool {
	// Check for negative multipliers
	for _, multiplier := range mp.Multipliers {
		if multiplier < 0 {
			return false
		}
	}
	return true
}

// GetAdditivePercent returns the additive percentage
func (mp *ModifierPack) GetAdditivePercent() float64 {
	return mp.AdditivePercent
}

// GetMultipliers returns the multipliers
func (mp *ModifierPack) GetMultipliers() []float64 {
	return mp.Multipliers
}

// GetPostAdd returns the post-add value
func (mp *ModifierPack) GetPostAdd() float64 {
	return mp.PostAdd
}

// SetAdditivePercent sets the additive percentage
func (mp *ModifierPack) SetAdditivePercent(percent float64) {
	mp.AdditivePercent = percent
}

// SetMultipliers sets the multipliers
func (mp *ModifierPack) SetMultipliers(multipliers []float64) {
	mp.Multipliers = multipliers
}

// SetPostAdd sets the post-add value
func (mp *ModifierPack) SetPostAdd(postAdd float64) {
	mp.PostAdd = postAdd
}

// AddMultiplier adds a multiplier
func (mp *ModifierPack) AddMultiplier(multiplier float64) {
	mp.Multipliers = append(mp.Multipliers, multiplier)
}

// GetTotalMultiplier returns the total multiplier
func (mp *ModifierPack) GetTotalMultiplier() float64 {
	total := 1.0
	for _, multiplier := range mp.Multipliers {
		total *= multiplier
	}
	return total
}

// Apply applies the modifier pack to a base value
func (mp *ModifierPack) Apply(baseValue float64) float64 {
	// Step 1: Apply additive percentage
	additiveValue := baseValue * (1.0 + mp.AdditivePercent)

	// Step 2: Apply all multipliers
	result := additiveValue
	for _, multiplier := range mp.Multipliers {
		result *= multiplier
	}

	// Step 3: Apply post-add
	result += mp.PostAdd

	return result
}
