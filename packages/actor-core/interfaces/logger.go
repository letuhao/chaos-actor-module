package interfaces

import (
	"context"
	"time"
)

// Logger represents a logger interface
type Logger interface {
	// Debug logs a debug message
	Debug(msg string, fields ...Field)
	
	// Info logs an info message
	Info(msg string, fields ...Field)
	
	// Warn logs a warning message
	Warn(msg string, fields ...Field)
	
	// Error logs an error message
	Error(msg string, fields ...Field)
	
	// Fatal logs a fatal message
	Fatal(msg string, fields ...Field)
	
	// WithContext returns a logger with context
	WithContext(ctx context.Context) Logger
	
	// WithFields returns a logger with fields
	WithFields(fields ...Field) Logger
	
	// SetLevel sets the log level
	SetLevel(level string) error
	
	// GetLevel returns the current log level
	GetLevel() string
}

// Field represents a log field
type Field struct {
	Key   string
	Value interface{}
}

// NewField creates a new log field
func NewField(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

// StringField creates a string field
func StringField(key, value string) Field {
	return Field{Key: key, Value: value}
}

// IntField creates an int field
func IntField(key string, value int) Field {
	return Field{Key: key, Value: value}
}

// Int64Field creates an int64 field
func Int64Field(key string, value int64) Field {
	return Field{Key: key, Value: value}
}

// Float64Field creates a float64 field
func Float64Field(key string, value float64) Field {
	return Field{Key: key, Value: value}
}

// BoolField creates a bool field
func BoolField(key string, value bool) Field {
	return Field{Key: key, Value: value}
}

// TimeField creates a time field
func TimeField(key string, value time.Time) Field {
	return Field{Key: key, Value: value}
}

// DurationField creates a duration field
func DurationField(key string, value time.Duration) Field {
	return Field{Key: key, Value: value}
}

// ErrorField creates an error field
func ErrorField(key string, value error) Field {
	return Field{Key: key, Value: value}
}

// LoggerConfig represents logger configuration
type LoggerConfig struct {
	// Level is the log level
	Level string `json:"level"`
	
	// Format is the log format (json, text)
	Format string `json:"format"`
	
	// Output is the output destination (stdout, stderr, file)
	Output string `json:"output"`
	
	// FilePath is the file path for file output
	FilePath string `json:"file_path,omitempty"`
	
	// MaxSize is the maximum file size in MB
	MaxSize int `json:"max_size,omitempty"`
	
	// MaxBackups is the maximum number of backup files
	MaxBackups int `json:"max_backups,omitempty"`
	
	// MaxAge is the maximum age of backup files in days
	MaxAge int `json:"max_age,omitempty"`
	
	// Compress indicates whether to compress backup files
	Compress bool `json:"compress,omitempty"`
	
	// Fields are default fields to include in all logs
	Fields map[string]interface{} `json:"fields,omitempty"`
}

// IsValid checks if the logger config is valid
func (lc *LoggerConfig) IsValid() bool {
	if lc.Level == "" || lc.Format == "" || lc.Output == "" {
		return false
	}
	
	if lc.Output == "file" && lc.FilePath == "" {
		return false
	}
	
	return true
}

// GetDefaultConfig returns the default logger config
func GetDefaultConfig() *LoggerConfig {
	return &LoggerConfig{
		Level:   "info",
		Format:  "json",
		Output:  "stdout",
		MaxSize: 100,
		MaxBackups: 3,
		MaxAge: 28,
		Compress: true,
		Fields: map[string]interface{}{
			"service": "actor-core",
			"version": "v3.0.0",
		},
	}
}

// LogEntry represents a log entry
type LogEntry struct {
	// Level is the log level
	Level string `json:"level"`
	
	// Message is the log message
	Message string `json:"message"`
	
	// Timestamp is the log timestamp
	Timestamp time.Time `json:"timestamp"`
	
	// Fields are additional fields
	Fields map[string]interface{} `json:"fields,omitempty"`
	
	// Error is the error (if any)
	Error error `json:"error,omitempty"`
}

// GetLevel returns the log level
func (le *LogEntry) GetLevel() string {
	return le.Level
}

// GetMessage returns the log message
func (le *LogEntry) GetMessage() string {
	return le.Message
}

// GetTimestamp returns the timestamp
func (le *LogEntry) GetTimestamp() time.Time {
	return le.Timestamp
}

// GetFields returns the fields
func (le *LogEntry) GetFields() map[string]interface{} {
	return le.Fields
}

// GetError returns the error
func (le *LogEntry) GetError() error {
	return le.Error
}

// HasError checks if the entry has an error
func (le *LogEntry) HasError() bool {
	return le.Error != nil
}
