package rest

import (
	"context"
	"go-grpc-api-gateway/src/business/entity"
	pb "go-grpc-api-gateway/src/proto/order"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

// @Summary Create Order
// @Description Create New Order
// @Tags Order
// @Security BearerAuth
// @Param order body entity.CreateOrderParam true "order info"
// @Produce json
// @Success 200 {object} entity.Response{}
// @Failure 400 {object} entity.Response{}
// @Failure 401 {object} entity.Response{}
// @Failure 404 {object} entity.Response{}
// @Failure 500 {object} entity.Response{}
// @Router /api/v1/order [POST]
func (r *rest) CreateOrder(ctx *gin.Context) {
	body := entity.CreateOrderParam{}

	if err := ctx.BindJSON(&body); err != nil {
		r.httpRespError(ctx, status.Error(http.StatusBadRequest, err.Error()))
		return
	}

	userId, _ := ctx.Get("userId")

	res, err := r.client.Order.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, int(res.Status), "successfully create new order", &res)
}
