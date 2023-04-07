package logger

import (
	"e-commerce/internal/models"
	"fmt"
)

// Add to cart repo error
func (m Messages) AddCartRepoError(cart *models.Cart, err error) (str string) {
	str = fmt.Sprintf("Error when creating user cart || Id: %v || UserId: %v || DateCreated: %v || Error: %v", cart.Id, cart.UserId, cart.DateCreated, err)
	return
}

// Get cart repo error
func (m Messages) GetCartRepoErrror(userId string, err error) (str string) {
	str = fmt.Sprintf("Error when trying to get user cart || UserId: %v || Error: %v", userId, err)
	return
}

// Get cart success message
func (m Messages) GetCartSuccess(items *models.Cart) (str string) {
	str = fmt.Sprintf("Cart gotten successfully || Id: %v || DateCreated: %v || DateUpdated: %v", items.Id, items.DateCreated, items.DateUpdated)
	return
}

// Clear cart repo error
func (m Messages) ClearCartRepoError(userId string, err error) (str string) {
	str = fmt.Sprintf("Error when trying to clear cart || UserId: %v || Error: %v", userId, err)
	return
}

// Clear cart success message
func (m Messages) ClearCartSuccess(userId string) (str string) {
	str = fmt.Sprintf("Cart cleared successfully || UserId: %v", userId)
	return
}

// Compare product owner/creator and logged in user before adding item to cart error
func (m Messages) AddItemCompareUser(productUserId, userId string) (str string) {
	str = fmt.Sprintf("Error forbidden when adding item to cart || ProductUserId: %v || UserId: %v", productUserId, userId)
	return
}

// Add item to cart repo error
func (m Messages) AddItemRepoError(productId, userId string, err error) (str string) {
	str = fmt.Sprintf("Error when adding product to cart || ProductId: %v || UserId: %v || Error: %v", productId, userId, err)
	return
}

// Add item to cart success message
func (m Messages) AddItemSuccess(item *models.Item) (str string) {
	str = fmt.Sprintf("Item successfully added to cart || ProductId: %v || Quantity: %v || DateCreated: %v || DateUpdated: %v", item.ProductId, item.Quantity, item.DateCreated, item.DateUpdated)
	return
}

// Get item repo error
func (m Messages) GetItemRepoError(productId, userId string, err error) (str string) {
	str = fmt.Sprintf("Error when getting item || ProductId: %v || UserId: %v || Error: %v", productId, userId, err)
	return
}

// Get item success message
func (m Messages) GetItemSuccess(item *models.Item) (str string) {
	str = fmt.Sprintf("Item gotten successfully || ProductId: %v || DateCreated: %v || DateUpdated: %v", item.ProductId, item.DateCreated, item.DateUpdated)
	return
}

// Delete item repo error
func (m Messages) DeleteItemRepoError(productId, userId string, err error) (str string) {
	str = fmt.Sprintf("Error when deleting item || ProductId: %v || UserId: %v || Error: %v", productId, userId, err)
	return
}

// Delete item success message
func (m Messages) DeleteItemSuccess(productId, userId string) (str string) {
	str = fmt.Sprintf("Item deleted successfully || ProductId: %v || UserId: %v", productId, userId)
	return
}
