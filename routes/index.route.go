package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) InitRoute(app *fiber.App) {
	v1Route(app)
}
