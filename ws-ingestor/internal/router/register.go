package router

import (
	"time"
	"ws-system/internal/config"
	"ws-system/internal/middleware"
	"ws-system/internal/services"
	"ws-system/ws-ingestor/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
	"github.com/rs/zerolog"
)

func Register(config *config.AppConfig,
	app *fiber.App,
	clients *services.Clients,
	logger zerolog.Logger) {

	var loginHandler = handlers.NewLoginHandler(config, logger)
	var wsHandler = handlers.NewWsHandler(config, clients, logger)
	var wsAuth = middleware.NewWsAuth(config)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowCredentials: true,
	}))
	app.Use(fiberLog.New())
	app.Use(recover.New())

	app.Post("/login", loginHandler.Login)
	app.Use(wsAuth.WsAuth)
	app.Get("/ws", websocket.New(wsHandler.HandleWebSocket, websocket.Config{
		HandshakeTimeout: 20 * time.Second,
	})).Name("subscriber")
}
