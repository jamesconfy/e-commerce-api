package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHome(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/", nil)

	r.ServeHTTP(w, req)

	resp, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(resp), `{"msg":"Home"}`)
}
