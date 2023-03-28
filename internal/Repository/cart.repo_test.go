package repo_test

import (
	"e-commerce/internal/models"
	"fmt"

	"testing"
)

func TestAddToCart(t *testing.T) {
	item := &models.CartItem{
		CartId:      "123",
		ProductId:   "006ae268-f2a3-4309-9fd9-ef58ca354335",
		Quantity:    10,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
	}

	item1 := &models.CartItem{
		CartId:      "12345",
		ProductId:   "006ae268-f2a3-4309-9fd9-ef58ca354335",
		Quantity:    10,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
	}

	item2 := &models.CartItem{
		CartId:      "12345",
		ProductId:   "006ae268-f2a3-4309",
		Quantity:    10,
		DateCreated: ti.CurrentTime(),
		DateUpdated: ti.CurrentTime(),
	}

	tests := []struct {
		name    string
		item    *models.CartItem
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with empty details", item: &models.CartItem{}, wantErr: true},
		{name: "Test with correct details", item: item, wantErr: false},
		{name: "Test with incorrect cart id details", item: item1, wantErr: true},
		{name: "Test with incorrect product id details", item: item2, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.Add(tt.item); (err != nil) != tt.wantErr {
				t.Errorf("cartSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetItem(t *testing.T) {
	tests := []struct {
		name      string
		productId string
		cartId    string
		wantErr   bool
	}{
		// TODO: Add test cases.
		{name: "Test with empty details", productId: "", cartId: "", wantErr: true},
		{name: "Test with correct details", productId: "006ae268-f2a3-4309-9fd9-ef58ca354335", cartId: "123", wantErr: false},
		{name: "Test with incorrect product id", productId: "006ae268-f2a3-4309", cartId: "123", wantErr: true},
		{name: "Test with incorrect cart id", productId: "006ae268-f2a3-4309-9fd9-ef58ca354335", cartId: "123456", wantErr: true},
		{name: "Test with interchanged id cart id", productId: "123", cartId: "006ae268-f2a3-4309-9fd9-ef58ca354335", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.Get(tt.productId, tt.cartId)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartSql.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCart(t *testing.T) {
	tests := []struct {
		name    string
		cartId  string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct cart id", cartId: "123", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			carts, err := c.GetCart(tt.cartId)
			fmt.Println(carts)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartSql.GetCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
