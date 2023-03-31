package models

type UserCart struct {
	User   *User  `json:"user"`
	UserId string `json:"-"`
	Cart   *Cart  `json:"cart"`
	CartId string `json:"-"`
}
