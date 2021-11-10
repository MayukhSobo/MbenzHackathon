package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	_ "mbenz_auth/docs"                   // load API Docs files (Swagger)
	"mbenz_auth/pkg/configs"
	"mbenz_auth/pkg/middlewares"
	"mbenz_auth/pkg/routes"
	"mbenz_auth/pkg/utils"
	"os"
	"strings"
)


// @title Mercedes-Benz Auth API
// @version 1.0
// @description This is the authentication API for mercedes-benz
// @termsOfService http://swagger.io/terms/
// @contact.name Mayukh Sarkar
// @contact.email mayukh2012@hotmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Define a fiber config
	config := configs.FiberConfig()


	// Define a fiber app with config
	app := fiber.New(config)

	// Add middlewares
	middlewares.FiberMiddlewaresCommon(app)

	// All routes
	routes.SwaggerRoute(app)  // For swagger docs
	routes.PublicRoutes(app)
	routes.NotFoundRoute(app) // For 404 not found page

	// Start the server
	if strings.ToLower(os.Getenv("DEPLOY_ENV")) == "dev" {
		utils.StartServer(app)
	}else {
		utils.StartServerWithGracefulShutdown(app)
	}
}


//func main() {
//	uri := "host=localhost port=5432 user=postgres dbname=mbenzdb sslmode=disable password=password"
//	db, err := gorm.Open("postgres", uri)
//	if err != nil{
//		panic(err)
//	}
//	var rows int
//	db.Debug().Table("users").Where("email = ?", "mayukh2012@hotmail.com").Count(&rows)
//	fmt.Println(rows)
//}
