package routes

import (
	producthandler "e-commerce/cmd/handlers/productHandler"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service/productService"
	"e-commerce/internal/service/tokenService"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(v1 *gin.RouterGroup, productSrv productService.ProductService, tokenSrv tokenService.TokenSrv) {
	handler := producthandler.NewProductHandler(productSrv)
	jwtMiddleWare := middleware.NewJWTMiddleWare(tokenSrv)

	product := v1.Group("/products")
	product.Use(jwtMiddleWare.ValidateJWT())
	{
		product.POST("", handler.AddProduct)
		product.GET("", handler.GetProducts)
		product.GET("/:product_id", handler.GetProduct)
		product.DELETE("/:product_id", handler.DeleteProduct)
		product.POST("/:product_id/ratings", handler.AddRating)
	}
}
