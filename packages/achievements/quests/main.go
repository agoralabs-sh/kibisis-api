package main

import (
	"fmt"
	internaltypes "quests/internal/types"
)

func Main(request internaltypes.Request) string {
	fmt.Println(fmt.Sprintf("method %s", request.Http.Method))
	fmt.Println(fmt.Sprintf("name %s", request.Name))

	return "Hello " + request.Name + request.Http.Method
}
