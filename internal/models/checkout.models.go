package models

type Checkout struct {
	Id        string  `json:"id"`
	Amount    float64 `json:"amount"`
	Quantity  int     `json:"total_quantity"`
	ProductId string  `json:"product_id"`
	UserId    string  `json:"user_id"`
	// Product       *Product `json:"product"`
	Status        string `json:"status"`
	DateCreated   string `json:"date_created"`
	DateUpdated   string `json:"date_updated"`
	PaymentMethod string `json:"payment_method"`
}

func (c Checkout) Price() float64 {
	return c.Amount * float64(c.Quantity)
}
