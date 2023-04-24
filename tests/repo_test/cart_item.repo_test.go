package repo_test

import (
	"e-commerce/internal/models"
	"testing"
)

func TestAddItem(t *testing.T) {
	user := createAndAddUser(nil)
	cart := createAndAddCart(user)
	item := generateItem(nil)

	tests := []struct {
		name    string
		item    *models.Item
		cart    *models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", item: item, cart: cart, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ci.Add(tt.cart, tt.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetItem(t *testing.T) {
	user := createAndAddUser(nil)
	product := createAndAddProduct(nil)
	cart := createAndAddCart(user)
	_ = createAndAddItem(cart, product)
	// fmt.Println(item)

	tests := []struct {
		name      string
		productId string
		cart      *models.Cart
		wantErr   bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", productId: product.Id, cart: cart, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ci.Get(tt.cart, tt.productId)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemSql.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetItems(t *testing.T) {
	user := createAndAddUser(nil)
	cart := createAndAddCart(user)
	for i := 0; i < 10; i++ {
		_ = createAndAddItem(cart, nil)
	}

	tests := []struct {
		name    string
		cart    *models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", cart: cart, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ci.GetItems(tt.cart)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemSql.GetItems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteItem(t *testing.T) {
	user := createAndAddUser(nil)
	cart := createAndAddCart(user)
	product := createAndAddProduct(nil)
	_ = createAndAddItem(cart, product)

	tests := []struct {
		name      string
		productId string
		cart      *models.Cart
		wantErr   bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", productId: product.Id, cart: cart, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ci.Delete(tt.cart, tt.productId); (err != nil) != tt.wantErr {
				t.Errorf("cartItemSql.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
