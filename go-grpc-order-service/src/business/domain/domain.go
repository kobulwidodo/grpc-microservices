package domain

import (
	"go-grpc-order-service/src/business/domain/order"

	"gorm.io/gorm"
)

type Domains struct {
	Order order.Interface
}

func Init(db *gorm.DB) *Domains {
	d := &Domains{
		Order: order.Init(db),
	}

	return d
}
