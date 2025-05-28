package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"ws-system/internal/config"
	"ws-system/internal/models"
	"ws-system/internal/services"

	redis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type ReceiverJob struct {
	redisClient *redis.Client
	channel     string
	logger      *zerolog.Logger
	clients     *services.Clients
}

func NewReceiverJob(config *config.AppConfig, clients *services.Clients, logger *zerolog.Logger) *ReceiverJob {
	redisClient := redis.NewClient(&redis.Options{Addr: config.Redis.Addr})
	return &ReceiverJob{
		redisClient: redisClient,
		channel:     config.Redis.Channel,
		clients:     clients,
		logger:      logger,
	}
}

func (b *ReceiverJob) Init(context context.Context) {

	if err := b.redisClient.Ping(context).Err(); err != nil {
		b.logger.Error().Err(err).Msg("Ping")
	}

	subscriber := b.redisClient.Subscribe(context, b.channel)

	for {

		msg, err := subscriber.ReceiveMessage(context)
		if err != nil {
			b.logger.Error().Err(err).Msg("ReceiveMessage")

			time.Sleep(time.Second)
			continue
		}

		var msgParsed models.Message

		if err := json.Unmarshal([]byte(msg.Payload), &msgParsed); err != nil {
			b.logger.Error().Err(err).Msg("Parsing message")

			break
		}

		b.logger.Info().Msgf("Message received from channel %s, payload %s", msg.Channel, msg.Payload)

		html := fmt.Sprintf(`<div id="txt">%s</div>`, msgParsed.Content)

		b.clients.Push([]byte(html))
	}
}
