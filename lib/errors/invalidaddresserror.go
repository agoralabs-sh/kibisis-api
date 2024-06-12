package errors

import (
	"fmt"
	"lib/constants"
)

type InvalidAddressError struct {
	Address string `json:"address"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

func NewInvalidAddressError(address string) *InvalidAddressError {
	return &InvalidAddressError{
		Address: address,
		Code:    constants.InvalidAddressError,
		Message: fmt.Sprintf("address \"%s\" is not a valid address", address),
		Name:    "InvalidAddressError",
	}
}
