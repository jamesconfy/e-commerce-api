package service_test

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
		Price:       float64(rand.Intn(5000)),
	}
}

func generateEditProductForm() *forms.EditProduct {
	return &forms.EditProduct{
		Name:        faker.Name(),
		Description: faker.Sentence(),
		Image:       faker.DomainName(),
		Price:       float64(rand.Intn(5000)),
	}
}

func createAndAddProduct(product *forms.Product, user *models.User) *models.Product {
	if product == nil {
		product = generateProductForm()
	}

	if user == nil {
		user = createAndRegisterUser(nil)
	}

	resultProduct, err := productSrv.Add(product, user.Id)
	if err != nil {
		panic(err)
	}

	return resultProduct
}

func generateRating() *forms.Rating {
	return &forms.Rating{
		Value: rand.Intn(5),
	}
}
