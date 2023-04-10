package repo_test

import (
	"e-commerce/internal/models"

	"github.com/bxcodec/faker/v4"
)

func generateCart(user *models.User) *models.Cart {
	if user == nil {
		user = createAndAddUser(nil)
	}

	return &models.Cart{
		Id:     faker.UUIDDigit(),
		UserId: user.Id,
	}
}

func createAndAddCart(user *models.User) *models.Cart {
	if user == nil {
		user = createAndAddUser(nil)
	}

	cart := generateCart(user)

	cart, err := c.Add(cart)
	if err != nil {
		panic(err)
	}

	return cart
}
