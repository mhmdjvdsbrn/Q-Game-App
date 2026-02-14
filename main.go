package main

import (
	"log"
	"q-game-app/repository/mysql/mysqlaccesscontrol"
	"q-game-app/repository/mysql/mysqluserrepository"
	"q-game-app/validator/uservalidator/matchingvalidator"

	"q-game-app/adapter/redis"
	cfg "q-game-app/config"
	"q-game-app/delivery/httpserver"
	"q-game-app/repository/migratior"
	"q-game-app/repository/redis/redismatching"
	"q-game-app/service/authorizationservice"
	"q-game-app/service/backofficeuserservice"
	"q-game-app/service/matchingservice"
	"q-game-app/validator/uservalidator"

	"q-game-app/repository/mysql"
	"q-game-app/service/authservice"
	"q-game-app/service/userservice"
)

func main() {
	// Load configuration
	cfg := cfg.Load("config.yml")

	// Run database migrations
	mgr := migratior.New(cfg.Mysql)
	mgr.Up()

	// Setup all services
	authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc, matchingSvc, matchingV := setupServices(cfg)

	// Start HTTP server

	server := httpserver.New(cfg, authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc,
		matchingSvc, matchingV)
	log.Println("Starting server on :8080...")
	server.Serve()
}

func setupServices(cfg cfg.Config) (
	authservice.Service,
	userservice.Service,
	uservalidator.Validator,
	backofficeuserservice.Service,
	authorizationservice.Service,
	matchingservice.Service, matchingvalidator.Validator,
) {
	// Auth and user services
	authSvc := authservice.New(cfg.Auth)
	MySqlRepo := mysql.New(cfg.Mysql)

	userMysql := mysqluserrepository.New(MySqlRepo)
	userSvc := userservice.New(userMysql, authSvc)

	userValidator := uservalidator.New(userMysql)
	backofficeUserSvc := backofficeuserservice.New()

	aclMysql := mysqlaccesscontrol.New(MySqlRepo)

	authorizationSvc := authorizationservice.New(aclMysql)

	// Redis and matching services
	matchingV := matchingvalidator.New()

	redisAdapter := redis.New(cfg.Redis)
	matchingRepo := redismatching.New(redisAdapter)
	matchingSvc := matchingservice.New(cfg.MatchingService, matchingRepo)

	return authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc, matchingSvc, matchingV
}
