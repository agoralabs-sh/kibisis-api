package utils

import (
	"encoding/json"
	"io"
)

func ParseResponseBody(responseBody io.Reader, output any) error {
	body, err := io.ReadAll(responseBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &output)
}
