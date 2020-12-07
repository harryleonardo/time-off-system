package errors

import (
	Errors "github.com/pkg/errors"
)

type (
	// ErrorWrapper ...
	ErrorWrapper struct {
		Code       string
		Message    string
		StatusCode int
		ErrorCode  string
	}
)

// Error returns error type as a string
func (q *ErrorWrapper) Error() string {
	return q.Message
}

// New returnns new error message in standard pkg errors new
func New(msg string) error {
	return Errors.New(msg)
}

// Wrap ...
func Wrap(err error, message string) error {
	return Errors.Wrap(&ErrorWrapper{
		Code:      "500",
		ErrorCode: "VP-GEN-001",
		Message:   message,
	}, err.Error())
}

// WrapError ...
func WrapError(code, message string, status int, original error) error {
	return Errors.Wrap(&ErrorWrapper{
		Code:       code,
		Message:    message,
		StatusCode: status,
	}, original.Error())
}
