package models

import "time"

type Cart struct {
	Id          string    `json:"cart_id"`
	UserId      string    `json:"-"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
