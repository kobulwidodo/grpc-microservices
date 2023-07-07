package entity

import (
	"go-grpc-auth-service/src/lib/auth"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserParam struct {
	Email string
}

func (u *User) ConvertToAuthUser() auth.User {
	return auth.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}
