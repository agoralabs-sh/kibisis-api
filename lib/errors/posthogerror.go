package errors

import "lib/constants"

type PostHogError struct {
	Code    int
	Error   error
	Message string
	Name    string
}

func NewPostHogError(message string, error error) *PostHogError {
	return &PostHogError{
		Code:    constants.PostHogError,
		Error:   error,
		Message: message,
		Name:    "PostHogError",
	}
}