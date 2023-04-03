package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(v1 *gin.RouterGroup, productSrv service.ProductService, tokenSrv service.TokenSrv) {
	handler := handler.NewProductHandler(productSrv)
	auth := middleware.Authentication(tokenSrv)

	product := v1.Group("/products")
	product.Use(auth.CheckJWT())
	{
		product.POST("", handler.Add)
		product.GET("", handler.GetAll)
		product.GET("/:product_id", handler.Get)
		product.PATCH("/:product_id", handler.Edit)
		product.DELETE("/:product_id", handler.Delete)
		product.POST("/:product_id/ratings", handler.AddRating)
	}
}
