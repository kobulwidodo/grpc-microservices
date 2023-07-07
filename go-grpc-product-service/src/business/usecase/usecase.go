package usecase

import (
	"go-grpc-product-service/src/business/domain"
	"go-grpc-product-service/src/business/usecase/product"
)

type Usecase struct {
	Product product.Interface
}

func Init(d *domain.Domains) *Usecase {
	uc := &Usecase{
		Product: product.Init(d.Product),
	}

	return uc
}
