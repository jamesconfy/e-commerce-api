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

func TestAddItem(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	// Generate user and login him in
	userForm := generateUserForm()
	_ = createAndRegisterUser(userForm)
	userLogin := generateLoginForm(userForm)
	authToken := loginUserAndGenerateAuth(userLogin)

	// Create item
	item := generateCartItem(nil)

	// Marshal item form into json
	obj, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", "/test/items/", bytes.NewBuffer(obj))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	r.ServeHTTP(w, req)

	resp, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
	assert.Contains(t, string(resp), "Item added successfully", "Response should contain a message")
}

func TestGetItems(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	// Generate user and login him in
	userForm := generateUserForm()
	user := createAndRegisterUser(userForm)
	authToken := loginUserAndGenerateAuth(generateLoginForm(userForm))

	// Create items
	for i := 0; i < 10; i++ {
		_ = createAndAddItem(user, nil)
	}

	req, _ := http.NewRequest("GET", "/test/items/", nil)
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	r.ServeHTTP(w, req)

	resp, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
	assert.Contains(t, string(resp), "Items gotten successfully", "Response should contain a message")
}

func TestGetItem(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	// Generate user and login him in
	userForm := generateUserForm()
	user := createAndRegisterUser(userForm)
	authToken := loginUserAndGenerateAuth(generateLoginForm(userForm))

	// Create product
	product := createAndAddProduct(nil, nil)

	// Create item
	_ = createAndAddItem(user, product)

	// Url
	url := fmt.Sprintf("/test/items/%v", product.Id)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", authToken)

	r.ServeHTTP(w, req)

	resp, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
	assert.Contains(t, string(resp), "Item gotten successfully", "Response should contain a message")
}

func TestDeleteItem(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	// Generate user and login him in
	userForm := generateUserForm()
	user := createAndRegisterUser(userForm)
	authToken := loginUserAndGenerateAuth(generateLoginForm(userForm))

	// Create product
	product := createAndAddProduct(nil, nil)

	// Create item
	_ = createAndAddItem(user, product)

	// Url
	url := fmt.Sprintf("/test/items/%v", product.Id)

	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", authToken)

	r.ServeHTTP(w, req)

	resp, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
	assert.Contains(t, string(resp), "Item deleted successfully", "Response should contain a message")
}
