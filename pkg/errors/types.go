package errors

import (
	"errors"
	"fmt"
)

// constants define error types
const (
	Success                    ErrorType = 200
	Created                    ErrorType = 201
	InvalidParams              ErrorType = 400
	ErrorNoPermission          ErrorType = 403
	ErrorNotFound              ErrorType = 404
	ErrorMethodNotAllow        ErrorType = 405
	ErrorInvalidParent         ErrorType = 409
	ErrorAllowDeleteWithChild  ErrorType = 410
	ErrorNotAllowDelete        ErrorType = 411
	ErrorUserDisabled          ErrorType = 412
	ErrorExistMenuName         ErrorType = 413
	ErrorExistRole             ErrorType = 414
	ErrorExistRoleUser         ErrorType = 415
	ErrorNotExistUser          ErrorType = 416
	ErrorBadRequest            ErrorType = 421
	ErrorLoginFailed           ErrorType = 422
	ErrorInvalidOldPass        ErrorType = 423
	ErrorPasswordRequired      ErrorType = 424
	ErrorTooManyRequest        ErrorType = 429
	ErrorAuthCheckTokenFail    ErrorType = 401
	ErrorAuthCheckTokenTimeout ErrorType = 402
	ErrorAuthToken             ErrorType = 408
	ErrorAuth                  ErrorType = 407
	ErrorExistEmail            ErrorType = 430
	ErrorNotExistRole          ErrorType = 431
	ErrorTokenExpired          ErrorType = 461
	ErrorTokenInvalid          ErrorType = 462
	ErrorTokenMalformed        ErrorType = 463
	ErrorInternalServer        ErrorType = 500

	// System errors
	ErrorMarshal ErrorType = iota + 1000
	ErrorUnmarshal
	ErrorPayload
	ErrorDatabaseGet
	ErrorDatabaseCreate
	ErrorDatabaseUpdate
	ErrorDatabaseDelete
	ErrorInvalidPassword
)

// ErrorType type
type ErrorType int

// New error
func (errType ErrorType) New() error {
	return CustomError{
		errType:      errType,
		wrappedError: fmt.Errorf(GetMsg(int(errType))),
	}
}

// Newm new error with message
func (errType ErrorType) Newm(msg string) error {
	return CustomError{
		errType:      errType,
		wrappedError: errors.New(msg),
	}
}

// Newf creates a new CustomError with formatted message
func (errType ErrorType) Newf(msg string, args ...interface{}) error {
	err := fmt.Errorf(msg, args...)

	return CustomError{
		errType:      errType,
		wrappedError: err,
	}
}

// Define some template errors
var (
	ErrMethodNotAllow = ErrorMethodNotAllow.New()
	ErrNoPermission   = ErrorNoPermission.New()
	ErrNotFound       = ErrorNotFound.New()
	ErrTokenExpired   = ErrorTokenExpired.New()
	ErrTokenInvalid   = ErrorTokenInvalid.New()
	ErrTokenMalformed = ErrorTokenMalformed.New()
)
