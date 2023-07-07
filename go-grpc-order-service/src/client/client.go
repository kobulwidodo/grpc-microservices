package client

import (
	"go-grpc-order-service/src/client/product"
	"go-grpc-order-service/src/utils/config"
)

type Client struct {
	ProductClient product.ProductInterface
}

func Init(cfg config.ServiceUrl) *Client {
	c := &Client{
		ProductClient: product.InitProductClient(cfg.ProductUrl),
	}

	return c
}
