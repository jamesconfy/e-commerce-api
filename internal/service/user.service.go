package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/logger"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/se"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	Validate(req any) error
	Create(req *forms.Signup) (*models.UserCart, *se.ServiceError)
	Login(req *forms.Login) (*models.Auth, *se.ServiceError)
	GetById(userId string) (*models.User, *se.ServiceError)
	//ResetPassword(req *models.PasswordReset) (*userModels.ResetPasswordRes, *serviceerror.ServiceError)
	// ValidateToken(userId, token string) (*userModels.ValidateTokenRes, *serviceerror.ServiceError)
	// ChangePassword(userId string, req *models.ChangePasswordReq) *serviceerror.ServiceError
}

type userSrv struct {
	repo      repo.UserRepo
	authRepo  repo.AuthRepo
	cartRepo  repo.CartRepo
	validator ValidationSrv
	crypto    CryptoSrv
	token     AuthSrv
	email     EmailService
	loggerSrv LogSrv
	message   logger.Messages
}

func (u *userSrv) Validate(req any) error {
	err := u.validator.Validate(req)
	if err != nil {
		u.loggerSrv.Error(u.message.ValidationError(req, err))
		return err
	}

	return nil
}

func (u *userSrv) Create(req *forms.Signup) (*models.UserCart, *se.ServiceError) {
	if err := u.Validate(req); err != nil {
		return nil, se.Validating(err)
	}

	// Check if email already exists in the database
	if u.repo.ExistsEmail(req.Email) {
		u.loggerSrv.Error(u.message.CreateUserExists(req.Email))
		return nil, se.Conflict("User already exists")
	}

	// Create a hash of the user password
	password, err := u.crypto.HashPassword(req.Password)
	if err != nil {
		u.loggerSrv.Error(u.message.CreatePasswordError(req, err))
		return nil, se.Internal(err)
	}

	// Creating User model
	var user models.User

	user.Id = uuid.New().String()
	user.Email = req.Email
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Password = req.Password
	user.PhoneNumber = req.PhoneNumber
	user.Password = password
	user.DateCreated = time.Now().Local()

	// Register user
	resultUser, err := u.repo.Register(&user)
	if err != nil {
		u.loggerSrv.Fatal(u.message.CreateRepoError(&user, err))
		return nil, se.NotFoundOrInternal(err)
	}

	// Create a Cart for the user
	var cart models.Cart

	cart.Id = uuid.New().String()
	cart.UserId = resultUser.Id

	// Add cart to the database
	resultCart, err := u.cartRepo.CreateCart(&cart)
	if err != nil {
		u.loggerSrv.Fatal(u.message.AddCartRepoError(&cart, err))
		return nil, se.NotFoundOrInternal(err)
	}

	// User created successfully return both user and user cart
	u.loggerSrv.Info(u.message.CreateSuccess(&user))
	return &models.UserCart{
		User: resultUser,
		Cart: resultCart,
	}, nil
}

func (u *userSrv) Login(req *forms.Login) (*models.Auth, *se.ServiceError) {
	if err := u.Validate(req); err != nil {
		return nil, se.Validating(err)
	}

	// Check if provided email exists in the database
	if !u.repo.ExistsEmail(req.Email) {
		u.loggerSrv.Error(u.message.LoginEmailExists(req.Email))
		return nil, se.NotFound("User does not exist")
	}

	// Get user by email
	user, err := u.repo.GetByEmail(req.Email)
	if err != nil {
		u.loggerSrv.Error(u.message.LoginGetError(req))
		return nil, se.NotFoundOrInternal(err, "user not found")
	}

	// Compare provided password and database password
	ok := u.crypto.ComparePassword(user.Password, req.Password)
	if !ok {
		// u.loggerSrv.Error(u.message.LoginPasswordError(req, user.Id))
		return nil, se.BadRequest("Passwords do not match")
	}

	// Creating auth models
	var auth models.Auth

	auth.Id = uuid.New().String()
	auth.UserId = user.Id

	// Create access and refresh token
	auth.AccessToken, auth.RefreshToken, err = u.token.Create(user.Id, user.Email)
	if err != nil {
		// u.loggerSrv.Error(u.message.CreateTokenError(user.Id, req.Email))
		return nil, se.Internal(err, "Error when creating tokens")
	}

	// Create/Update auth table
	resultAuth, err := u.authRepo.Create(&auth)
	if err != nil {
		// u.loggerSrv.Error(u.message.CreateTokenRepoError(resultAuth))
		return nil, se.Internal(err, "Error when adding/updating user token")
	}

	// User logged in successfully return user recently updated auth
	// u.loggerSrv.Info(u.message.LoginSuccess(resultAuth))
	return resultAuth, nil
}

func (u *userSrv) GetById(userId string) (*models.User, *se.ServiceError) {
	if !u.repo.ExistsId(userId) {
		u.loggerSrv.Error(u.message.GetRepoError(userId))
		return nil, se.NotFound("No user with that id")
	}

	user, err := u.repo.GetById(userId)
	if err != nil {
		u.loggerSrv.Error(u.message.GetFetchUserError(userId, err))
		return nil, se.NotFoundOrInternal(err, "user not found")
	}

	u.loggerSrv.Error(u.message.GetFetchUserSuccess(user))
	return user, nil
}

func NewUserService(repo repo.UserRepo, authRepo repo.AuthRepo, cartRepo repo.CartRepo, validator ValidationSrv, crypto CryptoSrv, token AuthSrv, email EmailService, logSrv LogSrv) UserService {
	return &userSrv{repo: repo, authRepo: authRepo, cartRepo: cartRepo, validator: validator, crypto: crypto, token: token, email: email, loggerSrv: logSrv}
}
