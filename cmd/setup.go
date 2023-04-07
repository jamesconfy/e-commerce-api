package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"e-commerce/cmd/middleware"
	route "e-commerce/cmd/routes"
	mysql "e-commerce/internal/database"
	"e-commerce/internal/logger"
	repo "e-commerce/internal/repository"
	service "e-commerce/internal/service"
	"e-commerce/utils"

	_ "e-commerce/docs"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() {
	// config, err := utils.LoadConfig("./")
	// if err != nil {
	// 	log.Println("Error loading configurations: ", err)
	// }
	// utils.MyConfig.ADDR

	addr := utils.AppConfig.ADDR
	if addr == "" {
		addr = "8000"
	}

	dsn := utils.AppConfig.DATA_SOURCE_NAME
	if dsn == "" {
		log.Println("DSN cannot be empty")
	}
	fmt.Println(dsn)

	secret := utils.AppConfig.SECRET_KEY_TOKEN
	if secret == "" {
		log.Println("Please provide a secret key token")
	}

	// host := utils.AppConfig.HOST
	// if host == "" {
	// 	log.Println("Please provide an email host name")
	// }

	// port := utils.AppConfig.PORT
	// if port == "" {
	// 	log.Println("Please provide an email port")
	// }

	// passwd := utils.AppConfig.PASSWD
	// if passwd == "" {
	// 	log.Println("Please provide an email password")
	// }

	// email := utils.AppConfig.EMAIL
	// if email == "" {
	// 	log.Println("Please provide an email address")
	// }

	connection, err := mysql.NewMySQLServer(dsn)
	if err != nil {
		log.Println("Error Connecting to DB: ", err)
		return
	}
	defer connection.Close()
	conn := connection.GetConn()

	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.New())
	gin.DisableConsoleColor()

	router := gin.New()
	router.SetTrustedProxies(nil)

	v1 := router.Group("/api/v1")
	v1.Use(gin.Logger())
	v1.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// User Repository
	userRepo := repo.NewUserRepo(conn)

	// Product Repository
	productRepo := repo.NewProductRepo(conn)

	// Token Repository
	authRepo := repo.NewAuthRepo(conn)

	// Cart Repository
	cartRepo := repo.NewCartRepo(conn)

	// Cart Item Repository
	cartItemRepo := repo.NewCartItemRepo(conn)

	// Logger Service
	loggerSrv := service.NewLoggerService()

	// Email Service
	emailSrv := service.NewEmailService("email", "passwd", "host", "port")

	// Validation Service
	validatorSrv := service.NewValidationService()

	// Cryptography Service
	cryptoSrv := service.NewCryptoService()

	// Auth Service
	authSrv := service.NewAuthService(authRepo, secret, loggerSrv)

	// Home Service
	homeSrv := service.NewHomeService(loggerSrv)

	// User Service
	userSrv := service.NewUserService(userRepo, authRepo, cartRepo, validatorSrv, cryptoSrv, authSrv, emailSrv, loggerSrv)

	// Product Service
	productSrv := service.NewProductService(productRepo, validatorSrv, loggerSrv)

	// Cart Service
	cartSrv := service.NewCartService(cartRepo, userRepo, productRepo, loggerSrv, validatorSrv)

	// Cart Item Service
	cartItemSrv := service.NewCartItemService(cartItemRepo, cartRepo, userRepo, productRepo, loggerSrv, validatorSrv)

	// Routes
	route.HomeRoute(v1, homeSrv)
	route.UserRoute(v1, userSrv, authSrv)
	route.ProductRoutes(v1, productSrv, authSrv)
	route.CartRoute(v1, cartSrv, authSrv)
	route.CartItemRoute(v1, cartItemSrv, authSrv)
	route.ErrorRoute(router)

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().Do(func() {
		conn.Ping()
	})

	// Documentation
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + addr)
}
