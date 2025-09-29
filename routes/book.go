package routes

import (
	"github.com/ADMex1/goweb/controller"
	"github.com/ADMex1/goweb/middleware"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	app.Get("/books", controller.BookIndex)
	app.Get("/books/:slug", controller.BookPerSlug)

	//Group by /Api, so Calling the Api will be like (/Api/Addbook)
	books := app.Group("/api", middleware.TokenBearer)
	books.Post("/Addbook", controller.AddBook)
	books.Put("/updatebook/:id", controller.UpdateBook)
	books.Delete("/deletebook/:id", controller.DeleteBook)
}
