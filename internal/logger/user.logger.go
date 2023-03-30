package logger

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	se "e-commerce/internal/serviceerror"
	"fmt"
)

// Internal server error
func (m Messages) InternalServerError(err *se.ServiceError) (str string) {
	str = fmt.Sprintf("Internal server error occurred || Error: %v", err)
	return
}

// Create user validation error
func (m Messages) CreateValidationError(user *forms.Signup, err error) (str string) {
	str = fmt.Sprintf("Error when validating create user request || Email: %s || First Name: %s || Last Name: %s || Password: %s || Phone Number: %s || Address: %s || Error: %s", user.Email, user.FirstName, user.LastName, user.Password, user.PhoneNumber, user.PhoneNumber, err)
	return
}

// Error when creating hashed password for provided password
func (m Messages) CreatePasswordError(user *models.User, err error) (str string) {
	str = fmt.Sprintf("Error when hashing password || Email: %s || Password: %s || Error: %s", user.Email, user.Password, err)
	return str
}

// Get user by email when creating user error
func (m Messages) CreateUserExists(email string) (str string) {
	str = fmt.Sprintf("Error when trying to register new user with email that already exists || Email: %s", email)
	return
}

// Error when adding user to database
func (m Messages) CreateAddToRepo(user *models.User, err error) (str string) {
	str = fmt.Sprintf("Error when adding created user to database || UserId: %s || Email: %s || Password: %s || Date_Created: %s || Error: %s", user.Id, user.Email, user.Password, user.DateCreated, err)
	return
}

// Create user success Messages
func (m Messages) CreateSuccess(user *models.User, accessToken, refreshToken string) (str string) {
	str = fmt.Sprintf("User created successfully || UserId: %s || Email: %s || Access_Token: %s || Refresh_Token: %s || Date_Created: %s", user.Id, user.Email, accessToken, refreshToken, user.DateCreated)
	return
}

// Login user validation error
func (m Messages) LoginValidationError(req *forms.Login, err error) (str string) {
	str = fmt.Sprintf("Error when validating login user request || Email: %s || Password: %s || Error: %v", req.Email, req.Password, err)
	return
}

// Login user check if user exists
func (m Messages) LoginEmailExists(email string) (str string) {
	str = fmt.Sprintf("No user with that email address || Email: %s", email)
	return
}

// Get user by email when logging in user error
func (m Messages) LoginGetError(req *forms.Login) (str string) {
	str = fmt.Sprintf("Error when trying to get user with that email, user don't exists || Email: %s", req.Email)
	return
}

// Password hash does not match when comparing user password and password in the database error
func (m Messages) LoginPasswordError(req *forms.Login, userId string) (str string) {
	str = fmt.Sprintf("Passwords provided do not match || UserId: %s || Email: %s || Password: %s", userId, req.Email, req.Password)
	return str
}

func (m Messages) UpdateTokensError(auth *models.Auth) (str string) {
	str = fmt.Sprintf("Error when trying to update users access and refresh token || UserId: %s || AccessToken: %s || RefreshToken: %s || DateUpdated: %s", auth.UserId, auth.AccessToken, auth.RefreshToken, auth.DateUpdated)
	return
}

// Error when creating access token or refresh token
func (m Messages) CreateTokenError(userId, email string) (str string) {
	str = fmt.Sprintf("Error creating access or refresh token || UserId: %s || Email: %s", userId, email)
	return
}

// Login user success Messages
func (m Messages) LoginSuccess(auth *models.Auth) (str string) {
	str = fmt.Sprintf("User logged in successfully || UserId: %s || Email: %s || Access_Token: %s || Refresh_Token: %s || Date_Created: %s", auth.UserId, auth.User.Email, auth.AccessToken, auth.RefreshToken, auth.DateUpdated)
	return
}

// Get by id user not exists
func (m Messages) GetRepoError(userId string) (str string) {
	str = fmt.Sprintf("No user witht that id || Id: %s", userId)
	return
}

// Get by id user not exists
func (m Messages) GetFetchUserError(userId string, err error) (str string) {
	str = fmt.Sprintf("No user witht that id || Id: %s || Error: %v", userId, err)
	return
}

func (m Messages) GetFetchUserSuccess(user *models.User) (str string) {
	str = fmt.Sprintf("User gotten successfully || Id: %v || Email: %v || DateCreated: %v || DateUpdated: %v", user.Id, user.Email, user.DateCreated, user.DateUpdated)
	return
}
