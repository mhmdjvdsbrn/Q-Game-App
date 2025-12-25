package migratior

import (
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"q-game-app/repository/mysql"
)

type Migrator struct {
	dialect    string
	dbConfig   mysql.Config
	migrations *migrate.FileMigrationSource
}

func New(cfg mysql.Config) Migrator {
	migrations := &migrate.FileMigrationSource{
		Dir: "./repository/mysql/migrations",
	}
	return Migrator{dbConfig: cfg, migrations: migrations, dialect: "mysql"}
}
func (m Migrator) Up() {
	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))
	if err != nil {
		panic(fmt.Errorf("Can't connect to mysql database: %v", err))
	}
	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("Can't migrate: %v", err))
	}
	fmt.Printf("Applied %d migrations\n", n)

}
func (m Migrator) Down() {
	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName))
	if err != nil {
		panic(fmt.Errorf("Can't connect to mysql database: %v", err))
	}
	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("Can't migrate: %v", err))
	}
	fmt.Printf("Applied %d migrations\n", n)
}
func (m Migrator) Status() {
	//ToDO->implement status  sql-migrate
}
