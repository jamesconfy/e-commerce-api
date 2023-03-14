package utils

import (
	"e-commerce/internal/models/userModels"
	"fmt"
)

var Messages message

type message struct {
}

func (m message) CreateUserValidationError(req *userModels.CreateUserReq) (str string) {
	str = fmt.Sprintf("Error when validating create user request || Email: %s || First Name: %s || Last Name: %s || Password: %s || Phone Number: %s || Address: %s", req.Email, req.FirstName, req.LastName, req.Password, req.PhoneNumber, req.PhoneNumber)
	return
}

func (m message) LoginUser(req *userModels.LoginReq) (str string) {
	return
}

func (m message) CreateUserGetByEmailError(req *userModels.GetByEmailRes) (str string) {
	str = fmt.Sprintf("Error when trying to register new user with email that already exists || Email: %s || UserId: %s || Date Created: %s", req.Email, req.UserId, req.DateCreated)
	return
}
