package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/internal/service"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, userSrv service.UserService, tokenSrv service.AuthService) {
	handler := handler.NewUserHandler(userSrv)
	user := router.Group("/users")
	{
		user.POST("/signup", handler.Create)
		user.POST("/login", handler.Login)
		user.GET("/:userId", handler.GetById)
		// user.POST("/reset-password", handler.ResetPassword)
		// user.POST("/reset-password/validate-token", handler.ValidateToken)
		// user.PATCH("/reset-password/change-password", handler.ChangePassword)
	}
}
