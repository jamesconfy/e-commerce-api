package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service"

	"github.com/gin-gonic/gin"
)

func CheckoutRoutes(v1 *gin.RouterGroup, checkoutSrv service.CheckoutService, tokenSrv service.AuthSrv) {
	_ = handler.NewCheckoutHandler(checkoutSrv)
	auth := middleware.Authentication(tokenSrv)
	checkout := v1.Group("/checkout")

	checkout.Use(auth.CheckJWT())
	{

	}
}
