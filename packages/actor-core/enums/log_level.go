package enums

import "chaos-actor-module/packages/actor-core/constants"

// LogLevel represents the log level
type LogLevel string

const (
	// LogLevelDebug represents debug level
	LogLevelDebug LogLevel = constants.LogLevelDebug

	// LogLevelInfo represents info level
	LogLevelInfo LogLevel = constants.LogLevelInfo

	// LogLevelWarn represents warn level
	LogLevelWarn LogLevel = constants.LogLevelWarn

	// LogLevelError represents error level
	LogLevelError LogLevel = constants.LogLevelError

	// LogLevelFatal represents fatal level
	LogLevelFatal LogLevel = constants.LogLevelFatal
)

// IsValid checks if the log level is valid
func (ll LogLevel) IsValid() bool {
	switch ll {
	case LogLevelDebug, LogLevelInfo, LogLevelWarn, LogLevelError, LogLevelFatal:
		return true
	default:
		return false
	}
}

// String returns the string representation of the log level
func (ll LogLevel) String() string {
	return string(ll)
}

// IsDebug checks if the log level is debug
func (ll LogLevel) IsDebug() bool {
	return ll == LogLevelDebug
}

// IsInfo checks if the log level is info
func (ll LogLevel) IsInfo() bool {
	return ll == LogLevelInfo
}

// IsWarn checks if the log level is warn
func (ll LogLevel) IsWarn() bool {
	return ll == LogLevelWarn
}

// IsError checks if the log level is error
func (ll LogLevel) IsError() bool {
	return ll == LogLevelError
}

// IsFatal checks if the log level is fatal
func (ll LogLevel) IsFatal() bool {
	return ll == LogLevelFatal
}

// GetLevel returns the numeric level for the log level
func (ll LogLevel) GetLevel() int64 {
	switch ll {
	case LogLevelDebug:
		return 1
	case LogLevelInfo:
		return 2
	case LogLevelWarn:
		return 3
	case LogLevelError:
		return 4
	case LogLevelFatal:
		return 5
	default:
		return 0
	}
}

// IsHigherThan checks if this log level is higher than another
func (ll LogLevel) IsHigherThan(other LogLevel) bool {
	return ll.GetLevel() > other.GetLevel()
}

// IsLowerThan checks if this log level is lower than another
func (ll LogLevel) IsLowerThan(other LogLevel) bool {
	return ll.GetLevel() < other.GetLevel()
}

// IsEqual checks if this log level is equal to another
func (ll LogLevel) IsEqual(other LogLevel) bool {
	return ll.GetLevel() == other.GetLevel()
}

// ShouldLog checks if a message with this level should be logged given the current level
func (ll LogLevel) ShouldLog(currentLevel LogLevel) bool {
	return ll.GetLevel() >= currentLevel.GetLevel()
}

// GetDefaultLevel returns the default log level
func GetDefaultLevel() LogLevel {
	return LogLevelInfo
}

// GetProductionLevel returns the production log level
func GetProductionLevel() LogLevel {
	return LogLevelWarn
}

// GetDevelopmentLevel returns the development log level
func GetDevelopmentLevel() LogLevel {
	return LogLevelDebug
}
