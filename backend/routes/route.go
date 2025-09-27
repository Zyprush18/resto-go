package routes

import (
	"github.com/Zyprush18/resto-go/backend/controllers"
	"github.com/Zyprush18/resto-go/backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func Route(c *fiber.App) {
	// v1
	v1 := c.Group("/api/v1")
	v1.Use(csrf.New(csrf.Config{
		// untuk setiap method get nggak gunain csrf token
		Next: func(c *fiber.Ctx) bool {
			return c.Route().Method == "GET"
		},
	}))

	// user
	v1.Get("/users", middleware.MiddlewareGlobal, middleware.MiddlewareAccess("admin"),controllers.UserControllerIndex)
	v1.Post("/users/create", middleware.MiddlewareGlobal, middleware.MiddlewareAccess("admin"),controllers.UserControllerCreate)
	v1.Get("/users/:id/show", middleware.MiddlewareGlobal, middleware.MiddlewareAccess("admin"),controllers.UserControllerShow)
	v1.Put("/users/:id/update", middleware.MiddlewareGlobal, middleware.MiddlewareAccess("admin"),controllers.UserControllerUpdate)
	v1.Delete("/users/:id/delete", middleware.MiddlewareGlobal, middleware.MiddlewareAccess("admin"),controllers.UserControllerDelete)

	// menu
	v1.Get("/menu", controllers.MenuControllerIndex)
	v1.Post("/menu/create", middleware.MiddlewareGlobal, middleware.MiddlewareAccess("admin"),controllers.MenuControllerCreate)
	v1.Get("/menu/:id/show", controllers.MenuControllerShow)
	v1.Put("/menu/:id/update", middleware.MiddlewareGlobal, middleware.MiddlewareAccess("admin"),controllers.MenuControllerUpdate)
	v1.Delete("/menu/:id/delete", middleware.MiddlewareGlobal, middleware.MiddlewareAccess("admin"),controllers.MenuControllerDelete)

	// order
	v1.Get("/order", controllers.OrderControllerIndex)
	v1.Post("/order/create", controllers.OrderControllerCreate)
	v1.Get("/order/:id/show", controllers.OrderControllerShow)
	v1.Put("/order/:id/update", controllers.OrderControllerUpdated)
	v1.Delete("/order/:id/delete", controllers.OrderControllerDelete)

	// order item
	v1.Get("/orderitem", controllers.OrderItemControllerIndex)
	v1.Post("/orderitem/create", controllers.OrderItemControllerCreate)
	v1.Get("/orderitem/:id/show", controllers.OrderItemControllerShow)
	v1.Put("/orderitem/:id/update", controllers.OrderItemControllerUpdate)
	v1.Delete("/orderitem/:id/delete", controllers.OrderItemControllerDelete)

	// reservation
	v1.Get("/reservation", controllers.ReservationControllerIndex)
	v1.Post("/reservation/create", controllers.ReservationControllerCreate)
	v1.Get("/reservation/:id/show", controllers.ReservationControllerShow)
	v1.Put("/reservation/:id/update", controllers.ReservationControllerUpdate)
	v1.Delete("/reservation/:id/delete", controllers.ReservationControllerDelete)

	// login
	c.Post("/api/login", controllers.LoginController)
}
