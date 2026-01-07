package userhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"q-game-app/param"
)

func (h Handler) userProfile(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	claim, err := h.authSvc.ParseToken(authToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	// call service
	response, err := h.userSvc.Profile(param.ProfileRequest{UserID: claim.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
