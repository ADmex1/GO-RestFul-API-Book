package controller

import (
	"fmt"
	"os"
	"strings"

	"github.com/ADMex1/goweb/database"
	"github.com/ADMex1/goweb/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func BookIndex(c *fiber.Ctx) error {
	//Declare variable, calling the model book, and then Preload the User first so the identity of the uploader revealed
	var books []model.Book
	database.DB.Preload("User").Find(&books)
	return c.JSON(books)
}

func AddBook(c *fiber.Ctx) error {
	//Making sure the folder exists
	os.MkdirAll("./storage", os.ModePerm)

	book := new(model.Book)

	book.Title = c.FormValue("Title")
	book.Author = c.FormValue("Author")
	book.Description = c.FormValue("Description")

	//JSON Format
	// if err := c.BodyParser(book); err != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"Error": "Cannot Parse JSON",
	// 	})
	// }

	// user := c.Locals("user").(jwt.MapClaims)
	// userID := int(user["id"].(float64))

	// book.CreatedBy = userID
	//User required
	user := c.Locals("user").(jwt.MapClaims)
	users := int(user["id"].(float64)) // token must have "id"

	file, err := c.FormFile("bookfile")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"{-}Error": "No file uploaded",
		})
	}
	savePath := fmt.Sprintf("./storage/%s", file.Filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"{-}Error": "Unable to upload a book",
		})
	}
	book.FileUpload = savePath

	// if err := database.DB.Save(&book).Error; err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"{-}Error": "Unable to locate uploaded book",
	// 	})
	// }
	//Checking if there are incoming problem (Db dead or Query Violation)
	book.CreatedBy = int64(users)
	if result := database.DB.Create(&book); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": "unable to add book",
		})
	}
	return c.Status(201).JSON(book)
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
func slugify(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "/", "-")
	slug = strings.ReplaceAll(slug, "?", " ")
	slug = strings.Trim(slug, "-")

	return slug
}
func UpdateBook(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var book model.Book
	if result := database.DB.Where("slug = ?", slug).First(&book); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"Error": "Book not found",
		})
	}
	claims := c.Locals("user").(jwt.MapClaims)
	role := claims["role"].(string)
	uid := int(claims["id"].(float64))

	if role != "admin" && book.CreatedBy != int64(uid) {
		return c.Status(403).JSON(fiber.Map{
			"{#}": "Unauthorized Access",
		})
	}
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Error": "Invalid Input",
		})
	}

	if newTitle, ok := updates["title"].(string); ok && newTitle != " " {
		book.Title = newTitle
		book.Slug = slugify(newTitle)
	}
	if len(updates) == 0 {
		return c.Status(400).JSON(fiber.Map{"Error": "No fields provided to update"})
	}

	database.DB.Model(&book).Updates(updates)
	database.DB.First(&book, slug)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var book model.Book
	if err := database.DB.Where("slug = ?", slug).First(&book).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Error": "Book not found",
		})
	}
	claims := c.Locals("user").(jwt.MapClaims)
	role := claims["role"].(string)
	uid := int(claims["id"].(float64))

	if role != "admin" && book.CreatedBy != int64(uid) {
		return c.Status(403).JSON(fiber.Map{
			"{#}": "Unauthorized Access",
		})
	}
	if result := database.DB.Where("slug = ?", slug).Delete(&model.Book{}); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": "Unable to Delete",
		})
	}
	return c.Status(204).JSON(fiber.Map{
		"{+}Success": "The Book has been Deleted!",
	})
}
