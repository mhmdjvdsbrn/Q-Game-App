package backofficeuserhandler

import (
	"github.com/labstack/echo/v4"
	"q-game-app/delivery/httpserver/middleware"
	"q-game-app/entity"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/backoffice/users")

	userGroup.GET("/", h.listUsers, middleware.Auth(h.authSvc, h.authConfig),
		middleware.AccessCheck(h.authorizationSvc, entity.UserListPermission))
}
