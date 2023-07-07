package client

import (
	"go-grpc-api-gateway/src/client/auth"
	"go-grpc-api-gateway/src/client/order"
	"go-grpc-api-gateway/src/client/product"
	pbAuth "go-grpc-api-gateway/src/proto/auth"
	pbOrder "go-grpc-api-gateway/src/proto/order"
	pbProduct "go-grpc-api-gateway/src/proto/product"
	"go-grpc-api-gateway/src/utils/config"
)

type Client struct {
	Auth    pbAuth.AuthServiceClient
	Order   pbOrder.OrderServiceClient
	Product pbProduct.ProductServiceClient
}

func Init(conf config.ServiceUrl) *Client {
	c := &Client{
		Auth:    auth.InitServiceClient(conf),
		Order:   order.InitServiceClient(conf),
		Product: product.InitServiceClient(conf),
	}

	return c
}
