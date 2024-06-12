package errors

import "lib/constants"

type PostHogError struct {
	Code    int    `json:"code"`
	Error   error  `json:"error"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

func NewPostHogError(message string, error error) *PostHogError {
	return &PostHogError{
		Code:    constants.PostHogError,
		Error:   error,
		Message: message,
		Name:    "PostHogError",
	}
}
