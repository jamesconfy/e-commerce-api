package models

import "time"

type Auth struct {
	Id           string    `json:"auth_id"`
	User         *User     `json:"-"`
	UserId       string    `json:"-"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	DateUpdated  time.Time `json:"date_updated"`
	DateCreated  time.Time `json:"date_created"`
	ExpiresAt    time.Time `json:"expires_at"`
}
