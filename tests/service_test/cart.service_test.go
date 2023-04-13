package service_test

import (
	"e-commerce/internal/models"
	"testing"
)

func TestClearCart(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)

	tests := []struct {
		name    string
		user    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct user", user: user, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cartSrv.Clear(tt.user.Id); (err != nil) != tt.wantErr {
				t.Errorf("cartSrv.Clear() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCart(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)

	tests := []struct {
		name    string
		user    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct user", user: user, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := cartSrv.Get(tt.user.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("cartSrv.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
