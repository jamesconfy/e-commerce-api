package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service"

	"github.com/gin-gonic/gin"
)

func CartItemRoute(router *gin.RouterGroup, cartItemSrv service.CartItemService, authSrv service.AuthService) {
	handler := handler.NewCartItemHandler(cartItemSrv)
	auth := middleware.Authentication(authSrv)
	item := router.Group("/items")

	item.Use(auth.CheckJWT())
	{
		item.POST("/", handler.Add)
		item.GET("/", handler.GetItems)
		item.GET("/:productId", handler.Get)
		item.DELETE("/:productId", handler.Delete)
	}
}
