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


//func main() {
//	a := fiber.AcquireAgent()
//	defer fiber.ReleaseAgent(a)
//	req := a.Request()
//
//	req.Header.SetMethod(fiber.MethodGet)
//	req.SetRequestURI("http://localhost:9000/api/internal/v1/status/W1K2062161F0046")
//	req.URI().SetQueryString("source=home&destination=Power Plant")
//	req.URI().QueryArgs().Add("source", "home")
//	req.URI().QueryArgs().Add("destination", "Power Plant")
//
//	if err := a.Parse(); err != nil {
//		panic(err)
//	}
//	httpCode, body, err := a.Bytes()
//	if err != nil {
//		panic(err)
//	}
//	if httpCode != fiber.StatusOK {
//		panic(fmt.Errorf("request was not successful. ErrorCode %d", httpCode))
//	}
//	fmt.Println(string(body))
//}