package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/logger"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/se"
)

type CartItemService interface {
	Add(req *forms.CartItem, userId string) (*models.Item, *se.ServiceError)
	Get(productId, userId string) (*models.Item, *se.ServiceError)
	GetItems(userId string) (*models.CartItem, *se.ServiceError)
	Delete(productId, userId string) *se.ServiceError
}

type cartItemSrv struct {
	cartItemRepo repo.CartItemRepo
	cartRepo     repo.CartRepo
	userRepo     repo.UserRepo
	productRepo  repo.ProductRepo
	loggerSrv    LogSrv
	message      logger.Messages
	validatorSrv ValidationSrv
}

func (ci *cartItemSrv) Validate(req any) error {
	err := ci.validatorSrv.Validate(req)
	if err != nil {
		ci.loggerSrv.Error(ci.message.ValidationError(req, err))
		return err
	}

	return nil
}

func (ci *cartItemSrv) Add(req *forms.CartItem, userId string) (*models.Item, *se.ServiceError) {
	if err := ci.Validate(req); err != nil {
		return nil, se.Validating(err)
	}

	cart, err := ci.cartRepo.Get(userId)
	if err != nil {
		ci.loggerSrv.Error(ci.message.GetCartRepoErrror(userId, err))
		return nil, se.Internal(err)
	}

	product, err := ci.productRepo.GetId(req.ProductId)
	if err != nil {
		ci.loggerSrv.Error(ci.message.GetProductRepoError(req.ProductId, err))
		return nil, se.NotFoundOrInternal(err, "product not found")
	}

	if product.Product.UserId == userId {
		ci.loggerSrv.Warning(ci.message.AddItemCompareUser(product.Product.UserId, userId))
		return nil, se.Forbidden("You cannot buy your own product")
	}

	var item models.Item

	item.ProductId = req.ProductId
	item.Product = product.Product
	item.Quantity = req.Quantity

	result, err := ci.cartItemRepo.Add(cart, &item)
	if err != nil {
		ci.loggerSrv.Error(ci.message.AddItemRepoError(req.ProductId, userId, err))
		return nil, se.NotFoundOrInternal(err, "item not found")
	}

	result.Product = product.Product
	ci.loggerSrv.Info(ci.message.AddItemSuccess(result))
	return result, nil
}

func (ci *cartItemSrv) Get(productId, userId string) (*models.Item, *se.ServiceError) {
	cart, err := ci.cartRepo.Get(userId)
	if err != nil {
		ci.loggerSrv.Error(ci.message.GetItemRepoError(productId, userId, err))
		return nil, se.NotFoundOrInternal(err, "cart not found")
	}

	product, err := ci.productRepo.GetId(productId)
	if err != nil {
		ci.loggerSrv.Error(ci.message.GetItemRepoError(productId, userId, err))
		return nil, se.NotFoundOrInternal(err, "product not found")
	}

	item, err := ci.cartItemRepo.Get(cart, productId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "item not found")
	}

	item.Product = product.Product

	ci.loggerSrv.Info(ci.message.GetItemSuccess(item))
	return item, nil
}

func (ci *cartItemSrv) GetItems(userId string) (*models.CartItem, *se.ServiceError) {
	cart, err := ci.cartRepo.Get(userId)
	if err != nil {
		// ci.loggerSrv.Error(ci.message.GetItemRepoError(productId, userId, err))
		return nil, se.NotFoundOrInternal(err, "cart not found")
	}

	cartItems, err := ci.cartItemRepo.GetItems(cart)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "cart items could not be fetched")
	}

	return cartItems, nil
}

func (ci *cartItemSrv) Delete(productId, userId string) *se.ServiceError {
	cart, err := ci.cartRepo.Get(userId)
	if err != nil {
		ci.loggerSrv.Error(ci.message.GetItemRepoError(productId, userId, err))
		return se.NotFoundOrInternal(err, "cart not found")
	}

	err = ci.cartItemRepo.Delete(cart, productId)
	if err != nil {
		ci.loggerSrv.Error(ci.message.DeleteItemRepoError(productId, userId, err))
		return se.Internal(err)
	}

	ci.loggerSrv.Error(ci.message.DeleteItemSuccess(productId, userId))
	return nil
}

func NewCartItemService(cartItemRepo repo.CartItemRepo, cartRepo repo.CartRepo, userRepo repo.UserRepo, productRepo repo.ProductRepo, loggerSrv LogSrv, validatorSrv ValidationSrv) CartItemService {
	return &cartItemSrv{cartItemRepo: cartItemRepo, cartRepo: cartRepo, userRepo: userRepo, productRepo: productRepo, loggerSrv: loggerSrv, validatorSrv: validatorSrv}
}
