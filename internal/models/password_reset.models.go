package models

import "time"

type PasswordReset struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
