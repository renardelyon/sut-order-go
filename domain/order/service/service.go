package service

import (
	"sut-order-go/domain/product"
	"sut-order-go/domain/storage"
	"sut-order-go/lib/pkg/db"

	"github.com/go-redis/redis/v8"
)

type Service struct {
	H                db.Handler
	redisDB          *redis.Client
	productInterface product.ProductRepoInterface
	storageInterface storage.StorageRepoInterface
}

func NewService(
	H db.Handler,
	redisDB *redis.Client,
	productInterface product.ProductRepoInterface,
	storageInterface storage.StorageRepoInterface,
) *Service {
	return &Service{
		H,
		redisDB,
		productInterface,
		storageInterface,
	}
}
