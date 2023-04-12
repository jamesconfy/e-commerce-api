package handler

import (
	"e-commerce/internal/response"
	"e-commerce/internal/service"

	"github.com/gin-gonic/gin"
)

type CartHanlder interface {
	Get(c *gin.Context)
	Clear(c *gin.Context)
}

type cartHandler struct {
	cartSrv service.CartService
}

// Get Cart godoc
// @Summary	Get cart route
// @Description	Get a user's cart details
// @Tags	Cart
// @Produce	json
// @Success	200  {object}	response.SuccessMessage{data=models.Cart}
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/carts/ [get]
// @Security ApiKeyAuth
func (ch *cartHandler) Get(c *gin.Context) {
	userId := c.GetString("userId")

	carts, err := ch.cartSrv.Get(userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Cart gotten successfully", carts)
}

// Clear Cart godoc
// @Summary	Clear cart route
// @Description	Clear a user's cart i.e delete all items in the cart
// @Tags	Cart
// @Produce	json
// @Success	200  {object}	response.SuccessMessage
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/carts/ [delete]
// @Security ApiKeyAuth
func (ch *cartHandler) Clear(c *gin.Context) {
	userId := c.GetString("userId")

	err := ch.cartSrv.Clear(userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success202(c, "Cart cleared successfully")
}

func NewCartHandler(cartSrv service.CartService) CartHanlder {
	return &cartHandler{cartSrv: cartSrv}
}
