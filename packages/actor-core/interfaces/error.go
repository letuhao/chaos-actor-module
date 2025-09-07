package interfaces

import (
	"time"
)

// ActorCoreError represents an error in the Actor Core system
type ActorCoreError struct {
	// Type is the error category
	Type string `json:"type"`

	// Code is the specific error code
	Code string `json:"code"`

	// Message is the human-readable message
	Message string `json:"message"`

	// System is the originating system
	System string `json:"system"`

	// Dimension is the affected dimension
	Dimension string `json:"dimension"`

	// Layer is the affected layer
	Layer string `json:"layer"`

	// Context contains additional context
	Context map[string]interface{} `json:"context"`

	// Timestamp is when the error occurred
	Timestamp time.Time `json:"timestamp"`

	// StackTrace is the stack trace (debug only)
	StackTrace string `json:"stack_trace,omitempty"`
}

// Error returns the error message
func (ace *ActorCoreError) Error() string {
	return ace.Message
}

// GetType returns the error type
func (ace *ActorCoreError) GetType() string {
	return ace.Type
}

// GetCode returns the error code
func (ace *ActorCoreError) GetCode() string {
	return ace.Code
}

// GetSystem returns the originating system
func (ace *ActorCoreError) GetSystem() string {
	return ace.System
}

// GetDimension returns the affected dimension
func (ace *ActorCoreError) GetDimension() string {
	return ace.Dimension
}

// GetLayer returns the affected layer
func (ace *ActorCoreError) GetLayer() string {
	return ace.Layer
}

// GetContext returns the additional context
func (ace *ActorCoreError) GetContext() map[string]interface{} {
	return ace.Context
}

// GetTimestamp returns the timestamp
func (ace *ActorCoreError) GetTimestamp() time.Time {
	return ace.Timestamp
}

// GetStackTrace returns the stack trace
func (ace *ActorCoreError) GetStackTrace() string {
	return ace.StackTrace
}

// IsValidationError checks if this is a validation error
func (ace *ActorCoreError) IsValidationError() bool {
	return ace.Type == "validation"
}

// IsSystemError checks if this is a system error
func (ace *ActorCoreError) IsSystemError() bool {
	return ace.Type == "system"
}

// IsPerformanceError checks if this is a performance error
func (ace *ActorCoreError) IsPerformanceError() bool {
	return ace.Type == "performance"
}

// ErrorHandler represents an error handler
type ErrorHandler interface {
	// Handle handles an error
	Handle(err *ActorCoreError) error

	// CanHandle checks if this handler can handle the error
	CanHandle(err *ActorCoreError) bool

	// GetPriority returns the handler priority
	GetPriority() int64
}

// ErrorRecovery represents an error recovery mechanism
type ErrorRecovery interface {
	// Recover attempts to recover from an error
	Recover(err *ActorCoreError) (interface{}, error)

	// CanRecover checks if this recovery can handle the error
	CanRecover(err *ActorCoreError) bool

	// GetRecoveryType returns the type of recovery
	GetRecoveryType() string
}

// ErrorLogger represents an error logger
type ErrorLogger interface {
	// Log logs an error
	Log(err *ActorCoreError) error

	// LogWithContext logs an error with additional context
	LogWithContext(err *ActorCoreError, context map[string]interface{}) error

	// SetLevel sets the log level
	SetLevel(level string) error

	// GetLevel returns the current log level
	GetLevel() string
}

// ErrorMetrics represents error metrics
type ErrorMetrics struct {
	// TotalErrors is the total number of errors
	TotalErrors int64

	// ErrorsByType is the count of errors by type
	ErrorsByType map[string]int64

	// ErrorsByCode is the count of errors by code
	ErrorsByCode map[string]int64

	// ErrorsBySystem is the count of errors by system
	ErrorsBySystem map[string]int64

	// LastError is the timestamp of the last error
	LastError time.Time

	// ErrorRate is the error rate per second
	ErrorRate float64
}

// GetErrorRate returns the error rate
func (em *ErrorMetrics) GetErrorRate() float64 {
	return em.ErrorRate
}

// GetTotalErrors returns the total number of errors
func (em *ErrorMetrics) GetTotalErrors() int64 {
	return em.TotalErrors
}

// GetErrorsByType returns the count of errors by type
func (em *ErrorMetrics) GetErrorsByType() map[string]int64 {
	return em.ErrorsByType
}

// GetErrorsByCode returns the count of errors by code
func (em *ErrorMetrics) GetErrorsByCode() map[string]int64 {
	return em.ErrorsByCode
}

// GetErrorsBySystem returns the count of errors by system
func (em *ErrorMetrics) GetErrorsBySystem() map[string]int64 {
	return em.ErrorsBySystem
}

// GetLastError returns the timestamp of the last error
func (em *ErrorMetrics) GetLastError() time.Time {
	return em.LastError
}
