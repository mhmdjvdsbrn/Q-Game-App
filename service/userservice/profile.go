package userservice

import (
	"fmt"
	"q-game-app/param"
)

func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return param.ProfileResponse{}, fmt.Errorf("unexpected error %w", err)
	}
	return param.ProfileResponse{Name: user.Name}, nil
}
