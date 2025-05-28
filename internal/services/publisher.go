package services

import (
	"context"
	"encoding/json"
	"ws-system/internal/config"
	"ws-system/internal/models"

	redis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type Publisher struct {
	redisClient *redis.Client
	channel     string
	logger      zerolog.Logger
}

func NewPublisher(config *config.AppConfig,
	logger zerolog.Logger) *Publisher {
	redisClient := redis.NewClient(&redis.Options{Addr: config.Redis.Addr})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		logger.Error().Err(err).Msg("Redis client")
	}
	return &Publisher{
		redisClient: redisClient,
		channel:     config.Redis.Channel,
	}
}

func (p *Publisher) Publish(msg *models.Message) error {
	jsonStr, err := json.Marshal(msg)
	if err != nil {
		p.logger.Error().Err(err).Msg("Marshal")
	}

	return p.redisClient.Publish(context.Background(), p.channel, jsonStr).Err()
}
