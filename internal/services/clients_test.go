package services_test

import (
	"net/http/httptest"
	"testing"
	"time"
	"ws-system/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/stretchr/testify/assert"
)

func TestAddAndDeleteClient(t *testing.T) {
	app := fiber.New()

	clients := services.NewClients()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		clients.Add(c)
		time.Sleep(100 * time.Millisecond)
		err := clients.Delete(c)
		assert.NoError(t, err)
	}))

	req := httptest.NewRequest("GET", "/ws", nil)
	req.Header.Add("Connection", "Upgrade")
	req.Header.Add("Upgrade", "websocket")
	req.Header.Add("Sec-WebSocket-Version", "13")
	req.Header.Add("Sec-WebSocket-Key", "x3JJHMbDL1EzLkh9GBhXDw==")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 101, resp.StatusCode)
}

func TestCreateMessage(t *testing.T) {
	clients := services.NewClients()

	msg, err := clients.CreateMessage("Hello")
	assert.NoError(t, err)
	assert.Contains(t, string(msg), "Hello")
}

func TestParseValidMessage(t *testing.T) {
	clients := services.NewClients()

	input := []byte(`{"message": "Hello from HTMX"}`)
	result, err := clients.Parse(input)

	assert.NoError(t, err)
	assert.Equal(t, "Hello from HTMX", result.Content)
	assert.WithinDuration(t, time.Now(), result.Timestamp, time.Second)
}

func TestParseInvalidMessage(t *testing.T) {
	clients := services.NewClients()

	input := []byte(`{invalid json}`)
	result, err := clients.Parse(input)

	assert.Error(t, err)
	assert.Nil(t, result)
}
