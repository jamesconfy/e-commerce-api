package utils

import (
	"e-commerce/internal/models/productModels"
	"e-commerce/internal/models/userModels"
	"fmt"
)

type Messages struct {
}

// Create user validation error
func (m Messages) CreateUserValidationError(req *userModels.CreateUserReq, err error) (str string) {
	str = fmt.Sprintf("Error when validating create user request || Email: %s || First Name: %s || Last Name: %s || Password: %s || Phone Number: %s || Address: %s || Error: %s", req.Email, req.FirstName, req.LastName, req.Password, req.PhoneNumber, req.PhoneNumber, err)
	return
}

// Error when creating hashed password for provided password
func (m Messages) CreateUserPasswordError(req *userModels.CreateUserReq, err error) (str string) {
	str = fmt.Sprintf("Passwords provided do not match || Email: %s || Password: %s || Error: %s", req.Email, req.Password, err)
	return str
}

// Get user by email when creating user error
func (m Messages) CreateUserGetByEmailError(req *userModels.GetByEmailRes, err error) (str string) {
	str = fmt.Sprintf("Error when trying to register new user with email that already exists || Email: %s || UserId: %s || Date Created: %s || Error: %s", req.Email, req.UserId, req.DateCreated, err)
	return
}

// Error when adding user to database
func (m Messages) CreateUserAddToRepo(req *userModels.CreateUserReq, err error) (str string) {
	str = fmt.Sprintf("Error when adding created user to database || UserId: %s || Email: %s || Password: %s || Date_Created: %s || Error: %s", req.UserId, req.Email, req.Password, req.DateCreated, err)
	return
}

// Create user success Messages
func (m Messages) CreateUserSuccess(req *userModels.CreateUserRes) (str string) {
	str = fmt.Sprintf("User created successfully || UserId: %s || Email: %s || Access_Token: %s || Refresh_Token: %s || Date_Created: %s", req.UserId, req.Email, req.Token, req.RefreshToken, req.DateCreated)
	return
}

// Login user validation error
func (m Messages) LoginUserValidationError(req *userModels.LoginReq) (str string) {
	str = fmt.Sprintf("Error when validating login user request || Email: %s || Password: %s", req.Email, req.Password)
	return
}

// Get user by email when logging in user error
func (m Messages) LoginUserGetByEmailError(req *userModels.LoginReq) (str string) {
	str = fmt.Sprintf("Error when trying to get user with that email, user don't exists || Email: %s", req.Email)
	return
}

// Password hash does not match when comparing user password and password in the database error
func (m Messages) LoginUserPasswordError(userId string, req *userModels.LoginReq) (str string) {
	str = fmt.Sprintf("Passwords provided do not match || UserId: %s || Email: %s || Password: %s", userId, req.Email, req.Password)
	return str
}

func (m Messages) UpdateTokensError(req *userModels.UpdateTokens) (str string) {
	str = fmt.Sprintf("Error when trying to update users access and refresh token || UserId: %s || AccessToken: %s || RefreshToken: %s || DateUpdated: %s", req.UserId, req.AccessToken, req.RefreshToken, req.DateUpdated)
	return
}

// Error when creating access token or refresh token
func (m Messages) CreateTokenError(userId, email string) (str string) {
	str = fmt.Sprintf("Error creating access or refresh token || UserId: %s || Email: %s", userId, email)
	return
}

// Login user success Messages
func (m Messages) LoginUserSuccess(req *userModels.LoginRes) (str string) {
	str = fmt.Sprintf("User logged in successfully || UserId: %s || Email: %s || Access_Token: %s || Refresh_Token: %s || Date_Created: %s", req.UserId, req.Email, req.AccessToken, req.RefreshToken, req.DateCreated)
	return
}

func (m Messages) AddProductValidationError(req *productModels.AddProductReq) (str string) {
	str = fmt.Sprintf("Error when validating add product request || UserId: %s || Product Name: %s || Product Description: %s", req.UserId, req.Name, req.Description)
	return
}

func (m Messages) AddProductSuccess(req *productModels.AddProductReq) (str string) {
	str = fmt.Sprintf("Product created successfully || ProductId: %s || ProductName: %s || ProductDescription: %s || DateCreated: %s", req.ProductId, req.Name, req.Description, req.DateCreated)
	return
}

func (m Messages) AddProductRepoError(req *productModels.AddProductReq, err error) (str string) {
	str = fmt.Sprintf("Error occured when adding product to database || ProductId: %s || Error: %s || DateCreated: %s", req.ProductId, err, req.DateCreated)
	return
}

func (m Messages) GetProductsRepoError(err error) (str string) {
	str = fmt.Sprintf("Error occured when getting all product || Error: %s", err)
	return
}

func (m Messages) GetProductsSuccess() (str string) {
	str = "Products successfully gotten"
	return
}

func (m Messages) GetProductRepoError(productId string, err error) (str string) {
	str = fmt.Sprintf("Error occured when getting product || ProductId: %s || Error: %s", productId, err)
	return
}

func (m Messages) GetProductSuccess(req *productModels.GetProduct) (str string) {
	str = fmt.Sprintf("Product successfully gotten || ProductId: %s || ProductName: %s || ProductDescription: %s || DateCreated: %s", req.ProductId, req.Name, req.Description, req.DateCreated)
	return
}

func (m Messages) DeleteProductRepoError(productId string, err error) (str string) {
	str = fmt.Sprintf("Error occured when deleting product || ProductId: %s || Error: %s", productId, err)
	return
}

func (m Messages) DeleteProductSuccess(req *productModels.DeleteProduct) (str string) {
	str = fmt.Sprintf("Product successfully deleted || ProductId: %s || ProductName: %s || ProductDescription: %s || DateCreated: %s", req.ProductId, req.Name, req.Description, req.DateCreated)
	return
}

func (m Messages) AddRatingValidationError(req *productModels.AddRatingsReq) (str string) {
	str = fmt.Sprintf("Error when validating add rating request || RatingId: %s || Rating: %v || ProductId: %s || UserId: %s || DateCreated: %s", req.RatingId, req.Rating, req.ProductId, req.UserId, req.DateCreated)
	return
}

func (m Messages) AddRatingRepoError(req *productModels.AddRatingsReq) (str string) {
	str = fmt.Sprintf("Error when adding rating request to database || RatingId: %s || Rating: %v || ProductId: %s || UserId: %s || DateCreated: %s", req.RatingId, req.Rating, req.ProductId, req.UserId, req.DateCreated)
	return
}

func (m Messages) AddRatingSuccess(req *productModels.AddRatingsRes) (str string) {
	str = fmt.Sprintf("Rating successfully added || RatingId: %s || Rating: %v || ProductId: %s || UserId: %s || DateCreated: %s", req.RatingId, req.Rating, req.ProductId, req.UserId, req.DateCreated)
	return
}

func (m Messages) VerifyUserRatingsRepoError(userId, productId string) (str string) {
	str = fmt.Sprintf("User tried to re-rate a product || UserId: %s || ProductId: %s", userId, productId)
	return
}

func (m Messages) VerifyUserRatingsSucess(userId, productId string) (str string) {
	str = fmt.Sprintf("Product rated successfully || UserId: %s || ProductId: %s", userId, productId)
	return
}

func NewMessageUtils() Messages {
	return Messages{}
}
