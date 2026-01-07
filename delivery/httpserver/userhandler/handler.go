package userhandler

import (
	"q-game-app/service/authservice"
	"q-game-app/service/userservice"
	"q-game-app/validator/uservalidator"
)

type Handler struct {
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}

func New(authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
