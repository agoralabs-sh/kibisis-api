package main

import (
	"fmt"
	_types "quests/internal/types"
)

func Main(request _types.Request) *_types.Response {
	return &_types.Response{
		StatusCode: 302,
		Headers: _types.ResponseHeaders{
			Location: fmt.Sprintf("/v1/quests/daily?account=%s&", request.Account),
		},
	}
}
