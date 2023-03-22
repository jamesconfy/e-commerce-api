package routes

import (
	"e-commerce/cmd/handlers/cartHandler"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service/cartService"
	"e-commerce/internal/service/tokenService"

	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.RouterGroup, cartSrv cartService.CartService, tokenSrv tokenService.TokenSrv) {
	handler := cartHandler.NewCartHandler(cartSrv)
	jwt := middleware.NewJWTMiddleWare(tokenSrv)
	cart := router.Group("/carts")
	cart.Use(jwt.ValidateJWT())
	{
		cart.POST("/add", handler.AddToCart)
	}
}
