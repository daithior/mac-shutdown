package main

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestShutdownHandler(test *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		test.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handleShutdown)
	handler.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != http.StatusOK {
		test.Fatal("Http response was not 200")
	}

	expectedResponse := "shutdown complete"
	matchCondition := regexp.MustCompile(`\b` + expectedResponse + `\b`)
	responseText := responseRecorder.Body.String()
	if !matchCondition.MatchString(responseText) {
		test.Fatal("Expected response body did not match")
	}

}

func TestExecutShellCommand(test *testing.T) {
	result := executeShellCommand()
	expectedResult := "Darwin"
	matchCondition := regexp.MustCompile(`\b` + expectedResult + `\b`)

	if !matchCondition.MatchString(result) {
		test.Fatal("Uname shell command did not match expected darwin")
	}

}
