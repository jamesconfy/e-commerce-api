package routes

import (
	producthandler "e-commerce/cmd/handlers/productHandler"
	"e-commerce/internal/service/tokenService"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(v1 *gin.RouterGroup, tokenSrv tokenService.TokenSrv) {
	handler := producthandler.NewProductHandler()
	//jwtMiddleWare := middleware.NewJWTMiddleWare(tokenSrv)

	product := v1.Group("/products")
	//product.Use(jwtMiddleWare.ValidateJWT())
	{
		product.GET("", handler.Product)
	}
}
