package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTGenerator() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == " " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"{-}Error": "Missing Token",
			})
		}
		split := strings.Split(auth, " ")
		if len(split) != 2 || split[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"{-}Error": "Invalid Token Format",
			})
		}
		tokenStr := split[1]

		secret := []byte("Sekret Token")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil || token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"{-}Error": "Invalid Token",
			})
		}
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user", claims)
		return c.Next()
	}
}
