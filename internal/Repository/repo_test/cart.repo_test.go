package repo_test

import (
	"e-commerce/internal/models"

	"testing"
)

func TestAddCart(t *testing.T) {
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
			_, err := c.Add(tt.cart)

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
			_, err := c.Get(tt.userId)

			if (err != nil) != tt.wantErr {
				t.Errorf("cartSql.GetCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClearCart(t *testing.T) {
	user := createAndAddUser(nil)
	cart := createAndAddCart(user)

	for i := 0; i < 10; i++ {
		product := createAndAddProduct(nil)
		_ = createAndAddItem(cart, product)
	}

	tests := []struct {
		name    string
		user    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test with correct details", user: user, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.Clear(tt.user.Id); (err != nil) != tt.wantErr {
				t.Errorf("cartSql.ClearCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
