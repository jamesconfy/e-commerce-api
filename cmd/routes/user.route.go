package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, userSrv service.UserService, authSrv service.AuthService) {
	handler := handler.NewUserHandler(userSrv)
	user := router.Group("/users")
	{
		user.POST("/signup", handler.Create)
		user.POST("/login", handler.Login)
		user.GET("/:userId", handler.GetById)
		user.GET("/all", handler.GetAll)
	}

	user1 := router.Group("/users")
	jwt := middleware.Authentication(authSrv)
	user1.Use(jwt.CheckJWT())
	{
		user1.GET("/profile", handler.Get)
		user1.PATCH("/profile", handler.Edit)
		user1.DELETE("/profile", handler.Delete)
		user1.POST("/logout", handler.Logout)
	}
}
