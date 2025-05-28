package handlers

import (
	"ws-system/internal/config"
	"ws-system/internal/services"

	"github.com/gofiber/websocket/v2"
	"github.com/rs/zerolog"
)

type WsHandler struct {
	clients *services.Clients
	logger  *zerolog.Logger
}

func NewWsHandler(config *config.AppConfig,
	clients *services.Clients,
	logger *zerolog.Logger) *WsHandler {
	return &WsHandler{
		clients: clients,
		logger:  logger,
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
		messageType, _, err := c.ReadMessage()
		if err != nil {
			h.logger.Error().Err(err).Msg("ReadMessage")
			break
		}
		if messageType == websocket.CloseMessage {
			break
		}
	}
}
