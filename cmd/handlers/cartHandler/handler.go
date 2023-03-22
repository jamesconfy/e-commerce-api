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
}

type cartHandler struct {
	cartSrv cartService.CartService
}

func (ch *cartHandler) AddToCart(c *gin.Context) {
	var req cartModels.AddToCart

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
		return
	}

	
}

func NewCartHandler(cartSrv cartService.CartService) CartHanlder {
	return &cartHandler{cartSrv: cartSrv}
}
