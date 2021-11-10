package routes

import (
	"github.com/gofiber/fiber/v2"
	"mbenz_poc/controllers"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/internal/v1")

	/// Routes for GET methods:
	route.Get("/status/:vin?", controllers.GetVehicleStatus)
	route.Get("/distance/", controllers.GetDistance)
	route.Get("/stations/", controllers.GetStations)
}
