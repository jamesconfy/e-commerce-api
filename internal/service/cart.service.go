package service

// import (
// 	"e-commerce/internal/models/cartModels"
// 	"e-commerce/internal/models/errorModels"
// 	"e-commerce/internal/models/productModels"
// 	"e-commerce/internal/models/responseModels"
// 	"e-commerce/internal/models/userModels"
// 	repo "e-commerce/internal/repository"
// 	"net/http"

// 	"github.com/google/uuid"
// )

// type CartService interface {
// 	Validate(req any) *responseModels.ResponseMessage
// 	CheckProduct(productId string) (*productModels.GetProductRes, *responseModels.ResponseMessage)
// 	GetUser(userId string) (*userModels.GetByIdRes, *responseModels.ResponseMessage)
// 	AddToCart(req *cartModels.AddToCartReq, product *productModels.GetProductRes, user *userModels.GetByIdRes) (*cartModels.AddToCartRes, *responseModels.ResponseMessage)
// 	CheckIfProductInCart(productId, cartId string) *responseModels.ResponseMessage
// 	GetItem(itemId string) (*cartModels.GetItemByIdRes, *responseModels.ResponseMessage)
// 	EditItem(req *cartModels.EditItemReq, item *cartModels.GetItemByIdRes) (*cartModels.EditItemRes, *responseModels.ResponseMessage)
// 	DeleteItem(itemId string) *responseModels.ResponseMessage
// }

// type cartSrv struct {
// 	loggerSrv    LogSrv
// 	validatorSrv ValidationSrv
// 	timeSrv      TimeService
// 	repo         repo.CartRepo
// }

// func (ch *cartSrv) Validate(req any) *responseModels.ResponseMessage {
// 	err := ch.validatorSrv.Validate(req)
// 	if err != nil {
// 		e := errorModels.NewValidatingError(err)
// 		return responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", e, nil)
// 	}

// 	return nil
// }

// func (ch *cartSrv) CheckProduct(productId string) (*productModels.GetProductRes, *responseModels.ResponseMessage) {
// 	product, err := ch.repo.GetProduct(productId)
// 	if product != nil && err != nil {
// 		return product, responseModels.BuildErrorResponse(http.StatusNotFound, "Product not found", err, nil)
// 	}

// 	if product == nil && err != nil {
// 		return nil, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Internal server error", err, nil)
// 	}

// 	return product, nil
// }

// func (ch *cartSrv) GetUser(userId string) (*userModels.GetByIdRes, *responseModels.ResponseMessage) {
// 	user, err := ch.repo.GetUser(userId)
// 	if user != nil && err != nil {
// 		return user, responseModels.BuildErrorResponse(http.StatusNotFound, "User not found", err, nil)
// 	}

// 	if user == nil && err != nil {
// 		return nil, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Internal server error", err, nil)
// 	}

// 	return user, nil
// }

// func (ch *cartSrv) AddToCart(req *cartModels.AddToCartReq, product *productModels.GetProductRes, user *userModels.GetByIdRes) (*cartModels.AddToCartRes, *responseModels.ResponseMessage) {
// 	req.ItemId = uuid.New().String()
// 	req.Price = product.Price * float64(req.Quantity)
// 	req.CartId = user.CartId
// 	req.DateCreated = ch.timeSrv.CurrentTime()
// 	req.DateUpdated = ch.timeSrv.CurrentTime()

// 	err := ch.repo.Add(req)
// 	if err != nil {
// 		return nil, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Internal server error", err, nil)
// 	}

// 	return &cartModels.AddToCartRes{
// 		CartItemId:  req.ItemId,
// 		CartId:      req.CartId,
// 		ProductId:   req.ProductId,
// 		UserId:      user.UserId,
// 		Quantity:    req.Quantity,
// 		Price:       req.Price,
// 		DateCreated: req.DateCreated,
// 		DateUpdated: req.DateUpdated,
// 	}, nil
// }

// func (ch *cartSrv) CheckIfProductInCart(productId, cartId string) *responseModels.ResponseMessage {
// 	err := ch.repo.CheckIfProductInCart(productId, cartId)
// 	if err != nil {
// 		return responseModels.BuildErrorResponse(http.StatusConflict, "Product already in cart, you can either remove it or change the quantity", err, nil)
// 	}

// 	return nil
// }

// func (ch *cartSrv) GetItem(itemId string) (*cartModels.GetItemByIdRes, *responseModels.ResponseMessage) {
// 	item, err := ch.repo.GetItem(itemId)
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

// func (ch *cartSrv) DeleteItem(itemId string) *responseModels.ResponseMessage {
// 	err := ch.repo.DeleteItem(itemId)
// 	if err != nil {
// 		return responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when removing item from cart", err, nil)
// 	}

// 	return nil
// }

// func NewCartService(repo repo.CartRepo, loggerSrv LogSrv, validatorSrv ValidationSrv, timeSrv TimeService) CartService {
// 	return &cartSrv{loggerSrv: loggerSrv, validatorSrv: validatorSrv, timeSrv: timeSrv, repo: repo}
// }

// // Auxillary Function
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
