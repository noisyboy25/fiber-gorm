package main

import (
	"errors"
	"log"
	"mime"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Todo struct {
	ID     uint   `json:"id" gorm:"primaryKey; not null"`
	Text   string `json:"text"`
	UserID *uint  `json:"userId"`
	User   *User  `json:"user"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Todos    []Todo `json:"todos"`
}

type AuthPair struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Todo{}, &User{})
}

func main() {
	mime.AddExtensionType(".js", "text/javascript")

	one := Todo{ID: 1, Text: "one"}
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&one)
	if result.Error != nil {
		log.Printf("Cannot create one: %s", result.Error)
	}

	app := fiber.New()

	app.Static("/", "vite-project/dist")

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

	api.Get("/users", func(c *fiber.Ctx) error {
		users := []User{}
		result := db.Preload("Todos").Find(&users)
		if result.Error != nil {
			return result.Error
		}

		return c.JSON(fiber.Map{"users": users})
	})

	auth := app.Group("/auth")
	auth.Post("/login", func(c *fiber.Ctx) error {
		a := &AuthPair{}
		if err := c.BodyParser(a); err != nil {
			return err
		}

		user := &User{}
		result := db.Where("username = ?", a.Username).First(user)
		if result.Error != nil || user.Password != a.Password {
			return errors.New("wrong username or password")
		}

		return c.JSON(fiber.Map{"username": user.Username, "message": "user logged in successfully"})
	})
	auth.Post("/register", func(c *fiber.Ctx) error {
		a := &AuthPair{}
		if err := c.BodyParser(a); err != nil {
			return err
		}

		user := &User{Username: a.Username, Password: a.Password}
		result := db.Create(&user)
		if result.Error != nil {
			var sqliteError sqlite3.Error
			if errors.As(result.Error, &sqliteError) {
				if errors.Is(sqliteError.ExtendedCode, sqlite3.ErrConstraintUnique) {
					return errors.New("user already exists")
				}
			}
			return result.Error
		}
		return c.JSON(fiber.Map{"message": "user created successfully"})
	})

	app.Listen(":3000")
}
