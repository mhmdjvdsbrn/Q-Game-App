package config

import "time"

const (
	jwtSecret                = "my_super_secret_key_123"
	AuthMiddlewareContextKey = "claim"

	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7

	AccessTokenSubject  = "access"
	RefreshTokenSubject = "refresh"

	DBUserName = "myuser"
	DBUserPass = "mypassword"
	DBHost     = "localhost"
	DBName     = "mydb"
	DBPort     = 3306
)
