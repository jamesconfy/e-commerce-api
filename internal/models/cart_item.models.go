package models

import "time"

type CartItem struct {
	CartId      string    `json:"cart_id"`
	Product     *Product  `json:"product"`
	ProductId   string    `json:"product_id"`
	Quantity    int       `json:"quantity"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
