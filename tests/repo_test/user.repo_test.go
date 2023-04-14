package repo_test

import (
	"e-commerce/internal/models"
	"testing"
)

func TestAddUser(t *testing.T) {
	// Create a new user object
	user := generateUser()

	tests := []struct {
		name    string
		user    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", user: user, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := u.Add(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userSql.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmailExists(t *testing.T) {
	user := createAndAddUser(nil)

	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{name: "Test with correct email", email: user.Email, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := u.ExistsEmail(tt.email)
			if got != tt.want {
				t.Errorf("userSql.ExistsEmail() got = %v, wantErr %v", got, tt.want)
			}
		})
	}
}

func TestIdExists(t *testing.T) {
	user := createAndAddUser(nil)

	tests := []struct {
		name string
		id   string
		want bool
	}{
		{name: "Test with correct id", id: user.Id, want: true},
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

func TestGetByEmail(t *testing.T) {
	user := createAndAddUser(nil)

	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{name: "Test with correct email", email: user.Email, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := u.GetByEmail(tt.email)
			if (err != nil) != tt.want {
				t.Errorf("userSql.GetByEmail() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestGetById(t *testing.T) {
	user := createAndAddUser(nil)

	tests := []struct {
		name string
		id   string
		want bool
	}{
		{name: "Test with correct id", id: user.Id, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := u.GetById(tt.id)
			if (err != nil) != tt.want {
				t.Errorf("userSql.GetById() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

// func TestUpdateToken(t *testing.T) {
// 	auth1 := &models.Auth{
// 		UserId:       "7d4b4910-9472-4003-8454-ba09d91ac4d7",
// 		AccessToken:  "The latest access token",
// 		RefreshToken: "The latest access token",
// 		DateUpdated:  ti.CurrentTime(),
// 	}

// 	auth2 := &models.Auth{
// 		UserId:       "7d4b4910-9472-4003",
// 		AccessToken:  "The latest access token",
// 		RefreshToken: "The latest access token",
// 		DateUpdated:  ti.CurrentTime(),
// 	}

// 	auth3 := &models.Auth{
// 		UserId:       "",
// 		AccessToken:  "The latest access token",
// 		RefreshToken: "The latest access token",
// 		DateUpdated:  ti.CurrentTime(),
// 	}

// 	tests := []struct {
// 		name string
// 		auth *models.Auth
// 		want bool
// 	}{
// 		{name: "Test with empty auth details", auth: &models.Auth{}, want: false},
// 		{name: "Test with correct details", auth: auth1, want: false},
// 		{name: "Test with incorrect user id", auth: auth2, want: false},
// 		{name: "Test with empty user id", auth: auth3, want: false},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := u.Update(tt.auth); (err != nil) != tt.want {
// 				t.Errorf("userSql.UpdateToken() err = %v, wantErr %v", err, tt.want)
// 			}
// 		})
// 	}

// }
