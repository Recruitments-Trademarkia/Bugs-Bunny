package middleware

import (
	"Bugs-Bunny/src/db"
	"log"

	"github.com/gofiber/fiber/v2"
	uuid2 "github.com/google/uuid"
)

// UserAuthMiddleware is a middleware that checks if the user is authenticated
func UserAuthMiddleware(c *fiber.Ctx) error {
	// TODO: Implement User Auth Middleware
	// Get the session token from the request headers
	sessionToken := c.Get("Authorization")

	// If the session token is not found, return an Unauthorized response
	if sessionToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized. Missing authorization header.",
		})
	}
	uuid, err := uuid2.Parse(sessionToken)
	// Call the GetUserBySessionToken method to retrieve the user from the database
	user, err := db.ApiService.GetApiKeyByUserId(&uuid)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized. Invalid authorization header.",
		})
	}

	// If the user is found, set the user in the context for later use
	c.Locals("user", user)

	// Continue to the next middleware in the chain

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
