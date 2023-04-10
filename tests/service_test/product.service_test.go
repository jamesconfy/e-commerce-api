package service_test

import (
	"e-commerce/internal/forms"
	"testing"
)

func TestAddProduct(t *testing.T) {
	// Create a new product object
	user := createAndRegisterUser(nil)
	product := generateProductForm()

	tests := []struct {
		name    string
		product *forms.Product
		userId  string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", product: product, userId: user.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := productSrv.Add(tt.product, tt.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("product.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
