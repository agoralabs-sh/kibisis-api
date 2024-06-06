package main

import (
  internaltypes "quests/internal/types"
)

func Main(request internaltypes.Request) string {

	return "Hello " + request.Account
}
