package types

type Response struct {
	Body       ResponseBody    `json:"body,omitempty"`
	Headers    ResponseHeaders `json:"headers,omitempty"`
	StatusCode int             `json:"statusCode,omitempty"`
}
