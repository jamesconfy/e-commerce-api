package repo_test

import (
	"e-commerce/internal/models"
	"math/rand"

	"github.com/bxcodec/faker/v4"
)

func generateCart(user *models.User) *models.Cart {
	if user == nil {
		user = createAndRegisterUser(nil)
	}

	return &models.Cart{
		Id:     faker.UUIDDigit(),
		UserId: user.Id,
	}
}

func createAndAddCart(user *models.User) *models.Cart {
	if user == nil {
		user = createAndRegisterUser(nil)
	}

	cart := generateCart(user)

	cart, err := c.CreateCart(cart)
	if err != nil {
		panic(err)
	}

	return cart
}

func generateItem(user *models.User, product *models.Product) *models.CartItem {
	if user == nil {
		user = createAndRegisterUser(nil)
	}

	cart := createAndAddCart(user)

	if product == nil {
		product = createAndAddProduct(nil)
	}

	req := &models.CartItem{
		CartId:    cart.Id,
		Product:   product,
		ProductId: product.Id,
		Quantity:  rand.Intn(100),
	}

	return req
}

func createAndAddItem(user *models.User, product *models.Product) *models.CartItem {
	if user == nil {
		user = createAndRegisterUser(nil)
	}

	if product == nil {
		product = createAndAddProduct(nil)
	}

	item := generateItem(user, product)

	item, err := c.AddItem(item, user.Id)
	if err != nil {
		panic(err)
	}

	return item
}
