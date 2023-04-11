package handler_test

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	"math/rand"
)

func generateCartItem(product *models.Product) *forms.CartItem {
	if product == nil {
		product = createAndAddProduct(nil, nil)
	}

	return &forms.CartItem{
		ProductId: product.Id,
		Quantity:  rand.Intn(1000-1) + 1,
	}
}

func createAndAddItem(user *models.User, product *models.Product) *models.Item {
	if user == nil {
		user = createAndRegisterUser(nil)
	}

	if product == nil {
		product = createAndAddProduct(nil, nil)
	}

	cartItem := generateCartItem(product)

	item, err := cartItemSrv.Add(cartItem, user.Id)
	if err != nil {
		panic(err)
	}

	return item
}
