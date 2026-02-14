package matchinghandler

import (
	"github.com/labstack/echo/v4"
	"q-game-app/delivery/httpserver/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/matching")

	userGroup.POST("/add-to-waiting-list", h.addToWaitingList,
		middleware.Auth(h.authSvc, h.authConfig))
}
