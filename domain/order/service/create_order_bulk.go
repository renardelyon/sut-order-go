package service

import (
	"bytes"
	"context"
	"log"
	"net/http"
	productModel "sut-order-go/domain/product/model"
	productpb "sut-order-go/pb/product"

	pb "sut-order-go/pb/order"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func (s *Service) CreateOrderBulk(ctx context.Context, req *pb.CreateOrderBulkRequest) (*pb.CreateOrderBulkResponse, error) {
	fileBytes, err := s.storageInterface.GetFileByUserId(req.UserId)
	if err != nil {
		log.Println(err)
		return &pb.CreateOrderBulkResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	reader := bytes.NewReader(fileBytes)
	xlsx, err := excelize.OpenReader(reader)
	if err != nil {
		log.Println(err)
		return &pb.CreateOrderBulkResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	sheetName := "Sheet1"

	rows, err := xlsx.Rows(sheetName)
	if err != nil {
		log.Println(err)
		return &pb.CreateOrderBulkResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	productInfos := make([]*productpb.ProductInfo, 0)

	index := 0
	for rows.Next() {
		row := rows.Columns()

		if len(row) <= 0 {
			break
		}

		if index > 0 {
			productInfos = append(productInfos, &productpb.ProductInfo{
				Productname: row[0],
				Url:         row[1],
			})
		}

		index++
	}

	if rows.Error() != nil {
		log.Println(err)
		return &pb.CreateOrderBulkResponse{
			Status: http.StatusInternalServerError,
			Error:  rows.Error().Error(),
		}, nil
	}

	info := productModel.ProductInfo{
		AdminId:  req.AdminId,
		Fullname: req.Fullname,
		Username: req.Username,
		UserId:   req.UserId,
	}

	res, _ := s.productInterface.SaveRequestedGiftBulk(info, productInfos)
	if res.Error != "" {
		log.Println(err)
		return &pb.CreateOrderBulkResponse{
			Status: http.StatusInternalServerError,
			Error:  rows.Error().Error(),
		}, nil
	}

	return &pb.CreateOrderBulkResponse{
		Status: http.StatusOK,
	}, nil
}
