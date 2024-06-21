package types

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type PairingCreateRequestBody struct {
	Description string `json:"description,omitempty"`
	Host        string `json:"host"`
	Icon        string `json:"icon,omitempty"`
	Name        string `json:"name"`
}

func (requestBody PairingCreateRequestBody) Validate() interface{} {
	return validation.ValidateStruct(&requestBody,
		validation.Field(&requestBody.Description, validation.Length(0, 255)),
		validation.Field(&requestBody.Host, validation.Required, is.Host),
		validation.Field(&requestBody.Icon, validation.Required, is.DataURI),
		validation.Field(&requestBody.Name, validation.Required),
	)
}
