package enums

// Operator represents the aggregation operator
type Operator string

const (
	// OperatorSum represents sum aggregation
	OperatorSum Operator = "SUM"
	
	// OperatorMax represents maximum aggregation
	OperatorMax Operator = "MAX"
	
	// OperatorMin represents minimum aggregation
	OperatorMin Operator = "MIN"
	
	// OperatorAverage represents average aggregation
	OperatorAverage Operator = "AVERAGE"
	
	// OperatorMultiply represents multiply aggregation
	OperatorMultiply Operator = "MULTIPLY"
	
	// OperatorIntersect represents intersection aggregation
	OperatorIntersect Operator = "INTERSECT"
)

// IsValid checks if the operator is valid
func (o Operator) IsValid() bool {
	switch o {
	case OperatorSum, OperatorMax, OperatorMin, OperatorAverage, OperatorMultiply, OperatorIntersect:
		return true
	default:
		return false
	}
}

// String returns the string representation of the operator
func (o Operator) String() string {
	return string(o)
}

// IsSum checks if the operator is sum
func (o Operator) IsSum() bool {
	return o == OperatorSum
}

// IsMax checks if the operator is max
func (o Operator) IsMax() bool {
	return o == OperatorMax
}

// IsMin checks if the operator is min
func (o Operator) IsMin() bool {
	return o == OperatorMin
}

// IsAverage checks if the operator is average
func (o Operator) IsAverage() bool {
	return o == OperatorAverage
}

// IsMultiply checks if the operator is multiply
func (o Operator) IsMultiply() bool {
	return o == OperatorMultiply
}

// IsIntersect checks if the operator is intersect
func (o Operator) IsIntersect() bool {
	return o == OperatorIntersect
}

// IsCommutative checks if the operator is commutative
func (o Operator) IsCommutative() bool {
	switch o {
	case OperatorSum, OperatorMax, OperatorMin, OperatorMultiply:
		return true
	case OperatorAverage, OperatorIntersect:
		return false
	default:
		return false
	}
}

// IsAssociative checks if the operator is associative
func (o Operator) IsAssociative() bool {
	switch o {
	case OperatorSum, OperatorMax, OperatorMin, OperatorMultiply:
		return true
	case OperatorAverage, OperatorIntersect:
		return false
	default:
		return false
	}
}

// GetDefaultValue returns the default value for the operator
func (o Operator) GetDefaultValue() float64 {
	switch o {
	case OperatorSum, OperatorAverage:
		return 0.0
	case OperatorMax:
		return -999999999.0 // Very small number
	case OperatorMin:
		return 999999999.0 // Very large number
	case OperatorMultiply:
		return 1.0
	case OperatorIntersect:
		return 0.0
	default:
		return 0.0
	}
}

// Apply applies the operator to two values
func (o Operator) Apply(a, b float64) float64 {
	switch o {
	case OperatorSum:
		return a + b
	case OperatorMax:
		if a > b {
			return a
		}
		return b
	case OperatorMin:
		if a < b {
			return a
		}
		return b
	case OperatorAverage:
		return (a + b) / 2.0
	case OperatorMultiply:
		return a * b
	case OperatorIntersect:
		// For intersection, we need to handle ranges
		// This is a simplified version - actual implementation would be more complex
		if a < b {
			return a
		}
		return b
	default:
		return a
	}
}
