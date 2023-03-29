package models

type Cart struct {
	Id          string      `json:"cart_id"`
	TotalPrice  float64     `json:"total_price"`
	DateCreated string      `json:"date_created"`
	DateUpdated string      `json:"date_updated"`
	Items       []*CartItem `json:"cart_items"`
}

func (c *Cart) Total() float64 {
	var total float64 = 0
	for _, cartItem := range c.Items {
		total += cartItem.Product.Price * float64(cartItem.Quantity)
	}

	return total
}

type CartItem struct {
	// Id          string   `json:"id"`
	CartId      string   `json:"cart_id"`
	Product     *Product `json:"product"`
	ProductId   string   `json:"product_id"`
	Quantity    int      `json:"quantity"`
	DateCreated string   `json:"date_created"`
	DateUpdated string   `json:"date_updated"`
}
