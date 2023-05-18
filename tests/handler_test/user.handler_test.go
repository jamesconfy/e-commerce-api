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

func TestCreateUser(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	user := generateUserForm()

	obj, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", "/test/auth/signup", bytes.NewReader(obj))
	req.Header.Set("Content-type", "application/json")

	r.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestLoginUser(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	user := generateUserForm()
	_ = createAndRegisterUser(user)
	user1 := generateLoginForm(user)

	obj, err := json.Marshal(user1)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", "/test/auth/login", bytes.NewReader(obj))
	req.Header.Set("Content-type", "application/json")

	r.ServeHTTP(w, req)

	_, err = io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestGetUser(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	user := createAndRegisterUser(nil)

	getUrl := fmt.Sprintf("/test/users/%v", user.Id)

	req, _ := http.NewRequest("GET", getUrl, nil)

	r.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}

func TestGetAllUsers(t *testing.T) {
	r := setupApp()
	w := httptest.NewRecorder()

	for i := 0; i < 10; i++ {
		_ = createAndRegisterUser(nil)
	}

	req, _ := http.NewRequest("GET", "/test/users", nil)

	r.ServeHTTP(w, req)

	_, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be the same")
}
