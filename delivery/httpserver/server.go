package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"q-game-app/config"
	"q-game-app/delivery/httpserver/backofficeuserhandler"
	"q-game-app/delivery/httpserver/matchinghandler"
	"q-game-app/delivery/httpserver/userhandler"
	"q-game-app/service/authorizationservice"
	"q-game-app/service/authservice"
	"q-game-app/service/matchingservice"
	"q-game-app/validator/uservalidator/matchingvalidator"

	"q-game-app/service/backofficeuserservice"
	"q-game-app/service/userservice"
	"q-game-app/validator/uservalidator"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
	matchingHandler       matchinghandler.Handler
}

func New(config config.Config, authSvc authservice.Service,
	userSvc userservice.Service, userValidator uservalidator.Validator,
	backofficeUserSvc backofficeuserservice.Service,
	authorizationSvc authorizationservice.Service,
	matchingSvc matchingservice.Service,
	matchingValidator matchingvalidator.Validator,

) Server {
	return Server{
		config:                config,
		userHandler:           userhandler.New(authSvc, userSvc, userValidator, config.Auth),
		backofficeUserHandler: backofficeuserhandler.New(config.Auth, authSvc, backofficeUserSvc, authorizationSvc),
		matchingHandler:       matchinghandler.New(config.Auth, authSvc, matchingSvc, matchingValidator),
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", s.healthCheck)
	s.userHandler.SetUserRoutes(e)
	s.backofficeUserHandler.SetRoutes(e)
	s.matchingHandler.SetRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)))
}
