package main

import (
	_ "embed"
	"fmt"
	"lib/constants"
	"net/http"
	"os"
	"strings"
	_types "versions/internal/types"
)

//go:embed VERSION
var version string

// Main godoc.
// @Summary API Information
// @Description Gets API information including the environment and version.
// @Produce json
// @Success 200 {object} _types.ResponseBody
// @Router /v1/versions [get]
func Main() *_types.Response {
	return &_types.Response{
		Body: _types.ResponseBody{
			APIVersion:  strings.TrimSpace(version),
			Environment: os.Getenv("ENVIRONMENT"),
		},
		Headers: _types.ResponseHeaders{
			CacheControl: fmt.Sprintf("public, max-age=%d", constants.TwentyFourHoursInSeconds),
		},
		StatusCode: http.StatusOK,
	}
}
