package models

type Auth struct {
	User         *User  `json:"-"`
	UserId       string `json:"-"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	DateUpdated  string `json:"date_updated"`
	// DateCreated  string `json:"date_created"`

	// DateDeleted  string `json:"date_deleted"`
}

type Token struct {
	
}
