package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(v1 *gin.RouterGroup, productSrv service.ProductService, tokenSrv service.TokenSrv) {
	handler := handler.NewProductHandler(productSrv)
	jwtMiddleWare := middleware.NewJWTMiddleWare(tokenSrv)

	product := v1.Group("/products")
	product.Use(jwtMiddleWare.CheckJWT())
	{
		product.POST("", handler.AddProduct)
		product.GET("", handler.GetProducts)
		product.GET("/:product_id", handler.GetProduct)
		product.PATCH("/:product_id", handler.EditProduct)
		product.DELETE("/:product_id", handler.DeleteProduct)
		product.POST("/:product_id/ratings", handler.AddRating)
	}
}
