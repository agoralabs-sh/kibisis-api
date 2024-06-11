package types

type Response struct {
	Body       string          `json:"body"`
	StatusCode int             `json:"statusCode"`
	Headers    ResponseHeaders `json:"headers"`
}
