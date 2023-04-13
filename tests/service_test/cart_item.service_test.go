package service_test

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	"testing"
)

func TestAddItem(t *testing.T) {
	user := createAndRegisterUser(nil)
	product := createAndAddProduct(nil, nil)
	cartItem := generateCartItem(product.Id)

	tests := []struct {
		name     string
		user     *models.User
		cartItem *forms.CartItem
		wantErr  bool
	}{
		{name: "Test with correct details", user: user, cartItem: cartItem, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := cartItemSrv.Add(tt.cartItem, tt.user.Id)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemSrv.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetItem(t *testing.T) {
	user := createAndRegisterUser(nil)
	product := createAndAddProduct(nil, nil)
	_ = createAndAddItem(user, product)

	tests := []struct {
		name    string
		user    *models.User
		product *models.Product
		wantErr bool
	}{
		{name: "Test with correct details", user: user, product: product, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := cartItemSrv.Get(tt.product.Id, tt.user.Id)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemSrv.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetItems(t *testing.T) {
	user := createAndRegisterUser(nil)
	for i := 0; i < 10; i++ {
		_ = createAndAddItem(user, nil)
	}

	tests := []struct {
		name    string
		user    *models.User
		wantErr bool
	}{
		{name: "Test with correct details", user: user, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := cartItemSrv.GetItems(tt.user.Id)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemSrv.GetItems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteItem(t *testing.T) {
	user := createAndRegisterUser(nil)
	product := createAndAddProduct(nil, nil)
	_ = createAndAddItem(user, product)

	tests := []struct {
		name    string
		user    *models.User
		product *models.Product
		wantErr bool
	}{
		{name: "Test with correct details", user: user, product: product, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cartItemSrv.Delete(tt.product.Id, tt.user.Id)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemSrv.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
