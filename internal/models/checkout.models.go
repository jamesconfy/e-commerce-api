package models

import "sync"

type Checkout struct {
	sync.Mutex
	Id            string `json:"id"`
	Quantity      int    `json:"total_quantity"`
	ProductId     string `json:"product_id"`
	CartId        string `json:"cart_id"`
	Status        string `json:"status"`
	DateCreated   string `json:"date_created"`
	DateUpdated   string `json:"date_updated"`
	PaymentMethod string `json:"payment_method"`
}
