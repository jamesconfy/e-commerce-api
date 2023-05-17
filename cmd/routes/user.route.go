package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup, userSrv service.UserService, authSrv service.AuthService, store *persistence.InMemoryStore) {
	handler := handler.NewUserHandler(userSrv)
	jwt := middleware.Authentication(authSrv)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", func(ctx *gin.Context) {
			defer func() {
				store.Delete(cache.CreateKey("/api/v1/users"))
			}()
			handler.Create(ctx)
		})
		auth.POST("/login", handler.Login)
	}

	user := router.Group("/users")
	{
		user.GET("/:userId", cache.CachePage(store, time.Duration(time.Hour*1), func(ctx *gin.Context) {
			handler.GetById(ctx)
		}))
		user.GET("", cache.CachePage(store, time.Duration(time.Hour*1), func(ctx *gin.Context) {
			handler.GetAll(ctx)
		}))
	}

	authJWT := router.Group("/auth")
	authJWT.Use(jwt.CheckJWT())
	{
		authJWT.POST("/logout", handler.Logout)
		authJWT.DELETE("/clear", handler.ClearAuth)
	}

	userJWT := router.Group("/users")
	userJWT.Use(jwt.CheckJWT())
	{
		userJWT.GET("/profile", handler.Get)
		userJWT.PATCH("/profile", handler.Edit)
		userJWT.DELETE("/profile", handler.Delete)
	}
}
