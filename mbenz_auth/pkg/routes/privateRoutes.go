package routes

import (
	"github.com/gofiber/fiber/v2"
	"mbenz_auth/app/controllers"
	"mbenz_auth/pkg/middlewares"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	// Routes for the POST method
	route.Post("/user/sign/out", middlewares.JWTProtected(), controllers.UserSignOut)
	route.Post("/token/renew", middlewares.JWTProtected(), controllers.RenewToken)

}
