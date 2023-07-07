package domain

import (
	"go-grpc-product-service/src/business/domain/product"

	"gorm.io/gorm"
)

type Domains struct {
	Product product.Interface
}

func Init(db *gorm.DB) *Domains {
	d := &Domains{
		Product: product.Init(db),
	}

	return d
}
