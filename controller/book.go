package controller

import (
	"github.com/ADMex1/goweb/database"
	"github.com/ADMex1/goweb/model"
	"github.com/gofiber/fiber/v2"
)

func BookIndex(c *fiber.Ctx) error {
	var books []model.Book
	database.DB.Find(&books)
	return c.JSON(books)
}

func AddBook(c *fiber.Ctx) error {
	book := new(model.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Error": "Cannot Parse JSON",
		})
	}
	if result := database.DB.Create(&book); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": "unable to add book",
		})
	}
	return c.JSON(book)
}
