package auth

import (
	"context"
	authDom "go-grpc-auth-service/src/business/domain/auth"
	"go-grpc-auth-service/src/business/entity"
	authLib "go-grpc-auth-service/src/lib/auth"
	"go-grpc-auth-service/src/pb"
	"net/http"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

type Interface interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error)
}

type auth struct {
	auth    authDom.Interface
	authLib authLib.Interface
}

func Init(ad authDom.Interface, al authLib.Interface) Interface {
	a := &auth{
		auth:    ad,
		authLib: al,
	}

	return a
}

func (a *auth) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, err := a.auth.Get(entity.UserParam{
		Email: req.Email,
	})
	if err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "email already registered",
		}, status.Error(http.StatusConflict, "email already registered")
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, status.Error(http.StatusInternalServerError, "failed to hash password")
	}

	_, err = a.auth.Create(entity.User{
		Email:    req.Email,
		Password: string(hashPass),
	})
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, status.Error(http.StatusInternalServerError, "failed to register new account")
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (a *auth) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := a.auth.Get(entity.UserParam{
		Email: req.Email,
	})
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, status.Error(http.StatusNotFound, "data not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, status.Error(http.StatusNotFound, "data not found")
	}

	token, err := a.authLib.GenerateToken(user.ConvertToAuthUser())
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, status.Error(http.StatusInternalServerError, "failed to generate token")
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (a *auth) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	token, err := a.authLib.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return &pb.ValidateResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	user, err := a.auth.Get(entity.UserParam{
		Email: claim["email"].(string),
	})
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, err
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: int64(user.ID),
	}, nil
}
