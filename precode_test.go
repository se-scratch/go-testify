package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	actualCount := len(strings.Split(responseRecorder.Body.String(), ","))
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, totalCount, actualCount)
}

func TestMainHandlerWhenCityNotFound(t *testing.T) {
	messageError := "wrong city value"
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=1&city=omsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	actualMessage := responseRecorder.Body.String()
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, messageError, actualMessage)
}

func TestMainHandlerWhenRequestCorrect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=1&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotNil(t, responseRecorder.Body)
}
