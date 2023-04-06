package models

import "time"

type Rating struct {
	Value       int       `json:"value"`
	ProductId   string    `json:"product_id"`
	UserId      string    `json:"user_id"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
