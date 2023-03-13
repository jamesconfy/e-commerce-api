package routes

import (
	"github.com/gin-gonic/gin"
)

func ProductRoutes(v1 *gin.RouterGroup) {
	product := v1.Group("/products")
	// product.Use(middleware.ValidateToken)
	{
		product.POST("")
	}
}
