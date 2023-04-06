package repo_test

import (
	"e-commerce/internal/models"
	"fmt"
	"testing"
)

func TestAddProduct(t *testing.T) {
	product1 := generateProduct()

	tests := []struct {
		name    string
		product *models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", product: product1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.Add(tt.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("productSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	product1 := createAndAddProduct(nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: product1.Id, wantErr: false},
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
	product1 := createAndAddProduct(nil)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", id: product1.Id, wantErr: false},
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
	for i := 0; i <= 10; i++ {
		_ = createAndAddProduct(nil)
	}

	tests := []struct {
		name    string
		page    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with product page of 1", page: 1, wantErr: false},
		{name: "Test with product page of 2", page: 2, wantErr: false},
		{name: "Test with product page of 3", page: 3, wantErr: false},
		{name: "Test with product page of 4", page: 4, wantErr: false},
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
	product1 := createAndAddProduct(nil)
	product2 := editProduct(product1)

	tests := []struct {
		name    string
		product *models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct product object", product: product2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.Edit(tt.product)

			if (err != nil) != tt.wantErr {
				t.Errorf("productSql.Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddRating(t *testing.T) {
	rating1 := generateRating([]int{3}, nil, nil)

	tests := []struct {
		name    string
		rating  *models.Rating
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct rating", rating: rating1, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rating, err := p.AddRating(tt.rating)
			fmt.Println(rating.Value)
			if (err != nil) != tt.wantErr {
				t.Errorf("productSql.AddRating() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
