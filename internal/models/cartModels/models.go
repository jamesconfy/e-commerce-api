package cartModels

type AddToCartReq struct {
	CartItemId  string  `json:"cart_item_id"`
	CartId      string  `json:"cart_id"`
	ProductId   string  `json:"product_id" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required,min=1"`
	Price       float64 `json:"price"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
}

type AddToCartRes struct {
	CartItemId  string  `json:"cart_item_id"`
	CartId      string  `json:"cart_id"`
	ProductId   string  `json:"product_id"`
	UserId      string  `json:"user_id"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
}

type GetItemByIdRes struct {
	CartItemId  string  `json:"cart_item_id"`
	CartId      string  `json:"cart_id"`
	ProductId   string  `json:"product_id"`
	UserId      string  `json:"user_id"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
}

type RemoveFromCart struct {
	CartItemId string `json:"cart_item_id"`
}
