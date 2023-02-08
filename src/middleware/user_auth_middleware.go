package middleware

import (
	"Bugs-Bunny/src/db"
	"github.com/gofiber/fiber/v2"
	uuid2 "github.com/google/uuid"
	"log"
)

// UserAuthMiddleware is a middleware that checks if the user is authenticated
func UserAuthMiddleware(c *fiber.Ctx) error {
	// TODO: Implement User Auth Middleware
	return c.Next()

}

func ApiAuthMiddleware(c *fiber.Ctx) error {
	RequestApiKey := c.Get("api-key", "")
	// Api Key not found
	if RequestApiKey == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Access Credentials",
		})
	}
	uuid, err := uuid2.Parse(RequestApiKey)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Access Credentials",
		})
	}
	if check, err := db.ApiService.ApiAuthCheck(&uuid); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	} else {
		if !check {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid Access Credentials",
			})
		}
		return c.Next()
	}

}
