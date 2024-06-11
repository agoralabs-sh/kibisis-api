package main

import (
	_types "docs/internal/types"
	_ "embed"
	"net/http"
)

//go:embed swagger.json
var swaggerJSON string

// @title Kibisis API
// @version 1.0
// @description The Kibisis API provides analytics queries and token data.

// @contact.name   Agora Labs Support
// @contact.email  support@agoralabs.sh

// @license.name  GNU AGPLv3+
// @license.url   https://github.com/agoralabs-sh/kibisis-api/blob/main/COPYING
func Main() *_types.Response {
	return &_types.Response{
		Body: swaggerJSON,
		Headers: _types.ResponseHeaders{
			ContentType: "application/json",
		},
		StatusCode: http.StatusOK,
	}
}
