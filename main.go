package main

import (
	"github.com/Zyprush18/resto/repositories/databases"
	"github.com/Zyprush18/resto/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	databases.Migrate()
	routes.Route(app)

	app.Listen(":3000")
}