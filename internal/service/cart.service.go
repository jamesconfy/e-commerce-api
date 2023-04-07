package service

import (
	"e-commerce/internal/logger"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/se"
)

type CartService interface {
	Get(userId string) (*models.Cart, *se.ServiceError)
	Clear(userId string) *se.ServiceError
}

type cartSrv struct {
	cartRepo     repo.CartRepo
	userRepo     repo.UserRepo
	productRepo  repo.ProductRepo
	loggerSrv    LogSrv
	message      logger.Messages
	validatorSrv ValidationSrv
}

func (ch *cartSrv) Validate(req any) error {
	err := ch.validatorSrv.Validate(req)
	if err != nil {
		ch.loggerSrv.Error(ch.message.ValidationError(req, err))
		return err
	}

	return nil
}

func (ch *cartSrv) Get(userId string) (*models.Cart, *se.ServiceError) {
	items, err := ch.cartRepo.Get(userId)
	if err != nil {
		ch.loggerSrv.Error(ch.message.GetCartRepoErrror(userId, err))
		return nil, se.Internal(err)
	}

	ch.loggerSrv.Info(ch.message.GetCartSuccess(items))
	return items, nil
}

func (ch *cartSrv) Clear(userId string) *se.ServiceError {
	err := ch.cartRepo.Clear(userId)
	if err != nil {
		ch.loggerSrv.Error(ch.message.ClearCartRepoError(userId, err))
		return se.Internal(err)
	}

	ch.loggerSrv.Error(ch.message.ClearCartSuccess(userId))
	return nil
}

func NewCartService(cartRepo repo.CartRepo, userRepo repo.UserRepo, productRepo repo.ProductRepo, loggerSrv LogSrv, validatorSrv ValidationSrv) CartService {
	return &cartSrv{cartRepo: cartRepo, userRepo: userRepo, productRepo: productRepo, loggerSrv: loggerSrv, validatorSrv: validatorSrv}
}
