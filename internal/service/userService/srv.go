package userService

import (
	"e-commerce/internal/Repository/userRepo"
	"e-commerce/internal/models/emailModels"
	"e-commerce/internal/models/errorModels"
	"e-commerce/internal/models/userModels"
	"e-commerce/internal/service/cryptoService"
	"e-commerce/internal/service/emailService"
	"e-commerce/internal/service/tokenService"
	validationService "e-commerce/internal/service/validatorService"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(req *userModels.CreateUserReq) (*userModels.CreateUserRes, *errorModels.ServiceError)
	ResetPassword(req *userModels.ResetPasswordReq) (*userModels.ResetPasswordRes, *errorModels.ServiceError)
	Login(req *userModels.LoginReq) (*userModels.LoginRes, *errorModels.ServiceError)
	ValidateToken(userId, token string) (*userModels.ValidateTokenRes, *errorModels.ServiceError)
	ChangePassword(userId string, req *userModels.ChangePasswordReq) *errorModels.ServiceError
}

type userSrv struct {
	repo      userRepo.UserRepo
	validator validationService.ValidationSrv
	crypto    cryptoService.CryptoSrv
	token     tokenService.TokenSrv
	email     emailService.EmailService
}

// Register User godoc
// @Summary	Register route
// @Description	Register route
// @Tags	Users
// @Accept	json
// @Produce	json
// @Param	request	body	userModels.CreateUserReq	true "Signup Details"
// @Success	200  {object}  userModels.CreateUserRes
// @Failure	400  {object}  errorModels.ServiceError
// @Failure	404  {object}  errorModels.ServiceError
// @Failure	500  {object}  errorModels.ServiceError
// @Router	/users [post]
func (u *userSrv) CreateUser(req *userModels.CreateUserReq) (*userModels.CreateUserRes, *errorModels.ServiceError) {
	err := u.validator.Validate(req)
	if err != nil {
		return nil, errorModels.NewValidatingError(err)
	}

	user, err := u.repo.GetByEmail(req.Email)
	if user != nil {
		return nil, errorModels.NewCustomServiceError("User already exists", err) //.NewInternalServiceError(err)
	}

	password, err := u.crypto.HashPassword(req.Password)
	if err != nil {
		return nil, errorModels.NewCustomServiceError("Could not hash password", err)
	}

	req.UserId = uuid.New().String()
	req.Password = password
	req.DateCreated = time.Now().Local().Format(time.RFC3339)

	err = u.repo.RegisterUser(req)
	if err != nil {
		return nil, errorModels.NewCustomServiceError("Error saving user to database", err)
	}

	token, refresh_token, err := u.token.CreateToken(req.UserId, req.Email)
	if err != nil {
		return nil, errorModels.NewCustomServiceError("Error when creating token", err)
	}

	return &userModels.CreateUserRes{
		UserId:       req.UserId,
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Token:        token,
		RefreshToken: refresh_token,
	}, nil
}

// Login User godoc
// @Summary	Login route
// @Description	Login route
// @Tags	Users
// @Accept	json
// @Produce	json
// @Param	request	body	userModels.LoginReq	true "Login Details"
// @Success	200  {object}  userModels.LoginRes
// @Failure	400  {object}  errorModels.ServiceError
// @Failure	404  {object}  errorModels.ServiceError
// @Failure	500  {object}  errorModels.ServiceError
// @Router	/users/login [post]
func (u *userSrv) Login(req *userModels.LoginReq) (*userModels.LoginRes, *errorModels.ServiceError) {
	err := u.validator.Validate(req)
	if err != nil {
		return nil, errorModels.NewValidatingError(err)
	}

	user, err := u.repo.GetByEmail(req.Email)
	if user == nil {
		return nil, errorModels.NewCustomServiceError("Email does not exists!", err) //.NewInternalServiceError(err)
	}

	ok := u.crypto.ComparePassword(user.Password, req.Password)
	if !ok {
		return nil, errorModels.NewCustomServiceError("Passwords do not match", err) //.NewInternalServiceError(err)
	}

	token, refresh_token, err := u.token.CreateToken(user.UserId, user.Email)
	if err != nil {
		return nil, errorModels.NewCustomServiceError("Error when creating token", err)
	}

	return &userModels.LoginRes{
		UserId:       user.UserId,
		Name:         user.FirstName + user.LastName,
		DateCreated:  user.DateCreated,
		Email:        user.Email,
		Token:        token,
		RefreshToken: refresh_token,
	}, nil
}

