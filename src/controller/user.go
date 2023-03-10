package controller

import (
	"Bugs-Bunny/src/db"
	"Bugs-Bunny/src/schemas"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController struct{}

func (u *UserController) GetUserByEmail(c *fiber.Ctx) error {
	// Initialize a request
	req := new(schemas.GetUserByEmail)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request",
		})
	}
	if err := req.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request",
			"errors":  err,
		})
	}
	user, err := db.UserService.GetUserByEmail(req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User Not Found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"body":    user,
	})
}

func (u *UserController) ReplaceApiKey(c *fiber.Ctx) error {
	// TODO: Implement ReplaceApiKey

	return nil
}
