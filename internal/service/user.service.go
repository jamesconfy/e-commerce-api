package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/serviceerror"
	"e-commerce/utils"

	"github.com/google/uuid"
)

type UserService interface {
	Validate(req any) error
	Create(req *forms.Signup) (*models.UserCart, *serviceerror.ServiceError)
	Login(req *forms.Login) (*models.Auth, *serviceerror.ServiceError)
	GetById(userId string) (*models.User, *serviceerror.ServiceError)
	//ResetPassword(req *models.PasswordReset) (*userModels.ResetPasswordRes, *serviceerror.ServiceError)
	// ValidateToken(userId, token string) (*userModels.ValidateTokenRes, *serviceerror.ServiceError)
	// ChangePassword(userId string, req *models.ChangePasswordReq) *serviceerror.ServiceError
}

type userSrv struct {
	repo      repo.UserRepo
	cartRepo  repo.CartRepo
	validator ValidationSrv
	crypto    CryptoSrv
	token     TokenSrv
	email     EmailService
	logSrv    LogSrv
	timeSrv   TimeService
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

func (u *userSrv) Create(req *forms.Signup) (*models.UserCart, *serviceerror.ServiceError) {
	if err := u.Validate(req); err != nil {
		return nil, serviceerror.Validating(err)
	}

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
		return nil, serviceerror.Conflict("User already exists")
	}

	password, err := u.crypto.HashPassword(user.Password)
	if err != nil {
		// u.logSrv.Error(u.message.CreateUserPasswordError(user, err))
		return nil, serviceerror.Internal(err) // serviceerror.New("Could not hash password", err)
	}

	// User
	user.Id = userId
	user.Password = password
	user.DateCreated = u.timeSrv.CurrentTime()

	// Cart
	cart.Id = cartId
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
		return nil, serviceerror.New("Error when creating token", err, serviceerror.ErrServer) // serviceerror.New("Error when creating token", err)
	}

	resultUser, err := u.repo.Register(userCart, auth.AccessToken, auth.RefreshToken)
	if err != nil {
		// u.logSrv.Fatal(u.message.CreateUserAddToRepo(user, err))
		return nil, serviceerror.NotFoundOrInternal(err)
	}

	resultCart, err := u.cartRepo.CreateCart(userCart)
	if err != nil {
		return nil, serviceerror.NotFoundOrInternal(err)
	}

	// u.logSrv.Info(u.message.CreateUserSuccess(user))
	return &models.UserCart{
		User: resultUser,
		Cart: resultCart,
	}, nil
}

func (u *userSrv) Login(req *forms.Login) (*models.Auth, *serviceerror.ServiceError) {
	if err := u.Validate(req); err != nil {
		return nil, serviceerror.Validating(err)
	}

	var auth models.Auth
	if !u.repo.ExistsEmail(req.Email) {
		return nil, serviceerror.NotFound("User does not exist")
	}

	user, err := u.repo.GetByEmail(req.Email)
	if err != nil {
		// u.logSrv.Error(u.message.LoginUserGetByEmailError(req))
		return nil, serviceerror.NotFoundOrInternal(err)
	}

	ok := u.crypto.ComparePassword(user.Password, req.Password)
	if !ok {
		// u.logSrv.Error(u.message.LoginUserPasswordError(user.UserId, req))
		return nil, serviceerror.New("Passwords do not match", err, serviceerror.ErrBadRequest) //.NewInternalServiceError(err)
	}

	auth.User = user
	auth.UserId = user.Id
	auth.DateUpdated = u.timeSrv.CurrentTime()
	// updatetoken.DateUpdated = u.timeSrv.CurrentTime()

	auth.AccessToken, auth.RefreshToken, err = u.token.CreateToken(user.Id, user.Email)
	if err != nil {
		// u.logSrv.Error(u.message.CreateTokenError(user.UserId, req.Email))
		return nil, serviceerror.New("Error when creating token", err, serviceerror.ErrServer)
	}

	if err := u.repo.UpdateTokens(&auth); err != nil {
		// u.logSrv.Error(u.message.UpdateTokensError(&updatetoken))
		return nil, serviceerror.New("Error when updating users token", err, serviceerror.ErrServer)
	}

	// u.logSrv.Info(u.message.LoginUserSuccess(data))
	return &auth, nil
}

func (u *userSrv) GetById(userId string) (*models.User, *serviceerror.ServiceError) {
	if !u.repo.ExistsId(userId) {
		return nil, serviceerror.NotFound("No user with that id")
	}

	user, err := u.repo.GetById(userId) // (userId)
	if err != nil {
		return nil, serviceerror.NotFoundOrInternal(err, "user not found")
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

func NewUserService(repo repo.UserRepo, cartRepo repo.CartRepo, validator ValidationSrv, crypto CryptoSrv, token TokenSrv, email EmailService, logSrv LogSrv, timeSrv TimeService, message utils.Messages) UserService {
	return &userSrv{repo: repo, cartRepo: cartRepo, validator: validator, crypto: crypto, token: token, email: email, logSrv: logSrv, timeSrv: timeSrv, message: message}
}

// Auxillary Function
// func GenerateToken(tokenLength int) string {
// 	rand.Seed(time.Now().UnixNano())
// 	const charset = "0123456789"
// 	b := make([]byte, tokenLength)
// 	for i := range b {
// 		b[i] = charset[rand.Intn(len(charset))]
// 	}
// 	return string(b)
// }

// func CreateMessageBody(firstName, lastName, token string) string {
// 	subject := fmt.Sprintf("Hi %v %v, \n\n", firstName, lastName)
// 	mainBody := fmt.Sprintf("You have requested to reset your password, this is your otp code %v\nBut if you did not request for a change of password, you can ignore this email.\n\nLink expires in 30 minutes!", token)

// 	message := subject + mainBody
// 	return string(message)
// }
