package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	require.NoError(t, err, "Request creation failed")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Response body should not be empty")

	expectedBody := strings.Join(cafeList["moscow"][:2], ",")
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Response body does not match expected result")
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=2&city=unknown_city", nil)
	require.NoError(t, err, "Request creation failed")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status code 400")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Expected error message in response body")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	require.NoError(t, err, "Request failed")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200")

	expectedBody := strings.Join(cafeList["moscow"], ",")
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Response body doesnt match expected result")
}