package handler

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/response"
	se "e-commerce/internal/se"
	"e-commerce/internal/service"

	"github.com/gin-gonic/gin"
)

type CartItemHanlder interface {
	Add(c *gin.Context)
	GetItems(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
}

type cartItemHandler struct {
	cartItemSrv service.CartItemService
}

func (ci *cartItemHandler) Add(c *gin.Context) {
	var req forms.CartItem

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	userId := c.GetString("userId")

	item, err := ci.cartItemSrv.Add(&req, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Item added successfully", item)
}

func (ci *cartItemHandler) GetItems(c *gin.Context) {
	userId := c.GetString("userId")

	items, err := ci.cartItemSrv.GetItems(userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Items gotten successfully", items)
}

func (ci *cartItemHandler) Get(c *gin.Context) {
	productId := c.Param("productId")
	if productId == "" {
		response.Error(c, *se.BadRequest("No product id was provided"))
		return
	}

	userId := c.GetString("userId")

	item, err := ci.cartItemSrv.Get(productId, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Item fetched successfully", item)
}

func (ci *cartItemHandler) Delete(c *gin.Context) {
	productId := c.Param("productId")
	if productId == "" {
		response.Error(c, *se.BadRequest("No product id was provided"))
		return
	}

	userId := c.GetString("userId")

	err := ci.cartItemSrv.Delete(productId, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success202(c, "Product deleted successfully")
}

func NewCartItemHandler(cartItemSrv service.CartItemService) CartItemHanlder {
	return &cartItemHandler{cartItemSrv: cartItemSrv}
}
