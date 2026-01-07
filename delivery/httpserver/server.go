package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"q-game-app/config"
	"q-game-app/delivery/httpserver/userhandler"
	"q-game-app/service/authservice"
	"q-game-app/service/userservice"
	"q-game-app/validator/uservalidator"
)

type Server struct {
	config      config.Config
	userHandler userhandler.Handler
}

func New(config config.Config, authSvc authservice.Service,
	userSvc userservice.Service, userValidator uservalidator.Validator) Server {
	return Server{
		config:      config,
		userHandler: userhandler.New(authSvc, userSvc, userValidator),
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", s.healthCheck)
	s.userHandler.SetUserRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)))
}
