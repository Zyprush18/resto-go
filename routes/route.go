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

	// order item
	c.Get("/api/v1/orderitem", controllers.OrderItemControllerIndex)
	c.Post("/api/v1/orderitem/create", controllers.OrderItemControllerCreate)
	c.Get("/api/v1/orderitem/:id/show", controllers.OrderItemControllerShow)
	c.Put("/api/v1/orderitem/:id/update", controllers.OrderItemControllerUpdate)
	c.Delete("/api/v1/orderitem/:id/delete", controllers.OrderItemControllerDelete)

	// reservation
	c.Get("/api/v1/reservation", controllers.ReservationControllerIndex)
	c.Post("/api/v1/reservation/create", controllers.ReservationControllerCreate)
	c.Get("/api/v1/reservation/:id/show", controllers.ReservationControllerShow)
}
