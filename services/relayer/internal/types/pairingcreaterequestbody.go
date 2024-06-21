package types

type PairingCreateRequestBody struct {
	Description string `json:"description,omitempty" validate:"lte=255"`
	Host        string `json:"host" validate:"required,hostname"`
	Icon        string `json:"icon,omitempty" validate:"datauri"`
	Name        string `json:"name" validate:"required"`
}
