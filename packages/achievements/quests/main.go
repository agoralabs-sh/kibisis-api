package main

import (
	"context"
	"fmt"
	internaltypes "quests/internal/types"
	"types"
)

func Main(ctx context.Context, event internaltypes.Event) *types.Response {
	fmt.Println(fmt.Sprintf("method %s", event.Http.Method))
	fmt.Println(fmt.Sprintf("name %s", event.Name))

	return &types.Response{
		Body: "Hello " + event.Name + event.Http.Method,
	}
}
