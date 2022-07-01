package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sut-order-go/domain/order/model"
	"time"

	"github.com/PuerkitoBio/goquery"

	"sut-order-go/lib/utils"
	pb "sut-order-go/pb/order"
)

func (s *Service) GetProductsToOrderByKeyword(ctx context.Context, reqProd *pb.GetProductsToOrderRequest) (*pb.GetProductsToOrderResponse, error) {
	redisData := s.redisDB.Get(ctx, reqProd.Keyword)
	res, _ := redisData.Result()

	if res != "" {
		return &pb.GetProductsToOrderResponse{
			Status: http.StatusOK,
			Data:   res,
		}, nil
	}

	doc, err := utils.Scraping(fmt.Sprint("https://www.bukalapak.com/products?search%5Bkeywords%5D=", reqProd.Keyword))
	if err != nil {
		return &pb.GetProductsToOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	rows := make([]model.Product, 0)
	doc.Find(".bl-product-card").Each(func(i int, sel *goquery.Selection) {
		row := new(model.Product)

		name := sel.Find(".bl-product-card__description-name").Text()

		regex, _ := regexp.Compile(`\n`)

		trimName := regex.ReplaceAllString(name, " ")
		trimName = strings.Trim(trimName, " ")
		row.Name = trimName

		url, _ := sel.Find("a").Attr("href")
		row.Url = url

		rows = append(rows, *row)
	})

	bts, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		return &pb.GetProductsToOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	ttl := time.Duration(5) * time.Minute

	op := s.redisDB.Set(ctx, reqProd.Keyword, string(bts), ttl)
	if err := op.Err(); err != nil {
		return &pb.GetProductsToOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.GetProductsToOrderResponse{
		Status: http.StatusOK,
		Data:   string(bts),
	}, nil
}
