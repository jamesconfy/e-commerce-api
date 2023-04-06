package repo_test

import (
	"e-commerce/internal/models"

	"testing"
)

func TestCreateCart(t *testing.T) {
	cart1 := generateCart(nil)

	tests := []struct {
		name    string
		cart    *models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct user id", cart: cart1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.CreateCart(tt.cart)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartSql.CreateCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCart(t *testing.T) {
	cart := createAndAddCart(nil)

	tests := []struct {
		name    string
		userId  string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct user id", userId: cart.UserId, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetCart(tt.userId)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartSql.GetCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddToCart(t *testing.T) {
	item := generateItem(nil, nil)

	tests := []struct {
		name    string
		item    *models.CartItem
		userId  string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", item: item, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.AddItem(tt.item, tt.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("cartSql.AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetItem(t *testing.T) {
	user := createAndRegisterUser(nil)
	product := createAndAddProduct(nil)
	_ = createAndAddItem(user, product)
	// fmt.Println(item)

	tests := []struct {
		name      string
		productId string
		cartId    string
		wantErr   bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", productId: product.Id, cartId: user.Id, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetItem(tt.productId, tt.cartId)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartSql.GetItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
