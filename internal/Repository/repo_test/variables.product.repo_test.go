package repo_test

import (
	"e-commerce/internal/models"
	"fmt"
	"math/rand"

	"github.com/bxcodec/faker/v4"
)

func generateProduct() *models.Product {
	user := createAndRegisterTestUser(nil)

	return &models.Product{
		Id:          faker.UUIDDigit(),
		UserId:      user.Id,
		Name:        faker.Name(),
		Description: faker.Word(),
		Price:       rand.Float64(),
		Image:       faker.URL(),
	}
}

func createAndAddProduct(product *models.Product) *models.Product {
	if product == nil {
		product = generateProduct()
	}

	product, err := p.Add(product)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return product
}

func editProduct(product *models.Product) *models.Product {
	amount, err := faker.RandomInt(100)
	if err != nil {
		panic(err)
	}

	product.Name = faker.Name()
	product.Description = faker.Word()
	product.Price = float64(amount[0])

	return product
}

func generateRating(rating []int, product *models.Product, user *models.User) *models.Rating {
	var value int

	if rating == nil {
		value = rand.Intn(5)
		rating = append(rating, value)
	}

	if user == nil {
		user = createAndRegisterTestUser(nil)
	}

	if product == nil {
		product = createAndAddProduct(nil)
	}

	return &models.Rating{
		Value:     rating[0],
		ProductId: product.Id,
		UserId:    user.Id,
	}
}
