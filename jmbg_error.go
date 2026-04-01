package jmbg

import (
	"errors"
	"fmt"
)

// Sentinel errors for each validation failure.
var (
	ErrInvalidLength   = errors.New("jmbg: invalid length")
	ErrInvalidFormat   = errors.New("jmbg: invalid format")
	ErrInvalidDate     = errors.New("jmbg: invalid date")
	ErrInvalidRegion   = errors.New("jmbg: invalid region")
	ErrInvalidChecksum = errors.New("jmbg: invalid checksum")
)

// ValidationError represents a JMBG validation failure.
// Use errors.Is to check the kind of failure and errors.As to access details.
type ValidationError struct {
	// Err is the underlying sentinel error (e.g., ErrInvalidLength).
	Err error
	// Detail provides additional context about the failure.
	Detail string
}

func (e *ValidationError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("%s: %s", e.Err, e.Detail)
	}
	return e.Err.Error()
}

func (e *ValidationError) Unwrap() error {
	return e.Err
}
