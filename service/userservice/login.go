package userservice

import (
	"q-game-app/param"
	"q-game-app/pkg/errmsg"
	"q-game-app/pkg/richerror"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	const op = "service.Login"
	//check phone number existence
	// get user by phone number
	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return param.LoginResponse{}, richerror.New(
			op,
		).WithMessage(err.Error())
	}

	//compare user.password with the req.password
	if user.Password != getMD5Hash(req.Password) {
		return param.LoginResponse{}, richerror.New(op).WithMessage(errmsg.ErrorMsgUserNameOrPasswordIsWrong)

	}
	//return ok
	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return param.LoginResponse{}, richerror.New(
			op,
		)
	}
	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return param.LoginResponse{}, richerror.New(
			op,
		)
	}
	return param.LoginResponse{
		User:   param.UserInfo{ID: user.ID, PhoneNumber: user.PhoneNumber, Name: user.Name},
		Tokens: param.Tokens{AccessToken: accessToken, RefreshToken: refreshToken},
	}, nil
}
