package route

import (
	handler "e-commerce/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func ErrorRoute(router *gin.Engine) {
	handler := handler.NewErrorHandler()
	router.NoRoute(handler.Error404)

	router.NoMethod(handler.MethodNotAllowed)
}
