package service

import (
	"sut-order-go/lib/pkg/db"

	"github.com/go-redis/redis/v8"
)

type Service struct {
	H       db.Handler
	redisDB *redis.Client
}

func NewService(H db.Handler, redisDB *redis.Client) *Service {
	return &Service{
		H,
		redisDB,
	}
}
