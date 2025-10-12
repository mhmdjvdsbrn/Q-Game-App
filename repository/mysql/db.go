package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MysqlDB struct {
	db *sql.DB
}

func New() *MysqlDB {
	db, err := sql.Open("mysql", "myuser:mypassword@(localhost:3308)/mydb")
	if err != nil {
		panic(fmt.Errorf("Can't connect to mysql database: %v", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &MysqlDB{db: db}
}
