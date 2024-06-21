package types

type ErrorField struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}
