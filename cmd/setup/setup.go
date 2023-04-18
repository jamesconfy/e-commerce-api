package cmd

import (
	"io"
	"log"
	"os"
	"time"

	"e-commerce/cmd/middleware"
	route "e-commerce/cmd/routes"
	db "e-commerce/internal/database"
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

var (
	addr           string
	dsn            string
	redis_addr     string
	redis_username string
	redis_password string
	mode           string
	secret         string
	cache          bool
)

func Setup() {
	mysqlDB, err := db.New(dsn)
	if err != nil {
		log.Println("Error Connecting to DB: ", err)
		return
	}
	defer mysqlDB.Close()
	conn := mysqlDB.Get()

	redisClient := db.NewRedisDB(redis_username, redis_password, redis_addr)
	if redisClient == nil {
		log.Println("Error when connecting to Redis")
		return
	}
	defer redisClient.Close()

	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.New())
	gin.DisableConsoleColor()

	router := gin.New()
	router.SetTrustedProxies(nil)

	v1 := router.Group("/api/v1")
	v1.Use(gin.Logger())
	v1.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// Redis Repository
	cacheRepo := repo.NewRedisCache(redisClient)

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
	loggerSrv := service.NewLoggerService("")

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

	// Check cache and implement it
	if cache && redisClient != nil {
		authSrv = service.NewCachedAuthService(authSrv, cacheRepo)
		userSrv = service.NewCachedUserService(userSrv, cacheRepo)
		cartItemSrv = service.NewCachedCartItemService(cartItemSrv, cacheRepo)
		cartSrv = service.NewCachedCartService(cartSrv, cacheRepo)
		productSrv = service.NewCachedProductService(productSrv, cacheRepo)
	}

	// Routes
	route.HomeRoute(v1, homeSrv)
	route.UserRoute(v1, userSrv, authSrv)
	route.ProductRoutes(v1, productSrv, authSrv)
	route.CartRoute(v1, cartSrv, authSrv)
	route.CartItemRoute(v1, cartItemSrv, authSrv)
	route.ErrorRoute(router)

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().Do(func() {
		mysqlDB.Ping()
	})

	// Documentation
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + addr)
}

func init() {
	addr = utils.AppConfig.ADDR
	mode = utils.AppConfig.MODE
	secret = utils.AppConfig.SECRET_KEY_TOKEN
	cache = utils.AppConfig.ENABLE_CACHE

	if addr == "" {
		addr = "8000"
	}

	if mode == "development" {
		gin.SetMode(gin.DebugMode)

		dsn = utils.AppConfig.DEVELOPMENT_DATABASE
		if dsn == "" {
			log.Println("DSN cannot be empty")
		}

		redis_addr = utils.AppConfig.DEVELOPMENT_REDIS_DATABASE
		if redis_addr == "" {
			log.Println("REDIS ADDRESS cannot be empty")
		}

		redis_username = utils.AppConfig.DEVELOPMENT_REDIS_DATABASE_USERNAME
		if redis_username == "" {
			log.Println("REDIS USERNAME cannot be empty")
		}

		redis_password = utils.AppConfig.DEVELOPMENT_REDIS_DATABASE_PASSWORD
		if redis_password == "" {
			log.Println("REDIS ADDRESS cannot be empty")
		}
	}

	if mode == "production" {
		gin.SetMode(gin.ReleaseMode)

		dsn = utils.AppConfig.PRODUCTION_DATABASE
		if dsn == "" {
			log.Println("DSN cannot be empty")
		}

		redis_addr = utils.AppConfig.PRODUCTION_REDIS_DATABASE_HOST
		if redis_addr == "" {
			log.Println("REDIS ADDRESS cannot be empty")
		}

		redis_username = utils.AppConfig.PRODUCTION_REDIS_DATABASE_USERNAME
		if redis_username == "" {
			log.Println("REDIS USERNAME cannot be empty")
		}

		redis_password = utils.AppConfig.PRODUCTION_REDIS_DATABASE_PASSWORD
		if redis_password == "" {
			log.Println("REDIS ADDRESS cannot be empty")
		}
	}

	if secret == "" {
		secret = "somerandomsecret"
		log.Println("Please provide a secret key token")
	}
}
