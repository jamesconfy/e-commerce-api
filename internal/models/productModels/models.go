package productModels

type AddProductReq struct {
	ProductId   string  `json:"product_id"`
	UserId      string  `json:"user_id"`
	Name        string  `json:"name" validate:"required,min=1"`
	Description string  `json:"description" validate:"required,min=10"`
	Price       float64 `json:"price" validate:"required"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
	Image       string  `json:"product_image"`
}

type AddProductRes struct {
	ProductId   string  `json:"product_id"`
	Name        string  `json:"name" validate:"required,min=1"`
	Description string  `json:"description" validate:"required,min=10"`
	Price       float64 `json:"price" validate:"required"`
	Image       string  `json:"product_image"`
}

type GetProductRes struct {
	ProductId   string  `json:"product_id"`
	UserId      string  `json:"user_id"`
	Name        string  `json:"name" validate:"required,min=1"`
	Description string  `json:"description" validate:"required,min=10"`
	Price       float64 `json:"price" validate:"required"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
	Image       string  `json:"product_image"`
	Rating      string  `json:"rating"`
}

type EditProductReq struct {
	ProductId   string  `json:"product_id"`
	UserId      string  `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	DateUpdated string  `json:"date_updated"`
	Image       string  `json:"image"`
}

type EditProductRes struct {
	ProductId   string  `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	DateUpdated string  `json:"date_updated"`
	Image       string  `json:"image"`
}

type AddRatingsReq struct {
	RatingId    string  `json:"rating_id"`
	Rating      float32 `json:"rating" validate:"required,min=0,max=5"`
	ProductId   string  `json:"product_id"`
	UserId      string  `json:"user_id"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
}

type AddRatingsRes struct {
	RatingId    string  `json:"rating_id"`
	Rating      float32 `json:"rating" validate:"required"`
	ProductId   string  `json:"product_id"`
	UserId      string  `json:"user_id"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
}
