package productModels

type AddProductReq struct {
	ProductId   string `json:"product_id"`
	UserId      string `json:"user_id"`
	Name        string `json:"name" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=10"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	Image       string `json:"product_image"`
}

type AddProductRes struct {
}
