package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMainWithSuccessfulResponse(t *testing.T) {
	// arrange
	err := os.Setenv("ENVIRONMENT", "development")
	if err != nil {
		t.Errorf("failed to set \"environment\" variable")
	}

	// act
	response := Main()

	// assert
	if response.Body.Environment != "development" {
		t.Errorf("expected \"environment\" to be \"development\", got %s", response.Body.Environment)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("expected \"statusCode\" to be \"%d\", got %d", http.StatusOK, response.StatusCode)
	}
}
