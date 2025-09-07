package enums

import "chaos-actor-module/packages/actor-core/constants"

// Bucket represents the type of contribution bucket
type Bucket string

const (
	// BucketFlat represents flat additive contributions
	BucketFlat Bucket = "FLAT"
	
	// BucketMult represents multiplicative contributions
	BucketMult Bucket = "MULT"
	
	// BucketPostAdd represents post-addition contributions
	BucketPostAdd Bucket = "POST_ADD"
	
	// BucketOverride represents override contributions
	BucketOverride Bucket = "OVERRIDE"
	
	// BucketExponential represents exponential contributions (value^exponent)
	BucketExponential Bucket = "EXPONENTIAL"
	
	// BucketLogarithmic represents logarithmic contributions (log(value))
	BucketLogarithmic Bucket = "LOGARITHMIC"
	
	// BucketConditional represents conditional contributions
	BucketConditional Bucket = "CONDITIONAL"
)

// IsValid checks if the bucket type is valid
func (b Bucket) IsValid() bool {
	switch b {
	case BucketFlat, BucketMult, BucketPostAdd, BucketOverride,
		 BucketExponential, BucketLogarithmic, BucketConditional:
		return true
	default:
		return false
	}
}

// String returns the string representation of the bucket
func (b Bucket) String() string {
	return string(b)
}

// GetDefaultClampRange returns the default clamp range for a bucket type
func (b Bucket) GetDefaultClampRange() (min, max float64) {
	switch b {
	case BucketFlat, BucketPostAdd:
		return constants.MinAttackPower, constants.MaxAttackPower
	case BucketMult:
		return 0.0, 10.0 // Multiplier range
	case BucketOverride:
		return constants.MinAttackPower, constants.MaxAttackPower
	case BucketExponential:
		return 0.0, 100.0 // Exponential range
	case BucketLogarithmic:
		return 0.0, 10.0 // Logarithmic range
	case BucketConditional:
		return 0.0, 1.0 // Conditional range (0-1 for boolean-like)
	default:
		return 0.0, 1.0
	}
}

// RequiresBaseValue checks if the bucket type requires a base value for calculation
func (b Bucket) RequiresBaseValue() bool {
	switch b {
	case BucketFlat, BucketMult, BucketPostAdd, BucketExponential, BucketLogarithmic:
		return true
	case BucketOverride, BucketConditional:
		return false
	default:
		return true
	}
}

// IsMultiplicative checks if the bucket type is multiplicative
func (b Bucket) IsMultiplicative() bool {
	return b == BucketMult || b == BucketExponential || b == BucketLogarithmic
}

// IsAdditive checks if the bucket type is additive
func (b Bucket) IsAdditive() bool {
	return b == BucketFlat || b == BucketPostAdd
}

// IsOverride checks if the bucket type overrides the base value
func (b Bucket) IsOverride() bool {
	return b == BucketOverride
}

// IsConditional checks if the bucket type is conditional
func (b Bucket) IsConditional() bool {
	return b == BucketConditional
}
