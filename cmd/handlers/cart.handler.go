package handler

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/response"
	"e-commerce/internal/service"
	se "e-commerce/internal/serviceerror"

	"github.com/gin-gonic/gin"
)

type CartHanlder interface {
	// Cart
	GetCart(c *gin.Context)
	ClearCart(c *gin.Context)

	// Item
	AddItem(c *gin.Context)
	// UpdateCart(c *gin.Context)
	GetItem(c *gin.Context)
	DeleteItem(c *gin.Context)
}

type cartHandler struct {
	cartSrv service.CartService
}

func (ch *cartHandler) GetCart(c *gin.Context) {
	userId := c.GetString("userId")

	carts, err := ch.cartSrv.GetCart(userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Cart gotten successfully", carts)
}

func (ch *cartHandler) ClearCart(c *gin.Context) {
	userId := c.GetString("userId")

	err := ch.cartSrv.ClearCart(userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success202(c, "Cart cleared successfully")
}

func (ch *cartHandler) AddItem(c *gin.Context) {
	var req forms.CartItem

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	productId := c.Param("productId")
	userId := c.GetString("userId")

	item, err := ch.cartSrv.AddItem(&req, productId, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Item added successfully", item)
}

func (ch *cartHandler) GetItem(c *gin.Context) {
	productId := c.Param("productId")
	if productId == "" {
		response.Error(c, *se.BadRequest("No product id was provided"))
		return
	}

	userId := c.GetString("userId")

	item, err := ch.cartSrv.GetItem(productId, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Item fetched successfully", item)
}

func (ch *cartHandler) DeleteItem(c *gin.Context) {
	productId := c.Param("productId")
	if productId == "" {
		response.Error(c, *se.BadRequest("No product id was provided"))
		return
	}

	userId := c.GetString("userId")

	err := ch.cartSrv.DeleteItem(productId, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success202(c, "Product deleted successfully")
}

func NewCartHandler(cartSrv service.CartService) CartHanlder {
	return &cartHandler{cartSrv: cartSrv}
}
