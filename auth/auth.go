package auth

import (
	"errors"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/noisyboy25/fiber-gorm/model"
	"gorm.io/gorm"
)

type Config struct {
	Db *gorm.DB
}

type AuthPair struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.GetReqHeaders()["Authorization"]
		authPair, err := parseAuthPair(auth)
		if err != nil {
			return err
		}
		var user model.User
		result := config.Db.Where("username = ?", authPair.Username).First(&user)
		if result.Error != nil {
			return result.Error
		}
		if user.Password != authPair.Password {
			return c.SendStatus(401)
		}
		c.Locals("auth", authPair)
		return c.Next()
	}
}

func parseAuthPair(s string) (AuthPair, error) {
	r := regexp.MustCompile(`Basic (\w+):(\w+)`)
	groups := r.FindStringSubmatch(s)
	if len(groups) < 3 {
		return AuthPair{}, errors.New("invalid auth pair")
	}
	return AuthPair{groups[1], groups[2]}, nil
}
