package usecase

import (
	"go-grpc-order-service/src/business/domain"
	"go-grpc-order-service/src/business/usecase/order"
	"go-grpc-order-service/src/client"
)

type Usecase struct {
	Order  order.Interface
	Client client.Client
}

func Init(d *domain.Domains, c *client.Client) *Usecase {
	u := &Usecase{
		Order: order.Init(d.Order, c.ProductClient),
	}

	return u
}
