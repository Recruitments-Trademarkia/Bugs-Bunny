package router

import (
	"Bugs-Bunny/src/controller"
	"Bugs-Bunny/src/middleware"

	"github.com/gofiber/fiber/v2"
)

// MountUserSideRoutes mounts all the routes for the user side
func MountUserSideRoutes(app fiber.Router) {
	userGroup := app.Group("/user", middleware.UserAuthMiddleware)
	{
		thisController := controller.UserController{}
		userGroup.Get("/fetch", thisController.GetUserByEmail)
		// TODO: Replace API key controller.UserController.ReplaceApiKey
		userGroup.Put("/replace-api-key", thisController.ReplaceApiKey)
	}

}
