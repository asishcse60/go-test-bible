// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	ww := httptest.NewRecorder()
	handler := http.HandlerFunc(HomePage)
	handler.ServeHTTP(ww, req)
	assert.Equal(t, http.StatusOK, ww.Code)
}
