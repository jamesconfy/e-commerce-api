package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/logger"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/serviceerror"
	"fmt"

	"github.com/google/uuid"
)

type CheckoutService interface {
	Add(req *forms.Checkout, userId string) (*models.Checkout, *serviceerror.ServiceError)
}

type checkoutSrv struct {
	repo        repo.CheckoutRepo
	cartRepo    repo.CartRepo
	loggerSrv   LogSrv
	validateSrv ValidationSrv
	message     logger.Messages
}

func (co *checkoutSrv) Validate(req any) error {
	if err := co.validateSrv.Validate(req); err != nil {
		co.loggerSrv.Error(co.message.ValidationError(req, err))
		return err
	}

	return nil
}

func (co *checkoutSrv) Add(req *forms.Checkout, userId string) (*models.Checkout, *serviceerror.ServiceError) {
	if err := co.Validate(&req); err != nil {
		return nil, serviceerror.Validating(err)
	}

	carts, err := co.cartRepo.GetCart(userId)
	if err != nil {
		return nil, serviceerror.NotFoundOrInternal(err)
	}

	defer co.addToDatabase(carts, req)

	return nil, nil
}

func NewCheckoutService(cartRepo repo.CartRepo) CheckoutService {
	return &checkoutSrv{cartRepo: cartRepo}
}

// Auxillary Function
func (co *checkoutSrv) addToDatabase(carts *models.Cart, req *forms.Checkout) {
	for _, item := range carts.Items {
		go func(item *models.CartItem) {
			var checkout *models.Checkout

			checkout.Mutex.Lock()
			checkout.Id = uuid.New().String()
			checkout.CartId = item.CartId
			checkout.Quantity = item.Quantity
			checkout.ProductId = item.ProductId
			checkout.PaymentMethod = req.PaymentMethod

			_, err := co.repo.Add(checkout)
			if err != nil {
				co.loggerSrv.Warning(fmt.Sprintf("Error when adding product to database || Id: %v || CartId: %v || ProductId: %v || Error: %v", checkout.Id, checkout.CartId, checkout.ProductId, err))
			}

			checkout.Mutex.Unlock()
		}(item)
	}
}
