package main

import (
	"fmt"
	"log"
	cfg "q-game-app/config"
	"q-game-app/delivery/httpserver"

	"q-game-app/validator/uservalidator"

	"q-game-app/repository/mysql"
	"q-game-app/service/authservice"
	"q-game-app/service/userservice"
)

//const (
//	jwtSecret = "my_super_secret_key_123"
//
//	AccessTokenExpireDuration  = time.Hour * 24
//	cfg.RefreshTokenExpireDuration = time.Hour * 24 * 7
//
//	AccessTokenSubject  = "access"
//	RefreshTokenSubject = "refresh"
//
//	DBUserName = "myuser"
//	DBUserPass = "mypassword"
//	DBHost     = "localhost"
//	DBName     = "mydb"
//	DBPort     = 3306
//)

func main() {
	cfg2 := cfg.Load()
	fmt.Println(cfg2)
	cfg := cfg.Config{
		HttpServer: cfg.HttpServer{Port: 8080},
		Auth: authservice.Config{
			AccessExpirationTime:  cfg.AccessTokenExpireDuration,
			RefreshExpirationTime: cfg.RefreshTokenExpireDuration,
			AccessSubject:         cfg.AccessTokenSubject,
			RefreshSubject:        cfg.RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Username: cfg.DBUserName,
			Password: cfg.DBUserPass,
			Port:     cfg.DBPort,
			DBName:   cfg.DBName,
			Host:     cfg.DBHost,
		},
	}
	//mgr := migratior.New(cfg.Mysql)
	//mgr.Up()
	authSvc, userSvc, userValidator := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc, userValidator)
	server.Serve()

	log.Println("Starting server on :8080...")
}
func setupServices(cfg cfg.Config) (authservice.Service, userservice.Service, uservalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)

	mysqlRepo := mysql.New(cfg.Mysql)

	userSvc := userservice.New(mysqlRepo, authSvc)
	userValidator := uservalidator.New(mysqlRepo)

	return authSvc, userSvc, userValidator
}
