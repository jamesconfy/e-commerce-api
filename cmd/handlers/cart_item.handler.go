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

// Add Cart Item godoc
// @Summary	Add cart item route
// @Description	Add cart item to user cart
// @Tags	Item
// @Produce	json
// @Accept	json
// @Param	request	body	forms.CartItem	true	"Cart item"
// @Success	200  {object}	response.SuccessMessage{data=models.Item}
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/items [post]
// @Security ApiKeyAuth
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

// Get Items godoc
// @Summary	Get items route
// @Description	Get all items in a user cart
// @Tags	Item
// @Produce	json
// @Success	200  {object}	response.SuccessMessage{data=models.CartItem}
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/items [get]
// @Security ApiKeyAuth
func (ci *cartItemHandler) GetItems(c *gin.Context) {
	userId := c.GetString("userId")

	items, err := ci.cartItemSrv.GetItems(userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Items gotten successfully", items)
}

// Get Item godoc
// @Summary	Get item route
// @Description	Get an item in a user cart using product id
// @Tags	Item
// @Produce	json
// @Param	productId	path	string	true	"Product id"
// @Success	200  {object}	response.SuccessMessage{data=models.Item}
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/items/:productId [get]
// @Security ApiKeyAuth
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

	response.Success(c, "Item gotten successfully", item)
}

// Delete Item godoc
// @Summary	Delete item route
// @Description	Delete an item in a user cart using product id
// @Tags	Item
// @Produce	json
// @Param	productId	path	string	true	"Product id"
// @Success	200  {object}	response.SuccessMessage
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/items/:productId [delete]
// @Security ApiKeyAuth
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

	response.Success202(c, "Item deleted successfully")
}

func NewCartItemHandler(cartItemSrv service.CartItemService) CartItemHanlder {
	return &cartItemHandler{cartItemSrv: cartItemSrv}
}
