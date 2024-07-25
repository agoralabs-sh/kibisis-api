package main

import (
	_types "leaderboard/internal/types"
	"lib/types"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := os.Setenv("ENVIRONMENT", "test")
	if err != nil {
		log.Fatalf("failed to set \"ENVIRONMENT\" variable to \"test\"")
	}

	// run tests
	code := m.Run()

	os.Exit(code)
}

func TestIncorrectHTTPMethod(t *testing.T) {
	// arrange
	request := _types.Request{
		Http: types.Http{
			Method: http.MethodPost,
		},
	}

	// act
	response := Main(request)

	// assert
	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected \"statusCode\" to be \"%d\", got %d", http.StatusMethodNotAllowed, response.StatusCode)
	}
}
