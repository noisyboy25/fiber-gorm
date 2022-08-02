package main

import (
	"log"
	"mime"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Todo struct {
	ID   uint   `json:"id" gorm:"primaryKey; not null"`
	Text string `json:"text"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Todo{})
}

func main() {
	mime.AddExtensionType(".js", "text/javascript")

	one := Todo{ID: 1, Text: "one"}
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&one)
	if result.Error != nil {
		log.Printf("Cannot create one: %s", result.Error)
	}

	app := fiber.New()

	app.Static("/", "/vite-project/dist")

	api := app.Group("/api")
	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})
	api.Get("/todos", func(c *fiber.Ctx) error {
		todos := []Todo{}
		result := db.Find(&todos)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(fiber.Map{"todos": todos})
	})
	api.Post("/todos", func(c *fiber.Ctx) error {
		t := &Todo{}
		if err := c.BodyParser(t); err != nil {
			log.Println(err)
			return err
		}
		result := db.Create(&t)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(fiber.Map{"create": t})
	})
	api.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		result := db.Delete(&Todo{}, id)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(fiber.Map{"deleteCount": result.RowsAffected})
	})

	app.Listen(":3000")
}
