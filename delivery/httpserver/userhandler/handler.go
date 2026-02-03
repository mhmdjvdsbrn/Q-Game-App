package userhandler

import (
	"q-game-app/service/authservice"
	"q-game-app/service/userservice"
	"q-game-app/validator/uservalidator"
)

type Handler struct {
	authConfig    authservice.Config
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator, config authservice.Config) Handler {
	return Handler{
		authConfig:    config,
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
