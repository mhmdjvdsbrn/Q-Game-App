package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

type MysqlDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *MysqlDB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s",
		config.Username, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		panic(fmt.Errorf("can't connect to mysql database: %v", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &MysqlDB{config: config, db: db}
}
