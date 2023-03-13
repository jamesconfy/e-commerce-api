package routes

import (
	"e-commerce/cmd/handlers/userHandler"
	"e-commerce/internal/service/userService"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, userSrv userService.UserService) {
	handler := userHandler.NewUserHandler(userSrv)
	user := router.Group("/users")
	{
		user.POST("", handler.CreateUser)
		user.POST("/login", handler.LoginUser)
		user.POST("/reset-password", handler.ResetPassword)
		user.POST("/reset-password/validate-token", handler.ValidateToken)
		user.PATCH("/reset-password/change-password", handler.ChangePassword)
	}
}
