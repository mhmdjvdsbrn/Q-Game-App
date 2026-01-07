package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"q-game-app/entity"
)

type Repository interface {
	RegisterUser(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUserByID(userID uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}
type Service struct {
	auth AuthGenerator
	repo Repository
}

// New now requires AuthGenerator
func New(repo Repository, auth AuthGenerator) Service {
	return Service{
		auth: auth,
		repo: repo,
	}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
