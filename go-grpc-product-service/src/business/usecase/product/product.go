package product

import (
	"context"
	productDom "go-grpc-product-service/src/business/domain/product"
	"go-grpc-product-service/src/business/entity"
	"go-grpc-product-service/src/pb"
	"net/http"

	"google.golang.org/grpc/status"
)

type Interface interface {
	CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
	Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error)
}

type product struct {
	product productDom.Interface
}

func Init(pd productDom.Interface) Interface {
	p := &product{
		product: pd,
	}

	return p
}

func (p *product) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := p.product.Create(entity.Product{
		Name:  req.Name,
		Price: req.Price,
	})
	if err != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, status.Error(http.StatusInternalServerError, "failed to create new product")
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     int64(product.ID),
	}, nil
}

func (p *product) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	product, err := p.product.Get(entity.ProductParam{
		ID: uint(req.Id),
	})
	if err != nil {
		return &pb.GetResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, status.Error(http.StatusNotFound, "product not found")
	}

	return &pb.GetResponse{
		Status: http.StatusOK,
		Data:   product.ConvertToGetData(),
	}, nil
}
