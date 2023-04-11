package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	// Generate user and login him in
	user := generateLoginForm(nil)
	authToken := loginUserAndGenerateAuth(user)

	// Generate product
	product := generateProductForm()

	obj, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", "/test/products/", bytes.NewReader(obj))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", authToken)

	r.ServeHTTP(w, req)

	resp, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
	assert.Contains(t, string(resp), "message", "Response should contain a message")
}

func TestGetProduct(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	// Generate user and login him in
	userForm := generateUserForm()
	user := createAndRegisterUser(userForm)
	userLogin := generateLoginForm(userForm)
	authToken := loginUserAndGenerateAuth(userLogin)

	// Generate product
	product := generateProductForm()
	resultProduct := createAndAddProduct(user, product)

	// Url
	getUrl := fmt.Sprintf("/test/products/%v", resultProduct.Id)

	req, _ := http.NewRequest("GET", getUrl, nil)
	req.Header.Set("Authorization", authToken)

	r.ServeHTTP(w, req)

	resp, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
	assert.Contains(t, string(resp), "message", "Response should contain a message")
}

func TestGetAllProduct(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	// Generate product
	for i := 0; i < 10; i++ {
		_ = createAndAddProduct(nil, nil)
	}

	req, _ := http.NewRequest("GET", "/test/products/", nil)

	r.ServeHTTP(w, req)

	resp, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
	assert.Contains(t, string(resp), "message", "Response should contain a message")
}
