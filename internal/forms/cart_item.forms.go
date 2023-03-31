package forms

type CartItem struct {
	// Id          string   `json:"id"`
	// ProductId string `json:"product_id"`
	Quantity int `json:"quantity" validate:"required,min=1"`
}
