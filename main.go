package main

import (
	"os"

	"github.com/ADMex1/goweb/database"
	"github.com/ADMex1/goweb/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// db, err := database.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	app := fiber.New()
	app.Use(logger.New())
	//Route Call
	database.Connect()
	database.Migration()

	routes.BookRoutes(app)
	routes.DatabaseRoute(app, database.DB)
	routes.UserRoute(app)
	routes.AuthRoute(app)
	//Default Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"{#}Status": "API is Running....",
		})
	})
	app.Static("/storage", "./storage")
	app.Get("/:3", func(c *fiber.Ctx) error {
		c.Type("html")
		return c.SendString(
			`	
<html>
  <head>
    <style>
      body {
        margin: 0;
        height: 100vh;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        text-align: center;
      }
    </style>
  </head>
  <body>
    <h1>404 Not Found</h1>
    <a href="/">Back</a>
    <a href="https://www.youtube.com/watch?v=UagP13Zgd7o">
      <h3>Marshi I UÇK</h3>
    </a>
  </body>
</html>`,
		)
	})

	app.Listen(":" + os.Getenv("PORT"))
}
