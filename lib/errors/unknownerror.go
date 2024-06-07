package errors

import "lib/constants"

type UnknownError struct {
	Code    int
	Error   error
	Message string
}

func NewUnknownError(message string, error error) *UnknownError {
	return &UnknownError{
		Code:    constants.UnknownError,
		Error:   error,
		Message: message,
	}
}
