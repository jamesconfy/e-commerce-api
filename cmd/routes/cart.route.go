package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service"

	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.RouterGroup, cartSrv service.CartService, tokenSrv service.TokenSrv) {
	handler := handler.NewCartHandler(cartSrv)
	auth := middleware.Authentication(tokenSrv)
	cart := router.Group("/carts")

	cart.Use(auth.CheckJWT())
	{
		cart.GET("", handler.GetCart)
		cart.DELETE("", handler.ClearCart)
		cart.POST("/items/:productId", handler.AddItem)
		// cart.PUT("/update", handler.AddItem)
		cart.GET("/items/:productId", handler.GetItem)
		// cart.PATCH("/:itemId", handler)
		cart.DELETE("/items/:productId", handler.DeleteItem)
	}
}
