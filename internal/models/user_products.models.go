package models

type UserProducts struct {
	User      *User    `json:"-"`
	UserId    string   `json:"user_id"`
	Product   *Product `json:"product"`
	ProductId string   `json:"product_id"`
}
