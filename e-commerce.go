package main

import "e-commerce/cmd"

func main() {
	// @title           Benny-Foods
	// @version         1.0
	// @description     Server for https://bennyfoodie.netlify.app/
	// @termsOfService  http://swagger.io/terms/

	// @contact.name   Confidence James
	// @contact.url    http://github.com/jamesconfy
	// @contact.email  bobdence@gmail.com

	// @license.name  Apache 2.0
	// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

	// @host      benny-foods.fly.dev
	// @schemes http https
	// @BasePath  /api/v1

	// @securityDefinitions.apiKey  ApiKeyAuth
	// @in header
	// @name Authorisation
	cmd.Setup()
}
