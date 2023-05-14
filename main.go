package main

import (
	"go-todo-back/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initial route
	routes.InitRoute(app)

	app.Listen(":8000")
}
