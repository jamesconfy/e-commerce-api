package forms

type Product struct {
	Name        string  `json:"name" validate:"required,min=1"`
	Description string  `json:"description" validate:"required,min=10"`
	Price       float64 `json:"price" validate:"required,min=0.1"`
	Image       string  `json:"image"`
}

type EditProduct struct {
	Name        string  `json:"name" validate:"omitempty,min=1"`
	Description string  `json:"description" validate:"omitempty,min=10"`
	Price       float64 `json:"price" validate:"omitempty,min=0.1"`
	Image       string  `json:"image"`
}
