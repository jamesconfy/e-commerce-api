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

func TestPhoneExists(t *testing.T) {
	user := createAndAddUser(nil)

	tests := []struct {
		name  string
		phone string
		want  bool
	}{
		{name: "Test with correct id", phone: user.PhoneNumber, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := u.ExistsPhone(tt.phone)
			if got != tt.want {
				t.Errorf("userSql.ExistsPhone() got = %v, wantErr %v", got, tt.want)
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

func TestEditUser(t *testing.T) {
	user := createAndAddUser(nil)
	editUser := generateUser()

	tests := []struct {
		name string
		id   string
		user *models.User
		want bool
	}{
		{name: "Test with correct id", id: user.Id, user: editUser, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := u.Edit(tt.user, tt.id)
			if (err != nil) != tt.want {
				t.Errorf("userSql.Edit() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestGetAll(t *testing.T) {

	for i := 0; i < 10; i++ {
		_ = createAndAddUser(nil)
	}

	tests := []struct {
		name string
		page int
		want bool
	}{
		{name: "Test with correct id", page: 1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := u.GetAll(tt.page)
			if (err != nil) != tt.want {
				t.Errorf("userSql.GetAll() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
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
			err := u.Delete(tt.id)
			if (err != nil) != tt.want {
				t.Errorf("userSql.Delete() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}
