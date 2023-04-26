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
	Add(req *forms.Signup) (*models.UserCart, *se.ServiceError)
	Login(req *forms.Login) (*models.Auth, *se.ServiceError)
	GetById(userId string) (*models.User, *se.ServiceError)
	GetAll(pageI int) ([]*models.User, *se.ServiceError)
	Edit(req *forms.EditUser, userId string) (*models.User, *se.ServiceError)
	Delete(userId string) *se.ServiceError
	DeleteAuth(userId, accessToken string) *se.ServiceError
	ClearAuth(userId, accessToken string) *se.ServiceError
}

var _ UserService = &userSrv{}

type userSrv struct {
	userRepo  repo.UserRepo
	authRepo  repo.AuthRepo
	cartRepo  repo.CartRepo
	validator ValidationSrv
	cryptoSrv CryptoSrv
	authSrv   AuthService
	emailSrv  EmailService
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

func (u *userSrv) Add(req *forms.Signup) (*models.UserCart, *se.ServiceError) {
	if err := u.Validate(req); err != nil {
		return nil, se.Validating(err)
	}

	// Check if email already exists in the database
	ok, err := u.userRepo.ExistsEmail(req.Email)
	if ok {
		u.loggerSrv.Error(u.message.CreateUserExists(req.Email))
		return nil, se.ConflictOrInternal(err, "User already exists")
	}

	ok, err = u.userRepo.ExistsPhone(req.PhoneNumber)
	if ok {
		u.loggerSrv.Error(u.message.CreateUserExists(req.Email))
		return nil, se.ConflictOrInternal(err, "Phone number in use")
	}

	// Create a hash of the user password
	password, err := u.cryptoSrv.HashPassword(req.Password)
	if err != nil {
		u.loggerSrv.Error(u.message.CreatePasswordError(req, err))
		return nil, se.Internal(err)
	}

	// Creating User model
	var user models.User

	user.Email = req.Email
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Password = req.Password
	user.PhoneNumber = req.PhoneNumber
	user.Password = password
	user.DateCreated = time.Now().Local()

	// Register user
	resultUser, err := u.userRepo.Add(&user)
	if err != nil {
		u.loggerSrv.Fatal(u.message.CreateRepoError(&user, err))
		return nil, se.NotFoundOrInternal(err)
	}

	// Create a Cart for the user
	var cart models.Cart

	cart.UserId = resultUser.Id

	// Add cart to the database
	resultCart, err := u.cartRepo.Add(&cart)
	if err != nil {
		// u.loggerSrv.Fatal(u.message.AddCartRepoError(&cart, err))
		return nil, se.NotFoundOrInternal(err)
	}

	// User created successfully return both user and user cart
	// u.loggerSrv.Info(u.message.CreateSuccess(&user))
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
	ok, _ := u.userRepo.ExistsEmail(req.Email)
	if !ok {
		u.loggerSrv.Error(u.message.LoginEmailExists(req.Email))
		return nil, se.NotFound("User does not exist")
	}

	// Get user by email
	user, err := u.userRepo.GetByEmail(req.Email)
	if err != nil {
		u.loggerSrv.Error(u.message.LoginGetError(req))
		return nil, se.NotFoundOrInternal(err, "user not found")
	}

	// Compare provided password and database password
	ok = u.cryptoSrv.ComparePassword(user.Password, req.Password)
	if !ok {
		// u.loggerSrv.Error(u.message.LoginPasswordError(req, user.Id))
		return nil, se.BadRequest("Passwords do not match")
	}

	// Creating auth models
	var auth models.Auth

	auth.UserId = user.Id

	// Create access and refresh token
	auth.AccessToken, auth.RefreshToken, err = u.authSrv.Create(user.Id, user.Email)
	if err != nil {
		// u.loggerSrv.Error(u.message.CreateTokenError(user.Id, req.Email))
		return nil, se.Internal(err, "Error when creating tokens")
	}

	// Create/Update auth table
	resultAuth, err := u.authRepo.Add(&auth)
	if err != nil {
		// u.loggerSrv.Error(u.message.CreateTokenRepoError(resultAuth))
		return nil, se.Internal(err, "Error when adding/updating user token")
	}

	// User logged in successfully return user recently updated auth
	// u.loggerSrv.Info(u.message.LoginSuccess(resultAuth))
	return resultAuth, nil
}

func (u *userSrv) GetById(userId string) (*models.User, *se.ServiceError) {
	if _, er := uuid.Parse(userId); er != nil {
		return nil, se.NotFound("user not found")
	}

	user, err := u.userRepo.GetById(userId)
	if err != nil {
		u.loggerSrv.Error(u.message.GetFetchUserError(userId, err))
		return nil, se.NotFoundOrInternal(err, "user not found")
	}

	u.loggerSrv.Info(u.message.GetFetchUserSuccess(user))
	return user, nil
}

func (u *userSrv) GetAll(pageI int) ([]*models.User, *se.ServiceError) {
	users, err := u.userRepo.GetAll(pageI)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "Could not fetch users")
	}

	return users, nil
}

func (u *userSrv) Edit(req *forms.EditUser, userId string) (*models.User, *se.ServiceError) {
	if err := u.Validate(req); err != nil {
		return nil, se.Validating(err)
	}

	editUser, err := u.editUser(req, userId)
	if err != nil {
		return nil, err
	}

	user, er := u.userRepo.Edit(editUser, userId)
	if er != nil {
		return nil, se.NotFoundOrInternal(er)
	}

	return user, nil
}

func (u *userSrv) Delete(userId string) *se.ServiceError {
	err := u.userRepo.Delete(userId)
	if err != nil {
		return se.NotFoundOrInternal(err)
	}

	return nil
}

func (u *userSrv) DeleteAuth(userId, accessToken string) *se.ServiceError {
	err := u.authRepo.Delete(userId, accessToken)
	if err != nil {
		return se.Internal(err)
	}

	return nil
}

func (u *userSrv) ClearAuth(userId, accessToken string) *se.ServiceError {
	err := u.authRepo.Clear(userId, accessToken)
	if err != nil {
		return se.Internal(err)
	}

	return nil
}

func NewUserService(repo repo.UserRepo, authRepo repo.AuthRepo, cartRepo repo.CartRepo, validator ValidationSrv, crypto CryptoSrv, authSrv AuthService, email EmailService, logSrv LogSrv) UserService {
	return &userSrv{userRepo: repo, authRepo: authRepo, cartRepo: cartRepo, validator: validator, cryptoSrv: crypto, authSrv: authSrv, emailSrv: email, loggerSrv: logSrv}
}

// Auxillary function
func (u *userSrv) editUser(req *forms.EditUser, userId string) (*models.User, *se.ServiceError) {
	user, err := u.userRepo.GetById(userId)
	if err != nil {
		return nil, se.Internal(err)
	}

	if req.Email != "" && req.Email != user.Email {
		ok, err := u.userRepo.ExistsEmail(req.Email)
		if ok {
			return nil, se.ConflictOrInternal(err, "User already exists")
		}

		user.Email = req.Email
	}

	if req.FirstName != "" && req.FirstName != user.FirstName {
		user.FirstName = req.FirstName
	}

	if req.LastName != "" && req.LastName != user.LastName {
		user.LastName = req.LastName
	}

	if req.PhoneNumber != "" && req.PhoneNumber != user.PhoneNumber {
		ok, err := u.userRepo.ExistsPhone(req.PhoneNumber)
		if ok {
			return nil, se.ConflictOrInternal(err, "Phone already exists")
		}

		user.PhoneNumber = req.PhoneNumber
	}

	return user, nil
}
