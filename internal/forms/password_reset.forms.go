package forms

type PasswordReset struct {
	Email string `json:"email" validate:"required"`
}
