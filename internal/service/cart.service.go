package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/serviceerror"
)

type CartService interface {
	Add(req *forms.CartItem, userId string) (*models.CartItem, *serviceerror.ServiceError)
	GetCart(userId string) (*models.Cart, *serviceerror.ServiceError)
	// GetItem(itemId string) (*models.Cart, *responseModels.ResponseMessage)
	// EditItem(req *cartModels.EditItemReq, item *cartModels.GetItemByIdRes) (*cartModels.EditItemRes, *responseModels.ResponseMessage)
	Delete(productId, userId string) *serviceerror.ServiceError
}

type cartSrv struct {
	loggerSrv    LogSrv
	validatorSrv ValidationSrv
	timeSrv      TimeService
	repo         repo.CartRepo
	userRepo     repo.UserRepo
	productRepo  repo.ProductRepo
}

func (ch *cartSrv) Validate(req any) error {
	err := ch.validatorSrv.Validate(req)
	if err != nil {
		return err
	}

	return nil
}

func (ch *cartSrv) Add(req *forms.CartItem, userId string) (*models.CartItem, *serviceerror.ServiceError) {
	if err := ch.Validate(req); err != nil {
		return nil, serviceerror.Validating(err)
	}

	cartId, err := ch.userRepo.GetCartId(userId)
	if err != nil {
		return nil, serviceerror.Internal(err)
	}

	product, err := ch.productRepo.GetId(req.ProductId)
	if err != nil {
		return nil, serviceerror.NotFoundOrInternal(err)
	}

	var item models.CartItem

	item.CartId = *cartId
	item.Product = product.Product
	item.Quantity = req.Quantity
	item.DateCreated = ch.timeSrv.CurrentTime()
	item.DateUpdated = ch.timeSrv.CurrentTime()

	err = ch.repo.Add(&item)
	if err != nil {
		return nil, serviceerror.Internal(err)
	}

	return &item, nil
}

func (ch *cartSrv) GetCart(userId string) (*models.Cart, *serviceerror.ServiceError) {
	cartId, err := ch.userRepo.GetCartId(userId)
	if err != nil {
		return nil, serviceerror.Internal(err)
	}

	items, err := ch.repo.GetCart(*cartId)
	if err != nil {
		return nil, serviceerror.Internal(err)
	}

	return items, nil
}

// func (ch *cartSrv) GetItem(itemId string) (*cartModels.GetItemByIdRes, *responseModels.ResponseMessage) {
// 	item, err := ch.repo.GetId(itemId)
// 	if item != nil && err != nil {
// 		return item, responseModels.BuildErrorResponse(http.StatusNotFound, "Item not found", err, nil)
// 	}

// 	if item == nil && err != nil {
// 		return nil, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Internal server error", err, nil)
// 	}

// 	return item, nil
// }

// func (ch *cartSrv) EditItem(req *cartModels.EditItemReq, item *cartModels.GetItemByIdRes) (*cartModels.EditItemRes, *responseModels.ResponseMessage) {
// 	editItem := ch.updateItem(*req, item)

// 	err := ch.repo.Edit(editItem)
// 	if err != nil {
// 		return nil, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when updating cart item", err, nil)
// 	}

// 	return &cartModels.EditItemRes{
// 		CartItemId:  editItem.CartItemId,
// 		CartId:      item.CartId,
// 		UserId:      item.UserId,
// 		ProductId:   item.ProductId,
// 		Quantity:    editItem.Quantity,
// 		Price:       editItem.Price,
// 		DateUpdated: editItem.DateUpdated,
// 	}, nil
// }

func (ch *cartSrv) Delete(productId, userId string) *serviceerror.ServiceError {
	cartId, err := ch.userRepo.GetCartId(userId)
	if err != nil {
		return serviceerror.Internal(err)
	}

	err = ch.repo.Delete(productId, *cartId)
	if err != nil {
		return serviceerror.Internal(err)
	}

	return nil
}

func NewCartService(repo repo.CartRepo, loggerSrv LogSrv, validatorSrv ValidationSrv, timeSrv TimeService, userRepo repo.UserRepo, productRepo repo.ProductRepo) CartService {
	return &cartSrv{loggerSrv: loggerSrv, validatorSrv: validatorSrv, timeSrv: timeSrv, repo: repo, userRepo: userRepo, productRepo: productRepo}
}

// Auxillary Function
// func (ch *cartSrv) updateItem(req cartModels.EditItemReq, item *cartModels.GetItemByIdRes) *cartModels.EditItemReq {
// 	price := item.Price / float64(item.Quantity)

// 	if req.Quantity != 0 && req.Quantity != item.Quantity {
// 		item.Quantity = req.Quantity
// 	}

// 	item.Price = price * float64(item.Quantity)

// 	item.DateUpdated = ch.timeSrv.CurrentTime()

// 	return &cartModels.EditItemReq{
// 		CartItemId:  item.CartItemId,
// 		UserId:      item.UserId,
// 		Quantity:    item.Quantity,
// 		Price:       item.Price,
// 		DateUpdated: item.DateUpdated,
// 	}
// }
