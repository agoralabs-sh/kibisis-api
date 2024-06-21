package errors

import (
	"fmt"
	"lib/enums"
)

type RequiredParamsError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Name    string   `json:"name"`
	Params  []string `json:"params"`
}

func NewRequiredParamsError(params []string) *RequiredParamsError {
	return &RequiredParamsError{
		Code:    enums.RequiredParamsError,
		Message: fmt.Sprintf("the params \"%s\" are required", params),
		Name:    "RequiredParamsError",
		Params:  params,
	}
}
