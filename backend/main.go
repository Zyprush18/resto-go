package main

import (
	"log"

	"github.com/Zyprush18/resto-go/backend/repositories/databases"
	"github.com/Zyprush18/resto-go/backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	databases.Migrate()
	routes.Route(app)

	log.Println("App Running On port 3000")
	app.Listen(":3000")
}
