package routes

import (
	"github.com/ADMex1/goweb/controller"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	app.Get("/books", controller.BookIndex)
	app.Post("/Addbook", controller.AddBook)
	app.Get("/books/:id", controller.BookPerId)
	app.Put("/books/:id", controller.UpdateBook)
	app.Delete("/books/:id", controller.DeleteBook)
}
