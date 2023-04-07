package repo_test

import (
	"e-commerce/internal/models"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
)

func generateAuth(user *models.User) *models.Auth {
	if user == nil {
		user = createAndAddUser(nil)
	}

	return &models.Auth{
		Id:           uuid.New().String(),
		User:         user,
		UserId:       user.Id,
		AccessToken:  faker.Jwt(),
		RefreshToken: faker.Jwt(),
	}
}

func createAndAddAuth(auth *models.Auth, user *models.User) *models.Auth {
	if auth == nil {
		auth = generateAuth(user)
	}

	auth, err := a.Add(auth)
	if err != nil {
		panic(err)
	}

	return auth
}
