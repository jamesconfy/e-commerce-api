package forms

type Rating struct {
	Value int `json:"value" validate:"omitempty,min=0,max=5"`
}
