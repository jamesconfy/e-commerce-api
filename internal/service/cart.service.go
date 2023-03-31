package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/logger"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/serviceerror"
)

type CartService interface {
	// Cart
	GetCart(userId string) (*models.Cart, *serviceerror.ServiceError)
	ClearCart(userId string) *serviceerror.ServiceError

	// CartItem
	AddItem(req *forms.CartItem, productId, userId string) (*models.CartItem, *serviceerror.ServiceError)
	GetItem(productId, userId string) (*models.CartItem, *serviceerror.ServiceError)
	DeleteItem(productId, userId string) *serviceerror.ServiceError
}

type cartSrv struct {
	loggerSrv    LogSrv
	message      logger.Messages
	validatorSrv ValidationSrv
	timeSrv      TimeService
	repo         repo.CartRepo
	userRepo     repo.UserRepo
	productRepo  repo.ProductRepo
}

func (ch *cartSrv) Validate(req any) error {
	err := ch.validatorSrv.Validate(req)
	if err != nil {
		ch.loggerSrv.Error(ch.message.ValidationError(req, err))
		return err
	}

	return nil
}

func (ch *cartSrv) GetCart(userId string) (*models.Cart, *serviceerror.ServiceError) {
	items, err := ch.repo.GetCart(userId)
	if err != nil {
		ch.loggerSrv.Error(ch.message.GetCartRepoErrror(userId, err))
		return nil, serviceerror.Internal(err)
	}

	ch.loggerSrv.Info(ch.message.GetCartSuccess(items))
	return items, nil
}

func (ch *cartSrv) ClearCart(userId string) *serviceerror.ServiceError {
	err := ch.repo.ClearCart(userId)
	if err != nil {
		return serviceerror.Internal(err)
	}
	return nil
}

func (ch *cartSrv) AddItem(req *forms.CartItem, productId, userId string) (*models.CartItem, *serviceerror.ServiceError) {
	if err := ch.Validate(req); err != nil {
		return nil, serviceerror.Validating(err)
	}

	cart, err := ch.repo.GetCart(userId)
	if err != nil {
		return nil, serviceerror.Internal(err)
	}

	product, err := ch.productRepo.GetId(productId)
	if err != nil {
		return nil, serviceerror.NotFoundOrInternal(err)
	}

	if product.Product.UserId == userId {
		return nil, serviceerror.Forbidden("You cannot buy your own product")
	}

	var item models.CartItem

	item.CartId = cart.Id
	item.ProductId = productId
	item.Product = product.Product
	item.Quantity = req.Quantity
	item.DateCreated = ch.timeSrv.CurrentTime()
	item.DateUpdated = ch.timeSrv.CurrentTime()

	result, err := ch.repo.AddItem(&item, userId)
	if err != nil {
		return nil, serviceerror.NotFoundOrInternal(err)
	}

	return result, nil
}

func (ch *cartSrv) GetItem(productId, userId string) (*models.CartItem, *serviceerror.ServiceError) {
	item, err := ch.repo.GetItem(productId, userId)
	if err != nil {
		return nil, serviceerror.NotFoundOrInternal(err)
	}

	return item, nil
}

func (ch *cartSrv) DeleteItem(productId, userId string) *serviceerror.ServiceError {
	err := ch.repo.DeleteItem(productId, userId)
	if err != nil {
		return serviceerror.Internal(err)
	}

	return nil
}

func NewCartService(repo repo.CartRepo, loggerSrv LogSrv, validatorSrv ValidationSrv, timeSrv TimeService, userRepo repo.UserRepo, productRepo repo.ProductRepo) CartService {
	return &cartSrv{loggerSrv: loggerSrv, validatorSrv: validatorSrv, timeSrv: timeSrv, repo: repo, userRepo: userRepo, productRepo: productRepo}
}
