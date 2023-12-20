package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCorrectDataAreCorrect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenCityDataAreWrong(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?count=5&city=Samara", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountAnother(t *testing.T) {
	totalCount := 4
	city := "moscow"
	count := "5"

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	q := req.URL.Query()
	q.Add("count", count)
	q.Add("city", city)
	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(mainHandle)
	handlerFunc.ServeHTTP(responseRecorder, req)

	res := responseRecorder.Body.String()
	arr := strings.Split(res, ",")
	assert.Equal(t, totalCount, len(arr))
}
