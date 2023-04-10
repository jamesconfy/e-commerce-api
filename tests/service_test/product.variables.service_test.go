package service_test

import (
	"e-commerce/internal/forms"
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
