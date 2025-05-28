package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
	"ws-system/internal/models"

	"github.com/gofiber/websocket/v2"
)

type Clients struct {
	sync.RWMutex
	connections map[*websocket.Conn]struct{}
}

func NewClients() *Clients {
	return &Clients{
		connections: make(map[*websocket.Conn]struct{}),
	}
}

func (c *Clients) Add(conn *websocket.Conn) {
	c.Lock()
	defer c.Unlock()
	c.connections[conn] = struct{}{}
	fmt.Printf("Add connection, %d connection\n", len(c.connections))
}

func (c *Clients) Delete(conn *websocket.Conn) error {
	fmt.Println("Delete Connection")

	c.Lock()
	defer c.Unlock()

	if _, ok := c.connections[conn]; ok {
		delete(c.connections, conn)
		conn.Close()
		return nil
	}

	return errors.New("connection not found")
}

func (c *Clients) Push(messageJSON []byte) {
	var wg sync.WaitGroup

	c.RLock()
	defer c.RUnlock()

	for conn := range c.connections {
		wg.Add(1)
		go func(conn *websocket.Conn) {
			defer wg.Done()

			conn.WriteMessage(websocket.TextMessage, messageJSON)
		}(conn)
	}

	wg.Wait()
}

func (c *Clients) CreateMessage(content string) ([]byte, error) {
	message := models.Message{
		Content:   content,
		Timestamp: time.Now(),
	}

	json, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message to JSON: %v", err)
	}

	return json, nil
}

func (c *Clients) Parse(content []byte) (*models.Message, error) {

	var htmxMessage models.HtmxMessage

	err := json.Unmarshal(content, &htmxMessage)

	if err != nil {
		return nil, err
	}

	message := models.Message{
		Content:   htmxMessage.Message,
		Timestamp: time.Now(),
	}

	return &message, nil
}
