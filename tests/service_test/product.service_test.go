package service_test

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	"fmt"
	"testing"
)

func TestAddProduct(t *testing.T) {
	// Create a new product object
	user := createAndRegisterUser(nil)
	product := generateProductForm()

	tests := []struct {
		name    string
		product *forms.Product
		user    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", product: product, user: user, wantErr: false},
		{name: "Test with empty user object", product: product, user: &models.User{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := productSrv.Add(tt.product, tt.user.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("productSrv.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllProduct(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)
	for i := 0; i < 10; i++ {
		// Create products using the created user
		_ = createAndAddProduct(nil, user)
	}

	tests := []struct {
		name    string
		page    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct page details", page: 1, wantErr: false},
		{name: "Test with correct page details", page: 2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := productSrv.GetAll(tt.page)

			if (err != nil) != tt.wantErr {
				t.Errorf("productSrv.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetProduct(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)
	product := createAndAddProduct(nil, user)

	tests := []struct {
		name    string
		product *models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct page details", product: product, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := productSrv.Get(tt.product.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("productSrv.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEditProduct(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)
	// Create and add product
	product := createAndAddProduct(nil, user)
	// Generate a edit product form
	edit := generateEditProductForm()

	tests := []struct {
		name    string
		product *models.Product
		user    *models.User
		edit    *forms.EditProduct
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct page details", product: product, user: user, edit: edit, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := productSrv.Edit(tt.edit, tt.product.Id, tt.user.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("productSrv.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	// Create a new user object
	user := createAndRegisterUser(nil)
	// Create and add product
	product := createAndAddProduct(nil, user)

	tests := []struct {
		name    string
		product *models.Product
		user    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct page details", product: product, user: user, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := productSrv.Delete(product.Id, user.Id); (err != nil) != tt.wantErr {
				t.Errorf("productSrv.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRateProduct(t *testing.T) {
	// Create a new user object
	user1 := createAndRegisterUser(nil)
	user2 := createAndRegisterUser(nil)
	// Create and add product
	product := createAndAddProduct(nil, user2)

	tests := []struct {
		name    string
		product *models.Product
		user    *models.User
		rating  *forms.Rating
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with different user that created the product", product: product, user: user1, rating: generateRating(), wantErr: false},
		{name: "Test with same user that created the product", product: product, user: user2, rating: generateRating(), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Value: ", tt.rating.Value)
			_, err := productSrv.AddRating(tt.rating, tt.product.Id, tt.user.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("productSrv.AddRating() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
