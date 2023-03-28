package serviceerror

import (
	"database/sql"
	"errors"
	"time"
)

type Type int

const (
	Unknown = iota
	ErrConflict
	ErrNotFound
	ErrServer
	ErrBadRequest
)

func (t Type) String() string {
	switch t {
	case ErrConflict:
		return "ErrConflict"
	case ErrNotFound:
		return "ErrNotFound"
	case ErrServer:
		return "ErrServer"
	case ErrBadRequest:
		return "BadRequest"
	default:
		return "Unknown"

	}
}

type ServiceError struct {
	Time        string `json:"time"`
	Description string `json:"description"`
	Error       any    `json:"error"`
	ErrorType   Type   `json:"type"`
}

func (se *ServiceError) Type() Type {
	return se.ErrorType
}

func New(description string, err error, errType Type) *ServiceError {
	return &ServiceError{Time: time.Now().Format(time.RFC3339), Description: description, Error: err, ErrorType: errType}
}

func Internal(err error) *ServiceError {
	return New("Internal server error", err, ErrServer)
}

func Validating(err error) *ServiceError {
	return New("Bad input request", err, ErrBadRequest)
}

func Conflict(description string) *ServiceError {
	return New(description, nil, ErrConflict)
}

func NotFound(description string) *ServiceError {
	return New(description, nil, ErrNotFound)
}
func NotFoundOrInternal(err error, descriptions ...string) *ServiceError {
	description := "not found"
	if len(descriptions) > 0 {
		description = descriptions[0]
	}
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NotFound(description)
	default:
		return Internal(err)
	}
}
