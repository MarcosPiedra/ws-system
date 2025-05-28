package main

import (
	"context"
	"fmt"
	"os"
	"ws-system/internal/config"
	"ws-system/internal/logger"
	"ws-system/internal/services"
	"ws-system/ws-ingestor/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("backend exitted abnormally: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	cfg, err = config.Setup()
	if err != nil {
		return err
	}

	logger := logger.NewLogger(&cfg)
	app := fiber.New()
	clients := services.NewClients()

	router.Register(&cfg, app, clients, logger)

	statusJob := services.NewStatusJob(&cfg, clients, "statusA", &logger)
	go statusJob.Init(context.Background())

	return app.Listen(cfg.Web.Port)
}
