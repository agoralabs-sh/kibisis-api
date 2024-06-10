package main

import (
	_ "embed"
	"net/http"
	"os"
	"strings"
	_types "versions/internal/types"
)

//go:embed VERSION
var version string

func Main() *_types.Response {
	return &_types.Response{
		Body: _types.ResponseBody{
			APIVersion:  strings.TrimSpace(version),
			Environment: os.Getenv("ENVIRONMENT"),
		},
		StatusCode: http.StatusOK,
	}
}
