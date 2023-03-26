package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service/cartService"
	"e-commerce/internal/service/tokenService"

	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.RouterGroup, cartSrv cartService.CartService, tokenSrv tokenService.TokenSrv) {
	handler := handler.NewCartHandler(cartSrv)
	jwt := middleware.NewJWTMiddleWare(tokenSrv)
	cart := router.Group("/carts")
	cart.Use(jwt.ValidateJWT())
	{
		cart.POST("", handler.AddToCart)
		cart.PUT("/update", handler.UpdateCart)
		cart.GET("/:itemId", handler.GetItem)
		cart.PATCH("/:itemId", handler.EditItem)
		cart.DELETE("/:itemId", handler.DeleteItem)
	}
}
