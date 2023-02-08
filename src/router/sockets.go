package router

import (
	"Bugs-Bunny/src/compiler"
	"Bugs-Bunny/src/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// MountSockets mounts all the routes for the user side
func MountSockets(app fiber.Router) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:lang", middleware.ApiAuthMiddleware, websocket.New(func(c *websocket.Conn) {
		lang := c.Params("lang")
		l := compiler.GetLanguage(lang)
		if l == nil {
			c.WriteMessage(websocket.TextMessage, []byte("Request language is not yet supported"))
			return
		}
		l.Handler(c)
	}))
}
