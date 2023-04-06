package repo_test

import (
	"e-commerce/internal/models"

	"github.com/bxcodec/faker/v4"
)

func generateUser() *models.User {
	return &models.User{
		Id:          faker.UUIDDigit(),
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		Email:       faker.Email(),
		PhoneNumber: faker.Phonenumber(),
		Password:    faker.Password(),
	}
}

func createAndRegisterTestUser(user *models.User) *models.User {
	if user == nil {
		user = generateUser()
	}

	user, err := u.Register(user)
	if err != nil {
		panic(err)
	}

	return user
}
