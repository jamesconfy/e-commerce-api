package forms

type Rating struct {
	Value int `json:"value" validate:"required,min=0,max=5"`
}
