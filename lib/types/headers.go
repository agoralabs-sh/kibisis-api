package types

type Headers struct {
	Accept          string `json:"accept"`
	AcceptEncoding  string `json:"accept-encoding"`
	UserAgent       string `json:"user-agent"`
	XForwardedFor   string `json:"x-forwarded-for"`
	XForwardedProto string `json:"x-forwarded-proto"`
	XRequestId      string `json:"x-request-id"`
}
