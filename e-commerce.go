package main

import cmd "e-commerce/cmd/setup"

func main() {
	// @title           E-Commerce-Api
	// @version         1.0
	// @description     An e-commerce-api.
	// @termsOfService  http://swagger.io/terms/

	// @contact.name   Confidence James
	// @contact.url    http://github.com/jamesconfy
	// @contact.email  bobdence@gmail.com

	// @license.name  Apache 2.0
	// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

	// @host      e-commerce-api.fly.dev
	// @schemes http https
	// @BasePath  /api/v1

	// @securityDefinitions.apiKey  ApiKeyAuth
	// @in header
	// @name Authorisation
	cmd.Setup()
}
