package cartModels

type CreateCart struct{
	CartId string `json:""`
}

type AddToCart struct {
	ProductId   string `json:"product_id" validate:"required"`
	UserId      string `json:"user_id"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	CartId      string `json:"cart_id"`
}
