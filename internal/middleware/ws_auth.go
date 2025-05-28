package middleware

import (
	"net/http"
	"ws-system/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type WsAuth struct {
	config *config.AppConfig
}

func NewWsAuth(config *config.AppConfig) *WsAuth {
	return &WsAuth{
		config: config,
	}
}

func (ws *WsAuth) WsAuth(c *fiber.Ctx) error {
	token := c.Query("t")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON("unauthorized")
	}

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(ws.config.Web.Secret), nil
	})
	if err != nil || !t.Valid {
		return c.Status(http.StatusUnauthorized).JSON("unauthorized")
	}

	return c.Next()
}
