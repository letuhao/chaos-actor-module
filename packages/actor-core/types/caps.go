package types

// Caps represents min/max caps for a dimension
type Caps struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

// IsValid checks if the caps are valid
func (c Caps) IsValid() bool {
	return c.Min <= c.Max
}

// Contains checks if the value is within the caps
func (c Caps) Contains(value float64) bool {
	return value >= c.Min && value <= c.Max
}

// Clamp clamps the value to the caps range
func (c Caps) Clamp(value float64) float64 {
	if value < c.Min {
		return c.Min
	}
	if value > c.Max {
		return c.Max
	}
	return value
}

// Intersect intersects this caps with another
func (c Caps) Intersect(other Caps) Caps {
	min := c.Min
	if other.Min > min {
		min = other.Min
	}

	max := c.Max
	if other.Max < max {
		max = other.Max
	}

	return Caps{Min: min, Max: max}
}

// Union unions this caps with another
func (c Caps) Union(other Caps) Caps {
	min := c.Min
	if other.Min < min {
		min = other.Min
	}

	max := c.Max
	if other.Max > max {
		max = other.Max
	}

	return Caps{Min: min, Max: max}
}

// GetMin returns the minimum value
func (c Caps) GetMin() float64 {
	return c.Min
}

// GetMax returns the maximum value
func (c Caps) GetMax() float64 {
	return c.Max
}

// SetMin sets the minimum value
func (c *Caps) SetMin(min float64) {
	c.Min = min
}

// SetMax sets the maximum value
func (c *Caps) SetMax(max float64) {
	c.Max = max
}

// Set sets both min and max values
func (c *Caps) Set(min, max float64) {
	c.Min = min
	c.Max = max
}

// IsEmpty checks if the caps are empty (min == max == 0)
func (c Caps) IsEmpty() bool {
	return c.Min == 0 && c.Max == 0
}

// GetRange returns the range (max - min)
func (c Caps) GetRange() float64 {
	return c.Max - c.Min
}

// GetCenter returns the center point
func (c Caps) GetCenter() float64 {
	return (c.Min + c.Max) / 2.0
}

// Expand expands the caps by the given amount
func (c *Caps) Expand(amount float64) {
	c.Min -= amount
	c.Max += amount
}

// Shrink shrinks the caps by the given amount
func (c *Caps) Shrink(amount float64) {
	c.Min += amount
	c.Max -= amount

	// Ensure min doesn't exceed max
	if c.Min > c.Max {
		center := c.GetCenter()
		c.Min = center
		c.Max = center
	}
}

// Clone creates a copy of the caps
func (c Caps) Clone() Caps {
	return Caps{Min: c.Min, Max: c.Max}
}
