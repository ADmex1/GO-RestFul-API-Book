package routes

import (
	"github.com/ADMex1/goweb/controller"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	app.Get("/books", controller.BookIndex)
	app.Post("/Addbook", controller.AddBook)
	app.Get("/books/:slug", controller.BookPerSlug)
	app.Put("/updatebook/:id", controller.UpdateBook)
	app.Delete("/deletebook/:id", controller.DeleteBook)
}
