package domain

import (
	"go-grpc-auth-service/src/business/domain/auth"

	"gorm.io/gorm"
)

type Domains struct {
	Auth auth.Interface
}

func Init(db *gorm.DB) *Domains {
	d := &Domains{
		Auth: auth.Init(db),
	}

	return d
}
