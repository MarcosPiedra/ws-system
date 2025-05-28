package services

import (
	"context"
	"fmt"
	"time"
	"ws-system/internal/config"

	"github.com/rs/zerolog"
)

type StatusJob struct {
	logger     *zerolog.Logger
	clients    *Clients
	statusName string
}

func NewStatusJob(config *config.AppConfig,
	clients *Clients,
	statusName string,
	logger *zerolog.Logger) *StatusJob {
	return &StatusJob{
		clients:    clients,
		logger:     logger,
		statusName: statusName,
	}
}

func (b *StatusJob) Init(context context.Context) {

	for {
		html := fmt.Sprintf(`<span id="%s" class="inline-block px-2 py-1 text-xs bg-green-100 text-green-800 rounded-full">Connected</span>`, b.statusName)

		b.clients.Push([]byte(html))

		time.Sleep(time.Second)
	}
}
