package types

type Http struct {
	Headers Headers `json:"headers"`
	Method  string  `json:"method"`
	Path    string  `json:"path"`
}
