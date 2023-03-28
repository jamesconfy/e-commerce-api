package models

type Product struct {
	Id          string  `json:"id"`
	UserId      string  `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
	Image       string  `json:"product_image"`
}
