package main

import (
	"log"
	"q-game-app/config"
	"q-game-app/delivery/httpserver"
	//"q-game-app/repository/migratior"
	"q-game-app/validator/uservalidator"
	"time"

	"q-game-app/repository/mysql"
	"q-game-app/service/authservice"
	"q-game-app/service/userservice"
)

const (
	jwtSecret = "my_super_secret_key_123"

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

func main() {
	cfg := config.Config{
		HttpServer: config.HttpServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:               jwtSecret,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Username: DBUserName,
			Password: DBUserPass,
			Port:     DBPort,
			DBName:   DBName,
			Host:     DBHost,
		},
	}
	//mgr := migratior.New(cfg.Mysql)
	//mgr.Up()
	authSvc, userSvc, userValidator := setupServices(cfg)
	
	server := httpserver.New(cfg, authSvc, userSvc, userValidator)
	server.Serve()

	log.Println("Starting server on :8080...")
}
func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)

	mysqlRepo := mysql.New(cfg.Mysql)

	userSvc := userservice.New(mysqlRepo, authSvc)
	userValidator := uservalidator.New(mysqlRepo)

	return authSvc, userSvc, userValidator
}
