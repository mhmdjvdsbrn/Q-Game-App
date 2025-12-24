package config

import (
	"q-game-app/repository/mysql"
	"q-game-app/service/authservice"
)

type HttpServer struct {
	Port int
}
type Config struct {
	HttpServer HttpServer
	Auth       authservice.Config
	Mysql      mysql.Config
}
