package routes

import (
	"github.com/gofiber/fiber/v2"
	"mbenz_planning/controllers"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/business/v1")

	/// Routes for GET methods:
	route.Post("/route/", controllers.GetRoutingDetails)
}
