package types

import "lib/types"

type Request struct {
	Http types.Http `json:"http,omitempty"`
	Page string     `json:"page,omitempty"`
}
