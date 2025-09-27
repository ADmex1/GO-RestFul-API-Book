package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// DatabaseRoute just checks the existing DB connection
func DatabaseRoute(app *fiber.App, db *gorm.DB) {
	app.Get("/connection", func(c *fiber.Ctx) error {
		sqlDB, err := db.DB()
		if err != nil {
			return c.Status(500).SendString("DB connection error: " + err.Error())
		}
		if err := sqlDB.Ping(); err != nil {
			return c.Status(500).SendString("DB not connected: " + err.Error())
		}
		return c.JSON(fiber.Map{
			"{+}Status": "Connected",
			"Database":  os.Getenv("DB_NAME"),
		})
	})
}
