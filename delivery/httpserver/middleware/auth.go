package middleware

import (
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	cfg "q-game-app/config"
	"q-game-app/service/authservice"
)

func AuthMiddleware(service authservice.Service, config authservice.Config) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey:    cfg.AuthMiddlewareContextKey,
		SigningKey:    []byte(config.SignKey),
		SigningMethod: "HS256",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claim, err := service.ParseToken(auth)
			if err != nil {
				return nil, err
			}
			return claim, nil
		},
	})
}

func GetClaim(c echo.Context) *authservice.Claims {
	claim := c.Get(cfg.AuthMiddlewareContextKey)
	return claim.(*authservice.Claims)
}
