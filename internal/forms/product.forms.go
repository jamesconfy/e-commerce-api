package forms

type Product struct {
	Name        string  `json:"name" validate:"required,min=1"`
	Description string  `json:"description" validate:"required,min=10"`
	Price       float64 `json:"price" validate:"required"`
	Image       string  `json:"image"`
}

type EditProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}
