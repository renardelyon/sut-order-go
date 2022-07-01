package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"sut-order-go/domain/order/model"
	"sut-order-go/lib/utils"
	pb "sut-order-go/pb/order"
)

func (s *Service) GetDetailProducts(ctx context.Context, reqDetail *pb.GetDetailProductsRequest) (*pb.GetDetailProductsResponse, error) {
	log.Println(reqDetail.Url)
	doc, err := utils.Scraping(fmt.Sprint(reqDetail.Url))
	if err != nil {
		return &pb.GetDetailProductsResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	data := doc.Find(".elysium").Find("script").Get(1).FirstChild.Data

	var detailProduct model.DetailProduct
	err = json.Unmarshal([]byte(data), &detailProduct)
	if err != nil {
		return &pb.GetDetailProductsResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.GetDetailProductsResponse{
		Status: http.StatusOK,
		Detailproduct: &pb.DetailProduct{
			Name:        detailProduct.Name,
			Description: detailProduct.Description,
			Image:       detailProduct.Image,
			Url:         detailProduct.Url,
		},
	}, nil
}
