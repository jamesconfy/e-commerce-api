package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service"

	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.RouterGroup, cartSrv service.CartService, tokenSrv service.AuthSrv) {
	handler := handler.NewCartHandler(cartSrv)
	auth := middleware.Authentication(tokenSrv)
	cart := router.Group("/carts")

	cart.Use(auth.CheckJWT())
	{
		cart.GET("", handler.Get)
		cart.DELETE("", handler.Clear)
	}
}
