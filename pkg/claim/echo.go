package claim

import (
	"github.com/labstack/echo/v4"
	"q-game-app/config"
	"q-game-app/service/authservice"
)

func GetClaimsFromEchoContext(c echo.Context) *authservice.Claims {
	return c.Get(config.AuthMiddlewareContextKey).(*authservice.Claims)
}
