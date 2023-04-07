package repo_test

import (
	"e-commerce/internal/models"
	"testing"
)

func TestCreateAuth(t *testing.T) {
	user := createAndAddUser(nil)
	auth := generateAuth(user)

	tests := []struct {
		name    string
		auth    *models.Auth
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct auth object", auth: auth, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := a.Add(tt.auth)

			if (err != nil) != tt.wantErr {
				t.Errorf("auth.CreateAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAuth(t *testing.T) {
	user := createAndAddUser(nil)
	auth := createAndAddAuth(nil, user)

	tests := []struct {
		name    string
		auth    *models.Auth
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct user id", auth: auth, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := a.Get(auth.UserId)

			if (err != nil) != tt.wantErr {
				t.Errorf("auth.GetAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
