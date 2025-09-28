package routes

import (
	"github.com/ADMex1/goweb/auth"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/register", auth.UserRegister)
}
