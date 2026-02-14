package config

import "time"

const (
	AuthMiddlewareContextKey   = "claim"
	AccessTokenSubject         = "access"
	RefreshTokenSubject        = "refresh"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)
