package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error { // GET /users
		return c.JSON(&fiber.Map{ // Return JSON
			"users": []string{"John", "Jane", "Joe"}, // Array of users
		})
	})

	app.Listen(":" + port)
	fmt.Println("Server started on port 3000")
}
