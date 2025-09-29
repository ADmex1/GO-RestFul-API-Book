package routes

import (
	"github.com/ADMex1/goweb/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRoute(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/user", middleware.TokenBearer, func(c *fiber.Ctx) error {
		claims := c.Locals("user").(jwt.MapClaims)
		email := claims["email"].(string)

		return c.JSON(fiber.Map{
			"{+}User": email,
			"{*}Data": claims,
		})
	})
}
