package order

import (
	"go-grpc-order-service/src/business/entity"

	"gorm.io/gorm"
)

type Interface interface {
	Create(order entity.Order) (entity.Order, error)
}

type order struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Interface {
	o := &order{
		db: db,
	}

	return o
}

func (o *order) Create(order entity.Order) (entity.Order, error) {
	if err := o.db.Create(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}
