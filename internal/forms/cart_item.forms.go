package forms

type CartItem struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}
