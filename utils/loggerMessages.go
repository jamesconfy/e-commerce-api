package utils

import (
	"e-commerce/internal/models/userModels"
	"fmt"
)

var Messages message

type message struct {
}

// Create user validation error
func (m message) CreateUserValidationError(req *userModels.CreateUserReq, err error) (str string) {
	str = fmt.Sprintf("Error when validating create user request || Email: %s || First Name: %s || Last Name: %s || Password: %s || Phone Number: %s || Address: %s || Error: %s", req.Email, req.FirstName, req.LastName, req.Password, req.PhoneNumber, req.PhoneNumber, err)
	return
}

// Error when creating hashed password for provided password
func (m message) CreateUserPasswordError(req *userModels.CreateUserReq, err error) (str string) {
	str = fmt.Sprintf("Passwords provided do not match || Email: %s || Password: %s || Error: %s", req.Email, req.Password, err)
	return str
}

// Get user by email when creating user error
func (m message) CreateUserGetByEmailError(req *userModels.GetByEmailRes, err error) (str string) {
	str = fmt.Sprintf("Error when trying to register new user with email that already exists || Email: %s || UserId: %s || Date Created: %s || Error: %s", req.Email, req.UserId, req.DateCreated, err)
	return
}

// Error when adding user to database
func (m message) CreateUserAddToRepo(req *userModels.CreateUserReq, err error) (str string) {
	str = fmt.Sprintf("Error when adding created user to database || UserId: %s || Email: %s || Password: %s || Date_Created: %s || Error: %s", req.UserId, req.Email, req.Password, req.DateCreated, err)
	return
}

// Create user success message
func (m message) CreateUserSuccess(req *userModels.CreateUserRes) (str string) {
	str = fmt.Sprintf("User created successfully || UserId: %s || Email: %s || Access_Token: %s || Refresh_Token: %s || Date_Created: %s", req.UserId, req.Email, req.Token, req.RefreshToken, req.DateCreated)
	return
}

// Login user validation error
func (m message) LoginUserValidationError(req *userModels.LoginReq) (str string) {
	str = fmt.Sprintf("Error when validating login user request || Email: %s || Password: %s", req.Email, req.Password)
	return
}

// Get user by email when logging in user error
func (m message) LoginUserGetByEmailError(req *userModels.LoginReq) (str string) {
	str = fmt.Sprintf("Error when trying to get user with that email, user don't exists || Email: %s", req.Email)
	return
}

// Password hash does not match when comparing user password and password in the database error
func (m message) LoginUserPasswordError(userId string, req *userModels.LoginReq) (str string) {
	str = fmt.Sprintf("Passwords provided do not match || UserId: %s || Email: %s || Password: %s", userId, req.Email, req.Password)
	return str
}

// Error when creating access token or refresh token
func (m message) CreateTokenError(userId, email string) (str string) {
	str = fmt.Sprintf("Error creating access or refresh token || UserId: %s || Email: %s", userId, email)
	return
}

// Login user success message
func (m message) LoginUserSuccess(req *userModels.LoginRes) (str string) {
	str = fmt.Sprintf("User logged in successfully || UserId: %s || Email: %s || Access_Token: %s || Refresh_Token: %s || Date_Created: %s", req.UserId, req.Email, req.Token, req.RefreshToken, req.DateCreated)
	return
}
