package main

import (
	"log"
	"mime"

	"github.com/gofiber/fiber/v2"
)

type Text struct {
	Text string `json:"text"`
}

func main() {
	mime.AddExtensionType(".js", "text/javascript")

	list := []string{"one"}

	app := fiber.New()

	app.Static("/", "/vite-project/dist")

	api := app.Group("/api")
	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})
	api.Get("/list", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"list": list})
	})
	api.Post("/list", func(c *fiber.Ctx) error {
		t := &Text{}
		if err := c.BodyParser(t); err != nil {
			log.Println(err)
			return err
		}
		list = append(list, t.Text)
		return c.JSON(fiber.Map{"create": t})
	})

	app.Listen(":3000")
}
