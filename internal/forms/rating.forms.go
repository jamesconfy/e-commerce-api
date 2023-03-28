package forms

type Rating struct {
	Value     int    `json:"value"`
	ProductId string `json:"product_id"`
}
