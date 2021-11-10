package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"mbenz_planning/configs"
	"mbenz_planning/middleware"
	"mbenz_planning/routes"
	"mbenz_planning/utils"
	"os"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
