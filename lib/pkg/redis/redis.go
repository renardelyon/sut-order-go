package redis

import (
	"sut-order-go/config"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConfig.Password,
		Password: cfg.RedisConfig.Password,
		DB:       cfg.RedisConfig.DB,
	})

	return client
}
