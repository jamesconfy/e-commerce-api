package userService

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/service/cryptoService"
	"e-commerce/internal/service/emailService"
	se "e-commerce/internal/service/errors"
	loggerService "e-commerce/internal/service/loggerService"
	"e-commerce/internal/service/timeService"
	"e-commerce/internal/service/tokenService"
	validationService "e-commerce/internal/service/validatorService"
	"e-commerce/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	Validate(req any) error
	Create(req *forms.Signup) (*models.UserCart, *se.ServiceError)
	Login(req *forms.Login) (*models.Auth, *se.ServiceError)
	GetId(userId string) (*models.User, *se.ServiceError)
	//ResetPassword(req *models.PasswordReset) (*userModels.ResetPasswordRes, *se.ServiceError)
	// ValidateToken(userId, token string) (*userModels.ValidateTokenRes, *se.ServiceError)
	// ChangePassword(userId string, req *models.ChangePasswordReq) *se.ServiceError
}

type userSrv struct {
	repo      repo.UserRepo
	validator validationService.ValidationSrv
	crypto    cryptoService.CryptoSrv
	token     tokenService.TokenSrv
	email     emailService.EmailService
	logSrv    loggerService.LogSrv
	timeSrv   timeService.TimeService
	message   utils.Messages
}

func (u *userSrv) Validate(req any) error {
	err := u.validator.Validate(req)
	if err != nil {
		//u.logSrv.Error(u.message.CreateUserValidationError(req, err))
		return err
	}

	return nil
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
func (u *userSrv) Create(req *forms.Signup) (*models.UserCart, *se.ServiceError) {
	var auth models.Auth
	var user models.User
	var cart models.Cart

	cartId := uuid.New().String()
	userId := uuid.New().String()
	user.Email = req.Email
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Password = req.Password
	user.PhoneNumber = req.PhoneNumber

	if u.repo.ExistsEmail(req.Email) {
		return nil, se.NewConflict("User already exists")
	}

	password, err := u.crypto.HashPassword(user.Password)
	if err != nil {
		// u.logSrv.Error(u.message.CreateUserPasswordError(user, err))
		return nil, se.NewInternal(err) // se.New("Could not hash password", err)
	}

	// User
	user.Id = userId
	user.Password = password
	user.DateCreated = u.timeSrv.CurrentTime()

	// Cart
	cart.Id = cartId
	cart.UserId = userId
	cart.DateCreated = u.timeSrv.CurrentTime()

	// UserCart
	userCart := &models.UserCart{
		User:   &user,
		Cart:   &cart,
		CartId: cartId,
		UserId: userId,
	}

	auth.AccessToken, auth.RefreshToken, err = u.token.CreateToken(user.Id, user.Email)
	if err != nil {
		// u.logSrv.Error(u.message.CreateTokenError(user.User.Id, user.User.Email))
		return nil, se.New("Error when creating token", err, se.ErrServer) // se.New("Error when creating token", err)
	}

	err = u.repo.Register(userCart, auth.AccessToken, auth.RefreshToken)
	if err != nil {
		// u.logSrv.Fatal(u.message.CreateUserAddToRepo(user, err))
		return nil, se.New("Error saving user to database", err, se.ErrServer) // se.New("Error saving user to database", err)
	}

	err = u.repo.CreateCart(userCart)
	if err != nil {
		return nil, se.NewInternal(err)
	}

	// u.logSrv.Info(u.message.CreateUserSuccess(user))
	return userCart, nil
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
func (u *userSrv) Login(req *forms.Login) (*models.Auth, *se.ServiceError) {
	var auth models.Auth

	if !u.repo.ExistsEmail(req.Email) {
		return nil, se.NewNotFound("User does not exist")
	}

	user, err := u.repo.GetByEmail(req.Email)
	if err != nil {
		// u.logSrv.Error(u.message.LoginUserGetByEmailError(req))
		return nil, se.NewInternal(err) //.NewInternalServiceError(err)
	}

	ok := u.crypto.ComparePassword(user.Password, req.Password)
	if !ok {
		// u.logSrv.Error(u.message.LoginUserPasswordError(user.UserId, req))
		return nil, se.New("Passwords do not match", err, se.ErrBadRequest) //.NewInternalServiceError(err)
	}

	auth.User = user
	auth.UserId = user.Id
	auth.DateUpdated = u.timeSrv.CurrentTime()
	// updatetoken.DateUpdated = u.timeSrv.CurrentTime()

	auth.AccessToken, auth.RefreshToken, err = u.token.CreateToken(user.Id, user.Email)
	if err != nil {
		// u.logSrv.Error(u.message.CreateTokenError(user.UserId, req.Email))
		return nil, se.New("Error when creating token", err, se.ErrServer)
	}

	if err := u.repo.UpdateTokens(&auth); err != nil {
		// u.logSrv.Error(u.message.UpdateTokensError(&updatetoken))
		return nil, se.New("Error when updating users token", err, se.ErrServer)
	}

	// u.logSrv.Info(u.message.LoginUserSuccess(data))
	return &auth, nil
}

func (u *userSrv) GetId(userId string) (*models.User, *se.ServiceError) {
	if !u.repo.ExistsId(userId) {
		return nil, se.NewNotFound("No user with that id")
	}

	user, err := u.repo.GetById(userId) // (userId)
	if user == nil && err != nil {
		return nil, se.NewInternal(err)
	}

	if user != nil && err != nil {
		//u.logSrv.Error(u.message.LoginUserGetByEmailError(req))
		return nil, se.NewNotFound("No user with that id")
	}

	return user, nil
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
// func (u *userSrv) ResetPassword(req *userModels.ResetPasswordReq) (*userModels.ResetPasswordRes, *errorModels.ServiceError) {
// 	var token userModels.ResetPasswordRes
// 	var message emailModels.SendEmailReq

// 	err := u.validator.Validate(req)
// 	if err != nil {
// 		return nil, errorModels.NewValidatingError(err)
// 	}

// 	user, err := u.repo.GetByEmail(req.Email)
// 	if user == nil {
// 		return nil, errorModels.NewCustomServiceError("Email does not exists!", err) //.NewInternalServiceError(err)
// 	}

// 	// Create token, add to database and then send to user's email address
// 	token.UserId = user.UserId
// 	token.TokenId = uuid.New().String()
// 	token.Token = GenerateToken(6)
// 	token.Expiry = time.Now().Add(time.Minute * 30).Format(time.RFC3339)

// 	err = u.repo.CreateToken(&token)
// 	if err != nil {
// 		return nil, errorModels.NewInternalServiceError(err)
// 	}

// 	// Send message to users email, if it exists
// 	message.EmailAddress = user.Email
// 	message.EmailSubject = "Subject: Reset Password Token\n"
// 	message.EmailBody = CreateMessageBody(user.FirstName, user.LastName, token.Token)

// 	defer func(message *emailModels.SendEmailReq) {
// 		err = u.email.SendMail(*message)
// 		log.Println(err)
// 	}(&message)

// 	return &token, nil
// }

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
// func (u *userSrv) ValidateToken(userId, token string) (*userModels.ValidateTokenRes, *errorModels.ServiceError) {
// 	// err := u.validator.Validate(req)
// 	// if err != nil {
// 	// 	return nil, errorModels.NewValidatingError(err)
// 	// }

// 	tokenDB, err := u.repo.ValidateToken(userId, token)
// 	if err != nil {
// 		return nil, errorModels.NewCustomServiceError("Unable to validate token, check the provided token or userId", err)
// 	}

// 	timeNow := time.Now().Format(time.RFC3339)
// 	if tokenDB.Expiry < timeNow {
// 		return nil, errorModels.NewCustomServiceError("Token has expired", nil)
// 	}

// 	return tokenDB, nil
// }

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
// func (u *userSrv) ChangePassword(userId string, req *userModels.ChangePasswordReq) *errorModels.ServiceError {
// 	err := u.validator.Validate(req)
// 	if err != nil {
// 		return errorModels.NewCustomServiceError("Password not equal to Confirm Password, please check!", err)
// 	}

// 	user, errU := u.repo.GetById(userId)
// 	if errU != nil {
// 		return errorModels.NewCustomServiceError("Cannot find user with that user id", err)
// 	}

// 	if ok := u.crypto.ComparePassword(user.Password, req.Password); ok {
// 		return errorModels.NewCustomServiceError("The new password cannot be the same as your old password!", nil)
// 	}

// 	password, err := u.crypto.HashPassword(req.Password)
// 	if err != nil {
// 		return errorModels.NewCustomServiceError("Error when hashing password", err)
// 	}

// 	if err := u.repo.ChangePassword(userId, password); err != nil {
// 		return errorModels.NewCustomServiceError("Error when changing password", err)
// 	}

// 	return nil
// }

func New(repo repo.UserRepo, validator validationService.ValidationSrv, crypto cryptoService.CryptoSrv, token tokenService.TokenSrv, email emailService.EmailService, logSrv loggerService.LogSrv, timeSrv timeService.TimeService, message utils.Messages) UserService {
	return &userSrv{repo: repo, validator: validator, crypto: crypto, token: token, email: email, logSrv: logSrv, timeSrv: timeSrv, message: message}
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
