package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"q-game-app/config"
	"q-game-app/service/authservice"
	"q-game-app/service/userservice"
)

type Server struct {
	config  config.Config
	authSvc *authservice.Service
	userSvc *userservice.Service
}

func New(config config.Config,
	authSvc *authservice.Service,
	userSvc *userservice.Service) *Server {
	return &Server{
		config:  config,
		authSvc: authSvc,
		userSvc: userSvc,
	}
}

func (s *Server) Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", s.healthCheck)

	e.POST("/users/register-user", s.userRegister)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)))
}
