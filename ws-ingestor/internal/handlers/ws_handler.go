package handlers

import (
	"ws-system/internal/config"
	"ws-system/internal/services"

	"github.com/gofiber/websocket/v2"
	"github.com/rs/zerolog"
)

type WsHandler struct {
	clients   *services.Clients
	publisher *services.Publisher
	logger    zerolog.Logger
}

func NewWsHandler(config *config.AppConfig,
	clients *services.Clients,
	logger zerolog.Logger) *WsHandler {
	return &WsHandler{
		clients:   clients,
		publisher: services.NewPublisher(config, logger),
		logger:    logger,
	}
}

func (h *WsHandler) HandleWebSocket(c *websocket.Conn) {

	h.clients.Add(c)

	defer func() {
		h.clients.Delete(c)
		c.Close()
		h.logger.Info().Msg("Client disconnected")
	}()

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			h.logger.Error().Err(err).Msg("Read message")
			break
		}

		if messageType == websocket.CloseMessage {
			break
		}

		msg, err := h.clients.Parse(message)
		if err != nil {
			h.logger.Error().Err(err).Msg("Parse message")
			continue
		}

		if err := h.publisher.Publish(msg); err != nil {
			h.logger.Error().Err(err).Msg("Publish")
			continue
		}

		h.logger.Debug().Msgf("Message sent: %s", msg.Content)
	}
}
