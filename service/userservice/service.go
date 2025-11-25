package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"q-game-app/entity"
	"q-game-app/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(u entity.User) (entity.User, error)
}
type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type RegisterResponse struct {
	User entity.User
}

func (s *Service) Register(req RegisterRequest) (RegisterResponse, error) {
	//validate phone number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, errors.New("invalid phone number")
	}
	//check uniq phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
		}
		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number %s is not unique", req.PhoneNumber)
		}
	}
	//validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name is too short")
	}

	//TODO -> check password with regex
	//validate passsword
	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password is too short, minimum is 8")
	}
	//create user to db
	createdUser, err := s.repo.RegisterUser(entity.User{Name: req.Name, PhoneNumber: req.PhoneNumber, Password: getMD5Hash(req.Password)})
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
	}
	//return user
	return RegisterResponse{User: createdUser}, nil
}

type LoginRequest struct {
	PhoneNumber string
	Password    string
}

type LoginResponse struct {
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	//check phone number existence

	// get user by phone number

	//compare user.password with the req.password

	//return ok
	panic("implement me")
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
