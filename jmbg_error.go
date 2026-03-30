package main

import "fmt"

// JmbgError represents a validation error.
type JmbgError struct {
	Message string
}

func (e *JmbgError) Error() string {
	return e.Message
}

func newError(format string, args ...any) *JmbgError {
	return &JmbgError{Message: fmt.Sprintf(format, args...)}
}
