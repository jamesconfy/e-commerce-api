package userModels

type CreateUserReq struct {
	UserId      string `json:"user_id"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	DateCreated string `json:"date_created"`
}

type CreateUserRes struct {
	UserId       string `json:"user_id"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	DateCreated  string `json:"date_created"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRes struct {
	UserId       string `json:"user_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	DateCreated  string `json:"date_created"`
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GetByEmailRes struct {
	UserId      string `json:"user_id"`
	FirstName   string `json:"first_name"`
	Password    string `json:"password"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	DateCreated string `json:"date_created"`
}

type GetByIdRes struct {
	UserId      string `json:"user_id"`
	FirstName   string `json:"first_name"`
	Password    string `json:"password"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	DateCreated string `json:"date_created"`
}

type ResetPasswordReq struct {
	Email string `json:"email" validate:"required"`
}

type ResetPasswordRes struct {
	UserId  string `json:"user_id"`
	TokenId string `json:"token_id"`
	Token   string `json:"token"`
	Expiry  string `json:"expiry"`
}

type ValidateTokenReq struct {
	Token  string `json:"token" validate:"required,len=6"`
	UserId string `json:"user_id" validate:"required"`
}

type ValidateTokenRes struct {
	Token   string `json:"token"`
	UserId  string `json:"user_id"`
	TokenId string `json:"token_id"`
	Expiry  string `json:"expiry"`
}

type ChangePasswordReq struct {
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
