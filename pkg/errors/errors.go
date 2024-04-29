package errors

import (
	"fmt"

	"errors"
)

//Reference article:
//https://hackernoon.com/golang-handling-errors-gracefully-8e27f1db729f

// CustomError struct
type CustomError struct {
	errType      ErrorType
	wrappedError error
	context      errorContext
}

// Error return error message
func (err CustomError) GetErrorType() ErrorType {
	return err.errType
}

// Error return error message
func (err CustomError) Error() string {
	return err.wrappedError.Error()
}

// Stacktrace return error stacktrace message
func (err CustomError) Stacktrace() string {
	return fmt.Sprintf("%+v\n", err.wrappedError)
}

// New creates a no type error
func New(ErrType ErrorType, msg string) error {
	return CustomError{errType: ErrType, wrappedError: errors.New(msg)}
}

// Newf creates a no type error with formatted message
func Newf(ErrType ErrorType, msg string, args ...interface{}) error {
	return CustomError{errType: ErrType, wrappedError: fmt.Errorf(msg, args...)}
}
