package types

import (
	"chaos-actor-module/packages/actor-core/enums"
)

// Contribution represents a contribution to a dimension
type Contribution struct {
	// Dimension is the dimension name
	Dimension string `json:"dimension"`

	// Bucket is the contribution bucket type
	Bucket enums.Bucket `json:"bucket"`

	// Value is the contribution value
	Value float64 `json:"value"`

	// System is the contributing system ID
	System string `json:"system"`

	// Priority is the processing priority
	Priority int64 `json:"priority,omitempty"`

	// Tags are additional tags
	Tags map[string]string `json:"tags,omitempty"`
}

// IsValid checks if the contribution is valid
func (c *Contribution) IsValid() bool {
	if c.Dimension == "" {
		return false
	}
	if !c.Bucket.IsValid() {
		return false
	}
	if c.System == "" {
		return false
	}
	if c.Priority < 0 {
		return false
	}
	return true
}

// GetDimension returns the dimension
func (c *Contribution) GetDimension() string {
	return c.Dimension
}

// GetBucket returns the bucket
func (c *Contribution) GetBucket() enums.Bucket {
	return c.Bucket
}

// GetValue returns the value
func (c *Contribution) GetValue() float64 {
	return c.Value
}

// GetSystem returns the system
func (c *Contribution) GetSystem() string {
	return c.System
}

// GetPriority returns the priority
func (c *Contribution) GetPriority() int64 {
	return c.Priority
}

// GetTags returns the tags
func (c *Contribution) GetTags() map[string]string {
	return c.Tags
}

// SetDimension sets the dimension
func (c *Contribution) SetDimension(dimension string) {
	c.Dimension = dimension
}

// SetBucket sets the bucket
func (c *Contribution) SetBucket(bucket enums.Bucket) {
	c.Bucket = bucket
}

// SetValue sets the value
func (c *Contribution) SetValue(value float64) {
	c.Value = value
}

// SetSystem sets the system
func (c *Contribution) SetSystem(system string) {
	c.System = system
}

// SetPriority sets the priority
func (c *Contribution) SetPriority(priority int64) {
	c.Priority = priority
}

// SetTags sets the tags
func (c *Contribution) SetTags(tags map[string]string) {
	c.Tags = tags
}

// AddTag adds a tag
func (c *Contribution) AddTag(key, value string) {
	if c.Tags == nil {
		c.Tags = make(map[string]string)
	}
	c.Tags[key] = value
}

// GetTag returns a tag value
func (c *Contribution) GetTag(key string) (string, bool) {
	if c.Tags == nil {
		return "", false
	}
	value, exists := c.Tags[key]
	return value, exists
}

// CapContribution represents a cap contribution
type CapContribution struct {
	// System is the contributing system ID
	System string `json:"system"`

	// Dimension is the dimension name
	Dimension string `json:"dimension"`

	// Mode is the cap mode
	Mode enums.CapMode `json:"mode"`

	// Kind is the cap kind (min, max)
	Kind string `json:"kind"`

	// Value is the cap value
	Value float64 `json:"value"`

	// Priority is the processing priority
	Priority int64 `json:"priority,omitempty"`

	// Scope is the cap scope
	Scope string `json:"scope,omitempty"`

	// Realm is the realm identifier
	Realm string `json:"realm,omitempty"`

	// Tags are additional tags
	Tags map[string]string `json:"tags,omitempty"`
}

// IsValid checks if the cap contribution is valid
func (cc *CapContribution) IsValid() bool {
	if cc.System == "" {
		return false
	}
	if cc.Dimension == "" {
		return false
	}
	if !cc.Mode.IsValid() {
		return false
	}
	if cc.Kind != "min" && cc.Kind != "max" {
		return false
	}
	if cc.Priority < 0 {
		return false
	}
	return true
}

// GetSystem returns the system
func (cc *CapContribution) GetSystem() string {
	return cc.System
}

// GetDimension returns the dimension
func (cc *CapContribution) GetDimension() string {
	return cc.Dimension
}

// GetMode returns the mode
func (cc *CapContribution) GetMode() enums.CapMode {
	return cc.Mode
}

// GetKind returns the kind
func (cc *CapContribution) GetKind() string {
	return cc.Kind
}

// GetValue returns the value
func (cc *CapContribution) GetValue() float64 {
	return cc.Value
}

// GetPriority returns the priority
func (cc *CapContribution) GetPriority() int64 {
	return cc.Priority
}

// GetScope returns the scope
func (cc *CapContribution) GetScope() string {
	return cc.Scope
}

// GetRealm returns the realm
func (cc *CapContribution) GetRealm() string {
	return cc.Realm
}

// GetTags returns the tags
func (cc *CapContribution) GetTags() map[string]string {
	return cc.Tags
}

// SetSystem sets the system
func (cc *CapContribution) SetSystem(system string) {
	cc.System = system
}

// SetDimension sets the dimension
func (cc *CapContribution) SetDimension(dimension string) {
	cc.Dimension = dimension
}

// SetMode sets the mode
func (cc *CapContribution) SetMode(mode enums.CapMode) {
	cc.Mode = mode
}

// SetKind sets the kind
func (cc *CapContribution) SetKind(kind string) {
	cc.Kind = kind
}

// SetValue sets the value
func (cc *CapContribution) SetValue(value float64) {
	cc.Value = value
}

// SetPriority sets the priority
func (cc *CapContribution) SetPriority(priority int64) {
	cc.Priority = priority
}

// SetScope sets the scope
func (cc *CapContribution) SetScope(scope string) {
	cc.Scope = scope
}

// SetRealm sets the realm
func (cc *CapContribution) SetRealm(realm string) {
	cc.Realm = realm
}

// SetTags sets the tags
func (cc *CapContribution) SetTags(tags map[string]string) {
	cc.Tags = tags
}

// AddTag adds a tag
func (cc *CapContribution) AddTag(key, value string) {
	if cc.Tags == nil {
		cc.Tags = make(map[string]string)
	}
	cc.Tags[key] = value
}

// GetTag returns a tag value
func (cc *CapContribution) GetTag(key string) (string, bool) {
	if cc.Tags == nil {
		return "", false
	}
	value, exists := cc.Tags[key]
	return value, exists
}

// IsMinCap checks if this is a minimum cap
func (cc *CapContribution) IsMinCap() bool {
	return cc.Kind == "min"
}

// IsMaxCap checks if this is a maximum cap
func (cc *CapContribution) IsMaxCap() bool {
	return cc.Kind == "max"
}

// GetSortKey returns the sort key for deterministic ordering
func (cc *CapContribution) GetSortKey() string {
	return cc.Dimension + ":" + string(cc.Mode) + ":" + cc.Kind
}
