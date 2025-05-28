package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"ws-system/internal/config"

	"ws-system/internal/services"
	"ws-system/ws-ingestor/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupApp() *fiber.App {
	app := fiber.New()
	conf := &config.AppConfig{
		Web: config.WebConfig{
			User: "User",
			Pass: "test123",
		},
	}
	logger := zerolog.Nop()
	clients := services.NewClients()

	router.Register(conf, app, clients, logger)
	return app
}

func TestLoginSuccess(t *testing.T) {
	app := setupApp()

	payload := map[string]string{
		"user": "User",
		"pass": "test123",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestLoginFail(t *testing.T) {
	app := setupApp()

	payload := map[string]string{
		"user": "wrong",
		"pass": "bad",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}

func TestWebSocketConnection(t *testing.T) {
	app := setupApp()

	payload := map[string]string{
		"user": "User",
		"pass": "test123",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	readReqBody, _ := io.ReadAll(resp.Body)
	var token map[string]any

	err = json.Unmarshal(readReqBody, &token)

	target := fmt.Sprintf("/ws?t=%s", token["token"])
	req = httptest.NewRequest(http.MethodGet, target, nil)
	req.Header.Add("Connection", "Upgrade")
	req.Header.Add("Upgrade", "websocket")
	req.Header.Add("Sec-WebSocket-Version", "13")
	req.Header.Add("Sec-WebSocket-Key", "x3JJHMbDL1EzLkh9GBhXDw==")

	resp, err = app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode)
}
