package controller

import (
	"github.com/ADMex1/goweb/database"
	"github.com/ADMex1/goweb/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func BookIndex(c *fiber.Ctx) error {
	var books []model.Book
	database.DB.Preload("User").Find(&books)
	return c.JSON(books)
}

func AddBook(c *fiber.Ctx) error {
	book := new(model.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Error": "Cannot Parse JSON",
		})
	}
	// user := c.Locals("user").(jwt.MapClaims)
	// userID := int(user["id"].(float64))

	// book.CreatedBy = userID
	user := c.Locals("user").(jwt.MapClaims)
	users := int(user["id"].(float64)) // token must have "id"

	book.CreatedBy = int64(users)
	if result := database.DB.Create(&book); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": "unable to add book",
		})
	}
	return c.JSON(book)
}

func BookPerSlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return c.Status(400).JSON(fiber.Map{
			"Error": "Slug parameter is required",
		})
	}

	var books model.Book
	if err := database.DB.Preload("User").Where("slug = ?", slug).First(&books); err.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"Error": "Book not found",
		})
	}
	return c.JSON(books)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var books model.Book
	if result := database.DB.First(&books, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"Error": "Book not found",
		})
	}
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Error": "Invalid Input",
		})
	}
	if len(updates) == 0 {
		return c.Status(400).JSON(fiber.Map{"Error": "No fields provided to update"})
	}

	database.DB.Model(&books).Updates(updates)
	database.DB.First(&books, id)
	return c.JSON(books)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if result := database.DB.Delete(&model.Book{}, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"Error": "Book not found",
		})
	}
	return c.SendStatus(204)
}
