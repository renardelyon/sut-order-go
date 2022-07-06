package grpc

import (
	"context"
	"sut-order-go/domain/product/model"
	productpb "sut-order-go/pb/product"
)

type repo struct {
	productGrpc productpb.ProductServiceClient
}

func NewGrpcRepo(productGrpc productpb.ProductServiceClient) *repo {
	return &repo{
		productGrpc: productGrpc,
	}
}

func (r *repo) SaveRequestedGift(info model.ProductInfo) (*productpb.SaveRequestedGiftResponse, error) {
	req := &productpb.SaveRequestedGiftRequest{
		AdminId:     info.AdminId,
		Fullname:    info.Fullname,
		Username:    info.Username,
		UserId:      info.UserId,
		Productname: info.Productname,
		Url:         info.Url,
	}
	return r.productGrpc.SaveRequestedGift(context.Background(), req)
}
