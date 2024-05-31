package types

type Response struct {
	Body       ResponseBody `json:"body,omitempty"`
	StatusCode int          `json:"statusCode,omitempty"`
}
