package errors

import (
	"fmt"
	"lib/enums"
)

type ValidationError struct {
	Code    int         `json:"code"`
	Fields  interface{} `json:"fields"`
	Message string      `json:"message"`
	Name    string      `json:"name"`
	Type    string      `json:"type"`
}

func NewValidationError(fieldType string, fields interface{}) *ValidationError {
	return &ValidationError{
		Code:    enums.ValidationError,
		Fields:  fields,
		Message: fmt.Sprintf("failed validation of the \"%s\" fields", fieldType),
		Name:    "ValidationError",
		Type:    fieldType,
	}
}
