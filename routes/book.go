package routes

import (
	"github.com/ADMex1/goweb/controller"
	"github.com/ADMex1/goweb/middleware"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	app.Get("/books", controller.BookIndex)
	app.Get("/books/:slug", controller.BookPerSlug)

	books := app.Group("/addbook", middleware.TokenBearer)
	books.Post("/Addbook", controller.AddBook)
	books.Put("/updatebook/:id", controller.UpdateBook)
	books.Delete("/deletebook/:id", controller.DeleteBook)
}
