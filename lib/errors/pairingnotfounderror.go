package errors

import (
	"fmt"
	"lib/enums"
)

type PairingNotFoundError struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Name      string `json:"name"`
	PairingId string `json:"pairingId"`
}

func NewPairingNotFoundError(pairingId string) *PairingNotFoundError {
	return &PairingNotFoundError{
		Code:      enums.PairingNotFoundError,
		Message:   fmt.Sprintf("pairing \"%s\" not found", pairingId),
		Name:      "PairingNotFoundError",
		PairingId: pairingId,
	}
}
