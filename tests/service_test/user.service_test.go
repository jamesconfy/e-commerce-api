package service_test

import (
	"e-commerce/internal/forms"
	"testing"
)

func TestAddUser(t *testing.T) {
	// Create a new user object
	user := generateUserForm()

	tests := []struct {
		name    string
		user    *forms.Signup
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", user: user, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userSrv.Add(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	// Create a new user object
	userForm := generateUserForm()
	user := createAndRegisterUser(userForm)

	tests := []struct {
		name    string
		user    *forms.Login
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", user: &forms.Login{Email: user.Email, Password: userForm.Password}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userSrv.Login(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	// Create a new user object
	userForm := generateUserForm()
	user := createAndRegisterUser(userForm)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: user.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userSrv.GetById(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSrv.GetById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
