package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(a *fiber.App) {
	route := a.Group("/docs")

	// Swagger docs method
	route.Get("*", swagger.Handler)

}
