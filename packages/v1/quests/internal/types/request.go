package types

import "lib/types"

type Request struct {
	Account string     `json:"account"`
	Http    types.Http `json:"http,omitempty"`
}
