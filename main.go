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
			`			<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>405 Method Not Allowed</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      text-align: center;
      background-color: #f8f9fa;
      color: #333;
      margin: 0;
      padding: 0;
    }
    .container {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      height: 100vh;
    }
    h1 {
      font-size: 6rem;
      color: #ff8800;
      margin: 0;
    }
    p {
      font-size: 1.2rem;
      margin: 10px 0 20px;
    }
    a {
      display: inline-block;
      text-decoration: none;
      background-color: #007bff;
      color: #fff;
      padding: 10px 20px;
      border-radius: 5px;
      transition: background-color 0.3s;
    }
    a:hover {
      background-color: #0056b3;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>405</h1>
    <p>Method Not Allowed – The request method is not supported for this resource.</p>
    <a href="/">Go Back Home</a>
  </div>
</body>
</html>`,
			// 		`<html>
			//     <body>
			//         <a href="https://www.youtube.com/watch?v=UagP13Zgd7o">
			// 		<h3> Marshi I UÇK </h3>
			//         </a>
			//     </body>
			// </html>`,
		)
	})
	app.Listen(":" + os.Getenv("PORT"))
}
