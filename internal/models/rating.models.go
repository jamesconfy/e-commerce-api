package models

type Rating struct {
	Value       int    `json:"value"`
	ProductId   string `json:"product_id"`
	UserId      string `json:"user_id"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

type ProductRating struct {
	Product   *Product `json:"product"`
	ProductId string   `json:"product_id"`
	Rating    float32  `json:"rating"`
}
