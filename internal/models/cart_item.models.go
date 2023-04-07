package models

import "time"

type Item struct {
	Product     *Product  `json:"product"`
	ProductId   string    `json:"-"`
	Quantity    int       `json:"quantity"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type CartItem struct {
	CartId     string  `json:"-"`
	Cart       *Cart   `json:"cart"`
	TotalPrice float64 `json:"total_price"`
	Items      []*Item `json:"cart_items"`
}

func (c *CartItem) Total() float64 {
	var total float64 = 0
	for _, cartItem := range c.Items {
		total += cartItem.Product.Price * float64(cartItem.Quantity)
	}

	return total
}
