package userservice

import (
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
	Name        string
	PhoneNumber string
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
	//create user to db
	createdUser, err := s.repo.RegisterUser(entity.User{Name: req.Name, PhoneNumber: req.PhoneNumber})
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
	}
	//return user
	return RegisterResponse{User: createdUser}, nil
}
