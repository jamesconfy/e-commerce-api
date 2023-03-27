package repo_test

import (
	"e-commerce/internal/models"
	"testing"

	"github.com/google/uuid"
)

func TestAddUser(t *testing.T) {
	// Create a new user object
	user1 := &models.User{
		Id:          uuid.New().String(),
		FirstName:   "Confidence",
		LastName:    "James",
		Email:       "bobdence@gmail11.com",
		PhoneNumber: "0814979537011",
		Password:    "123456",
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
	}

	// Create a new cart object
	cart1 := &models.Cart{
		Id:          uuid.New().String(),
		DateCreated: ti.CurrentTime(),
	}

	// Create a new user cart object
	tests := []struct {
		name        string
		usercart    *models.UserCart
		wantErrUser bool
		wantErrCart bool
	}{
		{name: "Test with correct details", usercart: &models.UserCart{User: user1, Cart: cart1}, wantErrUser: false, wantErrCart: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// u := &userSql{
			// 	conn: DB,
			// }
			if err := u.Register(tt.usercart, "", ""); (err != nil) != tt.wantErrUser {
				t.Errorf("userSql.Register() error = %v, wantErr %v", err, tt.wantErrUser)
			}
			if err := u.CreateCart(tt.usercart); (err != nil) != tt.wantErrCart {
				t.Errorf("userSql.CreateCart() error = %v, wantErr %v", err, tt.wantErrCart)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	// Create a new user cart object
	tests := []struct {
		name        string
		email       string
		password    string
		wantErrGet  bool
		wantErrPass bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", email: "bobdence@gmail.com", password: "123456", wantErrGet: false, wantErrPass: false},
		{name: "Test with wrong email details", email: "bobdence@live.com", password: "123456", wantErrGet: true, wantErrPass: true},
		{name: "Test with wrong password details", email: "bobdence@gmail.com", password: "12345", wantErrGet: false, wantErrPass: true},
		{name: "Test with empty email", email: "", password: "12346", wantErrGet: true, wantErrPass: true},
		{name: "Test with empty password", email: "bobdence@gmail.com", password: "", wantErrGet: false, wantErrPass: true},
		{name: "Test with empty details", email: "@gmail.com", password: "", wantErrGet: true, wantErrPass: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// u := &userSql{
			// 	conn: DB,
			// }
			user, err := u.GetByEmail(tt.email)
			if (err != nil) != tt.wantErrGet {
				t.Errorf("userSql.GetByEmail() error = %v, wantErr %v", err, tt.wantErrGet)
				// return
			}

			if user != nil && (user.Password != tt.password) != tt.wantErrPass {
				t.Fatalf("Compare password, err = %v", err)
			}
		})
	}
}

func TestEmailExists(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{name: "Test with correct email", email: "bobdence@gmail.com", want: true},
		{name: "Test with incorrect email", email: "bobdence@live.com", want: false},
		{name: "Test with empty email", email: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := u.ExistsEmail(tt.email)
			if got != tt.want {
				t.Errorf("userSql.ExistsEmail() got = %v, wantErr %v", got, tt.want)
			}
		})
	}
}

func TestIdExists(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want bool
	}{
		{name: "Test with correct id", id: "7d4b4910-9472-4003-8454-ba09d91ac4d7", want: true},
		{name: "Test with incorrect id", id: "thatisme", want: false},
		{name: "Test with empty id", id: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := u.ExistsId(tt.id)
			if got != tt.want {
				t.Errorf("userSql.ExistsId() got = %v, wantErr %v", got, tt.want)
			}
		})
	}
}

func TestUpdateToken(t *testing.T) {
	auth1 := &models.Auth{
		UserId:       "7d4b4910-9472-4003-8454-ba09d91ac4d7",
		AccessToken:  "The latest access token",
		RefreshToken: "The latest access token",
		DateUpdated:  ti.CurrentTime(),
	}

	auth2 := &models.Auth{
		UserId:       "7d4b4910-9472-4003",
		AccessToken:  "The latest access token",
		RefreshToken: "The latest access token",
		DateUpdated:  ti.CurrentTime(),
	}

	auth3 := &models.Auth{
		UserId:       "",
		AccessToken:  "The latest access token",
		RefreshToken: "The latest access token",
		DateUpdated:  ti.CurrentTime(),
	}

	tests := []struct {
		name string
		auth *models.Auth
		want bool
	}{
		{name: "Test with empty auth details", auth: &models.Auth{}, want: false},
		{name: "Test with correct details", auth: auth1, want: false},
		{name: "Test with incorrect user id", auth: auth2, want: false},
		{name: "Test with empty user id", auth: auth3, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := u.UpdateTokens(tt.auth); (err != nil) != tt.want {
				t.Errorf("userSql.UpdateToken() err = %v, wantErr %v", err, tt.want)
			}
		})
	}

}
