package handlers

import (
	"time"
	"ws-system/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
)

type LoginHandler struct {
	logger zerolog.Logger
	config *config.AppConfig
}

func NewLoginHandler(config *config.AppConfig,
	logger zerolog.Logger) *LoginHandler {
	return &LoginHandler{
		logger: logger,
		config: config,
	}
}

func (l *LoginHandler) Login(c *fiber.Ctx) error {
	payload := struct {
		User string `json:"user"`
		Pass string `json:"pass"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if l.config.Web.User != payload.User || l.config.Web.Pass != payload.Pass {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": l.config.Web.User,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	t, err := token.SignedString([]byte(l.config.Web.Secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
