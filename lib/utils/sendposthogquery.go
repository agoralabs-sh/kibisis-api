package utils

import (
	"bytes"
	"fmt"
	"lib/constants"
	"net/http"
	"os"
	"strings"
	"time"
)

func SendPostHogQuery(query string, output any) error {
	path := strings.Replace(constants.QueryPath, ":project_id", os.Getenv("POSTHOG_PROJECT_ID"), -1)
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", os.Getenv("POSTHOG_API_URL"), path), bytes.NewReader([]byte(query)))
	if err != nil {
		return err
	}

	// add headers
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("POSTHOG_API_KEY")))
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 20 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	err = ParseResponseBody(response.Body, &output)
	if err != nil {
		return err
	}

	return nil
}
