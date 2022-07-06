package main

import (
	"log"
	"sut-order-go/application"
	"sut-order-go/config"
	"sut-order-go/domain/order/service"
	productGrpc "sut-order-go/domain/product/repo/grpc"
	pb "sut-order-go/pb/order"
	productpb "sut-order-go/pb/product"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config: ", err.Error())
	}

	app, err := application.Setup(&c)
	if err != nil {
		log.Fatalln("Failed at application setup: ", err.Error())
	}

	p := productpb.NewProductServiceClient(app.GrpcClients["product-management-service"])
	productRepo := productGrpc.NewGrpcRepo(p)

	s := service.NewService(app.DbClients, app.RedisClient, productRepo)

	// grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(app.GrpcServer, s)

	err = app.Run(&c)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
