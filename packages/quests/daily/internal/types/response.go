package types

type Response struct {
	Headers    ResponseHeaders `json:"headers"`
	StatusCode int             `json:"statusCode,omitempty"`
}
