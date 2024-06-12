package errors

import (
	"fmt"
	"lib/constants"
)

type RequiredParamsError struct {
	Code    int
	Message string
	Name    string
	Params  []string
}

func NewRequiredParamsError(params []string) *RequiredParamsError {
	return &RequiredParamsError{
		Code:    constants.RequiredParamsError,
		Message: fmt.Sprintf("the params \"%s\" are required", params),
		Name:    "RequiredParamsError",
		Params:  params,
	}
}
