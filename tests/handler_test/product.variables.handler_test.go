package handler_test

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	"math/rand"

	"github.com/bxcodec/faker/v4"
)

func generateProductForm() *forms.Product {
	return &forms.Product{
		Name:        faker.Name(),
		Description: faker.Sentence(),
		Image:       faker.DomainName(),
		Price:       rand.Float64() * 1000,
	}
}

func createAndAddProduct(user *models.User, product *forms.Product) *models.Product {
	if user == nil {
		user = createAndRegisterUser(nil)
	}

	if product == nil {
		product = generateProductForm()
	}

	resultProduct, err := productSrv.Add(product, user.Id)
	if err != nil {
		panic(err)
	}

	return resultProduct
}

func generateRatingForm() *forms.Rating {
	return &forms.Rating{
		Value: rand.Intn(5-1) + 1,
	}
}
