package repo_test

import (
	"e-commerce/internal/models"
	"math/rand"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
)

func generateProduct(user *models.User) *models.Product {
	if user == nil {
		user = createAndAddUser(nil)
	}

	return &models.Product{
		Id:          uuid.New().String(),
		UserId:      user.Id,
		Name:        faker.Name(),
		Description: faker.Word(),
		Price:       rand.Float64() * 1000,
		Image:       faker.URL(),
	}
}

func createAndAddProduct(user *models.User) *models.Product {
	var product *models.Product

	if user == nil {
		product = generateProduct(nil)
	}

	product, err := p.Add(product)
	if err != nil {
		panic(err)
	}

	return product
}

func editProduct(product *models.Product) *models.Product {
	product.Name = faker.Name()
	product.Description = faker.Word()
	product.Price = rand.Float64() * 1000

	return product
}

func generateRating(rating []int, product *models.Product, user *models.User) *models.Rating {
	var value int

	if rating == nil {
		value = rand.Intn(5)
		rating = append(rating, value)
	}

	if user == nil {
		user = createAndAddUser(nil)
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

func createAndAddRating(value []int, product *models.Product, user *models.User) *models.Rating {
	rating := generateRating(value, product, user)

	rating, err := p.AddRating(rating)
	if err != nil {
		panic(err)
	}

	return rating
}
