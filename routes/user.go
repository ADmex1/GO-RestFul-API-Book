package routes

import (
	"github.com/ADMex1/goweb/auth"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {

	user := app.Group("/user")
	user.Post("/register", auth.UserRegister)
	user.Post("/login", auth.UserLogin)
}
