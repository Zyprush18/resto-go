package routes

import (
	"github.com/Zyprush18/resto-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func Route(c *fiber.App) {
	// v1

	// user
	c.Get("/api/v1/users", controllers.UserControllerIndex)
	c.Post("/api/v1/users/create", controllers.UserControllerCreate)
	c.Get("/api/v1/users/:id/show", controllers.UserControllerShow)
	c.Put("/api/v1/users/:id/update", controllers.UserControllerUpdate)
	c.Delete("/api/v1/users/:id/delete", controllers.UserControllerDelete)

	// menu
	c.Get("/api/v1/menu", controllers.MenuControllerIndex)
	c.Post("/api/v1/menu/create", controllers.MenuControllerCreate)
	c.Get("/api/v1/menu/:id/show", controllers.MenuControllerShow)
	c.Put("/api/v1/menu/:id/update", controllers.MenuControllerUpdate)
	c.Delete("/api/v1/menu/:id/delete", controllers.MenuControllerDelete)

	// order
	c.Get("/api/v1/order", controllers.OrderControllerIndex)
	c.Post("/api/v1/order/create", controllers.OrderControllerCreate)
	c.Get("/api/v1/order/:id/show", controllers.OrderControllerShow)
	c.Put("/api/v1/order/:id/update", controllers.OrderControllerUpdated)
	c.Delete("/api/v1/order/:id/delete", controllers.OrderControllerDelete)
}
