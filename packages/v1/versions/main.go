package main

import (
	_ "embed"
	"net/http"
	"os"
	"strings"
	internaltypes "versions/internal/types"
)

//go:embed VERSION
var version string

func Main() *internaltypes.Response {
	return &internaltypes.Response{
		Body: internaltypes.ResponseBody{
			APIVersion:  strings.TrimSpace(version),
			Environment: os.Getenv("ENVIRONMENT"),
		},
		StatusCode: http.StatusOK,
	}
}
