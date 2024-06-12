package errors

import "lib/constants"

type UnknownError struct {
	Code    int    `json:"code"`
	Error   error  `json:"error"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

func NewUnknownError(message string, error error) *UnknownError {
	return &UnknownError{
		Code:    constants.UnknownError,
		Error:   error,
		Message: message,
		Name:    "UnknownError",
	}
}
