package service

import (
	"sut-order-go/domain/product"
	"sut-order-go/lib/pkg/db"

	"github.com/go-redis/redis/v8"
)

type Service struct {
	H                db.Handler
	redisDB          *redis.Client
	productInterface product.ProductRepoInterface
}

func NewService(H db.Handler, redisDB *redis.Client, productInterface product.ProductRepoInterface) *Service {
	return &Service{
		H,
		redisDB,
		productInterface,
	}
}
