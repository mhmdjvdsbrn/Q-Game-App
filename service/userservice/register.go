package userservice

import (
	"fmt"
	"q-game-app/entity"
	"q-game-app/param"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {
	const op = "service.Register"

	//create user to db
	createdUser, err := s.repo.RegisterUser(entity.User{Name: req.Name, PhoneNumber: req.PhoneNumber, Password: getMD5Hash(req.Password)})
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
	}
	//return
	//user
	return param.RegisterResponse{
		User: struct {
			ID          uint   `json:"id"`
			PhoneNumber string `json:"phone_number"`
			Name        string `json:"name"`
		}{
			ID:          createdUser.ID,
			PhoneNumber: createdUser.PhoneNumber,
			Name:        createdUser.Name,
		},
	}, nil
}
