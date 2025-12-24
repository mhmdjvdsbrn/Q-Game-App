package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"q-game-app/service/userservice"
)

func (s Server) userRegister(c echo.Context) error {
	var uReq userservice.RegisterRequest
	if err := c.Bind(&uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := s.userSvc.Register(uReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