// Reset User Password godoc
// @Summary	Reset password route
// @Description	Reset password route
// @Tags	Users
// @Accept	json
// @Produce	json
// @Param	request	body	userModels.ResetPasswordReq	true "Reset Password Details"
// @Success	200  {object}  userModels.ResetPasswordRes
// @Failure	400  {object}  errorModels.ServiceError
// @Failure	404  {object}  errorModels.ServiceError
// @Failure	500  {object}  errorModels.ServiceError
// @Router	/users/reset-password [post]
func (u *userSrv) ResetPassword(req *userModels.ResetPasswordReq) (*userModels.ResetPasswordRes, *errorModels.ServiceError) {
	var token userModels.ResetPasswordRes
	var message emailModels.SendEmailReq

	err := u.validator.Validate(req)
	if err != nil {
		return nil, errorModels.NewValidatingError(err)
	}

	user, err := u.repo.GetByEmail(req.Email)
	if user == nil {
		return nil, errorModels.NewCustomServiceError("Email does not exists!", err) //.NewInternalServiceError(err)
	}

	// Create token, add to database and then send to user's email address
	token.UserId = user.UserId
	token.TokenId = uuid.New().String()
	token.Token = GenerateToken(6)
	token.Expiry = time.Now().Add(time.Minute * 30).Format(time.RFC3339)

	err = u.repo.CreateToken(&token)
	if err != nil {
		return nil, errorModels.NewInternalServiceError(err)
	}

	// Send message to users email, if it exists
	message.EmailAddress = user.Email
	message.EmailSubject = "Subject: Reset Password Token\n"
	message.EmailBody = CreateMessageBody(user.FirstName, user.LastName, token.Token)

	defer func(message *emailModels.SendEmailReq) {
		err = u.email.SendMail(*message)
		log.Println(err)
	}(&message)

	return &token, nil
}

// Validate Token godoc
// @Summary	Validate token route
// @Description	Validate token route
// @Tags	Users
// @Accept	json
// @Produce	json
// @Param	token	query	string	true	"Token"
// @Param	user_id	query	string	true	"User Id"
// @Success	200  {object}  userModels.ValidateTokenRes
// @Failure	400  {object}  errorModels.ServiceError
// @Failure	404  {object}  errorModels.ServiceError
// @Failure	500  {object}  errorModels.ServiceError
// @Router	/users/reset-password/validate-token [post]
func (u *userSrv) ValidateToken(userId, token string) (*userModels.ValidateTokenRes, *errorModels.ServiceError) {
	// err := u.validator.Validate(req)
	// if err != nil {
	// 	return nil, errorModels.NewValidatingError(err)
	// }

	tokenDB, err := u.repo.ValidateToken(userId, token)
	if err != nil {
		return nil, errorModels.NewInternalServiceError("Unable to validate token, check the provided token or userId")
	}

	timeNow := time.Now().Format(time.RFC3339)
	if tokenDB.Expiry < timeNow {
		return nil, errorModels.NewCustomServiceError("Token has expired", nil)
	}

	return tokenDB, nil
}

// Change Password godoc
// @Summary	Change password route
// @Description	Change password route
// @Tags	Users
// @Accept	json
// @Produce	json
// @Param	user_id	query	string	true	"User Id"
// @Param	request	body	userModels.ChangePasswordReq	true "Reset Password Details"
// @Success	200  {string}  string    "Password changed successfully"
// @Failure	400  {object}  errorModels.ServiceError
// @Failure	404  {object}  errorModels.ServiceError
// @Failure	500  {object}  errorModels.ServiceError
// @Router	/users/reset-password/change-password [patch]
func (u *userSrv) ChangePassword(userId string, req *userModels.ChangePasswordReq) *errorModels.ServiceError {
	err := u.validator.Validate(req)
	if err != nil {
		return errorModels.NewValidatingError("Password not equal to Confirm Password, please check!")
	}

	user, errU := u.repo.GetById(userId)
	if errU != nil {
		return errorModels.NewCustomServiceError("Cannot find user with that user id", err)
	}

	if ok := u.crypto.ComparePassword(user.Password, req.Password); ok {
		return errorModels.NewCustomServiceError("The new password cannot be the same as your old password!", nil)
	}

	password, err := u.crypto.HashPassword(req.Password)
	if err != nil {
		return errorModels.NewInternalServiceError("Error when hashing password")
	}

	if err := u.repo.ChangePassword(userId, password); err != nil {
		return errorModels.NewInternalServiceError("Error when changing password")
	}

	return nil
}

func NewUserSrv(repo userRepo.UserRepo, validator validationService.ValidationSrv, crypto cryptoService.CryptoSrv, token tokenService.TokenSrv, email emailService.EmailService) UserService {
	return &userSrv{repo: repo, validator: validator, crypto: crypto, token: token, email: email}
}

// Auxillary Function
func GenerateToken(tokenLength int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "0123456789"
	b := make([]byte, tokenLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func CreateMessageBody(firstName, lastName, token string) string {
	subject := fmt.Sprintf("Hi %v %v, \n\n", firstName, lastName)
	mainBody := fmt.Sprintf("You have requested to reset your password, this is your otp code %v\nBut if you did not request for a change of password, you can ignore this email.\n\nLink expires in 30 minutes!", token)

	message := subject + mainBody
	return string(message)
}
