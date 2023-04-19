package handler_test

import (
	"fmt"
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

	req, _ := http.NewRequest("GET", "/test/carts", nil)
	req.Header.Set("Authorization", authToken)

	r.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(http.StatusOK, w.Code)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestClearCart(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	// Generate user and login him in
	userForm := generateUserForm()
	_ = createAndRegisterUser(userForm)
	userLogin := generateLoginForm(userForm)
	authToken := loginUserAndGenerateAuth(userLogin)

	req, _ := http.NewRequest("DELETE", "/test/carts", nil)
	req.Header.Set("Authorization", authToken)

	r.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}
