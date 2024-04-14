package fiber

import (
	"Bugs-Bunny/src/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"strings"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func Start() {
	app := fiber.New(
		fiber.Config{
			Concurrency: 0,
			GETOnly:     true,
			AppName:     "Trademarkia - Bugs Bunny",
		},
	)

	// Health check
	app.Get("/health", healthCheck)

	// Initialise Logger
	app.Use(logger.New(logger.Config{Next: func(c *fiber.Ctx) bool {
		// Don't log health check. (To avoid ALB Health check Spam)
		return strings.HasPrefix(c.Path(), "api")
	}}))

	router.MountSockets(app)

	log.Fatal(app.Listen(":8080"))
}
