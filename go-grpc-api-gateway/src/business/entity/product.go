package entity

type CreateProductParam struct {
	Name  string `json:"name" binding:"required"`
	Price int64  `json:"price" binding:"required"`
}
