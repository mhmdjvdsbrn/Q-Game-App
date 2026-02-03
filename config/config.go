package config

import (
	"q-game-app/repository/mysql"
	"q-game-app/service/authservice"
)

type HttpServer struct {
	Port int `koanf:"port"`
}
type Config struct {
	HttpServer HttpServer         `koanf:"http_server"`
	Auth       authservice.Config `koanf:"auth"`
	Mysql      mysql.Config       `koanf:"mysql"`
}
