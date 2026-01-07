package userhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"q-game-app/param"
	"q-game-app/pkg/httpmsg"
	"q-game-app/pkg/richerror"
)

func (h Handler) userRegister(c echo.Context) error {
	var req param.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.userValidator.ValidateRegisterRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		fieldErrors := map[string][]string{}
		if rErr, ok := err.(richerror.RichError); ok {
			if meta, ok := rErr.GetMeta()["errors"].(map[string][]string); ok {
				fieldErrors = meta
			}
		}

		return c.JSON(code, echo.Map{
			"code":    code,
			"message": msg,
			"error":   fieldErrors,
		})
	}

	response, err := h.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
