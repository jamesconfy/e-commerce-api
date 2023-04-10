package service_test

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"

	"github.com/bxcodec/faker/v4"
)

func generateUserForm() *forms.Signup {
	return &forms.Signup{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		Password:    faker.Password(),
	}
}

func createAndRegisterUser(user *forms.Signup) *models.User {
	if user == nil {
		user = generateUserForm()
	}

	resultUser, err := userSrv.Add(user)
	if err != nil {
		panic(err)
	}

	return resultUser.User
}
