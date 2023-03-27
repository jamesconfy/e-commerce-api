package repo_test

import (
	"e-commerce/internal/models"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestAddProduct(t *testing.T) {
	product1 := &models.Product{
		Id:          uuid.New().String(),
		UserId:      "4567",
		Name:        "Test Product",
		Description: "This is a test product.",
		Price:       9.99,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
		Image:       "test.png",
	}

	product2 := &models.Product{
		Id:          uuid.New().String(),
		UserId:      "4567890",
		Name:        "Test Product",
		Description: "This is a test product.",
		Price:       9.99,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
		Image:       "test.png",
	}

	product3 := &models.Product{
		UserId:      "4567890",
		Name:        "Test Product",
		Description: "This is a test product.",
		Price:       9.99,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
		Image:       "test.png",
	}

	product4 := &models.Product{
		Id:          uuid.New().String(),
		Name:        "Test Product",
		Description: "This is a test product.",
		Price:       9.99,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
		Image:       "test.png",
	}

	product5 := &models.Product{
		Id:          uuid.New().String(),
		UserId:      "4567",
		Description: "This is a test product.",
		Price:       9.99,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
		Image:       "test.png",
	}

	product6 := &models.Product{
		Id:          uuid.New().String(),
		UserId:      "4567",
		Name:        "Test Product",
		Price:       9.99,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
		Image:       "test.png",
	}

	product7 := &models.Product{
		Id:          uuid.New().String(),
		UserId:      "4567",
		Name:        "Test Product",
		Description: "This is a test product.",
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
		Image:       "test.png",
	}

	product8 := &models.Product{
		Id:          uuid.New().String(),
		UserId:      "4567",
		Name:        "Test Product",
		Description: "This is a test product.",
		Price:       9.99,
		DateUpdated: ti.CurrentTime(),
		Image:       "test.png",
	}

	product9 := &models.Product{
		Id:          uuid.New().String(),
		UserId:      "4567",
		Name:        "Test Product",
		Description: "This is a test product.",
		Price:       9.99,
		DateCreated: ti.CurrentTime(),
		Image:       "test.png",
	}

	product10 := &models.Product{
		Id:          uuid.New().String(),
		UserId:      "4567",
		Name:        "Test Product",
		Description: "This is a test product.",
		Price:       9.99,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
	}

	product11 := &models.Product{
		Id:          uuid.New().String(),
		UserId:      "4567",
		Name:        "Test Product",
		Description: "This is a test product.",
		Price:       -9.99,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
		Image:       "test.png",
	}

	tests := []struct {
		name    string
		product *models.Product
		wantErr bool
	}{
		{name: "Test with empty product object", product: &models.Product{}, wantErr: true},
		{name: "Test with correct details", product: product1, wantErr: false},
		{name: "Test with non-existing userid", product: product2, wantErr: true},
		{name: "Test with empty id", product: product3, wantErr: true},
		{name: "Test with empty userid", product: product4, wantErr: true},
		{name: "Test with empty name", product: product5, wantErr: false},
		{name: "Test with empty description", product: product6, wantErr: false},
		{name: "Test with empty price", product: product7, wantErr: false},
		{name: "Test with empty date_created", product: product8, wantErr: true},
		{name: "Test with empty date_updated", product: product9, wantErr: true},
		{name: "Test with empty image", product: product10, wantErr: false},
		{name: "Test with negative pricee", product: product11, wantErr: true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := p.Add(tt.product); (err != nil) != tt.wantErr {
				t.Errorf("productSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with invalid id", id: "jdcah", wantErr: false},
		{name: "Test with empty id", id: "", wantErr: false},
		{name: "Test with correct id", id: "123", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := p.Delete(tt.id); (err != nil) != tt.wantErr {
				t.Errorf("productSql.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetProduct(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with incorrect id", id: "jdcah", wantErr: true},
		{name: "Test with empty id", id: "", wantErr: true},
		{name: "Test with correct id", id: "006ae268-f2a3-4309-9fd9-ef58ca354335", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.GetId(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("productSql.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetProducts(t *testing.T) {
	tests := []struct {
		name    string
		page    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with product page of 1", page: 1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.GetAll(tt.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("productSql.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEditProduct(t *testing.T) {
	product1 := &models.Product{
		Id:          "006ae268-f2a3-4309-9fd9-ef58ca354335",
		UserId:      "4567",
		Name:        "Test Product 1",
		Description: "This is a from testedit to test product.",
		Price:       20.99,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
		Image:       "test1.png",
	}

	product2 := &models.Product{
		Id:          "006ae268-f2a3-4309-9fd9-ef58ca354335",
		UserId:      "4567890",
		Name:        "Test Product",
		Description: "This is a test product.",
		Price:       9.99,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
		Image:       "test.png",
	}

	product3 := &models.Product{
		UserId:      "4567890",
		Name:        "Test Product",
		Description: "This is a test product.",
		Price:       9.99,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
		Image:       "test.png",
	}

	tests := []struct {
		name    string
		product *models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with empty product object", product: &models.Product{}, wantErr: false},
		{name: "Test with correct details", product: product1, wantErr: false},
		{name: "Test with non-existing userid", product: product2, wantErr: false},
		{name: "Test with empty id", product: product3, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := p.Edit(tt.product); (err != nil) != tt.wantErr {
				fmt.Println(err)
				t.Errorf("productSql.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddRating(t *testing.T) {
	rating1 := &models.Rating{
		Value:       4,
		ProductId:   "006ae268-f2a3-4309-9fd9-ef58ca354335",
		UserId:      "4567",
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
	}

	rating2 := &models.Rating{
		Value:       6,
		ProductId:   "006ae268-f2a3-4309-9fd9-ef58ca354335",
		UserId:      "4567",
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
	}

	rating3 := &models.Rating{
		Value:       -1,
		ProductId:   "006ae268-f2a3-4309-9fd9-ef58ca354335",
		UserId:      "4567",
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
	}

	rating4 := &models.Rating{
		Value:       0,
		UserId:      "4567",
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
	}

	rating5 := &models.Rating{
		Value:       4,
		ProductId:   "006ae268-f2a3-4309-9fd9-ef58ca354335",
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
	}

	tests := []struct {
		name    string
		rating  *models.Rating
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with empty rating", rating: &models.Rating{}, wantErr: true},
		{name: "Test with correct details", rating: rating1, wantErr: false},
		{name: "Test with value greater than 5", rating: rating2, wantErr: true},
		{name: "Test with value less than 1", rating: rating3, wantErr: true},
		{name: "Test with empty product id", rating: rating4, wantErr: true},
		{name: "Test with empty user id", rating: rating5, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := p.AddRating(tt.rating); (err != nil) != tt.wantErr {
				fmt.Println(err)
				t.Errorf("productSql.AddRating() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
