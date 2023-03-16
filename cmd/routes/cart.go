package routes

import (
	"e-commerce/cmd/handlers/userHandler"
	"e-commerce/internal/service/tokenService"
	"e-commerce/internal/service/userService"

	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.RouterGroup, cartSrv userService.UserService, tokenSrv tokenService.TokenSrv) {
	handler := userHandler.NewUserHandler(cartSrv)
	user := router.Group("/cart")
	{
		user.POST("/add-to-cart", handler.CreateUser)
		user.POST("/remove-from-cart", handler.LoginUser)
		// user.POST("/reset-password", handler.ResetPassword)
		// user.POST("/reset-password/validate-token", handler.ValidateToken)
		// user.PATCH("/reset-password/change-password", handler.ChangePassword)
	}
}
