package routes

import (
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service/cartService"
	"e-commerce/internal/service/tokenService"

	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.RouterGroup, cartSrv cartService.CartService, tokenSrv tokenService.TokenSrv) {
	// //handler := .NewUserHandler(cartSrv)
	jwt := middleware.NewJWTMiddleWare(tokenSrv)
	cart := router.Group("/cart")
	cart.Use(jwt.ValidateJWT())
	// {
	// 	user.POST("/add-item", handler.CreateUser)
	// 	user.POST("/remove-item", handler.LoginUser)
	// 	// user.POST("/reset-password", handler.ResetPassword)
	// 	// user.POST("/reset-password/validate-token", handler.ValidateToken)
	// 	// user.PATCH("/reset-password/change-password", handler.ChangePassword)
	// }
}
