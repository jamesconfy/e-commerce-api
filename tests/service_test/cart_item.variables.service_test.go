package service_test

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	"math/rand"
)

func generateCartItem(productId string) *forms.CartItem {
	return &forms.CartItem{
		ProductId: productId,
		Quantity:  rand.Intn(100-1) + 1,
	}
}

func createAndAddItem(user *models.User, product *models.Product) *models.Item {
	if user == nil {
		user = createAndRegisterUser(nil)
	}

	if product == nil {
		product = createAndAddProduct(nil, nil)
	}

	cartItem := generateCartItem(product.Id)

	item, err := cartItemSrv.Add(cartItem, user.Id)
	if err != nil {
		panic(err)
	}

	return item
}
