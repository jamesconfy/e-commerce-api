package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCart(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	// Generate user and login him in
	userForm := generateUserForm()
	_ = createAndRegisterUser(userForm)
	userLogin := generateLoginForm(userForm)
	authToken := loginUserAndGenerateAuth(userLogin)

	req, _ := http.NewRequest("GET", "/test/carts/", nil)
	req.Header.Set("Authorization", authToken)

	r.ServeHTTP(w, req)

	resp, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
	assert.Contains(t, string(resp), "Cart gotten successfully", "Response should contain a message")
}

func TestClearCart(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	// Generate user and login him in
	userForm := generateUserForm()
	_ = createAndRegisterUser(userForm)
	userLogin := generateLoginForm(userForm)
	authToken := loginUserAndGenerateAuth(userLogin)

	req, _ := http.NewRequest("DELETE", "/test/carts/", nil)
	req.Header.Set("Authorization", authToken)

	r.ServeHTTP(w, req)

	resp, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
	assert.Contains(t, string(resp), "Cart cleared successfully", "Response should contain a message")
}
