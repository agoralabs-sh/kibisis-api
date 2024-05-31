package types

import "types"

type Request struct {
	Http types.Http `json:"http,omitempty"`
	Name string     `json:"name"`
}
