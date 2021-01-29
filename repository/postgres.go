package repository

import (
	"fmt"

	server "github.com/andy-smoker/wh-server"
	"github.com/jmoiron/sqlx"

	// postgres driver
	_ "github.com/lib/pq"
)

const (
	// название таблиц в базе данных
	userTable = "users"
)

// создаём новое подключение к postgresql
func NewPostgresDB(cfg server.PostgresCFG) (*sqlx.DB, error) {
	// connect to postgresdb
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
