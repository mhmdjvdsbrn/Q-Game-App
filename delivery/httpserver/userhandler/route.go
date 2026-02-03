package userhandler

import (
	"github.com/labstack/echo/v4"
	"q-game-app/delivery/httpserver/middleware"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.POST("/register-user", h.userRegister)
	userGroup.POST("/login", h.userLogin)
	userGroup.GET("/profile", h.userProfile, middleware.AuthMiddleware(h.authSvc, h.authConfig))
}
