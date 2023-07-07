package entity

import (
	"go-grpc-product-service/src/pb"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type ProductParam struct {
	ID uint
}

func (p *Product) ConvertToGetData() *pb.GetData {
	return &pb.GetData{
		Id:    int64(p.ID),
		Name:  p.Name,
		Price: p.Price,
	}
}
