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
	user := router.Group("/users")
	{
		user.POST("/signup", func(ctx *gin.Context) {
			defer func() {
				store.Delete(cache.CreateKey("/api/v1/users"))
			}()
			handler.Create(ctx)
		})
		user.POST("/login", handler.Login)
		user.GET("/:userId", cache.CachePage(store, time.Duration(time.Hour*1), func(ctx *gin.Context) {
			handler.GetById(ctx)
		}))
		user.GET("", cache.CachePage(store, time.Duration(time.Hour*1), func(ctx *gin.Context) {
			handler.GetAll(ctx)
		}))
	}

	user1 := router.Group("/users")
	jwt := middleware.Authentication(authSrv)
	user1.Use(jwt.CheckJWT())
	{
		user1.GET("/profile", handler.Get)
		user1.PATCH("/profile", handler.Edit)
		user1.DELETE("/profile", handler.Delete)
		user1.POST("/logout", handler.Logout)
		user1.DELETE("/profile/clear", handler.ClearAuth)
	}
}
