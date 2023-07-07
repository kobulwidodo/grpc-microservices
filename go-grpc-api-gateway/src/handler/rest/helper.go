package rest

import (
	"context"
	"go-grpc-api-gateway/src/business/entity"
	pb "go-grpc-api-gateway/src/proto/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

func (r *rest) httpRespSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	resp := entity.Response{
		Meta: entity.Meta{
			Message: message,
			Code:    code,
			IsError: false,
		},
		Data: data,
	}
	ctx.JSON(code, resp)
}

func (r *rest) httpRespError(ctx *gin.Context, err error) {
	code := http.StatusInternalServerError
	st, ok := status.FromError(err)
	if ok {
		code = int(st.Code())
	}

	resp := entity.Response{
		Meta: entity.Meta{
			Message: st.Message(),
			Code:    int(st.Code()),
			IsError: true,
		},
		Data: nil,
	}
	ctx.AbortWithStatusJSON(code, resp)
}

func (r *rest) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		r.httpRespError(ctx, status.Error(http.StatusUnauthorized, "unauthorized"))
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		r.httpRespError(ctx, status.Error(http.StatusUnauthorized, "unauthorized"))
		return
	}

	res, err := r.client.Auth.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		r.httpRespError(ctx, status.Error(http.StatusUnauthorized, "unauthorized"))
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}
