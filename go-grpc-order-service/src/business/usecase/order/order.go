package order

import (
	"context"
	orderDom "go-grpc-order-service/src/business/domain/order"
	"go-grpc-order-service/src/business/entity"
	"go-grpc-order-service/src/client/product"
	"go-grpc-order-service/src/pb"
	"net/http"

	"google.golang.org/grpc/status"
)

type Interface interface {
	CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
}

type order struct {
	order         orderDom.Interface
	productClient product.ProductInterface
}

func Init(od orderDom.Interface, pc product.ProductInterface) Interface {
	o := &order{
		order:         od,
		productClient: pc,
	}

	return o
}

func (o *order) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := o.productClient.Get(ctx, req.ProductId)
	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, status.Error(http.StatusNotFound, "product does not exist")
	}

	order, err := o.order.Create(entity.Order{
		Price:     product.Data.Price * req.Quantity,
		ProductId: product.Data.Id,
		Quantity:  req.Quantity,
		UserId:    req.UserId,
	})
	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, status.Error(http.StatusInternalServerError, "failed to create new product")
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     int64(order.ID),
	}, nil
}
