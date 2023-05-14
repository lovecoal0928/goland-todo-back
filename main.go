package main

import (
	"go-todo-back/configs"
	"go-todo-back/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.BootApp()
	app := fiber.New()

	// Initial route
	routes.InitRoute(app)

	app.Listen(":8000")
}
