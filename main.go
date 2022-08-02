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

type Text struct {
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
	db.AutoMigrate(&Text{})
}

func main() {
	mime.AddExtensionType(".js", "text/javascript")

	one := Text{ID: 1, Text: "one"}
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
	api.Get("/list", func(c *fiber.Ctx) error {
		list := []Text{}
		result := db.Find(&list)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(fiber.Map{"list": list})
	})
	api.Post("/list", func(c *fiber.Ctx) error {
		t := &Text{}
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
	api.Delete("/list/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}
		result := db.Delete(&Text{}, id)
		if result.Error != nil {
			return result.Error
		}
		return c.JSON(fiber.Map{"deleteCount": result.RowsAffected})
	})

	app.Listen(":3000")
}
