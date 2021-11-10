package routes

import (
	"github.com/gofiber/fiber/v2"
	"mbenz_auth/app/controllers"
)

// PublicRoutes func for describe group of public routes
func PublicRoutes(a *fiber.App) {
	// Create a route group
	route := a.Group("/api/v1")

	// Routes for POST methods
	route.Post("/user/sign/up", controllers.UserSignUP)
	route.Post("user/sign/in", controllers.UserSignIN)
}
