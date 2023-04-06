package models

type ProductRating struct {
	Product   *Product `json:"product"`
	ProductId string   `json:"-"`
	Rating    float32  `json:"rating"`
}
