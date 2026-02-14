package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"q-game-app/entity"
	"q-game-app/pkg/claim"
	"q-game-app/pkg/errmsg"
	"q-game-app/service/authorizationservice"
)

func AccessCheck(service authorizationservice.Service,
	permissions ...entity.PermissionsTitle) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			claims := claim.GetClaimsFromEchoContext(c)
			isAllowed, err := service.CheckAccess(claims.UserID, claims.Role, permissions...)
			if err != nil {
				// TODO - log unexpected error
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": errmsg.ErrorMsgSomethingWentWrong,
				})
			}

			if !isAllowed {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": errmsg.ErrorMsgUserNotAllowed,
				})
			}

			return next(c)
		}
	}
}
