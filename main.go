package main

import (
	"github.com/framadhita4/quranlet-api/platform/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Api Is UP!")

		return err
	})

	app.Listen(":3000")
}
