package controllers

import (
	"github.com/Zyprush18/resto-go/service"
	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {
	loginPayload := new(service.Login)

	if err := c.BodyParser(loginPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed request",
		})
	}

	user, err := service.LoginService(loginPayload.Email, loginPayload.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid email or password",
		})
	}

	token, err := service.CreateToken(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"token":   token,
	})
}
