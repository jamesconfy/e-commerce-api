package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/logger"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/serviceerror"

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
	loggerSrv LogSrv
	timeSrv   TimeService
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

func (u *userSrv) Create(req *forms.Signup) (*models.UserCart, *serviceerror.ServiceError) {
	if err := u.Validate(req); err != nil {
		return nil, serviceerror.Validating(err)
	}

	var user models.User

	userId := uuid.New().String()
	user.Email = req.Email
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Password = req.Password
	user.PhoneNumber = req.PhoneNumber

	if u.repo.ExistsEmail(req.Email) {
		u.loggerSrv.Error(u.message.CreateUserExists(req.Email))
		return nil, serviceerror.Conflict("User already exists")
	}

	password, err := u.crypto.HashPassword(user.Password)
	if err != nil {
		u.loggerSrv.Error(u.message.CreatePasswordError(&user, err))
		return nil, serviceerror.Internal(err)
	}

	// User
	user.Id = userId
	user.Password = password
	user.DateCreated = u.timeSrv.CurrentTime()

	resultUser, err := u.repo.Register(&user)
	if err != nil {
		u.loggerSrv.Fatal(u.message.CreateRepoError(&user, err))
		return nil, serviceerror.NotFoundOrInternal(err)
	}

	// Cart
	var cart models.Cart

	cartId := uuid.New().String()
	cart.Id = cartId
	cart.UserId = resultUser.Id
	cart.DateCreated = u.timeSrv.CurrentTime()

	resultCart, err := u.cartRepo.CreateCart(&cart)
	if err != nil {
		u.loggerSrv.Fatal(u.message.AddCartRepoError(&cart, err))
		return nil, serviceerror.NotFoundOrInternal(err)
	}

	u.loggerSrv.Info(u.message.CreateSuccess(&user))
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
		u.loggerSrv.Error(u.message.LoginEmailExists(req.Email))
		return nil, serviceerror.NotFound("User does not exist")
	}

	user, err := u.repo.GetByEmail(req.Email)
	if err != nil {
		u.loggerSrv.Error(u.message.LoginGetError(req))
		return nil, serviceerror.NotFoundOrInternal(err, "user not found")
	}

	ok := u.crypto.ComparePassword(user.Password, req.Password)
	if !ok {
		u.loggerSrv.Error(u.message.LoginPasswordError(req, user.Id))
		return nil, serviceerror.BadRequest("Passwords do not match")
	}

	auth.User = user
	auth.UserId = user.Id
	auth.DateUpdated = u.timeSrv.CurrentTime()

	auth.AccessToken, auth.RefreshToken, err = u.token.CreateToken(user.Id, user.Email)
	if err != nil {
		u.loggerSrv.Error(u.message.CreateTokenError(user.Id, req.Email))
		return nil, serviceerror.Internal(err, "Error when creating tokens")
	}

	if err := u.repo.UpdateTokens(&auth); err != nil {
		u.loggerSrv.Error(u.message.UpdateTokensError(&auth))
		return nil, serviceerror.Internal(err, "Error when updating users token")
	}

	u.loggerSrv.Info(u.message.LoginSuccess(&auth))
	return &auth, nil
}

func (u *userSrv) GetById(userId string) (*models.User, *serviceerror.ServiceError) {
	if !u.repo.ExistsId(userId) {
		u.loggerSrv.Error(u.message.GetRepoError(userId))
		return nil, serviceerror.NotFound("No user with that id")
	}

	user, err := u.repo.GetById(userId)
	if err != nil {
		u.loggerSrv.Error(u.message.GetFetchUserError(userId, err))
		return nil, serviceerror.NotFoundOrInternal(err, "user not found")
	}

	u.loggerSrv.Error(u.message.GetFetchUserSuccess(user))
	return user, nil
}

func NewUserService(repo repo.UserRepo, cartRepo repo.CartRepo, validator ValidationSrv, crypto CryptoSrv, token TokenSrv, email EmailService, logSrv LogSrv, timeSrv TimeService) UserService {
	return &userSrv{repo: repo, cartRepo: cartRepo, validator: validator, crypto: crypto, token: token, email: email, loggerSrv: logSrv, timeSrv: timeSrv}
}
