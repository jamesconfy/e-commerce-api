package models

import "time"

type Cart struct {
	Id          string      `json:"cart_id"`
	UserId      string      `json:"-"`
	TotalPrice  float64     `json:"total_price"`
	DateCreated time.Time   `json:"date_created"`
	DateUpdated time.Time   `json:"date_updated"`
	Items       []*CartItem `json:"cart_items"`
}

func (c *Cart) Total() float64 {
	var total float64 = 0
	for _, cartItem := range c.Items {
		total += cartItem.Product.Price * float64(cartItem.Quantity)
	}

	return total
}
