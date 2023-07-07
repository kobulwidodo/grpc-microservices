package usecase

import (
	"go-grpc-auth-service/src/business/domain"
	"go-grpc-auth-service/src/business/usecase/auth"
	authLib "go-grpc-auth-service/src/lib/auth"
)

type Usecase struct {
	Auth auth.Interface
}

func Init(al authLib.Interface, d *domain.Domains) *Usecase {
	uc := &Usecase{
		Auth: auth.Init(d.Auth, al),
	}

	return uc
}
