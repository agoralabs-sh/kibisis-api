package types

import "types"

type Event struct {
	types.BaseEvent
	Name string `json:"name"`
}
