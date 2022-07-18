package service

import (
	"context"
	"net/http"
	omodel "sut-order-go/domain/order/model"
	pmodel "sut-order-go/domain/product/model"
	pb "sut-order-go/pb/order"
)

func (s *Service) CreateOrder(ctx context.Context, reqCreated *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	var temp omodel.Order
	if result := s.H.DB.Where(&omodel.Order{Id: reqCreated.UserId, ProductName: reqCreated.ProductName}).First(&temp); result.Error == nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  "Product have been ordered previously",
		}, nil
	}

	s.H.DB.Create(&omodel.Order{
		Id:          reqCreated.UserId,
		ProductName: reqCreated.ProductName,
	})

	if res, _ := s.productInterface.SaveRequestedGift(pmodel.ProductInfo{
		AdminId:     reqCreated.AdminId,
		Fullname:    reqCreated.Fullname,
		Username:    reqCreated.Username,
		UserId:      reqCreated.UserId,
		Productname: reqCreated.ProductName,
		Url:         reqCreated.Url,
	}); res.Error != "" {
		return &pb.CreateOrderResponse{
			Status: res.Status,
			Error:  res.Error,
		}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusOK,
	}, nil
}
