package forms

type Checkout struct {
	Amount        float64 `json:"amount" validate:"required,min=0"`
	PaymentMethod string  `json:"payment_method" validate:"omitempty,oneof=CARD PAY-ON-DELIVERY PICKUP-STATION"`
	Products      []*CheckoutProduct
}

type CheckoutProduct struct {
	ProductId string `json:"product_id"`
	Quantity  string `json:"quantity"`
}
