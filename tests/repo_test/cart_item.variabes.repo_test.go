package repo_test

import (
	"e-commerce/internal/models"
	"math/rand"
)

func generateItem(product *models.Product) *models.Item {
	if product == nil {
		product = createAndAddProduct(nil)
	}

	req := &models.Item{
		Product:   product,
		ProductId: product.Id,
		Quantity:  rand.Intn(500-1) + 1,
	}

	return req
}

func createAndAddItem(cart *models.Cart, product *models.Product) *models.Item {
	if cart == nil {
		cart = generateCart(nil)
	}

	if product == nil {
		product = createAndAddProduct(nil)
	}

	item := generateItem(product)

	item, err := ci.Add(cart, item)
	if err != nil {
		panic(err)
	}

	return item
}
