package routes

import (
	"github.com/ADMex1/goweb/controller"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	app.Get("/books", controller.BookIndex)
}
