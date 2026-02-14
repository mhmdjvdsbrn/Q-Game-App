package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"q-game-app/entity"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint
	Role   entity.Role
}

func (c Claims) Validate() error {
	validator := jwt.NewValidator()
	return validator.Validate(&c.RegisteredClaims)
}
