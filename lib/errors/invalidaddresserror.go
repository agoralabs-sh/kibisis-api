package errors

import (
	"fmt"
	"lib/constants"
)

type InvalidAddressError struct {
	Address string
	Code    int
	Message string
}

func NewInvalidAddressError(address string) *InvalidAddressError {
	return &InvalidAddressError{
		Address: address,
		Code:    constants.InvalidAddressError,
		Message: fmt.Sprintf("address \"%s\" is not a valid address", address),
	}
}
