package logger

import (
	"e-commerce/internal/models"
	"fmt"
)

func (m Messages) GetCartRepoErrror(userId string, err error) (str string) {
	str = fmt.Sprintf("Error when trying to get user cart || UserId: %v || Error: %v", userId, err)
	return
}

func (m Messages) GetCartSuccess(items *models.Cart) (str string) {
	str = fmt.Sprintf("Cart gotten successfully || Id: %v || DateCreated: %v || DateUpdated: %v || Total: %v", items.Id, items.DateCreated, items.DateUpdated, items.TotalPrice)
	return
}
