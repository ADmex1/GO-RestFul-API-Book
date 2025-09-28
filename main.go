package main

import (
	"os"

	"github.com/ADMex1/goweb/database"
	"github.com/ADMex1/goweb/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// db, err := database.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	app := fiber.New()

	//Route Call
	database.Connect()
	database.Migration()

	routes.BookRoutes(app)
	routes.DatabaseRoute(app, database.DB)
	routes.UserRoute(app)
	//Default Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"{#}Status": "API is Running....",
		})
	})
	app.Get("/:3", func(c *fiber.Ctx) error {
		c.Type("html")
		return c.SendString(
			`<html>
	    <body>
	        <a href="https://www.youtube.com/watch?v=UagP13Zgd7o">
			<h3> Marshi I UÇK </h3>
	        </a>
	    </body>
	</html>`,
		)
	})
	app.Listen(":" + os.Getenv("PORT"))
}
