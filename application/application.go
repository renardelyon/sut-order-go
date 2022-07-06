package application

import (
	"context"
	"sut-order-go/lib/pkg/db"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

type Application struct {
	DbClients   db.Handler
	RedisClient *redis.Client
	GrpcServer  *grpc.Server
	GrpcClients map[string]*grpc.ClientConn
	Context     context.Context
}
