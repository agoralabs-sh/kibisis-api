package main

import (
	"lib/enums"
	"lib/errors"
	"lib/types"
	"log"
	"net/http"
	"os"
	internaltypes "quests/internal/types"
	"testing"
)

var account = "VSVZ7LMZMKV33XMABJHBCZVO325WFTHZXJZFLS6PN2XLLZC43GOEX4ZYFQ"

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
	request := internaltypes.Request{
		Account: account,
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

func TestInvalidAccount(t *testing.T) {
	// arrange
	request := internaltypes.Request{
		Account: "not a valid account",
		Http: types.Http{
			Method: http.MethodGet,
		},
	}

	// act
	response := Main(request)

	// assert
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("expected \"statusCode\" to be \"%d\", got %d", http.StatusBadRequest, response.StatusCode)
	}
	if response.Body.Error.(*errors.InvalidAddressError).Code != enums.InvalidAddressError {
		t.Errorf("expected \"body.error.code\" to be \"%d\", got %d", enums.InvalidAddressError, response.Body.Error.(*errors.InvalidAddressError).Code)
	}
}
