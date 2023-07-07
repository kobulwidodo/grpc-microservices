package product

import (
	"context"
	"fmt"
	"go-grpc-order-service/src/pb"

	"google.golang.org/grpc"
)

type ProductInterface interface {
	Get(ctx context.Context, productId int64) (*pb.GetResponse, error)
}

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductClient(url string) ProductInterface {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := &ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}

	return c
}

func (c *ProductServiceClient) Get(ctx context.Context, productId int64) (*pb.GetResponse, error) {
	req := &pb.GetRequest{
		Id: productId,
	}

	return c.Client.Get(ctx, req)
}
