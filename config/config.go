package config

import (
	"q-game-app/adapter/redis"
	"q-game-app/repository/mysql"
	"q-game-app/service/authservice"
	"q-game-app/service/matchingservice"
)

type HttpServer struct {
	Port int `koanf:"port"`
}
type Config struct {
	HttpServer      HttpServer             `koanf:"http_server"`
	Auth            authservice.Config     `koanf:"auth"`
	Mysql           mysql.Config           `koanf:"mysql"`
	MatchingService matchingservice.Config `koanf:"matching_service"`
	Redis           redis.Config           `koanf:"redis"`
}
