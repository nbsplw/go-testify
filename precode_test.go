package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var correctCity = "moscow"
var incorrectCity = "unknown"

func TestMainHandlerWhenStatusOkAndBodyNotEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=1&city=%s", correctCity), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=1&city=%s", incorrectCity), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	testCount := len(cafeList[correctCity]) + 1

	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=%s", testCount, correctCity), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	cities := strings.Split(responseRecorder.Body.String(), ",")

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, cafeList[correctCity], cities)
}
