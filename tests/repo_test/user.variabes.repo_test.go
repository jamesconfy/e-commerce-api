package repo_test

import (
	"e-commerce/internal/models"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
)

func generateUser() *models.User {
	return &models.User{
		Id:          uuid.New().String(),
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		Email:       faker.Email(),
		PhoneNumber: faker.Phonenumber(),
		Password:    faker.Password(),
	}
}

func createAndAddUser(user *models.User) *models.User {
	if user == nil {
		user = generateUser()
	}

	user, err := u.Add(user)
	if err != nil {
		panic(err)
	}

	return user
}
