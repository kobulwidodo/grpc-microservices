package rest

import (
	"context"
	"go-grpc-api-gateway/src/business/entity"
	pb "go-grpc-api-gateway/src/proto/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

// @Summary Create Product
// @Description Create New Product
// @Tags Product
// @Security BearerAuth
// @Param product body entity.CreateProductParam true "product info"
// @Produce json
// @Success 200 {object} entity.Response{}
// @Failure 400 {object} entity.Response{}
// @Failure 401 {object} entity.Response{}
// @Failure 404 {object} entity.Response{}
// @Failure 500 {object} entity.Response{}
// @Router /api/v1/product [POST]
func (r *rest) CreateProduct(ctx *gin.Context) {
	body := entity.CreateProductParam{}

	if err := ctx.BindJSON(&body); err != nil {
		r.httpRespError(ctx, status.Error(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := r.client.Product.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Price: body.Price,
	})

	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, int(res.Status), "sucessfully created new product", &res)
}

// @Summary Find One Product
// @Description Find One Product
// @Tags Product
// @Security BearerAuth
// @Param id path integer true "product id"
// @Produce json
// @Success 200 {object} entity.Response{}
// @Failure 400 {object} entity.Response{}
// @Failure 401 {object} entity.Response{}
// @Failure 404 {object} entity.Response{}
// @Failure 500 {object} entity.Response{}
// @Router /api/v1/product/{id} [GET]
func (r *rest) FindOneProduct(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := r.client.Product.Get(context.Background(), &pb.GetRequest{
		Id: int64(id),
	})

	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, int(res.Status), "successfully fine one product", &res.Data)
}
