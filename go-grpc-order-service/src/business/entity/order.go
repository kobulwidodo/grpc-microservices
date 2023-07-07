package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Price     int64 `json:"price"`
	Quantity  int64 `json:"quantity"`
	ProductId int64 `json:"product_id"`
	UserId    int64 `json:"user_id"`
}
