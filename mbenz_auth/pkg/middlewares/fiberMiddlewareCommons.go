package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddlewaresCommon(a *fiber.App) {
	a.Use(
		// Add CORS to each route
		cors.New(),

		// Adding a simple logger
		logger.New(),
	)
}