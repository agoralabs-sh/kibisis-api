package types

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

type PairingAddRequestBody struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (requestBody PairingAddRequestBody) Validate() interface{} {
	return validation.ValidateStruct(&requestBody,
		validation.Field(&requestBody.ID, validation.Required),
		validation.Field(&requestBody.Name, validation.Required),
	)
}
