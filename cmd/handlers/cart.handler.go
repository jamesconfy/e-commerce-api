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

func (ch *cartHandler) Get(c *gin.Context) {
	userId := c.GetString("userId")

	carts, err := ch.cartSrv.Get(userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Cart gotten successfully", carts)
}

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
