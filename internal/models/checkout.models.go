package models

import (
	"sync"
	"time"
)

type Checkout struct {
	sync.Mutex
	Id            string    `json:"id"`
	Quantity      int       `json:"total_quantity"`
	ProductId     string    `json:"product_id"`
	CartId        string    `json:"cart_id"`
	Status        string    `json:"status"`
	DateCreated   time.Time `json:"date_created"`
	DateUpdated   time.Time `json:"date_updated"`
	PaymentMethod string    `json:"payment_method"`
}
