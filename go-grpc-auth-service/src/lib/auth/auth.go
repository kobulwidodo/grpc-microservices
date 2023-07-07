package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type Interface interface {
	GenerateToken(user User) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type auth struct {
}

func Init() Interface {
	return &auth{}
}

func (a *auth) GenerateToken(user User) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = user.ID
	claim["email"] = user.Email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil

}

func (a *auth) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("token invalid")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
