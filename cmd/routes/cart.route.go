package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service"

	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.RouterGroup, cartSrv service.CartService, tokenSrv service.TokenSrv) {
	handler := handler.NewCartHandler(cartSrv)
	jwt := middleware.NewJWTMiddleWare(tokenSrv)
	cart := router.Group("/carts")
	cart.Use(jwt.CheckJWT())
	{
		cart.POST("", handler.AddToCart)
		cart.PUT("/update", handler.UpdateCart)
		cart.GET("/:itemId", handler.GetItem)
		cart.PATCH("/:itemId", handler.EditItem)
		cart.DELETE("/:itemId", handler.DeleteItem)
	}
}
