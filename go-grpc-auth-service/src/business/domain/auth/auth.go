package auth

import (
	"go-grpc-auth-service/src/business/entity"

	"gorm.io/gorm"
)

type Interface interface {
	Create(user entity.User) (entity.User, error)
	Get(param entity.UserParam) (entity.User, error)
}

type auth struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Interface {
	a := &auth{
		db: db,
	}

	return a
}

func (a *auth) Create(user entity.User) (entity.User, error) {
	if err := a.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (a *auth) Get(param entity.UserParam) (entity.User, error) {
	user := entity.User{}

	if err := a.db.Where(param).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
