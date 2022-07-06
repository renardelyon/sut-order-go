package product

import (
	"sut-order-go/domain/product/model"
	productpb "sut-order-go/pb/product"
)

type ProductRepoInterface interface {
	SaveRequestedGift(model.ProductInfo) (*productpb.SaveRequestedGiftResponse, error)
}
