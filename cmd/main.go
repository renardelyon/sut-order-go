package main

import (
	"log"
	"sut-order-go/application"
	"sut-order-go/config"
	"sut-order-go/domain/order/service"
	productGrpc "sut-order-go/domain/product/repo/grpc"
	storageGrpc "sut-order-go/domain/storage/repo/grpc"
	pb "sut-order-go/pb/order"
	productpb "sut-order-go/pb/product"
	storagepb "sut-order-go/pb/storage"
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

	pClient := productpb.NewProductServiceClient(app.GrpcClients["product-management-service"])
	productRepo := productGrpc.NewGrpcRepo(pClient)

	sClient := storagepb.NewStorageServiceClient(app.GrpcClients["storage-service"])
	storageRepo := storageGrpc.NewGrpcRepo(sClient)

	s := service.NewService(app.DbClients, app.RedisClient, productRepo, storageRepo)

	pb.RegisterOrderServiceServer(app.GrpcServer, s)

	err = app.Run(&c)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
