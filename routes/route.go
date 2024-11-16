package routes

import (
	"github.com/Zyprush18/resto/controllers"
	"github.com/gofiber/fiber/v2"
)


func Route(c *fiber.App)  {
	// v1

	// user
	c.Get("/api/v1/users", controllers.UserControllerIndex)
	c.Post("/api/v1/users/create", controllers.UserControllerCreate)
}