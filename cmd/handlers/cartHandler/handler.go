package cartHandler

import (
	"e-commerce/internal/models/cartModels"
	"e-commerce/internal/models/responseModels"
	"e-commerce/internal/service/cartService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHanlder interface {
	AddToCart(c *gin.Context)
	GetItem(c *gin.Context)
	DeleteItem(c *gin.Context)
}

type cartHandler struct {
	cartSrv cartService.CartService
}

func (ch *cartHandler) AddToCart(c *gin.Context) {
	var req cartModels.AddToCartReq

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
		return
	}

	if err := ch.cartSrv.Validate(req); err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	userId := c.GetString("userId")
	user, err := ch.cartSrv.GetUser(userId)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	product, err := ch.cartSrv.CheckProduct(req.ProductId)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	if product.UserId == user.UserId {
		c.AbortWithStatusJSON(http.StatusForbidden, responseModels.BuildErrorResponse(http.StatusForbidden, "You cannot buy your own product", nil, nil))
		return
	}

	err = ch.cartSrv.CheckIfProductInCart(product.ProductId, user.CartId)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	item, err := ch.cartSrv.AddToCart(&req, product, user)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Item added to cart successfully", item, nil))
}

func (ch *cartHandler) GetItem(c *gin.Context) {
	itemId := c.Param("cart_item_id")
	if itemId == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, responseModels.BuildErrorResponse(http.StatusNotFound, "No cart item id was provided", nil, nil))
		return
	}

	userId := c.GetString("userId")
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responseModels.BuildErrorResponse(http.StatusUnauthorized, "You need to be logged in to access this resource", nil, nil))
		return
	}

	item, err := ch.cartSrv.GetItem(itemId)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	if item.UserId != userId {
		c.AbortWithStatusJSON(http.StatusForbidden, responseModels.BuildErrorResponse(http.StatusForbidden, "You are not authorized to view this product", nil, nil))
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Item fetched successfully", item, nil))
}

func (ch *cartHandler) DeleteItem(c *gin.Context) {
	itemId := c.Param("cart_item_id")
	if itemId == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, responseModels.BuildErrorResponse(http.StatusNotFound, "No cart item id was provided", nil, nil))
		return
	}

	userId := c.GetString("userId")
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responseModels.BuildErrorResponse(http.StatusUnauthorized, "You need to be logged in to access this resource", nil, nil))
		return
	}

	item, err := ch.cartSrv.GetItem(itemId)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	if item.UserId != userId {
		c.AbortWithStatusJSON(http.StatusForbidden, responseModels.BuildErrorResponse(http.StatusForbidden, "You are not authorized to delete this resource", nil, nil))
		return
	}

	err = ch.cartSrv.DeleteItem(itemId)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Item deleted successfully", item, nil))
}

func NewCartHandler(cartSrv cartService.CartService) CartHanlder {
	return &cartHandler{cartSrv: cartSrv}
}
