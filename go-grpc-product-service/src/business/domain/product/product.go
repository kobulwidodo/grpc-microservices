package product

import (
	"go-grpc-product-service/src/business/entity"

	"gorm.io/gorm"
)

type Interface interface {
	Create(product entity.Product) (entity.Product, error)
	Get(param entity.ProductParam) (entity.Product, error)
}

type product struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Interface {
	p := &product{
		db: db,
	}

	return p
}

func (p *product) Create(product entity.Product) (entity.Product, error) {
	if err := p.db.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *product) Get(param entity.ProductParam) (entity.Product, error) {
	product := entity.Product{}

	if err := p.db.Where(param).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}
