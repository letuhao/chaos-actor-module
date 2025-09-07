package stat

// DataType represents the data type of a stat value
type DataType int

const (
	// DataTypeInt64 represents 64-bit integer values
	DataTypeInt64 DataType = iota

	// DataTypeFloat64 represents 64-bit floating point values
	DataTypeFloat64

	// DataTypeBoolean represents boolean values
	DataTypeBoolean

	// DataTypeString represents string values
	DataTypeString

	// DataTypeMap represents map/dictionary values
	DataTypeMap
)

// String returns the string representation of DataType
func (dt DataType) String() string {
	switch dt {
	case DataTypeInt64:
		return "int64"
	case DataTypeFloat64:
		return "float64"
	case DataTypeBoolean:
		return "bool"
	case DataTypeString:
		return "string"
	case DataTypeMap:
		return "map"
	default:
		return "unknown"
	}
}

// IsValid checks if the DataType is valid
func (dt DataType) IsValid() bool {
	return dt >= DataTypeInt64 && dt <= DataTypeMap
}

// GetDisplayName returns the display name of DataType
func (dt DataType) GetDisplayName() string {
	switch dt {
	case DataTypeInt64:
		return "Integer (64-bit)"
	case DataTypeFloat64:
		return "Float (64-bit)"
	case DataTypeBoolean:
		return "Boolean"
	case DataTypeString:
		return "String"
	case DataTypeMap:
		return "Map/Dictionary"
	default:
		return "Unknown Data Type"
	}
}

// GetDescription returns the description of DataType
func (dt DataType) GetDescription() string {
	switch dt {
	case DataTypeInt64:
		return "64-bit signed integer, suitable for counts, levels, and discrete values"
	case DataTypeFloat64:
		return "64-bit floating point number, suitable for percentages, ratios, and continuous values"
	case DataTypeBoolean:
		return "Boolean value (true/false), suitable for flags and binary states"
	case DataTypeString:
		return "String value, suitable for names, descriptions, and text data"
	case DataTypeMap:
		return "Map/dictionary value, suitable for complex structured data"
	default:
		return "Unknown data type"
	}
}

// GetGoType returns the Go type string for DataType
func (dt DataType) GetGoType() string {
	switch dt {
	case DataTypeInt64:
		return "int64"
	case DataTypeFloat64:
		return "float64"
	case DataTypeBoolean:
		return "bool"
	case DataTypeString:
		return "string"
	case DataTypeMap:
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}

// GetDefaultValue returns the default value for DataType
func (dt DataType) GetDefaultValue() interface{} {
	switch dt {
	case DataTypeInt64:
		return int64(0)
	case DataTypeFloat64:
		return float64(0.0)
	case DataTypeBoolean:
		return false
	case DataTypeString:
		return ""
	case DataTypeMap:
		return make(map[string]interface{})
	default:
		return nil
	}
}

// GetAllDataTypes returns all valid DataType values
func GetAllDataTypes() []DataType {
	return []DataType{
		DataTypeInt64,
		DataTypeFloat64,
		DataTypeBoolean,
		DataTypeString,
		DataTypeMap,
	}
}

// GetDataTypeFromString converts a string to DataType
func GetDataTypeFromString(s string) (DataType, bool) {
	switch s {
	case "int64":
		return DataTypeInt64, true
	case "float64":
		return DataTypeFloat64, true
	case "bool":
		return DataTypeBoolean, true
	case "string":
		return DataTypeString, true
	case "map":
		return DataTypeMap, true
	default:
		return DataTypeInt64, false
	}
}

// IsNumeric checks if the DataType is numeric
func (dt DataType) IsNumeric() bool {
	return dt == DataTypeInt64 || dt == DataTypeFloat64
}

// IsComparable checks if the DataType supports comparison operations
func (dt DataType) IsComparable() bool {
	return dt == DataTypeInt64 || dt == DataTypeFloat64 || dt == DataTypeBoolean || dt == DataTypeString
}

// SupportsArithmetic checks if the DataType supports arithmetic operations
func (dt DataType) SupportsArithmetic() bool {
	return dt == DataTypeInt64 || dt == DataTypeFloat64
}
