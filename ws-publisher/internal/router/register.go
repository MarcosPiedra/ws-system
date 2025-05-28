package router

import (
	"ws-system/internal/config"
	"ws-system/internal/middleware"
	"ws-system/internal/services"
	"ws-system/ws-publisher/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
	"github.com/rs/zerolog"
)

func Register(config *config.AppConfig,
	app *fiber.App,
	logger *zerolog.Logger,
	clients *services.Clients) {

	var wsHandler = handlers.NewWsHandler(config, clients, logger)
	var wsAuth = middleware.NewWsAuth(config)

	app.Use(fiberLog.New())
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(wsAuth.WsAuth)
	app.Get("/ws", websocket.New(wsHandler.HandleWebSocket))
}
