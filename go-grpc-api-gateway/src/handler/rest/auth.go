package rest

import (
	"context"
	"go-grpc-api-gateway/src/business/entity"
	pb "go-grpc-api-gateway/src/proto/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

// @Summary Register User
// @Description Register New User
// @Tags Auth
// @Param user body entity.CreateUserParam true "user info"
// @Produce json
// @Success 200 {object} entity.Response{}
// @Failure 400 {object} entity.Response{}
// @Failure 401 {object} entity.Response{}
// @Failure 404 {object} entity.Response{}
// @Failure 500 {object} entity.Response{}
// @Router /api/v1/auth/register [POST]
func (r *rest) RegisterUser(ctx *gin.Context) {
	body := entity.CreateUserParam{}

	if err := ctx.BindJSON(&body); err != nil {
		r.httpRespError(ctx, status.Error(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := r.client.Auth.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, int(res.Status), "successfully register new user", nil)
}

// @Summary Login User
// @Description Login User
// @Tags Auth
// @Param user body entity.LoginUserParam true "user info"
// @Produce json
// @Success 200 {object} entity.Response{}
// @Failure 400 {object} entity.Response{}
// @Failure 401 {object} entity.Response{}
// @Failure 404 {object} entity.Response{}
// @Failure 500 {object} entity.Response{}
// @Router /api/v1/auth/login [POST]
func (r *rest) LoginUser(ctx *gin.Context) {
	b := entity.LoginUserParam{}

	if err := ctx.BindJSON(&b); err != nil {
		r.httpRespError(ctx, status.Error(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := r.client.Auth.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, int(res.Status), "successfully login", &res)
}
