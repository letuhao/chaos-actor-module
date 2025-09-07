package core

import "errors"

// Error definitions
var (
	ErrStatNotFound = errors.New("stat not found")
	ErrInvalidValue = errors.New("invalid value")
	ErrOutOfBounds  = errors.New("value out of bounds")
)

// ValidationError represents a validation error
type ValidationError struct {
	Field     string `json:"field"`
	Message   string `json:"message"`
	Severity  string `json:"severity"`
	Timestamp int64  `json:"timestamp"`
}
