package models

type Product struct {
	Id          string  `json:"id"`
	UserId      string  `json:"user_id"`
	Name        string  `json:"name" validate:"required,min=1"`
	Description string  `json:"description" validate:"required,min=10"`
	Price       float64 `json:"price" validate:"required"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
	Image       string  `json:"product_image"`
}
