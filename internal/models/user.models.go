package models

type User struct {
	Id          string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}
