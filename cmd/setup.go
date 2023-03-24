package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"e-commerce/cmd/middleware"
	"e-commerce/cmd/routes"
	"e-commerce/internal/Repository/cartRepo"
	"e-commerce/internal/Repository/productRepo"
	"e-commerce/internal/Repository/tokenRepo"
	"e-commerce/internal/Repository/userRepo"
	mysql "e-commerce/internal/database"
	"e-commerce/internal/logger"
	"e-commerce/internal/service/cartService"
	"e-commerce/internal/service/cryptoService"
	"e-commerce/internal/service/emailService"
	"e-commerce/internal/service/homeService"
	"e-commerce/internal/service/loggerService"
	"e-commerce/internal/service/productService"
	"e-commerce/internal/service/timeService"
	"e-commerce/internal/service/tokenService"
	"e-commerce/internal/service/userService"
	validationService "e-commerce/internal/service/validatorService"
	"e-commerce/utils"

	_ "e-commerce/docs"

	"github.com/gin-gonic/gin"
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
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.NewLogger())
	gin.DisableConsoleColor()

	router := gin.New()
	router.SetTrustedProxies(nil)

	v1 := router.Group("/api/v1")
	v1.Use(gin.Logger())
	v1.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// User Repository
	userRepo := userRepo.NewMySqlUserRepo(conn)

	// Product Repository
	productRepo := productRepo.NewMySqlProductRepo(conn)

	// Token Repository
	tokenRepo := tokenRepo.NewMySqlTokenRepo(conn)

	// Cart Repository
	cartRepo := cartRepo.NewMySqlCartRepo(conn)

	// Message Utility
	message := utils.NewMessageUtils()

	// Logger Service
	loggerSrv := loggerService.NewLogger()

	// Time Service
	timeSrv := timeService.NewTimeService()

	// Email Service
	emailSrv := emailService.NewEmailSrv("email", "passwd", "host", "port")

	// Validation Service
	validatorSrv := validationService.NewValidationService()

	// Cryptography Service
	cryptoSrv := cryptoService.NewCryptoSrv()

	// Token Service
	tokenSrv := tokenService.NewTokenSrv(secret, loggerSrv, tokenRepo)

	// Home Service
	homeSrv := homeService.NewHomeSrv(loggerSrv)

	// User Service
	userSrv := userService.NewUserSrv(userRepo, validatorSrv, cryptoSrv, tokenSrv, emailSrv, loggerSrv, timeSrv, message)

	// Product Service
	productSrv := productService.NewProductService(productRepo, validatorSrv, loggerSrv, timeSrv, message)

	// Cart Service
	cartSrv := cartService.NewCartService(cartRepo, loggerSrv, validatorSrv, timeSrv)

	// Routes
	routes.HomeRoute(v1, homeSrv)
	routes.UserRoute(v1, userSrv, tokenSrv)
	routes.ProductRoutes(v1, productSrv, tokenSrv)
	routes.CartRoute(v1, cartSrv, tokenSrv)
	routes.ErrorRoute(router)

	// Documentation
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + addr)
}
